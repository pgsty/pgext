## 用法

来源：

- [官方项目或服务商页面](https://Cognitum.One/RuVector)

`ruvector` — 面向 PostgreSQL 的 SIMD 加速向量类型与索引，提供图计算、求解器、TDA、本地嵌入及 AI SQL 函数。

已复核目录快照记录的版本为 `0.3.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ruvector";
SELECT extversion
FROM pg_extension
WHERE extname = 'ruvector';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
