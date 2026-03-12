


## 用法

> [pgmb: 构建在 PostgreSQL 内部的轻量级消息代理系统](https://github.com/fraruiz/pgmb)

`pgmb` 扩展提供数据库内置的消息代理，支持基于 HTTP 的工作进程分发、自动重试、死信队列和基于模式的路由。

```sql
CREATE EXTENSION pgmb;  -- 需要 pg_cron 和 http 扩展
```

### 注册工作进程

```sql
SELECT pgmb.worker(
    'order_processor',                     -- 工作进程名称
    'http://localhost:8080/process',       -- 端点 URL
    100                                    -- 每秒请求数限制
);
-- 返回: 工作进程 UUID
```

### 创建队列

```sql
SELECT pgmb.create(
    'order_queue',                         -- 队列名称
    'order.*',                             -- 绑定键模式（支持 * 通配符）
    5,                                     -- 最大重试次数
    '550e8400-e29b-41d4-a716-446655440000' -- 工作进程 UUID
);
-- 返回: 队列 UUID
```

### 发送消息

```sql
-- 简单消息
SELECT pgmb.send(
    gen_random_uuid(),
    'order.created',
    '{"order_id": 123, "amount": 45.67}'::jsonb
);

-- 带自定义头部
SELECT pgmb.send(
    gen_random_uuid(),
    'order.created',
    '{"order_id": 123}'::jsonb,
    '{"source": "web", "priority": "high"}'::jsonb
);

-- 延迟消息（按时间戳或秒数）
SELECT pgmb.send(
    gen_random_uuid(),
    'order.created',
    '{"order_id": 123}'::jsonb,
    '{}'::jsonb,
    now() + interval '10 minutes'
);
```

### API 参考

| 函数 | 描述 |
|------|------|
| `pgmb.worker(name, endpoint, rps)` | 注册 HTTP 工作进程端点 |
| `pgmb.create(name, binding_key, max_retries, worker_id)` | 创建带路由模式的队列 |
| `pgmb.send(id, routing_key, body)` | 发送消息 |
| `pgmb.send(id, routing_key, body, headers)` | 发送带头部的消息 |
| `pgmb.send(id, routing_key, body, headers, delay)` | 发送延迟消息 |

### 工作原理

1. 消息通过 `pgmb.send()` 插入到 `pgmb.messages` 表
2. 触发器根据路由键模式将消息路由到匹配的队列
3. `pg_cron` 每秒通过 HTTP POST 将消息分发到工作进程端点
4. 失败的消息会被重试；超过最大重试次数后移入死信队列

### 监控

```sql
SELECT * FROM pgmb.workers;
SELECT * FROM pgmb.queues;
SELECT COUNT(*) FROM pgmb.order_queue WHERE acknoledge = false;
SELECT * FROM pgmb.order_dead_letter_queue;
```
