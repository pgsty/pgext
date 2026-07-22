## Usage

Sources:

- [Official README](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/README.md)
- [Version 1.0 SQL API](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/pgotel--1.0.sql)
- [GUC and SQL-call implementation](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/pgotel.cpp)
- [OTLP gRPC metric exporter implementation](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/metrics_grpc.cpp)
- [Native link dependencies](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/Makefile)

`pgotel` 1.0 is a proof-of-concept C/C++ bridge to the OpenTelemetry C++ SDK. The current SQL surface emits double-counter metrics through an OTLP gRPC exporter. It does not expose trace or log APIs, despite the broader project description.

### Configure and Emit a Counter

The native library links the OpenTelemetry metrics SDK and OTLP gRPC exporter. Configure the collector endpoint before the library is first loaded in a backend, then create the extension and emit nonnegative counter values:

```conf
pgotel.endpoint = 'localhost:4317'
pgotel.interval = 2000
pgotel.timeout = 500
```

```sql
CREATE EXTENSION pgotel;

SELECT pgotel_counter('requests', 1.0);
SELECT pgotel_counter(
  'rows_processed',
  42.0,
  '{"database":"app","worker":"import-1"}'::jsonb
);
```

The two overloads are `pgotel_counter(text, float8)` and `pgotel_counter(text, float8, jsonb)`. Labels should be a nonempty, flat JSON object with string values.

### Current Implementation Boundaries

The README discusses choosing HTTP transport and shows a three-argument call with an endpoint, but the reviewed code actually links and uses the OTLP gRPC metrics exporter, and the endpoint is a GUC rather than a function argument. Follow the SQL and C++ implementation for version 1.0.

`pgotel.enabled` is registered but never consulted by the metric path, so changing it does not disable emission in this source. Changing `pgotel.endpoint` has no assignment hook; set it before first library initialization. Changes to `pgotel.interval` or `pgotel.timeout` rebuild the per-backend metric provider.

Label iteration assumes string values and does not safely handle arbitrary nested, numeric, boolean, or null JSON. Empty labels are rejected. Counter names and labels can create unbounded metric cardinality, while a fresh counter instrument is created on every call; restrict both to a small controlled vocabulary.

Metric export is asynchronous external I/O and is not transactional: a metric may be exported even if the surrounding SQL transaction rolls back, or lost if a backend exits before a periodic export. The extension requires a compatible C++ ABI plus OpenTelemetry/gRPC shared libraries. Pin those dependencies, test collector outage and backend shutdown, and use it only as experimental telemetry rather than durable accounting.
