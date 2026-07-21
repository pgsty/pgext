## Usage

Sources:

- [pg_search v0.24.3 README](https://github.com/paradedb/paradedb/blob/v0.24.3/pg_search/README.md)
- [pg_search v0.24.3 release notes](https://github.com/paradedb/paradedb/releases/tag/v0.24.3)
- [ParadeDB documentation](https://docs.paradedb.com/)

pg_search adds ParadeDB's BM25 access method and query operators to PostgreSQL for ranked full-text, structured, and hybrid search. Use it when search must remain transactional with PostgreSQL data and must support BM25 scoring, highlighting, filters, aggregates, and joins.

### Create the Extension

    CREATE EXTENSION pg_search;

Upstream v0.24.3 supports PostgreSQL 15 and later. Use a build packaged for the exact PostgreSQL major version. The extension participates in planner and executor paths, so test query plans and resource use before production upgrades.

### Build a BM25 Index

Create a BM25 index with a stable unique key field:

    CREATE INDEX products_search_idx
    ON products
    USING bm25 (
      id,
      title,
      description,
      category,
      rating
    )
    WITH (key_field = 'id');

Keep the key field unique and non-null. Index only fields used by search or filtering; every indexed field increases build time, disk use, and write amplification.

### Query, Rank, and Highlight

The @@@ operator matches a field or indexed row against a ParadeDB query:

    SELECT id,
           title,
           paradedb.score(id) AS score,
           paradedb.snippet(description) AS snippet
    FROM products
    WHERE description @@@ 'wireless keyboard'
      AND category = 'electronics'
    ORDER BY score DESC
    LIMIT 20;

Use field-qualified query strings or the paradedb query constructors when user input must be constrained. Do not concatenate untrusted input into query syntax without validation.

For boolean queries, paradedb.boolean() can combine must, should, and must_not clauses and can set minimum_should_match. The extension also exposes index_created_at() for inspecting index creation time.

### User-Facing Object Index

- bm25: index access method for text and structured fields.
- @@@: search-match operator used in WHERE clauses.
- paradedb.score(key): BM25 relevance score for a matching row.
- paradedb.snippet(field): highlighted excerpt for the current match.
- paradedb.parse(...), paradedb.term(...), paradedb.boolean(...): typed query constructors.
- paradedb.index_info(...): index metadata and field configuration.
- paradedb.index_created_at(...): index creation timestamp.

### Version 0.24.3 Operational Changes

The 0.24.x line enables more aggregate and join scan paths and adds time and timetz support. Version 0.24.3 also bounds sequential-scan buffering, caps index-build worker memory, checks available disk space earlier, fixes GROUP BY cardinality routing, and raises an error when Tantivy would truncate a value.

These safeguards reduce runaway resource use but do not eliminate capacity planning. Monitor temporary space, index size, build duration, and query memory. Re-run representative EXPLAIN ANALYZE plans after upgrading because planner behavior can change.

### Compatibility and Caveats

- pg_search uses its own BM25 index implementation. Do not assume an index created by another extension is interchangeable.
- Local catalog metadata reports a bm25 access-method conflict with pg_textsearch and vchord_bm25; avoid loading competing implementations in the same database unless their documentation explicitly supports coexistence.
- Search indexes must be maintained with the table and can materially increase update cost.
- Ranking is query- and corpus-dependent. Benchmark with production-like text and filters rather than treating example scores as portable.
