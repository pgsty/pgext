

## Usage

> [pg_tracing: distributed tracing for PostgreSQL](https://github.com/DataDog/pg_tracing)

pg_tracing generates server-side spans for distributed tracing. It creates spans for planner, executor, plan nodes, nested queries, triggers, parallel workers, and transaction commits.

### Trace Propagation via SQLCommenter

Pass trace context as a SQL comment. Queries with a sampled flag will generate spans:

```sql
/*traceparent='00-00000000000000000000000000000123-0000000000000123-01'*/ SELECT 1;
```

### Trace Propagation via GUC

```sql
BEGIN;
SET LOCAL pg_tracing.trace_context='traceparent=''00-00000000000000000000000000000005-0000000000000005-01''';
UPDATE pgbench_accounts SET abalance=1 WHERE aid=1;
COMMIT;
```

### Standalone Sampling

Randomly sample queries without external trace context:

```sql
SET pg_tracing.sample_rate = 1.0;  -- trace all queries
SELECT 1;
```

### Viewing Spans

```sql
-- Consume spans (removes them from buffer)
SELECT trace_id, parent_id, span_id, span_start, span_end, span_type, span_operation
FROM pg_tracing_consume_spans ORDER BY span_start;

-- Peek at spans (non-destructive)
SELECT * FROM pg_tracing_peek_spans;

-- Export as OTLP JSON
SELECT pg_tracing_json_spans();
```

### Statistics

```sql
SELECT * FROM pg_tracing_info;
SELECT pg_tracing_reset();
```

### Sending Spans to OpenTelemetry Collector

Configure in `postgresql.conf`:

```
pg_tracing.otel_endpoint = http://127.0.0.1:4318/v1/traces
pg_tracing.otel_naptime = 2000
```

### Key GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_tracing.max_span` | 10000 | Maximum spans in shared memory |
| `pg_tracing.track` | `all` | Which statements to track |
| `pg_tracing.sample_rate` | 0 | Fraction of queries to sample randomly |
| `pg_tracing.otel_endpoint` | (empty) | OTLP HTTP endpoint for span export |
