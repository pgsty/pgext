---
title: "documentdb_core"
linkTitle: "documentdb_core"
description: "Core API surface for DocumentDB on PostgreSQL"
weight: 9010
categories: ["SIM"]
width: full
---

[**documentdb**](https://github.com/documentdb/documentdb) : Core API surface for DocumentDB on PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9010** | {{< badge content="documentdb_core" link="https://github.com/documentdb/documentdb" >}} | {{< ext "documentdb_core" "documentdb" >}} | `0.114` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} {{< ext "documentdb_distributed" >}} |
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_cron" >}} {{< ext "postgis" >}} {{< ext "vector" >}} |
|    **Siblings**   | {{< ext "documentdb" >}} {{< ext "documentdb_distributed" >}} {{< ext "documentdb_extended_rum" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `documentdb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "documentdb_18" "green" >}} {{< bg "17" "documentdb_17" "green" >}} {{< bg "16" "documentdb_16" "green" >}} {{< bg "15" "documentdb_15" "green" >}} {{< bg "14" "documentdb_14" "red" >}} | `documentdb_$v` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v`, `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.114` | {{< bg "18" "postgresql-18-documentdb" "green" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum`, `postgresql-$v-postgis-3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.114" "documentdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.114" "postgresql-18-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.114" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.114" "postgresql-18-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-17-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-16-documentdb : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.114" "postgresql-15-documentdb : AVAIL 4" "blue" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |


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
pig install documentdb_core;		# install by extension name, for the current active PG version

pig install documentdb_core -v 18;   # install for PG 18
pig install documentdb_core -v 17;   # install for PG 17
pig install documentdb_core -v 16;   # install for PG 16
pig install documentdb_core -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_documentdb, pg_documentdb_core';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION documentdb_core;
```

## Usage

Sources:

- [DocumentDB v0.114-0 README](https://github.com/documentdb/documentdb/blob/v0.114-0/README.md)
- [`documentdb_core` control file](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_core/documentdb_core.control)
- [BSON SQL definitions](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_core/sql/udfs/bson_io/bson_io--latest.sql)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v0.114-0/scripts/preload_libraries.sh)

`documentdb_core` is the low-level BSON type and operator layer used by DocumentDB. It is normally installed as a dependency of `documentdb`; by itself it does not provide collection CRUD, the MongoDB wire protocol, or the gateway.

### Configure and Install

`pg_documentdb_core` must be loaded through `shared_preload_libraries`, followed by a PostgreSQL restart:

```conf
shared_preload_libraries = 'pg_documentdb_core'
```

For a complete single-node stack, the official helper also preloads `pg_cron` and `pg_documentdb`. Install the parent extension in normal deployments:

```sql
CREATE EXTENSION documentdb CASCADE;
```

Direct installation is useful only for low-level BSON work:

```sql
CREATE EXTENSION documentdb_core;
```

The extension is superuser-only and non-relocatable.

### BSON Workflow

```sql
SELECT '{"name":"Ada","score":42}'::documentdb_core.bson;

SELECT documentdb_core.bson_get_value_text(
  '{"name":"Ada","score":42}'::documentdb_core.bson,
  'name'
);
```

Use explicit schema qualification unless `documentdb_core` is in `search_path`.

### Important Objects

- `documentdb_core.bson` stores BSON documents.
- `documentdb_core.bsonquery` represents BSON query values used by the DocumentDB planner and operator layer.
- `documentdb_core.bsonsequence` represents sequences of BSON values.
- `bson_get_value` and `bson_get_value_text`, also exposed through `->` and `->>`, extract a path from a BSON document.
- `bson_from_bytea`, `bson_to_bytea`, `bson_json_to_bson`, and `bson_to_json_string` support serialization boundaries.
- `bson_btree_ops` and `bson_hash_ops` provide comparison and hashing support required by higher layers.

### Operational Boundaries

BSON comparison, indexing, and numeric semantics follow DocumentDB's implementation and should not be assumed to match PostgreSQL `jsonb`. Most objects are infrastructure for `documentdb`; applications seeking collections and MongoDB commands should use the parent extension or gateway rather than building directly on internal types.

Version 0.114-0 keeps `documentdb_core` aligned with the rest of the DocumentDB stack. The upstream changelog does not identify a separate end-user core API migration for this release, so no new standalone workflow is claimed.
