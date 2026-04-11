
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_command_fw;
> ALTER SYSTEM SET pg_command_fw.block_truncate = on;
> ALTER SYSTEM SET pg_command_fw.production_schemas = 'public,payments';
> SELECT pg_reload_conf();
> ```
>
> Source: [README](https://github.com/rustwizard/pg_command_fw)

`pg_command_fw` is a PostgreSQL command firewall. It intercepts DDL and utility commands through the `ProcessUtility` hook and blocks selected built-in file-reading functions through the post-parse analyze hook. Each command category is controlled by its own GUC.

## Setup

The extension must be preloaded:

```ini
shared_preload_libraries = 'pg_command_fw'
```

Then enable it in the database:

```sql
CREATE EXTENSION pg_command_fw;
```

## Command Categories

The upstream README documents these firewall categories:

- `TRUNCATE`
- `DROP TABLE`
- `ALTER SYSTEM`
- `LOAD`
- `COPY ... PROGRAM`
- plain `COPY`
- `pg_read_file()`, `pg_read_binary_file()`, and `pg_stat_file()`

Some categories block only non-superusers, while others block everyone including superusers. Superusers are only exempt from non-superuser categories unless they are explicitly listed in `pg_command_fw.blocked_roles`.

## Important GUCs

- `pg_command_fw.enabled` to enable or disable all checks
- `pg_command_fw.block_truncate`
- `pg_command_fw.block_drop_table`
- `pg_command_fw.production_schemas`
- `pg_command_fw.block_alter_system`
- `pg_command_fw.block_load`
- `pg_command_fw.block_copy_program`
- `pg_command_fw.block_copy`
- `pg_command_fw.block_read_file`
- `pg_command_fw.blocked_roles`
- `pg_command_fw.hint`
- `pg_command_fw.audit_log_enabled`

## Audit Log

The extension records intercepted commands in `command_fw.audit_log`. The README documents columns such as:

- timestamp
- session and current user names
- original query text
- command type
- target schema or object
- client address
- whether the command was blocked
- internal block reason

## Examples

Block `TRUNCATE` and `DROP TABLE` in production schemas:

```sql
ALTER SYSTEM SET pg_command_fw.block_truncate = on;
ALTER SYSTEM SET pg_command_fw.block_drop_table = on;
ALTER SYSTEM SET pg_command_fw.production_schemas = 'public,payments';
ALTER SYSTEM SET pg_command_fw.hint = 'Contact your DBA to request access';
SELECT pg_reload_conf();
```

Block a specific role from any governed command:

```sql
ALTER SYSTEM SET pg_command_fw.blocked_roles = 'app_deploy';
SELECT pg_reload_conf();
```
