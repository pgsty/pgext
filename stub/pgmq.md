## Usage

Sources:

- [PGMQ v1.12.0 README](https://github.com/pgmq/pgmq/blob/v1.12.0/README.md)
- [PGMQ v1.12.0 release notes](https://github.com/pgmq/pgmq/releases/tag/v1.12.0)
- [Version 1.12.0 migration SQL](https://github.com/pgmq/pgmq/blob/v1.12.0/pgmq-extension/sql/pgmq--1.11.1--1.12.0.sql)
- [PGMQ documentation](https://pgmq.github.io/pgmq/)

pgmq implements durable message queues as PostgreSQL tables and SQL functions. It supports delayed delivery, visibility timeouts, FIFO groups, message headers, polling, topics, and archival. Use it when queue transactions should be coordinated with relational changes in the same database.

### Create a Queue and Send

    CREATE EXTENSION pgmq;
    SELECT pgmq.create('jobs');

    SELECT *
    FROM pgmq.send(
      queue_name => 'jobs',
      msg        => '{"task":"refresh"}'::jsonb,
      delay      => 0
    );

send returns the message identifier. send_batch inserts multiple JSONB messages. Headers can carry routing or tracing metadata separately from the body where the selected overload supports them.

### Read with a Visibility Timeout

    SELECT *
    FROM pgmq.read(
      queue_name => 'jobs',
      vt         => 30,
      qty        => 10
    );

Reading hides each message for vt seconds. On success, delete or archive it:

    SELECT pgmq.delete('jobs', 42);
    SELECT pgmq.archive('jobs', 43);

If processing fails or the consumer disappears, an unacknowledged message becomes visible again. Consumers must therefore be idempotent; pgmq does not make arbitrary application side effects globally exactly once.

pop reads and deletes immediately and is appropriate only when losing a message after the call is acceptable.

### FIFO Group Head Polling

Version 1.12.0 adds a polling read for the head message of multiple FIFO groups:

    SELECT *
    FROM pgmq.read_grouped_head_with_poll(
      queue_name          => 'jobs',
      vt                  => 30,
      qty                 => 10,
      max_poll_seconds    => 5,
      poll_interval_ms    => 100
    );

It selects head-of-group messages while polling for up to max_poll_seconds. This preserves per-group order while allowing different groups to be processed concurrently.

### Queue Administration Index

- pgmq.create(queue_name): create queue and archive structures.
- pgmq.send and pgmq.send_batch: enqueue JSONB messages, optionally delayed.
- pgmq.read: claim visible messages for a visibility timeout.
- pgmq.read_grouped_head_with_poll: poll FIFO group heads.
- pgmq.pop: read and immediately delete.
- pgmq.delete: acknowledge by removing a message.
- pgmq.archive: move a message to the queue archive table.
- pgmq.drop_queue: remove queue objects.
- pgmq.metrics and related helpers: inspect queue depth and age where available.

For queue jobs, archive rows are stored in pgmq.a_<queue_name>. Treat those tables as extension-managed objects.

### Operational Notes

- Set vt longer than normal processing time and design for redelivery after timeouts.
- Queue and archive tables consume ordinary PostgreSQL WAL, storage, vacuum, and backup capacity.
- Archive or delete completed messages and enforce an archive-retention policy.
- Long polling holds a database connection. Size connection pools and polling intervals for the number of consumers.
- Keep queue names within pgmq's identifier rules; call the API rather than constructing table names from untrusted input.
