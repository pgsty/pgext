

## Usage

> [pg_partman: Extension to manage partitioned tables by time or ID](https://github.com/pgpartman/pg_partman)

`pg_partman` automates creation and management of both time-based and number-based partition sets
using PostgreSQL's native declarative partitioning (v5.0+). It handles adding new partitions and
removing old ones per retention policies, with an optional background worker for automatic maintenance.

### Create the Extension

```sql
CREATE SCHEMA partman;
CREATE EXTENSION pg_partman SCHEMA partman;
```

### Create a Time-Based Partition Set

```sql
CREATE TABLE public.measurements (
    id          bigserial,
    created_at  timestamptz NOT NULL DEFAULT now(),
    value       numeric
) PARTITION BY RANGE (created_at);

SELECT partman.create_parent(
    p_parent_table  := 'public.measurements',
    p_control       := 'created_at',
    p_interval      := '1 day'
);
```

### Create a Serial/ID-Based Partition Set

```sql
CREATE TABLE public.events (
    id      bigserial,
    data    text
) PARTITION BY RANGE (id);

SELECT partman.create_parent(
    p_parent_table  := 'public.events',
    p_control       := 'id',
    p_interval      := '100000'
);
```

### Run Maintenance

Manually trigger partition maintenance (create new partitions, drop expired ones):

```sql
SELECT partman.run_maintenance();
```

Or for a specific table:

```sql
SELECT partman.run_maintenance(p_parent_table := 'public.measurements');
```

### Configure Retention

Update the configuration to set retention policy:

```sql
UPDATE partman.part_config
SET    retention = '30 days',
       retention_keep_table = false
WHERE  parent_table = 'public.measurements';
```

### Background Worker

Enable automatic maintenance in `postgresql.conf`:

```
shared_preload_libraries = 'pg_partman_bgw'
pg_partman_bgw.interval = 3600          -- run every hour (seconds)
pg_partman_bgw.dbname = 'mydb'
```

### Migrate Existing Data into Partitions

```sql
CALL partman.partition_data_proc('public.measurements');
```

### Show Partitions

```sql
SELECT * FROM partman.show_partitions('public.measurements');
```

### Undo Partitioning

```sql
CALL partman.undo_partition_proc('public.measurements');
```
