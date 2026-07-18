## Usage

Sources:

- [Pinned upstream README](https://github.com/slardiere/streaminglag/blob/8f8b9889f94a1b1685f5078f3540b470fc335e24/README.md)
- [Pinned extension control file](https://github.com/slardiere/streaminglag/blob/8f8b9889f94a1b1685f5078f3540b470fc335e24/streaminglag.control)

`streaminglag` is a pure-SQL helper that converts WAL log sequence numbers to integers and subtracts two positions through `pg_streaming_lag(text,text)` or `pg_streaming_lag(pg_lsn,pg_lsn)`. It was written for old `pg_stat_replication` column names and pre-PostgreSQL-9.4 text LSNs.

```sql
CREATE EXTENSION streaminglag;

SELECT pid,
       pg_size_pretty(pg_streaming_lag(sent_lsn, replay_lsn)) AS extension_lag,
       pg_size_pretty(pg_wal_lsn_diff(sent_lsn, replay_lsn)::bigint) AS builtin_lag
FROM pg_stat_replication;
```

Do not use the extension result for correctness-sensitive monitoring. The reviewed SQL multiplies the high LSN half by hexadecimal `ff000000`, not `2^32`, so it produces a wrong difference when the high halves differ. Modern PostgreSQL provides `pg_wal_lsn_diff(pg_lsn,pg_lsn)` and current replication views directly expose LSN positions.

The project is an abandoned 2014 version `1.0` with no release or maintenance history. Prefer the built-in function, retain this extension only for historical compatibility, and compare every result during migration before removing it.
