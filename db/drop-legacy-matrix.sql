-- Optional maintenance for an upgraded standalone PGEXT database.
-- Check external SQL consumers and save the view definition before running.
BEGIN;
SET LOCAL lock_timeout = '3s';
DROP MATERIALIZED VIEW IF EXISTS pgext.matrix RESTRICT;
COMMIT;
