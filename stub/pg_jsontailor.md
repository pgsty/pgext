## Usage

Sources:

- [Upstream source 1](https://github.com/baremonconsulting/pg_jsontailor/blob/2e884e1cac72cbd89aab0fa0293db0d77885768c/README.md)
- [Upstream source 2](https://github.com/baremonconsulting/pg_jsontailor/blob/2e884e1cac72cbd89aab0fa0293db0d77885768c/pg_jsontailor.control)

`pg_jsontailor` development version 0.2.0 generates writable JSON views over relational tables with optimistic concurrency checks.

### Core Workflow

`jsontailor.duality_view_ddl` returns the DDL for review before execution. The supported shape is one root table, an optional read-only many-to-one object and an optional writable one-to-many array. Updates replace complete documents and must present the preceding ETag; guarded deletion also checks that tag.

```sql
CREATE EXTENSION pg_jsontailor;
CREATE SCHEMA shop;
CREATE SCHEMA api;
CREATE TABLE shop.orders (order_id bigint PRIMARY KEY, status text);
SELECT jsontailor.create_duality_view(
  'api.orders_dv',
  '{"root":{"table":"shop.orders","key":"order_id",
    "fields":{"status":"status"},"insert":true,"update":true,"delete":true}}'::jsonb
);
SELECT data FROM api.orders_dv;
```

### Operational Boundaries

This is an unreleased evaluation/controlled-pilot version for PostgreSQL 17–19, implemented in SQL and PL/pgSQL without preload. Generated access follows invoker privileges and row-level security; creating a view does not authorize broader table access. Generated DDL has no runtime extension dependency. It is not a complete Oracle duality-view or ORDS compatibility layer.
