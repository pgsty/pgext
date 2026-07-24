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
	"net/url"
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
	VisitsFile  string        // page hit counter persistence path (empty: memory only)
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

	visits := newVisitStore(opts.VisitsFile)
	visits.startFlusher(ctx)
	a := &api{store: store, cache: newTTLCache(opts.CacheTTL, 2048), pool: opts.Pool, reloadToken: opts.ReloadToken, visits: visits}

	mux := newMux(a)

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
	mux.HandleFunc("GET /api/v1/ext/{name}/visit", a.handleVisit)
	mux.HandleFunc("POST /api/v1/ext/{name}/visit", a.handleVisit)
	mux.HandleFunc("GET /api/v1/dim/{key}", a.handleDim)
	mux.HandleFunc("GET /api/v1/bootstrap", a.handleBootstrap)
	mux.HandleFunc("POST /api/v1/reload", a.handleReload)
	mux.HandleFunc("GET /healthz", a.handleHealth)
	mux.HandleFunc("GET /assets/{file}", handleAsset)
	// Canonical deep links are /ext/{name}, /pkg/{name} and /cate/{code}; the
	// short legacy forms redirect permanently so old shared URLs keep working.
	for _, route := range []struct{ from, to string }{
		{"/e", "/ext/"}, {"/p", "/pkg/"}, {"/c", "/cate/"},
	} {
		mux.HandleFunc("GET "+route.from+"/{name}", redirectPrefix(route.to, http.StatusMovedPermanently))
		mux.HandleFunc("GET "+route.from+"/{name}/{$}", redirectPrefix(route.to, http.StatusMovedPermanently))
	}
	registerLegacyRedirects(mux)
	mux.HandleFunc("/", handleSPA)
	return mux
}

/* ---------------- legacy Hugo URL compatibility ---------------- */

// pigDocsURL is where the retired /pig/* CLI handbook lives now.
const pigDocsURL = "https://pigsty.io/docs/pig"

// legacyRedirects maps retired Hugo-era pages onto their dynamic equivalents.
// Each entry is registered with and without a trailing slash and answers with
// a 302 that carries the query string over. Prefix families (/zh, /en, /pig
// and the taxonomy term pages) are wired in registerLegacyRedirects.
var legacyRedirects = map[string]string{
	"/e":          "/",              // full extension table → catalog home
	"/list/ext":   "/",              // ditto (lang/license are native dims already)
	"/list/pkg":   "/list/package",  // package family index
	"/list/cate":  "/list/category", // category index
	"/os":         "/list/os",       // Linux target index; /os/{target} is a live SPA route
	"/os/matrix":  "/matrix",        // global build matrix
	"/categories": "/list/category", // Hugo taxonomy index
	"/tags":       "/list/tag",      // Hugo taxonomy index
	"/repo":       "/list/repo",     // repo doc index; /repo/{value} is a live SPA route
	"/repo/pgsql": "/repo/PIGSTY",   // old Pigsty PGSQL repo doc → its repo page
}

func registerLegacyRedirects(mux *http.ServeMux) {
	for from, to := range legacyRedirects {
		mux.HandleFunc("GET "+from, redirectTo(to))
		mux.HandleFunc("GET "+from+"/{$}", redirectTo(to))
	}
	// The /zh mirror (and a defensive /en) drops its language prefix; the SPA
	// switches UI language via ?lang= instead.
	for _, lang := range []string{"/zh", "/en"} {
		mux.HandleFunc("GET "+lang, redirectStripPrefix(lang))
		mux.HandleFunc("GET "+lang+"/", redirectStripPrefix(lang))
	}
	// The whole retired pig CLI handbook moved to the Pigsty documentation.
	mux.HandleFunc("GET /pig", redirectTo(pigDocsURL))
	mux.HandleFunc("GET /pig/", redirectTo(pigDocsURL))
	// Hugo taxonomy term pages: categories have dedicated landing pages while
	// tag terms map onto the catalog tag filter.
	mux.HandleFunc("GET /categories/{name}", redirectPrefix("/cate/", http.StatusFound))
	mux.HandleFunc("GET /categories/{name}/{$}", redirectPrefix("/cate/", http.StatusFound))
	mux.HandleFunc("GET /tags/{name}", redirectTagTerm)
	mux.HandleFunc("GET /tags/{name}/{$}", redirectTagTerm)
}

// redirectPrefix maps a one-segment legacy route onto its canonical prefix,
// preserving the query string.
func redirectPrefix(prefix string, code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target := prefix + url.PathEscape(r.PathValue("name"))
		if r.URL.RawQuery != "" {
			target += "?" + r.URL.RawQuery
		}
		http.Redirect(w, r, target, code)
	}
}

// redirectTo answers a retired page with a temporary redirect onto its fixed
// replacement, carrying the query string over.
func redirectTo(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest := target
		if q := r.URL.RawQuery; q != "" {
			if strings.Contains(dest, "?") {
				dest += "&" + q
			} else {
				dest += "?" + q
			}
		}
		http.Redirect(w, r, dest, http.StatusFound)
	}
}

// redirectStripPrefix drops a legacy language prefix and redirects to the same
// path on the default site, keeping the query string. The target is always a
// local absolute path (never protocol-relative), and stripping one prefix per
// hop cannot loop.
func redirectStripPrefix(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target := strings.TrimPrefix(r.URL.EscapedPath(), prefix)
		for strings.HasPrefix(target, "//") {
			target = target[1:]
		}
		if !strings.HasPrefix(target, "/") {
			target = "/" + target
		}
		if q := r.URL.RawQuery; q != "" {
			target += "?" + q
		}
		http.Redirect(w, r, target, http.StatusFound)
	}
}

// redirectTagTerm maps a Hugo /tags/{term} taxonomy page onto the catalog tag
// filter, merging any extra query parameters behind it.
func redirectTagTerm(w http.ResponseWriter, r *http.Request) {
	target := "/?tag=" + url.QueryEscape(r.PathValue("name"))
	if r.URL.RawQuery != "" {
		target += "&" + r.URL.RawQuery
	}
	http.Redirect(w, r, target, http.StatusFound)
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
