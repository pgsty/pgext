package cli

import (
	"context"
	"strings"
	"testing"
)

func TestAPTParserRejectsMalformedNonEmptyMetadata(t *testing.T) {
	parser := newAPTParser(context.Background(), liveTable("apt"))

	for _, tt := range []struct {
		name string
		data string
		want string
	}{
		{name: "not control data", data: "invalid", want: "Package"},
		{name: "missing version", data: "Package: demo\nArchitecture: amd64\nFilename: demo.deb", want: "Version"},
		{name: "missing architecture", data: "Package: demo\nVersion: 1.0\nFilename: demo.deb", want: "Architecture"},
		{name: "missing filename", data: "Package: demo\nVersion: 1.0\nArchitecture: amd64", want: "Filename"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			packages, err := parser.parsePackages(tt.data)
			if err == nil {
				t.Fatalf("malformed metadata parsed as %d package(s)", len(packages))
			}
			if !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("error %q does not mention %q", err, tt.want)
			}
		})
	}
}

func TestAPTParserAcceptsEmptyAndValidMetadata(t *testing.T) {
	parser := newAPTParser(context.Background(), liveTable("apt"))
	for _, empty := range []string{"", " \n\t", "\r\n"} {
		packages, err := parser.parsePackages(empty)
		if err != nil || len(packages) != 0 {
			t.Fatalf("empty metadata = (%d, %v), want (0, nil)", len(packages), err)
		}
	}

	packages, err := parser.parsePackages("Package: demo\r\nVersion: 1.0\r\nArchitecture: amd64\r\nFilename: pool/demo.deb\r\n")
	if err != nil {
		t.Fatalf("valid metadata: %v", err)
	}
	if len(packages) != 1 || packages[0]["Package"] != "demo" {
		t.Fatalf("valid metadata parsed as %#v", packages)
	}
}

func TestAPTParserAcceptsCaseInsensitiveFieldsAndTabContinuation(t *testing.T) {
	parser := newAPTParser(context.Background(), liveTable("apt"))
	packages, err := parser.parsePackages(strings.Join([]string{
		"package: first",
		"VERSION: 1.0",
		"architecture: amd64",
		"filename: pool/first.deb",
		"description: first line",
		"\tsecond line",
		" \t",
		"Package: second",
		"Version: 2.0",
		"Architecture: arm64",
		"Filename: pool/second.deb",
	}, "\n"))
	if err != nil {
		t.Fatalf("parse Debian control variants: %v", err)
	}
	if len(packages) != 2 {
		t.Fatalf("parsed %d packages, want 2", len(packages))
	}
	if got := packages[0]["Description"]; got != "first line\nsecond line" {
		t.Fatalf("multiline Description = %#v, want two joined lines", got)
	}
	if got := packages[1]["Package"]; got != "second" {
		t.Fatalf("second Package = %#v, want second", got)
	}
}
