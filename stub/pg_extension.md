## Usage

Sources:

- [indices developer guide at the reviewed commit](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/documents/dev_guide.md)
- [pg_extension Rust SQL API at the reviewed commit](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/src/lib.rs)
- [pg_extension Cargo configuration at the reviewed commit](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/Cargo.toml)

`pg_extension` 0.1.0 is the in-database research component of the indices project. It exposes Rust/pgrx functions for model selection, model slicing, profiling, and inference. The documented `inference` call accepts a dataset, JSON condition, configuration and cardinality files, a model path, a SQL filter fragment, and a batch size.

```sql
CREATE EXTENSION pg_extension;

SELECT inference(
  'frappe',
  '{"2":977}',
  '/project/Trails/internal/ml/model_selection/config.ini',
  '/project/Trails/frappe_col_cardinalities',
  '/project/tensor_log/frappe/dnn_K16_alpha4',
  'WHERE col2=''977:1''',
  10000
);
```

This example is repository-specific and works only after reproducing its Python environment, tables, model files, and configuration layout.

### Caveats

- The name collides with PostgreSQL's `pg_extension` system catalog; qualify catalog references when discussing or querying both.
- The control file requires superuser. The default build embeds Python support and declares features only for PostgreSQL 11 through 15, with PostgreSQL 14 as the default.
- Source and guide contain deployment-specific absolute paths and assumptions about dataset table names and column counts. This is not a portable general-purpose inference API.
- Several APIs are declared immutable and parallel safe even though their implementations query tables, load models, or use shared memory. Treat those planner labels as unsafe until independently audited.
- The SQL filter argument is incorporated into a dynamically built query. Accept it only from trusted, validated configuration.
