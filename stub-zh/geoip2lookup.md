## 用法

来源：

- [官方扩展控制文件](https://github.com/adjust/pg-geoip2lookup/blob/5c8144578f1c1461cc6ec0a12e18e9708de77220/geoip2lookup.control)
- [官方上游文档](https://pgxn.org/dist/geoip2lookup/0.0.3/)
- [官方 PGXN 分发页](https://pgxn.org/dist/geoip2lookup/)

`geoip2lookup` — 使用 MaxMind GeoIP2 MMDB 文件查询 IP 地址地理信息的 PL/PerlU 扩展。

已复核目录快照记录的版本为 `0.0.3`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plperlu`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "geoip2lookup";
SELECT extversion
FROM pg_extension
WHERE extname = 'geoip2lookup';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
