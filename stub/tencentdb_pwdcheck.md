## Usage

Sources:

- [Official TencentDB for PostgreSQL extension matrix](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_pwdcheck` is a provider-only extension available on TencentDB for PostgreSQL. Use it only within that managed-service boundary and obtain the password-policy semantics from the TencentDB console or support channel: the public matrix confirms availability but does not document SQL functions, policy parameters, preload behavior, or source code.

### Core Workflow

The current TencentDB matrix lists version `1.0` for PostgreSQL 10 through 17 and no entry for PostgreSQL 18. First verify that the extension is offered by the exact instance engine version, then enable it using TencentDB's supported extension procedure:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'tencentdb_pwdcheck';

CREATE EXTENSION tencentdb_pwdcheck;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'tencentdb_pwdcheck';
```

If the first query returns no row, do not copy files or substitute the community `passwordcheck` module; choose a supported engine version or contact Tencent Cloud.

### Provider Boundary

The cited TencentDB page documents package availability only. It does not specify which password operations are checked, how rules are configured, whether existing credentials are evaluated, or whether a restart is required. Treat those details as provider-managed and revision-specific, and verify them against the console help or a Tencent support response for the target instance.

### Operational Notes

Test activation and password changes on a disposable role before applying a policy to production users. Coordinate database policy changes with application secret rotation and preserve an independently tested administrative recovery path; a stricter or misunderstood password rule can block user creation, password rotation, or emergency access.
