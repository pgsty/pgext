## Usage

Sources:

- [FEP 18 quick-start guide](https://www.postgresql.fastware.com/knowledge-base/quick-start-guides)
- [FEP 18 knowledge data management manual](https://www.postgresql.fastware.com/hubfs/_Global/Manuals/FEP-v18forx86-KnowledgeDataManagementUserGuide.pdf)

`pgx_inference` 1.0 in Fujitsu Enterprise Postgres 18 manages ONNX text-embedding pipeline models and executes them through Triton on the database host. Models must satisfy the vendor's input/output and operator requirements.

### Core Workflow

Configure Triton, model paths, gRPC/TLS and worker capacity first. Add `pgx_inference` to `shared_preload_libraries` and restart. Enable the extension and start its load launcher as the extension owner:

```sql
CREATE EXTENSION pgx_inference CASCADE;
SELECT pgx_inference.pgx_launch_load_launcher();
```

Import a compatible model with the vendor `pgx_aimodel_tool`. For a model imported as `sample-model-v1`, request loading and inspect its status:

```sql
SELECT * FROM pgx_inference.pgx_list_model_metadata();
SELECT pgx_inference.pgx_load_model('sample-model-v1');
SELECT * FROM pgx_inference.pgx_triton_model_status;
```

Loading is asynchronous. Once the model reports `is_available = true`, generate an embedding:

```sql
SELECT pgx_inference.pgx_onnx_embed('sample-model-v1', 'PostgreSQL replication');
```

### Operational Boundaries

Model import, loading, execution and schema access have separate privileges. Restrict Triton access with mTLS to enforce per-model access control. Backups must include large objects; unload models and confirm completion before deleting them. This feature requires the Fujitsu runtime.
