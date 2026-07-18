## 用法

来源：

- [官方扩展控制文件](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/pgext/uuid47/uuid47.control)
- [官方上游文档](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/README.md)
- [官方项目或服务商页面](https://uuidv47.stateless.me/)

`uuid47` — 以 UUIDv7 存储并通过密钥输出 UUIDv4 外观的类型

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "uuid47";
SELECT extversion
FROM pg_extension
WHERE extname = 'uuid47';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
