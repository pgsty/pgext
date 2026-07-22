---
title: "pg_extension_base"
linkTitle: "pg_extension_base"
description: "Extension development kit by Snowflake"
weight: 2561
categories: ["OLAP"]
width: full
---

[**pg_lake**](https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_base) : Extension development kit by Snowflake


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2561** | {{< badge content="pg_extension_base" link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_base" >}} | {{< ext "pg_extension_base" "pg_lake" >}} | `3.4` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `extension_base` `pg_catalog` |
|    **Need By**    | {{< ext "pg_extension_updater" >}} {{< ext "pg_lake_engine" >}} |
|    **Siblings**   | {{< ext "pg_lake" >}} {{< ext "pg_extension_updater" >}} {{< ext "pg_map" >}} {{< ext "pg_lake_engine" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_lake_table" >}} {{< ext "pg_lake_copy" >}} |

> [!Note] Set shared_preload_libraries=pg_extension_base. It auto-loads dependent libraries declared through the pg_lake control-file marker and manages extension lifecycle background workers.
Extension SQL/control version is 3.4; source and DEB/RPM package version is 3.4.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lake` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.0` | {{< bg "18" "pg_lake_18" "green" >}} {{< bg "17" "pg_lake_17" "green" >}} {{< bg "16" "pg_lake_16" "green" >}} {{< bg "15" "pg_lake_15" "red" >}} {{< bg "14" "pg_lake_14" "red" >}} | `pg_lake_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.0` | {{< bg "18" "postgresql-18-pg-lake" "green" >}} {{< bg "17" "postgresql-17-pg-lake" "green" >}} {{< bg "16" "postgresql-16-pg-lake" "green" >}} {{< bg "15" "postgresql-15-pg-lake" "red" >}} {{< bg "14" "postgresql-14-pg-lake" "red" >}} | `postgresql-$v-pg-lake` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_lake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_lake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "pg_lake_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_lake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_lake_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-18-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-17-pg-lake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.0" "postgresql-16-pg-lake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-lake : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-lake : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_base" title="Repository" icon="github" subtitle="github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_base" >}}
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
pig install pg_extension_base;		# install by extension name, for the current active PG version

pig install pg_extension_base -v 18;   # install for PG 18
pig install pg_extension_base -v 17;   # install for PG 17
pig install pg_extension_base -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_extension_base';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_extension_base;
```

## Usage

Sources:

- [Official pg_extension_base README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/README.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/pg_extension_base.control)
- [SQL API definition](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_extension_base/pg_extension_base--1.6.sql)
- [pg_lake build and preload guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/building-from-source.md)

`pg_extension_base` is Snowflake's infrastructure extension for other PostgreSQL extensions. It provides control-file-driven library preloading, database-scoped lifecycle workers, dependency-aware updates, and short-lived attached workers. Application users normally install it as a dependency of `pg_lake`; extension developers use its C and SQL APIs directly.

### Enable the Infrastructure

`pg_extension_base` must be preloaded before extensions can use its startup machinery:

```conf
shared_preload_libraries = 'pg_extension_base'
```

Restart PostgreSQL, then create it in a database that hosts dependent extensions:

```sql
CREATE EXTENSION pg_extension_base;

SELECT * FROM extension_base.list_preload_libraries();
SELECT * FROM extension_base.list_base_workers();
```

### Extension-Developer Workflow

A dependent extension can request startup loading with a control-file directive:

```text
requires = 'pg_extension_base'
module_pathname = '$libdir/my_extension'
#!shared_preload_libraries
```

To attach a database lifecycle worker, define an internal C function and register it during extension installation:

```sql
CREATE FUNCTION my_extension.main_worker(internal)
RETURNS internal
LANGUAGE C
AS 'MODULE_PATHNAME', 'my_extension_main_worker';

SELECT extension_base.register_worker(
    'my_extension_main',
    'my_extension.main_worker'
);
```

The base infrastructure starts the worker after server startup, `CREATE EXTENSION`, or creation of a database from a template, and attempts to stop it for `DROP EXTENSION` or `DROP DATABASE`.

### SQL API Index

- `extension_base.list_preload_libraries()`: reports extension/library pairs discovered for startup loading.
- `extension_base.register_worker(name, regproc)` and `deregister_worker(name)`: manage lifecycle-worker registrations; public execution is revoked.
- `extension_base.list_base_workers()`: reports database, extension, PID, and restart state for base workers.
- `extension_base.list_database_starters()`: reports per-database starter processes.
- `extension_base.run_attached(command, dbname)`: runs a command in a short-lived worker, optionally in another database, and returns command IDs and tags.

### Operational Boundaries

- Lifecycle worker functions run as superuser and outside a transaction. They must start their own transactions, check interrupts, and avoid trusting user-controlled SQL.
- Worker termination around failed `DROP` operations is best effort; worker code must tolerate the extension briefly disappearing or a stop being reversed.
- `run_attached` is for bounded work that may commit independently of the caller. It is not a detached job queue for long-running tasks.
- Version `3.4` changes no user-facing SQL objects relative to `3.3`; its upgrade script is intentionally empty.
