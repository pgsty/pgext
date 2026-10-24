---
title: "pg_relusage"
linkTitle: "pg_relusage"
description: "Log all the queries that reference a particular column"
weight: 6850
categories: ["STAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_relusage**](https://github.com/adept/pg_relusage) : Log all the queries that reference a particular column


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6850** | {{< badge content="pg_relusage" link="https://github.com/adept/pg_relusage" >}} | {{< ext "pg_relusage" >}} | `1.0.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL--r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_qualstats" >}} {{< ext "online_advisor" >}} {{< ext "hypopg" >}} {{< ext "pg_stat_statements" >}} {{< ext "index_advisor" >}} {{< ext "pg_stat_plans" >}} {{< ext "pg_upless" >}} {{< ext "pg_hint_plan" >}} {{< ext "pg_track_optimizer" >}} {{< ext "plan_filter" >}} |

> [!Note] LOAD/preload enables logging; CREATE EXTENSION only records installation.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_relusage` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "pg_relusage_18" "green" >}} {{< bg "17" "pg_relusage_17" "green" >}} {{< bg "16" "pg_relusage_16" "green" >}} {{< bg "15" "pg_relusage_15" "green" >}} {{< bg "14" "pg_relusage_14" "green" >}} | `pg_relusage_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "postgresql-18-pg-relusage" "green" >}} {{< bg "17" "postgresql-17-pg-relusage" "green" >}} {{< bg "16" "postgresql-16-pg-relusage" "green" >}} {{< bg "15" "postgresql-15-pg-relusage" "green" >}} {{< bg "14" "postgresql-14-pg-relusage" "green" >}} | `postgresql-$v-pg-relusage` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_relusage_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_18` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 39.3 KiB | [pg_relusage_18-1.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_18-1.0.2-1PGSTY.el8.x86_64.rpm) |
| `pg_relusage_18` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 39.0 KiB | [pg_relusage_18-1.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_18-1.0.2-1PGSTY.el8.aarch64.rpm) |
| `pg_relusage_18` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 39.8 KiB | [pg_relusage_18-1.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_18-1.0.2-1PGSTY.el9.x86_64.rpm) |
| `pg_relusage_18` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 39.6 KiB | [pg_relusage_18-1.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_18-1.0.2-1PGSTY.el9.aarch64.rpm) |
| `pg_relusage_18` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 40.0 KiB | [pg_relusage_18-1.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_18-1.0.2-1PGSTY.el10.x86_64.rpm) |
| `pg_relusage_18` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 39.9 KiB | [pg_relusage_18-1.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_18-1.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-relusage` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.4 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.4 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.8 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.4 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.3 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.7 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.9 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.3 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-relusage` | `1.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.2 KiB | [postgresql-18-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-18-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_17` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 39.2 KiB | [pg_relusage_17-1.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_17-1.0.2-1PGSTY.el8.x86_64.rpm) |
| `pg_relusage_17` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 39.0 KiB | [pg_relusage_17-1.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_17-1.0.2-1PGSTY.el8.aarch64.rpm) |
| `pg_relusage_17` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 39.7 KiB | [pg_relusage_17-1.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_17-1.0.2-1PGSTY.el9.x86_64.rpm) |
| `pg_relusage_17` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 39.6 KiB | [pg_relusage_17-1.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_17-1.0.2-1PGSTY.el9.aarch64.rpm) |
| `pg_relusage_17` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 40.0 KiB | [pg_relusage_17-1.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_17-1.0.2-1PGSTY.el10.x86_64.rpm) |
| `pg_relusage_17` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 39.9 KiB | [pg_relusage_17-1.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_17-1.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-relusage` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.3 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.2 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.8 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.2 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.2 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.5 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.1 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.2 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-relusage` | `1.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.1 KiB | [postgresql-17-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-17-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_16` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 39.2 KiB | [pg_relusage_16-1.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_16-1.0.2-1PGSTY.el8.x86_64.rpm) |
| `pg_relusage_16` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 39.0 KiB | [pg_relusage_16-1.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_16-1.0.2-1PGSTY.el8.aarch64.rpm) |
| `pg_relusage_16` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 39.7 KiB | [pg_relusage_16-1.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_16-1.0.2-1PGSTY.el9.x86_64.rpm) |
| `pg_relusage_16` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 39.6 KiB | [pg_relusage_16-1.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_16-1.0.2-1PGSTY.el9.aarch64.rpm) |
| `pg_relusage_16` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 40.0 KiB | [pg_relusage_16-1.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_16-1.0.2-1PGSTY.el10.x86_64.rpm) |
| `pg_relusage_16` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 39.9 KiB | [pg_relusage_16-1.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_16-1.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-relusage` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.3 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.2 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.8 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.8 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.8 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.5 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.9 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.3 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-relusage` | `1.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.2 KiB | [postgresql-16-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-16-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_15` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 39.2 KiB | [pg_relusage_15-1.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_15-1.0.2-1PGSTY.el8.x86_64.rpm) |
| `pg_relusage_15` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 39.0 KiB | [pg_relusage_15-1.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_15-1.0.2-1PGSTY.el8.aarch64.rpm) |
| `pg_relusage_15` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 39.7 KiB | [pg_relusage_15-1.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_15-1.0.2-1PGSTY.el9.x86_64.rpm) |
| `pg_relusage_15` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 39.6 KiB | [pg_relusage_15-1.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_15-1.0.2-1PGSTY.el9.aarch64.rpm) |
| `pg_relusage_15` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 39.9 KiB | [pg_relusage_15-1.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_15-1.0.2-1PGSTY.el10.x86_64.rpm) |
| `pg_relusage_15` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 39.9 KiB | [pg_relusage_15-1.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_15-1.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-relusage` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.3 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.1 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.2 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.8 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.7 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.7 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.5 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.0 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.2 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-relusage` | `1.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.2 KiB | [postgresql-15-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-15-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_14` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 39.2 KiB | [pg_relusage_14-1.0.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_14-1.0.2-1PGSTY.el8.x86_64.rpm) |
| `pg_relusage_14` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 39.0 KiB | [pg_relusage_14-1.0.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_14-1.0.2-1PGSTY.el8.aarch64.rpm) |
| `pg_relusage_14` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 39.6 KiB | [pg_relusage_14-1.0.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_14-1.0.2-1PGSTY.el9.x86_64.rpm) |
| `pg_relusage_14` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 39.5 KiB | [pg_relusage_14-1.0.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_14-1.0.2-1PGSTY.el9.aarch64.rpm) |
| `pg_relusage_14` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 39.9 KiB | [pg_relusage_14-1.0.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_14-1.0.2-1PGSTY.el10.x86_64.rpm) |
| `pg_relusage_14` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 39.8 KiB | [pg_relusage_14-1.0.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_14-1.0.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-relusage` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.2 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.6 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.2 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.8 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.6 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.7 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.4 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.8 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.2 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-relusage` | `1.0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.1 KiB | [postgresql-14-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-relusage/postgresql-14-pg-relusage_1.0.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adept/pg_relusage" title="Repository" icon="github" subtitle="github.com/adept/pg_relusage" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_relusage-1.0.2.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_relusage;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_relusage;		# install via package name, for the active PG version

pig install pg_relusage -v 18;   # install for PG 18
pig install pg_relusage -v 17;   # install for PG 17
pig install pg_relusage -v 16;   # install for PG 16
pig install pg_relusage -v 15;   # install for PG 15
pig install pg_relusage -v 14;   # install for PG 14

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_relusage';
```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources:

- [README.md](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/README.md)
- [Control file](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/pg_relusage.control)
- [Implementation](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/pg_relusage.c)

`pg_relusage` 1.0.2 logs relations used by statements after plan simplification. It supports PostgreSQL 10–18 and defines no SQL objects. Loading the library enables it; optional CREATE EXTENSION only records installation.

### Core Workflow

```sql
LOAD 'pg_relusage';
SET pg_relusage.nested_statements = 'fold';
SET pg_relusage.log_level = 'notice';
CREATE TABLE relusage_demo (id integer);
SELECT * FROM relusage_demo;
```

### Operational Boundaries

All settings require superuser: `pg_relusage.enabled` defaults on, `pg_relusage.log_level` defaults log and `pg_relusage.rel_kinds` defaults riSvmfp. `pg_relusage.nested_statements` accepts `all` (default: each statement), `top` (ignore nested statements) or `fold` (merge nested relations into the outer log). Version 1.0.1 fixes fold handling after a failed subtransaction.

Use `session_preload_libraries` for new sessions without a restart, or `shared_preload_libraries` for all sessions with a restart; preserve existing preload entries. LOAD normally requires superuser. Log messages omit SQL text, so configure `log_line_prefix` to identify the client.

System catalogs are excluded. Relations eliminated by planning or startup pruning are excluded, but partitions skipped later during execution can still appear. INSERT into a partitioned table reports only its parent. Deferred triggers at COMMIT, COPY FROM triggers and planner-invoked functions are treated as top-level statements. DDL and maintenance commands are excluded. A quiet log alone does not prove an object can be dropped.

### 1.0.2 and Partitioned Tables

Version 1.0.2 reduces relation-membership and partition-matching overhead for statements involving many partitions. Log format and settings are unchanged. Replace the library and restart if using shared preload, or reconnect when using session preload or LOAD. If installation was recorded with CREATE EXTENSION, also run `ALTER EXTENSION pg_relusage UPDATE`.
