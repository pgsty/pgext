## Usage
> Sources: [README](https://github.com/yuiseki/pg_rrf/blob/main/README.md) and [project repo](https://github.com/yuiseki/pg_rrf).

`pg_rrf` provides Reciprocal Rank Fusion functions for hybrid search score fusion.
It is focused on combining ranked candidate lists without hand-written `FULL OUTER JOIN` / `COALESCE` plumbing.

### Core Functions

- `rrf(rank_a, rank_b, k)`
- `rrf3(rank_a, rank_b, rank_c, k)`
- `rrf_fuse(ids_a bigint[], ids_b bigint[], k int default 60)`
- `rrfn(ranks bigint[], k int)`

The README also documents the behavior of the score helpers:

- missing ranks are ignored
- ranks `<= 0` are ignored
- `k <= 0` raises an error

### Example

```sql
CREATE EXTENSION pg_rrf;

SELECT rrf(1, 2, 60) AS rrf_12;
SELECT rrf3(1, 2, 3, 60) AS rrf_123;
SELECT rrfn(ARRAY[1, 2, 3], 60) AS rrfn_123;
SELECT *
FROM rrf_fuse(ARRAY[10, 20, 30], ARRAY[20, 40], 60)
ORDER BY score DESC;
```

### Hybrid Search Pattern

The upstream README shows `rrf_fuse` as a replacement for a manual fusion query:

```sql
WITH fused AS (
  SELECT *
  FROM rrf_fuse(
    ARRAY(SELECT id FROM docs ORDER BY bm25_score DESC LIMIT 100),
    ARRAY(SELECT id FROM docs ORDER BY embedding <=> :qvec LIMIT 100),
    60
  )
)
SELECT d.*, fused.score
FROM fused
JOIN docs d USING (id)
ORDER BY fused.score DESC
LIMIT 20;
```

### Requirements

- PostgreSQL 14-17
- Docker and Docker Compose v2

The README says the build and test flow runs in Docker, so local PostgreSQL and Rust toolchains are not required for the package workflow.
