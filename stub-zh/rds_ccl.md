## 用法

来源：

- [阿里云 rds_ccl SQL 限流文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-rds-ccl-extension-to-perform-sql-throttling-on-an-apsaradb-rds-for-postgresql-instance)
- [阿里云 RDS 产品概览](https://www.alibabacloud.com/help/en/rds/product-overview/what-is-apsaradb-rds-what-is-apsaradb-rds)

`rds_ccl` 是阿里云 ApsaraDB RDS for PostgreSQL 的 SQL 限流扩展，按规范化 `query_id` 匹配语句。规则限制并发并让超量执行排队，用于保护实例免受突发或重复高开销查询影响。

```sql
CREATE EXTENSION rds_ccl;
SELECT * FROM rds_show_current_db_ccl_rule();
```

使用前需启用实例参数 `rds_enable_ccl=on`，并把 `compute_query_id=auto` 或 `on`，然后用高权限 RDS 账号创建规则。重启或主备切换后必须重新加载规则。并发或队列上限过低会阻塞正常流量，因此应从实测 query ID 制定规则，逐步上线、监控队列，并保留立即禁用路径。
