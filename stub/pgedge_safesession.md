## Usage

Sources:

- [Official documentation](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/docs/index.md)
- [Official README](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/README.md)
- [Hook and GUC implementation](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/src/pgedge_safesession.c)

`pgedge_safesession` 1.0-alpha1 enforces read-only sessions for configured PostgreSQL roles. Executor and utility hooks block DML, DDL, dangerous utility commands, and selected C functions, while PostgreSQL's read-only transaction state adds a second layer. It is a preview security control for PostgreSQL 14+, not a substitute for normal least-privilege grants.

### Preload and Configure

The hooks work only when loaded at server startup:

```conf
shared_preload_libraries = 'pgedge_safesession'
```

Restart PostgreSQL, then configure as a superuser. Registering the already loaded module in `pg_extension` is optional and does not activate the hooks:

```sql
ALTER SYSTEM SET pgedge_safesession.roles =
  'readonly_user,reporting_role';
SELECT pg_reload_conf();
```

Restrictions are anchored to `session_user`. Members of a configured role are restricted too; `SET ROLE` and a privileged `SECURITY DEFINER` owner cannot escape the check. A session whose login role is superuser is always exempt.

### Protection Controls

All controls are `SUSET`. `pgedge_safesession.block_dml`, `block_ddl`, `block_c_functions`, and `force_read_only` default on. `block_all_c_functions` defaults off, so only volatile C functions are blocked by default; enabling it also blocks stable/immutable extension calculations.

With defaults, restricted sessions may use SELECT, non-executing EXPLAIN, transaction control, safe SET/SHOW, LISTEN/NOTIFY, cursors, read-only PL/pgSQL/SQL functions, and ordinary `COPY TO`. Writes, MERGE, DDL, COPY FROM, COPY TO PROGRAM, CTAS/SELECT INTO, GRANT/REVOKE, VACUUM/ANALYZE, strong table locks, protected read-only GUC changes, and volatile C functions are blocked.

### Security Boundaries

Role grants remain the primary control: SafeSession does not make readable data safe, filter rows, mask values, authenticate users, or protect against exempt superusers. Review every allowed utility operation for the deployment and use `block_all_c_functions` if C volatility labels are not trusted. Because this is alpha software built around server hooks, validate command coverage, extension interactions, prepared statements, connection pooling, reloads, and failover on the exact PostgreSQL major before treating it as a security boundary.
