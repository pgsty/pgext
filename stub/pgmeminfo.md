

## Usage

> [pgmeminfo: PostgreSQL memory context information](https://github.com/okbob/pgmeminfo)

pgmeminfo provides functions to inspect PostgreSQL backend memory usage and memory context hierarchies.

### Functions

**Memory information overview:**

```sql
-- Show overall memory info
SELECT * FROM pgmeminfo();
```

**Memory context hierarchy:**

```sql
-- Show cumulative memory context sizes
SELECT * FROM pgmeminfo_contexts();

-- Show memory contexts to a specific depth
SELECT * FROM pgmeminfo_contexts(deep => 1);

-- Show all contexts without accumulation
SELECT * FROM pgmeminfo_contexts(deep => -1, accum_mode => 'off');
```
