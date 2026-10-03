## Usage

Sources:

- [README.md](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/README.md)
- [tdigest.control](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/tdigest.control)
- [Changes](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/Changes)
- [tdigest--1.4.6--1.4.7.sql](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/tdigest--1.4.6--1.4.7.sql)

`tdigest` 1.4.7 supplies approximate, mergeable rank statistics. Store partial digests and combine them to calculate quantiles without sorting the complete original data.

### Core Workflow

```sql
CREATE EXTENSION tdigest;
SELECT tdigest_percentile(v, 100, ARRAY[0.5, 0.95, 0.99])
FROM generate_series(1, 1000) AS g(v);
CREATE TABLE digest_daily AS
SELECT current_date AS day, tdigest(v, 100) AS digest
FROM generate_series(1, 1000) AS g(v);
SELECT tdigest_percentile(digest, 0.95) FROM digest_daily;
```

### Functions and Accuracy

`tdigest(value, compression)` builds a digest; `tdigest(digest)` merges stored digests. `tdigest_percentile` estimates quantiles and `tdigest_percentile_of` estimates ranks, with scalar and array forms. `tdigest_add` and `tdigest_union` update or combine digests; `tdigest_count`, `tdigest_sum` and `tdigest_avg` report counts and trimmed aggregates. The low and high arguments to trimmed aggregates are quantile thresholds, not raw value bounds. `tdigest_is_valid` checks serialized values.

`compression` must be between 10 and 10000. Higher values trade memory and CPU for generally better accuracy, but do not provide a fixed error bound. Validate estimates against exact results on representative data and use consistent compression when merging states.

### Upgrade and Boundaries

Version 1.4.7 fixes closely spaced percentile interpolation, large-count inverse ranks, reusable-state finalization and NULL-compression handling, and marks more scalar functions parallel safe. Upgrade the installed library and SQL together with `ALTER EXTENSION tdigest UPDATE TO '1.4.7'`. Review stored digests with `tdigest_is_valid`; older malformed values can be rejected by the stricter validation. Installation requires superuser rights; the extension is relocatable and requires no preload. The release metadata sets PostgreSQL 13 as the minimum, and the changelog records PostgreSQL 19 build fixes without making a broad performance guarantee.
