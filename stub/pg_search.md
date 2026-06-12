## Usage

Sources: [ParadeDB extension install docs](https://docs.paradedb.com/deploy/self-hosted/extension), [create-index docs](https://docs.paradedb.com/documentation/indexing/create-index.md), [match docs](https://docs.paradedb.com/documentation/full-text/match.md), [score docs](https://docs.paradedb.com/documentation/sorting/score.md), [highlight docs](https://docs.paradedb.com/documentation/full-text/highlight.md), [v0.24.0 release](https://github.com/paradedb/paradedb/releases/tag/v0.24.0), [pg_search README](https://github.com/paradedb/paradedb/blob/v0.24.0/pg_search/README.md)

`pg_search` is ParadeDB's BM25-based search extension for PostgreSQL. The upstream README says support starts at PostgreSQL 15; Pigsty packages `0.24.0` for PostgreSQL 15-18 and builds it with `cargo-pgrx` 0.18.1.

### Enable And Create The Extension

```conf
shared_preload_libraries = 'pg_search'
```

```sql
CREATE EXTENSION pg_search;
```

The self-hosted extension docs require `shared_preload_libraries = 'pg_search'` before `CREATE EXTENSION`.

### Create A BM25 Index

Current examples use the `bm25` access method with a unique key field:

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating)
WITH (key_field = 'id');
```

Only one BM25 index is supported per table. `key_field` is mandatory, must be unique, and must be the first indexed column; text key fields must be untokenized.

### Query Operators And Helpers

The current docs use these query operators:

- `|||`: match disjunction, equivalent to `term1 OR term2`.
- `&&&`: match conjunction, equivalent to `term1 AND term2`.

Examples:

```sql
SELECT description, rating
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY rating
LIMIT 5;

SELECT description, pdb.score(id) AS score
FROM mock_items
WHERE description &&& 'running shoes'
ORDER BY score DESC
LIMIT 5;

SELECT description, pdb.snippet(description) AS snippet, pdb.score(id) AS score
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY score DESC
LIMIT 5;
```

Useful result helpers include `pdb.score(id)`, `pdb.snippet(field)`, `pdb.snippets(field)`, and `pdb.snippet_positions(field)`. Highlighting is relatively expensive and is not supported for fuzzy search queries.

### Notes

- The old quickstart URL was removed; use the versioned docs pages above for current `|||`, `&&&`, scoring, and highlighting syntax.
- Release `0.24.0` requires preloading `pg_search`, upgrades pgrx to 0.18.1, and documents crash-recovery, `ltree`, and inline-tokenizer work without changing the basic BM25 query examples above.
- The Pigsty metadata notes that the `bm25` access method conflicts with `pg_textsearch` and `vchord_bm25`; do not preload competing BM25 access-method extensions in the same cluster without testing the target combination.
