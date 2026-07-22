---
title: "pg_lake_table"
linkTitle: "pg_lake_table"
description: "Data lake tables and Iceberg tables"
weight: 2566
categories: ["OLAP"]
width: full
---

[**pg_lake**](https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_table) : Data lake tables and Iceberg tables


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2566** | {{< badge content="pg_lake_table" link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_table" >}} | {{< ext "pg_lake_table" "pg_lake" >}} | `3.4` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `__pg_lake_table_writes` `lake_file` `lake_file_cache` `lake_table` `pg_catalog` |
|   **Requires**    | {{< ext "btree_gist" >}} {{< ext "pg_lake_engine" >}} {{< ext "pg_lake_iceberg" >}} |
|    **Need By**    | {{< ext "pg_lake" >}} {{< ext "pg_lake_copy" >}} |
|    **Siblings**   | {{< ext "pg_lake" >}} {{< ext "pg_extension_base" >}} {{< ext "pg_extension_updater" >}} {{< ext "pg_map" >}} {{< ext "pg_lake_engine" >}} {{< ext "pg_lake_iceberg" >}} {{< ext "pg_lake_copy" >}} |

> [!Note] pg_extension_base auto-loads pg_lake_engine, pg_lake_iceberg, and pg_lake_table in dependency order.
Extension SQL/control version is 3.4; source and DEB/RPM package version is 3.4.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lake` | `btree_gist`, `pg_lake_engine`, `pg_lake_iceberg` |
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
{{< card link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_table" title="Repository" icon="github" subtitle="github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_table" >}}
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
pig install pg_lake_table;		# install by extension name, for the current active PG version

pig install pg_lake_table -v 18;   # install for PG 18
pig install pg_lake_table -v 17;   # install for PG 17
pig install pg_lake_table -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_extension_base';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_lake_table CASCADE; -- requires btree_gist, pg_lake_engine, pg_lake_iceberg
```

## Usage

Sources:

- [Official data-lake file query guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/query-data-lake-files.md)
- [Official Iceberg table guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/iceberg-tables.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_table/pg_lake_table.control)
- [FDW, server, utility, and access-method SQL](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_table/pg_lake_table--3.0.sql)

`pg_lake_table` exposes object-store files as PostgreSQL foreign tables and provides the `USING iceberg` table syntax. It owns the `pg_lake` and `pg_lake_iceberg` foreign servers, file inspection/cache utilities, table catalogs, and transaction hooks; Iceberg metadata encoding is delegated to `pg_lake_iceberg`.

### Query External Files

Install the complete stack, including the required `pg_lake_engine`, `pg_lake_iceberg`, and `btree_gist` dependencies:

```sql
CREATE EXTENSION pg_lake CASCADE;
```

An empty column list asks pg_lake to infer the file schema:

```sql
CREATE FOREIGN TABLE external_events ()
SERVER pg_lake
OPTIONS (
    path 's3://analytics-bucket/events/*.parquet',
    filename 'true'
);

SELECT count(*) FROM external_events;
```

Create a writable Iceberg table through the extension-provided table access method:

```sql
CREATE TABLE managed_events (
    event_time timestamptz,
    payload jsonb
) USING iceberg;
```

### File and Table Utility Index

- `lake_file.list(pattern)`: lists matching object paths, sizes, modification times, and ETags.
- `lake_file.size(path)` and `lake_file.exists(path)`: inspect one remote object.
- `lake_file.preview(url, format, compression)`: returns inferred column names and types.
- `lake_file.delete(url)`: deletes a remote object; restrict it to roles that are allowed to remove data.
- `lake_file_cache.add(path, refresh)`, `remove(path)`, and `list()`: manage the local file cache for members of `lake_read`.
- `lake_iceberg.table_size(regclass)`: totals the current data-file sizes of an Iceberg table.
- Foreign servers `pg_lake` and `pg_lake_iceberg`: read-file and writable-Iceberg entry points.
- Access methods `iceberg` and `pg_lake_iceberg`: aliases used to translate `CREATE TABLE ... USING` into the extension's foreign-table representation.

```sql
SELECT path, file_size
FROM lake_file.list('s3://analytics-bucket/events/**/*.parquet');

SELECT *
FROM lake_file.preview('s3://analytics-bucket/events/sample.parquet');
```

### Operational Caveats

- `pg_extension_base` must be preloaded and `pgduck_server` must be running with credentials for every referenced location.
- `lake_read`, `lake_write`, and `lake_read_write` control access to servers, schemas, and utilities. Grant the narrowest role required by each application.
- External tables are references to files, not imported copies. File replacement, cross-region access, and cache invalidation can change latency or results independently of PostgreSQL catalog state.
- Iceberg inserts are optimized for batches rather than single rows. Use a staging heap table for high-rate row-at-a-time ingestion and periodically flush batches.
- Internal `lake_table.*` catalogs track files, field IDs, partitions, and recovery state. Do not modify them directly.
