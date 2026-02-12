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
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.5` | {{< bg "18" "postgresql-18-pljs" "green" >}} {{< bg "17" "postgresql-17-pljs" "green" >}} {{< bg "16" "postgresql-16-pljs" "green" >}} {{< bg "15" "postgresql-15-pljs" "green" >}} {{< bg "14" "postgresql-14-pljs" "green" >}} {{< bg "13" "postgresql-13-pljs" "red" >}} | `postgresql-$v-pljs` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.5" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.5" "pljs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.5" "pljs_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pljs_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pljs : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-13-pljs : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pljs : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-13-pljs : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pljs : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-13-pljs : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pljs : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-13-pljs : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pljs : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-13-pljs : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pljs : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-13-pljs : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pljs : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-13-pljs : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0.5" "postgresql-18-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-17-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-16-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-15-pljs : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.5" "postgresql-14-pljs : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.5" "postgresql-13-pljs : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_18` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 381.9 KiB | [pljs_18-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_18-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pljs_18` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.2 KiB | [pljs_18-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_18-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pljs_18` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 388.2 KiB | [pljs_18-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_18-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pljs_18` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 370.5 KiB | [pljs_18-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_18-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pljs_18` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 413.8 KiB | [pljs_18-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_18-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pljs_18` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 380.1 KiB | [pljs_18-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_18-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.1 KiB | [postgresql-18-pljs_1.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 409.1 KiB | [postgresql-18-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.3 KiB | [postgresql-18-pljs_1.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 374.6 KiB | [postgresql-18-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.3 KiB | [postgresql-18-pljs_1.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 428.3 KiB | [postgresql-18-pljs_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.8 KiB | [postgresql-18-pljs_1.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 380.5 KiB | [postgresql-18-pljs_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 407.3 KiB | [postgresql-18-pljs_1.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 434.7 KiB | [postgresql-18-pljs_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 373.2 KiB | [postgresql-18-pljs_1.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 424.5 KiB | [postgresql-18-pljs_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.6 KiB | [postgresql-18-pljs_1.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 432.1 KiB | [postgresql-18-pljs_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.9 KiB | [postgresql-18-pljs_1.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 422.6 KiB | [postgresql-18-pljs_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-18-pljs_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_17` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 381.7 KiB | [pljs_17-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_17-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pljs_17` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.3 KiB | [pljs_17-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_17-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pljs_17` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 388.1 KiB | [pljs_17-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_17-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pljs_17` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 370.4 KiB | [pljs_17-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_17-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pljs_17` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 413.6 KiB | [pljs_17-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_17-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pljs_17` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 379.4 KiB | [pljs_17-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_17-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.0 KiB | [postgresql-17-pljs_1.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 408.8 KiB | [postgresql-17-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.2 KiB | [postgresql-17-pljs_1.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 373.9 KiB | [postgresql-17-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.3 KiB | [postgresql-17-pljs_1.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 427.8 KiB | [postgresql-17-pljs_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.7 KiB | [postgresql-17-pljs_1.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 380.7 KiB | [postgresql-17-pljs_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 422.9 KiB | [postgresql-17-pljs_1.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 450.8 KiB | [postgresql-17-pljs_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 389.6 KiB | [postgresql-17-pljs_1.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 440.0 KiB | [postgresql-17-pljs_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 407.9 KiB | [postgresql-17-pljs_1.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 431.6 KiB | [postgresql-17-pljs_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 375.8 KiB | [postgresql-17-pljs_1.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 422.4 KiB | [postgresql-17-pljs_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-17-pljs_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_16` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 381.8 KiB | [pljs_16-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_16-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pljs_16` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.2 KiB | [pljs_16-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_16-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pljs_16` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 388.0 KiB | [pljs_16-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_16-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pljs_16` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 370.6 KiB | [pljs_16-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_16-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pljs_16` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 413.5 KiB | [pljs_16-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_16-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pljs_16` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 379.5 KiB | [pljs_16-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_16-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.0 KiB | [postgresql-16-pljs_1.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 408.8 KiB | [postgresql-16-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.2 KiB | [postgresql-16-pljs_1.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 374.9 KiB | [postgresql-16-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.1 KiB | [postgresql-16-pljs_1.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 427.8 KiB | [postgresql-16-pljs_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.8 KiB | [postgresql-16-pljs_1.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 380.8 KiB | [postgresql-16-pljs_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 421.7 KiB | [postgresql-16-pljs_1.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 449.2 KiB | [postgresql-16-pljs_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.7 KiB | [postgresql-16-pljs_1.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 439.1 KiB | [postgresql-16-pljs_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.0 KiB | [postgresql-16-pljs_1.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 431.7 KiB | [postgresql-16-pljs_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 375.9 KiB | [postgresql-16-pljs_1.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 422.5 KiB | [postgresql-16-pljs_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-16-pljs_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_15` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 382.4 KiB | [pljs_15-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_15-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pljs_15` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.6 KiB | [pljs_15-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_15-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pljs_15` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 389.2 KiB | [pljs_15-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_15-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pljs_15` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 372.4 KiB | [pljs_15-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_15-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pljs_15` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 414.1 KiB | [pljs_15-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_15-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pljs_15` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 381.2 KiB | [pljs_15-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_15-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 409.9 KiB | [postgresql-15-pljs_1.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 408.6 KiB | [postgresql-15-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.5 KiB | [postgresql-15-pljs_1.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 374.3 KiB | [postgresql-15-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.4 KiB | [postgresql-15-pljs_1.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 428.6 KiB | [postgresql-15-pljs_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.9 KiB | [postgresql-15-pljs_1.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 380.8 KiB | [postgresql-15-pljs_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 421.9 KiB | [postgresql-15-pljs_1.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 449.6 KiB | [postgresql-15-pljs_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.0 KiB | [postgresql-15-pljs_1.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 439.3 KiB | [postgresql-15-pljs_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.1 KiB | [postgresql-15-pljs_1.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 432.0 KiB | [postgresql-15-pljs_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.5 KiB | [postgresql-15-pljs_1.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 423.0 KiB | [postgresql-15-pljs_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-15-pljs_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pljs_14` | `1.0.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 382.4 KiB | [pljs_14-1.0.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pljs_14-1.0.5-1PIGSTY.el8.x86_64.rpm) |
| `pljs_14` | `1.0.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 349.6 KiB | [pljs_14-1.0.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pljs_14-1.0.5-1PIGSTY.el8.aarch64.rpm) |
| `pljs_14` | `1.0.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 388.8 KiB | [pljs_14-1.0.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pljs_14-1.0.5-1PIGSTY.el9.x86_64.rpm) |
| `pljs_14` | `1.0.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 372.2 KiB | [pljs_14-1.0.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pljs_14-1.0.5-1PIGSTY.el9.aarch64.rpm) |
| `pljs_14` | `1.0.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 414.2 KiB | [pljs_14-1.0.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pljs_14-1.0.5-1PIGSTY.el10.x86_64.rpm) |
| `pljs_14` | `1.0.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 381.1 KiB | [pljs_14-1.0.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pljs_14-1.0.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.2 KiB | [postgresql-14-pljs_1.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 409.0 KiB | [postgresql-14-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.5 KiB | [postgresql-14-pljs_1.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 374.2 KiB | [postgresql-14-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.4 KiB | [postgresql-14-pljs_1.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 428.7 KiB | [postgresql-14-pljs_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.7 KiB | [postgresql-14-pljs_1.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 381.6 KiB | [postgresql-14-pljs_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 422.1 KiB | [postgresql-14-pljs_1.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 450.1 KiB | [postgresql-14-pljs_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.1 KiB | [postgresql-14-pljs_1.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 439.6 KiB | [postgresql-14-pljs_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.2 KiB | [postgresql-14-pljs_1.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 432.2 KiB | [postgresql-14-pljs_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.6 KiB | [postgresql-14-pljs_1.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 423.0 KiB | [postgresql-14-pljs_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-14-pljs_1.0.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-pljs` | `1.0.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 409.2 KiB | [postgresql-13-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-13-pljs_1.0.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pljs` | `1.0.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 373.8 KiB | [postgresql-13-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pljs/postgresql-13-pljs_1.0.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pljs` | `1.0.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 427.9 KiB | [postgresql-13-pljs_1.0.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-13-pljs_1.0.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pljs` | `1.0.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 380.3 KiB | [postgresql-13-pljs_1.0.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pljs/postgresql-13-pljs_1.0.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pljs` | `1.0.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 448.8 KiB | [postgresql-13-pljs_1.0.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-13-pljs_1.0.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pljs` | `1.0.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 438.7 KiB | [postgresql-13-pljs_1.0.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pljs/postgresql-13-pljs_1.0.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pljs` | `1.0.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 432.0 KiB | [postgresql-13-pljs_1.0.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-13-pljs_1.0.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pljs` | `1.0.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 422.5 KiB | [postgresql-13-pljs_1.0.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pljs/postgresql-13-pljs_1.0.5-1PIGSTY~noble_arm64.deb) |

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
