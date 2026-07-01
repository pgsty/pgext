


## 用法

来源：[README](https://github.com/Treedo/pg_accumulator/blob/main/README.MD)，[documentation](https://github.com/Treedo/pg_accumulator/blob/main/doc/pg_accumulator.md)，[control file](https://github.com/Treedo/pg_accumulator/blob/main/pg_accumulator.control)

`pg_accumulator` 为 PostgreSQL 提供声明式累积寄存器。寄存器按任意维度和资源记录业务流水，并生成余额、发生额、流水历史、校验和重建等查询与维护函数。

### 启用

```sql
CREATE EXTENSION pg_accumulator;
```

如果高写入场景需要 delta buffer 和后台 worker，需要预加载共享库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_accumulator'
```

### 创建寄存器

```sql
SELECT accum.register_create(
  name := 'inventory',
  dimensions := '{"warehouse": "int", "product": "int"}',
  resources := '{"quantity": "numeric", "amount": "numeric"}',
  kind := 'balance'
);
```

一次调用会创建寄存器元数据、流水表、汇总表、余额缓存、触发器、生成查询函数和配套索引。

### 记账与查询

```sql
SELECT accum.register_post('inventory', '{
  "period": "2026-04-01",
  "document": "receipt:1",
  "dimensions": {"warehouse": 1, "product": 42},
  "resources": {"quantity": 10, "amount": 250}
}');

SELECT * FROM accum.inventory_balance(
  dimensions := '{"warehouse": 1, "product": 42}'
);

SELECT * FROM accum.inventory_turnover(
  from_date := '2026-04-01',
  to_date := '2026-04-30',
  dimensions := '{"warehouse": 1}',
  group_by := '["product"]'
);
```

### 更正与维护

使用 `register_unpost` 和 `register_repost` 处理追溯更正。必要时使用 `register_verify`、`register_rebuild_totals`、`register_rebuild_cache` 校验和修复派生状态。

```sql
SELECT accum.register_unpost('inventory', 'receipt:1');
SELECT accum.register_verify('inventory');
```

### 配置

文档中的 GUC 包括 `pg_accumulator.background_workers`、`pg_accumulator.delta_merge_interval`、`pg_accumulator.delta_merge_delay`、`pg_accumulator.delta_merge_batch_size`、`pg_accumulator.partitions_ahead`、`pg_accumulator.maintenance_interval`。
