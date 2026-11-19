## 用法

来源：

- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.0.1/README.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.0.1/pg_sift.control)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.0.1/CHANGELOG.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.0.1/docs/USAGE.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.0.1/docs/UPGRADING.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.0.1/docs/SUPPORT.md)
- [当前版本资料](https://github.com/brauliobo/pg_sift/blob/v5.0.1/sql/pg_sift--5.0--5.0.1.sql)

`pg_sift` 5.0.1 提供 BM25 文本检索、压缩列索引与列式表访问方法。上游支持 PostgreSQL 17–19，首次安装需要超级用户。普通检索和列式分析无需预加载；通过逻辑复制发布列式表时，须另外遵循上游指南的预加载要求。

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

运行对应版本的预检，为包括备库在内的全部读取节点安装匹配动态库。按要求重启或重连后，再升级各数据库：

```sql
ALTER EXTENSION pg_sift UPDATE TO '5.0.1';
SELECT * FROM sift.upgrade_commands();
SELECT * FROM sift.col_repair('public.orders_col', true);
```

5.0 到 5.0.1 的 SQL 脚本为空，只对齐报告的扩展版本，不改变对象。本补丁修复并发构建锁、内存保留、取消与崩溃路径，不改变磁盘格式，也不增加 REINDEX 要求。

跨越至 5.0 时仍须完整执行升级指南：先将读取端至少升至 4.3.6，并在主库写入新格式前升级备库。4.x 读取端无法解释 5.0 索引数据。`sift_table` 逻辑变化记录使用 WAL resource-manager ID 152；5.0 同时读取旧 ID 128，但 4.x 无法回放 152。回退要求遵循干净关闭及格式重建流程；格式不匹配本身不能证明损坏。

旧版修复义务继续适用：除 `analyzer = keyword` 外，3.3 之前的文本索引需为长词匹配正确性重建；受影响的 4.2.1–4.3.5 列段若存在存活位图损坏，可能需要 `sift.col_repair` 或 REINDEX。先执行上面的 dry-run，并发重建命令须在事务块外逐条运行。移除废弃的 `sift.enable_col` 和 `sift.table_auto_defrag_dead_pct`；`sift.eval_methods` 仅保留 `penalty_maintenance`。选项重命名及读取端格式限制应按当前指南处理，不可直接原地降级动态库。
