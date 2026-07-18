## Usage

Sources:

- [Pinned upstream README and warning](https://github.com/RhodiumToad/pg-toastutils/blob/2cb6af18a866af847a672ad18b3c5e9c681b64da/README.md)
- [Version 1.0 installation SQL](https://github.com/RhodiumToad/pg-toastutils/blob/2cb6af18a866af847a672ad18b3c5e9c681b64da/pg_toastutils--1.0.sql)
- [Pinned TOAST implementation](https://github.com/RhodiumToad/pg-toastutils/blob/2cb6af18a866af847a672ad18b3c5e9c681b64da/toastutils.c)
- [Pinned extension control file](https://github.com/RhodiumToad/pg-toastutils/blob/2cb6af18a866af847a672ad18b3c5e9c681b64da/pg_toastutils.control)

pg_toastutils version 1.0 is an experimental physical-storage diagnostics module. It classifies varlena representations, exposes on-disk TOAST pointer details, checks a referenced TOAST item, and validates all TOAST references in a table including dead physical rows.

### Inspect a disposable table

```sql
CREATE EXTENSION pg_toastutils;

CREATE TABLE toast_demo (id integer, payload text);
ALTER TABLE toast_demo
    ALTER COLUMN payload SET STORAGE EXTERNAL;

INSERT INTO toast_demo
SELECT 1, repeat(md5(random()::text), 1000);

SELECT toast_type(payload), toast_ptr_detail(payload)
FROM toast_demo;

SELECT *
FROM toast_validate_table('toast_demo', false);
```

The table validator may only be called by the relation owner or qualifying database owner. A clean result is diagnostic evidence, not a repair operation or corruption-proof guarantee.

### Crash and blocking boundary

Upstream warns that this experimental code may break, crash, or destroy a database. It directly depends on private heap, TOAST, tuple-visibility, snapshot, buffer, and index APIs from PostgreSQL 2018-era source. These layouts and APIs change across majors; never load a binary built for a different exact server.

The validator takes ExclusiveLock on the base table and AccessExclusiveLock on its TOAST table, scans physical rows with SnapshotAny, and reads every referenced chunk while checking transaction metadata. It can block reads or writes that need conflicting locks and perform heavy I/O on large or damaged tables. Use a replica or restored copy, set lock and statement timeouts, and validate one small relation at a time.

Pointer-detail functions expose physical relation and value OIDs and are executable by PUBLIC unless privileges are changed. Revoke all functions, grant only a specialist diagnostics role, preserve forensic copies before investigation, and use supported PostgreSQL recovery procedures rather than attempting manual TOAST modification. The repository is classified as abandoned and has not changed since 2018.
