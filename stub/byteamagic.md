## Usage

> Sources: [GitHub repo](https://github.com/nmandery/pg_byteamagic), [byteamagic docs](https://raw.githubusercontent.com/nmandery/pg_byteamagic/master/doc/byteamagic.md)
> Extension name: `byteamagic`
> The CSV package row is `pg_byteamagic`; the upstream extension name is `byteamagic`.

`byteamagic` uses `libmagic` on `bytea` values to identify the MIME type and the file type text inside PostgreSQL.

```sql
CREATE EXTENSION byteamagic;
SELECT byteamagic_mime(data);
SELECT byteamagic_text(data);
```

### Functions

- `byteamagic_mime(bytea)` returns the MIME type as text and matches `file --mime-type`.
- `byteamagic_text(bytea)` returns a human-readable file type description and matches `file`.

### Notes

- The extension needs PostgreSQL development headers and the `libmagic` development package.
- It is intended for databases that store files or blobs in `bytea` and need in-database type detection.
