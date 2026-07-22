## 用法

来源：

- [官方 README](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/README.md)
- [官方用法文档](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/doc/amqp.md)
- [官方控制文件](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/amqp.control)

`amqp` 让 PostgreSQL 后端可以直接向 AMQP 代理发布消息。它适合简单的数据库触发通知，但发布属于外部副作用，不会随外围数据库事务回滚。

### 核心流程

把 `pg_amqp` 加入 `shared_preload_libraries`，重启 PostgreSQL，创建扩展，再按照已安装版本随附的表定义把代理主机、端口、用户与密码写入 `amqp.broker`。然后用返回的代理 ID 发布消息：

```sql
CREATE EXTENSION amqp;

SELECT amqp.publish(
    1,
    'amqp.direct',
    'jobs.created',
    '{"job_id":42}',
    2,
    'application/json',
    NULL,
    'job-42'
);

SELECT amqp.disconnect(1);
```

`amqp.publish` 最后四个参数是可选消息属性：投递模式、内容类型、回复地址与关联 ID。投递模式 1 表示非持久，2 请求持久投递；端到端保证仍由代理持久性与消费者确认共同决定。

### 重要对象

- `amqp.broker` 保存代理连接信息，包括凭据。
- `amqp.publish` 向交换机发送带路由键与载荷的消息。
- `amqp.disconnect` 关闭当前后端针对某个代理的连接；后端退出时会关闭其余代理连接。

### 运维注意事项

严格限制 `amqp.broker` 与发布函数的所有权限，避免凭据经转储、副本、日志或宽泛只读角色泄露。网络调用在数据库后端中执行，因此代理延迟或故障会拉长 SQL 延迟并占用连接。

SQL 成功提交并不能证明消息恰好投递一次。消费者应实现幂等，消息中应带持久的应用标识；需要数据库状态与消息原子协调时应采用 outbox 工作进程。上游项目与 API 已较老，因此必须针对选定代理与 PostgreSQL 版本直接测试兼容性、TLS/认证能力、重连行为及失败语义。
