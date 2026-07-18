## 用法

来源：

- [官方扩展控制文件](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/plv8geo.control)
- [官方上游文档](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/README.md)

`plv8geo` — plv8geo 为 PostGIS 工作流提供 PL/v8 地理空间函数，并内置 D3、TopoJSON、JSTS、GeoTIFF 等 JavaScript 库。

已复核目录快照记录的版本为 `0.0.2`、类型为 `puresql`、实现语言为 `JavaScript`。
应先安装并验证声明的扩展依赖：`plv8`, `postgis`。

```sql
CREATE EXTENSION "plv8geo";
SELECT extversion
FROM pg_extension
WHERE extname = 'plv8geo';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
