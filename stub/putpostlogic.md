## Usage

Sources:

- [Official output-plugin implementation](https://github.com/rektide/putpostlogic/blob/0d71733aa77085b4f62c49c0bba972c241d0fcbc/putpostlogic.c)
- [Official change decoder](https://github.com/rektide/putpostlogic/blob/0d71733aa77085b4f62c49c0bba972c241d0fcbc/decode.c)
- [Official extension control file](https://github.com/rektide/putpostlogic/blob/0d71733aa77085b4f62c49c0bba972c241d0fcbc/putpostlogic.control)

`putpostlogic` is a historical logical-decoding output plugin that groups row changes into transaction JSON and can additionally send that JSON to Kafka or nanomsg sockets. It is headless: the pinned source has no extension SQL script and is loaded by a logical replication slot, not by `CREATE EXTENSION`.

### Historical Invocation

The server must be configured for logical decoding and the plugin's json-c, librdkafka, and nanomsg libraries must be present. On a compatible old PostgreSQL build, a slot invocation follows this shape:

```sql
SELECT *
FROM pg_create_logical_replication_slot(
  'putpostlogic_demo',
  'putpostlogic'
);

SELECT *
FROM pg_logical_slot_get_changes(
  'putpostlogic_demo', NULL, NULL,
  'include-transaction', 'true',
  'json-topic', 'database_changes',
  'kafka-brokers', 'localhost:9092',
  'kafka-topic', 'postgres_changes'
);
```

Recognized options in the reviewed source include `include-transaction`, `json-topic`, `kafka-brokers`, `kafka-topic`, `nanomsg-pair`, `nanomsg-pair-bind`, `nanomsg-pub`, `nanomsg-pub-topic`, `nanomsg-pub-bind`, `nanomsg-push`, and `nanomsg-push-bind`.

### Output and Identity

The plugin emits one JSON object per commit with numbered change members, relation identifiers, new values for inserts/updates, and an `@where` identity object for updates/deletes. Correct old-row identity depends on a primary/replica-identity index or `REPLICA IDENTITY FULL`; unchanged TOAST values use a sentinel string rather than the original value.

### Critical Boundaries

The fixed commit is unfinished 2014-era code and contains multiple unsafe paths, including unchecked or undersized allocations, NULL-value handling defects, an invalid replica-index subscript, and an unconditional Kafka poll even when Kafka is not configured. It also targets obsolete logical-decoding APIs and may not compile or run safely on modern PostgreSQL. Do not deploy this build without a source audit, fixes, sanitizer testing, and end-to-end crash tests.

Kafka and nanomsg sends are external side effects of decoding, not transactional delivery. Re-reading a slot can duplicate messages, failures can diverge slot output from broker state, and decoded rows may expose secrets or personal data. Treat the example as interface documentation only; use a maintained logical-decoding/CDC system with explicit delivery, retry, ordering, schema evolution, access control, and slot-retention guarantees for production.
