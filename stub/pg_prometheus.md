## Usage

Sources:

- [Archived upstream README and sunset notice](https://github.com/timescale/pg_prometheus/blob/b5a113ef014597e9adcb1021d40bba005a8d741a/README.md)
- [Pinned extension SQL](https://github.com/timescale/pg_prometheus/blob/b5a113ef014597e9adcb1021d40bba005a8d741a/sql/prometheus.sql)
- [Pinned sample parser](https://github.com/timescale/pg_prometheus/blob/b5a113ef014597e9adcb1021d40bba005a8d741a/src/parse.c)
- [Pinned extension control file](https://github.com/timescale/pg_prometheus/blob/b5a113ef014597e9adcb1021d40bba005a8d741a/pg_prometheus.control)

pg_prometheus version 0.2.2 defines a prom_sample type for Prometheus exposition-format samples and creates raw or normalized PostgreSQL storage with query views. It was designed for the old Prometheus PostgreSQL remote-storage adapter, with optional TimescaleDB hypertables.

### Create and query storage

Upstream requires the library in shared_preload_libraries and a server restart before installation:

```conf
shared_preload_libraries = 'pg_prometheus'
```

```sql
CREATE EXTENSION pg_prometheus;

SELECT create_prometheus_table(
    'metrics',
    normalized_tables => true,
    use_timescaledb => false
);

INSERT INTO metrics (sample)
VALUES ('cpu_usage{service="nginx",host="machine1"} 34.6 1494595898000');

SELECT time, value, labels
FROM metrics
WHERE name = 'cpu_usage'
ORDER BY time DESC;
```

Normalized storage separates label sets from timestamp/value rows; raw storage keeps each complete prom_sample. The helper creates several tables, indexes, views, and triggers in the caller's schema. COPY integrations use the generated copy table rather than the view.

### Archived integration

Timescale archived this repository after putting it in sunset and maintenance mode. The reviewed source last changed in 2020, its Docker example targets PostgreSQL 11, and both the adapter and TimescaleDB integration reflect that era. The README points to a successor project, but that does not make this extension current.

Treat pg_prometheus as a legacy data format and migration source. Validate its C parser with untrusted exposition input, inspect all generated object names and privileges, define explicit retention and partition management, and measure label-cardinality growth. Do not deploy the old adapter or archived container without a complete security, compatibility, and migration review.
