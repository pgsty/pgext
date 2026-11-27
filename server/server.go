/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"context"
	"crypto/sha1"
	"embed"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

//go:embed web
var webFS embed.FS

// Options configures the pgext web server.
type Options struct {
	Listen      string        // listen address, e.g. ":8432"
	CacheTTL    time.Duration // snapshot refresh interval & response cache TTL
	Pool        *pgxpool.Pool // connected database pool (pgext schema expected)
	ReloadToken string        // enables authenticated POST /api/v1/reload when non-empty
	VisitsFile  string        // retired file-backed counter path, imported into pgext.counter once
}

// Serve runs the web server until ctx is cancelled.
func Serve(ctx context.Context, opts Options) error {
	if opts.Listen == "" {
		opts.Listen = ":8432"
	}
	if opts.CacheTTL <= 0 {
		opts.CacheTTL = 5 * time.Minute
	}

	store := NewStore(opts.Pool)
	bootCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	snap, err := store.Reload(bootCtx)
	if err != nil {
		return fmt.Errorf("initial catalog load failed: %w", err)
	}
	store.StartRefresher(ctx, opts.CacheTTL)

	counter := newCounterStore(opts.Pool)
	names := make([]string, len(snap.Exts))
	for i, e := range snap.Exts {
		names[i] = e.Name
	}
	counter.prime(bootCtx, names)
	counter.importLegacy(opts.VisitsFile)
	flushCtx, stopFlusher := context.WithCancel(context.Background())
	counter.startFlusher(flushCtx, time.Minute)

	a := newAPI(store, opts.Pool, opts.CacheTTL, opts.ReloadToken, counter)
	go a.warmMatrixCache(ctx)

	srv := &http.Server{
		Addr:              opts.Listen,
		Handler:           withLogging(withSecurityHeaders(withCORS(withGzip(withRecover(newMux(a)))))),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      120 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logrus.Infof("pgext.cloud serving at http://localhost%s (api: /api/v1, health: /healthz)", displayAddr(opts.Listen))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		stopFlusher()
		counter.Wait()
		return err
	case <-ctx.Done():
		logrus.Info("shutting down...")
		shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := srv.Shutdown(shutCtx)
		stopFlusher()
		counter.Wait()
		return err
	}
}

// newMux wires every route of the pgext web server onto a fresh ServeMux.
func newMux(a *api) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/meta", a.handleMeta)
	mux.HandleFunc("GET /api/v1/ext", a.handleList)
	mux.HandleFunc("GET /api/v1/ext/{name}", a.handleExt)
	mux.HandleFunc("GET /api/v1/ext/{name}/matrix", a.handleMatrix)
	mux.HandleFunc("GET /api/v1/matrix", a.handleGlobalMatrix)
	mux.HandleFunc("GET /api/v1/ext/{name}/files", a.handleFiles)
	mux.HandleFunc("GET /api/v1/ext/{name}/doc", a.handleDoc)
	mux.HandleFunc("GET /api/v1/ext/{name}/changelog", a.handleChangelog)
	mux.HandleFunc("GET /api/v1/ext/{name}/visit", a.handleVisit)
	mux.HandleFunc("POST /api/v1/ext/{name}/visit", a.handleVisit)
	mux.HandleFunc("GET /api/v1/dim/{key}", a.handleDim)
	mux.HandleFunc("GET /api/v1/bootstrap", a.handleBootstrap)
	mux.HandleFunc("POST /api/v1/reload", a.handleReload)
	mux.HandleFunc("GET /healthz", a.handleHealth)
	mux.HandleFunc("GET /assets/{path...}", handleAsset)
	registerLegacyRedirects(mux)
	mux.HandleFunc("/", handleSPA)
	return mux
}

func displayAddr(listen string) string {
	if strings.HasPrefix(listen, ":") {
		return listen
	}
	if _, port, ok := strings.Cut(listen, ":"); ok {
		return ":" + port
	}
	return listen
}

/* ---------------- embedded assets & SPA shell ---------------- */

type embeddedAsset struct {
	data        []byte
	contentType string
	etag        string
}

