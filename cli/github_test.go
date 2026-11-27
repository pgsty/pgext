package cli

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNormalizeGitHubRepo(t *testing.T) {
	cases := []struct {
		raw   string
		owner string
		repo  string
		norm  string
		bad   bool
	}{
		{"https://github.com/pgvector/pgvector", "pgvector", "pgvector", "https://github.com/pgvector/pgvector", false},
		{"https://www.github.com/Timescale/TimescaleDB.git", "Timescale", "TimescaleDB", "https://github.com/timescale/timescaledb", false},
		{"github.com/PgSty/PGExt/tree/main", "PgSty", "PGExt", "https://github.com/pgsty/pgext", false},
		{"git@github.com:citusdata/pg_cron.git", "citusdata", "pg_cron", "https://github.com/citusdata/pg_cron", false},
		{"https://gitlab.com/foo/bar", "", "", "", true},
		{"ftp://github.com/foo/bar", "", "", "", true},
		{"https://github.com/onlyowner", "", "", "", true},
		{"", "", "", "", true},
	}
	for _, tc := range cases {
		o, r, n, err := NormalizeGitHubRepo(tc.raw)
		if tc.bad {
			if err == nil {
				t.Errorf("NormalizeGitHubRepo(%q) expected error", tc.raw)
			}
			continue
		}
		if err != nil || o != tc.owner || r != tc.repo || n != tc.norm {
			t.Errorf("NormalizeGitHubRepo(%q) = (%q, %q, %q, %v), want (%q, %q, %q, nil)",
				tc.raw, o, r, n, err, tc.owner, tc.repo, tc.norm)
		}
	}
}

func TestGitHubSyncRejectsInvalidURLBeforeSelection(t *testing.T) {
	for _, urls := range [][]string{
		{"https://example.invalid/owner/repo"},
		{"https://github.com/pgvector/pgvector", ""},
	} {
		_, err := SyncGitHub(context.Background(), nil, GitHubSyncOptions{URLs: urls, DryRun: true})
		if err == nil || !strings.Contains(err.Error(), "--url") {
			t.Fatalf("invalid URL filters %v: %v", urls, err)
		}
	}
	if _, err := SyncGitHub(context.Background(), nil, GitHubSyncOptions{BackfillOnly: true}); err == nil {
		t.Fatal("a cache-only backfill requires a freshness boundary")
	}
}

