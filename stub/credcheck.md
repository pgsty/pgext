


## Usage

> [credcheck: Credential checking for PostgreSQL usernames and passwords](https://github.com/MigOpsRepos/credcheck)

`credcheck` enforces configurable rules for username and password strength during `CREATE ROLE`, `ALTER ROLE`, and password changes. It also supports password reuse policies and authentication failure banning.

### Configuration Parameters

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'credcheck'
```

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
| `credcheck.password_valid_until` | Minimum days for VALID UNTIL | `60` |
| `credcheck.password_valid_max` | Maximum days for VALID UNTIL | `365` |
| `credcheck.whitelist` | Usernames excluded from checks | `admin,super` |

### Examples

```sql
-- Rejected: username too short
CREATE USER abc WITH PASSWORD 'pass';
-- ERROR: username length should match the configured credcheck.username_min_length

-- Rejected: password contains username
CREATE USER abcd$ WITH PASSWORD 'abcd$xyz';
-- ERROR: password should not contain username
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

### Authentication Failure Ban

```sql
SET credcheck.max_auth_failure = 3;  -- ban after 3 failures
```

Reset banned users:

```sql
SELECT pg_banned_role_reset();              -- reset all
SELECT pg_banned_role_reset('username');     -- reset specific user
```
