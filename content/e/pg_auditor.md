---
title: "pg_auditor"
linkTitle: "pg_auditor"
description: "Audit data changes and provide flashback ability"
weight: 7190
categories: ["SEC"]
width: full
---

Audit data changes and provide flashback ability


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7190** | {{< badge content="pg_auditor" link="https://github.com/kouber/pg_auditor" >}} | {{< ext "pg_auditor" >}} | `0.2` | {{< category "SEC" >}} | {{< license "BSD 3-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_drop_events" >}} {{< ext "table_log" >}} {{< ext "pgaudit" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "pg_savior" >}} {{< ext "pg_upless" >}} {{< ext "pgauditlogtofile" >}} |

> [!Note] pg15 rpm pkg name is pgaudit17_$v*


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_auditor" >}} | `0.2` | {{< bg "18" "pg_auditor_18" "green" >}} {{< bg "17" "pg_auditor_17" "green" >}} {{< bg "16" "pg_auditor_16" "green" >}} {{< bg "15" "pg_auditor_15" "green" >}} {{< bg "14" "pg_auditor_14" "green" >}} {{< bg "13" "pg_auditor_13" "green" >}} | `pg_auditor_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_auditor" >}} | `0.2` | {{< bg "18" "postgresql-18-pg-auditor" "red" >}} {{< bg "17" "postgresql-17-pg-auditor" "green" >}} {{< bg "16" "postgresql-16-pg-auditor" "green" >}} {{< bg "15" "postgresql-15-pg-auditor" "green" >}} {{< bg "14" "postgresql-14-pg-auditor" "green" >}} {{< bg "13" "postgresql-13-pg-auditor" "green" >}} | `postgresql-$v-pg-auditor` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 0.2" "pg_auditor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 0.2" "pg_auditor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 0.2" "pg_auditor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 0.2" "pg_auditor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "pg_auditor_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_auditor_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_auditor_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_auditor_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-auditor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2" "postgresql-17-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-16-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-15-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-14-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-13-pg-auditor : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-auditor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2" "postgresql-17-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-16-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-15-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-14-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-13-pg-auditor : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-auditor : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-auditor : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-auditor : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-auditor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2" "postgresql-17-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-16-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-15-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-14-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-13-pg-auditor : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-auditor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2" "postgresql-17-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-16-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-15-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-14-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-13-pg-auditor : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-auditor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2" "postgresql-17-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-16-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-15-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-14-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-13-pg-auditor : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-auditor : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.2" "postgresql-17-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-16-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-15-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-14-pg-auditor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2" "postgresql-13-pg-auditor : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auditor_18` | 0.2 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_auditor_18-0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_auditor_18-0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_auditor_18` | 0.2 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_auditor_18-0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_auditor_18-0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_auditor_18` | 0.2 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_auditor_18-0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_auditor_18-0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_auditor_18` | 0.2 | `el9.aarch64` | pigsty | 10.9 KiB | [pg_auditor_18-0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_auditor_18-0.2-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auditor_17` | 0.2 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_auditor_17-0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_auditor_17-0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_auditor_17` | 0.2 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_auditor_17-0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_auditor_17-0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_auditor_17` | 0.2 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_auditor_17-0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_auditor_17-0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_auditor_17` | 0.2 | `el9.aarch64` | pigsty | 10.9 KiB | [pg_auditor_17-0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_auditor_17-0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-auditor` | 0.2 | `d12.x86_64` | pigsty | 8.1 KiB | [postgresql-17-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-17-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-auditor` | 0.2 | `d12.aarch64` | pigsty | 8.1 KiB | [postgresql-17-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-17-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-auditor` | 0.2 | `u22.x86_64` | pigsty | 8.0 KiB | [postgresql-17-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-17-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-auditor` | 0.2 | `u22.aarch64` | pigsty | 8.0 KiB | [postgresql-17-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-17-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-auditor` | 0.2 | `u24.x86_64` | pigsty | 8.0 KiB | [postgresql-17-pg-auditor_0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-17-pg-auditor_0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-auditor` | 0.2 | `u24.aarch64` | pigsty | 8.0 KiB | [postgresql-17-pg-auditor_0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-17-pg-auditor_0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auditor_16` | 0.2 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_auditor_16-0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_auditor_16-0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_auditor_16` | 0.2 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_auditor_16-0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_auditor_16-0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_auditor_16` | 0.2 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_auditor_16-0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_auditor_16-0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_auditor_16` | 0.2 | `el9.aarch64` | pigsty | 10.9 KiB | [pg_auditor_16-0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_auditor_16-0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-auditor` | 0.2 | `d12.x86_64` | pigsty | 8.1 KiB | [postgresql-16-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-16-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-auditor` | 0.2 | `d12.aarch64` | pigsty | 8.1 KiB | [postgresql-16-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-16-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-auditor` | 0.2 | `u22.x86_64` | pigsty | 8.0 KiB | [postgresql-16-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-16-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-auditor` | 0.2 | `u22.aarch64` | pigsty | 8.0 KiB | [postgresql-16-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-16-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-auditor` | 0.2 | `u24.x86_64` | pigsty | 8.0 KiB | [postgresql-16-pg-auditor_0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-16-pg-auditor_0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-auditor` | 0.2 | `u24.aarch64` | pigsty | 8.0 KiB | [postgresql-16-pg-auditor_0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-16-pg-auditor_0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auditor_15` | 0.2 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_auditor_15-0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_auditor_15-0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_auditor_15` | 0.2 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_auditor_15-0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_auditor_15-0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_auditor_15` | 0.2 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_auditor_15-0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_auditor_15-0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_auditor_15` | 0.2 | `el9.aarch64` | pigsty | 10.9 KiB | [pg_auditor_15-0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_auditor_15-0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-auditor` | 0.2 | `d12.x86_64` | pigsty | 8.1 KiB | [postgresql-15-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-15-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-auditor` | 0.2 | `d12.aarch64` | pigsty | 8.1 KiB | [postgresql-15-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-15-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-auditor` | 0.2 | `u22.x86_64` | pigsty | 8.0 KiB | [postgresql-15-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-15-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-auditor` | 0.2 | `u22.aarch64` | pigsty | 8.0 KiB | [postgresql-15-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-15-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-auditor` | 0.2 | `u24.x86_64` | pigsty | 8.0 KiB | [postgresql-15-pg-auditor_0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-15-pg-auditor_0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-auditor` | 0.2 | `u24.aarch64` | pigsty | 8.0 KiB | [postgresql-15-pg-auditor_0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-15-pg-auditor_0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auditor_14` | 0.2 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_auditor_14-0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_auditor_14-0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_auditor_14` | 0.2 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_auditor_14-0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_auditor_14-0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_auditor_14` | 0.2 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_auditor_14-0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_auditor_14-0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_auditor_14` | 0.2 | `el9.aarch64` | pigsty | 10.9 KiB | [pg_auditor_14-0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_auditor_14-0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-auditor` | 0.2 | `d12.x86_64` | pigsty | 8.1 KiB | [postgresql-14-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-14-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-auditor` | 0.2 | `d12.aarch64` | pigsty | 8.1 KiB | [postgresql-14-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-14-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-auditor` | 0.2 | `u22.x86_64` | pigsty | 8.0 KiB | [postgresql-14-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-14-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-auditor` | 0.2 | `u22.aarch64` | pigsty | 8.0 KiB | [postgresql-14-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-14-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-auditor` | 0.2 | `u24.x86_64` | pigsty | 8.0 KiB | [postgresql-14-pg-auditor_0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-14-pg-auditor_0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-auditor` | 0.2 | `u24.aarch64` | pigsty | 8.0 KiB | [postgresql-14-pg-auditor_0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-14-pg-auditor_0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auditor_13` | 0.2 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_auditor_13-0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_auditor_13-0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_auditor_13` | 0.2 | `el8.aarch64` | pigsty | 10.9 KiB | [pg_auditor_13-0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_auditor_13-0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_auditor_13` | 0.2 | `el9.x86_64` | pigsty | 10.9 KiB | [pg_auditor_13-0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_auditor_13-0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_auditor_13` | 0.2 | `el9.aarch64` | pigsty | 10.9 KiB | [pg_auditor_13-0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_auditor_13-0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-auditor` | 0.2 | `d12.x86_64` | pigsty | 8.1 KiB | [postgresql-13-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-13-pg-auditor_0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-auditor` | 0.2 | `d12.aarch64` | pigsty | 8.1 KiB | [postgresql-13-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-auditor/postgresql-13-pg-auditor_0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-auditor` | 0.2 | `u22.x86_64` | pigsty | 8.0 KiB | [postgresql-13-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-13-pg-auditor_0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-auditor` | 0.2 | `u22.aarch64` | pigsty | 8.0 KiB | [postgresql-13-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-auditor/postgresql-13-pg-auditor_0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-auditor` | 0.2 | `u24.x86_64` | pigsty | 8.0 KiB | [postgresql-13-pg-auditor_0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-13-pg-auditor_0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-auditor` | 0.2 | `u24.aarch64` | pigsty | 8.0 KiB | [postgresql-13-pg-auditor_0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-auditor/postgresql-13-pg-auditor_0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kouber/pg_auditor" title="Repository" icon="github" subtitle="github.com/kouber/pg_auditor" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_auditor-0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_auditor; # get pg_auditor source code
pig build dep pg_auditor; # install build dependencies
pig build pkg pg_auditor; # build extension rpm or deb
pig build ext pg_auditor; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_auditor; # install by extension name, for the current active PG version
pig ext install pg_auditor; # install via package alias, for the active PG version
pig ext install pg_auditor -v 18;   # install for PG 18
pig ext install pg_auditor -v 17;   # install for PG 17
pig ext install pg_auditor -v 16;   # install for PG 16
pig ext install pg_auditor -v 15;   # install for PG 15
pig ext install pg_auditor -v 14;   # install for PG 14
pig ext install pg_auditor -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_auditor;
```

