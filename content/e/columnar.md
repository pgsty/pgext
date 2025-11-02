---
title: "columnar"
linkTitle: "columnar"
description: "Hydra Columnar extension"
weight: 2410
categories: ["OLAP"]
width: full
---

Hydra Columnar extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2410** | {{< badge content="columnar" link="https://github.com/hydradatabase/hydra" >}} | {{< ext "columnar" "hydra" >}} | `1.1.2` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "timeseries" >}} |
|   **See Also**    | {{< ext "citus" >}} {{< ext "citus_columnar" >}} {{< ext "pg_mooncake" >}} {{< ext "timescaledb" >}} {{< ext "pg_analytics" >}} {{< ext "pg_parquet" >}} {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} |

> [!Note] conflict with citus columnar, no pg17


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/columnar" >}} | `1.1.2` | {{< bg "18" "hydra_18*" "red" >}} {{< bg "17" "hydra_17*" "red" >}} {{< bg "16" "hydra_16*" "green" >}} {{< bg "15" "hydra_15*" "green" >}} {{< bg "14" "hydra_14*" "green" >}} {{< bg "13" "hydra_13*" "green" >}} | `hydra_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/columnar" >}} | `1.1.2` | {{< bg "18" "postgresql-18-hydra" "red" >}} {{< bg "17" "postgresql-17-hydra" "red" >}} {{< bg "16" "postgresql-16-hydra" "green" >}} {{< bg "15" "postgresql-15-hydra" "green" >}} {{< bg "14" "postgresql-14-hydra" "green" >}} {{< bg "13" "postgresql-13-hydra" "green" >}} | `postgresql-$v-hydra` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |  {{< bg "MISS" "hydra_18 : HIDE 0" "red" >}}   |  {{< bg "MISS" "hydra_17 : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_14 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_13 : HIDE 1" >}}   |
| {{< os "el8.aarch64" >}} |  {{< bg "MISS" "hydra_18 : HIDE 0" "red" >}}   |  {{< bg "MISS" "hydra_17 : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_14 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_13 : HIDE 1" >}}   |
| {{< os "el9.x86_64" >}} |  {{< bg "MISS" "hydra_18 : HIDE 0" "red" >}}   |  {{< bg "MISS" "hydra_17 : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_14 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_13 : HIDE 1" >}}   |
| {{< os "el9.aarch64" >}} |  {{< bg "MISS" "hydra_18 : HIDE 0" "red" >}}   |  {{< bg "MISS" "hydra_17 : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_14 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_13 : HIDE 1" >}}   |
| {{< os "el10.x86_64" >}} |  {{< bg "MISS" "hydra_18 : HIDE 0" "red" >}}   |  {{< bg "MISS" "hydra_17 : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_14 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_13 : HIDE 1" >}}   |
| {{< os "el10.aarch64" >}} |  {{< bg "MISS" "hydra_18 : HIDE 0" "red" >}}   |  {{< bg "MISS" "hydra_17 : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_16 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_15 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_14 : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "hydra_13 : HIDE 1" >}}   |
| {{< os "d12.x86_64" >}} |  {{< bg "MISS" "postgresql-18-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-hydra : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-13-hydra : HIDE 1" >}}   |
| {{< os "d12.aarch64" >}} |  {{< bg "MISS" "postgresql-18-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-hydra : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-13-hydra : HIDE 1" >}}   |
| {{< os "d13.x86_64" >}} |  {{< bg "MISS" "postgresql-18-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-16-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-15-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-14-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-13-hydra : HIDE 0" "red" >}}   |
| {{< os "d13.aarch64" >}} |  {{< bg "MISS" "postgresql-18-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-16-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-15-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-14-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-13-hydra : HIDE 0" "red" >}}   |
| {{< os "u22.x86_64" >}} |  {{< bg "MISS" "postgresql-18-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-hydra : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-13-hydra : HIDE 1" >}}   |
| {{< os "u22.aarch64" >}} |  {{< bg "MISS" "postgresql-18-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-hydra : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-13-hydra : HIDE 1" >}}   |
| {{< os "u24.x86_64" >}} |  {{< bg "MISS" "postgresql-18-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-hydra : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-13-hydra : HIDE 1" >}}   |
| {{< os "u24.aarch64" >}} |  {{< bg "MISS" "postgresql-18-hydra : HIDE 0" "red" >}}   |  {{< bg "MISS" "postgresql-17-hydra : HIDE 0" "red" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : HIDE 1" >}}   |  {{< bg "PIGSTY 1.1.2" "postgresql-13-hydra : HIDE 1" >}}   |


{{< tabs items="PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hydra_16` | 1.1.2 | `el8.x86_64` | pigsty | 143.1 KiB | [hydra_16-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hydra_16-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `hydra_16` | 1.1.2 | `el8.aarch64` | pigsty | 136.7 KiB | [hydra_16-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hydra_16-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `hydra_16` | 1.1.2 | `el9.x86_64` | pigsty | 116.2 KiB | [hydra_16-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hydra_16-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `hydra_16` | 1.1.2 | `el9.aarch64` | pigsty | 112.5 KiB | [hydra_16-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hydra_16-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `hydra_16` | 1.1.2 | `el10.x86_64` | pigsty | 118.3 KiB | [hydra_16-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hydra_16-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `hydra_16` | 1.1.2 | `el10.aarch64` | pigsty | 113.9 KiB | [hydra_16-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hydra_16-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-hydra` | 1.1.2 | `d12.x86_64` | pigsty | 419.1 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-hydra` | 1.1.2 | `d12.aarch64` | pigsty | 409.0 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-hydra` | 1.1.2 | `u22.x86_64` | pigsty | 430.5 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-hydra` | 1.1.2 | `u22.aarch64` | pigsty | 425.2 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-hydra` | 1.1.2 | `u24.x86_64` | pigsty | 367.7 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-hydra` | 1.1.2 | `u24.aarch64` | pigsty | 363.6 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hydra_15` | 1.1.2 | `el8.x86_64` | pigsty | 146.7 KiB | [hydra_15-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hydra_15-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `hydra_15` | 1.1.2 | `el8.aarch64` | pigsty | 140.2 KiB | [hydra_15-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hydra_15-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `hydra_15` | 1.1.2 | `el9.x86_64` | pigsty | 136.8 KiB | [hydra_15-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hydra_15-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `hydra_15` | 1.1.2 | `el9.aarch64` | pigsty | 131.2 KiB | [hydra_15-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hydra_15-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `hydra_15` | 1.1.2 | `el10.x86_64` | pigsty | 138.1 KiB | [hydra_15-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hydra_15-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `hydra_15` | 1.1.2 | `el10.aarch64` | pigsty | 132.3 KiB | [hydra_15-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hydra_15-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-hydra` | 1.1.2 | `d12.x86_64` | pigsty | 423.6 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-hydra` | 1.1.2 | `d12.aarch64` | pigsty | 412.7 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-hydra` | 1.1.2 | `u22.x86_64` | pigsty | 449.6 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-hydra` | 1.1.2 | `u22.aarch64` | pigsty | 443.0 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-hydra` | 1.1.2 | `u24.x86_64` | pigsty | 385.7 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-hydra` | 1.1.2 | `u24.aarch64` | pigsty | 381.2 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hydra_14` | 1.1.2 | `el8.x86_64` | pigsty | 146.8 KiB | [hydra_14-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hydra_14-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `hydra_14` | 1.1.2 | `el8.aarch64` | pigsty | 140.2 KiB | [hydra_14-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hydra_14-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `hydra_14` | 1.1.2 | `el9.x86_64` | pigsty | 137.1 KiB | [hydra_14-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hydra_14-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `hydra_14` | 1.1.2 | `el9.aarch64` | pigsty | 131.5 KiB | [hydra_14-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hydra_14-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `hydra_14` | 1.1.2 | `el10.x86_64` | pigsty | 138.2 KiB | [hydra_14-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hydra_14-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `hydra_14` | 1.1.2 | `el10.aarch64` | pigsty | 133.0 KiB | [hydra_14-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hydra_14-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-hydra` | 1.1.2 | `d12.x86_64` | pigsty | 425.3 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-hydra` | 1.1.2 | `d12.aarch64` | pigsty | 414.1 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-hydra` | 1.1.2 | `u22.x86_64` | pigsty | 451.5 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-hydra` | 1.1.2 | `u22.aarch64` | pigsty | 444.8 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-hydra` | 1.1.2 | `u24.x86_64` | pigsty | 386.9 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-hydra` | 1.1.2 | `u24.aarch64` | pigsty | 382.5 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hydra_13` | 1.1.2 | `el8.x86_64` | pigsty | 127.1 KiB | [hydra_13-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hydra_13-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `hydra_13` | 1.1.2 | `el8.aarch64` | pigsty | 123.5 KiB | [hydra_13-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hydra_13-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `hydra_13` | 1.1.2 | `el9.x86_64` | pigsty | 119.5 KiB | [hydra_13-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hydra_13-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `hydra_13` | 1.1.2 | `el9.aarch64` | pigsty | 115.0 KiB | [hydra_13-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hydra_13-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `hydra_13` | 1.1.2 | `el10.x86_64` | pigsty | 120.1 KiB | [hydra_13-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hydra_13-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `hydra_13` | 1.1.2 | `el10.aarch64` | pigsty | 115.7 KiB | [hydra_13-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hydra_13-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-hydra` | 1.1.2 | `d12.x86_64` | pigsty | 364.6 KiB | [postgresql-13-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-13-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-hydra` | 1.1.2 | `d12.aarch64` | pigsty | 357.2 KiB | [postgresql-13-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-13-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-hydra` | 1.1.2 | `u22.x86_64` | pigsty | 387.0 KiB | [postgresql-13-hydra_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-13-hydra_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-hydra` | 1.1.2 | `u22.aarch64` | pigsty | 382.7 KiB | [postgresql-13-hydra_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-13-hydra_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-hydra` | 1.1.2 | `u24.x86_64` | pigsty | 332.0 KiB | [postgresql-13-hydra_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-13-hydra_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-hydra` | 1.1.2 | `u24.aarch64` | pigsty | 328.9 KiB | [postgresql-13-hydra_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-13-hydra_1.1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hydradatabase/hydra" title="Repository" icon="github" subtitle="github.com/hydradatabase/hydra" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hydra-1.1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get columnar; # get columnar source code
pig build dep columnar; # install build dependencies
pig build pkg columnar; # build extension rpm or deb
pig build ext columnar; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install columnar; # install by extension name, for the current active PG version
pig ext install hydra; # install via package alias, for the active PG version
pig ext install columnar -v 16;   # install for PG 16
pig ext install columnar -v 15;   # install for PG 15
pig ext install columnar -v 14;   # install for PG 14
pig ext install columnar -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION columnar;
```

