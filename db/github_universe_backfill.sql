\set ON_ERROR_STOP on
\pset pager off

-- Required invocation for an accepted full refresh:
--   psql -v refresh_after='2026-08-10T08:24:00Z' \
--     -f db/github_universe_backfill.sql
-- refresh_after must be the start of the accepted full collection run.
\if :{?refresh_after}
\else
\echo 'ERROR: refresh_after psql variable is required'
\quit 2
\endif

BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;
LOCK TABLE pgext.universe IN SHARE ROW EXCLUSIVE MODE;
LOCK TABLE pgext.extension IN SHARE ROW EXCLUSIVE MODE;
LOCK TABLE pgext.gh_repo IN SHARE MODE;

CREATE TEMP TABLE q_github_refresh_parameter ON COMMIT DROP AS
SELECT :'refresh_after'::timestamptz AS refresh_after;

CREATE TEMP TABLE q_live_github_repo ON COMMIT DROP AS
WITH normalized AS (
    SELECT id, name, pgext.repo_url_norm(repo_url) AS url_norm
    FROM pgext.universe
    WHERE repo_url ~* '^https?://(www\.)?github\.com/'
      AND pgext.repo_url_norm(repo_url) IS NOT NULL
)
SELECT url_norm,
       array_agg(id ORDER BY id) AS extension_ids,
       array_agg(name ORDER BY name COLLATE "C") AS extension_names,
       count(*)::integer AS extension_count
FROM normalized
GROUP BY url_norm;

DO $guard$
DECLARE
    backup_rows bigint;
    backup_fingerprint text;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM q_live_github_repo) THEN
        RAISE EXCEPTION 'current Universe GitHub scope is empty';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM pgext.extension e
        LEFT JOIN pgext.universe u ON u.name = e.name
        WHERE u.name IS NULL OR NOT u.packaged OR u.id <> e.id
    ) THEN
        RAISE EXCEPTION 'packaged Extension/Universe identity mapping is incomplete or inconsistent';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM q_live_github_repo l
        LEFT JOIN pgext.gh_repo g USING (url_norm)
        WHERE g.url_norm IS NULL
           OR (g.extension_ids, g.extension_names, g.extension_count)
              IS DISTINCT FROM
              (l.extension_ids, l.extension_names, l.extension_count)
    ) THEN
        RAISE EXCEPTION 'gh_repo rows are missing or mapping has drifted from Universe';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM q_live_github_repo l
        JOIN pgext.gh_repo g USING (url_norm)
        WHERE g.status NOT IN ('fetched', 'blocked')
    ) THEN
        RAISE EXCEPTION 'current GitHub scope still contains non-terminal rows';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM q_live_github_repo l
        JOIN pgext.gh_repo g USING (url_norm)
        CROSS JOIN q_github_refresh_parameter p
        WHERE g.updated_at < p.refresh_after
    ) THEN
        RAISE EXCEPTION 'current GitHub scope contains rows older than refresh_after';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM q_live_github_repo l
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
          )
    ) THEN
        RAISE EXCEPTION 'fetched GitHub rows have missing or raw-mismatched repository fields';
    END IF;

    IF to_regclass('pgext.gh_repo_20260507') IS NOT NULL THEN
        EXECUTE $sql$
            SELECT count(*),
                   md5(string_agg(md5(to_jsonb(b)::text), '' ORDER BY url_norm))
            FROM pgext.gh_repo_20260507 b
        $sql$ INTO backup_rows, backup_fingerprint;
        IF backup_rows <> 1342
           OR backup_fingerprint <> '9354ad62777856b2d60dc7d68f110e2b' THEN
            RAISE EXCEPTION 'protected gh_repo backup changed: rows %, fingerprint %',
                            backup_rows, backup_fingerprint;
        END IF;
    END IF;
END
$guard$;

-- The audited legacy fingerprint was produced in the database's normal
-- session timezone. Switch to UTC only after that byte-level check; all source
-- timestamp-to-date conversions below are also explicitly UTC.
SET LOCAL TIME ZONE 'UTC';

CREATE TEMP TABLE q_universe_github_source ON COMMIT DROP AS
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
       (g.fetched_at AT TIME ZONE 'UTC')::date AS checked_at
FROM pgext.universe u
JOIN pgext.gh_repo g
  ON u.repo_url ~* '^https?://(www\.)?github\.com/'
 AND pgext.repo_url_norm(u.repo_url) = g.url_norm
WHERE g.status = 'fetched';

