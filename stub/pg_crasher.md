## Usage

Sources:

- [Official README](https://github.com/davidcrawford/pg_crasher/blob/fdfb24d7b7087a211c940cf96473069103c42a95/README)
- [Official extension SQL](https://github.com/davidcrawford/pg_crasher/blob/fdfb24d7b7087a211c940cf96473069103c42a95/pg_crasher--1.0.sql)
- [Official C implementation](https://github.com/davidcrawford/pg_crasher/blob/fdfb24d7b7087a211c940cf96473069103c42a95/pg_crasher.c)

`pg_crasher` deliberately crashes the PostgreSQL backend process that calls it. Its only legitimate use is controlled testing of client disconnect handling, transaction recovery, connection pools, monitoring, and failover behavior; never install version `1.0` in a shared production database.

### Safe Installation Boundary

The extension creates a function that is executable by `PUBLIC` under PostgreSQL's normal defaults. Revoke that privilege in the same transaction as installation and grant it only to a dedicated test role if required:

```sql
BEGIN;
CREATE EXTENSION pg_crasher;
REVOKE ALL ON FUNCTION pg_crasher() FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_crasher() TO crash_test_role;
COMMIT;
```

### Deliberate Crash Test

From an isolated client connected to a disposable instance, invoke the function once:

```sql
SELECT pg_crasher();
```

The C function writes through a null pointer. The session should terminate abruptly, its open transaction should roll back, and the postmaster should start a replacement backend for later connections.

### Operational Caveats

This is not an error-raising helper: it causes a native process crash and can trigger crash alerts, core dumps, pool churn, or broader recovery behavior. Do not run it in a transaction containing data that must commit. Confirm the exact host, cluster, database, role, and maintenance window before testing, monitor server logs and recovery, and remove the extension afterward. The reviewed project has no compatibility matrix and its last source change was in 2013, so even laboratory use requires an exact-version build and an expendable environment.
