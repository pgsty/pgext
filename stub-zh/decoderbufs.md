

## 用法

> [decoderbufs: 使用 Protocol Buffer 格式传递 WAL 流变更的逻辑解码插件](https://github.com/debezium/postgres-decoderbufs)

一个 PostgreSQL 逻辑解码输出插件，将 WAL 变更序列化为 Protocol Buffers 格式，主要由 Debezium PostgreSQL 连接器用于变更数据捕获。

### 配置

在 `postgresql.conf` 中：

```ini
shared_preload_libraries = 'decoderbufs'
wal_level = logical
max_wal_senders = 8
max_replication_slots = 4
```

### 使用 SQL（调试模式）

```sql
-- 创建逻辑复制槽
SELECT * FROM pg_create_logical_replication_slot('decoderbufs_demo', 'decoderbufs');

-- 执行表修改
INSERT INTO my_table VALUES (1, 'test');
UPDATE my_table SET col = 'updated' WHERE id = 1;

-- 以调试文本模式查看变更
SELECT data FROM pg_logical_slot_peek_changes(
    'decoderbufs_demo', NULL, NULL, 'debug-mode', '1');

-- 消费变更
SELECT data FROM pg_logical_slot_get_changes(
    'decoderbufs_demo', NULL, NULL, 'debug-mode', '1');

-- 检查槽状态
SELECT * FROM pg_replication_slots WHERE slot_type = 'logical';
```

### 类型映射

| PostgreSQL 类型    | Protobuf 字段   |
|--------------------|------------------|
| BOOL               | datum_boolean    |
| INT2, INT4         | datum_int32      |
| INT8, OID          | datum_int64      |
| FLOAT4             | datum_float      |
| FLOAT8, NUMERIC    | datum_double     |
| CHAR, VARCHAR, TEXT | datum_string    |
| JSON, XML, UUID    | datum_string     |
| TIMESTAMP(TZ)      | datum_string     |
| BYTEA              | datum_bytes      |
| POINT, PostGIS     | datum_point      |

### 注意事项

- 对于 UPDATE/DELETE，需要适当设置 [REPLICA IDENTITY](https://www.postgresql.org/docs/current/sql-altertable.html#SQL-CREATETABLE-REPLICA-IDENTITY)
- 二进制 Protocol Buffer 输出由 Debezium Postgres Connector 消费
- `debug-mode` 选项提供人类可读的文本输出，用于 SQL 控制台测试
- 编译需要 `protobuf-c` 库和 PostGIS 开发包
