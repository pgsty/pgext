## Usage

Sources: [official README](https://github.com/RayElg/pgmqtt/blob/main/README.md), [official repo](https://github.com/RayElg/pgmqtt)

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
  '{{ columns | tojson }}'
);
```

With that mapping, `INSERT`, `UPDATE`, and `DELETE` publish JSON payloads to topics such as `topics/insert`.

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

### MQTT Client Examples

```bash
mosquitto_sub -h localhost -t 'topics/#'
mosquitto_pub -h localhost -t 'sensor/site-1/temperature' -m '{"temperature": 22.5}'
```

### Caveats

- The README requires `wal_level = logical`; without logical decoding the CDC side will not work.
- Upstream documentation is currently README-level only, so the documented SQL surface is limited to the inbound and outbound mapping workflow.
