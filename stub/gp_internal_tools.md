## Usage

Sources:

- [Official upstream README](https://github.com/gridgentoo/gpdb/blob/f3dc101a7b4fa3d392f79cc5146b20c83894eb19/contrib/gp_internal_tools/README)
- [Official extension control file (gp_internal_tools.control)](https://github.com/gridgentoo/gpdb/blob/f3dc101a7b4fa3d392f79cc5146b20c83894eb19/contrib/gp_internal_tools/gp_internal_tools.control)
- [Official extension SQL (gp_internal_tools--1.0.0.sql)](https://github.com/gridgentoo/gpdb/blob/f3dc101a7b4fa3d392f79cc5146b20c83894eb19/contrib/gp_internal_tools/gp_internal_tools--1.0.0.sql)

`gp_internal_tools` — Different internal tools for Greenplum. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gp_internal_tools;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `session_state_memory_entries_f_on_master()` is an extension function and returns `SETOF`.
- `session_state_memory_entries_f_on_segments()` is an extension function and returns `SETOF`.
- `session_level_memory_consumption` is an extension-defined view.
- `session_state` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
