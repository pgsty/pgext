## Usage

Sources:

- [README 0.1.12](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/README.md)
- [Control](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/pg_tci.control)
- [SQL 0.1.12](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/sql/pg_tci--0.1.12.sql)
- [Operations](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/docs/operations.md)
- [Limitations](https://github.com/hghukasyan/pg_tci/blob/pg_tci-v0.1.12/docs/limitations.md)

`pg_tci` 0.1.12 is an external preview for counting time intervals covering an instant on PostgreSQL 18. It exposes explicit functions rather than an index access method or an automatic query rewrite.

### Core Workflow

Installation requires a superuser. Intervals use a start-inclusive, end-exclusive boundary; a null end is open-ended.

```sql
CREATE EXTENSION pg_tci;
CREATE TABLE sessions (
    id bigserial PRIMARY KEY,
    started_at timestamptz NOT NULL,
    ended_at timestamptz
);
SELECT tci.create('sessions'::regclass, 'started_at', 'ended_at');
SELECT tci.count_at('sessions'::regclass, now());
SELECT * FROM tci.stats('sessions'::regclass);
```

`tci.create_filtered` and `tci.count_at_filtered` support a registered equality filter, not an arbitrary predicate. `tci.create_partitioned` and `tci.count_at_partitioned` work across leaf partitions; register again after attaching a partition.

### Maintenance and Privileges

Use `tci.health`, `tci.seal_bounded`, `tci.compact_pairwise` and `tci.maintain` to monitor delta tails, seal rows and compact runs:

```sql
SELECT * FROM tci.health('sessions'::regclass);
SELECT tci.seal_bounded('sessions'::regclass, 8192);
SELECT * FROM tci.maintain('sessions'::regclass);
```

Manual or external scheduling does not require preload. The optional worker requires adding `pg_tci` to `shared_preload_libraries`, enabling `pg_tci.enable_maintenance_worker`, and restarting the server. Set `pg_tci.maintenance_databases` to select databases.

Read APIs enforce table access. Packed counts refuse requests that would bypass row security; use `tci.count_at_safe` or `tci.count_at_filtered_safe` for policy-aware native counting. Administrative functions are not executable by PUBLIC; table owners receive grants through setup, and a DBA can use `tci.grant_admin`.

### Limits and Recovery

Only one start/end timestamp pair per table is supported. Writes add delta rows and require maintenance; high write concurrency and compaction can dominate latency. This preview is not a production GA release or a formal MVCC proof.

After schema changes or restore, invalid indexes fall back to native counting until repaired. Run the documented rebuild procedure, such as `tci.rebuild_invalid`, and check health before relying on packed results.
