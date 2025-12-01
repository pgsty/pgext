---
title: "documentdb"
linkTitle: "documentdb"
description: "API surface for DocumentDB for PostgreSQL"
weight: 9000
categories: ["SIM"]
width: full
---

[**documentdb**](https://github.com/microsoft/documentdb) : API surface for DocumentDB for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9000** | {{< badge content="documentdb" link="https://github.com/microsoft/documentdb" >}} | {{< ext "documentdb" >}} | `0.107` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "documentdb_core" >}} {{< ext "pg_cron" >}} {{< ext "tsm_system_rows" >}} {{< ext "vector" >}} {{< ext "postgis" >}} {{< ext "rum" >}} |
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "wal2mongo" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} |
|    **Siblings**   | {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] FerretDB Fork, work with FerretDB 2.7


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.107` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `documentdb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.107` | {{< bg "18" "documentdb_18*" "red" >}} {{< bg "17" "documentdb_17*" "green" >}} {{< bg "16" "documentdb_16*" "green" >}} {{< bg "15" "documentdb_15*" "green" >}} {{< bg "14" "documentdb_14*" "red" >}} {{< bg "13" "documentdb_13*" "red" >}} | `documentdb_$v*` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.107` | {{< bg "18" "postgresql-18-documentdb" "red" >}} {{< bg "17" "postgresql-17-documentdb" "green" >}} {{< bg "16" "postgresql-16-documentdb" "green" >}} {{< bg "15" "postgresql-15-documentdb" "green" >}} {{< bg "14" "postgresql-14-documentdb" "red" >}} {{< bg "13" "postgresql-13-documentdb" "red" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "documentdb_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "documentdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "documentdb_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "documentdb_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "documentdb_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-documentdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.107" "postgresql-17-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-16-documentdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.107" "postgresql-15-documentdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-documentdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-documentdb : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_17` | `0.107` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [documentdb_17-0.107-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-0.107-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_17` | `0.107` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.7 MiB | [documentdb_17-0.107-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-0.107-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_17` | `0.107` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [documentdb_17-0.107-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-0.107-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_17` | `0.107` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [documentdb_17-0.107-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-0.107-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_17` | `0.107` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [documentdb_17-0.107-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_17-0.107-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_17` | `0.107` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [documentdb_17-0.107-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_17-0.107-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-documentdb` | `0.107` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.7 MiB | [postgresql-17-documentdb_0.107-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.107-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-documentdb` | `0.107` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.5 MiB | [postgresql-17-documentdb_0.107-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.107-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-documentdb` | `0.107` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.7 MiB | [postgresql-17-documentdb_0.107-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-17-documentdb_0.107-0PIGSTY~trixie_amd64.deb) |
| `postgresql-17-documentdb` | `0.107` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.5 MiB | [postgresql-17-documentdb_0.107-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-17-documentdb_0.107-0PIGSTY~trixie_arm64.deb) |
| `postgresql-17-documentdb` | `0.107` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.3 MiB | [postgresql-17-documentdb_0.107-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.107-0PIGSTY~jammy_amd64.deb) |
| `postgresql-17-documentdb` | `0.107` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 MiB | [postgresql-17-documentdb_0.107-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.107-0PIGSTY~jammy_arm64.deb) |
| `postgresql-17-documentdb` | `0.107` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 MiB | [postgresql-17-documentdb_0.107-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.107-0PIGSTY~noble_amd64.deb) |
| `postgresql-17-documentdb` | `0.107` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.8 MiB | [postgresql-17-documentdb_0.107-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.107-0PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_16` | `0.107` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [documentdb_16-0.107-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-0.107-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_16` | `0.107` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.7 MiB | [documentdb_16-0.107-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-0.107-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_16` | `0.107` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [documentdb_16-0.107-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-0.107-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_16` | `0.107` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [documentdb_16-0.107-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-0.107-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_16` | `0.107` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [documentdb_16-0.107-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_16-0.107-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_16` | `0.107` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.6 MiB | [documentdb_16-0.107-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_16-0.107-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-documentdb` | `0.107` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.7 MiB | [postgresql-16-documentdb_0.107-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.107-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-documentdb` | `0.107` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.5 MiB | [postgresql-16-documentdb_0.107-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.107-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-documentdb` | `0.107` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.7 MiB | [postgresql-16-documentdb_0.107-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-16-documentdb_0.107-0PIGSTY~trixie_amd64.deb) |
| `postgresql-16-documentdb` | `0.107` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.6 MiB | [postgresql-16-documentdb_0.107-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-16-documentdb_0.107-0PIGSTY~trixie_arm64.deb) |
| `postgresql-16-documentdb` | `0.107` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.3 MiB | [postgresql-16-documentdb_0.107-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.107-0PIGSTY~jammy_amd64.deb) |
| `postgresql-16-documentdb` | `0.107` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 MiB | [postgresql-16-documentdb_0.107-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.107-0PIGSTY~jammy_arm64.deb) |
| `postgresql-16-documentdb` | `0.107` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 MiB | [postgresql-16-documentdb_0.107-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.107-0PIGSTY~noble_amd64.deb) |
| `postgresql-16-documentdb` | `0.107` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.8 MiB | [postgresql-16-documentdb_0.107-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.107-0PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `documentdb_15` | `0.107` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.9 MiB | [documentdb_15-0.107-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-0.107-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_15` | `0.107` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.7 MiB | [documentdb_15-0.107-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-0.107-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_15` | `0.107` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [documentdb_15-0.107-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-0.107-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_15` | `0.107` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [documentdb_15-0.107-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-0.107-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_15` | `0.107` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [documentdb_15-0.107-0PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/documentdb_15-0.107-0PIGSTY.el10.x86_64.rpm) |
| `documentdb_15` | `0.107` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.6 MiB | [documentdb_15-0.107-0PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/documentdb_15-0.107-0PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-documentdb` | `0.107` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.7 MiB | [postgresql-15-documentdb_0.107-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.107-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-documentdb` | `0.107` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.6 MiB | [postgresql-15-documentdb_0.107-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.107-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-documentdb` | `0.107` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.8 MiB | [postgresql-15-documentdb_0.107-0PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-15-documentdb_0.107-0PIGSTY~trixie_amd64.deb) |
| `postgresql-15-documentdb` | `0.107` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.6 MiB | [postgresql-15-documentdb_0.107-0PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/documentdb/postgresql-15-documentdb_0.107-0PIGSTY~trixie_arm64.deb) |
| `postgresql-15-documentdb` | `0.107` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.4 MiB | [postgresql-15-documentdb_0.107-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.107-0PIGSTY~jammy_amd64.deb) |
| `postgresql-15-documentdb` | `0.107` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.3 MiB | [postgresql-15-documentdb_0.107-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.107-0PIGSTY~jammy_arm64.deb) |
| `postgresql-15-documentdb` | `0.107` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 MiB | [postgresql-15-documentdb_0.107-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.107-0PIGSTY~noble_amd64.deb) |
| `postgresql-15-documentdb` | `0.107` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.8 MiB | [postgresql-15-documentdb_0.107-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.107-0PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/microsoft/documentdb" title="Repository" icon="github" subtitle="github.com/microsoft/documentdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="documentdb-0.107.0-ferretdb-2.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg documentdb;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install documentdb;		# install via package name, for the active PG version

pig install documentdb -v 17;   # install for PG 17
pig install documentdb -v 16;   # install for PG 16
pig install documentdb -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_documentdb, pg_documentdb_core, pg_cron';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION documentdb CASCADE; -- requires documentdb_core, pg_cron, tsm_system_rows, vector, postgis, rum
```


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