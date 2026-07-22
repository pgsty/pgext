## Usage

Sources:

- [Official README](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/README.md)
- [Official usage documentation](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/doc/amqp.md)
- [Official control file](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/amqp.control)

`amqp` lets a PostgreSQL backend publish messages directly to an AMQP broker. It is useful for simple database-triggered notifications, but publishing is an external side effect and is not rolled back with the surrounding database transaction.

### Core Workflow

Add `pg_amqp` to `shared_preload_libraries`, restart PostgreSQL, create the extension, and insert the broker host, port, user, and password into `amqp.broker` using the table definition shipped by the installed version. Then publish with the returned broker ID:

```sql
CREATE EXTENSION amqp;

SELECT amqp.publish(
    1,
    'amqp.direct',
    'jobs.created',
    '{"job_id":42}',
    2,
    'application/json',
    NULL,
    'job-42'
);

SELECT amqp.disconnect(1);
```

The final four `amqp.publish` arguments are optional message properties: delivery mode, content type, reply-to, and correlation ID. Delivery mode 1 is non-persistent and 2 requests persistent delivery; broker durability and consumer acknowledgement still determine end-to-end guarantees.

### Important Objects

- `amqp.broker` stores broker connection details, including credentials.
- `amqp.publish` sends a message to an exchange with a routing key and payload.
- `amqp.disconnect` closes the current backend's connection for one broker; backend termination closes its remaining broker connections.

### Operational Notes

Restrict all privileges on `amqp.broker` and on publish functions, and avoid exposing credentials through dumps, replicas, logs, or broad read roles. Network calls execute in database backends, so broker latency or failure can extend SQL latency and consume connections.

Do not use a successful SQL commit as proof that a message was delivered exactly once. Design consumers to be idempotent, include a durable application identifier, and use an outbox worker when atomic database-state/message coordination is required. The upstream project and API are old; compatibility, TLS/authentication support, reconnect behavior, and failure semantics need direct testing against the chosen broker and PostgreSQL release.
