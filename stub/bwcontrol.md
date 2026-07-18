## Usage

Sources:

- [Upstream README](https://github.com/kayform/bwcontrol/blob/5adfb169fb675a69ce3b9ea39275d59582e9ee6b/README.md)
- [Extension control file](https://github.com/kayform/bwcontrol/blob/5adfb169fb675a69ce3b9ea39275d59582e9ee6b/bwcontrol.control)
- [SQL installation script](https://github.com/kayform/bwcontrol/blob/5adfb169fb675a69ce3b9ea39275d59582e9ee6b/bwcontrol--1.0.sql)
- [C implementation](https://github.com/kayform/bwcontrol/blob/5adfb169fb675a69ce3b9ea39275d59582e9ee6b/bwcontrol.c)

`bwcontrol` is a PostgreSQL-side controller for a customized `Bottledwater` change-data-capture process and its `Kafka Connect` integration. It keeps a table whitelist in extension-owned metadata, starts or stops the external process, inspects its status, and updates connector configuration over HTTP. The control file and catalog both identify version `1.0`.

### Minimal status check

Configure the external paths and endpoints first, then create the extension and use the read-only status entry point:

```sql
CREATE EXTENSION bwcontrol;
SELECT pg_get_status_ingest();
```

The implementation reads `bw.bwpath`, `bw.kafka_broker`, `bw.schema_registry`, `bw.consumer`, and `bw.consumer_sub`. The upstream workflow also requires replication access, a customized Bottledwater build, Kafka, a schema registry, and Kafka Connect.

Whitelist management is exposed through `pg_add_ingest_table()`, `pg_del_ingest_table()`, `pg_add_ingest_column()`, and `pg_del_ingest_column()`. Process and connector operations use `pg_resume_ingest()`, `pg_suspend_ingest()`, `pg_make_kafka_connect()`, and `pg_delete_kafka_connect()`.

This is old, environment-specific C code with no current PostgreSQL compatibility matrix. The control functions launch shell commands, manage a replication slot, issue HTTP requests, and build dynamic SQL, yet the install script marks every function `IMMUTABLE` and does not revoke default execution rights. Treat all entry points as privileged administration APIs, restrict `EXECUTE`, validate every identifier and connector name, and test recovery outside production before use.
