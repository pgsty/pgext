---
title: "columnar"
linkTitle: "columnar"
description: "Hydra Columnar extension"
weight: 2410
categories: ["OLAP"]
width: full
---

[**hydra**](https://github.com/hydradatabase/hydra) : Hydra Columnar extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2410** | {{< badge content="columnar" link="https://github.com/hydradatabase/hydra" >}} | {{< ext "columnar" "hydra" >}} | `1.1.2` | {{< category "OLAP" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "citus" >}} {{< ext "citus_columnar" >}} {{< ext "pg_mooncake" >}} {{< ext "timescaledb" >}} {{< ext "pg_analytics" >}} {{< ext "pg_parquet" >}} {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} |

> [!Note] conflict with citus columnar, obsolete, no longer maintained


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "" "red" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `hydra` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "hydra_18" "red" >}} {{< bg "17" "hydra_17" "red" >}} {{< bg "16" "hydra_16" "green" >}} {{< bg "15" "hydra_15" "green" >}} {{< bg "14" "hydra_14" "green" >}} | `hydra_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "postgresql-18-hydra" "red" >}} {{< bg "17" "postgresql-17-hydra" "red" >}} {{< bg "16" "postgresql-16-hydra" "green" >}} {{< bg "15" "postgresql-15-hydra" "green" >}} {{< bg "14" "postgresql-14-hydra" "green" >}} | `postgresql-$v-hydra` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "hydra_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "hydra_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "hydra_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "hydra_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "hydra_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "hydra_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "hydra_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "hydra_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "hydra_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "hydra_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "hydra_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "hydra_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "hydra_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "hydra_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "hydra_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "hydra_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "hydra_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "hydra_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "hydra_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-hydra : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hydra : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1.2" "postgresql-16-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-hydra : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-hydra : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hydra_16` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 143.1 KiB | [hydra_16-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hydra_16-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `hydra_16` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 136.7 KiB | [hydra_16-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hydra_16-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `hydra_16` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 116.2 KiB | [hydra_16-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hydra_16-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `hydra_16` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 112.5 KiB | [hydra_16-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hydra_16-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `hydra_16` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 118.3 KiB | [hydra_16-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hydra_16-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `hydra_16` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 113.9 KiB | [hydra_16-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hydra_16-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-hydra` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 354.6 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 409.0 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 355.8 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 347.6 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 430.5 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 425.2 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 359.9 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 363.6 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 358.0 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-hydra` | `1.1.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 355.8 KiB | [postgresql-16-hydra_1.1.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hydra/postgresql-16-hydra_1.1.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hydra_15` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 146.7 KiB | [hydra_15-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hydra_15-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `hydra_15` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 140.2 KiB | [hydra_15-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hydra_15-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `hydra_15` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 136.8 KiB | [hydra_15-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hydra_15-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `hydra_15` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 131.2 KiB | [hydra_15-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hydra_15-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `hydra_15` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 138.1 KiB | [hydra_15-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hydra_15-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `hydra_15` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 132.3 KiB | [hydra_15-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hydra_15-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-hydra` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 358.3 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 412.7 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 359.5 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 349.7 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 449.6 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 443.4 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 377.5 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 381.2 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 375.7 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-hydra` | `1.1.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 371.3 KiB | [postgresql-15-hydra_1.1.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hydra/postgresql-15-hydra_1.1.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hydra_14` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 146.8 KiB | [hydra_14-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hydra_14-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `hydra_14` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 140.2 KiB | [hydra_14-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hydra_14-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `hydra_14` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 137.1 KiB | [hydra_14-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hydra_14-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `hydra_14` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 131.5 KiB | [hydra_14-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hydra_14-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `hydra_14` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 138.2 KiB | [hydra_14-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hydra_14-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `hydra_14` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 133.0 KiB | [hydra_14-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hydra_14-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-hydra` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 359.9 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 414.1 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 360.6 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 350.6 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 451.4 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 444.9 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 378.5 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 382.5 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 376.6 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-hydra` | `1.1.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 372.3 KiB | [postgresql-14-hydra_1.1.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hydra/postgresql-14-hydra_1.1.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hydradatabase/hydra" title="Repository" icon="github" subtitle="github.com/hydradatabase/hydra" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hydra-1.1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg hydra;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install hydra;		# install via package name, for the active PG version
pig install columnar;		# install by extension name, for the current active PG version

pig install columnar -v 16;   # install for PG 16
pig install columnar -v 15;   # install for PG 15
pig install columnar -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION columnar;
```




## Usage

Sources:

- [Hydra v1.1.2 README](https://github.com/hydradatabase/columnar/blob/v1.1.2/README.md)
- [Hydra Columnar README](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/README.md)
- [Columnar storage README](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/src/backend/columnar/README.md)
- [columnar.control](https://github.com/hydradatabase/columnar/blob/v1.1.2/columnar/src/backend/columnar/columnar.control)
- [CHANGELOG 1.1.2](https://github.com/hydradatabase/columnar/blob/v1.1.2/CHANGELOG.md)

> [!WARNING]
> `columnar` is the PostgreSQL extension name, while Pigsty packages it as `hydra`. Pigsty metadata marks this extension as obsolete and no longer maintained; local packages are retained for PostgreSQL 14-16 only. Prefer newer actively maintained analytics extensions for new deployments.

Hydra Columnar is a PostgreSQL table access method for column-oriented storage. It stores selected tables in a columnar format to reduce I/O for analytical scans, compression-heavy datasets, projections over a subset of columns, and aggregate queries. It originated from Citus Columnar and is exposed through `CREATE EXTENSION columnar`.

### Create Columnar Tables

```sql
CREATE EXTENSION IF NOT EXISTS columnar;

CREATE TABLE events_columnar (
  event_id     bigint,
  account_id   bigint,
  event_time   timestamptz,
  event_type   text,
  amount       numeric
) USING columnar;

INSERT INTO events_columnar
SELECT
  g,
  g % 10000,
  now() - (g || ' seconds')::interval,
  CASE WHEN g % 10 = 0 THEN 'purchase' ELSE 'view' END,
  (g % 1000)::numeric / 10
FROM generate_series(1, 1000000) AS g;

SELECT event_type, count(*), sum(amount)
FROM events_columnar
WHERE event_time >= now() - interval '1 day'
GROUP BY event_type;
```

Use `USING columnar` when creating a table or materialized view. Reads and bulk inserts use normal PostgreSQL SQL, but the storage format is optimized for large analytical scans rather than high-churn OLTP tables.

### Table Options

```sql
SELECT columnar.alter_columnar_table_set(
  'events_columnar'::regclass,
  compression           => 'zstd',
  compression_level     => 3,
  stripe_row_limit      => 150000,
  chunk_group_row_limit => 10000
);

SELECT * FROM columnar.options;

SELECT columnar.alter_columnar_table_reset(
  'events_columnar'::regclass,
  compression => true,
  stripe_row_limit => true
);
```

Available table options include `compression`, `compression_level`, `stripe_row_limit`, and `chunk_group_row_limit`. Compression choices depend on build support, but documented values include `none`, `pglz`, `zstd`, `lz4`, and `lz4hc`. Option changes apply to newly inserted data; existing stripes are not automatically rewritten.

You can also set defaults for newly created columnar tables with GUCs:

```sql
SET columnar.compression = 'zstd';
SET columnar.compression_level = 3;
SET columnar.stripe_row_limit = 150000;
SET columnar.chunk_group_row_limit = 10000;
```

These GUCs affect newly created tables, not new stripes on an already existing table.

### Convert Existing Tables

```sql
CREATE TABLE events_heap (
  event_id bigint,
  payload  jsonb
);

INSERT INTO events_heap
SELECT g, jsonb_build_object('kind', 'view', 'seq', g)
FROM generate_series(1, 10000) AS g;

SELECT columnar.alter_table_set_access_method('events_heap', 'columnar');
SELECT columnar.alter_table_set_access_method('events_heap', 'heap');
```

`columnar.alter_table_set_access_method(table, method)` rewrites a heap table as columnar storage or a columnar table back to heap storage. Review foreign keys, identity columns, row-level security, triggers, indexes, constraints, inheritance, and storage options before conversion; the helper rejects or skips features it cannot safely recreate.

### Partitioning

```sql
CREATE TABLE measurements (
  ts timestamptz,
  device_id bigint,
  value double precision
) PARTITION BY RANGE (ts);

CREATE TABLE measurements_2024 PARTITION OF measurements
  FOR VALUES FROM ('2024-01-01') TO ('2025-01-01')
  USING columnar;

CREATE TABLE measurements_hot PARTITION OF measurements
  FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');
```

Partitioned tables can mix row and columnar partitions. Operations that target only row partitions can use row-table behavior, while operations that touch columnar partitions must respect columnar limitations. This is useful for keeping recent mutable data in heap partitions and older analytical history in columnar partitions.

### Maintenance and Introspection

```sql
VACUUM VERBOSE events_columnar;
VACUUM FULL events_columnar;

SELECT * FROM columnar.stats('events_columnar'::regclass);
SELECT columnar.vacuum('events_columnar'::regclass);
SELECT columnar.vacuum_full('public', 0.1, 25);
```

`VACUUM VERBOSE` reports columnar storage statistics such as file size, compression rate, row count, stripe count, and chunk count. `columnar.stats()` exposes stripe-level metadata. `columnar.vacuum()` and `columnar.vacuum_full()` compact columnar storage incrementally; ordinary `VACUUM` is metadata-oriented and cheaper than a full rewrite.

### Caveats

- This extension is obsolete in Pigsty metadata and conflicts with `citus`/`citus_columnar` style columnar storage. Avoid installing conflicting columnar table access methods in the same PostgreSQL major unless you have tested the exact combination.
- Pigsty packages `hydra`/`columnar` for PostgreSQL 14-16; PostgreSQL 17 and 18 are marked unsupported locally.
- Hydra 1.1.x added update/delete and upsert improvements, but the project itself still describes columnar storage as unsuitable for frequent large updates, small transactions, and OLTP-style single-row workloads.
- Unsupported or limited areas include logical decoding, unlogged columnar tables, serializable isolation, some scan types, and many non-btree/non-hash indexes. Check constraints and index-backed constraints carefully before relying on them.
- The `columnar` schema contains internal metadata tables such as `columnar.options`, `columnar.stripe`, `columnar.chunk_group`, and `columnar.chunk`. Query public views/functions for inspection, but do not mutate metadata tables directly.
