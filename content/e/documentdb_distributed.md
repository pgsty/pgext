---
title: "documentdb_distributed"
linkTitle: "documentdb_distributed"
description: "Multi-Node API surface for DocumentDB"
weight: 9020
categories: ["SIM"]
width: full
---

[**documentdb**](https://github.com/microsoft/documentdb)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9020** | {{< badge content="documentdb_distributed" link="https://github.com/microsoft/documentdb" >}} | {{< ext "documentdb_distributed" "documentdb" >}} | `0.106` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "citus" >}} {{< ext "mongo_fdw" >}} {{< ext "plproxy" >}} {{< ext "postgres_fdw" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} |
|    **Siblings**   | {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} |

> [!Note] FerretDB Fork, work with FerretDB 2.5


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/documentdb" >}} | `0.106` | {{< bg "18" "documentdb_18*" "red" >}} {{< bg "17" "documentdb_17*" "green" >}} {{< bg "16" "documentdb_16*" "green" >}} {{< bg "15" "documentdb_15*" "green" >}} {{< bg "14" "documentdb_14*" "red" >}} {{< bg "13" "documentdb_13*" "red" >}} | `documentdb_$v*` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/documentdb" >}} | `0.106` | {{< bg "18" "postgresql-18-documentdb" "red" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} {{< bg "13" "postgresql-13-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/microsoft/documentdb" title="Repository" icon="github" subtitle="github.com/microsoft/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.105.0-ferretdb-2.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get documentdb_distributed; # get documentdb_distributed source code
pig build dep documentdb_distributed; # install build dependencies
pig build pkg documentdb_distributed; # build extension rpm or deb
pig build ext documentdb_distributed; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install documentdb_distributed; # install by extension name, for the current active PG version
pig ext install documentdb; # install via package alias, for the active PG version
pig ext install documentdb_distributed -v 17;   # install for PG 17
pig ext install documentdb_distributed -v 16;   # install for PG 16
pig ext install documentdb_distributed -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION documentdb_distributed;
```

