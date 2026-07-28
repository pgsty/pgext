## Usage

Sources:

- [Official upstream README](https://github.com/hyz1840/pg_jiebaparser/blob/a244c1332f02b1383a822c0c952d55b6949da329/README.md)
- [Official extension control file (jiebaparser.control)](https://github.com/hyz1840/pg_jiebaparser/blob/a244c1332f02b1383a822c0c952d55b6949da329/jiebaparser.control)
- [Official extension SQL (jiebaparser.sql)](https://github.com/hyz1840/pg_jiebaparser/blob/a244c1332f02b1383a822c0c952d55b6949da329/jiebaparser.sql)

`jiebaparser` — Postgresql full-text search extension for chinese (jiaba engine) using shared mem. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION jiebaparser;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `jbprs_end(internal)` is an extension function and returns `void`.
- `jbprs_getlexeme(internal, internal, internal)` is an extension function and returns `internal`.
- `jbprs_lextype(internal)` is an extension function and returns `internal`.
- `jbprs_start(internal, int4)` is an extension function and returns `internal`.
- `jbprs_start_q(internal, int4)` is an extension function and returns `internal`.
- `jiebaparser_reset()` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
