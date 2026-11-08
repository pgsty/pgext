package proxy

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"pgext/cli/build"
)

type Plan struct {
	API          int        `json:"api"`
	Command      string     `json:"command"`
	Boundary     string     `json:"boundary"`
	Confirmation string     `json:"confirmation"`
	Actions      []Action   `json:"actions"`
	Affects      []Resource `json:"affects"`
	Expected     string     `json:"expected"`
}
type Action struct {
	Step        int    `json:"step"`
	Description string `json:"description"`
}
type Resource struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
type limitedBuffer struct {
	bytes.Buffer
	remaining int
}

func newLimitedBuffer(n int) *limitedBuffer { return &limitedBuffer{remaining: n} }
func (b *limitedBuffer) Write(v []byte) (int, error) {
	n := len(v)
	if len(v) > b.remaining {
		v = v[:b.remaining]
	}
	_, _ = b.Buffer.Write(v)
	b.remaining -= len(v)
	return n, nil
}
func packageManager() string {
	p, err := build.DetectPlatform()
	if err != nil {
		return ""
	}
	switch p.Family {
	case "rpm":
		return "dnf"
	case "deb":
		return "apt-get"
	}
	return ""
}
func atomicWrite(p string, b []byte, m os.FileMode) error {
	if i, err := os.Lstat(p); err == nil && !i.Mode().IsRegular() {
		return fmt.Errorf("refusing non-regular destination")
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}
	return build.AtomicWrite(p, b, m)
}
func putFile(p string, b []byte) error { return atomicWrite(p, b, 0644) }
func runRoot(args []string) error {
	return (build.NativeExecutor{}).Run(context.Background(), build.Command{Args: args, Root: true}, os.Stderr)
}
