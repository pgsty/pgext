## Usage

Sources:

- [pg_qpubsub/README.md](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/README.md)
- [pg_qpubsub/Makefile](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/Makefile)
- [pg_qpubsub/build.sh](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/build.sh)
- [pg_qpubsub/wrappers.sql](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/wrappers.sql)
- [lib/schema/schema.sql](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/lib/schema/schema.sql)
- [pg_qpubsub/pg_qpubsub.control](https://github.com/smartpricing/queen/blob/c0cf3b3640a4280ef997bde8adb8b9bda16f59e3/pg_qpubsub/pg_qpubsub.control)
- [Removal commit](https://github.com/smartpricing/queen/commit/cc213db9c507f6b1869a03b7ba45486506208bd6)

`pg_qpubsub` 1.0 is a historical SQL extension for Queen message queues, partitions and consumer groups. Upstream removed its extension directory on July 19, 2026. These instructions describe the linked historical revision; current Queen releases do not supply this extension.

### Core Workflow

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_qpubsub;
SELECT queen.configure('orders', 60, 3, true);
SELECT queen.produce_one('orders', '{"orderId":123}'::jsonb);
SELECT * FROM queen.consume_one('orders', '__QUEUE_MODE__', 10);
```

After processing a consumed message, call `queen.commit_one` with the actual transaction ID, partition UUID and lease ID returned by consumption. `queen.nack` retries a message, `queen.reject` sends it to the dead-letter queue, and `queen.renew_one` extends a lease. Use a named consumer group for independent fan-out; `__QUEUE_MODE__` selects queue consumption.

### API and Operational Boundaries

The JSONB APIs `queen.produce`, `queen.consume`, `queen.commit`, `queen.renew` and `queen.transaction` support batches. The SQL convenience layer also exposes `queen.configure`, `queen.lag`, `queen.has_messages`, `queen.seek` and `queen.delete_consumer_group`. Batch consumption returns immediately; only the scalar consumption API supports a polling timeout. Application notification wiring remains the application's responsibility.

The historical README declares PostgreSQL 14 or later and requires `pgcrypto`. Objects reside in the fixed `queen` schema. This is pure SQL with no preload or server restart, and the control file sets `superuser=false`; the installer still needs the object-creation and dependency privileges. The versioned install SQL is assembled from the same revision's schema, procedures and wrappers.

The supplied build grants PUBLIC access to the schema, all its tables, sequences and functions. Review and restrict those grants before handling sensitive messages. Reconcile leases, retries, retention and dead-letter handling with application failure behavior. Do not install into an existing Queen schema without assessing object conflicts, and do not assume an upgrade path from this removed extension to the current broker. No current Pigsty package or runtime support is claimed.
