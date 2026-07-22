## Usage

Sources:

- [Official README and protocol documentation](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/README.rst)
- [Extension SQL](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/sql/replisome.sql)
- [Output-plugin implementation](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/src/replisome.c)

`replisome` 0.1 is a pre-production logical-decoding output plugin that streams INSERT, UPDATE, and DELETE changes as JSON. It can select tables, columns, and rows and has an optional Python consumer framework. It is a change-integration prototype, not a complete replication solution.

### Configure Logical Decoding

```conf
wal_level = logical
max_replication_slots = 1
max_wal_senders = 1
```

Restart after changing these settings, configure a least-privilege replication entry in `pg_hba.conf`, install the output-plugin library, and optionally register the SQL extension. The SQL extension only exposes `replisome_version()` for protocol checks.

```sql
CREATE EXTENSION replisome;
SELECT replisome_version();

SELECT *
FROM pg_create_logical_replication_slot('integration_slot', 'replisome');
```

Consume with the replication protocol, for example:

```sh
pg_recvlogical -d appdb --slot integration_slot --start \
  -o pretty-print=1 -f -
```

Each committed transaction is represented as JSON with operation code, schema/table, values, and old key data where available.

### Filtering and Output Options

Repeated `include` and `exclude` JSON options are processed in order, with the last match winning. Includes can match table/schema names or regular expressions, select/skip columns, and apply a row `where` expression. Other options include `include-xids`, `include-lsn`, `include-timestamp`, `include-schemas`, `include-types`, `include-empty-xacts`, and `write-in-chunks`. Chunked output may not be valid JSON until the consumer reassembles it.

### Reliability and Compatibility Boundaries

Replication slots retain WAL until a consumer advances them; monitor lag and disk use, and drop abandoned slots deliberately. Consumers must be idempotent because reconnect/retry can replay data. Filters and omitted columns can make output unsuitable for reconstructing rows, and row filters must be tested across UPDATE/DELETE key changes.

The plugin does not handle TRUNCATE, DDL, sequences, schema evolution, or conflict resolution. Upstream labels it pre-production, targets PostgreSQL 9.4-era logical decoding, and documents a Python 2.7 consumer. Its old private server APIs require a maintained port for current PostgreSQL. Do not use it as disaster recovery or lossless replication without full protocol, crash, upgrade, failover, and replay testing.
