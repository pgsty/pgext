## Usage

Sources:

- [README.md](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/README.md)
- [extension/pqnai--0.3.0.sql](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/extension/pqnai--0.3.0.sql)
- [docs/architecture.md](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/docs/architecture.md)
- [.github/workflows/ci.yml](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/.github/workflows/ci.yml)
- [extension/pqnai.control](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/extension/pqnai.control)

`pqnai` exposes forecasting and retrieval procedures backed by the external `pqnaid` Go worker. This early version implements moving-average forecasts and chunked retrieval using pgvector and Ollama; ARIMA, hybrid retrieval and reranking remain planned.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pqnai;
CALL pqnai.forecast(ARRAY[10,20,30,40,50]::double precision[], 3);
CALL pqnai.embed('The worker receives jobs using LISTEN and NOTIFY.');
CALL pqnai.ask('How does the worker receive jobs?', 1);
```

### Operational Boundaries

Install `vector` first and create `pqnai` as a superuser. The fixed schema holds `jobs`, `sources` and `chunks`. The worker must use the same database and be able to reach Ollama. No PostgreSQL preload is declared; running the SQL extension alone does not provide inference. Upstream CI targets PostgreSQL 14–17.

Use top-level `CALL` outside an explicit transaction block: the procedures commit queued jobs before waiting for the worker. A timeout does not undo that committed work. `pqnai.enqueue_job` and `pqnai.wait_for_job` are lower-level queue helpers. Restrict job and stored-document access; per-tenant access control is not implemented in this version. The installed vector column has 384 dimensions, so changing the embedding dimension requires a schema migration. Plan retention for queue and document data.
