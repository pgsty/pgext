

## 用法

> [mimeo: PostgreSQL 实例之间专用的按表复制扩展](https://github.com/omniti-labs/mimeo)

提供 PostgreSQL 实例之间的按表复制，支持快照（全量复制）、增量（基于时间戳/ID）和 DML（插入/更新/删除）模式。

### 启用

```sql
CREATE SCHEMA mimeo;
CREATE EXTENSION mimeo SCHEMA mimeo;
```

需要 `dblink` 扩展。可选安装 `pg_jobmon` 用于监控。

### 设置数据源

```sql
-- 创建到源数据库的 dblink 连接
SELECT mimeo.dblink_mapping_create(
    p_mapping_name := 'source_db',
    p_data_source := 'host=sourcehost dbname=sourcedb user=replicator password=secret',
    p_superuser := true
);
```

### 快照复制（全表复制）

每次运行时复制整个源表：

```sql
SELECT mimeo.snapshot_maker(
    p_src_table := 'public.my_table',
    p_dblink_id := 1  -- 来自 dblink_mapping
);

-- 刷新快照
SELECT mimeo.refresh_snap('public.my_table');
```

### 增量复制（基于时间戳）

基于递增时间戳列复制行：

```sql
SELECT mimeo.inserter_maker(
    p_src_table := 'public.events',
    p_control := 'created_at',  -- 时间戳列
    p_dblink_id := 1
);

-- 增量刷新
SELECT mimeo.refresh_inserter('public.events');
```

对于有更新的表（不仅仅是插入）：

```sql
SELECT mimeo.updater_maker(
    p_src_table := 'public.orders',
    p_control := 'updated_at',
    p_dblink_id := 1
);

SELECT mimeo.refresh_updater('public.orders');
```

### DML 复制（插入/更新/删除）

通过源表上的触发器实现完整 DML 追踪：

```sql
SELECT mimeo.dml_maker(
    p_src_table := 'public.accounts',
    p_dblink_id := 1
);

SELECT mimeo.refresh_dml('public.accounts');
```

### 调度刷新

使用 `pg_jobmon` 或 cron 调度定期调用相应的 `refresh_*` 函数。

### 关键特性

- 三种复制模式：快照、增量、DML
- 按表复制（无需复制整个数据库）
- 在不同 PostgreSQL 版本之间工作
- 基于 `dblink` 实现跨数据库通信
