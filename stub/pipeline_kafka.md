## Usage

Sources:

- [Official README](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/README.md)
- [Extension SQL API](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/pipeline_kafka--1.0.0.sql)
- [Kafka worker implementation](https://github.com/pipelinedb/pipeline_kafka/blob/2c2e6e9d71463427dca641b0657e12e35b255734/pipeline_kafka.c)

`pipeline_kafka` 1.0.0 is a historical Kafka integration for PipelineDB. Background workers consume Kafka topics into PipelineDB streams, and SQL functions or triggers produce messages. It requires the obsolete `pipelinedb` extension and old librdkafka interfaces; it is not compatible by design with an ordinary current PostgreSQL installation.

### Preload and Broker Setup

Build against the upstream-documented librdkafka version, preload both modules, and restart PostgreSQL:

```conf
shared_preload_libraries = 'pipelinedb,pipeline_kafka'
```

Create the extension and register brokers:

```sql
CREATE EXTENSION pipeline_kafka CASCADE;

SELECT pipeline_kafka.add_broker('kafka-1:9092');
SELECT pipeline_kafka.add_broker('kafka-2:9092');
```

The extension uses a fixed `pipeline_kafka` schema and stores brokers, consumer definitions, and offsets in its tables.

### Consume and Produce

```sql
SELECT pipeline_kafka.consume_begin(
  'events',
  'public.event_stream',
  format      => 'text',
  delimiter   => E'\t',
  batchsize   => 10000,
  parallelism => 4
);

SELECT * FROM pipeline_kafka.consumer_lag;

SELECT pipeline_kafka.produce_message(
  'notifications',
  convert_to('hello', 'UTF8'),
  partition => 2
);

SELECT pipeline_kafka.consume_end('events', 'public.event_stream');
```

Consumers target PipelineDB streams, not arbitrary modern tables. `consume_begin()` and `consume_end()` without arguments start or stop all recorded consumers. Specialized functions exist for partitioned streams and `topic_watermarks(topic, partition)`; `consumer_lag` compares saved offsets with broker watermarks.

To emit changed rows, define an `AFTER INSERT OR UPDATE FOR EACH ROW` trigger using `pipeline_kafka.emit_tuple('topic')`. It serializes the tuple as JSON. Trigger-produced messages and `produce_message` are external side effects and cannot be rolled back atomically with the PostgreSQL transaction.

### Legacy and Security Boundaries

Upstream specifies PipelineDB and librdkafka 0.11.6-era APIs. The project does not document current Kafka TLS/SASL configuration, idempotent production, exactly-once delivery, or modern consumer-group guarantees. Its side-effecting functions are even declared `IMMUTABLE` in the historical SQL, so do not rely on planner semantics intended for pure functions.

PipelineDB is discontinued and its PostgreSQL fork, stream catalogs, and background-worker APIs have diverged from supported PostgreSQL. Treat this extension as archival unless a maintained downstream port supplies a tested compatibility matrix, secure broker configuration, offset recovery, duplicate handling, and shutdown/rebalance testing.
