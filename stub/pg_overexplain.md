

## Usage

> [pg_overexplain: extended EXPLAIN with internal planner details](https://www.postgresql.org/docs/current/pgoverexplain.html)

pg_overexplain extends the `EXPLAIN` command with additional debugging options to display internal planner data structures. Primarily intended for planner debugging and development.

### Loading

```sql
LOAD 'pg_overexplain';
-- Or in postgresql.conf:
-- session_preload_libraries = 'pg_overexplain'
```

### EXPLAIN (DEBUG)

Displays internal plan tree information:

```sql
EXPLAIN (DEBUG) SELECT * FROM my_table WHERE id = 1;
```

Shows per-node fields:
- **Disabled Nodes** -- raw counter of disabled nodes
- **Parallel Safe** -- whether the node can appear under Gather
- **Plan Node ID** -- internal ID for parallel query coordination
- **extParam / allParam** -- parameters affecting the node

Shows per-query fields:
- **Command Type** -- query type (select, update, etc.)
- **Flags** -- hasReturning, hasModifyingCTE, canSetTag, transientPlan, etc.
- **Subplans Needing Rewind** -- subplan IDs requiring rewind
- **Relation OIDs** -- OIDs the plan depends on
- **Parse Location** -- location in the query string

### EXPLAIN (RANGE_TABLE)

Displays information about the query's range table entries:

```sql
EXPLAIN (RANGE_TABLE) SELECT * FROM t1 JOIN t2 ON t1.id = t2.id;
```

Shows range table index references (`Scan RTI`, `Nominal RTI`, `Append RTIs`, etc.) and dumps each range table entry with its kind (relation, subquery, join, cte, etc.) and entry-specific fields.

### Notes

- Output reflects internal planner data structures and may require reading source code to interpret
- Output format may change across PostgreSQL versions
- Not recommended for general production use
