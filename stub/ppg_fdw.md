## Usage

Sources:

- [Upstream README](https://github.com/scarbrofair/ppg_fdw/blob/7f8d675da440400ad73d6cfbc397cb1e57a4e4ee/README.md)
- [Version 1.0 install SQL](https://github.com/scarbrofair/ppg_fdw/blob/7f8d675da440400ad73d6cfbc397cb1e57a4e4ee/ppg_fdw--1.0.sql)
- [Distributed planner hook](https://github.com/scarbrofair/ppg_fdw/blob/7f8d675da440400ad73d6cfbc397cb1e57a4e4ee/ppg_planner.c)

ppg_fdw is a legacy prototype for parallel queries over horizontally sharded PostgreSQL servers. It combines a fork of postgres_fdw with a global planner hook and was demonstrated with TPC-H fact-table shards and replicated dimension tables in a shared-nothing layout.

### Minimal FDW objects

The extension installs the ppg_fdw wrapper, handler, and validator. A single-server definition uses familiar foreign-data-wrapper objects:

```sql
CREATE EXTENSION ppg_fdw;

CREATE SERVER shard_group
FOREIGN DATA WRAPPER ppg_fdw
OPTIONS (
    host '127.0.0.1',
    dbname 'analytics',
    port '5432'
);

CREATE USER MAPPING FOR CURRENT_USER
SERVER shard_group
OPTIONS (
    user 'remote_user',
    password 'replace-with-managed-secret'
);

CREATE FOREIGN TABLE remote_orders (
    order_id bigint,
    customer_id bigint,
    total numeric(15,2)
)
SERVER shard_group;

SELECT * FROM remote_orders LIMIT 10;
```

Use a secret-management process for user-mapping credentials and require encrypted, authenticated network connections. The upstream multi-host example replaces the server's catalog option array with custom host1 and host2 entries by directly updating pg_catalog; that unsupported mutation is deliberately omitted here.

### Caveats

- The code and installation notes are from the PostgreSQL 9.3 era and copy PostgreSQL planner and FDW internals. There is no evidence of compatibility with current PostgreSQL releases.
- Upstream's distributed configuration requires direct system-catalog updates that bypass DDL validation, dependency tracking, and supported administration interfaces. Do not use that method on a production cluster.
- The global planner hook handles only a subset of plan nodes and query shapes. Unsupported plans can error, return unexpected behavior, or lose any intended parallelism.
- Correct results depend on the application's sharding key, data placement, duplicated dimension tables, and cross-shard transaction assumptions. The extension does not provide a complete distributed database control plane.
- Treat ppg_fdw as historical research code. For a real deployment, evaluate maintained FDWs or distributed PostgreSQL systems with documented failure handling, security, compatibility, and upgrade paths.
