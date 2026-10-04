## Usage

Sources:

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [extensions/pg_mqtt/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_mqtt/pgbrew.toml)
- [extensions/pg_mqtt/pg_mqtt.control](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_mqtt/pg_mqtt.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/README.md)
- [extensions/pg_mqtt/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_mqtt/Cargo.toml)
- [extensions/pg_mqtt/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_mqtt/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/LICENSE)
- [extensions/pg_mqtt/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_mqtt/README.md)
- [extensions/pg_mqtt/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_mqtt/pgbrew.toml)
- [extensions/pg_mqtt/src/protocol/codec.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_mqtt/src/protocol/codec.rs)
- [extensions/pg_mqtt/src/protocol/handlers.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_mqtt/src/protocol/handlers.rs)

`pg_mqtt` 0.3.1 implements an MQTT broker backed by `pgmqtt` tables. The reviewed decoder accepts protocol level 5 and rejects other levels; the extension README’s older 3.1.1 claim is stale.

### Core Workflow

```ini
shared_preload_libraries = 'pg_mqtt'
pg_mqtt.database = 'postgres'
pg_mqtt.port = 1883
```

```sql
CREATE EXTENSION pg_mqtt;
SELECT pgmqtt.mqtt_publish('sensors/room1/temperature', '22.5');
SELECT * FROM pgmqtt.mqtt_messages('sensors/%', 10);
```

### Operational Boundaries

The control requires superuser installation. Preload and restart, creating the extension in the configured worker database. `pg_mqtt.database`, `pg_mqtt.port` and `pg_mqtt.worker_count` configure workers. The source binds all interfaces by default on port 1883, so isolate this experimental listener. `mqtt_publish`, `mqtt_messages`, `mqtt_sessions` and schema-binding helpers expose stored data; `mqtt_messages` uses SQL LIKE patterns. `mqtt_status` reports configuration and is not a socket health probe. Do not infer complete MQTT security or delivery guarantees from the project description. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

### Current Release and Upgrade

Extension version 0.3.1 uses `pg_mqtt.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above. Withdrawn repository 0.4.0 bottles must be replaced by 0.4.1; upstream bottles do not imply Pigsty package availability.
