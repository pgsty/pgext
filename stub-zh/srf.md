## 用法

来源：

- [官方扩展控制文件](https://github.com/zombodb/srf/blob/350fb8962cf5fdd06d4bdca1ed1a01a7a798a498/srf.control)

`srf` — 用 C 实现闭区间整数集合返回函数的最小示例扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "srf";
SELECT extversion
FROM pg_extension
WHERE extname = 'srf';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
