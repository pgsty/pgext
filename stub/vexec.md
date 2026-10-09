## Usage

Sources:

- [README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/README.md)
- [modules/vexec/vexec--1.0.sql](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec/vexec--1.0.sql)
- [modules/vexec/vexec.control](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec/vexec.control)

`vexec` plans and executes batches inside PostgreSQL 19. It can use the PostgreSQL planner or the separately installed ORCA integration. Unsupported expressions fall back to PostgreSQL evaluation; the Cloudberry storage and interconnect integrations have additional fork-specific requirements.

### Core Workflow

```ini
shared_preload_libraries = 'vexec'
vexec.mode = auto
```

```sql
EXPLAIN (VEXEC) SELECT count(*) FROM pg_catalog.pg_class;
```

### Operational Boundaries

Append the library to the existing preload list and restart. `vexec.mode` defaults to off; auto uses the cost model and force selects supported vector paths more aggressively. Creating the SQL extension as a superuser is optional for ordinary vector execution, but supplies `vexec.sources`, `vexec.sinks`, `vexec.type_layouts`, `vexec.kernels`, `vexec.kernel_packs`, `vexec.declared_calls` and `vexec.ingest_stream`. The last function is required for Flight bulk ingestion.

This early source release targets PostgreSQL 19. Compare results and plans with mode off before adoption. ORCA requires the compatible gp_orca module; PAX, ao_column and distributed execution require the separate Cloudberry port. These optional integrations are not supplied by installing the base extension.
