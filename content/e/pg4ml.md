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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg4ml` | `plpgsql`, `tablefunc`, `cube`, `plpython3u` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "18" "pg4ml_18" "green" >}} {{< bg "17" "pg4ml_17" "green" >}} {{< bg "16" "pg4ml_16" "green" >}} {{< bg "15" "pg4ml_15" "green" >}} {{< bg "14" "pg4ml_14" "green" >}} | `pg4ml_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0` | {{< bg "18" "postgresql-18-pg4ml" "green" >}} {{< bg "17" "postgresql-17-pg4ml" "green" >}} {{< bg "16" "postgresql-16-pg4ml" "green" >}} {{< bg "15" "postgresql-15-pg4ml" "green" >}} {{< bg "14" "postgresql-14-pg4ml" "green" >}} | `postgresql-$v-pg4ml` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0" "pg4ml_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "pg4ml_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0" "postgresql-18-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-17-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-16-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-15-pg4ml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0" "postgresql-14-pg4ml : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg4ml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg4ml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg4ml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg4ml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg4ml : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg4ml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg4ml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg4ml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg4ml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg4ml : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
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

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitee.com/guotiecheng/plpgsql_pg4ml" title="Repository" icon="link" subtitle="gitee.com/guotiecheng/plpgsql_pg4ml" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg4ml-2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg4ml;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg4ml;		# install via package name, for the active PG version

pig install pg4ml -v 18;   # install for PG 18
pig install pg4ml -v 17;   # install for PG 17
pig install pg4ml -v 16;   # install for PG 16
pig install pg4ml -v 15;   # install for PG 15
pig install pg4ml -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg4ml CASCADE; -- requires plpgsql, tablefunc, cube, plpython3u
```



## Usage

> [pg4ml](https://gitee.com/guotiecheng/plpgsql_pg4ml): Machine learning framework for PostgreSQL.
> Source: [README.md](https://gitee.com/guotiecheng/plpgsql_pg4ml)

`pg4ml` is a PostgreSQL extension that implements a machine learning framework entirely within the database using PL/pgSQL and PL/Python. It provides matrix operations, neural network construction and training, clustering algorithms, and scientific computing -- all through SQL.


--------

### Prerequisites

- PostgreSQL >= 14 with Python3 support
- Required extensions: `plpgsql`, `tablefunc`, `cube`, `plpython3u`

### Getting Started

```sql
CREATE EXTENSION pg4ml CASCADE;
-- This will also create the required dependencies: plpgsql, tablefunc, cube, plpython3u
```


--------

### Features

#### Matrix Operations

The framework provides a comprehensive matrix operation library under the `sm_sc` schema:

- **Element-wise operations**: arithmetic, comparison, rounding, concatenation, boolean, bitwise, complex number, and broadcast operations
- **Matrix operations**: multiplication, transpose, flip, rotate, concatenation
- **Construction**: sampling, replacement, padding, character matching, random generation
- **Trigonometric functions**: broadcast operations on matrices
- **Aggregation**: slice-level aggregation, matrix-level aggregation, sorting by slice values, locating extremum positions

#### Slice Aggregation Examples

Average over vertical slices (groups of 2):

```sql
SELECT sm_sc.fv_aggr_slice_avg(
    array[[1.5, 11.5],
          [2.1, 12.1],
          [3.3, 13.3],
          [4.3, 14.3],
          [5.5, 15.5],
          [6.1, 16.1]],
    array[2, 1]
);
-- Returns: array[[1.8, 11.8],[3.8, 13.8],[5.8, 15.8]]
```

Max pooling over 2x3 blocks:

```sql
SELECT sm_sc.fv_aggr_slice_max(
    array[[2.3, 5.1, 8.2, 2.56, 3.33, -1.9],
          [3.25, 6.4, 6.6, 6.9, -2.65, -4.6],
          [-2.3, 5.1, -8.2, 2.56, -3.33, -1.9],
          [3.25, -6.4, -6.6, 6.9, -2.65, -4.6]],
    array[2, 3]
);
-- Returns: array[[8.2, 6.9],[5.1, 6.9]]
```

#### Neural Networks

The framework supports deep neural network construction and training:

- **Node and Path tables**: `sm_sc.tb_nn_node` / `sm_sc.tb_nn_path` for defining network structure
- **Training input buffer**: `sm_sc.tb_nn_train_input_buff` for receiving training data
- **Task management**: `sm_sc.tb_classify_task` for deploying and managing training tasks
- **Activation functions**, **convolution**, **pooling**, **lambda operations**
- **Loss functions**, **derivative computation**, **backpropagation**
- **Inference**: `sm_sc.ft_nn_in_out` for running test/validation data through a trained model

#### Clustering

- **K-means++**: via `sm_sc.prc_kmeans_pp` procedure
- **DBSCAN**: via `sm_sc.prc_dbscan_pp` procedure

Both use `sm_sc.tb_cluster_task` for task deployment and management.

#### Scientific Computing

- Waveform processing
- Computational graph JSON serialization/deserialization
- Complex number operations
- Linear algebra


--------

### Performance Tips

- Enable debug mode with: `SET session pg4ml._v_is_debug_check = '1';`
- Matrix multiplication uses `plpython3u` to call numpy for optimization
- Adjust PostgreSQL parallel parameters for multi-threaded training:
  - `max_parallel_workers_per_gather`
  - `force_parallel_mode`
  - `parallel_setup_cost`, `parallel_tuple_cost`
- Consider using `pg_strom` extension for GPU acceleration
