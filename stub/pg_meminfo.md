## Usage

Sources:

- [Official documentation](https://github.com/bonesmoses/pg_meminfo/blob/da919d5ff0eb8da28de934915a745526739d2cd2/README.md)
- [Extension control file](https://github.com/bonesmoses/pg_meminfo/blob/da919d5ff0eb8da28de934915a745526739d2cd2/pg_meminfo.control)
- [Official repository](https://github.com/bonesmoses/pg_meminfo)

`pg_meminfo` Backend memory and process usage inspection from SQL.

### Enablement

Install the files for the intended server, then create `pg_meminfo` in the target database:

```sql
CREATE EXTENSION pg_meminfo;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT pss FROM smap_summary WHERE pid = 4242;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `smap_summary` | VIEW | Inspection or query view created by the extension. |
| `get_all_smaps` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
