## Usage

Sources:

- [README](https://github.com/planetscale/lead/blob/e3ed2f4ce1388167b0254a0c0c706a61d0a37d83/README.md)
- [Control file](https://github.com/planetscale/lead/blob/e3ed2f4ce1388167b0254a0c0c706a61d0a37d83/postgres/tin.control)
- [Cargo.toml](https://github.com/planetscale/lead/blob/e3ed2f4ce1388167b0254a0c0c706a61d0a37d83/Cargo.toml)
- [postgres/src/lib.rs](https://github.com/planetscale/lead/blob/e3ed2f4ce1388167b0254a0c0c706a61d0a37d83/postgres/src/lib.rs)
- [PlanetScale TIN and Lead](https://planetscale.com/docs/postgres/search)

`tin` is the extension name of Lead, a TIN-compatible text-search implementation for development, CI, and staging. The reviewed 1.0.6 source snapshot deliberately trades performance for a small reference implementation and is unsuitable for production search.

### Core Workflow

```sql
CREATE EXTENSION tin;
CREATE TABLE documents (id integer, body text);
INSERT INTO documents VALUES (1, 'craft beer'), (2, 'wine'), (3, 'beer festival');
CREATE INDEX documents_search ON documents USING tin (body);
SELECT id, tin.full_score(ctid) AS score
FROM documents WHERE body ==> 'beer' ORDER BY score DESC;
```

### Objects and Behavior

The `tin` access method and `==>` operator evaluate TINQL predicates. `tin.score`, `tin.full_score`, `tin.max_score`, and `tin.score_inspect` expose scoring; `tin.highlight` and `tin.highlight_ansi` mark matches. Scoring calls bind to the matching search expression and index.

Every index scan returns all heap pages as lossy candidates. PostgreSQL rechecks the visible tuples, expressions, and partial-index predicates. The index stores no search data; scoring may rescan visible heap rows. This preserves query semantics without providing a production inverted index.

### Requirements

The reviewed source targets PostgreSQL 17 and 18. Installation requires a superuser; objects use the fixed `tin` schema. The library loads on demand and needs no preload or restart. Keep realistic data volumes in mind when using it in tests.

PlanetScale documents TIN as its production search implementation and Lead as the local development substitute. Both use the canonical extension name `tin`; this catalog entry tracks the open-source Lead implementation. Its source-snapshot version is not a claim about the version deployed by PlanetScale.

The current source adds a trailing `stemmer` argument to `tin.tokenize`, `tin.ql_parse`, `tin.highlight`, and `tin.highlight_ansi`, while keeping the earlier signatures under names ending in `_v1_0_3`. Changing an index's `stemmer` requires `REINDEX`. A score call must be at the same query level as its matching `==>` predicate.
