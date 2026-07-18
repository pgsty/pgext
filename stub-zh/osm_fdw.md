## 用法

来源：

- [官方扩展控制文件](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/osm_fdw.control)
- [官方上游文档](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/doc/osm_fdw.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/osm_fdw/)

`osm_fdw` — osm_fdw：用于读取 OpenStreetMap PBF 文件的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `4.1.2`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "osm_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'osm_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
