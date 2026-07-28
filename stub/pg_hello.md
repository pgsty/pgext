## Usage

Sources:

- [Official upstream README](https://github.com/magnusp/pg_hello/blob/83ab1b18a51ff0d0d42d3639caf79b5885bb7040/README.md)
- [Official extension control file (pg_hello.control)](https://github.com/magnusp/pg_hello/blob/83ab1b18a51ff0d0d42d3639caf79b5885bb7040/pg_hello.control)
- [Official extension SQL (pg_hello--1.0.sql)](https://github.com/magnusp/pg_hello/blob/83ab1b18a51ff0d0d42d3639caf79b5885bb7040/pg_hello--1.0.sql)

`pg_hello` — pg_hello, a very basic Postgres extension =========================================. Use it when an application needs this specific database capability. The reviewed upstream material marks this capability deprecated.

### Core Workflow

```sql
CREATE EXTENSION pg_hello;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hello(TEXT)` is an extension function and returns `TEXT`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream material contains an explicit deprecation boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
