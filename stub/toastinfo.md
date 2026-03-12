

## Usage

> [toastinfo: inspect TOAST storage details of varlena columns](https://github.com/credativ/toastinfo)

toastinfo exposes the internal storage form of variable-length (varlena) data types, showing how PostgreSQL stores each datum.

### Functions

**`pg_toastinfo(anyelement)`** -- describes the storage form of a datum:

```sql
SELECT a, length(b), pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;

        a         | length  | pg_column_size |              pg_toastinfo              | pg_toastpointer
------------------+---------+----------------+----------------------------------------+-----------------
 null             |       * |              * | null                                   |               *
 default          |       7 |              8 | short inline varlena                   |               *
 external-200     |     200 |            204 | long inline varlena, uncompressed      |               *
 external-10000   |   10000 |          10000 | toasted varlena, uncompressed          |           16427
 extended-10000   |   10000 |            125 | long inline varlena, compressed (pglz) |               *
 extended-1000000 | 1000000 |          11452 | toasted varlena, compressed (pglz)     |           16429
 extended-1000000 | 1000000 |           3936 | toasted varlena, compressed (lz4)      |           16430
```

Possible storage forms:
- `null` -- NULL values
- `ordinary` -- non-varlena datatypes
- `short inline varlena` -- up to 126 bytes (1-byte header)
- `long inline varlena, (un)compressed` -- up to 1GiB (4-byte header)
- `toasted varlena, (un)compressed` -- up to 1GiB stored in TOAST table
- Compressed varlenas show method (`pglz`, `lz4`) on PG14+

**`pg_toastpointer(anyelement)`** -- returns the `chunk_id` OID in the TOAST table, or NULL for non-toasted data:

```sql
SELECT pg_toastpointer(large_column) FROM my_table;
```
