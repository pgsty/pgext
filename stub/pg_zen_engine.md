## Usage

Sources:

- [Official README](https://github.com/foxflow/pg_zen_engine/blob/e4924af4f5a0ad7da3d0031ee5e54ef6698db0b8/README.md)
- [Rust implementation](https://github.com/foxflow/pg_zen_engine/blob/e4924af4f5a0ad7da3d0031ee5e54ef6698db0b8/src/lib.rs)
- [Extension control file](https://github.com/foxflow/pg_zen_engine/blob/e4924af4f5a0ad7da3d0031ee5e54ef6698db0b8/pg_zen_engine.control)

`pg_zen_engine` evaluates GoRules JSON Decision Model graphs against `jsonb` input inside PostgreSQL. It exposes one Rust-backed function and is useful when the same stored decision graph must be applied close to relational data.

### Core Workflow

Store a valid JDM graph as `jsonb`, pass the per-row input document to `evaluate_jdm()`, and extract the returned engine result.

```sql
CREATE EXTENSION pg_zen_engine;

SELECT name,
       evaluate_jdm(graph, jsonb_build_object('input', facts)) AS decision
FROM decision_models
JOIN decision_inputs USING (name);
```

`evaluate_jdm(graph jsonb, data jsonb) RETURNS jsonb` is declared strict, immutable, and parallel-safe. Each call deserializes the graph, creates a new decision engine, evaluates the input, and serializes the complete result; the extension does not provide a model cache, management table, or validation function.

### Operational Boundaries

Installing the extension requires superuser privileges according to its control file. Invalid graph JSON, unsupported model content, or evaluation failure reaches Rust paths that use `unwrap()`, so validate models outside critical queries and exercise failure cases on the exact build. Treat model expressions as executable business logic and limit who may change stored graphs.

The reviewed `0.0.1` crate pins Zen Engine `0.22.0` and advertises PostgreSQL 12 through 16 features. Decision-model semantics can change with the engine dependency even when the SQL signature does not; pin the complete build, retain model fixtures with expected results, and benchmark repeated evaluation before applying it to large scans.
