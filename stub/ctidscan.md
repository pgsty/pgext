## Usage

Sources:

- [Official README](https://github.com/XiaYingyin/ctidscan/blob/101a177a42ee94c0aac5bffa8ca63ae9f0790d5d/README.md)
- [Module control file](https://github.com/XiaYingyin/ctidscan/blob/101a177a42ee94c0aac5bffa8ca63ae9f0790d5d/ctidscan.control)
- [Custom-scan implementation](https://github.com/XiaYingyin/ctidscan/blob/101a177a42ee94c0aac5bffa8ca63ae9f0790d5d/ctidscan.c)

`ctidscan` is a headless custom-scan demo that adds a planner path for heap scans constrained by `ctid` inequalities or ranges. It has no extension SQL and should be loaded only in sessions used to study physical tuple-range scans.

### Core Workflow

```sql
LOAD 'ctidscan';

CREATE TABLE ctid_demo AS
SELECT g AS id, repeat('x', 100) AS payload
FROM generate_series(1, 10000) AS g;

SET enable_ctidscan = on;

EXPLAIN (COSTS OFF)
SELECT ctid, id
FROM ctid_demo
WHERE ctid BETWEEN '(2,1)'::tid AND '(3,20)'::tid;
```

The planner extracts lower and upper `ctid` bounds and starts or stops the heap scan around the corresponding blocks. Compare plans and results with `SET enable_ctidscan = off` when evaluating the custom path.

### Objects and Configuration

- `LOAD 'ctidscan'` initializes the module for the current backend.
- `enable_ctidscan` enables or disables its planner path.
- The module emits informational start/end messages while executing the custom scan.

The README also shows `shared_preload_libraries = 'ctidscan'`, but recommends local loading so the prototype does not affect unrelated users. A server restart is required after changing `shared_preload_libraries`.

### Safety and Compatibility

A `ctid` is a physical tuple address, not a durable row identifier. It can change after `UPDATE`, `VACUUM FULL`, `CLUSTER`, or any table rewrite; never store it as an application key or use an old value to authorize or mutate a row.

The implementation copies executor structures and uses heap APIs from the PostgreSQL 9.x era. It has no current compatibility matrix and can break across major versions. Treat `ctidscan` as planner-extension sample code, validate correctness against a normal scan, and do not preload it on production clusters.
