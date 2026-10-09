## Usage

Sources:

- [modules/vexec_postgis/README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_postgis/README.md)
- [modules/vexec_postgis/vexec_postgis--1.0.sql](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_postgis/vexec_postgis--1.0.sql)
- [README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/README.md)
- [modules/vexec_postgis/vexec_postgis.control](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_postgis/vexec_postgis.control)

`vexec_postgis` is a PostgreSQL 19 companion to the `vexec` vectorized executor. It registers declarations for the underlying extension’s functions so batch execution can reuse their behavior.

### Core Workflow

```ini
shared_preload_libraries = 'vexec,vexec_postgis'
vexec.mode = auto
```

```sql
SELECT * FROM vexec_postgis.declarations();
SELECT * FROM vexec.declared_calls();
```

### Operational Boundaries

Preload both `vexec` and `vexec_postgis` on every participating server and restart; preload order is unrestricted. `vexec.mode` must enable vector execution. `CREATE EXTENSION vexec_postgis` is optional inspection support and requires the underlying `postgis` extension; it does not activate hooks. `vexec.declared_calls()` shows the bound calls. A superuser creates the non-relocatable SQL objects.

The pack binds only PostGIS 3.7.0rc2. Other versions run without pack acceleration. Bounding-box prefilters and point checks decide only supported cases; remaining rows use the original PostGIS functions. It does not add spatial storage or a new index method.
