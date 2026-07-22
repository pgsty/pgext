## Usage

Sources:

- [Official README](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/README.md)
- [Extension SQL](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/file_textarray_fdw.sql)
- [FDW implementation](https://github.com/adunstan/file_text_array_fdw/blob/dde60712f74a249e9d3675f8c33ee18c33d23bd1/file_textarray_fdw.c)

`file_textarray_fdw` version `1.0.1` reads server-side files whose rows may have different numbers of fields. Each parsed row is returned through one `text[]` column instead of requiring a fixed foreign-table column list.

### Core Workflow

```sql
CREATE EXTENSION file_textarray_fdw;
CREATE SERVER file_server FOREIGN DATA WRAPPER file_textarray_fdw;

CREATE FOREIGN TABLE agg_csv_array (t text[])
SERVER file_server
OPTIONS (
  format 'csv',
  filename '/path/to/agg.csv',
  header 'true',
  delimiter ';',
  quote '@',
  escape '"',
  null ''
);

SELECT t[1]::int2 AS a, t[2]::float4 AS b
FROM agg_csv_array;
```

The sole foreign-table column must be `text[]`. Fields are placed into the array in file order, so callers can index and cast only the positions they need.

### Options and Boundaries

The wrapper accepts the file parsing options implemented by PostgreSQL's file FDW machinery, including `filename`, `format`, `header`, `delimiter`, `quote`, `escape`, and `null`. The filename is on the PostgreSQL server, not the client. The server operating-system account must be able to read it.

This FDW implements scans only: it does not write back to the file. Grant server and foreign-table privileges carefully, because access can expose files readable by the PostgreSQL service account. No user mapping, preload, or restart is required.
