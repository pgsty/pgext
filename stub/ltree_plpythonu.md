## Usage

Sources:

- [Official upstream README](https://github.com/masahikosawada/gmapbench/blob/a2425a83a59454d931793febaa52fb04d8912492/source/pg-garbagemap-background/contrib/README)
- [Official extension control file (ltree_plpythonu.control)](https://github.com/masahikosawada/gmapbench/blob/a2425a83a59454d931793febaa52fb04d8912492/source/pg-garbagemap-background/contrib/ltree_plpython/ltree_plpythonu.control)
- [Official extension SQL (ltree_plpythonu--1.0.sql)](https://github.com/masahikosawada/gmapbench/blob/a2425a83a59454d931793febaa52fb04d8912492/source/pg-garbagemap-background/contrib/ltree_plpython/ltree_plpythonu--1.0.sql)

`ltree_plpythonu` — transform between ltree and plpythonu. Use it when database code must run in or interoperate with this procedural language. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION ltree_plpythonu;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ltree_to_plpython(val internal)` is an extension function and returns `internal`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `ltree`, `plpythonu`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
