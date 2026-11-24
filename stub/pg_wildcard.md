## Usage

Sources:

- [Published crate 0.2.1](https://crates.io/crates/pg_wildcard/0.2.1)
- [Extension README](https://github.com/f1b0-gh/pg_wildcard/blob/14db6f91c988a23ac3f2cd6664986198045743b8/pg_wildcard/README.md)
- [Control](https://github.com/f1b0-gh/pg_wildcard/blob/14db6f91c988a23ac3f2cd6664986198045743b8/pg_wildcard/pg_wildcard.control)
- [SQL API](https://github.com/f1b0-gh/pg_wildcard/blob/14db6f91c988a23ac3f2cd6664986198045743b8/docs/sql-api.md)
- [Foreign data wrapper](https://github.com/f1b0-gh/pg_wildcard/blob/14db6f91c988a23ac3f2cd6664986198045743b8/docs/fdw.md)
- [Upgrade SQL](https://github.com/f1b0-gh/pg_wildcard/blob/14db6f91c988a23ac3f2cd6664986198045743b8/pg_wildcard/sql/pg_wildcard--0.2.0--0.2.1.sql)

`pg_wildcard` 0.2.1 searches telemetry strings with wildcard and regular-expression predicates, derives indicators and generates ClickHouse/Antalya queries. It also supplies the scan-only `wildcard_fdw`. The published crate uses `antalya-rs` 0.1.0; PostgreSQL 18 is the qualified server major. Declared PostgreSQL 14–17 build features require separate qualification.

### Local String Search

A superuser installs this untrusted, non-relocatable extension. No shared preload or server restart is documented for ordinary use.

```sql
CREATE EXTENSION pg_wildcard;
SELECT wc_match('https://example.org/admin', '*/admin', 'raw');
SELECT wc_match('GET /%77p-ADMIN/a.PHP', '*/wp-admin/*.php', 'decoded');
SELECT * FROM wc_indicator('hxxps://Login.Evil[.]com/');
SELECT wc_explain('*/wp-admin/*.php', 'wildcard', 'decoded');
```

Views include `raw`, `lower`, `decoded`, `revdomain` and `winpath`. `wc_match`, `wc_like` and `wc_regex` match text; `wc_regex_bytes` handles bytes. `wc_url_decode` preserves decoded bytes, while `wc_view` produces display text. `wc_explain` describes conservative candidates; an exact predicate must still decide the match. Regex semantics follow the supported RE2 dialect.

### Remote Queries and Lake Metadata

`wc_ch_predicate` and related helpers return SQL for the caller to execute. Their optional filter expressions are trusted SQL fragments. `wc_lake_index`, `wc_lake_file` and `wc_lake_gram` hold imported pruning metadata; candidate files always require an exact recheck, and missing coverage must widen the scan.

`wildcard_fdw` requires an Antalya endpoint, a user mapping with `bearer_token`, compatible scalar columns and a UTF-8 PostgreSQL database. Generic ClickHouse HTTP availability is insufficient. The mapping principal needs source access and the documented query-completion evidence, normally from `system.query_log`. Protect mapping credentials and refresh tokens separately.

The wrapper pushes supported predicates but retains local checks. Joins and aggregates stay local; foreign tables have no local PostgreSQL index. Unsupported collation or query shapes limit pushdown. A late remote or completion error fails the statement even after some rows have been streamed.

### Resource Limits and Upgrades

`pg_wildcard.matcher_cache_entries` and `pg_wildcard.matcher_cache_kb` bound retained patterns per backend; `wc_matcher_cache_stats()` reports this cache. These limits do not cap total backend memory.

From 0.2.0, install the matching 0.2.1 library and SQL files, reconnect, then run:

```sql
ALTER EXTENSION pg_wildcard UPDATE TO '0.2.1';
```

The update preserves SQL signatures and records the new extension version. Standalone library, lake CLI and remote-engine configuration are separate components.
