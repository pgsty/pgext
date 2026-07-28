## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_set_acl/pg_set_acl-0.0.2/README.md)
- [Official extension control file (pg_set_acl.control)](https://api.pgxn.org/src/pg_set_acl/pg_set_acl-0.0.2/pg_set_acl.control)
- [Official extension SQL (pg_set_acl--0.0.1.sql)](https://api.pgxn.org/src/pg_set_acl/pg_set_acl-0.0.2/pg_set_acl--0.0.1.sql)

`pg_set_acl` — PostgreSQL extension that implements a SET command access control list. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_set_acl;

select set_acl.grant(setting, user);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `set_acl.grant(cstring, cstring)` is an extension function.
- `set_acl.read_acl(cstring, cstring)` is an extension function.
- `set_acl.revoke(cstring, cstring)` is an extension function.
- `set_acl.privs` is a table installed or managed by the extension.
- `set_acl` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
