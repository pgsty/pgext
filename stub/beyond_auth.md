## Usage

Sources:

- [Official upstream README](https://github.com/beyondoss/auth/blob/50766c44c0d5a06037741b6c1e80b54f13bb54e0/beyond-auth-extension/README.md)
- [Official extension control file (beyond_auth.control)](https://github.com/beyondoss/auth/blob/50766c44c0d5a06037741b6c1e80b54f13bb54e0/beyond-auth-extension/beyond_auth.control)
- [Official implementation source](https://github.com/beyondoss/auth/blob/50766c44c0d5a06037741b6c1e80b54f13bb54e0/beyond-auth-extension/src/lib.rs)

`beyond_auth` — Evaluate transitive permissions inside PostgreSQL. Replaces N×depth round-trips with depth+1 queries via BFS over the auth.authz_relations graph. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION beyond_auth;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `authz_check_array` is an extension function.
- `authz_check_batch` is an extension function.
- `authz_check_multi` is an extension function.
- `authz_check_parallel_batch` is an extension function.
- `authz_check_path_batch` is an extension function.
- `authz_check_single` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
