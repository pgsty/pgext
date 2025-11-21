---
title: "oracle_fdw"
linkTitle: "oracle_fdw"
description: "foreign data wrapper for Oracle access"
weight: 8610
categories: ["FDW"]
width: full
---

[**oracle_fdw**](https://github.com/laurenz/oracle_fdw) : foreign data wrapper for Oracle access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8610** | {{< badge content="oracle_fdw" link="https://github.com/laurenz/oracle_fdw" >}} | {{< ext "oracle_fdw" >}} | `2.8.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "firebird_fdw" >}} {{< ext "orafce" >}} {{< ext "wrappers" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} |

> [!Note] require oracle-libs


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.8.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `oracle_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.8.0` | {{< bg "18" "oracle_fdw_18*" "green" >}} {{< bg "17" "oracle_fdw_17*" "green" >}} {{< bg "16" "oracle_fdw_16*" "green" >}} {{< bg "15" "oracle_fdw_15*" "green" >}} {{< bg "14" "oracle_fdw_14*" "green" >}} {{< bg "13" "oracle_fdw_13*" "green" >}} | `oracle_fdw_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.8.0` | {{< bg "18" "postgresql-18-oracle-fdw" "green" >}} {{< bg "17" "postgresql-17-oracle-fdw" "green" >}} {{< bg "16" "postgresql-16-oracle-fdw" "green" >}} {{< bg "15" "postgresql-15-oracle-fdw" "green" >}} {{< bg "14" "postgresql-14-oracle-fdw" "green" >}} {{< bg "13" "postgresql-13-oracle-fdw" "green" >}} | `postgresql-$v-oracle-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "oracle_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.8.0" "oracle_fdw_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_15 : AVAIL 11" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_14 : AVAIL 12" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_13 : AVAIL 12" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "oracle_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_15 : AVAIL 12" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_14 : AVAIL 12" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_13 : AVAIL 12" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "oracle_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.8.0" "oracle_fdw_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "oracle_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "oracle_fdw_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-13-oracle-fdw : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-13-oracle-fdw : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-13-oracle-fdw : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-13-oracle-fdw : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-13-oracle-fdw : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-13-oracle-fdw : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-13-oracle-fdw : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.8.0" "postgresql-18-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-17-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-16-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-15-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-14-oracle-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.0" "postgresql-13-oracle-fdw : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_18` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.2 KiB | [oracle_fdw_18-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/oracle_fdw_18-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_18` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.1 KiB | [oracle_fdw_18-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/oracle_fdw_18-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_18` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.1 KiB | [oracle_fdw_18-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/oracle_fdw_18-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_18` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.3 KiB | [oracle_fdw_18-2.8.0-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/oracle_fdw_18-2.8.0-8PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_18` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.2 KiB | [oracle_fdw_18-2.8.0-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/oracle_fdw_18-2.8.0-7PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_18` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.3 KiB | [oracle_fdw_18-2.8.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/oracle_fdw_18-2.8.0-6PGDG.rhel10.x86_64.rpm) |
| `postgresql-18-oracle-fdw` | `2.8.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 85.3 KiB | [postgresql-18-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-18-oracle-fdw` | `2.8.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.0 KiB | [postgresql-18-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-18-oracle-fdw` | `2.8.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 85.3 KiB | [postgresql-18-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-18-oracle-fdw` | `2.8.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.1 KiB | [postgresql-18-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-18-oracle-fdw` | `2.8.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 73.4 KiB | [postgresql-18-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-oracle-fdw` | `2.8.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 67.7 KiB | [postgresql-18-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-oracle-fdw` | `2.8.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.4 KiB | [postgresql-18-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-oracle-fdw` | `2.8.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.7 KiB | [postgresql-18-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-18-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_17` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.9 KiB | [oracle_fdw_17-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.7 KiB | [oracle_fdw_17-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_17-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_17-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.8 KiB | [oracle_fdw_17-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.8 KiB | [oracle_fdw_17-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/oracle_fdw_17-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.0 KiB | [oracle_fdw_17-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.9 KiB | [oracle_fdw_17-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.1 KiB | [oracle_fdw_17-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.0 KiB | [oracle_fdw_17-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.2 KiB | [oracle_fdw_17-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.4 KiB | [oracle_fdw_17-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.3 KiB | [oracle_fdw_17-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/oracle_fdw_17-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.3 KiB | [oracle_fdw_17-2.8.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/oracle_fdw_17-2.8.0-6PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.2 KiB | [oracle_fdw_17-2.8.0-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/oracle_fdw_17-2.8.0-7PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_17` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.3 KiB | [oracle_fdw_17-2.8.0-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/oracle_fdw_17-2.8.0-8PGDG.rhel10.x86_64.rpm) |
| `postgresql-17-oracle-fdw` | `2.8.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 85.3 KiB | [postgresql-17-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-17-oracle-fdw` | `2.8.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.0 KiB | [postgresql-17-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-17-oracle-fdw` | `2.8.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 85.4 KiB | [postgresql-17-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-17-oracle-fdw` | `2.8.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.2 KiB | [postgresql-17-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-17-oracle-fdw` | `2.8.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 73.4 KiB | [postgresql-17-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-oracle-fdw` | `2.8.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 67.6 KiB | [postgresql-17-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-oracle-fdw` | `2.8.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.7 KiB | [postgresql-17-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-oracle-fdw` | `2.8.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.7 KiB | [postgresql-17-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-17-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_16` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.7 KiB | [oracle_fdw_16-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_16-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_16-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.9 KiB | [oracle_fdw_16-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.8 KiB | [oracle_fdw_16-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.9 KiB | [oracle_fdw_16-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 85.9 KiB | [oracle_fdw_16-2.6.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.6.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 85.9 KiB | [oracle_fdw_16-2.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/oracle_fdw_16-2.6.0-2PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.1 KiB | [oracle_fdw_16-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.0 KiB | [oracle_fdw_16-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.1 KiB | [oracle_fdw_16-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.0 KiB | [oracle_fdw_16-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.4 KiB | [oracle_fdw_16-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.5 KiB | [oracle_fdw_16-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.3 KiB | [oracle_fdw_16-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.5 KiB | [oracle_fdw_16-2.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.6.0-2PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.4 KiB | [oracle_fdw_16-2.6.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/oracle_fdw_16-2.6.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.4 KiB | [oracle_fdw_16-2.8.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/oracle_fdw_16-2.8.0-6PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.4 KiB | [oracle_fdw_16-2.8.0-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/oracle_fdw_16-2.8.0-8PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_16` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 87.4 KiB | [oracle_fdw_16-2.8.0-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/oracle_fdw_16-2.8.0-7PGDG.rhel10.x86_64.rpm) |
| `postgresql-16-oracle-fdw` | `2.8.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 85.4 KiB | [postgresql-16-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-16-oracle-fdw` | `2.8.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.1 KiB | [postgresql-16-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-16-oracle-fdw` | `2.8.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 85.4 KiB | [postgresql-16-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-16-oracle-fdw` | `2.8.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.2 KiB | [postgresql-16-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-16-oracle-fdw` | `2.8.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 73.3 KiB | [postgresql-16-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-oracle-fdw` | `2.8.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 67.7 KiB | [postgresql-16-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-oracle-fdw` | `2.8.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.9 KiB | [postgresql-16-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-oracle-fdw` | `2.8.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.8 KiB | [postgresql-16-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-16-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_15` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.5 KiB | [oracle_fdw_15-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.5 KiB | [oracle_fdw_15-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.6 KiB | [oracle_fdw_15-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.4 KiB | [oracle_fdw_15-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.4 KiB | [oracle_fdw_15-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.5 KiB | [oracle_fdw_15-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.6 KiB | [oracle_fdw_15-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.7 KiB | [oracle_fdw_15-2.6.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.6.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 84.1 KiB | [oracle_fdw_15-2.5.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.5.0-2.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 84.0 KiB | [oracle_fdw_15-2.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.5.0-1.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 84.2 KiB | [oracle_fdw_15-2.5.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/oracle_fdw_15-2.5.0-3.rhel8.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.3 KiB | [oracle_fdw_15-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.4 KiB | [oracle_fdw_15-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.5 KiB | [oracle_fdw_15-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.4 KiB | [oracle_fdw_15-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 86.5 KiB | [oracle_fdw_15-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 86.7 KiB | [oracle_fdw_15-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 86.7 KiB | [oracle_fdw_15-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.6 KiB | [oracle_fdw_15-2.6.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.6.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.6 KiB | [oracle_fdw_15-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.8 KiB | [oracle_fdw_15-2.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.5.0-1.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.6 KiB | [oracle_fdw_15-2.5.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.5.0-3.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.5 KiB | [oracle_fdw_15-2.5.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/oracle_fdw_15-2.5.0-2.rhel9.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.4 KiB | [oracle_fdw_15-2.8.0-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/oracle_fdw_15-2.8.0-7PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.4 KiB | [oracle_fdw_15-2.8.0-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/oracle_fdw_15-2.8.0-8PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_15` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.5 KiB | [oracle_fdw_15-2.8.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/oracle_fdw_15-2.8.0-6PGDG.rhel10.x86_64.rpm) |
| `postgresql-15-oracle-fdw` | `2.8.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 86.2 KiB | [postgresql-15-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-15-oracle-fdw` | `2.8.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.8 KiB | [postgresql-15-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-15-oracle-fdw` | `2.8.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 86.3 KiB | [postgresql-15-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-15-oracle-fdw` | `2.8.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 80.0 KiB | [postgresql-15-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-15-oracle-fdw` | `2.8.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.6 KiB | [postgresql-15-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-oracle-fdw` | `2.8.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 69.6 KiB | [postgresql-15-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-oracle-fdw` | `2.8.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.5 KiB | [postgresql-15-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-oracle-fdw` | `2.8.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 69.6 KiB | [postgresql-15-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-15-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_14` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.9 KiB | [oracle_fdw_14-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.8 KiB | [oracle_fdw_14-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.7 KiB | [oracle_fdw_14-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.8 KiB | [oracle_fdw_14-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.7 KiB | [oracle_fdw_14-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_14-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.0 KiB | [oracle_fdw_14-2.6.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.6.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.9 KiB | [oracle_fdw_14-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 84.5 KiB | [oracle_fdw_14-2.5.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.5.0-3.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 84.3 KiB | [oracle_fdw_14-2.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.5.0-1.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 84.4 KiB | [oracle_fdw_14-2.5.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.5.0-2.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 83.6 KiB | [oracle_fdw_14-2.4.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/oracle_fdw_14-2.4.0-3.rhel8.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_14-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.7 KiB | [oracle_fdw_14-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_14-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.7 KiB | [oracle_fdw_14-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.0 KiB | [oracle_fdw_14-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 86.8 KiB | [oracle_fdw_14-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 86.9 KiB | [oracle_fdw_14-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.8 KiB | [oracle_fdw_14-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.8 KiB | [oracle_fdw_14-2.6.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.6.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.9 KiB | [oracle_fdw_14-2.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.5.0-1.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.7 KiB | [oracle_fdw_14-2.5.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.5.0-2.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.8 KiB | [oracle_fdw_14-2.5.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/oracle_fdw_14-2.5.0-3.rhel9.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.8 KiB | [oracle_fdw_14-2.8.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/oracle_fdw_14-2.8.0-6PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.7 KiB | [oracle_fdw_14-2.8.0-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/oracle_fdw_14-2.8.0-8PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_14` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.7 KiB | [oracle_fdw_14-2.8.0-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/oracle_fdw_14-2.8.0-7PGDG.rhel10.x86_64.rpm) |
| `postgresql-14-oracle-fdw` | `2.8.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 86.5 KiB | [postgresql-14-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-14-oracle-fdw` | `2.8.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 80.1 KiB | [postgresql-14-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-14-oracle-fdw` | `2.8.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 86.8 KiB | [postgresql-14-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-14-oracle-fdw` | `2.8.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 80.6 KiB | [postgresql-14-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-14-oracle-fdw` | `2.8.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.7 KiB | [postgresql-14-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-oracle-fdw` | `2.8.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 70.0 KiB | [postgresql-14-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-oracle-fdw` | `2.8.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 75.9 KiB | [postgresql-14-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-oracle-fdw` | `2.8.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 69.9 KiB | [postgresql-14-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-14-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `oracle_fdw_13` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.2 KiB | [oracle_fdw_13-2.8.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.8.0-5PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.4 KiB | [oracle_fdw_13-2.8.0-8PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.8.0-8PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.3 KiB | [oracle_fdw_13-2.8.0-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.8.0-7PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 88.3 KiB | [oracle_fdw_13-2.8.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.8.0-6PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.4 KiB | [oracle_fdw_13-2.7.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.7.0-4PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 87.3 KiB | [oracle_fdw_13-2.7.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.7.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.4 KiB | [oracle_fdw_13-2.6.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.6.0-3PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 86.3 KiB | [oracle_fdw_13-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 83.9 KiB | [oracle_fdw_13-2.5.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.5.0-2.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 83.9 KiB | [oracle_fdw_13-2.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.5.0-1.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 84.0 KiB | [oracle_fdw_13-2.5.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.5.0-3.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 83.2 KiB | [oracle_fdw_13-2.4.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/oracle_fdw_13-2.4.0-3.rhel8.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_13-2.8.0-8PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.8.0-8PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.7 KiB | [oracle_fdw_13-2.8.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.8.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.7 KiB | [oracle_fdw_13-2.8.0-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.8.0-7PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.8 KiB | [oracle_fdw_13-2.8.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.8.0-6PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 87.1 KiB | [oracle_fdw_13-2.7.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.7.0-5PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 86.9 KiB | [oracle_fdw_13-2.7.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.7.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 86.8 KiB | [oracle_fdw_13-2.7.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.7.0-4PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.8 KiB | [oracle_fdw_13-2.6.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.6.0-3PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 85.8 KiB | [oracle_fdw_13-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.8 KiB | [oracle_fdw_13-2.5.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.5.0-3.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.0 KiB | [oracle_fdw_13-2.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.5.0-1.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 83.8 KiB | [oracle_fdw_13-2.5.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/oracle_fdw_13-2.5.0-2.rhel9.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.5 KiB | [oracle_fdw_13-2.8.0-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-10-x86_64/oracle_fdw_13-2.8.0-8PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.5 KiB | [oracle_fdw_13-2.8.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-10-x86_64/oracle_fdw_13-2.8.0-6PGDG.rhel10.x86_64.rpm) |
| `oracle_fdw_13` | `2.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 88.4 KiB | [oracle_fdw_13-2.8.0-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-10-x86_64/oracle_fdw_13-2.8.0-7PGDG.rhel10.x86_64.rpm) |
| `postgresql-13-oracle-fdw` | `2.8.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 86.6 KiB | [postgresql-13-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-13-oracle-fdw_2.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-13-oracle-fdw` | `2.8.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 80.0 KiB | [postgresql-13-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-13-oracle-fdw_2.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-13-oracle-fdw` | `2.8.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 86.8 KiB | [postgresql-13-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-13-oracle-fdw_2.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-13-oracle-fdw` | `2.8.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 80.3 KiB | [postgresql-13-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-13-oracle-fdw_2.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-13-oracle-fdw` | `2.8.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.7 KiB | [postgresql-13-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-13-oracle-fdw_2.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-oracle-fdw` | `2.8.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 70.0 KiB | [postgresql-13-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-13-oracle-fdw_2.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-oracle-fdw` | `2.8.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.1 KiB | [postgresql-13-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-13-oracle-fdw_2.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-oracle-fdw` | `2.8.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 69.9 KiB | [postgresql-13-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/oracle-fdw/postgresql-13-oracle-fdw_2.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/laurenz/oracle_fdw" title="Repository" icon="github" subtitle="github.com/laurenz/oracle_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="oracle_fdw-ORACLE_FDW_2_8_0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg oracle_fdw;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install oracle_fdw;		# install via package name, for the active PG version

pig install oracle_fdw -v 18;   # install for PG 18
pig install oracle_fdw -v 17;   # install for PG 17
pig install oracle_fdw -v 16;   # install for PG 16
pig install oracle_fdw -v 15;   # install for PG 15
pig install oracle_fdw -v 14;   # install for PG 14
pig install oracle_fdw -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION oracle_fdw;
```
