package cli

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// primaryXMLPackage is the standard RPM primary metadata representation.
type primaryXMLPackage struct {
	Type    string `xml:"type,attr"`
	Name    string `xml:"name"`
	Arch    string `xml:"arch"`
	Version struct {
		Epoch string `xml:"epoch,attr"`
		Ver   string `xml:"ver,attr"`
		Rel   string `xml:"rel,attr"`
	} `xml:"version"`
	Checksum struct {
		Type  string `xml:"type,attr"`
		Value string `xml:",chardata"`
	} `xml:"checksum"`
	Summary     string `xml:"summary"`
	Description string `xml:"description"`
	Packager    string `xml:"packager"`
	URL         string `xml:"url"`
	Time        struct {
		File  *int64 `xml:"file,attr"`
		Build *int64 `xml:"build,attr"`
	} `xml:"time"`
	Size struct {
		Package   *int64 `xml:"package,attr"`
		Installed *int64 `xml:"installed,attr"`
		Archive   *int64 `xml:"archive,attr"`
	} `xml:"size"`
	Location struct {
		Href string `xml:"href,attr"`
		Base string `xml:"base,attr"`
	} `xml:"location"`
	Format struct {
		License   string `xml:"license"`
		Vendor    string `xml:"vendor"`
		Group     string `xml:"group"`
		BuildHost string `xml:"buildhost"`
		SourceRPM string `xml:"sourcerpm"`
		Header    struct {
			Start *int64 `xml:"start,attr"`
			End   *int64 `xml:"end,attr"`
		} `xml:"header-range"`
	} `xml:"format"`
}

func (p *DNFParser) extractPrimaryXML(data []byte) ([]DNFPackage, error) {
	var packages []DNFPackage
	_, err := readPrimaryXML(p.ctx, data, func(pkg DNFPackage) {
		packages = append(packages, pkg)
	})
	return packages, err
}

// readPrimaryXML processes one package at a time. A nil visitor validates a
// fetched payload without retaining another copy of every package in memory.
func readPrimaryXML(ctx context.Context, data []byte, visit func(DNFPackage)) (int, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	var root xml.StartElement
	for {
		if err := ctx.Err(); err != nil {
			return 0, err
		}
		token, err := decoder.Token()
		if err != nil {
			return 0, fmt.Errorf("parse primary XML: %w", err)
		}
		if start, ok := token.(xml.StartElement); ok {
			root = start
			break
		}
		if text, ok := token.(xml.CharData); ok && strings.TrimSpace(string(text)) != "" {
			return 0, fmt.Errorf("primary XML contains text before its root element")
		}
	}
	if root.Name.Local != "metadata" || (root.Name.Space != "" && root.Name.Space != "http://linux.duke.edu/metadata/common") {
		return 0, fmt.Errorf("expected primary XML metadata root, got %s", root.Name.Local)
	}
	expected := -1
	for _, attr := range root.Attr {
		if attr.Name.Local == "packages" && attr.Name.Space == "" {
			n, err := strconv.Atoi(attr.Value)
			if err != nil || n < 0 || expected >= 0 {
				return 0, fmt.Errorf("primary XML has an invalid package count")
			}
			expected = n
		}
	}
	if expected < 0 {
		return 0, fmt.Errorf("primary XML package count missing")
	}
	str := func(s string) sql.NullString { return sql.NullString{String: s, Valid: s != ""} }
	num := func(n *int64) sql.NullInt64 {
		if n == nil {
			return sql.NullInt64{}
		}
		return sql.NullInt64{Int64: *n, Valid: true}
	}
	count := 0
	for {
		if err := ctx.Err(); err != nil {
			return 0, err
		}
		token, err := decoder.Token()
		if err != nil {
			return 0, fmt.Errorf("parse primary XML: %w", err)
		}
		if end, ok := token.(xml.EndElement); ok && end.Name == root.Name {
			break
		}
		start, ok := token.(xml.StartElement)
		if !ok {
			if text, ok := token.(xml.CharData); ok && strings.TrimSpace(string(text)) != "" {
				return 0, fmt.Errorf("primary XML contains text outside a package")
			}
			continue
		}
		if start.Name.Local != "package" || start.Name.Space != root.Name.Space {
			if err := decoder.Skip(); err != nil {
				return 0, fmt.Errorf("parse primary XML: %w", err)
			}
			continue
		}
		var x primaryXMLPackage
		if err := decoder.DecodeElement(&x, &start); err != nil {
			return 0, fmt.Errorf("parse primary XML package %d: %w", count+1, err)
		}
		x.Checksum.Value = strings.TrimSpace(x.Checksum.Value)
		if x.Type != "rpm" || x.Name == "" || x.Arch == "" || x.Version.Ver == "" || x.Checksum.Type == "" || x.Checksum.Value == "" || x.Location.Href == "" {
			return 0, fmt.Errorf("primary XML package %d lacks required RPM identity fields", count+1)
		}
		count++
		if count > expected {
			return 0, fmt.Errorf("primary XML contains more than %d declared packages", expected)
		}
		if visit == nil {
			continue
		}
		visit(DNFPackage{
			PkgKey: count, PkgId: x.Checksum.Value, Name: x.Name,
			Arch: str(x.Arch), Version: str(x.Version.Ver), Epoch: str(x.Version.Epoch), Release: str(x.Version.Rel),
			Summary: str(x.Summary), Description: str(x.Description), URL: str(x.URL), RPMPackager: str(x.Packager),
			TimeFile: num(x.Time.File), TimeBuild: num(x.Time.Build),
			RPMLicense: str(x.Format.License), RPMVendor: str(x.Format.Vendor), RPMGroup: str(x.Format.Group),
			RPMBuildHost: str(x.Format.BuildHost), RPMSourceRPM: str(x.Format.SourceRPM),
			RPMHeaderStart: num(x.Format.Header.Start), RPMHeaderEnd: num(x.Format.Header.End),
			SizePackage: num(x.Size.Package), SizeInstalled: num(x.Size.Installed), SizeArchive: num(x.Size.Archive),
			LocationHref: str(x.Location.Href), LocationBase: str(x.Location.Base), ChecksumType: str(x.Checksum.Type),
		})
	}
	if count != expected {
		return 0, fmt.Errorf("primary XML declares %d packages but contains %d", expected, count)
	}
	// Require EOF after the single root; otherwise a valid prefix could conceal
	// a second document or corrupt trailing data.
	for {
		if err := ctx.Err(); err != nil {
			return 0, err
		}
		token, err := decoder.Token()
		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return 0, fmt.Errorf("parse primary XML trailer: %w", err)
		}
		switch x := token.(type) {
		case xml.Comment, xml.ProcInst:
		case xml.CharData:
			if strings.TrimSpace(string(x)) != "" {
				return 0, fmt.Errorf("primary XML contains trailing text")
			}
		default:
			return 0, fmt.Errorf("primary XML contains content after its root element")
		}
	}
}
