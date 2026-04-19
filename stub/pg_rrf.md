## Usage
> Sources: [README](https://github.com/yuiseki/pg_rrf/blob/main/README.md), [v0.0.3 release](https://github.com/yuiseki/pg_rrf/releases/tag/v0.0.3)

`pg_rrf` provides Reciprocal Rank Fusion functions for hybrid search score fusion.
It is focused on combining ranked candidate lists without hand-written `FULL OUTER JOIN` / `COALESCE` plumbing.

### Core Functions

- `rrf(rank_a, rank_b, k)`
- `rrf3(rank_a, rank_b, rank_c, k)`
- `rrf_fuse(ids_a bigint[], ids_b bigint[], k int default 60)`
- `rrfn(ranks bigint[], k int)`

The `v0.0.3` release explicitly adds `rrfn` while keeping `rrf` and `rrf3` as compatibility wrappers. The README documents the same score behavior:

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

### Notes

The README targets PostgreSQL 14-17 and documents Docker-based build and test flows. The extension surface remains intentionally small: score helpers plus `rrf_fuse` for the common two-list hybrid-search pattern.
