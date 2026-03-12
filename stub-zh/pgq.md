


## 用法

> [pgq: PostgreSQL 通用高性能无锁队列](https://github.com/pgq/pgq)

PgQ 是一个 PostgreSQL 扩展，提供通用的高性能无锁队列，带有简单的 SQL 函数 API。它使用生产者-消费者模型，基于批次进行事件处理。

```sql
CREATE EXTENSION pgq;
```

### 核心概念

- **队列（Queue）**：命名的事件流。生产者插入事件，消费者按批次消费。
- **消费者（Consumer）**：注册到队列上的命名订阅者。每个消费者跟踪自己的位置。
- **批次（Batch）**：一组一起获取的事件。消费者逐批处理事件。
- **心跳进程（Ticker）**：后台进程，定期创建批次边界（tick）。

### 队列管理

```sql
-- 创建队列
SELECT pgq.create_queue('myqueue');

-- 删除队列
SELECT pgq.drop_queue('myqueue');

-- 获取队列信息
SELECT * FROM pgq.get_queue_info();
SELECT * FROM pgq.get_queue_info('myqueue');
```

### 消费者注册

```sql
-- 在队列上注册消费者
SELECT pgq.register_consumer('myqueue', 'myconsumer');

-- 注销消费者
SELECT pgq.unregister_consumer('myqueue', 'myconsumer');

-- 获取消费者信息
SELECT * FROM pgq.get_consumer_info('myqueue');
```

### 生产事件

```sql
-- 向队列插入事件
SELECT pgq.insert_event('myqueue', 'event_type', 'event_data');

-- 插入带额外字段的事件
SELECT pgq.insert_event('myqueue', 'event_type', 'event_data',
                         'extra1', 'extra2', 'extra3', 'extra4');
```

### 消费事件

```sql
-- 获取下一批事件（返回 batch_id，无新批次时返回 NULL）
SELECT pgq.next_batch('myqueue', 'myconsumer');

-- 从批次中获取事件
SELECT * FROM pgq.get_batch_events(:batch_id);

-- 重试失败的事件（在指定间隔后重新出现）
SELECT pgq.event_retry(:batch_id, :event_id, :retry_seconds);

-- 标记批次完成
SELECT pgq.finish_batch(:batch_id);
```

### 典型消费者循环

```sql
-- 1. 获取下一批次
SELECT pgq.next_batch('myqueue', 'myconsumer') AS batch_id;

-- 2. 如果 batch_id 不为 NULL，获取事件
SELECT * FROM pgq.get_batch_events(:batch_id);

-- 3. 处理事件，重试失败的
SELECT pgq.event_retry(:batch_id, :event_id, 60);

-- 4. 完成批次
SELECT pgq.finish_batch(:batch_id);
```

### 维护

PgQ 需要在后台运行心跳守护进程（`pgqd`），用于创建批次边界并执行表轮转和重试事件处理等维护任务。

### 主要函数

| 函数 | 描述 |
|------|------|
| `pgq.create_queue(name)` | 创建新队列 |
| `pgq.drop_queue(name)` | 删除队列 |
| `pgq.register_consumer(queue, consumer)` | 注册消费者 |
| `pgq.unregister_consumer(queue, consumer)` | 注销消费者 |
| `pgq.insert_event(queue, type, data, ...)` | 插入事件 |
| `pgq.next_batch(queue, consumer)` | 获取下一批次 ID |
| `pgq.get_batch_events(batch_id)` | 从批次获取事件 |
| `pgq.event_retry(batch_id, event_id, seconds)` | 安排事件重试 |
| `pgq.finish_batch(batch_id)` | 标记批次已处理 |
| `pgq.get_queue_info([name])` | 获取队列统计信息 |
| `pgq.get_consumer_info(queue)` | 获取消费者统计信息 |
