## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/prometheus_fdw/)
- [Official project or provider page](https://tembo.io/blog/monitoring-with-prometheus-fdw)

`prometheus_fdw` — A Rust/pgrx foreign data wrapper that queries Prometheus metrics through PostgreSQL foreign tables.

The reviewed catalog snapshot records version `0.2.1`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "prometheus_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'prometheus_fdw';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
