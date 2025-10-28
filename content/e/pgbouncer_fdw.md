---
title: "pgbouncer_fdw"
linkTitle: "pgbouncer_fdw"
description: "Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions"
weight: 8650
categories: ["FDW"]
width: full
---

Extension for querying PgBouncer stats from normal SQL views & running pgbouncer commands from normal SQL functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8650** | {{< badge content="pgbouncer_fdw" link="https://github.com/CrunchyData/pgbouncer_fdw" >}} | {{< ext "pgbouncer_fdw" >}} | `1.4.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "dblink" >}} {{< ext "postgres_fdw" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_stat_statements" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgbouncer_fdw" >}} | `1.4.0` | {{< bg "18" "pgbouncer_fdw_18" "green" >}} {{< bg "17" "pgbouncer_fdw_17" "green" >}} {{< bg "16" "pgbouncer_fdw_16" "green" >}} {{< bg "15" "pgbouncer_fdw_15" "green" >}} {{< bg "14" "pgbouncer_fdw_14" "green" >}} {{< bg "13" "pgbouncer_fdw_13" "green" >}} | `pgbouncer_fdw_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_13 : AVAIL 8" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_13 : AVAIL 6" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_13 : AVAIL 6" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_13 : AVAIL 6" "blue" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_13 : AVAIL 1" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "pgbouncer_fdw_13 : AVAIL 1" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "pgbouncer_fdw : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_18` | 1.4.0 | `el8.x86_64` | pgdg | 24.0 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_18` | 1.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_18` | 1.4.0 | `el9.x86_64` | pgdg | 21.9 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_18` | 1.4.0 | `el9.aarch64` | pgdg | 21.8 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_18` | 1.4.0 | `el10.x86_64` | pgdg | 22.4 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_18` | 1.4.0 | `el10.aarch64` | pgdg | 22.4 KiB | [pgbouncer_fdw_18-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgbouncer_fdw_18-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_17` | 1.4.0 | `el8.x86_64` | pgdg | 24.0 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.3.0 | `el8.x86_64` | pgdg | 23.5 KiB | [pgbouncer_fdw_17-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgbouncer_fdw_17-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.2.0 | `el8.x86_64` | pgdg | 21.3 KiB | [pgbouncer_fdw_17-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgbouncer_fdw_17-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.1.0 | `el8.x86_64` | pgdg | 19.7 KiB | [pgbouncer_fdw_17-1.1.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgbouncer_fdw_17-1.1.0-2PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_17` | 1.3.0 | `el8.aarch64` | pgdg | 23.4 KiB | [pgbouncer_fdw_17-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgbouncer_fdw_17-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_17` | 1.2.0 | `el8.aarch64` | pgdg | 21.2 KiB | [pgbouncer_fdw_17-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgbouncer_fdw_17-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_17` | 1.1.0 | `el8.aarch64` | pgdg | 19.7 KiB | [pgbouncer_fdw_17-1.1.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgbouncer_fdw_17-1.1.0-2PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_17` | 1.4.0 | `el9.x86_64` | pgdg | 21.9 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.3.0 | `el9.x86_64` | pgdg | 21.5 KiB | [pgbouncer_fdw_17-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgbouncer_fdw_17-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.2.0 | `el9.x86_64` | pgdg | 19.8 KiB | [pgbouncer_fdw_17-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgbouncer_fdw_17-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.1.0 | `el9.x86_64` | pgdg | 18.6 KiB | [pgbouncer_fdw_17-1.1.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgbouncer_fdw_17-1.1.0-2PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.4.0 | `el9.aarch64` | pgdg | 21.8 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_17` | 1.3.0 | `el9.aarch64` | pgdg | 21.4 KiB | [pgbouncer_fdw_17-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgbouncer_fdw_17-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_17` | 1.2.0 | `el9.aarch64` | pgdg | 19.8 KiB | [pgbouncer_fdw_17-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgbouncer_fdw_17-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_17` | 1.1.0 | `el9.aarch64` | pgdg | 18.5 KiB | [pgbouncer_fdw_17-1.1.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgbouncer_fdw_17-1.1.0-2PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_17` | 1.4.0 | `el10.x86_64` | pgdg | 22.4 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_17` | 1.4.0 | `el10.aarch64` | pgdg | 22.4 KiB | [pgbouncer_fdw_17-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgbouncer_fdw_17-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_16` | 1.4.0 | `el8.x86_64` | pgdg | 24.0 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.3.0 | `el8.x86_64` | pgdg | 23.5 KiB | [pgbouncer_fdw_16-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.2.0 | `el8.x86_64` | pgdg | 21.3 KiB | [pgbouncer_fdw_16-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.1.0 | `el8.x86_64` | pgdg | 19.6 KiB | [pgbouncer_fdw_16-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.0.1 | `el8.x86_64` | pgdg | 18.9 KiB | [pgbouncer_fdw_16-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | 0.4 | `el8.x86_64` | pgdg | 13.5 KiB | [pgbouncer_fdw_16-0.4-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgbouncer_fdw_16-0.4-3.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.3.0 | `el8.aarch64` | pgdg | 23.4 KiB | [pgbouncer_fdw_16-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.2.0 | `el8.aarch64` | pgdg | 21.2 KiB | [pgbouncer_fdw_16-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.1.0 | `el8.aarch64` | pgdg | 19.6 KiB | [pgbouncer_fdw_16-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.0.1 | `el8.aarch64` | pgdg | 18.9 KiB | [pgbouncer_fdw_16-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | 0.4 | `el8.aarch64` | pgdg | 13.4 KiB | [pgbouncer_fdw_16-0.4-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgbouncer_fdw_16-0.4-3.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.4.0 | `el9.x86_64` | pgdg | 21.9 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.3.0 | `el9.x86_64` | pgdg | 21.5 KiB | [pgbouncer_fdw_16-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.2.0 | `el9.x86_64` | pgdg | 19.8 KiB | [pgbouncer_fdw_16-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.1.0 | `el9.x86_64` | pgdg | 18.5 KiB | [pgbouncer_fdw_16-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.0.1 | `el9.x86_64` | pgdg | 18.0 KiB | [pgbouncer_fdw_16-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | 0.4 | `el9.x86_64` | pgdg | 13.0 KiB | [pgbouncer_fdw_16-0.4-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgbouncer_fdw_16-0.4-3.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.4.0 | `el9.aarch64` | pgdg | 21.8 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.3.0 | `el9.aarch64` | pgdg | 21.4 KiB | [pgbouncer_fdw_16-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.2.0 | `el9.aarch64` | pgdg | 19.7 KiB | [pgbouncer_fdw_16-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.1.0 | `el9.aarch64` | pgdg | 18.3 KiB | [pgbouncer_fdw_16-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.0.1 | `el9.aarch64` | pgdg | 17.9 KiB | [pgbouncer_fdw_16-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | 0.4 | `el9.aarch64` | pgdg | 12.8 KiB | [pgbouncer_fdw_16-0.4-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgbouncer_fdw_16-0.4-3.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_16` | 1.4.0 | `el10.x86_64` | pgdg | 22.4 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_16` | 1.4.0 | `el10.aarch64` | pgdg | 22.4 KiB | [pgbouncer_fdw_16-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgbouncer_fdw_16-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_15` | 1.4.0 | `el8.x86_64` | pgdg | 24.0 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.3.0 | `el8.x86_64` | pgdg | 23.5 KiB | [pgbouncer_fdw_15-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.2.0 | `el8.x86_64` | pgdg | 21.3 KiB | [pgbouncer_fdw_15-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.1.0 | `el8.x86_64` | pgdg | 19.6 KiB | [pgbouncer_fdw_15-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.0.1 | `el8.x86_64` | pgdg | 18.9 KiB | [pgbouncer_fdw_15-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | 0.4 | `el8.x86_64` | pgdg | 13.4 KiB | [pgbouncer_fdw_15-0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgbouncer_fdw_15-0.4-2.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.3.0 | `el8.aarch64` | pgdg | 23.4 KiB | [pgbouncer_fdw_15-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.2.0 | `el8.aarch64` | pgdg | 21.2 KiB | [pgbouncer_fdw_15-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.1.0 | `el8.aarch64` | pgdg | 19.6 KiB | [pgbouncer_fdw_15-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.0.1 | `el8.aarch64` | pgdg | 18.9 KiB | [pgbouncer_fdw_15-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | 0.4 | `el8.aarch64` | pgdg | 13.3 KiB | [pgbouncer_fdw_15-0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgbouncer_fdw_15-0.4-2.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.4.0 | `el9.x86_64` | pgdg | 21.9 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.3.0 | `el9.x86_64` | pgdg | 21.5 KiB | [pgbouncer_fdw_15-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.2.0 | `el9.x86_64` | pgdg | 19.8 KiB | [pgbouncer_fdw_15-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.1.0 | `el9.x86_64` | pgdg | 18.5 KiB | [pgbouncer_fdw_15-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.0.1 | `el9.x86_64` | pgdg | 18.0 KiB | [pgbouncer_fdw_15-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | 0.4 | `el9.x86_64` | pgdg | 13.3 KiB | [pgbouncer_fdw_15-0.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgbouncer_fdw_15-0.4-2.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.4.0 | `el9.aarch64` | pgdg | 21.8 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.3.0 | `el9.aarch64` | pgdg | 21.4 KiB | [pgbouncer_fdw_15-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.2.0 | `el9.aarch64` | pgdg | 19.8 KiB | [pgbouncer_fdw_15-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.1.0 | `el9.aarch64` | pgdg | 18.3 KiB | [pgbouncer_fdw_15-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.0.1 | `el9.aarch64` | pgdg | 17.9 KiB | [pgbouncer_fdw_15-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | 0.4 | `el9.aarch64` | pgdg | 13.1 KiB | [pgbouncer_fdw_15-0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgbouncer_fdw_15-0.4-2.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_15` | 1.4.0 | `el10.x86_64` | pgdg | 22.4 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_15` | 1.4.0 | `el10.aarch64` | pgdg | 22.4 KiB | [pgbouncer_fdw_15-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgbouncer_fdw_15-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_14` | 1.4.0 | `el8.x86_64` | pgdg | 24.0 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.3.0 | `el8.x86_64` | pgdg | 23.5 KiB | [pgbouncer_fdw_14-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.2.0 | `el8.x86_64` | pgdg | 21.3 KiB | [pgbouncer_fdw_14-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.1.0 | `el8.x86_64` | pgdg | 19.6 KiB | [pgbouncer_fdw_14-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.0.1 | `el8.x86_64` | pgdg | 19.0 KiB | [pgbouncer_fdw_14-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | 0.4 | `el8.x86_64` | pgdg | 13.4 KiB | [pgbouncer_fdw_14-0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-0.4-2.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | 0.4 | `el8.x86_64` | pgdg | 13.3 KiB | [pgbouncer_fdw_14-0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-0.4-1.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | 0.3 | `el8.x86_64` | pgdg | 12.2 KiB | [pgbouncer_fdw_14-0.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgbouncer_fdw_14-0.3-2.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.3.0 | `el8.aarch64` | pgdg | 23.4 KiB | [pgbouncer_fdw_14-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.2.0 | `el8.aarch64` | pgdg | 21.2 KiB | [pgbouncer_fdw_14-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.1.0 | `el8.aarch64` | pgdg | 19.6 KiB | [pgbouncer_fdw_14-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.0.1 | `el8.aarch64` | pgdg | 18.9 KiB | [pgbouncer_fdw_14-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | 0.4 | `el8.aarch64` | pgdg | 13.3 KiB | [pgbouncer_fdw_14-0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgbouncer_fdw_14-0.4-2.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.4.0 | `el9.x86_64` | pgdg | 21.9 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.3.0 | `el9.x86_64` | pgdg | 21.5 KiB | [pgbouncer_fdw_14-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.2.0 | `el9.x86_64` | pgdg | 19.8 KiB | [pgbouncer_fdw_14-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.1.0 | `el9.x86_64` | pgdg | 18.5 KiB | [pgbouncer_fdw_14-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.0.1 | `el9.x86_64` | pgdg | 18.0 KiB | [pgbouncer_fdw_14-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | 0.4 | `el9.x86_64` | pgdg | 13.3 KiB | [pgbouncer_fdw_14-0.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgbouncer_fdw_14-0.4-2.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.4.0 | `el9.aarch64` | pgdg | 21.8 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.3.0 | `el9.aarch64` | pgdg | 21.4 KiB | [pgbouncer_fdw_14-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.2.0 | `el9.aarch64` | pgdg | 19.7 KiB | [pgbouncer_fdw_14-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.1.0 | `el9.aarch64` | pgdg | 18.4 KiB | [pgbouncer_fdw_14-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.0.1 | `el9.aarch64` | pgdg | 17.9 KiB | [pgbouncer_fdw_14-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | 0.4 | `el9.aarch64` | pgdg | 13.1 KiB | [pgbouncer_fdw_14-0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgbouncer_fdw_14-0.4-2.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_14` | 1.4.0 | `el10.x86_64` | pgdg | 22.4 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_14` | 1.4.0 | `el10.aarch64` | pgdg | 22.4 KiB | [pgbouncer_fdw_14-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgbouncer_fdw_14-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgbouncer_fdw_13` | 1.4.0 | `el8.x86_64` | pgdg | 24.0 KiB | [pgbouncer_fdw_13-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgbouncer_fdw_13-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.3.0 | `el8.x86_64` | pgdg | 23.5 KiB | [pgbouncer_fdw_13-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgbouncer_fdw_13-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.2.0 | `el8.x86_64` | pgdg | 21.3 KiB | [pgbouncer_fdw_13-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgbouncer_fdw_13-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.1.0 | `el8.x86_64` | pgdg | 19.6 KiB | [pgbouncer_fdw_13-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgbouncer_fdw_13-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.0.1 | `el8.x86_64` | pgdg | 19.0 KiB | [pgbouncer_fdw_13-1.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgbouncer_fdw_13-1.0.1-1PGDG.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_13` | 0.4 | `el8.x86_64` | pgdg | 13.3 KiB | [pgbouncer_fdw_13-0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgbouncer_fdw_13-0.4-1.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_13` | 0.4 | `el8.x86_64` | pgdg | 13.4 KiB | [pgbouncer_fdw_13-0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgbouncer_fdw_13-0.4-2.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_13` | 0.3 | `el8.x86_64` | pgdg | 12.2 KiB | [pgbouncer_fdw_13-0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgbouncer_fdw_13-0.3-1.rhel8.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.4.0 | `el8.aarch64` | pgdg | 23.9 KiB | [pgbouncer_fdw_13-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgbouncer_fdw_13-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.3.0 | `el8.aarch64` | pgdg | 23.4 KiB | [pgbouncer_fdw_13-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgbouncer_fdw_13-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.2.0 | `el8.aarch64` | pgdg | 21.2 KiB | [pgbouncer_fdw_13-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgbouncer_fdw_13-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.1.0 | `el8.aarch64` | pgdg | 19.6 KiB | [pgbouncer_fdw_13-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgbouncer_fdw_13-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.0.1 | `el8.aarch64` | pgdg | 18.9 KiB | [pgbouncer_fdw_13-1.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgbouncer_fdw_13-1.0.1-1PGDG.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_13` | 0.4 | `el8.aarch64` | pgdg | 13.3 KiB | [pgbouncer_fdw_13-0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgbouncer_fdw_13-0.4-2.rhel8.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.4.0 | `el9.x86_64` | pgdg | 21.9 KiB | [pgbouncer_fdw_13-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgbouncer_fdw_13-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.3.0 | `el9.x86_64` | pgdg | 21.5 KiB | [pgbouncer_fdw_13-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgbouncer_fdw_13-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.2.0 | `el9.x86_64` | pgdg | 19.8 KiB | [pgbouncer_fdw_13-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgbouncer_fdw_13-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.1.0 | `el9.x86_64` | pgdg | 18.5 KiB | [pgbouncer_fdw_13-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgbouncer_fdw_13-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.0.1 | `el9.x86_64` | pgdg | 18.0 KiB | [pgbouncer_fdw_13-1.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgbouncer_fdw_13-1.0.1-1PGDG.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_13` | 0.4 | `el9.x86_64` | pgdg | 13.3 KiB | [pgbouncer_fdw_13-0.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgbouncer_fdw_13-0.4-2.rhel9.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.4.0 | `el9.aarch64` | pgdg | 21.8 KiB | [pgbouncer_fdw_13-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgbouncer_fdw_13-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.3.0 | `el9.aarch64` | pgdg | 21.4 KiB | [pgbouncer_fdw_13-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgbouncer_fdw_13-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.2.0 | `el9.aarch64` | pgdg | 19.8 KiB | [pgbouncer_fdw_13-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgbouncer_fdw_13-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.1.0 | `el9.aarch64` | pgdg | 18.3 KiB | [pgbouncer_fdw_13-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgbouncer_fdw_13-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.0.1 | `el9.aarch64` | pgdg | 17.9 KiB | [pgbouncer_fdw_13-1.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgbouncer_fdw_13-1.0.1-1PGDG.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_13` | 0.4 | `el9.aarch64` | pgdg | 13.1 KiB | [pgbouncer_fdw_13-0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgbouncer_fdw_13-0.4-2.rhel9.aarch64.rpm) |
| `pgbouncer_fdw_13` | 1.4.0 | `el10.x86_64` | pgdg | 22.4 KiB | [pgbouncer_fdw_13-1.4.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgbouncer_fdw_13-1.4.0-1PGDG.rhel10.x86_64.rpm) |
| `pgbouncer_fdw_13` | 1.4.0 | `el10.aarch64` | pgdg | 22.4 KiB | [pgbouncer_fdw_13-1.4.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgbouncer_fdw_13-1.4.0-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pgbouncer_fdw" title="Repository" icon="github" subtitle="github.com/CrunchyData/pgbouncer_fdw" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgbouncer_fdw; # install by extension name, for the current active PG version
pig ext install pgbouncer_fdw; # install via package alias, for the active PG version
pig ext install pgbouncer_fdw -v 18;   # install for PG 18
pig ext install pgbouncer_fdw -v 17;   # install for PG 17
pig ext install pgbouncer_fdw -v 16;   # install for PG 16
pig ext install pgbouncer_fdw -v 15;   # install for PG 15
pig ext install pgbouncer_fdw -v 14;   # install for PG 14
pig ext install pgbouncer_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgbouncer_fdw;
```

