


## Usage

> [set_user: User switching with enhanced logging and control](https://github.com/pgaudit/set_user)

`set_user` allows switching users and optional privilege escalation with enhanced audit logging. It provides an additional layer of control when unprivileged users must escalate to superuser or object owner roles for maintenance tasks.

```sql
CREATE EXTENSION set_user;
```

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'set_user'
```

| Parameter | Default | Description |
|-----------|---------|-------------|
| `set_user.block_alter_system` | `on` | Block ALTER SYSTEM when escalated |
| `set_user.block_copy_program` | `on` | Block COPY PROGRAM when escalated |
| `set_user.block_log_statement` | `on` | Block SET log_statement; force `log_statement=all` for superusers |
| `set_user.superuser_allowlist` | `*` | Roles allowed to escalate to superuser |
| `set_user.nosuperuser_target_allowlist` | `*` | Roles allowed as non-superuser targets |
| `set_user.superuser_audit_tag` | `AUDIT` | Tag appended to log_line_prefix on escalation |

### Functions

```sql
-- Switch to a non-superuser role
SELECT set_user('dbclient');

-- Escalate to superuser (requires EXECUTE on set_user_u)
SELECT set_user_u('postgres');

-- Switch with a token (token required for reset)
SELECT set_user('dbclient', 'my_secret_token');

-- Reset back to original user
SELECT reset_user();
SELECT reset_user('my_secret_token');  -- if token was used

-- Irrevocable session auth switch
SELECT set_session_auth('target_role');
```

### Permission Setup

```sql
-- Allow role to switch to non-superuser roles
GRANT EXECUTE ON FUNCTION set_user(text) TO admin;

-- Allow role to escalate to superuser
GRANT EXECUTE ON FUNCTION set_user_u(text) TO dba;
```

### Behavior on Escalation

When escalating to a superuser role:
- The role transition is logged with a specific notation
- `ALTER SYSTEM` and `COPY PROGRAM` are blocked (if configured)
- `log_statement` is forced to `all` for full audit trail
- The `AUDIT` tag is appended to `log_line_prefix`
