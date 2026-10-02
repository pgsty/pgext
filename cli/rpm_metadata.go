package cli

import (
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"hash"
	"io"
	"net/url"
	"path"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/zstd"
	"github.com/ulikunitz/xz"
)

const rpmPrimaryFormat = "rpm-primary-xml"

func parseRPMManifest(reader io.Reader) (*RepoMD, error) {
	data, err := readMetadata(reader, 8*1024*1024)
	if err != nil {
		return nil, fmt.Errorf("read repomd.xml: %w", err)
	}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	var manifest RepoMD
	if err := decoder.Decode(&manifest); err != nil {
		return nil, fmt.Errorf("parse repomd.xml: %w", err)
	}
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return &manifest, nil
		}
		if err != nil {
			return nil, fmt.Errorf("parse repomd.xml trailer: %w", err)
		}
		switch x := token.(type) {
		case xml.Comment, xml.ProcInst:
		case xml.CharData:
			if strings.TrimSpace(string(x)) != "" {
				return nil, fmt.Errorf("repomd.xml contains trailing text")
			}
		default:
			return nil, fmt.Errorf("repomd.xml contains content after its root element")
		}
	}
}

// RepoMD is the manifest that identifies a repository's published metadata.
type RepoMD struct {
	XMLName   xml.Name     `xml:"repomd"`
	Data      []RepoMDData `xml:"data"`
	SourceURL string       `xml:"-"`
}

type MetadataChecksum struct {
	Type string `xml:"type,attr" json:"type"`
	Text string `xml:",chardata" json:"value"`
}

type RepoMDData struct {
	Type     string `xml:"type,attr"`
	Location struct {
		Href string `xml:"href,attr"`
	} `xml:"location"`
	Checksum     MetadataChecksum `xml:"checksum"`
	OpenChecksum MetadataChecksum `xml:"open-checksum"`
	Size         *int64           `xml:"size"`
	OpenSize     *int64           `xml:"open-size"`
}

// RPMMetadataCache binds the cached XML to a particular manifest entry and URL.
// Old cache rows without this descriptor are refreshed without HTTP validators.
type RPMMetadataCache struct {
	Format        string           `json:"format"`
	SourceURL     string           `json:"source_url"`
	Compression   string           `json:"compression"`
	Checksum      MetadataChecksum `json:"checksum"`
	OpenChecksum  MetadataChecksum `json:"open_checksum"`
	Size          *int64           `json:"compressed_size,omitempty"`
	OpenSize      *int64           `json:"open_size,omitempty"`
	ContentSHA256 string           `json:"content_sha256"`
	PackageCount  int              `json:"package_count"`
}

func (r *RepoMD) FindPrimary() (*RepoMDData, error) {
	var primary *RepoMDData
	for i := range r.Data {
		if r.Data[i].Type != "primary" {
			continue
		}
		if primary != nil {
			return nil, fmt.Errorf("multiple primary XML entries in repomd.xml")
		}
		primary = &r.Data[i]
	}
	if primary == nil {
		return nil, fmt.Errorf("primary XML entry not found in repomd.xml; this repository must publish primary XML metadata")
	}
	if strings.TrimSpace(primary.Location.Href) == "" {
		return nil, fmt.Errorf("primary XML location missing in repomd.xml")
	}
	return primary, nil
}

func rpmPrimaryURL(metadataURL, href string) (string, error) {
	base, err := url.Parse(metadataURL)
	if err != nil {
		return "", fmt.Errorf("parse repository URL: %w", err)
	}
	base.Path = path.Dir(path.Dir(base.Path))
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawPath, base.RawQuery, base.Fragment = "", "", ""
	ref, err := url.Parse(href)
	if err != nil {
		return "", fmt.Errorf("parse primary XML location: %w", err)
	}
	resolved := base.ResolveReference(ref)
	if (resolved.Scheme != "http" && resolved.Scheme != "https") || resolved.Host == "" {
		return "", fmt.Errorf("primary XML location must resolve to an HTTP URL")
	}
	return resolved.String(), nil
}

func rpmPrimaryPath(repomdPath, href string) (string, error) {
	ref, err := url.Parse(href)
	if err != nil {
		return "", fmt.Errorf("parse primary XML location: %w", err)
	}
	if ref.IsAbs() || ref.Host != "" || ref.RawQuery != "" || ref.Fragment != "" || filepath.IsAbs(ref.Path) {
		return "", fmt.Errorf("local primary XML location must be relative to the repository root")
	}
	root := filepath.Dir(filepath.Dir(repomdPath))
	file := filepath.Join(root, filepath.FromSlash(ref.Path))
	rel, err := filepath.Rel(root, file)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("primary XML location escapes the repository root")
	}
	return file, nil
}

func readMetadata(reader io.Reader, limit int64) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(reader, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(data)) > limit {
		return nil, fmt.Errorf("repository metadata exceeds %d bytes", limit)
	}
	return data, nil
}

