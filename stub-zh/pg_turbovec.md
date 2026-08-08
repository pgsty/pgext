## 用法

来源：

- [pg_turbovec v1.29.0 README](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/README.md)
- [pg_turbovec v1.29.0 变更日志](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/CHANGELOG.md)
- [pg_turbovec v1.29.0 控制文件](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/pg_turbovec.control)
- [分区扩展指南](https://codeberg.org/gregburd/pg_turbovec/src/tag/v1.29.0/docs/PARTITIONED_SCALE.md)
- [Pigsty 软件包矩阵](https://pgext.cloud/ext/pg_turbovec)

`pg_turbovec` 提供稠密的 `turbovec.vector` 类型和 `turbovec` 近似最近邻索引访问方法。它将浮点坐标量化为 2、3 或 4 位，并使用堆向量对候选项重新排序。适用于索引占用空间与搜索吞吐量比精确最近邻召回更重要的场景。

### 创建与查询向量

```sql
CREATE EXTENSION pg_turbovec;
SET search_path = public, turbovec;

CREATE TABLE items (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  embedding turbovec.vector
);

INSERT INTO items (embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT id, embedding <=> '[3,1,2]'::turbovec.vector AS cosine_distance
FROM items
ORDER BY embedding <=> '[3,1,2]'::turbovec.vector
LIMIT 10;
```

距离操作符包括：`<->` 表示 L2，`<#>` 表示负内积，`<=>` 表示余弦距离，`<+>` 表示 L1。当前索引支持按内积和余弦排序；L2 与 L1 只能执行精确运算。

### 构建 TurboVec 索引

```sql
CREATE INDEX items_embedding_idx ON items
USING turbovec (embedding vec_cosine_ops)
WITH (bit_width = 4);

SET turbovec.probes = 32;

SELECT id
FROM items
ORDER BY embedding <=> '[3,1,2]'::turbovec.vector
LIMIT 10;
```

将 `vec_cosine_ops` 与 `<=>` 搭配，或将 `vec_ip_ops` 与 `<#>` 搭配。`bit_width = 4` 是默认值，通常更有利于召回率；2 位索引更小，但需要针对具体工作负载测试召回率。支持 `CREATE INDEX CONCURRENTLY`。

重要的调优控制项包括 `turbovec.probes`、`turbovec.search_k`、`turbovec.oversample`、`turbovec.hi_dim_rerank`、`turbovec.iterative_scan` 和 `turbovec.cache_size_mb`。每次只改变一个维度，并将近似结果与精确基线进行比较。

### 过滤与分区

对于稳定的过滤值，使用 PostgreSQL 部分索引；对于显式候选允许列表，使用文档所述的 `turbovec.knn(..., allowed)` 接口；对于普通的带过滤条件 `ORDER BY ... LIMIT` 查询，则使用迭代扫描。

版本 1.29 记录了如何使用原生 PostgreSQL 分区来处理超出单表规模的数据集。父表查询可在各分区的 TurboVec 索引之间使用 `Merge Append`：

```sql
SELECT id
FROM partitioned_items
ORDER BY embedding <=> $1::turbovec.vector
LIMIT 20;
```

应分别为每个分区执行构建、vacuum 和 reindex。基于粗粒度向量量化器的分区裁剪在 1.29.0 中仅是设计方案，并非已经交付的功能。

### 版本与完整性边界

- 控制文件将对象安装到模式 `turbovec`，不可重定位，且不要求 `shared_preload_libraries` 或重启服务器。
- 上游 v1.29 面向 PostgreSQL 13-18，并将 PostgreSQL 19 支持标记为实验性；当前 Pigsty 软件包覆盖 PostgreSQL 14-18，并提供匹配的 OpenBLAS 链接二进制文件。
- 当前 Pigsty 软件包/源码元数据仍为 `1.28.3`，而上游版本为 `1.29.0`。在打包的 1.28.3 安装上，不要假定已经具备 1.28.4 完整性检查器或 1.29 功能。
- 上游 1.28.4 修复了持久化行数漂移问题，该问题可能破坏索引 ID 表；同时新增 `turbovec.turbovec_check(regclass)`。对于高写入生产场景，应优先将软件包至少升级到 1.28.4。已经损坏的索引仍须通过 `REINDEX` 或删除后重建来恢复。
- 从健康的 1.28.4 索引升级时，版本 1.29.0 仅包含增量变更，不要求重建索引。安装新文件后，执行 `ALTER EXTENSION pg_turbovec UPDATE TO '1.29.0'` 即可。
- 虽然 1.29 的 reloption 解析器接受 `bit_width = 1`，但端到端一位索引尚未实现，`CREATE INDEX` 会主动报错。请使用 2、3 或 4。
- 磁盘上的 ID 表在非正常关机后仍存在有文档记录的崩溃安全缺口。应认真处理完整性错误，并遵循上游恢复指南。
