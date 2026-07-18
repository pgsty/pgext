## Usage

Sources:

- [postgresql-migrations README at the reviewed commit](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/README.org)
- [migration.control at the reviewed commit](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/migration.control)
- [Version 0.0.1 install SQL](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/migration--0.0.1.sql)
- [Upstream migration design and examples](https://github.com/drewc/postgresql-migrations/blob/3da0fe400bfa37f92a78e8854f08522f3989079d/doc/migration.org)

`migration` is a pure-SQL migration runner. It records numbered up/down script paths in `migration_registry`, reads those files from the database server, executes their contents, and records outcomes in `migration_log` and `migration_error`.

### Register and Run Scripts

```sql
CREATE EXTENSION migration;

SELECT migration_setenv('ROOT_DIRECTORY', '/srv/postgresql/migrations');
SELECT register_migration(0, '000-up.sql', '000-down.sql');
SELECT register_migration(1, '001-up.sql', '001-down.sql');

SELECT migration_migrate_database(1);
SELECT migration_current_number();
SELECT migration_migrate_database(0);
```

The paths are resolved on the server, not on the client. The reviewed SQL exposes `migration_migrate_database`; the README's older `migration.migrate_database()` spelling does not match the installed API.

### Caveats

- `migration_read_file` implements server-side file reads through dynamic `COPY`, and `migration_eval` executes the resulting text as SQL. Restrict function privileges and the script directory to trusted database operators.
- Up and down scripts run with the caller's database privileges and can make arbitrary schema or data changes. Review, checksum, back up, and rehearse them before execution.
- Version `0.0.1` is 2018-era code with no current compatibility matrix. Its recursive runner and custom error logging are not a substitute for release orchestration, locking, or an external audit trail.
