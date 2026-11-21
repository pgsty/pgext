## 用法

来源：

- [Official source](https://www.paradedb.com/docs/project/changelog/0.27.0.md)
- [Official source](https://api.pgxn.org/src/pg_search/pg_search-0.27.0/pg_search/pg_search.control)
- [Official source](https://api.pgxn.org/src/pg_search/pg_search-0.27.0/Cargo.toml)

- [0.26.8 变更日志](https://www.paradedb.com/docs/project/changelog/0.26.8.md)
- [pg_search 0.26.8 README](https://api.pgxn.org/src/pg_search/pg_search-0.26.8/pg_search/README.md)
- [pg_search 0.26.8 控制文件](https://api.pgxn.org/src/pg_search/pg_search-0.26.8/pg_search/pg_search.control)
- [0.26.7 到 0.26.8 SQL](https://api.pgxn.org/src/pg_search/pg_search-0.26.8/pg_search/sql/pg_search--0.26.7--0.26.8.sql)
- [Upstream source](https://api.pgxn.org/src/pg_search/pg_search-0.26.7/docs/project/changelog/0.26.7.mdx)
- [Upstream source](https://api.pgxn.org/src/pg_search/pg_search-0.26.7/pg_search/pg_search.control)

- [0.26.6 changelog](https://github.com/paradedb/paradedb/blob/v0.26.6/docs/project/changelog/0.26.6.mdx)
- [pg_search 0.26.6 control](https://github.com/paradedb/paradedb/blob/v0.26.6/pg_search/pg_search.control)
- [pg_search 0.26.6 Cargo.toml](https://github.com/paradedb/paradedb/blob/v0.26.6/pg_search/Cargo.toml)

- [docs/project/changelog/0.26.2.mdx](https://api.pgxn.org/src/pg_search/pg_search-0.26.2/docs/project/changelog/0.26.2.mdx)
- [pg_search/sql/pg_search--0.26.1--0.26.2.sql](https://api.pgxn.org/src/pg_search/pg_search-0.26.2/pg_search/sql/pg_search--0.26.1--0.26.2.sql)
- [pg_search/Cargo.toml](https://api.pgxn.org/src/pg_search/pg_search-0.26.2/pg_search/Cargo.toml)

- [Release 0.26.1](https://github.com/paradedb/paradedb/releases/tag/v0.26.1)
- [Control](https://github.com/paradedb/paradedb/blob/529c77e9dffe4bd86b6b7971c019bb26b8eadbfa/pg_search/pg_search.control)

- [0.26.0 migration notes](https://github.com/paradedb/paradedb/blob/aa7f9dd407018036d144a54a3f875c5a00d20cda/docs/project/changelog/0.26.0.mdx)
- [0.26.0 SQL upgrade](https://github.com/paradedb/paradedb/blob/aa7f9dd407018036d144a54a3f875c5a00d20cda/pg_search/sql/pg_search--0.25.11--0.26.0.sql)
- [pg_search v0.25.11 README](https://github.com/paradedb/paradedb/blob/v0.25.11/pg_search/README.md)
- [pg_search v0.25.11 发行说明](https://github.com/paradedb/paradedb/releases/tag/v0.25.11)
- [PGXN 0.25.11 元数据](https://api.pgxn.org/src/pg_search/pg_search-0.25.11/META.json)
- [pg_search v0.25.1 迁移说明](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/changelog/0.25.1.mdx)
- [创建 ParadeDB 索引](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/indexing/create-index.mdx)
- [全文匹配操作符](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/full-text/match.mdx)
- [BM25 评分](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/sorting/score.mdx)
- [高亮与摘要](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/full-text/highlight.mdx)
- [索引向量](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/indexing/indexing-vectors.mdx)
- [查询向量](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/vector/querying.mdx)
- [混合搜索概述](https://github.com/paradedb/paradedb/blob/v0.25.11/docs/documentation/hybrid/overview.mdx)

`pg_search` 0.27.0 为 PostgreSQL 增加 ParadeDB 的全文、结构化、向量和混合搜索索引。版本 0.25 使用 `paradedb` 索引访问方法；旧的 `bm25` 访问方法名称仍保留为兼容别名。该扩展依赖 `vector`，上游支持 PostgreSQL 15-18，且必须通过 `shared_preload_libraries` 加载。

### 安装并构建索引

```ini
shared_preload_libraries = 'pg_search'
```

重启 PostgreSQL，然后创建扩展以及一张拥有稳定唯一键的表：

```sql
CREATE EXTENSION pg_search CASCADE;

CREATE TABLE documents (
  id          bigint PRIMARY KEY,
  title       text,
  body        text,
  category    text,
  embedding   vector(768)
);

CREATE INDEX documents_search_idx ON documents
USING paradedb (
  id,
  title,
  body,
  category,
  embedding vector_cosine_ops
)
WITH (key_field = 'id');
```

`key_field` 必须是第一个索引列，并能唯一标识每一行。文本键必须以不分词的方式建立索引。一张表只能拥有一个 ParadeDB 索引，因此应将所有需要搜索的字段都纳入该索引。

### 全文搜索

使用 `|||` 匹配任意词元，使用 `&&&` 要求匹配所有词元：

```sql
SELECT id, title, pdb.score(id) AS score
FROM documents
WHERE body ||| 'postgresql search'
ORDER BY score DESC, id;

SELECT id, pdb.snippet(body) AS excerpt
FROM documents
WHERE body &&& 'postgresql indexing';
```

`pdb.score(key_field)` 提供当前行的相关性分数。`pdb.snippet(indexed_text_column)` 返回高亮摘要。这些辅助函数只有在 ParadeDB 搜索谓词驱动的查询中才有意义。

### 向量搜索

向量索引在 0.25 系列中处于 beta 阶段，使用 pgvector 的 `vector` 类型。创建索引时应选择操作符类；更改距离度量需要重建索引。

```sql
SELECT id, title, embedding <=> $1::vector AS distance
FROM documents
WHERE id @@@ pdb.all()
ORDER BY embedding <=> $1::vector, id
LIMIT 20;
```

支持的索引操作符类为 `vector_l2_ops`、`vector_ip_ops` 和 `vector_cosine_ops`。0.25 向量索引不为 `halfvec`、`sparsevec` 或 `bit` 列建立索引。

### 混合搜索

单个 ParadeDB 索引可以组合词法谓词、结构化过滤与向量排序。需要更复杂的融合时，应使用文档中的 RRF 和加权混合搜索函数，而不是直接将量纲不同的分数相加。

```sql
SELECT id, title, pdb.score(id) AS lexical_score
FROM documents
WHERE body ||| 'postgresql extension'
  AND category === 'database'
ORDER BY embedding <=> $1::vector, id
LIMIT 20;
```

### 版本 0.25.11 与注意事项

- 版本 0.25 将主要索引访问方法从 `bm25` 重命名为 `paradedb`。现有的 `USING bm25` 定义仍受支持，但新示例应使用 `USING paradedb`。
- 版本 0.25.1 支持确定性的向量并列结果排序，并将倒数排名融合查询的向量分支下推到索引中。它还新增 `paradedb.vector_clustering_threshold`，默认值为 500，并将向量索引构建并行度上限设为四个工作进程。
- 版本 0.25.1 移除了 `paradedb.vector_cluster_probe_epsilon`，并更改了向量索引的边界门控。从 0.25.0 升级数据库后，必须对所有包含向量字段的 ParadeDB 索引执行 `REINDEX`；对于这些索引，仅安装新的共享库并执行 `ALTER EXTENSION` 并不充分。
- 0.25.2 是稳定性与正确性版本：它修复带向量列的无字段 `more_like_this`、通用预备计划中的 `pdb.fuzzy`、遗留动态过滤器、多种并行子计划和 MPP 计划形态错误，并收紧 typemod 定义的访问控制。除了继承自 0.25.0 的向量索引重建要求外，没有新增索引迁移。
- 0.25.4 至 0.25.6 新增 `paradedb.vector_clusters`、分区感知索引构建、聚合分数连接、统一规划器告警控制与位图扫描交集；同时修复被丢弃的位图交集子节点，以及负高精度 `numeric` 值的可排序编码。即便没有记录新的显式索引迁移，升级后也应验证计划与排序。
- 0.25.7–0.25.10 扩展了感知分区的并发构建、数组横向展开、连接聚合及并行聚合能力。0.25.10 恢复了并行 top-K 规划，并修复 HOT 重定向导致的丢行、缺失 JSON 路径值以及无法准确表示默认值的聚合类型问题。升级后应检查结果正确性与查询计划。
- 0.25.11 修复取消或终止搜索扫描时的后端崩溃，并将 `STRING_AGG` 的空分隔符按空字符串处理。
- `CREATE EXTENSION pg_search CASCADE` 可以安装所需的 `vector` 扩展，但仍须先为所有服务器进程配置预加载并重启。仅通过 `LOAD` 或 `session_preload_libraries` 加载并不充分。
- 使用不同字段选项重建索引后，查询计划、分词和排名都可能变化。在上线前，请使用符合生产形态的数据测试相关性与向量召回率。

### 0.26.0 迁移

0.26.0 将向量存储改为格式 4。更新库与 SQL 扩展后，必须重建所有包含向量列的索引，包括关闭量化的索引。在完成重建前，向量查询会报需重建错误，写入仍可继续，并支持 `REINDEX CONCURRENTLY`。至少 64 维的向量默认启用量化，其设置仅在创建或重建索引时生效。保存了 `centroid_ratio` 或 `training_samples_per_centroid` 的索引需使用 `training_sample_ratio` 和 `max_leaf_size` 重新创建。含多个正则分词列的索引也应重建以修正模式选择。

此版本新增 beta 状态的 `partition_by` 索引分区、连接上的窗口聚合下推，以及可选的 `paradedb.spill_to_disk`。聚合的 `visibility` 默认是 `transaction`，`raw` 跳过 MVCC 检查，不能将其结果当作事务可见结果。应用上线前应复核排序、召回率、执行计划与存储需求。

### 0.26.1 版本

0.26.1 修复聚合与过滤结果、通用计划中的 RLS 行为及缓存处理。它未取消 0.26.0 的迁移要求：跨越该版本升级时，应重建每个含向量字段的索引，使用已移除向量选项的索引须重新创建。仍须保留 vector 依赖及服务器预加载、重启配置。

### 0.26.2 聚合与位图更新

`pdb.combine_agg(pdb.agg(...)) OVER ()` 可在 GROUP BY 分组间计算全局分面与聚合，聚合 FILTER 中的搜索操作符也可下推。新的 posting bitmaps 默认启用，用于符合条件的常见词 AND/OR 计数查询；字段选项 `bitmap_postings=false` 控制存储，`paradedb.enable_bitmap_postings = off` 可在查询时禁用而不重建索引。

此版本修复通用预备计划中的搜索操作符、并行扫描重复执行时的错误结果或崩溃，以及同步备库缺失时后台合并阻碍主库快速关闭的问题。范围过滤的得分修正也可能改变排名。安装匹配库后更新 SQL 扩展，并按预加载库的维护流程重启；跨越 0.26.0 时仍须遵循前述向量索引迁移要求。

### 0.26.3 行为修复

0.26.3 修复了整数范围边界、空关系、较大 LIMIT、MPP 连接分面，以及扫描、向量训练和 VACUUM 中的取消处理；索引构建和写入现在拒绝超出范围的日期。升级恢复旧范围分词器，并修正表情符号过滤的配置行为。连接扫描新增窗口函数 FILTER 与易变投影支持。检查依赖这些边界行为的查询，并按上游升级流程更新每个数据库中的 SQL 扩展。上游版本与当前已构建软件包版本应分别核对。

### 0.26.4 性能与稳定性

0.26.4 为数组列增加延迟物化，使符合条件的连接扫描在 `ORDER BY` 与 `LIMIT` 场景使用迭代式 rank-join，并避免 CJK Lindera 分词消耗呈二次增长。该版本修复分区表追加扫描中的错误文件路径、窗口函数场景及分区索引辅助函数问题。发行说明没有新增索引格式迁移；跨越 0.26.0 时仍须执行前述向量索引重建。

### 0.26.5 权限与正确性

0.26.5 为独立读取索引的函数实施表权限检查，并修复涉及 NULL 或空 ALL 数组的下推结果。它还改进并行 worker 清理与合并背压。应安装匹配文件，按预加载库维护流程重启，并逐库更新 SQL 扩展。此次发布不取消 0.26.0 的向量索引迁移要求。

### 0.26.6 版本

Join Scan 现在支持在 `INNER` 和 `SEMI`（`IN` / `EXISTS`）连接中对向量距离操作符 `<->`、`<=>`、`<#>` 执行 top-K `ORDER BY`。当叶子节点包含可由索引回答的单表过滤条件（包括行级安全策略、等值、`= ANY` 和范围条件）时，合并得分的 top-K 连接也可使用 rank join。本版本修复 Join Scan 选列同时包含 `pdb.score()` 与标量子查询或查询参数时的崩溃，并在未使用 `RankJoinExec` 时通过 `EXPLAIN` 给出警告。升级后应验证计划与 top-K 结果；软件包坐标需与上游版本分别核对。

### 0.26.7 分区与分数

数据倾斜的 partition_by 索引现在使用堆 ctid 回退拆分原本无法再切分的分区，已有索引需 REINDEX 才能获益。若显式把 ctid 加入 partition_by，仍沿用它与其他列竞争的旧行为，不启用该回退。

行级安全过滤现在始终对 pdb.score 贡献 0.0，移除了此前随计划变化的常数，因此报告分数可能改变。合并反压统计全部可见段，并修复单段包含超过 4 GiB 不同字符串值时的构建与合并错误。这些变化不取消早先 0.26 向量格式的索引重建要求。

### 0.26.8 前台合并行为

写入路径成功启动后台合并器后，现在会跳过前台合并。若无法启动后台合并器，仍回退到前台合并。按既有流程安装匹配库、更新 SQL 扩展并执行预加载维护；本版本没有新的索引格式迁移。

### 版本 0.27.0：升级与查询变化

先升级备用节点，再升级主节点。Community 物理副本现在可以继续重放 WAL，但只读副本上使用 ParadeDB 索引的查询会失败。副本提升为主节点，或从备用节点取得的基础备份恢复后，复制来的 ParadeDB 索引须执行 `REINDEX INDEX`。主节点取得的备份不受此限制。旧版备用节点库无法重放新的 ParadeDB WAL 记录。

所有现有向量索引都须重建：0.27.0 使用向量格式 6，拒绝旧格式。将 `vector_fields` 索引选项替换为 `quantization`；显式保存了已删除选项的索引须重新创建。通过 ALTER INDEX 或重建，将 `max_leaf_size` 改为 `target_leaf_size`。量化默认值为 `[1,1]`，零表示禁用。新的默认值 `replication = 2` 将向量额外分配到第二个簇，影响存储及构建工作量；`replication = 1` 可关闭额外分配。

整行的裸搜索现在检索已索引字段，不再隐式指向键字段。字段专用构造器须绑定已索引列，例如 `body @@@ pdb.term('postgres')`，不能使用整行 term 构造器。直接索引 ltree 已弃用；文档建议使用独立 GiST 索引并通过位图求交组合检索。

`paradedb.mpp_queue_size` 现表示字节预算，默认为 8MB；`paradedb.mpp_request_timeout` 已删除。`paradedb.vector_cluster_max_probe` 改为绝对工作预算，默认为 100。`paradedb.max_term_set_size` 默认 262144，超限词项集合会报错。部署前须复核已有配置。

新增 `pdb.expand` 在查询时应用同义词规则，`pdb.explain_score` 返回 JSONB 分数解释。L2 和内积向量的模超过约 6.5e18 时会被拒绝，包括 REINDEX 期间；有限值余弦向量仍被接受。上游预编译二进制要求 PostgreSQL 15.20、16.16、17.12 或 18.7，或同主版本内更高的小版本。这些上游要求并不表示目录中的 RPM/DEB 包已经更新。
