---
title: "pljs"
linkTitle: "pljs"
description: "PL/JS trusted procedural language"
weight: 3011
categories: ["LANG"]
width: full
---

[**pljs**](https://github.com/plv8/pljs) : PL/JS trusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3011** | {{< badge content="pljs" link="https://github.com/plv8/pljs" >}} | {{< ext "pljs" >}} | `1.0.5` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "plv8" >}} {{< ext "jsquery" >}} {{< ext "pllua" >}} {{< ext "pg_tle" >}} {{< ext "plpgsql" >}} {{< ext "pg_jsonschema" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} |

> [!Note] with submodules, hot fix with CONFIG_VERSION


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.0.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pljs` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.5` | {{< bg "18" "pljs_18" "green" >}} {{< bg "17" "pljs_17" "green" >}} {{< bg "16" "pljs_16" "green" >}} {{< bg "15" "pljs_15" "green" >}} {{< bg "14" "pljs_14" "green" >}} {{< bg "13" "pljs_13" "red" >}} | `pljs_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.5` | {{< bg "18" "postgresql-18-pljs" "green" >}} {{< bg "17" "postgresql-17-pljs" "green" >}} {{< bg "16" "postgresql-16-pljs" "green" >}} {{< bg "15" "postgresql-15-pljs" "green" >}} {{< bg "14" "postgresql-14-pljs" "green" >}} {{< bg "13" "postgresql-13-pljs" "red" >}} | `postgresql-$v-pljs` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_18` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 381.8 KiB | [pljs_18-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_18-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pljs_18` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.1 KiB | [pljs_18-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_18-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pljs_18` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 388.1 KiB | [pljs_18-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_18-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pljs_18` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 370.4 KiB | [pljs_18-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_18-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pljs_18` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 413.1 KiB | [pljs_18-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_18-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pljs_18` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 380.1 KiB | [pljs_18-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_18-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.1 KiB | [postgresql-18-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 374.9 KiB | [postgresql-18-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.9 KiB | [postgresql-18-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.6 KiB | [postgresql-18-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 406.8 KiB | [postgresql-18-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 372.9 KiB | [postgresql-18-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.7 KiB | [postgresql-18-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 375.4 KiB | [postgresql-18-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_17` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 381.7 KiB | [pljs_17-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_17-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pljs_17` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.1 KiB | [pljs_17-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_17-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pljs_17` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 388.0 KiB | [pljs_17-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_17-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pljs_17` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 370.5 KiB | [pljs_17-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_17-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pljs_17` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 413.1 KiB | [pljs_17-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_17-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pljs_17` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 379.8 KiB | [pljs_17-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_17-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.0 KiB | [postgresql-17-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.1 KiB | [postgresql-17-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 428.8 KiB | [postgresql-17-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.6 KiB | [postgresql-17-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 423.1 KiB | [postgresql-17-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.6 KiB | [postgresql-17-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.3 KiB | [postgresql-17-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 375.9 KiB | [postgresql-17-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_16` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 381.7 KiB | [pljs_16-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_16-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pljs_16` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.1 KiB | [pljs_16-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_16-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pljs_16` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 388.4 KiB | [pljs_16-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_16-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pljs_16` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 370.3 KiB | [pljs_16-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_16-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pljs_16` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 413.4 KiB | [pljs_16-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_16-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pljs_16` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 379.3 KiB | [pljs_16-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_16-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.0 KiB | [postgresql-16-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.0 KiB | [postgresql-16-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 428.9 KiB | [postgresql-16-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.8 KiB | [postgresql-16-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 421.6 KiB | [postgresql-16-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.3 KiB | [postgresql-16-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.6 KiB | [postgresql-16-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.0 KiB | [postgresql-16-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_15` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 382.4 KiB | [pljs_15-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_15-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pljs_15` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.4 KiB | [pljs_15-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_15-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pljs_15` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 388.5 KiB | [pljs_15-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_15-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pljs_15` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 372.3 KiB | [pljs_15-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_15-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pljs_15` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 413.6 KiB | [pljs_15-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_15-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pljs_15` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 394.6 KiB | [pljs_15-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_15-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.1 KiB | [postgresql-15-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.7 KiB | [postgresql-15-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.5 KiB | [postgresql-15-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 382.8 KiB | [postgresql-15-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 421.6 KiB | [postgresql-15-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.9 KiB | [postgresql-15-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.5 KiB | [postgresql-15-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.5 KiB | [postgresql-15-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_14` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 382.3 KiB | [pljs_14-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_14-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `pljs_14` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.5 KiB | [pljs_14-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_14-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `pljs_14` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 389.1 KiB | [pljs_14-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_14-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `pljs_14` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 372.3 KiB | [pljs_14-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_14-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `pljs_14` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 421.8 KiB | [pljs_14-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_14-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `pljs_14` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 394.4 KiB | [pljs_14-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_14-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.1 KiB | [postgresql-14-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.3 KiB | [postgresql-14-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.3 KiB | [postgresql-14-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.5 KiB | [postgresql-14-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 422.0 KiB | [postgresql-14-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.1 KiB | [postgresql-14-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.2 KiB | [postgresql-14-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.4 KiB | [postgresql-14-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/plv8/pljs" title="Repository" icon="github" subtitle="github.com/plv8/pljs" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pljs-1.0.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pljs;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pljs;		# install via package name, for the active PG version

pig install pljs -v 18;   # install for PG 18
pig install pljs -v 17;   # install for PG 17
pig install pljs -v 16;   # install for PG 16
pig install pljs -v 15;   # install for PG 15
pig install pljs -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pljs;
```
