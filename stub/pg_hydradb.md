## Usage

Sources:

- [README.md](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/README.md)
- [REFERENCE.md](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/REFERENCE.md)
- [pg_hydradb.control](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/pg_hydradb.control)
- [sql/pg_hydradb--0.1.0.sql](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/sql/pg_hydradb--0.1.0.sql)
- [Cargo.toml](https://github.com/hydra-db/hydradb/blob/v0.4.0/crates/pg-hydradb/Cargo.toml)
- [License](https://github.com/hydra-db/hydradb/blob/v0.4.0/LICENSE)

`pg_hydradb` 0.1.0, included in HydraDB distribution 0.4.0, traverses graphs registered over ordinary PostgreSQL vertex and edge tables. Triggers maintain internal adjacency data in the same database. It supports PostgreSQL 16–18; it does not expose the separate HydraDB server's entire query surface.

### Register and Traverse

A superuser installs the non-relocatable extension into its fixed `hydradb` schema. In a fresh database, register tables with non-null integer, text or UUID keys, primary/unique constraints and edge-to-vertex foreign keys:

```sql
CREATE EXTENSION pg_hydradb;
CREATE TABLE people (id bigint PRIMARY KEY);
CREATE TABLE friendships (
    id bigint PRIMARY KEY,
    person_a bigint NOT NULL REFERENCES people(id),
    person_b bigint NOT NULL REFERENCES people(id)
);
INSERT INTO people VALUES (1), (2), (3);
INSERT INTO friendships VALUES (1, 1, 2), (2, 2, 3);
SELECT hydradb.create_graph('social',
    '[{"label":"person","table":"public.people","key":"id"}]',
    '[{"label":"knows","table":"public.friendships","key":"id", "source":"person_a","source_label":"person", "target":"person_b","target_label":"person"}]');
SELECT hydradb.reachable_vertices('social', '{"person":[1]}', '{"max_depth":3}');
SELECT hydradb.paths_between('social', '{"person":1}', '{"person":3}', '{}');
```

`hydradb.reachable_edges` returns reached edges; `hydradb.paths_from` returns paths from seeds. Options include `max_depth`, `direction`, `edge_kind`, `edge_kinds` and `paths_per_pair`. Results use registered labels and keys. Query functions check source-table access; a caller subject to RLS is refused rather than receiving a filtered graph. Use a table owner or a role that legitimately bypasses RLS.

### Maintenance and Recovery

Run `hydradb.create_graph`, `hydradb.refresh_graph` and `hydradb.drop_graph` at READ COMMITTED. Views, partitioned tables and inheritance are unsupported. Relevant destructive DDL is refused; rewrites, missing triggers, missing foreign keys or text keys over 2,000 bytes can mark a graph as needing refresh. Correct the cause, then run `hydradb.refresh_graph`. Source data and extension adjacency/configuration data must both be included in backups.

The documented upgrade path drops each graph, runs `DROP EXTENSION pg_hydradb` without CASCADE, installs matching files, recreates the extension and registers graphs again. REPEATABLE READ writers may receive SQLSTATE 40001 after a refresh and must retry.

### Optional Preload and Telemetry

Queries work without preload. Adding `pg_hydradb` to the existing shared preload list and restarting enables shared activity counters and an hourly usage-report worker. Reporting is on by default and includes a stable cluster-derived install ID, versions, rounded counts and usage/error histograms; the upstream contract excludes query text, identifiers, keys and results. `hydradb.usage_report()` shows the current database report.

Set `pg_hydradb.usage_report = off` and reload to disable reporting. Re-enabling requires a restart. `pg_hydradb.usage_report_url` and `pg_hydradb.usage_report_interval` control the endpoint and interval. Without preload, nothing is sent and activity counters are unavailable. The source is licensed under Elastic-2.0; review its deployment restrictions.
