## Usage

Sources:

- [Release 4.1.0](https://github.com/pgroonga/pgroonga/releases/tag/4.1.0)
- [Control](https://github.com/pgroonga/pgroonga/blob/47d0252f76dccbdc2ebeee9ed6cf941dd9552bcc/pgroonga.control)

- [Version 4.1.0 SQL](https://github.com/pgroonga/pgroonga/blob/4.1.0/data/pgroonga.sql)
- [Version 4.1.0 control](https://github.com/pgroonga/pgroonga/blob/4.1.0/pgroonga.control)
- [Official tutorial](https://pgroonga.github.io/tutorial/)
- [Version 4.1.0 release](https://github.com/pgroonga/pgroonga/releases/tag/4.1.0)
- [Upgrade guidance](https://pgroonga.github.io/upgrade/)

`pgroonga` 4.1.0 provides Groonga-backed indexes for multilingual full-text search. It installs the `pgroonga` access method and SQL operators; ordinary use does not require shared preload.

### Core Workflow

Install compatible PGroonga and Groonga libraries, then create the extension as an administrator:

```sql
CREATE EXTENSION pgroonga;
CREATE TABLE search_notes (id bigint PRIMARY KEY, body text);
CREATE INDEX search_notes_body_idx ON search_notes USING pgroonga (body);
INSERT INTO search_notes VALUES (1, 'PostgreSQL supports full text search');
SELECT id, body FROM search_notes WHERE body &@ 'PostgreSQL';
SELECT id, body, pgroonga_score(tableoid, ctid) AS score
FROM search_notes WHERE body &@~ 'PostgreSQL OR Groonga'
ORDER BY score DESC;
```

### Important Objects

- `&@` matches a keyword; `&@~` accepts Groonga query syntax. Supported LIKE/ILIKE searches can also use the index, with rechecks where required.
- `pgroonga_score(tableoid, ctid)` retrieves search scores. Confirm the intended index plan when using score-based ordering.
- `pgroonga_highlight_html()` and `pgroonga_query_extract_keywords()` produce highlighted search results; `pgroonga_snippet_html()` provides surrounding text.
- New in 4.0.9, `pgroonga_physical_table_names(partitioned_index, prefix)` returns a text array of Groonga command arguments identifying the physical tables behind partition indexes. This release also increments `pg_stat_user_indexes.idx_scan` for PGroonga scans.

### Maintenance and Privileges

The 4.1.0 control file declares neither trusted installation nor relocatability; do not assume ordinary users can install it or move it between schemas. Index creation and queries follow the relevant table privileges. Match the extension and Groonga libraries to the target PostgreSQL build and follow upstream upgrade guidance before replacing binaries.

PGroonga manages derived index files in addition to table data. Plan disk capacity and backup/recovery procedures accordingly. Use REINDEX for index repair where appropriate. The separate `pgroonga_database` module is a recovery tool for damaged internal Groonga databases and is not needed for normal searches. Do not disable sequential scans globally merely to force an index in production.

### Pigsty Runtime Compatibility

The current Pigsty EL8 and EL9 packages have a confirmed coexistence limit with PostGIS Raster on both x86_64 and aarch64, across PostgreSQL 14–18. Groonga uses Arrow 22, while the tested GDAL/Raster stacks use Arrow 8 on EL8 and Arrow 9 on EL9. Loading both stacks into one PostgreSQL backend can crash it; a successful SQL query does not prove a normal backend exit.

Changing load order is not a complete fix: automatic session preload still produced backend-exit crashes on EL9 aarch64. Avoid enabling both stacks in the same backend until a compatible dependency combination is verified; isolate their use when necessary. The tested EL10 and Debian/Ubuntu combinations did not reproduce this failure, which does not establish compatibility for arbitrary other dependency versions. This is a Pigsty package-stack boundary, not an upstream requirement to preload PGroonga for ordinary search.

MeCab tokenization additionally needs the matching Groonga tokenizer plugin and dictionary. Installing the PostgreSQL extension alone does not supply every optional tokenizer; verify the requested tokenizer before building an index that names it.

### Version 4.1.0

Version 4.1.0 fixes another session running VACUUM/ANALYZE before CREATE INDEX commits, which could remove index data. An index already affected may need REINDEX. Direct Groonga calls can use `pgroonga_tuple_is_alive` in `post_filter` with Groonga 16.1.5 or later. This source release does not change the separately documented Pigsty dependency-stack limitation.
