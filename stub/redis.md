


## Usage

> [redis: Send Redis pub/sub messages to Redis from PostgreSQL directly](https://github.com/brettlaforge/pg_redis_pubsub)

The `redis` extension (pg_redis_pubsub) allows PostgreSQL to publish messages to Redis pub/sub channels. Currently, only the publish side is supported.

### Configuration

Set the Redis connection parameters via GUC variables:

```sql
ALTER SYSTEM SET redis.host = '127.0.0.1';
ALTER SYSTEM SET redis.port = '6379';
SELECT pg_reload_conf();
```

These can also be set at the database, role, or session level:

```sql
SET redis.host = '127.0.0.1';
SET redis.port = '6379';
```

### Basic Usage

```sql
CREATE EXTENSION redis;

SELECT redis_connect();
SELECT redis_publish('mychannel', 'Hello World');
SELECT redis_disconnect();
```

### Available Functions

| Function | Description |
|----------|-------------|
| `redis_connect()` | Connect to Redis using `redis.host` and `redis.port` settings |
| `redis_disconnect()` | Disconnect from Redis |
| `redis_publish(channel text, message text)` | Publish a message on the specified channel |
| `redis_status()` | Return the status of the Redis client |

Note: `redis_publish` automatically connects if no connection exists.

### Trigger Example

Publish change events to Redis whenever a table is modified:

```sql
CREATE OR REPLACE FUNCTION notify_changes()
RETURNS TRIGGER AS $$
BEGIN
  PERFORM redis_publish(
    'products:' || NEW.id::text,
    to_jsonb(NEW)::text
  );
  RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER on_product_change
  AFTER INSERT OR UPDATE ON products
  FOR EACH ROW EXECUTE PROCEDURE notify_changes();
```

This allows external subscribers listening on Redis channels to react to PostgreSQL data changes in real time.
