## Usage

Sources:

- [I-Rex extension guide](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/README.md)
- [Control file](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/blmfl.control)
- [Version 0.0.1 SQL](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/blmfl--0.0.1.sql)
- [Usage tests](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/sql/blmfl_test.sql)
- [C implementation](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/blmfl.c)
- [PostgreSQL 16 environment](https://github.com/yihaoh/I-Rex/blob/f7b781a3cd31a2b8363b7c6fda59fb335c00091b/hnrq-db/sargsum/docker/postgres/Dockerfile)

`blmfl` 0.0.1 is the experimental bloom-filter and range-summary extension shipped with the I-Rex SQL-debugging research artifact. It supplies aggregates and membership probes; it does not automatically optimize ordinary PostgreSQL queries.

### Enable and Create a Summary

The native extension requires administrative installation. The supplied environment uses PostgreSQL 16; other majors are not established by that artifact. No shared preload is required.

```sql
CREATE EXTENSION blmfl;

SET blmfl.bloomfilter_bitsize = 256;
SET blmfl.estimated_count = 10;
SET blmfl.num_hashes = 15;

CREATE TEMP TABLE example_values (value integer);
INSERT INTO example_values VALUES (5), (24), (90);

CREATE TEMP TABLE example_summary AS
SELECT blmfl_any(value) AS summary FROM example_values;

SELECT blmfl_test_any(summary, 24),
       blmfl_fpr(summary)
FROM example_summary;
```

Set positive sizes and counts explicitly: all three GUCs default to zero, and the source contains divisions and modulo operations using them. These session settings determine subsequently constructed summaries, not existing values.

### SQL Surface

- `BLMFL_RESULT` stores filter bytes, size, estimated count, hash count and accumulated-item count.
- `blmfl_any` and `blmfl_test_any` accept typed arguments; `blmfl` and `blmfl_test` expose the bytea variant. Keep construction and probe types consistent.
- `blmfl_optimal_k`, `blmfl_fpr` and `blmfl_approx_unique_count` expose sizing and estimate helpers.
- `blmfl_merge` and `blmfl_agg` combine summaries. The implementation checks matching bit lengths but does not establish every compatibility invariant.
- `sarg_agg` builds or combines lower/upper summaries; the `is_iid_*`, `min_iid` and `max_iid` helpers support I-Rex's internal identifiers.

### Research Boundaries

Source inspection shows that the typed hash loop retains only its final argument's hash and ignores the per-hash seed. Do not treat it as a validated multi-column bloom filter. The merge buffer is allocated without explicit zeroing, and empty-input finalization raises an error. Estimates and merged summaries therefore need independent correctness checks.

Keep this artifact in an isolated research environment. Do not use its summaries to exclude application rows without exact validation. The supplied revision has no identified license grant, supported upgrade chain or production compatibility guarantee.
