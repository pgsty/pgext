## Usage

Sources:

- [Project README](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/README.md)
- [Extension control file](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/pg_stat_sysinfo.control)
- [Package metadata](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/Cargo.toml)
- [SQL functions and view](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/src/lib.rs)
- [Collector implementation](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/src/collector.rs)
- [Background cache worker](https://github.com/postgresml/pg_stat_sysinfo/blob/b92c5161c22272aa955ad18e3d2e091d3ca6948c/src/cache_worker.rs)

`pg_stat_sysinfo` 0.0.1 is a pgrx system-metrics prototype. It reports load averages, CPU use, memory and swap capacity and use, and mounted-disk capacity. `pg_stat_sysinfo_collect()` samples synchronously; an optional background worker stores recent samples in a shared-memory ring buffer exposed through the `pg_stat_sysinfo` view.

### Direct and cached collection

Direct collection does not require preloading:

```sql
CREATE EXTENSION pg_stat_sysinfo;

SELECT * FROM pg_stat_sysinfo_collect();
```

For cached collection, configure the library and interval, then restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_stat_sysinfo.so'
pg_stat_sysinfo.interval = '1s'
```

```sql
SELECT metric, dimensions, at, value
FROM pg_stat_sysinfo
ORDER BY at DESC;

SELECT * FROM pg_stat_sysinfo_cache_summary();
```

The fixed cache is about 1280 KiB, so retention varies with the number of filesystems and collection interval. A configuration reload can change the interval and periodically refresh the disk list.

### Visibility, overhead, and compatibility

Metrics reveal host load, memory, swap, filesystem mount names, and capacities. Review default function execution privileges and grant access only to monitoring roles. Direct calls refresh operating-system counters and may briefly wait for the system library's minimum CPU sampling interval; aggressive polling adds backend and filesystem-enumeration work.

The reviewed Cargo features cover PostgreSQL 11 through 15 with pgrx 0.8.3, and the last code change is from 2023. There is no evidence for current majors or containers with unusual mount namespaces. Values are point-in-time host observations, not PostgreSQL attribution and not durable monitoring. Validate units, cgroup visibility, disk discovery, failover behavior, and sampling overhead, then export important metrics to a maintained monitoring system.
