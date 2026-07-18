## 用法

来源：

- [官方扩展控制文件](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/geohistorical_objects.control)
- [官方上游文档](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/README.md)

`geohistorical_objects` — geohistorical_objects：地理空间类 PostgreSQL 扩展；上游说明为“提供地理历史对象支持”。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pg_trgm`, `pgsfti`, `postgis`, `unaccent`。

```sql
CREATE EXTENSION "geohistorical_objects";
SELECT extversion
FROM pg_extension
WHERE extname = 'geohistorical_objects';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
