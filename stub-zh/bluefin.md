## 用法

来源：

- [Advanced Storage Pack 官方文档](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/)
- [官方配置指南](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/configuring/)
- [官方使用指南](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/using/)

`bluefin` 是 append-only EDB table access method，通过压缩 tuple header 并对相邻数据做 delta compression，面向时序与监控 workload。

### 启用

Bluefin 要求 PostgreSQL family 15 或更高版本。安装 Advanced Storage Pack 1.7.0，预加载、重启并创建扩展：

```ini
shared_preload_libraries = 'bluefin'
```

```sql
CREATE EXTENSION bluefin;
```

### 使用分区 Append-Only 存储

```sql
CREATE TABLE truck_logs (
  ts        timestamptz,
  truck_id  integer,
  latitude  float8,
  longitude float8
) PARTITION BY RANGE (ts);

CREATE TABLE truck_logs_2026_08
PARTITION OF truck_logs
FOR VALUES FROM ('2026-08-01') TO ('2026-09-01')
USING bluefin;
```

Bluefin 不允许 `UPDATE` 或 `DELETE`。保留策略应使用 partition detach/drop，而不是逐行删除。

### 估算与运维

```sql
SELECT *
FROM bluefin.estimate_compression('existing_heap', 10.0);
```

压缩估算只采样 heap page，不会移动数据。迁移前应验证插入顺序、分区轮转、索引支持、崩溃恢复、backup/restore 与 replica 行为。Append-only contract 是 schema-design 边界；需要纠错的应用必须写补偿行或选择其他 TAM。

