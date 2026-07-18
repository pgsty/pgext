## 用法

来源：

- [官方扩展控制文件](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/nanomsgtopdb.control)
- [官方上游文档](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/README.md)

`nanomsgtopdb` — 通过预加载后台工作进程将 nanomsg 消息写入 PipelineDB 流。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "nanomsgtopdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'nanomsgtopdb';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
