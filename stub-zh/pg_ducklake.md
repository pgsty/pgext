


## 用法

来源：[README](https://github.com/relytcloud/pg_ducklake/blob/v1.0.0/README.md)、[v1.0.0 release](https://github.com/relytcloud/pg_ducklake/releases/tag/v1.0.0)、[project docs](https://github.com/relytcloud/pg_ducklake/tree/v1.0.0/pg_ducklake/docs)

`pg_ducklake` 为 PostgreSQL 增加 DuckLake tables。DuckLake metadata 存储在 PostgreSQL 中，表数据以 Parquet 存储并通过 DuckDB 查询，让 PostgreSQL SQL clients 可以访问 snapshots、time travel、partitioning、sort keys 和外部对象存储等 lakehouse 功能。

### 创建 DuckLake 表

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

当数据需要放在默认路径之外时，显式设置 table path：

```sql
CREATE TABLE lake_events (
  id int,
  payload jsonb
) WITH (
  ducklake.table_path = 's3://my-bucket/prefix/'
) USING ducklake;
```

### Time Travel

每次 commit 都会创建 snapshot。在修改前记录 snapshot id，然后查询旧状态：

```sql
SELECT max(snapshot_id) AS before_delete
FROM ducklake.ducklake_snapshot \gset

DELETE FROM events WHERE id = 1;

SELECT * FROM ducklake.time_travel('events'::regclass, :before_delete);
```

### 转换与加载数据

可从已有 PostgreSQL heap tables 或外部 data readers 创建 DuckLake tables：

```sql
CREATE TABLE row_store AS
SELECT i AS id, 'hello pg_ducklake' AS msg
FROM generate_series(1, 10000) AS i;

CREATE TABLE col_store USING ducklake AS
SELECT * FROM row_store;

CREATE TABLE titanic USING ducklake AS
SELECT * FROM ducklake.read_csv('https://raw.githubusercontent.com/datasciencedojo/datasets/master/titanic.csv');
```

### Inlining、分区与维护

小批写入默认 inline 到 metadata 中，避免产生大量小 Parquet files。可调整行数限制或显式 flush：

```sql
CALL ducklake.set_option('data_inlining_row_limit', 100);
SELECT * FROM ducklake.flush_inlined_data('events'::regclass);
```

为表设置 partition 和 sort keys，以便 pruning 和 analytics：

```sql
CALL ducklake.set_partition('events'::regclass, 'bucket(4, id)', 'month(ts)');
CREATE INDEX ON events USING ducklake_sorted (id, ts);
```

当自动后台维护不够时，可以按需执行维护：

```sql
SELECT * FROM ducklake.merge_adjacent_files('events'::regclass);
CALL ducklake.set_option('expire_older_than', '7 days');
SELECT * FROM ducklake.expire_snapshots();
SELECT * FROM ducklake.cleanup_old_files();
```

### 外部 DuckDB 访问

DuckDB clients 可以 attach 同一份 DuckLake metadata：

```sql
INSTALL ducklake;
LOAD ducklake;
ATTACH 'ducklake:postgres:dbname=postgres host=localhost' AS my_ducklake
  (METADATA_SCHEMA 'ducklake');

SELECT * FROM my_ducklake.public.events;
```

### 注意事项

- 版本 1.0.0 支持 PostgreSQL 14-18。
- README 列出的源码构建目标包括 Ubuntu 22.04-24.04 和 macOS。
- Cloud credentials 通过 `ducklake_secret` foreign server 和 per-user mappings 存储；应像保护其他数据库 secrets 一样保护这些 catalog objects。
- 对于 incremental heap-to-DuckLake conversion，上游指向单独的 `pg_duckpipe` 项目。
