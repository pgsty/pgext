## Usage

Sources:

- [Official upstream README](https://github.com/kalamdb/koldstore/blob/2a50565936bc97377a439dcbb4cc2ae5c07db1a5/README.md)
- [Official extension control file (koldstore.control)](https://github.com/kalamdb/koldstore/blob/2a50565936bc97377a439dcbb4cc2ae5c07db1a5/crates/pg_koldstore/koldstore.control)
- [Official extension SQL (koldstore--0.1.0.sql)](https://github.com/kalamdb/koldstore/blob/2a50565936bc97377a439dcbb4cc2ae5c07db1a5/crates/pg_koldstore/sql/koldstore--0.1.0.sql)

`koldstore` — PostgreSQL tiered-storage that moves historical rows to Parquet while keeping the original table fully queryable and supporting updates and deletes. Use it for the corresponding analytical or storage workflow. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION koldstore;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `koldstore.internal_apply_flush_row_counts(p_table_oid oid, p_mirror_pruned bigint, p_hot_pruned bigint, p_cold_rows_added bigint)` is an extension function and returns `void`.
- `koldstore.internal_bump_row_counts(p_table_oid oid, p_hot_delta bigint, p_mirror_delta bigint)` is an extension function and returns `void`.
- `koldstore.internal_ensure_manifest_row(p_table_oid oid)` is an extension function and returns `void`.
- `koldstore.internal_refresh_row_counts(p_table_oid oid, p_hot_rows bigint, p_mirror_rows bigint)` is an extension function and returns `void`.
- `koldstore.change_event` is an extension-defined type.
- `koldstore.dml_result` is an extension-defined type.
- `koldstore.managed_table_info` is an extension-defined type.
- `koldstore.async_mirror_state` is a table installed or managed by the extension.
- `koldstore.cold_segment_stats` is a table installed or managed by the extension.
- `koldstore.cold_segments` is a table installed or managed by the extension.
- `koldstore.jobs` is a table installed or managed by the extension.
- `koldstore.manifest` is a table installed or managed by the extension.
- `koldstore.schemas` is a table installed or managed by the extension.
- `koldstore.storage` is a table installed or managed by the extension.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Upstream explicitly says the project is not production-ready.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
