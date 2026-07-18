## 用法

来源：

- [官方扩展控制文件](https://github.com/jasonmassie01/pg_sage/blob/a77458aaa3bed42adb7e47e481f476b13413cc3b/pg_sage.control)

`pg_sage` — 提供预加载式 C DBA 代理扩展，以及更新的独立 Go 监控与受控操作 sidecar。

已复核目录快照记录的版本为 `0.5.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_sage";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_sage';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
