## Usage

Sources: [official repo](https://github.com/nmandery/pg_byteamagic), [official doc](https://raw.githubusercontent.com/nmandery/pg_byteamagic/master/doc/byteamagic.md)

`byteamagic` runs `libmagic` on `bytea` values, so PostgreSQL can identify the MIME type or human-readable file type of blobs stored in the database. The package name is `pg_byteamagic`, but the extension name is `byteamagic`.

```sql
CREATE EXTENSION byteamagic;

SELECT byteamagic_mime(data) FROM files;
SELECT byteamagic_text(data) FROM files;
```

### Functions

- `byteamagic_mime(bytea) returns text`: MIME type, equivalent to `file --mime-type`.
- `byteamagic_text(bytea) returns text`: descriptive file type text, equivalent to `file`.

### Common Use

```sql
SELECT
  id,
  byteamagic_mime(blob) AS mime,
  byteamagic_text(blob) AS kind
FROM uploads;
```

### Caveats

- It only inspects `bytea` content; there are no operators, custom types, or extra SQL management objects.
- Build/install requires PostgreSQL development headers plus the system `libmagic` development package.
- The upstream documentation is minimal; current user-facing behavior is unchanged from the long-standing doc page.
