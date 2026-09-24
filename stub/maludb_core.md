## Usage

Sources:

- [Source snapshot README](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/README.md)
- [Control file](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/maludb_core.control)
- [Version 0.106.0 SQL](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/sql/extension/maludb_core--0.106.0.sql)
- [Principals and scopes](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/docs/principal-scoping.md)
- [Changelog](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/CHANGELOG.md)
- [Ingest and retrieval example](https://github.com/maludb/maludb-core/blob/71cf34de6986281b36f8bec5b2c924907da97cb7/examples/01-ingest-to-replay.sql)

`maludb_core` stores institutional memory as documents, subjects, claims, facts, episodes, and graph relationships, with text and vector retrieval. This page describes the pinned source snapshot declaring version 0.106.0 on PostgreSQL 17. The changelog marks it unreleased; the latest tagged distribution is v4.5.0, containing SQL extension version 0.100.0.

### Enable a Tenant Schema

Install the extension files and enable it as an administrator in the target database. Dependencies are `vector`, `btree_gist`, `pg_trgm`, `pgcrypto`, and PL/pgSQL. The extension uses a C shared library but does not require shared preloading. Its core schema is fixed; application schemas must be enabled individually.

```sql
CREATE EXTENSION maludb_core CASCADE;
CREATE ROLE app LOGIN;
GRANT maludb_user TO app;
CREATE SCHEMA app AUTHORIZATION app;
ALTER ROLE app SET search_path TO app, maludb_core, public;
SELECT * FROM maludb_core.enable_memory_schema('app');

SET ROLE app;
SET search_path TO app, maludb_core, public;
SELECT * FROM maludb_subject;
RESET ROLE;
```

Configure authentication separately before remote login. `SET ROLE` does not apply a role’s login defaults, which is why the example sets the path explicitly. Normal application connections use the tenant-local views and functions. `maludb_user` enables application access; use `maludb_read` for read-only access.

### Ingest, Retrieve, and Maintain Memory

Tenant APIs include `maludb_memory_ingest_edge` and `maludb_memory_ingest_extraction` for graph ingestion, `maludb_memory_search` and `maludb_vector_search` for retrieval, and `maludb_graph_import` for bounded node-link imports. The extension stores model outputs; extraction and embedding workers are separate clients. The official ingest-to-replay example demonstrates source registration, claims, verified facts, text retrieval, episode replay, and correction by supersession.

`maludb_forget_document` and `maludb_forget_chunk` remove selected material. In 0.106.0 all three retrieval paths filter tombstoned chunks. A source under legal hold refuses deletion, and a still-referenced source is retained and reported. Review the returned result instead of assuming that every related object was deleted.

### Bind a Principal to Each Request

Version 0.106.0 introduces tenant-local principals and scope grants. From an authorized tenant session with no bound principal, create the principal and grants; a trusted service can then narrow each request:

```sql
SELECT maludb_principal_upsert('agent:44', 'agent', 'Sasha', 'agent:44');
SELECT maludb_principal_grant_scope('agent:44', 'dept:3', 'read');
BEGIN;
SET LOCAL maludb_core.principal_ref = 'agent:44';
SET LOCAL maludb_core.principal_scopes = '["agent:44", "dept:3"]';
SET LOCAL maludb_core.principal_readonly = 'on';
SELECT maludb_principal_whoami();
COMMIT;
```

An unset principal preserves unrestricted tenant behavior; an unknown or disabled principal receives no access. A principal has write access to its home scope, write grants imply read, and request scopes can narrow grants but cannot widen them. Sensitivity limits further restrict retrieval. Shared tenant vocabulary can still reveal that a subject name exists.

These settings are a contract with a trusted service, not a boundary against a client holding the tenant’s database login: that client can set them itself. Use database roles and authentication for that boundary. Principal-bound sessions cannot administer principals or their grants.

### Upgrade and Backup

After installing matching library and SQL files on the host, update every database and regenerate interfaces in every enabled tenant schema:

```sql
ALTER EXTENSION maludb_core UPDATE TO '0.106.0';
SELECT maludb_core.maludb_core_version();
SELECT * FROM maludb_core.enable_memory_schema('app');
SELECT schema_name, enabled_version FROM maludb_core.malu$enabled_schema;
```

Tenant-owned interface objects are not replaced by extension migration alone. The 0.105.2 index migrations can block writes while indexes are built. Version 0.105.0 registers stored data for logical dumps, but the database master key and authentication pepper are deliberately excluded; preserve secrets separately and follow the changelog’s restore procedure. Do not assume an older dump contains all memory data.
