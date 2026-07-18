## 用法

来源：

- [官方扩展控制文件](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_jina_reranker_v1_tiny_en/rag_jina_reranker_v1_tiny_en.control)
- [官方上游文档](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [官方 Rust 包清单](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_jina_reranker_v1_tiny_en/Cargo.toml)

`rag_jina_reranker_v1_tiny_en` — 在预加载后台工作进程中本地运行 jina-reranker-v1-tiny-en，并提供 SQL 重排分数与距离函数。

已复核目录快照记录的版本为 `0.0.0`、类型为 `preload`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`vector`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "rag_jina_reranker_v1_tiny_en";
SELECT extversion
FROM pg_extension
WHERE extname = 'rag_jina_reranker_v1_tiny_en';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
