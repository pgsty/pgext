## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/vibhorkum/edb_block_commands/blob/9cf36b955c54b47578b303bb27bae16718352863/README.md)
- [Extension control file](https://github.com/vibhorkum/edb_block_commands/blob/9cf36b955c54b47578b303bb27bae16718352863/edb_block_commands.control)
- [C implementation](https://github.com/vibhorkum/edb_block_commands/blob/9cf36b955c54b47578b303bb27bae16718352863/edb_block_commands.c)

`edb_block_commands` is an EDB Advanced Server-only policy hook that can restrict utility, DML, and read commands for superusers. The reviewed upstream says it was tested with EPAS 10 and later; it is not a portable PostgreSQL extension. Load it for selected roles through `session_preload_libraries`, then reload role sessions.

```sql
CREATE EXTENSION edb_block_commands;
ALTER ROLE guarded_admin
  SET session_preload_libraries = '$libdir/edb_block_commands';
```

Command-specific allowlists include `edb_block_commands.su_read_whitelist`, `edb_block_commands.su_write_whitelist`, and `edb_block_commands.su_alter_system_whitelist`; `edb_block_commands.su_whitelist` bypasses all restrictions for named superusers. Treat every allowlist change as a privileged security-policy change and test it in a separate EPAS instance.

The functions `edb_switch_user(text)` and especially `edb_switch_user_u(text)` cross role boundaries; upstream describes the latter as switching a normal user to a superuser. Do not grant their execution to untrusted roles. This project is experimental and has no current support or compatibility matrix, so review its C privilege checks before deployment and keep native role separation as the primary control.
