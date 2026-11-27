package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	githubDefaultREST    = "https://api.github.com"
	githubDefaultGraphQL = "https://api.github.com/graphql"
	githubUserAgent      = "pgext-github-sync"
)

const githubReleaseTagQuery = `
query RepoReleaseTagActivity($owner: String!, $name: String!) {
  repository(owner: $owner, name: $name, followRenames: true) {
    nameWithOwner
    isArchived
    releases(first: 1, orderBy: {field: CREATED_AT, direction: DESC}) {
      nodes {
        tagName
        publishedAt
        createdAt
        url
      }
    }
    refs(refPrefix: "refs/tags/", first: 1, orderBy: {field: TAG_COMMIT_DATE, direction: DESC}) {
      nodes {
        name
        target {
          __typename
          oid
          ... on Commit { committedDate }
          ... on Tag {
            tagger { date }
            target {
              __typename
              oid
              ... on Commit { committedDate }
            }
          }
        }
      }
    }
  }
  rateLimit {
    cost
    remaining
    resetAt
  }
}
`

// GitHubSyncOptions configures `pgext sync github`.
type GitHubSyncOptions struct {
	Extensions   []string
	URLs         []string
	Statuses     []string
	StaleDays    int
	Limit        int
	Reserve      int
	Delay        time.Duration
	Token        string
	Proxy        string
	Backfill     bool
	BackfillOnly bool
	RefreshAfter time.Time
	DryRun       bool
}

// GitHubSyncSummary reports the outcome of a GitHub metadata sync.
type GitHubSyncSummary struct {
	Targets          int
	Fetched          int
	Blocked          int
	Errors           int
	StoppedOnReserve bool
	UniverseUpdated  int64
	ExtensionUpdated int64
	CoreRemaining    int
	GraphQLRemaining int
}

type githubTarget struct {
	URLNorm        string
	Owner          string
	Repo           string
	APIURL         string
	ExtensionIDs   []int32
	ExtensionNames []string
}

// GitHubActivity holds normalized repository metadata and activity timestamps.
type GitHubActivity struct {
	URLNorm                  string
	Status                   string
	HTTPStatus               int
	RateLimitRemaining       *int
	RateLimitReset           *time.Time
	Error                    string
	ETag                     string
	Owner                    string
	Repo                     string
	APIURL                   string
	APIJSON                  json.RawMessage
	Stars                    *int
	Forks                    *int
	Watchers                 *int
	Subscribers              *int
	PushedAt                 *time.Time
	UpdatedAtAPI             *time.Time
	DefaultBranch            string
	Archived                 *bool
	LastCommitAt             *time.Time
	LastCommitSHA            string
	LastCommitHTMLURL        string
	LatestReleaseTag         string
	LatestReleasePublishedAt *time.Time
	LatestReleaseCreatedAt   *time.Time
	LastReleaseAt            *time.Time
	LatestTagName            string
	LatestTagAt              *time.Time
	LatestTagDateSource      string
	LatestTagTargetType      string
	LatestTagCommitSHA       string
	LatestTagCommitAt        *time.Time
	LastReleaseOrTagAt       *time.Time
	LastReleaseOrTagSource   string
	LastUpdateAt             *time.Time
	LastUpdateSource         string
	ActivityJSON             json.RawMessage
}

// NormalizeGitHubRepo normalizes a GitHub repository URL to (owner, repo, canonicalURL).
func NormalizeGitHubRepo(rawURL string) (owner, repo, urlNorm string, err error) {
	v := strings.TrimSpace(rawURL)
	if v == "" {
		return "", "", "", errors.New("empty GitHub repository URL")
	}
	var path string
	if strings.HasPrefix(v, "git@github.com:") {
		path = strings.TrimPrefix(v, "git@github.com:")
	} else {
		if strings.HasPrefix(v, "github.com/") || strings.HasPrefix(v, "www.github.com/") {
			v = "https://" + v
		}
		u, parseErr := url.Parse(v)
		if parseErr != nil {
			return "", "", "", fmt.Errorf("invalid GitHub URL %q: %w", rawURL, parseErr)
		}
		host := strings.ToLower(u.Hostname())
		if (host != "github.com" && host != "www.github.com") ||
			(u.Scheme != "https" && u.Scheme != "http") || u.User != nil || u.Port() != "" {
			return "", "", "", fmt.Errorf("not a github.com repository URL: %s", rawURL)
		}
		path = strings.TrimPrefix(u.Path, "/")
	}

	var parts []string
	for _, seg := range strings.Split(path, "/") {
		if seg != "" {
			if decoded, decErr := url.PathUnescape(seg); decErr == nil {
				seg = decoded
			}
			parts = append(parts, seg)
		}
	}
	if len(parts) < 2 {
		return "", "", "", fmt.Errorf("not a GitHub owner/repo URL: %s", rawURL)
	}
	owner, repo = parts[0], parts[1]
	if strings.HasSuffix(strings.ToLower(repo), ".git") {
		repo = repo[:len(repo)-4]
	}
	if owner == "" || repo == "" {
		return "", "", "", fmt.Errorf("not a GitHub owner/repo URL: %s", rawURL)
	}
	urlNorm = fmt.Sprintf("https://github.com/%s/%s", strings.ToLower(owner), strings.ToLower(repo))
	return owner, repo, urlNorm, nil
}

