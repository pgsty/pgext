\set ON_ERROR_STOP on
\pset pager off

-- Reusable read-only analysis for the current pgext.universe GitHub scope.
-- GitHub timestamps are converted to UTC calendar dates explicitly so the
-- result does not depend on the psql session timezone.
BEGIN TRANSACTION READ ONLY;
SET LOCAL TIME ZONE 'UTC';

-- 1. Scope and collection coverage.
WITH normalized AS (
    SELECT id, name, pgext.repo_url_norm(repo_url) AS url_norm
    FROM pgext.universe
    WHERE repo_url ~* '^https?://(www\.)?github\.com/'
      AND pgext.repo_url_norm(repo_url) IS NOT NULL
), live_repo AS (
    SELECT url_norm,
           array_agg(id ORDER BY id) AS extension_ids,
           array_agg(name ORDER BY name COLLATE "C") AS extension_names,
           count(*)::integer AS extension_count
    FROM normalized
    GROUP BY url_norm
), scoped AS (
    SELECT l.*, g.status, g.api_json, g.activity_json,
           g.stargazers_count, g.subscribers_count, g.forks_count,
           g.last_commit_at, g.last_release_or_tag_at, g.last_update_at,
           g.fetched_at, g.updated_at
    FROM live_repo l
    LEFT JOIN pgext.gh_repo g USING (url_norm)
)
SELECT (SELECT count(*) FROM normalized) AS github_extensions,
       count(*) AS github_repositories,
       count(*) FILTER (WHERE status = 'fetched') AS fetched,
       count(*) FILTER (WHERE status = 'blocked') AS blocked,
       count(*) FILTER (WHERE status = 'error') AS error,
       count(*) FILTER (WHERE status = 'rate_limited') AS rate_limited,
       count(*) FILTER (WHERE status IN ('pending', 'collecting')) AS pending,
       count(*) FILTER (WHERE status IS NULL) AS missing,
       count(api_json) AS with_repo_json,
       count(activity_json) AS with_activity_json,
       count(stargazers_count) AS with_stars,
       count(subscribers_count) AS with_explicit_watchers,
       count(forks_count) AS with_forks,
       count(last_commit_at) AS with_commit,
       count(last_release_or_tag_at) AS with_release_or_tag,
       count(last_update_at) AS with_activity_date,
       min(fetched_at) FILTER (WHERE status = 'fetched') AS oldest_fetch,
       max(fetched_at) FILTER (WHERE status = 'fetched') AS newest_fetch
FROM scoped;

-- 2. Missing rows or repository-to-extension mapping drift. Expected: 0 rows.
WITH normalized AS (
    SELECT id, name, pgext.repo_url_norm(repo_url) AS url_norm
    FROM pgext.universe
    WHERE repo_url ~* '^https?://(www\.)?github\.com/'
      AND pgext.repo_url_norm(repo_url) IS NOT NULL
), live_repo AS (
    SELECT url_norm,
           array_agg(id ORDER BY id) AS extension_ids,
           array_agg(name ORDER BY name COLLATE "C") AS extension_names,
           count(*)::integer AS extension_count
    FROM normalized
    GROUP BY url_norm
)
SELECT l.url_norm,
       CASE
           WHEN g.url_norm IS NULL THEN 'missing gh_repo row'
           ELSE 'extension mapping mismatch'
       END AS anomaly,
       l.extension_ids AS expected_ids,
       g.extension_ids AS actual_ids,
       l.extension_names AS expected_names,
       g.extension_names AS actual_names
FROM live_repo l
LEFT JOIN pgext.gh_repo g USING (url_norm)
WHERE g.url_norm IS NULL
   OR (g.extension_ids, g.extension_names, g.extension_count)
      IS DISTINCT FROM
      (l.extension_ids, l.extension_names, l.extension_count)
ORDER BY l.url_norm;

-- 3. Parsed/raw inconsistencies and implausible timestamps. Expected: 0 rows.
WITH live_repo AS (
    SELECT DISTINCT pgext.repo_url_norm(repo_url) AS url_norm
    FROM pgext.universe
    WHERE repo_url ~* '^https?://(www\.)?github\.com/'
      AND pgext.repo_url_norm(repo_url) IS NOT NULL
)
SELECT g.url_norm, g.status, g.http_status,
       concat_ws('; ',
           CASE WHEN g.api_json IS NULL THEN 'missing api_json' END,
           CASE WHEN g.activity_json IS NULL THEN 'missing activity_json' END,
           CASE WHEN g.stargazers_count IS NULL THEN 'missing stars' END,
           CASE WHEN g.subscribers_count IS NULL THEN 'missing subscribers' END,
           CASE WHEN g.forks_count IS NULL THEN 'missing forks' END,
           CASE WHEN g.api_json->>'created_at' IS NULL THEN 'missing created_at' END,
           CASE WHEN g.fetched_at IS NULL THEN 'missing fetched_at' END,
           CASE WHEN g.stargazers_count IS DISTINCT FROM (g.api_json->>'stargazers_count')::integer THEN 'stars/raw mismatch' END,
           CASE WHEN g.watchers_count IS DISTINCT FROM (g.api_json->>'watchers_count')::integer THEN 'watchers/raw mismatch' END,
           CASE WHEN g.subscribers_count IS DISTINCT FROM (g.api_json->>'subscribers_count')::integer THEN 'subscribers/raw mismatch' END,
           CASE WHEN g.forks_count IS DISTINCT FROM (g.api_json->>'forks_count')::integer THEN 'forks/raw mismatch' END,
           CASE WHEN g.last_commit_at > now() + interval '1 day' THEN 'future commit' END,
           CASE WHEN g.last_release_or_tag_at > now() + interval '1 day' THEN 'future release/tag' END,
           CASE WHEN g.last_update_at > now() + interval '1 day' THEN 'future activity' END,
           CASE WHEN g.last_update_at < g.last_commit_at THEN 'activity before commit' END,
           CASE WHEN g.last_update_at < g.last_release_or_tag_at THEN 'activity before release/tag' END
       ) AS anomaly
