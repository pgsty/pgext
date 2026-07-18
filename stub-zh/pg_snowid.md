## 用法

来源：

- [官方扩展控制文件](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/pg_snowid.control)
- [官方上游文档](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/README.md)
- [官方 Rust 包清单](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/Cargo.toml)

`pg_snowid` — 在 PostgreSQL 中生成线程安全且按时间排序的 Snowflake 风格 ID

已复核目录快照记录的版本为 `3.1.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_snowid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_snowid';
```

该上游项目与 `Rixl` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
