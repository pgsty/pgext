## Usage

Sources:

- [Official upstream README](https://github.com/softvisio/postgresql-softvisio-types/blob/d126b0f53d517a0c5677327618659f4de6b3d257/README.md)
- [Official extension control file (softvisio_types.control)](https://github.com/softvisio/postgresql-softvisio-types/blob/d126b0f53d517a0c5677327618659f4de6b3d257/softvisio_types.control)
- [Official extension SQL (softvisio_types--1.1.0.sql)](https://github.com/softvisio/postgresql-softvisio-types/blob/d126b0f53d517a0c5677327618659f4de6b3d257/softvisio_types--1.1.0.sql)

`softvisio_types` — PostgreSQL additional types extension. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION softvisio_types;

CREATE EXTENSION IF NOT EXISTS softvisio_types;

ALTER EXTENSION softvisio_types UPDATE;

DROP EXTENSION IF EXISTS softvisio_types;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `lo_size(oid)` is an extension function and returns `int53`.
- `int53` is an extension-defined domain.
- `number1` is an extension-defined domain.
- `number10` is an extension-defined domain.
- `number11` is an extension-defined domain.
- `number12` is an extension-defined domain.
- `number13` is an extension-defined domain.
- `number14` is an extension-defined domain.
- `number15` is an extension-defined domain.
- `number16` is an extension-defined domain.
- `number2` is an extension-defined domain.
- `number3` is an extension-defined domain.
- `number4` is an extension-defined domain.
- `number5` is an extension-defined domain.

### Requirements and Caveats

- The reviewed control file declares default version `1.2.0`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
