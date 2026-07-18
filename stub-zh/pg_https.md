## 用法

来源：

- [官方扩展控制文件](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/pg_https.control)
- [官方上游文档](https://github.com/Preet-Govind/pg_https/blob/cb3d887433b7118bd42f3a216b16db4ae2788033/readme.md)

`pg_https` — 基于 libcurl 的 PostgreSQL HTTP/HTTPS 客户端函数

已复核目录快照记录的版本为 `1.1`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_https";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_https';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
