## Usage

Sources:

- [Official upstream README](https://github.com/babdulhakim2/pgpm-test/blob/8ce14c436c1ace0eb846844ed78b22b777931036/extensions/@pgpm/verify/README.md)
- [Official extension control file (pgpm-verify.control)](https://github.com/babdulhakim2/pgpm-test/blob/8ce14c436c1ace0eb846844ed78b22b777931036/extensions/@pgpm/verify/pgpm-verify.control)
- [Official extension SQL (pgpm-verify--0.15.3.sql)](https://github.com/babdulhakim2/pgpm-test/blob/8ce14c436c1ace0eb846844ed78b22b777931036/extensions/@pgpm/verify/sql/pgpm-verify--0.15.3.sql)

`pgpm-verify` — Verification utilities for PostgreSQL modules. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-verify";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_entity_from_str(qualified_name text)` is an extension function and returns `text`.
- `get_schema_from_str(qualified_name text)` is an extension function and returns `text`.
- `list_indexes(_table text, _index text)` is an extension function and returns `TABLE`.
- `list_memberships(_user text)` is an extension function and returns `TABLE`.
- `verify_constraint(_table text, _name text)` is an extension function and returns `boolean`.
- `verify_domain(_type text)` is an extension function and returns `boolean`.
- `verify_extension(_extname text)` is an extension function and returns `boolean`.
- `verify_function(_name text, _user text DEFAULT NULL)` is an extension function and returns `boolean`.
- `verify_index(_table text, _index text)` is an extension function and returns `boolean`.
- `verify_membership(_user text, _role text)` is an extension function and returns `boolean`.
- `verify_policy(_policy text, _table text)` is an extension function and returns `boolean`.
- `verify_role(_user text)` is an extension function and returns `boolean`.
- `verify_schema(_schema text)` is an extension function and returns `boolean`.
- `verify_security(_table text)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `0.15.3`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
