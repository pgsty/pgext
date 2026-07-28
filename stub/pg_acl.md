## Usage

Sources:

- [Official extension control file (pg_acl.control)](https://api.pgxn.org/src/pg_acl/pg_acl-0.1.3/pg_acl.control)
- [Official extension SQL (pg_acl.in.sql)](https://api.pgxn.org/src/pg_acl/pg_acl-0.1.3/sql/pg_acl.in.sql)

`pg_acl` — Utilities for handling aclitems. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_acl;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `acl(input aclitem)` is an extension function and returns `acl`.
- `acl(input aclitem[])` is an extension function and returns `acl[]`.
- `acl` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.3`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
