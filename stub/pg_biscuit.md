## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_biscuit/pg_biscuit-1.0.0/README.md)
- [Official extension control file (pg_biscuit.control)](https://api.pgxn.org/src/pg_biscuit/pg_biscuit-1.0.0/pg_biscuit.control)
- [Official extension SQL (pg_biscuit--1.0.sql)](https://api.pgxn.org/src/pg_biscuit/pg_biscuit-1.0.0/sql/pg_biscuit--1.0.sql)
- [Current upstream continuation repository](https://github.com/CrystallineCore/Biscuit)

`pg_biscuit` — A PostgreSQL Index Access Method (IAM) for high-performance pattern matching on text columns. Biscuit indexes are specifically designed to accelerate LIKE queries with arbitrary wildcards. Use it for the corresponding text-search, parsing, or linguistic workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_biscuit;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `biscuit_handler(internal)` is an extension function and returns `index_am_handler`.
- `biscuit_index_stats(oid)` is an extension function and returns `text`.
- `biscuit_like_support(internal)` is an extension function.
- `biscuit_multicolumn_enabled()` is an extension function and returns `boolean`.
- `biscuit_indexes` is an extension-defined view.
- `biscuit_indexes_detailed` is an extension-defined view.
- `biscuit_version` is a table installed or managed by the extension.
- `products` is a table installed or managed by the extension.
- `biscuit` is an extension-defined access method.
- `biscuit_bpchar_ops` is an extension-defined operator class.
- `biscuit_text_ops` is an extension-defined operator class.
- `biscuit_varchar_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as non-relocatable.
- The PGXN sources above document the `pg_biscuit` identity; the current upstream repository publishes the renamed `biscuit` extension, so review its migration boundary before substituting it.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
