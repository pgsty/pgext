
## 用法

> [README](https://github.com/adjust/kafka_fdw)

`kafka_fdw` 是一个将 Kafka 消息暴露为 PostgreSQL 外部表的外部数据包装器。上游 README 明确提醒，这个项目尚未达到生产级就绪状态。

### 服务器与映射

先定义带 Kafka broker 列表的外部服务器，再创建用户映射：

```sql
CREATE EXTENSION kafka_fdw;

CREATE SERVER kafka_server
FOREIGN DATA WRAPPER kafka_fdw
OPTIONS (brokers 'localhost:9092');

CREATE USER MAPPING FOR PUBLIC SERVER kafka_server;
```

## 外部表

Kafka 外部表必须包含两个元数据列，一个标记为 `partition 'true'`，另一个标记为 `offset 'true'`。其余列用于描述消息负载。

### CSV 消息

```sql
CREATE FOREIGN TABLE kafka_test (
    part int OPTIONS (partition 'true'),
    offs bigint OPTIONS (offset 'true'),
    some_int int,
    some_text text,
    some_date date,
    some_time timestamp
)
SERVER kafka_server
OPTIONS (
    format 'csv',
    topic 'contrib_regress',
    batch_size '30',
    buffer_delay '100'
);
```

对于 CSV，列按位置映射。上游说明，字段校验强度取决于消息写入方，因此在数据质量不稳定时，严格解析和 junk 处理选项很重要。

### JSON 消息

```sql
CREATE FOREIGN TABLE kafka_test_json (
    part int OPTIONS (partition 'true'),
    offs bigint OPTIONS (offset 'true'),
    some_int int OPTIONS (json 'int_val'),
    some_text text OPTIONS (json 'text_val'),
    some_date date OPTIONS (json 'date_val'),
    some_time timestamp OPTIONS (json 'time_val')
)
SERVER kafka_server
OPTIONS (
    format 'json',
    topic 'contrib_regress_json',
    batch_size '30',
    buffer_delay '100'
);
```

对于 JSON，每个列都可以通过 `json` 选项映射到对象键。当前实现支持 JSON 对象，不支持顶层 JSON 数组。

## 查询与写入

偏移量列和分区列是特殊列，上游 README 建议在查询中尽可能显式指定它们：

```sql
SELECT * FROM kafka_test WHERE part = 0 AND offs > 1000 LIMIT 60;

SELECT *
FROM kafka_test
WHERE (part = 0 AND offs > 100)
   OR (part = 1 AND offs > 300)
   OR (part = 3 AND offs > 700);
```

也可以通过 `INSERT` 发送消息。如果指定了分区值，就使用该分区；否则由 Kafka 内置分区器决定：

```sql
INSERT INTO kafka_test(part, some_int, some_text)
VALUES
    (0, 5464565, 'some text goes into partition 0'),
    (NULL, 5464565, 'some text goes into partition selected by kafka');
```

## 错误处理

默认行为较为宽松：

- 缺失尾部列会视为 `NULL`
- 多余字段会被忽略
- 但无法解析的值默认仍会报错

相关表选项和辅助列包括：

- `strict 'true'`，拒绝列数不匹配
- `ignore_junk 'true'`，将格式错误的值设为 `NULL`
- 标记为 `junk 'true'` 的列，用于捕获原始负载
- 标记为 `junk_error 'true'` 的列，用于捕获解析错误

## 构建说明

该扩展使用 `librdkafka`，上游构建步骤很标准：

```bash
make && make install
make installcheck
```

测试环境假定 Kafka 运行在 `localhost:9092`，ZooKeeper 运行在 `localhost:2181`。
