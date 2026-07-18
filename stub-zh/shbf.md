## 用法

来源：

- [官方扩展控制文件](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/shbf.control)
- [官方上游文档](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/README.md)

`shbf` — 提供布隆过滤器、移位布隆过滤器和 Count-Min Sketch PostgreSQL 数据类型及函数的 C 扩展。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "shbf";
SELECT extversion
FROM pg_extension
WHERE extname = 'shbf';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
