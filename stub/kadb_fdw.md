## Usage

Sources:

- [Upstream README](https://github.com/arenadata/kafka-adb-os/blob/23400f5b2ef5e1fe8bb50e3b0c859af9ee4dcb8f/README.md)
- [Detailed usage guide](https://github.com/arenadata/kafka-adb-os/blob/23400f5b2ef5e1fe8bb50e3b0c859af9ee4dcb8f/USAGE.md)
- [Version 0.10.2 control file](https://github.com/arenadata/kafka-adb-os/blob/23400f5b2ef5e1fe8bb50e3b0c859af9ee4dcb8f/kadb_fdw.control)
- [Offset storage implementation](https://github.com/arenadata/kafka-adb-os/blob/23400f5b2ef5e1fe8bb50e3b0c859af9ee4dcb8f/src/offsets.c)

kadb_fdw 0.10.2 is an archived Kafka foreign data wrapper built specifically for Arenadata DB and Greenplum Database. It distributes Kafka partitions across MPP segments and stores consumer offsets in kadb.offsets.

### ADB or Greenplum example

Use only on a compatible disposable MPP cluster with a test topic:

```sql
CREATE EXTENSION kadb_fdw;
CREATE SERVER ka_server
FOREIGN DATA WRAPPER kadb_fdw
OPTIONS (k_brokers 'localhost:9092');
CREATE FOREIGN TABLE ka_table(payload text)
SERVER ka_server
OPTIONS (
  format 'text',
  k_topic 'test_topic',
  k_consumer_group 'test_group',
  k_seg_batch '5',
  k_timeout_ms '1000'
);
SELECT * FROM ka_table;
```

A successful SELECT advances the offsets stored for later reads. It is a consuming operation rather than an idempotent view of a remote table.

### Caveats

- The install and upgrade SQL uses Greenplum-specific syntax such as distributed tables, segment execution, and MPP options. It is not a standard PostgreSQL FDW and will not install unchanged on community PostgreSQL.
- LIMIT does not limit Kafka retrieval: the wrapper fetches batches and advances offsets through the last retrieved message, including messages omitted from the returned rows.
- Offset helper functions do not give transactional guarantees in Kafka, and several partition synchronization operations are explicitly non-atomic. Coordinate retries and exactly-once expectations outside this FDW.
- Some polling phases cannot be cancelled before the configured timeout, and total query duration can be several timeout intervals across partitions.
- A user-supplied Avro schema is not validated against messages; upstream documents mismatches reaching invalid memory-allocation errors.
- Kerberos support uses SASL plaintext and requires the same server-readable keytab path on every segment. Protect server options, catalogs, files, and logs.
- The repository was archived in 2021 and targets old ADB, Greenplum, librdkafka, Avro, and CSV library versions. Its install script also resets the session client_min_messages setting.
