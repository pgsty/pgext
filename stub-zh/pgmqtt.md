## 用法
- GitHub 仓库: [`RayElg/pgmqtt`](https://github.com/RayElg/pgmqtt)
- README: [RayElg/pgmqtt/blob/main/README.md](https://github.com/RayElg/pgmqtt/blob/main/README.md)

`pgmqtt` 是一个基于 `pgrx` 的 PostgreSQL 扩展，可以把 CDC 变更接入 MQTT broker。数据库变更可以通过 SQL 定义的映射发布到 MQTT topic，MQTT 发布也可以写回 PostgreSQL 表。

README 的快速上手要求将 `wal_level` 设为 `logical`，这样才能正确捕获 CDC 事件。

### 出站映射

```sql
SELECT pgmqtt_add_outbound_mapping(
    'public',
    'my_table',
    'topics/{{ op | lower }}',
    '{{ columns | tojson }}'
);
```

使用这条映射后，`INSERT`、`UPDATE` 和 `DELETE` 对该表的变更会作为 MQTT 消息发布出去。订阅 `topics/insert`、`topics/update` 或 `topics/delete` 的客户端会收到 JSON 负载。

### 入站映射

```sql
SELECT pgmqtt_add_inbound_mapping(
    'sensor/{site_id}/temperature',
    'sensor_readings',
    '{"site_id": "{site_id}", "value": "$.temperature"}'::jsonb
);
```

当客户端向 `sensor/site-1/temperature` 发布负载 `{"temperature": 22.5}` 时，README 说明会向 `sensor_readings` 插入一行，并写入提取出的字段值。

### 客户端示例

```bash
mosquitto_sub -h localhost -t 'topics/#'
mosquitto_pub -h localhost -t 'sensor/site-1/temperature' -m '{"temperature": 22.5}'
```

### 范围

上游 README 已覆盖 broker 模型、出站/入站映射示例以及基础 MQTT 客户端用法；它没有单独的项目主页文档，因此这个 stub 仅保留 README 范围。