// DiscoverGitHubToken resolves a token from explicit flag, env vars, or `gh auth token`.
func DiscoverGitHubToken(explicit string) (string, string) {
	if t := strings.TrimSpace(explicit); t != "" {
		return t, "flag"
	}
	for _, env := range []string{"GITHUB_TOKEN", "GH_TOKEN"} {
		if t := strings.TrimSpace(os.Getenv(env)); t != "" {
			return t, env
		}
	}
	if out, err := exec.Command("gh", "auth", "token").Output(); err == nil {
		if t := strings.TrimSpace(string(out)); t != "" {
			return t, "gh"
		}
	}
	return "", "none"
}

// GitHubClient performs authenticated REST and GraphQL calls with same-origin redirect enforcement.
type GitHubClient struct {
	RestBase         string
	GraphQLURL       string
	Token            string
	Reserve          int
	CoreRemaining    int
	CoreReset        *time.Time
	GraphQLRemaining int
	GraphQLReset     *time.Time
	httpClient       *http.Client
}

// NewGitHubClient creates a client honoring optional proxy and refusing off-origin redirects.
func NewGitHubClient(token, proxyURL string, reserve int) (*GitHubClient, error) {
	return newGitHubClientWithBases(githubDefaultREST, githubDefaultGraphQL, token, proxyURL, reserve)
}

func newGitHubClientWithBases(restBase, graphqlURL, token, proxyURL string, reserve int) (*GitHubClient, error) {
	parsedBase, err := url.Parse(restBase)
	if err != nil {
		return nil, fmt.Errorf("parse GitHub REST base URL: %w", err)
	}
	allowedHost := strings.ToLower(parsedBase.Host)
	allowedScheme := strings.ToLower(parsedBase.Scheme)

	transport := http.DefaultTransport.(*http.Transport).Clone()
	if strings.TrimSpace(proxyURL) != "" {
		pu, err := url.Parse(proxyURL)
		if err != nil {
			return nil, fmt.Errorf("parse proxy URL %q: %w", proxyURL, err)
		}
		transport.Proxy = http.ProxyURL(pu)
	}

	client := &http.Client{
		Timeout:   45 * time.Second,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return errors.New("stopped after 10 redirects")
			}
			if strings.ToLower(req.URL.Scheme) != allowedScheme ||
				strings.ToLower(req.URL.Host) != allowedHost ||
				req.URL.User != nil {
				return fmt.Errorf("refusing cross-origin GitHub API redirect to %s", req.URL)
			}
			return nil
		},
	}

	return &GitHubClient{
		RestBase:         strings.TrimRight(restBase, "/"),
		GraphQLURL:       graphqlURL,
		Token:            token,
		Reserve:          reserve,
		CoreRemaining:    -1,
		GraphQLRemaining: -1,
		httpClient:       client,
	}, nil
}

func (c *GitHubClient) doJSON(ctx context.Context, method, targetURL string, body any) (int, http.Header, []byte, error) {
	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return 0, nil, nil, err
		}
	}

	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return 0, nil, nil, ctx.Err()
			case <-time.After(time.Duration(attempt) * 750 * time.Millisecond):
			}
		}
		var reader io.Reader
		if bodyBytes != nil {
			reader = bytes.NewReader(bodyBytes)
		}
		req, err := http.NewRequestWithContext(ctx, method, targetURL, reader)
		if err != nil {
			return 0, nil, nil, err
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("User-Agent", githubUserAgent)
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
		if c.Token != "" {
			req.Header.Set("Authorization", "Bearer "+c.Token)
		}
		if bodyBytes != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if strings.Contains(err.Error(), "refusing cross-origin") {
				return 0, nil, nil, err
			}
			continue
		}
		data, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			lastErr = readErr
			continue
		}
		if resp.StatusCode >= 500 && attempt < 2 {
			lastErr = fmt.Errorf("HTTP %d", resp.StatusCode)
			continue
		}
		return resp.StatusCode, resp.Header, data, nil
	}
	return 0, nil, nil, lastErr
}

// ReserveReached reports whether either REST or GraphQL rate-limit budget is below Reserve.
func (c *GitHubClient) ReserveReached() bool {
	if c.Reserve <= 0 {
		return false
	}
	return (c.CoreRemaining >= 0 && c.CoreRemaining <= c.Reserve) ||
		(c.GraphQLRemaining >= 0 && c.GraphQLRemaining <= c.Reserve)
}

