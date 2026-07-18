## 用法

来源：

- [官方扩展控制文件](https://github.com/cybertec-postgresql/generic-plan/blob/f998a7eefc298a1e77f98e46464e20d96a2c1c1a/generic_plan.control)
- [官方上游文档](https://github.com/cybertec-postgresql/generic-plan/blob/f998a7eefc298a1e77f98e46464e20d96a2c1c1a/README.md)

`generic_plan` — 为带参数 SQL 语句生成不依赖参数值的通用执行计划。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "generic_plan";
SELECT extversion
FROM pg_extension
WHERE extname = 'generic_plan';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
