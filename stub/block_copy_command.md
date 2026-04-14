## Usage
- GitHub Repo: [`rustwizard/block_copy_command`](https://github.com/rustwizard/block_copy_command)
- README: [rustwizard/block_copy_command/blob/master/README.md](https://github.com/rustwizard/block_copy_command/blob/master/README.md)

`block_copy_command` blocks `COPY` commands cluster-wide by installing a `ProcessUtility` hook. It is loaded with `shared_preload_libraries`, and `CREATE EXTENSION` only registers the extension metadata in each database.

This extension is intended for deployments that want to stop `COPY TO` and `COPY FROM` by default for non-superusers, while still allowing finer-grained policy through GUCs and an audit table.

### Setup

```conf
shared_preload_libraries = 'block_copy_command'
```

```sql
CREATE EXTENSION block_copy_command;
```

The README says the hook becomes active for the whole cluster as soon as the library is loaded.

### Blocking Rules

By default, non-superusers are blocked from running `COPY`.

```sql
COPY my_table TO STDOUT;
COPY my_table FROM STDIN;
COPY (SELECT * FROM my_table) TO '/tmp/out.csv';
```

Superusers bypass the block unless they are listed in `block_copy_command.blocked_roles` or `block_copy_command.block_program` is enabled. `COPY ... PROGRAM` is blocked for everyone by default.

### Settings

- `block_copy_command.enabled` toggles blocking for non-superusers.
- `block_copy_command.block_to` controls whether `COPY TO` is blocked.
- `block_copy_command.block_from` controls whether `COPY FROM` is blocked.
- `block_copy_command.block_program` blocks `COPY TO/FROM PROGRAM` for all users.
- `block_copy_command.hint` appends a custom `HINT:` to blocked commands.
- `block_copy_command.blocked_roles` permanently blocks named roles, including superusers.
- `block_copy_command.audit_log_enabled` controls whether intercepted `COPY` events are written to `block_copy_command.audit_log`.

### Audit Log

The extension records intercepted `COPY` activity in `block_copy_command.audit_log` and also writes blocked events to the PostgreSQL server log at `LOG` level.

Typical monitoring queries from the README include listing recent events, filtering blocked events, and grouping by user.

### Scope

The upstream README covers requirements, enablement, blocking behavior, the main GUCs, the audit table, and test coverage. No separate project homepage or docs site was needed for this stub.
