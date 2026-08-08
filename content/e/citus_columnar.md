---
title: "citus_columnar"
linkTitle: "citus_columnar"
description: "Citus columnar storage engine"
weight: 2401
categories: ["OLAP"]
width: full
---

[**citus**](https://github.com/citusdata/citus) : Citus columnar storage engine


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2401** | {{< badge content="citus_columnar" link="https://github.com/citusdata/citus" >}} | {{< ext "citus_columnar" "citus" >}} | `14.2.0` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "pg_mooncake" >}} {{< ext "columnar" >}} {{< ext "storage_engine" >}} {{< ext "orioledb" >}} {{< ext "pg_sorted_heap" >}} |
|    **Siblings**   | {{< ext "citus" >}} |

> [!Note] Packaged with Citus 14.2.0; the control default_version is 14.2-1; citus_columnar itself does not require preload and conflicts with Hydra Columnar.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `citus` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.2.0` | {{< bg "18" "citus_18" "green" >}} {{< bg "17" "citus_17" "green" >}} {{< bg "16" "citus_16" "green" >}} {{< bg "15" "citus_15" "red" >}} {{< bg "14" "citus_14" "red" >}} | `citus_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `14.2.0` | {{< bg "18" "postgresql-18-citus" "green" >}} {{< bg "17" "postgresql-17-citus" "green" >}} {{< bg "16" "postgresql-16-citus" "green" >}} {{< bg "15" "postgresql-15-citus" "red" >}} {{< bg "14" "postgresql-14-citus" "red" >}} | `postgresql-$v-citus` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 14.2.0" "citus_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_16 : AVAIL 16" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 21" "blue" >}} | {{< bg "PGDG 13.0.0" "citus_14 : AVAIL 28" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 14.2.0" "citus_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_16 : AVAIL 16" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 20" "blue" >}} | {{< bg "PGDG 13.0.0" "citus_14 : AVAIL 15" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 14.2.0" "citus_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_17 : AVAIL 11" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_16 : AVAIL 18" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 21" "blue" >}} | {{< bg "PGDG 13.0.0" "citus_14 : AVAIL 25" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 14.2.0" "citus_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_17 : AVAIL 11" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_16 : AVAIL 18" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 21" "blue" >}} | {{< bg "PGDG 13.0.0" "citus_14 : AVAIL 15" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 14.2.0" "citus_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_16 : AVAIL 9" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 4" "blue" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 14.2.0" "citus_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 14.2.0" "citus_16 : AVAIL 9" "green" >}} | {{< bg "PGDG 13.2.0" "citus_15 : AVAIL 4" "blue" >}} | {{< bg "PIGSTY 13.0.0" "citus_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-18-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-17-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 14.2.0" "postgresql-16-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.2.0" "postgresql-15-citus : AVAIL 1" "green" >}} | {{< bg "PIGSTY 13.0.0" "postgresql-14-citus : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/citus" title="Repository" icon="github" subtitle="github.com/citusdata/citus" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="citus-14.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg citus;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install citus;		# install via package name, for the active PG version
pig install citus_columnar;		# install by extension name, for the current active PG version

pig install citus_columnar -v 18;   # install for PG 18
pig install citus_columnar -v 17;   # install for PG 17
pig install citus_columnar -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION citus_columnar;
```

## Usage

Sources:

- [Citus v14.2.0 columnar control file](https://github.com/citusdata/citus/blob/v14.2.0/src/backend/columnar/citus_columnar.control)
- [Citus v14.2.0 columnar option helper](https://github.com/citusdata/citus/blob/v14.2.0/src/backend/columnar/sql/udfs/alter_columnar_table_set/latest.sql)
- [Citus columnar-storage documentation](https://docs.citusdata.com/en/stable/admin_guide/table_management.html#columnar-storage)
- [Citus v14.2.0 release](https://github.com/citusdata/citus/releases/tag/v14.2.0)

`citus_columnar` provides an append-oriented columnar table access method for PostgreSQL. It is shipped by the Citus 14.2 package but is a separate extension: the package release is `14.2.0`, while the extension control version is `14.2-1`. Use it for scan-heavy archival or analytical tables whose workload fits its write and feature restrictions.

### Create a Columnar Table

```sql
CREATE EXTENSION citus_columnar;

CREATE TABLE events_archive (
  event_at timestamptz NOT NULL,
  tenant_id bigint NOT NULL,
  kind text,
  payload jsonb
) USING columnar;
```

`citus_columnar` itself does not require `shared_preload_libraries`. Preloading `citus` is still required when the database also uses the distributed `citus` extension.

### Load and Query Data

Columnar storage groups rows into stripes and compresses columns in chunks. Bulk inserts in reasonably sized transactions produce better stripes than a stream of tiny transactions.

```sql
INSERT INTO events_archive
SELECT event_at, tenant_id, kind, payload
FROM events
WHERE event_at < now() - interval '90 days';

SELECT tenant_id, count(*), min(event_at), max(event_at)
FROM events_archive
GROUP BY tenant_id;
```

### Convert with the Citus Extension

When the main `citus` extension is also preloaded and installed, use its helper to convert a local or distributed table:

```sql
SELECT alter_table_set_access_method('events_archive', 'columnar');
SELECT alter_table_set_access_method('events_archive', 'heap');
```

Conversion rewrites the table. Converting to columnar drops existing indexes, so inventory dependent indexes and constraints before running it and schedule enough disk and lock time for the rewrite.

`alter_table_set_access_method()` belongs to `citus`, not to standalone `citus_columnar`. Without the main extension, create a new `USING columnar` table and copy data into it instead of assuming this helper exists.

### Tune Compression

Inspect and change table-level columnar options with the documented helpers:

```sql
SELECT alter_columnar_table_set(
  'events_archive',
  compression => 'zstd',
  compression_level => 3,
  stripe_row_limit => 150000,
  chunk_group_row_limit => 10000
);
```

New settings affect newly written stripes. Rewrite existing data when old stripes also need the new layout.

### Operational Boundaries

- Columnar tables are intended for append-heavy use. `UPDATE` and `DELETE` are not supported, and space left by rolled-back writes is not reclaimed through ordinary heap-style maintenance.
- TOAST is not available; large values remain inline and can hit PostgreSQL's row-size limits.
- Row locks, `AFTER ... FOR EACH ROW` triggers, serializable isolation, logical decoding, foreign keys, unlogged tables, and several scan types are unsupported. Check the current upstream limitation list before adopting the access method.
- Ordinary heap assumptions about indexes, vacuum, replication, triggers, and constraints do not automatically apply. Validate every required database feature against a representative columnar table.
- The extension installs in `pg_catalog`, is not relocatable, and has SQL version `14.2-1`; use that version when checking or updating `pg_extension`, not the package version `14.2.0`.
