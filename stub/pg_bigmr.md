## Usage

Sources:

- [Official extension control file](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/pg_bigmr.control)
- [Official upstream documentation](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/README.md)
- [Official SQL definition](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/sql/pg_bigmr--0.1.0.sql)
- [Official Rust implementation](https://github.com/shinyaaa/pg_bigmr/blob/cb85c4eab90a8a41aa81492751713b683e20e084/src/lib.rs)

`pg_bigmr` 0.1.0 is an experimental Rust and pgrx reimplementation of `pg_bigm`. It provides bigram-based text similarity and a GIN operator class for accelerating `LIKE` and similarity searches, but upstream explicitly labels it work in progress.

### Core Workflow

Add the module to `shared_preload_libraries`, restart PostgreSQL, create the extension, and build a GIN index with `gin_bigm_ops`.

```conf
shared_preload_libraries = 'pg_bigmr'
```

```sql
CREATE EXTENSION pg_bigmr;

CREATE INDEX documents_body_bigm_idx
  ON documents USING gin (body gin_bigm_ops);

SELECT * FROM documents WHERE body LIKE '%database%';
SELECT bigm_similarity('database', 'databass');
```

`show_bigm(text)` displays extracted bigrams, `likequery(text)` escapes a literal for contains-style `LIKE`, and `pg_gin_pending_stats(oid)` reports pending-list pages and tuples for a GIN index. The `=%` operator applies the similarity threshold.

### Settings and Caveats

`pg_bigmr.similarity_limit` defaults to 0.3, `pg_bigmr.gin_key_limit` defaults to no limit, and `pg_bigmr.enable_recheck` defaults to on. Changing limits or disabling heap rechecks can trade correctness or selectivity for speed. This is not a drop-in assurance of `pg_bigm` compatibility; benchmark representative text, verify Unicode and wildcard behavior, and avoid production use without workload-specific validation.
