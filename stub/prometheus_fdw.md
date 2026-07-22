## Usage

Sources:

- [Official documentation](https://github.com/Jayko001/prometheus_fdw/blob/d9794a9b89bf00e65bf9d31475b740d7f5fdf9f0/docs/index.md)
- [Official FDW implementation](https://github.com/Jayko001/prometheus_fdw/blob/d9794a9b89bf00e65bf9d31475b740d7f5fdf9f0/src/lib.rs)
- [Official PGXN release](https://pgxn.org/dist/prometheus_fdw/0.2.1/)

`prometheus_fdw` is a read-only foreign data wrapper that turns a Prometheus range-query response into PostgreSQL rows. Version `0.2.1` supports one `metrics` object and requires exact metric-name and time-bound predicates to construct the Prometheus request.

### Configure and Query

```sql
CREATE EXTENSION prometheus_fdw;

CREATE FOREIGN DATA WRAPPER prometheus_wrapper
  HANDLER prometheus_fdw_handler
  VALIDATOR prometheus_fdw_validator;

CREATE SERVER my_prometheus_server
  FOREIGN DATA WRAPPER prometheus_wrapper
  OPTIONS (base_url 'https://prometheus.example.com');

CREATE FOREIGN TABLE metrics (
  metric_name text,
  metric_labels jsonb,
  metric_time bigint,
  metric_value float8
)
SERVER my_prometheus_server
OPTIONS (object 'metrics', step '10m');

SELECT metric_name, metric_labels, metric_time, metric_value
FROM metrics
WHERE metric_name = 'container_cpu_usage_seconds_total'
  AND metric_time > 1696046800
  AND metric_time < 1696133000;
```

The three predicates are operationally required by the implementation: `metric_name =`, `metric_time >`, and `metric_time <`. Missing them produces no valid request. `step` is passed to Prometheus and defaults to `10m`.

### Returned Columns

- `metric_name` repeats the requested metric name.
- `metric_labels` contains the Prometheus series label object.
- `metric_time` is the Unix timestamp returned by the range query.
- `metric_value` is the sample parsed as `float8`.

### Operational Boundaries

Each scan performs a synchronous outbound HTTP request with a 30-second client timeout and materializes all returned samples in backend memory; SQL sort, limit, and most filters are not pushed down. Restrict the endpoint, query window, step, row count, statement timeout, and network egress. Although the source defines username, password, and bearer-token fields, the reviewed constructor does not populate them from server or user-mapping options, so do not assume authenticated Prometheus works without verifying the packaged build. It also unwraps JSON parsing and does not explicitly reject HTTP error status. Test malformed/error responses, TLS trust, planner rescans, parallel queries, timestamp boundary semantics, large series sets, and credential handling before production use.
