## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/hstore_ops.control)
- [官方上游文档](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/README.md)

`hstore_ops` — 面向 hstore 的替代 GIN 操作符类，侧重更小的索引与更快的包含查询。

已复核目录快照记录的版本为 `1.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`hstore`。

```sql
CREATE EXTENSION "hstore_ops";
SELECT extversion
FROM pg_extension
WHERE extname = 'hstore_ops';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
