## Usage

Sources:

- [Official README](https://github.com/LogicOcean/pgokf/blob/v0.1.13/README.md)
- [Extension control file](https://github.com/LogicOcean/pgokf/blob/v0.1.13/crates/extension/pgokf.control)
- [pgrx manifest](https://github.com/LogicOcean/pgokf/blob/v0.1.13/crates/extension/Cargo.toml)

`pgokf` materializes Open Knowledge Format bundles into a transactional PostgreSQL catalog with full-text search, a link graph, provenance, optional semantic search, and tenant isolation.

### Enablement and Roles

Release 0.1.13 has pgrx features for PostgreSQL 15–19. Install the exact-major build and create the non-relocatable extension as a superuser:

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

Core lexical search works without dependencies. Optional `vector` enables semantic/hybrid functions, `pg_search` enables BM25, and `pg_cron` enables scheduled refresh. Their absence degrades or errors as documented rather than blocking base installation.

Registration reads server-side files, and some administration functions use `SECURITY DEFINER`; restrict writer/admin membership and review path containment. Version 0.1.x is pre-1.0, so read its upgrade scripts and changelog before changing minor versions. Companion ingest/embed/MCP binaries run outside PostgreSQL and are not required by the core extension.

