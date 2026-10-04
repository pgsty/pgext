---
title: "db2_fdw"
linkTitle: "db2_fdw"
description: "foreign data wrapper for DB2 access"
weight: 8630
categories: ["FDW"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**db2_fdw**](https://github.com/pg-fdw/db2_fdw) : foreign data wrapper for DB2 access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8630** | {{< badge content="db2_fdw" link="https://github.com/pg-fdw/db2_fdw" >}} | {{< ext "db2_fdw" >}} | `18.2.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "db_migrator" >}} {{< ext "db2fce" >}} {{< ext "pg_statement_rollback" >}} {{< ext "mysql_fdw" >}} {{< ext "orafce" >}} {{< ext "postgres_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "sqlite_fdw" >}} {{< ext "informix_fdw" >}} |

> [!Note] Latest PGDG RPM/catalog version is 18.2.0; no DEB package is available.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `18.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `db2_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `18.2.0` | {{< bg "18" "db2_fdw_18" "green" >}} {{< bg "17" "db2_fdw_17" "green" >}} {{< bg "16" "db2_fdw_16" "green" >}} {{< bg "15" "db2_fdw_15" "green" >}} {{< bg "14" "db2_fdw_14" "green" >}} | `db2_fdw_$v` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_14 : AVAIL 8" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "db2_fdw_18 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_17 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_16 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_15 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_14 : AVAIL 8" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "db2_fdw_18 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_17 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_16 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_15 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 18.2.0" "db2_fdw_14 : AVAIL 6" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "db2_fdw_18 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_17 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_16 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_15 : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} | {{< bg "N/A" "db2_fdw : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_18` | `18.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 124.1 KiB | [db2_fdw_18-18.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/db2_fdw_18-18.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_18` | `18.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 79.3 KiB | [db2_fdw_18-18.1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/db2_fdw_18-18.1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_18` | `18.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 79.1 KiB | [db2_fdw_18-18.1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/db2_fdw_18-18.1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.6 KiB | [db2_fdw_18-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/db2_fdw_18-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.6 KiB | [db2_fdw_18-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/db2_fdw_18-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_18` | `18.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 115.0 KiB | [db2_fdw_18-18.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/db2_fdw_18-18.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `db2_fdw_18` | `18.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.5 KiB | [db2_fdw_18-18.1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/db2_fdw_18-18.1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_18` | `18.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.2 KiB | [db2_fdw_18-18.1.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/db2_fdw_18-18.1.1-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.4 KiB | [db2_fdw_18-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/db2_fdw_18-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.3 KiB | [db2_fdw_18-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/db2_fdw_18-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_18` | `18.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.1 KiB | [db2_fdw_18-18.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/db2_fdw_18-18.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `db2_fdw_18` | `18.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 73.4 KiB | [db2_fdw_18-18.1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/db2_fdw_18-18.1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_18` | `18.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 73.1 KiB | [db2_fdw_18-18.1.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/db2_fdw_18-18.1.1-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.4 KiB | [db2_fdw_18-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/db2_fdw_18-18.0.1-2PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_18` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.3 KiB | [db2_fdw_18-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/db2_fdw_18-18.0.1-1PGDG.rhel10.x86_64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_17` | `18.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 124.0 KiB | [db2_fdw_17-18.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-18.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_17` | `18.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 79.2 KiB | [db2_fdw_17-18.1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-18.1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_17` | `18.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 79.1 KiB | [db2_fdw_17-18.1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-18.1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.6 KiB | [db2_fdw_17-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.5 KiB | [db2_fdw_17-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_17` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 59.6 KiB | [db2_fdw_17-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/db2_fdw_17-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_17` | `18.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.7 KiB | [db2_fdw_17-18.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-18.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `db2_fdw_17` | `18.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.5 KiB | [db2_fdw_17-18.1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-18.1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_17` | `18.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.2 KiB | [db2_fdw_17-18.1.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-18.1.1-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.4 KiB | [db2_fdw_17-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.3 KiB | [db2_fdw_17-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_17` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [db2_fdw_17-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/db2_fdw_17-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_17` | `18.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.1 KiB | [db2_fdw_17-18.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-18.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `db2_fdw_17` | `18.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 73.4 KiB | [db2_fdw_17-18.1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-18.1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_17` | `18.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 73.1 KiB | [db2_fdw_17-18.1.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-18.1.1-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.4 KiB | [db2_fdw_17-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-18.0.1-2PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_17` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.3 KiB | [db2_fdw_17-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-18.0.1-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_17` | `7.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.8 KiB | [db2_fdw_17-7.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/db2_fdw_17-7.0.0-1PGDG.rhel10.x86_64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_16` | `18.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 124.0 KiB | [db2_fdw_16-18.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-18.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_16` | `18.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 79.4 KiB | [db2_fdw_16-18.1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-18.1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_16` | `18.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 79.1 KiB | [db2_fdw_16-18.1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-18.1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.6 KiB | [db2_fdw_16-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.5 KiB | [db2_fdw_16-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_16` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 59.6 KiB | [db2_fdw_16-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_16` | `6.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 59.4 KiB | [db2_fdw_16-6.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/db2_fdw_16-6.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_16` | `18.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.7 KiB | [db2_fdw_16-18.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-18.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `db2_fdw_16` | `18.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.5 KiB | [db2_fdw_16-18.1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-18.1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_16` | `18.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.2 KiB | [db2_fdw_16-18.1.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-18.1.1-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.4 KiB | [db2_fdw_16-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.3 KiB | [db2_fdw_16-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_16` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [db2_fdw_16-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_16` | `6.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 58.0 KiB | [db2_fdw_16-6.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/db2_fdw_16-6.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_16` | `18.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 117.2 KiB | [db2_fdw_16-18.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-18.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `db2_fdw_16` | `18.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 73.4 KiB | [db2_fdw_16-18.1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-18.1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_16` | `18.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 73.1 KiB | [db2_fdw_16-18.1.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-18.1.1-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.5 KiB | [db2_fdw_16-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-18.0.1-2PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_16` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.4 KiB | [db2_fdw_16-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-18.0.1-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_16` | `7.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.9 KiB | [db2_fdw_16-7.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/db2_fdw_16-7.0.0-1PGDG.rhel10.x86_64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_15` | `18.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 126.3 KiB | [db2_fdw_15-18.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-18.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_15` | `18.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 82.0 KiB | [db2_fdw_15-18.1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-18.1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_15` | `18.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 81.9 KiB | [db2_fdw_15-18.1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-18.1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.1 KiB | [db2_fdw_15-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.0 KiB | [db2_fdw_15-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_15` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.7 KiB | [db2_fdw_15-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_15` | `6.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.4 KiB | [db2_fdw_15-6.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/db2_fdw_15-6.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_15` | `18.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 119.6 KiB | [db2_fdw_15-18.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-18.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `db2_fdw_15` | `18.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.9 KiB | [db2_fdw_15-18.1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-18.1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_15` | `18.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.7 KiB | [db2_fdw_15-18.1.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-18.1.1-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.3 KiB | [db2_fdw_15-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.2 KiB | [db2_fdw_15-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_15` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 60.2 KiB | [db2_fdw_15-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_15` | `6.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.4 KiB | [db2_fdw_15-6.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/db2_fdw_15-6.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_15` | `18.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.7 KiB | [db2_fdw_15-18.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-18.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `db2_fdw_15` | `18.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.4 KiB | [db2_fdw_15-18.1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-18.1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_15` | `18.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.3 KiB | [db2_fdw_15-18.1.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-18.1.1-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.0 KiB | [db2_fdw_15-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-18.0.1-2PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_15` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 69.9 KiB | [db2_fdw_15-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-18.0.1-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_15` | `7.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 60.9 KiB | [db2_fdw_15-7.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/db2_fdw_15-7.0.0-1PGDG.rhel10.x86_64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db2_fdw_14` | `18.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 126.4 KiB | [db2_fdw_14-18.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-18.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_14` | `18.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 82.1 KiB | [db2_fdw_14-18.1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-18.1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_14` | `18.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 81.9 KiB | [db2_fdw_14-18.1.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-18.1.1-1PGDG.rhel8.10.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.1 KiB | [db2_fdw_14-18.0.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-18.0.1-2PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.0 KiB | [db2_fdw_14-18.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-18.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `7.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.8 KiB | [db2_fdw_14-7.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-7.0.0-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `6.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.6 KiB | [db2_fdw_14-6.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-6.0.1-1PGDG.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `5.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 357.6 KiB | [db2_fdw_14-5.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/db2_fdw_14-5.0.0-1.rhel8.x86_64.rpm) |
| `db2_fdw_14` | `18.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 119.6 KiB | [db2_fdw_14-18.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-18.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `db2_fdw_14` | `18.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 78.0 KiB | [db2_fdw_14-18.1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-18.1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_14` | `18.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.7 KiB | [db2_fdw_14-18.1.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-18.1.1-1PGDG.rhel9.7.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.4 KiB | [db2_fdw_14-18.0.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-18.0.1-2PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.3 KiB | [db2_fdw_14-18.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-18.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `7.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 60.2 KiB | [db2_fdw_14-7.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-7.0.0-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `6.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.5 KiB | [db2_fdw_14-6.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-6.0.1-1PGDG.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `5.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 364.1 KiB | [db2_fdw_14-5.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/db2_fdw_14-5.0.0-1.rhel9.x86_64.rpm) |
| `db2_fdw_14` | `18.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.6 KiB | [db2_fdw_14-18.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-18.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `db2_fdw_14` | `18.1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.5 KiB | [db2_fdw_14-18.1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-18.1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_14` | `18.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.3 KiB | [db2_fdw_14-18.1.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-18.1.1-1PGDG.rhel10.1.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.1 KiB | [db2_fdw_14-18.0.1-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-18.0.1-2PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_14` | `18.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 70.0 KiB | [db2_fdw_14-18.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-18.0.1-1PGDG.rhel10.x86_64.rpm) |
| `db2_fdw_14` | `7.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 61.0 KiB | [db2_fdw_14-7.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/db2_fdw_14-7.0.0-1PGDG.rhel10.x86_64.rpm) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pg-fdw/db2_fdw" title="Repository" icon="github" subtitle="github.com/pg-fdw/db2_fdw" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install db2_fdw;		# install via package name, for the active PG version

pig install db2_fdw -v 18;   # install for PG 18
pig install db2_fdw -v 17;   # install for PG 17
pig install db2_fdw -v 16;   # install for PG 16
pig install db2_fdw -v 15;   # install for PG 15
pig install db2_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION db2_fdw;
```

## Usage

Sources:

- [18.2.0 README](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/README.md)
- [18.2.0 control file](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/db2_fdw.control)
- [18.2.0 SQL API](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/sql/db2_fdw--18.2.0.sql)

`db2_fdw` 18.2.0 queries and modifies IBM Db2 tables through PostgreSQL foreign tables. It pushes down supported filters and only the required columns. Upstream requires PostgreSQL 10.1 or later and an IBM Db2 client 11.1 or later with the same architecture as PostgreSQL. The server process must be able to load that client and access the configured database; installing this extension does not supply the external client or a Db2 server.

### Connect and Import Tables

The example assumes that the Db2 client can connect to the catalogued SAMPLE database and that DB2INST1.EMPLOYEE exists. Install the extension as a superuser, then create the server and a role-specific mapping. Use the actual Db2 credentials instead of the example password:

```sql
CREATE EXTENSION db2_fdw;
CREATE SERVER db2srv FOREIGN DATA WRAPPER db2_fdw
  OPTIONS (dbserver 'SAMPLE');
CREATE USER MAPPING FOR CURRENT_USER SERVER db2srv
  OPTIONS (user 'db2inst1', password 'change-me');
CREATE SCHEMA db2_remote;
IMPORT FOREIGN SCHEMA "DB2INST1" LIMIT TO ("EMPLOYEE")
  FROM SERVER db2srv INTO db2_remote;
SELECT empno, firstname, lastname, salary
FROM db2_remote.employee
WHERE empno = '000010';
```

Importing obtains the Db2 column metadata required by the wrapper. Default smart case folding lowercases all-uppercase names. A separate application role needs USAGE on the foreign server, its own user mapping, and appropriate schema/table privileges. Avoid a PUBLIC mapping when credentials are meant for one role. Empty user/password strings select the upstream external-authentication path and depend on the Db2 client environment.

### Options and Writes

- Server `dbserver` selects the Db2 connection; `no_encoding_error` controls handling of encoding-conversion errors. `batch_size` is reserved in this release and should not be treated as working batch insertion.
- Table `schema` and `table` identify the remote object; `readonly` prohibits modifications. `prefetch` defaults to 100 and accepts 0–1024; `fetch_size` is accepted but currently fixed to 1. `sample_percent` controls ANALYZE sampling.
- Column `key` must identify every remote primary-key column for `UPDATE` and `DELETE`. Imported metadata includes `db2type`, `db2size`, `db2bytes`, `db2chars`, `db2scale`, `db2null` and `db2ccsid`; preserve it when altering imported tables.
- IMPORT FOREIGN SCHEMA options `case` and `readonly` control name folding and whether imported tables permit writes.

INSERT, UPDATE and DELETE additionally require Db2-side privileges. Review column mapping and remote key definitions before enabling writes; PostgreSQL declarations alone do not create a remote primary key. Unsupported filters are evaluated locally, so inspect EXPLAIN before assuming pushdown.

### Diagnostics and Transactions

```sql
SELECT db2_diag();
SELECT db2_diag('db2srv');
SELECT db2_close_connections();
```

`db2_diag()` reports local client/build diagnostics and optionally remote-server details. `db2_close_connections()` closes connections cached by the current backend; do not call it within a transaction that has modified Db2 data. Long-lived PostgreSQL sessions can retain remote connections and transaction resources.

### Types and Maintenance

Common mappings include character types to text/character, BLOB to bytea, integer types to PostgreSQL integer types, and DATE/TIMESTAMP/TIME to their corresponding types. The declared PostgreSQL type and width must accommodate the remote values; failed conversions surface at query time. The control version is 18.2.0 and is relocatable. Normal use requires no shared preload. Follow the versioned README for IBM client environment and connectivity requirements, and distinguish a package upgrade from validation against the actual external database.
