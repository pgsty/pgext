

## 用法

- https://github.com/tensorchord/VectorChord
- 发布博客：[VectorChord: 在 PostgreSQL 中以 1 美元存储 40 万向量](https://blog.pgvecto.rs/vectorchord-store-400k-vectors-for-1-in-postgresql)

将此扩展添加到 postgresql.conf 中的 shared_preload_libraries 配置项

```sql
CREATE EXTENSION vchord CASCADE;
```

在 embedding 列上创建索引：

```sql
CREATE INDEX ON gist_train USING vchordrq (embedding vector_l2_ops) WITH (options = $$
residual_quantization = true
[build.internal]
lists = [4096]
spherical_centroids = false
build_threads = 8
$$);
```

--------

## 文档

### 查询

查询语句与 pgvector 完全相同。VectorChord 支持任意过滤操作以及 WHERE/JOIN 子句，就像 pgvecto.rs 的 VBASE 一样。

```sql
SELECT * FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

支持的距离函数包括：

- `<->` - L2 距离
- `<#>` - （负）内积
- `<=>` - 余弦距离


> 由于 PostgreSQL 查询规划器的限制，我们无法直接支持类似 `SELECT embedding <-> '[3,1,2]' as distance WHERE distance < 0.1 ORDER BY distance` 的范围查询。

要查询特定距离范围内的向量，可以使用以下语法。

```sql
-- 查询特定距离范围内的向量
-- sphere(center, radius) 表示以 center 为圆心、radius 为半径的球体内的向量，即范围查询
-- <<->> 是 L2 距离，<<#>> 是内积，<<=>> 是余弦距离
SELECT vec FROM t WHERE vec <<->> sphere('[0.24, 0.24, 0.24]'::vector, 0.012)
```

### 查询性能调优

可以通过调整 `probes` 和 `epsilon` 参数来优化搜索性能：

```sql
-- 设置 probes 以控制扫描的列表数量。
-- 建议范围：总 `lists` 值的 3%–10%。
SET vchordrq.probes = 100;

-- 设置 epsilon 以控制重排序精度。
-- 值越小重排序越少速度越快，值越大召回率越高。
-- 建议范围：0.0–4.0。默认值为 1.9。
SET vchordrq.epsilon = 1.9;
```

以及 PostgreSQL 本身的设置：
```sql
-- 如果使用 SSD，将 `effective_io_concurrency` 设置为 200 以获得更快的磁盘 I/O。
SET effective_io_concurrency = 200;

-- 禁用 JIT（即时编译），因为它收益甚微（1–2%），
-- 且对单查询负载会增加额外开销。
SET jit = off;

-- 至少将总内存的 25% 分配给 `shared_buffers`。
-- 对于磁盘密集型负载，可以将其增加到总内存的 90%。使用网络存储时，建议禁用交换分区以避免 I/O 挂起。
-- 注意：此设置需要重启才能生效。
ALTER SYSTEM SET shared_buffers = '8GB';
```

### 索引预热

要预热索引，可以使用以下 SQL。在内存有限的情况下，这将显著提升性能。

```sql
-- vchordrq_prewarm(index_name::regclass) 将索引预热到共享缓冲区中
SELECT vchordrq_prewarm('gist_train_embedding_idx'::regclass);
```


### 索引构建时间

索引构建可以通过索引选项中的 `build_threads` 参数和 PostgreSQL 设置来实现并行化。使用以下设置优化并行度：

```sql
-- 将此值设置为可用于并行操作的 CPU 核心数。
SET max_parallel_maintenance_workers = 8;
SET max_parallel_workers = 8;

-- 调整工作进程总数。
-- 注意：此设置需要重启才能生效。
ALTER SYSTEM SET max_worker_processes = 8;
```

### 索引构建进度


可以通过查询 `pg_stat_progress_create_index` 视图来查看索引构建进度。

```sql
SELECT phase, round(100.0 * blocks_done / nullif(blocks_total, 0), 1) AS "%" FROM pg_stat_progress_create_index;
```

### 外部索引预计算

与纯 SQL 方式不同，外部索引预计算会先在外部进行聚类，然后将质心插入 PostgreSQL 表中。虽然过程更为复杂，但在大规模数据集（>500 万条）上，外部构建明显更快。

首先，需要使用 `faiss`、`scikit-learn` 或其他聚类库对向量进行聚类。

质心需要预先存储在一个任意名称的表中，该表包含 3 列：
- id（integer）：每个质心的 ID，必须唯一
- parent（integer，可为空）：每个质心的父 ID，普通聚类时应为 NULL
- vector（vector）：每个质心的向量表示，使用 `pgvector` 的 vector 类型

示例如下：

```sql
-- 创建质心表
CREATE TABLE public.centroids (id integer NOT NULL UNIQUE, parent integer, vector vector(768));
-- 插入质心数据
INSERT INTO public.centroids (id, parent, vector) VALUES (1, NULL, '{0.1, 0.2, 0.3, ..., 0.768}');
INSERT INTO public.centroids (id, parent, vector) VALUES (2, NULL, '{0.4, 0.5, 0.6, ..., 0.768}');
INSERT INTO public.centroids (id, parent, vector) VALUES (3, NULL, '{0.7, 0.8, 0.9, ..., 0.768}');
-- ...

-- 使用外部质心表创建索引
CREATE INDEX ON gist_train USING vchordrq (embedding vector_l2_ops) WITH (options = $$
[build.external]
table = 'public.centroids'
$$);
```

为了简化工作流程，我们提供了端到端的外部索引预计算脚本，详见 [scripts](./scripts/README.md#run-external-index-precomputation-toolkit)。



------

## 限制

- 架构兼容性：快速扫描内核针对 x86_64 架构进行了优化。虽然可以在 aarch64 上运行，但性能可能较低。


------

## 构建

构建此扩展需要 [clang-17+](https://github.com/tensorchord/VectorChord/issues/188)。

在 EL 8/9、Ubuntu 24.04 上可直接使用，但在 Ubuntu 22.04 / Debian 12 上需要手动安装。

例如，在 Ubuntu 22 / Debian 12 上安装 clang-18 并将其设置为默认的 clang：

```bash
curl --proto '=https' --tlsv1.2 -sSf https://apt.llvm.org/llvm.sh | bash -s -- 18
sudo update-alternatives --install /usr/bin/clang clang $(which clang-18) 255
```
