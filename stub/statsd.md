## Usage

Sources:

- [Official README](https://github.com/gmr/pg-statsd/blob/f9612228787f1366d6ad0b57252a1f8e221591fa/README.md)
- [Official SQL API](https://github.com/gmr/pg-statsd/blob/f9612228787f1366d6ad0b57252a1f8e221591fa/sql/statsd.sql)
- [Official UDP client implementation](https://github.com/gmr/pg-statsd/blob/f9612228787f1366d6ad0b57252a1f8e221591fa/src/statsd.c)
- [Official PGXN release](https://pgxn.org/dist/statsd/0.2.0/)

`statsd` version `0.2.0` sends counters, gauges, and timing metrics directly from a PostgreSQL backend to a StatsD endpoint over UDP. It is an archived 2014 client with fire-and-forget network behavior; use it only where losing metrics is acceptable and database roles cannot abuse outbound access.

### Core Workflow

Install the extension and emit metrics to a reachable StatsD listener:

```sql
CREATE EXTENSION statsd;

SELECT statsd.add_timing('127.0.0.1', 8125, 'checkout.duration_ms', 70);
SELECT statsd.increment_counter('127.0.0.1', 8125, 'checkout.completed');
SELECT statsd.increment_counter('127.0.0.1', 8125, 'checkout.items', 5);
SELECT statsd.increment_counter('127.0.0.1', 8125, 'checkout.sampled', 5, 0.25);
SELECT statsd.set_gauge('127.0.0.1', 8125, 'workers.busy', 3);
```

All functions return void. The host and port are supplied on every call rather than through extension configuration.

### Function Forms

- `statsd.add_timing(host, port, metric, milliseconds)` sends an `ms` metric.
- `statsd.increment_counter(host, port, metric)` increments by one.
- `statsd.increment_counter(host, port, metric, value)` sends an explicit counter delta.
- `statsd.increment_counter(host, port, metric, value, sample)` adds a sample rate.
- `statsd.set_gauge(host, port, metric, value)` has integer and double-precision overloads.

### Operational Caveats

Each call performs DNS resolution, opens a UDP socket, and sends one datagram from the querying backend. UDP provides no delivery acknowledgement, and the implementation reports resolution or send failures as server warnings while still returning void, so SQL success does not prove metric delivery. The SQL incorrectly declares these network-side-effecting functions `IMMUTABLE`; never use them in indexes, generated columns, constraints, or expressions the planner may fold. Revoke broad EXECUTE privileges to prevent arbitrary DNS lookups, network egress, or metric injection. Metric names are not sanitized. Prefer a maintained asynchronous exporter or log/telemetry pipeline for production workloads.
