## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/types/README.md)
- [Official extension control file (pgpm-types.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/types/pgpm-types.control)
- [Official extension SQL (pgpm-types--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/types/sql/pgpm-types--0.15.5.sql)

`pgpm-types` — Core PostgreSQL data types with SQL scripts. Use it when application data needs this type, domain, or its operators. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-types";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `attachment` is an extension-defined domain.
- `email` is an extension-defined domain.
- `hostname` is an extension-defined domain.
- `image` is an extension-defined domain.
- `origin` is an extension-defined domain.
- `upload` is an extension-defined domain.
- `url` is an extension-defined domain.

### Requirements and Caveats

- The reviewed control file declares default version `0.15.5`.
- Install the confirmed extension dependencies first: `plpgsql`, `citext`, `pgpm-verify`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
