


## Usage

> [pg_cheat_funcs: Provides cheat (but useful) functions](https://github.com/MasaoFujii/pg_cheat_funcs)

The `pg_cheat_funcs` extension provides a collection of utility functions for debugging, diagnostics, and low-level PostgreSQL operations. Many are superuser-restricted.

### Process Control

```sql
SELECT pg_signal_process(12345, 'TERM');       -- send signal to a PG process
SELECT pg_get_priority(pg_backend_pid());      -- get scheduling priority
SELECT pg_set_priority(pg_backend_pid(), 10);  -- set scheduling priority (-20..19)
SELECT pg_postmaster_pid();                    -- get postmaster PID
SELECT pg_backend_start_time();                -- server process start time
```

### Memory Context Inspection

```sql
-- Show memory context statistics (PG 9.6-13)
SELECT * FROM pg_stat_get_memory_context();
-- Columns: name, parent, level, total_bytes, total_nblocks, free_bytes, free_chunks, used_bytes
```

### Prepared Statement Inspection

```sql
-- Show cached plan info for a prepared statement
SELECT * FROM pg_cached_plan_source('my_stmt');
-- Columns: generic_cost, total_custom_cost, num_custom_plans, force_generic, force_custom
```

### Transaction & WAL Functions

```sql
SELECT pg_xlogfile_name('0/1234568'::pg_lsn, false);  -- LSN to WAL filename
SELECT pg_wait_syncrep('0/1234568'::pg_lsn);           -- wait for sync rep
SELECT * FROM pg_stat_get_syncrep_waiters();            -- list sync rep waiters
SELECT pg_set_next_xid('100'::xid);                    -- set next transaction ID (dangerous)
SELECT * FROM pg_xid_assignment();                      -- XID state info
```

### Checkpoint & Recovery

```sql
SELECT pg_checkpoint(true, true, true);  -- fast, wait, force
SELECT pg_promote(true);                 -- promote standby (PG <= 11)
SELECT * FROM pg_recovery_settings();    -- show recovery.conf parameters
SELECT pg_show_primary_conninfo();       -- show primary_conninfo
```

### File Operations

```sql
SELECT * FROM pg_list_relation_filepath('my_table'::regclass);  -- list segment files
SELECT pg_file_write_binary('/tmp/test', '\x48656c6c6f'::bytea); -- write binary file
SELECT pg_file_fsync('/tmp/test');                                -- fsync file
```

### Text & Encoding Conversion

```sql
SELECT to_octal(255);                   -- '377'
SELECT pg_text_to_hex('PostgreSQL');     -- '506f737467726553514c'
SELECT pg_hex_to_text('506f737467726553514c'); -- 'PostgreSQL'
SELECT pg_chr(9731);                     -- snowman character
```

### Compression

```sql
SELECT pglz_compress('some text data');        -- PGLZ compress text to bytea
SELECT pglz_decompress(compressed_data);       -- decompress back to text
SELECT pglz_compress_bytea(data);              -- compress bytea
SELECT pglz_decompress_bytea(compressed_data); -- decompress to bytea
```

### Advisory Lock Management

```sql
SELECT pg_advisory_xact_unlock(12345);              -- release exclusive advisory lock
SELECT pg_advisory_xact_unlock_shared(12345);       -- release shared advisory lock
```

### GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_cheat_funcs.log_memory_context` | `off` | Log memory context stats after query execution |
| `pg_cheat_funcs.hide_appname` | `false` | Hide client application_name |
| `pg_cheat_funcs.log_session_start_options` | `off` | Log connection startup options |
| `pg_cheat_funcs.scheduling_priority` | `0` | Process scheduling priority (-20..19) |
| `pg_cheat_funcs.exit_on_segv` | `off` | If on, segfault terminates session only |
