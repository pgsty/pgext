## Usage

Sources:

- [Official README](https://github.com/hyperiondb/search/blob/6dc33192199ed18a1be9d269b88a9027c76f2eac/README.md)
- [Extension control file](https://github.com/hyperiondb/search/blob/6dc33192199ed18a1be9d269b88a9027c76f2eac/hsearch.control)
- [pgrx manifest](https://github.com/hyperiondb/search/blob/6dc33192199ed18a1be9d269b88a9027c76f2eac/Cargo.toml)

`hsearch` provides a Tantivy-backed `bm25` index access method for PostgreSQL 18 while storing index bytes in PostgreSQL pages protected by generic WAL.

### Enablement

Install hsearch 0.4.1 for PostgreSQL 18, add it to `shared_preload_libraries`, and restart before creating the extension:

```ini
shared_preload_libraries = 'hsearch'
```

```sql
CREATE EXTENSION hsearch;
```

The extension is non-relocatable and superuser-only. Its SQL objects live in the `hyper` schema.

### Build and Query an Index

The first indexed column is the stable key named by `key_field`; other indexed expressions are searchable text fields.

```sql
CREATE TABLE items (
  id      varchar(24) PRIMARY KEY,
  name    text,
  summary text
);

CREATE INDEX items_bm25 ON items
USING bm25 (
  id,
  (name::hyper.ngram(2,5,'ascii_folding=true')),
  (summary::hyper.ngram(2,5,'ascii_folding=true'))
) WITH (key_field='id');

SELECT id, hyper.score(id) AS score
FROM items
WHERE name &&& 'postgres'
ORDER BY score DESC
LIMIT 20;
```

`&&&` requires all generated ngrams and rechecks heap visibility. Rows inserted in the current transaction become searchable at commit.

### Operations and Limits

The index is WAL-logged for crash recovery and physical replication, but it remains a derived structure: use `REINDEX` or `hyper.reindex_all()` after corruption or when operational policy calls for rebuilding. `hsearch.max_matches` caps candidates per scan key (default 1,000); increasing it improves deep filtered recall at higher memory and latency cost. Unlogged tables and tokenizer settings other than `(2,5,'ascii_folding=true')` are not supported in 0.4.1.

