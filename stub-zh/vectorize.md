

## 用法

> [pg_vectorize](https://github.com/ChuckHend/pg_vectorize)：在 Postgres 上进行向量搜索的最简方式。
> 来源：[README.md](https://raw.githubusercontent.com/ChuckHend/pg_vectorize/main/README.md)

一个 Postgres 扩展，自动化文本到嵌入向量的转换和编排，并提供与最流行 LLM 的集成。可以快速搭建和自动维护向量搜索、全文搜索和混合搜索，帮助你在 Postgres 上快速构建 RAG 和搜索引擎。

该项目重度依赖 [pgvector](https://github.com/pgvector/pgvector)（向量相似度搜索）、[pgmq](https://github.com/pgmq/pgmq)（后台工作进程编排）和 [SentenceTransformers](https://huggingface.co/sentence-transformers)。

**API 文档**：https://chuckhend.github.io/pg_vectorize/

--------

### 概览

pg_vectorize 提供两种方式为任意 Postgres 添加语义搜索、全文搜索和混合搜索，方便在 Postgres 上构建检索增强生成（RAG）。

模式概览：

- **HTTP 服务器**（适用于托管数据库）：运行一个独立服务连接 Postgres 并暴露 REST API（`POST /api/v1/table`、`GET /api/v1/search`）。
- **Postgres 扩展**（SQL）：将扩展安装到 Postgres 中，使用 `vectorize.table()` 和 `vectorize.search()` 等 SQL 函数（需要 Postgres 的文件系统访问权限）。

--------

### 快速开始 -- HTTP 服务器

使用 docker compose 在本地运行 Postgres 和 HTTP 服务：

```bash
# 运行 Postgres、嵌入向量服务器和管理 API
docker compose up -d
```

将示例数据集加载到 Postgres（可选）：

```bash
psql postgres://postgres:postgres@localhost:5432/postgres -f server/sql/example.sql
```

通过 HTTP API 创建嵌入向量任务。这会为现有数据生成嵌入向量，并持续监控更新或新数据：

```bash
curl -X POST http://localhost:8080/api/v1/table -d '{
		"job_name": "my_job",
		"src_table": "my_products",
		"src_schema": "public",
		"src_columns": ["product_name", "description"],
		"primary_key": "product_id",
		"update_time_col": "updated_at",
		"model": "sentence-transformers/all-MiniLM-L6-v2"
	}' -H "Content-Type: application/json"
```

```json
{"id":"16b80184-2e8e-4ee6-b7e2-1a068ff4b314"}
```

使用 HTTP API 搜索：

```bash
curl -G \
  "http://localhost:8080/api/v1/search" \
  --data-urlencode "job_name=my_job" \
  --data-urlencode "query=camping backpack" \
  --data-urlencode "limit=1" \
  | jq .
```

```json
[
  {
    "description": "Storage solution for carrying personal items on ones back",
    "fts_rank": 1,
    "price": 45.0,
    "product_category": "accessories",
    "product_id": 6,
    "product_name": "Backpack",
    "rrf_score": 0.03278688524590164,
    "semantic_rank": 1,
    "similarity_score": 0.6296013593673706,
    "updated_at": "2025-10-05T00:14:39.220893+00:00"
  }
]
```

--------

### 选择哪种模式？

- 当 Postgres 是托管的（RDS、Cloud SQL 等）或无法安装扩展时，使用 **HTTP 服务器**。只需数据库中有 `pgvector` 即可，HTTP 服务单独运行。
- 当自托管 Postgres 且可以安装扩展时，使用 **Postgres 扩展**。提供数据库内体验和直接的 SQL API 进行向量化和 RAG。

### 快速开始 -- Postgres 扩展（SQL）

```sql
CREATE EXTENSION vectorize CASCADE;
```

使用 `vectorize.table()` 创建嵌入向量任务，使用 `vectorize.search()` 直接从 SQL 进行语义搜索。
