## Usage

Sources:

- [modules/vexec_pgvector/README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_pgvector/README.md)
- [modules/vexec_pgvector/vexec_pgvector--1.0.sql](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_pgvector/vexec_pgvector--1.0.sql)
- [README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/README.md)
- [modules/vexec_pgvector/vexec_pgvector.control](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_pgvector/vexec_pgvector.control)

`vexec_pgvector` is a PostgreSQL 19 companion to the `vexec` vectorized executor. It registers declarations for the underlying extension’s functions so batch execution can reuse their behavior.

### Core Workflow

```ini
shared_preload_libraries = 'vexec,vexec_pgvector'
vexec.mode = auto
```

```sql
SELECT * FROM vexec_pgvector.declarations();
SELECT * FROM vexec.declared_calls();
```

### Operational Boundaries

Preload both `vexec` and `vexec_pgvector` on every participating server and restart; preload order is unrestricted. `vexec.mode` must enable vector execution. `CREATE EXTENSION vexec_pgvector` is optional inspection support and requires the underlying `vector` extension; it does not activate hooks. `vexec.declared_calls()` shows the bound calls. A superuser creates the non-relocatable SQL objects.

The pack binds only pgvector 0.8.6 and 0.8.7. Other versions run without pack acceleration. It calls pgvector functions through fmgr and checks matching dimensions; erroring rows fall back to PostgreSQL evaluation. It is not a new vector index method.
