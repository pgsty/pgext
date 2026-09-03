## Usage

Sources:

- [Official documentation](https://github.com/GreengageDB/greengage/blob/3253a8120ecdbafb29019740b90af5bef07790c2/gpcontrib/gg_tables_tracking/README.md)
- [Extension control file](https://github.com/GreengageDB/greengage/blob/3253a8120ecdbafb29019740b90af5bef07790c2/gpcontrib/gg_tables_tracking/gg_tables_tracking.control)
- [Official repository](https://github.com/GreengageDB/greengage)

`gg_tables_tracking` Greengage relation-size change tracking for monitoring and maintenance.

### Enablement

Merge `gg_tables_tracking` into the existing preload list, restart PostgreSQL, and then create `gg_tables_tracking` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'gg_tables_tracking'
```

```sql
CREATE EXTENSION gg_tables_tracking;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT * FROM gg_wait_sampling.gg_wait_sampling_reset_profile;
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- Preloading `gg_tables_tracking` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
