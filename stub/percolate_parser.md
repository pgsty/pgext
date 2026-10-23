## Usage

Sources:

- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/README.md)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate_parser/percolate_parser.control)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate_parser/Cargo.toml)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate_parser/sql/surface.sql)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate_parser/sql/percolate_parser--0.2.29--0.2.30.sql)

`percolate_parser` 0.2.30 contains the compiled query, YAML-workflow and graph routines used by Percolate. This Rust extension currently declares PostgreSQL 19 and pgrx 0.19.2. It separates native functions from the non-superuser-owned SQL application.

### Install and Call

```sql
CREATE EXTENSION percolate_parser;
SELECT extname, extversion FROM pg_extension WHERE extname = 'percolate_parser';
```

`aiq_parse(query)` parses the application's query dialect into JSONB; `p8_compile_workflow(yaml)` compiles a workflow into JSONB. Callers must supply valid Percolate syntax, rather than assuming arbitrary SQL or YAML is accepted. The matching SQL application defines the surrounding execution, identity and authorization workflow.

### Graph APIs and Boundaries

The SQL surface also exposes `p8_graph_components`, `p8_graph_communities`, `p8_graph_distance` and budgeted traversal functions. These graph operations depend on the application's graph tables and context; the parser extension alone is not a complete graph service. Budget arguments constrain work but do not establish universal latency guarantees.

Install with a superuser before installing `percolate` as its application owner. The parser is not relocatable and no preload requirement is declared. Keep parser and application versions aligned. Its generated install SQL and versioned upgrade SQL must match the native library; copying only the shared library is not a complete upgrade.
