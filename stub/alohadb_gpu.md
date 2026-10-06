## Usage

Sources:

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_gpu/alohadb_gpu--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_gpu/alohadb_gpu--1.0.sql)
- [contrib/alohadb_gpu/alohadb_gpu.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_gpu/alohadb_gpu.c)
- [contrib/alohadb_gpu/alohadb_gpu.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_gpu/alohadb_gpu.control)

`alohadb_gpu` 1.0 provides batch vector distances, top-K selection and matrix multiplication in the PostgreSQL 18-based AlohaDB fork. It uses the fork’s built-in vector type and is not a drop-in pgvector add-on for stock PostgreSQL.

### Core Workflow

```sql
CREATE EXTENSION alohadb_gpu;
SELECT alohadb_gpu_available();
SELECT * FROM alohadb_gpu_topk(
  '[0.1,0.2,0.3]'::vector,
  ARRAY['[0.1,0.0,0.2]'::vector, '[0.4,0.5,0.6]'::vector], 1);
```

### Operational Boundaries

`alohadb_gpu_batch_l2` and `alohadb_gpu_batch_cosine` return per-candidate distances; `alohadb_gpu_topk` returns candidate positions and L2 distances; `alohadb_gpu_matmul` multiplies two-dimensional float4 arrays. Inputs must have compatible dimensions.

`alohadb.gpu_device_id` selects the CUDA device and `alohadb.gpu_min_batch_size` controls when batches use the GPU. Small batches or unavailable CUDA support use the CPU implementation; `alohadb_gpu_available` reports availability. Large arrays still use database-process memory. Installation requires a superuser; no preload or restart is required by the reviewed module. Availability depends on how AlohaDB was built, not just the presence of a GPU.
