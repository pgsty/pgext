## 用法

来源：

- [官方扩展控制文件](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/pjpg.control)
- [官方上游文档](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/README.md)

`pjpg` — 提供时区转换、日期边界与时间范围等通用 SQL 工具函数。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "pjpg";
SELECT extversion
FROM pg_extension
WHERE extname = 'pjpg';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
