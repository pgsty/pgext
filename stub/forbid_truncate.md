## Usage

Sources:

- [forbid_truncate README at the reviewed commit](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/README.md)
- [forbid_truncate hook implementation at the reviewed commit](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/forbid_truncate.c)
- [forbid_truncate control file at the reviewed commit](https://github.com/NaokiNakamichi/postgres_extension/blob/47029106ccbc1fe00d0ec9a612ddfedd7abf6847/forbid_truncate.control)

`forbid_truncate` installs a ProcessUtility hook that rejects every TRUNCATE statement. The library must be loaded while PostgreSQL starts. After installing the binary, add it to the server configuration and restart PostgreSQL:

```ini
shared_preload_libraries = 'forbid_truncate'
```

Record the installed extension in each database where it is managed, then verify the guard in a disposable table. The final statement is expected to fail with “forbid truncate”.

```sql
CREATE EXTENSION forbid_truncate;

CREATE TABLE truncate_guard_demo (id integer);
TRUNCATE TABLE truncate_guard_demo;
```

### Caveats

- Changing `shared_preload_libraries` requires a server restart. Creating the extension without preloading the library does not activate the hook in existing backends.
- The implementation blocks all users, including superusers, and provides no per-role, per-table, or maintenance bypass.
- The hook calls PostgreSQL's standard utility handler directly instead of chaining to the previously installed hook. Test it with every other ProcessUtility-hook extension used by the cluster.
- This guard covers TRUNCATE only; DELETE, DROP TABLE, partition maintenance, and other data-removal paths remain unaffected.
