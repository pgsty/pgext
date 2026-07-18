## 用法

来源：

- [官方扩展控制文件](https://github.com/gfphoenix78/autofailover/blob/6494a6ed6b8b882f287331996b2ab075e334afa0/autofailover.control)

`autofailover` — 提供自动故障转移辅助函数的 C 语言 PostgreSQL 扩展，可检查角色和 WAL 状态、执行提升并调整同步复制设置。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "autofailover";
SELECT extversion
FROM pg_extension
WHERE extname = 'autofailover';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
