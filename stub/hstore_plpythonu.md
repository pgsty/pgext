## Usage

Sources:

- [Official upstream README](https://github.com/ssudarshaniitb/protectdb/blob/3e23a06f19785c72dd203b4d6bb6225e5cf6b9e3/safedb/contrib/README)
- [Official extension control file (hstore_plpythonu.control)](https://github.com/ssudarshaniitb/protectdb/blob/3e23a06f19785c72dd203b4d6bb6225e5cf6b9e3/safedb/contrib/hstore_plpython/hstore_plpythonu.control)
- [Official extension SQL (hstore_plpythonu--1.0.sql)](https://github.com/ssudarshaniitb/protectdb/blob/3e23a06f19785c72dd203b4d6bb6225e5cf6b9e3/safedb/contrib/hstore_plpython/hstore_plpythonu--1.0.sql)

`hstore_plpythonu` — transform between hstore and plpythonu. Use it when database code must run in or interoperate with this procedural language. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION hstore_plpythonu;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hstore_to_plpython(val internal)` is an extension function and returns `internal`.
- `plpython_to_hstore(val internal)` is an extension function and returns `hstore`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `hstore`, `plpythonu`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
