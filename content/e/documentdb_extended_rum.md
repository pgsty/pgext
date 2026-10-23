---
title: "documentdb_extended_rum"
linkTitle: "documentdb_extended_rum"
description: "DocumentDB Extended RUM index access method"
weight: 9030
categories: ["SIM"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**documentdb**](https://github.com/documentdb/documentdb) : DocumentDB Extended RUM index access method


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9030** | {{< badge content="documentdb_extended_rum" link="https://github.com/documentdb/documentdb" >}} | {{< ext "documentdb_extended_rum" "documentdb" >}} | `1.0` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `documentdb_extended_rum_catalog` |
|   **Requires**    | {{< ext "documentdb" >}} |
|   **See Also**    | {{< ext "rum" >}} {{< ext "pg_projection" >}} {{< ext "jsquery" >}} {{< ext "pgbson" >}} {{< ext "mongo_fdw" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsonschema" >}} {{< ext "pgjq" >}} {{< ext "hll" >}} {{< ext "roaringbitmap" >}} |
|    **Siblings**   | {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] DocumentDB 1.0 packages install SQL 1.0-0. PGSTY supports PostgreSQL 16-18; PG15 is excluded due to the libbson allocation alignment failure. The OSS package includes documentdb, documentdb_core and documentdb_extended_rum only. External RUM is a package dependency on PG16/17 only, not PG18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `documentdb` | `documentdb` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "documentdb_18" "green" >}} {{< bg "17" "documentdb_17" "green" >}} {{< bg "16" "documentdb_16" "green" >}} {{< bg "15" "documentdb_15" "red" >}} {{< bg "14" "documentdb_14" "red" >}} | `documentdb_$v` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v`, `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-documentdb" "green" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "red" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum`, `postgresql-$v-postgis-3` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 4" "green" >}} | {{< bg "PGDG 1.0" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 4" "green" >}} | {{< bg "PGDG 1.0" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.117" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 4" "green" >}} | {{< bg "PGDG 1.0" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-documentdb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-documentdb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-documentdb : AVAIL 4" "green" >}} | {{< bg "PGDG 1.0" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/documentdb/documentdb" title="Repository" icon="github" subtitle="github.com/documentdb/documentdb" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-1.0-0.tar.gz intelrdfpmath-applied-2.0u3-1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg documentdb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install documentdb;		# install via package name, for the active PG version
pig install documentdb_extended_rum;		# install by extension name, for the current active PG version

pig install documentdb_extended_rum -v 18;   # install for PG 18
pig install documentdb_extended_rum -v 17;   # install for PG 17
pig install documentdb_extended_rum -v 16;   # install for PG 16

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_extended_rum';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION documentdb_extended_rum CASCADE; -- requires documentdb
```

## Usage

Sources:

- [DocumentDB Extended RUM README](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/README.md)
- [`documentdb_extended_rum` control file](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)
- [Access-method SQL definitions](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/sql/documentdb_extended_rum--0.106-0.sql)
- [DocumentDB v1.0-0 changelog](https://github.com/documentdb/documentdb/blob/v1.0-0/CHANGELOG.md)

`documentdb_extended_rum` is DocumentDB's extended RUM index access method. It is an implementation component selected by DocumentDB's indexing layer, not a general-purpose application index or a replacement for installing `documentdb`.

### Configure and Install

The library can only be initialized from `shared_preload_libraries`. Preload it after the base DocumentDB libraries and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_extended_rum'
documentdb.rum_library_load_option = 'require_documentdb_extended_rum'
```

Then install the extension using the same release as the base stack:

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_extended_rum;
```

DocumentDB deployment tooling normally owns this configuration. Existing databases should follow the release-specific upgrade procedure rather than switching an index handler ad hoc.

### Important Objects

- `documentdb_extended_rum` is the index access method registered by the extension.
- `documentdb_extended_rum_catalog` contains BSON operator families and classes used by DocumentDB.
- `documentdb.rum_library_load_option = 'require_documentdb_extended_rum'` directs the DocumentDB index layer to the adapter.
- The implementation is a RUM fork whose on-disk layout and content are designed to remain backward compatible with upstream RUM while changing query and volatile paths for document workloads.

### Operational Boundaries

Install and upgrade this component with matching `documentdb` and `documentdb_core` binaries. Do not build indexes with its internal operator classes directly unless following upstream development guidance; create and manage indexes through the DocumentDB APIs so metadata stays consistent.

In 1.0-0, `documentdb.rum_library_load_option` continues to default to `require_documentdb_extended_rum`. The release fixes concurrent root-split handling in vacuum, restores vacuum delays and interrupt checks, and enables targeted posting-tree pruning and empty entry-leaf pruning by default. Apply the matching upgrade SQL and validate existing indexes through supported DocumentDB operations.
