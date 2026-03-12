


## Usage

> [pgq: Generic high-performance lockless queue for PostgreSQL](https://github.com/pgq/pgq)

PgQ is a PostgreSQL extension that provides a generic, high-performance lockless queue with a simple SQL function API. It uses a producer-consumer model with batch-based event processing.

```sql
CREATE EXTENSION pgq;
```

### Core Concepts

- **Queue**: A named event stream. Events are inserted by producers and consumed in batches.
- **Consumer**: A named subscriber registered to a queue. Each consumer tracks its own position.
- **Batch**: A group of events retrieved together. Consumers process events batch by batch.
- **Ticker**: A background process that creates batch boundaries (ticks) at regular intervals.

### Queue Management

```sql
-- Create a queue
SELECT pgq.create_queue('myqueue');

-- Drop a queue
SELECT pgq.drop_queue('myqueue');

-- Get queue info
SELECT * FROM pgq.get_queue_info();
SELECT * FROM pgq.get_queue_info('myqueue');
```

### Consumer Registration

```sql
-- Register a consumer on a queue
SELECT pgq.register_consumer('myqueue', 'myconsumer');

-- Unregister a consumer
SELECT pgq.unregister_consumer('myqueue', 'myconsumer');

-- Get consumer info
SELECT * FROM pgq.get_consumer_info('myqueue');
```

### Producing Events

```sql
-- Insert an event into a queue
SELECT pgq.insert_event('myqueue', 'event_type', 'event_data');

-- Insert with extra fields
SELECT pgq.insert_event('myqueue', 'event_type', 'event_data',
                         'extra1', 'extra2', 'extra3', 'extra4');
```

### Consuming Events

```sql
-- Get the next batch of events (returns batch_id or NULL if no new batches)
SELECT pgq.next_batch('myqueue', 'myconsumer');

-- Get events from the batch
SELECT * FROM pgq.get_batch_events(:batch_id);

-- Retry a failed event (will reappear after the specified interval)
SELECT pgq.event_retry(:batch_id, :event_id, :retry_seconds);

-- Mark batch as done
SELECT pgq.finish_batch(:batch_id);
```

### Typical Consumer Loop

```sql
-- 1. Get next batch
SELECT pgq.next_batch('myqueue', 'myconsumer') AS batch_id;

-- 2. If batch_id is not NULL, get events
SELECT * FROM pgq.get_batch_events(:batch_id);

-- 3. Process events, retry failures
SELECT pgq.event_retry(:batch_id, :event_id, 60);

-- 4. Finish the batch
SELECT pgq.finish_batch(:batch_id);
```

### Maintenance

PgQ requires a ticker daemon (`pgqd`) to run in the background for creating batch boundaries and performing maintenance tasks like table rotation and retry event processing.

### Key Functions

| Function | Description |
|----------|-------------|
| `pgq.create_queue(name)` | Create a new queue |
| `pgq.drop_queue(name)` | Remove a queue |
| `pgq.register_consumer(queue, consumer)` | Register a consumer |
| `pgq.unregister_consumer(queue, consumer)` | Unregister a consumer |
| `pgq.insert_event(queue, type, data, ...)` | Insert an event |
| `pgq.next_batch(queue, consumer)` | Get next batch ID |
| `pgq.get_batch_events(batch_id)` | Get events from a batch |
| `pgq.event_retry(batch_id, event_id, seconds)` | Schedule event retry |
| `pgq.finish_batch(batch_id)` | Mark batch as processed |
| `pgq.get_queue_info([name])` | Get queue statistics |
| `pgq.get_consumer_info(queue)` | Get consumer statistics |
