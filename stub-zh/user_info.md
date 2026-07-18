## 用法

来源：

- [官方上游文档](https://github.com/tvondra/user_info/blob/338fccb497543cf7321b020da0667b225327f8ea/doc/user_info.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/user_info/)

`user_info` — user_info：管理辅助类 PostgreSQL 扩展，提供用户拥有对象和授予角色等信息。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "user_info";
SELECT extversion
FROM pg_extension
WHERE extname = 'user_info';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
