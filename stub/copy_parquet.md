## Usage

Sources:

- [Official upstream README](https://github.com/kou/pg-copy-parquet/blob/7da367ea81d8964f5045fe0b1514a798d4ecbbc7/README.md)
- [Official extension control file (copy_parquet.control)](https://github.com/kou/pg-copy-parquet/blob/7da367ea81d8964f5045fe0b1514a798d4ecbbc7/copy_parquet.control)
- [Official extension SQL (copy_parquet--0.0.1.sql)](https://github.com/kou/pg-copy-parquet/blob/7da367ea81d8964f5045fe0b1514a798d4ecbbc7/copy_parquet--0.0.1.sql)

`copy_parquet` — PoC to support PostgreSQL COPY with Apache Parquet. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Upstream describes it as a proof of concept.

### Core Workflow

```sql
CREATE EXTENSION copy_parquet;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `copy_to_parquet(pg_catalog.regclass)` is an extension function and returns `bytea`.
- `format_parquet(bytea)` is an extension function and returns `text`.
- `scan_to_parquet(pg_catalog.regclass)` is an extension function and returns `bytea`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
