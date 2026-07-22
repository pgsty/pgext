## Usage

Sources:

- [Official extension SQL](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/hello_fdw--1.0.sql)
- [Official C implementation](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/hello_fdw.c)
- [Official README](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/README.md)

`hello_fdw` is a minimal teaching example for PostgreSQL's Foreign Data Wrapper API. Every scan returns exactly one row and feeds the string `Hello,World` to every declared column; it does not connect to an external system or provide a general-purpose data source.

### Core Workflow

```sql
CREATE EXTENSION hello_fdw;

CREATE SERVER hello_server
    FOREIGN DATA WRAPPER hello_fdw;

CREATE FOREIGN TABLE hello_demo (
    first_value text,
    second_value text
) SERVER hello_server;

SELECT * FROM hello_demo;
```

The result contains one row, with `Hello,World` in both columns. A rescan resets the one-row iterator.

### Implemented Surface

- `hello_fdw_handler` registers planning, scan, rescan, explain, and analyze callbacks.
- `hello_fdw_validator` is a no-op, so it does not validate wrapper, server, mapping, or table options.
- `EXPLAIN` adds a fixed `Hello` property.
- The implementation estimates one row and exposes no insert, update, delete, import-schema, or remote-join callback.

### Caveats

Values are passed through each column type's text input function. Non-text columns may accept `Hello,World` differently or reject it, so use text columns for the demonstration. Options do not alter the returned data even though the no-op validator accepts them.

This extension is sample code, not an operational FDW. It has only a handful of old commits, does no option validation, performs no remote I/O, and implements a historical FDW API. Use it to study callbacks in a disposable environment; do not use it as a template without updating memory, planner, API-compatibility, error-handling, and write-path assumptions for the target PostgreSQL version.
