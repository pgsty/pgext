

## Usage

> [pg_visibility: examine the visibility map](https://www.postgresql.org/docs/current/pgvisibility.html)

pg_visibility provides functions to examine and verify the visibility map (VM), which tracks which pages contain only tuples visible to all transactions.

### Functions

**Single page visibility:**

```sql
-- VM bits for a specific block
SELECT * FROM pg_visibility_map('my_table', 0);
-- Returns: all_visible, all_frozen

-- VM bits plus the page's PD_ALL_VISIBLE flag
SELECT * FROM pg_visibility('my_table', 0);
-- Returns: all_visible, all_frozen, pd_all_visible
```

**All pages visibility:**

```sql
-- VM bits for every page
SELECT * FROM pg_visibility_map('my_table');
-- Returns: blkno, all_visible, all_frozen

-- VM bits plus PD_ALL_VISIBLE for every page
SELECT * FROM pg_visibility('my_table');
-- Returns: blkno, all_visible, all_frozen, pd_all_visible
```

**Summary:**

```sql
SELECT * FROM pg_visibility_map_summary('my_table');
-- Returns: all_visible (count), all_frozen (count)
```

### Corruption Detection

```sql
-- Find tuples on all-frozen pages that aren't actually frozen
SELECT * FROM pg_check_frozen('my_table');

-- Find tuples on all-visible pages that aren't actually all-visible
SELECT * FROM pg_check_visible('my_table');
```

If either function returns rows, the visibility map is corrupt.

### Repair

```sql
-- Truncate the visibility map (forces full VACUUM rebuild)
SELECT pg_truncate_visibility_map('my_table');
-- Then run: VACUUM my_table;
```

### Access

Functions require superuser or `pg_stat_scan_tables` role. `pg_truncate_visibility_map` requires superuser.
