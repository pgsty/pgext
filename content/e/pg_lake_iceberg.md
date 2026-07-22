---
title: "pg_lake_iceberg"
linkTitle: "pg_lake_iceberg"
description: "Iceberg implementation in Postgres"
weight: 2565
categories: ["OLAP"]
width: full
---

[**pg_lake**](https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_iceberg) : Iceberg implementation in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2565** | {{< badge content="pg_lake_iceberg" link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_iceberg" >}} | {{< ext "pg_lake_iceberg" "pg_lake" >}} | `3.4` | {{< category "OLAP" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `lake_iceberg` `pg_catalog` |
|   **Requires**    | {{< ext "pg_lake_engine" >}} {{< ext "plpgsql" >}} |
|    **Need By**    | {{< ext "pg_lake_copy" >}} {{< ext "pg_lake_table" >}} |
|    **Siblings**   | {{< ext "pg_lake" >}} {{< ext "pg_extension_base" >}} {{< ext "pg_extension_updater" >}} {{< ext "pg_map" >}} {{< ext "pg_lake_engine" >}} {{< ext "pg_lake_table" >}} {{< ext "pg_lake_copy" >}} |

> [!Note] The control file declares pg_lake_engine; canonical metadata also retains plpgsql because the install SQL uses PL/pgSQL. plpgsql is installed by default in normal PostgreSQL databases.
Extension SQL/control version is 3.4; source and DEB/RPM package version is 3.4.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lake` | `pg_lake_engine`, `plpgsql` |
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
{{< card link="https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_iceberg" title="Repository" icon="github" subtitle="github.com/Snowflake-Labs/pg_lake/tree/main/pg_lake_iceberg" >}}
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
pig install pg_lake_iceberg;		# install by extension name, for the current active PG version

pig install pg_lake_iceberg -v 18;   # install for PG 18
pig install pg_lake_iceberg -v 17;   # install for PG 17
pig install pg_lake_iceberg -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_lake_iceberg CASCADE; -- requires pg_lake_engine, plpgsql
```

## Usage

Sources:

- [Official Iceberg table guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/iceberg-tables.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_iceberg/pg_lake_iceberg.control)
- [Iceberg metadata SQL API](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_iceberg/pg_lake_iceberg--3.0.sql)
- [Version 3.4 catalog FDW SQL](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake_iceberg/pg_lake_iceberg--3.3--3.4.sql)

`pg_lake_iceberg` implements Iceberg metadata, snapshots, manifests, partition specifications, and catalog integration inside PostgreSQL. The familiar `CREATE TABLE ... USING iceberg` syntax is exposed by the dependent `pg_lake_table` component; users normally install both through `pg_lake`.

### Create and Inspect an Iceberg Table

```sql
CREATE EXTENSION pg_lake CASCADE;

SET pg_lake_iceberg.default_location_prefix =
    's3://analytics-bucket/warehouse';

CREATE TABLE events (
    event_time timestamptz NOT NULL,
    user_id bigint NOT NULL,
    payload jsonb
) USING iceberg
WITH (partition_by = 'day(event_time), bucket(32, user_id)');

SELECT table_namespace, table_name, metadata_location
FROM iceberg_tables
WHERE table_name = 'events';
```

Inspect an Iceberg metadata file and its referenced state:

```sql
SELECT lake_iceberg.metadata(metadata_location)
FROM iceberg_tables
WHERE table_name = 'events';

SELECT f.*
FROM iceberg_tables AS t
CROSS JOIN LATERAL lake_iceberg.files(t.metadata_location) AS f
WHERE t.table_name = 'events';
```

### Metadata and Catalog API

- `iceberg_tables`: `pg_catalog` view combining local managed tables and external catalog entries.
- `iceberg_namespace_properties`: catalog namespace properties.
- `lake_iceberg.metadata(uri)`: raw Iceberg metadata JSON.
- `lake_iceberg.files(uri)`: manifest path, content kind, data-file path/format, spec ID, record count, and file size.
- `lake_iceberg.snapshots(uri)`: sequence number, snapshot ID, timestamp, and manifest-list path.
- `lake_iceberg.data_file_stats(uri)`: per-file sequence and lower/upper bounds; execution is granted to `lake_read` rather than `PUBLIC`.
- `iceberg_catalog`: version 3.4 FDW for named PostgreSQL, object-store, or REST catalog configurations.

Define a user-managed REST catalog server and keep credentials in a user mapping:

```sql
CREATE SERVER my_polaris TYPE 'rest'
FOREIGN DATA WRAPPER iceberg_catalog
OPTIONS (rest_endpoint 'https://polaris.example.com');

CREATE USER MAPPING FOR app_role SERVER my_polaris
OPTIONS (client_id 'app', client_secret 'secret');

CREATE TABLE catalog_events (id bigint)
USING iceberg
WITH (catalog = 'my_polaris');
```

### Catalog and Storage Caveats

- User-created catalog servers require their own `USER MAPPING` credentials and do not fall back to the built-in REST catalog credential GUCs.
- The built-in `postgres`, `object_store`, and `rest` catalogs map to immutable extension-owned servers. Configure them through the documented GUCs rather than altering those servers.
- External modifications to `iceberg_tables` are blocked by default because changing metadata behind pg_lake can break transaction and query-engine consistency.
- Iceberg writes should be batched. Each statement can add Parquet files and snapshots; regular `VACUUM` compacts small files and expires data according to table/GUC policy.
- Iceberg has narrower representations for some PostgreSQL values. The default `out_of_range_values = 'error'` preserves integrity; `clamp` silently changes out-of-range temporal values and replaces some unsupported values with `NULL`.
