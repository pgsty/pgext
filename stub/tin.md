## Usage

Sources:

- [README](https://github.com/planetscale/lead/blob/bd95c7e51b6afce81396790852ee2f2c169570ad/README.md)
- [Control file](https://github.com/planetscale/lead/blob/bd95c7e51b6afce81396790852ee2f2c169570ad/postgres/tin.control)
- [Cargo.toml](https://github.com/planetscale/lead/blob/bd95c7e51b6afce81396790852ee2f2c169570ad/Cargo.toml)
- [postgres/src/lib.rs](https://github.com/planetscale/lead/blob/bd95c7e51b6afce81396790852ee2f2c169570ad/postgres/src/lib.rs)

`tin` is the extension name of Lead, a TIN-compatible text-search implementation for development, CI, and staging. Version 1.0.3 deliberately trades performance for a small reference implementation and is unsuitable for production search.

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
