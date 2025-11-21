---
title: "db2_fdw"
linkTitle: "db2_fdw"
description: "foreign data wrapper for DB2 access"
weight: 8630
categories: ["FDW"]
width: full
---

[**db2_fdw**](https://github.com/wolfgangbrandl/db2_fdw) : foreign data wrapper for DB2 access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8630** | {{< badge content="db2_fdw" link="https://github.com/wolfgangbrandl/db2_fdw" >}} | {{< ext "db2_fdw" >}} | `18.0.1` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "odbc_fdw" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "jdbc_fdw" >}} {{< ext "postgres_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `18.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `db2_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `18.0.1` | {{< bg "18" "db2_fdw_18*" "red" >}} {{< bg "17" "db2_fdw_17*" "red" >}} {{< bg "16" "db2_fdw_16*" "green" >}} {{< bg "15" "db2_fdw_15*" "green" >}} {{< bg "14" "db2_fdw_14*" "green" >}} {{< bg "13" "db2_fdw_13*" "green" >}} | `db2_fdw_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "db2_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 18.0.1" "db2_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 7.0.0" "db2_fdw_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "db2_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 7.0.0" "db2_fdw_13 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "db2_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 18.0.1" "db2_fdw_14 : AVAIL 3" "blue" >}} |      {{< bg "MISS" "db2_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "db2_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "db2_fdw : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_18` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.3 KiB | [db2_fdw_18-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/db2_fdw_18-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.4 KiB | [db2_fdw_18-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/db2_fdw_18-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.4 KiB | [db2_fdw_18-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/db2_fdw_18-18.0.1-2PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.3 KiB | [db2_fdw_18-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/db2_fdw_18-18.0.1-1PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_17` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 59.6 KiB | [db2_fdw_17-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.6 KiB | [db2_fdw_17-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.5 KiB | [db2_fdw_17-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_17` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [db2_fdw_17-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.4 KiB | [db2_fdw_17-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.3 KiB | [db2_fdw_17-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_17` | `7.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.8 KiB | [db2_fdw_17-7.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-7.0.0-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.3 KiB | [db2_fdw_17-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-18.0.1-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.4 KiB | [db2_fdw_17-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-18.0.1-2PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_16` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 59.6 KiB | [db2_fdw_16-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_16` | `6.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 59.4 KiB | [db2_fdw_16-6.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-6.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.5 KiB | [db2_fdw_16-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.6 KiB | [db2_fdw_16-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_16` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [db2_fdw_16-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_16` | `6.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 58.0 KiB | [db2_fdw_16-6.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-6.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.4 KiB | [db2_fdw_16-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.3 KiB | [db2_fdw_16-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_16` | `7.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.9 KiB | [db2_fdw_16-7.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-7.0.0-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.5 KiB | [db2_fdw_16-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-18.0.1-2PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.4 KiB | [db2_fdw_16-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-18.0.1-1PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_15` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.7 KiB | [db2_fdw_15-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_15` | `6.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.4 KiB | [db2_fdw_15-6.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-6.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.1 KiB | [db2_fdw_15-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.0 KiB | [db2_fdw_15-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_15` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 60.2 KiB | [db2_fdw_15-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_15` | `6.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.4 KiB | [db2_fdw_15-6.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-6.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.2 KiB | [db2_fdw_15-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.3 KiB | [db2_fdw_15-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_15` | `7.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 60.9 KiB | [db2_fdw_15-7.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-7.0.0-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.0 KiB | [db2_fdw_15-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-18.0.1-2PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 69.9 KiB | [db2_fdw_15-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-18.0.1-1PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_14` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.8 KiB | [db2_fdw_14-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `6.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.6 KiB | [db2_fdw_14-6.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-6.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 357.6 KiB | [db2_fdw_14-5.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-5.0.0-1.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.0 KiB | [db2_fdw_14-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.1 KiB | [db2_fdw_14-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 60.2 KiB | [db2_fdw_14-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `6.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.5 KiB | [db2_fdw_14-6.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-6.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 364.1 KiB | [db2_fdw_14-5.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-5.0.0-1.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.3 KiB | [db2_fdw_14-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.4 KiB | [db2_fdw_14-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `7.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 61.0 KiB | [db2_fdw_14-7.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-7.0.0-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.0 KiB | [db2_fdw_14-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-18.0.1-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.1 KiB | [db2_fdw_14-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-18.0.1-2PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_13` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.5 KiB | [db2_fdw_13-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/db2_fdw_13-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_13` | `6.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.3 KiB | [db2_fdw_13-6.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/db2_fdw_13-6.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_13` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 355.2 KiB | [db2_fdw_13-5.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/db2_fdw_13-5.0.0-1.rhel8.x86_64.rpm) |
| `db2_fdw_13` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 60.4 KiB | [db2_fdw_13-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/db2_fdw_13-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_13` | `6.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.5 KiB | [db2_fdw_13-6.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/db2_fdw_13-6.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_13` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 362.2 KiB | [db2_fdw_13-5.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/db2_fdw_13-5.0.0-1.rhel9.x86_64.rpm) |
| `db2_fdw_13` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 362.2 KiB | [db2_fdw_13-4.0.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/db2_fdw_13-4.0.0-2.rhel9.x86_64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/wolfgangbrandl/db2_fdw" title="Repository" icon="github" subtitle="github.com/wolfgangbrandl/db2_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="db2_fdw-6.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg db2_fdw;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install db2_fdw;		# install via package name, for the active PG version

pig install db2_fdw -v 18;   # install for PG 18
pig install db2_fdw -v 17;   # install for PG 17
pig install db2_fdw -v 16;   # install for PG 16
pig install db2_fdw -v 15;   # install for PG 15
pig install db2_fdw -v 14;   # install for PG 14
pig install db2_fdw -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION db2_fdw;
```
