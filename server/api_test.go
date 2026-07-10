/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"net/url"
	"testing"
	"time"
)

func fixtureSnapshot() *Snapshot {
	stars := func(n int) *int { return &n }
	exts := []*Ext{
		{Name: "postgis", Pkg: "postgis", Category: "GIS", State: "available", Repo: "PGDG",
			License: "GPL-2.0", Lang: "C", PG: []int{18, 17, 16, 15, 14}, Stars: stars(9),
			EnDesc: "spatial types and functions", ExtType: "standard", LastCommit: "2026-05-11"},
		{Name: "timescaledb", Pkg: "timescaledb", Category: "TIME", State: "available", Repo: "PIGSTY",
			License: "Timescale", Lang: "C", PG: []int{18, 17, 16, 15}, Stars: stars(22656),
			EnDesc: "time-series database", ZhDesc: "时序数据库扩展插件", ExtType: "preload", LastCommit: "2026-05-18"},
		{Name: "vector", Pkg: "pgvector", Category: "RAG", State: "available", Repo: "PGDG",
			License: "PostgreSQL", Lang: "C", PG: []int{18, 17, 16, 15, 14}, Stars: stars(21351),
			EnDesc: "vector data type and ivfflat and hnsw", ExtType: "standard", LastCommit: "2026-04-01"},
		{Name: "aws_s3", Pkg: "aws_s3", Category: "ETL", State: "n/a", Repo: "n/a",
			License: "Unknown", Lang: "SQL", ExtKernel: "aws-rds-pg", ExtVendor: "AWS",
			EnDesc: "S3 import export for RDS", ExtType: "puresql"},
	}
	snap := &Snapshot{Exts: exts, ByName: map[string]*Ext{}, LoadedAt: time.Now()}
	for _, e := range exts {
		snap.ByName[e.Name] = e
		if e.State == "available" {
			snap.CountAvail++
		}
	}
	return snap
}

func TestParseFilterOperators(t *testing.T) {
	tests := []struct {
		name string
		q    string
		want Filter
	}{
		{"free words", "q=vector+index", Filter{Words: []string{"vector", "index"}}},
		{"cat operator", "q=cat:gis", Filter{Cat: "GIS"}},
		{"mixed", "q=cat:RAG+hnsw", Filter{Cat: "RAG", Words: []string{"hnsw"}}},
		{"pg operator", "q=pg:18", Filter{PG: 18}},
		{"is packaged", "q=is:packaged", Filter{Scope: "packaged"}},
		{"vendor", "q=vendor:aws", Filter{Vendor: "aws"}},
		{"explicit params", "cat=time&repo=pigsty&pg=17", Filter{Cat: "TIME", Repo: "PIGSTY", PG: 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := url.ParseQuery(tt.q)
			if err != nil {
				t.Fatal(err)
			}
			got := ParseFilter(v)
			if got.Cat != tt.want.Cat || got.Repo != tt.want.Repo || got.PG != tt.want.PG ||
				got.Scope != tt.want.Scope || got.Vendor != tt.want.Vendor ||
				len(got.Words) != len(tt.want.Words) {
				t.Errorf("ParseFilter(%q) = %+v, want %+v", tt.q, got, tt.want)
			}
		})
	}
}

func TestFilterApply(t *testing.T) {
	snap := fixtureSnapshot()
	tests := []struct {
		name  string
		query string
		sort  string
		want  []string
	}{
		{"all sorted by stars", "", "", []string{"timescaledb", "vector", "postgis", "aws_s3"}},
		{"by name", "", "name", []string{"aws_s3", "postgis", "timescaledb", "vector"}},
		{"by updated", "", "updated", []string{"timescaledb", "postgis", "vector", "aws_s3"}},
		{"category", "cat=GIS", "", []string{"postgis"}},
		{"pg 14 support", "pg=14", "", []string{"vector", "postgis"}},
		{"packaged only", "scope=packaged", "", []string{"timescaledb", "vector", "postgis"}},
		{"cloud only", "scope=cloud", "", []string{"aws_s3"}},
		{"exact name outranks stars", "q=postgis", "", []string{"postgis"}},
		{"desc match", "q=hnsw", "", []string{"vector"}},
		{"zh desc match", "q=时序", "", []string{"timescaledb"}},
		{"operator in q", "q=cat:RAG", "", []string{"vector"}},
		{"no match", "q=nosuchthing", "", []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, _ := url.ParseQuery(tt.query)
			got := ParseFilter(v).Apply(snap, tt.sort)
			names := make([]string, len(got))
			for i, e := range got {
				names[i] = e.Name
			}
			if len(names) != len(tt.want) {
				t.Fatalf("Apply(%q) = %v, want %v", tt.query, names, tt.want)
			}
			for i := range names {
				if names[i] != tt.want[i] {
					t.Fatalf("Apply(%q) = %v, want %v", tt.query, names, tt.want)
				}
			}
		})
	}
}

func TestSortOS(t *testing.T) {
	oss := []string{
		"u24.x86_64", "el8.aarch64", "d12.x86_64", "el10.x86_64", "u22.aarch64",
		"el8.x86_64", "d13.aarch64", "el9.x86_64", "u26.x86_64", "d12.aarch64",
	}
	sortOS(oss)
	want := []string{
		"el8.x86_64", "el8.aarch64", "el9.x86_64", "el10.x86_64",
		"d12.x86_64", "d12.aarch64", "d13.aarch64",
		"u22.aarch64", "u24.x86_64", "u26.x86_64",
	}
	for i := range want {
		if oss[i] != want[i] {
			t.Fatalf("sortOS = %v, want %v", oss, want)
		}
	}
}

func TestAtois(t *testing.T) {
	got := atois([]string{"14", "18", "junk", "16", ""})
	want := []int{18, 16, 14}
	if len(got) != len(want) {
		t.Fatalf("atois = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("atois = %v, want %v", got, want)
		}
	}
}

func TestTTLCache(t *testing.T) {
	c := newTTLCache(50*time.Millisecond, 2)
	c.Set("a", []byte("1"))
	if v, ok := c.Get("a"); !ok || string(v) != "1" {
		t.Fatal("expected cache hit for a")
	}
	c.Set("b", []byte("2"))
	c.Set("c", []byte("3")) // exceeds maxSize, evicts something
	if len(c.items) > 2 {
		t.Fatalf("cache exceeded max size: %d", len(c.items))
	}
	time.Sleep(60 * time.Millisecond)
	if _, ok := c.Get("a"); ok {
		t.Fatal("expected ttl expiry for a")
	}
}
