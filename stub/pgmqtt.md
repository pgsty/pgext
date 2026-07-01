


## Usage

Sources: [README](https://github.com/RayElg/pgmqtt/blob/0.3.0/README.md), [interfaces](https://github.com/RayElg/pgmqtt/blob/0.3.0/docs/interfaces.md), [configuration](https://github.com/RayElg/pgmqtt/blob/0.3.0/docs/configuration.md), [limitations](https://github.com/RayElg/pgmqtt/blob/0.3.0/docs/limitations.md), [Cargo.toml](https://github.com/RayElg/pgmqtt/blob/0.3.0/extension/Cargo.toml)

`pgmqtt` is a `pgrx` extension that embeds an MQTT broker into PostgreSQL and uses change data capture to turn table changes into MQTT messages. It also supports inbound topic mappings so MQTT publishes can insert rows into PostgreSQL tables.

```sql
CREATE EXTENSION pgmqtt;
```

### Outbound Mapping

Publish table changes to topics:

```sql
SELECT pgmqtt_add_outbound_mapping(
  'public',
  'my_table',
  'topics/{{ op | lower }}',
  '{{ columns | tojson }}',
  1
);
```

With that mapping, `INSERT`, `UPDATE`, and `DELETE` publish JSON payloads to topics such as `topics/insert`. The documented function signature also accepts optional `qos integer DEFAULT 0` and `template_type text DEFAULT 'jinja2'`.

### Inbound Mapping

Insert rows from MQTT publishes:

```sql
SELECT pgmqtt_add_inbound_mapping(
  'sensor/{site_id}/temperature',
  'sensor_readings',
  '{"site_id": "{site_id}", "value": "$.temperature"}'::jsonb
);
```

Publishing `{"temperature": 22.5}` to `sensor/site-1/temperature` inserts a row into `sensor_readings`.

Inbound mappings can also perform `upsert` and `delete` operations by passing `op`, `conflict_columns`, `target_schema`, `mapping_name`, and `template_type`. Topic patterns use `{variable}` captures; JSON payload fields use expressions such as `$.temperature`, `$payload`, and `$topic`.

### Inspect and Remove Mappings

```sql
SELECT * FROM pgmqtt_list_outbound_mappings();
SELECT pgmqtt_remove_outbound_mapping('public', 'my_table');

SELECT * FROM pgmqtt_list_inbound_mappings();
SELECT pgmqtt_remove_inbound_mapping('temp_readings');

SELECT * FROM pgmqtt_status();
```

`pgmqtt_status()` reports active connections, subscriptions, retained messages, pending session messages, CDC mappings, CDC slot state, inbound mappings, pending inbound writes, and dead letters.

Version 0.3.0 adds asynchronous admin commands:

```sql
SELECT pgmqtt_disconnect_client('device-42');
SELECT pgmqtt_disconnect_role('mqtt_devices');
SELECT pgmqtt_reload_acls('*');
```

The background worker drains `pgmqtt_admin_commands`; ACL reloads prune subscriptions that are no longer allowed.

### MQTT Client Examples

```bash
mosquitto_sub -h localhost -t 'topics/#'
mosquitto_pub -h localhost -t 'sensor/site-1/temperature' -m '{"temperature": 22.5}'
```

### Configuration

The documented GUCs live under the `pgmqtt` namespace:

```sql
ALTER SYSTEM SET pgmqtt.cdc_every_n_ticks = 16;
SELECT pg_reload_conf();
```

Listener GUCs include `pgmqtt.mqtt_enabled`, `pgmqtt.mqtt_port` (`1883`), `pgmqtt.ws_enabled`, `pgmqtt.ws_port` (`9001`), `pgmqtt.mqtts_enabled`, `pgmqtt.mqtts_port` (`8883`), `pgmqtt.wss_enabled`, and `pgmqtt.wss_port` (`9002`). TLS and authentication settings include `pgmqtt.tls_cert_file`, `pgmqtt.tls_key_file`, `pgmqtt.license_key`, `pgmqtt.jwt_public_key`, `pgmqtt.jwt_required`, `pgmqtt.jwt_required_ws`, `pgmqtt.password_auth_enabled`, `pgmqtt.password_auth_required`, and `pgmqtt.password_auth_role_filter`.

Performance and observability GUCs include `pgmqtt.tick_interval_ms`, `pgmqtt.max_client_buffer_bytes`, `pgmqtt.cdc_every_n_ticks`, `pgmqtt.debug_log`, `pgmqtt.metrics_snapshot_interval`, `pgmqtt.metrics_retention_days`, `pgmqtt.metrics_connections_cache_interval`, `pgmqtt.metrics_hook_function`, and `pgmqtt.metrics_notify_channel`. Listener and TLS settings are read when the MQTT background worker starts, so they require a worker restart rather than only `pg_reload_conf()`.

### Caveats

- The README requires `wal_level = logical`; without logical decoding the CDC side will not work.
- This project's CSV tracks version `0.3.0`, PostgreSQL versions 14-18, and a package-side `pgrx` `0.18.1` rebuild note.
- MQTT 5.0 and MQTT 3.1.1 clients are documented as supported. QoS 0 and QoS 1 are supported; QoS 2 is not implemented and subscriptions requesting QoS 2 are downgraded to QoS 1.
- CDC captures `INSERT`, `UPDATE`, and `DELETE`; DDL changes and `TRUNCATE` are not captured. `DELETE` may require `REPLICA IDENTITY FULL`.
- Version 0.3.0 adds PostgreSQL role/password authentication, per-topic ACLs, admin disconnect/reload commands, metrics/observability tables and functions, MQTT 3.1.1 protocol tests, UNLOGGED metrics/cache recovery tests, and additional flow-control limits.
