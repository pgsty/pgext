## Usage

Sources:

- [Official documentation](https://github.com/wahicona/pg_sag_rag/blob/9bd58193d84cff7762794b75c56f40b4cbdc387a/README.md)
- [Extension control file](https://github.com/wahicona/pg_sag_rag/blob/9bd58193d84cff7762794b75c56f40b4cbdc387a/pg_sag_rag.control)
- [Official repository](https://github.com/wahicona/pg_sag_rag)

`pg_sag_rag` Pure-SQL multi-hop event/entity retrieval, query routing, and in-database RAG evaluation.

### Enablement

Install the files for the intended server, then create `pg_sag_rag` in the target database:

```sql
CREATE EXTENSION pg_sag_rag CASCADE;
```

The reviewed control or official workflow requires `vector`, `pg_trgm`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT sag_rag.add_evaluation_set('my-rag-eval');
SELECT sag_rag.add_evaluation_question(1, 'How much is AGI Bar foam?', '[0.86,0.11,0.10]'::vector);
SELECT sag_rag.link_evaluation_answer_event(1, 2);

SELECT sag_rag.run_evaluation_hybrid(1, p_top_k => 1);
SELECT sag_rag.run_evaluation_multihop(1, p_seed_k => 1, p_top_k => 10);
SELECT sag_rag.run_evaluation_auto(1);

SELECT run_id, strategy, parameters
FROM sag_rag.evaluation_run
ORDER BY run_id;

SELECT * FROM sag_rag.recall_at_k(1, 1);
SELECT * FROM sag_rag.recall_at_k(2, 10);
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `sag_rag.event` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `sag_rag.entity` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `sag_rag.create_hnsw_indexes` | FUNCTION | Callable function from the reviewed install surface. |
| `sag_rag.document` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `sag_rag.event_entity` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `sag_rag.recall_at_k` | FUNCTION | Callable function from the reviewed install surface. |
| `sag_rag.run_evaluation_auto` | FUNCTION | Callable function from the reviewed install surface. |
| `sag_rag.search_events_auto` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 14, 15, 16, 17; do not infer unlisted majors.
