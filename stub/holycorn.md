## Usage

Sources:

- [Official README](https://github.com/franckverrot/holycorn/blob/6cdfa71caffa28ba3ca720f253c4d5e112d5fa1d/README.md)
- [Official extension SQL](https://github.com/franckverrot/holycorn/blob/6cdfa71caffa28ba3ca720f253c4d5e112d5fa1d/holycorn--1.0.sql)

`holycorn` is a foreign-data wrapper host built around the sandboxed mruby VM. It can read Redis through the bundled wrapper or execute a server-side Ruby producer supplied for a foreign table.

### Core Workflow

The bundled Redis wrapper is selected with `wrapper_class`; its remaining table options are passed to the Ruby wrapper constructor.

```sql
CREATE EXTENSION holycorn;
CREATE SERVER holycorn_server FOREIGN DATA WRAPPER holycorn;

CREATE FOREIGN TABLE redis_table (key text, value text)
  SERVER holycorn_server
  OPTIONS (
    wrapper_class 'HolycornRedis',
    host '127.0.0.1',
    port '6379',
    db '0'
  );

SELECT * FROM redis_table;
```

For bulk definition, `IMPORT FOREIGN SCHEMA` invokes the wrapper's import callback. The upstream Redis example requires the remote schema name `holycorn_schema` and recommends a `prefix` to avoid local name collisions.

```sql
CREATE SCHEMA holycorn_tables;
IMPORT FOREIGN SCHEMA holycorn_schema
  FROM SERVER holycorn_server
  INTO holycorn_tables
  OPTIONS (
    wrapper_class 'HolycornRedis',
    host '127.0.0.1',
    port '6379',
    db '0',
    prefix 'holycorn_'
  );
```

### Wrapper Contract

- `wrapper_class` names a Ruby object already compiled into the mruby runtime; the bundled class is `HolycornRedis`.
- `wrapper_path` instead points to a server-side Ruby file. Its object must accept `.new` with one environment argument and expose `each` with no arguments.
- The producer returns row values in foreign-table column order. Ruby scalars map to PostgreSQL booleans, integers, floating-point values, text, or timestamps as documented upstream.

### Limitations and Safety

The reviewed upstream requires PostgreSQL `9.4+` and documents reads only: `SELECT` works, but `INSERT` is not supported. A custom `wrapper_path` is code and file access on the database server, so only trusted administrators should control it. This is an old, narrow FDW implementation; validate its type conversions and Redis failure behavior on the exact PostgreSQL build before adoption.
