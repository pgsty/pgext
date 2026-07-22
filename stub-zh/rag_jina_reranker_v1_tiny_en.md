## 用法

来源：

- [官方控制文件](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_jina_reranker_v1_tiny_en/rag_jina_reranker_v1_tiny_en.control)
- [官方 pgrag README](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [官方包清单](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_jina_reranker_v1_tiny_en/Cargo.toml)

`rag_jina_reranker_v1_tiny_en` 在本地运行 3300 万参数的 jina-reranker-v1-tiny-en 模型，并暴露标量与批量重排分数或距离。它已经弃用且不再维护。

### 核心流程

模型在后台 worker 中执行，因此创建扩展前必须预加载库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'rag_jina_reranker_v1_tiny_en.so'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION rag_jina_reranker_v1_tiny_en;

SELECT rag_jina_reranker_v1_tiny_en.rerank_distance(
    'What did the quick brown fox jump over?',
    'The quick brown fox jumps over the lazy dog'
);
```

在 macOS 上，上游说明要求在 `shared_preload_libraries` 中用 `.dylib` 替代 `.so`。

### SQL API

- `rag_jina_reranker_v1_tiny_en.rerank_score(text, text)` 返回一个分数；`text[]` 重载按输入顺序返回分数。
- `rag_jina_reranker_v1_tiny_en.rerank_distance(text, text)` 返回一个距离；`text[]` 重载按输入顺序返回距离。
- 对所有重载，距离等于分数的负值。因此距离越小，对应分数越高。

### 运维说明

所核验版本为 0.0.0，不可重定位、依赖 `vector`，并设置 `superuser = true`。worker 随 PostgreSQL 启动并在第一次使用时懒加载模型，因此要考虑首调延迟和内存占用。默认构建嵌入 ONNX 模型；可选的 `remote_onnx` 构建会在重启后下载模型，应指向受控存储。项目钉住了修补版 ONNX Runtime 组件并已归档弃用，因此应部署固定且测试过的产物，并规划替代方案。
