## 用法

来源：

- [官方扩展控制文件](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/pg_ripple.control)
- [官方上游文档](https://trickle-labs.github.io/pg-ripple/)
- [官方 Rust 包清单](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/Cargo.toml)

`pg_ripple` — 在 PostgreSQL 18 内提供支持 SPARQL、SHACL、Datalog 推理、联邦查询与图分析的 RDF 知识图谱引擎。

已复核目录快照记录的版本为 `0.128.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_ripple";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_ripple';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
