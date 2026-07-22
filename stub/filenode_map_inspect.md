## Usage

Sources:

- [Official extension control file](https://github.com/ibarwick/filenode_map_inspect/blob/8cffe378ce38a6c3d7bff828636fff33e29e46f4/filenode_map_inspect.control)
- [Official upstream documentation](https://github.com/ibarwick/filenode_map_inspect/blob/8cffe378ce38a6c3d7bff828636fff33e29e46f4/README.md)

`filenode_map_inspect` 0.1 reads PostgreSQL `pg_filenode.map` files for cluster-wide diagnostics. It can identify unreadable or incomplete mapping files and list mapped catalog objects, but it neither repairs corruption nor replaces normal recovery procedures.

### Core Workflow

```sql
CREATE EXTENSION filenode_map_inspect;

SELECT * FROM filenode_map_check();
SELECT * FROM filenode_map_list();
SELECT * FROM filenode_map_list_global();
```

- `filenode_map_check()` locates database and global map files and reports each file's state.
- `filenode_map_list()` returns relation name, OID, and filenode entries for the current database.
- `filenode_map_list_global()` returns the corresponding entries from the global map.

### Operational Boundaries

The upstream project labels the extension experimental and intended only for education and diagnostics. It reads server-side cluster files, so install it and grant execution only to trusted administrators. A reported error is evidence for further investigation; preserve the data directory and use supported backup, restore, or recovery procedures rather than editing mapping files by hand.
