## Usage

- Sources: [ParadeDB extension install docs](https://docs.paradedb.com/deploy/self-hosted/extension), [quickstart](https://docs.paradedb.com/documentation/getting-started/quickstart), [v0.23.0 release](https://github.com/paradedb/paradedb/releases/tag/v0.23.0), [pg_search README](https://github.com/paradedb/paradedb/blob/dev/pg_search/README.md)

`pg_search` is ParadeDB's BM25-based search extension for PostgreSQL. The upstream README says support starts at PostgreSQL 15, and the v0.23.0 self-hosted install docs still require preloading the library before `CREATE EXTENSION`.

### Enable And Create The Extension

```conf
shared_preload_libraries = 'pg_search'
```

```sql
CREATE EXTENSION pg_search;
```

The self-hosted extension docs for v0.23.0 describe prebuilt binaries for Postgres 15+.

### Create A BM25 Index

Quickstart examples use the `bm25` access method with a unique key field:

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating)
WITH (key_field = 'id');
```

The v0.23.0 release also notes newly tunable BM25 `k1` and `b` parameters per field.

### Query Operators And Helpers

The current quickstart uses these query operators:

- `|||`: match disjunction, equivalent to `term1 OR term2`.
- `&&&`: match conjunction, equivalent to `term1 AND term2`.

Examples:

```sql
SELECT description, rating
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY rating
LIMIT 5;

SELECT description, pdb.score(id)
FROM mock_items
WHERE description &&& 'running shoes'
ORDER BY score DESC
LIMIT 5;

SELECT description, pdb.snippet(description), pdb.score(id)
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY score DESC
LIMIT 5;
```

### Notes

The development README points users to the docs site for installation and usage instead of documenting SQL details inline. The quickstart is therefore the authoritative usage surface for current `pg_search` syntax, and it reflects the post-0.20 API rather than the older `@@@` examples still found in some secondary materials.
