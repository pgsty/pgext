## 用法

来源：

- [官方扩展控制文件](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/auth0.control)
- [官方上游文档](https://github.com/brunoenten/pg_auth0/blob/c6831e8eb821384fd5ef9950e31bfe35450913e7/README.md)

`auth0` — 在 PostgreSQL 中调用 Auth0 Management API 的扩展，提供用户、角色等管理函数。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`http`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "auth0";
SELECT extversion
FROM pg_extension
WHERE extname = 'auth0';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
