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

- [Official upstream documentation](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)

`pg_lake_engine` — Query engine for data lake queries

The reviewed catalog snapshot records version `3.4`, kind `preload`, and implementation language `C`.
Install and validate the declared extension dependencies first: `pg_extension_base`, `pg_map`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_lake_engine";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lake_engine';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
