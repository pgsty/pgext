## Usage

Sources:

- [README.md](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/README.md)
- [pg_tviews.control](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/pg_tviews.control)
- [Cargo.toml](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/Cargo.toml)
- [CHANGELOG.md](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/CHANGELOG.md)
- [scripts/migrate-from-0.1.0.sql](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/scripts/migrate-from-0.1.0.sql)
- [docs/reference/read-contract.md](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/docs/reference/read-contract.md)
- [docs/operations/replication.md](https://github.com/fraiseql/pg_tviews/blob/9cf848dddfca3eb83ec2f702db5bce3919ab27b9/docs/operations/replication.md)

`pg_tviews` 0.1.0-beta.20 maintains derived TVIEW tables transactionally through analyzed queries and base-table triggers. All extension objects now live in the fixed `tviews` schema, and the SQL version matches the release.

### Core Workflow

```sql
CREATE EXTENSION pg_tviews;
CREATE TABLE tb_post (pk_post bigint PRIMARY KEY, title text);
INSERT INTO tb_post VALUES (1, 'Example');
SELECT tviews.pg_tviews_create('post',
  $$ SELECT pk_post, jsonb_build_object('title', title) AS data FROM tb_post $$);
SELECT * FROM tv_post;
SELECT * FROM tviews.registry;
SELECT * FROM tviews.pg_tviews_health_check();
```

### Operational Boundaries

Basic creation/refresh can load lazily without shared preload in beta.20. Preload and restart remain necessary for automatic rebuild workers and their postmaster settings. The control is not superuser-only; callers need schema creation/trigger rights and must own TVIEWs they replace or drop. Cascades run as each TVIEW owner. The query must expose the entity key and JSONB data expected by the TVIEW contract. UNLOGGED TVIEWs remain unreadable on standbys and are emptied by crash/promotion; select LOGGED storage or plan rebuilding.

Use `tviews.registry` and `tviews.contract_version()` for tool discovery. `tviews.pg_tviews_create_or_replace` applies compatible definition/options changes, while `tviews.pg_tviews_reregister_all` re-derives metadata and triggers without replacing rows. A library/catalog mismatch blocks writes until migration is complete. Existing SQL version 0.1.0 (through beta.19) requires the linked transactional migration script, which preserves TVIEW rows but refuses external extension dependencies; ordinary ALTER EXTENSION alone is not that migration. Back up and rehearse first.