var (
	assetExtTypes = map[string]string{
		".css": "text/css; charset=utf-8",
		".js":  "text/javascript; charset=utf-8",
		".ico": "image/x-icon",
	}
	esmImportRe    = regexp.MustCompile(`(from\s+['"]\./[A-Za-z0-9_-]+\.js)(['"])`)
	embeddedAssets = map[string]embeddedAsset{}
	jsBundleVer    string
	indexHTML      []byte
	indexETag      string
)

func etagOf(data []byte) string {
	sum := sha1.Sum(data)
	return `"` + hex.EncodeToString(sum[:8]) + `"`
}

// etagMatches checks RFC 9110 If-None-Match headers, handling weak validators
// (W/"...") and comma-separated ETag lists.
func etagMatches(header, etag string) bool {
	if header == "" || etag == "" {
		return false
	}
	want := strings.TrimPrefix(strings.TrimSpace(etag), "W/")
	for _, part := range strings.Split(header, ",") {
		tok := strings.TrimSpace(part)
		if tok == "*" || strings.TrimPrefix(tok, "W/") == want {
			return true
		}
	}
	return false
}

func init() {
	rawJS := map[string][]byte{}
	var jsNames []string

	_ = fs.WalkDir(webFS, "web", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		rel := strings.TrimPrefix(p, "web/")
		ctype, ok := assetExtTypes[path.Ext(rel)]
		if !ok {
			return nil
		}
		b, err := webFS.ReadFile(p)
		if err != nil {
			return nil
		}
		if path.Ext(rel) == ".js" {
			rawJS[rel] = b
			jsNames = append(jsNames, rel)
			return nil
		}
		embeddedAssets[rel] = embeddedAsset{data: b, contentType: ctype, etag: etagOf(b)}
		return nil
	})

	// Fingerprint the whole ES module graph with one combined content hash and
	// append ?v=<hash> to relative imports so browser caches never mix modules
	// from two binary deployments (while keeping strict CSP script-src 'self').
	sort.Strings(jsNames)
	h := sha1.New()
	for _, rel := range jsNames {
		h.Write([]byte(rel))
		h.Write([]byte{0})
		h.Write(rawJS[rel])
	}
	jsBundleVer = hex.EncodeToString(h.Sum(nil)[:8])
	for _, rel := range jsNames {
		rewritten := esmImportRe.ReplaceAll(rawJS[rel], []byte("${1}?v="+jsBundleVer+"${2}"))
		embeddedAssets[rel] = embeddedAsset{
			data:        rewritten,
			contentType: assetExtTypes[".js"],
			etag:        etagOf(rewritten),
		}
	}

	// Precompute the fingerprinted SPA shell and its ETag once at startup.
	if b, err := webFS.ReadFile("web/index.html"); err == nil {
		styleVer := strings.Trim(embeddedAssets["style.css"].etag, `"`)
		html := string(b)
		html = strings.ReplaceAll(html, `/assets/style.css`, `/assets/style.css?v=`+styleVer)
		html = strings.ReplaceAll(html, `/assets/js/main.js`, `/assets/js/main.js?v=`+jsBundleVer)
		indexHTML = []byte(html)
		indexETag = etagOf(indexHTML)
	}
}

func handleAsset(w http.ResponseWriter, r *http.Request) {
	asset, ok := embeddedAssets[r.PathValue("path")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("ETag", asset.etag)
	if etagMatches(r.Header.Get("If-None-Match"), asset.etag) {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	w.Header().Set("Content-Type", asset.contentType)
	w.Header().Set("Cache-Control", "public, max-age=300")
	w.WriteHeader(http.StatusOK)
	w.Write(asset.data)
}

// serveIndex keeps the shell revalidated and fingerprints its embedded asset
// URLs so a newly deployed binary never pairs with stale cached CSS or JS.
func serveIndex(w http.ResponseWriter, r *http.Request) {
	if len(indexHTML) == 0 {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("ETag", indexETag)
	w.Header().Set("Cache-Control", "no-cache")
	if etagMatches(r.Header.Get("If-None-Match"), indexETag) {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(indexHTML)
}

// handleSPA serves the app shell for any non-API GET; unknown API paths get a JSON 404.
func handleSPA(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/") {
		writeErr(w, http.StatusNotFound, "no such endpoint: %s %s", r.Method, r.URL.Path)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		writeErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	serveIndex(w, r)
}
