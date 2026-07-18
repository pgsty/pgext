## 用法

来源：

- [官方扩展控制文件](https://github.com/gojuno/lostgis/blob/85f26871adef0eb5f8fed43704a9d28b2fb5b80b/lostgis.control)
- [官方上游文档](https://github.com/gojuno/lostgis/blob/85f26871adef0eb5f8fed43704a9d28b2fb5b80b/doc/lostgis.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/lostgis/)

`lostgis` — 为 PostGIS 提供时空类型、安全几何运算与空间辅助函数的纯 SQL 扩展。

已复核目录快照记录的版本为 `1.0.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`postgis`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "lostgis";
SELECT extversion
FROM pg_extension
WHERE extname = 'lostgis';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
