## Usage

Sources:

- [Official upstream README](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/README.md)
- [Official extension control file (pg_mentat.control)](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/pg_mentat.control)
- [Official extension SQL (pg_mentat--1.0.0.sql)](https://github.com/gburd/pg_mentat/blob/134015ebd3121e1a74eeff2de6a800143b33cb4e/pg_mentat/sql/pg_mentat--1.0.0.sql)

`pg_mentat` — PostgreSQL extension providing a Datomic-compatible Datalog query engine with a native EDN data type. Use it when porting or emulating the corresponding database API. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_mentat;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mentat.allocate_entid(partition_name TEXT)` is an extension function and returns `BIGINT`.
- `mentat.fulltext_update_trigger()` is an extension function and returns `trigger`.
- `mentat.resolve_ident(keyword TEXT)` is an extension function and returns `BIGINT`.
- `mentat.cardinality_type` is an extension-defined type.
- `mentat.EdnValue` is an extension-defined type.
- `mentat.unique_type` is an extension-defined type.
- `mentat.value_type` is an extension-defined type.
- `mentat.datoms` is a table installed or managed by the extension.
- `mentat.datoms_bool` is a table installed or managed by the extension.
- `mentat.datoms_bytes` is a table installed or managed by the extension.
- `mentat.datoms_default` is a table installed or managed by the extension.
- `mentat.datoms_double` is a table installed or managed by the extension.
- `mentat.datoms_instant` is a table installed or managed by the extension.
- `mentat.datoms_keyword` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.5.7`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
