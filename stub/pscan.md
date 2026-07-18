## Usage

Sources:

- [Pinned upstream README](https://github.com/MasahikoSawada/pscan/blob/fd5ff5c53cc80ea2625c27d469bf3c7a6a3eb232/README.md)
- [Pinned extension control file](https://github.com/MasahikoSawada/pscan/blob/fd5ff5c53cc80ea2625c27d469bf3c7a6a3eb232/pscan.control)

`pscan` is a PostgreSQL parallel heap-scan test tool. `p_tuplescan(regclass)` exercises the parallel tuple-scan path used by PostgreSQL 9.6 parallel sequential scan, while `p_brangescan(regclass)` divides the relation into consecutive block ranges assigned to workers. Overloads accept an explicit worker count, and GUC `pscan.n_workers` supplies the default.

```sql
CREATE EXTENSION pscan;
SET pscan.n_workers = 4;

SELECT p_tuplescan('public.fact_table'::regclass);
SELECT p_brangescan('public.fact_table'::regclass, 4);
```

These functions are not correctness-preserving replacements for SQL scans. The upstream README states that they skip visibility checks and filtering, fetch all tuples, and assume a garbage-free all-visible table. The block-range method can issue less-sequential I/O because workers read distant ranges concurrently.

Use version `1.0` only on disposable benchmark relations and verify the exact PostgreSQL build, because it depends on internal heap and parallel-worker APIs. Do not expose the functions to application roles, compare returned values with ordinary queries, or use timings without controlling cache state, table visibility, storage type, worker startup, and concurrent load.
