## Usage

Sources:

- [Extension control file](https://github.com/MIuserX/pgsword/blob/505a5bd993303dd6d470da93484843df768d260a/pgsword.control)
- [Empty extension SQL](https://github.com/MIuserX/pgsword/blob/505a5bd993303dd6d470da93484843df768d260a/pgsword--1.0.sql)
- [Hook implementation](https://github.com/MIuserX/pgsword/blob/505a5bd993303dd6d470da93484843df768d260a/pgsword.c)
- [Audit rules](https://github.com/MIuserX/pgsword/blob/505a5bd993303dd6d470da93484843df768d260a/rule.c)

`pgsword` version `1.0` is a PostgreSQL 10-era Qunar SQL-audit prototype. The install SQL creates no database objects; the library must be preloaded to register its parser and utility hooks. Keep auditing disabled globally:

```conf
shared_preload_libraries = 'pgsword'
pgsword.enabled = off
```

After restart, register the extension. A superuser can enable it only in an isolated audit session and submit candidate DDL:

```sql
CREATE EXTENSION pgsword;
SHOW pgsword.enabled;

SET pgsword.enabled = on;

-- Audit only: this statement is deliberately aborted after checks.
CREATE TABLE audit_candidate (
  id bigserial PRIMARY KEY,
  payload jsonb,
  created_at timestamptz
);
```

### Blocking behavior and compatibility

When `pgsword.enabled` is on, most utility statements end in the extension's `finishAudit()` path, which raises an error even when all audit rules pass; the candidate DDL is not executed. Rules reject various names and types, including non-lowercase identifiers, a primary key not named `id`, `timestamp` instead of `timestamptz`, and `json` instead of `jsonb`. The source and Makefile target PostgreSQL 10.1 internal APIs and provide no README, license, or regression suite. Do not enable this library on normal sessions, and port/review it before any modern-server use.
