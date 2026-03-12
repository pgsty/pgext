

## 用法

> [pg_partman: 按时间或 ID 管理分区表的扩展](https://github.com/pgpartman/pg_partman)

`pg_partman` 使用 PostgreSQL 原生声明式分区（v5.0+）自动创建和管理基于时间和基于数字的分区集。它处理新分区的添加和旧分区的按保留策略删除，并提供可选的后台工作进程用于自动维护。

### 创建扩展

```sql
CREATE SCHEMA partman;
CREATE EXTENSION pg_partman SCHEMA partman;
```

### 创建基于时间的分区集

```sql
CREATE TABLE public.measurements (
    id          bigserial,
    created_at  timestamptz NOT NULL DEFAULT now(),
    value       numeric
) PARTITION BY RANGE (created_at);

SELECT partman.create_parent(
    p_parent_table  := 'public.measurements',
    p_control       := 'created_at',
    p_interval      := '1 day'
);
```

### 创建基于序列/ID 的分区集

```sql
CREATE TABLE public.events (
    id      bigserial,
    data    text
) PARTITION BY RANGE (id);

SELECT partman.create_parent(
    p_parent_table  := 'public.events',
    p_control       := 'id',
    p_interval      := '100000'
);
```

### 运行维护

手动触发分区维护（创建新分区、删除过期分区）：

```sql
SELECT partman.run_maintenance();
```

或针对特定表：

```sql
SELECT partman.run_maintenance(p_parent_table := 'public.measurements');
```

### 配置保留策略

更新配置以设置保留策略：

```sql
UPDATE partman.part_config
SET    retention = '30 days',
       retention_keep_table = false
WHERE  parent_table = 'public.measurements';
```

### 后台工作进程

在 `postgresql.conf` 中启用自动维护：

```
shared_preload_libraries = 'pg_partman_bgw'
pg_partman_bgw.interval = 3600          -- 每小时运行一次（秒）
pg_partman_bgw.dbname = 'mydb'
```

### 将现有数据迁移到分区

```sql
CALL partman.partition_data_proc('public.measurements');
```

### 显示分区

```sql
SELECT * FROM partman.show_partitions('public.measurements');
```

### 撤销分区

```sql
CALL partman.undo_partition_proc('public.measurements');
```
