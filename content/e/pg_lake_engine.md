---
title: "pg_lake_engine"
linkTitle: "pg_lake_engine"
description: "Query engine for data lake queries"
weight: 2564
categories: ["OLAP"]
width: full
---

[**pg_lake**](https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_engine) : Query engine for data lake queries


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2564** | {{< badge content="pg_lake_engine" link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_engine" >}} | {{< ext "pg_lake_engine" "pg_lake" >}} | `3.4` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `__lake__internal__nsp__` `lake_engine` `lake_struct` `pg_catalog` |
|   **Requires**    | {{< ext "pg_extension_base" >}} {{< ext "pg_map" >}} |
|    **Need By**    | {{< ext "pg_lake_copy" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_lake_table" >}} |
|    **Siblings**   | {{< ext "pg_lake" >}} {{< ext "pg_extension_base" >}} {{< ext "pg_extension_updater" >}} {{< ext "pg_map" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_lake_table" >}} {{< ext "pg_lake_copy" >}} |

> [!Note] Query-engine component. pg_extension_base auto-loads its module; delegated DuckDB execution additionally requires the separately running PG-major pgduck_server.
Extension SQL/control version is 3.4; source and DEB/RPM package version is 3.4.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lake` | `pg_extension_base`, `pg_map` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.0` | {{< bg "18" "pg_lake_18" "green" >}} {{< bg "17" "pg_lake_17" "green" >}} {{< bg "16" "pg_lake_16" "green" >}} {{< bg "15" "pg_lake_15" "red" >}} {{< bg "14" "pg_lake_14" "red" >}} | `pg_lake_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.0` | {{< bg "18" "postgresql-18-pg-lake" "green" >}} {{< bg "17" "postgresql-17-pg-lake" "green" >}} {{< bg "16" "postgresql-16-pg-lake" "green" >}} {{< bg "15" "postgresql-15-pg-lake" "red" >}} {{< bg "14" "postgresql-14-pg-lake" "red" >}} | `postgresql-$v-pg-lake` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_lake_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_lake_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lake_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lake_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lake : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lake : N/A 0" "gray" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_engine" title="Repository" icon="github" subtitle="github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_engine" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_lake-3.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_lake;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_lake;		# install via package name, for the active PG version
pig install pg_lake_engine;		# install by extension name, for the current active PG version

pig install pg_lake_engine -v 18;   # install for PG 18
pig install pg_lake_engine -v 17;   # install for PG 17
pig install pg_lake_engine -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_extension_base';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_lake_engine CASCADE; -- requires pg_extension_base, pg_map
```

## Usage

Sources:

- [Official pg_lake architecture overview](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md#architecture)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine.control)
- [Base SQL objects](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine--3.0.sql)
- [Version 3.4 cleanup-queue change](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_engine/pg_lake_engine--3.3--3.4.sql)

`pg_lake_engine` is the shared execution layer used by the pg_lake table, copy, and Iceberg extensions. It rewrites eligible PostgreSQL work for `pgduck_server`, maps PostgreSQL and DuckDB values, and tracks remote files that must be removed after aborts or table changes. It is an internal dependency rather than a standalone analytics interface.

### Deployment Boundary

Install it through the top-level extension so its dependency graph and preload directives remain aligned:

```conf
shared_preload_libraries = 'pg_extension_base'
```

```sql
CREATE EXTENSION pg_lake CASCADE;
```

`pg_lake_engine` requires `pg_extension_base` and `pg_map`. Query execution also requires a running local `pgduck_server`; creating only this extension does not provide a usable lake table.

### User-Visible Objects

- Roles `lake_read`, `lake_write`, and `lake_read_write`: shared privilege groups consumed by the other components.
- `to_postgres(any)`: returns its input while forcing that expression to be evaluated in PostgreSQL instead of pushed to the lake engine.
- `to_date(double precision)`: converts a days-since-Unix-epoch value commonly found in Parquet to a PostgreSQL `date`.
- `lake_engine.deletion_queue`: tracks committed orphan-file cleanup; readable by `lake_write`.
- `lake_engine.in_progress_files`: tracks files produced by transactions that have not committed.
- `lake_engine.flush_deletion_queue(regclass)` and `flush_in_progress_queue()`: privileged cleanup functions used by maintenance paths.

```sql
SELECT to_postgres(application_only_function(payload))
FROM external_events;
```

Use `to_postgres()` only when an expression cannot or should not be pushed down; pulling data back into PostgreSQL may substantially increase transfer and execution cost.

### Internal State and Caveats

- The `__lake__internal__nsp__` functions are planner/deparser placeholders and are not a supported direct SQL API.
- Do not manually update or delete queue rows. Cleanup functions need the extension's object-store credentials and privilege roles and should be invoked only as documented by operational tooling.
- Version `3.4` adds `resolve_metadata` to the deletion queue so Iceberg metadata can be expanded into exact referenced files during `VACUUM`, moving object-store traversal off the `DROP` path.
- Roles are cluster-wide objects and can outlive an extension instance in one database; review memberships separately when removing pg_lake.
