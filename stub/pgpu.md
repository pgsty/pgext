## Usage

Sources:

- [Upstream README](https://github.com/EnterpriseDB/pgpu/blob/b39cc5e3fa49efb673c524d722ea7c812228bc7e/README.md)
- [Extension control file](https://github.com/EnterpriseDB/pgpu/blob/b39cc5e3fa49efb673c524d722ea7c812228bc7e/pgpu.control)
- [Cargo package metadata](https://github.com/EnterpriseDB/pgpu/blob/b39cc5e3fa49efb673c524d722ea7c812228bc7e/Cargo.toml)
- [VectorChord index creation implementation](https://github.com/EnterpriseDB/pgpu/blob/b39cc5e3fa49efb673c524d722ea7c812228bc7e/vectorchord-indexing/src/vectorchord_index.rs)

`pgpu` version `2.1.0` uses an NVIDIA GPU to compute vector centroids for a VectorChord index. It reads pgvector data in batches, runs cuVS k-means, writes a centroids table, and optionally drops/rebuilds a VectorChord index with those external centroids.

### Smoke test

The build requires the documented CUDA 13.1/cuVS environment plus matching PostgreSQL, `vector`, and `vchord` installations. After installation:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION vchord;
CREATE EXTENSION pgpu;
SELECT pgpu.ping();
SET pgpu.gpu_acceleration = 'enable';
```

The main entry point is `pgpu.create_vector_index_on_gpu(...)`; it is a schema-changing, GPU- and memory-intensive operation, so test parameters and recall on a copy first. Version `2.1.0` supports PostgreSQL 14–18 in Cargo features, but the reviewed index builder drops a generated index name and hardcodes the indexed column as `embedding` during the final VectorChord DDL. Restrict execution to table owners and verify generated names/DDL, GPU resource limits, distance operator, dimensions, and recovery behavior.
