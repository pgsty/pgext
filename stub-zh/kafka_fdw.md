

## 用法

> [kafka_fdw: CSV 格式消息的 Kafka 外部数据包装器](https://github.com/adjust/kafka_fdw)

### 创建服务器

```sql
CREATE EXTENSION kafka_fdw;

CREATE SERVER kafka_server FOREIGN DATA WRAPPER kafka_fdw
  OPTIONS (brokers 'localhost:9092');
```

**服务器选项：** `brokers`（必填，逗号分隔的 Kafka broker 端点）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR PUBLIC SERVER kafka_server;
```

### 创建外部表（CSV 格式）

```sql
CREATE FOREIGN TABLE kafka_csv (
  part int OPTIONS (partition 'true'),
  offs bigint OPTIONS (offset 'true'),
  some_int int,
  some_text text,
  some_date date,
  some_time timestamp
)
SERVER kafka_server
OPTIONS (format 'csv', topic 'my_topic', batch_size '30', buffer_delay '100');
```

需要两个元数据列：一个带 `partition 'true'`，一个带 `offset 'true'`。其余列与消息格式匹配。

**表选项：** `format`（`csv` 或 `json`）、`topic`（Kafka 主题名）、`batch_size`、`buffer_delay`（毫秒）、`strict`（强制严格模式验证）、`ignore_junk`（将格式错误的列设为 NULL）。

### 创建外部表（JSON 格式）

```sql
CREATE FOREIGN TABLE kafka_json (
  part int OPTIONS (partition 'true'),
  offs bigint OPTIONS (offset 'true'),
  some_int int OPTIONS (json 'int_val'),
  some_text text OPTIONS (json 'text_val')
)
SERVER kafka_server
OPTIONS (format 'json', topic 'my_json_topic', batch_size '30', buffer_delay '100');
```

使用 `json` 列选项将列名映射到 JSON 键。

### 消费消息

```sql
-- 从特定分区和偏移量读取
SELECT * FROM kafka_csv WHERE part = 0 AND offs > 1000 LIMIT 60;

-- 从多个分区读取
SELECT * FROM kafka_csv
WHERE (part = 0 AND offs > 100) OR (part = 1 AND offs > 300);
```

注意：`offset` 在 SQL 中是保留关键字；在某些上下文中引用偏移量列时使用双引号。

### 生产消息

```sql
-- 指定分区插入
INSERT INTO kafka_csv (part, some_int, some_text)
  VALUES (0, 42, 'hello from partition 0');

-- 自动分区选择插入
INSERT INTO kafka_csv (some_int, some_text)
  VALUES (42, 'auto-partitioned message');
```
