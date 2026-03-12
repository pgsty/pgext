


## Usage

> [pg_orphaned: Deal with orphaned files](https://github.com/bdrouvot/pg_orphaned)

pg_orphaned provides functions to detect and manage orphaned data files in PostgreSQL. It handles corner cases like in-progress transactions that could cause false positives by using a dirty snapshot.

### List Orphaned Files

```sql
-- List orphaned files (default: older than 1 day marked as "older")
SELECT * FROM pg_list_orphaned();

-- Custom age threshold
SELECT * FROM pg_list_orphaned('10 seconds');
SELECT * FROM pg_list_orphaned('1 minute');
```

Returns: `dbname`, `path`, `name`, `size`, `mod_time`, `relfilenode`, `reloid`, `older` (boolean).

### Move Orphaned Files to Backup

```sql
-- Move files older than the threshold to orphaned_backup directory
SELECT pg_move_orphaned('1 minute');
```

### List Moved Files

```sql
SELECT * FROM pg_list_orphaned_moved();
```

### Move Files Back (if still orphaned)

```sql
SELECT pg_move_back_orphaned();
```

### Remove Moved Files

```sql
SELECT pg_remove_moved_orphaned();
```

### Typical Workflow

```sql
-- 1. Check for orphaned files
SELECT * FROM pg_list_orphaned('1 minute');

-- 2. Move them to backup (only those older than threshold)
SELECT pg_move_orphaned('1 minute');

-- 3. Verify what was moved
SELECT * FROM pg_list_orphaned_moved();

-- 4. After confirming, remove the backup files
SELECT pg_remove_moved_orphaned();
```

Note: functions operate on orphaned files for the database you are connected to. Always double-check carefully before moving or removing files.
