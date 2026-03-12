

## Usage

> [pg_freespacemap: examine the free space map](https://www.postgresql.org/docs/current/pgfreespacemap.html)

pg_freespacemap provides functions to examine the free space map (FSM), which tracks available space on each page of a relation.

### Functions

**Free space on a single page:**

```sql
SELECT pg_freespace('my_table'::regclass, 0);  -- block 0
```

**Free space on all pages:**

```sql
SELECT * FROM pg_freespace('my_table'::regclass);

 blkno | avail
-------+-------
     0 |     0
     1 |     0
     2 |   224
     3 |  3456
     4 |  8160
```

### Example: Table Bloat Analysis

```sql
-- Pages with significant free space
SELECT blkno, avail
FROM pg_freespace('my_table'::regclass)
WHERE avail > 1000
ORDER BY avail DESC;

-- Total free space in a relation
SELECT sum(avail) AS total_free_bytes,
       count(*) AS total_pages,
       count(*) FILTER (WHERE avail > 0) AS pages_with_free_space
FROM pg_freespace('my_table'::regclass);
```

### Notes

- FSM values are rounded to 1/256th of `BLCKSZ` (typically 32 bytes)
- FSM is not kept fully up-to-date; values may lag behind actual free space
- For indexes, only entirely unused pages are tracked
- Access restricted to superusers and `pg_stat_scan_tables` members
