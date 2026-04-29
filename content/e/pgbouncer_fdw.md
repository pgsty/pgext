---
title: "pgbouncer_fdw"
linkTitle: "pgbouncer_fdw"
description: "Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions"
weight: 8650
categories: ["FDW"]
width: full
---

[**pgbouncer_fdw**](https://github.com/CrunchyData/pgbouncer_fdw) : Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8650** | {{< badge content="pgbouncer_fdw" link="https://github.com/CrunchyData/pgbouncer_fdw" >}} | {{< ext "pgbouncer_fdw" >}} | `1.4.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} |
|   **See Also**    | {{< ext "dblink" >}} {{< ext "postgres_fdw" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_stat_statements" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgbouncer_fdw` | `dblink` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.0` | {{< bg "18" "pgbouncer_fdw_18" "green" >}} {{< bg "17" "pgbouncer_fdw_17" "green" >}} {{< bg "16" "pgbouncer_fdw_16" "green" >}} {{< bg "15" "pgbouncer_fdw_15" "green" >}} {{< bg "14" "pgbouncer_fdw_14" "green" >}} | `pgbouncer_fdw_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 8" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_18` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_18` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_18` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.9 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_18` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.8 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_18` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.4 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_18` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_17` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.5 KiB | [pgbouncer_fdw_17-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgbouncer_fdw_17-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.3 KiB | [pgbouncer_fdw_17-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgbouncer_fdw_17-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [pgbouncer_fdw_17-1.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgbouncer_fdw_17-1.1.0-2PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_17` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgbouncer_fdw_17-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgbouncer_fdw_17-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pgbouncer_fdw_17-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgbouncer_fdw_17-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.7 KiB | [pgbouncer_fdw_17-1.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgbouncer_fdw_17-1.1.0-2PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.9 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.5 KiB | [pgbouncer_fdw_17-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgbouncer_fdw_17-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pgbouncer_fdw_17-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgbouncer_fdw_17-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.6 KiB | [pgbouncer_fdw_17-1.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgbouncer_fdw_17-1.1.0-2PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.8 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_17` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.4 KiB | [pgbouncer_fdw_17-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgbouncer_fdw_17-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.8 KiB | [pgbouncer_fdw_17-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgbouncer_fdw_17-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.5 KiB | [pgbouncer_fdw_17-1.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgbouncer_fdw_17-1.1.0-2PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_17` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.4 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_17` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_16` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.5 KiB | [pgbouncer_fdw_16-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.3 KiB | [pgbouncer_fdw_16-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pgbouncer_fdw_16-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.9 KiB | [pgbouncer_fdw_16-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | `0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.5 KiB | [pgbouncer_fdw_16-0.4-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-0.4-3.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgbouncer_fdw_16-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pgbouncer_fdw_16-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.6 KiB | [pgbouncer_fdw_16-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.9 KiB | [pgbouncer_fdw_16-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | `0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.4 KiB | [pgbouncer_fdw_16-0.4-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-0.4-3.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.9 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.5 KiB | [pgbouncer_fdw_16-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pgbouncer_fdw_16-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.5 KiB | [pgbouncer_fdw_16-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.0 KiB | [pgbouncer_fdw_16-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | `0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [pgbouncer_fdw_16-0.4-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-0.4-3.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.8 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.4 KiB | [pgbouncer_fdw_16-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.7 KiB | [pgbouncer_fdw_16-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.3 KiB | [pgbouncer_fdw_16-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.9 KiB | [pgbouncer_fdw_16-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | `0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [pgbouncer_fdw_16-0.4-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-0.4-3.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.4 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_16` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.5 KiB | [pgbouncer_fdw_15-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.3 KiB | [pgbouncer_fdw_15-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pgbouncer_fdw_15-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.9 KiB | [pgbouncer_fdw_15-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | `0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.4 KiB | [pgbouncer_fdw_15-0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-0.4-2.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgbouncer_fdw_15-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pgbouncer_fdw_15-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.6 KiB | [pgbouncer_fdw_15-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.9 KiB | [pgbouncer_fdw_15-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | `0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.3 KiB | [pgbouncer_fdw_15-0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-0.4-2.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.9 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.5 KiB | [pgbouncer_fdw_15-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pgbouncer_fdw_15-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.5 KiB | [pgbouncer_fdw_15-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.0 KiB | [pgbouncer_fdw_15-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | `0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.3 KiB | [pgbouncer_fdw_15-0.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-0.4-2.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.8 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.4 KiB | [pgbouncer_fdw_15-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.8 KiB | [pgbouncer_fdw_15-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.3 KiB | [pgbouncer_fdw_15-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.9 KiB | [pgbouncer_fdw_15-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | `0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.1 KiB | [pgbouncer_fdw_15-0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-0.4-2.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.4 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_15` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.5 KiB | [pgbouncer_fdw_14-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.3 KiB | [pgbouncer_fdw_14-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pgbouncer_fdw_14-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.0 KiB | [pgbouncer_fdw_14-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | `0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.4 KiB | [pgbouncer_fdw_14-0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-0.4-2.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | `0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [pgbouncer_fdw_14-0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-0.4-1.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | `0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.2 KiB | [pgbouncer_fdw_14-0.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-0.3-2.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [pgbouncer_fdw_14-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pgbouncer_fdw_14-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.6 KiB | [pgbouncer_fdw_14-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.9 KiB | [pgbouncer_fdw_14-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | `0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.3 KiB | [pgbouncer_fdw_14-0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-0.4-2.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.9 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.5 KiB | [pgbouncer_fdw_14-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pgbouncer_fdw_14-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.5 KiB | [pgbouncer_fdw_14-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.0 KiB | [pgbouncer_fdw_14-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | `0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.3 KiB | [pgbouncer_fdw_14-0.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-0.4-2.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.8 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.4 KiB | [pgbouncer_fdw_14-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.7 KiB | [pgbouncer_fdw_14-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.4 KiB | [pgbouncer_fdw_14-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.9 KiB | [pgbouncer_fdw_14-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | `0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.1 KiB | [pgbouncer_fdw_14-0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-0.4-2.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.4 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_14` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pgbouncer_fdw" title="Repository" icon="github" subtitle="github.com/CrunchyData/pgbouncer_fdw" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgbouncer_fdw;		# install via package name, for the active PG version

pig install pgbouncer_fdw -v 18;   # install for PG 18
pig install pgbouncer_fdw -v 17;   # install for PG 17
pig install pgbouncer_fdw -v 16;   # install for PG 16
pig install pgbouncer_fdw -v 15;   # install for PG 15
pig install pgbouncer_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgbouncer_fdw CASCADE; -- requires dblink
```




## Usage

> [pgbouncer_fdw: Extension for querying PgBouncer stats from normal SQL views and running PgBouncer commands from normal SQL functions](https://github.com/CrunchyData/pgbouncer_fdw)

### Create Server

```sql
CREATE EXTENSION pgbouncer_fdw;

CREATE SERVER pgbouncer FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host 'localhost', port '6432', dbname 'pgbouncer');
```

For multiple PgBouncer instances:

```sql
CREATE SERVER pgbouncer1 FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host '192.168.1.10', port '6432', dbname 'pgbouncer');
CREATE SERVER pgbouncer2 FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (host '192.168.1.11', port '6432', dbname 'pgbouncer');

INSERT INTO pgbouncer_fdw_targets (target_host) VALUES ('pgbouncer1'), ('pgbouncer2');
UPDATE pgbouncer_fdw_targets SET active = false WHERE target_host = 'pgbouncer';
```

### Create User Mapping

```sql
CREATE USER MAPPING FOR PUBLIC SERVER pgbouncer
  OPTIONS (user 'ccp_monitoring', password 'mypassword');
```

### Available Views

| View | Description |
|------|-------------|
| `pgbouncer_clients` | Client connection details |
| `pgbouncer_pools` | Connection pool statistics |
| `pgbouncer_servers` | Backend server status |
| `pgbouncer_stats` | Statistics summary |
| `pgbouncer_databases` | Database definitions |
| `pgbouncer_config` | Configuration parameters |
| `pgbouncer_lists` | Internal lists |
| `pgbouncer_dns_hosts` | DNS host cache |
| `pgbouncer_dns_zones` | DNS zone cache |
| `pgbouncer_sockets` | Socket information |
| `pgbouncer_users` | User configuration |

```sql
SELECT * FROM pgbouncer_pools;
SELECT * FROM pgbouncer_stats;
SELECT database, cl_active, cl_waiting, sv_active FROM pgbouncer_pools;
```

When monitoring multiple instances, each row includes a `pgbouncer_target_host` column identifying the source.

### Command Functions

Administrative functions (require explicit `GRANT EXECUTE`):

```sql
SELECT pgbouncer_command_reload();              -- Reload configuration
SELECT pgbouncer_command_pause('mydb');          -- Pause a database
SELECT pgbouncer_command_resume('mydb');         -- Resume a database
SELECT pgbouncer_command_kill('mydb');           -- Kill connections
SELECT pgbouncer_command_disable('mydb');        -- Disable a database
SELECT pgbouncer_command_enable('mydb');         -- Enable a database
SELECT pgbouncer_command_reconnect('mydb');      -- Reconnect to backend
SELECT pgbouncer_command_set('key', 'value');    -- Set a parameter
SELECT pgbouncer_command_shutdown();             -- Shutdown PgBouncer
SELECT pgbouncer_command_suspend();              -- Suspend operations
SELECT pgbouncer_command_wait_close('mydb');     -- Wait for connections to close
```

### Permissions

```sql
GRANT USAGE ON FOREIGN SERVER pgbouncer TO monitoring_user;
GRANT SELECT ON pgbouncer_pools TO monitoring_user;
GRANT SELECT ON pgbouncer_stats TO monitoring_user;
GRANT EXECUTE ON FUNCTION pgbouncer_command_reload() TO admin_user;
```
