## Usage

Sources:

- [Official source](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/README.md)
- [Official source](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/pg_vxvector/pg_vxvector.control)
- [Official source](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/pg_vxvector/sql/pg_vxvector--0.1.0.sql)
- [Official source](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/docs/pg_vxvector.md)
- [Official source](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/pg_vxvector/Cargo.toml)

`pg_vxvector` 0.1.0 trains centroids in an external VxVector worker and can use them to build a VectorChord index. PostgreSQL 17, `vector` and `vchord` are required. This is an experimental training integration, not query-time GPU search.

### Train centroids

Install matching extension and worker files, configure `vxvector.worker_path`, preload `vchord` and restart. Installation needs a superuser; calls need SELECT on the source column and ordinary privileges for creating output objects. The example assumes an existing vector-valued column and deliberately requests centroids without index creation.

```sql
CREATE EXTENSION vector;
CREATE EXTENSION vchord;
CREATE EXTENSION pg_vxvector;
SELECT vxvector.create_vector_index(
    'public.items'::regclass, 'embedding', ARRAY[100],
    distance_operator => 'l2', skip_index_build => true
);
```

### Controls and failure boundaries

`vxvector.create_vector_index` returns JSON describing its outputs. Parameters include `lists`, `sampling_factor`, `batch_size`, `kmeans_iterations`, `kmeans_nredo`, `seed`, `memory_budget_bytes`, `worker_timeout_ms`, `centroid_table` and `index_name`. Review the dimensionality, sample size and memory budget before a large run. RLS is respected; `halfvec`, `sparsevec` and `bit` are unsupported.

The superuser setting `vxvector.backend_mode` accepts `auto`, `cpu` and `gpu`. Automatic selection may choose CPU before training if no supported GPU is available; a GPU failure after training starts is reported rather than silently retried on CPU. The current dated roadmap and README supersede the older native-macOS blocker: one macOS/Apple Metal PostgreSQL 17 setup was validated. Linux Docker uses CPU; the CUDA backend has not been validated on NVIDIA hardware. Keep the documented compiler/worker revision matched; no production readiness or general GPU compatibility is implied.
