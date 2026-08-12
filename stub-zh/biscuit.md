## 用法

来源：

- [PGXN 上的 Biscuit 3.0.0](https://pgxn.org/dist/biscuit/3.0.0/)
- [Biscuit 3.0.0 发行说明](https://github.com/CrystallineCore/Biscuit/releases/tag/v3.0.0)
- [Biscuit 3.0.0 README](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/README.md)
- [Biscuit 3.0.0 变更日志](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/CHANGELOG.md)
- [Biscuit 3.0.0 元数据](https://api.pgxn.org/dist/biscuit/3.0.0/META.json)
- [Biscuit 3.0.0 控制文件](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/biscuit.control)
- [Biscuit 3.0.0 Makefile](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/Makefile)
- [Biscuit 3.0.0 安装 SQL](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/sql/biscuit.sql)
- [Biscuit 2.5.0 至 3.0.0 升级 SQL](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/sql/biscuit--2.5.0--3.0.0.sql)

`biscuit` 3.0.0 是面向 PostgreSQL 16 及以上版本的定位位图索引访问方法，用于精确执行 `LIKE` 与 `ILIKE` 过滤。它尤其适合锚定模式、`_` 通配符、长度谓词和多列合取条件。3.0.0 把索引状态保存在有 WAL 日志的关系页面中，因此崩溃恢复、时间点恢复、物理复制和热备读取都使用 PostgreSQL 的常规恢复路径。它不需要 `shared_preload_libraries`，也无需重启。

项目仍处于积极开发阶段，并建议使用有代表性的负载进行预发布测试。其每连接内存、写放大和缓存重载特性更适合读多写少的分析负载，而不适合持续更新的 OLTP 表或超大连接池。

### 创建并查询索引

应先装载数据，再创建索引。默认 `biscuit_ops` 同时支持区分与不区分大小写的谓词。只需要一种模式时，应使用 `biscuit_like_ops` 或 `biscuit_ilike_ops`，避免构建不使用的结构集合。

```sql
CREATE EXTENSION biscuit;

CREATE INDEX message_body_biscuit_idx
ON message USING biscuit (body biscuit_like_ops);

ANALYZE message;

EXPLAIN (ANALYZE, BUFFERS)
SELECT id, body
FROM message
WHERE body LIKE 'timeout%';
```

它支持表达式索引和多列索引。查询必须使用与所选操作符类兼容的表达式和操作符。装载统计信息后，应检查有代表性的执行计划，尤其是非锚定模式。

### 操作符类与查询边界

- `biscuit_ops` 是默认文本操作符类，索引 `LIKE`、`NOT LIKE`、`ILIKE` 和 `NOT ILIKE`。
- `biscuit_like_ops` 只索引 `LIKE` 和 `NOT LIKE`。
- `biscuit_ilike_ops` 只索引 `ILIKE` 和 `NOT ILIKE`。

Biscuit 返回无需堆表复查的精确结果，但它是过滤索引：不提供有序、反向、仅索引或唯一扫描，不能用于 `CLUSTER`，也不支持正则表达式、相似度搜索、模糊搜索或区域设置感知的排序规则。对选择性前缀查询，带 `text_pattern_ops` 的 B-tree 通常更合适；`pg_trgm` 则专用于非锚定子串、正则表达式和相似度搜索。

### 诊断与配置

重要的检查对象包括 `biscuit_indexes`、`biscuit_status`、`biscuit_index_stats(oid)`、`biscuit_index_memory_size()`、`biscuit_pending_list_stats(oid)` 与 `biscuit_pending_list_usage`。内存函数报告当前后端的会话本地副本。`total_pending_bytes` 在 `VACUUM` 时刷新，因此待处理列表数值最多可能比实时写入落后一个清理周期。

- `biscuit.delta_compaction_slots` 默认为 20000，控制压缩前允许积累的待处理行数。提高它会增加其他会话的重载工作，因此这是受权限控制的设置。
- `biscuit.diag_scan_trace` 默认为关闭，会发出详细的逐扫描候选集统计。只应在聚焦复现问题时启用。

每个后端都会惰性加载自己的索引副本，并在连接生命周期内保留。任一已提交写入都会使其他缓存副本失效；下次访问会重载整个索引，而不是增量刷新。连接池容量必须计入这种内存行为，并应避免把频繁写入与延迟敏感的读取交错在同一索引上。

对现有索引执行 `INSERT` 与 `UPDATE` 会产生大量 WAL；应监控 `pg_wal`、复制延迟与复制槽保留，并考虑设置有限的 `max_slot_wal_keep_size`。先批量装载再建索引的成本低得多。`VACUUM` 会排空待处理工作，但不会缩小索引；回收索引空间需要使用 `REINDEX`。

### 升级至 3.0.0

3.0.0 是不兼容的磁盘格式变更。更新扩展目录不会转换现有索引页面：所有由 2.x 创建的 Biscuit 索引都必须重建。应为重建预留足够的维护时间和 WAL 容量。

```sql
ALTER EXTENSION biscuit UPDATE TO '3.0.0';

SELECT schema_name, index_name
FROM biscuit_indexes;

REINDEX INDEX CONCURRENTLY public.message_body_biscuit_idx;
```

未经修补的上游 3.0.0 归档只携带并安装 `2.5.0--3.0.0` 这一步，而早期稳定软件包暴露的目录版本为 `2.4.0` 或 `2.4.1`。Pigsty 的 3.0.0 RPM 与 DEB 软件包会先恢复缺失的目录升级路径，再应用上游步骤。使用其他源码构建或软件包时，应在 `ALTER EXTENSION` 前检查 `pg_extension_update_paths('biscuit')`；无论 SQL 路径是否可用，强制要求的 `REINDEX` 或 `REINDEX CONCURRENTLY` 都仍是独立的手工操作。
