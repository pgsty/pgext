---
title: "pllua"
linkTitle: "pllua"
description: "Lua as a procedural language"
weight: 3020
categories: ["LANG"]
width: full
---

Lua as a procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3020** | {{< badge content="pllua" link="https://github.com/pllua/pllua" >}} | {{< ext "pllua" >}} | `2.0.12` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "hstore_pllua" >}} |
|   **See Also**    | {{< ext "plperl" >}} {{< ext "plpgsql" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pljava" >}} {{< ext "plperlu" >}} |
|    **Siblings**   | {{< ext "hstore_pllua" >}} {{< ext "plluau" >}} {{< ext "hstore_plluau" >}} |

> [!Note] missing pg12-15 on el.aarch64


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pllua" >}} | `2.0.12` | {{< bg "18" "pllua_18*" "red" >}} {{< bg "17" "pllua_17*" "green" >}} {{< bg "16" "pllua_16*" "green" >}} {{< bg "15" "pllua_15*" "green" >}} {{< bg "14" "pllua_14*" "green" >}} {{< bg "13" "pllua_13*" "green" >}} | `pllua_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pllua" >}} | `2.0.12` | {{< bg "18" "postgresql-18-pllua" "green" >}} {{< bg "17" "postgresql-17-pllua" "green" >}} {{< bg "16" "postgresql-16-pllua" "green" >}} {{< bg "15" "postgresql-15-pllua" "green" >}} {{< bg "14" "postgresql-14-pllua" "green" >}} {{< bg "13" "postgresql-13-pllua" "green" >}} | `postgresql-$v-pllua` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.9" "pllua_13 : AVAIL 3" "blue" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pllua_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_13 : AVAIL 1" "blue" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pllua_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_13 : AVAIL 1" "blue" >}} |
|    `el10.aarch64`    |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_13 : AVAIL 1" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 347.1 KiB | [postgresql-18-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 336.4 KiB | [postgresql-18-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `d13.x86_64` | pgdg | 347.8 KiB | [postgresql-18-pllua_2.0.12-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg13+1_amd64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `d13.aarch64` | pgdg | 336.4 KiB | [postgresql-18-pllua_2.0.12-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg13+1_arm64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 354.6 KiB | [postgresql-18-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 342.1 KiB | [postgresql-18-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 347.2 KiB | [postgresql-18-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 335.3 KiB | [postgresql-18-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_17` | 2.0.12 | `el8.x86_64` | pgdg | 119.4 KiB | [pllua_17-2.0.12-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pllua_17-2.0.12-3PGDG.rhel8.x86_64.rpm) |
| `pllua_17` | 2.0.12 | `el8.aarch64` | pgdg | 110.9 KiB | [pllua_17-2.0.12-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pllua_17-2.0.12-3PGDG.rhel8.aarch64.rpm) |
| `pllua_17` | 2.0.12 | `el9.x86_64` | pgdg | 120.4 KiB | [pllua_17-2.0.12-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pllua_17-2.0.12-3PGDG.rhel9.x86_64.rpm) |
| `pllua_17` | 2.0.12 | `el9.aarch64` | pgdg | 115.3 KiB | [pllua_17-2.0.12-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pllua_17-2.0.12-3PGDG.rhel9.aarch64.rpm) |
| `pllua_17` | 2.0.12 | `el10.x86_64` | pgdg | 122.4 KiB | [pllua_17-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pllua_17-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_17` | 2.0.12 | `el10.aarch64` | pgdg | 117.8 KiB | [pllua_17-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pllua_17-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 346.9 KiB | [postgresql-17-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 335.7 KiB | [postgresql-17-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `d13.x86_64` | pgdg | 347.4 KiB | [postgresql-17-pllua_2.0.12-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg13+1_amd64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `d13.aarch64` | pgdg | 335.9 KiB | [postgresql-17-pllua_2.0.12-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg13+1_arm64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 391.2 KiB | [postgresql-17-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 378.7 KiB | [postgresql-17-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 347.0 KiB | [postgresql-17-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 335.3 KiB | [postgresql-17-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_16` | 2.0.12 | `el8.x86_64` | pgdg | 119.2 KiB | [pllua_16-2.0.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pllua_16-2.0.12-1PGDG.rhel8.x86_64.rpm) |
| `pllua_16` | 2.0.12 | `el8.aarch64` | pgdg | 110.6 KiB | [pllua_16-2.0.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pllua_16-2.0.12-1PGDG.rhel8.aarch64.rpm) |
| `pllua_16` | 2.0.12 | `el9.x86_64` | pgdg | 120.1 KiB | [pllua_16-2.0.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pllua_16-2.0.12-1PGDG.rhel9.x86_64.rpm) |
| `pllua_16` | 2.0.12 | `el9.aarch64` | pgdg | 115.5 KiB | [pllua_16-2.0.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pllua_16-2.0.12-1PGDG.rhel9.aarch64.rpm) |
| `pllua_16` | 2.0.12 | `el10.x86_64` | pgdg | 122.7 KiB | [pllua_16-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pllua_16-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_16` | 2.0.12 | `el10.aarch64` | pgdg | 117.7 KiB | [pllua_16-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pllua_16-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 347.3 KiB | [postgresql-16-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 335.9 KiB | [postgresql-16-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `d13.x86_64` | pgdg | 347.5 KiB | [postgresql-16-pllua_2.0.12-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg13+1_amd64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `d13.aarch64` | pgdg | 335.9 KiB | [postgresql-16-pllua_2.0.12-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg13+1_arm64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 389.6 KiB | [postgresql-16-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 377.1 KiB | [postgresql-16-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 346.9 KiB | [postgresql-16-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 335.1 KiB | [postgresql-16-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_15` | 2.0.11 | `el8.x86_64` | pgdg | 120.9 KiB | [pllua_15-2.0.11-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pllua_15-2.0.11-1.rhel8.x86_64.rpm) |
| `pllua_15` | 2.0.10 | `el8.x86_64` | pgdg | 120.9 KiB | [pllua_15-2.0.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pllua_15-2.0.10-1.rhel8.x86_64.rpm) |
| `pllua_15` | 2.0.11 | `el9.x86_64` | pgdg | 123.5 KiB | [pllua_15-2.0.11-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pllua_15-2.0.11-1.rhel9.x86_64.rpm) |
| `pllua_15` | 2.0.10 | `el9.x86_64` | pgdg | 123.8 KiB | [pllua_15-2.0.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pllua_15-2.0.10-1.rhel9.x86_64.rpm) |
| `pllua_15` | 2.0.12 | `el10.x86_64` | pgdg | 125.8 KiB | [pllua_15-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pllua_15-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_15` | 2.0.12 | `el10.aarch64` | pgdg | 120.4 KiB | [pllua_15-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pllua_15-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 349.0 KiB | [postgresql-15-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 336.6 KiB | [postgresql-15-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `d13.x86_64` | pgdg | 349.5 KiB | [postgresql-15-pllua_2.0.12-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg13+1_amd64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `d13.aarch64` | pgdg | 337.3 KiB | [postgresql-15-pllua_2.0.12-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg13+1_arm64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 392.6 KiB | [postgresql-15-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 379.9 KiB | [postgresql-15-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 348.6 KiB | [postgresql-15-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 337.6 KiB | [postgresql-15-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_14` | 2.0.11 | `el8.x86_64` | pgdg | 121.1 KiB | [pllua_14-2.0.11-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pllua_14-2.0.11-1.rhel8.x86_64.rpm) |
| `pllua_14` | 2.0.10 | `el8.x86_64` | pgdg | 120.9 KiB | [pllua_14-2.0.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pllua_14-2.0.10-1.rhel8.x86_64.rpm) |
| `pllua_14` | 2.0.11 | `el9.x86_64` | pgdg | 123.2 KiB | [pllua_14-2.0.11-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pllua_14-2.0.11-1.rhel9.x86_64.rpm) |
| `pllua_14` | 2.0.12 | `el10.x86_64` | pgdg | 125.8 KiB | [pllua_14-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pllua_14-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_14` | 2.0.12 | `el10.aarch64` | pgdg | 120.6 KiB | [pllua_14-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pllua_14-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 349.0 KiB | [postgresql-14-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 336.7 KiB | [postgresql-14-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `d13.x86_64` | pgdg | 349.2 KiB | [postgresql-14-pllua_2.0.12-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg13+1_amd64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `d13.aarch64` | pgdg | 336.8 KiB | [postgresql-14-pllua_2.0.12-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg13+1_arm64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 388.0 KiB | [postgresql-14-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 375.9 KiB | [postgresql-14-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 348.5 KiB | [postgresql-14-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 336.9 KiB | [postgresql-14-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_13` | 2.0.9 | `el8.x86_64` | pgdg | 120.1 KiB | [pllua_13-2.0.9-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pllua_13-2.0.9-1.rhel8.x86_64.rpm) |
| `pllua_13` | 2.0.11 | `el8.x86_64` | pgdg | 120.4 KiB | [pllua_13-2.0.11-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pllua_13-2.0.11-1.rhel8.x86_64.rpm) |
| `pllua_13` | 2.0.10 | `el8.x86_64` | pgdg | 120.2 KiB | [pllua_13-2.0.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pllua_13-2.0.10-1.rhel8.x86_64.rpm) |
| `pllua_13` | 2.0.11 | `el9.x86_64` | pgdg | 123.2 KiB | [pllua_13-2.0.11-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pllua_13-2.0.11-1.rhel9.x86_64.rpm) |
| `pllua_13` | 2.0.12 | `el10.x86_64` | pgdg | 125.6 KiB | [pllua_13-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pllua_13-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_13` | 2.0.12 | `el10.aarch64` | pgdg | 120.6 KiB | [pllua_13-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pllua_13-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pllua` | 2.0.12 | `d12.x86_64` | pgdg | 347.8 KiB | [postgresql-13-pllua_2.0.12-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg12+1_amd64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `d12.aarch64` | pgdg | 335.8 KiB | [postgresql-13-pllua_2.0.12-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg12+1_arm64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `d13.x86_64` | pgdg | 348.4 KiB | [postgresql-13-pllua_2.0.12-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg13+1_amd64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `d13.aarch64` | pgdg | 336.5 KiB | [postgresql-13-pllua_2.0.12-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg13+1_arm64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `u22.x86_64` | pgdg | 386.8 KiB | [postgresql-13-pllua_2.0.12-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `u22.aarch64` | pgdg | 374.4 KiB | [postgresql-13-pllua_2.0.12-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `u24.x86_64` | pgdg | 348.0 KiB | [postgresql-13-pllua_2.0.12-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pllua` | 2.0.12 | `u24.aarch64` | pgdg | 336.5 KiB | [postgresql-13-pllua_2.0.12-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-13-pllua_2.0.12-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pllua/pllua" title="Repository" icon="github" subtitle="github.com/pllua/pllua" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pllua; # install by extension name, for the current active PG version
pig ext install pllua; # install via package alias, for the active PG version
pig ext install pllua -v 18;   # install for PG 18
pig ext install pllua -v 17;   # install for PG 17
pig ext install pllua -v 16;   # install for PG 16
pig ext install pllua -v 15;   # install for PG 15
pig ext install pllua -v 14;   # install for PG 14
pig ext install pllua -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pllua CASCADE SCHEMA pg_catalog;
```

