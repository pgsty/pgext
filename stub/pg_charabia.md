## Usage

Sources:

- [README.md](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/README.md)
- [pg_charabia.control](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/pg_charabia.control)
- [Cargo.toml](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/Cargo.toml)
- [reference.md](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/docs/reference.md)
- [operations.md](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/docs/operations.md)
- [pg_charabia--0.4.1--0.4.2.sql](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/sql/pg_charabia--0.4.1--0.4.2.sql)

`pg_charabia` 0.4.2 adds multilingual analysis to PostgreSQL 18 native full-text search. It requires a UTF8 database. Its resources are embedded; native FTS does not require an external service or the optional BM25 extension.

### Core workflow

A database superuser installs the nonrelocatable extension into its fixed schema:

```sql
CREATE EXTENSION pg_charabia;
SELECT charabia.analysis_info();
SELECT * FROM charabia.analyze('PostgreSQL multilingual search');
CREATE TABLE documents (id bigint PRIMARY KEY, body text NOT NULL);
INSERT INTO documents VALUES (1, 'PostgreSQL multilingual search');
CREATE INDEX documents_gin ON documents
USING gin (to_tsvector('public.charabia', body));
SELECT id FROM documents
WHERE to_tsvector('public.charabia', body)
   @@ plainto_tsquery('public.charabia', 'PostgreSQL');
```

Use the same configuration for documents, queries, indexes and highlighting. `ts_headline` uses the original text. `charabia.query_any` provides OR-style plain query analysis.

### Configurations and privileges

`public.charabia` and `public.charabia_linguistic` provide ordinary linguistic analysis. Select `public.charabia_structured` explicitly for URLs, hosts, email addresses, paths and mixed letter/number tokens. The extension does not change `default_text_search_config`.

The fixed `charabia` schema must not allow untrusted creation. Installation reserves its SQL objects and the three public configurations. Create custom configurations and dictionaries in an application-owned schema. Copying a configuration shares its dictionaries; later dictionary changes affect every configuration that references them.

### Maintenance and limits

Native FTS needs no preload. Optional BM25 requires a separately installed compatible `pg_textsearch`, its dependencies, and preload followed by a server restart. Linux AMD64 is the release validation target; the macOS ARM64 runner covers local development with narrower test coverage.

The supported 0.4.1 upgrade requires stopping PostgreSQL before binary replacement, restarting, then running `ALTER EXTENSION pg_charabia UPDATE TO '0.4.2'` in each database. That upgrade preserves SQL objects and analysis, so it alone needs no reindex. Earlier releases have no direct upgrade path.

Analysis or dictionary changes can require recomputing stored vectors and rebuilding indexes from original text. Retain exact library/resource versions and external dictionary files for recovery. Rust allocations are not limited by `work_mem`; allocation failure can abort a backend and trigger recovery. Review the upstream language-analysis limits before choosing an indexing policy.
