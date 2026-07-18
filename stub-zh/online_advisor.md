## 用法

来源：

- [官方扩展控制文件](https://github.com/knizhnik/online_advisor/blob/e24652f824ffac872141634394ea76320fbae4f0/online_advisor.control)
- [官方上游文档](https://github.com/knizhnik/online_advisor/blob/e24652f824ffac872141634394ea76320fbae4f0/README.md)

`online_advisor` — 基于真实查询负载在线建议索引、扩展统计信息和预备语句优化的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "online_advisor";
SELECT extversion
FROM pg_extension
WHERE extname = 'online_advisor';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
