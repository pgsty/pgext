## Usage

Sources:

- [pg_steadytext/README.md](https://github.com/julep-ai/steadytext/blob/e4d877620a6801de9705e362eda8d3a4f80cc5d1/pg_steadytext/README.md)
- [pg_steadytext/sql/pg_steadytext--2025.11.25.sql](https://github.com/julep-ai/steadytext/blob/e4d877620a6801de9705e362eda8d3a4f80cc5d1/pg_steadytext/sql/pg_steadytext--2025.11.25.sql)
- [pg_steadytext/Makefile](https://github.com/julep-ai/steadytext/blob/e4d877620a6801de9705e362eda8d3a4f80cc5d1/pg_steadytext/Makefile)
- [pg_steadytext/pg_steadytext.control](https://github.com/julep-ai/steadytext/blob/e4d877620a6801de9705e362eda8d3a4f80cc5d1/pg_steadytext/pg_steadytext.control)

`pg_steadytext` exposes SteadyText generation and embeddings through PL/Python, with SQL caches, asynchronous queues and prompt templates. The extension version is 2025.11.25, distinct from the Python library’s version.

### Core Workflow

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION vector;
CREATE EXTENSION pg_steadytext;
SELECT steadytext_generate('Write a haiku about PostgreSQL');
SELECT steadytext_embed('PostgreSQL is a database');
SELECT * FROM steadytext_cache_stats();
```

### Operational Boundaries

Upstream requires PostgreSQL 14+, Python 3.10+, `plpython3u` and `vector`. Install the Python dependencies into the interpreter linked to PostgreSQL, not merely the shell interpreter. Creating untrusted-language functions requires a superuser despite the control’s superuser=false setting. Objects are fixed in `public`; the Makefile installs SQL and Python files rather than a native shared library. Review installation SQL against existing data: it drops `steadytext_cache` with CASCADE. Generation may load large models or contact a configured daemon. Remote models are an explicit unsafe mode and can send prompts outside the database. Higher-level input limits are the application’s responsibility. Asynchronous processing requires separately operated workers; optional `pg_cron` can schedule cache eviction. No shared preload requirement is documented.
