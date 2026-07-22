## Usage

Sources:

- [Official pg_bleve README](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/README.md)
- [Version 1.0.0 extension SQL](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/pg_bleve--1.0.0.sql)
- [Go implementation](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/main.go)
- [PostgreSQL index access-method wrapper](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/pg_wrapper.c)

`pg_bleve` is an early PostgreSQL 17 experiment that calls the Go Bleve search library through cgo. Its functional surface is a set of manual index-management functions backed by files in `/tmp/pg_bleve_indexes`; the declared PostgreSQL `bleve` index access method is only a nonfunctional skeleton and does not maintain a searchable SQL index.

### Manual indexing workflow

```sql
CREATE EXTENSION pg_bleve;

SELECT bleve_create_index('products');

SELECT bleve_index_document(
    'products',
    'sku-42',
    '{"title":"Wireless headphones","category":"electronics"}'
);

SELECT bleve_search('products', 'category:electronics');
SELECT 'products' @@@ 'title:wireless';
```

`bleve_create_index(index_name, config_json)` creates an on-disk Bleve index, but this release ignores `config_json`. `bleve_index_document(index_name, doc_id, document_json)` inserts or replaces one JSON document. `bleve_search(index_name, query_string)` returns search results as JSON text and hard-codes a maximum of ten hits.

The `@@@` operator performs a direct query. `|||` combines whitespace-separated terms as a disjunction, and `&&&` combines them as a conjunction. `bleve_populate_index(index_name, table_name)` serializes table rows but processes at most 1,000 rows; `bleve_populate_index_batch(index_name, table_name, batch_size)` is the experimental batch variant. Verify actual counts and results after either helper.

### Critical limitations

Do not rely on `CREATE INDEX ... USING bleve`. In this revision the build callback creates an empty external index without indexing heap tuples, the insert callback does no indexing, scans return no tuples, and vacuum callbacks do no maintenance. PostgreSQL updates and deletes are therefore not synchronized automatically.

The Bleve files are outside PostgreSQL transactions, MVCC, WAL, backup, recovery, and privilege enforcement. `/tmp` may be cleared, index names are used in filesystem paths, and names can collide across databases or clusters on the same host. Several side-effecting or filesystem-reading functions are also declared `IMMUTABLE`, which does not match their behavior. Treat the extension as disposable lab code only; never infer database durability or consistency from a successful function call.
