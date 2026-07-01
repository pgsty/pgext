


## Usage

Sources: [PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md), [pgh_consistency SQL](https://github.com/pghydro/pghydro/blob/master/pgh_consistency--6.6.sql)

`pgh_consistency` is a PgHydro companion extension for checking and cleaning hydrographic network data. It creates the `pgh_consistency` schema with quality-control tables and functions for drainage-line and drainage-area geometry, topology, and Pfafstetter consistency checks.

### Enable

Install it after PostGIS and the core `pghydro` extension:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_consistency;
```

### Inspect Check Tables

The extension stores findings in tables named after the issue being detected:

```sql
SELECT tablename
FROM pg_tables
WHERE schemaname = 'pgh_consistency'
ORDER BY tablename;
```

Examples include `pghft_drainagelineisdisconnected`, `pghft_drainagelineisnotvalid`, `pghft_drainageareaisnotvalid`, `pghft_drainageareaoverlapdrainagearea`, `pghft_pointvalencevalue2`, and `pghft_pointdivergent`.

### Main Workflows

Use the functions in `pgh_consistency` to prepare and validate the drainage network before running analytical PgHydro routines. The SQL objects include helpers to split drainage lines at multipoints, snap drainage lines to a grid, remove repeated points, create drainage-line vertex intersections, break drainage lines, and detect invalid, non-simple, disconnected, overlapping, looping, or misclassified drainage features.

```sql
SELECT nspname, proname
FROM pg_proc p
JOIN pg_namespace n ON n.oid = p.pronamespace
WHERE nspname = 'pgh_consistency'
ORDER BY proname;
```

### Notes

The extension assumes PgHydro's base drainage tables are already populated. Run consistency checks before exporting or calculating higher-level hydrological products so invalid geometries and broken network links are visible early.
