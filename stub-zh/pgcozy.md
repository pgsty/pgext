

## 用法

> [pgcozy: 根据之前的 pg_buffercache 快照预热 PostgreSQL 共享缓冲区](https://github.com/vventirozos/pgcozy)

pgcozy 对当前共享缓冲区状态进行快照，并可稍后从这些快照恢复（预热）缓冲区。需要 `pg_buffercache` 和 `pg_prewarm` 扩展。

### 初始化

```sql
SELECT pgcozy_init();
```

创建 `pgcozy` 模式，包含 `snapshots` 表和 `cozy_type` 类型。

### 拍摄快照

```sql
-- 快照使用次数 >= 3 的缓冲页（热度 1-5）
SELECT pgcozy_snapshot(3);

-- 快照所有缓冲页（热度 = 0）
SELECT pgcozy_snapshot(0);
```

快照以 JSONB 格式存储在 `pgcozy.snapshots` 中，包含列：`id`、`snapshot_date`、`snapshot`。每条记录包含 `table_name`、`block_no` 和 `popularity`。

### 从快照预热

```sql
-- 从特定快照 ID 预热
SELECT pgcozy_warm(1);

-- 从最新快照预热
SELECT pgcozy_warm(0);
```

### 查看快照

```sql
SELECT id, snapshot_date FROM pgcozy.snapshots;
```

快照以 JSONB 存储，可以查看、备份或在服务器之间传输。
