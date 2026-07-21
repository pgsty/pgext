## 用法

来源：

- [PGMQ v1.12.0 README](https://github.com/pgmq/pgmq/blob/v1.12.0/README.md)
- [PGMQ v1.12.0 发行说明](https://github.com/pgmq/pgmq/releases/tag/v1.12.0)
- [版本 1.12.0 迁移 SQL](https://github.com/pgmq/pgmq/blob/v1.12.0/pgmq-extension/sql/pgmq--1.11.1--1.12.0.sql)
- [PGMQ 文档](https://pgmq.github.io/pgmq/)

pgmq 实现了持久化消息队列，作为 PostgreSQL 表和 SQL 函数。它支持延迟投递、可见性超时、FIFO 组、消息头、轮询、主题和归档功能。当需要将队列事务与同一数据库中的关系变化协调起来时，请使用此扩展。

### 创建队列并发送消息

    CREATE EXTENSION pgmq;
    SELECT pgmq.create('jobs');

    SELECT *
    FROM pgmq.send(
      queue_name => 'jobs',
      msg        => '{"task":"refresh"}'::jsonb,
      delay      => 0
    );

send 返回消息标识符。send_batch 可插入多个 JSONB 消息。头部可以携带路由或跟踪元数据，这些元数据与主体分开存储，并且在某些重载中支持它们。

### 使用可见性超时读取消息

    SELECT *
    FROM pgmq.read(
      queue_name => 'jobs',
      vt         => 30,
      qty        => 10
    );

读取操作会隐藏每条消息 vt 秒。成功后，可以删除或归档它：

    SELECT pgmq.delete('jobs', 42);
    SELECT pgmq.archive('jobs', 43);

如果处理失败或者消费者消失，则未确认的消息将再次可见。因此，消费者必须是幂等的；pgmq 不会保证任意应用程序副作用在全局范围内恰好执行一次。

pop 读取消息并立即删除，仅当允许在调用后丢失消息时才适用。

### FIFO 组头轮询

1.12.0 版本增加了对多个 FIFO 组头部的消息轮询：

    SELECT *
    FROM pgmq.read_grouped_head_with_poll(
      queue_name          => 'jobs',
      vt                  => 30,
      qty                 => 10,
      max_poll_seconds    => 5,
      poll_interval_ms    => 100
    );

此操作会选择组头部消息并进行轮询，直到达到最大轮询时间。这保留了每个组内的顺序，并允许不同组并发处理。

### 队列管理索引

- pgmq.create(queue_name): 创建队列和归档结构。
- pgmq.send 和 pgmq.send_batch: 入队 JSONB 消息，可选延迟。
- pgmq.read: 为可见性超时声明消息。
- pgmq.read_grouped_head_with_poll: 轮询 FIFO 组头部。
- pgmq.pop: 读取消息并立即删除。
- pgmq.delete: 通过移除消息来确认。
- pgmq.archive: 将消息移动到队列归档表中。
- pgmq.drop_queue: 移除队列对象。
- pgmq.metrics 和相关辅助函数: 检查可用时的队列深度和年龄。

对于队列作业，归档行存储在 pgmq.a_<queue_name> 中。将这些表视为由扩展管理的对象。

### 运营注意事项

- 将 vt 设定为比正常处理时间更长，并设计好超时后的重新投递。
- 队列和归档表消耗普通 PostgreSQL 的 WAL、存储、真空和备份容量。
- 归档或删除已完成的消息并强制执行归档保留策略。
- 长期轮询会占用数据库连接。根据消费者数量调整连接池大小和轮询间隔。
- 保持队列名称符合 pgmq 的标识符规则；调用 API 而不是从不可信输入构造表名。
