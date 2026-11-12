## Usage

Sources:

- [Official README](https://github.com/LogicOcean/pgokf/blob/v0.4.0/README.md)
- [Extension control file](https://github.com/LogicOcean/pgokf/blob/v0.4.0/crates/extension/pgokf.control)
- [SQL API](https://github.com/LogicOcean/pgokf/blob/v0.4.0/docs/sql-api.md)
- [Upgrade contract](https://github.com/LogicOcean/pgokf/blob/v0.4.0/docs/api-stability.md)
- [Release and upgrade details](https://github.com/LogicOcean/pgokf/releases/tag/v0.4.0)
- [pgrx manifest](https://github.com/LogicOcean/pgokf/blob/v0.4.0/crates/extension/Cargo.toml)

`pgokf` materializes Open Knowledge Format bundles into a transactional PostgreSQL catalog with full-text search, a link graph, provenance, optional semantic search, and tenant isolation.

### Enablement and Roles

Release 0.4.0 supports PostgreSQL 15–18; PostgreSQL 19 RC testing is a separate compatibility target. Install the exact-major build and create the non-relocatable extension as a superuser:

```sql
CREATE EXTENSION pgokf;
GRANT pgokf_writer TO app_user;
```

No preload is required. The extension creates `pgokf_reader`, `pgokf_writer`, and `pgokf_admin` roles; grant the least-capable role needed.

### Register and Search a Bundle

The filesystem path is resolved by the PostgreSQL server and must be absolute and server-readable.

```sql
SELECT *
FROM pgokf.register_bundle('/srv/okf/runbooks');

SELECT concept_id, title, rank
FROM pgokf.concept_search(
  'postgres failover',
  concept_type => 'runbook'
);

SELECT *
FROM pgokf.concept_neighbors('runbooks/database-failover', 2);
```

The bundle remains the portable source of truth; use `catalog_stats`, `health`, `search_index_status`, and the sync log to operate the materialized projection.

### Optional Search Backends and Boundaries

Core lexical search works without dependencies. Optional `vector` enables semantic/hybrid functions, `pg_textsearch` or `pg_search` enables BM25, and `pg_cron` enables scheduled refresh. Their absence degrades or errors as documented rather than blocking base installation.

Registration reads server-side files, and some administration functions use `SECURITY DEFINER`; restrict writer/admin membership and review path containment. Version 0.4.0 is pre-1.0, so read its upgrade scripts and changelog before changing minor versions. Companion ingest/embed/MCP binaries run outside PostgreSQL and are not required by the core extension.


### Generations, Freshness and Upgrades

Catalog generations, a durable change-event outbox and typed relationships let producers reconcile derived knowledge. `concept_search_fresh` exposes freshness filtering; `mark_reconciling` and `mark_fresh` provide guarded producer transitions. Semantic ranking requires matching file, model and embedding-contract provenance plus effective freshness. Legacy bundles become stale during the 0.3 upgrade, and old embeddings must be regenerated before they can rank.

Back up the database and install the 0.4.0 files before updating:

```sql
ALTER EXTENSION pgokf UPDATE TO '0.4.0';
SELECT extversion FROM pg_extension WHERE extname = 'pgokf';
```

The supplied chain covers earlier 0.1, 0.2, 0.3 and development versions. It reconciles superseded search overloads only when their exact definitions, grants and dependencies match the expected old catalog; an unexpected shape aborts the update. Do not force-drop dependent objects to bypass that guard. The declared older web-table column ordering is retained. SQL catalog changes and the separately deployed MCP/web companions must be upgraded according to their own contracts.
