## 用法

来源：

- [官方扩展控制文件](https://github.com/semenikhind/complex_contrib/blob/5b815cbf896ba3713862c2302202411d27c7c421/complex.control)

`complex` — 为 PostgreSQL 提供复数数据类型。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "complex";
SELECT extversion
FROM pg_extension
WHERE extname = 'complex';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
