## 用法

来源：

- [官方扩展控制文件](https://github.com/CrunchyData/pgmonitor-extension/blob/e6859da8306cca419e16fc509d888e1da36f172e/pgmonitor.control)
- [官方上游文档](https://github.com/CrunchyData/pgmonitor-extension/blob/e6859da8306cca419e16fc509d888e1da36f172e/README.md)

`pgmonitor` — 为外部采集器提供 PostgreSQL 指标，并可通过预加载后台工作进程自动刷新物化视图和指标表。

已复核目录快照记录的版本为 `2.2.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgmonitor";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmonitor';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
