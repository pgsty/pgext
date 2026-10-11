## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.9/README.md)
- [pg_recall_guard.control](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.9/pg_recall_guard.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.9/CHANGELOG.md)
- [pg_recall_guard--0.2.8--0.2.9.sql](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.9/pg_recall_guard--0.2.8--0.2.9.sql)

`pg_recall_guard` 0.2.9 measures approximate-nearest-neighbour recall against exact scans and records approved index baselines. It discovers distance operators from PostgreSQL catalogs and requires `tsm_system_rows`.

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

### Operational Boundaries

`recall_guard.vector_indexes` lists eligible indexes. `recall_guard.approve()` stores a measurement and search settings; `recall_guard.check()` reuses those settings, and `recall_guard.measure()` performs an individual comparison. Equal-distance neighbours count as hits, so baselines approved under the old duplicate-vector behaviour may need re-approval. The sample is drawn from non-NULL vectors; `k` is 1–1000 and `sample_size` is 1–10000. Partial, expression and temporary indexes are refused. Row security limits baseline/measurement access to readable tables. Each sample performs exact work and holds a relation lock, so schedule conservatively. Samples are not real query traffic, and name-based baselines can outlive index replacement. Upstream tests cover PostgreSQL 18.6/19 beta despite the PG13+ metadata claim. This is pure SQL without preload; current Apache-2.0 licensing starts at 0.2.3.

### Verdict Changes in 0.2.9

Version 0.2.9 changes the values returned by `recall_guard.check()`: `degradado` becomes `degraded`, `critico` becomes `critical`, and `NO SE PUDO MEDIR: <error>` becomes `COULD NOT MEASURE: <error>`; `ok` is unchanged. Update monitoring filters that match the old words. The upgrade restates functions and comments in English, without adding or removing public objects.
