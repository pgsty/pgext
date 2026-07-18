## Usage

Sources:

- [Project README](https://github.com/tvondra/query_histogram/blob/45c3f2e24ae67d81354c10a4f638ac43bfa18d19/README.md)
- [Extension control file](https://github.com/tvondra/query_histogram/blob/45c3f2e24ae67d81354c10a4f638ac43bfa18d19/query_histogram.control)
- [Version 1.1 SQL API](https://github.com/tvondra/query_histogram/blob/45c3f2e24ae67d81354c10a4f638ac43bfa18d19/sql/query_histogram--1.1.sql)
- [Executor and shared-memory implementation](https://github.com/tvondra/query_histogram/tree/45c3f2e24ae67d81354c10a4f638ac43bfa18d19/src)

`query_histogram` 1.1 records cluster-wide histograms of statement and transaction duration in shared memory. It reports counts and accumulated time by duration bin, but does not retain query identity or text.

### Preload and read histograms

Configure the library and collection parameters, then restart PostgreSQL:

```conf
shared_preload_libraries = 'query_histogram'
query_histogram.sample_pct = 5
query_histogram.bin_width = 10
query_histogram.bin_count = 1000
query_histogram.dynamic = true
```

Install the SQL objects and query the views:

```sql
CREATE EXTENSION query_histogram;

SELECT * FROM query_histogram;
SELECT * FROM xact_histogram;
SELECT query_histogram_get_reset();
```

`query_histogram.bin_width` is measured in milliseconds, and `query_histogram.bin_count` is capped at 1000. A zero bin count disables collection but leaves hooks installed. Dynamic mode permits runtime changes at the cost of additional shared-memory locking.

### Sampling, resets, and overhead

Each sampled execution updates a shared segment protected by a System V semaphore. Start with a low `query_histogram.sample_pct`, benchmark representative concurrency, and remember that sampled counts are not exact totals. Bin boundaries also discard distribution detail, and the extension cannot identify which normalized statement caused a slow bin.

`query_histogram_reset()` clears shared observations for all users, so restrict execution and coordinate collectors before resetting. The SQL declares state-reading and reset functions `IMMUTABLE` even though their results depend on mutable shared state; do not rely on planner caching semantics around them. The source predates current PostgreSQL majors and publishes no current compatibility matrix, so compile and load-test the exact target version and every other executor-hook extension.
