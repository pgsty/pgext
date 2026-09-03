## Usage

Sources:

- [Official documentation](https://github.com/BRupireddy2/pg_misc_functions/blob/9e4718148093375fd5b563a5c4fea987558a0502/README.md)
- [Extension control file](https://github.com/BRupireddy2/pg_misc_functions/blob/9e4718148093375fd5b563a5c4fea987558a0502/pg_misc_functions.control)
- [Official repository](https://github.com/BRupireddy2/pg_misc_functions)

`pg_misc_functions` Administrative test functions for backend errors, PANIC, crashes, and failover exercises.

### Enablement

Install the files for the intended server, then create `pg_misc_functions` in the target database:

```sql
CREATE EXTENSION pg_misc_functions;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT pg_current_wal_tli();
SELECT pg_control_checkpoint_tli();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_signal_backend` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_cause_fatal` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_cause_panic` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_control_checkpoint_previous_tli` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_control_checkpoint_tli` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_current_wal_tli` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_last_wal_receive_tli` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_last_wal_replay_tli` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
