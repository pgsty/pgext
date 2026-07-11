package cmd

import (
	"errors"
	"strings"
	"testing"
)

func TestRootCommandIncludesRescan(t *testing.T) {
	cmd, _, err := rootCmd.Find([]string{"rescan"})
	if err != nil {
		t.Fatalf("find rescan command: %v", err)
	}
	if cmd == nil || cmd.Use != "rescan" {
		t.Fatalf("expected rescan command, got %#v", cmd)
	}
}

func TestRescanCommandDocumentsCompleteParsePipeline(t *testing.T) {
	cmd, _, err := rootCmd.Find([]string{"rescan"})
	if err != nil {
		t.Fatalf("find rescan command: %v", err)
	}
	if cmd == nil {
		t.Fatal("rescan command not found")
	}

	for _, want := range []string{
		"pgext scan",
		"pgext parse",
		"pgext.pkg",
		"published atomically",
	} {
		if !strings.Contains(cmd.Long, want) {
			t.Fatalf("rescan long help missing %q: %s", want, cmd.Long)
		}
	}
	if strings.Contains(cmd.Long, "pgext recap") {
		t.Fatalf("rescan help still documents recap as a required separate command: %s", cmd.Long)
	}
}

type fakeParsePipeline struct {
	coreCalls int
	fullCalls int
	coreErr   error
	fullErr   error
}

func (p *fakeParsePipeline) ParseAll() error {
	p.coreCalls++
	return p.coreErr
}

func (p *fakeParsePipeline) ParseAndRecap() error {
	p.fullCalls++
	return p.fullErr
}

func TestRunParsePipelineDefaultsToCompleteCatalog(t *testing.T) {
	parser := &fakeParsePipeline{}
	if err := runParsePipeline(parser, false); err != nil {
		t.Fatalf("run complete parse: %v", err)
	}
	if parser.fullCalls != 1 || parser.coreCalls != 0 {
		t.Fatalf("parse calls = full:%d core:%d, want full:1 core:0", parser.fullCalls, parser.coreCalls)
	}
}

func TestRunParsePipelineNoPkg(t *testing.T) {
	parser := &fakeParsePipeline{}
	if err := runParsePipeline(parser, true); err != nil {
		t.Fatalf("run no-pkg parse: %v", err)
	}
	if parser.fullCalls != 0 || parser.coreCalls != 1 {
		t.Fatalf("parse calls = full:%d core:%d, want full:0 core:1", parser.fullCalls, parser.coreCalls)
	}
}

func TestRunParsePipelinePreservesErrors(t *testing.T) {
	sentinel := errors.New("parse failed")
	for _, tt := range []struct {
		name   string
		noPkg  bool
		parser *fakeParsePipeline
	}{
		{name: "complete", parser: &fakeParsePipeline{fullErr: sentinel}},
		{name: "no pkg", noPkg: true, parser: &fakeParsePipeline{coreErr: sentinel}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := runParsePipeline(tt.parser, tt.noPkg); !errors.Is(err, sentinel) {
				t.Fatalf("runParsePipeline() error = %v, want wrapped sentinel", err)
			}
		})
	}
}
