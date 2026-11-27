/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

/* ----------------------------------------------------------------
   Global build-matrix payload (GET /api/v1/matrix).

   The builder collapses pgext.pkg into one row per package with a single
   aggregation query; Go lays the cells out on the canonical OS × PG grid of
   the active snapshot and caches each view in memory per snapshot version
   with in-flight deduplication (no runtime DDL or database cache table).

   Wire format "matrix-row.v2": each row is {p, e, c[, v, i]} where c holds
   one status byte per cell — G = AVAIL via PGDG, P = AVAIL via Pigsty/other,
   M = MISS, N = N/A. v is the row-local version dictionary and i its
   per-cell base36 index ('.' = none).

   Besides the global view, ?repo=pgdg|pigsty serves single-repository
   what-if views (mirroring pgext.pkg_pgdg / pkg_pigsty): rows shrink to the
   packages that organization ships, AVAIL cells carry its own latest version,
   every other valid slot degrades to M, and stats.global sizes the full grid.
   ---------------------------------------------------------------- */

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const matrixFormat = "matrix-row.v2"

// matrixView is one /api/v1/matrix variant: the ?repo= query value and the
// status letter its AVAIL cells carry ("" is the global multi-repo view).
type matrixView struct {
	name   string
	letter string
}

// matrixViews lists every variant in warmup/reload rebuild order (global first).
var matrixViews = []matrixView{
	{name: ""},
	{name: "pgdg", letter: "G"},
	{name: "pigsty", letter: "P"},
}

func matrixViewOf(name string) (matrixView, bool) {
	for _, mv := range matrixViews {
		if mv.name == name {
			return mv, true
		}
	}
	return matrixView{}, false
}

// label names the view in logs and error chains.
func (mv matrixView) label() string {
	if mv.name == "" {
		return "global"
	}
	return mv.name
}

type matrixCacheEntry struct {
	ver      string
	payload  []byte
	building chan struct{}
}

// matrixRow is one package row of the global matrix payload.
type matrixRow struct {
	Pkg      string   `json:"p"`
	Ext      string   `json:"e"`
	Codes    string   `json:"c"`
	Versions []string `json:"v,omitempty"`
	VIdx     string   `json:"i,omitempty"`
}

// parseMatrixCell splits one encoded cell value "<version>@<state>".
func parseMatrixCell(value string) (version string, state byte) {
	if len(value) < 2 {
		return "", 'N'
	}
	state = value[len(value)-1]
	switch state {
	case 'G', 'P', 'M', 'N':
	default:
		state = 'N'
	}
	return value[:len(value)-2], state
}

// matrixColumnKeys returns the aggregation key ("el8i.17") for every cell of
// the canonical os × pg grid in row-major (os outer, pg inner) order.
func matrixColumnKeys(oss []string, pgs []int) []string {
	keys := make([]string, 0, len(oss)*len(pgs))
	for _, os := range oss {
		short, _, _ := strings.Cut(os, ".")
		arch := "a"
		if strings.Contains(os, "x86") {
			arch = "i"
		}
		for _, pg := range pgs {
			keys = append(keys, short+arch+"."+strconv.Itoa(pg))
		}
	}
	return keys
}

const matrixBase36 = "0123456789abcdefghijklmnopqrstuvwxyz"

// layoutMatrixRow places one package's aggregated cells onto the canonical
// grid, producing the status string, the row-local version dictionary and its
// per-cell index. counts is keyed by state letter.
func layoutMatrixRow(cells map[string]string, keys []string, counts map[string]int) matrixRow {
	var m matrixRow
	codes := make([]byte, len(keys))
	vidx := make([]byte, len(keys))
	seen := map[string]int{}
	hasVersion := false
	for ci, key := range keys {
		codes[ci], vidx[ci] = 'N', '.'
		value, ok := cells[key]
		if !ok {
			counts["N"]++
			continue
		}
		version, state := parseMatrixCell(value)
		counts[string(state)]++
		codes[ci] = state
		if version != "" && (state == 'G' || state == 'P') {
			idx, known := seen[version]
			if !known && len(m.Versions) < len(matrixBase36) {
				idx = len(m.Versions)
				seen[version] = idx
				m.Versions = append(m.Versions, version)
				known = true
			}
			if known {
				vidx[ci] = matrixBase36[idx]
				hasVersion = true
			}
		}
	}
	m.Codes = string(codes)
	if hasVersion {
		m.VIdx = string(vidx)
	} else {
		m.Versions = nil
	}
	return m
}

