## Usage

Sources:

- [Official database.dev package page](https://database.dev/dventimi/pg_partition_magician)

`dventimi@pg_partition_magician` — Pure-SQL online RANGE-partition manager (time / id / uuidv7), pg_cron-driven. Use it for the corresponding analytical or storage workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "dventimi@pg_partition_magician";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgpm._adopt` is an extension function.
- `pgpm._create_partition` is an extension function.
- `pgpm._decode` is an extension function.
- `pgpm._encode` is an extension function.
- `pgpm._frontier_native` is an extension function.
- `pgpm._grid_floor` is an extension function.
- `pgpm._grid_next` is an extension function.
- `pgpm._native_gt` is an extension function.
- `pgpm._native_type` is an extension function.
- `pgpm._part_name` is an extension function.
- `pgpm._ts_to_uuid` is an extension function.
- `pgpm._uuid_to_ts` is an extension function.
- `pgpm.adopt(p_parent regclass, p_control name, p_interval interval, p_premake int default 4, p_retention interval default null, p_keep_default boolean default true, p_drain_batch int default 5000, p_anchor timestamptz default '2000-01-01 00:00:00+00', p_paused boolean def…)` is an extension function and returns `regclass`.
- `pgpm.adopt_by_id(p_parent regclass, p_control name, p_step bigint, p_premake int default 4, p_retention bigint default null, p_keep_default boolean default true, p_drain_batch int default 5000, p_anchor bigint default 0, p_paused boolean default true, p_incoming_fks text defau…)` is an extension function and returns `regclass`.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- Install the confirmed extension dependencies first: `pg_cron`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
