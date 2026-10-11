## Usage

Sources:

- [README.md](https://github.com/toiletbril/pg_what_is_happening/blob/e04e7ca309cfb475a6927d2b93cfbf3892290313/README.md)
- [pg_what_is_happening--1.1.sql](https://github.com/toiletbril/pg_what_is_happening/blob/e04e7ca309cfb475a6927d2b93cfbf3892290313/pg_what_is_happening--1.1.sql)
- [src/pg_what_is_happening.c](https://github.com/toiletbril/pg_what_is_happening/blob/e04e7ca309cfb475a6927d2b93cfbf3892290313/src/pg_what_is_happening.c)
- [pg_what_is_happening.control](https://github.com/toiletbril/pg_what_is_happening/blob/e04e7ca309cfb475a6927d2b93cfbf3892290313/pg_what_is_happening.control)

`pg_what_is_happening` exposes live execution and buffer metrics for active queries. Upstream claims PostgreSQL 9.4 through 18 and newer, but explicitly does not guarantee a stable library ABI.

### Core Workflow

```ini
shared_preload_libraries = 'pg_what_is_happening'
```

```sql
CREATE EXTENSION pg_what_is_happening;
SELECT * FROM what_is_happening.v1_status;
```

### Operational Boundaries

Preload and restart are mandatory; extension creation requires a superuser. `what_is_happening.v1_status` reports one row per active plan node. The shipped SQL grants the view and its backing function to PUBLIC, so review access to query text. The optional background worker exposes `/metrics`, defaulting to localhost port 9187; protect that endpoint. `what_is_happening.min_cost_to_track` filters cheap queries. Limits on tracked queries, nodes and query-text length require restart, while enablement and signal timeout can reload. Completed queries disappear; this is not a historical statistics store.
