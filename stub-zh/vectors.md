## 用法

来源：

- [官方扩展控制文件](https://github.com/tensorchord/pgvecto.rs/blob/2b290b34e8ba69104ea2f800fa53328c6ed6c236/vectors.control)
- [官方上游文档](https://github.com/tensorchord/pgvecto.rs/blob/2b290b34e8ba69104ea2f800fa53328c6ed6c236/README.md)
- [官方项目或服务商页面](https://docs.vectorchord.ai/getting-started/overview.html)

`vectors` — 用 Rust 编写的 PostgreSQL 向量数据库插件，面向 LLM 场景。

已复核目录快照记录的版本为 `0.4.0`、类型为 `preload`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "vectors";
SELECT extversion
FROM pg_extension
WHERE extname = 'vectors';
```

该上游项目与 `TensorChord` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
