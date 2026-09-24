## Usage

Sources:

- [Control file](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/pg_weave.control)
- [Makefile](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/Makefile)
- [README](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/README.md)
- [Installation SQL](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/sql/pg_weave--0.1.0.sql)
- [Vector storage boundary](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/sql/pg_weave--0.7.0--0.8.0.sql)
- [Latest upgrade SQL](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/sql/pg_weave--0.9.0--0.10.0.sql)
- [Regression examples](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/sql/weave.sql)
- [Module initialization and settings](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/src/am/customscan.c)
- [Licensing and provenance](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/doc/LICENSING.md)
- [Migration status](https://codeberg.org/gregburd/pg_weave/src/commit/d00b21d145eb3390e9e85576d44ee4e573038861/doc/MIGRATION.md)
- [PostgreSQL extension trust](https://www.postgresql.org/docs/18/extend-extensions.html)

`pg_weave` provides the `weave` index access method for BM25 lexical search. The reviewed control version is 0.10.0. This is an experimental extension: the lexical channel is queryable, while the vector channel is storage-only and the advertised fuzzy and fused search workflows remain unfinished.

### Enablement and Core Workflow

Upstream targets PostgreSQL 17 or later and reports testing on 17 and 18. After installing the matching extension files, enable it and index analyzed documents:

```sql
CREATE EXTENSION pg_weave;

CREATE TABLE docs (id serial PRIMARY KEY, d wdoc);
INSERT INTO docs (d) VALUES
  (to_wdoc('english', 'postgres streaming replication')),
  (to_wdoc('english', 'a slow green turtle'));

CREATE INDEX docs_weave ON docs USING weave (d);

SELECT id FROM docs WHERE d @@@ to_wquery('english', 'postgres') ORDER BY id;
SELECT weave_count('docs_weave', to_wquery('english', 'postgres'));
SELECT ndocs, avgdl, nterms FROM weave_index_stats('docs_weave');
```

The default `wdoc_lex_ops` operator class indexes `wdoc` values. `to_wdoc` and `to_wquery` accept a text-search configuration; use the same configuration for documents and queries. The `@@@` operator matches a query, and `<=>` supports ranked ordering.

No explicit preload or restart is required. The control file declares `trusted = true` and `relocatable = true`; installation is therefore available to users with database `CREATE` privilege when the server files are installed. This is an upstream trust declaration, not a claim that the project is production-ready.

### Inspection and Maintenance

- `weave_count` counts matches through an index.
- `weave_index_stats` reports document count, average document length, and vocabulary size.
- `weave_check` checks index structures.
- `weave_merge` folds pending data into segments; `weave_vacuum` performs index compaction.

`pg_weave.wand_initial_k` defaults to 32 and controls the initial ranked-scan width. `pg_weave.build_collapse_max_mb` defaults to 4096 MiB; larger builds retain bounded tiers instead of forcing a single segment. Both are user-settable settings. Inspect representative query plans and maintenance costs before using the extension for a workload.

### Version and Feature Boundaries

Installation follows the 0.1.0 baseline and nine upgrade scripts to 0.10.0. The in-tree metadata still says 0.7.0 and the README status says 0.3.0; the control and complete SQL chain establish the reviewed installation version.

`wvec` exists, but `wvec_weave_ops` registers storage without query operators. The README's `score()` and `fuse()` functions and its proposed multi-channel operator classes are absent from the installed SQL. Use the verified lexical workflow above. The automatic migration paths from other search extensions are also documented as unimplemented.

The project declares the PostgreSQL License and documents separate MIT and BSD licenses for vendored components. Its canonical repository is on Codeberg, with a synchronized GitHub mirror. Source availability and the declared trust setting do not imply a packaged or production-supported release.
