---
title: "pg_lake_copy"
linkTitle: "pg_lake_copy"
description: "Copy to/from data lake files"
weight: 2567
categories: ["OLAP"]
width: full
---

[**pg_lake**](https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_copy) : Copy to/from data lake files


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2567** | {{< badge content="pg_lake_copy" link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_copy" >}} | {{< ext "pg_lake_copy" "pg_lake" >}} | `3.4` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **Requires**    | {{< ext "pg_lake_engine" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_lake_table" >}} |
|    **Need By**    | {{< ext "pg_lake" >}} |
|    **Siblings**   | {{< ext "pg_lake" >}} {{< ext "pg_extension_base" >}} {{< ext "pg_extension_updater" >}} {{< ext "pg_map" >}} {{< ext "pg_lake_engine" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_lake_table" >}} |

> [!Note] pg_extension_base auto-loads pg_lake_engine, pg_lake_iceberg, pg_lake_table, and pg_lake_copy in dependency order.
Extension SQL/control version is 3.4; source and DEB/RPM package version is 3.4.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lake` | `pg_lake_engine`, `pg_lake_iceberg`, `pg_lake_table` |
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
{{< card link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_copy" title="Repository" icon="github" subtitle="github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_copy" >}}
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
pig install pg_lake_copy;		# install by extension name, for the current active PG version

pig install pg_lake_copy -v 18;   # install for PG 18
pig install pg_lake_copy -v 17;   # install for PG 17
pig install pg_lake_copy -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_extension_base';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_lake_copy CASCADE; -- requires pg_lake_engine, pg_lake_iceberg, pg_lake_table
```

## Usage

Sources:

- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_copy/pg_lake_copy.control)
- [Official data-lake import and export guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/data-lake-import-export.md)
- [Official file-format reference](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/file-formats-reference.md)
- [Version 3.4 SQL file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_copy/pg_lake_copy--3.3--3.4.sql)

`pg_lake_copy` extends PostgreSQL `COPY` so queries, heap tables, external lake tables, and Iceberg tables can exchange Parquet, CSV, or newline-delimited JSON files with local paths, HTTP endpoints, and configured object stores. It adds behavior through hooks and has no standalone SQL function API.

### Enable the Component

The normal entry point installs `pg_lake_copy` and its exact dependencies together:

```sql
CREATE EXTENSION pg_lake CASCADE;
```

Its control file requires `pg_lake_engine`, `pg_lake_iceberg`, and `pg_lake_table`. The deployment also needs `pg_extension_base` in `shared_preload_libraries` and a running `pgduck_server`.

### Export and Import

Format is inferred from the path suffix or selected explicitly:

```sql
COPY (
    SELECT event_id, event_time, payload
    FROM events
    WHERE event_time >= DATE '2026-07-01'
)
TO 's3://analytics-bucket/events/july.parquet'
WITH (format 'parquet');

COPY events_archive
FROM 's3://analytics-bucket/events/july.parquet'
WITH (format 'parquet');
```

CSV and compressed output use standard-looking `COPY` options extended for the lake writer:

```sql
COPY (SELECT * FROM daily_summary)
TO 's3://analytics-bucket/summary/daily.csv.gz'
WITH (format 'csv', header true, compression 'gzip');
```

The destination can be a PostgreSQL heap table or an Iceberg table; the source can likewise be any query supported by the installed pg_lake stack.

### Format and Runtime Boundaries

- Parquet is columnar and preserves supported typed values; CSV and newline-delimited JSON have format-specific inference and conversion options documented upstream.
- Object-store access runs through `pgduck_server`. Its credential chain, network access, and bucket permissions must permit the requested read or write.
- `COPY` is one statement and participates in the surrounding PostgreSQL transaction, but remote files and cleanup also depend on the pg_lake transaction/queue machinery. Inspect failed operations and orphan cleanup before retrying large exports.
- Version `3.4` adds no user-visible SQL objects in `pg_lake_copy`; its `3.3` to `3.4` upgrade script is empty.
