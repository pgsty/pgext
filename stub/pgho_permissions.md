## Usage

Sources:

- [Official upstream README](https://github.com/asghonim/pgho_permissions/blob/077ab341aade6e33fc7c60eee0297cdb37b67ed2/README.md)
- [Official extension control file (pgho_permissions.control)](https://github.com/asghonim/pgho_permissions/blob/077ab341aade6e33fc7c60eee0297cdb37b67ed2/pgho_permissions.control)
- [Official extension SQL (pgho_permissions--0.0.18.sql)](https://github.com/asghonim/pgho_permissions/blob/077ab341aade6e33fc7c60eee0297cdb37b67ed2/pgho_permissions--0.0.18.sql)

`pgho_permissions` — A PostgreSQL extension that brings flexible, hierarchical access control to your database. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgho_permissions;

-- Install the extension
SELECT dbdev.install('asghonim@pgho_permissions');

-- (Re)create the schema and extension
DROP EXTENSION  IF EXISTS "asghonim@pgho_permissions";
DROP SCHEMA     IF EXISTS pgho_permissions;
CREATE SCHEMA   IF NOT EXISTS pgho_permissions;
CREATE EXTENSION IF NOT EXISTS "asghonim@pgho_permissions"
  SCHEMA pgho_permissions
  VERSION '0.0.18';
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.18`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