func (c *GitHubClient) updateHeaderLimit(hdr http.Header, isGraphQL bool) (*int, *time.Time) {
	var remPtr *int
	var resetPtr *time.Time
	if v := hdr.Get("X-RateLimit-Remaining"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			remPtr = &n
			if isGraphQL {
				c.GraphQLRemaining = n
			} else {
				c.CoreRemaining = n
			}
		}
	}
	if v := hdr.Get("X-RateLimit-Reset"); v != "" {
		if sec, err := strconv.ParseInt(v, 10, 64); err == nil {
			t := time.Unix(sec, 0).UTC()
			resetPtr = &t
			if isGraphQL {
				c.GraphQLReset = &t
			} else {
				c.CoreReset = &t
			}
		}
	}
	return remPtr, resetPtr
}

// FetchRepoActivity queries REST repo, REST latest commit, and GraphQL release/tag activity.
func (c *GitHubClient) FetchRepoActivity(ctx context.Context, owner, repo, urlNorm string) GitHubActivity {
	act := GitHubActivity{
		URLNorm: urlNorm,
		Owner:   owner,
		Repo:    repo,
		APIURL:  fmt.Sprintf("%s/repos/%s/%s", c.RestBase, url.PathEscape(owner), url.PathEscape(repo)),
	}

	// 1. REST repository metadata
	status, hdr, body, err := c.doJSON(ctx, http.MethodGet, act.APIURL, nil)
	if err != nil {
		act.Status = "error"
		act.Error = err.Error()
		return act
	}
	act.HTTPStatus = status
	act.ETag = hdr.Get("ETag")
	act.RateLimitRemaining, act.RateLimitReset = c.updateHeaderLimit(hdr, false)

	if isGitHubRateLimited(status, hdr, body) {
		act.Status = "rate_limited"
		act.Error = fmt.Sprintf("HTTP %d rate limited", status)
		return act
	}
	if status == http.StatusNotFound || status == http.StatusGone || status == http.StatusUnavailableForLegalReasons || status == http.StatusForbidden {
		act.Status = "blocked"
		act.Error = fmt.Sprintf("HTTP %d", status)
		return act
	}
	if status != http.StatusOK {
		act.Status = "error"
		act.Error = fmt.Sprintf("HTTP %d", status)
		return act
	}

	var repoMeta struct {
		FullName         string     `json:"full_name"`
		StargazersCount  *int       `json:"stargazers_count"`
		ForksCount       *int       `json:"forks_count"`
		WatchersCount    *int       `json:"watchers_count"`
		SubscribersCount *int       `json:"subscribers_count"`
		DefaultBranch    string     `json:"default_branch"`
		Archived         bool       `json:"archived"`
		CreatedAt        *time.Time `json:"created_at"`
		PushedAt         *string    `json:"pushed_at"`
		UpdatedAt        *string    `json:"updated_at"`
	}
	if err := json.Unmarshal(body, &repoMeta); err != nil {
		act.Status = "error"
		act.Error = fmt.Sprintf("decode repo JSON: %v", err)
		return act
	}
	if repoMeta.StargazersCount == nil || repoMeta.ForksCount == nil ||
		repoMeta.WatchersCount == nil || repoMeta.SubscribersCount == nil || repoMeta.CreatedAt == nil {
		act.Status = "error"
		act.Error = "incomplete GitHub repository metadata"
		return act
	}
	if parts := strings.SplitN(repoMeta.FullName, "/", 2); len(parts) == 2 && parts[0] != "" && parts[1] != "" {
		act.Owner, act.Repo = parts[0], parts[1]
	}
	act.APIJSON = append(json.RawMessage(nil), body...)
	act.Stars = repoMeta.StargazersCount
	act.Forks = repoMeta.ForksCount
	act.Watchers = repoMeta.WatchersCount
	act.Subscribers = repoMeta.SubscribersCount
	act.DefaultBranch = repoMeta.DefaultBranch
	act.Archived = &repoMeta.Archived
	act.PushedAt = parseOptionalTime(repoMeta.PushedAt)
	act.UpdatedAtAPI = parseOptionalTime(repoMeta.UpdatedAt)

	audit := map[string]any{"collected_at": time.Now().UTC().Format(time.RFC3339)}

	// 2. REST latest commit on default branch
	commitURL := fmt.Sprintf("%s/repos/%s/%s/commits?per_page=1", c.RestBase, url.PathEscape(act.Owner), url.PathEscape(act.Repo))
	if act.DefaultBranch != "" {
		commitURL += "&sha=" + url.QueryEscape(act.DefaultBranch)
	}
	cStatus, cHdr, cBody, cErr := c.doJSON(ctx, http.MethodGet, commitURL, nil)
	if cErr != nil {
		act.Status = "error"
		act.Error = fmt.Sprintf("fetch commits: %v", cErr)
		return act
	}
	act.RateLimitRemaining, act.RateLimitReset = c.updateHeaderLimit(cHdr, false)
	if isGitHubRateLimited(cStatus, cHdr, cBody) {
		act.Status = "rate_limited"
		act.Error = fmt.Sprintf("commits HTTP %d rate limited", cStatus)
		return act
	}
	audit["commit_status"] = cStatus
	if cStatus == http.StatusOK {
		var commits []struct {
			SHA     string `json:"sha"`
			HTMLURL string `json:"html_url"`
			Commit  struct {
				Committer struct{ Date string } `json:"committer"`
				Author    struct{ Date string } `json:"author"`
			} `json:"commit"`
		}
		if err := json.Unmarshal(cBody, &commits); err == nil && len(commits) > 0 {
			act.LastCommitSHA = commits[0].SHA
			act.LastCommitHTMLURL = commits[0].HTMLURL
			dateStr := commits[0].Commit.Committer.Date
			if dateStr == "" {
				dateStr = commits[0].Commit.Author.Date
			}
			act.LastCommitAt = parseOptionalTime(&dateStr)
		}
	} else if cStatus != http.StatusConflict {
		// 409 indicates an empty repository with no commits.
		act.Status = "error"
		act.Error = fmt.Sprintf("commits HTTP %d", cStatus)
		return act
	}

	// 3. GraphQL latest release + latest tag (when authenticated)
	if c.Token != "" {
		gqlReq := map[string]any{
			"query":     githubReleaseTagQuery,
			"variables": map[string]string{"owner": act.Owner, "name": act.Repo},
		}
		gStatus, gHdr, gBody, gErr := c.doJSON(ctx, http.MethodPost, c.GraphQLURL, gqlReq)
		if gErr != nil {
			act.Status = "error"
			act.Error = fmt.Sprintf("fetch graphql: %v", gErr)
			return act
		}
		c.updateHeaderLimit(gHdr, true)
		if isGitHubRateLimited(gStatus, gHdr, gBody) {
			act.Status = "rate_limited"
			act.Error = fmt.Sprintf("graphql HTTP %d rate limited", gStatus)
			return act
		}
		audit["graphql_status"] = gStatus
		if gStatus == http.StatusOK {
			if err := c.applyGraphQLActivity(&act, gBody); err != nil {
				act.Status = "error"
				act.Error = err.Error()
				return act
			}
			audit["graphql"] = json.RawMessage(gBody)
		}
	}

	act.LastReleaseOrTagSource, act.LastReleaseOrTagAt = latestTimeSource(
		timeSource{"release", act.LastReleaseAt},
		timeSource{"tag", act.LatestTagAt},
	)
	act.LastUpdateSource, act.LastUpdateAt = latestTimeSource(
		timeSource{"commit", act.LastCommitAt},
		timeSource{"release", act.LastReleaseAt},
		timeSource{"tag", act.LatestTagAt},
	)
	act.ActivityJSON, _ = json.Marshal(audit)
	act.Status = "fetched"
	return act
}

