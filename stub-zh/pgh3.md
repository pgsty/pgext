## 用法

来源：

- [官方扩展控制文件](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/pgh3.control)
- [官方上游文档](https://github.com/dlr-eoc/pgh3/blob/a225ac6da928cb0668780221da64061d3c4b1d54/README.md)

`pgh3` — 为 PostGIS 几何对象提供 Uber H3 分层地理空间索引函数

已复核目录快照记录的版本为 `0.3.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`postgis`。
整理后的兼容版本集合为 `10,11,12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgh3";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgh3';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
