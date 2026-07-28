## Usage

Sources:

- [Official upstream README](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_auth_multi/README)
- [Official extension control file (ext3.control)](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_auth_multi/ext3.control)
- [Official extension SQL (ext3--1.0.sql)](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_auth_multi/ext3--1.0.sql)

`ext3` — ClientAuthentication hook example that logs authentication start and end events. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION ext3;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ext3` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
