## Usage

Sources:

- [Official upstream README](https://github.com/ai-db-uom/dbtune/blob/997916b3ed9aff19b15bd9a6a8379ddbfe52cdb0/README.md)
- [Official extension control file (dbtune_mab.control)](https://github.com/ai-db-uom/dbtune/blob/997916b3ed9aff19b15bd9a6a8379ddbfe52cdb0/dbtune_pg_mab_extension/dbtune_mab.control)
- [Official extension SQL (dbtune_mab--0.0.1.sql)](https://github.com/ai-db-uom/dbtune/blob/997916b3ed9aff19b15bd9a6a8379ddbfe52cdb0/dbtune_pg_mab_extension/dbtune_mab--0.0.1.sql)

`dbtune_mab` — DBTune MAB advisor for PostgreSQL. Use it when an application needs this specific database capability. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION dbtune_mab;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `dbtune_mab_tune(tablename TEXT, columns TEXT[])` is an extension function and returns `TEXT`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
