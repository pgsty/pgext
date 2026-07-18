## Usage

Sources:

- [pg_kafka_events README at the reviewed commit](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/README.md)
- [pg_kafka_events.control at the reviewed commit](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/pg_kafka_events.control)
- [Background-worker implementation](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/src/pg_kafka_events.c)
- [Build dependencies](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/Makefile)
- [Version 1.0 install SQL](https://github.com/alaisi/pg_kafka_events/blob/0d4476c1d6a587fc571aae5b1ded246d29fcefb2/pg_kafka_events--1.0.sql)

`pg_kafka_events` is a preloaded background worker that forks `pg_recvlogical`, reads newline-delimited logical-decoding output, and publishes each change through librdkafka. The default decoder is `decoding_json`; one configured database is streamed to one Kafka topic.

### Server Configuration

```conf
shared_preload_libraries = 'pg_kafka_events'
wal_level = logical
max_replication_slots = 4
max_wal_senders = 4
kafka.database_name = 'appdb'
kafka.topic = 'row_changes'
kafka.bootstrap_servers = 'kafka1:9092'
```

Install librdkafka, the chosen logical-decoding plugin, and the configured `pg_recvlogical` binary before restarting PostgreSQL. After restart, register the extension as a superuser:

```sql
CREATE EXTENSION pg_kafka_events;
```

The install SQL is empty; preloading starts the worker, while `CREATE EXTENSION` only records version `1.0` in that database.

### Caveats

- The worker always uses the replication slot name `kafka_events`. A stopped or failing consumer can retain WAL indefinitely; monitor slot lag and define an operator recovery procedure.
- This 2017 implementation does not provide transactional grouping, acknowledgements exposed to PostgreSQL, retries coordinated with the replication position, or an exactly-once guarantee. Validate duplicates and loss behavior under Kafka and process failures.
- Decoder lines are accumulated in a fixed 64 KiB static buffer without a length guard. Large row-change messages can overrun it; do not deploy the reviewed code without remediation.
- Published payloads are also emitted as PostgreSQL notices, potentially copying row data into server logs. The only upstream tag is `0.1` even though control and catalog use `1.0`.
