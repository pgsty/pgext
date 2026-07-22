## 用法

来源：

- [官方 README](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/README.md)
- [0.1.0 版扩展 SQL](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/sql/pg_consul--0.1.0.sql)
- [官方 Consul KV 说明](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/doc/consul-kv.md)
- [C 实现与 GUC 定义](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/src/pg_consul.cpp)

`pg_consul` 让 PostgreSQL 后端向 Consul agent 发起同步 HTTP 请求。`0.1.0` 版可以探测 agent、读取 leader 和 peer 状态以及获取 KV 条目；命令行工具文档中的 KV 写入、删除、CAS 或加锁操作并未作为 SQL 函数暴露。

### 核心流程

安装共享库及其 cURL 依赖后，创建扩展，并把当前会话指向可访问的 Consul agent。

```sql
CREATE EXTENSION pg_consul;

SET consul.agent_host = '127.0.0.1';
SET consul.agent_port = 8500;
SET consul.agent_timeout = 1000;

SELECT consul_agent_ping();
SELECT consul_status_leader();
SELECT * FROM consul_status_peers();
```

可以读取单个键，或在第二个参数传入 `true` 以递归读取键前缀。可选的第三个参数选择 Consul 数据中心。

```sql
SELECT *
FROM consul_kv_get('app/config', false, NULL);

SELECT *
FROM consul_kv_get('services/', true, 'dc1');
```

### 重要对象

- `consul_agent_ping()`：探测已配置的 agent。
- `consul_agent_ping(text, integer)`：探测显式指定的主机和端口。
- `consul_status_leader()`：以文本返回 Consul leader 端点。
- `consul_status_peers()`：返回 peer 的 `host`、`port` 与 `leader` 状态行。
- `consul_kv_get(text, boolean, text)`：返回键、解码后的值、flags、create/modify/lock 索引和 session 信息。
- `consul.agent_host`：用户可设置的 agent 主机，默认 `127.0.0.1`。
- `consul.agent_port`：用户可设置的 agent 端口，默认 `8500`。
- `consul.agent_timeout`：用户可设置的请求超时毫秒数，默认 `1000`。

### 安全与故障行为

每次调用都在 PostgreSQL 后端内运行并等待网络请求。因此超时、异常响应或 agent 不可用都可能让 SQL 语句失败或延迟。应限制超时；未经测试，不要把这些调用放在延迟敏感的查询路径上。

公开接口没有用于 Consul ACL token、HTTPS CA 或客户端证书的 SQL 参数或 GUC。只能连接可信本地 agent，或使用扩展之外已有保护的端点；同时应限制 EXECUTE 权限，因为 KV 值与集群拓扑可能敏感。README 标题写作 `0.10.0`，而 control 文件、扩展 SQL 和 PGXN 元数据都将本次复核版本标为 `0.1.0`。
