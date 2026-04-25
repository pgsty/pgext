---
title: "pg_dbms_errlog"
linkTitle: "pg_dbms_errlog"
description: "Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table."
weight: 9270
categories: ["SIM"]
width: full
---

[**pg_dbms_errlog**](https://github.com/HexaCluster/pg_dbms_errlog) : Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9270** | {{< badge content="pg_dbms_errlog" link="https://github.com/HexaCluster/pg_dbms_errlog" >}} | {{< ext "pg_dbms_errlog" >}} | `2.2` | {{< category "SIM" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `dbms_errlog` |
|   **See Also**    | {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_dbms_errlog` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2` | {{< bg "18" "pg_dbms_errlog_18" "green" >}} {{< bg "17" "pg_dbms_errlog_17" "green" >}} {{< bg "16" "pg_dbms_errlog_16" "green" >}} {{< bg "15" "pg_dbms_errlog_15" "green" >}} {{< bg "14" "pg_dbms_errlog_14" "green" >}} | `pg_dbms_errlog_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_18` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.6 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_18` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_17` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.6 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_17` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_16` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.0 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.6 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_16` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.2 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_15` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.6 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.8 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.7 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.5 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_15` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_14` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.5 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.8 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.6 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.4 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_14` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.1 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_dbms_errlog" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_dbms_errlog" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_dbms_errlog;		# install via package name, for the active PG version

pig install pg_dbms_errlog -v 18;   # install for PG 18
pig install pg_dbms_errlog -v 17;   # install for PG 17
pig install pg_dbms_errlog -v 16;   # install for PG 16
pig install pg_dbms_errlog -v 15;   # install for PG 15
pig install pg_dbms_errlog -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_dbms_errlog';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_dbms_errlog;
```



## Usage

> [pg_dbms_errlog: Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table](https://github.com/HexaCluster/pg_dbms_errlog)

Enables DML operations to continue after encountering errors by logging failures to an error table, rather than aborting the transaction.

### Enabling

Add to `shared_preload_libraries` in `postgresql.conf`:

```ini
shared_preload_libraries = 'pg_dbms_errlog'
```

```sql
CREATE EXTENSION pg_dbms_errlog;
LOAD 'pg_dbms_errlog';
```

### Create Error Log Table

```sql
BEGIN;
CALL dbms_errlog.create_error_log('employees');
END;
-- Creates table "ERR$_employees" with error logging columns

-- With custom name and schema:
BEGIN;
CALL dbms_errlog.create_error_log('hr.employees', '"ERRORS"."ERR$_EMPTABLE"');
END;
```

### Configuration

```sql
SET pg_dbms_errlog.enabled TO true;       -- enable error logging
SET pg_dbms_errlog.query_tag TO 'daily_load';  -- tag for identifying statements
SET pg_dbms_errlog.reject_limit TO 10;    -- max errors before rollback (-1=unlimited)
SET pg_dbms_errlog.synchronous TO 'transaction'; -- 'transaction', 'query', or 'off'
SET pg_dbms_errlog.no_client_error TO true;      -- suppress client error messages
```

### Usage with pg_statement_rollback

```sql
LOAD 'pg_dbms_errlog';
LOAD 'pg_statement_rollback';

CREATE TABLE hr.raises (emp_id integer, sal integer CHECK(sal > 8000));

BEGIN;
CALL dbms_errlog.create_error_log('hr.raises');
END;

SET pg_dbms_errlog.query_tag TO 'daily_load';
SET pg_dbms_errlog.reject_limit TO 10;
SET pg_dbms_errlog.enabled TO true;

BEGIN;
SET pg_statement_rollback.enabled TO on;
INSERT INTO hr.raises VALUES (145, 15400);  -- Success
INSERT INTO hr.raises VALUES (161, 7700);   -- Failure (logged)
ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
INSERT INTO hr.raises VALUES (175, 9680);   -- Success
COMMIT;
```

### Viewing Error Logs

```sql
SELECT * FROM "ERR$_raises";
-- pg_err_number$  | 23514
-- pg_err_mesg$    | new row for relation "raises" violates check constraint
-- pg_err_optyp$   | I
-- pg_err_tag$     | daily_load
-- pg_err_query$   | INSERT INTO hr.raises VALUES (161, 7700);
```

### Flush Queued Errors

```sql
SELECT dbms_errlog.publish_queue();
```