CREATE TEMP TABLE q_universe_github_delta ON COMMIT DROP AS
SELECT s.id,
       u.stars IS DISTINCT FROM s.stars AS stars_changed,
       u.watchers IS DISTINCT FROM s.watchers AS watchers_changed,
       u.forks IS DISTINCT FROM s.forks AS forks_changed,
       u.repo_created_at IS DISTINCT FROM s.repo_created_at AS created_changed,
       u.last_commit IS DISTINCT FROM s.last_commit AS commit_changed,
       u.last_release IS DISTINCT FROM s.last_release AS release_changed,
       u.last_active IS DISTINCT FROM s.last_active AS active_changed,
       u.checked_at IS DISTINCT FROM s.checked_at AS checked_changed
FROM q_universe_github_source s
JOIN pgext.universe u USING (id);

CREATE TEMP TABLE q_universe_github_changed(
    id integer PRIMARY KEY
) ON COMMIT DROP;

WITH changed AS (
    UPDATE pgext.universe u
    SET stars = s.stars,
        watchers = s.watchers,
        forks = s.forks,
        repo_created_at = s.repo_created_at,
        last_commit = s.last_commit,
        last_release = s.last_release,
        last_active = s.last_active,
        checked_at = s.checked_at,
        mtime = CURRENT_DATE
    FROM q_universe_github_source s
    WHERE u.id = s.id
      AND (u.stars, u.watchers, u.forks, u.repo_created_at,
           u.last_commit, u.last_release, u.last_active, u.checked_at)
          IS DISTINCT FROM
          (s.stars, s.watchers, s.forks, s.repo_created_at,
           s.last_commit, s.last_release, s.last_active, s.checked_at)
    RETURNING u.id
)
INSERT INTO q_universe_github_changed(id)
SELECT id FROM changed;

CREATE TEMP TABLE q_extension_star_changed(
    id integer PRIMARY KEY
) ON COMMIT DROP;

WITH changed AS (
    UPDATE pgext.extension e
    SET extra = jsonb_set(
            coalesce(e.extra, '{}'::jsonb),
            '{star}',
            to_jsonb(u.stars),
            true
        ),
        mtime = CURRENT_DATE
    FROM pgext.universe u
    WHERE u.name = e.name
      AND u.id = e.id
      AND u.packaged
      AND u.stars IS NOT NULL
      AND e.extra->'star' IS DISTINCT FROM to_jsonb(u.stars)
    RETURNING e.id
)
INSERT INTO q_extension_star_changed(id)
SELECT id FROM changed;

DO $postflight$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM q_universe_github_source s
        JOIN pgext.universe u USING (id)
        WHERE (u.stars, u.watchers, u.forks, u.repo_created_at,
               u.last_commit, u.last_release, u.last_active, u.checked_at)
              IS DISTINCT FROM
              (s.stars, s.watchers, s.forks, s.repo_created_at,
               s.last_commit, s.last_release, s.last_active, s.checked_at)
    ) THEN
        RAISE EXCEPTION 'Universe postflight mismatch after GitHub backfill';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM pgext.extension e
        JOIN pgext.universe u ON u.name = e.name AND u.id = e.id
        WHERE u.stars IS NOT NULL
          AND e.extra->'star' IS DISTINCT FROM to_jsonb(u.stars)
    ) THEN
        RAISE EXCEPTION 'Extension extra.star differs from packaged Universe stars after backfill';
    END IF;
END
$postflight$;

SELECT jsonb_pretty(jsonb_build_object(
    'refresh_after', (SELECT refresh_after FROM q_github_refresh_parameter),
    'repositories', (SELECT count(*) FROM q_live_github_repo),
    'github_extensions', (SELECT sum(extension_count) FROM q_live_github_repo),
    'fetched_extensions', (SELECT count(*) FROM q_universe_github_source),
    'blocked_repositories', (
        SELECT count(*) FROM q_live_github_repo l
        JOIN pgext.gh_repo g USING (url_norm) WHERE g.status = 'blocked'
    ),
    'changed_extensions', (SELECT count(*) FROM q_universe_github_changed),
    'extension_star_changed', (SELECT count(*) FROM q_extension_star_changed),
    'stars_changed', (SELECT count(*) FROM q_universe_github_delta WHERE stars_changed),
    'watchers_changed', (SELECT count(*) FROM q_universe_github_delta WHERE watchers_changed),
    'forks_changed', (SELECT count(*) FROM q_universe_github_delta WHERE forks_changed),
    'created_changed', (SELECT count(*) FROM q_universe_github_delta WHERE created_changed),
    'commit_changed', (SELECT count(*) FROM q_universe_github_delta WHERE commit_changed),
    'release_changed', (SELECT count(*) FROM q_universe_github_delta WHERE release_changed),
    'active_changed', (SELECT count(*) FROM q_universe_github_delta WHERE active_changed),
    'checked_changed', (SELECT count(*) FROM q_universe_github_delta WHERE checked_changed)
)) AS github_universe_backfill;

COMMIT;
