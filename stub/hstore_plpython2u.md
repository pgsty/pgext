## Usage

Sources:

- [Official upstream README](https://github.com/skyrise-l/queryartisan/blob/29613aa7bbfd5108a64ac4b5115bc3f94310985b/optimization/contrib/README)
- [Official extension control file (hstore_plpython2u.control)](https://github.com/skyrise-l/queryartisan/blob/29613aa7bbfd5108a64ac4b5115bc3f94310985b/optimization/contrib/hstore_plpython/hstore_plpython2u.control)
- [Official extension SQL (hstore_plpython2u--1.0.sql)](https://github.com/skyrise-l/queryartisan/blob/29613aa7bbfd5108a64ac4b5115bc3f94310985b/optimization/contrib/hstore_plpython/hstore_plpython2u--1.0.sql)

`hstore_plpython2u` — transform between hstore and plpython2u. Use it when database code must run in or interoperate with this procedural language. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION hstore_plpython2u;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hstore_to_plpython2(val internal)` is an extension function and returns `internal`.
- `plpython2_to_hstore(val internal)` is an extension function and returns `hstore`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `hstore`, `plpython2u`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
