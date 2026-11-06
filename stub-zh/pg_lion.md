## 用法

来源：

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [参考手册](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [扩展控制文件](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion.control)

`pg_lion` 0.3.0 为 PostgreSQL 16 及以上版本提供 Roaring 位图倒排索引，用于过滤、计数与分组聚合。它支持堆表，查询沿用 PostgreSQL 原有运算符；默认的通用 WAL 模式不需要预加载。

### 核心流程

为过滤与分组使用的列创建 `lion` 索引。经过清理、所有行均可见的页面允许计数避免访问堆；最近修改过的页面可能仍需到堆中检查可见性。

```sql
CREATE EXTENSION pg_lion;
CREATE TABLE lion_events (country text, amount integer);
INSERT INTO lion_events VALUES ('JP', 10), ('JP', 20), ('NZ', 30);
CREATE INDEX lion_country ON lion_events USING lion (country);
CREATE INDEX lion_amount ON lion_events USING lion (amount) WITH (slices = on);
VACUUM ANALYZE lion_events;
SELECT count(*), sum(amount) FROM lion_events WHERE country = 'JP';
SELECT country, count(*) FROM lion_events GROUP BY country;
```

支持标量相等、范围、`IN`、空值判断及组合条件。`lion` 也可索引数组、`jsonb` 和 `tsvector`；`lion_bm25_score` 用于全文匹配排序。针对 JSON 字段的表达式索引可以支持默认 `jsonb` 索引无法处理的分组。

### 选项与诊断

| 对象 | 用途 |
| --- | --- |
| `summaries = auto` | 预先合并位图，加速包含大量不同值的范围查询。 |
| `slices = on` | 保存位切片，支持整数、日期与时间值的过滤聚合；增加写入和存储成本。 |
| `lion_index_count(index, key)` | 通过指定索引统计某个键的行数。 |
| `lion_index_stats(index)` | 查看索引统计信息；初始授权给 `pg_stat_scan_tables`。 |
| `lion_index_verify(index, true)` | 检查结构与堆覆盖情况；属于受限诊断函数。 |
| `lion_index_wal_mode(index)` | 返回 `generic` 或 `rmgr`；属于受限诊断函数。 |

`pg_lion_citext` 添加 `citext` 运算符类。`pg_lion_vector` 添加 `lionvec` 访问方法，依赖 pgvector 0.7 或更新版本。这两个伴随扩展随父扩展一起分发，分别启用。

### WAL 与运行边界

创建这些扩展需要超级用户。可选的 `shared_preload_libraries = 'pg_lion'` 配置加服务器重启会启用自定义 WAL 资源管理器（默认 ID 为 170）。以该模式创建的索引在重建前，写入与 `VACUUM` 均依赖预加载；修改设置前先检查 `lion_index_wal_mode`。

只要旧的自定义 WAL 仍可能被回放，就需要在相关实例上保留预加载，包括备用库和归档恢复。`REINDEX` 不会改写历史 WAL。移除预加载前，应在动态库仍预加载时，将所有 `rmgr` 索引以 `wal_mode = generic` 重建，并核对各 WAL 回放使用方。

不支持唯一索引或 `INCLUDE` 列，且仅支持堆表。唯一性约束与普通有序扫描仍应使用 B-tree。加载动态库会安装规划器钩子，并可能影响其他查询的估算。为写密集型负载添加索引或位切片前应先测量。

1.0 之前的磁盘格式可能变化，应遵循各版本的 `REINDEX` 要求。0.3.0 之前的构建没有扩展升级路径：上游要求重新创建扩展及其索引。0.3.0 只提供安装脚本。
