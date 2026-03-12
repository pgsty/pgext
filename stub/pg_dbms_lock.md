

## Usage

> [pg_dbms_lock: Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL](https://github.com/HexaCluster/pg_dbms_lock)

Uses PostgreSQL advisory locks to emulate Oracle DBMS_LOCK behavior.

### Enabling

```sql
CREATE EXTENSION pg_dbms_lock;
```

### ALLOCATE_UNIQUE

Allocate a unique lock identifier for a named lock:

```sql
DO $$
DECLARE
    lockhandle varchar;
BEGIN
    CALL dbms_lock.allocate_unique(
        lockname => 'printer_lock',
        lockhandle => lockhandle
    );
    RAISE NOTICE 'Handle: %', lockhandle;
END;
$$;
```

### REQUEST

Request a lock with a specific mode (Exclusive=6, Shared=4):

```sql
DO $$
DECLARE
    lock_res int;
BEGIN
    lock_res := dbms_lock.request(
        id => 123,
        lockmode => 6,           -- Exclusive
        timeout => 300,
        release_on_commit => false
    );
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'Lock request failed: %', lock_res;
    END IF;
END;
$$;
```

Return values: 0=Success, 1=Timeout, 3=Parameter error, 4=Already own lock, 5=Illegal handle.

### RELEASE

Explicitly release a previously acquired lock:

```sql
DO $$
DECLARE
    lock_res int;
BEGIN
    lock_res := dbms_lock.release(id => 123);
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'Release failed: %', lock_res;
    END IF;
END;
$$;
```

### SLEEP

Suspend the session for a given duration:

```sql
CALL dbms_lock.sleep(0.70);  -- sleep 0.7 seconds
```

### Complete Example

```sql
DO $$
DECLARE
    lock_res int;
    printer_lockhandle varchar;
BEGIN
    CALL dbms_lock.allocate_unique('printer_lock', printer_lockhandle);
    lock_res := dbms_lock.request(lockhandle => printer_lockhandle, lockmode => 6, timeout => 5);
    IF lock_res <> 0 THEN
        RAISE EXCEPTION 'REQUEST failed: %', lock_res;
    END IF;
    -- do exclusive work here
    lock_res := dbms_lock.release(lockhandle => printer_lockhandle);
END;
$$;
```
