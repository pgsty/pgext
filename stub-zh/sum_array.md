## 用法

来源：

- [官方扩展控制文件](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/sum_array.control)
- [官方上游文档](https://github.com/scottjustin5000/pg-sum-array/blob/4845e1e605773ea51ca51f8044dbd25da2698a0b/README.md)

`sum_array` — 用 C 将 PostgreSQL 整数或浮点数组求和并返回 float8。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "sum_array";
SELECT extversion
FROM pg_extension
WHERE extname = 'sum_array';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
