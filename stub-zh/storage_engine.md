## 用法

来源：[README](https://github.com/saulojb/storage_engine/blob/main/README.md)，[release 1.3.4](https://github.com/saulojb/storage_engine/releases/tag/v1.3.4)，[PGXN 1.3.4](https://pgxn.org/dist/storage_engine/1.3.4/)，[PGXN changelog](https://pgxn.org/dist/storage_engine/1.3.4/CHANGELOG.html)，[META.json](https://github.com/saulojb/storage_engine/blob/main/META.json)

`storage_engine` 在 `engine` schema 中提供两种 PostgreSQL table access method：

- `colcompress`：面向列式压缩存储，支持 vectorized execution、min/max pruning 与 parallel scan。
- `rowcompress`：面向行批压缩，支持 parallel scan。

```sql
CREATE EXTENSION storage_engine;
```

### 快速开始

使用任一 access method 建表：

```sql
CREATE TABLE events (
  ts timestamptz NOT NULL,
  user_id bigint,
  event_type text,
  value float8
) USING colcompress;

CREATE TABLE logs (
  id bigserial,
  logged_at timestamptz NOT NULL,
  message text
) USING rowcompress;
```

### 主要调优旋钮

上游文档列出的 session 级 GUC 包括：

- `storage_engine.enable_parallel_execution`
- `storage_engine.min_parallel_processes`
- `storage_engine.enable_vectorization`
- `storage_engine.enable_custom_scan`
- `storage_engine.enable_column_cache`
- `storage_engine.enable_columnar_index_scan`
- `storage_engine.enable_dml`
- `storage_engine.stripe_row_limit`
- `storage_engine.chunk_group_row_limit`
- `storage_engine.compression_level`

README 说明这些 GUC 会在库被加载后可见；如果希望每个 session 一开始就可用，可把 `storage_engine` 加入 `shared_preload_libraries`。

### 类型与操作符

`engine.uint8` 为需要完整 `0` 到 `2^64 - 1` 范围的 `colcompress` 工作负载存储无符号 64 位值。上游文档列出了 comparison operators（`=`、`<>`、`<`、`<=`、`>`、`>=`）、B-tree 与 hash opclasses、与 `bigint`、`numeric`、`text` 的双向 casts，以及 `engine.min`、`engine.max`、`engine.sum` aggregates。

### 常用管理函数

用于 `colcompress` 表：

```sql
SELECT engine.alter_colcompress_table_set(
  'events'::regclass,
  orderby => 'ts ASC, user_id ASC',
  compression => 'zstd',
  compression_level => 9
);

SELECT engine.colcompress_merge('events');
CALL engine.colcompress_repack('events');
CALL engine.colcompress_repack('events', 0.7);
```

在 1.3.4 中，`engine.colcompress_repack(table_name regclass, min_fill_ratio float8 DEFAULT 0.9)` 是用于 `colcompress` 表在线逐 stripe defragmentation 的 procedure。它会 repack live-row ratio 低于阈值的 stripes。需要按 `orderby` key 做旧式 full-table rewrite 和全局排序时，使用 `engine.colcompress_merge()`。

用于 `rowcompress` 表：

```sql
SELECT engine.alter_rowcompress_table_set(
  'logs'::regclass,
  batch_size => 10000,
  compression => 'zstd',
  compression_level => 5
);

SELECT engine.rowcompress_repack('logs');
```

### 何时选择哪种 AM

- `colcompress` 适合分析扫描、聚合与范围谓词，这类场景能受益于列裁剪、vectorization 以及 stripe/chunk pruning。
- `rowcompress` 适合追加型日志或通常会整行一起读取的宽表，此时压缩收益往往高于列投影。
- 对 `colcompress` 的点查，上游建议启用 `storage_engine.enable_columnar_index_scan` 或 per-table `index_scan`。

### 注意事项

- 使用 `ALTER EXTENSION storage_engine UPDATE TO '1.3.4';` 升级已有安装。
- 旧的 `FUNCTION engine.colcompress_repack(regclass)` alias 在 1.3.4 中已移除。已有调用方应改用 `CALL engine.colcompress_repack('table')` 做 stripe defragmentation，或用 `SELECT engine.colcompress_merge('table')` 执行旧的 full rewrite 行为。
- `colcompress` 与 `rowcompress` 都不支持 foreign key 或 `AFTER ROW` triggers。
- 不支持 `VACUUM FULL`，也不支持 `CREATE UNLOGGED TABLE ... USING colcompress`；上游建议改用扩展自带的 repack 函数。
- 在 `colcompress` 上，`orderby` 与 B-tree index 组合可能禁用 sort-on-write 路径，而 ordered columns 上的 B-tree indexes 可能让 range query 失去 stripe pruning 收益。全局顺序重要时，请在装载数据后运行 `engine.colcompress_merge()`；分析型表优先使用 `index_scan => false`。
