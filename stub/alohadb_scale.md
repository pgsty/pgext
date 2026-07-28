## Usage

Sources:

- [Official upstream README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [Official extension control file (alohadb_scale.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_scale/alohadb_scale.control)
- [Official extension SQL (alohadb_scale--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_scale/alohadb_scale--1.0.sql)

`alohadb_scale` — AlohaDB Scale - scale-to-zero management. Use it when administering or automating the database behavior described above. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION alohadb_scale;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `scale_activity()` is an extension function and returns `TABLE`.
- `scale_configure(suspend_after interval DEFAULT NULL, min_connections int DEFAULT NULL)` is an extension function and returns `void`.
- `scale_status()` is an extension function and returns `TABLE`.
- `scale_suspend()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
