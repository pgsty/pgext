## 用法

来源：

- [官方扩展控制文件](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/avocado.control)
- [官方 Rust 包清单](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/Cargo.toml)
- [官方项目或服务商页面](https://avocadodb.ai)

`avocado` — AvocadoDB 的 PostgreSQL 扩展，使用 pgvector/pgrx 在数据库内提供面向 AI 代理的确定性上下文编译。

已复核目录快照记录的版本为 `2.2.0`、类型为 `standard`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`vector`。
整理后的兼容版本集合为 `16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "avocado";
SELECT extversion
FROM pg_extension
WHERE extname = 'avocado';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
