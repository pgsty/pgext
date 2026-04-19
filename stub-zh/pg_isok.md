## 用法

来源: [official repo](https://codeberg.org/kop/pg_isok), [official docs home](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/index.html), [official reference source](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/isok.xml)

`pg_isok` 是一个基于查询的数据完整性与监控扩展。它不会只报告当前看起来可疑的行，而是保存先前结果，并在后续运行时聚焦于尚未解决或尚未延期的变化。

```sql
CREATE SCHEMA isok;
CREATE EXTENSION pg_isok SCHEMA isok;

SELECT *
FROM isok.run_isok_queries()
AS problems;
```

### 核心对象

- `ISOK_QUERIES`：存储监控查询及其执行设置。
- `ISOK_RESULTS`：存储被报告的行，以及它们是否已解决或已延期。
- `run_isok_queries()`：运行所有启用的检查。
- `run_isok_queries($$VALUES ('check_name')$$)`：只运行指定的检查。

### 典型流程

运行一个具名检查：

```sql
SELECT *
FROM isok.run_isok_queries($$VALUES ('new_countries')$$)
AS problems;
```

通过更新 `ISOK_RESULTS` 接受或推迟一个已知告警：

```sql
UPDATE isok.isok_results
SET deferred_to = 'infinity'
WHERE iqname = 'new_countries';
```

当某个条件不再需要关注时使用 `resolved`；当它应该隐藏到将来某个时间时使用 `deferred_to`。

### 适用场景

- 导入后的数据清理
- 监控不常见但有时可接受的模式
- “软触发器”式的审查流程，即硬约束过于严格时的替代方案

### 注意事项

- 上游建议将它安装在独立 schema 中，并在调用时显式带上 schema。
- 文档将其描述为纯 SQL 扩展，这在受管 PostgreSQL 服务限制 C 扩展时很有价值。
- 本仓库的包元数据显示 `superuser=false`，但上游并未把它记录为 trusted extension；应保守对待安装权限。
