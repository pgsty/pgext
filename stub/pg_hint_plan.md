## Usage

Sources:

- [README.md](https://github.com/ossc-db/pg_hint_plan/blob/4d3e9164dc2dc1b10dfe0c5e7a3cfe7983a8688d/README.md)
- [pg_hint_plan.control](https://github.com/ossc-db/pg_hint_plan/blob/4d3e9164dc2dc1b10dfe0c5e7a3cfe7983a8688d/pg_hint_plan.control)
- [docs/hint_list.md](https://github.com/ossc-db/pg_hint_plan/blob/4d3e9164dc2dc1b10dfe0c5e7a3cfe7983a8688d/docs/hint_list.md)
- [docs/installation.md](https://github.com/ossc-db/pg_hint_plan/blob/4d3e9164dc2dc1b10dfe0c5e7a3cfe7983a8688d/docs/installation.md)
- [docs/requirements.md](https://github.com/ossc-db/pg_hint_plan/blob/4d3e9164dc2dc1b10dfe0c5e7a3cfe7983a8688d/docs/requirements.md)
- [docs/functional_limitations.md](https://github.com/ossc-db/pg_hint_plan/blob/4d3e9164dc2dc1b10dfe0c5e7a3cfe7983a8688d/docs/functional_limitations.md)


`pg_hint_plan` overrides the PostgreSQL query planner's decisions using SQL comment hints, allowing you to force specific scan methods, join strategies, and join orders.

### Hint Syntax

Hints are embedded in SQL comments prefixed with `/*+` and closed with `*/`:

```sql
/*+
  HashJoin(a b)
  SeqScan(a)
*/
SELECT *
FROM pgbench_branches b
JOIN pgbench_accounts a ON b.bid = a.bid
ORDER BY a.aid;
```

### Scan Method Hints

| Hint | Description |
|------|-------------|
| `SeqScan(table)` | Force sequential scan |
| `IndexScan(table [index...])` | Force index scan (optionally restrict to specific indexes) |
| `IndexOnlyScan(table [index...])` | Force index-only scan |
| `BitmapScan(table [index...])` | Force bitmap scan |
| `TidScan(table)` | Force TID scan |
| `NoSeqScan(table)` | Prevent sequential scan |
| `NoIndexScan(table)` | Prevent index scan and index-only scan |
| `NoIndexOnlyScan(table)` | Prevent index-only scan |
| `NoBitmapScan(table)` | Prevent bitmap scan |
| `NoTidScan(table)` | Prevent TID scan |
| `IndexScanRegexp(table [regexp...])` | Force index scan with regex-matched index names |
| `DisableIndex(table index...)` | Disable specific indexes during planning |

### Join Method Hints

| Hint | Description |
|------|-------------|
| `NestLoop(t1 t2 [t3...])` | Force nested loop join |
| `HashJoin(t1 t2 [t3...])` | Force hash join |
| `MergeJoin(t1 t2 [t3...])` | Force merge join |
| `NoNestLoop(t1 t2 [t3...])` | Prevent nested loop join |
| `NoHashJoin(t1 t2 [t3...])` | Prevent hash join |
| `NoMergeJoin(t1 t2 [t3...])` | Prevent merge join |
| `Memoize(t1 t2 [t3...])` | Allow memoization of inner result |
| `NoMemoize(t1 t2 [t3...])` | Prevent memoization |

### Join Order Hints

```sql
-- Simple left-deep join order
/*+ Leading(a b c) */

-- Explicit join tree with nesting
/*+ Leading((a (b c))) */
```

### Row Number Correction

```sql
/*+ Rows(a b #100) */    -- Set to absolute number
/*+ Rows(a b +100) */    -- Add to estimate
/*+ Rows(a b -100) */    -- Subtract from estimate
/*+ Rows(a b *2.0) */    -- Multiply estimate
```

### Parallel Query Control

```sql
/*+ Parallel(table 4 hard) */   -- Force 4 parallel workers
/*+ Parallel(table 0 hard) */   -- Disable parallelism
```

### GUC Parameter Override

```sql
/*+ Set(random_page_cost 1.0) Set(seq_page_cost 1.0) */
SELECT * FROM my_table WHERE id = 42;
```

### GUC Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `pg_hint_plan.enable_hint` | Enable or disable hints globally | on |
| `pg_hint_plan.enable_hint_table` | Enable hint table for query-based hints | off |
| `pg_hint_plan.debug_print` | Print debug info for applied hints | off |
| `pg_hint_plan.parse_messages` | Log level for hint parsing messages | INFO |
| `pg_hint_plan.message_level` | Log level for debug messages | LOG |

### Loading and PostgreSQL Version Branches

Version 1.9.0 requires PostgreSQL 19. Use 1.8.1 for PG18, 1.7.2 for PG17, 1.6.3 for PG16, 1.5.4 for PG15 or 1.4.5 for PG14; do not install the PG19 build on an older server.

`LOAD 'pg_hint_plan'` activates hints in the current session without CREATE EXTENSION. Global preload requires a restart. To use the hint table, additionally run `CREATE EXTENSION pg_hint_plan` and enable `pg_hint_plan.enable_hint_table`. Hints cannot force a non-executable plan and query IDs ignore hint comments; verify actual plans with EXPLAIN. The current releases fix parser-loop, interrupt and planner error-path handling.
