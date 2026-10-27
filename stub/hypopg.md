


## Usage

Sources:

- [README](https://github.com/HypoPG/hypopg/blob/4e564ebbd46e1ab365ad0d05ca4a2f19cf316ba7/README.md)
- [Control](https://github.com/HypoPG/hypopg/blob/4e564ebbd46e1ab365ad0d05ca4a2f19cf316ba7/hypopg.control)
- [Release 1.4.4](https://github.com/HypoPG/hypopg/releases/tag/1.4.4)

HypoPG lets you create hypothetical (virtual) indexes that exist only in the current session and are considered by `EXPLAIN` (without `ANALYZE`) for query planning. This enables testing the impact of indexes without the cost of actually creating them.

### Functions

| Function | Description |
|----------|-------------|
| `hypopg_create_index(query text)` | Create a hypothetical index using CREATE INDEX syntax |
| `hypopg_list_indexes` | List all hypothetical indexes in the session |
| `hypopg_drop_index(oid)` | Drop a specific hypothetical index by OID |
| `hypopg_reset()` | Drop all hypothetical indexes |
| `hypopg()` | Return hypothetical indexes in pg_index-like format |

### Workflow

Create a test table and check the baseline plan:

```sql
CREATE TABLE hypo AS SELECT id, 'line ' || id AS val FROM generate_series(1, 10000) id;
ANALYZE hypo;
EXPLAIN SELECT * FROM hypo WHERE id = 1;
-- Seq Scan on hypo (cost=0.00..170.00 rows=1 width=15)
```

Create a hypothetical index:

```sql
SELECT * FROM hypopg_create_index('CREATE INDEX ON hypo (id)');
--  indexrelid |      indexname
-- ------------+----------------------
--       13543 | <13543>btree_hypo_id
```

Check the plan with the hypothetical index:

```sql
EXPLAIN SELECT * FROM hypo WHERE id = 1;
-- Index Scan using <13543>btree_hypo_id on hypo (cost=0.04..8.06 rows=1 width=15)
```

List and manage hypothetical indexes:

```sql
SELECT * FROM hypopg_list_indexes();
SELECT * FROM hypopg_drop_index(13543);
SELECT * FROM hypopg_reset();
```

### Limitations

- Only `EXPLAIN` without `ANALYZE` will consider hypothetical indexes
- Hypothetical indexes exist only in the current backend session
- Other concurrent connections are not affected
- Index names and some CREATE INDEX options are ignored


### 1.4.4 and Upgrade Boundaries

Upstream 1.4.4 fixes hypothetical index naming for large OIDs, dependent-object invalidation, INCLUDE-column limits and failed EXPLAIN state, with additional PostgreSQL 19 compatibility fixes. Hypothetical indexes remain local to one backend and affect EXPLAIN without ANALYZE. No extension update scripts are provided: after installing matching files, drop and recreate the extension in each database. Reconnect or reset the backend to discard hypothetical state. This upstream release does not assert a matching packaged build.
