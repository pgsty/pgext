## Usage

Sources:

- [Fork README](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/README.md)
- [Extension control file](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/pg_bm25.control)
- [Index option implementation](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/src/index_access/options.rs)

pg_bm25 is an early ParadeDB/Tantivy full-text-search extension that adds a BM25 index access method and Lucene-style queries to PostgreSQL. This stub describes the cataloged historical fork, whose API is materially different from current ParadeDB extensions; do not substitute current pg_search documentation for this version.

### Core Workflow

The control file requires superuser installation and fixes objects in the `paradedb` schema. Build an index over the table's composite row and declare the fields to index:

```sql
CREATE EXTENSION pg_bm25;

CREATE TABLE articles (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text,
  body text,
  rating integer
);

CREATE INDEX articles_bm25
ON articles
USING bm25 ((articles.*))
WITH (text_fields='{"title": {}, "body": {}}',
      numeric_fields='{"rating": {}}');

SELECT id, title, paradedb.rank_bm25(ctid)
FROM articles
WHERE articles @@@ 'body:postgres AND rating:[4 TO *]'
ORDER BY paradedb.rank_bm25(ctid) DESC;
```

### Important Objects

- `bm25` is the index access method backed by Tantivy files.
- `@@@` executes a text query against the BM25 index attached to the row's table.
- `text_fields`, `numeric_fields`, `boolean_fields`, and `json_fields` are JSON index options.
- `paradedb.rank_bm25` returns the current query's score for a tuple.
- `paradedb.minmax_bm25` normalizes scores for a named index and query.
- `paradedb.schema_bm25` reports the Tantivy field schema.
- `paradedb.aggregation` runs a Tantivy aggregation request and returns JSON.

### Operational Notes

The fork README says PostgreSQL 12 or newer, while its packaged binaries and tests were much narrower. Index data is maintained outside ordinary heap storage by extension code, so test crash recovery, backup/restore, replication, upgrades, VACUUM behavior, and concurrent writes before relying on it. The repository is a historical fork and uses an old 0.3.x API; prefer a maintained search extension for new deployments.
