## Usage

Sources:

- [Official upstream README](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/readme.md)
- [Official extension control file (launchql-extension-utils.control)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/utils/launchql-extension-utils.control)
- [Official extension SQL (launchql-extension-utils--0.0.1.sql)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/utils/sql/launchql-extension-utils--0.0.1.sql)

`launchql-extension-utils` — PostgreSQL utilities. Use it for the corresponding SQL or database utility workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "launchql-extension-utils";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_entity_from_str(qualified_name text)` is an extension function and returns `text`.
- `get_schema_from_str(qualified_name text)` is an extension function and returns `text`.
- `list_indexes(_table text, _index text)` is an extension function and returns `TABLE`.
- `list_memberships(_user text)` is an extension function and returns `TABLE`.
- `tg_update_timestamps()` is an extension function and returns `trigger`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
