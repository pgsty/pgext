## Usage

Sources:

- [postgres/resources/parquet_io.md](https://github.com/minyeamer/linkmerce/blob/a662f8ab7646c677bfd9d2112b0c0a6556ca396b/postgres/resources/parquet_io.md)
- [postgres/extension/parquet_io.control](https://github.com/minyeamer/linkmerce/blob/a662f8ab7646c677bfd9d2112b0c0a6556ca396b/postgres/extension/parquet_io.control)
- [postgres/extension/parquet_io--1.0.sql](https://github.com/minyeamer/linkmerce/blob/a662f8ab7646c677bfd9d2112b0c0a6556ca396b/postgres/extension/parquet_io--1.0.sql)
- [postgres/extension/parquet_io.cpp](https://github.com/minyeamer/linkmerce/blob/a662f8ab7646c677bfd9d2112b0c0a6556ca396b/postgres/extension/parquet_io.cpp)

`parquet_io` 1.0 is the LinkMerce component for moving data between PostgreSQL and Parquet using Apache Arrow C++. The upstream environment targets PostgreSQL 18. Installation requires a superuser; no shared preload is needed.

### Core workflow

```sql
CREATE EXTENSION parquet_io;
SELECT parquet_create('/srv/import/input.parquet', 'public.imported_rows');
SELECT parquet_read('/srv/import/input.parquet', 'public.imported_rows');
SELECT octet_length(parquet_write('SELECT * FROM public.imported_rows'));
```

### Interface and boundaries

`parquet_create` creates an empty table and returns the number of columns; it does not import rows. Its default conflict mode is `error`; `replace` drops the target table. `parquet_read` appends rows to an existing table and returns the row count. Both accept either a server path or a `bytea` value. `parquet_write` returns Parquet bytes when given only SQL, or writes a server file and returns a row count when also given a path.

Paths are interpreted by the database server, with its OS permissions. File output is outside PostgreSQL transaction rollback. Restrict function execution to trusted roles and review SQL arguments and replacement targets. The installation script does not revoke the default function grants. Large byte values and Arrow conversion require server memory; type mappings may simplify the input schema.
