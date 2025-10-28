---
title: "schedoc"
linkTitle: "schedoc"
description: "Cross documentation between Django and DBT projects"
weight: 4330
categories: ["UTIL"]
width: full
---

Cross documentation between Django and DBT projects


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4330** | {{< badge content="schedoc" link="https://github.com/ZeroGachis/pg_schedoc" >}} | {{< ext "schedoc" "pg_schedoc" >}} | `0.0.1` | {{< category "UTIL" >}} | {{< license "GPL-3.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "ddl_historization" >}} |
|   **See Also**    | {{< ext "pg_readme_test_extension" >}} {{< ext "pg_readme" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/schedoc" >}} | `0.0.1` | {{< bg "18" "pg_schedoc_18" "green" >}} {{< bg "17" "pg_schedoc_17" "green" >}} {{< bg "16" "pg_schedoc_16" "green" >}} {{< bg "15" "pg_schedoc_15" "green" >}} {{< bg "14" "pg_schedoc_14" "green" >}} {{< bg "13" "pg_schedoc_13" "green" >}} | `pg_schedoc_$v` | `ddl_historization_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/schedoc" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-schedoc" "green" >}} {{< bg "17" "postgresql-17-pg-schedoc" "green" >}} {{< bg "16" "postgresql-16-pg-schedoc" "green" >}} {{< bg "15" "postgresql-15-pg-schedoc" "green" >}} {{< bg "14" "postgresql-14-pg-schedoc" "green" >}} {{< bg "13" "postgresql-13-pg-schedoc" "green" >}} | `postgresql-$v-pg-schedoc` | `postgresql-$v-ddl-historization` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_schedoc_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_schedoc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_schedoc_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_schedoc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_schedoc_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_schedoc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_schedoc_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_schedoc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_schedoc_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_schedoc_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_schedoc_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_schedoc_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-schedoc : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-schedoc : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-schedoc : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-schedoc : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-schedoc : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-schedoc : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-schedoc : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-schedoc : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-schedoc : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-schedoc : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-schedoc : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-schedoc : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-schedoc : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-schedoc : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-schedoc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-schedoc : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_schedoc_17` | 0.0.1 | `el8.x86_64` | pigsty | 22.9 KiB | [pg_schedoc_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_schedoc_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_schedoc_17` | 0.0.1 | `el8.aarch64` | pigsty | 22.9 KiB | [pg_schedoc_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_schedoc_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_schedoc_17` | 0.0.1 | `el9.x86_64` | pigsty | 22.5 KiB | [pg_schedoc_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_schedoc_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_schedoc_17` | 0.0.1 | `el9.aarch64` | pigsty | 22.5 KiB | [pg_schedoc_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_schedoc_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-schedoc` | 0.0.1 | `d12.x86_64` | pigsty | 4.5 KiB | [postgresql-17-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-17-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-schedoc` | 0.0.1 | `d12.aarch64` | pigsty | 4.5 KiB | [postgresql-17-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-17-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-schedoc` | 0.0.1 | `u22.x86_64` | pigsty | 4.4 KiB | [postgresql-17-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-17-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-schedoc` | 0.0.1 | `u22.aarch64` | pigsty | 4.4 KiB | [postgresql-17-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-17-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-schedoc` | 0.0.1 | `u24.x86_64` | pigsty | 4.4 KiB | [postgresql-17-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-17-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-schedoc` | 0.0.1 | `u24.aarch64` | pigsty | 4.4 KiB | [postgresql-17-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-17-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_schedoc_16` | 0.0.1 | `el8.x86_64` | pigsty | 22.9 KiB | [pg_schedoc_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_schedoc_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_schedoc_16` | 0.0.1 | `el8.aarch64` | pigsty | 22.9 KiB | [pg_schedoc_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_schedoc_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_schedoc_16` | 0.0.1 | `el9.x86_64` | pigsty | 22.5 KiB | [pg_schedoc_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_schedoc_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_schedoc_16` | 0.0.1 | `el9.aarch64` | pigsty | 22.5 KiB | [pg_schedoc_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_schedoc_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-schedoc` | 0.0.1 | `d12.x86_64` | pigsty | 4.5 KiB | [postgresql-16-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-16-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-schedoc` | 0.0.1 | `d12.aarch64` | pigsty | 4.5 KiB | [postgresql-16-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-16-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-schedoc` | 0.0.1 | `u22.x86_64` | pigsty | 4.4 KiB | [postgresql-16-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-16-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-schedoc` | 0.0.1 | `u22.aarch64` | pigsty | 4.4 KiB | [postgresql-16-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-16-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-schedoc` | 0.0.1 | `u24.x86_64` | pigsty | 4.4 KiB | [postgresql-16-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-16-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-schedoc` | 0.0.1 | `u24.aarch64` | pigsty | 4.4 KiB | [postgresql-16-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-16-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_schedoc_15` | 0.0.1 | `el8.x86_64` | pigsty | 22.9 KiB | [pg_schedoc_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_schedoc_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_schedoc_15` | 0.0.1 | `el8.aarch64` | pigsty | 22.9 KiB | [pg_schedoc_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_schedoc_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_schedoc_15` | 0.0.1 | `el9.x86_64` | pigsty | 22.5 KiB | [pg_schedoc_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_schedoc_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_schedoc_15` | 0.0.1 | `el9.aarch64` | pigsty | 22.5 KiB | [pg_schedoc_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_schedoc_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-schedoc` | 0.0.1 | `d12.x86_64` | pigsty | 4.5 KiB | [postgresql-15-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-15-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-schedoc` | 0.0.1 | `d12.aarch64` | pigsty | 4.5 KiB | [postgresql-15-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-15-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-schedoc` | 0.0.1 | `u22.x86_64` | pigsty | 4.4 KiB | [postgresql-15-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-15-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-schedoc` | 0.0.1 | `u22.aarch64` | pigsty | 4.4 KiB | [postgresql-15-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-15-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-schedoc` | 0.0.1 | `u24.x86_64` | pigsty | 4.4 KiB | [postgresql-15-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-15-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-schedoc` | 0.0.1 | `u24.aarch64` | pigsty | 4.4 KiB | [postgresql-15-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-15-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_schedoc_14` | 0.0.1 | `el8.x86_64` | pigsty | 22.9 KiB | [pg_schedoc_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_schedoc_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_schedoc_14` | 0.0.1 | `el8.aarch64` | pigsty | 22.9 KiB | [pg_schedoc_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_schedoc_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_schedoc_14` | 0.0.1 | `el9.x86_64` | pigsty | 22.5 KiB | [pg_schedoc_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_schedoc_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_schedoc_14` | 0.0.1 | `el9.aarch64` | pigsty | 22.5 KiB | [pg_schedoc_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_schedoc_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-schedoc` | 0.0.1 | `d12.x86_64` | pigsty | 4.5 KiB | [postgresql-14-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-14-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-schedoc` | 0.0.1 | `d12.aarch64` | pigsty | 4.5 KiB | [postgresql-14-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-14-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-schedoc` | 0.0.1 | `u22.x86_64` | pigsty | 4.4 KiB | [postgresql-14-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-14-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-schedoc` | 0.0.1 | `u22.aarch64` | pigsty | 4.4 KiB | [postgresql-14-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-14-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-schedoc` | 0.0.1 | `u24.x86_64` | pigsty | 4.4 KiB | [postgresql-14-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-14-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-schedoc` | 0.0.1 | `u24.aarch64` | pigsty | 4.4 KiB | [postgresql-14-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-14-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_schedoc_13` | 0.0.1 | `el8.x86_64` | pigsty | 22.9 KiB | [pg_schedoc_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_schedoc_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_schedoc_13` | 0.0.1 | `el8.aarch64` | pigsty | 22.9 KiB | [pg_schedoc_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_schedoc_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_schedoc_13` | 0.0.1 | `el9.x86_64` | pigsty | 22.5 KiB | [pg_schedoc_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_schedoc_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_schedoc_13` | 0.0.1 | `el9.aarch64` | pigsty | 22.5 KiB | [pg_schedoc_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_schedoc_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-schedoc` | 0.0.1 | `d12.x86_64` | pigsty | 4.5 KiB | [postgresql-13-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-13-pg-schedoc_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-schedoc` | 0.0.1 | `d12.aarch64` | pigsty | 4.5 KiB | [postgresql-13-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-schedoc/postgresql-13-pg-schedoc_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-schedoc` | 0.0.1 | `u22.x86_64` | pigsty | 4.4 KiB | [postgresql-13-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-13-pg-schedoc_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-schedoc` | 0.0.1 | `u22.aarch64` | pigsty | 4.4 KiB | [postgresql-13-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-schedoc/postgresql-13-pg-schedoc_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-schedoc` | 0.0.1 | `u24.x86_64` | pigsty | 4.4 KiB | [postgresql-13-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-13-pg-schedoc_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-schedoc` | 0.0.1 | `u24.aarch64` | pigsty | 4.4 KiB | [postgresql-13-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-schedoc/postgresql-13-pg-schedoc_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ZeroGachis/pg_schedoc" title="Repository" icon="github" subtitle="github.com/ZeroGachis/pg_schedoc" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_schedoc-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get schedoc; # get schedoc source code
pig build dep schedoc; # install build dependencies
pig build pkg schedoc; # build extension rpm or deb
pig build ext schedoc; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install schedoc; # install by extension name, for the current active PG version
pig ext install pg_schedoc; # install via package alias, for the active PG version
pig ext install schedoc -v 18;   # install for PG 18
pig ext install schedoc -v 17;   # install for PG 17
pig ext install schedoc -v 16;   # install for PG 16
pig ext install schedoc -v 15;   # install for PG 15
pig ext install schedoc -v 14;   # install for PG 14
pig ext install schedoc -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION schedoc;
```

