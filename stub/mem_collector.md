## Usage

Sources:

- [Official source](https://github.com/aotenjou/mem_collector/blob/ed19a382495eeffe279497820340a2eb8a78b978/README.md)
- [Official source](https://github.com/aotenjou/mem_collector/blob/ed19a382495eeffe279497820340a2eb8a78b978/mem_collector.control)
- [Official source](https://github.com/aotenjou/mem_collector/blob/ed19a382495eeffe279497820340a2eb8a78b978/mem_collector--1.0.sql)

`mem_collector` 1.0 collects query memory features and window summaries for a research memory-tuning pipeline. Upstream targets PostgreSQL 12. Planner/executor hooks and shared-memory buffers require preloading and a server restart.

### Configure and Read Samples

Configure the database in which the worker persists samples:

```ini
shared_preload_libraries = 'mem_collector'
mem_collector.window_seconds = 30
mem_collector.enable_bgw = on
mem_collector.bgw_dbname = 'postgres'
```

After restarting, install as a superuser in that database:

```sql
CREATE EXTENSION mem_collector;
SELECT * FROM mem_collector_features();
SELECT * FROM mem_collector_pending_samples();
SELECT mem_collector_flush_query_samples(1024);
SELECT * FROM mem_query_samples ORDER BY id DESC LIMIT 5;
SELECT * FROM mem_window_samples ORDER BY id DESC LIMIT 5;
```

The versioned SQL implements the first two read functions as recent persisted-table views; their names do not establish a direct view of unflushed ring contents. `mem_collector_generate_window_sample` creates an explicit window sample, and `mem_collector_reset` resets collection state. Protect access to collected query/workload metadata and schedule retention for the persistent sample tables.

### Interpretation and Operations

Features cover sort/hash operators and simplified parallel ratios. Peak memory and spill signals are estimates derived from public executor structures. Several pressure and pool-usage fields remain unimplemented, and ratios use host physical memory rather than a PostgreSQL memory-pool budget. This is not a production memory accounting authority or an automatic memory limiter.

Shared-memory rings are lost on restart; persisted sample tables survive. A worker aimed at another database needs the SQL objects there. Check that samples actually advance after enabling collection and after upgrades, and monitor table growth and collection overhead.
