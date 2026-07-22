## Usage

Sources:

- [Official README](https://github.com/edpratomo/pg_redispub/blob/e550997426a2f66accf6a28354bc6d741a92cfa2/README.md)
- [Official extension SQL](https://github.com/edpratomo/pg_redispub/blob/e550997426a2f66accf6a28354bc6d741a92cfa2/pg_redispub--1.0.sql)
- [Official implementation](https://github.com/edpratomo/pg_redispub/blob/e550997426a2f66accf6a28354bc6d741a92cfa2/pg_redispub.c)

`pg_redispub` 1.0 sends a Redis `PUBLISH` command from a PostgreSQL backend. Its single function always connects synchronously to Redis at `127.0.0.1:6379`; the endpoint, authentication, and TLS settings are not configurable in the pinned source.

### Core Workflow

```sql
CREATE EXTENSION pg_redispub;

SELECT redispub(
  'order_events',
  '{"event":"order_paid","order_id":42}'
);
```

The first argument is the Redis channel and the second is the message. The function returns false and emits a warning when the hiredis connection or command context reports an error; otherwise it returns true.

### Important Objects

- `redispub(text, text)` opens a connection, issues `PUBLISH channel message`, closes the connection, and returns a boolean status.

### Delivery and Transaction Boundaries

Redis Pub/Sub is transient and `redispub` is not transaction-aware. A message is sent immediately even if the surrounding PostgreSQL transaction later rolls back, and a retried transaction can publish duplicates. A true return value does not report the Redis subscriber count and does not prove that any subscriber consumed or persisted the message. Use an outbox table and a separate delivery worker when atomic, durable, or retryable delivery is required.

### Operational Notes

The extension must be built with hiredis. Each call creates a new connection with a 500 ms connect timeout, so row-by-row use can add substantial latency. The implementation has no Redis password, ACL username, TLS, alternate host, alternate port, pooling, or reconnect configuration. PostgreSQL's operating-system security policy must permit connections to the Redis port; upstream documents SELinux denial troubleshooting. Restrict execution privileges and test Redis outage behavior before using the function in triggers or write paths.
