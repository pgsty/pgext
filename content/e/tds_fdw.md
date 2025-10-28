---
title: "tds_fdw"
linkTitle: "tds_fdw"
description: "Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server)"
weight: 8620
categories: ["FDW"]
width: full
---

Foreign data wrapper for querying a TDS database (Sybase or Microsoft SQL Server)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8620** | {{< badge content="tds_fdw" link="https://github.com/tds-fdw/tds_fdw" >}} | {{< ext "tds_fdw" >}} | `2.0.4` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "wrappers" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "db2_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/tds_fdw" >}} | `2.0.4` | {{< bg "18" "tds_fdw_18*" "red" >}} {{< bg "17" "tds_fdw_17*" "green" >}} {{< bg "16" "tds_fdw_16*" "green" >}} {{< bg "15" "tds_fdw_15*" "green" >}} {{< bg "14" "tds_fdw_14*" "green" >}} {{< bg "13" "tds_fdw_13*" "green" >}} | `tds_fdw_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/tds_fdw" >}} | `2.0.4` | {{< bg "18" "postgresql-18-tds-fdw" "red" >}} {{< bg "17" "postgresql-17-tds-fdw" "green" >}} {{< bg "16" "postgresql-16-tds-fdw" "green" >}} {{< bg "15" "postgresql-15-tds-fdw" "green" >}} {{< bg "14" "postgresql-14-tds-fdw" "green" >}} {{< bg "13" "postgresql-13-tds-fdw" "green" >}} | `postgresql-$v-tds-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 2.0.5" "tds_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_13 : AVAIL 3" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 2.0.5" "tds_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_13 : AVAIL 3" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 2.0.5" "tds_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_13 : AVAIL 3" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 2.0.5" "tds_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_13 : AVAIL 3" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 2.0.5" "tds_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_13 : AVAIL 2" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 2.0.5" "tds_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.5" "tds_fdw_13 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 2.0.5" "postgresql-18-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-17-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-16-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-15-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-14-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-13-tds-fdw : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 2.0.5" "postgresql-18-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-17-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-16-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-15-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-14-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-13-tds-fdw : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 2.0.5" "postgresql-18-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-17-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-16-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-15-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-14-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-13-tds-fdw : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 2.0.5" "postgresql-18-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-17-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-16-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-15-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-14-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-13-tds-fdw : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 2.0.5" "postgresql-18-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-17-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-16-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-15-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-14-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-13-tds-fdw : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 2.0.5" "postgresql-18-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-17-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-16-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-15-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-14-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-13-tds-fdw : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 2.0.5" "postgresql-18-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-17-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-16-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-15-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-14-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-13-tds-fdw : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 2.0.5" "postgresql-18-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-17-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-16-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-15-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-14-tds-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.5" "postgresql-13-tds-fdw : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tds_fdw_18` | 2.0.5 | `el8.x86_64` | pgdg | 49.7 KiB | [tds_fdw_18-2.0.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/tds_fdw_18-2.0.5-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_18` | 2.0.5 | `el8.aarch64` | pgdg | 47.1 KiB | [tds_fdw_18-2.0.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/tds_fdw_18-2.0.5-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_18` | 2.0.5 | `el9.x86_64` | pgdg | 47.8 KiB | [tds_fdw_18-2.0.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/tds_fdw_18-2.0.5-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_18` | 2.0.5 | `el9.aarch64` | pgdg | 45.7 KiB | [tds_fdw_18-2.0.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/tds_fdw_18-2.0.5-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_18` | 2.0.5 | `el10.x86_64` | pgdg | 48.4 KiB | [tds_fdw_18-2.0.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/tds_fdw_18-2.0.5-1PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_18` | 2.0.5 | `el10.aarch64` | pgdg | 47.3 KiB | [tds_fdw_18-2.0.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/tds_fdw_18-2.0.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-tds-fdw` | 2.0.5 | `d12.x86_64` | pgdg | 111.7 KiB | [postgresql-18-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-18-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-18-tds-fdw` | 2.0.5 | `d12.aarch64` | pgdg | 109.0 KiB | [postgresql-18-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-18-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-18-tds-fdw` | 2.0.5 | `d13.x86_64` | pgdg | 111.8 KiB | [postgresql-18-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-18-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-18-tds-fdw` | 2.0.5 | `d13.aarch64` | pgdg | 109.3 KiB | [postgresql-18-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-18-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-18-tds-fdw` | 2.0.5 | `u22.x86_64` | pgdg | 112.5 KiB | [postgresql-18-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-18-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-tds-fdw` | 2.0.5 | `u22.aarch64` | pgdg | 109.3 KiB | [postgresql-18-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-18-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-tds-fdw` | 2.0.5 | `u24.x86_64` | pgdg | 109.9 KiB | [postgresql-18-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-18-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-tds-fdw` | 2.0.5 | `u24.aarch64` | pgdg | 107.2 KiB | [postgresql-18-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-18-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tds_fdw_17` | 2.0.5 | `el8.x86_64` | pgdg | 49.7 KiB | [tds_fdw_17-2.0.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tds_fdw_17-2.0.5-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_17` | 2.0.4 | `el8.x86_64` | pgdg | 49.3 KiB | [tds_fdw_17-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tds_fdw_17-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_17` | 2.0.5 | `el8.aarch64` | pgdg | 47.1 KiB | [tds_fdw_17-2.0.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tds_fdw_17-2.0.5-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_17` | 2.0.4 | `el8.aarch64` | pgdg | 46.6 KiB | [tds_fdw_17-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tds_fdw_17-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_17` | 2.0.5 | `el9.x86_64` | pgdg | 47.8 KiB | [tds_fdw_17-2.0.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tds_fdw_17-2.0.5-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_17` | 2.0.4 | `el9.x86_64` | pgdg | 47.8 KiB | [tds_fdw_17-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tds_fdw_17-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_17` | 2.0.5 | `el9.aarch64` | pgdg | 45.7 KiB | [tds_fdw_17-2.0.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tds_fdw_17-2.0.5-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_17` | 2.0.4 | `el9.aarch64` | pgdg | 45.9 KiB | [tds_fdw_17-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tds_fdw_17-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_17` | 2.0.5 | `el10.x86_64` | pgdg | 48.5 KiB | [tds_fdw_17-2.0.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tds_fdw_17-2.0.5-1PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_17` | 2.0.4 | `el10.x86_64` | pgdg | 48.2 KiB | [tds_fdw_17-2.0.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tds_fdw_17-2.0.4-2PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_17` | 2.0.5 | `el10.aarch64` | pgdg | 47.1 KiB | [tds_fdw_17-2.0.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tds_fdw_17-2.0.5-1PGDG.rhel10.aarch64.rpm) |
| `tds_fdw_17` | 2.0.4 | `el10.aarch64` | pgdg | 46.8 KiB | [tds_fdw_17-2.0.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tds_fdw_17-2.0.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-tds-fdw` | 2.0.5 | `d12.x86_64` | pgdg | 111.6 KiB | [postgresql-17-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-17-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-17-tds-fdw` | 2.0.5 | `d12.aarch64` | pgdg | 108.8 KiB | [postgresql-17-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-17-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-17-tds-fdw` | 2.0.5 | `d13.x86_64` | pgdg | 111.4 KiB | [postgresql-17-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-17-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-17-tds-fdw` | 2.0.5 | `d13.aarch64` | pgdg | 109.3 KiB | [postgresql-17-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-17-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-17-tds-fdw` | 2.0.5 | `u22.x86_64` | pgdg | 127.2 KiB | [postgresql-17-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-17-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-tds-fdw` | 2.0.5 | `u22.aarch64` | pgdg | 124.0 KiB | [postgresql-17-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-17-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-tds-fdw` | 2.0.5 | `u24.x86_64` | pgdg | 109.7 KiB | [postgresql-17-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-17-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-tds-fdw` | 2.0.5 | `u24.aarch64` | pgdg | 106.9 KiB | [postgresql-17-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-17-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tds_fdw_16` | 2.0.5 | `el8.x86_64` | pgdg | 49.7 KiB | [tds_fdw_16-2.0.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tds_fdw_16-2.0.5-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_16` | 2.0.4 | `el8.x86_64` | pgdg | 49.3 KiB | [tds_fdw_16-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tds_fdw_16-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_16` | 2.0.3 | `el8.x86_64` | pgdg | 47.4 KiB | [tds_fdw_16-2.0.3-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tds_fdw_16-2.0.3-4PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_16` | 2.0.5 | `el8.aarch64` | pgdg | 47.2 KiB | [tds_fdw_16-2.0.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tds_fdw_16-2.0.5-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_16` | 2.0.4 | `el8.aarch64` | pgdg | 46.7 KiB | [tds_fdw_16-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tds_fdw_16-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_16` | 2.0.3 | `el8.aarch64` | pgdg | 44.8 KiB | [tds_fdw_16-2.0.3-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tds_fdw_16-2.0.3-4PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_16` | 2.0.3 | `el8.aarch64` | pgdg | 44.7 KiB | [tds_fdw_16-2.0.3-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tds_fdw_16-2.0.3-3.rhel8.aarch64.rpm) |
| `tds_fdw_16` | 2.0.5 | `el9.x86_64` | pgdg | 47.8 KiB | [tds_fdw_16-2.0.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tds_fdw_16-2.0.5-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_16` | 2.0.4 | `el9.x86_64` | pgdg | 47.8 KiB | [tds_fdw_16-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tds_fdw_16-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_16` | 2.0.3 | `el9.x86_64` | pgdg | 45.7 KiB | [tds_fdw_16-2.0.3-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tds_fdw_16-2.0.3-4PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_16` | 2.0.5 | `el9.aarch64` | pgdg | 45.7 KiB | [tds_fdw_16-2.0.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tds_fdw_16-2.0.5-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_16` | 2.0.4 | `el9.aarch64` | pgdg | 45.9 KiB | [tds_fdw_16-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tds_fdw_16-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_16` | 2.0.3 | `el9.aarch64` | pgdg | 43.6 KiB | [tds_fdw_16-2.0.3-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tds_fdw_16-2.0.3-3.rhel9.aarch64.rpm) |
| `tds_fdw_16` | 2.0.3 | `el9.aarch64` | pgdg | 43.7 KiB | [tds_fdw_16-2.0.3-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tds_fdw_16-2.0.3-4PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_16` | 2.0.5 | `el10.x86_64` | pgdg | 48.5 KiB | [tds_fdw_16-2.0.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tds_fdw_16-2.0.5-1PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_16` | 2.0.4 | `el10.x86_64` | pgdg | 48.2 KiB | [tds_fdw_16-2.0.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tds_fdw_16-2.0.4-2PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_16` | 2.0.5 | `el10.aarch64` | pgdg | 47.8 KiB | [tds_fdw_16-2.0.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tds_fdw_16-2.0.5-1PGDG.rhel10.aarch64.rpm) |
| `tds_fdw_16` | 2.0.4 | `el10.aarch64` | pgdg | 46.9 KiB | [tds_fdw_16-2.0.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tds_fdw_16-2.0.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-tds-fdw` | 2.0.5 | `d12.x86_64` | pgdg | 111.4 KiB | [postgresql-16-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-16-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-16-tds-fdw` | 2.0.5 | `d12.aarch64` | pgdg | 108.8 KiB | [postgresql-16-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-16-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-16-tds-fdw` | 2.0.5 | `d13.x86_64` | pgdg | 111.1 KiB | [postgresql-16-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-16-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-16-tds-fdw` | 2.0.5 | `d13.aarch64` | pgdg | 109.1 KiB | [postgresql-16-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-16-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-16-tds-fdw` | 2.0.5 | `u22.x86_64` | pgdg | 126.1 KiB | [postgresql-16-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-16-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-tds-fdw` | 2.0.5 | `u22.aarch64` | pgdg | 122.6 KiB | [postgresql-16-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-16-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-tds-fdw` | 2.0.5 | `u24.x86_64` | pgdg | 109.7 KiB | [postgresql-16-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-16-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-tds-fdw` | 2.0.5 | `u24.aarch64` | pgdg | 107.0 KiB | [postgresql-16-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-16-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tds_fdw_15` | 2.0.5 | `el8.x86_64` | pgdg | 49.7 KiB | [tds_fdw_15-2.0.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tds_fdw_15-2.0.5-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_15` | 2.0.4 | `el8.x86_64` | pgdg | 49.3 KiB | [tds_fdw_15-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tds_fdw_15-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_15` | 2.0.3 | `el8.x86_64` | pgdg | 47.1 KiB | [tds_fdw_15-2.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tds_fdw_15-2.0.3-1.rhel8.x86_64.rpm) |
| `tds_fdw_15` | 2.0.5 | `el8.aarch64` | pgdg | 47.2 KiB | [tds_fdw_15-2.0.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tds_fdw_15-2.0.5-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_15` | 2.0.4 | `el8.aarch64` | pgdg | 46.7 KiB | [tds_fdw_15-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tds_fdw_15-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_15` | 2.0.3 | `el8.aarch64` | pgdg | 44.6 KiB | [tds_fdw_15-2.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tds_fdw_15-2.0.3-1.rhel8.aarch64.rpm) |
| `tds_fdw_15` | 2.0.5 | `el9.x86_64` | pgdg | 47.8 KiB | [tds_fdw_15-2.0.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tds_fdw_15-2.0.5-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_15` | 2.0.4 | `el9.x86_64` | pgdg | 47.7 KiB | [tds_fdw_15-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tds_fdw_15-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_15` | 2.0.3 | `el9.x86_64` | pgdg | 46.0 KiB | [tds_fdw_15-2.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tds_fdw_15-2.0.3-1.rhel9.x86_64.rpm) |
| `tds_fdw_15` | 2.0.5 | `el9.aarch64` | pgdg | 45.7 KiB | [tds_fdw_15-2.0.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tds_fdw_15-2.0.5-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_15` | 2.0.4 | `el9.aarch64` | pgdg | 45.9 KiB | [tds_fdw_15-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tds_fdw_15-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_15` | 2.0.3 | `el9.aarch64` | pgdg | 44.6 KiB | [tds_fdw_15-2.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tds_fdw_15-2.0.3-1.rhel9.aarch64.rpm) |
| `tds_fdw_15` | 2.0.5 | `el10.x86_64` | pgdg | 48.5 KiB | [tds_fdw_15-2.0.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tds_fdw_15-2.0.5-1PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_15` | 2.0.4 | `el10.x86_64` | pgdg | 48.3 KiB | [tds_fdw_15-2.0.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tds_fdw_15-2.0.4-2PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_15` | 2.0.5 | `el10.aarch64` | pgdg | 47.1 KiB | [tds_fdw_15-2.0.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tds_fdw_15-2.0.5-1PGDG.rhel10.aarch64.rpm) |
| `tds_fdw_15` | 2.0.4 | `el10.aarch64` | pgdg | 46.9 KiB | [tds_fdw_15-2.0.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tds_fdw_15-2.0.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-tds-fdw` | 2.0.5 | `d12.x86_64` | pgdg | 111.5 KiB | [postgresql-15-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-15-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-15-tds-fdw` | 2.0.5 | `d12.aarch64` | pgdg | 108.7 KiB | [postgresql-15-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-15-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-15-tds-fdw` | 2.0.5 | `d13.x86_64` | pgdg | 111.2 KiB | [postgresql-15-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-15-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-15-tds-fdw` | 2.0.5 | `d13.aarch64` | pgdg | 109.1 KiB | [postgresql-15-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-15-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-15-tds-fdw` | 2.0.5 | `u22.x86_64` | pgdg | 126.3 KiB | [postgresql-15-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-15-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-tds-fdw` | 2.0.5 | `u22.aarch64` | pgdg | 122.9 KiB | [postgresql-15-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-15-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-tds-fdw` | 2.0.5 | `u24.x86_64` | pgdg | 109.8 KiB | [postgresql-15-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-15-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-tds-fdw` | 2.0.5 | `u24.aarch64` | pgdg | 107.0 KiB | [postgresql-15-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-15-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tds_fdw_14` | 2.0.5 | `el8.x86_64` | pgdg | 49.7 KiB | [tds_fdw_14-2.0.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tds_fdw_14-2.0.5-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_14` | 2.0.4 | `el8.x86_64` | pgdg | 49.3 KiB | [tds_fdw_14-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tds_fdw_14-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_14` | 2.0.3 | `el8.x86_64` | pgdg | 47.1 KiB | [tds_fdw_14-2.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tds_fdw_14-2.0.3-1.rhel8.x86_64.rpm) |
| `tds_fdw_14` | 2.0.2 | `el8.x86_64` | pgdg | 136.3 KiB | [tds_fdw_14-2.0.2-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tds_fdw_14-2.0.2-3.rhel8.x86_64.rpm) |
| `tds_fdw_14` | 2.0.5 | `el8.aarch64` | pgdg | 47.2 KiB | [tds_fdw_14-2.0.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tds_fdw_14-2.0.5-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_14` | 2.0.4 | `el8.aarch64` | pgdg | 46.7 KiB | [tds_fdw_14-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tds_fdw_14-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_14` | 2.0.3 | `el8.aarch64` | pgdg | 44.5 KiB | [tds_fdw_14-2.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tds_fdw_14-2.0.3-1.rhel8.aarch64.rpm) |
| `tds_fdw_14` | 2.0.5 | `el9.x86_64` | pgdg | 47.7 KiB | [tds_fdw_14-2.0.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tds_fdw_14-2.0.5-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_14` | 2.0.4 | `el9.x86_64` | pgdg | 47.7 KiB | [tds_fdw_14-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tds_fdw_14-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_14` | 2.0.3 | `el9.x86_64` | pgdg | 46.0 KiB | [tds_fdw_14-2.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tds_fdw_14-2.0.3-1.rhel9.x86_64.rpm) |
| `tds_fdw_14` | 2.0.5 | `el9.aarch64` | pgdg | 45.7 KiB | [tds_fdw_14-2.0.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tds_fdw_14-2.0.5-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_14` | 2.0.4 | `el9.aarch64` | pgdg | 46.5 KiB | [tds_fdw_14-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tds_fdw_14-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_14` | 2.0.3 | `el9.aarch64` | pgdg | 44.6 KiB | [tds_fdw_14-2.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tds_fdw_14-2.0.3-1.rhel9.aarch64.rpm) |
| `tds_fdw_14` | 2.0.5 | `el10.x86_64` | pgdg | 48.5 KiB | [tds_fdw_14-2.0.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tds_fdw_14-2.0.5-1PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_14` | 2.0.4 | `el10.x86_64` | pgdg | 48.3 KiB | [tds_fdw_14-2.0.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tds_fdw_14-2.0.4-2PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_14` | 2.0.5 | `el10.aarch64` | pgdg | 47.1 KiB | [tds_fdw_14-2.0.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tds_fdw_14-2.0.5-1PGDG.rhel10.aarch64.rpm) |
| `tds_fdw_14` | 2.0.4 | `el10.aarch64` | pgdg | 46.9 KiB | [tds_fdw_14-2.0.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tds_fdw_14-2.0.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-tds-fdw` | 2.0.5 | `d12.x86_64` | pgdg | 111.5 KiB | [postgresql-14-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-14-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-14-tds-fdw` | 2.0.5 | `d12.aarch64` | pgdg | 108.8 KiB | [postgresql-14-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-14-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-14-tds-fdw` | 2.0.5 | `d13.x86_64` | pgdg | 111.2 KiB | [postgresql-14-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-14-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-14-tds-fdw` | 2.0.5 | `d13.aarch64` | pgdg | 109.1 KiB | [postgresql-14-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-14-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-14-tds-fdw` | 2.0.5 | `u22.x86_64` | pgdg | 126.2 KiB | [postgresql-14-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-14-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-tds-fdw` | 2.0.5 | `u22.aarch64` | pgdg | 123.0 KiB | [postgresql-14-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-14-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-tds-fdw` | 2.0.5 | `u24.x86_64` | pgdg | 109.8 KiB | [postgresql-14-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-14-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-tds-fdw` | 2.0.5 | `u24.aarch64` | pgdg | 106.9 KiB | [postgresql-14-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-14-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tds_fdw_13` | 2.0.5 | `el8.x86_64` | pgdg | 49.1 KiB | [tds_fdw_13-2.0.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/tds_fdw_13-2.0.5-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_13` | 2.0.4 | `el8.x86_64` | pgdg | 48.7 KiB | [tds_fdw_13-2.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/tds_fdw_13-2.0.4-1PGDG.rhel8.x86_64.rpm) |
| `tds_fdw_13` | 2.0.3 | `el8.x86_64` | pgdg | 46.5 KiB | [tds_fdw_13-2.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/tds_fdw_13-2.0.3-1.rhel8.x86_64.rpm) |
| `tds_fdw_13` | 2.0.5 | `el8.aarch64` | pgdg | 47.1 KiB | [tds_fdw_13-2.0.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/tds_fdw_13-2.0.5-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_13` | 2.0.4 | `el8.aarch64` | pgdg | 46.6 KiB | [tds_fdw_13-2.0.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/tds_fdw_13-2.0.4-1PGDG.rhel8.aarch64.rpm) |
| `tds_fdw_13` | 2.0.3 | `el8.aarch64` | pgdg | 44.4 KiB | [tds_fdw_13-2.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/tds_fdw_13-2.0.3-1.rhel8.aarch64.rpm) |
| `tds_fdw_13` | 2.0.5 | `el9.x86_64` | pgdg | 47.7 KiB | [tds_fdw_13-2.0.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/tds_fdw_13-2.0.5-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_13` | 2.0.4 | `el9.x86_64` | pgdg | 47.7 KiB | [tds_fdw_13-2.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/tds_fdw_13-2.0.4-1PGDG.rhel9.x86_64.rpm) |
| `tds_fdw_13` | 2.0.3 | `el9.x86_64` | pgdg | 45.9 KiB | [tds_fdw_13-2.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/tds_fdw_13-2.0.3-1.rhel9.x86_64.rpm) |
| `tds_fdw_13` | 2.0.5 | `el9.aarch64` | pgdg | 45.8 KiB | [tds_fdw_13-2.0.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/tds_fdw_13-2.0.5-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_13` | 2.0.4 | `el9.aarch64` | pgdg | 45.9 KiB | [tds_fdw_13-2.0.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/tds_fdw_13-2.0.4-1PGDG.rhel9.aarch64.rpm) |
| `tds_fdw_13` | 2.0.3 | `el9.aarch64` | pgdg | 43.9 KiB | [tds_fdw_13-2.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/tds_fdw_13-2.0.3-1.rhel9.aarch64.rpm) |
| `tds_fdw_13` | 2.0.5 | `el10.x86_64` | pgdg | 48.5 KiB | [tds_fdw_13-2.0.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/tds_fdw_13-2.0.5-1PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_13` | 2.0.4 | `el10.x86_64` | pgdg | 48.2 KiB | [tds_fdw_13-2.0.4-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/tds_fdw_13-2.0.4-2PGDG.rhel10.x86_64.rpm) |
| `tds_fdw_13` | 2.0.5 | `el10.aarch64` | pgdg | 47.2 KiB | [tds_fdw_13-2.0.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/tds_fdw_13-2.0.5-1PGDG.rhel10.aarch64.rpm) |
| `tds_fdw_13` | 2.0.4 | `el10.aarch64` | pgdg | 47.0 KiB | [tds_fdw_13-2.0.4-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/tds_fdw_13-2.0.4-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-tds-fdw` | 2.0.5 | `d12.x86_64` | pgdg | 111.0 KiB | [postgresql-13-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-13-tds-fdw_2.0.5-1.pgdg12+1_amd64.deb) |
| `postgresql-13-tds-fdw` | 2.0.5 | `d12.aarch64` | pgdg | 108.8 KiB | [postgresql-13-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-13-tds-fdw_2.0.5-1.pgdg12+1_arm64.deb) |
| `postgresql-13-tds-fdw` | 2.0.5 | `d13.x86_64` | pgdg | 111.3 KiB | [postgresql-13-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-13-tds-fdw_2.0.5-1.pgdg13+1_amd64.deb) |
| `postgresql-13-tds-fdw` | 2.0.5 | `d13.aarch64` | pgdg | 109.1 KiB | [postgresql-13-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-13-tds-fdw_2.0.5-1.pgdg13+1_arm64.deb) |
| `postgresql-13-tds-fdw` | 2.0.5 | `u22.x86_64` | pgdg | 126.1 KiB | [postgresql-13-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-13-tds-fdw_2.0.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-tds-fdw` | 2.0.5 | `u22.aarch64` | pgdg | 122.4 KiB | [postgresql-13-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-13-tds-fdw_2.0.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-tds-fdw` | 2.0.5 | `u24.x86_64` | pgdg | 109.6 KiB | [postgresql-13-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-13-tds-fdw_2.0.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-tds-fdw` | 2.0.5 | `u24.aarch64` | pgdg | 106.8 KiB | [postgresql-13-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tds-fdw/postgresql-13-tds-fdw_2.0.5-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tds-fdw/tds_fdw" title="Repository" icon="github" subtitle="github.com/tds-fdw/tds_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="tds_fdw-2.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get tds_fdw; # get tds_fdw source code
pig build dep tds_fdw; # install build dependencies
pig build pkg tds_fdw; # build extension rpm or deb
pig build ext tds_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install tds_fdw; # install by extension name, for the current active PG version
pig ext install tds_fdw; # install via package alias, for the active PG version
pig ext install tds_fdw -v 17;   # install for PG 17
pig ext install tds_fdw -v 16;   # install for PG 16
pig ext install tds_fdw -v 15;   # install for PG 15
pig ext install tds_fdw -v 14;   # install for PG 14
pig ext install tds_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION tds_fdw;
```

