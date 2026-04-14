## 用法
> 来源: [README](https://github.com/yuiseki/pg_rrf/blob/main/README.md) 和 [项目仓库](https://github.com/yuiseki/pg_rrf)。

`pg_rrf` 提供互惠排名融合（Reciprocal Rank Fusion，RRF）函数，用于混合检索场景下的分数融合。
它的重点是合并多个有序候选列表，而不需要手写 `FULL OUTER JOIN` / `COALESCE` 之类的连接逻辑。

### 核心函数

- `rrf(rank_a, rank_b, k)`
- `rrf3(rank_a, rank_b, rank_c, k)`
- `rrf_fuse(ids_a bigint[], ids_b bigint[], k int default 60)`
- `rrfn(ranks bigint[], k int)`

README 还说明了这些分数辅助函数的行为：

- 缺失的排名会被忽略
- `<= 0` 的排名会被忽略
- `k <= 0` 时会报错

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

上游 README 将 `rrf_fuse` 作为手工融合查询的替代方案：

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

### 需求

- PostgreSQL 14-17
- Docker 和 Docker Compose v2

README 说明构建与测试流程都在 Docker 中运行，因此该包工作流不依赖本地 PostgreSQL 或 Rust 工具链。
