## Usage

Sources:

- [pgmqtt 0.4.1 README](https://github.com/RayElg/pgmqtt/blob/0.4.1/README.md)
- [pgmqtt 0.4.1 interfaces](https://github.com/RayElg/pgmqtt/blob/0.4.1/docs/interfaces.md)
- [pgmqtt 0.4.1 configuration](https://github.com/RayElg/pgmqtt/blob/0.4.1/docs/configuration.md)
- [pgmqtt 0.4.1 limitations](https://github.com/RayElg/pgmqtt/blob/0.4.1/docs/limitations.md)
- [pgmqtt 0.4.1 release notes](https://github.com/RayElg/pgmqtt/releases/tag/0.4.1)

pgmqtt embeds an MQTT broker in PostgreSQL. It can publish INSERT, UPDATE, and DELETE changes through logical decoding and can map inbound MQTT topics and JSON payloads to table writes. Use it when database and MQTT integration justify running a network broker inside the PostgreSQL server process.

### Preload and Create the Extension

Set logical WAL and preload the worker, then restart PostgreSQL:

    wal_level = logical
    shared_preload_libraries = 'pgmqtt'

Create the extension after restart:

    CREATE EXTENSION pgmqtt;

Listener addresses, ports, authentication, and TLS settings are read by the background worker. Settings documented as startup-only require a worker/server restart, not just pg_reload_conf().

### Publish Table Changes

Create an outbound mapping:

    SELECT pgmqtt_add_outbound_mapping(
      'public',
      'orders',
      'orders/{{ op | lower }}',
      '{{ columns | tojson }}',
      1
    );

The mapping publishes row changes to topics such as orders/insert. The interface also accepts a QoS and template type where supported. Version 0.4.1 drains CDC changes in batches of up to 4096 records.

Inspect or remove outbound mappings:

    SELECT * FROM pgmqtt_list_outbound_mappings();
    SELECT pgmqtt_remove_outbound_mapping('public', 'orders');

### Write Rows from MQTT

Map captured topic segments and JSON fields to a target table:

    SELECT pgmqtt_add_inbound_mapping(
      'sensor/{site_id}/temperature',
      'sensor_readings',
      '{"site_id":"{site_id}","value":"$.temperature"}'::jsonb
    );

Inbound mappings support insert and documented upsert/delete modes with options such as target_schema, conflict_columns, mapping_name, and template_type. Grant the worker role only the required table privileges and validate payload types and constraints.

    SELECT * FROM pgmqtt_list_inbound_mappings();
    SELECT pgmqtt_remove_inbound_mapping('temp_readings');

### Administration and Status

    SELECT * FROM pgmqtt_status();
    SELECT pgmqtt_disconnect_client('device-42');
    SELECT pgmqtt_disconnect_role('mqtt_devices');
    SELECT pgmqtt_reload_acls('*');

pgmqtt_status reports listener, client, subscription, retained-message, CDC, inbound-write, and dead-letter state. Administrative calls are queued for asynchronous processing by the worker.

### Configuration Index

- pgmqtt.mqtt_enabled and pgmqtt.mqtt_port: TCP MQTT listener.
- pgmqtt.ws_enabled and pgmqtt.ws_port: WebSocket listener.
- pgmqtt.tick_interval_ms and pgmqtt.cdc_every_n_ticks: worker cadence.
- pgmqtt.max_client_buffer_bytes: per-client flow-control boundary.
- pgmqtt.debug_log and pgmqtt.metrics_*: diagnostics and metrics integration.
- pgmqtt TLS, JWT, password-authentication, and ACL settings: transport and client access controls; availability differs between community and enterprise features.

### Protocol and CDC Boundaries

- MQTT 5.0 and 3.1.1 are supported. QoS 0 and 1 are implemented; requested QoS 2 is downgraded to QoS 1.
- CDC covers INSERT, UPDATE, and DELETE, not DDL or TRUNCATE. DELETE payloads may require REPLICA IDENTITY FULL.
- The CDC ring has a finite capacity of 8192 and drops the oldest records on overflow. The QoS 0 topic buffer is capped at 4096 and also drops oldest entries; QoS 1 buffering can grow without a fixed bound.
- The community edition documents TLS through a proxy, while native TLS and some JWT features are enterprise boundaries. Verify the edition before setting listener expectations.

### Version 0.4.1 and Operations

The 0.4 line consolidates HTTP/worker handling and reduces panic paths; 0.4.1 raises CDC batch processing to 4096. These changes improve throughput and structure but do not make the embedded broker lossless under every overload or crash.

Running a broker inside PostgreSQL expands the database network and resource boundary. Isolate listener interfaces, enforce authentication and topic ACLs, monitor worker lag and dropped buffers, and test failover and restart behavior before production use.
