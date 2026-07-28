## Usage

Sources:

- [Official upstream README](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/docs/README.md)
- [Official extension control file (veil2.control)](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/veil2.control)
- [Official extension SQL (veil2--0.9.1.sql)](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/sql/veil2--0.9.1.sql)

`veil2` — Provides the basis for a Virtual Private Database implementation. Use it when implementing the corresponding security, audit, or access-control workflow. The reviewed upstream material marks this capability deprecated.

### Core Workflow

```sql
CREATE EXTENSION veil2;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `veil2.add_session_privileges(scope_type_id integer, scope_id integer, roles bitmap, privs bitmap)` is an extension function and returns `void`.
- `veil2.all_accessor_roles(accessor_id in out integer, session_context_type_id in integer, session_context_id in integer, role_id out integer, context_type_id out integer, context_id out integer)` is an extension function and returns `setof`.
- `veil2.always_true(integer)` is an extension function and returns `boolean`.
- `veil2.authenticate(accessor_id integer, authent_type text, token text)` is an extension function and returns `boolean`.
- `veil2.authenticate_bcrypt(accessor_id integer, token text)` is an extension function and returns `boolean`.
- `veil2.authenticate_false(accessor_id integer, token text)` is an extension function and returns `boolean`.
- `veil2.authenticate_plaintext(accessor_id integer, token text)` is an extension function and returns `boolean`.
- `veil2.base_accessor_roleprivs(accessor_id in out integer, session_context_type_id in integer, session_context_id in integer, mapping_context_type_id in out integer, mapping_context_id in out integer, assignment_context_type_id out integer, assignment_context_id out integer, role_id out int…)` is an extension function and returns `setof`.
- `veil2.bcrypt(passwd text)` is an extension function and returns `text`.
- `veil2.become_accessor(accessor_id in integer, login_context_type_id in integer, login_context_id in integer, session_context_type_id in integer, session_context_id in integer, session_id out bigint, session_token out text, success out boolean, errmsg out text)` is an extension function and returns `record`.
- `veil2.become_accessor(accessor_id in integer, login_context_type_id in integer, login_context_id in integer, session_context_type_id in integer, session_context_id in integer, session_id out integer, session_token out text, success out boolean, errmsg out text)` is an extension function and returns `record`.
- `veil2.become_user(username in text, login_context_type_id in integer, login_context_id in integer, session_context_type_id in integer default null, session_context_id in integer default null, session_id out bigint, session_token out text, success out boolean, errmsg out text)` is an extension function and returns `record`.
- `veil2.become_user(username in text, login_context_type_id in integer, login_context_id in integer, session_context_type_id in integer default null, session_context_id in integer default null, session_id out integer, session_token out text, success out boolean, errmsg out text)` is an extension function and returns `record`.
- `veil2.check_accessor_context(label text, accessor_id integer, context_type_id integer, context_id integer)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.9.3`.
- Install the confirmed extension dependencies first: `pgbitmap`, `pgcrypto`.
- The control file requires a superuser for installation.
- Upstream material contains an explicit deprecation boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
