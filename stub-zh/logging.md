## 用法

来源：

- [官方扩展控制文件](https://github.com/snaga/xlogging/blob/fe28bf2c8ef71e214e2472fb6e17310ca5cfd017/logging.control)

`logging` — 切换表及其索引的 WAL 日志模式。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "logging";
SELECT extversion
FROM pg_extension
WHERE extname = 'logging';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
