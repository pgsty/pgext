---
title: "omni_seq"
linkTitle: "omni_seq"
description: "Distributed integer sequences"
weight: 2964
categories: ["FEAT"]
width: full
---

[**omnigres**](https://docs.omnigres.org/omni_seq/id/) : Distributed integer sequences


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2964** | {{< badge content="omni_seq" link="https://docs.omnigres.org/omni_seq/id/" >}} | {{< ext "omni_seq" "omnigres" >}} | `0.1.1` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `omni_seq` |
|    **Siblings**   | {{< ext "omni" >}} {{< ext "omni_auth" >}} {{< ext "omni_aws" >}} {{< ext "omni_cloudevents" >}} {{< ext "omni_containers" >}} {{< ext "omni_credentials" >}} {{< ext "omni_email" >}} {{< ext "omni_http" >}} {{< ext "omni_httpc" >}} {{< ext "omni_httpd" >}} {{< ext "omni_id" >}} {{< ext "omni_json" >}} {{< ext "omni_kube" >}} {{< ext "omni_ledger" >}} {{< ext "omni_manifest" >}} {{< ext "omni_mimetypes" >}} {{< ext "omni_os" >}} {{< ext "omni_polyfill" >}} {{< ext "omni_python" >}} {{< ext "omni_regex" >}} {{< ext "omni_rest" >}} {{< ext "omni_schema" >}} {{< ext "omni_service" >}} {{< ext "omni_session" >}} {{< ext "omni_sql" >}} {{< ext "omni_sqlite" >}} {{< ext "omni_test" >}} {{< ext "omni_txn" >}} {{< ext "omni_types" >}} {{< ext "omni_var" >}} {{< ext "omni_vfs" >}} {{< ext "omni_vfs_types_v1" >}} {{< ext "omni_web" >}} {{< ext "omni_worker" >}} {{< ext "omni_xml" >}} {{< ext "omni_yaml" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `omnigres` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "omnigres_18" "green" >}} {{< bg "17" "omnigres_17" "green" >}} {{< bg "16" "omnigres_16" "green" >}} {{< bg "15" "omnigres_15" "green" >}} {{< bg "14" "omnigres_14" "green" >}} | `omnigres_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-omnigres" "green" >}} {{< bg "17" "postgresql-17-omnigres" "green" >}} {{< bg "16" "postgresql-16-omnigres" "green" >}} {{< bg "15" "postgresql-15-omnigres" "green" >}} {{< bg "14" "postgresql-14-omnigres" "green" >}} | `postgresql-$v-omnigres` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "omnigres_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250507" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250507" "omnigres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250507" "omnigres_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "omnigres_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "omnigres_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-omnigres : MISS 0" "red" >}}      | {{< bg "PIGSTY 20250120" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20250120" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 20251108" "postgresql-18-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-17-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-16-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-15-omnigres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 20251108" "postgresql-14-omnigres : AVAIL 1" "green" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://docs.omnigres.org/omni_seq/id/" title="Repository" icon="link" subtitle="docs.omnigres.org/omni_seq/id/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="omnigres-20251108.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg omnigres;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install omnigres;		# install via package name, for the active PG version
pig install omni_seq;		# install by extension name, for the current active PG version

pig install omni_seq -v 18;   # install for PG 18
pig install omni_seq -v 17;   # install for PG 17
pig install omni_seq -v 16;   # install for PG 16
pig install omni_seq -v 15;   # install for PG 15
pig install omni_seq -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION omni_seq;
```




## Usage

> [omni_seq: Distributed integer sequences](https://docs.omnigres.org/omni_seq/id/)

The `omni_seq` extension provides distributed (prefixed) identifier types to work around PostgreSQL's limitation where sequence data is not replicated in logical replication.

### Types

Types follow the pattern `omni_seq.id_<PREFIX_TYPE>_<ID_TYPE>`. Supported base types: `int16`, `int32`, `int64`. When prefix and identifier types match, the shorthand `omni_seq.id_int64` is available.

### Create a Table with Distributed IDs

```sql
CREATE SEQUENCE seq;
CREATE TABLE t (
    id omni_seq.id_int64 PRIMARY KEY NOT NULL
       DEFAULT omni_seq.id_int64_nextval(omni_seq.system_identifier(), 'seq')
);
```

IDs display as `PREFIX:IDENTIFIER` (e.g., `7222168279780171472:1`).

### Functions

- **`omni_seq.id_int64_nextval(node_id, sequence_name)`** -- Generates the next distributed ID
- **`omni_seq.system_identifier()`** -- Returns the unique 64-bit system identifier from pg_control
- **`omni_seq.id_int64_int32_make(prefix, identifier)`** -- Constructs a distributed ID from components (useful during migrations)

### Migration Example

```sql
BEGIN;
LOCK TABLE my_table;
ALTER TABLE my_table ALTER COLUMN id DROP IDENTITY IF EXISTS;
CREATE SEQUENCE my_table_id_seq;
ALTER TABLE my_table
    ALTER COLUMN id TYPE omni_seq.id_int64_int32
        USING omni_seq.id_int64_int32_make(0, id),
    ALTER COLUMN id SET DEFAULT omni_seq.id_int64_int32_nextval(
        omni_seq.system_identifier(), 'my_table_id_seq');
COMMIT;
```
