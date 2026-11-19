## Usage

Sources:

- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/docs/pg_bigm_en.md)
- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/pg_bigm.control)
- [Official source](https://github.com/pgbigm/pg_bigm/blob/v1.2-20261118/README.md)
- [Official source](https://github.com/pgbigm/pg_bigm/releases/tag/v1.2-20261118)

`pg_bigm` builds GIN bigram indexes for LIKE and similarity searches, including multibyte text. Release 1.2-20261118 keeps SQL version 1.2 and adds PostgreSQL 19 support.

### Search

```sql
CREATE EXTENSION pg_bigm;
CREATE TABLE documents (id bigint PRIMARY KEY, content text);
INSERT INTO documents VALUES (1, 'PostgreSQL database search');
CREATE INDEX docs_bigm_idx ON documents USING gin (content gin_bigm_ops);
SELECT * FROM documents WHERE content LIKE likequery('database');
SELECT show_bigm('PostgreSQL');
SET pg_bigm.similarity_limit = 0.2;
SELECT * FROM documents WHERE content =% 'database search';
```

`likequery(text)` escapes a literal search keyword for LIKE. `show_bigm(text)` displays its bigrams, and `bigm_similarity(text,text)` returns a similarity value. The `=%` operator applies `pg_bigm.similarity_limit`, normally 0.3. An index is useful only when the planner selects it.

### Settings and maintenance

`pg_bigm.enable_recheck` normally remains on: disabling rechecks can return false positives. `pg_bigm.gin_key_limit` defaults to 0 (no limit); a positive value bounds keys used in a search. `pg_gin_pending_stats(regclass)` reports GIN pending-list pages and tuples. `pg_bigm.last_update` identifies the library revision.

The 20261118 build fixes incorrect multibyte truncation, LIKE underscore handling and missing similarity matches. The truncation fix requires compilation with headers from PostgreSQL 14.21, 15.16, 16.12, 17.8, 18.2 or later, or PostgreSQL 19; the same extension SQL version does not prove the library contains the correction. Deploy the matching library to all servers and replace processes using the old library. Base use requires no preload.
