## Usage

Sources:

- [Official upstream README](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/readme.md)
- [Official extension control file (launchql-extension-verify.control)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/verify/launchql-extension-verify.control)
- [Official extension SQL (launchql-extension-verify--0.0.1.sql)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/verify/sql/launchql-extension-verify--0.0.1.sql)

`launchql-extension-verify` — PostgreSQL verification utilities. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "launchql-extension-verify";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

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
- `verify_table(_table text)` is an extension function and returns `boolean`.
- `verify_table_grant(_table text, _privilege text, _role text)` is an extension function and returns `boolean`.
- `verify_trigger(_trigger text)` is an extension function and returns `boolean`.
- `verify_type(_type text)` is an extension function and returns `boolean`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `plpgsql`, `uuid-ossp`, `launchql-extension-utils`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
