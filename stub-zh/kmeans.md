## 用法

来源：

- [官方扩展控制文件](https://github.com/umitanuki/kmeans-postgresql/blob/d6a57eda9e4908ef5768ad2a5c70bd92d0abb3ae/kmeans.control)
- [官方上游文档](https://pgxn.org/dist/kmeans/1.1.0/doc/kmeans.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/kmeans/)

`kmeans` — 以 PostgreSQL 用户自定义窗口函数实现 K-means 聚类。

已复核目录快照记录的版本为 `1.1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "kmeans";
SELECT extversion
FROM pg_extension
WHERE extname = 'kmeans';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
