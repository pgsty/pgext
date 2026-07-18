## 用法

来源：

- [官方扩展控制文件](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/statsmgr.control)
- [官方上游文档](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/README.statsmgr.md)
- [官方上游来源](https://codeberg.org/Data-Bene/StatsMgr)

`statsmgr` — 在动态共享内存中快照 PostgreSQL WAL、SLRU、检查点、归档、I/O 与后台写进程累计统计信息的 C 扩展。

已复核目录快照记录的版本为 `0.1-alpha`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "statsmgr";
SELECT extversion
FROM pg_extension
WHERE extname = 'statsmgr';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
