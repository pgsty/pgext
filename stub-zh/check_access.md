## 用法

来源：

- [官方扩展控制文件](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/check_access.control)
- [官方上游文档](https://github.com/CrunchyData/crunchy_check_access/blob/c8f01b40d7507715c4f0fda18a651916b010653a/README.md)

`check_access` — 用于检查 PostgreSQL 对象访问权限的函数与视图。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "check_access";
SELECT extversion
FROM pg_extension
WHERE extname = 'check_access';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
