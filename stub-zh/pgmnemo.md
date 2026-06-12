来源：[pgmnemo v0.8.3 README](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/README.md)、[Usage Guide](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/docs/USAGE.md)、[extension control file](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/extension/pgmnemo.control)、[SQL definition](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/extension/pgmnemo--0.8.3.sql)。

## 用法

`pgmnemo` 在 PostgreSQL 中保存带有来源门控的 agent 经验，并通过向量、BM25 风格文本、图边、JSONB 元数据和关系过滤进行检索。扩展控制文件要求 `vector`，因此在创建 `pgmnemo` 前必须先提供 `pgvector`。本地包元数据面向 PostgreSQL 14-18。

### 创建并摄入经验

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

SELECT pgmnemo.ingest(
  p_role        := 'developer',
  p_project_id  := 1,
  p_topic       := 'security',
  p_lesson_text := 'Rotate JWT secrets after any key-compromise incident.',
  p_importance  := 4,
  p_embedding   := NULL,
  p_commit_sha  := 'abc1234',
  p_metadata    := '{"source":"incident-runbook"}'::jsonb
);
```

`pgmnemo.ingest()` 是推荐的写入路径。它会在提供 embedding 时验证其 1024 维度，在存在来源信息时把行标记为已验证，并应用来源门控。

### 来源门控

```sql
SHOW pgmnemo.gate_strict;

SET pgmnemo.gate_strict = 'warn';
SET pgmnemo.gate_strict = 'enforce';
```

`pgmnemo.gate_strict` 接受 `enforce`、`warn` 或 `off`。在默认强制模式下，如果 `p_commit_sha` 和 `p_artifact_hash` 同时为 NULL，插入会失败。`pgmnemo.include_unverified` 是独立设置：它控制未验证行是否可参与 recall，而不是控制是否允许写入。

### 召回

```sql
-- Text-only recall.
SELECT topic, lesson_text, score
FROM pgmnemo.recall_lessons(
  NULL::vector(1024),
  5,
  'developer',
  1,
  'JWT secret rotation'
);

-- Hybrid vector and text recall.
SELECT lesson_id, topic, score, vec_score, bm25_score, rrf_score
FROM pgmnemo.recall_hybrid(
  '<1024-dimensional vector literal>'::vector(1024),
  'JWT rotation key compromise',
  10,
  'developer',
  1
);
```

`recall_lessons()` 中的混合路由要求 `pgmnemo.disable_hybrid` 关闭、`query_text` 非空，且 embedding 非 NULL。如果需要显式诊断分数，可直接使用 `recall_hybrid()`。

### 导航与扩展

```sql
SELECT *
FROM pgmnemo.navigate_locate(
  NULL::vector(1024),
  'JWT rotation',
  10,
  'developer',
  1,
  '{"topic":"security"}'::jsonb,
  2000
);

SELECT *
FROM pgmnemo.navigate_expand(
  ARRAY[1001, 1002]::bigint[],
  include_edges := true
);
```

`navigate_locate()` 会在字符预算内返回排序后的 lesson ID 和短预览。调用方选择值得展开的 ID 后，`navigate_expand()` 会取回选中的完整经验，并可选地带上图邻居。

### 边与结果学习

```sql
SELECT pgmnemo.add_edge(1001, 1002, 'CAUSED_BY', 0.85, '{"run_id":7320}'::jsonb);

SELECT pgmnemo.reinforce(1001, 'success');
SELECT pgmnemo.reinforce(1002, 'failure');
```

`pgmnemo.add_edge()` 是维护经验关系的幂等辅助函数。`reinforce()` 会根据观察到的结果调整置信度，并影响后续匹配置信度。

### 维护与 GUC

```sql
SELECT * FROM pgmnemo.stats();

SELECT pgmnemo.reembed(
  p_lesson_id  := 1001,
  p_new_vector := '<1024-dimensional vector literal>'::vector(1024)
);

SELECT pgmnemo.recompute_content(
  p_lesson_id       := 1001,
  p_new_lesson_text := 'Rotate JWT secrets within 24 hours after compromise.'
);
```

常用设置包括 `pgmnemo.gate_strict`、`pgmnemo.include_unverified`、`pgmnemo.ef_search`、`pgmnemo.disable_hybrid`、`pgmnemo.recency_weight`、`pgmnemo.importance_weight`、`pgmnemo.graph_proximity_weight`、`pgmnemo.temporal_boost` 和 `pgmnemo.max_query_text_chars`。
