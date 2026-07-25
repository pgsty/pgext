/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

/* ----------------------------------------------------------------
   Global build-matrix payload (GET /api/v1/matrix).

   The payload is computed at write time and cached in pgext.matrix_cache
   (one mtime-stamped row per view), so the read path is one primary-key
   lookup. The builder collapses pgext.pkg into one row per package with a
   single aggregation query; Go lays the cells out on the canonical OS × PG
   grid of the active snapshot.

   Wire format "matrix-row.v2": each row is {p, e, c[, v, i]} where c holds
   one status byte per cell — G = AVAIL via PGDG, P = AVAIL via Pigsty/other,
   M = MISS, N = N/A. v is the row-local version dictionary and i its
   per-cell base36 index ('.' = none). The pkg.hide flag is deliberately
   ignored: rendering follows the four states only.

   Besides the global view, ?repo=pgdg|pigsty serves single-repository
   what-if views (mirroring the pgext.pkg_pgdg / pkg_pigsty schema views):
   the same grid recomputed as if only that organization's repositories
   existed. Rows shrink to the packages that organization actually ships,
   AVAIL cells carry its own latest version, every other currently-valid
   slot degrades to M. stats.global carries the real-world package/slot
   totals so clients can state what the missing complement amounts to.

   Freshness: a cache row is reused while it is newer than the catalog
   load recorded in pgext.status and less than a day old; POST /api/v1/reload
   rebuilds all views eagerly. Rebuilds are cheap (tens of ms) but never sit
   on the request path for regular visitors.
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

// matrixViews maps the ?repo= query value onto its cache row and the status
// letter its AVAIL cells carry ("" is the global three-color view).
var matrixViews = map[string]struct {
	id     int
	letter string
}{
	"":       {id: 1},
	"pgdg":   {id: 2, letter: "G"},
	"pigsty": {id: 3, letter: "P"},
}

// matrixViewOrder fixes the warmup/reload rebuild order (global first).
var matrixViewOrder = []string{"", "pgdg", "pigsty"}

const matrixCacheDDL = `
	CREATE TABLE IF NOT EXISTS pgext.matrix_cache (
	    id      int PRIMARY KEY CONSTRAINT matrix_cache_id_range CHECK (id BETWEEN 1 AND 3),
	    mtime   timestamptz NOT NULL DEFAULT now(),
	    payload jsonb NOT NULL
	)`

// Databases from the one-row era carry a CHECK (id = 1) under this name;
// dropping it lets the repo view rows in without touching fresh installs.
const matrixCacheLegacyCheckDDL = `
	ALTER TABLE pgext.matrix_cache DROP CONSTRAINT IF EXISTS matrix_cache_id_check`

// matrixRow is one package row of the global matrix payload.
type matrixRow struct {
	Pkg      string   `json:"p"`
	Ext      string   `json:"e"`
	Codes    string   `json:"c"`
	Versions []string `json:"v,omitempty"`
	VIdx     string   `json:"i,omitempty"`
}

// parseMatrixCell splits one encoded cell value "<version>(#|@)<state>",
// where '#' marks a hidden package row.
func parseMatrixCell(value string) (version string, hidden bool, state byte) {
	if len(value) < 2 {
		return "", false, 'N'
	}
	state = value[len(value)-1]
	switch state {
	case 'G', 'P', 'M', 'N':
	default:
		state = 'N'
	}
	return value[:len(value)-2], value[len(value)-2] == '#', state
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
		version, _, state := parseMatrixCell(value)
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
		      ORDER BY pkg, id) e ON e.pkg = p.pkg`

// matrixGlobalQuery folds live pgext.pkg into one aggregated row per package.
const matrixGlobalQuery = `
		SELECT p.pkg, e.name,
		       jsonb_object_agg(
		           split_part(p.os, '.', 1)
		             || CASE WHEN p.os LIKE '%x86%' THEN 'i' ELSE 'a' END
		             || '.' || p.pg,
		           coalesce(p.version, '')
		             || CASE WHEN p.hide THEN '#' ELSE '@' END
		             || CASE p.state
		                  WHEN 'AVAIL' THEN CASE p.org WHEN 'pgdg' THEN 'G' ELSE 'P' END
		                  WHEN 'MISS'  THEN 'M'
		                  ELSE 'N'
		                END) AS cells
		FROM pgext.pkg p` + matrixLeadJoin + `
		GROUP BY e.id, p.pkg, e.name
		ORDER BY e.id`

// matrixRepoQuery recomputes the grid as if only one repository organization
// ($1) existed — the SQL twin of the pgext.pkg_pgdg / pkg_pigsty views. Rows
// are limited to packages that organization ships at least once; AVAIL cells
// carry its own latest version (as the given status letter, always one of the
// matrixViews constants) and everything else valid degrades to MISS.
func matrixRepoQuery(letter string) string {
	return `
		SELECT p.pkg, e.name,
		       jsonb_object_agg(
		           split_part(p.os, '.', 1)
		             || CASE WHEN p.os LIKE '%x86%' THEN 'i' ELSE 'a' END
		             || '.' || p.pg,
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

// buildMatrixPayload computes one /api/v1/matrix response variant from
// pgext.pkg with one aggregation query, ordered by lead-extension catalog id.
func (a *api) buildMatrixPayload(ctx context.Context, snap *Snapshot, view string) ([]byte, error) {
	mv, ok := matrixViews[view]
	if !ok {
		return nil, fmt.Errorf("unknown matrix view %q", view)
	}
	var generated time.Time
	if err := a.pool.QueryRow(ctx, `
		SELECT coalesce(recap_time, parse_time, init_time, now()) FROM pgext.status WHERE id = 0`,
	).Scan(&generated); err != nil {
		generated = snap.LoadedAt
	}
	query, args := matrixGlobalQuery, []any(nil)
	if view != "" {
		query, args = matrixRepoQuery(mv.letter), []any{view}
	}
	rows, err := a.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("aggregate %s matrix: %w", matrixViewName(view), err)
	}
	defer rows.Close()

	keys := matrixColumnKeys(snap.OSs, snap.PGs)
	counts := map[string]int{"G": 0, "P": 0, "M": 0, "N": 0}
	out := make([]matrixRow, 0, 400)
	for rows.Next() {
		var pkg, ext string
		var raw []byte
		if err := rows.Scan(&pkg, &ext, &raw); err != nil {
			return nil, fmt.Errorf("scan %s matrix row: %w", matrixViewName(view), err)
		}
		cells := map[string]string{}
		if err := json.Unmarshal(raw, &cells); err != nil {
			return nil, fmt.Errorf("decode %s matrix row %s: %w", matrixViewName(view), pkg, err)
		}
		row := layoutMatrixRow(cells, keys, counts)
		row.Pkg, row.Ext = pkg, ext
		out = append(out, row)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate %s matrix rows: %w", matrixViewName(view), err)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("pgext.pkg holds no %s packages for the active lead extensions", matrixViewName(view))
	}
	stats := map[string]any{
		"rows": len(out), "os": len(snap.OSs), "pg": len(snap.PGs),
		"cells": len(out) * len(keys), "counts": counts,
	}
	payload := map[string]any{
		"generated": generated.UTC().Format(time.RFC3339),
		"source":    "pgext.matrix_cache",
		"format":    matrixFormat,
		"pg":        snap.PGs,
		"os":        snap.OSs,
		"stats":     stats,
		"rows":      out,
	}
	if view != "" {
		// size the real-world matrix so clients can state the complement:
		// packages this organization does not ship and slots it leaves missing
		var globalRows, globalAvail int
		if err := a.pool.QueryRow(ctx, matrixGlobalStats).Scan(&globalRows, &globalAvail); err != nil {
			return nil, fmt.Errorf("size global matrix for %s view: %w", matrixViewName(view), err)
		}
		payload["repo"] = view
		stats["global"] = map[string]any{"rows": globalRows, "avail": globalAvail}
	}
	return json.Marshal(payload)
}

// matrixViewName labels a view in logs and error chains.
func matrixViewName(view string) string {
	if view == "" {
		return "global"
	}
	return view
}

// ensureMatrixCache creates the cache table; a read-only role just loses the
// persistent cache, never the endpoint.
func (a *api) ensureMatrixCache(ctx context.Context) {
	if _, err := a.pool.Exec(ctx, matrixCacheDDL); err != nil {
		logrus.Debugf("matrix cache table unavailable (read-only role?): %v", err)
		return
	}
	if _, err := a.pool.Exec(ctx, matrixCacheLegacyCheckDDL); err != nil {
		logrus.Debugf("matrix cache legacy check not dropped: %v", err)
	}
}

// matrixPayload returns the cached payload of one view when it is still
// current, and rebuilds + upserts it otherwise. force skips the freshness
// check (reload).
func (a *api) matrixPayload(ctx context.Context, force bool, view string) ([]byte, error) {
	mv, ok := matrixViews[view]
	if !ok {
		return nil, fmt.Errorf("unknown matrix view %q", view)
	}
	snap := a.store.Get()
	if snap == nil {
		return nil, fmt.Errorf("catalog snapshot not loaded yet")
	}
	var cached []byte
	var mtime *time.Time
	if err := a.pool.QueryRow(ctx,
		`SELECT payload::text, mtime FROM pgext.matrix_cache WHERE id = $1`, mv.id,
	).Scan(&cached, &mtime); err != nil {
		cached, mtime = nil, nil
	}
	if !force && len(cached) > 0 && mtime != nil && time.Since(*mtime) < 24*time.Hour {
		var dataMtime *time.Time
		if err := a.pool.QueryRow(ctx,
			`SELECT coalesce(recap_time, parse_time, init_time) FROM pgext.status WHERE id = 0`,
		).Scan(&dataMtime); err == nil && (dataMtime == nil || !mtime.Before(*dataMtime)) {
			return cached, nil
		}
	}
	payload, err := a.buildMatrixPayload(ctx, snap, view)
	if err != nil {
		if len(cached) > 0 {
			logrus.Warnf("%s matrix rebuild failed, serving cached payload: %v", matrixViewName(view), err)
			return cached, nil
		}
		return nil, err
	}
	if _, err := a.pool.Exec(ctx, `
		INSERT INTO pgext.matrix_cache (id, mtime, payload) VALUES ($1, now(), $2::jsonb)
		ON CONFLICT (id) DO UPDATE SET mtime = EXCLUDED.mtime, payload = EXCLUDED.payload`,
		mv.id, string(payload)); err != nil {
		logrus.Debugf("matrix cache write skipped: %v", err)
	}
	return payload, nil
}

// refreshMatrixCaches rebuilds/refreshes every matrix view in order; used by
// startup warmup and POST /api/v1/reload so no visitor pays for a build.
func (a *api) refreshMatrixCaches(ctx context.Context, force bool) {
	for _, view := range matrixViewOrder {
		if _, err := a.matrixPayload(ctx, force, view); err != nil {
			logrus.Warnf("%s matrix cache refresh failed: %v", matrixViewName(view), err)
		}
	}
}

// warmMatrixCache runs at startup so the first /matrix visitor gets the
// precomputed payload instead of paying for the initial build.
func (a *api) warmMatrixCache(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	a.ensureMatrixCache(ctx)
	a.refreshMatrixCaches(ctx, false)
}
