
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION parray_gin;
> CREATE INDEX test_tags_idx ON test_table USING gin (val parray_gin_ops);
> SELECT * FROM test_table WHERE val @> ARRAY['must','contain'];
> SELECT * FROM test_table WHERE val @@> ARRAY['what%like%'];
> ```
>
> Sources: [README](https://github.com/theirix/parray_gin), [Reference](https://github.com/theirix/parray_gin/blob/master/doc/parray_gin.md)

`parray_gin` adds GIN indexing and operators for `text[]` arrays with both strict and partial matching. The upstream docs describe it as trigram-based array indexing built on `pg_trgm`'s trigram implementation.

## Indexing Arrays

The extension provides the `parray_gin_ops` operator class for GIN indexes on text arrays:

```sql
CREATE TABLE test_table(id bigserial, val text[]);
CREATE INDEX test_tags_idx ON test_table USING gin (val parray_gin_ops);
```

According to the reference docs, indexed values and queries are split into trigrams. Because the trigram index can return false positives, operator matches are rechecked after index lookup.

## Operators

### Strict Matching

`@> (text[], text[]) -> bool`

Returns true when the left-hand array contains all items from the right-hand array.

```sql
SELECT * FROM test_table WHERE val @> ARRAY['far'];
```

`<@ (text[], text[]) -> bool`

Returns true when the left-hand array is contained by the right-hand array.

```sql
SELECT * FROM test_table WHERE val <@ ARRAY['galaxy','ago','vader'];
```

### Partial Matching

`@@> (text[], text[]) -> bool`

Returns true when the left-hand array contains all right-hand items using partial matching, for example `'foobar' ~~ 'foo%'` or `'foobar' ~~ '%oo%'`.

```sql
SELECT * FROM test_table WHERE val @@> ARRAY['%ar%'];
```

`<@@ (text[], text[]) -> bool`

Returns true when the left-hand array is contained by the right-hand patterns using partial matching.

```sql
SELECT * FROM test_table WHERE val <@@ ARRAY['%ar%','vader'];
```

## Notes

The upstream docs say GIN can be used with `@>`, `<@`, `@@>`, and `<@@`. They also mention successful use on JSON arrays extracted from JSON text fields via the `json_accessors` extension.
