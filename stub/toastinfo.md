


## Usage

Sources: [upstream README](https://github.com/credativ/toastinfo), [upstream tags](https://github.com/credativ/toastinfo/tags), [PGDG package metadata via local `db/extension.csv`](../db/extension.csv).

`toastinfo` exposes how PostgreSQL stores variable-length (`varlena`) values, including inline, compressed, and out-of-line TOAST storage forms.

### Functions

`pg_toastinfo(anyelement)` describes the storage form of a datum:

```sql
CREATE EXTENSION toastinfo;

SELECT a, length(b), pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;
```

Possible storage forms include:

- `null`, for NULL values.
- `ordinary`, for non-varlena data types.
- `short inline varlena`, up to 126 bytes with a 1-byte header.
- `long inline varlena, uncompressed`, up to 1 GiB with a 4-byte header.
- `long inline varlena, compressed (pglz|lz4)`.
- `toasted varlena, uncompressed`.
- `toasted varlena, compressed (pglz|lz4)`.

Compressed varlenas show the compression method on PostgreSQL 14+.

`pg_toastpointer(anyelement)` returns the TOAST table `chunk_id` OID for out-of-line toasted values, or NULL for non-toasted input:

```sql
SELECT pg_toastpointer(large_column)
FROM my_table;
```

### Storage Example

```sql
CREATE TABLE t (a text, b text);

ALTER TABLE t ALTER COLUMN b SET STORAGE EXTERNAL;
INSERT INTO t VALUES ('external-10000', repeat('x', 10000));

ALTER TABLE t ALTER COLUMN b SET STORAGE EXTENDED;
INSERT INTO t VALUES ('extended-1000000', repeat('x', 1000000));

ALTER TABLE t ALTER COLUMN b SET COMPRESSION lz4;
INSERT INTO t VALUES ('extended-lz4', repeat('x', 1000000));

SELECT a, pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;
```

### Caveats

- Pigsty metadata carries `toastinfo` 1.6 for PostgreSQL 14-18, with Pigsty RPMs and PGDG DEBs.
- The upstream GitHub README documents the same SQL surface, but the public GitHub tags visible during review stop at `v1.5`; no upstream 1.6 changelog was found in that repository.
- `pg_toastpointer` is meaningful only for out-of-line toasted data; ordinary, inline, or NULL values return NULL.
