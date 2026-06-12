Source: [pgmnemo v0.8.3 README](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/README.md), [Usage Guide](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/docs/USAGE.md), [extension control file](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/extension/pgmnemo.control), [SQL definition](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/extension/pgmnemo--0.8.3.sql).

## Usage

`pgmnemo` stores provenance-gated agent lessons in PostgreSQL and retrieves them through vector, BM25-style text, graph-edge, JSONB metadata, and relational filters. The extension control file requires `vector`, so `pgvector` must be available before creating `pgmnemo`. The local package metadata targets PostgreSQL 14-18.

### Create and Ingest Lessons

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

`pgmnemo.ingest()` is the preferred write path. It validates the 1024-dimensional embedding when supplied, stamps verified rows when provenance is present, and applies the provenance gate.

### Provenance Gate

```sql
SHOW pgmnemo.gate_strict;

SET pgmnemo.gate_strict = 'warn';
SET pgmnemo.gate_strict = 'enforce';
```

`pgmnemo.gate_strict` accepts `enforce`, `warn`, or `off`. In the default enforced mode, inserts fail when both `p_commit_sha` and `p_artifact_hash` are NULL. `pgmnemo.include_unverified` is separate: it controls whether unverified rows are eligible for recall, not whether writes are allowed.

### Recall

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

Hybrid routing in `recall_lessons()` requires `pgmnemo.disable_hybrid` to be off, non-empty `query_text`, and a non-NULL embedding. Use `recall_hybrid()` directly when you want explicit diagnostic scores.

### Navigation and Expansion

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

`navigate_locate()` returns ranked lesson IDs and short previews within a character budget. `navigate_expand()` fetches selected full lessons and optional graph neighbors after the caller chooses which IDs are worth expanding.

### Edges and Outcome Learning

```sql
SELECT pgmnemo.add_edge(1001, 1002, 'CAUSED_BY', 0.85, '{"run_id":7320}'::jsonb);

SELECT pgmnemo.reinforce(1001, 'success');
SELECT pgmnemo.reinforce(1002, 'failure');
```

`pgmnemo.add_edge()` is the idempotent helper for lesson relationships. `reinforce()` adjusts confidence after observed outcomes and feeds later match confidence.

### Maintenance and GUCs

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

Useful settings include `pgmnemo.gate_strict`, `pgmnemo.include_unverified`, `pgmnemo.ef_search`, `pgmnemo.disable_hybrid`, `pgmnemo.recency_weight`, `pgmnemo.importance_weight`, `pgmnemo.graph_proximity_weight`, `pgmnemo.temporal_boost`, and `pgmnemo.max_query_text_chars`.
