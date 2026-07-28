## Usage

Sources:

- [Official documentation](https://www.postgresql.org/docs/12/datatype-json.html#JSON-TRANSFORM)
- [Official extension control file (jsonb_plpython2u.control)](https://github.com/postgres/postgres/blob/REL_12_STABLE/contrib/jsonb_plpython/jsonb_plpython2u.control)
- [Official extension SQL (jsonb_plpython2u--1.0.sql)](https://github.com/postgres/postgres/blob/REL_12_STABLE/contrib/jsonb_plpython/jsonb_plpython2u--1.0.sql)

`jsonb_plpython2u` — Historical transform between jsonb and the untrusted PL/Python 2 language. Use it when database code must run in or interoperate with this procedural language. The reviewed upstream project is archived or no longer maintained.

### Core Workflow

```sql
CREATE EXTENSION jsonb_plpython2u;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `jsonb_to_plpython2(val internal)` is an extension function and returns `internal`.
- `plpython2_to_jsonb(val internal)` is an extension function and returns `jsonb`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `plpython2u`.
- The control file marks the extension as relocatable.
- This historical transform depends on `plpython2u`; current PostgreSQL uses the Python 3 transform variant instead.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
