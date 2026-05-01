---
title: "pg_isok"
linkTitle: "pg_isok"
description: "Query-based data integrity management and soft alerting for PostgreSQL"
weight: 4340
categories: ["UTIL"]
width: full
---

[**pg_isok**](https://codeberg.org/kop/pg_isok) : Query-based data integrity management and soft alerting for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4340** | {{< badge content="pg_isok" link="https://codeberg.org/kop/pg_isok" >}} | {{< ext "pg_isok" >}} | `1.4.1` | {{< category "UTIL" >}} | {{< license "AGPL-3.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] superuser=false, but this is not a trusted extension.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_isok` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "pg_isok_18" "green" >}} {{< bg "17" "pg_isok_17" "green" >}} {{< bg "16" "pg_isok_16" "green" >}} {{< bg "15" "pg_isok_15" "green" >}} {{< bg "14" "pg_isok_14" "green" >}} | `pg_isok_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.1` | {{< bg "18" "postgresql-18-pg-isok" "green" >}} {{< bg "17" "postgresql-17-pg-isok" "green" >}} {{< bg "16" "postgresql-16-pg-isok" "green" >}} {{< bg "15" "postgresql-15-pg-isok" "green" >}} {{< bg "14" "postgresql-14-pg-isok" "green" >}} | `postgresql-$v-pg-isok` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "pg_isok_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-18-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-17-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-16-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-15-pg-isok : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.1" "postgresql-14-pg-isok : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_isok_18` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_isok_18-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_isok_18-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_isok_18` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.6 KiB | [pg_isok_18-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_isok_18-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_isok_18` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.6 KiB | [pg_isok_18-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_isok_18-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_isok_18` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 60.6 KiB | [pg_isok_18-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_isok_18-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_isok_18` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.7 KiB | [pg_isok_18-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_isok_18-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_isok_18` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.7 KiB | [pg_isok_18-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_isok_18-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-isok` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 56.8 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.8 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 56.8 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 56.8 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.0 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 57.0 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 56.9 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 56.9 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 56.9 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-isok` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 56.9 KiB | [postgresql-18-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-18-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_isok_17` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_isok_17-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_isok_17-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_isok_17` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.6 KiB | [pg_isok_17-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_isok_17-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_isok_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.6 KiB | [pg_isok_17-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_isok_17-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_isok_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 60.6 KiB | [pg_isok_17-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_isok_17-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_isok_17` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.7 KiB | [pg_isok_17-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_isok_17-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_isok_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.7 KiB | [pg_isok_17-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_isok_17-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-isok` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 56.8 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.8 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 56.8 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 56.8 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.0 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 57.0 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 56.9 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 56.9 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 56.9 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-isok` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 56.9 KiB | [postgresql-17-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-17-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_isok_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_isok_16-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_isok_16-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_isok_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.6 KiB | [pg_isok_16-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_isok_16-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_isok_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.6 KiB | [pg_isok_16-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_isok_16-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_isok_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 60.6 KiB | [pg_isok_16-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_isok_16-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_isok_16` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.7 KiB | [pg_isok_16-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_isok_16-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_isok_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.7 KiB | [pg_isok_16-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_isok_16-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-isok` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 56.8 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.8 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 56.8 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 56.8 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.0 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 57.0 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 56.9 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 56.9 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 56.9 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-isok` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 56.9 KiB | [postgresql-16-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-16-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_isok_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_isok_15-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_isok_15-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_isok_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.6 KiB | [pg_isok_15-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_isok_15-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_isok_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.6 KiB | [pg_isok_15-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_isok_15-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_isok_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 60.6 KiB | [pg_isok_15-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_isok_15-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_isok_15` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.7 KiB | [pg_isok_15-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_isok_15-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_isok_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.7 KiB | [pg_isok_15-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_isok_15-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-isok` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 56.8 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.8 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 56.8 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 56.8 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.0 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 57.0 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 56.9 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 56.9 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 56.9 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-isok` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 56.9 KiB | [postgresql-15-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-15-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_isok_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.7 KiB | [pg_isok_14-1.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_isok_14-1.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_isok_14` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 63.6 KiB | [pg_isok_14-1.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_isok_14-1.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_isok_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.6 KiB | [pg_isok_14-1.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_isok_14-1.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_isok_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 60.6 KiB | [pg_isok_14-1.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_isok_14-1.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_isok_14` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 60.7 KiB | [pg_isok_14-1.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_isok_14-1.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_isok_14` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 60.7 KiB | [pg_isok_14-1.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_isok_14-1.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-isok` | `1.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 56.8 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 56.8 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 56.8 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 56.8 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.0 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 57.0 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 56.9 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 56.9 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 56.9 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-isok` | `1.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 56.9 KiB | [postgresql-14-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-isok/postgresql-14-pg-isok_1.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/kop/pg_isok" title="Repository" icon="link" subtitle="codeberg.org/kop/pg_isok" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_isok-1.4.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_isok;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_isok;		# install via package name, for the active PG version

pig install pg_isok -v 18;   # install for PG 18
pig install pg_isok -v 17;   # install for PG 17
pig install pg_isok -v 16;   # install for PG 16
pig install pg_isok -v 15;   # install for PG 15
pig install pg_isok -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_isok;
```

## Usage

Sources: [official repo](https://codeberg.org/kop/pg_isok), [official docs home](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/index.html), [official reference source](https://codeberg.org/kop/pg_isok/src/branch/main/doc_src/isok.xml)

`pg_isok` is a query-based data integrity and monitoring extension. Instead of only reporting rows that currently look questionable, it stores prior results and focuses later runs on unresolved or undeferred changes.

```sql
CREATE SCHEMA isok;
CREATE EXTENSION pg_isok SCHEMA isok;

SELECT *
FROM isok.run_isok_queries()
AS problems;
```

### Core Objects

- `ISOK_QUERIES` stores the monitoring queries and their execution settings.
- `ISOK_RESULTS` stores the reported rows, including whether they were resolved or deferred.
- `run_isok_queries()` runs every active check.
- `run_isok_queries($$VALUES ('check_name')$$)` runs only selected checks.

### Typical Workflow

Run one named check:

```sql
SELECT *
FROM isok.run_isok_queries($$VALUES ('new_countries')$$)
AS problems;
```

Accept or postpone a known warning by updating `ISOK_RESULTS`:

```sql
UPDATE isok.isok_results
SET deferred_to = 'infinity'
WHERE iqname = 'new_countries';
```

Use `resolved` when the condition is no longer a concern, or `deferred_to` when it should stay hidden until a later date.

### Where It Fits

- data cleanup after imports
- monitoring unusual but sometimes acceptable patterns
- "soft trigger" style review workflows where hard constraints are too strict

### Caveats

- Upstream recommends installing it in a dedicated schema and qualifying calls accordingly.
- The docs describe it as pure SQL, which is useful on managed PostgreSQL services where C extensions may be restricted.
- The package metadata in this repo says `superuser=false`, but this is not documented upstream as a trusted extension; treat installation privileges conservatively.
