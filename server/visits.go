/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// visitStore is a tiny file-backed page hit counter. Counts live in memory and
// are flushed to disk every 30 seconds plus once on shutdown; it is best-effort
// by design — a lost flush costs a few counts, never a request. With an empty
// path counts still work for the process lifetime but are not persisted.
type visitStore struct {
	mu    sync.Mutex
	path  string
	dirty bool

	Visits    map[string]int64 `json:"visits"`
	Downloads map[string]int64 `json:"downloads"`
}

func newVisitStore(path string) *visitStore {
	s := &visitStore{path: path, Visits: map[string]int64{}, Downloads: map[string]int64{}}
	if path == "" {
		return s
	}
	if b, err := os.ReadFile(path); err == nil {
		if err := json.Unmarshal(b, s); err != nil {
			logrus.Warnf("visit counter file %s is corrupt, starting fresh: %v", path, err)
		}
		if s.Visits == nil {
			s.Visits = map[string]int64{}
		}
		if s.Downloads == nil {
			s.Downloads = map[string]int64{}
		}
	}
	return s
}

func (s *visitStore) hit(name string) (visits, downloads int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Visits[name]++
	s.dirty = true
	return s.Visits[name], s.Downloads[name]
}

func (s *visitStore) get(name string) (visits, downloads int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Visits[name], s.Downloads[name]
}

func (s *visitStore) flush() {
	s.mu.Lock()
	if !s.dirty || s.path == "" {
		s.mu.Unlock()
		return
	}
	b, err := json.Marshal(s)
	s.dirty = false
	s.mu.Unlock()
	if err != nil {
		return
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		logrus.Debugf("visit counter dir: %v", err)
		return
	}
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		logrus.Debugf("visit counter save: %v", err)
		return
	}
	if err := os.Rename(tmp, s.path); err != nil {
		logrus.Debugf("visit counter save: %v", err)
	}
}

func (s *visitStore) startFlusher(ctx context.Context) {
	if s.path == "" {
		return
	}
	go func() {
		t := time.NewTicker(30 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				s.flush()
				return
			case <-t.C:
				s.flush()
			}
		}
	}()
}
