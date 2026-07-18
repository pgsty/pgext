## Usage

Sources:

- [Project README and check coverage](https://github.com/tvondra/scrub/blob/8d39b4a2c7419715c9b7726c5d7b478ee5e0a9c6/README.md)
- [Extension control file](https://github.com/tvondra/scrub/blob/8d39b4a2c7419715c9b7726c5d7b478ee5e0a9c6/scrub.control)
- [Version 1.0 SQL API](https://github.com/tvondra/scrub/blob/8d39b4a2c7419715c9b7726c5d7b478ee5e0a9c6/scrub--1.0.sql)
- [Background-worker implementation](https://github.com/tvondra/scrub/blob/8d39b4a2c7419715c9b7726c5d7b478ee5e0a9c6/scrub.c)

`scrub` 1.0 starts a background worker that reads database pages and checks for corruption. It can verify page checksums when data checksums are enabled, generic page-header bounds, heap pages and tuples including TOAST values, and B-tree pages and tuples. It reports findings but does not repair damage.

### Preload and start a throttled run

Add the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'scrub'
```

Install the extension, start one database, and monitor counters:

```sql
CREATE EXTENSION scrub;

SELECT scrub_start(
  dbname => 'appdb',
  cost_delay => 10,
  cost_limit => 1000,
  reset_status => true
);

SELECT * FROM scrub_status();
SELECT scrub_is_running();
```

The cost settings throttle work similarly to vacuum cost controls. Only one scrub can run across the cluster at a time. `scrub_stop()` requests termination, while `scrub_reset()` clears counters; status otherwise accumulates across runs.

### Detection boundary and incident handling

The worker reads all selected pages without intentionally writing them, but it can still generate substantial storage I/O, consume cache, and compete with backups, vacuum, replicas, and foreground queries. Establish an I/O budget, monitor latency and replica lag, and schedule initial runs outside peak periods.

Details of detected failures are written to the PostgreSQL server log. Preserve those logs and storage evidence, stop avoidable writes, and follow a tested corruption-response procedure; do not assume a failed counter identifies the only damaged object. The checker does not cover every access method, cross-check every heap row against every index, or prove the absence of latent corruption. Its low-level page code has no current PostgreSQL compatibility matrix, so test the exact major, checksum setting, storage stack, and backup workflow before production use.
