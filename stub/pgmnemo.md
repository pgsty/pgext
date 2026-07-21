## Usage

Sources:

- [pgmnemo v0.13.0 README](https://github.com/pgmnemo/pgmnemo/blob/v0.13.0/README.md)
- [pgmnemo v0.13.0 usage guide](https://github.com/pgmnemo/pgmnemo/blob/v0.13.0/docs/USAGE.md)
- [pgmnemo v0.13.0 SQL reference](https://github.com/pgmnemo/pgmnemo/blob/v0.13.0/docs/SQL_REFERENCE.md)
- [pgmnemo v0.13.0 release notes](https://github.com/pgmnemo/pgmnemo/releases/tag/v0.13.0)
- [pgmnemo v0.13.0 control file](https://github.com/pgmnemo/pgmnemo/blob/v0.13.0/extension/pgmnemo.control)

pgmnemo stores agent memory in PostgreSQL and retrieves it through vector, BM25-style text, graph, metadata, temporal, provenance, and outcome-confidence signals. It installs into schema pgmnemo, requires the vector extension, and expects 1024-dimensional embeddings in its current SQL API.

Version 0.13.0 changes confidence to a Bayesian posterior by default, records whether recalled memories were actually used, and adds minimum-confidence filters to recall.

### Install

    CREATE EXTENSION IF NOT EXISTS vector;
    CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

    SELECT pgmnemo.version();
    SELECT * FROM pgmnemo.stats();

The v0.13.0 control file marks pgmnemo as trusted and non-superuser-installable when the required vector extension is available.

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

    SELECT lesson_id, topic, score, match_confidence
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

### Configuration Index

- pgmnemo.confidence_mode: posterior by default; additive retains the legacy calculation.
- pgmnemo.confidence_prior_alpha and pgmnemo.confidence_prior_beta: Bayesian prior parameters.
- pgmnemo.confidence_boost_weight: contribution of confidence to ranking; defaults to 0, so confidence does not change rank unless enabled.
- pgmnemo.gate_strict and pgmnemo.include_unverified: provenance enforcement and retrieval.
- pgmnemo.disable_hybrid and pgmnemo.ef_search: recall strategy and HNSW search breadth.
- pgmnemo.track_recall_recency: whether recall updates last_recalled_at and recall_count.
- pgmnemo.max_query_text_chars, pgmnemo.tenant_id, and pgmnemo.test_project_floor: text, tenancy, and optional test-project controls.

The older confidence-delta settings are deprecated and ignored in posterior mode.

### Caveats

- Recall can write recency metadata. Disable pgmnemo.track_recall_recency for read-only analysis.
- The confidence model is only as reliable as reinforcement feedback. Avoid treating posterior values as calibrated probabilities without evaluation.
- HNSW, text, graph, and metadata indexes increase write and maintenance cost.
- The default confidence_boost_weight of 0 means p_min_score can filter results while confidence still contributes nothing to ranking.
