---
title: "oracle_fdw"
linkTitle: "oracle_fdw"
description: "foreign data wrapper for Oracle access"
weight: 8610
categories: ["FDW"]
width: full
---

foreign data wrapper for Oracle access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8610** | {{< badge content="oracle_fdw" link="https://github.com/laurenz/oracle_fdw" >}} | {{< ext "oracle_fdw" >}} | `2.8.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "firebird_fdw" >}} {{< ext "orafce" >}} {{< ext "wrappers" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} |

> [!Note] require oracle-libs


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/oracle_fdw" >}} | `2.8.0` | {{< bg "18" "oracle_fdw_18*" "red" >}} {{< bg "17" "oracle_fdw_17*" "green" >}} {{< bg "16" "oracle_fdw_16*" "green" >}} {{< bg "15" "oracle_fdw_15*" "green" >}} {{< bg "14" "oracle_fdw_14*" "green" >}} | `oracle_fdw_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/oracle_fdw" >}} | `2.8.0` | {{< bg "18" "postgresql-18-oracle-fdw" "green" >}} {{< bg "17" "postgresql-17-oracle-fdw" "green" >}} {{< bg "16" "postgresql-16-oracle-fdw" "green" >}} {{< bg "15" "postgresql-15-oracle-fdw" "green" >}} {{< bg "14" "postgresql-14-oracle-fdw" "green" >}} | `postgresql-$v-oracle-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "MISS" "oracle_fdw_18 : HIDE 0" >}}   |  {{< bg "PGDG 2.8.0" "oracle_fdw_17 : HIDE 6" >}}   |  {{< bg "PGDG 2.8.0" "oracle_fdw_16 : HIDE 8" >}}   |  {{< bg "PGDG 2.8.0" "oracle_fdw_15 : HIDE 11" >}}   |  {{< bg "PGDG 2.8.0" "oracle_fdw_14 : HIDE 12" >}}   |
|    `el8.aarch64`    |  {{< bg "MISS" "oracle_fdw_18 : HIDE 0" >}}   |  {{< bg "MISS" "oracle_fdw_17 : HIDE 0" >}}   |  {{< bg "MISS" "oracle_fdw_16 : HIDE 0" >}}   |  {{< bg "MISS" "oracle_fdw_15 : HIDE 0" >}}   |  {{< bg "MISS" "oracle_fdw_14 : HIDE 0" >}}   |
|    `el9.x86_64`    |  {{< bg "PGDG 2.8.0" "oracle_fdw_18 : HIDE 3" >}}   |  {{< bg "PGDG 2.8.0" "oracle_fdw_17 : HIDE 7" >}}   |  {{< bg "PGDG 2.8.0" "oracle_fdw_16 : HIDE 9" >}}   |  {{< bg "PGDG 2.8.0" "oracle_fdw_15 : HIDE 12" >}}   |  {{< bg "PGDG 2.8.0" "oracle_fdw_14 : HIDE 12" >}}   |
|    `el9.aarch64`    |  {{< bg "MISS" "oracle_fdw_18 : HIDE 0" >}}   |  {{< bg "MISS" "oracle_fdw_17 : HIDE 0" >}}   |  {{< bg "MISS" "oracle_fdw_16 : HIDE 0" >}}   |  {{< bg "MISS" "oracle_fdw_15 : HIDE 0" >}}   |  {{< bg "MISS" "oracle_fdw_14 : HIDE 0" >}}   |
|    `d12.x86_64`    |  {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : HIDE 1" >}}   |
|    `d12.aarch64`    |  {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : HIDE 1" >}}   |
|    `u22.x86_64`    |  {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : HIDE 1" >}}   |
|    `u22.aarch64`    |  {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : HIDE 1" >}}   |
|    `u24.x86_64`    |  {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : HIDE 1" >}}   |
|    `u24.aarch64`    |  {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : HIDE 1" >}}   |  {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : HIDE 1" >}}   |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_18` | 2.8.0 | `el9.x86_64` | pgdg | 85.1 KiB | [oracle_fdw_18-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/oracle_fdw_18-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_18` | 2.8.0 | `el9.x86_64` | pgdg | 85.2 KiB | [oracle_fdw_18-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/oracle_fdw_18-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_18` | 2.8.0 | `el9.x86_64` | pgdg | 85.1 KiB | [oracle_fdw_18-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/oracle_fdw_18-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `postgresql-18-oracle-fdw` | 2.8.0 | `d12.x86_64` | pgdg | 85.2 KiB | [postgresql-18-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb) |
| `postgresql-18-oracle-fdw` | 2.8.0 | `d12.aarch64` | pgdg | 79.0 KiB | [postgresql-18-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb) |
| `postgresql-18-oracle-fdw` | 2.8.0 | `u22.x86_64` | pgdg | 73.4 KiB | [postgresql-18-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-oracle-fdw` | 2.8.0 | `u22.aarch64` | pgdg | 67.6 KiB | [postgresql-18-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-oracle-fdw` | 2.8.0 | `u24.x86_64` | pgdg | 73.6 KiB | [postgresql-18-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-oracle-fdw` | 2.8.0 | `u24.aarch64` | pgdg | 67.6 KiB | [postgresql-18-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_17` | 2.8.0 | `el8.x86_64` | pgdg | 87.7 KiB | [oracle_fdw_17-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | 2.8.0 | `el8.x86_64` | pgdg | 87.9 KiB | [oracle_fdw_17-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | 2.8.0 | `el8.x86_64` | pgdg | 87.8 KiB | [oracle_fdw_17-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | 2.8.0 | `el8.x86_64` | pgdg | 87.8 KiB | [oracle_fdw_17-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | 2.7.0 | `el8.x86_64` | pgdg | 86.8 KiB | [oracle_fdw_17-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | 2.7.0 | `el8.x86_64` | pgdg | 86.8 KiB | [oracle_fdw_17-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | 2.8.0 | `el9.x86_64` | pgdg | 84.9 KiB | [oracle_fdw_17-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | 2.8.0 | `el9.x86_64` | pgdg | 85.1 KiB | [oracle_fdw_17-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | 2.8.0 | `el9.x86_64` | pgdg | 85.0 KiB | [oracle_fdw_17-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | 2.8.0 | `el9.x86_64` | pgdg | 85.0 KiB | [oracle_fdw_17-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | 2.7.0 | `el9.x86_64` | pgdg | 84.2 KiB | [oracle_fdw_17-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | 2.7.0 | `el9.x86_64` | pgdg | 84.4 KiB | [oracle_fdw_17-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | 2.7.0 | `el9.x86_64` | pgdg | 84.3 KiB | [oracle_fdw_17-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-oracle-fdw` | 2.8.0 | `d12.x86_64` | pgdg | 85.2 KiB | [postgresql-17-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb) |
| `postgresql-17-oracle-fdw` | 2.8.0 | `d12.aarch64` | pgdg | 79.0 KiB | [postgresql-17-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb) |
| `postgresql-17-oracle-fdw` | 2.8.0 | `u22.x86_64` | pgdg | 73.2 KiB | [postgresql-17-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-oracle-fdw` | 2.8.0 | `u22.aarch64` | pgdg | 67.7 KiB | [postgresql-17-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-oracle-fdw` | 2.8.0 | `u24.x86_64` | pgdg | 73.6 KiB | [postgresql-17-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-oracle-fdw` | 2.8.0 | `u24.aarch64` | pgdg | 67.6 KiB | [postgresql-17-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_16` | 2.8.0 | `el8.x86_64` | pgdg | 87.8 KiB | [oracle_fdw_16-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | 2.8.0 | `el8.x86_64` | pgdg | 87.7 KiB | [oracle_fdw_16-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | 2.8.0 | `el8.x86_64` | pgdg | 87.8 KiB | [oracle_fdw_16-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | 2.8.0 | `el8.x86_64` | pgdg | 87.9 KiB | [oracle_fdw_16-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | 2.7.0 | `el8.x86_64` | pgdg | 86.8 KiB | [oracle_fdw_16-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | 2.7.0 | `el8.x86_64` | pgdg | 86.9 KiB | [oracle_fdw_16-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | 2.6.0 | `el8.x86_64` | pgdg | 85.9 KiB | [oracle_fdw_16-2.6.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.6.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | 2.6.0 | `el8.x86_64` | pgdg | 85.9 KiB | [oracle_fdw_16-2.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.6.0-2PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | 2.8.0 | `el9.x86_64` | pgdg | 85.1 KiB | [oracle_fdw_16-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | 2.8.0 | `el9.x86_64` | pgdg | 85.0 KiB | [oracle_fdw_16-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | 2.8.0 | `el9.x86_64` | pgdg | 85.0 KiB | [oracle_fdw_16-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | 2.8.0 | `el9.x86_64` | pgdg | 85.1 KiB | [oracle_fdw_16-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | 2.7.0 | `el9.x86_64` | pgdg | 84.4 KiB | [oracle_fdw_16-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | 2.7.0 | `el9.x86_64` | pgdg | 84.3 KiB | [oracle_fdw_16-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | 2.7.0 | `el9.x86_64` | pgdg | 84.5 KiB | [oracle_fdw_16-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | 2.6.0 | `el9.x86_64` | pgdg | 83.4 KiB | [oracle_fdw_16-2.6.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.6.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | 2.6.0 | `el9.x86_64` | pgdg | 83.5 KiB | [oracle_fdw_16-2.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.6.0-2PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-oracle-fdw` | 2.8.0 | `d12.x86_64` | pgdg | 85.3 KiB | [postgresql-16-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb) |
| `postgresql-16-oracle-fdw` | 2.8.0 | `d12.aarch64` | pgdg | 79.1 KiB | [postgresql-16-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb) |
| `postgresql-16-oracle-fdw` | 2.8.0 | `u22.x86_64` | pgdg | 73.3 KiB | [postgresql-16-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-oracle-fdw` | 2.8.0 | `u22.aarch64` | pgdg | 67.6 KiB | [postgresql-16-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-oracle-fdw` | 2.8.0 | `u24.x86_64` | pgdg | 73.7 KiB | [postgresql-16-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-oracle-fdw` | 2.8.0 | `u24.aarch64` | pgdg | 67.8 KiB | [postgresql-16-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_15` | 2.8.0 | `el8.x86_64` | pgdg | 88.5 KiB | [oracle_fdw_15-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.8.0 | `el8.x86_64` | pgdg | 88.6 KiB | [oracle_fdw_15-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.8.0 | `el8.x86_64` | pgdg | 88.5 KiB | [oracle_fdw_15-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.8.0 | `el8.x86_64` | pgdg | 88.4 KiB | [oracle_fdw_15-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.7.0 | `el8.x86_64` | pgdg | 87.4 KiB | [oracle_fdw_15-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.7.0 | `el8.x86_64` | pgdg | 87.5 KiB | [oracle_fdw_15-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.6.0 | `el8.x86_64` | pgdg | 86.7 KiB | [oracle_fdw_15-2.6.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.6.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.6.0 | `el8.x86_64` | pgdg | 86.6 KiB | [oracle_fdw_15-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.5.0 | `el8.x86_64` | pgdg | 84.1 KiB | [oracle_fdw_15-2.5.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.5.0-2.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.5.0 | `el8.x86_64` | pgdg | 84.0 KiB | [oracle_fdw_15-2.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.5.0-1.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.5.0 | `el8.x86_64` | pgdg | 84.2 KiB | [oracle_fdw_15-2.5.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.5.0-3.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | 2.8.0 | `el9.x86_64` | pgdg | 87.4 KiB | [oracle_fdw_15-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.8.0 | `el9.x86_64` | pgdg | 87.4 KiB | [oracle_fdw_15-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.8.0 | `el9.x86_64` | pgdg | 87.5 KiB | [oracle_fdw_15-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.8.0 | `el9.x86_64` | pgdg | 87.3 KiB | [oracle_fdw_15-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.7.0 | `el9.x86_64` | pgdg | 86.7 KiB | [oracle_fdw_15-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.7.0 | `el9.x86_64` | pgdg | 86.5 KiB | [oracle_fdw_15-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.7.0 | `el9.x86_64` | pgdg | 86.7 KiB | [oracle_fdw_15-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.6.0 | `el9.x86_64` | pgdg | 85.6 KiB | [oracle_fdw_15-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.6.0 | `el9.x86_64` | pgdg | 85.6 KiB | [oracle_fdw_15-2.6.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.6.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.5.0 | `el9.x86_64` | pgdg | 83.8 KiB | [oracle_fdw_15-2.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.5.0-1.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.5.0 | `el9.x86_64` | pgdg | 83.5 KiB | [oracle_fdw_15-2.5.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.5.0-2.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | 2.5.0 | `el9.x86_64` | pgdg | 83.6 KiB | [oracle_fdw_15-2.5.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.5.0-3.rhel9.x86_64.rpm) |
| `postgresql-15-oracle-fdw` | 2.8.0 | `d12.x86_64` | pgdg | 86.2 KiB | [postgresql-15-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb) |
| `postgresql-15-oracle-fdw` | 2.8.0 | `d12.aarch64` | pgdg | 79.8 KiB | [postgresql-15-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb) |
| `postgresql-15-oracle-fdw` | 2.8.0 | `u22.x86_64` | pgdg | 75.5 KiB | [postgresql-15-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-oracle-fdw` | 2.8.0 | `u22.aarch64` | pgdg | 69.7 KiB | [postgresql-15-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-oracle-fdw` | 2.8.0 | `u24.x86_64` | pgdg | 75.6 KiB | [postgresql-15-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-oracle-fdw` | 2.8.0 | `u24.aarch64` | pgdg | 69.5 KiB | [postgresql-15-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_14` | 2.8.0 | `el8.x86_64` | pgdg | 88.8 KiB | [oracle_fdw_14-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.8.0 | `el8.x86_64` | pgdg | 88.7 KiB | [oracle_fdw_14-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.8.0 | `el8.x86_64` | pgdg | 88.9 KiB | [oracle_fdw_14-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.8.0 | `el8.x86_64` | pgdg | 88.8 KiB | [oracle_fdw_14-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.7.0 | `el8.x86_64` | pgdg | 87.8 KiB | [oracle_fdw_14-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.7.0 | `el8.x86_64` | pgdg | 87.7 KiB | [oracle_fdw_14-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.6.0 | `el8.x86_64` | pgdg | 86.9 KiB | [oracle_fdw_14-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.6.0 | `el8.x86_64` | pgdg | 87.0 KiB | [oracle_fdw_14-2.6.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.6.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.5.0 | `el8.x86_64` | pgdg | 84.3 KiB | [oracle_fdw_14-2.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.5.0-1.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.5.0 | `el8.x86_64` | pgdg | 84.5 KiB | [oracle_fdw_14-2.5.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.5.0-3.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.5.0 | `el8.x86_64` | pgdg | 84.4 KiB | [oracle_fdw_14-2.5.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.5.0-2.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.4.0 | `el8.x86_64` | pgdg | 83.6 KiB | [oracle_fdw_14-2.4.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.4.0-3.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | 2.8.0 | `el9.x86_64` | pgdg | 87.7 KiB | [oracle_fdw_14-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.8.0 | `el9.x86_64` | pgdg | 87.8 KiB | [oracle_fdw_14-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.8.0 | `el9.x86_64` | pgdg | 87.7 KiB | [oracle_fdw_14-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.8.0 | `el9.x86_64` | pgdg | 87.8 KiB | [oracle_fdw_14-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.7.0 | `el9.x86_64` | pgdg | 86.8 KiB | [oracle_fdw_14-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.7.0 | `el9.x86_64` | pgdg | 87.0 KiB | [oracle_fdw_14-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.7.0 | `el9.x86_64` | pgdg | 86.9 KiB | [oracle_fdw_14-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.6.0 | `el9.x86_64` | pgdg | 85.8 KiB | [oracle_fdw_14-2.6.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.6.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.6.0 | `el9.x86_64` | pgdg | 85.8 KiB | [oracle_fdw_14-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.5.0 | `el9.x86_64` | pgdg | 83.8 KiB | [oracle_fdw_14-2.5.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.5.0-3.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.5.0 | `el9.x86_64` | pgdg | 83.7 KiB | [oracle_fdw_14-2.5.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.5.0-2.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | 2.5.0 | `el9.x86_64` | pgdg | 83.9 KiB | [oracle_fdw_14-2.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.5.0-1.rhel9.x86_64.rpm) |
| `postgresql-14-oracle-fdw` | 2.8.0 | `d12.x86_64` | pgdg | 86.5 KiB | [postgresql-14-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-1.pgdg120+1_amd64.deb) |
| `postgresql-14-oracle-fdw` | 2.8.0 | `d12.aarch64` | pgdg | 80.1 KiB | [postgresql-14-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-1.pgdg120+1_arm64.deb) |
| `postgresql-14-oracle-fdw` | 2.8.0 | `u22.x86_64` | pgdg | 75.7 KiB | [postgresql-14-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-oracle-fdw` | 2.8.0 | `u22.aarch64` | pgdg | 70.0 KiB | [postgresql-14-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-oracle-fdw` | 2.8.0 | `u24.x86_64` | pgdg | 75.7 KiB | [postgresql-14-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-oracle-fdw` | 2.8.0 | `u24.aarch64` | pgdg | 69.9 KiB | [postgresql-14-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/laurenz/oracle_fdw" title="Repository" icon="github" subtitle="github.com/laurenz/oracle_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="oracle_fdw-ORACLE_FDW_2_8_0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get oracle_fdw; # get oracle_fdw source code
pig build dep oracle_fdw; # install build dependencies
pig build pkg oracle_fdw; # build extension rpm or deb
pig build ext oracle_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install oracle_fdw; # install by extension name, for the current active PG version
pig ext install oracle_fdw; # install via package alias, for the active PG version
pig ext install oracle_fdw -v 18;   # install for PG 18
pig ext install oracle_fdw -v 17;   # install for PG 17
pig ext install oracle_fdw -v 16;   # install for PG 16
pig ext install oracle_fdw -v 15;   # install for PG 15
pig ext install oracle_fdw -v 14;   # install for PG 14
pig ext install oracle_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION oracle_fdw;
```

