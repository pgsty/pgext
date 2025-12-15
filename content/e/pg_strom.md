---
title: "pg_strom"
linkTitle: "pg_strom"
description: "PG-Strom - big-data processing acceleration using GPU and NVME"
weight: 2530
categories: ["OLAP"]
width: full
---

[**pg_strom**](https://github.com/heterodb/pg-strom) : PG-Strom - big-data processing acceleration using GPU and NVME


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2530** | {{< badge content="pg_strom" link="https://github.com/heterodb/pg-strom" >}} | {{< ext "pg_strom" >}} | `6.0` | {{< category "OLAP" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "columnar" >}} {{< ext "citus" >}} {{< ext "pg_analytics" >}} {{< ext "citus_columnar" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_mooncake" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `6.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_strom` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `6.0` | {{< bg "18" "pg_strom_18*" "red" >}} {{< bg "17" "pg_strom_17*" "green" >}} {{< bg "16" "pg_strom_16*" "green" >}} {{< bg "15" "pg_strom_15*" "green" >}} {{< bg "14" "pg_strom_14*" "green" >}} {{< bg "13" "pg_strom_13*" "green" >}} | `pg_strom_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_strom_18 : MISS 0" "red" >}}      | {{< bg "PGDG 6.0" "pg_strom_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 6.0" "pg_strom_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 6.0" "pg_strom_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 3.5" "pg_strom_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.5" "pg_strom_13 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_strom_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_strom_18 : MISS 0" "red" >}}      | {{< bg "PGDG 6.0" "pg_strom_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 6.0" "pg_strom_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 6.0" "pg_strom_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 3.5" "pg_strom_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.5" "pg_strom_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_strom_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 6.1" "pg_strom_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.1" "pg_strom_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.1" "pg_strom_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 6.1" "pg_strom_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_strom_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pg_strom_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strom : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strom_18` | `6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 705.8 KiB | [pg_strom_18-6.1-1PGDG.el10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/pg_strom_18-6.1-1PGDG.el10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strom_17` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 565.1 KiB | [pg_strom_17-6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/pg_strom_17-6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_17` | `5.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 472.4 KiB | [pg_strom_17-5.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/pg_strom_17-5.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_17` | `6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 539.0 KiB | [pg_strom_17-6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/pg_strom_17-6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_17` | `5.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 467.3 KiB | [pg_strom_17-5.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/pg_strom_17-5.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_17` | `6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 705.4 KiB | [pg_strom_17-6.1-1PGDG.el10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/pg_strom_17-6.1-1PGDG.el10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strom_16` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 565.4 KiB | [pg_strom_16-6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/pg_strom_16-6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_16` | `5.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 472.8 KiB | [pg_strom_16-5.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/pg_strom_16-5.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_16` | `5.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 462.2 KiB | [pg_strom_16-5.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/pg_strom_16-5.1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_16` | `5.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 461.7 KiB | [pg_strom_16-5.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/pg_strom_16-5.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_16` | `5.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 MiB | [pg_strom_16-5.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/pg_strom_16-5.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_16` | `5.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 MiB | [pg_strom_16-5.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/pg_strom_16-5.0.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_16` | `5.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 MiB | [pg_strom_16-5.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/pg_strom_16-5.0.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_16` | `5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 MiB | [pg_strom_16-5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/pg_strom_16-5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_16` | `6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 539.0 KiB | [pg_strom_16-6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/pg_strom_16-6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_16` | `5.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 467.3 KiB | [pg_strom_16-5.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/pg_strom_16-5.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_16` | `5.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 455.8 KiB | [pg_strom_16-5.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/pg_strom_16-5.1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_16` | `5.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 455.2 KiB | [pg_strom_16-5.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/pg_strom_16-5.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_16` | `5.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 8.6 MiB | [pg_strom_16-5.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/pg_strom_16-5.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_16` | `5.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 8.6 MiB | [pg_strom_16-5.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/pg_strom_16-5.0.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_16` | `5.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 8.6 MiB | [pg_strom_16-5.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/pg_strom_16-5.0.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_16` | `5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 8.6 MiB | [pg_strom_16-5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/pg_strom_16-5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_16` | `6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 705.7 KiB | [pg_strom_16-6.1-1PGDG.el10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/pg_strom_16-6.1-1PGDG.el10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strom_15` | `6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 569.4 KiB | [pg_strom_15-6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_15` | `5.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 475.6 KiB | [pg_strom_15-5.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-5.2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_15` | `5.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 464.9 KiB | [pg_strom_15-5.1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-5.1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_15` | `5.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 464.7 KiB | [pg_strom_15-5.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-5.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_15` | `5.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 MiB | [pg_strom_15-5.0.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-5.0.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_15` | `5.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 MiB | [pg_strom_15-5.0.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-5.0.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_15` | `5.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 MiB | [pg_strom_15-5.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-5.0.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_15` | `5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 MiB | [pg_strom_15-5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_strom_15` | `3.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.6 MiB | [pg_strom_15-3.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-3.5-1.rhel8.x86_64.rpm) |
| `pg_strom_15` | `3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.4 MiB | [pg_strom_15-3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/pg_strom_15-3.4-1.rhel8.x86_64.rpm) |
| `pg_strom_15` | `6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 545.7 KiB | [pg_strom_15-6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_15` | `5.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 474.2 KiB | [pg_strom_15-5.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-5.2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_15` | `5.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 464.7 KiB | [pg_strom_15-5.1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-5.1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_15` | `5.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 463.8 KiB | [pg_strom_15-5.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-5.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_15` | `5.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 8.6 MiB | [pg_strom_15-5.0.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-5.0.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_15` | `5.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 8.6 MiB | [pg_strom_15-5.0.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-5.0.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_15` | `5.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 8.6 MiB | [pg_strom_15-5.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-5.0.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_15` | `5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 8.6 MiB | [pg_strom_15-5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_strom_15` | `3.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 MiB | [pg_strom_15-3.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-3.5-1.rhel9.x86_64.rpm) |
| `pg_strom_15` | `3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.0 MiB | [pg_strom_15-3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/pg_strom_15-3.4-1.rhel9.x86_64.rpm) |
| `pg_strom_15` | `6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 730.1 KiB | [pg_strom_15-6.1-1PGDG.el10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/pg_strom_15-6.1-1PGDG.el10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strom_14` | `3.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.6 MiB | [pg_strom_14-3.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/pg_strom_14-3.5-1.rhel8.x86_64.rpm) |
| `pg_strom_14` | `3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.4 MiB | [pg_strom_14-3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/pg_strom_14-3.4-1.rhel8.x86_64.rpm) |
| `pg_strom_14` | `3.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 MiB | [pg_strom_14-3.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/pg_strom_14-3.5-1.rhel9.x86_64.rpm) |
| `pg_strom_14` | `3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.0 MiB | [pg_strom_14-3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/pg_strom_14-3.4-1.rhel9.x86_64.rpm) |
| `pg_strom_14` | `3.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.0 MiB | [pg_strom_14-3.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/pg_strom_14-3.3.2-1.rhel9.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strom_13` | `3.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.6 MiB | [pg_strom_13-3.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/pg_strom_13-3.5-1.rhel8.x86_64.rpm) |
| `pg_strom_13` | `3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 28.4 MiB | [pg_strom_13-3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/pg_strom_13-3.4-1.rhel8.x86_64.rpm) |
| `pg_strom_13` | `3.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.4 MiB | [pg_strom_13-3.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/pg_strom_13-3.3.2-1.rhel8.x86_64.rpm) |
| `pg_strom_13` | `3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.0 MiB | [pg_strom_13-3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/pg_strom_13-3.1-1.rhel8.x86_64.rpm) |
| `pg_strom_13` | `3.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 MiB | [pg_strom_13-3.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/pg_strom_13-3.5-1.rhel9.x86_64.rpm) |
| `pg_strom_13` | `3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.0 MiB | [pg_strom_13-3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/pg_strom_13-3.4-1.rhel9.x86_64.rpm) |
| `pg_strom_13` | `3.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.0 MiB | [pg_strom_13-3.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/pg_strom_13-3.3.2-1.rhel9.x86_64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/heterodb/pg-strom" title="Repository" icon="github" subtitle="github.com/heterodb/pg-strom" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_strom;		# install via package name, for the active PG version

pig install pg_strom -v 17;   # install for PG 17
pig install pg_strom -v 16;   # install for PG 16
pig install pg_strom -v 15;   # install for PG 15
pig install pg_strom -v 14;   # install for PG 14
pig install pg_strom -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_strom;
```
