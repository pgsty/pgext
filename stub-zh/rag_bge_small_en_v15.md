## 用法

来源：

- [官方控制文件](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_bge_small_en_v15/rag_bge_small_en_v15.control)
- [官方 pgrag README](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/README.md)
- [官方包清单](https://github.com/neondatabase/pgrag/blob/53b0d8da7a310b4a28c40bdf3d4ae5efd156673b/exts/rag_bge_small_en_v15/Cargo.toml)

`rag_bge_small_en_v15` 在本地运行 3300 万参数的 bge-small-en-v1.5 模型，用于分词、按模型 token 分块，并生成 384 维段落或查询嵌入。它已经弃用且不再维护。

### 核心流程

模型在后台 worker 中执行，因此创建扩展前必须预加载库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'rag_bge_small_en_v15.so'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION rag_bge_small_en_v15;

SELECT rag_bge_small_en_v15.embedding_for_query(
    'What did the quick brown fox jump over?'
);
```

在 macOS 上，上游说明要求在 `shared_preload_libraries` 中用 `.dylib` 替代 `.so`。

### SQL API

- `rag_bge_small_en_v15.chunks_by_token_count(text, integer, integer)` 按模型 token 上限和指定重叠返回文本块。
- `rag_bge_small_en_v15.embedding_for_passage(text)` 为持久化段落返回 `vector(384)`。
- `rag_bge_small_en_v15.embedding_for_query(text)` 为搜索查询返回 `vector(384)`。

### 运维说明

所核验版本为 0.0.0，不可重定位、依赖 `vector`，并设置 `superuser = true`。worker 随 PostgreSQL 启动，但在第一次调用时才懒加载模型；应预期首调延迟和可观内存占用。默认构建嵌入大型 ONNX 模型；`remote_onnx` 构建会在重启后下载模型，应使用受控 URL。仓库钉住了修补版 ONNX Runtime 依赖，并记录了平台相关的符号错误风险。项目已归档并弃用，必须钉住完整构建。
