## Usage

Sources:

- [Upstream source 1](https://github.com/landstats-us/postgis-cuda/blob/43fbd87c90e225ab49d6ff16a55ec7f8d00c3edf/README.md)
- [Upstream source 2](https://github.com/landstats-us/postgis-cuda/blob/43fbd87c90e225ab49d6ff16a55ec7f8d00c3edf/ext/gpu_spatial.control)
- [Upstream source 3](https://github.com/landstats-us/postgis-cuda/blob/43fbd87c90e225ab49d6ff16a55ec7f8d00c3edf/ext/gpu_spatial--0.1.sql)

`gpu_spatial` 0.1 exposes CUDA spatial operations over layers held by the external gpujoind service.

### Core Workflow

With the service running and a prepared spatial source table, `gpu_load` loads a named snapshot and `gpu_unload` removes it. `gpu_join` returns matching identifiers and optional intersection area; `gpu_relate`, `gpu_dwithin` and `gpu_cluster_intersecting` cover additional spatial operations. Inspect `gpu_last_stats` for the last operation.

```sql
CREATE EXTENSION gpu_spatial;
SELECT gpu_status();
SELECT * FROM gpu_layers();
```

### Operational Boundaries

Installation is superuser-only. Loading/unloading functions revoke PUBLIC execution; query access still needs deliberate privilege review. The documented environment uses NVIDIA Hopper sm90 and a separate CUDA daemon, not arbitrary graphics hardware. Resident layers do not follow source-table changes automatically: reload them explicitly. Coordinate-system and geography distance semantics must match the upstream function contract.
