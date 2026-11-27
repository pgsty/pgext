-- Apply the 2026-11-21 build policy to an existing catalog without reloading it.
-- Run with psql -X -v ON_ERROR_STOP=1 against the intended database.
-- PG19 release metadata must already exist; do not invent it in this migration.
BEGIN;
LOCK TABLE pgext.pg, pgext.os IN SHARE ROW EXCLUSIVE MODE;
COMMENT ON COLUMN pgext.pg.active IS 'Whether this major version is an active build and availability-matrix target; historical rows remain inactive';
COMMENT ON COLUMN pgext.os.active IS 'Whether this OS/architecture is an active build and availability-matrix target; historical rows remain inactive';

DO $$
DECLARE
    changed_pg integer;
    changed_os integer;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pgext.pg WHERE pg = 19) THEN
        RAISE EXCEPTION 'PG19 metadata is missing; curate pgext.pg before applying build targets';
    END IF;

    UPDATE pgext.pg SET active = (pg BETWEEN 15 AND 19)
    WHERE pg BETWEEN 14 AND 19 AND active IS DISTINCT FROM (pg BETWEEN 15 AND 19);
    GET DIAGNOSTICS changed_pg = ROW_COUNT;
    UPDATE pgext.os SET active = true, comment = 'Near EOL'
    WHERE os_code = 'el8' AND (NOT active OR comment IS DISTINCT FROM 'Near EOL');
    GET DIAGNOSTICS changed_os = ROW_COUNT;

    IF (SELECT array_agg(pg ORDER BY pg) FROM pgext.active_pg)
        IS DISTINCT FROM ARRAY[15,16,17,18,19] THEN
        RAISE EXCEPTION 'Unexpected active PostgreSQL dimensions';
    END IF;
    IF (SELECT array_agg(os ORDER BY os) FROM pgext.active_os) IS DISTINCT FROM ARRAY[
        'd12.aarch64','d12.x86_64','d13.aarch64','d13.x86_64',
        'el10.aarch64','el10.x86_64','el8.aarch64','el8.x86_64','el9.aarch64','el9.x86_64',
        'u22.aarch64','u22.x86_64','u24.aarch64','u24.x86_64','u26.aarch64','u26.x86_64'
    ]::text[] THEN
        RAISE EXCEPTION 'Unexpected active operating-system dimensions';
    END IF;

    IF changed_pg + changed_os > 0 THEN
        -- Invalidate derived summaries without changing historical pkg/bin rows.
        UPDATE pgext.status SET recap_time = now() WHERE id = 0;
    END IF;
END $$;
COMMIT;
