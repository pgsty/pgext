## Usage

Sources:

- [PGXN 0.2.0 README](https://pgxn.org/dist/pg_recall_guard/0.2.0/README.html)
- [pg_recall_guard control file](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.0/pg_recall_guard.control)
- [pg_recall_guard 0.2.0 SQL definitions](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.0/pg_recall_guard--0.2.0.sql)
- [PostgreSQL license](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.0/LICENSE)

`pg_recall_guard` measures approximate-nearest-neighbor index recall against an exact baseline and records the value an operator approved. It discovers indexable distance-ordering operators from PostgreSQL catalogs instead of hard-coding one vector extension.

### Core Workflow

```sql
CREATE EXTENSION pg_recall_guard CASCADE;

SELECT index_name, table_name, access_method, operator
FROM recall_guard.vector_indexes;

SELECT recall_guard.approve(
    'items_hnsw',
    p_k => 10,
    p_sample_size => 30
);

SELECT * FROM recall_guard.check();
```

`CASCADE` installs the required `tsm_system_rows` contrib extension when available. `recall_guard.approve(...)` measures the selected index and stores an accepted baseline. `recall_guard.check()` remeasures all baselines and returns the baseline, current recall, drift, and verdict.

### Objects and Measurement

- `recall_guard.vector_indexes` lists indexes that expose an ordering operator through `pg_amop`.
- `recall_guard.measure(...)` samples query vectors and compares indexed nearest neighbors with an exact sequential-scan result.
- `recall_guard.baselines` stores approved settings and recall values.
- `recall_guard.measurements` stores later observations.
- `recall_guard.evaluate_query(...)` is the lower-level comparison helper.

The implementation verifies that the approximate side actually uses the target index and that the ground-truth side does not. It excludes each sampled row's self-match so a trivial zero-distance hit cannot inflate recall.

### Cost and Boundaries

Each sampled query runs an exact scan, so measurement can be expensive on large tables. Start with a small `p_sample_size`, schedule checks away from peak load, and observe I/O and CPU. Samples come from the indexed table rather than production query traffic; they are a proxy, not a workload trace.

PGXN metadata claims PostgreSQL 13 or newer, while upstream testing covers 18.6 and 19 beta. PostgreSQL 13 through 17 remain untested. Version 0.2.0 has not been field-validated on a production workload, and its published experiments cover a limited set of vector dimensions and index settings. Treat thresholds as workload-specific and approve a baseline only after independent quality testing.
