## 用法

来源：

- [扩展 SQL 定义](https://github.com/johto/pg_metrics/blob/ce7dcbf8ad5bec4bf27906c28f404f702f89cbaf/pg_metrics--1.0.sql)
- [共享内存实现](https://github.com/johto/pg_metrics/blob/ce7dcbf8ad5bec4bf27906c28f404f702f89cbaf/pg_metrics.c)
- [扩展控制文件](https://github.com/johto/pg_metrics/blob/ce7dcbf8ad5bec4bf27906c28f404f702f89cbaf/pg_metrics.control)

`pg_metrics` 在 PostgreSQL 共享内存中保存具名整数计数器。后端可以原子地增加计数器，监控会话则可以枚举当前值与容量。它是构建数据库内自定义指标的轻量组件，并不是通用时序存储。

### 核心流程

该库必须在 postmaster 启动期间初始化。将其加入 `shared_preload_libraries`，按需设置计数器容量，并重启 PostgreSQL，然后再创建 SQL 扩展。

```conf
shared_preload_libraries = 'pg_metrics'
pg_metrics.max = 100
```

```sql
CREATE EXTENSION pg_metrics;

SELECT counter_add('jobs_processed', 1);
SELECT * FROM metrics();
SELECT * FROM metrics_stats();
```

`counter_add` 返回计数器之前的值，然后应用增量。`int8` 签名接受负增量，因此也可以减小计数器。

### 用户接口

- `metric_type`：目前只包含 `COUNTER` 的枚举。
- `counter_add(counter text, increment int8)`：首次使用时创建具名计数器，并原子地累加增量。
- `metrics()`：为每个共享计数器返回 `metric_name`、`metric_type` 和 `counter_value`。
- `metrics_stats()`：返回 `max_metrics` 和 `num_metrics`。
- `pg_metrics.max`：计数器上限的 postmaster 参数；源码默认值为 `50`，最小值为 `10`。

### 运维说明

预加载和重启服务器是强制要求：共享内存未初始化时，`counter_add` 返回 `NULL`。名称超过 127 字节会被拒绝；共享哈希达到 `pg_metrics.max` 后再增加新计数器也会返回 `NULL`。值只存在于共享内存中，共享内存状态重建时会被清零；版本 `1.0` 没有持久化或重置 SQL API。计数器由各后端共享，而调用它们的每个数据库都需要创建相应 SQL 对象。如需历史指标，应从外部导出或定期快照这些值。
