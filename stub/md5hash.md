

## Usage

> [md5hash: native MD5 hash data type for PostgreSQL](https://github.com/tvondra/md5hash)

The `md5hash` extension provides an efficient 128-bit data type for storing MD5 hashes in binary format (16 bytes) instead of text (32+ bytes).

```sql
CREATE EXTENSION md5hash;

CREATE TABLE test_table (
    id md5hash PRIMARY KEY
);

INSERT INTO test_table VALUES ('c4ca4238a0b923820dcc509a6f75849b');

SELECT * FROM test_table
WHERE id = 'c4ca4238a0b923820dcc509a6f75849b';
```

### Operators

Standard comparison operators are supported: `=`, `<>`, `<`, `>`, `<=`, `>=`.

### Index Support

Btree index operator class is included, enabling efficient lookups and primary key constraints on `md5hash` columns.

### Storage Benefits

Compared to storing MD5 as `text`, the `md5hash` type uses approximately 60% of the storage space and provides faster indexed lookups.
