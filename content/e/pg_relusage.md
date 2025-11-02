---
title: "pg_relusage"
linkTitle: "pg_relusage"
description: "Log all the queries that reference a particular column"
weight: 6380
categories: ["STAT"]
width: full
---

Log all the queries that reference a particular column


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6380** | {{< badge content="pg_relusage" link="https://github.com/adept/pg_relusage" >}} | {{< ext "pg_relusage" >}} | `0.0.1` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_freespacemap" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "toastinfo" >}} {{< ext "pageinspect" >}} {{< ext "pg_buffercache" >}} {{< ext "pgfincore" >}} {{< ext "old_snapshot" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_relusage" >}} | `0.0.1` | {{< bg "18" "pg_relusage_18" "green" >}} {{< bg "17" "pg_relusage_17" "green" >}} {{< bg "16" "pg_relusage_16" "green" >}} {{< bg "15" "pg_relusage_15" "green" >}} {{< bg "14" "pg_relusage_14" "green" >}} {{< bg "13" "pg_relusage_13" "green" >}} | `pg_relusage_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_relusage" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-relusage" "red" >}} {{< bg "17" "postgresql-17-pg-relusage" "green" >}} {{< bg "16" "postgresql-16-pg-relusage" "green" >}} {{< bg "15" "postgresql-15-pg-relusage" "green" >}} {{< bg "14" "postgresql-14-pg-relusage" "green" >}} {{< bg "13" "postgresql-13-pg-relusage" "green" >}} | `postgresql-$v-pg-relusage` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_relusage_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-relusage : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-relusage : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-relusage : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-relusage : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-relusage : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-relusage : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-relusage : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-relusage : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-relusage : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-relusage : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-relusage : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-relusage : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_18` | 0.0.1 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_relusage_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_relusage_18` | 0.0.1 | `el8.aarch64` | pigsty | 13.5 KiB | [pg_relusage_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_relusage_18` | 0.0.1 | `el9.x86_64` | pigsty | 13.0 KiB | [pg_relusage_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_relusage_18` | 0.0.1 | `el9.aarch64` | pigsty | 13.2 KiB | [pg_relusage_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_relusage_18` | 0.0.1 | `el10.x86_64` | pigsty | 13.0 KiB | [pg_relusage_18-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_18-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_relusage_18` | 0.0.1 | `el10.aarch64` | pigsty | 13.4 KiB | [pg_relusage_18-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_18-0.0.1-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_17` | 0.0.1 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_relusage_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_relusage_17` | 0.0.1 | `el8.aarch64` | pigsty | 13.5 KiB | [pg_relusage_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_relusage_17` | 0.0.1 | `el9.x86_64` | pigsty | 13.1 KiB | [pg_relusage_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_relusage_17` | 0.0.1 | `el9.aarch64` | pigsty | 13.3 KiB | [pg_relusage_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_relusage_17` | 0.0.1 | `el10.x86_64` | pigsty | 13.1 KiB | [pg_relusage_17-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_17-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_relusage_17` | 0.0.1 | `el10.aarch64` | pigsty | 13.4 KiB | [pg_relusage_17-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_17-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-relusage` | 0.0.1 | `d12.x86_64` | pigsty | 14.3 KiB | [postgresql-17-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-17-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-relusage` | 0.0.1 | `d12.aarch64` | pigsty | 14.3 KiB | [postgresql-17-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-17-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-relusage` | 0.0.1 | `u22.x86_64` | pigsty | 14.6 KiB | [postgresql-17-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-17-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-relusage` | 0.0.1 | `u22.aarch64` | pigsty | 14.7 KiB | [postgresql-17-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-17-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-relusage` | 0.0.1 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-17-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-17-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-relusage` | 0.0.1 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-17-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-17-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_16` | 0.0.1 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_relusage_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_relusage_16` | 0.0.1 | `el8.aarch64` | pigsty | 13.5 KiB | [pg_relusage_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_relusage_16` | 0.0.1 | `el9.x86_64` | pigsty | 13.1 KiB | [pg_relusage_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_relusage_16` | 0.0.1 | `el9.aarch64` | pigsty | 13.3 KiB | [pg_relusage_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_relusage_16` | 0.0.1 | `el10.x86_64` | pigsty | 13.1 KiB | [pg_relusage_16-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_16-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_relusage_16` | 0.0.1 | `el10.aarch64` | pigsty | 13.4 KiB | [pg_relusage_16-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_16-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-relusage` | 0.0.1 | `d12.x86_64` | pigsty | 14.3 KiB | [postgresql-16-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-16-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-relusage` | 0.0.1 | `d12.aarch64` | pigsty | 14.3 KiB | [postgresql-16-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-16-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-relusage` | 0.0.1 | `u22.x86_64` | pigsty | 14.5 KiB | [postgresql-16-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-16-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-relusage` | 0.0.1 | `u22.aarch64` | pigsty | 14.6 KiB | [postgresql-16-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-16-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-relusage` | 0.0.1 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-16-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-16-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-relusage` | 0.0.1 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-16-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-16-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_15` | 0.0.1 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_relusage_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_relusage_15` | 0.0.1 | `el8.aarch64` | pigsty | 13.5 KiB | [pg_relusage_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_relusage_15` | 0.0.1 | `el9.x86_64` | pigsty | 13.1 KiB | [pg_relusage_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_relusage_15` | 0.0.1 | `el9.aarch64` | pigsty | 13.3 KiB | [pg_relusage_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_relusage_15` | 0.0.1 | `el10.x86_64` | pigsty | 13.1 KiB | [pg_relusage_15-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_15-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_relusage_15` | 0.0.1 | `el10.aarch64` | pigsty | 13.4 KiB | [pg_relusage_15-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_15-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-relusage` | 0.0.1 | `d12.x86_64` | pigsty | 14.3 KiB | [postgresql-15-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-15-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-relusage` | 0.0.1 | `d12.aarch64` | pigsty | 14.3 KiB | [postgresql-15-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-15-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-relusage` | 0.0.1 | `u22.x86_64` | pigsty | 14.5 KiB | [postgresql-15-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-15-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-relusage` | 0.0.1 | `u22.aarch64` | pigsty | 14.6 KiB | [postgresql-15-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-15-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-relusage` | 0.0.1 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-15-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-15-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-relusage` | 0.0.1 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-15-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-15-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_14` | 0.0.1 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_relusage_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_relusage_14` | 0.0.1 | `el8.aarch64` | pigsty | 13.5 KiB | [pg_relusage_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_relusage_14` | 0.0.1 | `el9.x86_64` | pigsty | 13.1 KiB | [pg_relusage_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_relusage_14` | 0.0.1 | `el9.aarch64` | pigsty | 13.3 KiB | [pg_relusage_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_relusage_14` | 0.0.1 | `el10.x86_64` | pigsty | 13.1 KiB | [pg_relusage_14-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_14-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_relusage_14` | 0.0.1 | `el10.aarch64` | pigsty | 13.4 KiB | [pg_relusage_14-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_14-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-relusage` | 0.0.1 | `d12.x86_64` | pigsty | 14.2 KiB | [postgresql-14-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-14-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-relusage` | 0.0.1 | `d12.aarch64` | pigsty | 14.2 KiB | [postgresql-14-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-14-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-relusage` | 0.0.1 | `u22.x86_64` | pigsty | 14.5 KiB | [postgresql-14-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-14-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-relusage` | 0.0.1 | `u22.aarch64` | pigsty | 14.6 KiB | [postgresql-14-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-14-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-relusage` | 0.0.1 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-14-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-14-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-relusage` | 0.0.1 | `u24.aarch64` | pigsty | 12.5 KiB | [postgresql-14-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-14-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_relusage_13` | 0.0.1 | `el8.x86_64` | pigsty | 13.1 KiB | [pg_relusage_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_relusage_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_relusage_13` | 0.0.1 | `el8.aarch64` | pigsty | 13.5 KiB | [pg_relusage_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_relusage_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_relusage_13` | 0.0.1 | `el9.x86_64` | pigsty | 13.1 KiB | [pg_relusage_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_relusage_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_relusage_13` | 0.0.1 | `el9.aarch64` | pigsty | 13.3 KiB | [pg_relusage_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_relusage_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_relusage_13` | 0.0.1 | `el10.x86_64` | pigsty | 13.1 KiB | [pg_relusage_13-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_relusage_13-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_relusage_13` | 0.0.1 | `el10.aarch64` | pigsty | 13.5 KiB | [pg_relusage_13-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_relusage_13-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-relusage` | 0.0.1 | `d12.x86_64` | pigsty | 13.9 KiB | [postgresql-13-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-13-pg-relusage_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-relusage` | 0.0.1 | `d12.aarch64` | pigsty | 14.1 KiB | [postgresql-13-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-relusage/postgresql-13-pg-relusage_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-relusage` | 0.0.1 | `u22.x86_64` | pigsty | 14.2 KiB | [postgresql-13-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-13-pg-relusage_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-relusage` | 0.0.1 | `u22.aarch64` | pigsty | 14.4 KiB | [postgresql-13-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-relusage/postgresql-13-pg-relusage_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-relusage` | 0.0.1 | `u24.x86_64` | pigsty | 12.5 KiB | [postgresql-13-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-13-pg-relusage_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-relusage` | 0.0.1 | `u24.aarch64` | pigsty | 12.5 KiB | [postgresql-13-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-relusage/postgresql-13-pg-relusage_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adept/pg_relusage" title="Repository" icon="github" subtitle="github.com/adept/pg_relusage" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_relusage-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_relusage; # get pg_relusage source code
pig build dep pg_relusage; # install build dependencies
pig build pkg pg_relusage; # build extension rpm or deb
pig build ext pg_relusage; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_relusage; # install by extension name, for the current active PG version
pig ext install pg_relusage; # install via package alias, for the active PG version
pig ext install pg_relusage -v 18;   # install for PG 18
pig ext install pg_relusage -v 17;   # install for PG 17
pig ext install pg_relusage -v 16;   # install for PG 16
pig ext install pg_relusage -v 15;   # install for PG 15
pig ext install pg_relusage -v 14;   # install for PG 14
pig ext install pg_relusage -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_relusage;
```

