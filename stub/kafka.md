## Usage

Sources:

- [Official README](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/README.md)
- [Usage documentation](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/doc/kafka.md)
- [Extension SQL](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/sql/kafka.sql)
- [Producer implementation](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/src/pg_kafka.c)
- [Extension control file](https://github.com/xstevens/pg_kafka/blob/e83efadcd3522938d43a9fbbf5846fbd03b70c37/kafka.control)

`kafka` 0.0.1, from the archived `pg_kafka` project, sends text messages directly from a PostgreSQL backend to Apache Kafka through librdkafka. It was written for PostgreSQL 9.2+ and Kafka/librdkafka 0.8-era APIs. The implementation does not provide reliable transactional delivery despite the SQL comment claiming commit coupling; use a transactional outbox and an external producer instead.

### Historical Workflow

Upstream instructs users to preload the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_kafka.so'
```

Install the extension, register brokers, and enqueue a message:

```sql
CREATE EXTENSION kafka;

INSERT INTO kafka.broker (host, port)
VALUES ('kafka.internal', 9092);

SELECT kafka.produce('audit-events', '{"event":"example"}');
SELECT kafka.close();
```

`kafka.produce(varchar, varchar)` returns true only when librdkafka accepts the message into its local producer queue. `kafka.close()` tears down the producer held by the current database backend. Broker configuration is read only when that backend first creates its producer; close or reconnect after changing `kafka.broker`.

### Delivery and Transaction Caveats

The C function calls `rd_kafka_produce` immediately. Its transaction callback does nothing on commit and merely destroys the producer on abort; there is no PostgreSQL commit record, outbox, Kafka transaction, idempotency key, retry ledger, or two-phase protocol. Successful enqueue can still fail later, an aborted SQL transaction can already have caused an external effect, and backend or broker failure can lose or duplicate messages.

The reviewed code does not poll or flush after a successful enqueue, so delivery callbacks and queue draining are not a reliable acknowledgement path. It also declares the externally mutating `kafka.produce` and `kafka.close` functions `IMMUTABLE`, allowing PostgreSQL to fold, cache, reorder, or omit calls in ways that are invalid for side effects. `kafka.close` is declared to return boolean in SQL but returns `void` in C. These are source defects, not application-level knobs.

### Operational Boundary

Do not deploy this archived native module on a modern cluster without a maintained fork that corrects volatility, return types, lifecycle, polling, delivery confirmation, memory/error paths, TLS/authentication, and current librdkafka/PostgreSQL compatibility. Triggers make database writes depend on Kafka latency and failure behavior but still do not create atomic delivery. For durable change publication, write an outbox row in the same transaction and publish it asynchronously, or use supported logical decoding/CDC tooling with monitored offsets and replay semantics.
