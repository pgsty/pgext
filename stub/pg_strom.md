

## Usage

> [pg_strom: Big-data processing acceleration using GPU and NVME](https://github.com/heterodb/pg-strom)

PG-Strom is a PostgreSQL extension that offloads analytical SQL workloads to GPU processors.
It automatically generates GPU code from SQL, accelerating scan, join, aggregation, and sort operations.
Full documentation is available at [https://heterodb.github.io/pg-strom/](https://heterodb.github.io/pg-strom/).

### Documentation Index

| Topic | Description |
|:------|:------------|
| [Install](https://heterodb.github.io/pg-strom/install/) | Prerequisites, CUDA setup, RPM/source installation |
| [Basic Operations](https://heterodb.github.io/pg-strom/operations/) | GpuScan, GpuJoin, GpuPreAgg fundamentals |
| [BRIN Index](https://heterodb.github.io/pg-strom/brin/) | BRIN index support for GPU table scans |
| [Partitioning](https://heterodb.github.io/pg-strom/partition/) | Partition-wise GPU execution |
| [PostGIS](https://heterodb.github.io/pg-strom/postgis/) | GPU-accelerated PostGIS functions |
| [GPU-Sort](https://heterodb.github.io/pg-strom/gpusort/) | GPU-accelerated sorting with bitonic sort |
| [GPUDirect SQL](https://heterodb.github.io/pg-strom/ssd2gpu/) | NVMe-to-GPU direct data transfer |
| [Apache Arrow / Arrow_Fdw](https://heterodb.github.io/pg-strom/arrow_fdw/) | Columnar Arrow file foreign data wrapper |
| [GPU Cache](https://heterodb.github.io/pg-strom/gpucache/) | In-GPU-memory table caching |
| [Pinned Inner Buffer](https://heterodb.github.io/pg-strom/pinned_buffer/) | Retaining inner-table results on GPU |
| [Data Types](https://heterodb.github.io/pg-strom/ref_types/) | GPU-supported data types (int1, float2, etc.) |
| [Functions & Operators](https://heterodb.github.io/pg-strom/ref_devfuncs/) | GPU-executable functions and operators |
| [SQL Objects](https://heterodb.github.io/pg-strom/ref_sqlfuncs/) | System views and utility functions |
| [GUC Parameters](https://heterodb.github.io/pg-strom/ref_params/) | All configuration parameters |
| [Troubleshooting](https://heterodb.github.io/pg-strom/troubles/) | Diagnostics and common issues |


### Prerequisites

PG-Strom requires an NVIDIA GPU (compute capability 7.5+ / Turing or later), CUDA Toolkit 12.2+, and PostgreSQL 15+.
It must be loaded via `shared_preload_libraries`.

```ini
# postgresql.conf
shared_preload_libraries = 'pg_strom'
max_worker_processes = 100
shared_buffers = 10GB
work_mem = 1GB
```

```sql
CREATE EXTENSION pg_strom;
SELECT pgstrom.device_info;  -- verify GPU detection
```


### GpuScan: GPU-Accelerated Table Scan

PG-Strom replaces sequential scans with GPU-parallel scans. WHERE clause filters
are compiled into GPU kernels and executed on the device. This appears as `GpuScan`
in EXPLAIN output.

```sql
-- GPU will accelerate this filtered scan automatically
EXPLAIN (ANALYZE, COSTS OFF)
SELECT * FROM lineorder WHERE lo_quantity > 40 AND lo_discount BETWEEN 1 AND 3;
```

```
 Custom Scan (GpuScan) on lineorder
   GPU Filter: ((lo_quantity > 40) AND (lo_discount >= 1) AND (lo_discount <= 3))
   Rows Removed by GPU Filter: 59532300
   ...
```


### GpuJoin: GPU-Accelerated Hash Join

Multi-table joins are offloaded as GPU hash joins. PG-Strom consolidates
sequential scan-join-aggregate chains into a single GPU kernel to minimize
CPU-GPU data transfer.

```sql
EXPLAIN (ANALYZE, COSTS OFF)
SELECT d_year, s_nation, sum(lo_revenue) AS revenue
  FROM lineorder, date1, supplier
 WHERE lo_orderdate = d_datekey
   AND lo_suppkey = s_suppkey
   AND s_region = 'ASIA'
 GROUP BY d_year, s_nation
 ORDER BY d_year, s_nation;
```

```
 Sort
   Sort Key: date1.d_year, supplier.s_nation
   ->  Custom Scan (GpuPreAgg) on lineorder
         GPU Projection: ...
         GPU Join Quals [1]: (lo_orderdate = d_datekey)  ... [HashJoin]
         GPU Join Quals [2]: (lo_suppkey = s_suppkey)    ... [HashJoin]
         GPU Outer Quals: (s_region = 'ASIA')
         GPU Group Key: d_year, s_nation
         ->  Seq Scan on date1
         ->  Seq Scan on supplier
```


### GpuPreAgg: GPU-Accelerated Aggregation

GROUP BY and aggregate functions (SUM, AVG, COUNT, MIN, MAX, STDDEV, etc.) are
executed on the GPU. PG-Strom performs a preliminary aggregation on the device,
then the CPU handles the final merge.

```sql
EXPLAIN (ANALYZE, COSTS OFF)
SELECT lo_shipmode, count(*), avg(lo_quantity)
  FROM lineorder
 GROUP BY lo_shipmode;
```

```
 Finalize GroupAggregate
   Group Key: lo_shipmode
   ->  Custom Scan (GpuPreAgg) on lineorder
         GPU Projection: lo_shipmode, pgstrom.nrows(), pgstrom.psum(lo_quantity)
         GPU Group Key: lo_shipmode
```


### GPU-Sort

GPU-Sort uses a bitonic sorting algorithm for ORDER BY and window functions.
It requires CPU fallback to be disabled so all data resides in GPU memory.

```sql
SET pg_strom.cpu_fallback = off;

EXPLAIN (ANALYZE, COSTS OFF)
SELECT *, rank() OVER (PARTITION BY lo_shipmode ORDER BY lo_revenue DESC) AS rnk
  FROM lineorder
 WHERE lo_quantity > 45;
```


### Arrow_Fdw: Apache Arrow Columnar Store

Arrow_Fdw maps Apache Arrow files as foreign tables, enabling GPU-accelerated
queries on columnar data. Combined with GPUDirect SQL, data flows directly
from NVMe storage to the GPU.

```sql
-- Create a foreign table from an Arrow file
CREATE FOREIGN TABLE arrow_logs (
    ts         timestamp,
    sensor_id  int,
    signal     float4
) SERVER arrow_fdw
  OPTIONS (file '/data/logs/sensor_2024.arrow');

-- Query with GPU acceleration
SELECT sensor_id, avg(signal)
  FROM arrow_logs
 WHERE ts BETWEEN '2024-01-01' AND '2024-06-30'
 GROUP BY sensor_id;

-- Map an entire directory of Arrow files
CREATE FOREIGN TABLE arrow_archive (
    ts         timestamp,
    device_id  int,
    value      float8
) SERVER arrow_fdw
  OPTIONS (dir '/data/archive', suffix '.arrow');
```

Export PostgreSQL data to Arrow format using the `pg2arrow` CLI tool:

```bash
pg2arrow -d mydb -c "SELECT * FROM sensor_data" -o /data/sensor_data.arrow
```


### GPU Cache

GPU Cache keeps a copy of a PostgreSQL table in GPU device memory for
ultra-fast analytical queries on frequently updated data (suitable for tables up to ~10GB).
Synchronization is log-based via a row-level trigger.

```sql
-- Enable GPU Cache on a table
CREATE TRIGGER row_sync
  AFTER INSERT OR UPDATE OR DELETE ON realtime_metrics
  FOR ROW EXECUTE FUNCTION pgstrom.gpucache_sync_trigger();
ALTER TABLE realtime_metrics ENABLE ALWAYS TRIGGER row_sync;

-- With custom parameters
CREATE TRIGGER row_sync
  AFTER INSERT OR UPDATE OR DELETE ON dpoints
  FOR ROW EXECUTE FUNCTION pgstrom.gpucache_sync_trigger(
    'gpu_device_id=0,max_num_rows=5000000,redo_buffer_size=200m,gpu_sync_interval=4'
  );

-- Monitor GPU Cache status
SELECT * FROM pgstrom.gpucache_info;

-- Preload tables at startup (postgresql.conf)
-- pg_strom.gpucache_auto_preload = 'mydb.public.realtime_metrics'
```


### GPUDirect SQL

GPUDirect SQL enables peer-to-peer DMA transfers from NVMe-SSD storage directly
to GPU memory, bypassing CPU and system RAM entirely. This is ideal for large
analytical scans where the GPU filters data before the CPU sees it.

Requirements: NVMe-SSD storage, NVIDIA GPUDirect Storage support, `nvidia-fs` kernel module.

```sql
-- GPUDirect SQL activates automatically for large tables on NVMe
-- The threshold is controlled by:
SET pg_strom.gpudirect_threshold = '2GB';
SET pg_strom.gpudirect_enabled = on;

-- Check GPUDirect driver status
SHOW pg_strom.gpudirect_driver;
```

In EXPLAIN output, GPUDirect SQL appears as a note on the scan node indicating
direct NVMe-to-GPU data transfer.


### PostGIS GPU Acceleration

PG-Strom accelerates several PostGIS functions on the GPU, including
`st_contains`, `st_crosses`, `st_relate`, and `st_makepoint`.
GpuJoin can leverage GiST (R-Tree) indexes for spatial join acceleration.

```sql
-- GPU-accelerated spatial filter
EXPLAIN (ANALYZE, COSTS OFF)
SELECT * FROM gps_points
 WHERE st_contains(
    'POLYGON((139.6 35.5, 139.8 35.5, 139.8 35.7, 139.6 35.7, 139.6 35.5))'::geometry,
    st_makepoint(longitude, latitude)
 );

-- GPU-accelerated spatial join with GiST index
EXPLAIN (ANALYZE, COSTS OFF)
SELECT a.id, b.name
  FROM gps_points a, city_boundaries b
 WHERE st_contains(b.geom, st_makepoint(a.longitude, a.latitude));
```

```
 Custom Scan (GpuJoin) on gps_points a
   GPU Projection: a.id, b.name
   GPU GiST Join Quals [1]: st_contains(b.geom, st_makepoint(a.longitude, a.latitude))
   ->  Custom Scan (GpuScan) on city_boundaries b
```


### Custom Data Types

PG-Strom provides additional data types optimized for GPU processing:

| Type | Size | Description |
|:-----|:-----|:------------|
| `int1` | 1 byte | 8-bit integer |
| `float2` | 2 bytes | Half-precision floating-point (IEEE 754) |

These types are especially useful for compact storage of large datasets where
precision can be reduced to save memory and increase GPU throughput.


### System Views and Functions

```sql
-- GPU device properties
SELECT * FROM pgstrom.device_info;

-- GPU Cache status
SELECT * FROM pgstrom.gpucache_info;

-- Module version info
SELECT pgstrom.githash();

-- Import Arrow file as foreign table
SELECT pgstrom.arrow_fdw_import_file('arrow_table', '/data/file.arrow');

-- GPU Cache management
SELECT pgstrom.gpucache_apply_redo('realtime_metrics'::regclass);
SELECT pgstrom.gpucache_compaction('realtime_metrics'::regclass);
SELECT pgstrom.gpucache_recovery('realtime_metrics'::regclass);
```


### Key GUC Parameters

**Feature toggles:**

| Parameter | Default | Description |
|:----------|:--------|:------------|
| `pg_strom.enabled` | `on` | Master switch for all PG-Strom features |
| `pg_strom.enable_gpuscan` | `on` | Enable/disable GpuScan |
| `pg_strom.enable_gpujoin` | `on` | Enable/disable GpuJoin |
| `pg_strom.enable_gpuhashjoin` | `on` | Enable/disable GpuHashJoin |
| `pg_strom.enable_gpugistindex` | `on` | Enable/disable GpuGiSTIndex join |
| `pg_strom.enable_gpupreagg` | `on` | Enable/disable GpuPreAgg |
| `pg_strom.enable_gpusort` | `on` | Enable/disable GPU-Sort |
| `pg_strom.enable_brin` | `on` | Enable/disable BRIN index on GPU scans |
| `pg_strom.enable_gpucache` | `on` | Enable/disable GPU Cache usage |
| `pg_strom.cpu_fallback` | `notice` | CPU fallback behavior on GPU recheck errors |

**Optimizer cost parameters:**

| Parameter | Default | Description |
|:----------|:--------|:------------|
| `pg_strom.gpu_setup_cost` | `100 * seq_page_cost` | Startup cost for GPU device initialization |
| `pg_strom.gpu_tuple_cost` | `cpu_tuple_cost` | Per-tuple CPU-GPU transfer cost |
| `pg_strom.gpu_operator_cost` | `cpu_operator_cost / 16` | Per-operator GPU execution cost |

**GPUDirect SQL:**

| Parameter | Default | Description |
|:----------|:--------|:------------|
| `pg_strom.gpudirect_enabled` | `on` | Enable/disable GPUDirect SQL |
| `pg_strom.gpudirect_threshold` | auto | Minimum table size to trigger GPUDirect SQL |
| `pg_strom.gpu_direct_seq_page_cost` | `seq_page_cost / 4` | Scan cost via GPUDirect SQL |

**Arrow_Fdw:**

| Parameter | Default | Description |
|:----------|:--------|:------------|
| `arrow_fdw.enabled` | `on` | Enable/disable Arrow_Fdw |
| `arrow_fdw.metadata_cache_size` | `512MB` | Shared memory for Arrow metadata cache |

**GPU memory management:**

| Parameter | Default | Description |
|:----------|:--------|:------------|
| `pg_strom.gpu_mempool_segment_sz` | `1GB` | GPU memory pool segment size |
| `pg_strom.gpu_mempool_max_ratio` | `50%` | Max device memory for the pool |
| `pg_strom.gpu_mempool_min_ratio` | `5%` | Min preserved memory pool ratio |
| `pg_strom.cuda_visible_devices` | (all) | Restrict to specific GPU device IDs |

**Execution:**

| Parameter | Default | Description |
|:----------|:--------|:------------|
| `pg_strom.max_async_tasks` | `12` | Max concurrent GPU tasks per query |
| `pg_strom.enable_partitionwise_gpujoin` | `on` | Push GpuJoin into partitions |
| `pg_strom.enable_partitionwise_gpupreagg` | `on` | Push GpuPreAgg into partitions |
