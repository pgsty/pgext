## Usage

Sources:

- [Official usage documentation](https://github.com/yazun/range_partitioning/blob/34853779f37b25a253b3c6889b6aaafe2ab56968/doc/range_partitioning.md)
- [Official README](https://github.com/yazun/range_partitioning/blob/34853779f37b25a253b3c6889b6aaafe2ab56968/README.md)
- [PGXN distribution metadata](https://pgxn.org/dist/range_partitioning/)

`range_partitioning` implements trigger-based table partitioning with PostgreSQL range types. It converts an ordinary table into a managed parent, routes inserted rows to child tables, and can split or merge adjacent ranges. It predates PostgreSQL's declarative partitioning and should be treated as a legacy alternative, not mixed casually with native `PARTITION BY` tables.

### Core Workflow

Install the relocatable extension in a dedicated schema, grant its management role only to partition administrators, and convert an existing table by naming its partition key:

```sql
CREATE SCHEMA range_partitioning;
CREATE EXTENSION range_partitioning SCHEMA range_partitioning;

CREATE SCHEMA app;
CREATE TABLE app.events (
  event_id bigint,
  occurred_on date NOT NULL,
  payload jsonb
);

SELECT range_partitioning.create_parent(
  'app.events',
  'occurred_on'
);

SELECT range_partitioning.create_partition(
  'app.events',
  '[2026-01-01,2027-01-01)'
);
```

`create_parent` creates metadata, an initial unbounded child, constraints, and the insert-routing trigger. If more than one compatible range type exists, pass the qualified range type explicitly. `create_partition` requires the new range to be a complete subset of one existing partition and splits that partition.

### Management and Inspection

- `drop_partition` merges a child into a named adjacent partition; it is not simply a detach operation.
- `master` and `partition` hold management metadata, while `master_partition` exposes a joined reporting view.
- `get_destination_partition` identifies which child would receive a candidate key.
- `where_clause` derives a predicate for a range or an existing partition.
- `value_in_range`, `is_subrange`, `range_add`, `range_subtract`, and `constructor_clause` manipulate dynamically specified range values.

Inspect the managed layout before and after every split or merge:

```sql
SELECT master_class::regclass AS parent,
       partition_class::regclass AS child,
       range,
       where_clause(partition_class) AS predicate
FROM range_partitioning.partition
ORDER BY 1, 2;
```

### Privileges and Caveats

The upstream README says installation creates a role named `range_partitioning`; grant that role only to users who must create or remove partitions. Child creation carries table structure and permissions according to the extension's legacy implementation, so verify indexes, grants, triggers, ownership, and dump/restore behavior on representative tables.

Version 1.2.2 requires PostgreSQL 9.2 or later and was released in 2016. The control file declares no preload requirement or extension dependency. Modern PostgreSQL offers declarative range partitioning with planner and maintenance integration; compare it before adopting this trigger-and-metadata design. Partition changes move or rewrite data and acquire locks, so size and rehearse them. Use the extension's upgrade scripts with `ALTER EXTENSION`, and do not manually edit its metadata tables.
