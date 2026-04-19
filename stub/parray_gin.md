## Usage

- Sources: [README](https://github.com/theirix/parray_gin/blob/master/README.md), [reference doc](https://github.com/theirix/parray_gin/blob/master/doc/parray_gin.md)

`parray_gin` adds a GIN operator class for `text[]` plus strict and partial-match operators. Upstream describes it as trigram-based array indexing built on `pg_trgm` trigram decomposition.

### Create The Extension And Index

```sql
CREATE EXTENSION parray_gin;

CREATE TABLE test_table (
  id  bigserial,
  val text[]
);

CREATE INDEX test_tags_idx
ON test_table
USING gin (val parray_gin_ops);
```

### Indexed Operators

The reference doc says `parray_gin_ops` can support these operators:

- `@>`: strict contains.
- `<@`: strict contained-by.
- `@@>`: partial contains, where array elements may use `LIKE` patterns.
- `<@@`: partial contained-by.

Examples:

```sql
SELECT * FROM test_table WHERE val @> ARRAY['must','contain'];
SELECT * FROM test_table WHERE val @@> ARRAY['what%like%'];
SELECT * FROM test_table WHERE val <@ ARRAY['galaxy','ago','vader'];
SELECT * FROM test_table WHERE val <@@ ARRAY['%ar%','vader'];
```

### Matching Behavior

Strict matching requires array-item equality. Partial matching allows patterns such as `'foo%'` or `'%oo%'`. Because the trigram index can return false positives, the docs note that matches are rechecked after index lookup.

### Caveats

The README says support extends through PostgreSQL 18, while the reference document still says 9.1-14. The operator and opclass behavior is consistent across both docs, but the version note is not fully synchronized upstream.
