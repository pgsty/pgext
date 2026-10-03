## 用法

来源：

- [contrib/alohadb_queue/alohadb_queue.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_queue/alohadb_queue.control)
- [contrib/alohadb_queue/alohadb_queue--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_queue/alohadb_queue--1.0.sql)
- [contrib/alohadb_queue/queue_ops.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_queue/queue_ops.c)

`alohadb_queue` 1.0 在所引用的 AlohaDB 源码中提供消息队列与消费组，提供显式确认，并存储可见性时间戳。

### 核心工作流

```sql
CREATE EXTENSION alohadb_queue;
SELECT queue_create('work');
SELECT queue_send('work', '{"task":"example"}'::jsonb);
SELECT * FROM queue_receive('work', 1);
```

### 对象与维护

应用处理成功后，用返回的消息 ID 调用 `queue_ack(queue_name, msg_id)`。`queue_nack` 释放一次投递；`queue_send_batch` 提交多条消息。`queue_subscribe`、`queue_poll` 与 `queue_commit_offset` 管理消费组偏移。`queue_stats` 返回队列状态；`queue_purge` 删除队列中的全部消息，`queue_drop` 删除队列。

扩展在固定的 `public` 模式中创建 `alohadb_queue_queues`、`alohadb_queue_messages` 与 `alohadb_queue_consumers`。安装会创建 C 函数，需要超级用户权限；所核验源码不要求预加载。所引用接收路径将消息改为已投递，后续接收只选择就绪消息；时间戳到期本身不会将该投递重新变为就绪。应使用 `queue_nack` 显式恢复，并使消费者能幂等处理这种重复投递。保留时长会被保存，但所核验 API 提供的是全量清空，并非自动按消息年龄清理。这些说明仅针对 AlohaDB 源码，不表示兼容普通 PostgreSQL。
