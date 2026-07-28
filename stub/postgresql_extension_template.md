## Usage

Sources:

- [Official upstream README](https://github.com/cybertec-postgresql/postgresql_extension_template/blob/8dcf2fa61c23185e2dedec0bc21b1402ff109def/README.md)
- [Official extension control file (postgresql_extension_template.control)](https://github.com/cybertec-postgresql/postgresql_extension_template/blob/8dcf2fa61c23185e2dedec0bc21b1402ff109def/postgresql_extension_template.control)
- [Official extension SQL (postgresql_extension_template--1.0.sql)](https://github.com/cybertec-postgresql/postgresql_extension_template/blob/8dcf2fa61c23185e2dedec0bc21b1402ff109def/postgresql_extension_template--1.0.sql)

`postgresql_extension_template` — This is a template repository for a PostgreSQL C extension. This repository includes:. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION postgresql_extension_template;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `my_function()` is an extension function and returns `cstring`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
