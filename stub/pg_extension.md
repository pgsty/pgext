## Usage

Sources:

- [indices developer guide](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/documents/dev_guide.md)
- [pg_extension control file](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/pg_extension.control)
- [SQL-facing Rust API](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/src/lib.rs)
- [Inference and SPI implementation](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/src/bindings/inference.rs)
- [Cargo features and dependencies](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/Cargo.toml)

`pg_extension` 0.1.0 is the in-database research component of the `indices` project. It embeds Python through pyo3 and exposes experimental model-selection, profiling, model initialization, and inference functions. It is tied to the repository's datasets, Python modules, model artifacts, configuration files, and filesystem layout; it is not a portable general-purpose ML extension.

### Repository-Specific Workflow

The control file requires superuser, and the default Cargo features build for PostgreSQL 14 with embedded Python. The developer guide creates the extension only after reproducing its container, Python dependencies, database tables, and absolute `/project/...` paths:

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

This example is a reproduction target for the upstream experiment, not a safe template for arbitrary application input.

### API Groups

- Model selection: `model_selection`, `model_selection_workloads`, `model_selection_trails`, and coordinator/filter/refinement helpers.
- Profiling and benchmarks: `profiling_filtering_phase`, `profiling_refinement_phase`, and latency benchmark functions.
- Inference: `inference`, `inference_shared`, `inference_shared_write_once`, integer/shared-memory variants, and `model_init`.
- Build matrix: Cargo declares PostgreSQL 11–15 features; default is PostgreSQL 14 plus `python`.

### Security and Correctness Boundaries

The name `pg_extension` collides with PostgreSQL's `pg_catalog.pg_extension` system catalog. Qualify catalog queries explicitly.

Many functions are declared `IMMUTABLE PARALLEL SAFE` even though they read database tables, load and mutate Python/model state, access files, create fixed-name OS shared memory such as `my_shared_memory`, write logs, and measure time. Those planner labels do not match the implementation. Do not use these functions in indexes, generated columns, constraints, or parallel production queries without a full audit and corrected volatility/parallel declarations.

Dataset names and a caller-supplied SQL fragment are concatenated into SPI queries. Fixed shared-memory names and unchecked `unwrap` paths also make concurrent calls unsafe. Treat every argument as administrator-controlled, isolate the experiment from production data, and review filesystem, Python import, native-library, memory, and SQL-injection boundaries before execution.
