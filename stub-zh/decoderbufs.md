## 用法

来源：

- [Debezium decoderbufs 3.6.0.Final README](https://github.com/debezium/postgres-decoderbufs/blob/v3.6.0.Final/README.md)
- [Output-plugin控制文件](https://github.com/debezium/postgres-decoderbufs/blob/v3.6.0.Final/decoderbufs.control)
- [Protocol Buffers模式](https://github.com/debezium/postgres-decoderbufs/blob/v3.6.0.Final/proto/pg_logicaldec.proto)

`decoderbufs` 是一个无头的PostgreSQL逻辑复制输出插件，由Debezium PostgreSQL连接器使用。它将WAL更改转换为Protocol Buffers消息；它不会创建用户SQL模式，并不需要 `CREATE EXTENSION`。

### 配置PostgreSQL

在 `postgresql.conf` 中启用该插件和逻辑复制，根据预期的消费者调整发送者和槽限制，然后重启PostgreSQL：

```conf
shared_preload_libraries = 'decoderbufs'
wal_level = logical
max_wal_senders = 8
max_replication_slots = 4
```

复制登录也需要具有 `REPLICATION` 属性，并且需要一个匹配的 `pg_hba.conf` 规则。使用适用于网络的安全认证方式，而不是README中的本地演示设置。

### 核心工作流程

创建一个输出插件为 `decoderbufs` 的逻辑槽：

```sql
SELECT *
FROM pg_create_logical_replication_slot('decoderbufs_demo', 'decoderbufs');
```

在 `psql` 中检查时，请向插件请求调试文本：

```sql
SELECT data
FROM pg_logical_slot_peek_changes(
  'decoderbufs_demo', NULL, NULL, 'debug-mode', '1'
);

SELECT data
FROM pg_logical_slot_get_changes(
  'decoderbufs_demo', NULL, NULL, 'debug-mode', '1'
);
```

`peek` 不改变已确认的位置；`get` 会将其向前推进。正常的Debezium操作将消费由 `pg_logicaldec.proto` 定义的二进制消息，而不是启用调试模式。

### 关键对象和边界

- `decoderbufs` 是创建槽时传递的逻辑复制输出插件名称。
- `debug-mode = 1` 只用于故障排除的人类可读输出。
- Protobuf 消息携带事务元数据、关系和列信息、操作类型、旧键以及带类型的值。
- 需要为 `UPDATE` 和 `DELETE` 表提供足够的数据的表需要适当的 `REPLICA IDENTITY`。

逻辑槽保留WAL直到消费者确认进度。监控 `pg_replication_slots` 并故意删除废弃的槽以防止磁盘耗尽。模式变更、复制标识符、不支持的数据类型映射和大事务应与匹配的Debezium连接器版本一起进行测试。

上游构建需要PostgreSQL 9.6或更高版本以及protobuf-c；当可用时，会编译PostGIS支持。包发布跟随Debezium到3.6.0.Final，而插件控制元数据保持SQL版本0.1.0，因为这是一个输出插件而不是基于迁移的SQL扩展。
