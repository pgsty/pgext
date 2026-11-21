## Usage

Sources:

- [README.md](https://github.com/junhkang/kor_search/blob/v2.0.0/localization/en/README.md)
- [CHANGELOG.md](https://github.com/junhkang/kor_search/blob/v2.0.0/CHANGELOG.md)
- [kor_search.control](https://github.com/junhkang/kor_search/blob/v2.0.0/kor_search.control)
- [kor_search--2.0.0.control](https://github.com/junhkang/kor_search/blob/v2.0.0/kor_search--2.0.0.control)
- [kor_search--1.1.0--2.0.0.sql](https://github.com/junhkang/kor_search/blob/v2.0.0/sql/kor_search--1.1.0--2.0.0.sql)

`kor_search` 2.0.0 analyzes mixed Korean/English text with Rust and Lindera/ko-dic, builds PostgreSQL tsvectors and expands query aliases from an in-database dictionary. It is tested on PostgreSQL 16–18 on Linux with UTF-8 databases; runtime does not call an external service.

### Indexed Search

Install matching native files, then create the extension as a superuser:

```sql
CREATE EXTENSION kor_search;
CREATE TABLE documents (body text NOT NULL);
INSERT INTO documents VALUES ('삼성에서 PostgreSQL을 사용한다');
CREATE INDEX documents_search ON documents USING gin (kor_vector_v1(body));
SELECT body FROM documents
WHERE kor_vector_v1(body) @@ kor_search_tsquery('PostgreSQL', 'default');
SELECT * FROM kor_analyze_v1('삼성에서 PostgreSQL을 사용한다');
```

`kor_vector_v1()` creates a positional tsvector for ordinary GIN indexing. `kor_analyze_v1()` exposes tokens. `kor_search_tsquery(query, scope)` expands the longest dictionary phrase first, ORs aliases and ANDs groups; input that produces no tokens returns NULL. A stored generated vector can avoid expensive phrase rechecks.

### Alias Dictionary

Add domain terms to `kor_search_aliases(scope, concept, term, kind)` or load the optional example dictionary. A normalized term belongs to one concept per scope. Updating aliases does not require reindexing documents. Scope selects a dictionary, not documents or user permissions. Search users need schema USAGE and dictionary SELECT.

The analyzer does not provide arbitrary substring search, automatic translation or English stemming. Limits are 512 KiB per document, 4 KiB or 64 analyzed tokens per query, eight tokens per alias and 32 expansions per concept.

### Upgrade and Privileges

New 2.0 installations do not require `pg_trgm`; the unused legacy trigram index was removed in 1.1. The four legacy search functions remain, but the indexed API is preferred for new applications. The 1.1 transition fixes interpolated dictionary SQL and fixed-buffer handling from the original release.

Install new files, reconnect existing sessions in a maintenance window, then run:

```sql
ALTER EXTENSION kor_search UPDATE TO '2.0.0';
```

The tested path passes through 1.1.0. Review and copy legacy domain terms into the new dictionary explicitly, and separately back up edits/deletions of legacy seed rows. Version 2.0 permits a custom schema at installation but cannot be relocated afterwards; its version-specific control overrides the root control. No preload is required. The native build is unavailable on managed services that prohibit module installation.
