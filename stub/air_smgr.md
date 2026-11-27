## Usage

Sources:

- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/README.md)
- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/extensions/air_smgr/air_smgr.control)
- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/extensions/air_smgr/air_smgr--1.0.sql)
- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/extensions/air_smgr/air_smgr.c)
- [Official source](https://github.com/Mahadi-rsio/air/blob/a19a53bea07a4f490937db50deec11cb64566525/docs/architecture/pageserver.md)

`air_smgr` 1.0 is AIR's experimental pageserver storage-manager module for its patched PostgreSQL 17. It depends on the project's storage-manager hook patch and cooperating AIR pageserver; it is not a drop-in module for stock PostgreSQL.

### Activation Boundary

The supplied SQL script contains no database objects. Activation is through startup preload, with administrative control of the server configuration. Preserve existing preload entries and restart the dedicated AIR cluster after configuration changes:

```ini
shared_preload_libraries = 'supautils,air_smgr'
air_smgr.enabled = off
```

```sql
SHOW air_smgr.enabled;
SHOW air_smgr.pageserver_socket;
SHOW air_smgr.tenant_id;
```

The example deliberately leaves storage redirection disabled. Enable it only in an AIR-provisioned cluster whose tenant identity, pageserver state, WAL flow and recovery procedure have been configured together. Loading the library during an ordinary session does not install its storage hooks.

### Settings and Data Path

The implementation registers startup-only settings `air_smgr.enabled`, `air_smgr.read_only`, `air_smgr.pageserver_socket`, `air_smgr.tenant_id`, `air_smgr.lfc_dir` and `air_smgr.lfc_max_mb`. Source defaults include a disabled module, writable mode, socket `/var/lib/air/pageserver.sock`, empty tenant identity, cache directory `/var/lib/air/lfc` and a 256 MB cache. A zero cache limit disables it. Match the socket to the actual pageserver; architecture examples use a different path.

Persistent relations use the remote storage manager when enabled; temporary relations retain the ordinary manager. Reads address pages at an LSN, with a local file cache. Writes and metadata operations require the cooperating pageserver and WAL path. The read-only setting rejects persistent writes; it is not a backup or failover protocol.

### Recovery and Compatibility

Treat tenant data, pageserver layers, WAL and configuration as one recovery unit. This module changes physical storage behavior; neither removing the library nor restoring only ordinary local relation files constitutes a documented migration back to stock PostgreSQL. The sources do not provide a SQL upgrade path or establish support for other majors. No project license file was found in the reviewed revision.
