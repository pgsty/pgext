## Usage

Sources:

- [Archived snapshot of the official upstream repository](https://github.pkg.st/dmtkfs/pg-tamper-log)
- [Last-known official repository URL](https://github.com/dmtkfs/pg-tamper-log)

`pg_tamperlog` is an educational pure-SQL/PL/pgSQL extension for append-only audit events. Each `audit_log` row stores JSON event data, a timestamp, the previous hash, and a SHA-256 current hash; a `BEFORE INSERT` trigger extends the chain, while `tamper_log_verify` reports integrity failures.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_tamperlog;

INSERT INTO audit_log (event)
VALUES ('{"user":"alice","action":"login"}');

SELECT *
FROM tamper_log_verify
WHERE integrity_check IS NOT NULL;
```

Version 1.1 can optionally use `pg_tamperlog_rust` for faster hashing. Verification scans the whole table, so schedule it with explicit I/O and latency budgets. A hash chain makes edits or deletions detectable only when a trustworthy checkpoint or copy of the expected chain head exists; a database superuser can alter tables, functions, triggers, and verification code.

Upstream explicitly describes the project as educational. Separate log writers from administrators, revoke update/delete and DDL privileges, serialize concurrent appends, export signed checkpoints and backups off the database host, and test rollback, restore, truncation, reordering, failover, and keyless-hash threat assumptions.

The official GitHub repository now returns HTTP 404; the source above is an archived rendering of that repository. Do not install an unverified binary or fork. Recover a content-addressed source tree and audit the exact SQL before production use.
