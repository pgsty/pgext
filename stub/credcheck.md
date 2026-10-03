## Usage

Sources:

- [v5.0 README](https://github.com/HexaCluster/credcheck/blob/v5.0/README.md)
- [v5.0 changelog](https://github.com/HexaCluster/credcheck/blob/v5.0/ChangeLog)
- [SQL objects 5.0.0](https://github.com/HexaCluster/credcheck/blob/v5.0/sql/credcheck--5.0.0.sql)
- [Password history WAL implementation](https://github.com/HexaCluster/credcheck/blob/v5.0/credcheck.c)
- [Login-event setup](https://github.com/HexaCluster/credcheck/blob/v5.0/event_trigger.sql)

`credcheck` enforces username and plaintext-password rules during role creation, password changes and role renames. It also tracks password reuse, bans repeated authentication failures and can require a password change at first login. Configure policy as a superuser; the defaults do not enforce a comprehensive password-strength policy.

### Enable and Set Policy

Add the library to the existing preload list and restart PostgreSQL. Install SQL objects in each database where administrators need its views and reset functions:

```ini
shared_preload_libraries = 'credcheck'
credcheck.password_min_length = 12
credcheck.password_contain_username = on
credcheck.password_reuse_history = 2
credcheck.password_reuse_interval = 365
```

```sql
CREATE EXTENSION credcheck;
CREATE ROLE app_user LOGIN PASSWORD 'example-Strong-Pass#123';
SELECT rolename, password_date FROM pg_password_history;
```

The interval is in days. Installing the SQL extension is separate from loading the server-wide hooks. Upstream release 5.0 uses SQL extension version 5.0.0; installing upgraded files requires a restart to reload the library.

### Policy Index

| Settings | Purpose |
|---|---|
| `credcheck.username_min_length`, `credcheck.username_min_special`, `credcheck.username_min_digit`, `credcheck.username_min_upper`, `credcheck.username_min_lower` | Username length and character requirements |
| `credcheck.password_min_length`, `credcheck.password_min_special`, `credcheck.password_min_digit`, `credcheck.password_min_upper`, `credcheck.password_min_lower` | Password length and character requirements |
| `credcheck.username_min_repeat`, `credcheck.password_min_repeat` | Maximum adjacent repetitions, despite the parameter names |
| `credcheck.username_contain`, `credcheck.username_not_contain`, `credcheck.password_contain`, `credcheck.password_not_contain` | Required or forbidden content |
| `credcheck.username_contain_password`, `credcheck.password_contain_username` | Reject credentials containing one another |
| `credcheck.username_ignore_case`, `credcheck.password_ignore_case` | Case handling |
| `credcheck.password_min_length_su`, `credcheck.password_valid_until_su` | Separate superuser requirements |
| `credcheck.password_valid_until`, `credcheck.password_valid_max` | Minimum and maximum password lifetime; the minimum also supplies an omitted expiry when a password is changed |
| `credcheck.whitelist`, `credcheck.superuser_nocheck` | Explicit policy exemptions |
| `credcheck.no_password_logging` | Suppress passwords in policy-error logs; enabled by default |

CrackLib strength checking is available only when the library was built with that support and its dictionary is available.

### Password History and Replication

History contains SHA-256 password hashes, is shared across databases, and is persisted in `$PGDATA/pg_password_history`. Include that file in backup planning and protect access to the SQL history view, which is granted to PUBLIC by default. `credcheck.history_max_size` changes the shared-memory capacity and requires a restart.

Version 5.0 replicates history changes through custom WAL resource manager ID 150 on PostgreSQL 15 and later. Keep the matching library preloaded on replicas and recovery servers that replay its WAL. Earlier PostgreSQL versions retain the file-backed history path without this replication support. History-reset and timestamp-test functions reject execution during recovery.

```sql
SELECT pg_password_history_reset('app_user');
```

Resetting history removes reuse protection for those records; reserve it for administrators.

### Authentication and Password Changes

```ini
credcheck.max_auth_failure = 3
credcheck.auth_delay_ms = 1000
credcheck.whitelist_auth_failure = 'service_user'
credcheck.password_change_first_login = true
```

```sql
SELECT * FROM pg_banned_role;
SELECT pg_banned_role_reset('app_user');
ALTER ROLE app_user SET credcheck_internal.force_change_password = true;
```

Bans remain until reset and their cache is lost at restart. `credcheck.reset_superuser` provides the documented superuser recovery path; `credcheck.auth_failure_cache_size` requires a restart.

Use the actual parameter `credcheck.disallow_change_password` to prohibit password changes. Even superusers are affected unless they enable `credcheck.superuser_nocheck` in their session. This exemption bypasses all corresponding role checks and must be controlled.

`credcheck.password_valid_warning` needs PostgreSQL 17 or later and the official login event trigger installed separately in every relevant database; SQL extension creation does not install that trigger.

### Plaintext Boundary

Strength and reuse checks need plaintext at password-change time. Already hashed passwords are rejected by default, including passwords sent by psql's `\password`. Setting `credcheck.encrypted_password_allowed` accepts them without providing equivalent plaintext checks. Protect the password-change connection and do not assume existing credentials are scanned retroactively. Username checks are skipped when creating a role without a password or renaming a role with no password.
