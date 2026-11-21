## 用法

来源：

- [Official source](https://github.com/brauliobo/pg_sift/releases/tag/v5.1.0)
- [Official source](https://github.com/brauliobo/pg_sift/blob/v5.0.2/sql/pg_sift--5.0.1--5.0.2.sql)

- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.1.0/README.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.1.0/pg_sift.control)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.1.0/CHANGELOG.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.1.0/docs/USAGE.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.1.0/docs/UPGRADING.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.1.0/docs/SUPPORT.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.0.2/sql/pg_sift--5.0--5.0.1.sql)

`pg_sift` 5.1.0 提供 BM25 文本检索、压缩列索引与列式表访问方法。上游支持 PostgreSQL 17–19，首次安装需要超级用户。普通检索和列式分析无需预加载；通过逻辑复制发布列式表时，须另外遵循上游指南的预加载要求。

### 文本检索

```sql
CREATE EXTENSION pg_sift;
CREATE TABLE articles (id bigint PRIMARY KEY, body text);
INSERT INTO articles VALUES
  (1, 'PostgreSQL extensions add full text search'),
  (2, 'PostgreSQL indexing and query planning');
CREATE INDEX articles_search ON articles USING sift (body);
SELECT id, sift.score(ctid) AS score
FROM articles
WHERE body @@> 'postgresql AND search'
ORDER BY score DESC LIMIT 10;
```

`sift` 访问方法为单个文本列或表达式建立索引，支持布尔、短语、前缀、通配符、正则与模糊查询。应在扫描查询中使用 `sift.score(ctid)`。`sift.tokenize`、`sift.highlight` 和 `sift.facets` 提供分析、高亮与分面；`sift.sections` 和 `sift.field_query` 实现文档所述的字段权重流程。

