## Usage

Sources:

- [Official upstream README](https://github.com/guepard-corp/gfs/blob/461d07820e253828d085a3b387f6912753e25d7c/crates/extensions/gfs/README.md)
- [Official extension control file (gfs.control)](https://github.com/guepard-corp/gfs/blob/461d07820e253828d085a3b387f6912753e25d7c/crates/extensions/gfs/gfs.control)
- [Official extension SQL (gfs--0.0.1.sql)](https://github.com/guepard-corp/gfs/blob/461d07820e253828d085a3b387f6912753e25d7c/crates/extensions/gfs/c-ref/gfs--0.0.1.sql)

`gfs` — Clones a remote PostgreSQL **copy-on-read**: an empty local database that fetches data from the source only when a query touches it, so a multi-TB source can be "cloned" instantly and the clone stays **partial** (it never has to pull what the app doesn't read). Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gfs;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gfs.register_clone(local regclass, source_ref text, key_col text DEFAULT 'id')` is an extension function and returns `void`.
- `gfs.unregister_clone(local regclass)` is an extension function and returns `void`.
- `gfs.warm(local regclass)` is an extension function and returns `bigint`.
- `gfs_handler(internal)` is an extension function and returns `table_am_handler`.
- `gfs.clones` is an extension-defined view.
- `gfs.clone_source` is a table installed or managed by the extension.
- `gfs.clone_stats` is a table installed or managed by the extension.
- `gfs` is a schema created by the extension.
- `gfs` is an extension-defined access method.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
