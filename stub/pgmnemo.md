## Usage

Sources:

- [pgmnemo v0.20.0 README](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/README.md)
- [pgmnemo v0.20.0 release notes](https://github.com/pgmnemo/pgmnemo/releases/tag/v0.20.0)
- [pgmnemo v0.20.0 usage guide](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/docs/USAGE.md)
- [pgmnemo v0.20.0 SQL reference](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/docs/SQL_REFERENCE.md)
- [pgmnemo v0.20.0 changelog](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/CHANGELOG.md)
- [pgmnemo v0.20.0 control file](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/extension/pgmnemo.control)
- [v0.19.1 to v0.20.0 upgrade SQL](https://github.com/pgmnemo/pgmnemo/blob/v0.20.0/extension/pgmnemo--0.19.1--0.20.0.sql)

pgmnemo stores agent memory in PostgreSQL and retrieves it through vector, BM25-style text, graph, metadata, temporal, provenance, and outcome-confidence signals. It installs into schema pgmnemo, requires the vector extension, and expects 1024-dimensional embeddings in its current SQL API.

Version 0.20.0 retains the corpus-maintenance, situation, and entity recall surfaces and adds two opt-in candidate-pool expanders to `recall_hybrid()`: causal-edge breadth-first expansion and entity-key GIN expansion. Both are disabled by default, and the result now identifies each candidate's retrieval path.

### Install

    CREATE EXTENSION IF NOT EXISTS vector;
    CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

    SELECT pgmnemo.version();
    SELECT * FROM pgmnemo.stats();

The v0.20.0 control file marks pgmnemo as trusted, installs it in schema `pgmnemo`, requires `vector`, and is not relocatable.

### Ingest a Lesson

    SELECT pgmnemo.ingest(
      p_role        := 'developer',
      p_project_id  := 1,
      p_topic       := 'security',
      p_lesson_text := 'Rotate signing keys after a compromise.',
      p_importance  := 4,
      p_embedding   := NULL,
      p_commit_sha  := 'abc1234',
      p_metadata    := '{"source":"incident-runbook"}'::jsonb
    );

When pgmnemo.gate_strict is enforce, commit_sha or artifact_hash provenance is required. warn accepts an unverified write with an audit warning; off disables the gate.

### Recall with Confidence Filtering

Hybrid recall combines embedding and text signals:

    SELECT lesson_id, topic, score, match_confidence, retrieval_source
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
      ARRAY['note', 'fact'],
      0.40
    );

The final p_min_score argument, added in 0.13.0, removes candidates whose match_confidence is below the threshold before LIMIT is applied. NULL preserves pre-0.13 behavior. The release notes suggest 0.40 as a starting point, not a universal value; calibrate it for the embedding model and feedback quality.

The same p_min_score concept is available in recall_fast, recall_lessons, and pooled recall entry points. recall_lessons routes to hybrid recall when both text and embedding are supplied and pgmnemo.disable_hybrid is off.

### Record Outcomes

    SELECT pgmnemo.reinforce(1001, 'success', true);
    SELECT pgmnemo.reinforce(
      ARRAY[1001, 1002]::bigint[],
      'failure',
      false
    );

The third p_used argument records whether the recalled memory was actually used. true or NULL increments use_count; false records the outcome without counting a use. Prefer an explicit value so analytics can distinguish ignored advice from used advice.

Under the default posterior mode, match confidence is:

    (success_count + alpha)
    / (success_count + failure_count + alpha + beta)

The default Beta prior is alpha 1 and beta 1. Set pgmnemo.confidence_prior_alpha and pgmnemo.confidence_prior_beta between 0.01 and 100 when a different prior is justified.

### Typed Memory and Navigation

Important write helpers include remember_fact, remember_event, remember_relation, add_edge, reembed, and recompute_content. remember_fact supersedes the active fact for an entity/property pair; events remain append-oriented; relations also populate the graph surface.

Use navigate_locate or navigate_locate_dispatch to select candidate IDs within a character budget, then navigate_expand_typed to fetch content and neighboring graph edges.

### Situation and Entity Recall

The 0.15 line adds deterministic situation fingerprints and a dedicated recall path:

```sql
SELECT pgmnemo.extract_sit_fp(
  'security',
  'failure_class=KEY_ROTATION outcome=COMPLETED'
);
SELECT *
FROM pgmnemo.recall_situation(
  pgmnemo.extract_sit_fp(
    'security',
    'failure_class=KEY_ROTATION outcome=COMPLETED'
  ),
  1,
  'developer',
  10
);
```

Starting with 0.15.1, `recall_situation` returns verified memories by default. Set `pgmnemo.include_unverified = on` only when the caller deliberately accepts memories without provenance verification.

The 0.16 line also extracts stable entity keys into `metadata.entity_keys` during ingestion and exposes entity-centered recall:

```sql
SELECT pgmnemo.extract_entity_keys('The run failed with INFRA_FAILURE.');
SELECT * FROM pgmnemo.recall_entity('failure:INFRA_FAILURE', 10);
```

These extractors are deterministic classifiers, not semantic entity resolution. Normalize application vocabulary and inspect the generated keys before relying on them for tenancy or authorization decisions.

### Opt-in Graph and Entity Pool Expansion

Version 0.20.0 can add candidates that lie outside the original ANN and BM25 pool. The graph variant starts from top ANN anchors and follows only `causal` edges, with a depth cap and a per-node hub cap. The entity variant uses the GIN-indexed keys in `metadata.entity_keys`. Both master weights default to `0.0`, so upgrading does not enable either expansion path.

Use transaction-local settings while evaluating the new paths:

```sql
BEGIN;
SET LOCAL pgmnemo.graph_expand_weight = '0.15';
SET LOCAL pgmnemo.graph_expand_depth = '1';
SET LOCAL pgmnemo.graph_expand_ann_k = '15';
SET LOCAL pgmnemo.graph_expand_per_node = '10';
SET LOCAL pgmnemo.graph_entity_expand_weight = '0.10';
SET LOCAL pgmnemo.graph_entity_min_overlap = '1';
SET LOCAL pgmnemo.graph_entity_max_expansion = '50';

SELECT lesson_id, score, retrieval_source
FROM pgmnemo.recall_hybrid(
  query_embedding := '<1024-dimensional vector literal>'::vector(1024),
  query_text := 'JWT rotation key compromise',
  k := 10
);
ROLLBACK;
```

`retrieval_source` is the 18th output column and reports `ann`, `graph`, or `entity`. Graph depth is limited to 1 or 2, ANN anchor oversampling to 10-50, and the per-node hub cap to 3-50. Entity expansion requires at least one overlapping key and applies its configured per-key candidate limit. Treat the weights as workload-specific ranking controls: the v0.20.0 release deliberately makes no general recall or latency guarantee for these opt-in paths.

### Configuration Index

- pgmnemo.confidence_mode: posterior by default; additive retains the legacy calculation.
- pgmnemo.confidence_prior_alpha and pgmnemo.confidence_prior_beta: Bayesian prior parameters.
- pgmnemo.confidence_boost_weight: contribution of confidence to ranking; defaults to 0, so confidence does not change rank unless enabled.
- pgmnemo.gate_strict and pgmnemo.include_unverified: provenance enforcement and retrieval.
- pgmnemo.disable_hybrid and pgmnemo.ef_search: recall strategy and HNSW search breadth.
- pgmnemo.track_recall_recency: whether recall updates last_recalled_at and recall_count.
- pgmnemo.max_query_text_chars, pgmnemo.tenant_id, and pgmnemo.test_project_floor: text, tenancy, and optional test-project controls.
- pgmnemo.graph_expand_weight, pgmnemo.graph_expand_depth, pgmnemo.graph_expand_ann_k, and pgmnemo.graph_expand_per_node: causal-edge candidate expansion, depth, ANN anchors, and per-node hub cap.
- pgmnemo.graph_entity_expand_weight, pgmnemo.graph_entity_min_overlap, and pgmnemo.graph_entity_max_expansion: entity-key GIN expansion, minimum overlap, and per-key candidate limit.

The older confidence-delta settings are deprecated and ignored in posterior mode.

### Upgrade to 0.20.0

Upgrade from 0.19.1 with the packaged extension update path:

```sql
ALTER EXTENSION pgmnemo UPDATE TO '0.20.0';

SELECT extversion
FROM pg_extension
WHERE extname = 'pgmnemo';
```

The upgrade script drops and recreates the 11-argument `recall_hybrid()` because its return shape changes from 17 to 18 columns by adding `retrieval_source`. It also drops and recreates `stats()`, whose return shape grows from 19 to 26 columns with the seven graph-expansion settings. Review positional row mappings, wrappers, prepared consumers, and exact column-count checks before upgrading. Callers that select stable columns by name are unaffected.

### Caveats

- Use PostgreSQL 17 or 18 for pgmnemo 0.20.0. The tagged changelog notes that syntax introduced in the 0.10 line makes older PostgreSQL 14-16 compatibility claims inaccurate; current Pigsty packages target 17-18.

Corpus-maintenance operations are read-only by default:

```sql
SELECT * FROM pgmnemo.reclassify_corpus();
SELECT * FROM pgmnemo.consolidate(
  p_threshold := 0.92,
  p_dry_run := true,
  p_role := NULL,
  p_limit := 100
);
SELECT * FROM pgmnemo.undo_consolidate(
  p_canonical_id := 42,
  p_dry_run := true
);
```

Set `p_dry_run := false` only after reviewing the result inside a transaction. In 0.14.2, reclassification touches only null or classifier-owned types and preserves curator-owned types such as event and relation. Consolidation marks noncanonical lessons superseded, writes edges, and accumulates evidence counts; `undo_consolidate` uses those edges to restore a selected cluster.

- Recall can write recency metadata. Disable pgmnemo.track_recall_recency for read-only analysis.
- The confidence model is only as reliable as reinforcement feedback. Avoid treating posterior values as calibrated probabilities without evaluation.
- HNSW, text, graph, and metadata indexes increase write and maintenance cost.
- The default confidence_boost_weight of 0 means p_min_score can filter results while confidence still contributes nothing to ranking.
- Classification is a deterministic keyword and regular-expression heuristic, not semantic review. Always inspect dry-run distributions and proposed duplicate clusters before applying corpus changes.
