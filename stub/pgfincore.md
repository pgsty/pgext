


## Usage

> [pgfincore: examine and manage the os buffer cache](https://github.com/klando/pgfincore)

pgfincore provides functions to inspect and manage OS page cache contents for PostgreSQL relations using `mincore` and `POSIX_FADVISE`.

### Inspect Cache State

```sql
SELECT * FROM pgfincore('pgbench_accounts');
```

Returns per-segment info: `relpath`, `segment`, `os_page_size`, `rel_os_pages`, `pages_mem` (pages in OS cache), `group_mem`, `os_pages_free`, `pages_dirty`, `group_dirty`.

Use `pgfincore('relation', true)` to include the `databit` varbit map for snapshot/restore.

### System Info

```sql
SELECT * FROM pgsysconf();          -- os_page_size, os_pages_free, os_total_pages
SELECT * FROM pgsysconf_pretty();   -- same with human-readable output
```

### Preload into OS Cache

```sql
SELECT * FROM pgfadvise_willneed('pgbench_accounts');
```

### Evict from OS Cache

```sql
SELECT * FROM pgfadvise_dontneed('pgbench_accounts');
```

### Other POSIX_FADVISE Flags

```sql
SELECT * FROM pgfadvise_normal('relation');
SELECT * FROM pgfadvise_sequential('relation');
SELECT * FROM pgfadvise_random('relation');
```

### Snapshot and Restore Cache State

```sql
-- Snapshot
CREATE TABLE pgfincore_snapshot AS
  SELECT 'pgbench_accounts'::text AS relname, *, now() AS date_snapshot
  FROM pgfincore('pgbench_accounts', true);

-- Restore
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, true,
               (SELECT databit FROM pgfincore_snapshot
                WHERE relname = 'pgbench_accounts' AND segment = 0));
```

### Direct Page Cache Control

```sql
-- Load first 3 pages, unload next 3
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, true, B'111000');
-- Load only
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, false, B'111000');
-- Unload only
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, false, true, B'111000');
```
