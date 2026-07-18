## 用法

来源：

- [官方扩展控制文件](https://github.com/neurondb/neurondb/blob/66284d0476a864bbe074017b517a34daef454fba/neurondb.control)
- [官方上游文档](https://www.neurondb.ai/docs)
- [官方项目或服务商页面](https://www.neurondb.ai)

`neurondb` — 面向 PostgreSQL 的向量检索、机器学习和 RAG 扩展。

已复核目录快照记录的版本为 `4.0.0-devel`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "neurondb";
SELECT extversion
FROM pg_extension
WHERE extname = 'neurondb';
```

该上游项目与 `NeuronDB` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
