


## Usage

> [pgcozy: Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL.](https://github.com/vventirozos/pgcozy)

pgcozy snapshots the current shared buffer state and can later restore (pre-warm) buffers from those snapshots. Requires `pg_buffercache` and `pg_prewarm` extensions.

### Initialize

```sql
SELECT pgcozy_init();
```

Creates the `pgcozy` schema with a `snapshots` table and the `cozy_type` type.

### Take a Snapshot

```sql
-- Snapshot buffer pages with usagecount >= 3 (popularity 1-5)
SELECT pgcozy_snapshot(3);

-- Snapshot all buffered pages (popularity = 0)
SELECT pgcozy_snapshot(0);
```

Snapshots are stored as JSONB in `pgcozy.snapshots` with columns: `id`, `snapshot_date`, `snapshot`. Each entry contains `table_name`, `block_no`, and `popularity`.

### Warm from a Snapshot

```sql
-- Warm from a specific snapshot by ID
SELECT pgcozy_warm(1);

-- Warm from the latest snapshot
SELECT pgcozy_warm(0);
```

### Review Snapshots

```sql
SELECT id, snapshot_date FROM pgcozy.snapshots;
```

Snapshots are stored in JSONB and can be reviewed, backed up, or transferred between servers.
