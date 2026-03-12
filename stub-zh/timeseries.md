

## 用法

> [pg_timeseries: 时间序列便捷 API](https://github.com/ChuckHend/pg_timeseries)

本扩展为时间序列表的创建、维护和使用提供了一致的用户体验。

### 快速开始

假设你已经有一张分区表，只需调用 `enable_ts_table` 函数即可启用时间序列功能：

```sql
CREATE EXTENSION timeseries CASCADE;

SELECT enable_ts_table('sensor_readings');
```

这一调用会自动完成以下操作：

  * 使用 PostgreSQL 原生分区功能将表重构为一系列分区
  * 每个分区覆盖一段时间范围（默认一周）
  * 自动为未来创建新分区（默认提前一个月）
  * 每小时运行一次维护任务，创建缺失的分区和未来所需的分区


## 使用表

### 索引

时间序列表本质上就是标准的 PostgreSQL 分区表，所有 PostgreSQL 的现有功能都能正常使用。

传统 B-Tree 索引适用于时间序列数据，但在某些查询场景下（通常是返回大量结果的查询），BRIN 索引可能表现更优。如果每个分区不超过一百万条记录，建议先用 B-Tree。

### 分区大小

扩展提供了 `ts_part_info` 视图，可方便查看各分区的数据大小、索引大小和总大小。一般建议每个分区能容纳在可用内存的约四分之一以内。

### 数据保留

调用 `set_ts_retention_policy` 设置保留策略，超出保留窗口的分区每小时会被自动删除。使用 `clear_ts_retention_policy` 恢复默认行为（无限保留）。

### 压缩

调用 `set_ts_compression_policy` 可对超过指定时间间隔的分区使用列式存储进行压缩。压缩功能依赖 citus 和 citus_columnar 扩展：

```sql
CREATE EXTENSION citus;
CREATE EXTENSION citus_columnar;
```


## 分析辅助函数

### `first` 和 `last`

按一个维度分组，但按另一个维度取首条/末条记录：

```sql
SELECT machine_id,
       last(cpu_util, recorded_at)
FROM events
GROUP BY machine_id;
```

### `date_bin_table`

自动将时间序列值对齐到指定步长，并为无数据的时间段补充 NULL 行：

```sql
SELECT * FROM date_bin_table(NULL::target_table, '1 hour', '[2024-02-01 00:00, 2024-02-02 15:00]');
```

输出结果与直接查表相比有三点不同：

  * 按时间升序排列
  * 时间列的值对齐到指定步长
  * 无数据的时间段会补充 NULL 行

### `make_view_incremental`

将普通视图转换为自动保持更新的增量物化视图，无需手动 `REFRESH`。该功能基于 pg_ivm 实现：

```sql
CREATE EXTENSION pg_ivm;
```


## 依赖

* [pg_cron](https://github.com/citusdata/pg_cron)
* [pg_partman](https://github.com/pgpartman/pg_partman)

### 可选依赖

* [pg_ivm](https://github.com/sraoss/pg_ivm) — 增量物化视图
* [Citus & Citus Columnar](https://github.com/citusdata/citus) — 压缩功能
