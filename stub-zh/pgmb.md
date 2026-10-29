## 用法

来源：
- [README](https://github.com/fraruiz/pgmb/blob/release/README.md)
- [Control](https://github.com/fraruiz/pgmb/blob/release/pgmb.control)
- [SQL](https://github.com/fraruiz/pgmb/blob/release/pgmb--1.0.0.sql)
- [META.json](https://github.com/fraruiz/pgmb/blob/release/META.json)

`pgmb` 是一个 SQL 消息代理，提供工作端点、路由键模式、队列，以及通过 `pg_cron` 调度的 HTTP 投递。上游仓库已归档。

### 安装与兼容性

RPM 修订版 `1.0.0-2PGSTY` 将 `pgsql_http_N` 声明为必需依赖，并移除了重复创建 schema 的语句。较早的 RPM 与现有 DEB 包仍使用上游 SQL，control 文件与安装脚本都会创建 schema，因而可能在 `CREATE EXTENSION pgmb` 时失败。软件包安装成功不代表扩展创建成功。

使用超级用户安装。必需扩展为 `pg_cron` 和 `http`；`pgmb` 本身不需要预加载。将调度器配置为扩展所在数据库：

```ini
shared_preload_libraries = 'pg_cron'
cron.database_name = 'postgres'
```

重启 PostgreSQL，连接该数据库，先创建依赖，再创建扩展：

```sql
CREATE EXTENSION pg_cron;
CREATE EXTENSION http;
CREATE EXTENSION pgmb;
```

### 注册队列并发送消息

工作端点接收来自数据库服务器的 HTTP 请求。使用可信端点，并限制可注册工作端点的角色。

```sql
SELECT pgmb.create(
    'orders', 'order.*', 3,
    pgmb.worker('order_worker', 'http://127.0.0.1:8080/worker', 100)
);

SELECT pgmb.send(
    gen_random_uuid(), 'order.created', '{"order_id": 123}'::jsonb
);

SELECT * FROM pgmb.workers;
SELECT * FROM pgmb.queues;
SELECT * FROM pgmb.orders_queue;
SELECT * FROM pgmb.orders_dead_letter_queue;
```

`pgmb.create` 的参数顺序是名称、绑定模式、最大重试次数、工作端点 UUID。触发器将消息路由到匹配的队列，创建队列时会注册调度任务。

### API

| 函数 | 用途 |
|---|---|
| `pgmb.worker(name, endpoint, rps)` | 注册工作端点并返回 UUID；rps 限制每次投递的批量大小。 |
| `pgmb.create(name, binding_key, max_retries, worker_id)` | 创建队列、死信表和调度任务，返回队列 UUID。 |
| `pgmb.send(id, routing_key, body)` | 保存 JSONB 消息，并路由到匹配的队列。 |
| `pgmb.send` 重载 | 接受可选的消息头，以及整数秒数或时间戳形式的延迟。 |
| `pgmb.dispatch_messages(queue_id)` | 批量投递；成功的 HTTP 响应标记为已确认，失败则重试，达到限制后移入死信表。 |

### 1.0.0 版本限制

- 延迟参数会填写 `enqueued_at`，但投递函数不排除未来时间戳，因此这些重载不能保证延迟投递。
- 消息头会随消息保存，但投递函数仅通过 `http_post` 发送消息体，不会转发为自定义 HTTP 请求头。
- 重试在后续投递时使用计数器；SQL 没有实现指数退避。
- HTTP 操作的副作用不受数据库事务控制。使用这个已归档项目之前，应核对投递与数据保留需求。
