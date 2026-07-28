## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/soundex-type/soundex-type-1.0.0/README.md)
- [Official extension control file (soundex-type.control)](https://api.pgxn.org/src/soundex-type/soundex-type-1.0.0/soundex-type.control)
- [Official extension SQL (soundex-type--1.0.0.sql)](https://api.pgxn.org/src/soundex-type/soundex-type-1.0.0/soundex-type--1.0.0.sql)

`soundex-type` — This repository contains a template for creating a Postgres type. The goal of this template is to provide a starting point for the creation of types for Postgres. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "soundex-type";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `soundex_in(cstring)` is an extension function and returns `soundex`.
- `soundex_out(soundex)` is an extension function and returns `cstring`.
- `soundex` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
