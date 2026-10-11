## Usage

Sources:

- [README.md](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/README.md)
- [Cargo.toml](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/Cargo.toml)
- [src/lib.rs](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/src/lib.rs)
- [sql/test_bm25.sql](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/sql/test_bm25.sql)
- [pg_gse.control](https://github.com/smoothdvd/pg_gse/blob/54346cf5cae9620116aba8744eef1b9bdf808ef0/pg_gse.control)

`pg_gse` 0.1.0 supplies a Chinese text search parser and configuration using Jieba. Tokenization works on its own; BM25 ranking is an optional integration with `pg_textsearch`. Upstream reports live verification on PostgreSQL 17 and 18.

### Core Workflow

```sql
CREATE EXTENSION pg_gse;
SELECT * FROM gse_tokenize('自然语言处理');
SELECT * FROM gse_tokenize('自然语言处理', 'search');
SELECT to_tsvector('public.gse_chinese', '自然语言处理');
```

### Operational Boundaries

`gse_tokenize()` returns tokens and byte offsets. Its two-argument overload accepts `precise`, `search` or `all`; `gse_set_mode()` and `gse_get_mode()` manage the current backend mode. The `gse` parser and `gse_chinese` configuration feed normal PostgreSQL text search. Installation requires a superuser; the extension is not relocatable and has no preload requirement itself. For BM25, install `pg_textsearch`, preload that library and restart as its documentation requires, then use a schema-qualified configuration such as `public.gse_chinese`. The custom dictionary table `pg_gse_custom_word` exists, but `gse_sync_custom_word()` is still a no-op for dictionary changes in this revision. Do not rely on the README dictionary example as a working feature. Cargo exposes PG14–18 features, while the explicit runtime verification covers 17/18.