func TestGitHubClientRejectsIncompleteMetrics(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/pgvector/pgvector" {
			t.Errorf("unexpected activity request %s", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"created_at":"2021-04-01T00:00:00Z","stargazers_count":42,"forks_count":0,"watchers_count":42}`))
	}))
	defer srv.Close()
	client, err := newGitHubClientWithBases(srv.URL, srv.URL+"/graphql", "", "", 200)
	if err != nil {
		t.Fatal(err)
	}
	act := client.FetchRepoActivity(context.Background(), "pgvector", "pgvector", "https://github.com/pgvector/pgvector")
	if act.Status != "error" || !strings.Contains(act.Error, "incomplete") || act.Subscribers != nil {
		t.Fatalf("accepted incomplete GitHub metrics: %+v", act)
	}
}

func TestGitHubBackfillScopeAndFreshnessIntegration(t *testing.T) {
	setupDisposableIntegrationDB(t)
	ctx := context.Background()
	refreshAfter := time.Date(2026, 1, 20, 12, 0, 0, 0, time.UTC)
	fresh := refreshAfter.Add(time.Minute)
	for _, fixture := range []struct {
		name      string
		fetchedAt time.Time
		checkedAt time.Time
		complete  bool
	}{
		{"vector", fresh, refreshAfter.AddDate(0, 0, -1), true},
		{"pg_cron", fresh, refreshAfter.AddDate(0, 0, -1), true},
		{"pg_hint_plan", refreshAfter.Add(-time.Minute), refreshAfter.AddDate(0, 0, -1), true},
		{"rum", fresh, refreshAfter.AddDate(0, 0, 1), true},
		{"pgaudit", fresh, refreshAfter.AddDate(0, 0, -1), false},
	} {
		url := "https://github.com/pgext-review/" + fixture.name
		if _, err := DB.Exec(ctx, `UPDATE pgext.universe SET repo_url = $2, stars = 11, watchers = 5, checked_at = $3 WHERE name = $1`,
			fixture.name, url, fixture.checkedAt); err != nil {
			t.Fatal(err)
		}
		if _, err := DB.Exec(ctx, `UPDATE pgext.extension SET extra = jsonb_set(coalesce(extra, '{}'::jsonb), '{star}', '11') WHERE name = $1`,
			fixture.name); err != nil {
			t.Fatal(err)
		}
		var subscribers *int
		if fixture.complete {
			count := 3
			subscribers = &count
		}
		if _, err := DB.Exec(ctx, `
			INSERT INTO pgext.gh_repo (
				url_norm, repo_owner, repo_name, api_url, status, api_json, activity_json, fetched_at,
				stargazers_count, forks_count, subscribers_count, watchers_count
			) VALUES (
				$1, 'pgext-review', $2, '', 'fetched',
				jsonb_build_object('created_at', '2021-04-01T00:00:00Z', 'stargazers_count', 42, 'forks_count', 0, 'subscribers_count', 3, 'watchers_count', 42),
				'{}', $3, 42, 0, $4, 42
			)`, url, fixture.name, fixture.fetchedAt, subscribers); err != nil {
			t.Fatal(err)
		}
	}
	summary, err := SyncGitHub(ctx, DB, GitHubSyncOptions{Extensions: []string{"no-such-fixture"}, Backfill: true})
	if err != nil || summary.Targets != 0 || summary.UniverseUpdated != 0 || summary.ExtensionUpdated != 0 {
		t.Fatalf("empty fetch scope: %+v, %v", summary, err)
	}

	scoped := []string{
		"https://github.com/pgext-review/vector",
		"https://github.com/pgext-review/pg_hint_plan",
		"https://github.com/pgext-review/rum",
		"https://github.com/pgext-review/pgaudit",
	}
	universe, extensions, err := backfillGitHubUniverse(ctx, DB, scoped, refreshAfter)
	if err != nil || universe != 1 || extensions != 1 {
		t.Fatalf("scoped backfill updated %d Universe and %d Extension rows: %v", universe, extensions, err)
	}
	for _, name := range []string{"vector", "pg_cron", "pg_hint_plan", "rum", "pgaudit"} {
		var stars, watchers, packagedStars int
		if err := DB.QueryRow(ctx, `SELECT u.stars, u.watchers, (e.extra->>'star')::int
			FROM pgext.universe u JOIN pgext.extension e USING (name) WHERE name = $1`, name).Scan(&stars, &watchers, &packagedStars); err != nil {
			t.Fatal(err)
		}
		wantStars, wantWatchers := 11, 5
		if name == "vector" {
			wantStars, wantWatchers = 42, 3
		}
		if stars != wantStars || packagedStars != wantStars || watchers != wantWatchers {
			t.Fatalf("%s stats = (%d, %d, %d), want (%d, %d, %d)", name, stars, watchers, packagedStars, wantStars, wantWatchers, wantStars)
		}
	}

	summary, err = SyncGitHub(ctx, DB, GitHubSyncOptions{BackfillOnly: true, RefreshAfter: refreshAfter})
	if err != nil || summary.UniverseUpdated != 1 || summary.ExtensionUpdated != 1 {
		t.Fatalf("explicit cache-only backfill: %+v, %v", summary, err)
	}
}

func TestGitHubClientFetchActivity(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-RateLimit-Remaining", "4800")
		w.Header().Set("X-RateLimit-Reset", "1900000000")
		switch r.URL.Path {
		case "/repos/pgvector/pgvector":
			w.Header().Set("ETag", `"abc"`)
			_, _ = w.Write([]byte(`{
				"full_name": "pgvector/pgvector",
				"stargazers_count": 20000,
				"forks_count": 1000,
				"watchers_count": 20000,
				"subscribers_count": 120,
				"default_branch": "master",
				"archived": false,
				"created_at": "2021-04-01T00:00:00Z",
				"pushed_at": "2026-08-01T12:00:00Z",
				"updated_at": "2026-08-02T00:00:00Z"
			}`))
		case "/repos/pgvector/pgvector/commits":
			_, _ = w.Write([]byte(`[{
				"sha": "deadbeef",
				"html_url": "https://github.com/pgvector/pgvector/commit/deadbeef",
				"commit": {"committer": {"date": "2026-08-01T12:00:00Z"}}
			}]`))
		case "/graphql":
			_, _ = w.Write([]byte(`{
				"data": {
					"repository": {
						"releases": {"nodes": [{"tagName": "v0.8.0", "publishedAt": "2026-07-15T10:00:00Z", "createdAt": "2026-07-15T09:00:00Z"}]},
						"refs": {"nodes": [{"name": "v0.8.1", "target": {"__typename": "Tag", "oid": "111", "tagger": {"date": "2026-08-05T00:00:00Z"}, "target": {"__typename": "Commit", "oid": "222", "committedDate": "2026-08-04T00:00:00Z"}}}]}
					},
					"rateLimit": {"remaining": 4900, "resetAt": "2026-08-06T00:00:00Z"}
				}
			}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	client, err := newGitHubClientWithBases(srv.URL, srv.URL+"/graphql", "test-token", "", 200)
	if err != nil {
		t.Fatal(err)
	}
	act := client.FetchRepoActivity(context.Background(), "pgvector", "pgvector", "https://github.com/pgvector/pgvector")
	if act.Status != "fetched" {
		t.Fatalf("status = %q, err = %q", act.Status, act.Error)
	}
	if act.Stars == nil || *act.Stars != 20000 || act.Subscribers == nil || *act.Subscribers != 120 {
		t.Fatalf("unexpected counts: stars=%v subscribers=%v", act.Stars, act.Subscribers)
	}
	if act.LastCommitSHA != "deadbeef" || act.LatestReleaseTag != "v0.8.0" || act.LatestTagName != "v0.8.1" {
		t.Fatalf("unexpected activity metadata: %+v", act)
	}
	if act.LastReleaseOrTagSource != "tag" || act.LastUpdateSource != "tag" {
		t.Fatalf("sources = (%q, %q), want (tag, tag)", act.LastReleaseOrTagSource, act.LastUpdateSource)
	}
	if client.CoreRemaining != 4800 || client.GraphQLRemaining != 4900 || client.ReserveReached() {
		t.Fatalf("unexpected rate limit state: core=%d gql=%d", client.CoreRemaining, client.GraphQLRemaining)
	}

	// Verify 404 maps to blocked
	blocked := client.FetchRepoActivity(context.Background(), "pgvector", "missing", "https://github.com/pgvector/missing")
	if blocked.Status != "blocked" || blocked.HTTPStatus != 404 {
		t.Fatalf("404 mapped to status=%q http=%d, want blocked/404", blocked.Status, blocked.HTTPStatus)
	}
}

func TestGitHubClientRefusesCrossOriginRedirect(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://evil.example.com/steal", http.StatusMovedPermanently)
	}))
	defer srv.Close()

	client, err := newGitHubClientWithBases(srv.URL, srv.URL+"/graphql", "secret-token", "", 200)
	if err != nil {
		t.Fatal(err)
	}
	act := client.FetchRepoActivity(context.Background(), "a", "b", "https://github.com/a/b")
	if act.Status != "error" || !strings.Contains(act.Error, "refusing cross-origin") {
		t.Fatalf("expected cross-origin refusal, got status=%q err=%q", act.Status, act.Error)
	}

	// Ensure ReserveReached triggers when remaining <= reserve.
	client.CoreRemaining = 150
	if !client.ReserveReached() {
		t.Fatal("expected ReserveReached() == true when CoreRemaining <= Reserve")
	}
	var raw map[string]any
	_ = json.Unmarshal([]byte(`{}`), &raw)
}
