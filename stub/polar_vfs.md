## Usage

Sources:

- [Official upstream README](https://github.com/polardb/polardb-for-postgresql/blob/b5f78e9e5af18ea3ffb54b8d638208ed7648edca/README.md)
- [Official extension control file (polar_vfs.control)](https://github.com/polardb/polardb-for-postgresql/blob/b5f78e9e5af18ea3ffb54b8d638208ed7648edca/src/polar_vfs/polar_vfs.control)
- [Official extension SQL (polar_vfs--1.0.sql)](https://github.com/polardb/polardb-for-postgresql/blob/b5f78e9e5af18ea3ffb54b8d638208ed7648edca/src/polar_vfs/polar_vfs--1.0.sql)

`polar_vfs` — A cloud-native database based on PostgreSQL developed by Alibaba Cloud. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION polar_vfs;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `polar_libpfs_version()` is an extension function and returns `text`.
- `polar_vfs_disk_expansion(text)` is an extension function.
- `polar_vfs_mem_status()` is an extension function and returns `setof`.
- `mm_type` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
