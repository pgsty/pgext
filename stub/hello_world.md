## Usage

Sources:

- [Official upstream README](https://github.com/andrelandgraf/russian-roulette-hello-world-postgres-extension/blob/0ab27d4357ce50a66d930c08ed5f16ba6dfbbd1b/README.md)
- [Official extension control file (hello_world.control)](https://github.com/andrelandgraf/russian-roulette-hello-world-postgres-extension/blob/0ab27d4357ce50a66d930c08ed5f16ba6dfbbd1b/hello_world.control)
- [Official extension SQL (hello_world--1.0.sql)](https://github.com/andrelandgraf/russian-roulette-hello-world-postgres-extension/blob/0ab27d4357ce50a66d930c08ed5f16ba6dfbbd1b/hello_world--1.0.sql)

`hello_world` — A simple PostgreSQL extension that demonstrates basic extension functionality using SQL only (no C code compilation required). Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION hello_world;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hello_world()` is an extension function and returns `TEXT`.
- `hello_world(name TEXT)` is an extension function and returns `TEXT`.
- `log_hello_world(message TEXT DEFAULT 'Hello, World!')` is an extension function and returns `INTEGER`.
- `russian_roulette(target_user_id INTEGER, odds INTEGER DEFAULT 6)` is an extension function and returns `TEXT`.
- `hello_world_log` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
