## 用法

来源：

- [官方 README](https://github.com/edpratomo/pg_redispub/blob/e550997426a2f66accf6a28354bc6d741a92cfa2/README.md)
- [官方扩展 SQL](https://github.com/edpratomo/pg_redispub/blob/e550997426a2f66accf6a28354bc6d741a92cfa2/pg_redispub--1.0.sql)
- [官方实现](https://github.com/edpratomo/pg_redispub/blob/e550997426a2f66accf6a28354bc6d741a92cfa2/pg_redispub.c)

`pg_redispub` 1.0 从 PostgreSQL backend 发送 Redis `PUBLISH` 命令。它唯一的函数始终同步连接 `127.0.0.1:6379` 的 Redis；固定源码无法配置 endpoint、认证或 TLS。

### 核心流程

```sql
CREATE EXTENSION pg_redispub;

SELECT redispub(
  'order_events',
  '{"event":"order_paid","order_id":42}'
);
```

第一个参数是 Redis channel，第二个参数是消息。hiredis 连接或命令 context 报错时，函数发出 warning 并返回 false；否则返回 true。

### 重要对象

- `redispub(text, text)` 打开连接，执行 `PUBLISH channel message`，关闭连接并返回布尔状态。

### 投递与事务边界

Redis Pub/Sub 是瞬时机制，`redispub` 不感知事务。即使外围 PostgreSQL 事务随后回滚，消息也已立即发送；事务重试还可能重复发布。true 返回值不包含 Redis subscriber 数量，也不能证明任何 subscriber 已消费或持久化消息。需要原子、持久或可重试投递时，应使用 outbox table 和独立 delivery worker。

### 运维说明

扩展构建时必须链接 hiredis。每次调用都会新建连接，connect timeout 为 500 ms，因此逐行调用可能增加大量延迟。实现没有 Redis password、ACL username、TLS、备用 host、备用 port、pooling 或 reconnect 配置。PostgreSQL 的操作系统安全策略必须允许连接 Redis 端口；上游文档提供了 SELinux 拒绝问题的排查方法。把函数用于 trigger 或写路径前，应限制执行权限并测试 Redis 故障行为。
