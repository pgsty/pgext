## Usage

Sources:

- [Prototype README at the reviewed commit](https://github.com/jessepinkman9900/code-snippets/blob/7b05250d7fe220efbbdee7c4aeeae27156d04cd8/pgrx-data-retention/README.md)
- [Extension control file](https://github.com/jessepinkman9900/code-snippets/blob/7b05250d7fe220efbbdee7c4aeeae27156d04cd8/pgrx-data-retention/pg_data_retention/pg_data_retention.control)
- [Background-worker source](https://github.com/jessepinkman9900/code-snippets/blob/7b05250d7fe220efbbdee7c4aeeae27156d04cd8/pgrx-data-retention/pg_data_retention/src/lib.rs)
- [Cargo feature matrix](https://github.com/jessepinkman9900/code-snippets/blob/7b05250d7fe220efbbdee7c4aeeae27156d04cd8/pgrx-data-retention/pg_data_retention/Cargo.toml)

`pg_data_retention` is a version 0.0.0 background-worker experiment that deletes rows according to entries in `public.data_retention_policy`. It must be present in `shared_preload_libraries`, followed by a server restart.

```sql
CREATE EXTENSION pg_data_retention;
SELECT *
FROM public.data_retention_policy
ORDER BY database_name, schema_name, table_name;
```

### Destructive Prototype Warning

At startup, current source connects to the `postgres` database, drops `public.data_retention_policy` if it exists, and recreates it. The orchestrator wakes every ten seconds, launches deletion workers sequentially, and panics on an SQL failure. The documented `cron_schedule` field is not implemented, and the repository explicitly lists missing worker and cross-version tests.

Do not preload this code on a server containing valuable data. Rebuilding the policy table destroys configuration, and a mistaken table, timestamp column, retention period, or batch size can delete application rows. Treat the source only as a prototype to audit and redesign in an isolated environment.
