## 用法

来源：

- [官方资料](https://github.com/brauliobo/pg_sift/blob/288643904639a2834a1cf92c6cfa7d4bbcaf41bf/README.md)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/288643904639a2834a1cf92c6cfa7d4bbcaf41bf/docs/UPGRADING.md)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/288643904639a2834a1cf92c6cfa7d4bbcaf41bf/pg_sift.control)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/288643904639a2834a1cf92c6cfa7d4bbcaf41bf/sql/pg_sift--3.4.sql)
- [用法与安全](https://github.com/brauliobo/pg_sift/blob/288643904639a2834a1cf92c6cfa7d4bbcaf41bf/docs/USAGE.md)
- [支持版本](https://github.com/brauliobo/pg_sift/blob/288643904639a2834a1cf92c6cfa7d4bbcaf41bf/docs/SUPPORT.md)

`pg_sift` 发行版 3.4.2 安装的 SQL 扩展版本为 3.4，同时提供 BM25 文本检索、压缩列索引与列式表访问方法。上游支持 PostgreSQL 17–19，但这不代表每个大版本都有对应的软件包。首次安装需要超级用户，普通用法不强制要求预加载。

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

`sift` 索引处理文本列或表达式，支持布尔、短语、前缀、通配符、正则与模糊查询。`sift.score(ctid)` 可用于扫描查询的输出或排序；也支持在过滤条件中与常量分数阈值比较。`sift.tokenize`、`sift.highlight` 与 `sift.facets` 分别提供分析、高亮和分面能力。`sift.sections` 与 `sift.field_query` 支持文档所述的字段权重流程，但不代表可以创建多列 BM25 索引。

### 列式分析

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, category text, amount numeric(10,2));
CREATE INDEX orders_col ON orders USING sift_col (category, amount);
SELECT category, count(*), sum(amount) FROM orders GROUP BY category;
CREATE TABLE events (id bigint, kind text, amount numeric(10,2)) USING sift_table;
```

`sift_col` 为普通表增加索引，`sift_table` 则直接以列格式存储表。符合条件的聚合与过滤使用 `Custom Scan (SiftAgg)`，其他情况保留普通执行计划，应检查实际计划。`sift.enable_agg = off` 可关闭聚合下推以作比较。3.3 扩展了 top-N 排序、点查与小表连接，新建列索引默认启用预聚合。可选 Vulkan 路径仍属实验功能，默认关闭，上游软件包也未启用。

### 安全与维护

受行级安全策略约束的角色默认不使用排名扫描。超级用户可通过 `sift.rls_scores = on` 显式启用；返回行仍遵循策略，但评分依赖整个索引的统计信息，包括隐藏行。启用前应评估这一信息边界。安全屏障及其他不支持的查询结构仍受文档所述限制。

3.3 的安装与升级会将 `sift` 模式的使用权限授予 `PUBLIC`，各函数仍检查自身的表或索引访问权限。`sift.verify`、`sift.segments`、`sift.merge` 等维护操作需要表所有权或 MAINTAIN 权限。修改分析器或评分选项需要重建受影响的检索索引。

`sift.col_auto_merge` 与 `sift.table_auto_defrag` 现在默认开启。后台维护可能清理表并重写列索引，从而产生 WAL。列合并预算由 `sift.col_auto_merge_wal_pct` 控制，`sift.col_maintenance_status` 与 `sift.col_index_status` 可查看维护状态。关闭这两个自动维护开关可恢复旧版行为。

`write_buffer_limit` 替代旧名 `mutable_limit`，旧名保留为别名。备库回放索引截断时可能取消查询。通过逻辑复制发布列式表需要遵循上游预加载配置；基本文本检索和列式分析不要求预加载。

### 从 3.3 之前版本升级

先运行上游升级预检查，安装匹配文件并更新每个数据库，再重启已加载旧动态库的进程，使所有读取方使用同一发布版本：

```sql
ALTER EXTENSION pg_sift UPDATE TO '3.4';
SELECT * FROM sift.upgrade_commands();
```

除使用 `analyzer = keyword` 的索引外，必须重建所有在 3.3 之前创建的文本索引：超过 255 字节的词的处理方式已改变，未重建的索引可能漏掉匹配行。升级过程会打印命令，并发索引重建应在事务块外逐条执行。建议的列索引重建可能改善性能，而建议的列式表重写可能持有排他锁；执行前需逐项评估。

3.3 默认的 `sift.col_format_minor = 15` 无法被 3.2 动态库读取。在写入新格式前，应升级所有读取方，包括备库。如果必须保留 3.2 兼容性，应遵循对应版本的降级指南，并在升级前设置服务器级格式上限；会话级限制无法影响后台工作进程。可选的点查 TID 目录需要格式 16。数据文件不会自动降级。

3.3 的 SQL 迁移修复了旧版 DROP 事件触发器权限缺陷，之前的手动修正不再是单独的安装步骤。它也处理新后端首条带绑定参数的查询，扩展了 3.2.3 对普通查询的修复。

### 3.4 升级与规划变化

发行版 3.4.2 使用 SQL 扩展版本 3.4；安装匹配文件后执行所提供的 3.3 到 3.4 SQL 升级。新版本加入外键辅助的列扫描，调整 IN 列表与扫描启动成本；不能证明等价或成本更高的查询仍使用 PostgreSQL 普通计划。修复涵盖聚合路径排序元数据、达到有符号 128 位上限的 numeric 求和和文本字典处理。已有索引及预加载进程应按该发行版升级指南处理，基准测试耗时不能作为实际负载保证。

### 3.4.2 维护修复

3.4.2 继续使用 SQL 扩展版本 3.4，修复在 `ALTER TABLE` 添加列后进行 `ORDER BY ... LIMIT` 时的崩溃，改进分组聚合及范围/IN 扫描，并修复字典与列维护的边界问题。所有读取方都应先安装匹配文件，再写入更新后的列格式；3.3 到 3.4 的升级边界保持不变。
