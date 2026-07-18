## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/range_partitioning/)

`range_partitioning` — 为 PostgreSQL-XL 管理基于触发器的静态范围分区，包括分区拆分与合并。

已复核目录快照记录的版本为 `1.2.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "range_partitioning";
SELECT extversion
FROM pg_extension
WHERE extname = 'range_partitioning';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
