## Usage

Sources:

- [README](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/README.md)
- [Control](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/pgx_permission_sync.control)
- [Makefile](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/Makefile)
- [Base SQL](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/sql/pgx_permission_sync--1.0.sql)
- [Security upgrade](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/sql/pgx_permission_sync--1.0--1.1.sql)
- [Drift and reconciliation SQL](https://github.com/valehdba/pgx_permission_sync/blob/7f8069caf38f534a0f62f08721a6505805d1b317/sql/pgx_permission_sync--1.1--1.2.sql)

`pgx_permission_sync` 1.2 synchronizes permission DDL to configured PostgreSQL databases through dblink and provides drift detection, reconciliation, audit logs and retries. It is pure SQL/PL/pgSQL, supports PostgreSQL 14 and later, and requires `dblink`, `pgcrypto` and a superuser to install. It has no shared-preload requirement.

### Install and configure

The fixed schema is `dba_pgx_permission_sync`. Installation starts from 1.0 SQL and follows the supplied upgrade chain to 1.2. Disable live synchronization while configuring targets. Set an encryption key of at least 16 characters before storing passwords; without a key, the implementation warns and stores the password without encryption. In psql, supply `encryption_key` and `sync_password` variables privately. The target account needs `CREATEROLE` and relevant grant privileges.

```sql
CREATE EXTENSION pgx_permission_sync CASCADE;
SELECT dba_pgx_permission_sync.disable_sync();
SELECT dba_pgx_permission_sync.enable_dry_run();
SELECT dba_pgx_permission_sync.set_encryption_key(:'encryption_key');
SELECT dba_pgx_permission_sync.add_environment(
  p_env_name := 'staging', p_host := 'staging-db.internal',
  p_dbname := 'appdb', p_username := 'sync_admin',
  p_password := :'sync_password', p_sslmode := 'verify-full'
);
SELECT * FROM dba_pgx_permission_sync.v_environment_status;
```

### Review drift before applying

`detect_drift()` compares roles, memberships, object privileges and default privileges. `reconcile()` defaults to a dry run. Only an administrator may apply it; the global dry-run setting also forces a rehearsal. Revocation requires a separate opt-in, target-only roles are never dropped, and privileged attributes remain blocked unless explicitly allowed. Missing roles are created without passwords.

```sql
SELECT dba_pgx_permission_sync.detect_drift('staging');
SELECT * FROM dba_pgx_permission_sync.v_drift_open_items;
SELECT kind, drift_type, action, status, fix_sql
FROM dba_pgx_permission_sync.reconcile('staging');
```

### Live synchronization and operations

Live event triggers cover supported privilege, policy and schema commands. `CREATE ROLE`, `ALTER ROLE` and `DROP ROLE` are cluster-wide and do not fire these triggers; use drift reconciliation or reviewed `manual_sync()` calls. After reviewing rules and dry-run results, enable actual delivery. Remote calls are synchronous inside the source DDL transaction; a slow target delays it. dblink changes on remote databases are not a distributed atomic transaction with the source.

```sql
SELECT dba_pgx_permission_sync.disable_dry_run();
SELECT dba_pgx_permission_sync.enable_sync();
SELECT * FROM dba_pgx_permission_sync.get_sync_status();
SELECT * FROM dba_pgx_permission_sync.v_failed_syncs;
SELECT * FROM dba_pgx_permission_sync.process_retry_queue();
```

### Access and limitations

Treat `pgsync_admin` and `pgsync_operator` as trusted administrative roles: operators can invoke manual remote SQL. `pgsync_reader` has read-only visibility. Encryption keys are held in the restricted `secrets` table; encryption does not protect them from its owner. Drift comparison excludes ownership, database privileges, large objects, passwords, per-role settings and PostgreSQL 16 membership options. `role_pattern` does not filter grantees of grant/revoke commands. Retry and log cleanup require explicit invocation or an independently configured scheduler.