func verifyMetadataChecksum(data []byte, checksum MetadataChecksum) error {
	if strings.TrimSpace(checksum.Text) == "" {
		if checksum.Type != "" {
			return fmt.Errorf("%s metadata checksum is empty", checksum.Type)
		}
		return nil
	}
	var h hash.Hash
	switch strings.ToLower(checksum.Type) {
	case "md5":
		h = md5.New()
	case "sha", "sha1":
		h = sha1.New()
	case "sha224":
		h = sha256.New224()
	case "sha256":
		h = sha256.New()
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		return fmt.Errorf("unsupported metadata checksum algorithm: %q", checksum.Type)
	}
	_, _ = h.Write(data)
	if !strings.EqualFold(hex.EncodeToString(h.Sum(nil)), strings.TrimSpace(checksum.Text)) {
		return fmt.Errorf("%s checksum verification failed", checksum.Type)
	}
	return nil
}

func decompressRPMMetadata(data []byte) ([]byte, string, error) {
	var reader io.Reader = bytes.NewReader(data)
	compression := "none"
	// Detect the actual compression independently of the metadata's XML format.
	switch {
	case bytes.HasPrefix(data, []byte{0x1f, 0x8b}):
		gz, err := gzip.NewReader(reader)
		if err != nil {
			return nil, "gzip", err
		}
		defer gz.Close()
		reader, compression = gz, "gzip"
	case bytes.HasPrefix(data, []byte("BZh")):
		reader, compression = bzip2.NewReader(reader), "bzip2"
	case bytes.HasPrefix(data, []byte{0xfd, 0x37, 0x7a, 0x58, 0x5a, 0x00}):
		xr, err := (xz.ReaderConfig{DictCap: 64 << 20}).NewReader(reader)
		if err != nil {
			return nil, "xz", err
		}
		reader, compression = xr, "xz"
	case bytes.HasPrefix(data, []byte{0x28, 0xb5, 0x2f, 0xfd}) ||
		(len(data) >= 4 && data[0]&0xf0 == 0x50 && bytes.Equal(data[1:4], []byte{0x2a, 0x4d, 0x18})):
		zr, err := zstd.NewReader(reader, zstd.WithDecoderConcurrency(1), zstd.WithDecoderMaxMemory(MaxFileSize))
		if err != nil {
			return nil, "zstd", err
		}
		defer zr.Close()
		reader, compression = zr, "zstd"
	}
	result, err := readMetadata(reader, MaxFileSize)
	return result, compression, err
}

// decodeRPMPrimary verifies both representations and validates the entire XML
// before either fetch or scan can replace a cached repository payload.
func decodeRPMPrimary(ctx context.Context, compressed []byte, primary *RepoMDData, sourceURL string) ([]byte, *RPMMetadataCache, error) {
	if primary.Size != nil && *primary.Size != int64(len(compressed)) {
		return nil, nil, fmt.Errorf("primary XML compressed size mismatch")
	}
	if err := verifyMetadataChecksum(compressed, primary.Checksum); err != nil {
		return nil, nil, fmt.Errorf("primary XML compressed checksum: %w", err)
	}
	data, compression, err := decompressRPMMetadata(compressed)
	if err != nil {
		return nil, nil, fmt.Errorf("decompress primary XML: %w", err)
	}
	if primary.OpenSize != nil && *primary.OpenSize != int64(len(data)) {
		return nil, nil, fmt.Errorf("primary XML open size mismatch")
	}
	if err := verifyMetadataChecksum(data, primary.OpenChecksum); err != nil {
		return nil, nil, fmt.Errorf("primary XML open checksum: %w", err)
	}
	count, err := readPrimaryXML(ctx, data, nil)
	if err != nil {
		return nil, nil, err
	}
	return data, &RPMMetadataCache{
		Format: rpmPrimaryFormat, SourceURL: sourceURL, Compression: compression,
		Checksum: primary.Checksum, OpenChecksum: primary.OpenChecksum,
		Size: primary.Size, OpenSize: primary.OpenSize,
		ContentSHA256: fmt.Sprintf("%x", sha256.Sum256(data)), PackageCount: count,
	}, nil
}

func (r *RepoMetadata) hasCurrentPrimary(primary *RepoMDData, sourceURL string) bool {
	c := r.CachedPrimary
	return c != nil && c.Format == rpmPrimaryFormat && c.SourceURL == sourceURL &&
		c.Checksum.Text != "" && c.Checksum == primary.Checksum && c.OpenChecksum == primary.OpenChecksum &&
		equalMetadataSize(c.Size, primary.Size) && equalMetadataSize(c.OpenSize, primary.OpenSize) &&
		r.CachedDataSHA256.Valid && c.ContentSHA256 == r.CachedDataSHA256.String
}

func equalMetadataSize(a, b *int64) bool {
	return (a == nil && b == nil) || (a != nil && b != nil && *a == *b)
}
