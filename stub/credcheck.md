## Usage

Sources:

- [v6.0 README](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/README.md)
- [v6.0 changelog](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/ChangeLog)
- [SQL objects 6.0.0](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/sql/credcheck--6.0.0.sql)
- [Password history WAL implementation](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/credcheck.c)
- [Login-event setup](https://github.com/HexaCluster/credcheck/blob/ac9127dfe0892579a42bd2255638f9948074a151/event_trigger.sql)

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

The interval is in days. Installing the SQL extension is separate from loading the server-wide hooks. Upstream release 6.0 uses SQL extension version 6.0.0; installing upgraded files requires a restart to reload the library.

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

History contains SHA-256 password hashes, is shared across databases, and is persisted in `$PGDATA/pg_password_history`. Include that file in backup planning and protect access to the SQL history view, which in 6.0 is revoked from PUBLIC and granted to pg_read_all_stats. `credcheck.history_max_size` changes the history capacity and requires a restart.

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

### Upgrade and New Interfaces in 6.0

6.0 moves password history into an integer-indexed hash table in an mmap file. The first upgraded startup imports 5.x history and renames the old file to pg_password_history.legacy-migrated; this happens on both primary and standby. Preserve and verify the history backup before upgrading. After installing matching libraries and restarting, run `ALTER EXTENSION credcheck UPDATE TO '6.0.0'` in databases using its SQL objects. Current SQL revokes PUBLIC access to the history function/view and grants view access to pg_read_all_stats; review existing custom grants.

`credcheck.password_valid_min` sets a short initial password lifetime for new roles; the default 0 disables it. Existing roles changing their password retain the previous lifetime policy. New `pg_check_password(username, password)` checks candidate credentials, raising an error on a violation and returning true on acceptance; it does not evaluate password reuse history.

`credcheck.lastlog` defaults off. When enabled, pg_lastlog exposes sessions and server startup/shutdown events; optional SQL tracking records query text, so define access and retention deliberately. PgBouncer integration for failed-attempt limits and first-login password changes needs the documented extra SQL and pooler hooks; preload alone does not configure it. `credcheck.no_password_logging` now masks password literals while retaining the error statement structure. This release also fixes self-ALTER ROLE privilege escalation and memory defects.
