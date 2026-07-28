## Usage

Sources:

- [Official upstream README](https://github.com/chadcatlett/pg-jq/blob/99b40b8d9d4924bab3465f90c655457cb2d9c50c/README.md)
- [Official extension control file (pg_jq.control)](https://github.com/chadcatlett/pg-jq/blob/99b40b8d9d4924bab3465f90c655457cb2d9c50c/pg_jq.control)
- [Official implementation source](https://github.com/chadcatlett/pg-jq/blob/99b40b8d9d4924bab3465f90c655457cb2d9c50c/src/lib.rs)

`pg_jq` — This is a toy PostgreSQL extension that exposes basic libjq functionality to PostgreSQL. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_jq;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `json_jq` is an extension function.
- `jsonb_jq` is an extension function.
- `what_is_something_carlson_likes()` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
