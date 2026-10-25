## Usage

Sources:

- [Official README](https://github.com/ibrohim238/postgresql-leapfrog/blob/88fe4e69f578db196613ce68d12349cf0042947a/README.md)
- [Control file](https://github.com/ibrohim238/postgresql-leapfrog/blob/88fe4e69f578db196613ce68d12349cf0042947a/leapfrog.control)
- [Extension SQL](https://github.com/ibrohim238/postgresql-leapfrog/blob/88fe4e69f578db196613ce68d12349cf0042947a/leapfrog--0.0.3.sql)

`leapfrog` 0.0.3 is an experimental PostgreSQL 18 planner extension. It combines ordinary B-tree indexes into an ordered intersection, allowing suitable filtered queries with ordering and a limit to use fewer specialized indexes. Other PostgreSQL majors are unsupported.

### Enable and Inspect

Install the SQL helpers as a superuser and load the library in the session that plans queries:

```sql
CREATE EXTENSION leapfrog;
LOAD 'leapfrog';
SELECT leapfrog_version(), leapfrog_is_loaded();

CREATE TABLE orders (
    id bigint PRIMARY KEY,
    customer_id bigint NOT NULL,
    product_id bigint NOT NULL
);
CREATE INDEX ON orders (customer_id, id);
CREATE INDEX ON orders (product_id, id);
ANALYZE orders;
EXPLAIN
SELECT * FROM orders
WHERE customer_id = 100 AND product_id = 200
ORDER BY id DESC LIMIT 100;
```

The planner may choose `Custom Scan (Index Merge Intersect)` when its estimated cost is favorable. An empty demonstration table normally keeps a native plan. The outer index ends with the sort column; leading index columns must have equality conditions. Each additional index must contribute an independent equality or partial-index predicate.

### Loading and Settings

`CREATE EXTENSION` registers the SQL helpers; every subsequent session still needs to load the library. `LOAD 'leapfrog'` applies to the current session. An administrator can configure `session_preload_libraries` for new sessions, or `shared_preload_libraries` for all backends, with a server restart for the latter.

| Setting | Default | Purpose |
| --- | --- | --- |
| `leapfrog.enabled` | on | Offer the custom path |
| `leapfrog.force` | off | Prefer eligible custom paths for diagnostics |
| `leapfrog.max_indexes` | 4 | Maximum streams per intersection, from 2 to 16 |
| `leapfrog.min_indexes` | 2 | Developer control for minimum stream count |
| `leapfrog.skip_after` | 64 | Initial skip-ahead threshold; 0 disables skipping |

`leapfrog.force` can select slower plans. Use actual query plans and timings before changing defaults. The normal planner retains its other paths.

### Limits and Maintenance

Sort columns must be non-null integers: `smallint`, `integer` or `bigint`. Index columns must use ascending order and nbtree version 4; descending query order uses a backward scan. Expression index columns and ranges or lists on leading equality columns are unsupported. Indexes carried from older PostgreSQL versions may need rebuilding to meet the nbtree requirement.

Duplicate-key TID groups are bounded by `work_mem`; an overflowing group falls back to row-condition checks. Open scans can hold index pages pinned and delay vacuum. This extension creates no application tables and does not replace existing indexes.

`DROP EXTENSION` does not unload a library already loaded into a backend. Disable its planner contribution with `SET leapfrog.enabled = off`, and reconnect sessions after replacing the library. Use a proper installation or package replacement; overwriting a loaded shared-library file in place can crash its users.
