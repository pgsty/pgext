## Usage

Sources:

- [Official upstream README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [Official extension control file (alohadb_branch.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_branch/alohadb_branch.control)
- [Official extension SQL (alohadb_branch--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_branch/alohadb_branch--1.0.sql)

`alohadb_branch` — Lightweight database branching for testing migrations and experiments. Use it when an application needs this specific database capability. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION alohadb_branch;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `alohadb_create_branch(name text, from_lsn pg_lsn DEFAULT NULL, OUT branch_name text, OUT port int, OUT data_dir text)` is an extension function and returns `record`.
- `alohadb_drop_branch(name text)` is an extension function and returns `void`.
- `alohadb_list_branches()` is an extension function and returns `TABLE`.
- `alohadb_branches` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
