## 用法

来源：

- [PgQ 3.5.2 README](https://github.com/pgq/pgq/blob/v3.5.2/README.rst)
- [控制文件与安装权限](https://github.com/pgq/pgq/blob/v3.5.2/pgq.control)
- [批次获取 API](https://github.com/pgq/pgq/blob/v3.5.2/functions/pgq.next_batch.sql)
- [批次事件 API](https://github.com/pgq/pgq/blob/v3.5.2/functions/pgq.get_batch_events.sql)
- [重试 API](https://github.com/pgq/pgq/blob/v3.5.2/functions/pgq.event_retry.sql)
- [3.5.2 版发布说明](https://github.com/pgq/pgq/releases/tag/v3.5.2)

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
-- Create a queue
SELECT pgq.create_queue('myqueue');

-- Drop a queue
SELECT pgq.drop_queue('myqueue');

-- Get queue info
SELECT * FROM pgq.get_queue_info();
SELECT * FROM pgq.get_queue_info('myqueue');
```

### 消费者注册

```sql
-- Register a consumer on a queue
SELECT pgq.register_consumer('myqueue', 'myconsumer');

-- Unregister a consumer
SELECT pgq.unregister_consumer('myqueue', 'myconsumer');

-- Get consumer info
SELECT * FROM pgq.get_consumer_info('myqueue');
```

### 生产事件

```sql
-- Insert an event into a queue
SELECT pgq.insert_event('myqueue', 'event_type', 'event_data');

-- Insert with extra fields
SELECT pgq.insert_event('myqueue', 'event_type', 'event_data',
                         'extra1', 'extra2', 'extra3', 'extra4');
```

### 消费事件

```sql
-- Get the next batch of events (returns batch_id or NULL if no new batches)
SELECT pgq.next_batch('myqueue', 'myconsumer');

-- Get events from the batch
SELECT * FROM pgq.get_batch_events(:batch_id);

-- Retry a failed event (will reappear after the specified interval)
SELECT pgq.event_retry(:batch_id, :event_id, :retry_seconds);

-- Mark batch as done
SELECT pgq.finish_batch(:batch_id);
```

### 典型消费者循环

```sql
-- 1. Get next batch
SELECT pgq.next_batch('myqueue', 'myconsumer') AS batch_id;

-- 2. If batch_id is not NULL, get events
SELECT * FROM pgq.get_batch_events(:batch_id);

-- 3. Process events, retry failures
SELECT pgq.event_retry(:batch_id, :event_id, 60);

-- 4. Finish the batch
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

### 版本与访问边界

上游 `3.5.2` 支持 PostgreSQL 10 至 19；3.5.2 新增 PostgreSQL 19 支持，未记载 SQL API 变化。目录当前记录的软件包版本为 `3.5.1`，选择扩展升级目标前须确认已安装文件。控制文件要求由超级用户安装，且不允许迁移模式。API 使用 `pgq` 模式；应逐一检查应用角色的队列管理与消费权限。

只有 `pgq.finish_batch` 成功后，消费者位置才会前进。尽可能让事件处理与确认在同一事务内完成；外部副作用无法加入数据库事务时，应处理重复投递。
