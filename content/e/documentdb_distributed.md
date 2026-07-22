---
title: "documentdb_distributed"
linkTitle: "documentdb_distributed"
description: "Multi-Node API surface for DocumentDB"
weight: 9020
categories: ["SIM"]
width: full
---

[**documentdb**](https://github.com/documentdb/documentdb) : Multi-Node API surface for DocumentDB


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9020** | {{< badge content="documentdb_distributed" link="https://github.com/documentdb/documentdb" >}} | {{< ext "documentdb_distributed" "documentdb" >}} | `0.114` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "citus" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb" >}} |
|   **See Also**    | {{< ext "citus" >}} {{< ext "mongo_fdw" >}} {{< ext "plproxy" >}} {{< ext "postgres_fdw" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} |
|    **Siblings**   | {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_extended_rum" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `documentdb` | `citus`, `documentdb_core`, `documentdb` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "documentdb_18" "green" >}} {{< bg "17" "documentdb_17" "green" >}} {{< bg "16" "documentdb_16" "green" >}} {{< bg "15" "documentdb_15" "green" >}} {{< bg "14" "documentdb_14" "red" >}} | `documentdb_$v` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v`, `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "postgresql-18-documentdb" "green" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum`, `postgresql-$v-postgis-3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "documentdb_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} | {{< bg "N/A" "postgresql-14-documentdb : N/A 0" "gray" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/documentdb/documentdb" title="Repository" icon="github" subtitle="github.com/documentdb/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.114-0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg documentdb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install documentdb;		# install via package name, for the active PG version
pig install documentdb_distributed;		# install by extension name, for the current active PG version

pig install documentdb_distributed -v 18;   # install for PG 18
pig install documentdb_distributed -v 17;   # install for PG 17
pig install documentdb_distributed -v 16;   # install for PG 16
pig install documentdb_distributed -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'citus, pg_documentdb, pg_documentdb_core';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION documentdb_distributed CASCADE; -- requires citus, documentdb_core, documentdb
```

## Usage

Sources:

- [`documentdb_distributed` v0.114-0 control file](https://github.com/documentdb/documentdb/blob/v0.114-0/internal/pg_documentdb_distributed/documentdb_distributed.control)
- [DocumentDB package scope](https://github.com/documentdb/documentdb/blob/v0.114-0/packaging/README.md)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v0.114-0/scripts/preload_libraries.sh)
- [DocumentDB v0.114-0 changelog](https://github.com/documentdb/documentdb/blob/v0.114-0/CHANGELOG.md)

`documentdb_distributed` is DocumentDB's internal multi-node execution layer. It integrates the public `documentdb` API with Citus; it is not a standalone document API and does not add a separate client workflow.

### Prerequisites and Install

All nodes need matching builds of Citus, `documentdb_core`, `documentdb`, and `documentdb_distributed`. The official helper places the libraries in preload order; restart every node after changing it:

```conf
shared_preload_libraries = 'citus, pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_distributed'
```

After the Citus topology and the base DocumentDB stack are configured, install the distributed component as a superuser:

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_distributed;

SELECT extname, extversion
FROM pg_extension
WHERE extname IN ('citus', 'documentdb_core', 'documentdb', 'documentdb_distributed');
```

Use the normal DocumentDB gateway or `documentdb_api` functions after installation. Collection placement, shard topology, worker availability, and metadata consistency must be managed as part of the cluster deployment.

### Important Boundaries

- The control file requires `citus`, `documentdb_core`, and `documentdb` and marks the extension superuser-only and non-relocatable.
- `documentdb_distributed` supplies distributed planner and execution support to existing DocumentDB commands; it is not an alternative to the public API extension.
- Release versions should remain synchronized across coordinators and workers before extension upgrades are attempted.
- Backup, restore, failover, and rolling-upgrade procedures must include both Citus metadata and DocumentDB data.

The upstream `packaging/README.md` explicitly says the standard packages do not include the `internal/pg_documentdb_distributed` component. Confirm that a distribution actually ships this extension before adding it to configuration. Version 0.114-0 includes a sharded `$sample` optimization fix behind a feature flag; it should not be treated as unconditional behavior.
