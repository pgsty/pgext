## Usage

Sources:

- [Official README](https://github.com/pgEdge/pgedge-vectorizer/blob/7d020be5399f3812fc38633cc287b5a613ed8c64/README.md)
- [Official control file](https://github.com/pgEdge/pgedge-vectorizer/blob/7d020be5399f3812fc38633cc287b5a613ed8c64/pgedge_vectorizer.control)
- [Official documentation](https://docs.pgedge.com/pgedge-vectorizer/)

`pgedge_vectorizer` asynchronously chunks table text and generates pgvector embeddings with OpenAI, Voyage AI, or Ollama. Triggers create queue work and background workers populate generated chunk tables, so source-row commit and embedding availability are separate events.

### Prerequisites and Configuration

The extension requires PostgreSQL 14 or later, `vector`, libcurl, and an embedding provider. Add the library and target databases before restarting PostgreSQL:

```ini
shared_preload_libraries = 'pgedge_vectorizer'
pgedge_vectorizer.databases = 'appdb'
pgedge_vectorizer.provider = 'openai'
pgedge_vectorizer.api_key_file = '/run/secrets/pgedge-vectorizer-key'
pgedge_vectorizer.model = 'text-embedding-3-small'
pgedge_vectorizer.num_workers = 2
```

The PostgreSQL service account must be able to read the key file; keep it out of SQL settings, backups, and logs. A preload or worker-count change requires restart.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgedge_vectorizer;

CREATE TABLE articles (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    content text NOT NULL
);

SELECT pgedge_vectorizer.enable_vectorization(
    source_table := 'articles',
    source_column := 'content',
    chunk_strategy := 'token_based',
    chunk_size := 400,
    chunk_overlap := 50
);

INSERT INTO articles (title, content)
VALUES ('Example', 'Text to split and embed asynchronously.');

SELECT * FROM pgedge_vectorizer.queue_status;
SELECT * FROM pgedge_vectorizer.failed_items;
```

`enable_vectorization` creates supporting triggers, queue state, and a chunk table. Use `disable_vectorization`, `retry_failed`, and `clear_completed` for lifecycle operations; review the `drop_chunk_table` choice before disabling.

### Operational Notes

Provider calls are external, asynchronous side effects. Expect rate limits, retries, duplicate attempts, model outages, and a period in which source data exists without embeddings. Monitor queue age and failures, make downstream reads tolerate lag, and re-embed deliberately after model, dimension, or chunking changes.

The control file requires superuser installation, is non-relocatable, and depends on `vector`. Generated chunk tables, triggers, queue metadata, GUCs, key files, provider connectivity, and background workers all belong in backup/restore and disaster-recovery tests. Restrict configuration and monitoring APIs because text chunks, errors, and outbound requests can contain sensitive data.
