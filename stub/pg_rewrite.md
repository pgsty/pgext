


## Usage

> [pg_rewrite: Tool allows read write to the table during the rewriting](https://github.com/cybertec-postgresql/pg_rewrite)

`pg_rewrite` requires `wal_level = logical` and must be added to `shared_preload_libraries`. Common use cases include changing column data types, partitioning tables, reordering columns, and moving tables to different tablespaces -- all while allowing concurrent reads and writes.

### Rewrite a Table

Create the target table with the desired schema, then call `rewrite_table()`:

```sql
-- Source table
CREATE TABLE measurement (id int, city_id int NOT NULL, logdate date NOT NULL, peaktemp int, PRIMARY KEY(id, logdate));

-- Target table with new schema (e.g., bigint id + partitioning)
CREATE TABLE measurement_aux (id bigint, city_id int NOT NULL, logdate date NOT NULL, peaktemp int, PRIMARY KEY(id, logdate))
  PARTITION BY RANGE (logdate);

CREATE TABLE measurement_y2006m02 PARTITION OF measurement_aux FOR VALUES FROM ('2006-02-01') TO ('2006-03-01');

-- Perform the rewrite (copies data, applies concurrent changes, then renames)
SELECT rewrite_table('measurement', 'measurement_aux', 'measurement_old');
```

Both source and target tables must have an identity index (typically a primary key). The function copies all rows, replays concurrent changes via logical decoding, then atomically renames the tables.

### Progress Monitoring

```sql
SELECT * FROM pg_rewrite_progress;
```

Shows `ins_initial` (initial rows copied), `ins`, `upd`, `del` (concurrent changes applied).

### Configuration

- **`rewrite.max_xlock_time`** -- Maximum time (ms) the exclusive lock is held during the final rename stage. Default `0` (unlimited). Set to e.g. `100` to limit lock duration to 0.1s; the function will retry if exceeded.

```sql
SET rewrite.max_xlock_time TO 100;
```

### Constraints Handling

- PRIMARY KEY, UNIQUE, EXCLUDE: add to target table before calling `rewrite_table()`
- CHECK, NOT NULL (PG18+), FOREIGN KEY: created automatically as NOT VALID; validate with `ALTER TABLE ... VALIDATE CONSTRAINT ...`
