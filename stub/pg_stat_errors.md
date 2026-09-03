## Usage

Sources:

- [Official documentation](https://github.com/akonorev/pg_stat_errors/blob/af98390e57f5cc529765c5efcf20a4bd6ddb04af/README.rst)
- [Extension control file](https://github.com/akonorev/pg_stat_errors/blob/af98390e57f5cc529765c5efcf20a4bd6ddb04af/pg_stat_errors.control)
- [Official repository](https://github.com/akonorev/pg_stat_errors)

`pg_stat_errors` Cluster-wide error-class counters with a bounded sample of recent errors.

### Enablement

Merge `pg_stat_errors` into the existing preload list, restart PostgreSQL, and then create `pg_stat_errors` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_stat_errors'
```

```sql
CREATE EXTENSION pg_stat_errors;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
displays the last errors that occured in the database. This view contains up to
``pg_stat_errors.max_last`` rows.

+---------------+----------------+-------------------------------------------------------+
| Name          | Type           | Description                                           |
+===============+================+=======================================================+
| error_time    | timestamp with | Time of occurrence of the error                       |
|               | time zone      |                                                       |
+---------------+----------------+-------------------------------------------------------+
| userid        | oid            | User OID                                              |
+---------------+----------------+-------------------------------------------------------+
| dbid          | oid            | Database OID                                          |
+---------------+----------------+-------------------------------------------------------+
| query         | text           | Text of the query                                     |
+---------------+----------------+-------------------------------------------------------+
| error_level   | text           | Error level (WARNING, ERROR, FATAL and PANIC)         |
+---------------+----------------+-------------------------------------------------------+
| error_state   | text           | Error state as a five-character code                  |
+---------------+----------------+-------------------------------------------------------+
| error_message | text           | Error message                                         |
+---------------+----------------+-------------------------------------------------------+


dba_stat_errors_last view
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_stat_errors_total_errors` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_stat_errors_info` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_stat_errors_last` | FUNCTION | Callable function from the reviewed install surface. |
| `dba_stat_errors` | VIEW | Inspection or query view created by the extension. |
| `dba_stat_errors_last` | VIEW | Inspection or query view created by the extension. |
| `pg_stat_errors_reset` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Preloading `pg_stat_errors` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
