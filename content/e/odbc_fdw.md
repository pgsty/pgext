---
title: "odbc_fdw"
linkTitle: "odbc_fdw"
description: "Foreign data wrapper for accessing remote databases using ODBC"
weight: 8520
categories: ["FDW"]
width: full
---

[**odbc_fdw**](https://github.com/CartoDB/odbc_fdw) : Foreign data wrapper for accessing remote databases using ODBC


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8520** | {{< badge content="odbc_fdw" link="https://github.com/CartoDB/odbc_fdw" >}} | {{< ext "odbc_fdw" >}} | `0.5.1` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "jdbc_fdw" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "postgres_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.5.1` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `odbc_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.5.1` | {{< bg "18" "odbc_fdw_18" "red" >}} {{< bg "17" "odbc_fdw_17" "green" >}} {{< bg "16" "odbc_fdw_16" "green" >}} {{< bg "15" "odbc_fdw_15" "green" >}} {{< bg "14" "odbc_fdw_14" "green" >}} {{< bg "13" "odbc_fdw_13" "green" >}} | `odbc_fdw_$v` | `unixODBC` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "odbc_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 0.5.1" "odbc_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "odbc_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 0.5.1" "odbc_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "odbc_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 0.5.1" "odbc_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "odbc_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 0.5.1" "odbc_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "odbc_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 0.5.1" "odbc_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "odbc_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 0.5.1" "odbc_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.1" "odbc_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "odbc_fdw : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_17` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.4 KiB | [odbc_fdw_17-0.5.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/odbc_fdw_17-0.5.1-2PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.6 KiB | [odbc_fdw_17-0.5.1-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/odbc_fdw_17-0.5.1-2PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.4 KiB | [odbc_fdw_17-0.5.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/odbc_fdw_17-0.5.1-2PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.9 KiB | [odbc_fdw_17-0.5.1-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/odbc_fdw_17-0.5.1-2PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [odbc_fdw_17-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/odbc_fdw_17-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_17-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/odbc_fdw_17-0.5.1-3PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_16-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/odbc_fdw_16-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.5 KiB | [odbc_fdw_16-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/odbc_fdw_16-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_16-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/odbc_fdw_16-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [odbc_fdw_16-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/odbc_fdw_16-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [odbc_fdw_16-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/odbc_fdw_16-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_16-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/odbc_fdw_16-0.5.1-3PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_15-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/odbc_fdw_15-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.5 KiB | [odbc_fdw_15-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/odbc_fdw_15-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_15-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/odbc_fdw_15-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.7 KiB | [odbc_fdw_15-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/odbc_fdw_15-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [odbc_fdw_15-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/odbc_fdw_15-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_15-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/odbc_fdw_15-0.5.1-3PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_14` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_14-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/odbc_fdw_14-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.5 KiB | [odbc_fdw_14-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/odbc_fdw_14-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [odbc_fdw_14-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/odbc_fdw_14-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [odbc_fdw_14-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/odbc_fdw_14-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.4 KiB | [odbc_fdw_14-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/odbc_fdw_14-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_14` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [odbc_fdw_14-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/odbc_fdw_14-0.5.1-3PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `odbc_fdw_13` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.0 KiB | [odbc_fdw_13-0.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/odbc_fdw_13-0.5.1-1PGDG.rhel8.x86_64.rpm) |
| `odbc_fdw_13` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.5 KiB | [odbc_fdw_13-0.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/odbc_fdw_13-0.5.1-1PGDG.rhel8.aarch64.rpm) |
| `odbc_fdw_13` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.2 KiB | [odbc_fdw_13-0.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/odbc_fdw_13-0.5.1-1PGDG.rhel9.x86_64.rpm) |
| `odbc_fdw_13` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.7 KiB | [odbc_fdw_13-0.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/odbc_fdw_13-0.5.1-1PGDG.rhel9.aarch64.rpm) |
| `odbc_fdw_13` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.5 KiB | [odbc_fdw_13-0.5.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/odbc_fdw_13-0.5.1-3PGDG.rhel10.x86_64.rpm) |
| `odbc_fdw_13` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.7 KiB | [odbc_fdw_13-0.5.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/odbc_fdw_13-0.5.1-3PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CartoDB/odbc_fdw" title="Repository" icon="github" subtitle="github.com/CartoDB/odbc_fdw" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install odbc_fdw;		# install via package name, for the active PG version

pig install odbc_fdw -v 17;   # install for PG 17
pig install odbc_fdw -v 16;   # install for PG 16
pig install odbc_fdw -v 15;   # install for PG 15
pig install odbc_fdw -v 14;   # install for PG 14
pig install odbc_fdw -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION odbc_fdw;
```
