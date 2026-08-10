## Usage

Sources:

- [pg_local_cache 1.2.0 on PGXN](https://pgxn.org/dist/pg_local_cache/1.2.0/)
- [pg_local_cache v1.2.0 README](https://github.com/profundium/pg_local_cache/blob/v1.2.0/README.md)
- [pg_local_cache v1.2.0 control file](https://github.com/profundium/pg_local_cache/blob/v1.2.0/pg_local_cache.control)
- [pg_local_cache 1.2.0 extension SQL](https://github.com/profundium/pg_local_cache/blob/v1.2.0/sql/pg_local_cache--1.2.0.sql)
- [Technical reference](https://github.com/profundium/pg_local_cache/blob/v1.2.0/docs/TECHNICAL.md)
- [Existing-server installation guide](https://github.com/profundium/pg_local_cache/blob/v1.2.0/docs/INSTALL_EXISTING.md)
- [pg_local_cache v1.2.0 release](https://github.com/profundium/pg_local_cache/releases/tag/v1.2.0)

`pg_local_cache` 1.2.0 is a transaction-aware, in-process cache for repeated PostgreSQL primary-key reads. It keeps bounded whole-row entries in shared memory and can transparently accelerate eligible ordinary `SELECT` statements while retaining the original PostgreSQL primary-key plan as the authoritative fallback. Use it for a hot working set on one writable primary; it is not a general query-result cache, a durability layer, or a distributed Redis/Valkey replacement.

### Core Workflow

The library must be loaded at postmaster startup. This SQL-only configuration disables the optional RESP listener and serves one application database:

```conf
shared_preload_libraries = 'pg_local_cache'
pg_local_cache.database = 'app'
pg_local_cache.port = 0
pg_local_cache.cache_entries = 16384
pg_local_cache.memory_budget_mb = 384
```

Add `pg_local_cache` to any existing comma-separated preload list instead of replacing other libraries, validate the configuration, and perform a controlled PostgreSQL restart. The control file fixes the extension in schema `local_cache`, sets `superuser=true`, and is not relocatable, so a superuser must create it in each database where it is used:

```sql
CREATE EXTENSION pg_local_cache;
```

Create an eligible permanent table, then attach it. `attach_table` takes a `ShareRowExclusiveLock`, records the complete primary key in `local_cache.mapping`, installs extension-owned invalidation triggers, and publishes the mapping to shared memory. Use a bounded lock timeout on a live system:

```sql
CREATE TABLE public.items (
    id bigint PRIMARY KEY,
    value text NOT NULL,
    enabled boolean NOT NULL DEFAULT true,
    metadata jsonb
);

INSERT INTO public.items VALUES
    (1, 'hello', true, '{"source":"postgres"}');

BEGIN;
SET LOCAL lock_timeout = '2s';
SELECT local_cache.attach_table('public.items'::regclass);
COMMIT;
```

The default `p_writable=false` disables RESP `SET` and `DEL`; it does not prevent normal PostgreSQL DML. Applications keep using their existing PostgreSQL connection, row types, and SQL:

```sql
SELECT * FROM public.items WHERE id = $1::bigint;

SELECT value, metadata
FROM public.items
WHERE id = ANY($1::bigint[]);

EXPLAIN (ANALYZE, COSTS OFF)
SELECT * FROM public.items WHERE id = 1;
```

An eligible plan appears as `Custom Scan (pg_local_cache_sql)`. A cache miss or any unsafe or unsupported condition executes the retained primary-key index plan; PostgreSQL remains the source of truth.

### Explicit JSON APIs

Ordinary SQL is the canonical tuple-returning interface. Callers that deliberately want a cache-shaped JSON API can use these `SECURITY INVOKER` functions:

```sql
SELECT local_cache.get('public.items'::regclass, 1::bigint);

SELECT local_cache.mget(
    'public.items'::regclass,
    ARRAY[1, 7, 1]::bigint[]
);
```

`get(regclass, anyelement)` returns complete-row JSON as `text`; `mget(regclass, anyarray)` returns a `text[]` aligned with its input, preserving duplicate and `NULL` positions. For a composite or heterogeneous primary key, call `get(regclass, text[])` with components in the key order recorded by `attach_table`. Explicit API callers need `USAGE` on schema `local_cache`, `EXECUTE` on the chosen overload, and normal `SELECT` privilege on the source table.

### Important Objects and Controls

- `local_cache.attach_table(regclass, boolean, text)` validates and registers a table. Set `p_writable=true` only when the optional RESP worker should be allowed to write the source relation; `p_namespace` overrides the generated mapping name.
- `local_cache.detach_table(regclass)` removes the mapping, managed triggers, shared entry state, and direct worker-role privileges. It returns `false` when the relation was not attached.
- `local_cache.reconcile_table(regclass)` and `local_cache.reconcile_all()` revalidate relation shape, primary keys, trigger provenance, and worker grants after controlled DDL or privilege changes.
- `local_cache.mapping` is the extension-owned mapping registry and is included in extension configuration dumps. Do not edit it as a substitute for the administrative functions.
- `local_cache.metrics()` returns typed counters and memory/worker gauges, `local_cache.health()` returns a compact JSON readiness assessment, and `local_cache.stats()` returns detailed JSON diagnostics. These and the administrative functions are revoked from `PUBLIC`; grant them only to designated deploy or monitoring roles.
- `local_cache.invalidate(namespace)` invalidates one mapping namespace and returns the affected-entry count. Normal DML, `TRUNCATE`, and relevant DDL use automatic transactional invalidation.

Key settings are:

| Setting | Default | Effect |
|---|---:|---|
| `pg_local_cache.port` | `6380` | RESP2 port; set `0` for SQL-only mode. |
| `pg_local_cache.database` | `postgres` | The one database served by this extension instance. |
| `pg_local_cache.cache_entries` | `16384` | Shared row-cache entry capacity. |
| `pg_local_cache.relation_states` | `1024` | Shared relation-version state capacity. |
| `pg_local_cache.memory_budget_mb` | `384` | Startup budget for deterministic extension allocations. |
| `pg_local_cache.max_dirty_keys` | `4096` | Per-transaction key bound before invalidation widens to the relation. |
| `pg_local_cache.sql_cache` | `on` | `USERSET` switch for the ordinary-SQL fast path; no restart is required. |

Except for `pg_local_cache.sql_cache`, the documented GUCs are postmaster settings. The memory budget covers the extension's deterministic shared hashes and optional RESP buffers, not `shared_buffers`, backend memory, the operating system, or other services.

### Fast-Path and Consistency Boundaries

The transparent path is deliberately narrow. It requires `READ COMMITTED`, one attached base table, direct column projections, and equality predicates for every primary-key column. A single-column primary key also supports bounded `IN` and `= ANY(array)` queries. Joins, CTEs, subqueries, aggregates, grouping, ordering, row locks, extra predicates, recovery, parallel execution, `REPEATABLE READ`, and `SERIALIZABLE` use normal PostgreSQL plans. Scalar lookups may use no `LIMIT` or constant `LIMIT 1`; batch lookups may not use `LIMIT`.

For `IN`/`ANY`, the executor deduplicates at most 1,024 non-null keys and copies at most 16 MiB of validated rows into query-local memory. The batch is all-or-nothing: one miss, unsafe snapshot, malformed entry, or budget overflow runs the complete source plan rather than merging cached and source rows.

Source-table writes remain ordinary PostgreSQL transactions. Managed triggers collect changed keys, and the pre-commit callback publishes invalidation fences before the transaction becomes visible. A rollback never publishes uncommitted row data. After the current transaction writes an attached relation, subsequent reads in that transaction bypass the cache to preserve read-your-own-write behavior. `PREPARE TRANSACTION` is rejected after such a write.

Entries have no TTL. They remain until invalidation, eviction, replacement, corruption detection, or an MVCC safety check retires them. Encoded cache values are limited to 8 KiB; a wider row simply uses PostgreSQL instead of becoming an entry.

### Table and Deployment Requirements

Version 1.2.0 supports PostgreSQL 14–18 on the published Linux amd64 builds, one configured database, and one writable primary. Attached relations must be permanent, non-partitioned tables with no inheritance or row-level security and with an immediate, non-partial B-tree primary key. Supported key columns are `int2`, `int4`, `int8`, `text`, `varchar`, `bpchar`, and `uuid`; composite keys may contain 1–16 columns. Temporary or unlogged tables, views, partitioned tables, expression or partial primary keys, nondeterministic key collations, and non-default primary-key operator classes are rejected.

At most 128 mappings are published per instance. Dropping a table forgets its mapping; recreating a table with the same name does not reattach it. The cache is not served on standbys and provides no multi-primary coordination, TTL, clustering, Pub/Sub, or general range/join/aggregate caching.

### Optional RESP2 Security Boundary

RESP mode exposes whole-row `GET`, `SET`, and `DEL` through a limited RESP2 protocol, but it uses one shared token and one `LOGIN NOSUPERUSER NOINHERIT` worker role for every accepted mapping. It has no TLS and no per-client PostgreSQL identity or ACL context. Keep `pg_local_cache.port=0` unless this interface is required. If enabled, retain the default loopback bind or place remote access behind network isolation and authenticated TLS, store the token in a PostgreSQL OS-user-owned mode `0400` or `0600` file through `pg_local_cache.auth_token_file`, and never treat a lost write reply as proof that the PostgreSQL transaction did not commit.
