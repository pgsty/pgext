


## 用法

来源：

- [Hydra v1.1.2 README](https://github.com/hydradatabase/columnar/blob/v1.1.2/README.md)
- [Hydra Columnar README](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/README.md)
- [Columnar storage README](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/src/backend/columnar/README.md)
- [columnar.control](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/src/backend/columnar/columnar.control)
- [CHANGELOG 1.1.2](https://github.com/hydradatabase/columnar/blob/v1.1.2/CHANGELOG.md)

> [!WARNING]
> `columnar` 是 PostgreSQL 扩展名，Pigsty 中的包名是 `hydra`。Pigsty 元数据已将该扩展标记为 obsolete / no longer maintained；本地包仅保留 PostgreSQL 14-16。新项目应优先选择仍在维护的分析类扩展。

Hydra Columnar 是 PostgreSQL 的列式表访问方法。它把指定表存为列式格式，用于降低分析扫描、压缩型数据集、少量列投影和聚合查询的 I/O 成本。该扩展源自 Citus Columnar，并通过 `CREATE EXTENSION columnar` 暴露。

### 创建列式表

```sql
CREATE EXTENSION IF NOT EXISTS columnar;

CREATE TABLE events_columnar (
  event_id     bigint,
  account_id   bigint,
  event_time   timestamptz,
  event_type   text,
  amount       numeric
) USING columnar;

INSERT INTO events_columnar
SELECT
  g,
  g % 10000,
  now() - (g || ' seconds')::interval,
  CASE WHEN g % 10 = 0 THEN 'purchase' ELSE 'view' END,
  (g % 1000)::numeric / 10
FROM generate_series(1, 1000000) AS g;

SELECT event_type, count(*), sum(amount)
FROM events_columnar
WHERE event_time >= now() - interval '1 day'
GROUP BY event_type;
```

创建表或物化视图时使用 `USING columnar`。读取和批量写入仍使用普通 PostgreSQL SQL，但底层存储更适合大规模分析扫描，不适合作为高频 OLTP 表的默认选择。

### 表选项

```sql
SELECT columnar.alter_columnar_table_set(
  'events_columnar'::regclass,
  compression           => 'zstd',
  compression_level     => 3,
  stripe_row_limit      => 150000,
  chunk_group_row_limit => 10000
);

SELECT * FROM columnar.options;

SELECT columnar.alter_columnar_table_reset(
  'events_columnar'::regclass,
  compression => true,
  stripe_row_limit => true
);
```

可配置项包括 `compression`、`compression_level`、`stripe_row_limit` 和 `chunk_group_row_limit`。可用压缩算法取决于构建方式，文档中列出的值包括 `none`、`pglz`、`zstd`、`lz4` 和 `lz4hc`。表选项变化只影响之后写入的数据；已有 stripe 不会自动重写。

也可以用 GUC 设置新建列式表的默认值：

```sql
SET columnar.compression = 'zstd';
SET columnar.compression_level = 3;
SET columnar.stripe_row_limit = 150000;
SET columnar.chunk_group_row_limit = 10000;
```

这些 GUC 影响新建表，而不是已有表后续生成的新 stripe。

### 转换已有表

```sql
CREATE TABLE events_heap (
  event_id bigint,
  payload  jsonb
);

INSERT INTO events_heap
SELECT g, jsonb_build_object('kind', 'view', 'seq', g)
FROM generate_series(1, 10000) AS g;

SELECT columnar.alter_table_set_access_method('events_heap', 'columnar');
SELECT columnar.alter_table_set_access_method('events_heap', 'heap');
```

`columnar.alter_table_set_access_method(table, method)` 会把 heap 表重写为 columnar 存储，或把 columnar 表改回 heap 存储。转换前应检查外键、identity 列、行级安全策略、触发器、索引、约束、继承关系和存储选项；辅助函数会拒绝或跳过它无法安全重建的特性。

### 分区表

```sql
CREATE TABLE measurements (
  ts timestamptz,
  device_id bigint,
  value double precision
) PARTITION BY RANGE (ts);

CREATE TABLE measurements_2024 PARTITION OF measurements
  FOR VALUES FROM ('2024-01-01') TO ('2025-01-01')
  USING columnar;

CREATE TABLE measurements_hot PARTITION OF measurements
  FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');
```

分区表可以混用行存与列存分区。只命中行存分区的操作可以使用普通行表行为；如果操作可能触及列存分区，就必须遵守 columnar 的限制。常见做法是把近期可变数据放在 heap 分区，把较老的分析历史放在 columnar 分区。

### 维护与查看

```sql
VACUUM VERBOSE events_columnar;
VACUUM FULL events_columnar;

SELECT * FROM columnar.stats('events_columnar'::regclass);
SELECT columnar.vacuum('events_columnar'::regclass);
SELECT columnar.vacuum_full('public', 0.1, 25);
```

`VACUUM VERBOSE` 会报告列式存储统计信息，例如文件大小、压缩率、行数、stripe 数和 chunk 数。`columnar.stats()` 暴露 stripe 级元数据。`columnar.vacuum()` 与 `columnar.vacuum_full()` 用于增量压缩列式存储；普通 `VACUUM` 主要扫描元数据，成本低于完整重写。

### 注意事项

- Pigsty 元数据中该扩展已 obsolete，并且与 `citus` / `citus_columnar` 一类列式存储存在冲突。除非已验证具体组合，否则不要在同一 PostgreSQL 大版本中混装多个冲突的列式表访问方法。
- Pigsty 仅为 PostgreSQL 14-16 保留 `hydra` / `columnar` 包；本地已将 PostgreSQL 17 和 18 标记为不支持。
- Hydra 1.1.x 已改进 update/delete 与 upsert 支持，但项目文档仍明确列式存储不适合频繁大规模更新、小事务和 OLTP 单行读写负载。
- 受限或不支持的场景包括 logical decoding、unlogged columnar table、serializable isolation、部分 scan 类型，以及许多非 btree / hash 索引。依赖约束和索引型约束前应先实测。
- `columnar` schema 中包含 `columnar.options`、`columnar.stripe`、`columnar.chunk_group`、`columnar.chunk` 等内部元数据表。可以通过公开函数查看，但不要直接修改这些元数据表。
