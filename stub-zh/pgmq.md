


## 用法

> [pgmq: PostgreSQL 轻量级消息队列](https://github.com/pgmq/pgmq)

PGMQ 是基于 PostgreSQL 构建的轻量级消息队列，提供可见性超时内的"恰好一次"投递保证、FIFO 队列、基于主题的路由和消息归档。

```sql
CREATE EXTENSION pgmq;
```

### 创建队列

```sql
SELECT pgmq.create('my_queue');
```

### 发送消息

```sql
-- 发送单条消息（返回 msg_id）
SELECT * FROM pgmq.send(
  queue_name => 'my_queue',
  msg        => '{"foo": "bar"}'
);

-- 延迟发送（5 秒内不可见）
SELECT * FROM pgmq.send(
  queue_name => 'my_queue',
  msg        => '{"foo": "bar"}',
  delay      => 5
);

-- 批量发送消息
SELECT pgmq.send_batch(
  queue_name => 'my_queue',
  msgs       => ARRAY['{"a":1}','{"b":2}','{"c":3}']::jsonb[]
);
```

### 读取消息

读取消息并在可见性超时期间（以秒为单位）使其不可见：

```sql
SELECT * FROM pgmq.read(
  queue_name => 'my_queue',
  vt         => 30,    -- 可见性超时（秒）
  qty        => 2      -- 读取消息数量
);
```

### 弹出消息

读取并立即删除一条消息：

```sql
SELECT * FROM pgmq.pop('my_queue');
```

### 删除消息

```sql
SELECT pgmq.delete('my_queue', 6);
```

### 归档消息

将消息从队列移动到归档表以便长期保留：

```sql
SELECT pgmq.archive(queue_name => 'my_queue', msg_id => 2);

-- 批量归档
SELECT pgmq.archive(queue_name => 'my_queue', msg_ids => ARRAY[3, 4, 5]);
```

查看已归档的消息：

```sql
SELECT * FROM pgmq.a_my_queue;
```

### 删除队列

```sql
SELECT pgmq.drop_queue('my_queue');
```

### 可见性超时

消息在被读取后，在可见性超时（`vt`）期间变为不可见。如果在此时间内未被删除或归档，它们将重新变为可见以供其他消费者处理。请将 `vt` 设置为大于预期处理时间。

### 主要函数

| 函数 | 描述 |
|------|------|
| `pgmq.create(queue_name)` | 创建新队列 |
| `pgmq.send(queue_name, msg, [delay])` | 发送消息 |
| `pgmq.send_batch(queue_name, msgs)` | 批量发送消息 |
| `pgmq.read(queue_name, vt, qty)` | 读取消息并设置可见性超时 |
| `pgmq.pop(queue_name)` | 原子性地读取并删除消息 |
| `pgmq.delete(queue_name, msg_id)` | 删除消息 |
| `pgmq.archive(queue_name, msg_id/msg_ids)` | 归档消息 |
| `pgmq.drop_queue(queue_name)` | 删除队列 |
