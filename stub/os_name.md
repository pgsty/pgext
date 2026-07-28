## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/os_name/os_name-0.0.3/README.md)
- [Official extension control file (os_name.control)](https://api.pgxn.org/src/os_name/os_name-0.0.3/os_name.control)
- [Official extension SQL (os_name--0.0.2.sql)](https://api.pgxn.org/src/os_name/os_name-0.0.3/os_name--0.0.2.sql)

`os_name` — Enumerable mobile os type, stored in a single-byte, fixed-length type. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION os_name;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hash_os_name(os_name)` is an extension function and returns `integer`.
- `os_name_cmp(os_name, os_name)` is an extension function and returns `integer`.
- `os_name_eq(os_name, os_name)` is an extension function and returns `boolean`.
- `os_name_ge(os_name, os_name)` is an extension function and returns `boolean`.
- `os_name_gt(os_name, os_name)` is an extension function and returns `boolean`.
- `os_name_in(cstring)` is an extension function and returns `os_name`.
- `os_name_le(os_name, os_name)` is an extension function and returns `boolean`.
- `os_name_lt(os_name, os_name)` is an extension function and returns `boolean`.
- `os_name_ne(os_name, os_name)` is an extension function and returns `boolean`.
- `os_name_out(os_name)` is an extension function and returns `cstring`.
- `os_name_recv(internal)` is an extension function and returns `os_name`.
- `os_name_send(os_name)` is an extension function and returns `bytea`.
- `os_name` is an extension-defined type.
- `btree_os_name_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.3`.
- The control file marks the extension as relocatable.
- The former GitHub repository URL returned 404 during the 2026-07-28 review; treat the pinned PGXN distribution above as the available source boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
