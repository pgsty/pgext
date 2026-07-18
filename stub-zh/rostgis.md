## 用法

来源：

- [官方扩展控制文件](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/rostgis.control)
- [官方上游文档](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/docs/README.md)
- [官方 Rust 包清单](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/Cargo.toml)

`rostgis` — 使用 Rust 编写的实验性 PostGIS 兼容空间扩展。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "rostgis";
SELECT extversion
FROM pg_extension
WHERE extname = 'rostgis';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
