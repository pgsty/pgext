package build

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

var buildWorkspaceName = regexp.MustCompile(`^(?:build(?:-[0-9]+)?|target)$`)
var specDirFilePattern = regexp.MustCompile(`%\{_specdir\}/([A-Za-z0-9_][A-Za-z0-9_.+-]*)(?:[ \t\r"';|&()<>]|$)`)

func fileExists(path string) bool { i, err := os.Stat(path); return err == nil && i.Mode().IsRegular() }

func AtomicWrite(path string, b []byte, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.CreateTemp(filepath.Dir(path), ".pgext-*")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err = f.Write(b); err != nil {
		f.Close()
		return err
	}
	if err = f.Chmod(mode); err != nil {
		f.Close()
		return err
	}
	if err = f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}

func fileInfo(path string) (File, error) {
	f, err := os.Open(path)
	if err != nil {
		return File{}, err
	}
	defer f.Close()
	h := sha256.New()
	n, err := io.Copy(h, f)
	return File{Path: path, SHA256: hex.EncodeToString(h.Sum(nil)), Size: n}, err
}

func copyFile(src, dst string, mode os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	if err = os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_EXCL, mode)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, in)
	if err == nil {
		err = out.Chmod(mode)
	}
	closeErr := out.Close()
	if err != nil {
		return err
	}
	return closeErr
}

// Copy recipe inputs only. Never copy shared source caches or old artifacts.
func copyRecipe(src, dst string) error {
	return walkRecipe(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		to := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(to, 0755)
		}
		i, err := d.Info()
		if err != nil {
			return err
		}
		if i.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(path)
			if err != nil {
				return err
			}
			resolved := filepath.Clean(filepath.Join(filepath.Dir(rel), target))
			if filepath.IsAbs(target) || resolved == ".." || strings.HasPrefix(resolved, "../") {
				return fmt.Errorf("recipe symlink escapes source tree: %s", path)
			}
			return os.Symlink(target, to)
		}
		if !i.Mode().IsRegular() {
			return fmt.Errorf("unsupported recipe file: %s", path)
		}
		return copyFile(path, to, i.Mode().Perm())
	})
}

func walkRecipe(src string, visit fs.WalkDirFunc) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path != src {
			skip := d.IsDir() && buildWorkspaceName.MatchString(d.Name())
			switch d.Name() {
			case ".git", "SOURCES", "RPMS", "SRPMS", "DEBS", "BUILD", "BUILDROOT":
				skip = true
			}
			for _, suffix := range []string{".deb", ".ddeb", ".rpm", ".buildinfo", ".changes", ".log"} {
				skip = skip || strings.HasSuffix(d.Name(), suffix)
			}
			if skip {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		return visit(path, d, nil)
	})
}

// Include packaging controls, patches, helpers, permissions and symlink targets
// in the input fingerprint, using exactly the same exclusions as recipe copies.
func recipeTreeHash(root string) (string, error) {
	h := sha256.New()
	err := walkRecipe(root, func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		var digest string
		if info.Mode()&os.ModeSymlink != 0 {
			digest, err = os.Readlink(path)
		} else if info.Mode().IsRegular() {
			var f File
			f, err = fileInfo(path)
			digest = f.SHA256
		} else {
			return fmt.Errorf("unsupported recipe file: %s", path)
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(h, "%s\x00%d\x00%s\x00", rel, info.Mode(), digest)
		return nil
	})
	return hex.EncodeToString(h.Sum(nil)), err
}

// rpmRecipeFiles lists the spec and its literal top-level SPECS/ inputs.
// Paths under patches/ are covered by the shared patch directory.
func rpmRecipeFiles(recipePath string) ([]string, error) {
	content, err := os.ReadFile(recipePath)
	if err != nil {
		return nil, err
	}
	spec := filepath.Base(recipePath)
	var auxiliary []string
	for line := range strings.SplitSeq(string(content), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}
		for _, match := range specDirFilePattern.FindAllStringSubmatch(line, -1) {
			if match[1] != spec && match[1] != "patches" {
				auxiliary = append(auxiliary, match[1])
			}
		}
	}
	slices.Sort(auxiliary)
	return append([]string{spec}, slices.Compact(auxiliary)...), nil
}

// copyRPMRecipe stages the target recipe, its local inputs and patches.
func copyRPMRecipe(specsDir, pkg, dstSpecsDir string) error {
	names, err := rpmRecipeFiles(filepath.Join(specsDir, pkg+".spec"))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dstSpecsDir, 0755); err != nil {
		return err
	}
	for _, name := range names {
		path := filepath.Join(specsDir, name)
		info, err := rpmRecipeFileInfo(path)
		if err != nil {
			return err
		}
		if err := copyFile(path, filepath.Join(dstSpecsDir, name), info.Mode().Perm()); err != nil {
			return err
		}
	}
	patchDir := filepath.Join(specsDir, "patches")
	if pi, err := os.Stat(patchDir); err == nil && pi.IsDir() {
		return copyRecipe(patchDir, filepath.Join(dstSpecsDir, "patches"))
	}
	return nil
}

func rpmRecipeFileInfo(path string) (os.FileInfo, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("RPM recipe input is not a regular file: %s", path)
	}
	return info, nil
}

// recipeTreeHashFor fingerprints the same files copied into the build workspace.
func recipeTreeHashFor(family, recipePath string) (string, error) {
	dir := filepath.Dir(recipePath)
	if family != "rpm" {
		return recipeTreeHash(dir)
	}
	names, err := rpmRecipeFiles(recipePath)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	for _, name := range names {
		path := filepath.Join(dir, name)
		info, err := rpmRecipeFileInfo(path)
		if err != nil {
			return "", err
		}
		f, err := fileInfo(path)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(h, "%s\x00%d\x00%s\x00", name, info.Mode(), f.SHA256)
	}
	patchDir := filepath.Join(dir, "patches")
	if pi, err := os.Stat(patchDir); err == nil && pi.IsDir() {
		ph, err := recipeTreeHash(patchDir)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(h, "patches\x00%s\x00", ph)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
