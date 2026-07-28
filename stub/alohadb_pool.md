## Usage

Sources:

- [Official upstream README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [Official extension control file (alohadb_pool.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_pool/alohadb_pool.control)
- [Official extension SQL (alohadb_pool--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_pool/alohadb_pool--1.0.sql)

`alohadb_pool` — AlohaDB Built-in Connection Pooling. Use it when administering or automating the database behavior described above. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION alohadb_pool;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pool_reset(pool_name text DEFAULT NULL)` is an extension function and returns `void`.
- `pool_settings()` is an extension function and returns `TABLE`.
- `pool_status()` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
