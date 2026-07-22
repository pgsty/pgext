## Usage

Sources:

- [Official extension SQL](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/pg_oltp_bench--1.0.sql)
- [Official initialization script](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/oltp_init.sql)
- [Official read-only workload](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/oltp_ro.sql)
- [Official read-write workload](https://github.com/postgrespro/pg_oltp_bench/blob/fbf92a58e12f6e155a4d2c09969063dc26e8a2a1/oltp_rw.sql)

`pg_oltp_bench` combines one random-string helper with `pgbench` scripts that approximate a sysbench-style OLTP workload. It is a benchmark fixture, not an application schema or a standardized score by itself.

### Prepare the Fixture

Create the extension, then run the upstream initialization script in a disposable benchmark database:

```sql
CREATE EXTENSION pg_oltp_bench;

SELECT sb_rand_str('order-########-@@@@');
```

In `sb_rand_str(text)`, each ASCII `#` is replaced with a random digit and each ASCII `@` with a random lowercase letter; other bytes are preserved. The shipped `oltp_init.sql` drops and recreates `sbtest`, inserts ten million rows, and builds an index on `k`.

```bash
psql "$BENCHMARK_URL" -f oltp_init.sql
pgbench "$BENCHMARK_URL" -n -c 16 -j 8 -T 300 -f oltp_ro.sql
pgbench "$BENCHMARK_URL" -n -c 16 -j 8 -T 300 -f oltp_rw.sql
```

`oltp_ro.sql` runs ten point lookups and four range/aggregate/order queries per transaction script. `oltp_rw.sql` adds updates, a delete, and a replacement insert inside an explicit transaction.

### Benchmark Discipline

- `oltp_init.sql` begins with `DROP TABLE IF EXISTS sbtest`; never point it at a database containing a table you need to preserve.
- The scripts hard-code `table_size` as ten million and `range_size` as 100. If the fixture size changes, update the scripts consistently or random IDs and range bounds no longer describe the data set.
- Record PostgreSQL settings, hardware, client count, threads, duration, warm-up, data scale, durability, and whether data fits in memory. Results without those controls are not comparable.
- Run read-only and read-write workloads separately, monitor saturation and errors, and validate table state after failures. The read-write script can leave workload-dependent churn and consumes WAL and storage.
- `sb_rand_str` is volatile and uses the extension's simple pseudo-random generation. It is suitable for benchmark filler, not secrets, identifiers, or cryptographic material.
