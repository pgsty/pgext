---
title: "plr"
linkTitle: "plr"
description: "load R interpreter and execute R script from within a database"
weight: 3100
categories: ["LANG"]
width: full
---

load R interpreter and execute R script from within a database


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3100** | {{< badge content="plr" link="https://github.com/postgres-plr/plr" >}} | {{< ext "plr" >}} | `8.4.8` | {{< category "LANG" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pgml" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pljava" >}} {{< ext "plperl" >}} {{< ext "pllua" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/plr" >}} | `8.4.8` | {{< bg "18" "plr_18*" "green" >}} {{< bg "17" "plr_17*" "green" >}} {{< bg "16" "plr_16*" "green" >}} {{< bg "15" "plr_15*" "green" >}} {{< bg "14" "plr_14*" "green" >}} {{< bg "13" "plr_13*" "green" >}} | `plr_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/plr" >}} | `8.4.8` | {{< bg "18" "postgresql-18-plr" "green" >}} {{< bg "17" "postgresql-17-plr" "green" >}} {{< bg "16" "postgresql-16-plr" "green" >}} {{< bg "15" "postgresql-15-plr" "green" >}} {{< bg "14" "postgresql-14-plr" "green" >}} {{< bg "13" "postgresql-13-plr" "green" >}} | `postgresql-$v-plr` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "PGDG 8.4.8" "plr_18 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_17 : HIDE 2" >}}   |  {{< bg "PGDG 8.4.8" "plr_16 : HIDE 3" >}}   |  {{< bg "PGDG 8.4.8" "plr_15 : HIDE 4" >}}   |  {{< bg "PGDG 8.4.8" "plr_14 : HIDE 5" >}}   |  {{< bg "PGDG 8.4.8" "plr_13 : HIDE 5" >}}   |
|    `el8.aarch64`    |  {{< bg "PGDG 8.4.8" "plr_18 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_17 : HIDE 2" >}}   |  {{< bg "PGDG 8.4.8" "plr_16 : HIDE 3" >}}   |  {{< bg "PGDG 8.4.8" "plr_15 : HIDE 3" >}}   |  {{< bg "PGDG 8.4.8" "plr_14 : HIDE 3" >}}   |  {{< bg "PGDG 8.4.8" "plr_13 : HIDE 3" >}}   |
|    `el9.x86_64`    |  {{< bg "PGDG 8.4.8" "plr_18 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_17 : HIDE 2" >}}   |  {{< bg "PGDG 8.4.8" "plr_16 : HIDE 6" >}}   |  {{< bg "PGDG 8.4.8" "plr_15 : HIDE 4" >}}   |  {{< bg "PGDG 8.4.8" "plr_14 : HIDE 4" >}}   |  {{< bg "PGDG 8.4.8" "plr_13 : HIDE 4" >}}   |
|    `el9.aarch64`    |  {{< bg "PGDG 8.4.8" "plr_18 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_17 : HIDE 2" >}}   |  {{< bg "PGDG 8.4.8" "plr_16 : HIDE 3" >}}   |  {{< bg "PGDG 8.4.8" "plr_15 : HIDE 3" >}}   |  {{< bg "PGDG 8.4.8" "plr_14 : HIDE 3" >}}   |  {{< bg "PGDG 8.4.7" "plr_13 : HIDE 2" >}}   |
|    `el10.x86_64`    |  {{< bg "MISS" "plr_18 : HIDE 0" "red" >}}   |  {{< bg "MISS" "plr_17 : HIDE 0" "red" >}}   |  {{< bg "MISS" "plr_16 : HIDE 0" "red" >}}   |  {{< bg "MISS" "plr_15 : HIDE 0" "red" >}}   |  {{< bg "MISS" "plr_14 : HIDE 0" "red" >}}   |  {{< bg "MISS" "plr_13 : HIDE 0" "red" >}}   |
|    `el10.aarch64`    |  {{< bg "PGDG 8.4.8" "plr_18 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_17 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_16 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_15 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_14 : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8" "plr_13 : HIDE 1" >}}   |
|    `d12.x86_64`    |  {{< bg "PGDG 8.4.8.1" "postgresql-18-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-17-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-16-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-15-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-14-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-13-plr : HIDE 1" >}}   |
|    `d12.aarch64`    |  {{< bg "PGDG 8.4.8.1" "postgresql-18-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-17-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-16-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-15-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-14-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-13-plr : HIDE 1" >}}   |
|    `d13.x86_64`    |  {{< bg "PGDG 8.4.8.1" "postgresql-18-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-17-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-16-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-15-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-14-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-13-plr : HIDE 1" >}}   |
|    `d13.aarch64`    |  {{< bg "PGDG 8.4.8.1" "postgresql-18-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-17-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-16-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-15-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-14-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-13-plr : HIDE 1" >}}   |
|    `u22.x86_64`    |  {{< bg "PGDG 8.4.8.1" "postgresql-18-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-17-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-16-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-15-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-14-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-13-plr : HIDE 1" >}}   |
|    `u22.aarch64`    |  {{< bg "PGDG 8.4.8.1" "postgresql-18-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-17-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-16-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-15-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-14-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-13-plr : HIDE 1" >}}   |
|    `u24.x86_64`    |  {{< bg "PGDG 8.4.8.1" "postgresql-18-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-17-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-16-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-15-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-14-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-13-plr : HIDE 1" >}}   |
|    `u24.aarch64`    |  {{< bg "PGDG 8.4.8.1" "postgresql-18-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-17-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-16-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-15-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-14-plr : HIDE 1" >}}   |  {{< bg "PGDG 8.4.8.1" "postgresql-13-plr : HIDE 1" >}}   |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_18` | 8.4.8 | `el8.x86_64` | pgdg | 76.1 KiB | [plr_18-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plr_18-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_18` | 8.4.8 | `el8.aarch64` | pgdg | 73.7 KiB | [plr_18-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plr_18-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_18` | 8.4.8 | `el9.x86_64` | pgdg | 73.9 KiB | [plr_18-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_18` | 8.4.8 | `el9.aarch64` | pgdg | 72.1 KiB | [plr_18-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_18` | 8.4.8 | `el10.aarch64` | pgdg | 73.7 KiB | [plr_18-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plr` | 8.4.8.1 | `d12.x86_64` | pgdg | 136.0 KiB | [postgresql-18-plr_8.4.8.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plr` | 8.4.8.1 | `d12.aarch64` | pgdg | 132.2 KiB | [postgresql-18-plr_8.4.8.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plr` | 8.4.8.1 | `d13.x86_64` | pgdg | 136.0 KiB | [postgresql-18-plr_8.4.8.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plr` | 8.4.8.1 | `d13.aarch64` | pgdg | 132.6 KiB | [postgresql-18-plr_8.4.8.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plr` | 8.4.8.1 | `u22.x86_64` | pgdg | 131.5 KiB | [postgresql-18-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plr` | 8.4.8.1 | `u22.aarch64` | pgdg | 128.2 KiB | [postgresql-18-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plr` | 8.4.8.1 | `u24.x86_64` | pgdg | 127.1 KiB | [postgresql-18-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plr` | 8.4.8.1 | `u24.aarch64` | pgdg | 123.5 KiB | [postgresql-18-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_17` | 8.4.8 | `el8.x86_64` | pgdg | 76.0 KiB | [plr_17-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_17` | 8.4.7 | `el8.x86_64` | pgdg | 75.7 KiB | [plr_17-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_17` | 8.4.8 | `el8.aarch64` | pgdg | 73.7 KiB | [plr_17-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_17` | 8.4.7 | `el8.aarch64` | pgdg | 73.4 KiB | [plr_17-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_17` | 8.4.8 | `el9.x86_64` | pgdg | 73.8 KiB | [plr_17-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_17` | 8.4.7 | `el9.x86_64` | pgdg | 73.6 KiB | [plr_17-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_17` | 8.4.8 | `el9.aarch64` | pgdg | 72.2 KiB | [plr_17-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_17` | 8.4.7 | `el9.aarch64` | pgdg | 72.0 KiB | [plr_17-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_17` | 8.4.8 | `el10.aarch64` | pgdg | 73.5 KiB | [plr_17-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-plr` | 8.4.8.1 | `d12.x86_64` | pgdg | 135.6 KiB | [postgresql-17-plr_8.4.8.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plr` | 8.4.8.1 | `d12.aarch64` | pgdg | 132.0 KiB | [postgresql-17-plr_8.4.8.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plr` | 8.4.8.1 | `d13.x86_64` | pgdg | 135.8 KiB | [postgresql-17-plr_8.4.8.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plr` | 8.4.8.1 | `d13.aarch64` | pgdg | 132.4 KiB | [postgresql-17-plr_8.4.8.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plr` | 8.4.8.1 | `u22.x86_64` | pgdg | 155.6 KiB | [postgresql-17-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plr` | 8.4.8.1 | `u22.aarch64` | pgdg | 152.1 KiB | [postgresql-17-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plr` | 8.4.8.1 | `u24.x86_64` | pgdg | 126.9 KiB | [postgresql-17-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plr` | 8.4.8.1 | `u24.aarch64` | pgdg | 123.3 KiB | [postgresql-17-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_16` | 8.4.8 | `el8.x86_64` | pgdg | 76.0 KiB | [plr_16-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | 8.4.7 | `el8.x86_64` | pgdg | 75.7 KiB | [plr_16-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | 8.4.6 | `el8.x86_64` | pgdg | 74.8 KiB | [plr_16-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | 8.4.8 | `el8.aarch64` | pgdg | 73.6 KiB | [plr_16-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | 8.4.7 | `el8.aarch64` | pgdg | 73.4 KiB | [plr_16-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | 8.4.6 | `el8.aarch64` | pgdg | 72.6 KiB | [plr_16-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | 8.4.8 | `el9.x86_64` | pgdg | 73.8 KiB | [plr_16-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | 8.4.7 | `el9.x86_64` | pgdg | 73.6 KiB | [plr_16-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | 8.4.6 | `el9.x86_64` | pgdg | 73.6 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | 8.4.6 | `el9.x86_64` | pgdg | 73.6 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | 8.4.6 | `el9.x86_64` | pgdg | 73.4 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | 8.4.6 | `el9.x86_64` | pgdg | 72.8 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | 8.4.8 | `el9.aarch64` | pgdg | 72.2 KiB | [plr_16-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | 8.4.7 | `el9.aarch64` | pgdg | 72.0 KiB | [plr_16-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | 8.4.6 | `el9.aarch64` | pgdg | 71.1 KiB | [plr_16-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | 8.4.8 | `el10.aarch64` | pgdg | 73.5 KiB | [plr_16-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-plr` | 8.4.8.1 | `d12.x86_64` | pgdg | 135.6 KiB | [postgresql-16-plr_8.4.8.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plr` | 8.4.8.1 | `d12.aarch64` | pgdg | 132.0 KiB | [postgresql-16-plr_8.4.8.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plr` | 8.4.8.1 | `d13.x86_64` | pgdg | 135.7 KiB | [postgresql-16-plr_8.4.8.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plr` | 8.4.8.1 | `d13.aarch64` | pgdg | 132.4 KiB | [postgresql-16-plr_8.4.8.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plr` | 8.4.8.1 | `u22.x86_64` | pgdg | 151.6 KiB | [postgresql-16-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plr` | 8.4.8.1 | `u22.aarch64` | pgdg | 148.2 KiB | [postgresql-16-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plr` | 8.4.8.1 | `u24.x86_64` | pgdg | 127.0 KiB | [postgresql-16-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plr` | 8.4.8.1 | `u24.aarch64` | pgdg | 123.4 KiB | [postgresql-16-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_15` | 8.4.8 | `el8.x86_64` | pgdg | 76.5 KiB | [plr_15-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | 8.4.7 | `el8.x86_64` | pgdg | 76.2 KiB | [plr_15-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | 8.4.6 | `el8.x86_64` | pgdg | 75.4 KiB | [plr_15-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | 8.4.5 | `el8.x86_64` | pgdg | 167.4 KiB | [plr_15-8.4.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.5-1.rhel8.x86_64.rpm) |
| `plr_15` | 8.4.8 | `el8.aarch64` | pgdg | 74.1 KiB | [plr_15-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | 8.4.7 | `el8.aarch64` | pgdg | 73.9 KiB | [plr_15-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | 8.4.6 | `el8.aarch64` | pgdg | 73.1 KiB | [plr_15-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | 8.4.8 | `el9.x86_64` | pgdg | 74.5 KiB | [plr_15-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | 8.4.7 | `el9.x86_64` | pgdg | 74.4 KiB | [plr_15-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | 8.4.6 | `el9.x86_64` | pgdg | 73.6 KiB | [plr_15-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | 8.4.5 | `el9.x86_64` | pgdg | 167.8 KiB | [plr_15-8.4.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.5-1.rhel9.x86_64.rpm) |
| `plr_15` | 8.4.8 | `el9.aarch64` | pgdg | 72.8 KiB | [plr_15-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | 8.4.7 | `el9.aarch64` | pgdg | 72.6 KiB | [plr_15-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | 8.4.6 | `el9.aarch64` | pgdg | 71.7 KiB | [plr_15-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | 8.4.8 | `el10.aarch64` | pgdg | 74.1 KiB | [plr_15-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-plr` | 8.4.8.1 | `d12.x86_64` | pgdg | 136.1 KiB | [postgresql-15-plr_8.4.8.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plr` | 8.4.8.1 | `d12.aarch64` | pgdg | 132.4 KiB | [postgresql-15-plr_8.4.8.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plr` | 8.4.8.1 | `d13.x86_64` | pgdg | 136.1 KiB | [postgresql-15-plr_8.4.8.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plr` | 8.4.8.1 | `d13.aarch64` | pgdg | 132.8 KiB | [postgresql-15-plr_8.4.8.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plr` | 8.4.8.1 | `u22.x86_64` | pgdg | 151.4 KiB | [postgresql-15-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plr` | 8.4.8.1 | `u22.aarch64` | pgdg | 148.1 KiB | [postgresql-15-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plr` | 8.4.8.1 | `u24.x86_64` | pgdg | 127.2 KiB | [postgresql-15-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plr` | 8.4.8.1 | `u24.aarch64` | pgdg | 123.6 KiB | [postgresql-15-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_14` | 8.4.8 | `el8.x86_64` | pgdg | 76.5 KiB | [plr_14-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | 8.4.7 | `el8.x86_64` | pgdg | 76.2 KiB | [plr_14-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | 8.4.6 | `el8.x86_64` | pgdg | 75.4 KiB | [plr_14-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | 8.4.5 | `el8.x86_64` | pgdg | 166.7 KiB | [plr_14-8.4.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.5-1.rhel8.x86_64.rpm) |
| `plr_14` | 8.4.3 | `el8.x86_64` | pgdg | 166.5 KiB | [plr_14-8.4.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.3-1.rhel8.x86_64.rpm) |
| `plr_14` | 8.4.8 | `el8.aarch64` | pgdg | 74.1 KiB | [plr_14-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | 8.4.7 | `el8.aarch64` | pgdg | 73.9 KiB | [plr_14-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | 8.4.6 | `el8.aarch64` | pgdg | 73.1 KiB | [plr_14-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | 8.4.8 | `el9.x86_64` | pgdg | 74.5 KiB | [plr_14-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | 8.4.7 | `el9.x86_64` | pgdg | 74.3 KiB | [plr_14-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | 8.4.6 | `el9.x86_64` | pgdg | 73.6 KiB | [plr_14-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | 8.4.5 | `el9.x86_64` | pgdg | 167.4 KiB | [plr_14-8.4.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.5-1.rhel9.x86_64.rpm) |
| `plr_14` | 8.4.8 | `el9.aarch64` | pgdg | 72.8 KiB | [plr_14-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | 8.4.7 | `el9.aarch64` | pgdg | 72.6 KiB | [plr_14-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | 8.4.6 | `el9.aarch64` | pgdg | 71.7 KiB | [plr_14-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | 8.4.8 | `el10.aarch64` | pgdg | 74.1 KiB | [plr_14-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-plr` | 8.4.8.1 | `d12.x86_64` | pgdg | 136.0 KiB | [postgresql-14-plr_8.4.8.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plr` | 8.4.8.1 | `d12.aarch64` | pgdg | 132.4 KiB | [postgresql-14-plr_8.4.8.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plr` | 8.4.8.1 | `d13.x86_64` | pgdg | 136.3 KiB | [postgresql-14-plr_8.4.8.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plr` | 8.4.8.1 | `d13.aarch64` | pgdg | 132.8 KiB | [postgresql-14-plr_8.4.8.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plr` | 8.4.8.1 | `u22.x86_64` | pgdg | 151.2 KiB | [postgresql-14-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plr` | 8.4.8.1 | `u22.aarch64` | pgdg | 148.3 KiB | [postgresql-14-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plr` | 8.4.8.1 | `u24.x86_64` | pgdg | 127.3 KiB | [postgresql-14-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plr` | 8.4.8.1 | `u24.aarch64` | pgdg | 123.6 KiB | [postgresql-14-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_13` | 8.4.8 | `el8.x86_64` | pgdg | 75.8 KiB | [plr_13-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plr_13-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_13` | 8.4.7 | `el8.x86_64` | pgdg | 75.6 KiB | [plr_13-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plr_13-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_13` | 8.4.6 | `el8.x86_64` | pgdg | 74.7 KiB | [plr_13-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plr_13-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_13` | 8.4.5 | `el8.x86_64` | pgdg | 165.5 KiB | [plr_13-8.4.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plr_13-8.4.5-1.rhel8.x86_64.rpm) |
| `plr_13` | 8.4.2 | `el8.x86_64` | pgdg | 164.3 KiB | [plr_13-8.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plr_13-8.4.2-1.rhel8.x86_64.rpm) |
| `plr_13` | 8.4.8 | `el8.aarch64` | pgdg | 74.1 KiB | [plr_13-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plr_13-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_13` | 8.4.7 | `el8.aarch64` | pgdg | 73.8 KiB | [plr_13-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plr_13-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_13` | 8.4.6 | `el8.aarch64` | pgdg | 73.0 KiB | [plr_13-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plr_13-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_13` | 8.4.8 | `el9.x86_64` | pgdg | 74.3 KiB | [plr_13-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plr_13-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_13` | 8.4.7 | `el9.x86_64` | pgdg | 74.2 KiB | [plr_13-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plr_13-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_13` | 8.4.6 | `el9.x86_64` | pgdg | 73.4 KiB | [plr_13-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plr_13-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_13` | 8.4.5 | `el9.x86_64` | pgdg | 166.6 KiB | [plr_13-8.4.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plr_13-8.4.5-1.rhel9.x86_64.rpm) |
| `plr_13` | 8.4.7 | `el9.aarch64` | pgdg | 72.6 KiB | [plr_13-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plr_13-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_13` | 8.4.6 | `el9.aarch64` | pgdg | 71.6 KiB | [plr_13-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plr_13-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_13` | 8.4.8 | `el10.aarch64` | pgdg | 74.0 KiB | [plr_13-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/plr_13-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-plr` | 8.4.8.1 | `d12.x86_64` | pgdg | 135.4 KiB | [postgresql-13-plr_8.4.8.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-13-plr_8.4.8.1-1.pgdg12+1_amd64.deb) |
| `postgresql-13-plr` | 8.4.8.1 | `d12.aarch64` | pgdg | 132.0 KiB | [postgresql-13-plr_8.4.8.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-13-plr_8.4.8.1-1.pgdg12+1_arm64.deb) |
| `postgresql-13-plr` | 8.4.8.1 | `d13.x86_64` | pgdg | 136.0 KiB | [postgresql-13-plr_8.4.8.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-13-plr_8.4.8.1-1.pgdg13+1_amd64.deb) |
| `postgresql-13-plr` | 8.4.8.1 | `d13.aarch64` | pgdg | 132.4 KiB | [postgresql-13-plr_8.4.8.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-13-plr_8.4.8.1-1.pgdg13+1_arm64.deb) |
| `postgresql-13-plr` | 8.4.8.1 | `u22.x86_64` | pgdg | 150.5 KiB | [postgresql-13-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-13-plr_8.4.8.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-plr` | 8.4.8.1 | `u22.aarch64` | pgdg | 147.0 KiB | [postgresql-13-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-13-plr_8.4.8.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-plr` | 8.4.8.1 | `u24.x86_64` | pgdg | 126.8 KiB | [postgresql-13-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-13-plr_8.4.8.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-plr` | 8.4.8.1 | `u24.aarch64` | pgdg | 123.2 KiB | [postgresql-13-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-13-plr_8.4.8.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgres-plr/plr" title="Repository" icon="github" subtitle="github.com/postgres-plr/plr" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install plr; # install by extension name, for the current active PG version
pig ext install plr; # install via package alias, for the active PG version
pig ext install plr -v 18;   # install for PG 18
pig ext install plr -v 17;   # install for PG 17
pig ext install plr -v 16;   # install for PG 16
pig ext install plr -v 15;   # install for PG 15
pig ext install plr -v 14;   # install for PG 14
pig ext install plr -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION plr;
```

