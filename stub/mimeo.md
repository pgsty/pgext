

## Usage

> [mimeo: Extension for specialized, per-table replication between PostgreSQL instances](https://github.com/omniti-labs/mimeo)

Provides per-table replication between PostgreSQL instances with snapshot (full copy), incremental (timestamp/id based), and DML (insert/update/delete) modes.

### Enabling

```sql
CREATE SCHEMA mimeo;
CREATE EXTENSION mimeo SCHEMA mimeo;
```

Requires the `dblink` extension. Optionally install `pg_jobmon` for monitoring.

### Setting Up a Data Source

```sql
-- Create a dblink connection to the source database
SELECT mimeo.dblink_mapping_create(
    p_mapping_name := 'source_db',
    p_data_source := 'host=sourcehost dbname=sourcedb user=replicator password=secret',
    p_superuser := true
);
```

### Snapshot Replication (Full Table Copy)

Copies the entire source table each time it runs:

```sql
SELECT mimeo.snapshot_maker(
    p_src_table := 'public.my_table',
    p_dblink_id := 1  -- from dblink_mapping
);

-- Refresh the snapshot
SELECT mimeo.refresh_snap('public.my_table');
```

### Incremental Replication (Timestamp-Based)

Replicates rows based on an incrementing timestamp column:

```sql
SELECT mimeo.inserter_maker(
    p_src_table := 'public.events',
    p_control := 'created_at',  -- timestamp column
    p_dblink_id := 1
);

-- Refresh incrementally
SELECT mimeo.refresh_inserter('public.events');
```

For tables with updates (not just inserts):

```sql
SELECT mimeo.updater_maker(
    p_src_table := 'public.orders',
    p_control := 'updated_at',
    p_dblink_id := 1
);

SELECT mimeo.refresh_updater('public.orders');
```

### DML Replication (Insert/Update/Delete)

Full DML tracking via triggers on the source:

```sql
SELECT mimeo.dml_maker(
    p_src_table := 'public.accounts',
    p_dblink_id := 1
);

SELECT mimeo.refresh_dml('public.accounts');
```

### Scheduling Refreshes

Use `pg_jobmon` or cron to schedule periodic calls to the appropriate `refresh_*` function.

### Key Features

- Three replication modes: snapshot, incremental, DML
- Per-table replication (no need to replicate entire database)
- Works between different PostgreSQL versions
- Built on `dblink` for cross-database communication
