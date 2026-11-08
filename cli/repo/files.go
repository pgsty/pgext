package repo

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"syscall"
	"time"
)

const stateDir = "var/lib/pgext/repo"

type fileImage struct {
	Content string `json:"content"`
	Mode    uint32 `json:"mode"`
}

type Change struct {
	Path   string     `json:"path"`
	Before *fileImage `json:"before"`
	After  *fileImage `json:"after"`
}

type InstalledFile struct {
	Path    string `json:"path"`
	Managed bool   `json:"managed"`
	SHA256  string `json:"sha256"`
}

type backup struct {
	Version  int       `json:"version"`
	Platform Platform  `json:"platform"`
	Created  time.Time `json:"created"`
	Changes  []Change  `json:"changes"`
}

func (m *Manager) path(relative string) (string, error) {
	if !filepath.IsLocal(relative) || filepath.Clean(relative) != relative {
		return "", fmt.Errorf("unsafe repository path %q", relative)
	}
	p := m.Root
	for _, part := range strings.Split(relative, string(filepath.Separator)) {
		p = filepath.Join(p, part)
		info, err := os.Lstat(p)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("refusing symlink %s", p)
		}
	}
	return p, nil
}

func (m *Manager) read(path string) (*fileImage, error) {
	p, err := m.path(path)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(p)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("not a regular file: %s", p)
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return &fileImage{Content: string(data), Mode: uint32(info.Mode().Perm())}, nil
}

func sameImage(a, b *fileImage) bool {
	return a == nil && b == nil || a != nil && b != nil && *a == *b
}

func owned(path string, image *fileImage) bool {
	return image != nil && strings.HasPrefix(filepath.Base(path), "pgext-") && strings.HasPrefix(image.Content, managedHeader)
}

func validRepoPath(path, family string) bool {
	if family == "rpm" {
		return filepath.Dir(path) == "etc/yum.repos.d" && strings.HasSuffix(path, ".repo")
	}
	return path == "etc/apt/sources.list" || filepath.Dir(path) == "etc/apt/sources.list.d" && (strings.HasSuffix(path, ".list") || strings.HasSuffix(path, ".sources"))
}

func (m *Manager) scan() (map[string]*fileImage, error) {
	dir := "etc/apt/sources.list.d"
	if m.Platform.Family == "rpm" {
		dir = "etc/yum.repos.d"
	}
	p, err := m.path(dir)
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(p)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	paths := []string{}
	if m.Platform.Family == "deb" {
		paths = append(paths, "etc/apt/sources.list")
	}
	for _, ent := range entries {
		path := filepath.Join(dir, ent.Name())
		if validRepoPath(path, m.Platform.Family) {
			paths = append(paths, path)
		}
	}
	files := map[string]*fileImage{}
	for _, path := range paths {
		image, err := m.read(path)
		if err != nil {
			return nil, err
		}
		if image != nil {
			files[path] = image
		}
	}
	return files, nil
}

func installed(root string, files map[string]*fileImage) []InstalledFile {
	result := []InstalledFile{}
	for path, f := range files {
		sum := sha256.Sum256([]byte(f.Content))
		result = append(result, InstalledFile{Path: filepath.Join(root, path), Managed: owned(path, f), SHA256: hex.EncodeToString(sum[:])})
	}
	slices.SortFunc(result, func(a, b InstalledFile) int { return strings.Compare(a.Path, b.Path) })
	return result
}

func (m *Manager) write(path string, image *fileImage) error {
	p, err := m.path(path)
	if err != nil {
		return err
	}
	if image == nil {
		err := os.Remove(p)
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return err
	}
	f, err := os.CreateTemp(filepath.Dir(p), ".pgext-")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	defer f.Close()
	if _, err := f.WriteString(image.Content); err != nil {
		return err
	}
	if err := f.Chmod(os.FileMode(image.Mode)); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), p)
}

func (m *Manager) lock() (func(), error) {
	p, err := m.path(stateDir)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(p, 0700); err != nil {
		return nil, err
	}
	p, err = m.path(stateDir + "/lock")
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		f.Close()
		return nil, fmt.Errorf("another pgext repo operation is running: %w", err)
	}
	return func() { _ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN); _ = f.Close() }, nil
}

func (m *Manager) saveBackup(changes []Change) (string, error) {
	dir, err := m.path(stateDir + "/backups")
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}
	f, err := os.CreateTemp(dir, time.Now().UTC().Format("20060102T150405Z-")+"*.json")
	if err != nil {
		return "", err
	}
	defer f.Close()
	data := backup{Version: 1, Platform: m.Platform, Created: time.Now().UTC(), Changes: changes}
	if err := json.NewEncoder(f).Encode(data); err != nil {
		return "", err
	}
	if err := f.Sync(); err != nil {
		return "", err
	}
	return filepath.Base(f.Name()), nil
}

func (m *Manager) restoreChanges(id string) ([]Change, error) {
	if filepath.Base(id) != id || !strings.HasSuffix(id, ".json") {
		return nil, fmt.Errorf("restore requires a backup ID from pgext repo, not a path")
	}
	f, err := m.read(stateDir + "/backups/" + id)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, fmt.Errorf("backup %q does not exist", id)
	}
	var saved backup
	dec := json.NewDecoder(bytes.NewBufferString(f.Content))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&saved); err != nil {
		return nil, err
	}
	if saved.Version != 1 || saved.Platform != m.Platform || len(saved.Changes) == 0 {
		return nil, fmt.Errorf("backup version/platform does not match this target")
	}
	seen := map[string]bool{}
	var changes []Change
	for _, c := range saved.Changes {
		if !validRepoPath(c.Path, m.Platform.Family) || seen[c.Path] {
			return nil, fmt.Errorf("invalid or duplicate backup path %q", c.Path)
		}
		seen[c.Path] = true
		for _, image := range []*fileImage{c.Before, c.After} {
			if image != nil && image.Mode & ^uint32(0777) != 0 {
				return nil, fmt.Errorf("invalid file mode in backup")
			}
		}
		current, err := m.read(c.Path)
		if err != nil {
			return nil, err
		}
		if !sameImage(current, c.After) {
			return nil, fmt.Errorf("%s changed since backup %s; refusing to overwrite it", c.Path, id)
		}
		changes = append(changes, Change{Path: c.Path, Before: c.After, After: c.Before})
	}
	return changes, nil
}

func (m *Manager) rollback(changes []Change) error {
	var failures []error
	for i := len(changes) - 1; i >= 0; i-- {
		c := changes[i]
		current, err := m.read(c.Path)
		if err == nil && sameImage(current, c.Before) {
			continue
		}
		if err == nil && !sameImage(current, c.After) {
			err = fmt.Errorf("%s changed concurrently during rollback", c.Path)
		}
		if err == nil {
			err = m.write(c.Path, c.Before)
		}
		if err != nil {
			failures = append(failures, err)
		}
	}
	return errors.Join(failures...)
}
