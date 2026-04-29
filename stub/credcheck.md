


## Usage

Sources: [README](https://github.com/HexaCluster/credcheck#readme), [release 4.7](https://github.com/HexaCluster/credcheck/releases/tag/v4.7)

`credcheck` enforces configurable rules for PostgreSQL usernames and passwords during `CREATE ROLE`, `ALTER ROLE`, password changes, and role renames. It can reject weak credentials, enforce password expiration windows, track password reuse, ban users after repeated authentication failures, delay failed authentication responses, force first-login password changes, and block password changes for ordinary users.

### Required Setup

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'credcheck'
```

Restart PostgreSQL after changing preload libraries. Password reuse history, authentication failure banning, first-login password changes, and login-time expiry warnings depend on preload or login-event support described in the upstream README.

### Configuration Parameters

#### Username Checks

| Parameter | Description | Example |
|-----------|-------------|---------|
| `credcheck.username_min_length` | Minimum username length | `4` |
| `credcheck.username_min_special` | Minimum special characters | `1` |
| `credcheck.username_min_digit` | Minimum digit characters | `1` |
| `credcheck.username_min_upper` | Minimum uppercase characters | `2` |
| `credcheck.username_min_lower` | Minimum lowercase characters | `1` |
| `credcheck.username_min_repeat` | Max adjacent repeat characters | `2` |
| `credcheck.username_contain` | Must contain one of these chars | `a,b,c` |
| `credcheck.username_not_contain` | Must not contain these chars | `x,y,z` |
| `credcheck.username_contain_password` | Username must not contain password | `on` |
| `credcheck.username_ignore_case` | Ignore case for username checks | `on` |

#### Password Checks

| Parameter | Description | Example |
|-----------|-------------|---------|
| `credcheck.password_min_length` | Minimum password length | `8` |
| `credcheck.password_min_special` | Minimum special characters | `1` |
| `credcheck.password_min_digit` | Minimum digit characters | `1` |
| `credcheck.password_min_upper` | Minimum uppercase characters | `1` |
| `credcheck.password_min_lower` | Minimum lowercase characters | `1` |
| `credcheck.password_min_repeat` | Max adjacent repeat characters | `3` |
| `credcheck.password_contain_username` | Password must not contain username | `on` |
| `credcheck.password_contain` | Must contain one of these chars | `a,b,c` |
| `credcheck.password_not_contain` | Must not contain these chars | `!@=$#` |
| `credcheck.password_ignore_case` | Ignore case for password checks | `on` |
| `credcheck.password_valid_until` | Minimum days for VALID UNTIL | `60` |
| `credcheck.password_valid_max` | Maximum days for VALID UNTIL | `365` |
| `credcheck.password_valid_warning` | Warn before password expiry; PostgreSQL 17+ login event trigger | `7` |
| `credcheck.password_change_first_login` | Force a new user to change password before normal queries | `true` |
| `credcheck.whitelist` | Usernames excluded from checks | `admin,super` |
| `credcheck.superuser_nocheck` | Skip policy checks for changes made by a superuser | `on` |
| `credcheck.disallow_password_change` | Disallow users from changing their own password | `on` |

If built with cracklib support, `credcheck` can also reject passwords that are easy to crack.

### Examples

```sql
-- Rejected: username too short
CREATE USER abc WITH PASSWORD 'pass';
-- ERROR: username length should match the configured credcheck.username_min_length

-- Rejected: password contains username
CREATE USER abcd$ WITH PASSWORD 'abcd$xyz';
-- ERROR: password should not contain username
```

Enforce password lifetime bounds:

```sql
SET credcheck.password_valid_until = 30;
SET credcheck.password_valid_max = 180;

CREATE USER abcd$;
-- ERROR: require a VALID UNTIL option with a date older than 30 days
```

### Password Reuse Policy

```sql
SET credcheck.password_reuse_history = 2;
SET credcheck.password_reuse_interval = 365;  -- days
```

View password history:

```sql
SELECT rolename, password_hash FROM pg_password_history;
```

The upstream README says password hashes are kept in shared memory and saved to `$PGDATA/pg_password_history`, so include that file in backup planning. Use `credcheck.history_max_size` to size the cache; changing it requires a PostgreSQL restart.

### Authentication Failure Ban

```sql
SET credcheck.max_auth_failure = 3;  -- ban after 3 failures
SET credcheck.auth_delay_ms = 1000;  -- delay failed authentication
SET credcheck.whitelist_auth_failure = 'appuser1,appuser2';
```

Reset banned users:

```sql
SELECT pg_banned_role_reset();              -- reset all
SELECT pg_banned_role_reset('username');     -- reset specific user
```

`credcheck.reset_superuser` can force superusers to be exempt from banning or reset a banned superuser.

### First-Login And Password-Change Controls

Force a new user to change the password before running normal queries:

```sql
SET credcheck.password_change_first_login = true;
CREATE USER user1 PASSWORD 'Rkd89,34' VALID UNTIL '2050-12-31';
-- first login:
-- ERROR: you must change your password first.
ALTER USER user1 PASSWORD 'Zkd89,34';
```

Force the same behavior later:

```sql
ALTER USER user1 SET credcheck_internal.force_change_password = true;
```

Version 4.7 adds `credcheck.disallow_password_change` for sites where users must not change their own password:

```sql
SET credcheck.disallow_password_change = on;
ALTER ROLE user1 PASSWORD 'My-New-Pass#123';
-- ERROR: you are not allowed to change your password.
```