// matrixLeadJoin resolves each package family to its lead catalog extension;
// every per-view aggregation shares it so rows keep the same identity/order.
const matrixLeadJoin = `
		JOIN (SELECT DISTINCT ON (pkg) pkg, id, name
		      FROM pgext.extension
		      WHERE lead AND (state IS NULL OR state <> 'not-ready')
		      ORDER BY pkg, id) e ON e.pkg = p.pkg
		JOIN pgext.active_pg ap ON ap.pg = p.pg
		JOIN pgext.active_os ao ON ao.os = p.os`

// matrixCellAggKey is the jsonb_object_agg key expression, producing the same
// "el8i.17"-style cell keys matrixColumnKeys lays out on the Go side.
const matrixCellAggKey = `split_part(p.os, '.', 1)
		             || CASE WHEN p.os LIKE '%x86%' THEN 'i' ELSE 'a' END
		             || '.' || p.pg`

// matrixGlobalQuery folds live pgext.pkg into one aggregated row per package.
const matrixGlobalQuery = `
		SELECT p.pkg, e.name,
		       jsonb_object_agg(` + matrixCellAggKey + `,
		           coalesce(p.version, '') || '@'
		             || CASE p.state
		                  WHEN 'AVAIL' THEN CASE p.org WHEN 'pgdg' THEN 'G' ELSE 'P' END
		                  WHEN 'MISS'  THEN 'M'
		                  ELSE 'N'
		                END) AS cells
		FROM pgext.pkg p` + matrixLeadJoin + `
		GROUP BY e.id, p.pkg, e.name
		ORDER BY e.id`

// matrixRepoQuery recomputes the grid as if only one repository organization
// ($1) existed — the SQL twin of pgext.pkg_pgdg / pkg_pigsty.
func matrixRepoQuery(letter string) string {
	return `
		SELECT p.pkg, e.name,
		       jsonb_object_agg(` + matrixCellAggKey + `,
		           CASE WHEN o.count > 0 THEN coalesce(o.version, '') || '@` + letter + `'
		                WHEN p.state = 'N/A' THEN '@N'
		                ELSE '@M'
		           END) AS cells
		FROM pgext.pkg p` + matrixLeadJoin + `
		LEFT JOIN (SELECT b.pg, b.os, b.name, count(*) AS count,
		                  (array_agg(b.version ORDER BY b.ver::pgext.version USING OPERATOR (pgext.>)))[1] AS version
		           FROM pgext.bin b JOIN pgext.repository r ON r.id = b.repo
		           WHERE r.org = $1
		           GROUP BY b.pg, b.os, b.name) o
		       ON o.pg = p.pg AND o.os = p.os AND o.name = p.name
		GROUP BY e.id, p.pkg, e.name
		HAVING count(*) FILTER (WHERE o.count > 0) > 0
		ORDER BY e.id`
}

// matrixGlobalStats sizes the real-world matrix so what-if payloads can say
// how much of it their organization covers.
const matrixGlobalStats = `
		SELECT count(DISTINCT p.pkg), count(*) FILTER (WHERE p.state = 'AVAIL')
		FROM pgext.pkg p` + matrixLeadJoin

// catalogMtime reads the last catalog load time recorded in pgext.status.
func (a *api) catalogMtime(ctx context.Context) (*time.Time, error) {
	var t *time.Time
	err := a.pool.QueryRow(ctx,
		`SELECT coalesce(recap_time, parse_time, init_time) FROM pgext.status WHERE id = 0`,
	).Scan(&t)
	return t, err
}

