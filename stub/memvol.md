## Usage

Sources:

- [Official extension control file (memvol.control)](https://github.com/grundez/pg_extensions/blob/59ffdb504865ff0838a8fce127cd66a4c1952121/MemoryVolume/memvol.control)
- [Official extension SQL (memvol--1.0.sql)](https://github.com/grundez/pg_extensions/blob/59ffdb504865ff0838a8fce127cd66a4c1952121/MemoryVolume/memvol--1.0.sql)

`memvol` — Memory units. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION memvol;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `display_vlm_ref()` is an extension function and returns `SETOF`.
- `init_vlm_ref()` is an extension function and returns `void`.
- `vlm2vlm(numeric, text, text)` is an extension function and returns `numeric`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
