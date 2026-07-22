## Usage

Sources:

- [Official README](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/README.md)
- [Official installation guide](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/Install.md)
- [Official control file](https://github.com/csudata/gogudb/blob/490294508c8df14e29f312a1720e53690ef24de4/gogudb.control)

`gogudb` is a Postgres-XL-era distributed table-partitioning extension. It routes hash- or range-partitioned data to PostgreSQL servers through its foreign-data wrapper and keeps partition rules in extension-owned metadata. Use it only when maintaining a compatible legacy deployment; it is not a substitute for modern declarative partitioning.

### Core Workflow

The library must be preloaded. Add it to the server configuration and restart PostgreSQL before creating the extension:

```conf
shared_preload_libraries = 'gogudb'
```

```sql
CREATE EXTENSION gogudb;

CREATE SERVER shard1
FOREIGN DATA WRAPPER gogudb_fdw
OPTIONS (host '10.0.0.11', port '5432', dbname 'app');

CREATE USER MAPPING FOR CURRENT_USER
SERVER shard1 OPTIONS (user 'app', password 'secret');
```

Define the server/range map in `_gogu.server_map`, reload it with `_gogu.reload_range_server_set()`, then register a parent table's hash or range rule in `_gogu.table_partition_rule`. Follow the upstream examples exactly for the installed revision because these metadata layouts are part of the extension's implementation contract.

### Important Objects

- `gogudb_fdw`: foreign-data wrapper used for remote shards.
- `_gogu.server_map`: mapping from partition ranges to foreign servers.
- `_gogu.table_partition_rule`: table-level partition rules; upstream uses partition type `1` for hash and `2` for range.
- `_gogu.reload_range_server_set()`: reloads the in-memory range/server mapping after metadata changes.
- `_gogu` and `gogudb_partition_table`: fixed extension schemas; the extension is not relocatable.

### Caveats

The official material targets old PostgreSQL 9.6/10 and CentOS 7-era environments. Confirm binary compatibility rather than assuming support for a current PostgreSQL server. The extension installs hooks through `shared_preload_libraries`, so all databases in the instance share the operational risk of the loaded library.

Upstream explicitly does not support renaming partition parents, children, remote tables, or their schemas; direct modification or dropping of child/remote tables; partition split/merge; or read/write splitting. Treat the metadata tables as privileged control data, protect credentials in user mappings, and rehearse failure/recovery with disposable shards before adopting it.
