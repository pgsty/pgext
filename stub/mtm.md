## Usage

Sources:

- [Official upstream README](https://github.com/skilyazhnev/mtm/blob/470bc005b180469c1246e0ead8cfeed703a8a6e3/README.md)
- [Official extension control file (mtm.control)](https://github.com/skilyazhnev/mtm/blob/470bc005b180469c1246e0ead8cfeed703a8a6e3/mtm/mtm.control)
- [Official extension SQL (mtm--0.1.sql)](https://github.com/skilyazhnev/mtm/blob/470bc005b180469c1246e0ead8cfeed703a8a6e3/mtm/mtm--0.1.sql)

`mtm` — aggregate to find min and max. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION mtm;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `final_mtm(state state_mtm)` is an extension function and returns `text`.
- `final_mtm(state state_mtm_dp)` is an extension function and returns `text`.
- `transition_mtm(state state_mtm, val numeric)` is an extension function and returns `state_mtm`.
- `transition_mtm(state state_mtm_dp, val double precision)` is an extension function and returns `state_mtm_dp`.
- `max_to_min` is an aggregate exposed by the extension.
- `state_mtm` is an extension-defined type.
- `state_mtm_dp` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
