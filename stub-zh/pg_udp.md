## 用法

来源：

- [官方 README](https://github.com/zilder/pg_udp/blob/ab0f84f8820996bf58eb597eeb5c299476f9137d/README.md)
- [扩展 SQL](https://github.com/zilder/pg_udp/blob/ab0f84f8820996bf58eb597eeb5c299476f9137d/pg_udp--0.1.sql)
- [UDP 实现](https://github.com/zilder/pg_udp/blob/ab0f84f8820996bf58eb597eeb5c299476f9137d/pg_udp.c)

`pg_udp` 0.1 暴露一个 C 函数，从 PostgreSQL 后端发送 UDP 数据报。它可以向可信主机发送小型尽力通知，但不提供响应、送达确认、重试、队列或事务集成。

### 核心流程

```sql
CREATE EXTENSION pg_udp;

SELECT udp_send('127.0.0.1', 9999, 'cache-invalidate:customer-42');
```

唯一对象是 `udp_send(host cstring, port int, data cstring) RETURNS void`。它通过 `gethostbyname` 解析 IPv4 主机，打开数据报套接字，再调用 `sendto` 发送首个 NUL 字节之前的数据。

### 安全与可靠性边界

UDP 会丢包且不保证顺序。SQL 调用成功不能证明接收方已收到或处理消息，之后的事务回滚也无法撤回已经发送的数据报。绝不能把它作为审计、复制、支付或缓存一致性的唯一机制。

固定版本的实现会在未检查 DNS 失败时解引用解析结果，不检查 `sendto` 返回值，也不关闭套接字。无效主机可能导致后端崩溃，发送错误可能被静默忽略，重复调用则可能耗尽文件描述符。它还使用旧式、仅 IPv4 的解析 API 与 `cstring` 伪类型参数。应限制执行权限和网络出口，只接受可信的固定主机、端口和数据；生产负载应改用维护中的事务发件箱或消息客户端。
