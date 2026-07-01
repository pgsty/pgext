


## Usage

- Source: [README](https://github.com/rustwizard/pg_command_fw/blob/master/README.md)

`pg_command_fw` is a PostgreSQL command firewall. It intercepts DDL and utility commands through the `ProcessUtility` hook and blocks selected built-in file-reading functions through the post-parse analyze hook. Each command category is controlled by its own GUC.

### Enable It

The extension must be preloaded:

```ini
shared_preload_libraries = 'pg_command_fw'
```

Then enable it in the database:

```sql
CREATE EXTENSION pg_command_fw;
```

Pigsty package metadata is version `0.1.0` for PostgreSQL 15-18 and notes that preloading is required to activate hooks for all sessions. The upstream README also documents PostgreSQL 15-18 support.

### Command Categories

The upstream README documents these firewall categories:

- `TRUNCATE`: `pg_command_fw.block_truncate`, default `on`, blocks non-superusers.
- `DROP TABLE`: `pg_command_fw.block_drop_table`, default `off`, blocks non-superusers when enabled.
- `ALTER SYSTEM`: `pg_command_fw.block_alter_system`, default `on`, blocks everyone.
- `LOAD`: `pg_command_fw.block_load`, default `on`, blocks everyone.
- `COPY ... PROGRAM`: `pg_command_fw.block_copy_program`, default `on`, blocks everyone.
- plain `COPY`: `pg_command_fw.block_copy`, default `off`, blocks non-superusers when enabled.
- `pg_read_file()`, `pg_read_binary_file()`, and `pg_stat_file()`: `pg_command_fw.block_read_file`, default `on`, blocks everyone.

Some categories block only non-superusers, while others block everyone including superusers. Superusers are only exempt from non-superuser categories unless they are explicitly listed in `pg_command_fw.blocked_roles`.

### Important GUCs

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

When `production_schemas` is set, `DROP TABLE` checks are limited to schema-qualified table names in those schemas; the README says unqualified names are not resolved through `search_path`.

### Audit Log

The extension records intercepted commands in `command_fw.audit_log`. The README documents columns such as:

- timestamp
- session and current user names
- original query text
- command type
- target schema or object
- client address
- whether the command was blocked
- internal block reason

Blocked audit inserts are best-effort because the row is rolled back with the blocked transaction; use the PostgreSQL server log as the authoritative record for blocked events.

### Examples

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

Temporarily disable the firewall in a maintenance session:

```sql
SET pg_command_fw.enabled = off;
TRUNCATE big_table;
SET pg_command_fw.enabled = on;
```
