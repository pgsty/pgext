## 用法

来源：

- [pg_qpubsub/README.md](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/README.md)
- [pg_qpubsub/Makefile](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/Makefile)
- [pg_qpubsub/build.sh](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/build.sh)
- [pg_qpubsub/wrappers.sql](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/wrappers.sql)
- [lib/schema/schema.sql](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/lib/schema/schema.sql)
- [pg_qpubsub/pg_qpubsub.control](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/pg_qpubsub.control)
- [Removal commit](https://github.com/smartpricing/queen/commit/cc213db9c507f6b1869a03b7ba45486506208bd6)

`pg_qpubsub` 1.0 是为 Queen 消息队列、分区与消费组提供接口的历史 SQL 扩展。上游已于 2026 年 7 月 19 日移除扩展目录。本文描述所链接的历史修订；当前 Queen 发行版不再提供此扩展。

### 核心用法

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_qpubsub;
SELECT queen.configure('orders', 60, 3, true);
SELECT queen.produce_one('orders', '{"orderId":123}'::jsonb);
SELECT * FROM queen.consume_one('orders', '__QUEUE_MODE__', 10);
```

处理消息后，将消费结果中的实际事务 ID、分区 UUID 与租约 ID 传给 `queen.commit_one`。`queen.nack` 重试消息，`queen.reject` 将消息送入死信队列，`queen.renew_one` 延长租约。具名消费组用于独立扇出，`__QUEUE_MODE__` 表示队列消费。

### 接口与运行边界

JSONB 接口 `queen.produce`、`queen.consume`、`queen.commit`、`queen.renew` 和 `queen.transaction` 支持批处理。SQL 便捷接口还包括 `queen.configure`、`queen.lag`、`queen.has_messages`、`queen.seek` 与 `queen.delete_consumer_group`。批量消费立即返回，只有标量消费接口支持轮询超时；应用通知机制仍需由应用连接。

历史 README 声明支持 PostgreSQL 14 及以上版本，依赖 `pgcrypto`。对象固定在 `queen` 模式中。扩展使用纯 SQL，不需要预加载或服务器重启；控制文件设置 `superuser=false`，安装者仍需创建对象及安装依赖所要求的权限。带版本号的安装 SQL 由同一修订中的模式、过程与包装函数组合生成。

原构建会向 PUBLIC 授予模式、全部表、序列和函数的访问权限。处理敏感消息前须审查并收紧这些授权，并结合应用故障行为设计租约、重试、保留与死信处理。在已有 Queen 模式中安装前须评估对象冲突，也不要假定该已移除扩展能直接升级为当前消息代理。本条目不宣称存在当前 Pigsty 软件包或运行支持。
