## Usage

Sources:

- [contrib/alohadb_queue/alohadb_queue.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_queue/alohadb_queue.control)
- [contrib/alohadb_queue/alohadb_queue--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_queue/alohadb_queue--1.0.sql)
- [contrib/alohadb_queue/queue_ops.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_queue/queue_ops.c)

`alohadb_queue` 1.0 supplies message queues and consumer groups in the pinned AlohaDB source, with explicit acknowledgements and stored visibility timestamps.

### Core Workflow

```sql
CREATE EXTENSION alohadb_queue;
SELECT queue_create('work');
SELECT queue_send('work', '{"task":"example"}'::jsonb);
SELECT * FROM queue_receive('work', 1);
```

### Objects and Maintenance

After successful application processing, call `queue_ack(queue_name, msg_id)` with the returned message ID. `queue_nack` releases a delivery; `queue_send_batch` submits multiple messages. `queue_subscribe`, `queue_poll` and `queue_commit_offset` manage consumer-group offsets. `queue_stats` reports queue state; `queue_purge` deletes all messages from a queue, and `queue_drop` removes the queue.

The extension creates `alohadb_queue_queues`, `alohadb_queue_messages` and `alohadb_queue_consumers` in fixed `public`. Installation uses C functions and requires superuser rights; the reviewed source does not require preload. The reviewed receive path changes a message to delivered and only selects ready messages on later receives; an expired timestamp alone does not make that delivery ready again. Use `queue_nack` for explicit recovery and make consumers idempotent for such redelivery. Retention configuration is stored, but the reviewed API provides a full purge rather than automatic age-based cleanup. This is an AlohaDB source boundary, not a stock PostgreSQL compatibility claim.
