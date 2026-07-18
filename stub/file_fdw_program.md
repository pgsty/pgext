## Usage

Sources:

- [file_fdw_program documentation at the reviewed commit](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/doc/file_fdw_program.md)
- [file_fdw_program README at the reviewed commit](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/README.md)
- [Current PostgreSQL file_fdw documentation](https://www.postgresql.org/docs/current/file-fdw.html)

`file_fdw_program` is a historical backport of the `program` option for `file_fdw`. Instead of reading a named file, a foreign table runs a server-side command and parses its standard output with COPY-style format options.

### Create a Program-Backed Foreign Table

```sql
CREATE EXTENSION file_fdw_program;
CREATE SERVER file_program_server
  FOREIGN DATA WRAPPER file_fdw_program;

CREATE FOREIGN TABLE generated_rows (
  one text,
  two text,
  three text
)
SERVER file_program_server
OPTIONS (program 'echo one,two,three', format 'csv');

SELECT * FROM generated_rows;
```

The command runs when the foreign table is scanned. The available format options follow the extension's `file_fdw`-like interface.

### Caveats

- PostgreSQL 10 and later include `program` support in the bundled `file_fdw`; use the bundled extension on modern PostgreSQL instead of this version 1.0.1 backport.
- A `program` is executed by the PostgreSQL server operating-system account. Treat server and foreign-table ownership as code-execution privileges, use fixed commands, and do not interpolate untrusted input.
- Upstream describes this package as a PostgreSQL 9.4-or-newer backport from 2016. It does not document support for current PostgreSQL releases.