// buildMatrixPayload computes one /api/v1/matrix response variant from
// pgext.pkg with one aggregation query, ordered by lead-extension catalog id.
func (a *api) buildMatrixPayload(ctx context.Context, snap *Snapshot, mv matrixView) ([]byte, error) {
	generated := snap.LoadedAt
	if t, err := a.catalogMtime(ctx); err == nil && t != nil {
		generated = *t
	}
	query, args := matrixGlobalQuery, []any(nil)
	if mv.name != "" {
		query, args = matrixRepoQuery(mv.letter), []any{mv.name}
	}
	rows, err := a.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("aggregate %s matrix: %w", mv.label(), err)
	}
	defer rows.Close()

	keys := matrixColumnKeys(snap.OSs, snap.PGs)
	counts := map[string]int{"G": 0, "P": 0, "M": 0, "N": 0}
	out := make([]matrixRow, 0, 400)
	for rows.Next() {
		var pkg, ext string
		var raw []byte
		if err := rows.Scan(&pkg, &ext, &raw); err != nil {
			return nil, fmt.Errorf("scan %s matrix row: %w", mv.label(), err)
		}
		cells := map[string]string{}
		if err := json.Unmarshal(raw, &cells); err != nil {
			return nil, fmt.Errorf("decode %s matrix row %s: %w", mv.label(), pkg, err)
		}
		row := layoutMatrixRow(cells, keys, counts)
		row.Pkg, row.Ext = pkg, ext
		out = append(out, row)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate %s matrix rows: %w", mv.label(), err)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("pgext.pkg holds no %s packages for the active lead extensions", mv.label())
	}
	stats := map[string]any{
		"rows": len(out), "os": len(snap.OSs), "pg": len(snap.PGs),
		"cells": len(out) * len(keys), "counts": counts,
	}
	payload := map[string]any{
		"generated": generated.UTC().Format(time.RFC3339),
		"source":    "pgext.pkg",
		"format":    matrixFormat,
		"pg":        snap.PGs,
		"os":        snap.OSs,
		"stats":     stats,
		"rows":      out,
	}
	if mv.name != "" {
		var globalRows, globalAvail int
		if err := a.pool.QueryRow(ctx, matrixGlobalStats).Scan(&globalRows, &globalAvail); err != nil {
			return nil, fmt.Errorf("size global matrix for %s view: %w", mv.label(), err)
		}
		payload["repo"] = mv.name
		stats["global"] = map[string]any{"rows": globalRows, "avail": globalAvail}
	}
	return json.Marshal(payload)
}

// matrixPayload returns the in-memory cached payload of one view for the
// active snapshot, coalescing concurrent rebuilds and falling back to the last
// valid payload if a transient database error occurs.
func (a *api) matrixPayload(ctx context.Context, force bool, mv matrixView) ([]byte, error) {
	snap := a.store.Get()
	if snap == nil {
		return nil, fmt.Errorf("catalog snapshot not loaded yet")
	}
	ver := snap.queryVersion()

	for {
		a.matrixMu.Lock()
		if a.matrixCache == nil {
			a.matrixCache = map[string]*matrixCacheEntry{}
		}
		entry := a.matrixCache[mv.name]
		if !force && entry != nil && entry.ver == ver && len(entry.payload) > 0 {
			out := entry.payload
			a.matrixMu.Unlock()
			return out, nil
		}
		if entry != nil && entry.building != nil {
			wait := entry.building
			a.matrixMu.Unlock()
			select {
			case <-wait:
				force = false // reuse the result of the build we waited for
				continue
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		if entry == nil {
			entry = &matrixCacheEntry{}
			a.matrixCache[mv.name] = entry
		}
		done := make(chan struct{})
		entry.building = done
		stale := entry.payload
		a.matrixMu.Unlock()

		payload, err := a.buildMatrixPayload(ctx, snap, mv)

		a.matrixMu.Lock()
		if err == nil {
			entry.ver = ver
			entry.payload = payload
		}
		entry.building = nil
		close(done)
		a.matrixMu.Unlock()

		if err != nil {
			if len(stale) > 0 {
				logrus.Warnf("%s matrix rebuild failed, serving stale payload: %v", mv.label(), err)
				return stale, nil
			}
			return nil, err
		}
		return payload, nil
	}
}

// refreshMatrixCaches rebuilds every matrix view in order; used by startup
// warmup and POST /api/v1/reload so visitors hit warm memory.
func (a *api) refreshMatrixCaches(ctx context.Context, force bool) {
	for _, mv := range matrixViews {
		if _, err := a.matrixPayload(ctx, force, mv); err != nil {
			logrus.Warnf("%s matrix cache refresh failed: %v", mv.label(), err)
		}
	}
}

// warmMatrixCache precomputes all views in the background at startup.
func (a *api) warmMatrixCache(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	a.refreshMatrixCaches(ctx, false)
}
