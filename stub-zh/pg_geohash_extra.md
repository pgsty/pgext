## 用法

来源：

- [官方扩展控制文件](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/pg_geohash_extra.control)
- [官方上游文档](https://github.com/siose-innova/pg_geohash_extra/blob/ea5fcbb319503599ef26f67bb165f99b39add763/README.md)
- [官方项目或服务商页面](http://www.siose-innova.es/)

`pg_geohash_extra` — 为 PostGIS 补充几何对象转 geohash、相邻 geohash 和规则网格等函数。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`postgis`。

```sql
CREATE EXTENSION "pg_geohash_extra";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_geohash_extra';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
