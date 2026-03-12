


## Usage

> [pgmb: A lightweight message broker system built inside PostgreSQL](https://github.com/fraruiz/pgmb)

The `pgmb` extension provides an in-database message broker with HTTP-based worker dispatch, automatic retries, dead letter queues, and pattern-based routing.

```sql
CREATE EXTENSION pgmb;  -- requires pg_cron and http extensions
```

### Register a Worker

```sql
SELECT pgmb.worker(
    'order_processor',                     -- worker name
    'http://localhost:8080/process',       -- endpoint URL
    100                                    -- requests per second limit
);
-- Returns: worker UUID
```

### Create a Queue

```sql
SELECT pgmb.create(
    'order_queue',                         -- queue name
    'order.*',                             -- binding key pattern (supports * wildcard)
    5,                                     -- max retries
    '550e8400-e29b-41d4-a716-446655440000' -- worker UUID
);
-- Returns: queue UUID
```

### Send Messages

```sql
-- Simple message
SELECT pgmb.send(
    gen_random_uuid(),
    'order.created',
    '{"order_id": 123, "amount": 45.67}'::jsonb
);

-- With custom headers
SELECT pgmb.send(
    gen_random_uuid(),
    'order.created',
    '{"order_id": 123}'::jsonb,
    '{"source": "web", "priority": "high"}'::jsonb
);

-- Delayed message (by timestamp or seconds)
SELECT pgmb.send(
    gen_random_uuid(),
    'order.created',
    '{"order_id": 123}'::jsonb,
    '{}'::jsonb,
    now() + interval '10 minutes'
);
```

### API Reference

| Function | Description |
|----------|-------------|
| `pgmb.worker(name, endpoint, rps)` | Register an HTTP worker endpoint |
| `pgmb.create(name, binding_key, max_retries, worker_id)` | Create a queue with routing pattern |
| `pgmb.send(id, routing_key, body)` | Send a message |
| `pgmb.send(id, routing_key, body, headers)` | Send a message with headers |
| `pgmb.send(id, routing_key, body, headers, delay)` | Send a delayed message |

### How It Works

1. Messages are inserted into `pgmb.messages` via `pgmb.send()`
2. A trigger routes messages to matching queues based on routing key patterns
3. `pg_cron` dispatches messages via HTTP POST to worker endpoints every second
4. Failed messages are retried; after max retries they move to a dead letter queue

### Monitoring

```sql
SELECT * FROM pgmb.workers;
SELECT * FROM pgmb.queues;
SELECT COUNT(*) FROM pgmb.order_queue WHERE acknoledge = false;
SELECT * FROM pgmb.order_dead_letter_queue;
```
