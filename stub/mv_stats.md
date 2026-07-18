## Usage

Sources:

- [Pinned upstream README](https://github.com/asotolongo/mv_stats/blob/7654de05ef9a02509315d7465a64cba884f66abe/README.md)
- [Version 0.3.0 installation SQL](https://github.com/asotolongo/mv_stats/blob/7654de05ef9a02509315d7465a64cba884f66abe/mv_stats--0.3.0.sql)
- [Version 0.3.0 destructive upgrade script](https://github.com/asotolongo/mv_stats/blob/7654de05ef9a02509315d7465a64cba884f66abe/mv_stats--0.2.0--0.3.0.sql)
- [Pinned extension control file](https://github.com/asotolongo/mv_stats/blob/7654de05ef9a02509315d7465a64cba884f66abe/mv_stats.control)

mv_stats version 0.3.0 uses database-wide event triggers to record creation, alteration, refresh count, and refresh duration for materialized views. It exposes the collected rows through the mv_stats view and initializes an empty row for each materialized view that predates installation. Installation requires a superuser because it creates event triggers.

### Inspect and reset statistics

```sql
CREATE EXTENSION mv_stats;

CREATE MATERIALIZED VIEW public.mv_example AS
SELECT current_database() AS database_name;

REFRESH MATERIALIZED VIEW public.mv_example;

SELECT mv_name, refresh_count, refresh_mv_last,
       refresh_mv_time_last, refresh_mv_time_total
FROM mv_stats
WHERE mv_name = 'public.mv_example';

SELECT *
FROM mv_activity_reset_stats('public.mv_example');
```

Rows for pre-existing materialized views have a NULL creation timestamp. The reset function accepts one schema-qualified name or the default asterisk for every row, and clears refresh counters and durations while recording reset_last.

### Event trigger, exposure, and upgrade limits

Three event triggers run around every CREATE, ALTER, REFRESH, or DROP MATERIALIZED VIEW in the database. Refresh timing is stored in a transaction-local custom setting and the results are ordinary rows, not crash-safe cumulative server statistics. Rolled-back DDL rolls back the corresponding row changes. Concurrent refreshes update the same statistics row, so validate counts and totals under the workload rather than using them for billing or strict auditing.

The extension grants PUBLIC select access to both the view and its backing table, exposing materialized-view names and timing to every database role. Revoke those grants if object inventory or workload duration is sensitive. The event functions are SECURITY DEFINER and do not set a protected search_path. Review ownership and function privileges, keep untrusted schemas out of privileged search paths, and restrict reset and destructive helper execution.

The supplied 0.2.0-to-0.3.0 upgrade script explicitly removes and recreates the table, view, functions, and triggers, deleting all collected statistics. The upstream README tells users to copy the view before updating. Capture and verify that copy, run the exact versioned ALTER EXTENSION command in a maintenance window, and confirm all event triggers afterward.

The README documents PostgreSQL 10 and newer, but no current-major matrix or performance characterization is provided. Measure DDL overhead, long and failed refreshes, schema renames, quoted identifiers, concurrent DDL, extension relocation, backup/restore, and upgrades on the exact PostgreSQL release before production deployment.
