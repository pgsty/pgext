## Usage

Sources:

- [Official upstream README](https://github.com/marcelomendoncasoares/pg_hlc/blob/dc7691ef484e6016f5de38eb9ee0ad35dbbdd39b/README.md)
- [Official extension control file (pg_hlc.control)](https://github.com/marcelomendoncasoares/pg_hlc/blob/dc7691ef484e6016f5de38eb9ee0ad35dbbdd39b/pg_hlc.control)
- [Official extension SQL (pg_hlc--0.1.0.sql)](https://github.com/marcelomendoncasoares/pg_hlc/blob/dc7691ef484e6016f5de38eb9ee0ad35dbbdd39b/pg_hlc--0.1.0.sql)

`pg_hlc` — A PostgreSQL extension (pg_hlc) that provides Hybrid Logical Clock (HLC) functionality with **100% compatibility** with the Dart CRDT library. The extension is built using the pgrx framework and implements the exact same HLC algorithm and API as the Dart reference implementation. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_hlc;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hlc_compare(hlctimestamp, hlctimestamp)` is an extension function and returns `integer`.
- `hlc_eq(hlctimestamp, hlctimestamp)` is an extension function and returns `boolean`.
- `hlc_from_date(text, text)` is an extension function and returns `hlctimestamp`.
- `hlc_gt(hlctimestamp, hlctimestamp)` is an extension function and returns `boolean`.
- `hlc_gte(hlctimestamp, hlctimestamp)` is an extension function and returns `boolean`.
- `hlc_increment(text)` is an extension function and returns `hlctimestamp`.
- `hlc_lt(hlctimestamp, hlctimestamp)` is an extension function and returns `boolean`.
- `hlc_lte(hlctimestamp, hlctimestamp)` is an extension function and returns `boolean`.
- `hlc_merge(text, hlctimestamp)` is an extension function and returns `hlctimestamp`.
- `hlc_ne(hlctimestamp, hlctimestamp)` is an extension function and returns `boolean`.
- `hlc_now(text)` is an extension function and returns `hlctimestamp`.
- `hlc_parse(text)` is an extension function and returns `hlctimestamp`.
- `hlc_reset(text)` is an extension function and returns `boolean`.
- `hlc_to_string(hlctimestamp)` is an extension function and returns `text`.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