### 列式分析

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, category text, amount numeric(10,2));
CREATE INDEX orders_col ON orders USING sift_col (category, amount);
SELECT category, count(*), sum(amount) FROM orders GROUP BY category;
CREATE TABLE events (id bigint, kind text, amount numeric(10,2)) USING sift_table;
```

`sift_col` 为普通表增加索引，`sift_table` 则直接按列存储表。符合条件的计划使用 `Custom Scan (SiftAgg)`，其他查询结构保留 PostgreSQL 普通计划。应检查实际计划。`sift.enable_agg = off` 关闭聚合下推，`sift.enable_col_scan = off` 关闭列扫描。可选 Vulkan 路径仍属实验功能，默认关闭。

符合条件的新建或重建列索引使用 `track_changes = 'auto'`，会安装语句级触发器并增加写入工作。跟踪不完整时会回退到堆表检查。变更分区拓扑或在禁用触发器的情况下恢复数据前，应阅读升级指南。

### 安全与维护

受行级安全约束的角色默认不使用排名扫描。超级用户可启用 `sift.rls_scores = on`；返回行遵循 RLS，但评分依赖包括隐藏行在内的整个索引统计信息，应评估这一泄露边界。`sift.verify`、`sift.segments`、`sift.merge` 等维护函数要求表所有权或 MAINTAIN 权限。修改分析器或评分选项可能需要重建检索索引。

5.0.1 停止自动完整重写 `sift_col` 索引。`sift.col_maintenance_status` 与 `sift.col_maintenance_due` 提供重建建议，由操作者选择 `REINDEX INDEX CONCURRENTLY` 或安排 `sift.col_maintain(index)`。增量写缓冲提升、小段合并、日志折叠、清理与值存储维护仍自动进行，`sift.col_auto_merge` 继续控制合并。应移除废弃的 `sift.col_auto_merge_max_active` 与 `sift.col_auto_merge_max_workers` 设置。上游仍未证明部分构建与合并内存范围有界，不能假设存在严格内存上限。

默认 `sift.on_corruption = quarantine` 会使发现损坏的语句失败，并将受损的 `sift` 或 `sift_col` 索引退出服务。之后查询采用准确的堆表结果，旧计划可能一次性报 `40001`。应检查 `sift.quarantine` 并主动安排重建。该机制不能修复 PostgreSQL 系统目录，也不涵盖 `sift_table` 损坏。

### 升级边界

运行对应版本的预检，为包括备库在内的全部读取节点安装匹配动态库。先升级备库，以匹配文件重启服务器并重连连接池客户端，再升级各数据库：

```sql
ALTER EXTENSION pg_sift UPDATE TO '5.1';
SELECT * FROM sift.upgrade_commands();
SELECT * FROM sift.col_repair('public.orders_col', true);
```

5.0 到 5.0.1 的 SQL 脚本为空，只对齐报告的扩展版本，不改变对象。本补丁修复并发构建锁、内存保留、取消与崩溃路径，不改变磁盘格式，也不增加 REINDEX 要求。

跨越至 5.0 时仍须完整执行升级指南：先将读取端至少升至 4.3.6，并在主库写入新格式前升级备库。4.x 读取端无法解释 5.0 索引数据。`sift_table` 逻辑变化记录使用 WAL resource-manager ID 152；5.0 同时读取旧 ID 128，但 4.x 无法回放 152。回退要求遵循干净关闭及格式重建流程；格式不匹配本身不能证明损坏。

旧版修复义务继续适用：除 `analyzer = keyword` 外，3.3 之前的文本索引需为长词匹配正确性重建；受影响的 4.2.1–4.3.5 列段若存在存活位图损坏，可能需要 `sift.col_repair` 或 REINDEX。先执行上面的 dry-run，并发重建命令须在事务块外逐条运行。移除废弃的 `sift.enable_col` 和 `sift.table_auto_defrag_dead_pct`；`sift.eval_methods` 仅保留 `penalty_maintenance`。选项重命名及读取端格式限制应按当前指南处理，不可直接原地降级动态库。

### 版本 5.0.2

5.0.2 终止可能让索引持续增长直至磁盘写满的重复列合并，修复规划器及退出阶段崩溃，并避免 VACUUM 与插入并行时误报重复行。维护检查器拒绝分区父索引，应检查实际拥有存储的子索引。

从持久配置中移除 `sift.col_build_key_ranges`，此选项已删除。键范围构建会预估临时空间；若超过 `temp_file_limit` 或临时表空间剩余容量的一半，会以 SQLSTATE 54000 失败。重试前应检查 `temp_tablespaces` 和可用容量。

新库及空的 5.0.1 到 5.0.2 更新脚本不改变磁盘格式，健康的 5.0 索引无需例行 REINDEX。失控合并造成的索引膨胀须主动重建才能回收；`REINDEX INDEX CONCURRENTLY` 在切换前需要同时容纳两份索引。安装匹配文件、重启或回收所有加载旧库的会话，再更新各数据库中的 SQL 扩展。此前 4.x 的迁移与修复要求继续适用。

### 版本 5.1.0

发行版 5.1.0 使用 SQL 扩展版本 5.1。已有 5.0 索引可继续读取，无需强制 REINDEX。新写入可能引入删除记录、键范围写缓冲与链式目录，5.0 会拒绝读取这些格式。为采用新布局而重建属于可选操作，并非升级兼容性的前提。

`sift.col_build_source` 默认为 auto，在堆表、合适的 btree 与被替换的旧索引中选择，heap 则强制使用旧策略。键范围临时文件空间不足时，构建可改用合适的 btree 或开放单元，不再总是失败。`sift.col_wbuf_key_range` 默认启用键范围，由超级用户控制。`sift.backend_cache_kb` 的默认值改为每后端 128 MB。

降级前检查 `sift.upgrade_check('5.0')`，按指南依次执行 VACUUM、SQL 版本回退、停库和旧库安装，再使用旧版本重建列出的索引。在 5.1 下 REINDEX 可能再次写入不兼容格式。从受影响旧版本保留至今的索引仍须履行此前的修复要求。
