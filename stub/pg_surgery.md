


## Usage

> [pg_surgery: extension to perform surgery on a damaged relation](https://www.postgresql.org/docs/current/pgsurgery.html)

The `pg_surgery` extension provides functions to perform low-level surgery on damaged relations. These are last-resort tools that can corrupt data if misused.

### Functions

#### heap_force_kill

Marks line pointers as "dead" without examining the tuples, forcibly removing inaccessible tuples.

```sql
heap_force_kill(regclass, tid[]) RETURNS void
```

```sql
-- Tuple causes error when accessed
SELECT * FROM t1 WHERE ctid = '(0, 1)';
-- ERROR: could not access status of transaction 4007513275

-- Force kill the problematic tuple
SELECT heap_force_kill('t1'::regclass, ARRAY['(0, 1)']::tid[]);

-- Tuple is now removed
SELECT * FROM t1 WHERE ctid = '(0, 1)';
-- (0 rows)
```

#### heap_force_freeze

Marks tuples as frozen without examining tuple data, making tuples accessible when visibility information is corrupted.

```sql
heap_force_freeze(regclass, tid[]) RETURNS void
```

```sql
-- VACUUM fails on corrupted visibility info
VACUUM t1;
-- ERROR: found xmin 507 from before relfrozenxid 515

-- Find the problematic tuple
SELECT ctid FROM t1 WHERE xmin = 507;
--  ctid
-- -------
--  (0,3)

-- Force freeze the tuple
SELECT heap_force_freeze('t1'::regclass, ARRAY['(0, 3)']::tid[]);

-- Tuple is now frozen (xmin becomes 2 = FrozenTransactionId)
SELECT ctid FROM t1 WHERE xmin = 2;
--  ctid
-- -------
--  (0,3)
```

### When to Use

- `heap_force_kill`: When tuples cause errors on access due to corrupted transaction status, and the data can be discarded
- `heap_force_freeze`: When VACUUM fails due to tuples with xmin before relfrozenxid, or when tuples are invisible due to corrupted visibility information

These functions are unsafe by design and should only be used as a last resort when normal recovery methods have failed.
