


## Usage

> [amcheck: functions for verifying relation integrity](https://www.postgresql.org/docs/current/amcheck.html)

The `amcheck` extension provides functions to verify the logical consistency of B-tree indexes, GIN indexes, and heap (table) data, detecting corruption without modifying data.

### B-Tree Index Verification

```sql
-- Lightweight check (AccessShareLock, safe for production)
SELECT bt_index_check('my_index');

-- With heap-all-indexed verification
SELECT bt_index_check('my_index', heapallindexed => true);

-- Thorough check including parent/child invariants (ShareLock, blocks writes)
SELECT bt_index_parent_check('my_index');

-- Most thorough: rootdescend re-finds each tuple from root
SELECT bt_index_parent_check('my_index',
    heapallindexed => true,
    rootdescend => true,
    checkunique => true);
```

### Check All Catalog Indexes

```sql
SELECT bt_index_check(c.oid), c.relname, c.relpages
FROM pg_index i
JOIN pg_opclass op ON i.indclass[0] = op.oid
JOIN pg_am am ON op.opcmethod = am.oid
JOIN pg_class c ON i.indexrelid = c.oid
JOIN pg_namespace n ON c.relnamespace = n.oid
WHERE am.amname = 'btree' AND n.nspname = 'pg_catalog'
  AND c.relpersistence != 't' AND c.relkind = 'i'
  AND i.indisready AND i.indisvalid
ORDER BY c.relpages DESC LIMIT 10;
```

### GIN Index Verification

```sql
SELECT gin_index_check('my_gin_index');
```

### Heap (Table) Verification

```sql
-- Basic heap check
SELECT * FROM verify_heapam('my_table');

-- With TOAST verification (slower)
SELECT * FROM verify_heapam('my_table', check_toast => true);

-- Check specific block range
SELECT * FROM verify_heapam('my_table', startblock => 0, endblock => 1000);

-- Stop at first corrupted block
SELECT * FROM verify_heapam('my_table', on_error_stop => true);
```

Returns rows for each detected problem:

| Column | Type | Description |
|--------|------|-------------|
| `blkno` | bigint | Block number with corruption |
| `offnum` | integer | Offset of corrupt tuple |
| `attnum` | integer | Attribute number (if column-specific) |
| `msg` | text | Description of the problem |

### Function Summary

| Function | Lock | Use Case |
|----------|------|----------|
| `bt_index_check(index, heapallindexed, checkunique)` | AccessShareLock | Routine production checks |
| `bt_index_parent_check(index, heapallindexed, rootdescend, checkunique)` | ShareLock | Comprehensive verification |
| `gin_index_check(index)` | AccessShareLock | GIN index verification |
| `verify_heapam(relation, on_error_stop, check_toast, skip, startblock, endblock)` | AccessShareLock | Table/heap corruption detection |

All `amcheck` errors are true positives. Use `REINDEX` or point-in-time recovery for repair after diagnosing with `pageinspect`.
