## 用法

来源：

- [上游来源 1](https://github.com/justusepostgres/pg_broker/blob/18f080f57de5e8c2bff310f93bc3366a1f0e7ade/README.md)
- [上游来源 2](https://github.com/justusepostgres/pg_broker/blob/18f080f57de5e8c2bff310f93bc3366a1f0e7ade/crates/extension/pg_broker.control)

`pg_broker` 的 SQL 版本 0.1.1 为独立 Kafka 协议网关提供持久化分区事件日志。

### 核心流程

需部署相互匹配的网关、二进制及 SQL 文件。网关通过 `public.broker_call(jsonb)` 创建主题、追加和读取记录、管理组偏移量。`broker` 模式保存留存记录与协调状态，消费记录不会删除它。

```sql
CREATE EXTENSION pg_broker;
SELECT extversion FROM pg_extension WHERE extname = 'pg_broker';
```

### 运行边界

受支持的构建面向 PostgreSQL 17，安装需要特权角色。仅向网关授予文档规定的入口权限。生产成功确认表示 PostgreSQL 已提交，并不代表 Kafka 副本确认。此 MVP 仅实现有限的 Kafka API，不提供身份认证、TLS、幂等生产者或 Kafka 事务保证。留存策略可能移除落后消费组所需的数据，应单独监控。
