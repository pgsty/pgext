## 用法

来源：

- [官方上游文档](https://github.com/AbdulYadi/pgcodec7/blob/b54cdcbadb03042396a65b8db7f9913fb0686fc1/README.md)

`pgcodec7` — pgcodec7：使用 7 位或 6 位文本分组对 bytea 数据进行编码与解码的 C 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgcodec7";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcodec7';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
