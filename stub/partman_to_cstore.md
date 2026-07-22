## Usage

Sources:

- [Official README](https://github.com/dturon/partman_to_cstore/blob/80c6411318494837e5ef92825dd7a6c80b61fb66/README.md)
- [Extension control file](https://github.com/dturon/partman_to_cstore/blob/80c6411318494837e5ef92825dd7a6c80b61fb66/partman_to_cstore.control)
- [Partition migration function](https://github.com/dturon/partman_to_cstore/blob/80c6411318494837e5ef92825dd7a6c80b61fb66/sql/functions/part_to_cstore.sql)

`partman_to_cstore` is an experimental archival bridge for old `pg_partman` time-based partitions. It copies a heap partition into a `cstore_fdw` foreign table, reproduces inheritance, check constraints, ownership, and grants, then drops the original heap. It can later drop aged cstore partitions according to a second interval.

### Core Workflow

Install the two declared dependencies and place the non-relocatable extension in a dedicated schema:

```sql
CREATE EXTENSION cstore_fdw;
CREATE SCHEMA partman;
CREATE EXTENSION pg_partman WITH SCHEMA partman;
CREATE SCHEMA partman_to_cstore;
CREATE EXTENSION partman_to_cstore WITH SCHEMA partman_to_cstore;

INSERT INTO partman_to_cstore.move_config
    (parent_table, move_int, drop_int, compression)
VALUES
    ('archive.events', '7 days', '180 days', 'pglz');

SELECT partman_to_cstore.part_to_cstore(
    p_parent_table := 'archive.events',
    move_int := '7 days',
    drop_int := '180 days'
);
```

Run the function from a controlled scheduler only after verifying the candidate partitions. `move_singlepart_to_cstore(part_schema, part_table, ...)` performs the same destructive conversion for one inherited partition.

### Object Index

- `move_config` stores each parent, migration age, retention age, compression, stripe/block sizing, and last-check time.
- `part_to_cstore(...)` discovers eligible `pg_partman` children, creates cstore foreign tables, copies data, and removes the heap tables.
- `move_singlepart_to_cstore(...)` converts one inherited child.
- `partman_to_cstore_server` is the `cstore_fdw` foreign server created by the extension.
- `partman_to_cstore_move_data` is the supplied cron-oriented Python driver.

### Operational Notes

Version `1.1.0` requires `pg_partman` and `cstore_fdw`, is non-relocatable, and declares no preload requirement. Upstream documents support only time-based `pg_partman` partitions, including epoch-based naming. The SQL assumes old `pg_partman` catalog and function shapes, while the scheduler script uses Python 2 syntax; test compatibility before using it with current releases.

The conversion copies rows and then issues `DROP TABLE` for the original heap. The upstream warning states that `pg_dump` captures cstore foreign-table schema but not its data. Prove row counts, constraints, grants, query behavior, physical backups, and restore before scheduling migration; retain an independently restorable copy of archived data.
