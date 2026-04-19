## 用法

来源: [official README](https://github.com/RayElg/pgmqtt/blob/main/README.md), [official repo](https://github.com/RayElg/pgmqtt)

`pgmqtt` 是一个 `pgrx` 扩展，它把 MQTT broker 嵌入 PostgreSQL，并使用 change data capture 将表变更转换为 MQTT 消息。它也支持入站 topic 映射，使 MQTT publish 可以向 PostgreSQL 表插入行。

```sql
CREATE EXTENSION pgmqtt;
```

### 出站映射

把表变更发布到 topic：

```sql
SELECT pgmqtt_add_outbound_mapping(
  'public',
  'my_table',
  'topics/{{ op | lower }}',
  '{{ columns | tojson }}'
);
```

使用这条映射后，`INSERT`、`UPDATE` 和 `DELETE` 会把 JSON 负载发布到 `topics/insert` 之类的 topic。

### 入站映射

把 MQTT publish 写入表：

```sql
SELECT pgmqtt_add_inbound_mapping(
  'sensor/{site_id}/temperature',
  'sensor_readings',
  '{"site_id": "{site_id}", "value": "$.temperature"}'::jsonb
);
```

向 `sensor/site-1/temperature` 发布 `{"temperature": 22.5}` 会向 `sensor_readings` 插入一行。

### MQTT 客户端示例

```bash
mosquitto_sub -h localhost -t 'topics/#'
mosquitto_pub -h localhost -t 'sensor/site-1/temperature' -m '{"temperature": 22.5}'
```

### 注意事项

- README 要求 `wal_level = logical`；没有 logical decoding，CDC 部分就无法工作。
- 上游文档目前只有 README 级别，因此已记录的 SQL 接口主要限于入站和出站映射工作流。
