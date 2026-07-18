


## 用法

来源：[README](https://github.com/relytcloud/pg_ducklake/blob/v1.0.0/README.md)、[v1.0.0 版本](https://github.com/relytcloud/pg_ducklake/releases/tag/v1.0.0)、[项目文档](https://github.com/relytcloud/pg_ducklake/tree/v1.0.0/pg_ducklake/docs)

`pg_ducklake` 为 PostgreSQL 增加 DuckLake 表。DuckLake 元数据存储在 PostgreSQL 中，表数据以 Parquet 格式存储并通过 DuckDB 查询，让 PostgreSQL SQL 客户端可以使用快照、时间旅行查询、分区、排序键和外部对象存储等湖仓功能。

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

当数据需要放在默认路径之外时，应显式设置表路径：

```sql
CREATE TABLE lake_events (
  id int,
  payload jsonb
) WITH (
  ducklake.table_path = 's3://my-bucket/prefix/'
) USING ducklake;
```

### 时间旅行查询

每次提交都会创建快照。在修改前记录快照 ID，然后查询旧状态：

```sql
SELECT max(snapshot_id) AS before_delete
FROM ducklake.ducklake_snapshot \gset

DELETE FROM events WHERE id = 1;

SELECT * FROM ducklake.time_travel('events'::regclass, :before_delete);
```

### 转换与加载数据

可以从已有的 PostgreSQL 堆表或外部数据读取器创建 DuckLake 表：

```sql
CREATE TABLE row_store AS
SELECT i AS id, 'hello pg_ducklake' AS msg
FROM generate_series(1, 10000) AS i;

CREATE TABLE col_store USING ducklake AS
SELECT * FROM row_store;

CREATE TABLE titanic USING ducklake AS
SELECT * FROM ducklake.read_csv('https://raw.githubusercontent.com/datasciencedojo/datasets/master/titanic.csv');
```

### 内联、分区与维护

小批量写入默认内联到元数据中，避免产生大量小 Parquet 文件。可以调整行数限制或显式刷出数据：

```sql
CALL ducklake.set_option('data_inlining_row_limit', 100);
SELECT * FROM ducklake.flush_inlined_data('events'::regclass);
```

为表设置分区键和排序键，以便执行裁剪和分析：

```sql
CALL ducklake.set_partition('events'::regclass, 'bucket(4, id)', 'month(ts)');
CREATE INDEX ON events USING ducklake_sorted (id, ts);
```

当自动后台维护不足时，可以按需执行维护：

```sql
SELECT * FROM ducklake.merge_adjacent_files('events'::regclass);
CALL ducklake.set_option('expire_older_than', '7 days');
SELECT * FROM ducklake.expire_snapshots();
SELECT * FROM ducklake.cleanup_old_files();
```

### 外部 DuckDB 访问

DuckDB 客户端可以挂载同一份 DuckLake 元数据：

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
- 云凭据通过 `ducklake_secret` 外部服务器和逐用户映射存储；应像保护其他数据库密钥一样保护这些目录对象。
- 对于从堆表到 DuckLake 的增量转换，上游指向单独的 `pg_duckpipe` 项目。
