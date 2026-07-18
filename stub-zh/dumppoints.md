## 用法

来源：

- [官方扩展控制文件](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/dumppoints.control)
- [官方上游文档](https://github.com/nmandery/dumppoints/blob/96343c7e4eb76765da69737f79611fa7355c9e5b/README)

`dumppoints` — 面向 PostGIS 几何集合的 ST_DumpPoints C 语言实现。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`postgis`。

```sql
CREATE EXTENSION "dumppoints";
SELECT extversion
FROM pg_extension
WHERE extname = 'dumppoints';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
