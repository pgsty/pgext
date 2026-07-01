


## 用法

来源：

- [PGXN pgmnemo 0.12.1](https://pgxn.org/dist/pgmnemo/0.12.1/)
- [pgmnemo README](https://github.com/pgmnemo/pgmnemo/blob/v0.12.1/README.md)
- [pgmnemo CHANGELOG](https://github.com/pgmnemo/pgmnemo/blob/v0.12.1/CHANGELOG.md)
- [pgmnemo control file](https://github.com/pgmnemo/pgmnemo/blob/v0.12.1/extension/pgmnemo.control)
- [本地包元数据](../db/extension.csv)

`pgmnemo` 在 PostgreSQL 中存储 agent memory，并通过一个多模态计划结合 pgvector HNSW 搜索、BM25 风格文本匹配、图边邻近度、JSONB 元数据过滤、时间过滤和 outcome confidence 进行检索。它是 SQL-only 扩展，依赖 `vector`，安装到 `pgmnemo` schema；v0.12.1 control 文件中标记为 trusted，且不需要 superuser 权限。

v0.12.1 保留 v0.12.0 的 typed write API，并调整 `guard_no_test_project()`：项目 ID 下限检查改为通过 `pgmnemo.test_project_floor` 显式开启；默认值 `0` 表示禁用该检查。

### 安装

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

SELECT pgmnemo.version();
SELECT * FROM pgmnemo.stats();
```

### 摄入 Lessons

```sql
SELECT pgmnemo.ingest(
  p_role        := 'developer',
  p_project_id  := 1,
  p_topic       := 'security',
  p_lesson_text := 'Rotate JWT signing keys after a key-compromise incident.',
  p_importance  := 4,
  p_embedding   := NULL,
  p_commit_sha  := 'abc1234',
  p_metadata    := '{"source":"incident-runbook"}'::jsonb
);
```

`pgmnemo.ingest()` 是基础写入路径。它会应用 provenance gate，在提供 embedding 时校验 1024 维度，按 `pgmnemo.max_query_text_chars` 截断过长 lesson 文本，并在存在来源信息时写入 `verified_at`。

### Provenance Gate

```sql
SHOW pgmnemo.gate_strict;

SET pgmnemo.gate_strict = 'enforce';
SET pgmnemo.gate_strict = 'warn';
SET pgmnemo.gate_strict = 'off';
```

当 `pgmnemo.gate_strict = 'enforce'` 时，缺少 `commit_sha` 和 `artifact_hash` 的写入会被拒绝。`warn` 会接受写入但产生审计警告，`off` 则关闭 gate。

v0.12.1 的 test-project guard 需要显式开启：

```sql
SHOW pgmnemo.test_project_floor;  -- default 0
SET pgmnemo.test_project_floor = '1000000';
```

默认 `0` 表示禁用生产项目下限检查。测试框架如果保留低 project ID，可设置为正数。

### 召回

```sql
SELECT lesson_id, topic, lesson_text, score
FROM pgmnemo.recall_fast(
  '<1024-dimensional vector literal>'::vector(1024),
  10,
  'developer',
  1,
  'dag-2026-abc',
  ARRAY['note', 'fact']
);

SELECT lesson_id, topic, score, vec_score, bm25_score, rrf_score, match_confidence
FROM pgmnemo.recall_hybrid(
  '<1024-dimensional vector literal>'::vector(1024),
  'JWT rotation key compromise',
  10,
  'developer',
  1,
  0.4,
  0.4,
  60,
  'dag-2026-abc',
  ARRAY['note', 'fact']
);
```

当同时提供文本和 embedding，且 `pgmnemo.disable_hybrid` 为 false 时，`recall_lessons()` 会路由到 hybrid recall。`exclude_dag_id` 用于避免召回同一 workflow run 产生的记录，`p_content_types` 则过滤 typed memories。

召回函数默认会更新 `last_recalled_at` 并递增 `recall_count`。只读分析时可以关闭该副作用：

```sql
SET pgmnemo.track_recall_recency = 'off';
```

### Locate 与 Expand

```sql
SELECT *
FROM pgmnemo.navigate_locate(
  NULL::vector(1024),
  'JWT rotation',
  2000,
  '{"topic":"security"}'::jsonb,
  1
);

SELECT *
FROM pgmnemo.navigate_locate_dispatch(
  query_text            := 'JWT rotation',
  content_type_dispatch := 'entity',
  project_id_filter     := 1,
  token_budget_chars    := 2000
);

SELECT *
FROM pgmnemo.navigate_expand_typed(
  ARRAY[1001, 1002]::bigint[],
  graph_expand_depth := 1,
  relation_types     := ARRAY['CAUSED_BY', 'DERIVED_FROM']
);
```

用 `navigate_locate()` 在字符预算内找候选 ID，再用 expand 函数只获取被选中 ID 的完整内容和图邻居。

### Typed Writes

```sql
SELECT pgmnemo.canonical_slug('concept', 'JWT Rotation');

SELECT *
FROM pgmnemo.remember_fact(
  p_role            := 'developer',
  p_entity_key      := 'concept:jwt_rotation',
  p_property        := 'runbook',
  p_value           := 'Rotate signing secrets within 24 hours after compromise.',
  p_confidence      := 0.82,
  p_has_contact_pii := false,
  p_source_type     := 'agent_authored',
  p_project_id      := 1,
  p_commit_sha      := 'abc1234'
);

SELECT pgmnemo.remember_event(
  p_role        := 'developer',
  p_entity_key  := 'concept:jwt_rotation',
  p_event_label := 'incident_response',
  p_event_body  := 'Rotation procedure validated during the July drill.',
  p_project_id  := 1,
  p_commit_sha  := 'abc1234'
);

SELECT pgmnemo.remember_relation(
  p_role          := 'developer',
  p_from_key      := 'concept:jwt_rotation',
  p_to_key        := 'concept:key_compromise',
  p_relation_type := 'MITIGATES',
  p_project_id    := 1,
  p_commit_sha    := 'abc1234'
);
```

`remember_fact()` 会替换同一 entity/property 的旧 active fact，`remember_event()` 是 append-only，`remember_relation()` 写入关系记忆和图边。

### 图边、强化与维护

```sql
SELECT pgmnemo.add_edge(1001, 1002, 'CAUSED_BY', 0.85, '{"run_id":7320}'::jsonb);

SELECT pgmnemo.reinforce(1001, 'success');
SELECT pgmnemo.reinforce(ARRAY[1001, 1002]::bigint[], 'failure');

SELECT pgmnemo.reembed(1001, '<1024-dimensional vector literal>'::vector(1024));
SELECT pgmnemo.recompute_content(1001, 'Updated lesson text.');
```

常用设置包括 `pgmnemo.gate_strict`、`pgmnemo.include_unverified`、`pgmnemo.ef_search`、`pgmnemo.disable_hybrid`、`pgmnemo.recency_weight`、`pgmnemo.importance_weight`、`pgmnemo.graph_proximity_weight`、`pgmnemo.temporal_boost`、`pgmnemo.confidence_boost_weight`、`pgmnemo.track_recall_recency`、`pgmnemo.max_query_text_chars`、`pgmnemo.tenant_id` 和 `pgmnemo.test_project_floor`。

### 注意事项

- PGXN 元数据声明 `pgmnemo` 依赖 `vector >= 0.7.0`。
- 当前 SQL 定义中的 embedding 预期为 1024 维。
- 默认 provenance gate 是有意设计。迁移旧 memory 行时优先使用 `warn`，不要直接长期使用 `off`。
- 召回函数可能写入 recency 元数据；只读诊断时应关闭 `pgmnemo.track_recall_recency`。
