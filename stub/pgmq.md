


## Usage

> [pgmq: A lightweight message queue for PostgreSQL](https://github.com/pgmq/pgmq)

PGMQ is a lightweight message queue built on PostgreSQL, providing guaranteed "exactly once" delivery within a visibility timeout, FIFO queues, topic-based routing, and message archival.

```sql
CREATE EXTENSION pgmq;
```

### Create a Queue

```sql
SELECT pgmq.create('my_queue');
```

### Send Messages

```sql
-- Send a single message (returns msg_id)
SELECT * FROM pgmq.send(
  queue_name => 'my_queue',
  msg        => '{"foo": "bar"}'
);

-- Send with delay (invisible for 5 seconds)
SELECT * FROM pgmq.send(
  queue_name => 'my_queue',
  msg        => '{"foo": "bar"}',
  delay      => 5
);

-- Send a batch of messages
SELECT pgmq.send_batch(
  queue_name => 'my_queue',
  msgs       => ARRAY['{"a":1}','{"b":2}','{"c":3}']::jsonb[]
);
```

### Read Messages

Read messages and make them invisible for a visibility timeout (in seconds):

```sql
SELECT * FROM pgmq.read(
  queue_name => 'my_queue',
  vt         => 30,    -- visibility timeout in seconds
  qty        => 2      -- number of messages to read
);
```

### Pop a Message

Read and immediately delete a message:

```sql
SELECT * FROM pgmq.pop('my_queue');
```

### Delete a Message

```sql
SELECT pgmq.delete('my_queue', 6);
```

### Archive a Message

Move a message from the queue to the archive table for long-term retention:

```sql
SELECT pgmq.archive(queue_name => 'my_queue', msg_id => 2);

-- Archive multiple messages
SELECT pgmq.archive(queue_name => 'my_queue', msg_ids => ARRAY[3, 4, 5]);
```

Inspect archived messages:

```sql
SELECT * FROM pgmq.a_my_queue;
```

### Drop a Queue

```sql
SELECT pgmq.drop_queue('my_queue');
```

### Visibility Timeout

Messages become invisible after being read for the duration of the visibility timeout (`vt`). If not deleted or archived within that time, they become visible again for other consumers. Set `vt` greater than the expected processing time.

### Key Functions

| Function | Description |
|----------|-------------|
| `pgmq.create(queue_name)` | Create a new queue |
| `pgmq.send(queue_name, msg, [delay])` | Send a message |
| `pgmq.send_batch(queue_name, msgs)` | Send multiple messages |
| `pgmq.read(queue_name, vt, qty)` | Read messages with visibility timeout |
| `pgmq.pop(queue_name)` | Read and delete a message atomically |
| `pgmq.delete(queue_name, msg_id)` | Delete a message |
| `pgmq.archive(queue_name, msg_id/msg_ids)` | Archive message(s) |
| `pgmq.drop_queue(queue_name)` | Delete a queue |
