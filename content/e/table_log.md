---
title: "table_log"
linkTitle: "table_log"
description: "record table modification logs and PITR for table/row"
weight: 5840
categories: ["ADMIN"]
width: full
---

[**table_log**](https://github.com/df7cb/table_log) : record table modification logs and PITR for table/row


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5840** | {{< badge content="table_log" link="https://github.com/df7cb/table_log" >}} | {{< ext "table_log" >}} | `0.6.4` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "pg_drop_events" >}} {{< ext "pg_auditor" >}} {{< ext "pg_upless" >}} {{< ext "pg_savior" >}} {{< ext "pgaudit" >}} {{< ext "pgauditlogtofile" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.6.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `table_log` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.4` | {{< bg "18" "table_log_18" "green" >}} {{< bg "17" "table_log_17" "green" >}} {{< bg "16" "table_log_16" "green" >}} {{< bg "15" "table_log_15" "green" >}} {{< bg "14" "table_log_14" "green" >}} {{< bg "13" "table_log_13" "green" >}} | `table_log_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.6.4` | {{< bg "18" "postgresql-18-tablelog" "green" >}} {{< bg "17" "postgresql-17-tablelog" "green" >}} {{< bg "16" "postgresql-16-tablelog" "green" >}} {{< bg "15" "postgresql-15-tablelog" "green" >}} {{< bg "14" "postgresql-14-tablelog" "green" >}} {{< bg "13" "postgresql-13-tablelog" "green" >}} | `postgresql-$v-tablelog` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.6.4" "table_log_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.6.4" "table_log_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.6.4" "table_log_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.6.4" "table_log_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.6.4" "table_log_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.6.4" "table_log_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "table_log_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.6.4" "postgresql-18-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-17-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-16-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-15-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-14-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-13-tablelog : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.6.4" "postgresql-18-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-17-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-16-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-15-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-14-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-13-tablelog : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.6.4" "postgresql-18-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-17-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-16-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-15-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-14-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-13-tablelog : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.6.4" "postgresql-18-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-17-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-16-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-15-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-14-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-13-tablelog : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.6.4" "postgresql-18-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-17-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-16-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-15-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-14-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-13-tablelog : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.6.4" "postgresql-18-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-17-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-16-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-15-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-14-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-13-tablelog : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.6.4" "postgresql-18-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-17-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-16-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-15-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-14-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-13-tablelog : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.6.4" "postgresql-18-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-17-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-16-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-15-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-14-tablelog : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.4" "postgresql-13-tablelog : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_log_18` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.6 KiB | [table_log_18-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_18-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_18` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.8 KiB | [table_log_18-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_18-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_18` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [table_log_18-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_18-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_18` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [table_log_18-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_18-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_18` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [table_log_18-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_log_18-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `table_log_18` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.0 KiB | [table_log_18-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_log_18-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-tablelog` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.3 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-18-tablelog` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.6 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-18-tablelog` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.2 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg13+1_amd64.deb) |
| `postgresql-18-tablelog` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.7 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg13+1_arm64.deb) |
| `postgresql-18-tablelog` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 45.9 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-tablelog` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 44.0 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-tablelog` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.3 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-tablelog` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.7 KiB | [postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-18-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_log_17` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [table_log_17-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_17-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_17` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.7 KiB | [table_log_17-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_17-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_17` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.8 KiB | [table_log_17-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_17-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_17` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [table_log_17-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_17-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_17` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [table_log_17-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_log_17-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `table_log_17` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.0 KiB | [table_log_17-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_log_17-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-tablelog` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.2 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-17-tablelog` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.6 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-17-tablelog` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.1 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg13+1_amd64.deb) |
| `postgresql-17-tablelog` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.7 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg13+1_arm64.deb) |
| `postgresql-17-tablelog` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.6 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-tablelog` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.8 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-tablelog` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.1 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-tablelog` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.6 KiB | [postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-17-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_log_16` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.5 KiB | [table_log_16-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_16-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_16` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.8 KiB | [table_log_16-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_16-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_16` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.8 KiB | [table_log_16-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_16-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_16` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [table_log_16-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_16-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_16` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [table_log_16-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_log_16-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `table_log_16` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.0 KiB | [table_log_16-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_log_16-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-tablelog` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.2 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-16-tablelog` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.6 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-16-tablelog` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.1 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg13+1_amd64.deb) |
| `postgresql-16-tablelog` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.7 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg13+1_arm64.deb) |
| `postgresql-16-tablelog` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.1 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-tablelog` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.3 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-tablelog` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.2 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-tablelog` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.7 KiB | [postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-16-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_log_15` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.6 KiB | [table_log_15-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_15-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_15` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.8 KiB | [table_log_15-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_15-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_15` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [table_log_15-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_15-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_15` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [table_log_15-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_15-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_15` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [table_log_15-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_log_15-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `table_log_15` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.1 KiB | [table_log_15-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_log_15-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-tablelog` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.2 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-15-tablelog` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.6 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-15-tablelog` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.2 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg13+1_amd64.deb) |
| `postgresql-15-tablelog` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.7 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg13+1_arm64.deb) |
| `postgresql-15-tablelog` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.1 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-tablelog` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 48.3 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-tablelog` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.2 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-tablelog` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.6 KiB | [postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-15-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_log_14` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.6 KiB | [table_log_14-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_14-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_14` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.8 KiB | [table_log_14-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_14-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_14` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.9 KiB | [table_log_14-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_14-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_14` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [table_log_14-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_14-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_14` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.8 KiB | [table_log_14-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_log_14-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `table_log_14` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.1 KiB | [table_log_14-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_log_14-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-tablelog` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.2 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-14-tablelog` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.6 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-14-tablelog` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.1 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg13+1_amd64.deb) |
| `postgresql-14-tablelog` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.7 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg13+1_arm64.deb) |
| `postgresql-14-tablelog` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 48.5 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-tablelog` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 46.6 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-tablelog` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.2 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-tablelog` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.6 KiB | [postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-14-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `table_log_13` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.2 KiB | [table_log_13-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/table_log_13-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `table_log_13` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.6 KiB | [table_log_13-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/table_log_13-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `table_log_13` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.7 KiB | [table_log_13-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/table_log_13-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `table_log_13` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [table_log_13-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/table_log_13-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `table_log_13` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.7 KiB | [table_log_13-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/table_log_13-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `table_log_13` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.0 KiB | [table_log_13-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/table_log_13-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-tablelog` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 44.9 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg12+1_amd64.deb) |
| `postgresql-13-tablelog` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 43.4 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg12+1_arm64.deb) |
| `postgresql-13-tablelog` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 44.7 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg13+1_amd64.deb) |
| `postgresql-13-tablelog` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 43.5 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg13+1_arm64.deb) |
| `postgresql-13-tablelog` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 48.0 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg22.04+1_amd64.deb) |
| `postgresql-13-tablelog` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 46.5 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg22.04+1_arm64.deb) |
| `postgresql-13-tablelog` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 44.6 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg24.04+1_amd64.deb) |
| `postgresql-13-tablelog` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 43.5 KiB | [postgresql-13-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tablelog/postgresql-13-tablelog_0.6.4-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/table_log" title="Repository" icon="github" subtitle="github.com/df7cb/table_log" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="table_log-0.6.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg table_log;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install table_log;		# install via package name, for the active PG version

pig install table_log -v 18;   # install for PG 18
pig install table_log -v 17;   # install for PG 17
pig install table_log -v 16;   # install for PG 16
pig install table_log -v 15;   # install for PG 15
pig install table_log -v 14;   # install for PG 14
pig install table_log -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION table_log;
```
