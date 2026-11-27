/*
Copyright 2018-2026 Ruohang Feng <rh@vonng.com>
*/
package server

import (
	"net/http"
	"net/url"
	"strings"
)

// pigSiteURL is where the retired /pig/* CLI handbook lives now.
const pigSiteURL = "https://pig.pgsty.com"

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
	// Canonical deep links are /ext/{name}, /pkg/{name} and /cate/{code}; the
	// short legacy forms redirect permanently so old shared URLs keep working.
	for _, route := range []struct{ from, to string }{
		{"/e", "/ext/"}, {"/p", "/pkg/"}, {"/c", "/cate/"},
	} {
		mux.HandleFunc("GET "+route.from+"/{name}", redirectPrefix(route.to, http.StatusMovedPermanently))
		mux.HandleFunc("GET "+route.from+"/{name}/{$}", redirectPrefix(route.to, http.StatusMovedPermanently))
	}
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
	// The retired pig CLI handbook now has its own bilingual site. Preserve the
	// old page suffix so bookmarks land on the corresponding new document.
	mux.HandleFunc("GET /pig", redirectTo(pigSiteURL))
	mux.HandleFunc("GET /pig/", redirectPigDocs(false))
	mux.HandleFunc("GET /zh/pig", redirectTo(pigSiteURL+"/zh"))
	mux.HandleFunc("GET /zh/pig/", redirectPigDocs(true))
	mux.HandleFunc("GET /release/pig", redirectTo(pigSiteURL+"/release/"))
	mux.HandleFunc("GET /release/pig/", redirectTo(pigSiteURL+"/release/"))
	mux.HandleFunc("GET /zh/release/pig", redirectTo(pigSiteURL+"/zh/release/"))
	mux.HandleFunc("GET /zh/release/pig/", redirectTo(pigSiteURL+"/zh/release/"))
	// Hugo taxonomy term pages: categories have dedicated landing pages while
	// tag terms map onto the catalog tag filter.
	mux.HandleFunc("GET /categories/{name}", redirectPrefix("/cate/", http.StatusFound))
	mux.HandleFunc("GET /categories/{name}/{$}", redirectPrefix("/cate/", http.StatusFound))
	mux.HandleFunc("GET /tags/{name}", redirectTagTerm)
	mux.HandleFunc("GET /tags/{name}/{$}", redirectTagTerm)
}

// withQuery carries the request's query string over to a redirect target,
// appending with '&' when the target already has one of its own.
func withQuery(target string, r *http.Request) string {
	if q := r.URL.RawQuery; q != "" {
		sep := "?"
		if strings.Contains(target, "?") {
			sep = "&"
		}
		return target + sep + q
	}
	return target
}

// redirectPrefix maps a one-segment legacy route onto its canonical prefix.
func redirectPrefix(prefix string, code int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, withQuery(prefix+url.PathEscape(r.PathValue("name")), r), code)
	}
}

// redirectTo answers a retired page with a temporary redirect onto its fixed
// replacement.
func redirectTo(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, withQuery(target, r), http.StatusFound)
	}
}

// redirectPigDocs maps the retired pgext Hugo handbook onto pig.pgsty.com.
// Old command pages used /pig/cmd/<name>; the dedicated site exposes each
// command at the root. Other handbook pages retain their original slug.
func redirectPigDocs(zh bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/pig/")
		if zh {
			path = strings.TrimPrefix(r.URL.Path, "/zh/pig/")
		}
		commandIndex := path == "cmd" || path == "cmd/"
		if !commandIndex {
			path = strings.TrimPrefix(path, "cmd/")
		}
		target := pigSiteURL
		if zh {
			target += "/zh"
		}
		if commandIndex {
			target += "/cmd"
			if strings.HasSuffix(path, "/") {
				target += "/"
			}
		} else if path != "" {
			target += "/" + path
		}
		http.Redirect(w, r, withQuery(target, r), http.StatusFound)
	}
}

// redirectStripPrefix drops a legacy language prefix and redirects to the same
// path on the default site. The target is always a local absolute path (never
// protocol-relative), and stripping one prefix per hop cannot loop.
func redirectStripPrefix(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target := strings.TrimPrefix(r.URL.EscapedPath(), prefix)
		for strings.HasPrefix(target, "//") {
			target = target[1:]
		}
		if !strings.HasPrefix(target, "/") {
			target = "/" + target
		}
		http.Redirect(w, r, withQuery(target, r), http.StatusFound)
	}
}

// redirectTagTerm maps a Hugo /tags/{term} taxonomy page onto the catalog tag
// filter, merging any extra query parameters behind it.
func redirectTagTerm(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, withQuery("/?tag="+url.QueryEscape(r.PathValue("name")), r), http.StatusFound)
}
