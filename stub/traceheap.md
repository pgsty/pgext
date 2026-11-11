## Usage

Sources:

- [Teaching example documentation](https://github.com/styczynski/bioseqdb/blob/3c5ba2a27d4ef288e78d0f645868f811bcbfc64d/website/static/examples/traceheap/README.md)
- [Extension control](https://github.com/styczynski/bioseqdb/blob/3c5ba2a27d4ef288e78d0f645868f811bcbfc64d/website/static/examples/traceheap/traceheap.control)
- [SQL definitions](https://github.com/styczynski/bioseqdb/blob/3c5ba2a27d4ef288e78d0f645868f811bcbfc64d/website/static/examples/traceheap/traceheap--1.0.sql)
- [Executable teaching tests](https://github.com/styczynski/bioseqdb/blob/3c5ba2a27d4ef288e78d0f645868f811bcbfc64d/website/static/examples/traceheap/test.sql)

`traceheap` is an educational PostgreSQL 17 table access method distributed with the BioseqDB tutorial. Version `1.0` retains heap storage while counting selected table access callbacks. It is not BioseqDB's columnar engine and is not intended for production storage. The source explicitly rejects other PostgreSQL majors.

### Observe a scan

Install the library and extension files in a disposable PostgreSQL 17 environment. Creating the access method requires a superuser. No shared preload is required.

```sql
CREATE EXTENSION traceheap;
SET max_parallel_workers_per_gather = 0;

CREATE TABLE specimens (
    id integer NOT NULL,
    label text NOT NULL,
    read_count integer NOT NULL
) USING traceheap;

INSERT INTO specimens VALUES (1, 'soil', 120), (2, 'river', 900);
SELECT traceheap.reset();
SELECT id, label FROM specimens WHERE read_count >= 400 ORDER BY id;
SELECT traceheap.stats();
```

`traceheap.stats()` returns JSON counters for the current backend, and `traceheap.reset()` resets them. They are not transactional, persistent, per-table, or aggregated across parallel workers. A rolled-back write can still increase a counter. Sampling and delegated scan paths are not all counted.

### Supported teaching scope

The upstream tests exercise sequential scans, insertion, `COPY`, update/delete, transaction rollback, `ANALYZE`, ordinary `VACUUM`, heap TOAST and `TRUNCATE`. User indexes, index-backed primary or unique constraints, index validation and clustering rewrites including `VACUUM FULL` are explicitly unsupported.

The extension owns the `traceheap` schema and cannot be relocated. Use only disposable tables for experiments; support for one tested operation does not establish support for every heap operation. Review the upstream test script before extending the example.
