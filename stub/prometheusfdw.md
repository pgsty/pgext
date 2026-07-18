## Usage

Sources:

- [Pinned upstream README](https://github.com/alesharik/prometheus-fdw/blob/df0989e406128c3d7bab65a86f4004b8314a5956/README.md)
- [Pinned Cargo manifest](https://github.com/alesharik/prometheus-fdw/blob/df0989e406128c3d7bab65a86f4004b8314a5956/Cargo.toml)
- [Pinned extension control file](https://github.com/alesharik/prometheus-fdw/blob/df0989e406128c3d7bab65a86f4004b8314a5956/prometheusfdw.control)

`prometheusfdw` is a Rust/pgrx foreign data wrapper for reading instant-query results from Prometheus-compatible HTTP endpoints. The extension installs handler and validator functions; the wrapper, server, and foreign tables are created explicitly.

```sql
CREATE EXTENSION prometheusfdw;

CREATE FOREIGN DATA WRAPPER prometheus_wrapper
  HANDLER prometheus_fdw_handler
  VALIDATOR prometheus_fdw_validator;

CREATE SERVER metrics_server
  FOREIGN DATA WRAPPER prometheus_wrapper
  OPTIONS (address 'http://prometheus:9090/');

CREATE FOREIGN TABLE scrape_duration_seconds (
  __name__ text,
  instance text,
  job text,
  value double precision
)
SERVER metrics_server
OPTIONS (query 'scrape_duration_seconds');

SELECT * FROM scrape_duration_seconds LIMIT 20;
```

Foreign-table option `query` holds PromQL. The README documents `${var}` substitution from `WHERE` qualifiers, column `time timestamp` for instant-query time, and optional option `rate` for VictoriaMetrics. Range queries remain explicitly unfinished, so do not infer historical-window support from the FDW interface.

The pinned Cargo manifest and generated control line report `0.0.2`, with features for PostgreSQL 14, 15, and 16. Verify the exact artifact because other release packaging uses a different version line. Restrict server-definition privileges, prefer authenticated TLS endpoints, set query and statement timeouts, and test label-to-column mapping, missing samples, cardinality, remote errors, and planner behavior.
