## Usage

Sources:

- [Official documentation](https://github.com/hank-cp/pg_op_log/blob/731aeea2b8053a3dde031276729401375a598c4c/README.md)
- [Extension control file](https://github.com/hank-cp/pg_op_log/blob/731aeea2b8053a3dde031276729401375a598c4c/op_log.control)
- [Official repository](https://github.com/hank-cp/pg_op_log)

`op_log` Trigger-based row change history recorded after transaction commit.

### Enablement

Install the files for the intended server, then create `op_log` in the target database:

```sql
CREATE EXTENSION op_log CASCADE;
```

The reviewed control or official workflow requires `plv8`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT op_log_disable('your_table_name'::regclass);
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `op_log_enable` | FUNCTION | Callable function from the reviewed install surface. |
| `data_op_log` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `op_log_disable` | FUNCTION | Callable function from the reviewed install surface. |
| `op_log_diff_array` | FUNCTION | Callable function from the reviewed install surface. |
| `op_log_diff_normal` | FUNCTION | Callable function from the reviewed install surface. |
| `op_log_diff_object` | FUNCTION | Callable function from the reviewed install surface. |
| `op_log_diff_object_array` | FUNCTION | Callable function from the reviewed install surface. |
| `op_log_diff_other` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
