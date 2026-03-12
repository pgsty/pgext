

## Usage

> [pgrowlocks: display row-level locking information](https://www.postgresql.org/docs/current/pgrowlocks.html)

pgrowlocks shows which rows in a table are currently locked, by which transactions, and the lock modes.

### Function

```sql
SELECT * FROM pgrowlocks('my_table');

 locked_row | locker | multi | xids  |     modes      |  pids
------------+--------+-------+-------+----------------+--------
 (0,1)      |    609 | f     | {609} | {"For Share"}  | {3161}
 (0,2)      |    609 | f     | {609} | {"For Share"}  | {3161}
 (0,3)      |    607 | f     | {607} | {"For Update"} | {3107}
```

### Return Columns

| Column | Type | Description |
|--------|------|-------------|
| `locked_row` | tid | Tuple ID of the locked row |
| `locker` | xid | Transaction ID (or multixact ID) |
| `multi` | boolean | True if locker is a multitransaction |
| `xids` | xid[] | Transaction IDs of all lockers |
| `modes` | text[] | Lock modes: `For Key Share`, `For Share`, `For No Key Update`, `For Update`, etc. |
| `pids` | integer[] | Process IDs of locking backends |

### View Locked Row Contents

```sql
SELECT * FROM accounts AS a, pgrowlocks('accounts') AS p
WHERE p.locked_row = a.ctid;
```

### Access

Restricted to superusers, roles with `pg_stat_scan_tables`, and users with `SELECT` on the target table.

### Caveats

- Takes `AccessShareLock` on the target table
- Not guaranteed to produce a self-consistent snapshot
- Can be slow on large tables
