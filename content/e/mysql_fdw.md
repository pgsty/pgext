---
title: "mysql_fdw"
linkTitle: "mysql_fdw"
description: "Foreign data wrapper for querying a MySQL server"
weight: 8600
categories: ["FDW"]
width: full
---

[**mysql_fdw**](https://github.com/EnterpriseDB/mysql_fdw) : Foreign data wrapper for querying a MySQL server


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8600** | {{< badge content="mysql_fdw" link="https://github.com/EnterpriseDB/mysql_fdw" >}} | {{< ext "mysql_fdw" >}} | `2.9.3` | {{< category "FDW" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.9.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `mysql_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.9.3` | {{< bg "18" "mysql_fdw_18*" "green" >}} {{< bg "17" "mysql_fdw_17*" "green" >}} {{< bg "16" "mysql_fdw_16*" "green" >}} {{< bg "15" "mysql_fdw_15*" "green" >}} {{< bg "14" "mysql_fdw_14*" "green" >}} {{< bg "13" "mysql_fdw_13*" "green" >}} | `mysql_fdw_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.9.3` | {{< bg "18" "postgresql-18-mysql-fdw" "green" >}} {{< bg "17" "postgresql-17-mysql-fdw" "green" >}} {{< bg "16" "postgresql-16-mysql-fdw" "green" >}} {{< bg "15" "postgresql-15-mysql-fdw" "green" >}} {{< bg "14" "postgresql-14-mysql-fdw" "green" >}} {{< bg "13" "postgresql-13-mysql-fdw" "green" >}} | `postgresql-$v-mysql-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_14 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_13 : AVAIL 10" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_13 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_13 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.9.3" "mysql_fdw_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.9.3" "postgresql-18-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-17-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-16-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-15-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-14-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-13-mysql-fdw : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.9.3" "postgresql-18-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-17-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-16-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-15-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-14-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-13-mysql-fdw : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.9.3" "postgresql-18-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-17-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-16-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-15-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-14-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-13-mysql-fdw : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.9.3" "postgresql-18-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-17-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-16-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-15-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-14-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-13-mysql-fdw : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.9.3" "postgresql-18-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-17-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-16-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-15-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-14-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-13-mysql-fdw : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.9.3" "postgresql-18-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-17-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-16-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-15-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-14-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-13-mysql-fdw : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.9.3" "postgresql-18-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-17-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-16-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-15-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-14-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-13-mysql-fdw : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.9.3" "postgresql-18-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-17-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-16-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-15-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-14-mysql-fdw : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.9.3" "postgresql-13-mysql-fdw : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mysql_fdw_18` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.6 KiB | [mysql_fdw_18-2.9.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/mysql_fdw_18-2.9.3-3PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_18` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.4 KiB | [mysql_fdw_18-2.9.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/mysql_fdw_18-2.9.3-3PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_18` | `2.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 135.3 KiB | [mysql_fdw_18-2.9.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/mysql_fdw_18-2.9.3-3PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_18` | `2.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 133.7 KiB | [mysql_fdw_18-2.9.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/mysql_fdw_18-2.9.3-3PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_18` | `2.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 136.9 KiB | [mysql_fdw_18-2.9.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/mysql_fdw_18-2.9.3-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_18` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 135.3 KiB | [mysql_fdw_18-2.9.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/mysql_fdw_18-2.9.3-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-mysql-fdw` | `2.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.1 KiB | [postgresql-18-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-18-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-mysql-fdw` | `2.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.7 KiB | [postgresql-18-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-18-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-mysql-fdw` | `2.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.5 KiB | [postgresql-18-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-18-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-mysql-fdw` | `2.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 121.7 KiB | [postgresql-18-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-18-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-mysql-fdw` | `2.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 129.6 KiB | [postgresql-18-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-18-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-mysql-fdw` | `2.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 126.4 KiB | [postgresql-18-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-18-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-mysql-fdw` | `2.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 123.9 KiB | [postgresql-18-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-18-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-mysql-fdw` | `2.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.7 KiB | [postgresql-18-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-18-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mysql_fdw_17` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.4 KiB | [mysql_fdw_17-2.9.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/mysql_fdw_17-2.9.3-3PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_17` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.1 KiB | [mysql_fdw_17-2.9.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/mysql_fdw_17-2.9.2-2PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_17` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.0 KiB | [mysql_fdw_17-2.9.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/mysql_fdw_17-2.9.3-3PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_17` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 137.1 KiB | [mysql_fdw_17-2.9.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/mysql_fdw_17-2.9.2-2PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_17` | `2.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 135.0 KiB | [mysql_fdw_17-2.9.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/mysql_fdw_17-2.9.3-3PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_17` | `2.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 139.5 KiB | [mysql_fdw_17-2.9.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/mysql_fdw_17-2.9.2-2PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_17` | `2.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 133.3 KiB | [mysql_fdw_17-2.9.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/mysql_fdw_17-2.9.3-3PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_17` | `2.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 138.0 KiB | [mysql_fdw_17-2.9.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/mysql_fdw_17-2.9.2-2PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_17` | `2.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 136.5 KiB | [mysql_fdw_17-2.9.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/mysql_fdw_17-2.9.3-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_17` | `2.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 136.4 KiB | [mysql_fdw_17-2.9.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/mysql_fdw_17-2.9.2-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_17` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 134.8 KiB | [mysql_fdw_17-2.9.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/mysql_fdw_17-2.9.3-3PGDG.rhel10.aarch64.rpm) |
| `mysql_fdw_17` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 134.8 KiB | [mysql_fdw_17-2.9.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/mysql_fdw_17-2.9.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-mysql-fdw` | `2.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 123.9 KiB | [postgresql-17-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-17-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-mysql-fdw` | `2.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.5 KiB | [postgresql-17-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-17-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-mysql-fdw` | `2.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.2 KiB | [postgresql-17-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-17-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-mysql-fdw` | `2.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 121.5 KiB | [postgresql-17-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-17-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-mysql-fdw` | `2.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 146.1 KiB | [postgresql-17-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-17-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-mysql-fdw` | `2.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 143.0 KiB | [postgresql-17-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-17-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-mysql-fdw` | `2.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 123.6 KiB | [postgresql-17-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-17-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-mysql-fdw` | `2.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.5 KiB | [postgresql-17-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-17-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mysql_fdw_16` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.4 KiB | [mysql_fdw_16-2.9.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/mysql_fdw_16-2.9.3-3PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_16` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.0 KiB | [mysql_fdw_16-2.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/mysql_fdw_16-2.9.2-1PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_16` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 155.4 KiB | [mysql_fdw_16-2.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/mysql_fdw_16-2.9.1-1PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_16` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.0 KiB | [mysql_fdw_16-2.9.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/mysql_fdw_16-2.9.3-3PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_16` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 137.1 KiB | [mysql_fdw_16-2.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/mysql_fdw_16-2.9.2-1PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_16` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 152.9 KiB | [mysql_fdw_16-2.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/mysql_fdw_16-2.9.1-1PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_16` | `2.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 135.0 KiB | [mysql_fdw_16-2.9.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/mysql_fdw_16-2.9.3-3PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_16` | `2.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 139.2 KiB | [mysql_fdw_16-2.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/mysql_fdw_16-2.9.2-1PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_16` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 155.7 KiB | [mysql_fdw_16-2.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/mysql_fdw_16-2.9.1-1PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_16` | `2.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 133.2 KiB | [mysql_fdw_16-2.9.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/mysql_fdw_16-2.9.3-3PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_16` | `2.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 138.0 KiB | [mysql_fdw_16-2.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/mysql_fdw_16-2.9.2-1PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_16` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.0 KiB | [mysql_fdw_16-2.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/mysql_fdw_16-2.9.1-1PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_16` | `2.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 136.5 KiB | [mysql_fdw_16-2.9.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/mysql_fdw_16-2.9.3-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_16` | `2.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 136.4 KiB | [mysql_fdw_16-2.9.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/mysql_fdw_16-2.9.2-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_16` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 134.8 KiB | [mysql_fdw_16-2.9.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/mysql_fdw_16-2.9.3-3PGDG.rhel10.aarch64.rpm) |
| `mysql_fdw_16` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 134.7 KiB | [mysql_fdw_16-2.9.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/mysql_fdw_16-2.9.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-mysql-fdw` | `2.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 123.7 KiB | [postgresql-16-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-16-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-mysql-fdw` | `2.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.3 KiB | [postgresql-16-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-16-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-mysql-fdw` | `2.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.2 KiB | [postgresql-16-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-16-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-mysql-fdw` | `2.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 121.2 KiB | [postgresql-16-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-16-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-mysql-fdw` | `2.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 145.6 KiB | [postgresql-16-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-16-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-mysql-fdw` | `2.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 142.4 KiB | [postgresql-16-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-16-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-mysql-fdw` | `2.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 123.7 KiB | [postgresql-16-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-16-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-mysql-fdw` | `2.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.2 KiB | [postgresql-16-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-16-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mysql_fdw_15` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.7 KiB | [mysql_fdw_15-2.9.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mysql_fdw_15-2.9.3-3PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.2 KiB | [mysql_fdw_15-2.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mysql_fdw_15-2.9.2-1PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 155.6 KiB | [mysql_fdw_15-2.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mysql_fdw_15-2.9.1-1PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 155.3 KiB | [mysql_fdw_15-2.9.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mysql_fdw_15-2.9.0-1.rhel8.x86_64.rpm) |
| `mysql_fdw_15` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 138.8 KiB | [mysql_fdw_15-2.8.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mysql_fdw_15-2.8.0-2.rhel8.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.0 KiB | [mysql_fdw_15-2.9.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/mysql_fdw_15-2.9.3-3PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 137.0 KiB | [mysql_fdw_15-2.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/mysql_fdw_15-2.9.2-1PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 134.9 KiB | [mysql_fdw_15-2.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/mysql_fdw_15-2.9.1-1PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 152.7 KiB | [mysql_fdw_15-2.9.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/mysql_fdw_15-2.9.0-1.rhel8.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 135.8 KiB | [mysql_fdw_15-2.9.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mysql_fdw_15-2.9.3-3PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 140.1 KiB | [mysql_fdw_15-2.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mysql_fdw_15-2.9.2-1PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 156.9 KiB | [mysql_fdw_15-2.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mysql_fdw_15-2.9.1-1PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 156.4 KiB | [mysql_fdw_15-2.9.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mysql_fdw_15-2.9.0-1.rhel9.x86_64.rpm) |
| `mysql_fdw_15` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 140.2 KiB | [mysql_fdw_15-2.8.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mysql_fdw_15-2.8.0-2.rhel9.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 133.7 KiB | [mysql_fdw_15-2.9.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/mysql_fdw_15-2.9.3-3PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 138.6 KiB | [mysql_fdw_15-2.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/mysql_fdw_15-2.9.2-1PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 136.7 KiB | [mysql_fdw_15-2.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/mysql_fdw_15-2.9.1-1PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.2 KiB | [mysql_fdw_15-2.9.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/mysql_fdw_15-2.9.0-1.rhel9.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 137.6 KiB | [mysql_fdw_15-2.9.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/mysql_fdw_15-2.9.3-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 137.5 KiB | [mysql_fdw_15-2.9.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/mysql_fdw_15-2.9.2-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_15` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 135.6 KiB | [mysql_fdw_15-2.9.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/mysql_fdw_15-2.9.3-3PGDG.rhel10.aarch64.rpm) |
| `mysql_fdw_15` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 135.5 KiB | [mysql_fdw_15-2.9.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/mysql_fdw_15-2.9.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-mysql-fdw` | `2.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.1 KiB | [postgresql-15-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-15-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-mysql-fdw` | `2.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 121.0 KiB | [postgresql-15-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-15-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-mysql-fdw` | `2.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.7 KiB | [postgresql-15-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-15-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-mysql-fdw` | `2.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 121.6 KiB | [postgresql-15-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-15-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-mysql-fdw` | `2.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 146.0 KiB | [postgresql-15-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-15-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-mysql-fdw` | `2.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 142.7 KiB | [postgresql-15-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-15-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-mysql-fdw` | `2.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.3 KiB | [postgresql-15-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-15-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-mysql-fdw` | `2.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.8 KiB | [postgresql-15-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-15-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mysql_fdw_14` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 136.2 KiB | [mysql_fdw_14-2.9.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mysql_fdw_14-2.9.3-3PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.4 KiB | [mysql_fdw_14-2.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mysql_fdw_14-2.9.2-1PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 155.9 KiB | [mysql_fdw_14-2.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mysql_fdw_14-2.9.1-1PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 155.3 KiB | [mysql_fdw_14-2.9.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mysql_fdw_14-2.9.0-1.rhel8.x86_64.rpm) |
| `mysql_fdw_14` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 138.8 KiB | [mysql_fdw_14-2.8.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mysql_fdw_14-2.8.0-2.rhel8.x86_64.rpm) |
| `mysql_fdw_14` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.6 KiB | [mysql_fdw_14-2.7.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mysql_fdw_14-2.7.0-1.rhel8.x86_64.rpm) |
| `mysql_fdw_14` | `2.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 122.6 KiB | [mysql_fdw_14-2.6.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mysql_fdw_14-2.6.1-2.rhel8.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.3 KiB | [mysql_fdw_14-2.9.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/mysql_fdw_14-2.9.3-3PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 137.2 KiB | [mysql_fdw_14-2.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/mysql_fdw_14-2.9.2-1PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.1 KiB | [mysql_fdw_14-2.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/mysql_fdw_14-2.9.1-1PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 152.8 KiB | [mysql_fdw_14-2.9.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/mysql_fdw_14-2.9.0-1.rhel8.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 136.3 KiB | [mysql_fdw_14-2.9.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mysql_fdw_14-2.9.3-3PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 140.2 KiB | [mysql_fdw_14-2.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mysql_fdw_14-2.9.2-1PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 156.9 KiB | [mysql_fdw_14-2.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mysql_fdw_14-2.9.1-1PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 156.6 KiB | [mysql_fdw_14-2.9.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mysql_fdw_14-2.9.0-1.rhel9.x86_64.rpm) |
| `mysql_fdw_14` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 140.4 KiB | [mysql_fdw_14-2.8.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mysql_fdw_14-2.8.0-2.rhel9.x86_64.rpm) |
| `mysql_fdw_14` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 130.3 KiB | [mysql_fdw_14-2.7.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mysql_fdw_14-2.7.0-1.rhel9.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 134.3 KiB | [mysql_fdw_14-2.9.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/mysql_fdw_14-2.9.3-3PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 138.8 KiB | [mysql_fdw_14-2.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/mysql_fdw_14-2.9.2-1PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 136.9 KiB | [mysql_fdw_14-2.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/mysql_fdw_14-2.9.1-1PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.4 KiB | [mysql_fdw_14-2.9.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/mysql_fdw_14-2.9.0-1.rhel9.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 138.0 KiB | [mysql_fdw_14-2.9.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/mysql_fdw_14-2.9.3-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 138.0 KiB | [mysql_fdw_14-2.9.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/mysql_fdw_14-2.9.2-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_14` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 135.9 KiB | [mysql_fdw_14-2.9.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/mysql_fdw_14-2.9.3-3PGDG.rhel10.aarch64.rpm) |
| `mysql_fdw_14` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 135.9 KiB | [mysql_fdw_14-2.9.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/mysql_fdw_14-2.9.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-mysql-fdw` | `2.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.7 KiB | [postgresql-14-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-14-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-mysql-fdw` | `2.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 121.3 KiB | [postgresql-14-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-14-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-mysql-fdw` | `2.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 125.0 KiB | [postgresql-14-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-14-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-mysql-fdw` | `2.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 122.0 KiB | [postgresql-14-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-14-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-mysql-fdw` | `2.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 146.3 KiB | [postgresql-14-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-14-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-mysql-fdw` | `2.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 143.0 KiB | [postgresql-14-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-14-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-mysql-fdw` | `2.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.5 KiB | [postgresql-14-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-14-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-mysql-fdw` | `2.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 121.1 KiB | [postgresql-14-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-14-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mysql_fdw_13` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.4 KiB | [mysql_fdw_13-2.9.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.9.3-3PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 136.1 KiB | [mysql_fdw_13-2.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.9.2-1PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.4 KiB | [mysql_fdw_13-2.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.9.1-1PGDG.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.8 KiB | [mysql_fdw_13-2.9.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.9.0-1.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 137.5 KiB | [mysql_fdw_13-2.8.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.8.0-2.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 128.5 KiB | [mysql_fdw_13-2.7.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.7.0-1.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 121.3 KiB | [mysql_fdw_13-2.6.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.6.1-1.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.5 KiB | [mysql_fdw_13-2.6.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.6.0-1.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.4 KiB | [mysql_fdw_13-2.5.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.5.5-1.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.6 KiB | [mysql_fdw_13-2.5.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/mysql_fdw_13-2.5.4-1.rhel8.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 130.1 KiB | [mysql_fdw_13-2.9.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/mysql_fdw_13-2.9.3-3PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 134.1 KiB | [mysql_fdw_13-2.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/mysql_fdw_13-2.9.2-1PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 131.5 KiB | [mysql_fdw_13-2.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/mysql_fdw_13-2.9.1-1PGDG.rhel8.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.7 KiB | [mysql_fdw_13-2.9.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/mysql_fdw_13-2.9.0-1.rhel8.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 133.0 KiB | [mysql_fdw_13-2.9.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/mysql_fdw_13-2.9.3-3PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 137.1 KiB | [mysql_fdw_13-2.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/mysql_fdw_13-2.9.2-1PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 152.8 KiB | [mysql_fdw_13-2.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/mysql_fdw_13-2.9.1-1PGDG.rhel9.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 152.6 KiB | [mysql_fdw_13-2.9.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/mysql_fdw_13-2.9.0-1.rhel9.x86_64.rpm) |
| `mysql_fdw_13` | `2.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 139.5 KiB | [mysql_fdw_13-2.8.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/mysql_fdw_13-2.8.0-2.rhel9.x86_64.rpm) |
| `mysql_fdw_13` | `2.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 129.7 KiB | [mysql_fdw_13-2.7.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/mysql_fdw_13-2.7.0-1.rhel9.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 130.9 KiB | [mysql_fdw_13-2.9.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/mysql_fdw_13-2.9.3-3PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 135.6 KiB | [mysql_fdw_13-2.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/mysql_fdw_13-2.9.2-1PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 133.3 KiB | [mysql_fdw_13-2.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/mysql_fdw_13-2.9.1-1PGDG.rhel9.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 150.4 KiB | [mysql_fdw_13-2.9.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/mysql_fdw_13-2.9.0-1.rhel9.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 134.6 KiB | [mysql_fdw_13-2.9.3-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/mysql_fdw_13-2.9.3-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 134.6 KiB | [mysql_fdw_13-2.9.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/mysql_fdw_13-2.9.2-3PGDG.rhel10.x86_64.rpm) |
| `mysql_fdw_13` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 132.6 KiB | [mysql_fdw_13-2.9.3-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/mysql_fdw_13-2.9.3-3PGDG.rhel10.aarch64.rpm) |
| `mysql_fdw_13` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 132.6 KiB | [mysql_fdw_13-2.9.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/mysql_fdw_13-2.9.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-mysql-fdw` | `2.9.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 121.2 KiB | [postgresql-13-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-13-mysql-fdw_2.9.3-1.pgdg12+1_amd64.deb) |
| `postgresql-13-mysql-fdw` | `2.9.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 118.0 KiB | [postgresql-13-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-13-mysql-fdw_2.9.3-1.pgdg12+1_arm64.deb) |
| `postgresql-13-mysql-fdw` | `2.9.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 121.8 KiB | [postgresql-13-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-13-mysql-fdw_2.9.3-1.pgdg13+1_amd64.deb) |
| `postgresql-13-mysql-fdw` | `2.9.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 118.7 KiB | [postgresql-13-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-13-mysql-fdw_2.9.3-1.pgdg13+1_arm64.deb) |
| `postgresql-13-mysql-fdw` | `2.9.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 142.5 KiB | [postgresql-13-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-13-mysql-fdw_2.9.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-mysql-fdw` | `2.9.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 139.2 KiB | [postgresql-13-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-13-mysql-fdw_2.9.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-mysql-fdw` | `2.9.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 121.0 KiB | [postgresql-13-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-13-mysql-fdw_2.9.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-mysql-fdw` | `2.9.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 117.9 KiB | [postgresql-13-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-mysql-fdw/postgresql-13-mysql-fdw_2.9.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/mysql_fdw" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/mysql_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="mysql_fdw-REL-2_9_3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg mysql_fdw;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install mysql_fdw;		# install via package name, for the active PG version

pig install mysql_fdw -v 18;   # install for PG 18
pig install mysql_fdw -v 17;   # install for PG 17
pig install mysql_fdw -v 16;   # install for PG 16
pig install mysql_fdw -v 15;   # install for PG 15
pig install mysql_fdw -v 14;   # install for PG 14
pig install mysql_fdw -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION mysql_fdw;
```
