

## Usage

> [pg_vectorize](https://github.com/ChuckHend/pg_vectorize): The simplest way to do vector search on Postgres.
> Source: [README.md](https://raw.githubusercontent.com/ChuckHend/pg_vectorize/main/README.md)

A Postgres extension that automates the transformation and orchestration of text to embeddings and provides hooks into the most popular LLMs. This allows you to get up and running and automate maintenance for vector search, full text search, and hybrid search, which enables you to quickly build RAG and search engines on Postgres.

This project relies heavily on [pgvector](https://github.com/pgvector/pgvector) for vector similarity search, [pgmq](https://github.com/pgmq/pgmq) for orchestration in background workers, and [SentenceTransformers](https://huggingface.co/sentence-transformers).

**API Documentation**: https://chuckhend.github.io/pg_vectorize/

--------

### Overview

pg_vectorize provides two ways to add semantic, full text, and hybrid search to any Postgres, making it easy to build retrieval-augmented generation (RAG) on Postgres.

Modes at a glance:

- **HTTP server** (recommended for managed DBs): run a standalone service that connects to Postgres and exposes a REST API (`POST /api/v1/table`, `GET /api/v1/search`).
- **Postgres extension** (SQL): install the extension into Postgres and use SQL functions like `vectorize.table()` and `vectorize.search()` (requires filesystem access to Postgres).

--------

### Quick Start -- HTTP Server

Run Postgres and the HTTP servers locally using docker compose:

```bash
# runs Postgres, the embeddings server, and the management API
docker compose up -d
```

Load the example dataset into Postgres (optional):

```bash
psql postgres://postgres:postgres@localhost:5432/postgres -f server/sql/example.sql
```

Create an embedding job via the HTTP API. This generates embeddings for the existing data and continuously watches for updates or new data:

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

Search using the HTTP API:

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

### Which Mode Should I Pick?

- Use the **HTTP server** when your Postgres is managed (RDS, Cloud SQL, etc.) or you cannot install extensions. It requires only that `pgvector` is available in the database. You run the HTTP services separately.
- Use the **Postgres extension** when you self-host Postgres and can install extensions. This provides an in-database experience and direct SQL APIs for vectorization and RAG.

### Quick Start -- Postgres Extension (SQL)

```sql
CREATE EXTENSION vectorize CASCADE;
```

Use `vectorize.table()` to create an embedding job and `vectorize.search()` to perform semantic search directly from SQL.
