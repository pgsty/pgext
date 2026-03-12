


## 用法

> [pg_strom: 使用 GPU 和 NVMe 加速大数据处理](https://github.com/heterodb/pg-strom)

PG-Strom 是一个 PostgreSQL 扩展，可将分析型 SQL 工作负载卸载到 GPU 处理器上执行。
它能自动从 SQL 生成 GPU 代码，加速扫描、连接、聚合和排序操作。
完整文档请参阅 [https://heterodb.github.io/pg-strom/](https://heterodb.github.io/pg-strom/)。

### 文档索引

| 主题 | 说明 |
|:------|:------------|
| [安装](https://heterodb.github.io/pg-strom/install/) | 前置条件、CUDA 配置、RPM/源码安装 |
| [基础操作](https://heterodb.github.io/pg-strom/operations/) | GpuScan、GpuJoin、GpuPreAgg 基础 |
| [BRIN 索引](https://heterodb.github.io/pg-strom/brin/) | GPU 表扫描的 BRIN 索引支持 |
| [分区表](https://heterodb.github.io/pg-strom/partition/) | 分区级 GPU 执行 |
| [PostGIS](https://heterodb.github.io/pg-strom/postgis/) | GPU 加速的 PostGIS 函数 |
| [GPU 排序](https://heterodb.github.io/pg-strom/gpusort/) | 基于双调排序的 GPU 加速排序 |
| [GPUDirect SQL](https://heterodb.github.io/pg-strom/ssd2gpu/) | NVMe 到 GPU 的直接数据传输 |
| [Apache Arrow / Arrow_Fdw](https://heterodb.github.io/pg-strom/arrow_fdw/) | 列式 Arrow 文件外部数据包装器 |
| [GPU 缓存](https://heterodb.github.io/pg-strom/gpucache/) | GPU 设备内存中的表缓存 |
| [Pinned Inner Buffer](https://heterodb.github.io/pg-strom/pinned_buffer/) | 在 GPU 上保留内表结果 |
| [数据类型](https://heterodb.github.io/pg-strom/ref_types/) | GPU 支持的数据类型（int1、float2 等） |
| [函数与运算符](https://heterodb.github.io/pg-strom/ref_devfuncs/) | GPU 可执行的函数和运算符 |
| [SQL 对象](https://heterodb.github.io/pg-strom/ref_sqlfuncs/) | 系统视图和工具函数 |
| [GUC 参数](https://heterodb.github.io/pg-strom/ref_params/) | 所有配置参数 |
| [故障排查](https://heterodb.github.io/pg-strom/troubles/) | 诊断与常见问题 |


### 前置条件

PG-Strom 需要 NVIDIA GPU（计算能力 7.5+ / Turing 及以上）、CUDA Toolkit 12.2+ 和 PostgreSQL 15+。
必须通过 `shared_preload_libraries` 预加载。

```ini
# postgresql.conf
shared_preload_libraries = 'pg_strom'
max_worker_processes = 100
shared_buffers = 10GB
work_mem = 1GB
```

```sql
CREATE EXTENSION pg_strom;
SELECT pgstrom.device_info;  -- 验证 GPU 检测
```


### GpuScan：GPU 加速表扫描

PG-Strom 用 GPU 并行扫描替代顺序扫描。WHERE 子句过滤条件被编译为 GPU 内核并在设备上执行，
在 EXPLAIN 输出中显示为 `GpuScan`。

```sql
-- GPU 会自动加速此过滤扫描
EXPLAIN (ANALYZE, COSTS OFF)
SELECT * FROM lineorder WHERE lo_quantity > 40 AND lo_discount BETWEEN 1 AND 3;
```

```
 Custom Scan (GpuScan) on lineorder
   GPU Filter: ((lo_quantity > 40) AND (lo_discount >= 1) AND (lo_discount <= 3))
   Rows Removed by GPU Filter: 59532300
   ...
```


### GpuJoin：GPU 加速哈希连接

多表连接被卸载为 GPU 哈希连接。PG-Strom 将扫描-连接-聚合链合并为单个 GPU 内核，
以最小化 CPU-GPU 之间的数据传输。

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


### GpuPreAgg：GPU 加速聚合

GROUP BY 和聚合函数（SUM、AVG、COUNT、MIN、MAX、STDDEV 等）在 GPU 上执行。
PG-Strom 在设备上完成初步聚合，然后由 CPU 进行最终合并。

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


### GPU 排序

GPU-Sort 使用双调排序（bitonic sort）算法来加速 ORDER BY 和窗口函数。
需要禁用 CPU 回退，以确保所有数据驻留在 GPU 内存中。

```sql
SET pg_strom.cpu_fallback = off;

EXPLAIN (ANALYZE, COSTS OFF)
SELECT *, rank() OVER (PARTITION BY lo_shipmode ORDER BY lo_revenue DESC) AS rnk
  FROM lineorder
 WHERE lo_quantity > 45;
```


### Arrow_Fdw：Apache Arrow 列式存储

Arrow_Fdw 将 Apache Arrow 文件映射为外部表，支持对列式数据的 GPU 加速查询。
结合 GPUDirect SQL，数据可直接从 NVMe 存储传输到 GPU。

```sql
-- 从 Arrow 文件创建外部表
CREATE FOREIGN TABLE arrow_logs (
    ts         timestamp,
    sensor_id  int,
    signal     float4
) SERVER arrow_fdw
  OPTIONS (file '/data/logs/sensor_2024.arrow');

-- GPU 加速查询
SELECT sensor_id, avg(signal)
  FROM arrow_logs
 WHERE ts BETWEEN '2024-01-01' AND '2024-06-30'
 GROUP BY sensor_id;

-- 映射整个 Arrow 文件目录
CREATE FOREIGN TABLE arrow_archive (
    ts         timestamp,
    device_id  int,
    value      float8
) SERVER arrow_fdw
  OPTIONS (dir '/data/archive', suffix '.arrow');
```

使用 `pg2arrow` 命令行工具将 PostgreSQL 数据导出为 Arrow 格式：

```bash
pg2arrow -d mydb -c "SELECT * FROM sensor_data" -o /data/sensor_data.arrow
```


### GPU 缓存

GPU 缓存在 GPU 设备内存中保留 PostgreSQL 表的副本，用于对频繁更新数据的超快分析查询
（适合约 10GB 以内的表）。同步基于日志，通过行级触发器实现。

```sql
-- 在表上启用 GPU 缓存
CREATE TRIGGER row_sync
  AFTER INSERT OR UPDATE OR DELETE ON realtime_metrics
  FOR ROW EXECUTE FUNCTION pgstrom.gpucache_sync_trigger();
ALTER TABLE realtime_metrics ENABLE ALWAYS TRIGGER row_sync;

-- 使用自定义参数
CREATE TRIGGER row_sync
  AFTER INSERT OR UPDATE OR DELETE ON dpoints
  FOR ROW EXECUTE FUNCTION pgstrom.gpucache_sync_trigger(
    'gpu_device_id=0,max_num_rows=5000000,redo_buffer_size=200m,gpu_sync_interval=4'
  );

-- 监控 GPU 缓存状态
SELECT * FROM pgstrom.gpucache_info;

-- 启动时预加载表（postgresql.conf）
-- pg_strom.gpucache_auto_preload = 'mydb.public.realtime_metrics'
```


### GPUDirect SQL

GPUDirect SQL 支持从 NVMe-SSD 存储到 GPU 内存的点对点 DMA 传输，
完全绕过 CPU 和系统内存。适用于大型分析扫描场景，GPU 在 CPU 处理之前先过滤数据。

要求：NVMe-SSD 存储、NVIDIA GPUDirect Storage 支持、`nvidia-fs` 内核模块。

```sql
-- GPUDirect SQL 对 NVMe 上的大表自动激活
-- 阈值由以下参数控制：
SET pg_strom.gpudirect_threshold = '2GB';
SET pg_strom.gpudirect_enabled = on;

-- 检查 GPUDirect 驱动状态
SHOW pg_strom.gpudirect_driver;
```

在 EXPLAIN 输出中，GPUDirect SQL 会在扫描节点上显示注释，表示 NVMe 到 GPU 的直接数据传输。


### PostGIS GPU 加速

PG-Strom 可在 GPU 上加速多个 PostGIS 函数，包括
`st_contains`、`st_crosses`、`st_relate` 和 `st_makepoint`。
GpuJoin 可利用 GiST（R-Tree）索引加速空间连接。

```sql
-- GPU 加速的空间过滤
EXPLAIN (ANALYZE, COSTS OFF)
SELECT * FROM gps_points
 WHERE st_contains(
    'POLYGON((139.6 35.5, 139.8 35.5, 139.8 35.7, 139.6 35.7, 139.6 35.5))'::geometry,
    st_makepoint(longitude, latitude)
 );

-- GPU 加速的空间连接（带 GiST 索引）
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


### 自定义数据类型

PG-Strom 提供为 GPU 处理优化的额外数据类型：

| 类型 | 大小 | 说明 |
|:-----|:-----|:------------|
| `int1` | 1 字节 | 8 位整数 |
| `float2` | 2 字节 | 半精度浮点数（IEEE 754） |

这些类型适用于可降低精度以节省内存并提高 GPU 吞吐量的大型数据集的紧凑存储。


### 系统视图与函数

```sql
-- GPU 设备属性
SELECT * FROM pgstrom.device_info;

-- GPU 缓存状态
SELECT * FROM pgstrom.gpucache_info;

-- 模块版本信息
SELECT pgstrom.githash();

-- 将 Arrow 文件导入为外部表
SELECT pgstrom.arrow_fdw_import_file('arrow_table', '/data/file.arrow');

-- GPU 缓存管理
SELECT pgstrom.gpucache_apply_redo('realtime_metrics'::regclass);
SELECT pgstrom.gpucache_compaction('realtime_metrics'::regclass);
SELECT pgstrom.gpucache_recovery('realtime_metrics'::regclass);
```


### 关键 GUC 参数

**功能开关：**

| 参数 | 默认值 | 说明 |
|:----------|:--------|:------------|
| `pg_strom.enabled` | `on` | PG-Strom 所有功能的主开关 |
| `pg_strom.enable_gpuscan` | `on` | 启用/禁用 GpuScan |
| `pg_strom.enable_gpujoin` | `on` | 启用/禁用 GpuJoin |
| `pg_strom.enable_gpuhashjoin` | `on` | 启用/禁用 GpuHashJoin |
| `pg_strom.enable_gpugistindex` | `on` | 启用/禁用 GpuGiSTIndex 连接 |
| `pg_strom.enable_gpupreagg` | `on` | 启用/禁用 GpuPreAgg |
| `pg_strom.enable_gpusort` | `on` | 启用/禁用 GPU-Sort |
| `pg_strom.enable_brin` | `on` | 启用/禁用 GPU 扫描的 BRIN 索引 |
| `pg_strom.enable_gpucache` | `on` | 启用/禁用 GPU 缓存 |
| `pg_strom.cpu_fallback` | `notice` | GPU 复核错误时的 CPU 回退行为 |

**优化器代价参数：**

| 参数 | 默认值 | 说明 |
|:----------|:--------|:------------|
| `pg_strom.gpu_setup_cost` | `100 * seq_page_cost` | GPU 设备初始化的启动代价 |
| `pg_strom.gpu_tuple_cost` | `cpu_tuple_cost` | 每行 CPU-GPU 传输代价 |
| `pg_strom.gpu_operator_cost` | `cpu_operator_cost / 16` | 每个运算符的 GPU 执行代价 |

**GPUDirect SQL：**

| 参数 | 默认值 | 说明 |
|:----------|:--------|:------------|
| `pg_strom.gpudirect_enabled` | `on` | 启用/禁用 GPUDirect SQL |
| `pg_strom.gpudirect_threshold` | 自动 | 触发 GPUDirect SQL 的最小表大小 |
| `pg_strom.gpu_direct_seq_page_cost` | `seq_page_cost / 4` | 通过 GPUDirect SQL 的扫描代价 |

**Arrow_Fdw：**

| 参数 | 默认值 | 说明 |
|:----------|:--------|:------------|
| `arrow_fdw.enabled` | `on` | 启用/禁用 Arrow_Fdw |
| `arrow_fdw.metadata_cache_size` | `512MB` | Arrow 元数据缓存的共享内存 |

**GPU 内存管理：**

| 参数 | 默认值 | 说明 |
|:----------|:--------|:------------|
| `pg_strom.gpu_mempool_segment_sz` | `1GB` | GPU 内存池段大小 |
| `pg_strom.gpu_mempool_max_ratio` | `50%` | 内存池最大设备内存占比 |
| `pg_strom.gpu_mempool_min_ratio` | `5%` | 内存池最小保留比例 |
| `pg_strom.cuda_visible_devices` | （全部） | 限定使用特定 GPU 设备 ID |

**执行参数：**

| 参数 | 默认值 | 说明 |
|:----------|:--------|:------------|
| `pg_strom.max_async_tasks` | `12` | 每个查询的最大并发 GPU 任务数 |
| `pg_strom.enable_partitionwise_gpujoin` | `on` | 将 GpuJoin 下推到分区 |
| `pg_strom.enable_partitionwise_gpupreagg` | `on` | 将 GpuPreAgg 下推到分区 |
