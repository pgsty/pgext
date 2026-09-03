## Usage

Sources:

- [Official documentation](https://github.com/codarn/pg_mqtt_pub/blob/36ed725dc7aa9a621749a6c378caed84e64bc76e/README.md)
- [Extension control file](https://github.com/codarn/pg_mqtt_pub/blob/36ed725dc7aa9a621749a6c378caed84e64bc76e/pg_mqtt_pub.control)
- [Official repository](https://github.com/codarn/pg_mqtt_pub)

`pg_mqtt_pub` MQTT publishing from SQL, triggers, and scheduled jobs through a background worker.

### Enablement

Merge `pg_mqtt_pub` into the existing preload list, restart PostgreSQL, and then create `pg_mqtt_pub` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_mqtt_pub'
```

```sql
CREATE EXTENSION pg_mqtt_pub;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Broker connection status and message metrics
SELECT * FROM mqtt_status();

-- Example output:
--  host      | port | connected | messages_sent | messages_failed | dead_lettered | queue_depth | connected_since | disconnected_since | worker_pid
-- -----------+------+-----------+---------------+-----------------+---------------+-------------+-----------------+--------------------+----------
--  localhost | 1883 | true      |         12847 |               3 |             1 |           0 | 2026-02-18 09:15:00 | (null)             | 1234

-- Dead letter diagnostics
SELECT * FROM mqtt_pub.dead_letters ORDER BY failed_at DESC;

-- Summary view
SELECT * FROM mqtt_pub.dead_letter_summary;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `mqtt_publish` | FUNCTION | Callable function from the reviewed install surface. |
| `mqtt_pub.dead_letters` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `mqtt_status` | FUNCTION | Callable function from the reviewed install surface. |
| `mqtt_pub.dead_letter_summary` | VIEW | Inspection or query view created by the extension. |

### Operations and Boundaries

- Preloading `pg_mqtt_pub` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
- The extension fixes or creates schema objects under `mqtt_pub`; include them in privilege and backup review.
