## Usage

Sources:

- [Official documentation](https://github.com/psqlmaster/pgsyswatch/blob/0dde3e65122066df75f47005aeb6b0fc951b1223/readme.md)
- [Extension control file](https://github.com/psqlmaster/pgsyswatch/blob/0dde3e65122066df75f47005aeb6b0fc951b1223/pgsyswatch.control)
- [Official repository](https://github.com/psqlmaster/pgsyswatch)

`pgsyswatch` Host process, CPU, memory, swap, network, and load statistics exposed through SQL.

### Enablement

Merge `pgsyswatch` into the existing preload list, restart PostgreSQL, and then create `pgsyswatch` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pgsyswatch'
```

```sql
CREATE EXTENSION pgsyswatch;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
select * from  pgsyswatch.net_and_loadavg_snapshots;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `proc_activity_snapshots` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pgsyswatch.net_and_loadavg` | VIEW | Inspection or query view created by the extension. |
| `proc_monitor_all` | FUNCTION | Callable function from the reviewed install surface. |
| `net_and_loadavg_snapshots` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pg_proc_activity` | VIEW | Inspection or query view created by the extension. |
| `cpu_frequencies` | FUNCTION | Callable function from the reviewed install surface. |
| `manage_partitions_maintenance` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Preloading `pgsyswatch` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
