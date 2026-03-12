


## Usage

> [pg_prewarm: prewarm relation data](https://www.postgresql.org/docs/current/pgprewarm.html)

The `pg_prewarm` extension provides functions to load relation data into the operating system buffer cache or PostgreSQL buffer cache, reducing I/O latency for subsequent queries.

### Prewarm a Relation

```sql
-- Load entire table into PostgreSQL buffer cache (default mode)
SELECT pg_prewarm('my_table');

-- Load with specific mode
SELECT pg_prewarm('my_table', 'prefetch');  -- async OS prefetch
SELECT pg_prewarm('my_table', 'read');      -- sync read into OS cache
SELECT pg_prewarm('my_table', 'buffer');    -- load into PG buffer cache

-- Load specific block range
SELECT pg_prewarm('my_table', 'buffer', 'main', 0, 999);

-- Prewarm an index
SELECT pg_prewarm('my_table_pkey');
```

### Function Signature

```sql
pg_prewarm(regclass,
           mode text DEFAULT 'buffer',
           fork text DEFAULT 'main',
           first_block int8 DEFAULT NULL,
           last_block int8 DEFAULT NULL
) RETURNS int8
```

Returns the number of blocks prewarmed.

| Parameter | Description |
|-----------|-------------|
| `mode` | `prefetch` (async OS), `read` (sync OS), or `buffer` (PG cache) |
| `fork` | Relation fork to prewarm (e.g., `main`, `fsm`, `vm`) |
| `first_block` | First block number (default: 0) |
| `last_block` | Last block number (default: last block of relation) |

### Autoprewarm

When loaded via `shared_preload_libraries`, autoprewarm periodically saves the list of buffers in the shared buffer cache to disk and restores them on restart.

```sql
-- Manually launch autoprewarm worker
SELECT autoprewarm_start_worker();

-- Force immediate dump of buffer state
SELECT autoprewarm_dump_now();  -- returns number of records written
```

### GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_prewarm.autoprewarm` | `true` | Enable autoprewarm background worker |
| `pg_prewarm.autoprewarm_interval` | `300s` | Interval between `autoprewarm.blocks` file updates (0 = dump only at shutdown) |

Buffer state is saved to `autoprewarm.blocks` in the data directory. After restart, two background workers reload the saved buffers.