FROM live_repo l
JOIN pgext.gh_repo g USING (url_norm)
WHERE g.status = 'fetched'
  AND (
      g.api_json IS NULL
      OR g.activity_json IS NULL
      OR g.stargazers_count IS NULL
      OR g.subscribers_count IS NULL
      OR g.forks_count IS NULL
      OR g.api_json->>'created_at' IS NULL
      OR g.fetched_at IS NULL
      OR g.stargazers_count IS DISTINCT FROM (g.api_json->>'stargazers_count')::integer
      OR g.watchers_count IS DISTINCT FROM (g.api_json->>'watchers_count')::integer
      OR g.subscribers_count IS DISTINCT FROM (g.api_json->>'subscribers_count')::integer
      OR g.forks_count IS DISTINCT FROM (g.api_json->>'forks_count')::integer
      OR g.last_commit_at > now() + interval '1 day'
      OR g.last_release_or_tag_at > now() + interval '1 day'
      OR g.last_update_at > now() + interval '1 day'
      OR g.last_update_at < g.last_commit_at
      OR g.last_update_at < g.last_release_or_tag_at
  )
ORDER BY g.url_norm;

-- 4. Proposed Universe changes. Watch means GitHub subscribers_count, not the
-- repository API's legacy watchers_count (which aliases stars).
WITH source AS (
    SELECT u.id, u.name,
           g.stargazers_count AS stars,
           g.subscribers_count AS watchers,
           g.forks_count AS forks,
           ((g.api_json->>'created_at')::timestamptz AT TIME ZONE 'UTC')::date AS repo_created_at,
           coalesce((g.last_commit_at AT TIME ZONE 'UTC')::date, u.last_commit) AS last_commit,
           CASE
               WHEN u.last_release IS NULL THEN (g.last_release_or_tag_at AT TIME ZONE 'UTC')::date
               WHEN g.last_release_or_tag_at IS NULL THEN u.last_release
               ELSE greatest(u.last_release, (g.last_release_or_tag_at AT TIME ZONE 'UTC')::date)
           END AS last_release,
           CASE
               WHEN u.last_active IS NULL THEN (g.last_update_at AT TIME ZONE 'UTC')::date
               WHEN g.last_update_at IS NULL THEN u.last_active
               ELSE greatest(u.last_active, (g.last_update_at AT TIME ZONE 'UTC')::date)
           END AS last_active,
           (g.fetched_at AT TIME ZONE 'UTC')::date AS checked_at,
           u.stars AS old_stars, u.watchers AS old_watchers, u.forks AS old_forks,
           u.repo_created_at AS old_repo_created_at,
           u.last_commit AS old_last_commit,
           u.last_release AS old_last_release,
           u.last_active AS old_last_active,
           u.checked_at AS old_checked_at
    FROM pgext.universe u
    JOIN pgext.gh_repo g
      ON u.repo_url ~* '^https?://(www\.)?github\.com/'
     AND pgext.repo_url_norm(u.repo_url) = g.url_norm
    WHERE g.status = 'fetched'
)
SELECT count(*) AS fetched_extensions,
       count(*) FILTER (WHERE (old_stars, old_watchers, old_forks,
                               old_repo_created_at, old_last_commit,
                               old_last_release, old_last_active, old_checked_at)
                              IS DISTINCT FROM
                             (stars, watchers, forks, repo_created_at,
                              last_commit, last_release, last_active, checked_at)) AS changed_extensions,
       count(*) FILTER (WHERE old_stars IS DISTINCT FROM stars) AS stars_changed,
       count(*) FILTER (WHERE old_watchers IS DISTINCT FROM watchers) AS watchers_changed,
       count(*) FILTER (WHERE old_forks IS DISTINCT FROM forks) AS forks_changed,
       count(*) FILTER (WHERE old_repo_created_at IS DISTINCT FROM repo_created_at) AS created_changed,
       count(*) FILTER (WHERE old_last_commit IS DISTINCT FROM last_commit) AS commit_changed,
       count(*) FILTER (WHERE old_last_release IS DISTINCT FROM last_release) AS release_changed,
       count(*) FILTER (WHERE old_last_active IS DISTINCT FROM last_active) AS active_changed,
       count(*) FILTER (WHERE old_checked_at IS DISTINCT FROM checked_at) AS checked_changed
FROM source;

-- 5. Packaged-catalog compatibility changes. pgext.universe is canonical;
-- pgext.extension.extra.star remains synchronized for legacy CSV consumers.
SELECT count(*) AS packaged_extensions,
       count(*) FILTER (
           WHERE u.stars IS NOT NULL
             AND e.extra->'star' IS DISTINCT FROM to_jsonb(u.stars)
       ) AS extension_star_changes,
       count(*) FILTER (
           WHERE u.stars IS NULL AND e.extra ? 'star'
       ) AS legacy_stars_preserved_without_universe_value
FROM pgext.extension e
JOIN pgext.universe u ON u.name = e.name AND u.id = e.id;

COMMIT;
