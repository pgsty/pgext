## Usage

Sources:

- [Official README](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/README.md)
- [Official extension SQL](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/sql/pgparts.sql)
- [Official extension control file](https://github.com/dvarrazzo/pgparts/blob/58c5d12ce5530323ab022e17ca317c5db4e1089c/pgparts.control)

`pgparts` 0.0.15 is a PL/pgSQL toolkit for legacy inheritance-based table partitioning. It copies a base table's constraints, indexes, triggers, owner, and permissions into child tables, then maintains routing triggers and metadata. Use it for databases that intentionally retain this design, not as a substitute for PostgreSQL's current declarative partitioning.

### Core Workflow

Install the non-relocatable extension into a dedicated schema, prepare a base table, create the partition that covers a value, and then insert data:

```sql
CREATE SCHEMA parts;
CREATE EXTENSION pgparts WITH SCHEMA parts;

CREATE TABLE events (
  id bigint PRIMARY KEY,
  created_at timestamptz NOT NULL,
  payload jsonb
);

SELECT parts.setup(
  'events'::regclass,
  'created_at',
  'monthly'
);

SELECT parts.create_for(
  'events'::regclass,
  '2026-07-22 12:00:00+00'
);

INSERT INTO events VALUES (
  1,
  '2026-07-22 12:00:00+00',
  '{"event":"created"}'::jsonb
);

SELECT * FROM parts.existing_partition;
```

`parts.setup` records the partition key and strategy and builds the insert-routing function and trigger. `parts.create_for` is idempotent for an already-present range and returns the child table as `regclass`.

### Important Objects

- `setup(regclass, name, name, params)` prepares a table; the fourth argument defaults to an empty `params` value.
- `create_for(regclass, text)` creates the partition covering an example key value.
- `partition_for(regclass, text)`, `info(regclass, text)`, `start_for(regclass, text)`, and `end_for(regclass, text)` inspect routing and boundaries.
- `detach_for(regclass, text)` and `attach_for(regclass, text)` change inheritance without dropping the child.
- `create_archive(regclass)`, `archive_before(regclass, timestamptz)`, `archive_partition(regclass)`, and `unarchive_partition(regclass)` manage archived children.
- `partitioned_table`, `partition`, and `existing_partition` expose configuration and live child metadata.

### Strategies and Parameters

The built-in `monthly` and `daily` strategies support date, timestamp, timestamp with time zone, and related configured types. Parameters include `nmonths` or `ndays`, boundary `timezone`, `drop_old`, `on_conflict_drop`, and informational `retention`; `start_dow` applies to daily/weekly grouping. Inspect `parts.partition_schema` and `parts.schema_param` for the exact accepted combinations and descriptions.

### Operational Notes

Routing uses generated triggers and table inheritance. The base table is a blueprint, but later DDL changes are not automatically guaranteed to propagate to every existing child; manage schema changes deliberately. Trigger routing checks partitions from newest to oldest, so random historical inserts become slower as child count grows. Archive operations change inheritance rather than moving rows, and `retention` is not enforced by the extension. Test uniqueness, foreign keys, update routing, backup/restore, permissions, DDL propagation, and concurrent partition creation on the target PostgreSQL version.