func (c *GitHubClient) applyGraphQLActivity(act *GitHubActivity, body []byte) error {
	var resp struct {
		Data struct {
			Repository *struct {
				Releases struct {
					Nodes []struct {
						TagName     string  `json:"tagName"`
						PublishedAt *string `json:"publishedAt"`
						CreatedAt   *string `json:"createdAt"`
					} `json:"nodes"`
				} `json:"releases"`
				Refs struct {
					Nodes []struct {
						Name   string `json:"name"`
						Target *struct {
							Typename      string  `json:"__typename"`
							OID           string  `json:"oid"`
							CommittedDate *string `json:"committedDate"`
							Tagger        *struct {
								Date *string `json:"date"`
							} `json:"tagger"`
							Target *struct {
								Typename      string  `json:"__typename"`
								OID           string  `json:"oid"`
								CommittedDate *string `json:"committedDate"`
							} `json:"target"`
						} `json:"target"`
					} `json:"nodes"`
				} `json:"refs"`
			} `json:"repository"`
			RateLimit *struct {
				Remaining int     `json:"remaining"`
				ResetAt   *string `json:"resetAt"`
			} `json:"rateLimit"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("decode graphql JSON: %w", err)
	}
	if resp.Data.RateLimit != nil {
		c.GraphQLRemaining = resp.Data.RateLimit.Remaining
		if t := parseOptionalTime(resp.Data.RateLimit.ResetAt); t != nil {
			c.GraphQLReset = t
		}
	}
	if len(resp.Errors) > 0 && resp.Data.Repository == nil {
		return fmt.Errorf("graphql error: %s", resp.Errors[0].Message)
	}
	repo := resp.Data.Repository
	if repo == nil {
		return nil
	}
	if len(repo.Releases.Nodes) > 0 {
		rel := repo.Releases.Nodes[0]
		act.LatestReleaseTag = rel.TagName
		act.LatestReleasePublishedAt = parseOptionalTime(rel.PublishedAt)
		act.LatestReleaseCreatedAt = parseOptionalTime(rel.CreatedAt)
		if act.LatestReleasePublishedAt != nil {
			act.LastReleaseAt = act.LatestReleasePublishedAt
		} else {
			act.LastReleaseAt = act.LatestReleaseCreatedAt
		}
	}
	if len(repo.Refs.Nodes) > 0 && repo.Refs.Nodes[0].Target != nil {
		tag := repo.Refs.Nodes[0]
		tgt := tag.Target
		act.LatestTagName = tag.Name
		act.LatestTagTargetType = tgt.Typename
		switch tgt.Typename {
		case "Tag":
			if tgt.Tagger != nil {
				act.LatestTagAt = parseOptionalTime(tgt.Tagger.Date)
				if act.LatestTagAt != nil {
					act.LatestTagDateSource = "tagger.date"
				}
			}
			if tgt.Target != nil && tgt.Target.Typename == "Commit" {
				act.LatestTagCommitSHA = tgt.Target.OID
				act.LatestTagCommitAt = parseOptionalTime(tgt.Target.CommittedDate)
				if act.LatestTagAt == nil && act.LatestTagCommitAt != nil {
					act.LatestTagAt = act.LatestTagCommitAt
					act.LatestTagDateSource = "underlying_commit.committedDate"
				}
			}
		case "Commit":
			act.LatestTagCommitSHA = tgt.OID
			act.LatestTagCommitAt = parseOptionalTime(tgt.CommittedDate)
			act.LatestTagAt = act.LatestTagCommitAt
			if act.LatestTagAt != nil {
				act.LatestTagDateSource = "commit.committedDate"
			}
		}
	}
	return nil
}

// SyncGitHub refreshes pgext.gh_repo for Universe GitHub repositories and optionally backfills Universe & Extension.
func SyncGitHub(ctx context.Context, pool *pgxpool.Pool, opts GitHubSyncOptions) (*GitHubSyncSummary, error) {
	summary := &GitHubSyncSummary{CoreRemaining: -1, GraphQLRemaining: -1}

	if opts.BackfillOnly {
		if opts.RefreshAfter.IsZero() {
			return summary, fmt.Errorf("--backfill-only requires --refresh-after")
		}
		if opts.Backfill || len(opts.Extensions)+len(opts.URLs)+len(opts.Statuses) != 0 || opts.StaleDays != 0 || opts.Limit != 0 {
			return summary, fmt.Errorf("--backfill-only cannot be combined with fetch selectors or --backfill")
		}
		if opts.DryRun {
			return summary, nil
		}
		uCnt, eCnt, err := backfillGitHubUniverse(ctx, pool, nil, opts.RefreshAfter)
		if err != nil {
			return summary, err
		}
		summary.UniverseUpdated, summary.ExtensionUpdated = uCnt, eCnt
		return summary, nil
	}
	if !opts.RefreshAfter.IsZero() {
		return summary, fmt.Errorf("--refresh-after requires --backfill-only")
	}

	targets, err := loadGitHubTargets(ctx, pool, opts)
	if err != nil {
		return nil, err
	}
	summary.Targets = len(targets)
	if opts.DryRun || len(targets) == 0 {
		return summary, nil
	}

	var refreshAfter time.Time
	if err := pool.QueryRow(ctx, "SELECT clock_timestamp()").Scan(&refreshAfter); err != nil {
		return summary, fmt.Errorf("get GitHub refresh start time: %w", err)
	}
	token, source := DiscoverGitHubToken(opts.Token)
	client, err := NewGitHubClient(token, opts.Proxy, opts.Reserve)
	if err != nil {
		return summary, err
	}
	if err := syncGitHubMappings(ctx, pool, targets); err != nil {
		return summary, err
	}
	logrus.Infof("GitHub sync starting: %d targets (token=%s, reserve=%d, refresh_after=%s)",
		len(targets), source, opts.Reserve, refreshAfter.UTC().Format(time.RFC3339Nano))

	var fetchedURLs []string
	for i, tgt := range targets {
		if client.ReserveReached() {
			summary.StoppedOnReserve = true
			logrus.Warnf("stopping early: GitHub rate-limit reserve reached (core=%d, graphql=%d)",
				client.CoreRemaining, client.GraphQLRemaining)
			break
		}
		if i > 0 && opts.Delay > 0 {
			select {
			case <-ctx.Done():
				return summary, ctx.Err()
			case <-time.After(opts.Delay):
			}
		}
		act := client.FetchRepoActivity(ctx, tgt.Owner, tgt.Repo, tgt.URLNorm)
		if err := upsertGitHubActivity(ctx, pool, act); err != nil {
			return summary, err
		}
		switch act.Status {
		case "fetched":
			summary.Fetched++
			fetchedURLs = append(fetchedURLs, tgt.URLNorm)
		case "blocked":
			summary.Blocked++
		case "rate_limited":
			summary.StoppedOnReserve = true
			summary.CoreRemaining, summary.GraphQLRemaining = client.CoreRemaining, client.GraphQLRemaining
			return summary, nil
		default:
			summary.Errors++
		}
	}

	summary.CoreRemaining, summary.GraphQLRemaining = client.CoreRemaining, client.GraphQLRemaining
	if opts.Backfill && !summary.StoppedOnReserve && len(fetchedURLs) > 0 {
		uCnt, eCnt, err := backfillGitHubUniverse(ctx, pool, fetchedURLs, refreshAfter)
		if err != nil {
			return summary, err
		}
		summary.UniverseUpdated, summary.ExtensionUpdated = uCnt, eCnt
	}
	return summary, nil
}

func loadGitHubTargets(ctx context.Context, pool *pgxpool.Pool, opts GitHubSyncOptions) ([]githubTarget, error) {
	extFilter := toLowerSet(opts.Extensions)
	statusFilter := toLowerSet(opts.Statuses)
	if len(opts.Extensions) > 0 && len(extFilter) == 0 {
		return nil, fmt.Errorf("--ext requires an extension name")
	}
	if len(opts.Statuses) > 0 && len(statusFilter) == 0 {
		return nil, fmt.Errorf("--status requires a status")
	}
	urlFilter := make(map[string]struct{}, len(opts.URLs))
	for _, raw := range opts.URLs {
		_, _, norm, err := NormalizeGitHubRepo(raw)
		if err != nil {
			return nil, fmt.Errorf("--url: %w", err)
		}
		urlFilter[norm] = struct{}{}
	}

	rows, err := pool.Query(ctx, `
		WITH live AS (
			SELECT pgext.repo_url_norm(repo_url) AS url_norm,
			       array_agg(id ORDER BY id) AS extension_ids,
			       array_agg(name ORDER BY name COLLATE "C") AS extension_names
			FROM pgext.universe
			WHERE repo_url ~* '^https?://(www\.)?github\.com/'
			  AND pgext.repo_url_norm(repo_url) IS NOT NULL
			GROUP BY 1
		)
		SELECT l.url_norm, l.extension_ids, l.extension_names,
		       coalesce(g.status, 'missing') AS status,
		       g.updated_at
		FROM live l
		LEFT JOIN pgext.gh_repo g USING (url_norm)
		ORDER BY l.url_norm`)
	if err != nil {
		return nil, fmt.Errorf("load GitHub targets: %w", err)
	}
	defer rows.Close()

	staleCutoff := time.Now().UTC().AddDate(0, 0, -opts.StaleDays)

	var targets []githubTarget
	for rows.Next() {
		var urlNorm, status string
		var ids []int32
		var names []string
		var updatedAt *time.Time
		if err := rows.Scan(&urlNorm, &ids, &names, &status, &updatedAt); err != nil {
			return nil, err
		}
		owner, repo, norm, err := NormalizeGitHubRepo(urlNorm)
		if err != nil {
			continue
		}
		if len(urlFilter) > 0 {
			if _, ok := urlFilter[norm]; !ok {
				continue
			}
		}
		if len(extFilter) > 0 {
			matched := false
			for _, n := range names {
				if _, ok := extFilter[strings.ToLower(n)]; ok {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}
		if len(statusFilter) > 0 {
			if _, ok := statusFilter[strings.ToLower(status)]; !ok {
				continue
			}
		}
		if opts.StaleDays > 0 && status == "fetched" && updatedAt != nil && updatedAt.After(staleCutoff) {
			continue
		}
		targets = append(targets, githubTarget{
			URLNorm:        norm,
			Owner:          owner,
			Repo:           repo,
			APIURL:         fmt.Sprintf("%s/repos/%s/%s", githubDefaultREST, strings.ToLower(owner), strings.ToLower(repo)),
			ExtensionIDs:   ids,
			ExtensionNames: names,
		})
		if opts.Limit > 0 && len(targets) >= opts.Limit {
			break
		}
	}
	return targets, rows.Err()
}

func syncGitHubMappings(ctx context.Context, pool *pgxpool.Pool, targets []githubTarget) error {
	for _, t := range targets {
		_, err := pool.Exec(ctx, `
			INSERT INTO pgext.gh_repo (
				url_norm, repo_owner, repo_name, api_url,
				extension_ids, extension_names, extension_count
			) VALUES ($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT (url_norm) DO UPDATE SET
				extension_ids   = EXCLUDED.extension_ids,
				extension_names = EXCLUDED.extension_names,
				extension_count = EXCLUDED.extension_count`,
			t.URLNorm, t.Owner, t.Repo, t.APIURL, t.ExtensionIDs, t.ExtensionNames, len(t.ExtensionIDs))
		if err != nil {
			return fmt.Errorf("upsert gh_repo mapping %s: %w", t.URLNorm, err)
		}
	}
	return nil
}

func upsertGitHubActivity(ctx context.Context, pool *pgxpool.Pool, a GitHubActivity) error {
	_, err := pool.Exec(ctx, `
		UPDATE pgext.gh_repo SET
			status                      = $2,
			attempts                    = attempts + 1,
			http_status                 = $3,
			rate_limit_remaining        = $4,
			rate_limit_reset            = $5,
			error                       = NULLIF($6, ''),
			etag                        = COALESCE(NULLIF($7, ''), etag),
			repo_owner                  = $8,
			repo_name                   = $9,
			api_url                     = $10,
			api_json                    = COALESCE($11::jsonb, api_json),
			stargazers_count            = COALESCE($12, stargazers_count),
			forks_count                 = COALESCE($13, forks_count),
			watchers_count              = COALESCE($14, watchers_count),
			subscribers_count           = COALESCE($15, subscribers_count),
			pushed_at                   = COALESCE($16, pushed_at),
			updated_at_api              = COALESCE($17, updated_at_api),
			default_branch              = COALESCE(NULLIF($18, ''), default_branch),
			archived                    = COALESCE($19, archived),
			last_commit_at              = COALESCE($20, last_commit_at),
			last_commit_sha             = COALESCE(NULLIF($21, ''), last_commit_sha),
			last_commit_html_url        = COALESCE(NULLIF($22, ''), last_commit_html_url),
			latest_release_tag          = COALESCE(NULLIF($23, ''), latest_release_tag),
			latest_release_published_at = COALESCE($24, latest_release_published_at),
			latest_release_created_at   = COALESCE($25, latest_release_created_at),
			last_release_at             = COALESCE($26, last_release_at),
			latest_tag_name             = COALESCE(NULLIF($27, ''), latest_tag_name),
			latest_tag_at               = COALESCE($28, latest_tag_at),
			latest_tag_date_source      = COALESCE(NULLIF($29, ''), latest_tag_date_source),
			latest_tag_target_type      = COALESCE(NULLIF($30, ''), latest_tag_target_type),
			latest_tag_commit_sha       = COALESCE(NULLIF($31, ''), latest_tag_commit_sha),
			latest_tag_commit_at        = COALESCE($32, latest_tag_commit_at),
			last_release_or_tag_at      = COALESCE($33, last_release_or_tag_at),
			last_release_or_tag_source  = COALESCE(NULLIF($34, ''), last_release_or_tag_source),
			last_update_at              = COALESCE($35, last_update_at),
			last_update_source          = COALESCE(NULLIF($36, ''), last_update_source),
			activity_json               = COALESCE($37::jsonb, activity_json),
			fetched_at                  = CASE WHEN $2 = 'fetched' THEN now() ELSE fetched_at END,
			updated_at                  = now()
		WHERE url_norm = $1`,
		a.URLNorm, a.Status, a.HTTPStatus, a.RateLimitRemaining, a.RateLimitReset,
		a.Error, a.ETag, a.Owner, a.Repo, a.APIURL, nullableJSON(a.APIJSON),
		a.Stars, a.Forks, a.Watchers, a.Subscribers, a.PushedAt, a.UpdatedAtAPI,
		a.DefaultBranch, a.Archived, a.LastCommitAt, a.LastCommitSHA, a.LastCommitHTMLURL,
		a.LatestReleaseTag, a.LatestReleasePublishedAt, a.LatestReleaseCreatedAt, a.LastReleaseAt,
		a.LatestTagName, a.LatestTagAt, a.LatestTagDateSource, a.LatestTagTargetType,
		a.LatestTagCommitSHA, a.LatestTagCommitAt, a.LastReleaseOrTagAt, a.LastReleaseOrTagSource,
		a.LastUpdateAt, a.LastUpdateSource, nullableJSON(a.ActivityJSON))
	if err != nil {
		return fmt.Errorf("update pgext.gh_repo %s: %w", a.URLNorm, err)
	}
	return nil
}

// A nil URL set denotes an explicit cache-only backfill. A non-nil set scopes
// both catalog updates to repositories fetched by the current run.
func backfillGitHubUniverse(ctx context.Context, pool *pgxpool.Pool, urls []string, refreshAfter time.Time) (universeUpdated, extensionUpdated int64, err error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return 0, 0, err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `
		SET LOCAL TIME ZONE 'UTC';
		LOCK TABLE pgext.universe IN SHARE ROW EXCLUSIVE MODE;
		LOCK TABLE pgext.extension IN SHARE ROW EXCLUSIVE MODE;
		LOCK TABLE pgext.gh_repo IN SHARE MODE;`); err != nil {
		return 0, 0, err
	}

	err = tx.QueryRow(ctx, `
		WITH src AS (
			SELECT u.id, u.name,
			       g.stargazers_count AS stars,
			       g.subscribers_count AS watchers,
			       g.forks_count AS forks,
			       ((g.api_json->>'created_at')::timestamptz AT TIME ZONE 'UTC')::date AS repo_created_at,
			       coalesce((g.last_commit_at AT TIME ZONE 'UTC')::date, u.last_commit) AS last_commit,
			       greatest(u.last_release, (g.last_release_or_tag_at AT TIME ZONE 'UTC')::date) AS last_release,
			       greatest(u.last_active, (g.last_update_at AT TIME ZONE 'UTC')::date) AS last_active,
			       (g.fetched_at AT TIME ZONE 'UTC')::date AS checked_at
			FROM pgext.universe u
			JOIN pgext.gh_repo g
			  ON u.repo_url ~* '^https?://(www\.)?github\.com/'
			 AND pgext.repo_url_norm(u.repo_url) = g.url_norm
			WHERE g.status = 'fetched'
			  AND g.fetched_at >= $1
			  AND ($2::text[] IS NULL OR g.url_norm = ANY($2))
			  AND (u.checked_at IS NULL OR (g.fetched_at AT TIME ZONE 'UTC')::date >= u.checked_at)
			  AND g.activity_json IS NOT NULL
			  AND nullif(g.api_json->>'created_at', '') IS NOT NULL
			  AND g.api_json->'stargazers_count' = to_jsonb(g.stargazers_count)
			  AND g.api_json->'subscribers_count' = to_jsonb(g.subscribers_count)
			  AND g.api_json->'forks_count' = to_jsonb(g.forks_count)
			  AND g.api_json->'watchers_count' = to_jsonb(g.watchers_count)
		), updated_universe AS (
			UPDATE pgext.universe u
			SET stars           = s.stars,
			    watchers        = s.watchers,
			    forks           = s.forks,
			    repo_created_at = s.repo_created_at,
			    last_commit     = s.last_commit,
			    last_release    = s.last_release,
			    last_active     = s.last_active,
			    checked_at      = s.checked_at,
			    mtime           = CURRENT_DATE
			FROM src s
			WHERE u.id = s.id
			  AND (u.stars, u.watchers, u.forks, u.repo_created_at,
			       u.last_commit, u.last_release, u.last_active, u.checked_at)
			      IS DISTINCT FROM
			      (s.stars, s.watchers, s.forks, s.repo_created_at,
			       s.last_commit, s.last_release, s.last_active, s.checked_at)
			RETURNING 1
		), updated_extension AS (
			UPDATE pgext.extension e
			SET extra = jsonb_set(coalesce(e.extra, '{}'::jsonb), '{star}', to_jsonb(s.stars), true),
			    mtime = CURRENT_DATE
			FROM src s
			WHERE s.name = e.name
			  AND s.id = e.id
			  AND e.extra->'star' IS DISTINCT FROM to_jsonb(s.stars)
			RETURNING 1
		)
		SELECT (SELECT count(*) FROM updated_universe), (SELECT count(*) FROM updated_extension)`,
		refreshAfter, urls).Scan(&universeUpdated, &extensionUpdated)
	if err != nil {
		return 0, 0, fmt.Errorf("backfill GitHub metrics: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, 0, err
	}
	return universeUpdated, extensionUpdated, nil
}

type timeSource struct {
	source string
	ts     *time.Time
}

func latestTimeSource(items ...timeSource) (string, *time.Time) {
	var best timeSource
	for _, it := range items {
		if it.ts == nil {
			continue
		}
		if best.ts == nil || it.ts.After(*best.ts) {
			best = it
		}
	}
	return best.source, best.ts
}

func parseOptionalTime(s *string) *time.Time {
	if s == nil || strings.TrimSpace(*s) == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, strings.TrimSpace(*s))
	if err != nil {
		return nil
	}
	utc := t.UTC()
	return &utc
}

func isGitHubRateLimited(status int, hdr http.Header, body []byte) bool {
	if status == http.StatusTooManyRequests {
		return true
	}
	if status == http.StatusForbidden {
		if hdr.Get("X-RateLimit-Remaining") == "0" || hdr.Get("Retry-After") != "" {
			return true
		}
		lower := strings.ToLower(string(body))
		if strings.Contains(lower, "rate limit") || strings.Contains(lower, "secondary rate") {
			return true
		}
	}
	return false
}

func nullableJSON(raw json.RawMessage) *string {
	if len(raw) == 0 {
		return nil
	}
	s := string(raw)
	return &s
}

func toLowerSet(items []string) map[string]struct{} {
	set := make(map[string]struct{}, len(items))
	for _, s := range items {
		if s = strings.ToLower(strings.TrimSpace(s)); s != "" {
			set[s] = struct{}{}
		}
	}
	return set
}
