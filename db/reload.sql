-- update pgext.pkg with state/hide mark
SET search_path TO pgext,public;

-- 'step 1/5: trucnate pgext.pkg ...';
TRUNCATE pgext.pkg;

-- 'step 2/5: init pgext.pkg combinations ...';
-- Invalid extension/PG/OS combinations start as N/A. A real package found in
-- the repository metadata is authoritative and is promoted to AVAIL in step 5.
INSERT INTO pgext.pkg (pg,os,name,pkg,ext,state)
SELECT pg,os,CASE os.os_type
                 WHEN 'rpm' THEN replace(replace(ext.rpm_pkg, '*', ''), '$v', pg::text)
                 WHEN 'deb' THEN replace(replace(ext.deb_pkg, '*', ''), '$v', pg::text) ELSE NULL END AS name,
       ext.pkg,ext.name,
       CASE
           WHEN NOT coalesce(ext.pg_ver, ARRAY[]::text[]) @> ARRAY[pg.pg::text] THEN 'N/A'
           WHEN os.os_type = 'rpm' AND coalesce(ext.rpm_repo, '') = '' THEN 'N/A'
           WHEN os.os_type = 'deb' AND coalesce(ext.deb_repo, '') = '' THEN 'N/A'
           WHEN coalesce(ext.extra->'os_exclude', '[]'::jsonb) ? os.os THEN 'N/A'
           WHEN coalesce(ext.tags, ARRAY[]::text[]) @> ARRAY['fork'] THEN 'N/A'
           ELSE 'MISS'
       END::pgext.pkg_state AS state
FROM (SELECT DISTINCT ON (pkg) * FROM pgext.extension WHERE lead AND NOT contrib ORDER BY pkg,id) ext,
     (SELECT * FROM pgext.os WHERE active) os,
     (SELECT * FROM pgext.pg WHERE active) pg;
-- special case handling
UPDATE pgext.pkg SET name = replace(name, 'pgaudit', 'pgaudit' || (pg+2)::TEXT ) WHERE pkg = 'pgaudit' AND pg IN (14,15) AND os ~ 'el\d';
UPDATE pgext.pkg SET name = (regexp_split_to_array(name, ' '))[1] WHERE pkg = 'postgis' AND position(' ' in name) > 0;
UPDATE pgext.pkg SET name = (regexp_split_to_array(name, ' '))[1] WHERE pkg = 'pgrouting' AND position(' ' in name) > 0;

-- 'step 3/5: update pgext.pkg package count';
UPDATE pgext.pkg SET count = (SELECT COUNT(*) FROM pgext.bin b WHERE b.pg = pkg.pg AND b.os = pkg.os AND b.name = pkg.name);

-- 'step 4/5: update pgext.pkg org and version ...';
UPDATE pgext.pkg SET org = sub.org, version = sub.version, hide = sub.hide
FROM (SELECT DISTINCT ON (pg,os,name) pg,os,name,org,version,hide FROM pgext.bin b,LATERAL (SELECT org, position('pgnf' in id) > 0 AS hide FROM pgext.repository r WHERE r.id = b.repo) ORDER BY pg,os,name,ver::pgext.version USING OPERATOR (pgext.>)) sub
WHERE pkg.pg = sub.pg AND pkg.os = sub.os AND pkg.name = sub.name;

-- 'step 5/5: update pgext.pkg state ...';
UPDATE pgext.pkg SET state = 'AVAIL' WHERE count > 0;

-- conflict with other extension, hide in list
UPDATE pgext.pkg SET hide = true WHERE pkg IN ('hydra');

-- too big, non-free, heavy extensions
UPDATE pgext.pkg SET hide = true WHERE pkg IN ('plr', 'fbsql', 'informix_fdw' ,'oracle_fdw', 'db2_fdw', 'pg_utl_smtp' ,'pg_strom', 'repmgr', 'pgpool', 'pgagent', 'dbt2', 'pgml');

-- kernel-specific packages stay out of vanilla PostgreSQL package groups
UPDATE pgext.pkg SET hide = true WHERE pkg IN (SELECT distinct pkg FROM pgext.extension WHERE tags @> '{fork}');

-- 'pgext.refresh_pkg complete';
UPDATE pgext.status SET recap_time = now();
