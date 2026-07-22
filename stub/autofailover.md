## Usage

Sources:

- [Version 1.0 control file](https://github.com/gfphoenix78/autofailover/blob/6494a6ed6b8b882f287331996b2ab075e334afa0/autofailover.control)
- [Version 1.0 SQL objects](https://github.com/gfphoenix78/autofailover/blob/6494a6ed6b8b882f287331996b2ab075e334afa0/autofailover--1.0.sql)
- [Command implementation](https://github.com/gfphoenix78/autofailover/blob/6494a6ed6b8b882f287331996b2ab075e334afa0/autofailover_funcs.c)

`autofailover` is an experimental C extension that exposes PostgreSQL role/WAL state and can directly promote a standby or change synchronous-replication settings. It uses private server internals, so it must be built and tested against the exact PostgreSQL release; it is not a complete failover controller.

### Inspect State

```sql
CREATE EXTENSION autofailover;

SELECT *
FROM autofailover_execute('status', NULL);
```

`autofailover_execute(cmd, last_role)` returns `role`, `syncrep`, `sync_state`, `lsn`, and `walconn`. The implementation recognizes `status`, `info`, `promote`, `syncrep`, and `unsyncrep`; `last_role` is accepted but not used for validation.

### Operational Boundaries

`promote` writes PostgreSQL's promote signal and signals the postmaster. `syncrep` and `unsyncrep` execute `ALTER SYSTEM`, set `synchronous_standby_names` to `*` or an empty value, and reload the configuration. These actions affect the entire cluster and do not provide fencing, quorum, leader election, split-brain prevention, or rollback.

Both installed functions are `SECURITY DEFINER`, and `test_udf()` only emits synthetic debug rows. Revoke access from untrusted roles and do not expose either function through application connections:

```sql
REVOKE EXECUTE ON FUNCTION autofailover_execute(text, text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION test_udf() FROM PUBLIC;
```

Use the status operation only as a low-level probe. Before considering any mutating command, validate the exact PostgreSQL ABI in an isolated cluster and supply external health checks, fencing, durable orchestration, and recovery procedures.
