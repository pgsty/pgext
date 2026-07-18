


## 用法

来源：[README](https://github.com/RayElg/pgmqtt/blob/0.3.0/README.md)、[interfaces](https://github.com/RayElg/pgmqtt/blob/0.3.0/docs/interfaces.md)、[configuration](https://github.com/RayElg/pgmqtt/blob/0.3.0/docs/configuration.md)、[limitations](https://github.com/RayElg/pgmqtt/blob/0.3.0/docs/limitations.md)、[Cargo.toml](https://github.com/RayElg/pgmqtt/blob/0.3.0/extension/Cargo.toml)

`pgmqtt` 是一个 `pgrx` 扩展，会在 PostgreSQL 中嵌入 MQTT 代理，并使用变更数据捕获将表变更转为 MQTT 消息。它也支持入站主题映射，让 MQTT 发布操作可以向 PostgreSQL 表插入行。

```sql
CREATE EXTENSION pgmqtt;
```

### 出站映射

将表变更发布到主题：

```sql
SELECT pgmqtt_add_outbound_mapping(
  'public',
  'my_table',
  'topics/{{ op | lower }}',
  '{{ columns | tojson }}',
  1
);
```

有了该映射后，`INSERT`、`UPDATE` 和 `DELETE` 会把 JSON 载荷发布到 `topics/insert` 等主题。文档化函数签名还接受可选的 `qos integer DEFAULT 0` 和 `template_type text DEFAULT 'jinja2'`。

### 入站映射

通过 MQTT 发布操作插入行：

```sql
SELECT pgmqtt_add_inbound_mapping(
  'sensor/{site_id}/temperature',
  'sensor_readings',
  '{"site_id": "{site_id}", "value": "$.temperature"}'::jsonb
);
```

向 `sensor/site-1/temperature` 发布 `{"temperature": 22.5}` 会向 `sensor_readings` 插入一行。

入站映射也可通过传入 `op`、`conflict_columns`、`target_schema`、`mapping_name` 和 `template_type` 执行 `upsert` 与 `delete` 操作。主题模式使用 `{variable}` 捕获；JSON 载荷字段使用 `$.temperature`、`$payload`、`$topic` 等表达式。

### 查看和删除映射

```sql
SELECT * FROM pgmqtt_list_outbound_mappings();
SELECT pgmqtt_remove_outbound_mapping('public', 'my_table');

SELECT * FROM pgmqtt_list_inbound_mappings();
SELECT pgmqtt_remove_inbound_mapping('temp_readings');

SELECT * FROM pgmqtt_status();
```

`pgmqtt_status()` 报告活跃连接、订阅、保留消息、待处理会话消息、CDC 映射、CDC 复制槽状态、入站映射、待处理入站写入和死信。

版本 0.3.0 增加了异步管理命令：

```sql
SELECT pgmqtt_disconnect_client('device-42');
SELECT pgmqtt_disconnect_role('mqtt_devices');
SELECT pgmqtt_reload_acls('*');
```

后台工作进程会处理 `pgmqtt_admin_commands`；重新加载 ACL 时会移除不再允许的订阅。

### MQTT 客户端示例

```bash
mosquitto_sub -h localhost -t 'topics/#'
mosquitto_pub -h localhost -t 'sensor/site-1/temperature' -m '{"temperature": 22.5}'
```

### 配置

文档化 GUC 位于 `pgmqtt` 命名空间下：

```sql
ALTER SYSTEM SET pgmqtt.cdc_every_n_ticks = 16;
SELECT pg_reload_conf();
```

监听器 GUC 包括 `pgmqtt.mqtt_enabled`、`pgmqtt.mqtt_port`（`1883`）、`pgmqtt.ws_enabled`、`pgmqtt.ws_port`（`9001`）、`pgmqtt.mqtts_enabled`、`pgmqtt.mqtts_port`（`8883`）、`pgmqtt.wss_enabled`、`pgmqtt.wss_port`（`9002`）。TLS 和认证设置包括 `pgmqtt.tls_cert_file`、`pgmqtt.tls_key_file`、`pgmqtt.license_key`、`pgmqtt.jwt_public_key`、`pgmqtt.jwt_required`、`pgmqtt.jwt_required_ws`、`pgmqtt.password_auth_enabled`、`pgmqtt.password_auth_required` 和 `pgmqtt.password_auth_role_filter`。

性能和可观测性 GUC 包括 `pgmqtt.tick_interval_ms`、`pgmqtt.max_client_buffer_bytes`、`pgmqtt.cdc_every_n_ticks`、`pgmqtt.debug_log`、`pgmqtt.metrics_snapshot_interval`、`pgmqtt.metrics_retention_days`、`pgmqtt.metrics_connections_cache_interval`、`pgmqtt.metrics_hook_function`、`pgmqtt.metrics_notify_channel`。监听器和 TLS 设置在 MQTT 后台工作进程启动时读取，因此需要重启工作进程，而不是只执行 `pg_reload_conf()`。

### 注意事项

- README 要求 `wal_level = logical`；没有逻辑解码时 CDC 侧无法工作。
- 本项目 CSV 跟踪版本 `0.3.0`、PostgreSQL 14-18，以及软件包侧 `pgrx` `0.18.1` 重构说明。
- 文档记录支持 MQTT 5.0 和 MQTT 3.1.1 客户端。QoS 0 和 QoS 1 受支持；QoS 2 未实现，请求 QoS 2 的订阅会降级到 QoS 1。
- CDC 捕获 `INSERT`、`UPDATE` 和 `DELETE`；不捕获 DDL 变更和 `TRUNCATE`。`DELETE` 可能需要 `REPLICA IDENTITY FULL`。
- 版本 0.3.0 增加 PostgreSQL 角色/密码认证、按主题 ACL、管理员断开连接/重新加载命令、指标与可观测性表和函数、MQTT 3.1.1 协议测试、UNLOGGED 指标/缓存恢复测试，以及额外的流量控制限制。
