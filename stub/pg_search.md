
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_search;
> CREATE INDEX search_idx ON mock_items
> USING bm25 (id, description, category)
> WITH (key_field='id');
> SELECT * FROM mock_items WHERE description @@@ 'keyboard';
> ```
>
> Sources: [README](https://github.com/paradedb/paradedb/tree/dev/pg_search), [Quickstart](https://docs.paradedb.com/documentation/getting-started/quickstart), [Install docs](https://docs.paradedb.com/documentation/getting-started/install)

`pg_search` is ParadeDB's full text search extension for PostgreSQL. It provides BM25-based indexing and querying on heap tables, is built on Tantivy, and the current upstream README states support starts at PostgreSQL 15.

## Setup

The upstream installation docs highlight one critical requirement: `pg_search` must be included in `shared_preload_libraries` so its background worker can process index writes.

```ini
shared_preload_libraries = 'pg_search'
```

After that:

```sql
CREATE EXTENSION pg_search;
ALTER SYSTEM SET paradedb.pg_search_telemetry TO 'off';
```

## Creating a BM25 Index

The quickstart demonstrates creating a BM25 index on a heap table, with a unique key field:

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating, in_stock, created_at, metadata, weight_range)
WITH (key_field='id');
```

The existing docs emphasize that the `key_field` must be unique and that BM25 indexes are the core access method for search queries.

## Querying

The `@@@` operator performs search queries:

```sql
SELECT description, rating, category
FROM mock_items
WHERE description @@@ 'keyboard' AND rating > 2
ORDER BY rating
LIMIT 5;
```

ParadeDB also documents helper functions for relevance scoring and snippets:

```sql
SELECT description, paradedb.score(id)
FROM mock_items
WHERE description @@@ 'keyboard'
ORDER BY paradedb.score(id) DESC
LIMIT 5;

SELECT description, paradedb.snippet(description), paradedb.score(id)
FROM mock_items
WHERE description @@@ 'keyboard'
ORDER BY paradedb.score(id) DESC
LIMIT 5;
```

Phrase search is supported with quoted expressions:

```sql
SELECT description
FROM mock_items
WHERE description @@@ '"metal keyboard"';
```

## Text Configuration

The quickstart also shows that text fields can be wrapped with tokenizer configuration, for example English stemming:

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, (description::pdb.simple('stemmer=english')), category)
WITH (key_field='id');
```

For deeper setup and operational guidance, the upstream project points to the ParadeDB documentation site as the primary docs surface.
