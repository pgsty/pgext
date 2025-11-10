---
title: "pg4ml"
linkTitle: "pg4ml"
description: "Machine learning framework for PostgreSQL"
weight: 1880
categories: ["RAG"]
width: full
---

[**pg4ml**](https://gitee.com/guotiecheng/plpgsql_pg4ml) : Machine learning framework for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1880** | {{< badge content="pg4ml" link="https://gitee.com/guotiecheng/plpgsql_pg4ml" >}} | {{< ext "pg4ml" >}} | `2.0` | {{< category "RAG" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} {{< ext "tablefunc" >}} {{< ext "cube" >}} {{< ext "plpython3u" >}} |
|   **See Also**    | {{< ext "pgml" >}} {{< ext "vectorize" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "pg_strom" >}} |

> [!Note] require python3


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg4ml` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "18" "pg4ml_18" "green" >}} {{< bg "17" "pg4ml_17" "green" >}} {{< bg "16" "pg4ml_16" "green" >}} {{< bg "15" "pg4ml_15" "green" >}} {{< bg "14" "pg4ml_14" "green" >}} {{< bg "13" "pg4ml_13" "green" >}} | `pg4ml_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "18" "postgresql-18-pg4ml" "green" >}} {{< bg "17" "postgresql-17-pg4ml" "green" >}} {{< bg "16" "postgresql-16-pg4ml" "green" >}} {{< bg "15" "postgresql-15-pg4ml" "green" >}} {{< bg "14" "postgresql-14-pg4ml" "green" >}} {{< bg "13" "postgresql-13-pg4ml" "green" >}} | `postgresql-$v-pg4ml` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-13-pg4ml : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-13-pg4ml : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-13-pg4ml : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-13-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-13-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-13-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-13-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-13-pg4ml : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg4ml_18` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 341.2 KiB | [pg4ml_18-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg4ml_18-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg4ml_18` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 341.2 KiB | [pg4ml_18-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg4ml_18-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg4ml_18` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 294.8 KiB | [pg4ml_18-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg4ml_18-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg4ml_18` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 294.7 KiB | [pg4ml_18-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg4ml_18-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg4ml_18` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 294.9 KiB | [pg4ml_18-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg4ml_18-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg4ml_18` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 294.9 KiB | [pg4ml_18-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg4ml_18-2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg4ml` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 316.3 KiB | [postgresql-18-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-18-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg4ml` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 316.3 KiB | [postgresql-18-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-18-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg4ml` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 317.5 KiB | [postgresql-18-pg4ml_2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-18-pg4ml_2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg4ml` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 317.5 KiB | [postgresql-18-pg4ml_2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-18-pg4ml_2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg4ml` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 317.2 KiB | [postgresql-18-pg4ml_2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-18-pg4ml_2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg4ml` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 317.2 KiB | [postgresql-18-pg4ml_2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-18-pg4ml_2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg4ml` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 316.8 KiB | [postgresql-18-pg4ml_2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-18-pg4ml_2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg4ml` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 316.8 KiB | [postgresql-18-pg4ml_2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-18-pg4ml_2.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg4ml_17` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 341.2 KiB | [pg4ml_17-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg4ml_17-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg4ml_17` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 341.1 KiB | [pg4ml_17-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg4ml_17-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg4ml_17` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 294.8 KiB | [pg4ml_17-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg4ml_17-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg4ml_17` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 294.7 KiB | [pg4ml_17-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg4ml_17-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg4ml_17` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 294.9 KiB | [pg4ml_17-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg4ml_17-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg4ml_17` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 294.9 KiB | [pg4ml_17-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg4ml_17-2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg4ml` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 316.3 KiB | [postgresql-17-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-17-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg4ml` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 316.3 KiB | [postgresql-17-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-17-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg4ml` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 317.5 KiB | [postgresql-17-pg4ml_2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-17-pg4ml_2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg4ml` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 317.5 KiB | [postgresql-17-pg4ml_2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-17-pg4ml_2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg4ml` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 317.3 KiB | [postgresql-17-pg4ml_2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-17-pg4ml_2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg4ml` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 317.3 KiB | [postgresql-17-pg4ml_2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-17-pg4ml_2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg4ml` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 316.9 KiB | [postgresql-17-pg4ml_2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-17-pg4ml_2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg4ml` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 316.9 KiB | [postgresql-17-pg4ml_2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-17-pg4ml_2.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg4ml_16` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 341.2 KiB | [pg4ml_16-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg4ml_16-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg4ml_16` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 341.2 KiB | [pg4ml_16-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg4ml_16-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg4ml_16` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 294.8 KiB | [pg4ml_16-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg4ml_16-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg4ml_16` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 294.7 KiB | [pg4ml_16-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg4ml_16-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg4ml_16` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 294.9 KiB | [pg4ml_16-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg4ml_16-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg4ml_16` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 294.9 KiB | [pg4ml_16-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg4ml_16-2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg4ml` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 316.3 KiB | [postgresql-16-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-16-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg4ml` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 316.3 KiB | [postgresql-16-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-16-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg4ml` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 317.5 KiB | [postgresql-16-pg4ml_2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-16-pg4ml_2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg4ml` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 317.5 KiB | [postgresql-16-pg4ml_2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-16-pg4ml_2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg4ml` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 317.3 KiB | [postgresql-16-pg4ml_2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-16-pg4ml_2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg4ml` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 317.3 KiB | [postgresql-16-pg4ml_2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-16-pg4ml_2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg4ml` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 316.9 KiB | [postgresql-16-pg4ml_2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-16-pg4ml_2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg4ml` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 316.9 KiB | [postgresql-16-pg4ml_2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-16-pg4ml_2.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg4ml_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 341.2 KiB | [pg4ml_15-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg4ml_15-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg4ml_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 341.1 KiB | [pg4ml_15-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg4ml_15-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg4ml_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 294.8 KiB | [pg4ml_15-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg4ml_15-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg4ml_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 294.7 KiB | [pg4ml_15-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg4ml_15-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg4ml_15` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 294.9 KiB | [pg4ml_15-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg4ml_15-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg4ml_15` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 294.9 KiB | [pg4ml_15-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg4ml_15-2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg4ml` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 316.3 KiB | [postgresql-15-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-15-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg4ml` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 316.3 KiB | [postgresql-15-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-15-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg4ml` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 317.5 KiB | [postgresql-15-pg4ml_2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-15-pg4ml_2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg4ml` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 317.5 KiB | [postgresql-15-pg4ml_2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-15-pg4ml_2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg4ml` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 317.2 KiB | [postgresql-15-pg4ml_2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-15-pg4ml_2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg4ml` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 317.2 KiB | [postgresql-15-pg4ml_2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-15-pg4ml_2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg4ml` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 316.8 KiB | [postgresql-15-pg4ml_2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-15-pg4ml_2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg4ml` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 316.8 KiB | [postgresql-15-pg4ml_2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-15-pg4ml_2.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg4ml_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 341.2 KiB | [pg4ml_14-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg4ml_14-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg4ml_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 341.1 KiB | [pg4ml_14-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg4ml_14-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg4ml_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 294.8 KiB | [pg4ml_14-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg4ml_14-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg4ml_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 294.8 KiB | [pg4ml_14-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg4ml_14-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg4ml_14` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 294.9 KiB | [pg4ml_14-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg4ml_14-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg4ml_14` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 294.9 KiB | [pg4ml_14-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg4ml_14-2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg4ml` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 316.3 KiB | [postgresql-14-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-14-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg4ml` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 316.3 KiB | [postgresql-14-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-14-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg4ml` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 317.5 KiB | [postgresql-14-pg4ml_2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-14-pg4ml_2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg4ml` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 317.5 KiB | [postgresql-14-pg4ml_2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-14-pg4ml_2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg4ml` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 317.1 KiB | [postgresql-14-pg4ml_2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-14-pg4ml_2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg4ml` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 317.1 KiB | [postgresql-14-pg4ml_2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-14-pg4ml_2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg4ml` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 316.9 KiB | [postgresql-14-pg4ml_2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-14-pg4ml_2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg4ml` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 316.9 KiB | [postgresql-14-pg4ml_2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-14-pg4ml_2.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg4ml_13` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 341.2 KiB | [pg4ml_13-2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg4ml_13-2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg4ml_13` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 341.2 KiB | [pg4ml_13-2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg4ml_13-2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg4ml_13` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 294.8 KiB | [pg4ml_13-2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg4ml_13-2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg4ml_13` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 294.7 KiB | [pg4ml_13-2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg4ml_13-2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg4ml_13` | `2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 294.9 KiB | [pg4ml_13-2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg4ml_13-2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg4ml_13` | `2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 294.9 KiB | [pg4ml_13-2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg4ml_13-2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg4ml` | `2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 316.3 KiB | [postgresql-13-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-13-pg4ml_2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg4ml` | `2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 316.3 KiB | [postgresql-13-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg4ml/postgresql-13-pg4ml_2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg4ml` | `2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 317.5 KiB | [postgresql-13-pg4ml_2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-13-pg4ml_2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg4ml` | `2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 317.5 KiB | [postgresql-13-pg4ml_2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg4ml/postgresql-13-pg4ml_2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg4ml` | `2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 317.2 KiB | [postgresql-13-pg4ml_2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-13-pg4ml_2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg4ml` | `2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 317.2 KiB | [postgresql-13-pg4ml_2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg4ml/postgresql-13-pg4ml_2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg4ml` | `2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 316.9 KiB | [postgresql-13-pg4ml_2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-13-pg4ml_2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg4ml` | `2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 316.9 KiB | [postgresql-13-pg4ml_2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg4ml/postgresql-13-pg4ml_2.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitee.com/guotiecheng/plpgsql_pg4ml" title="Repository" icon="link" subtitle="gitee.com/guotiecheng/plpgsql_pg4ml" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg4ml-2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg4ml;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg4ml;		# install via package name, for the active PG version

pig install pg4ml -v 18;   # install for PG 18
pig install pg4ml -v 17;   # install for PG 17
pig install pg4ml -v 16;   # install for PG 16
pig install pg4ml -v 15;   # install for PG 15
pig install pg4ml -v 14;   # install for PG 14
pig install pg4ml -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg4ml CASCADE; -- requires plpgsql, tablefunc, cube, plpython3u
```
