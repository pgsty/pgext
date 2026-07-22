## Usage

Sources:

- [TencentDB for PostgreSQL role and privilege documentation](https://cloud.tencent.com/document/product/409/61551)
- [TencentDB for PostgreSQL extension support matrix](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_superuser` is a TencentDB for PostgreSQL provider extension associated with the managed `pg_tencentdb_superuser` role. The role supplies a controlled substitute for native superuser capabilities on Tencent Cloud; neither the extension nor its privilege model is portable to self-managed PostgreSQL.

### Core Workflow

Use the TencentDB extension-management surface or `CREATE EXTENSION` only when the current engine's support matrix exposes the extension. Then verify its version and the managed role before granting membership to a dedicated administrator:

```sql
CREATE EXTENSION tencentdb_superuser;

SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_superuser';

SELECT rolname, rolcreatedb, rolcreaterole, rolreplication, rolbypassrls
FROM pg_roles
WHERE rolname = 'pg_tencentdb_superuser';

GRANT pg_tencentdb_superuser TO dba_role;
```

The grant must be issued by an account that TencentDB permits to administer the managed role. Audit membership and test the required operation on the exact engine version rather than assuming native-superuser equivalence.

### Confirmed Privilege Boundary

- The managed role has `CREATEDB`, `CREATEROLE`, `REPLICATION`, and `BYPASSRLS`-style capabilities and broad access to objects owned by non-superuser roles.
- It can create supported publications, subscriptions, replication slots, and extensions. During supported extension creation, TencentDB temporarily elevates the current managed role through the required superuser checks.
- Native `superuser`, `pg_execute_server_program`, `pg_read_server_files`, and `pg_write_server_files` are not exposed to customers.
- Library loading is limited to the service plugin directory, and only another managed superuser can terminate a managed-superuser backend. Checkpoint and event-trigger capabilities also depend on the TencentDB engine revision.

### Service and Version Boundaries

The official matrix lists `tencentdb_superuser` versions 1.0 and 1.1 across different PostgreSQL engine majors; availability is not uniform. The catalog version `1.1` must therefore be treated as an engine-specific service version, not a generally downloadable package. Grant the role only to trusted human or automation administrators, never to an application login; separate ownership from day-to-day access, log membership changes and privileged DDL, and re-check TencentDB documentation after engine upgrades or service-policy changes.
