## Usage

Sources:

- [course_embedding_search.control](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docker/db/extension/course_embedding_search.control)
- [course_embedding_search--1.0.sql](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docker/db/extension/course_embedding_search--1.0.sql)
- [course_embedding_search.cpp](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docker/db/extension/course_embedding_search.cpp)
- [search-queries.md](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docs/backend/search-and-discovery/search-queries.md)
- [20261028_0022_course_embedding_search_blocks.py](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/migrations/versions/20261028_0022_course_embedding_search_blocks.py)
- [Dockerfile](https://github.com/aaronpiipponen/tuni-course-explorer/blob/70a9c040fd28e8806be4ff58a1381ce747f050cd/docker/db/Dockerfile)

`course_embedding_search` is the native scan extension embedded in the course-explorer application. SQL version 1.0 targets the project's PostgreSQL 18 and pgvector environment. It scans the application's packed vector blocks and returns each group's maximum cosine score.

### Enable and inspect

After installing the matching library and SQL files, a superuser enables both extensions:

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS course_embedding_search;
SELECT oid::regprocedure
FROM pg_proc WHERE proname = 'course_embedding_search_scan';
```

The application migration installs the extension before creating its search tables. `course_embedding_search_scan(regclass, vector, float8, int)` takes the block relation, an already hub-centred query vector, its squared norm and the group count. It returns `group_no` and `score`; application SQL joins group identities and applies the score threshold afterward.

### Data and operational boundaries

This function is not a generic scan over an arbitrary vector column. It expects the application's exact column order and packed layout: block number, row count, integer group array, squared-norm array and float32 byte payload. The serving-artifact builder prepares those blocks with the same embedding configuration and hub as the query. Invalid dimensions, norms, group IDs or array/payload sizes raise errors.

The function is strict, volatile and parallel unsafe. It allocates a best-score array proportional to the supplied group count and scans with the active snapshot. Its native relation scan does not perform ordinary SQL table-permission or row-security checks. Restrict function execution to trusted application roles and accept only the application's validated relation and inputs.

The extension is relocatable and does not require preload or a restart for fresh installation. It does not create the application tables, obtain embeddings or build serving artifacts. Preserve the matching extension, migration schema and artifact-generation configuration for recovery.
