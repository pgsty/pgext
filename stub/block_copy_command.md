## Usage

- Source: [README](https://github.com/rustwizard/block_copy_command/blob/master/README.md)

`block_copy_command` installs a `ProcessUtility` hook that intercepts `COPY` statements. The hook is cluster-wide once the library is loaded, while `CREATE EXTENSION` only registers metadata in a database.

### Enable It

```conf
shared_preload_libraries = 'block_copy_command'
```

```sql
CREATE EXTENSION block_copy_command;
```

The upstream README lists PostgreSQL 13-18 support.

### Blocking Rules

By default, non-superusers cannot run `COPY TO` or `COPY FROM`:

```sql
COPY my_table TO STDOUT;
COPY my_table FROM STDIN;
COPY (SELECT * FROM my_table) TO '/tmp/out.csv';
```

Priority is documented as:

- `block_copy_command.blocked_roles`: always blocked, even superusers.
- `block_copy_command.block_program = on`: blocks `COPY ... PROGRAM` for everyone.
- `block_copy_command.enabled = off`: allows `COPY` for roles not in `blocked_roles`.
- Superusers otherwise bypass direction blocking.
- `block_copy_command.block_to` and `block_copy_command.block_from` control export/import blocking for non-superusers.

### Main Settings

- `block_copy_command.enabled`: master switch for non-superuser blocking.
- `block_copy_command.block_to`: block `COPY TO`.
- `block_copy_command.block_from`: block `COPY FROM`.
- `block_copy_command.block_program`: block `COPY TO/FROM PROGRAM` for all users.
- `block_copy_command.hint`: append a custom `HINT` to blocked-command errors.
- `block_copy_command.blocked_roles`: comma-separated always-blocked roles.
- `block_copy_command.audit_log_enabled`: write intercepted events to the audit table.

### Audit And Caveats

Allowed and blocked attempts are intercepted, and the extension defines `block_copy_command.audit_log` plus server-log entries for blocked events. The README notes one important caveat: blocked audit rows are inserted before the error is raised, so they are rolled back with the transaction. In practice, PostgreSQL server logs are the authoritative record for blocked `COPY` attempts.
