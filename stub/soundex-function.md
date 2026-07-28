## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/soundex-function/soundex-function-1.0.0/README.md)
- [Official extension control file (soundex-function.control)](https://api.pgxn.org/src/soundex-function/soundex-function-1.0.0/soundex-function.control)
- [Official extension SQL (soundex-function--1.0.0.sql)](https://api.pgxn.org/src/soundex-function/soundex-function-1.0.0/soundex-function--1.0.0.sql)

`soundex-function` — This repository contains a template for creating a Postgres function. The goal of this template is to provide a starting point for the creation of functions for Postgres. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "soundex-function";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `soundex(input TEXT)` is an extension function and returns `TEXT`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
