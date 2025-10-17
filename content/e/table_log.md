---
title: "table_log"
linkTitle: "table_log"
description: "record table modification logs and PITR for table/row"
weight: 5840
categories: ["Admin"]
width: full
---

record table modification logs and PITR for table/row

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5840** | {{< badge content="table_log" link="https://github.com/df7cb/table_log" >}} | {{< ext "table_log" "table_log" >}} | `0.6.4` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "pg_drop_events" >}} {{< ext "pg_auditor" >}} {{< ext "pg_upless" >}} {{< ext "pg_savior" >}} {{< ext "pgaudit" >}} {{< ext "pgauditlogtofile" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/table_log" >}} | `0.6.4` | {{< badge content="18" color="red" alt="table_log_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `table_log_$v` | - |
| **Debian** | {{< badge content="PGDG" link="/e/table_log" >}} | `0.6.4` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-tablelog` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "table_log_18" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_18-0.6.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "table_log_17" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_17-0.6.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "table_log_16" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_16-0.6.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "table_log_15" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_15-0.6.4-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "table_log_14" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_14-0.6.4-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "table_log_18" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_18-0.6.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "table_log_17" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_17-0.6.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "table_log_16" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_16-0.6.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "table_log_15" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_15-0.6.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "table_log_14" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_14-0.6.4-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "table_log_18" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_18-0.6.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "table_log_17" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_17-0.6.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "table_log_16" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_16-0.6.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "table_log_15" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_15-0.6.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "table_log_14" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_14-0.6.4-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "table_log_18" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_18-0.6.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "table_log_17" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_17-0.6.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "table_log_16" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_16-0.6.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "table_log_15" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_15-0.6.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "table_log_14" "0.6.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_14-0.6.4-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-tablelog" "0.6.4" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `table_log_18` | 0.6.4 | `el8.aarch64` | pigsty | 28.4 KiB | [table_log_18-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_18-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_18` | 0.6.4 | `el8.x86_64` | pigsty | 29.3 KiB | [table_log_18-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_18-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_18` | 0.6.4 | `el9.aarch64` | pigsty | 28.3 KiB | [table_log_18-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_18-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_18` | 0.6.4 | `el9.x86_64` | pigsty | 29.8 KiB | [table_log_18-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_18-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-18-tablelog` | 0.6.4 | `d12.aarch64` | pgdg | 43.6 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-18-tablelog` | 0.6.4 | `d12.x86_64` | pgdg | 45.3 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-18-tablelog` | 0.6.4 | `u22.x86_64` | pgdg | 45.9 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-tablelog` | 0.6.4 | `u22.aarch64` | pgdg | 44.0 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-tablelog` | 0.6.4 | `u24.x86_64` | pgdg | 45.3 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-tablelog` | 0.6.4 | `u24.aarch64` | pgdg | 43.7 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `table_log_17` | 0.6.4 | `el8.x86_64` | pigsty | 29.3 KiB | [table_log_17-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_17-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_17` | 0.6.4 | `el8.aarch64` | pigsty | 28.4 KiB | [table_log_17-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_17-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_17` | 0.6.4 | `el9.aarch64` | pigsty | 28.5 KiB | [table_log_17-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_17-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_17` | 0.6.4 | `el9.x86_64` | pigsty | 29.7 KiB | [table_log_17-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_17-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-tablelog` | 0.6.4 | `d12.x86_64` | pgdg | 45.2 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-17-tablelog` | 0.6.4 | `d12.aarch64` | pgdg | 43.6 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-17-tablelog` | 0.6.4 | `u22.aarch64` | pgdg | 48.8 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-tablelog` | 0.6.4 | `u22.x86_64` | pgdg | 50.6 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-tablelog` | 0.6.4 | `u24.aarch64` | pgdg | 43.6 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |
| `postgresql-17-tablelog` | 0.6.4 | `u24.x86_64` | pgdg | 45.1 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `table_log_16` | 0.6.4 | `el8.x86_64` | pigsty | 29.3 KiB | [table_log_16-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_16-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_16` | 0.6.4 | `el8.aarch64` | pigsty | 28.4 KiB | [table_log_16-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_16-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_16` | 0.6.4 | `el9.x86_64` | pigsty | 29.7 KiB | [table_log_16-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_16-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_16` | 0.6.4 | `el9.aarch64` | pigsty | 28.5 KiB | [table_log_16-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_16-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-tablelog` | 0.6.4 | `d12.aarch64` | pgdg | 43.6 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-16-tablelog` | 0.6.4 | `d12.x86_64` | pgdg | 45.2 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-16-tablelog` | 0.6.4 | `u22.aarch64` | pgdg | 48.3 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-tablelog` | 0.6.4 | `u22.x86_64` | pgdg | 50.1 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-tablelog` | 0.6.4 | `u24.x86_64` | pgdg | 45.2 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-tablelog` | 0.6.4 | `u24.aarch64` | pgdg | 43.7 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `table_log_15` | 0.6.4 | `el8.aarch64` | pigsty | 28.5 KiB | [table_log_15-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_15-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_15` | 0.6.4 | `el8.x86_64` | pigsty | 29.3 KiB | [table_log_15-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_15-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_15` | 0.6.4 | `el9.x86_64` | pigsty | 29.8 KiB | [table_log_15-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_15-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_15` | 0.6.4 | `el9.aarch64` | pigsty | 28.5 KiB | [table_log_15-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_15-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-tablelog` | 0.6.4 | `d12.x86_64` | pgdg | 45.2 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-15-tablelog` | 0.6.4 | `d12.aarch64` | pgdg | 43.6 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-15-tablelog` | 0.6.4 | `u22.aarch64` | pgdg | 48.3 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-tablelog` | 0.6.4 | `u22.x86_64` | pgdg | 50.1 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-tablelog` | 0.6.4 | `u24.x86_64` | pgdg | 45.2 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-tablelog` | 0.6.4 | `u24.aarch64` | pgdg | 43.6 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `table_log_14` | 0.6.4 | `el8.x86_64` | pigsty | 29.3 KiB | [table_log_14-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_14-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_14` | 0.6.4 | `el8.aarch64` | pigsty | 28.5 KiB | [table_log_14-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_14-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_14` | 0.6.4 | `el9.x86_64` | pigsty | 29.8 KiB | [table_log_14-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_14-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_14` | 0.6.4 | `el9.aarch64` | pigsty | 28.5 KiB | [table_log_14-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_14-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-tablelog` | 0.6.4 | `d12.x86_64` | pgdg | 45.2 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-14-tablelog` | 0.6.4 | `d12.aarch64` | pgdg | 43.6 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-14-tablelog` | 0.6.4 | `u22.x86_64` | pgdg | 48.5 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-tablelog` | 0.6.4 | `u22.aarch64` | pgdg | 46.6 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-tablelog` | 0.6.4 | `u24.aarch64` | pgdg | 43.6 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |
| `postgresql-14-tablelog` | 0.6.4 | `u24.x86_64` | pgdg | 45.2 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `table_log_13` | 0.6.4 | `el8.aarch64` | pigsty | 28.3 KiB | [table_log_13-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_13-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_13` | 0.6.4 | `el8.x86_64` | pigsty | 29.0 KiB | [table_log_13-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_13-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_13` | 0.6.4 | `el9.aarch64` | pigsty | 28.5 KiB | [table_log_13-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_13-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_13` | 0.6.4 | `el9.x86_64` | pigsty | 29.6 KiB | [table_log_13-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_13-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-tablelog` | 0.6.4 | `d12.aarch64` | pgdg | 43.4 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-13-tablelog` | 0.6.4 | `d12.x86_64` | pgdg | 44.9 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-13-tablelog` | 0.6.4 | `u22.aarch64` | pgdg | 46.5 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-13-tablelog` | 0.6.4 | `u22.x86_64` | pgdg | 48.0 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-13-tablelog` | 0.6.4 | `u24.x86_64` | pgdg | 44.6 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-13-tablelog` | 0.6.4 | `u24.aarch64` | pgdg | 43.5 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/table_log" title="Repository" icon="github" subtitle="github.com/df7cb/table_log" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="table_log-0.6.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get table_log; # get table_log source code
pig build dep table_log; # install build dependencies
pig build pkg table_log; # build extension rpm or deb
pig build ext table_log; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install table_log; # install by extension name, for the current active PG version
pig ext install table_log; # install via package alias, for the active PG version
pig ext install table_log -v 18;   # install for PG 18
pig ext install table_log -v 17;   # install for PG 17
pig ext install table_log -v 16;   # install for PG 16
pig ext install table_log -v 15;   # install for PG 15
pig ext install table_log -v 14;   # install for PG 14
pig ext install table_log -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION table_log;
```

