## Usage

Sources:

- [README.md](https://github.com/pierreforstmann/pg_start_sql/blob/a94782661c3f94a7f779a3cd61200d771fdba04a/README.md)
- [pg_start_sql.control](https://github.com/pierreforstmann/pg_start_sql/blob/a94782661c3f94a7f779a3cd61200d771fdba04a/pg_start_sql.control)
- [pg_start_sql--1.0.0.sql](https://github.com/pierreforstmann/pg_start_sql/blob/a94782661c3f94a7f779a3cd61200d771fdba04a/pg_start_sql--1.0.0.sql)

`pg_start_sql` runs configured SQL when a PostgreSQL instance starts. The reviewed control/installation SQL is 1.0.0; upstream reports testing on PostgreSQL 9.5–19. All startup statements run as a superuser, with no setting to choose a less privileged execution role.

### Configure Startup SQL

Add the library to the existing preload list and configure the target database and statement:

```ini
shared_preload_libraries = 'pg_start_sql'
pg_start_sql.dbname = 'postgres'
pg_start_sql.stmt = 'SELECT 1'
```

Restart PostgreSQL for preload, and follow upstream's extension setup:

```sql
CREATE EXTENSION pg_start_sql;
```

The worker is instance-wide; the extension need not be created in every database. `pg_start_sql.dbname` defaults to postgres. At least one of `pg_start_sql.stmt` or `pg_start_sql.file` must be configured; the latter selects a file containing startup statements.

### Operational Boundary

Restrict configuration and SQL-file writes to trusted administrators because the statements execute with superuser privileges. Use startup work that is safe to repeat across server starts and verify its result and server logs. Preserve other preloaded libraries. The control file permits relocation of SQL objects, but that does not remove the cluster-level preload and restart requirements.
