---
title: "pgaudit"
linkTitle: "pgaudit"
description: "provides auditing functionality"
weight: 7080
categories: ["SEC"]
width: full
---

provides auditing functionality


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7080** | {{< badge content="pgaudit" link="https://github.com/pgaudit/pgaudit" >}} | {{< ext "pgaudit" >}} | `17.1` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgauditlogtofile" >}} {{< ext "set_user" >}} {{< ext "pg_permissions" >}} {{< ext "pg_auth_mon" >}} {{< ext "pg_auditor" >}} {{< ext "safeupdate" >}} {{< ext "pg_drop_events" >}} {{< ext "table_log" >}} |

> [!Note] pg15=pgaudit17, pg14=pgaudit16 pg13=pgaudit15 pg12=pgaudit14


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgaudit" >}} | `17.1` | {{< bg "18" "pgaudit_18*" "red" >}} {{< bg "17" "pgaudit_17*" "green" >}} {{< bg "16" "pgaudit_16*" "green" >}} {{< bg "15" "pgaudit_15*" "green" >}} {{< bg "14" "pgaudit_14*" "green" >}} {{< bg "13" "pgaudit_13*" "green" >}} | `pgaudit_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgaudit" >}} | `17.1` | {{< bg "18" "postgresql-18-pgaudit" "red" >}} {{< bg "17" "postgresql-17-pgaudit" "green" >}} {{< bg "16" "postgresql-16-pgaudit" "green" >}} {{< bg "15" "postgresql-15-pgaudit" "green" >}} {{< bg "14" "postgresql-14-pgaudit" "green" >}} {{< bg "13" "postgresql-13-pgaudit" "green" >}} | `postgresql-$v-pgaudit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 18.0" "pgaudit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "pgaudit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 16.1" "pgaudit_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.1" "pgaudit17_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.3" "pgaudit16_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pgaudit15_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 18.0" "pgaudit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "pgaudit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 16.1" "pgaudit_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.1" "pgaudit17_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.3" "pgaudit16_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pgaudit15_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 18.0" "pgaudit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "pgaudit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 16.1" "pgaudit_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.1" "pgaudit17_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.3" "pgaudit16_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pgaudit15_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 18.0" "pgaudit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "pgaudit_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 16.1" "pgaudit_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.1" "pgaudit17_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.3" "pgaudit16_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pgaudit15_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 18.0" "pgaudit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "pgaudit_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "pgaudit_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "pgaudit17_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "pgaudit16_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "pgaudit15_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 18.0" "pgaudit_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "pgaudit_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "pgaudit_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "pgaudit17_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "pgaudit16_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "pgaudit15_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 18.0" "postgresql-18-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "postgresql-17-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "postgresql-16-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "postgresql-14-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-13-pgaudit : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 18.0" "postgresql-18-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "postgresql-17-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "postgresql-16-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "postgresql-14-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-13-pgaudit : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 18.0" "postgresql-18-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "postgresql-17-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "postgresql-16-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "postgresql-14-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-13-pgaudit : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 18.0" "postgresql-18-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "postgresql-17-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "postgresql-16-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "postgresql-14-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-13-pgaudit : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 18.0" "postgresql-18-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "postgresql-17-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "postgresql-16-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "postgresql-14-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-13-pgaudit : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 18.0" "postgresql-18-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "postgresql-17-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "postgresql-16-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "postgresql-14-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-13-pgaudit : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 18.0" "postgresql-18-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "postgresql-17-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "postgresql-16-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "postgresql-14-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-13-pgaudit : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 18.0" "postgresql-18-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 17.1" "postgresql-17-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 16.1" "postgresql-16-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.1" "postgresql-15-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.3" "postgresql-14-pgaudit : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-13-pgaudit : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit_18` | 18.0 | `el8.x86_64` | pgdg | 27.5 KiB | [pgaudit_18-18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgaudit_18-18.0-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_18` | 18.0 | `el8.aarch64` | pgdg | 27.0 KiB | [pgaudit_18-18.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgaudit_18-18.0-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_18` | 18.0 | `el9.x86_64` | pgdg | 27.8 KiB | [pgaudit_18-18.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgaudit_18-18.0-1PGDG.rhel9.x86_64.rpm) |
| `pgaudit_18` | 18.0 | `el9.aarch64` | pgdg | 27.4 KiB | [pgaudit_18-18.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgaudit_18-18.0-1PGDG.rhel9.aarch64.rpm) |
| `pgaudit_18` | 18.0 | `el10.x86_64` | pgdg | 28.1 KiB | [pgaudit_18-18.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgaudit_18-18.0-1PGDG.rhel10.x86_64.rpm) |
| `pgaudit_18` | 18.0 | `el10.aarch64` | pgdg | 28.2 KiB | [pgaudit_18-18.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgaudit_18-18.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgaudit` | 18.0 | `d12.x86_64` | pgdg | 47.1 KiB | [postgresql-18-pgaudit_18.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgaudit` | 18.0 | `d12.aarch64` | pgdg | 46.4 KiB | [postgresql-18-pgaudit_18.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgaudit` | 18.0 | `d13.x86_64` | pgdg | 47.0 KiB | [postgresql-18-pgaudit_18.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgaudit` | 18.0 | `d13.aarch64` | pgdg | 46.6 KiB | [postgresql-18-pgaudit_18.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgaudit` | 18.0 | `u22.x86_64` | pgdg | 48.7 KiB | [postgresql-18-pgaudit_18.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgaudit` | 18.0 | `u22.aarch64` | pgdg | 47.9 KiB | [postgresql-18-pgaudit_18.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgaudit` | 18.0 | `u24.x86_64` | pgdg | 47.3 KiB | [postgresql-18-pgaudit_18.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgaudit` | 18.0 | `u24.aarch64` | pgdg | 46.5 KiB | [postgresql-18-pgaudit_18.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-18/postgresql-18-pgaudit_18.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit_17` | 17.1 | `el8.x86_64` | pgdg | 28.0 KiB | [pgaudit_17-17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgaudit_17-17.1-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_17` | 17.0 | `el8.x86_64` | pgdg | 27.5 KiB | [pgaudit_17-17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgaudit_17-17.0-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_17` | 17.1 | `el8.aarch64` | pgdg | 27.6 KiB | [pgaudit_17-17.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgaudit_17-17.1-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_17` | 17.0 | `el8.aarch64` | pgdg | 27.2 KiB | [pgaudit_17-17.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgaudit_17-17.0-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_17` | 17.1 | `el9.x86_64` | pgdg | 28.2 KiB | [pgaudit_17-17.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgaudit_17-17.1-1PGDG.rhel9.x86_64.rpm) |
| `pgaudit_17` | 17.0 | `el9.x86_64` | pgdg | 27.8 KiB | [pgaudit_17-17.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgaudit_17-17.0-1PGDG.rhel9.x86_64.rpm) |
| `pgaudit_17` | 17.1 | `el9.aarch64` | pgdg | 28.0 KiB | [pgaudit_17-17.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgaudit_17-17.1-1PGDG.rhel9.aarch64.rpm) |
| `pgaudit_17` | 17.0 | `el9.aarch64` | pgdg | 27.6 KiB | [pgaudit_17-17.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgaudit_17-17.0-1PGDG.rhel9.aarch64.rpm) |
| `pgaudit_17` | 17.1 | `el10.x86_64` | pgdg | 28.6 KiB | [pgaudit_17-17.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgaudit_17-17.1-1PGDG.rhel10.x86_64.rpm) |
| `pgaudit_17` | 17.1 | `el10.aarch64` | pgdg | 28.7 KiB | [pgaudit_17-17.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgaudit_17-17.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgaudit` | 17.1 | `d12.x86_64` | pgdg | 46.0 KiB | [postgresql-17-pgaudit_17.1-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-1.pgdg120+1_amd64.deb) |
| `postgresql-17-pgaudit` | 17.1 | `d12.aarch64` | pgdg | 45.6 KiB | [postgresql-17-pgaudit_17.1-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-1.pgdg120+1_arm64.deb) |
| `postgresql-17-pgaudit` | 17.1 | `d13.x86_64` | pgdg | 46.0 KiB | [postgresql-17-pgaudit_17.1-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-1.pgdg130+1_amd64.deb) |
| `postgresql-17-pgaudit` | 17.1 | `d13.aarch64` | pgdg | 45.8 KiB | [postgresql-17-pgaudit_17.1-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-1.pgdg130+1_arm64.deb) |
| `postgresql-17-pgaudit` | 17.1 | `u22.x86_64` | pgdg | 52.5 KiB | [postgresql-17-pgaudit_17.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgaudit` | 17.1 | `u22.aarch64` | pgdg | 51.9 KiB | [postgresql-17-pgaudit_17.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgaudit` | 17.1 | `u24.x86_64` | pgdg | 46.1 KiB | [postgresql-17-pgaudit_17.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgaudit` | 17.1 | `u24.aarch64` | pgdg | 45.7 KiB | [postgresql-17-pgaudit_17.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-17/postgresql-17-pgaudit_17.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit_16` | 16.1 | `el8.x86_64` | pgdg | 27.4 KiB | [pgaudit_16-16.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgaudit_16-16.1-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_16` | 16.0 | `el8.x86_64` | pgdg | 26.9 KiB | [pgaudit_16-16.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgaudit_16-16.0-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit_16` | 16.1 | `el8.aarch64` | pgdg | 27.1 KiB | [pgaudit_16-16.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgaudit_16-16.1-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_16` | 16.0 | `el8.aarch64` | pgdg | 26.6 KiB | [pgaudit_16-16.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgaudit_16-16.0-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit_16` | 16.1 | `el9.x86_64` | pgdg | 27.8 KiB | [pgaudit_16-16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgaudit_16-16.1-1PGDG.rhel9.x86_64.rpm) |
| `pgaudit_16` | 16.0 | `el9.x86_64` | pgdg | 27.1 KiB | [pgaudit_16-16.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgaudit_16-16.0-1PGDG.rhel9.x86_64.rpm) |
| `pgaudit_16` | 16.1 | `el9.aarch64` | pgdg | 27.6 KiB | [pgaudit_16-16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgaudit_16-16.1-1PGDG.rhel9.aarch64.rpm) |
| `pgaudit_16` | 16.0 | `el9.aarch64` | pgdg | 26.7 KiB | [pgaudit_16-16.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgaudit_16-16.0-1PGDG.rhel9.aarch64.rpm) |
| `pgaudit_16` | 16.1 | `el10.x86_64` | pgdg | 28.2 KiB | [pgaudit_16-16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgaudit_16-16.1-1PGDG.rhel10.x86_64.rpm) |
| `pgaudit_16` | 16.1 | `el10.aarch64` | pgdg | 28.3 KiB | [pgaudit_16-16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgaudit_16-16.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgaudit` | 16.1 | `d12.x86_64` | pgdg | 45.0 KiB | [postgresql-16-pgaudit_16.1-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-1.pgdg120+1_amd64.deb) |
| `postgresql-16-pgaudit` | 16.1 | `d12.aarch64` | pgdg | 44.7 KiB | [postgresql-16-pgaudit_16.1-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-1.pgdg120+1_arm64.deb) |
| `postgresql-16-pgaudit` | 16.1 | `d13.x86_64` | pgdg | 45.0 KiB | [postgresql-16-pgaudit_16.1-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-1.pgdg130+1_amd64.deb) |
| `postgresql-16-pgaudit` | 16.1 | `d13.aarch64` | pgdg | 44.9 KiB | [postgresql-16-pgaudit_16.1-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-1.pgdg130+1_arm64.deb) |
| `postgresql-16-pgaudit` | 16.1 | `u22.x86_64` | pgdg | 51.0 KiB | [postgresql-16-pgaudit_16.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgaudit` | 16.1 | `u22.aarch64` | pgdg | 50.5 KiB | [postgresql-16-pgaudit_16.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgaudit` | 16.1 | `u24.x86_64` | pgdg | 45.1 KiB | [postgresql-16-pgaudit_16.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgaudit` | 16.1 | `u24.aarch64` | pgdg | 44.9 KiB | [postgresql-16-pgaudit_16.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-16/postgresql-16-pgaudit_16.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit17_15` | 1.7.1 | `el8.x86_64` | pgdg | 27.4 KiB | [pgaudit17_15-1.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgaudit17_15-1.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit17_15` | 1.7.0 | `el8.x86_64` | pgdg | 55.7 KiB | [pgaudit17_15-1.7.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgaudit17_15-1.7.0-1.rhel8.x86_64.rpm) |
| `pgaudit17_15` | 1.7 | `el8.x86_64` | pgdg | 55.6 KiB | [pgaudit17_15-1.7-beta1_1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgaudit17_15-1.7-beta1_1.rhel8.x86_64.rpm) |
| `pgaudit17_15` | 1.7.1 | `el8.aarch64` | pgdg | 27.1 KiB | [pgaudit17_15-1.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgaudit17_15-1.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit17_15` | 1.7.0 | `el8.aarch64` | pgdg | 55.2 KiB | [pgaudit17_15-1.7.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgaudit17_15-1.7.0-1.rhel8.aarch64.rpm) |
| `pgaudit17_15` | 1.7.1 | `el9.x86_64` | pgdg | 27.7 KiB | [pgaudit17_15-1.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgaudit17_15-1.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgaudit17_15` | 1.7.0 | `el9.x86_64` | pgdg | 57.0 KiB | [pgaudit17_15-1.7.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgaudit17_15-1.7.0-1.rhel9.x86_64.rpm) |
| `pgaudit17_15` | 1.7 | `el9.x86_64` | pgdg | 56.9 KiB | [pgaudit17_15-1.7-beta1_1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgaudit17_15-1.7-beta1_1.rhel9.x86_64.rpm) |
| `pgaudit17_15` | 1.7.1 | `el9.aarch64` | pgdg | 27.5 KiB | [pgaudit17_15-1.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgaudit17_15-1.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgaudit17_15` | 1.7.0 | `el9.aarch64` | pgdg | 56.2 KiB | [pgaudit17_15-1.7.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgaudit17_15-1.7.0-1.rhel9.aarch64.rpm) |
| `pgaudit17_15` | 1.7.1 | `el10.x86_64` | pgdg | 28.1 KiB | [pgaudit17_15-1.7.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgaudit17_15-1.7.1-1PGDG.rhel10.x86_64.rpm) |
| `pgaudit17_15` | 1.7.1 | `el10.aarch64` | pgdg | 28.3 KiB | [pgaudit17_15-1.7.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgaudit17_15-1.7.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgaudit` | 1.7.1 | `d12.x86_64` | pgdg | 43.9 KiB | [postgresql-15-pgaudit_1.7.1-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-1.pgdg120+1_amd64.deb) |
| `postgresql-15-pgaudit` | 1.7.1 | `d12.aarch64` | pgdg | 43.3 KiB | [postgresql-15-pgaudit_1.7.1-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-1.pgdg120+1_arm64.deb) |
| `postgresql-15-pgaudit` | 1.7.1 | `d13.x86_64` | pgdg | 43.9 KiB | [postgresql-15-pgaudit_1.7.1-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-1.pgdg130+1_amd64.deb) |
| `postgresql-15-pgaudit` | 1.7.1 | `d13.aarch64` | pgdg | 43.6 KiB | [postgresql-15-pgaudit_1.7.1-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-1.pgdg130+1_arm64.deb) |
| `postgresql-15-pgaudit` | 1.7.1 | `u22.x86_64` | pgdg | 50.1 KiB | [postgresql-15-pgaudit_1.7.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgaudit` | 1.7.1 | `u22.aarch64` | pgdg | 49.2 KiB | [postgresql-15-pgaudit_1.7.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgaudit` | 1.7.1 | `u24.x86_64` | pgdg | 44.0 KiB | [postgresql-15-pgaudit_1.7.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgaudit` | 1.7.1 | `u24.aarch64` | pgdg | 43.5 KiB | [postgresql-15-pgaudit_1.7.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.7/postgresql-15-pgaudit_1.7.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit16_14` | 1.6.3 | `el8.x86_64` | pgdg | 27.8 KiB | [pgaudit16_14-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgaudit16_14-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit16_14` | 1.6.2 | `el8.x86_64` | pgdg | 56.3 KiB | [pgaudit16_14-1.6.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgaudit16_14-1.6.2-1.rhel8.x86_64.rpm) |
| `pgaudit16_14` | 1.6.0 | `el8.x86_64` | pgdg | 55.0 KiB | [pgaudit16_14-1.6.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgaudit16_14-1.6.0-1.rhel8.x86_64.rpm) |
| `pgaudit16_14` | 1.6 | `el8.x86_64` | pgdg | 55.0 KiB | [pgaudit16_14-1.6-beta2_1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgaudit16_14-1.6-beta2_1.rhel8.x86_64.rpm) |
| `pgaudit16_14` | 1.6.3 | `el8.aarch64` | pgdg | 27.5 KiB | [pgaudit16_14-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgaudit16_14-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit16_14` | 1.6.2 | `el8.aarch64` | pgdg | 54.7 KiB | [pgaudit16_14-1.6.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgaudit16_14-1.6.2-1.rhel8.aarch64.rpm) |
| `pgaudit16_14` | 1.6.3 | `el9.x86_64` | pgdg | 28.1 KiB | [pgaudit16_14-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgaudit16_14-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgaudit16_14` | 1.6.2 | `el9.x86_64` | pgdg | 56.6 KiB | [pgaudit16_14-1.6.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgaudit16_14-1.6.2-1.rhel9.x86_64.rpm) |
| `pgaudit16_14` | 1.6.3 | `el9.aarch64` | pgdg | 27.9 KiB | [pgaudit16_14-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgaudit16_14-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgaudit16_14` | 1.6.2 | `el9.aarch64` | pgdg | 55.7 KiB | [pgaudit16_14-1.6.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgaudit16_14-1.6.2-1.rhel9.aarch64.rpm) |
| `pgaudit16_14` | 1.6.3 | `el10.x86_64` | pgdg | 28.5 KiB | [pgaudit16_14-1.6.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgaudit16_14-1.6.3-1PGDG.rhel10.x86_64.rpm) |
| `pgaudit16_14` | 1.6.3 | `el10.aarch64` | pgdg | 28.6 KiB | [pgaudit16_14-1.6.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgaudit16_14-1.6.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgaudit` | 1.6.3 | `d12.x86_64` | pgdg | 43.9 KiB | [postgresql-14-pgaudit_1.6.3-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.6/postgresql-14-pgaudit_1.6.3-1.pgdg120+1_amd64.deb) |
| `postgresql-14-pgaudit` | 1.6.3 | `d12.aarch64` | pgdg | 43.3 KiB | [postgresql-14-pgaudit_1.6.3-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.6/postgresql-14-pgaudit_1.6.3-1.pgdg120+1_arm64.deb) |
| `postgresql-14-pgaudit` | 1.6.3 | `d13.x86_64` | pgdg | 43.9 KiB | [postgresql-14-pgaudit_1.6.3-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.6/postgresql-14-pgaudit_1.6.3-1.pgdg130+1_amd64.deb) |
| `postgresql-14-pgaudit` | 1.6.3 | `d13.aarch64` | pgdg | 43.5 KiB | [postgresql-14-pgaudit_1.6.3-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.6/postgresql-14-pgaudit_1.6.3-1.pgdg130+1_arm64.deb) |
| `postgresql-14-pgaudit` | 1.6.3 | `u22.x86_64` | pgdg | 49.3 KiB | [postgresql-14-pgaudit_1.6.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.6/postgresql-14-pgaudit_1.6.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgaudit` | 1.6.3 | `u22.aarch64` | pgdg | 48.6 KiB | [postgresql-14-pgaudit_1.6.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.6/postgresql-14-pgaudit_1.6.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgaudit` | 1.6.3 | `u24.x86_64` | pgdg | 44.0 KiB | [postgresql-14-pgaudit_1.6.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.6/postgresql-14-pgaudit_1.6.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgaudit` | 1.6.3 | `u24.aarch64` | pgdg | 43.5 KiB | [postgresql-14-pgaudit_1.6.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.6/postgresql-14-pgaudit_1.6.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgaudit15_13` | 1.5.3 | `el8.x86_64` | pgdg | 26.9 KiB | [pgaudit15_13-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgaudit15_13-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pgaudit15_13` | 1.5.2 | `el8.x86_64` | pgdg | 53.7 KiB | [pgaudit15_13-1.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgaudit15_13-1.5.2-1.rhel8.x86_64.rpm) |
| `pgaudit15_13` | 1.5.0 | `el8.x86_64` | pgdg | 52.0 KiB | [pgaudit15_13-1.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgaudit15_13-1.5.0-1.rhel8.x86_64.rpm) |
| `pgaudit15_13` | 1.5.3 | `el8.aarch64` | pgdg | 26.8 KiB | [pgaudit15_13-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgaudit15_13-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pgaudit15_13` | 1.5.2 | `el8.aarch64` | pgdg | 52.1 KiB | [pgaudit15_13-1.5.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgaudit15_13-1.5.2-1.rhel8.aarch64.rpm) |
| `pgaudit15_13` | 1.5.3 | `el9.x86_64` | pgdg | 27.4 KiB | [pgaudit15_13-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgaudit15_13-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pgaudit15_13` | 1.5.2 | `el9.x86_64` | pgdg | 53.8 KiB | [pgaudit15_13-1.5.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgaudit15_13-1.5.2-1.rhel9.x86_64.rpm) |
| `pgaudit15_13` | 1.5.3 | `el9.aarch64` | pgdg | 27.1 KiB | [pgaudit15_13-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgaudit15_13-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pgaudit15_13` | 1.5.2 | `el9.aarch64` | pgdg | 53.3 KiB | [pgaudit15_13-1.5.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgaudit15_13-1.5.2-1.rhel9.aarch64.rpm) |
| `pgaudit15_13` | 1.5.3 | `el10.x86_64` | pgdg | 27.8 KiB | [pgaudit15_13-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgaudit15_13-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pgaudit15_13` | 1.5.3 | `el10.aarch64` | pgdg | 27.8 KiB | [pgaudit15_13-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgaudit15_13-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgaudit` | 1.5.3 | `d12.x86_64` | pgdg | 41.5 KiB | [postgresql-13-pgaudit_1.5.3-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.5/postgresql-13-pgaudit_1.5.3-1.pgdg120+1_amd64.deb) |
| `postgresql-13-pgaudit` | 1.5.3 | `d12.aarch64` | pgdg | 41.4 KiB | [postgresql-13-pgaudit_1.5.3-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.5/postgresql-13-pgaudit_1.5.3-1.pgdg120+1_arm64.deb) |
| `postgresql-13-pgaudit` | 1.5.3 | `d13.x86_64` | pgdg | 41.5 KiB | [postgresql-13-pgaudit_1.5.3-1.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.5/postgresql-13-pgaudit_1.5.3-1.pgdg130+1_amd64.deb) |
| `postgresql-13-pgaudit` | 1.5.3 | `d13.aarch64` | pgdg | 41.4 KiB | [postgresql-13-pgaudit_1.5.3-1.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.5/postgresql-13-pgaudit_1.5.3-1.pgdg130+1_arm64.deb) |
| `postgresql-13-pgaudit` | 1.5.3 | `u22.x86_64` | pgdg | 46.5 KiB | [postgresql-13-pgaudit_1.5.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.5/postgresql-13-pgaudit_1.5.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgaudit` | 1.5.3 | `u22.aarch64` | pgdg | 45.5 KiB | [postgresql-13-pgaudit_1.5.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.5/postgresql-13-pgaudit_1.5.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgaudit` | 1.5.3 | `u24.x86_64` | pgdg | 41.6 KiB | [postgresql-13-pgaudit_1.5.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.5/postgresql-13-pgaudit_1.5.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgaudit` | 1.5.3 | `u24.aarch64` | pgdg | 41.6 KiB | [postgresql-13-pgaudit_1.5.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgaudit-1.5/postgresql-13-pgaudit_1.5.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgaudit/pgaudit" title="Repository" icon="github" subtitle="github.com/pgaudit/pgaudit" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgaudit; # install by extension name, for the current active PG version
pig ext install pgaudit; # install via package alias, for the active PG version
pig ext install pgaudit -v 17;   # install for PG 17
pig ext install pgaudit -v 16;   # install for PG 16
pig ext install pgaudit -v 15;   # install for PG 15
pig ext install pgaudit -v 14;   # install for PG 14
pig ext install pgaudit -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgaudit;
```

