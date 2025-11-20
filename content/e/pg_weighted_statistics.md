---
title: "pg_weighted_statistics"
linkTitle: "pg_weighted_statistics"
description: "High-performance weighted statistics functions for sparse data"
weight: 4680
categories: ["FUNC"]
width: full
---

[**pg_weighted_statistics**](https://github.com/schmidni/pg_weighted_statistics) : High-performance weighted statistics functions for sparse data


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4680** | {{< badge content="pg_weighted_statistics" link="https://github.com/schmidni/pg_weighted_statistics" >}} | {{< ext "pg_weighted_statistics" >}} | `1.0.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_weighted_statistics` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_weighted_statistics_18*" "green" >}} {{< bg "17" "pg_weighted_statistics_17*" "green" >}} {{< bg "16" "pg_weighted_statistics_16*" "green" >}} {{< bg "15" "pg_weighted_statistics_15*" "green" >}} {{< bg "14" "pg_weighted_statistics_14*" "green" >}} {{< bg "13" "pg_weighted_statistics_13*" "green" >}} | `pg_weighted_statistics_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-weighted-statistics" "green" >}} {{< bg "17" "postgresql-17-weighted-statistics" "green" >}} {{< bg "16" "postgresql-16-weighted-statistics" "green" >}} {{< bg "15" "postgresql-15-weighted-statistics" "green" >}} {{< bg "14" "postgresql-14-weighted-statistics" "green" >}} {{< bg "13" "postgresql-13-weighted-statistics" "green" >}} | `postgresql-$v-weighted-statistics` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_weighted_statistics_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-weighted-statistics : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-weighted-statistics : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-weighted-statistics : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-weighted-statistics : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-weighted-statistics : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-weighted-statistics : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-weighted-statistics : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-weighted-statistics : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-weighted-statistics : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_weighted_statistics_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.4 KiB | [pg_weighted_statistics_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_weighted_statistics_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_weighted_statistics_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.6 KiB | [pg_weighted_statistics_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_weighted_statistics_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_weighted_statistics_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.0 KiB | [pg_weighted_statistics_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_weighted_statistics_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_weighted_statistics_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.1 KiB | [pg_weighted_statistics_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_weighted_statistics_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_weighted_statistics_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.2 KiB | [pg_weighted_statistics_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_weighted_statistics_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_weighted_statistics_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 26.1 KiB | [pg_weighted_statistics_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_weighted_statistics_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-weighted-statistics` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.9 KiB | [postgresql-18-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-18-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-weighted-statistics` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.1 KiB | [postgresql-18-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-18-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-weighted-statistics` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.5 KiB | [postgresql-18-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-18-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-weighted-statistics` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.9 KiB | [postgresql-18-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-18-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-weighted-statistics` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.2 KiB | [postgresql-18-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-18-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-weighted-statistics` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.4 KiB | [postgresql-18-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-18-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-weighted-statistics` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.2 KiB | [postgresql-18-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-18-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-weighted-statistics` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 34.7 KiB | [postgresql-18-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-18-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_weighted_statistics_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.4 KiB | [pg_weighted_statistics_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_weighted_statistics_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_weighted_statistics_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.6 KiB | [pg_weighted_statistics_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_weighted_statistics_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_weighted_statistics_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.0 KiB | [pg_weighted_statistics_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_weighted_statistics_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_weighted_statistics_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.1 KiB | [pg_weighted_statistics_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_weighted_statistics_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_weighted_statistics_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.2 KiB | [pg_weighted_statistics_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_weighted_statistics_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_weighted_statistics_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 26.1 KiB | [pg_weighted_statistics_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_weighted_statistics_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-weighted-statistics` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.9 KiB | [postgresql-17-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-17-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-weighted-statistics` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.1 KiB | [postgresql-17-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-17-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-weighted-statistics` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.5 KiB | [postgresql-17-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-17-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-weighted-statistics` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.8 KiB | [postgresql-17-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-17-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-weighted-statistics` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.3 KiB | [postgresql-17-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-17-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-weighted-statistics` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.6 KiB | [postgresql-17-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-17-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-weighted-statistics` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.2 KiB | [postgresql-17-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-17-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-weighted-statistics` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 34.7 KiB | [postgresql-17-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-17-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_weighted_statistics_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.4 KiB | [pg_weighted_statistics_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_weighted_statistics_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_weighted_statistics_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.6 KiB | [pg_weighted_statistics_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_weighted_statistics_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_weighted_statistics_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 26.0 KiB | [pg_weighted_statistics_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_weighted_statistics_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_weighted_statistics_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 26.1 KiB | [pg_weighted_statistics_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_weighted_statistics_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_weighted_statistics_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.2 KiB | [pg_weighted_statistics_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_weighted_statistics_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_weighted_statistics_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 26.1 KiB | [pg_weighted_statistics_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_weighted_statistics_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-weighted-statistics` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.9 KiB | [postgresql-16-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-16-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-weighted-statistics` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.1 KiB | [postgresql-16-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-16-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-weighted-statistics` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.5 KiB | [postgresql-16-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-16-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-weighted-statistics` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.8 KiB | [postgresql-16-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-16-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-weighted-statistics` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.4 KiB | [postgresql-16-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-16-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-weighted-statistics` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.6 KiB | [postgresql-16-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-16-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-weighted-statistics` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.2 KiB | [postgresql-16-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-16-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-weighted-statistics` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 34.7 KiB | [postgresql-16-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-16-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_weighted_statistics_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.4 KiB | [pg_weighted_statistics_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_weighted_statistics_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_weighted_statistics_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.6 KiB | [pg_weighted_statistics_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_weighted_statistics_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_weighted_statistics_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.4 KiB | [pg_weighted_statistics_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_weighted_statistics_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_weighted_statistics_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.2 KiB | [pg_weighted_statistics_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_weighted_statistics_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_weighted_statistics_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.3 KiB | [pg_weighted_statistics_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_weighted_statistics_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_weighted_statistics_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.9 KiB | [pg_weighted_statistics_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_weighted_statistics_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-weighted-statistics` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.9 KiB | [postgresql-15-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-15-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-weighted-statistics` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.1 KiB | [postgresql-15-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-15-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-weighted-statistics` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.5 KiB | [postgresql-15-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-15-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-weighted-statistics` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.9 KiB | [postgresql-15-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-15-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-weighted-statistics` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.7 KiB | [postgresql-15-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-15-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-weighted-statistics` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.8 KiB | [postgresql-15-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-15-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-weighted-statistics` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.3 KiB | [postgresql-15-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-15-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-weighted-statistics` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 34.8 KiB | [postgresql-15-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-15-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_weighted_statistics_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.4 KiB | [pg_weighted_statistics_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_weighted_statistics_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_weighted_statistics_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.6 KiB | [pg_weighted_statistics_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_weighted_statistics_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_weighted_statistics_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.4 KiB | [pg_weighted_statistics_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_weighted_statistics_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_weighted_statistics_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.2 KiB | [pg_weighted_statistics_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_weighted_statistics_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_weighted_statistics_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.3 KiB | [pg_weighted_statistics_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_weighted_statistics_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_weighted_statistics_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.9 KiB | [pg_weighted_statistics_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_weighted_statistics_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-weighted-statistics` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.9 KiB | [postgresql-14-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-14-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-weighted-statistics` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.1 KiB | [postgresql-14-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-14-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-weighted-statistics` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.5 KiB | [postgresql-14-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-14-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-weighted-statistics` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.9 KiB | [postgresql-14-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-14-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-weighted-statistics` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.7 KiB | [postgresql-14-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-14-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-weighted-statistics` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.7 KiB | [postgresql-14-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-14-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-weighted-statistics` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.3 KiB | [postgresql-14-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-14-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-weighted-statistics` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 34.8 KiB | [postgresql-14-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-14-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_weighted_statistics_13` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.3 KiB | [pg_weighted_statistics_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_weighted_statistics_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_weighted_statistics_13` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.6 KiB | [pg_weighted_statistics_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_weighted_statistics_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_weighted_statistics_13` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.4 KiB | [pg_weighted_statistics_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_weighted_statistics_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_weighted_statistics_13` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.2 KiB | [pg_weighted_statistics_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_weighted_statistics_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_weighted_statistics_13` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.3 KiB | [pg_weighted_statistics_13-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_weighted_statistics_13-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_weighted_statistics_13` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.9 KiB | [pg_weighted_statistics_13-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_weighted_statistics_13-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-weighted-statistics` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.8 KiB | [postgresql-13-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-13-weighted-statistics_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-weighted-statistics` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 33.7 KiB | [postgresql-13-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-weighted-statistics/postgresql-13-weighted-statistics_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-weighted-statistics` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.4 KiB | [postgresql-13-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-13-weighted-statistics_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-weighted-statistics` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.5 KiB | [postgresql-13-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-weighted-statistics/postgresql-13-weighted-statistics_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-weighted-statistics` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.5 KiB | [postgresql-13-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-13-weighted-statistics_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-weighted-statistics` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.3 KiB | [postgresql-13-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-weighted-statistics/postgresql-13-weighted-statistics_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-weighted-statistics` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.2 KiB | [postgresql-13-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-13-weighted-statistics_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-weighted-statistics` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 34.4 KiB | [postgresql-13-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-weighted-statistics/postgresql-13-weighted-statistics_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/schmidni/pg_weighted_statistics" title="Repository" icon="github" subtitle="github.com/schmidni/pg_weighted_statistics" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_weighted_statistics-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_weighted_statistics;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_weighted_statistics;		# install via package name, for the active PG version

pig install pg_weighted_statistics -v 18;   # install for PG 18
pig install pg_weighted_statistics -v 17;   # install for PG 17
pig install pg_weighted_statistics -v 16;   # install for PG 16
pig install pg_weighted_statistics -v 15;   # install for PG 15
pig install pg_weighted_statistics -v 14;   # install for PG 14
pig install pg_weighted_statistics -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_weighted_statistics;
```
