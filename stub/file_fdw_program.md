## Usage

Sources:

- [Official upstream documentation](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/doc/file_fdw_program.md)
- [Official extension control file](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/file_fdw_program.control)
- [Official implementation](https://github.com/corlogic-code/file_fdw_program/blob/412b751c6af32882f335515e5af1337000624ce2/src/file_fdw_program.c)

`file_fdw_program` 1.0.1 is a PostgreSQL 9.4-era, read-only FDW that feeds an external command's standard output through COPY parsing. It backported the `program` option later incorporated into core `file_fdw` for PostgreSQL 10. On supported modern PostgreSQL releases, prefer the built-in `file_fdw` implementation.

### Core Workflow

Only a superuser can define or change the command-bearing foreign-table options:

```sql
CREATE EXTENSION file_fdw_program;
CREATE SERVER local_program FOREIGN DATA WRAPPER file_fdw_program;

CREATE FOREIGN TABLE generated_rows (
  one text,
  two text,
  three text
)
SERVER local_program
OPTIONS (
  program 'printf "one,two,three\n"',
  format 'csv'
);

SELECT * FROM generated_rows;
```

Use either `program` or `filename`, never both. Other table and column options follow the COPY FROM parsing options supported by the target server version.

### Execution Boundary

The command is executed by the PostgreSQL server under the operating-system account running the database, not by the SQL client. It inherits server-side filesystem, environment, network, and process privileges. The validator restricts foreign-table option changes to superusers, but selecting the table runs the saved command; tightly control ownership, grants, and dependency chains so untrusted roles cannot cause privileged command execution or alter referenced objects.

### Compatibility and Operations

The wrapper is read-only and cannot estimate command output from filesystem metadata, so plans use fallback estimates. Commands must produce stable rows matching the declared types and COPY format; nonzero exits, stderr, encoding errors, or malformed rows fail the scan. Avoid secrets in command text because catalog metadata and logs may expose it. Set operating-system limits, absolute executable paths, a minimal environment, and explicit timeouts outside the extension. This project is deprecated and targets old PostgreSQL APIs; use it only for a pinned legacy server after testing cancellation, repeated scans, parallel plans, command failure, and privilege behavior.
