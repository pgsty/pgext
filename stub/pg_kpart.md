## Usage

Sources:

- [Official v1.0 README](https://github.com/HexaCluster/pg_kpart/blob/v1.0/README.md)
- [v1.0 control file](https://github.com/HexaCluster/pg_kpart/blob/v1.0/pg_kpart.control)

`pg_kpart` prevents accidental queries that would scan every leaf partition of a partitioned table without effective partition pruning. Its planner hook can raise, warn, or log before execution. The functional unit is the preloaded library; there are no SQL objects to create, and upstream describes `CREATE EXTENSION` only as optional catalog registration.

### Enable and Roll Out

For cluster-wide enforcement, preload the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_kpart'
```

It can also be loaded for selected sessions or databases without a server restart:

```conf
session_preload_libraries = 'pg_kpart'
```

Start in audit mode before enforcing errors:

```sql
ALTER SYSTEM SET pg_kpart.message_level = 'warning';
SELECT pg_reload_conf();
```

Once the observed queries are understood, set `pg_kpart.message_level = 'error'`.

### Scope and Behavior

```sql
-- Check only these tables and their sub-partitions.
ALTER SYSTEM SET pg_kpart.blacklisted =
    'public.measurement, public.orders';

-- Or check all partitioned tables except selected hierarchies.
ALTER SYSTEM SET pg_kpart.whitelisted = 'public.audit_log';
SELECT pg_reload_conf();
```

```sql
-- Partition key is logdate.
SELECT * FROM measurement WHERE city_id = 5;              -- violation
SELECT * FROM measurement WHERE logdate = DATE '2026-07-01'; -- pruned, allowed
SELECT * FROM measurement WHERE logdate = $1;             -- runtime pruning, allowed
```

Violations use SQLSTATE `FS001`, which applications can trap when `message_level` is `error`.

### Configuration Index and Caveats

- `pg_kpart.enabled`: master switch; default `on`.
- `pg_kpart.message_level`: `error`, `warning`, `notice`, `log`, and other PostgreSQL message levels.
- `pg_kpart.min_partitions`: minimum leaf-partition count to check; default `2`.
- `pg_kpart.check_superuser`: superusers bypass checks by default.
- `pg_kpart.blacklisted`: when nonempty, only named hierarchies are checked and `whitelisted` is ignored.
- `pg_kpart.whitelisted`: hierarchies exempt from checking when no blacklist is set.
- A predicate whose range still includes every partition is treated as a full scan and rejected, even if it mentions the partition key.
- The hook also applies to `UPDATE`, `DELETE`, and `EXPLAIN` without `ANALYZE`. It relies on PostgreSQL's planned pruning result, not textual inspection of `WHERE` clauses.
- Upstream v1.0 is tested on PostgreSQL 14 and newer.

