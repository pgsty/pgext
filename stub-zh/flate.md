## 用法

来源：

- [官方扩展控制文件](https://github.com/grahamedgecombe/pgflate/blob/a6880b7f5c2304b286e944945482efc7226ec18f/flate.control)
- [官方上游文档](https://github.com/grahamedgecombe/pgflate/blob/a6880b7f5c2304b286e944945482efc7226ec18f/README.markdown)

`flate` — 为 PostgreSQL bytea 数据提供原始 DEFLATE 压缩和解压函数。

已复核目录快照记录的版本为 `1.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "flate";
SELECT extversion
FROM pg_extension
WHERE extname = 'flate';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
