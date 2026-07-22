## Usage

Sources:

- [Official extension control file](https://github.com/disqus/pg_partition/blob/6ace26cfc6ffb2fdb6bb2163762a847b553d7691/pg_partition.control)
- [Official extension SQL](https://github.com/disqus/pg_partition/blob/6ace26cfc6ffb2fdb6bb2163762a847b553d7691/sql/pg_partition--0.2.0.sql)
- [Official usage document](https://github.com/disqus/pg_partition/blob/6ace26cfc6ffb2fdb6bb2163762a847b553d7691/doc/pg_partition.md)

`pg_partition` 0.2.0 generates SQL for pre-declarative, inheritance-based time partitions. Given a parent table, a `timestamptz` column, a UTC range, and a fixed interval, it emits child-table DDL, statements that move existing rows, one index per child, and either trigger- or rule-based insert routing. It does not create native declarative partitions or run ongoing partition maintenance.

### Core Workflow

The generator refuses to run unless the session time zone is UTC, the partition key is `timestamp with time zone`, and the start is at UTC midnight. Generate the script, review it, and execute it separately in a controlled transaction.

```sql
SET TIME ZONE 'UTC';
CREATE EXTENSION pg_partition;

CREATE TABLE public.events (
    event_id bigint GENERATED ALWAYS AS IDENTITY,
    happened_at timestamptz NOT NULL,
    payload jsonb
);

SELECT create_partitions_trigger_when(
    'public',
    'events',
    'happened_at',
    '2026-01-01 00:00:00+00'::timestamptz,
    '2026-02-01 00:00:00+00'::timestamptz,
    interval '1 day'
);
```

Capture the returned rows without headings or formatting, inspect their order and identifiers, then execute them once. `create_partitions_rule` produces equivalent child tables and row moves but routes inserts with rules instead of per-range triggers.

### Generator Functions

- `validate_inputs` enforces UTC, the key type, midnight start, and an allowed interval.
- `unroll_partition_spec` expands the requested half-open time range into child specifications and suffixes.
- `create_partitions_trigger_when` emits child tables, data moves, indexes, trigger functions, and triggers.
- `create_partitions_rule` emits child tables, data moves, indexes, and insert rules.
- The lower-level `create_partition_table`, `move_to_partition`, `create_partition_index`, `create_trigger_function`, `create_trigger`, and `create_rule` functions each return one piece of DDL as text.

Supported intervals are exactly one year, six months, three months, one month, one week, one day, one hour, one minute, or one second.

### Operational Boundaries

Generated child tables use `INHERITS` and `CHECK` constraints. They are not attached through PostgreSQL's declarative `PARTITION BY` mechanism, so native partition pruning, default partitions, partition attach/detach, and modern partition-management tooling do not apply in the same way.

The generated script deletes matching rows from only the parent and inserts them into each child. Run it during a controlled write window or design explicit locking, because concurrent inserts and failures can leave rows in unexpected locations. Rows outside the generated ranges remain in the parent; the extension does not create future partitions automatically.

The project source is from 2014 and its short usage document refers to a nonexistent `create_partitions` wrapper; use the two functions named above from the actual 0.2.0 SQL. Test generated DDL, query plans, trigger search paths, backup/restore, and migration to native partitioning on the exact supported PostgreSQL release.
