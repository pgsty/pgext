## 用法

来源：

- [Citus v14.2.0 columnar 控制文件](https://github.com/citusdata/citus/blob/v14.2.0/src/backend/columnar/citus_columnar.control)
- [Citus v14.2.0 columnar 选项辅助函数](https://github.com/citusdata/citus/blob/v14.2.0/src/backend/columnar/sql/udfs/alter_columnar_table_set/latest.sql)
- [Citus 列存储文档](https://docs.citusdata.com/en/stable/admin_guide/table_management.html#columnar-storage)
- [Citus v14.2.0 发行说明](https://github.com/citusdata/citus/releases/tag/v14.2.0)

`citus_columnar` 为 PostgreSQL 提供面向追加写入的列式表访问方法。它随 Citus 14.2 软件包交付，但属于独立扩展：软件包版本为 `14.2.0`，而扩展控制版本为 `14.2-1`。适用于以扫描为主，且工作负载符合其写入与功能限制的归档或分析表。

### 创建列存储表

```sql
CREATE EXTENSION citus_columnar;

CREATE TABLE events_archive (
  event_at timestamptz NOT NULL,
  tenant_id bigint NOT NULL,
  kind text,
  payload jsonb
) USING columnar;
```

`citus_columnar` 本身不要求 `shared_preload_libraries`。如果数据库还使用分布式 `citus` 扩展，仍需预加载 `citus`。

### 加载与查询数据

列存储会将行组成条带，并按数据块压缩各列。使用大小合理的事务进行批量插入，比持续执行微型事务更有利于产生良好条带。

```sql
INSERT INTO events_archive
SELECT event_at, tenant_id, kind, payload
FROM events
WHERE event_at < now() - interval '90 days';

SELECT tenant_id, count(*), min(event_at), max(event_at)
FROM events_archive
GROUP BY tenant_id;
```

### 使用 Citus 扩展转换

如果主 `citus` 扩展也已预加载并安装，可以使用其辅助函数转换本地表或分布式表：

```sql
SELECT alter_table_set_access_method('events_archive', 'columnar');
SELECT alter_table_set_access_method('events_archive', 'heap');
```

转换会重写表。转换为列存储会删除现有索引，因此执行前应清点依赖的索引和约束，并为重写安排足够的磁盘空间和锁定时间。

`alter_table_set_access_method()` 属于 `citus`，而不是独立的 `citus_columnar`。没有主扩展时，应新建 `USING columnar` 表并将数据复制进去，而不要假定该辅助函数存在。

### 调整压缩

使用文档所述的辅助函数检查和修改表级列存储选项：

```sql
SELECT alter_columnar_table_set(
  'events_archive',
  compression => 'zstd',
  compression_level => 3,
  stripe_row_limit => 150000,
  chunk_group_row_limit => 10000
);
```

新设置只影响后续写入的条带。如果旧条带也需要采用新布局，请重写现有数据。

### 运维边界

- 列存储表面向追加型使用场景。它不支持 `UPDATE` 和 `DELETE`，回滚写入留下的空间也无法通过普通的堆表式维护回收。
- TOAST 不可用；大值会保持行内存储，并可能触及 PostgreSQL 行大小限制。
- 不支持行锁、`AFTER ... FOR EACH ROW` 触发器、可串行化隔离、逻辑解码、外键、非日志表以及多种扫描类型。采用该访问方法前，应检查当前上游限制列表。
- 不应把普通堆表在索引、vacuum、复制、触发器和约束方面的惯例直接套用到列存储上。应使用有代表性的列存储表验证每项必需的数据库功能。
- 扩展安装在 `pg_catalog` 中，不可重定位，SQL 版本为 `14.2-1`；检查或更新 `pg_extension` 时应使用该版本，而不是软件包版本 `14.2.0`。
