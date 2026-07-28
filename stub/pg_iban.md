## Usage

Sources:

- [Official upstream README](https://github.com/rom8726/pg_iban/blob/1b4ea9a37823a253512f61803b1429a7037b1099/README.md)
- [Official extension control file (pg_iban.control)](https://github.com/rom8726/pg_iban/blob/1b4ea9a37823a253512f61803b1429a7037b1099/pg_iban.control)
- [Official extension SQL (pg_iban--1.0--1.1.sql)](https://github.com/rom8726/pg_iban/blob/1b4ea9a37823a253512f61803b1429a7037b1099/pg_iban--1.0--1.1.sql)

`pg_iban` — pg_iban is a PostgreSQL extension providing an IBAN (International Bank Account Number) data type and several utility functions for validation and manipulation of IBANs. Use it when application data needs this type, domain, or its operators. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_iban;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `iban_bban(iban)` is an extension function and returns `text`.
- `iban_cmp(iban, iban)` is an extension function and returns `int4`.
- `iban_country(iban)` is an extension function and returns `text`.
- `iban_eq(iban, iban)` is an extension function.
- `iban_format(iban)` is an extension function and returns `text`.
- `iban_ge(iban, iban)` is an extension function.
- `iban_gt(iban, iban)` is an extension function.
- `iban_hash(iban)` is an extension function and returns `int4`.
- `iban_in(cstring)` is an extension function and returns `iban`.
- `iban_le(iban, iban)` is an extension function.
- `iban_lt(iban, iban)` is an extension function.
- `iban_out(iban)` is an extension function and returns `cstring`.
- `iban_valid(text)` is an extension function.
- `iban` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
