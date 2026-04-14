## Usage
- GitHub Repo: [`RayElg/pgmqtt`](https://github.com/RayElg/pgmqtt)
- README: [RayElg/pgmqtt/blob/main/README.md](https://github.com/RayElg/pgmqtt/blob/main/README.md)

`pgmqtt` is a `pgrx`-based PostgreSQL extension that turns CDC into an MQTT broker. Database changes can be published to MQTT topics through SQL-defined mappings, and MQTT publishes can be written back into PostgreSQL tables.

The README's quickstart requires `wal_level = logical` so CDC events can be captured correctly.

### Outbound Mapping

```sql
SELECT pgmqtt_add_outbound_mapping(
    'public',
    'my_table',
    'topics/{{ op | lower }}',
    '{{ columns | tojson }}'
);
```

With this mapping, `INSERT`, `UPDATE`, and `DELETE` on the table are published as MQTT messages. Subscribers to `topics/insert`, `topics/update`, or `topics/delete` receive JSON payloads.

### Inbound Mapping

```sql
SELECT pgmqtt_add_inbound_mapping(
    'sensor/{site_id}/temperature',
    'sensor_readings',
    '{"site_id": "{site_id}", "value": "$.temperature"}'::jsonb
);
```

When a client publishes to `sensor/site-1/temperature` with payload `{"temperature": 22.5}`, the README says a row is inserted into `sensor_readings` with the extracted values.

### Client Examples

```bash
mosquitto_sub -h localhost -t 'topics/#'
mosquitto_pub -h localhost -t 'sensor/site-1/temperature' -m '{"temperature": 22.5}'
```

### Scope

The upstream README covers the broker model, the outbound/inbound mapping examples, and basic MQTT client usage. It does not document a separate project homepage, so this stub stays at README scope.
