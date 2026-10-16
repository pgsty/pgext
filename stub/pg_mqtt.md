## Usage

Sources:

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/LICENSE)
- [extensions/pg_mqtt/pg_mqtt.control](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/pg_mqtt.control)
- [extensions/pg_mqtt/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/Cargo.toml)
- [extensions/pg_mqtt/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/src/lib.rs)
- [extensions/pg_mqtt/sql/pg_mqtt--0.3.2--0.4.0.sql](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/sql/pg_mqtt--0.3.2--0.4.0.sql)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/crates/pg_bgworker/src/supervision.rs)
- [extensions/pg_mqtt/src/protocol/codec.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/src/protocol/codec.rs)
- [extensions/pg_mqtt/src/protocol/handlers.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/src/protocol/handlers.rs)
- [extensions/pg_mqtt/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/pgbrew.toml)

`pg_mqtt` 0.4.2 implements an MQTT broker backed by `pgmqtt` tables. The reviewed decoder accepts protocol level 5 and rejects other levels; the extension README’s older 3.1.1 claim is stale.

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

Extension version 0.4.0 uses `pg_mqtt.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above.

### Worker Supervision

Version 0.3.2 adds `pgmqtt.worker_status()` and the superuser-only `pgmqtt.reset_workers()`. Status reports worker name, state, failures, restarts, PID, time of entry into the current state and last failure. `pg_mqtt.max_worker_failures` defaults to 10 consecutive failures; 0 retries indefinitely. Restart delays back off from 5 to 60 seconds, and reaching the limit leaves a failed worker idle until reset.

Install matching libraries, restart PostgreSQL to initialize the new shared memory, and update the extension SQL in the configured worker database. Inspect the failure before resetting; a reset does not repair its cause.

```sql
ALTER EXTENSION pg_mqtt UPDATE;
SELECT * FROM pgmqtt.worker_status();
```

### Changes in This Release

Source feature declarations cover PostgreSQL 14–18. The repository release is 0.7.2, while each extension retains its own SQL version; this does not establish Pigsty package availability.

Version 0.4.0 fans messages out to subscribers on all workers through the unlogged `pgmqtt.outbox`, polled every 50 ms. SQL `mqtt_publish` delivers after commit and never after rollback; other long-running write transactions in the cluster can delay fan-out. The unlogged queue does not provide crash-durable delivery. The upgrade script creates the queue; install matching libraries before updating extension SQL.

### 0.4.1 Typed-Topic Fix

Version 0.4.1 fixes `mqtt_publish` for typed topics whose backing table has no `id` column: the publish path no longer assumes `RETURNING id`. The SQL schema is unchanged. Use matching 0.4.1 files and the extension upgrade path; the transactional outbox and delivery limits above still apply.

### Idle Workers in 0.4.2

Repository release 0.7.2 contains extension 0.4.2. Read-only or empty-claim polls no longer allocate a transaction ID merely because they use the SPI bridge; an actual row write still does. This removes empty commit WAL from idle polling. New bridge requests wake the worker latch; the main loop can otherwise sleep for up to 100 ms. Install matching libraries, restart the preloaded workers, and update the SQL extension. The existing privilege, external-service and pre-0.3.0 migration boundaries remain.
