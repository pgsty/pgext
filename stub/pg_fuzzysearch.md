## Usage

Sources:

- [Upstream source 1](https://github.com/d-d-dev/pg_fuzzysearch/blob/3460d9769515fc03d554493a2aaba438611469a7/README.md)
- [Upstream source 2](https://github.com/d-d-dev/pg_fuzzysearch/blob/3460d9769515fc03d554493a2aaba438611469a7/pg/pg_fuzzysearch.control)

`pg_fuzzysearch` 0.2.0 implements fuzzysort-style subsequence matching and prepared fuzzy targets.

### Core Workflow

`fuzzy_prepare` returns a `fuzzy_target` reusable in generated columns. `fuzzy_match` can apply a minimum score; `fuzzy_score` returns NULL for no match and zero for an empty query. Normalization helpers expose the same Unicode normalization used by matching.

```sql
CREATE EXTENSION pg_fuzzysearch;
SELECT fuzzy_match('pg', 'PostgreSQL');
SELECT fuzzy_score('pg', fuzzy_prepare('PostgreSQL'));
```

### Operational Boundaries

PostgreSQL 17 and UTF-8 databases only. Matching is based on subsequences, not edit distance. The extension provides scan functions rather than an index access method; prepared targets add storage and need workload-specific measurement. The project's own license is not explicit in the reviewed source, although its fuzzysort dependency is identified as MIT.
