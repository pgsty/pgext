## Usage

Sources:

- [Official upstream README](https://github.com/kolharsam/pg_bm25/blob/2e6a875874e0adb2f1e8d735059c66e208e2ebc8/docs/src/README.md)
- [Official extension control file (parodydb.control)](https://github.com/kolharsam/pg_bm25/blob/2e6a875874e0adb2f1e8d735059c66e208e2ebc8/parodydb.control)
- [Official implementation source](https://github.com/kolharsam/pg_bm25/blob/2e6a875874e0adb2f1e8d735059c66e208e2ebc8/src/lib.rs)

`parodydb` — A toy full-text search extension for Postgres, written in Rust using pgrx. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION parodydb;

SELECT parodydb_search('The quick brown fox', 'quick');  -- true
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hello_parodydb()` is an extension function.
- `parodydb_index` is an extension function.
- `parodydb_index_clear()` is an extension function.
- `parodydb_index_info()` is an extension function.
- `parodydb_search` is an extension function.
- `parodydb_tokenize` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
