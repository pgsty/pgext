## 用法

来源：

- [官方资料](https://github.com/brauliobo/pg_sift/blob/3e7fbf4aab7e4b1042ac83b106677dbc7377d501/README.md)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/3e7fbf4aab7e4b1042ac83b106677dbc7377d501/CHANGELOG.md)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/3e7fbf4aab7e4b1042ac83b106677dbc7377d501/pg_sift.control)
- [官方资料](https://github.com/brauliobo/pg_sift/blob/3e7fbf4aab7e4b1042ac83b106677dbc7377d501/sql/pg_sift--3.2.sql)

`pg_sift` 发行版 3.2.2 安装的 SQL 扩展版本为 3.2，同时提供 BM25 文本检索、压缩列索引与列式表访问方法。支持 PostgreSQL 17–18，PG19 仍为实验支持。首次安装需要超级用户，但无需预加载或重启服务器。

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

`sift` 索引处理文本列或表达式，支持布尔、短语、前缀、通配符、正则与模糊查询。`sift.score(ctid)` 可用于扫描查询的输出或排序；3.2 也支持在过滤条件中与常量分数阈值比较。`sift.tokenize`、`sift.highlight` 与 `sift.facets` 分别提供分析、高亮和分面能力。`sift.sections` 与 `sift.field_query` 支持文档所述的字段权重流程，但不代表可以创建多列 BM25 索引。

### 列式分析

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, category text, amount numeric(10,2));
CREATE INDEX orders_col ON orders USING sift_col (category, amount);
SELECT category, count(*), sum(amount) FROM orders GROUP BY category;
CREATE TABLE events (id bigint, kind text, amount numeric(10,2)) USING sift_table;
```

`sift_col` 为普通表增加索引，`sift_table` 则直接以列格式存储表。符合条件的聚合与过滤使用 `Custom Scan (SiftAgg)`，其他情况保留普通执行计划，应检查实际计划。`sift.enable_agg = off` 可关闭聚合下推以作比较。3.1 扩展了日历分组、支持的 JSON 结构与元素汇总。可选 Vulkan 路径仍属实验功能，默认关闭，上游软件包也未启用。

### 安全与维护

受行级安全或安全屏障视图约束的角色，不能对相关关系使用索引全局 BM25 评分。旧分面接口拒绝 RLS 表，新的 SQL 分面接口有自己的限制。`sift.verify`、`sift.segments`、`sift.merge` 等维护操作需要表所有权或 MAINTAIN 权限。修改分析器或评分选项需要重建受影响的检索索引。

`write_buffer_limit` 替代旧名 `mutable_limit`，旧名保留为别名。清理操作管理无效行和索引段，备库回放索引截断时可能取消查询。原型测试配套模块不是生产扩展。编码、数据块大小、支持的列类型与访问方法限制应以对应版本文档为准。

### 升级到 3.2

安装匹配文件后，更新每个使用库：

```sql
ALTER EXTENSION pg_sift UPDATE TO '3.2';
```

已发布的 2.x/3.0 数据结构仍可读取，无需例行重建。新写入会采用新版 BM25、列索引和表格式，旧库会拒绝打开。在首次写入新格式之前，必须升级所有访问这些数据的服务器，包括备库。降级需要明确的重写与重建流程，没有反向 SQL 升级脚本。已加载旧库的会话应重新连接。若跨越较早的 1.x 版本，还应检查更新日志中已删除的短语索引选项，以及函数和视图依赖变更。

### 3.2.2 修复与已知限制

3.2.2 修复了分区或继承子表在行级安全下通过执行计划行数估计泄露隐藏数据统计的问题，影响上游 3.0.0–3.2.1。修复只需更换库，无需为此重建索引。3.2 增加聚合下推与列类型支持；新增存储格式可能不被旧库读取，升级前需检查所有读取这些数据的节点。

3.2.0–3.2.2 的事件触发器还有已知权限缺陷：不是超级用户或扩展所有者的角色可能无法执行 DROP。上游给出的临时处理是由超级用户在每个数据库执行以下语句；从备份恢复后需重新检查。不要把尚未安装的后续 3.3 版本当作已修复。

```sql
ALTER FUNCTION sift.table_paths_drop() SECURITY DEFINER SET search_path = pg_catalog, pg_temp;
```
