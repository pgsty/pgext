## Usage

Sources: [README](https://github.com/relytcloud/pg_ducklake/blob/v1.0.0/README.md), [v1.0.0 release](https://github.com/relytcloud/pg_ducklake/releases/tag/v1.0.0), [project docs](https://github.com/relytcloud/pg_ducklake/tree/v1.0.0/pg_ducklake/docs)

`pg_ducklake` adds DuckLake tables to PostgreSQL. DuckLake metadata lives in PostgreSQL while table data is stored in Parquet and queried through DuckDB, giving PostgreSQL SQL clients access to lakehouse features such as snapshots, time travel, partitioning, sort keys, and external object storage.

### Create A DuckLake Table

```sql
CREATE EXTENSION pg_ducklake;

CREATE TABLE events (
  id int,
  kind text,
  ts timestamptz
) USING ducklake;

INSERT INTO events VALUES
  (1, 'login', now()),
  (2, 'click', now());

SELECT * FROM events ORDER BY id;
```

Set a table path explicitly when data should live outside the default path:

```sql
CREATE TABLE lake_events (
  id int,
  payload jsonb
) WITH (
  ducklake.table_path = 's3://my-bucket/prefix/'
) USING ducklake;
```

### Time Travel

Each commit creates a snapshot. Capture a snapshot id before a change, then query the older state:

```sql
SELECT max(snapshot_id) AS before_delete
FROM ducklake.ducklake_snapshot \gset

DELETE FROM events WHERE id = 1;

SELECT * FROM ducklake.time_travel('events'::regclass, :before_delete);
```

### Convert And Load Data

Create DuckLake tables from existing PostgreSQL heap tables or external data readers:

```sql
CREATE TABLE row_store AS
SELECT i AS id, 'hello pg_ducklake' AS msg
FROM generate_series(1, 10000) AS i;

CREATE TABLE col_store USING ducklake AS
SELECT * FROM row_store;

CREATE TABLE titanic USING ducklake AS
SELECT * FROM ducklake.read_csv('https://raw.githubusercontent.com/datasciencedojo/datasets/master/titanic.csv');
```

### Inlining, Partitioning, And Maintenance

Small writes are inlined in metadata by default to avoid creating many tiny Parquet files. Tune the row limit or flush explicitly:

```sql
CALL ducklake.set_option('data_inlining_row_limit', 100);
SELECT * FROM ducklake.flush_inlined_data('events'::regclass);
```

Partition and sort tables for pruning and analytics:

```sql
CALL ducklake.set_partition('events'::regclass, 'bucket(4, id)', 'month(ts)');
CREATE INDEX ON events USING ducklake_sorted (id, ts);
```

Run maintenance on demand when automatic background maintenance is not enough:

```sql
SELECT * FROM ducklake.merge_adjacent_files('events'::regclass);
CALL ducklake.set_option('expire_older_than', '7 days');
SELECT * FROM ducklake.expire_snapshots();
SELECT * FROM ducklake.cleanup_old_files();
```

### External DuckDB Access

DuckDB clients can attach the same DuckLake metadata:

```sql
INSTALL ducklake;
LOAD ducklake;
ATTACH 'ducklake:postgres:dbname=postgres host=localhost' AS my_ducklake
  (METADATA_SCHEMA 'ducklake');

SELECT * FROM my_ducklake.public.events;
```

### Caveats

- Version 1.0.0 supports PostgreSQL 14-18.
- The README lists Ubuntu 22.04-24.04 and macOS as source-build targets.
- Cloud credentials are stored through a `ducklake_secret` foreign server and per-user mappings; protect those catalog objects like other database secrets.
- For incremental heap-to-DuckLake conversion, upstream points to the separate `pg_duckpipe` project.
