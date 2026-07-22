## 用法

来源：

- [官方输出插件实现](https://github.com/rektide/putpostlogic/blob/0d71733aa77085b4f62c49c0bba972c241d0fcbc/putpostlogic.c)
- [官方变更解码器](https://github.com/rektide/putpostlogic/blob/0d71733aa77085b4f62c49c0bba972c241d0fcbc/decode.c)
- [官方扩展控制文件](https://github.com/rektide/putpostlogic/blob/0d71733aa77085b4f62c49c0bba972c241d0fcbc/putpostlogic.control)

`putpostlogic` 是一个历史逻辑解码输出插件，把行变更按事务组成 JSON，并可额外把 JSON 发送到 Kafka 或 nanomsg 套接字。它是无 DDL 插件：固定提交中没有扩展 SQL 脚本，通过逻辑复制槽加载，而不是执行 `CREATE EXTENSION`。

### 历史调用方式

服务器必须启用逻辑解码，并安装插件所需的 json-c、librdkafka 与 nanomsg 库。在兼容的旧 PostgreSQL 构建上，复制槽调用形式如下：

```sql
SELECT *
FROM pg_create_logical_replication_slot(
  'putpostlogic_demo',
  'putpostlogic'
);

SELECT *
FROM pg_logical_slot_get_changes(
  'putpostlogic_demo', NULL, NULL,
  'include-transaction', 'true',
  'json-topic', 'database_changes',
  'kafka-brokers', 'localhost:9092',
  'kafka-topic', 'postgres_changes'
);
```

本次复核源码识别的选项包括 `include-transaction`、`json-topic`、`kafka-brokers`、`kafka-topic`、`nanomsg-pair`、`nanomsg-pair-bind`、`nanomsg-pub`、`nanomsg-pub-topic`、`nanomsg-pub-bind`、`nanomsg-push` 和 `nanomsg-push-bind`。

### 输出与行标识

插件在每次提交时输出一个 JSON 对象，其中包含编号变更成员、关系标识、插入/更新后的新值，以及更新/删除使用的 `@where` 行标识对象。正确的旧行标识依赖主键/副本标识索引或 `REPLICA IDENTITY FULL`；未变化的 TOAST 值使用哨兵字符串，而非原值。

### 关键边界

该固定提交是未完成的 2014 年前后源码，存在多条不安全路径，包括未检查或长度不足的内存分配、NULL 值处理缺陷、无效的副本索引下标，以及未配置 Kafka 时仍无条件轮询 Kafka。它还依赖过时的逻辑解码 API，在现代 PostgreSQL 上可能无法构建或安全运行。没有完成源码审计、修复、清理器测试和端到端崩溃测试前，不得部署此构建。

Kafka 与 nanomsg 发送是解码过程的外部副作用，不具备事务性交付语义。重复读取复制槽可能重复消息，故障会让复制槽输出与消息代理状态分叉，解码行也可能泄露秘密或个人数据。该示例只用于说明接口；生产环境应采用受维护的逻辑解码/CDC 系统，并明确保证交付、重试、顺序、模式演进、访问控制和复制槽保留策略。
