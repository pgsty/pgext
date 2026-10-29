## Usage

Sources:

- [README.md](https://github.com/gliatti/pg_pass_policy/blob/ad47a9b8902ad6375d6afa20aea38d383dfc37c8/README.md)
- [pg_pass_policy.control](https://github.com/gliatti/pg_pass_policy/blob/ad47a9b8902ad6375d6afa20aea38d383dfc37c8/pg_pass_policy.control)
- [pg_pass_policy--1.0.sql](https://github.com/gliatti/pg_pass_policy/blob/ad47a9b8902ad6375d6afa20aea38d383dfc37c8/sql/pg_pass_policy--1.0.sql)
- [pg_pass_policy.c](https://github.com/gliatti/pg_pass_policy/blob/ad47a9b8902ad6375d6afa20aea38d383dfc37c8/src/pg_pass_policy.c)

`pg_pass_policy` manages PostgreSQL role password complexity, reuse history, expiry, rotation and authentication lockout. This source snapshot installs SQL version 1.0 and documents PostgreSQL 14–18. It controls cluster-wide authentication through hooks, but stores its policy in one designated database.

### Enablement

Review an administrative exemption and rescue path before enabling it. Add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_pass_policy'
pg_pass_policy.database = 'postgres'
pg_pass_policy.exempt_roles = 'postgres'
```

Connect as a superuser to the designated database only:

```sql
CREATE EXTENSION pg_pass_policy;
SELECT pass_policy.worker_status();
SELECT * FROM pass_policy.v_role_status;
```

During startup, non-exempt logins can be refused with 57P03 until the worker loads lock state, including trust/peer connections. The separately configured rescue path admits only listed superusers using peer authentication. A missing worker slot or database can prevent initialization; do not enable this without a tested administrative recovery path.

### Profiles and Administration

```sql
SELECT pass_policy.create_profile('application',
  '{"min_length":24,"max_age":"30 days","min_age":"0"}');
SELECT pass_policy.assign_profile('app_user', 'application');
SELECT * FROM pass_policy.effective_profile('app_user');
SELECT * FROM pass_policy.v_expiring;
SELECT * FROM pass_policy.v_locked;
SELECT * FROM pass_policy.v_recent_audit;
```

The default profile combines complexity, history, expiry and failed-login thresholds. SCRAM-SHA-256 is required by default. Password changes belong in `pg_pass_policy.database`; changes elsewhere are rejected when enforcement is enabled.

Installation creates the fixed `pass_policy` schema plus `pass_policy_viewer` and `pass_policy_admin` roles. Administration functions are SECURITY DEFINER and check `session_user`: connect directly as the authorized member rather than using SET ROLE. Restrict administrative membership because it can rotate non-privileged users' credentials.

### Rotation and Security Boundaries

`pass_policy.rotate_password` generates and applies a password; `pass_policy.switch_rotation_pair` supports paired-role rotation. Both return plaintext to the caller once. Protect transport with verified TLS and avoid retaining results in client grids, logs or terminal history; server-side handling does not protect client copies.

`pass_policy.generate_password`, `pass_policy.check_password`, `pass_policy.failed_attempts` and `pass_policy.worker_status` expose generation, preflight and operating state. Exempt roles cannot be locked or expired by the extension. Dropping the extension leaves the schema and administrative roles; remove them separately only after reviewing dependencies. This source-only snapshot has no tagged release or established production acceptance in this catalog.
