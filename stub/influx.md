## Usage

Sources:

- [influx.control](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/influx.control)
- [README.md](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/README.md)
- [docs/getting-started.md](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/docs/getting-started.md)
- [docs/options.md](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/docs/options.md)
- [docs/procedures.md](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/docs/procedures.md)
- [influx--0.4.sql](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/influx--0.4.sql)

`influx` receives InfluxDB line-protocol datagrams and writes PostgreSQL tables through background workers. Upstream explicitly describes it as an educational experiment; UDP delivery can lose data.

### Core Workflow

```sql
CREATE SCHEMA metrics;
CREATE EXTENSION influx WITH SCHEMA metrics;
SELECT metrics.worker_launch('8089');
CALL metrics.send_packet('cpu,host=demo usage=1.5 1574753954000000000', '8089');
```

### Operational Boundaries

Install as a superuser. A manually launched worker does not need startup preload. For automatic workers, configure `shared_preload_libraries`, `influx.database`, `influx.schema`, `influx.role` and `influx.workers`, then restart. The role defaults to superuser; select a restricted role for ingestion. `worker_launch` returns a worker PID, `send_packet` sends test data, and `_create` controls automatic table creation. The listener binds a passive UDP address rather than an explicitly selected interface; limit network access. The documented development target is PostgreSQL 13.
