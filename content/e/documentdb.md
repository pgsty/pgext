---
title: "documentdb"
linkTitle: "documentdb"
description: "API surface for DocumentDB for PostgreSQL"
weight: 9000
categories: ["SIM"]
width: full
---

API surface for DocumentDB for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9000** | {{< badge content="documentdb" link="https://github.com/microsoft/documentdb" >}} | {{< ext "documentdb" >}} | `0.106` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "documentdb_core" >}} {{< ext "pg_cron" >}} {{< ext "tsm_system_rows" >}} {{< ext "vector" >}} {{< ext "postgis" >}} {{< ext "rum" >}} |
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "wal2mongo" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} |
|    **Siblings**   | {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] FerretDB Fork, work with FerretDB 2.5


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/documentdb" >}} | `0.106` | {{< bg "18" "documentdb_18*" "red" >}} {{< bg "17" "documentdb_17*" "green" >}} {{< bg "16" "documentdb_16*" "green" >}} {{< bg "15" "documentdb_15*" "green" >}} {{< bg "14" "documentdb_14*" "red" >}} {{< bg "13" "documentdb_13*" "red" >}} | `documentdb_$v*` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/documentdb" >}} | `0.106` | {{< bg "18" "postgresql-18-documentdb" "red" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} {{< bg "13" "postgresql-13-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 2" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 2" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 2" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 2" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.106" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.106" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_17` | 0.106 | `el8.x86_64` | pigsty | 2.6 MiB | [documentdb_17-0.106-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-0.106-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_17` | 0.105 | `el8.x86_64` | pigsty | 2.5 MiB | [documentdb_17-0.105-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-0.105-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_17` | 0.106 | `el8.aarch64` | pigsty | 2.4 MiB | [documentdb_17-0.106-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-0.106-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_17` | 0.105 | `el8.aarch64` | pigsty | 2.4 MiB | [documentdb_17-0.105-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-0.105-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_17` | 0.106 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_17-0.106-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-0.106-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_17` | 0.105 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_17-0.105-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-0.105-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_17` | 0.106 | `el9.aarch64` | pigsty | 2.3 MiB | [documentdb_17-0.106-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-0.106-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_17` | 0.105 | `el9.aarch64` | pigsty | 2.4 MiB | [documentdb_17-0.105-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-0.105-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_17` | 0.106 | `el10.x86_64` | pigsty | 2.4 MiB | [documentdb_17-0.106-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_17-0.106-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_17` | 0.106 | `el10.aarch64` | pigsty | 2.4 MiB | [documentdb_17-0.106-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_17-0.106-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-documentdb` | 0.106 | `d12.x86_64` | pigsty | 4.3 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-documentdb` | 0.106 | `d12.aarch64` | pigsty | 4.1 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-documentdb` | 0.106 | `u22.x86_64` | pigsty | 4.8 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~jammy_amd64.deb) |
| `postgresql-17-documentdb` | 0.106 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~jammy_arm64.deb) |
| `postgresql-17-documentdb` | 0.106 | `u24.x86_64` | pigsty | 4.4 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~noble_amd64.deb) |
| `postgresql-17-documentdb` | 0.106 | `u24.aarch64` | pigsty | 4.4 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_16` | 0.106 | `el8.x86_64` | pigsty | 2.6 MiB | [documentdb_16-0.106-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-0.106-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_16` | 0.105 | `el8.x86_64` | pigsty | 2.5 MiB | [documentdb_16-0.105-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-0.105-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_16` | 0.106 | `el8.aarch64` | pigsty | 2.5 MiB | [documentdb_16-0.106-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-0.106-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_16` | 0.105 | `el8.aarch64` | pigsty | 2.4 MiB | [documentdb_16-0.105-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-0.105-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_16` | 0.106 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_16-0.106-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-0.106-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_16` | 0.105 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_16-0.105-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-0.105-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_16` | 0.106 | `el9.aarch64` | pigsty | 2.3 MiB | [documentdb_16-0.106-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-0.106-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_16` | 0.105 | `el9.aarch64` | pigsty | 2.4 MiB | [documentdb_16-0.105-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-0.105-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_16` | 0.106 | `el10.x86_64` | pigsty | 2.4 MiB | [documentdb_16-0.106-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_16-0.106-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_16` | 0.106 | `el10.aarch64` | pigsty | 2.4 MiB | [documentdb_16-0.106-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_16-0.106-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-documentdb` | 0.106 | `d12.x86_64` | pigsty | 4.3 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-documentdb` | 0.106 | `d12.aarch64` | pigsty | 4.1 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-documentdb` | 0.106 | `u22.x86_64` | pigsty | 4.8 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~jammy_amd64.deb) |
| `postgresql-16-documentdb` | 0.106 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~jammy_arm64.deb) |
| `postgresql-16-documentdb` | 0.106 | `u24.x86_64` | pigsty | 4.4 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~noble_amd64.deb) |
| `postgresql-16-documentdb` | 0.106 | `u24.aarch64` | pigsty | 4.4 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_15` | 0.106 | `el8.x86_64` | pigsty | 2.6 MiB | [documentdb_15-0.106-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-0.106-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_15` | 0.105 | `el8.x86_64` | pigsty | 2.5 MiB | [documentdb_15-0.105-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-0.105-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_15` | 0.106 | `el8.aarch64` | pigsty | 2.5 MiB | [documentdb_15-0.106-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-0.106-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_15` | 0.105 | `el8.aarch64` | pigsty | 2.4 MiB | [documentdb_15-0.105-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-0.105-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_15` | 0.106 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_15-0.106-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-0.106-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_15` | 0.105 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_15-0.105-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-0.105-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_15` | 0.106 | `el9.aarch64` | pigsty | 2.3 MiB | [documentdb_15-0.106-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-0.106-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_15` | 0.105 | `el9.aarch64` | pigsty | 2.4 MiB | [documentdb_15-0.105-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-0.105-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_15` | 0.106 | `el10.x86_64` | pigsty | 2.4 MiB | [documentdb_15-0.106-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_15-0.106-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_15` | 0.106 | `el10.aarch64` | pigsty | 2.4 MiB | [documentdb_15-0.106-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_15-0.106-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-documentdb` | 0.106 | `d12.x86_64` | pigsty | 4.3 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-documentdb` | 0.106 | `d12.aarch64` | pigsty | 4.1 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-documentdb` | 0.106 | `u22.x86_64` | pigsty | 4.8 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~jammy_amd64.deb) |
| `postgresql-15-documentdb` | 0.106 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~jammy_arm64.deb) |
| `postgresql-15-documentdb` | 0.106 | `u24.x86_64` | pigsty | 4.4 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~noble_amd64.deb) |
| `postgresql-15-documentdb` | 0.106 | `u24.aarch64` | pigsty | 4.4 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/microsoft/documentdb" title="Repository" icon="github" subtitle="github.com/microsoft/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.106.0-ferretdb-2.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get documentdb; # get documentdb source code
pig build dep documentdb; # install build dependencies
pig build pkg documentdb; # build extension rpm or deb
pig build ext documentdb; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install documentdb; # install by extension name, for the current active PG version
pig ext install documentdb; # install via package alias, for the active PG version
pig ext install documentdb -v 17;   # install for PG 17
pig ext install documentdb -v 16;   # install for PG 16
pig ext install documentdb -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION documentdb;
```




--------

## Usage

Add to shared_preload_libraries first:

```bash
-    shared_preload_libraries: documentdb_core, pg_stat_statements, auto_explain
+    shared_preload_libraries: pg_documentdb_core, pg_stat_statements, auto_explain
```

Example, create extension and perform DDL & CRUD

```sql
CREATE EXTENSION documentdb_core ;
```

Currently we only have documentdb_core extension, It can be used with the FerretDB 2.0+