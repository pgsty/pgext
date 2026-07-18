## 用法

来源：

- [上游 README](https://github.com/rpdelaney/pg-message-queue/blob/05309d14ae8a61b958ba3d56d2a1bccf53739bf0/README.md)
- [0.2.1 版安装 SQL](https://github.com/rpdelaney/pg-message-queue/blob/05309d14ae8a61b958ba3d56d2a1bccf53739bf0/pg_message_queue--0.2.1.sql)
- [安全说明](https://github.com/rpdelaney/pg-message-queue/blob/05309d14ae8a61b958ba3d56d2a1bccf53739bf0/doc/security)

pg_message_queue 使用 PostgreSQL 表、PL/pgSQL、触发器和 LISTEN/NOTIFY 实现小型持久化的频道队列。每个频道拥有一张队列表；通知只携带消息标识符，载荷保留在表中，直到消费者取回。

### 基本流程

0.2.1 版可为每个队列接受一个 PostgreSQL 注册类型，包括 jsonb：

```sql
CREATE EXTENSION pg_message_queue;

SELECT pg_mq_create_queue('jobs', 'jsonb'::regtype);
LISTEN jobs;

SELECT pg_mq_send_message(
    'jobs',
    '{"task":"reindex","relation":"public.orders"}'::jsonb
);

BEGIN;
SELECT * FROM pg_mq_get_msg_text('jobs', 10);
COMMIT;
```

取回操作会在当前事务中将行标记为已投递。回滚后，这些行会重新可用。应用可以监听唤醒通知，也可以分批轮询；LISTEN 不是必需的。

队列创建是一项安全定义者操作，安装脚本会撤销 PUBLIC 的执行权。发送与接收还取决于生成队列表上的权限，因此只应向生产者或消费者授予其必需操作。

### 注意事项

- 投递具有事务性，但不是恰好一次的分布式协议。消费者必须使外部副作用幂等，并决定如何恢复被遗弃的工作。
- 取回函数会更新队列行，并可能在并发消费中阻塞。该设计适合中小规模集成队列，不是高吞吐消息中间件的替代品。
- PostgreSQL 通知是唤醒信号，而不是受保护的载荷投递。数据库用户通常都能监听频道；应保护底层表，不要在通知载荷中放入密密。
- 已发布的 0.2.1 SQL 用未声明的 t_table_name 引用定义 pg_mq_rebuild_triggers。在修补并测试前，不应依赖该恢复辅助函数。
- 虽然队列标识符以 bigint 为基础，部分内部函数签名仍使用 integer 消息标识符。对长期运行的队列，应测试溢出与保留行为。
