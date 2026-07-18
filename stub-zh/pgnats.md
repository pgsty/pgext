## 用法

来源：

- [上游 README](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/README.md)
- [安装指南](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/INSTALL.md)
- [扩展 control 文件](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/pgnats.control)
- [扩展 Cargo 清单](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/Cargo.toml)
- [1.1.0 版本升级 SQL](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/sql/pgnats--1.0.0--1.1.0.sql)

`pgnats` 通过 SQL 函数连接 NATS 服务器，提供发布或请求消息、JetStream、键值存储、对象存储和可选主题订阅。使用 `nats_subscribe` 与 `nats_unsubscribe` 的后台工作进程功能时明确要求预加载：

```conf
shared_preload_libraries = 'pgnats.so'
```

重启 PostgreSQL，创建扩展及数据库中唯一允许的 NATS 外部服务器，再调用消息 API：

```sql
CREATE EXTENSION pgnats;

CREATE SERVER nats_fdw_server
  FOREIGN DATA WRAPPER pgnats_fdw
  OPTIONS (host 'nats.example.net', port '4222', capacity '128');

SELECT nats_publish_jsonb(
  'events.order.created',
  '{"id":42}'::jsonb
);

SELECT nats_request_text('service.echo', 'ping', 1000);
```

### 构建与权限边界

版本 `1.1.0` 的 Cargo 特性覆盖 PostgreSQL 14 至 18。构建特性决定订阅、键值和对象存储 API 是否存在；默认构建会启用三者。订阅回调必须接收一个 `bytea` 参数。这些 SQL 函数能够发送数据、调用回调和访问外部 NATS 存储，因此应限制 `EXECUTE`、使用 NATS 端 ACL，并在需要时配置 TLS 证书路径。扩展强制每个数据库最多只能有一个 `pgnats_fdw` 服务器。
