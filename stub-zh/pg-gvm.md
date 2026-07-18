## 用法

来源：

- [官方扩展控制文件](https://github.com/greenbone/pg-gvm/blob/2427faa6e26ebb02d134b65efad8c9bf936dfa81/control.in)

`pg-gvm` — 为 Greenbone 漏洞管理守护进程提供 PostgreSQL 实用函数。

已复核目录快照记录的版本为 `22.6`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg-gvm";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg-gvm';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
