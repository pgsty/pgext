## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/postgrespro/pg_grab_statement/blob/31273939223f2cf55c0eb5d28a30ba229c20d501/README.md)
- [Version 1.0 SQL objects](https://github.com/postgrespro/pg_grab_statement/blob/31273939223f2cf55c0eb5d28a30ba229c20d501/pg_grab_statement--1.0.sql)
- [Extension control file](https://github.com/postgrespro/pg_grab_statement/blob/31273939223f2cf55c0eb5d28a30ba229c20d501/pg_grab_statement.control)

`pg_grab_statement` records detailed statements from successfully committed transactions. It hooks executor start and end and writes directly to the unlogged table `grab.statement_log`; the view `grab.statements` resolves user and query-type details.

```sql
CREATE EXTENSION pg_grab_statement;
LOAD 'pg_grab_statement';
SELECT transaction, query_number, username, query_source
FROM grab.statements
ORDER BY transaction, query_number;
```

A session can use `LOAD`; cluster-wide capture requires `shared_preload_libraries`, while role or database defaults may use `session_preload_libraries`. Upstream measured roughly 10–15 percent overhead in a SELECT-only benchmark, so benchmark the real workload.

Captured rows include query text, parameter values, parameter types, user IDs, and timing. Restrict access and retention because secrets and personal data can be recorded. The table is unlogged and can be truncated after a crash, so it is not an audit ledger. The project uses old PostgreSQL internals and has no current compatibility matrix; validate hooks, prepared statements, failed transactions, crash behavior, and storage growth on the exact server release.
