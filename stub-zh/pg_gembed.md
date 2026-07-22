## 用法

来源：

- [官方 README](https://github.com/JoelDiaz222/pg_gembed/blob/c7f4d005a2462a1af9464424a5418cdca40bd0a4/README.md)
- [官方扩展 SQL](https://github.com/JoelDiaz222/pg_gembed/blob/c7f4d005a2462a1af9464424a5418cdca40bd0a4/sql/pg_gembed--1.0.0.sql)
- [官方后台 worker 实现](https://github.com/JoelDiaz222/pg_gembed/blob/c7f4d005a2462a1af9464424a5418cdca40bd0a4/src/embedding_worker.c)

`pg_gembed` 1.0.0 通过调用本地 Rust `libgembed` 核心，从 SQL 生成文本、图像和多模态向量 embedding。它支持本地或远程的可插拔 backend，并返回 `vector` 类型，因此依赖 `vector` 扩展。

### 核心流程

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_gembed;

SELECT embed_text(
  'embed_anything',
  'Qdrant/all-MiniLM-L6-v2-onnx',
  'Hello world'
);

CREATE TABLE articles (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  content text NOT NULL,
  embedding vector(384)
);

INSERT INTO articles (content, embedding)
VALUES (
  'Transformers use attention mechanisms.',
  embed_text(
    'embed_anything',
    'Qdrant/all-MiniLM-L6-v2-onnx',
    'Transformers use attention mechanisms.'
  )
);

SELECT id, content
FROM articles
ORDER BY embedding <=> embed_text(
  'embed_anything',
  'Qdrant/all-MiniLM-L6-v2-onnx',
  'machine learning'
)
LIMIT 10;
```

backend 与 model 字符串选择 `libgembed` 已知的实现和模型。模型文件、网络访问、凭据和加速要求取决于具体 backend。

### 重要对象

- `embed_text(text, text, text)` 和 `embed_texts(text, text, text[])` 分别嵌入一条或多条文本。
- `embed_image(text, text, bytea)`、`embed_images(text, text, bytea[])` 以及对应的 with-ID 变体用于嵌入图像数据。
- `embed_multimodal(text, text, bytea[], text[])` 组合图像与文本输入。
- `embed_image_directory(text, text, text)` 和 `embed_image_directories(text, text, text[])` 读取数据库服务器上的图像路径。
- `input_type` 与标量/数组 `embed(...)` 重载负责分派文本、图像或图像目录输入。
- `gembed.embedding_jobs` 和 `gembed.job_status` 保存并报告异步 embedding job。

### 预加载与 Worker 边界

同步 embedding 函数可以在会话中加载库，无需重启。后台 job worker 不同：`_PG_init` 只在 PostgreSQL 处理 `shared_preload_libraries` 时注册它。要使用 `gembed.embedding_jobs`，必须预加载 `pg_gembed` 并重启 PostgreSQL。`gembed.embedding_worker_naptime` 默认为 10 秒，`gembed.embedding_worker_batch_size` 默认为 256；两者都是可 reload 的 GUC。

### 运维说明

控制文件要求由超级用户安装并依赖 `vector`；README 要求 PostgreSQL 17 或更高版本。固定 1.0.0 源码中的 worker 会打开名为 `joeldiaz` 的固定数据库，因此在其他数据库依赖异步 job 前必须复核或修改该代码。目录函数读取服务端本地路径，embedding backend 也可能加载大型模型或联系已配置的 gRPC/HTTP 服务。生产使用前应在目标系统上评估模型内存、CPU/GPU 消耗、网络暴露和并行查询行为。
