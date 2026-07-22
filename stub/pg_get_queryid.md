## Usage

Sources:

- [Official upstream documentation](https://github.com/pgsentinel/pg_get_queryid/blob/826bc9f71282d0047330e9d6cb4de71471299623/README.md)
- [Official extension SQL](https://github.com/pgsentinel/pg_get_queryid/blob/826bc9f71282d0047330e9d6cb4de71471299623/pg_get_queryid--1.0.sql)
- [Official implementation](https://github.com/pgsentinel/pg_get_queryid/blob/826bc9f71282d0047330e9d6cb4de71471299623/pg_get_queryid.c)

`pg_get_queryid` reports the most recent query identifier generated or used by `pg_stat_statements` for a PostgreSQL backend PID. It was designed to correlate `pg_stat_activity` with normalized statements on PostgreSQL versions that did not expose a query ID in the activity view.

### Preload and Install

Load `pg_stat_statements` before `pg_get_queryid`, restart PostgreSQL, and create both extensions in the database:

```conf
shared_preload_libraries = 'pg_stat_statements,pg_get_queryid'
pg_get_queryid.track_utility = on
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_get_queryid;
```

The order in `shared_preload_libraries` is part of the upstream requirement. Changing either preload entry requires a server restart.

### Correlate Activity

`pg_get_queryid(integer)` returns a `bigint`, or `0` when no query ID is available for that PID:

```sql
SELECT a.pid,
       a.state,
       a.query AS active_query,
       pg_get_queryid(a.pid) AS queryid,
       s.query AS normalized_query
FROM pg_stat_activity AS a
LEFT JOIN pg_stat_statements AS s
  ON s.queryid = pg_get_queryid(a.pid)
WHERE a.query_start IS NOT NULL
  AND a.pid <> pg_backend_pid();
```

`pg_get_queryid.track_utility` is a superuser-settable boolean, enabled by default, that generates identifiers for utility statements. Prepared execution is attributed to the underlying prepared statement, and nested tracking follows the `pg_stat_statements.track` policy.

### Timing and Compatibility

The first execution is not present in `pg_stat_statements` until it finishes, and an evicted or reset entry is likewise absent until the next execution completes. At very high query rates, `pg_stat_activity.query` can change before this extension's hooks update the stored identifier, so a join can momentarily pair different statements. Apply the normal `pg_stat_activity` visibility rules and do not treat a query ID as a security boundary.

Upstream targets PostgreSQL 9.6 and later, but the hook code uses old server APIs. Modern PostgreSQL exposes `pg_stat_activity.query_id`; prefer the core field where available and verify that the extension is still needed and ABI-compatible before building it.
