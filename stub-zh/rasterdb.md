## 用法

来源：

- [官方扩展控制文件](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/rasterdb.control)
- [官方上游文档](https://github.com/hexiaoting/rasterdb/blob/7241bd185509371b08651f9c04876144da59772c/README.md)

`rasterdb` — rasterdb：面向文件系统 PostGIS 栅格数据的外部数据包装器

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`postgis`。

```sql
CREATE EXTENSION "rasterdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'rasterdb';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
