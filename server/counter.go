/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// counterStore tracks per-extension visit/download counts backed by
// pgext.counter. Page requests only touch the in-memory maps; a background
// flusher persists the accumulated increments once per minute with an
// additive UPDATE (visit = visit + delta) and then reads the global totals
// back. Because only increments are written, any number of pgext processes
// can share one counter table without conflicting on absolute values.
//
// When the counter table is missing (schema not migrated yet) the store
// degrades to memory-only counting and keeps retrying on every flush tick,
// so applying the DDL later heals persistence without a restart.
type counterStore struct {
	mu    sync.Mutex
	pool  *pgxpool.Pool
	ready bool           // pgext.counter reachable and primed
	names []string       // known extension names, used to (re)prime the table
	base  map[string]counterVal // totals last read back from the database
	delta map[string]counterVal // local increments not yet flushed
}

type counterVal struct{ Visit, Download int64 }

func newCounterStore(pool *pgxpool.Pool) *counterStore {
	return &counterStore{
		pool:  pool,
		base:  map[string]counterVal{},
		delta: map[string]counterVal{},
	}
}

// prime makes sure every known extension has a counter row, then loads the
// current totals. All extensions are enumerable (~1.6k), so the whole table
// lives comfortably in memory.
func (s *counterStore) prime(ctx context.Context, names []string) {
	s.mu.Lock()
	if len(names) > 0 {
		s.names = names
	}
	names = s.names
	s.mu.Unlock()
	if s.pool == nil {
		return
	}
	if len(names) > 0 {
		if _, err := s.pool.Exec(ctx, `
			INSERT INTO pgext.counter (ext)
			SELECT unnest($1::text[]) ON CONFLICT (ext) DO NOTHING`, names); err != nil {
			logrus.Warnf("visit counter unavailable (memory-only until it appears): %v", err)
			return
		}
	}
	base, err := s.readTotals(ctx)
	if err != nil {
		logrus.Warnf("visit counter read failed (memory-only until it recovers): %v", err)
		return
	}
	s.mu.Lock()
	s.base = base
	s.ready = true
	s.mu.Unlock()
	logrus.Infof("visit counter primed: %d extensions tracked in pgext.counter", len(base))
}

func (s *counterStore) readTotals(ctx context.Context) (map[string]counterVal, error) {
	rows, err := s.pool.Query(ctx, `SELECT ext, visit, download FROM pgext.counter`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	totals := make(map[string]counterVal, 2048)
	for rows.Next() {
		var ext string
		var v counterVal
		if err := rows.Scan(&ext, &v.Visit, &v.Download); err != nil {
			return nil, err
		}
		totals[ext] = v
	}
	return totals, rows.Err()
}

// importLegacy migrates counts from the retired file-backed counter exactly
// once: the file is renamed first (claiming it against concurrent processes),
// then its counts enter the delta map and reach the database on the next
// flush.
func (s *counterStore) importLegacy(path string) {
	if path == "" {
		return
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return // no legacy file, nothing to migrate
	}
	var legacy struct {
		Visits    map[string]int64 `json:"visits"`
		Downloads map[string]int64 `json:"downloads"`
	}
	if err := json.Unmarshal(b, &legacy); err != nil {
		logrus.Warnf("legacy visit file %s is corrupt, skipping import: %v", path, err)
		return
	}
	if err := os.Rename(path, path+".imported"); err != nil {
		logrus.Warnf("cannot claim legacy visit file %s, skipping import: %v", path, err)
		return
	}
	s.mu.Lock()
	var total int64
	for ext, n := range legacy.Visits {
		v := s.delta[ext]
		v.Visit += n
		s.delta[ext] = v
		total += n
	}
	for ext, n := range legacy.Downloads {
		v := s.delta[ext]
		v.Download += n
		s.delta[ext] = v
	}
	s.mu.Unlock()
	logrus.Infof("imported %d visits from legacy counter file %s", total, path)
}

// hit records one page visit and returns the updated totals.
func (s *counterStore) hit(name string) (visits, downloads int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v := s.delta[name]
	v.Visit++
	s.delta[name] = v
	b := s.base[name]
	return b.Visit + v.Visit, b.Download + v.Download
}

// get returns the current totals (database baseline plus unflushed local increments).
func (s *counterStore) get(name string) (visits, downloads int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, d := s.base[name], s.delta[name]
	return b.Visit + d.Visit, b.Download + d.Download
}

// flush persists pending increments additively and refreshes the baseline
// with the global totals (which include increments from other processes).
// On failure the increments are merged back, so counts are never lost while
// the process lives.
func (s *counterStore) flush(ctx context.Context) {
	if s.pool == nil {
		return
	}
	s.mu.Lock()
	ready := s.ready
	pending := s.delta
	s.delta = map[string]counterVal{}
	s.mu.Unlock()

	restore := func() {
		s.mu.Lock()
		for ext, v := range pending {
			d := s.delta[ext]
			d.Visit += v.Visit
			d.Download += v.Download
			s.delta[ext] = d
		}
		s.mu.Unlock()
	}

	if !ready {
		restore()
		s.prime(ctx, nil) // keep retrying until the table shows up
		return
	}
	if len(pending) > 0 {
		exts := make([]string, 0, len(pending))
		visits := make([]int64, 0, len(pending))
		downloads := make([]int64, 0, len(pending))
		for ext, v := range pending {
			exts = append(exts, ext)
			visits = append(visits, v.Visit)
			downloads = append(downloads, v.Download)
		}
		if _, err := s.pool.Exec(ctx, `
			INSERT INTO pgext.counter AS c (ext, visit, download)
			SELECT * FROM unnest($1::text[], $2::bigint[], $3::bigint[])
			ON CONFLICT (ext) DO UPDATE
			SET visit = c.visit + EXCLUDED.visit, download = c.download + EXCLUDED.download`,
			exts, visits, downloads); err != nil {
			logrus.Warnf("visit counter flush failed (retaining %d increments): %v", len(pending), err)
			restore()
			s.mu.Lock()
			s.ready = false // revalidate the table on the next tick
			s.mu.Unlock()
			return
		}
	}
	base, err := s.readTotals(ctx)
	if err != nil {
		logrus.Debugf("visit counter read-back failed, keeping previous baseline: %v", err)
		return // increments are persisted; only the baseline refresh is skipped
	}
	s.mu.Lock()
	s.base = base
	s.mu.Unlock()
}

// startFlusher persists increments every interval and once more on shutdown.
func (s *counterStore) startFlusher(ctx context.Context, interval time.Duration) {
	if s.pool == nil {
		return
	}
	if interval <= 0 {
		interval = time.Minute
	}
	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				// ctx is already cancelled — the final flush gets its own deadline
				shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				s.flush(shutCtx)
				cancel()
				return
			case <-t.C:
				tickCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
				s.flush(tickCtx)
				cancel()
			}
		}
	}()
}
