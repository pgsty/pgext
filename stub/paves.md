## Usage

Sources:

- [Official README](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/README.md)
- [Control file](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/extension/paves/paves.control)
- [Version 0.1 SQL](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/extension/paves/sql/paves--0.1.sql)
- [Configuration and limits](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/docs/07-reference.md)
- [Index maintenance](https://github.com/SakuraMarble/PAVES/blob/ac0de668b48f62945c38099a36155e3a98aac181/docs/02-index-build.md)

`paves` is a PostgreSQL 14 research extension for filtered vector search. It combines scalar predicates with L2 vector ordering and chooses CPU or optional GPU execution paths. Its indexes are static snapshots: table changes require rebuilding the index.

### Core Workflow

Install the matching PostgreSQL 14 library and the `vector` dependency. Add `paves` to the preload list and restart PostgreSQL, then enable both extensions as an administrator:

```conf
shared_preload_libraries = 'paves'
```

```sql
CREATE EXTENSION vector;
CREATE EXTENSION paves;

CREATE TABLE items (
    id bigint PRIMARY KEY,
    embedding vector(3),
    price integer,
    category integer
);
INSERT INTO items VALUES
    (1, '[0.1,0.2,0.3]', 50, 3),
    (2, '[0.4,0.5,0.6]', 80, 7),
    (3, '[0.8,0.9,1.0]', 150, 1);

SELECT * FROM paves_capacity('items');
SELECT paves_build('items');

EXPLAIN SELECT id FROM items
WHERE price < 100 AND category IN (3, 7)
ORDER BY embedding <-> '[0.1,0.2,0.3]'::vector
LIMIT 10;
```

The first `vector` column is indexed. Vector dimensions must agree; null vectors are skipped and an empty build is rejected. Only a single L2 distance ordering using `<->` with `LIMIT` is eligible. Other query forms use normal PostgreSQL paths; inspect the plan to see which path was selected.

### Interfaces and Tuning

- `paves_build(regclass)` scans the table, atomically replaces its index file and returns the indexed row count.
- `paves_capacity(regclass)` reports dimension, row and scalar-column limits, estimated disk/memory use and GPU availability.
- `paves_feedback()` exposes routing feedback and CPU/GPU load counters.
- `paves.enable` controls custom paths; `paves.force_strategy` selects `auto`, `brute`, `hnsw`, `gpu_brute` or `gpu_hnsw`.
- `paves.ef_search` controls graph-search breadth; `paves.threads` controls CPU scan workers.
- `paves.build_m`, `paves.build_ef_construction` and `paves.build_threads` govern index construction.

CUDA is optional; a CPU build works without it. The GPU worker requires preloading and restart. `paves.gpu_enable` limits automatic routing, while `paves.gpu_batch_max` and `paves.gpu_batch_wait_us` tune worker microbatches. Unsupported sizes or unavailable worker capacity can fall back to CPU execution.

### Maintenance and Limits

Re-run `paves_build()` after writes. Heap visibility rechecks do not maintain the index or guarantee recall of changed data. Index files live under `$PGDATA/paves/`; `DROP TABLE` leaves orphan files for manual cleanup. Rebuilds temporarily need both old and new files, and old mappings can remain until sessions end.

The implementation targets PostgreSQL 14 and one resident GPU per index; no broader major-version or multi-GPU support is claimed. Approximate paths trade recall against work. Validate correctness and resource limits for the actual workload; published benchmark results do not establish general performance guarantees. No license declaration was found in the reviewed source revision.
