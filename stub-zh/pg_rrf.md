## 用法

> 来源：[README](https://github.com/yuiseki/pg_rrf/blob/main/README.md), [v0.0.3 release](https://github.com/yuiseki/pg_rrf/releases/tag/v0.0.3)

`pg_rrf` 提供用于混合检索分数融合的 Reciprocal Rank Fusion 函数。它的目标是把多个有序候选列表合并起来，而不需要手写 `FULL OUTER JOIN` / `COALESCE` 这类拼接逻辑。

### 核心函数

- `rrf(rank_a, rank_b, k)`
- `rrf3(rank_a, rank_b, rank_c, k)`
- `rrf_fuse(ids_a bigint[], ids_b bigint[], k int default 60)`
- `rrfn(ranks bigint[], k int)`

`v0.0.3` release 明确新增了 `rrfn`，同时保留 `rrf` 和 `rrf3` 作为兼容包装器。README 记录的分数行为也很直接：

- 缺失排名会被忽略
- `<= 0` 的排名会被忽略
- `k <= 0` 会报错

### 示例

```sql
CREATE EXTENSION pg_rrf;

SELECT rrf(1, 2, 60) AS rrf_12;
SELECT rrf3(1, 2, 3, 60) AS rrf_123;
SELECT rrfn(ARRAY[1, 2, 3], 60) AS rrfn_123;
SELECT *
FROM rrf_fuse(ARRAY[10, 20, 30], ARRAY[20, 40], 60)
ORDER BY score DESC;
```

### 混合检索模式

上游 README 把 `rrf_fuse` 作为手工融合查询的替代方案：

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

### 说明

README 目标版本是 PostgreSQL 14-17，并记录了基于 Docker 的构建与测试流程。扩展接口刻意保持很小：核心是分数辅助函数，以及覆盖常见双列表混合检索场景的 `rrf_fuse`。
