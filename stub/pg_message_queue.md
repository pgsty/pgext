## Usage

Sources:

- [Upstream README](https://github.com/rpdelaney/pg-message-queue/blob/05309d14ae8a61b958ba3d56d2a1bccf53739bf0/README.md)
- [Version 0.2.1 install SQL](https://github.com/rpdelaney/pg-message-queue/blob/05309d14ae8a61b958ba3d56d2a1bccf53739bf0/pg_message_queue--0.2.1.sql)
- [Security notes](https://github.com/rpdelaney/pg-message-queue/blob/05309d14ae8a61b958ba3d56d2a1bccf53739bf0/doc/security)

pg_message_queue implements small, durable, channel-based queues with PostgreSQL tables, PL/pgSQL, triggers, and LISTEN/NOTIFY. Each channel owns a queue table; the notification carries a message identifier while the payload remains in the table until a consumer retrieves it.

### Basic workflow

Version 0.2.1 accepts a PostgreSQL registered type for each queue, including jsonb:

```sql
CREATE EXTENSION pg_message_queue;

SELECT pg_mq_create_queue('jobs', 'jsonb'::regtype);
LISTEN jobs;

SELECT pg_mq_send_message(
    'jobs',
    '{"task":"reindex","relation":"public.orders"}'::jsonb
);

BEGIN;
SELECT * FROM pg_mq_get_msg_text('jobs', 10);
COMMIT;
```

Retrieval marks rows delivered inside the current transaction. A rollback makes them available again. Applications may listen for wakeups or poll in batches; LISTEN is optional.

Queue creation is a security-definer operation whose PUBLIC execute privilege is revoked by the install script. Sending and receiving also depend on privileges on the generated queue table, so grant only the operations each producer or consumer needs.

### Caveats

- Delivery is transactional but not an exactly-once distributed protocol. Consumers must make external side effects idempotent and decide how to recover abandoned work.
- The retrieval functions update queue rows and can block under concurrent consumption. This design is suitable for modest integration queues, not a high-throughput broker replacement.
- `pg_mq_create_queue` and `pg_mq_drop_queue` are security-definer functions without a fixed `SET search_path`. Although installation revokes their default public execution privilege, do not grant them to untrusted roles without first repairing and reviewing that boundary.
- PostgreSQL notifications are wakeups, not protected payload delivery. Database users can generally LISTEN on a channel; secure the underlying tables and do not put secrets in notification payloads.
- The published 0.2.1 SQL defines pg_mq_rebuild_triggers with an undeclared t_table_name reference. Do not rely on that restoration helper without patching and testing it.
- Some internal function signatures use integer message identifiers even though queue identifiers are based on bigint. Test rollover and retention behavior for long-lived queues.
