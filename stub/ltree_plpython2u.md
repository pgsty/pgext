## Usage

Sources:

- [Official upstream README](https://github.com/vicmisael/postgres-xl-fix/blob/0bca966cd19ca7be25676db32f0f50c014349f20/contrib/README)
- [Official extension control file (ltree_plpython2u.control)](https://github.com/vicmisael/postgres-xl-fix/blob/0bca966cd19ca7be25676db32f0f50c014349f20/contrib/ltree_plpython/ltree_plpython2u.control)
- [Official extension SQL (ltree_plpython2u--1.0.sql)](https://github.com/vicmisael/postgres-xl-fix/blob/0bca966cd19ca7be25676db32f0f50c014349f20/contrib/ltree_plpython/ltree_plpython2u--1.0.sql)

`ltree_plpython2u` — A fixed postgres xl so it can compile on newer versions of GCC and glibc. Use it when database code must run in or interoperate with this procedural language. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION ltree_plpython2u;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ltree_to_plpython2(val internal)` is an extension function and returns `internal`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `ltree`, `plpython2u`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
