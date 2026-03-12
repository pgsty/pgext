


## Usage

> [kafka_fdw: Kafka Foreign Data Wrapper for CSV formatted messages](https://github.com/adjust/kafka_fdw)

### Create Server

```sql
CREATE EXTENSION kafka_fdw;

CREATE SERVER kafka_server FOREIGN DATA WRAPPER kafka_fdw
  OPTIONS (brokers 'localhost:9092');
```

**Server Options:** `brokers` (required, comma-separated Kafka broker endpoints).

### Create User Mapping

```sql
CREATE USER MAPPING FOR PUBLIC SERVER kafka_server;
```

### Create Foreign Table (CSV Format)

```sql
CREATE FOREIGN TABLE kafka_csv (
  part int OPTIONS (partition 'true'),
  offs bigint OPTIONS (offset 'true'),
  some_int int,
  some_text text,
  some_date date,
  some_time timestamp
)
SERVER kafka_server
OPTIONS (format 'csv', topic 'my_topic', batch_size '30', buffer_delay '100');
```

Two metadata columns are required: one with `partition 'true'` and one with `offset 'true'`. The remaining columns match the message format.

**Table Options:** `format` (`csv` or `json`), `topic` (Kafka topic name), `batch_size`, `buffer_delay` (milliseconds), `strict` (enforce strict schema validation), `ignore_junk` (set malformed columns to NULL).

### Create Foreign Table (JSON Format)

```sql
CREATE FOREIGN TABLE kafka_json (
  part int OPTIONS (partition 'true'),
  offs bigint OPTIONS (offset 'true'),
  some_int int OPTIONS (json 'int_val'),
  some_text text OPTIONS (json 'text_val')
)
SERVER kafka_server
OPTIONS (format 'json', topic 'my_json_topic', batch_size '30', buffer_delay '100');
```

Use the `json` column option to map column names to JSON keys.

### Consuming Messages

```sql
-- Read from a specific partition and offset
SELECT * FROM kafka_csv WHERE part = 0 AND offs > 1000 LIMIT 60;

-- Read from multiple partitions
SELECT * FROM kafka_csv
WHERE (part = 0 AND offs > 100) OR (part = 1 AND offs > 300);
```

Note: The `offset` keyword is reserved in SQL; use double quotes when referencing the offset column in some contexts.

### Producing Messages

```sql
-- Insert with explicit partition
INSERT INTO kafka_csv (part, some_int, some_text)
  VALUES (0, 42, 'hello from partition 0');

-- Insert with auto-partition selection
INSERT INTO kafka_csv (some_int, some_text)
  VALUES (42, 'auto-partitioned message');
```
