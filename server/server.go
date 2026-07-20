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
	"net/http"
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
	if _, err := store.Reload(bootCtx); err != nil {
		return fmt.Errorf("initial catalog load failed: %w", err)
	}
	store.StartRefresher(ctx, opts.CacheTTL)

	a := &api{store: store, cache: newTTLCache(opts.CacheTTL, 2048), pool: opts.Pool, reloadToken: opts.ReloadToken}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/meta", a.handleMeta)
	mux.HandleFunc("GET /api/v1/ext", a.handleList)
	mux.HandleFunc("GET /api/v1/ext/{name}", a.handleExt)
	mux.HandleFunc("GET /api/v1/ext/{name}/matrix", a.handleMatrix)
	mux.HandleFunc("GET /api/v1/ext/{name}/files", a.handleFiles)
	mux.HandleFunc("GET /api/v1/ext/{name}/doc", a.handleDoc)
	mux.HandleFunc("GET /api/v1/dim/{key}", a.handleDim)
	mux.HandleFunc("GET /api/v1/bootstrap", a.handleBootstrap)
	mux.HandleFunc("POST /api/v1/reload", a.handleReload)
	mux.HandleFunc("GET /healthz", a.handleHealth)
	mux.HandleFunc("GET /assets/{file}", handleAsset)
	mux.HandleFunc("/", handleSPA)

	srv := &http.Server{
		Addr:              opts.Listen,
		Handler:           withLogging(withSecurityHeaders(withCORS(withGzip(withRecover(mux))))),
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
		return err
	case <-ctx.Done():
		logrus.Info("shutting down...")
		shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return srv.Shutdown(shutCtx)
	}
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

var assetETags = map[string]string{}

func init() {
	for _, name := range []string{"index.html", "style.css", "app.js"} {
		if b, err := webFS.ReadFile("web/" + name); err == nil {
			sum := sha1.Sum(b)
			assetETags[name] = `"` + hex.EncodeToString(sum[:8]) + `"`
		}
	}
	// The shell ETag covers its referenced assets as well as its own source.
	// This makes the fingerprint query strings advance on every UI change.
	sum := sha1.Sum([]byte(assetETags["index.html"] + assetETags["style.css"] + assetETags["app.js"]))
	assetETags["index.html"] = `"` + hex.EncodeToString(sum[:8]) + `"`
}

func serveEmbedded(w http.ResponseWriter, r *http.Request, name, ctype string) {
	b, err := webFS.ReadFile("web/" + name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	etag := assetETags[name]
	if etag != "" {
		w.Header().Set("ETag", etag)
		if r.Header.Get("If-None-Match") == etag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	w.Header().Set("Content-Type", ctype)
	w.Header().Set("Cache-Control", "public, max-age=300")
	w.WriteHeader(200)
	w.Write(b)
}

func handleAsset(w http.ResponseWriter, r *http.Request) {
	switch r.PathValue("file") {
	case "style.css":
		serveEmbedded(w, r, "style.css", "text/css; charset=utf-8")
	case "app.js":
		serveEmbedded(w, r, "app.js", "text/javascript; charset=utf-8")
	default:
		http.NotFound(w, r)
	}
}

// serveIndex keeps the shell revalidated and fingerprints its embedded asset
// URLs. A newly deployed binary therefore cannot be paired with a five-minute
// stale app.js from the previous catalog contract.
func serveIndex(w http.ResponseWriter, r *http.Request) {
	b, err := webFS.ReadFile("web/index.html")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	version := func(name string) string { return strings.Trim(assetETags[name], `"`) }
	html := string(b)
	html = strings.ReplaceAll(html, `/assets/style.css`, `/assets/style.css?v=`+version("style.css"))
	html = strings.ReplaceAll(html, `/assets/app.js`, `/assets/app.js?v=`+version("app.js"))
	etag := assetETags["index.html"]
	w.Header().Set("ETag", etag)
	w.Header().Set("Cache-Control", "no-cache")
	if r.Header.Get("If-None-Match") == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
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
