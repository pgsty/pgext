## Usage

Sources:

- [Official documentation](https://github.com/VladUl287/table_change_tracker/blob/407d7b365b1ec2af9a32ba0203d56b7e838e2386/README.md)
- [Extension control file](https://github.com/VladUl287/table_change_tracker/blob/407d7b365b1ec2af9a32ba0203d56b7e838e2386/table_change_tracker.control)
- [Official repository](https://github.com/VladUl287/table_change_tracker)

`table_change_tracker` Tracks last modification timestamps for selected tables in shared memory.

### Enablement

Merge `table_change_tracker` into the existing preload list, restart PostgreSQL, and then create `table_change_tracker` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'table_change_tracker'
```

```sql
CREATE EXTENSION table_change_tracker;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Get timestamps for multiple tables
SELECT get_last_timestamps(ARRAY['public.users', 'public.orders']::regclass[]);
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `enable_table_tracking` | FUNCTION | Callable function from the reviewed install surface. |
| `disable_table_tracking` | FUNCTION | Callable function from the reviewed install surface. |
| `get_last_timestamp` | FUNCTION | Callable function from the reviewed install surface. |
| `get_last_timestamps` | FUNCTION | Callable function from the reviewed install surface. |
| `is_table_tracked` | FUNCTION | Callable function from the reviewed install surface. |
| `set_last_timestamp` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Preloading `table_change_tracker` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
