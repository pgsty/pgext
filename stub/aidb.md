## Usage

Sources:

- [Official AIDB documentation](https://www.enterprisedb.com/docs/aidb/7/)
- [Official configuration guide](https://www.enterprisedb.com/docs/aidb/7/install/configure/)
- [Official pipeline guide](https://www.enterprisedb.com/docs/aidb/7/data-pipelines/)

`aidb` is EDB's in-database AI pipeline extension for parsing, chunking, embedding, indexing, retrieving, reranking, and generating over PostgreSQL data.

### Enablement and Access

Install AIDB 7, preload it, restart, and create it with `CASCADE` and its required `vector` dependency:

```ini
shared_preload_libraries = 'aidb'
```

```sql
CREATE EXTENSION aidb CASCADE;
GRANT aidb_users TO app_user;
```

`aidb_users` grants access to AIDB routines but not to application tables; grant source/target table privileges separately. PGD users must also call `aidb.bdr_setup()` after creation.

### Create and Query a Pipeline

```sql
SELECT aidb.create_pipeline(
  name               => 'support_kb',
  source             => 'support_articles',
  source_key_column  => 'id',
  source_data_column => 'body',
  step_1             => 'KnowledgeBase',
  step_1_options     => aidb.knowledge_base_config(
    model       => 'bge-small-en-v1.5-f16',
    data_format => 'Text'
  ),
  auto_processing    => 'Disabled'
);

SELECT aidb.run_pipeline('support_kb');
SELECT * FROM aidb.retrieve_text('support_kb', 'connection exhaustion', 3);
```

### Resource and Security Boundaries

Models may download on first use and inference consumes backend/background-worker CPU, memory, and storage. Limit pipeline roles, source-table privileges, model provenance, network egress, concurrency, and threading. External files or object storage additionally require EDB `pgfs`; because that provider identity conflicts with an unrelated open-source extension of the same name, verify the installed package rather than relying on name alone.
