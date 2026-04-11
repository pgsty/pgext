
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION kafka_fdw;
> CREATE SERVER kafka_server FOREIGN DATA WRAPPER kafka_fdw
>   OPTIONS (brokers 'localhost:9092');
> ```
>
> Source: [README](https://github.com/adjust/kafka_fdw)

`kafka_fdw` is a foreign data wrapper that exposes Kafka messages as PostgreSQL foreign tables. The upstream README explicitly warns that the project is not yet production ready.

### Server and Mapping

Define a foreign server with the Kafka broker list, then add a user mapping:

```sql
CREATE EXTENSION kafka_fdw;

CREATE SERVER kafka_server
FOREIGN DATA WRAPPER kafka_fdw
OPTIONS (brokers 'localhost:9092');

CREATE USER MAPPING FOR PUBLIC SERVER kafka_server;
```

## Foreign Tables

Kafka foreign tables must declare two metadata columns, one marked with `partition 'true'` and one marked with `offset 'true'`. The remaining columns describe the message payload.

### CSV Messages

```sql
CREATE FOREIGN TABLE kafka_test (
    part int OPTIONS (partition 'true'),
    offs bigint OPTIONS (offset 'true'),
    some_int int,
    some_text text,
    some_date date,
    some_time timestamp
)
SERVER kafka_server
OPTIONS (
    format 'csv',
    topic 'contrib_regress',
    batch_size '30',
    buffer_delay '100'
);
```

For CSV, columns are mapped by position. Upstream notes that schema enforcement depends on the message writer, so strict parsing and junk-handling options matter when input quality is uncertain.

### JSON Messages

```sql
CREATE FOREIGN TABLE kafka_test_json (
    part int OPTIONS (partition 'true'),
    offs bigint OPTIONS (offset 'true'),
    some_int int OPTIONS (json 'int_val'),
    some_text text OPTIONS (json 'text_val'),
    some_date date OPTIONS (json 'date_val'),
    some_time timestamp OPTIONS (json 'time_val')
)
SERVER kafka_server
OPTIONS (
    format 'json',
    topic 'contrib_regress_json',
    batch_size '30',
    buffer_delay '100'
);
```

For JSON, each column can map to an object key with the `json` option. The current implementation supports JSON objects, not top-level JSON arrays.

## Querying and Producing

The offset and partition columns are special, and the upstream README recommends specifying them in queries whenever possible:

```sql
SELECT * FROM kafka_test WHERE part = 0 AND offs > 1000 LIMIT 60;

SELECT *
FROM kafka_test
WHERE (part = 0 AND offs > 100)
   OR (part = 1 AND offs > 300)
   OR (part = 3 AND offs > 700);
```

Messages can also be produced with `INSERT` statements. If a partition value is supplied, it is used; otherwise Kafka's builtin partitioner chooses one:

```sql
INSERT INTO kafka_test(part, some_int, some_text)
VALUES
    (0, 5464565, 'some text goes into partition 0'),
    (NULL, 5464565, 'some text goes into partition selected by kafka');
```

## Error Handling

The default behavior is permissive:

- missing trailing columns are treated as `NULL`
- extra fields are ignored
- unparsable values still raise errors by default

Relevant table options and helper columns include:

- `strict 'true'` to reject column count mismatches
- `ignore_junk 'true'` to set malformed values to `NULL`
- columns marked `junk 'true'` to capture the original payload
- columns marked `junk_error 'true'` to capture parsing errors

## Build Notes

The extension uses `librdkafka` and the upstream build instructions are the standard:

```bash
make && make install
make installcheck
```

The test setup assumes Kafka on `localhost:9092` and ZooKeeper on `localhost:2181`.
