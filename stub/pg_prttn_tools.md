## Usage

Sources:

- [Official extension control file](https://github.com/alexandersamoylov/pg_prttn_tools/blob/c7a034b56cd8c2327779dbc0d8060fd2220eb010/pg_prttn_tools.control)
- [Official version 0.5 SQL](https://github.com/alexandersamoylov/pg_prttn_tools/blob/c7a034b56cd8c2327779dbc0d8060fd2220eb010/sql/pg_prttn_tools--0.5.sql)

`pg_prttn_tools` manages old-style PostgreSQL inheritance partitions with generated check constraints, rules, and insert triggers. It predates declarative partitioning and should be used only when an existing inheritance-based design requires its exact behavior.

### Core Workflow

The time helpers accept `year`, `month`, `day`, `hour`, or `minute`, derive the child-table name and bounds, then optionally create an insert-routing rule.

```sql
CREATE EXTENSION pg_prttn_tools;

CREATE TABLE public.events (
  id bigint,
  occurred_at timestamp without time zone NOT NULL,
  payload jsonb
);

SELECT *
FROM prttn_tools.part_time_add(
  'public',
  'events',
  'occurred_at',
  'month',
  '2026-07-10 12:00:00'::timestamp,
  true
);
```

### Function Families

- `create_child_table` clones a parent with `LIKE ... INCLUDING ALL`, adds inheritance, a check constraint, and optionally a routing rule.
- `part_list_check`, `part_list_add`, and `part_list_create_trigger` handle list-valued partitions.
- `part_time_check`, `part_time_add`, and `part_time_create_trigger` handle timestamp ranges.
- `part_list_time_check`, `part_list_time_add`, and `part_list_time_create_trigger` combine list and time dimensions.
- `drop_ins_trigger`, `part_merge`, and `part_time_cleanup` remove routing objects, merge a child back into its parent, or drop old time-named tables.

### Privileges and Destructive Operations

The control file requires a superuser, fixes the extension schema as `prttn_tools`, and the SQL changes function ownership to `postgres`. The functions build DDL by concatenating their text arguments, so pass only trusted identifiers and values.

`part_time_cleanup` executes `DROP TABLE` for every table in a schema whose name starts with the supplied prefix and sorts before the computed timestamp suffix. Confirm the exact prefix and inspect candidate tables before calling it. These helpers create inheritance children, not native `PARTITION BY` partitions; planner pruning, constraints, indexes, and routing behavior therefore follow the older inheritance model.
