---
title: "timescaledb"
linkTitle: "timescaledb"
description: "Enables scalable inserts and complex queries for time-series data"
weight: 1000
categories: ["TIME"]
languages: ["C"]
licenses: ["Timescale"]
repos: ["PIGSTY"]
page_width: full
---

[**timescaledb**](https://github.com/timescale/timescaledb) : Enables scalable inserts and complex queries for time-series data


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1000** | {{< badge content="timescaledb" link="https://github.com/timescale/timescaledb" >}} | {{< ext "timescaledb" >}} | `2.32.1` | {{< category "TIME" >}} | {{< license "Timescale" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLdt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `timescaledb_information` `timescaledb_experimental` `_timescaledb_catalog` `_timescaledb_config` `_timescaledb_functions` `_timescaledb_internal` `_timescaledb_cache` |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "pg_stl" >}} {{< ext "pg_partman" >}} {{< ext "pg_ttl_index" >}} {{< ext "citus" >}} {{< ext "pg_cron" >}} {{< ext "pg_ivm" >}} {{< ext "pg_incremental" >}} {{< ext "pg_trickle" >}} |

> [!Note] TimescaleDB 2.29 removes PostgreSQL 15 support


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.32.1` | {{< bg "19" "" "red" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} | `timescaledb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.32.1` | {{< bg "19" "timescaledb-tsl_19" "red" >}} {{< bg "18" "timescaledb-tsl_18" "green" >}} {{< bg "17" "timescaledb-tsl_17" "green" >}} {{< bg "16" "timescaledb-tsl_16" "green" >}} {{< bg "15" "timescaledb-tsl_15" "green" >}} | `timescaledb-tsl_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.32.1` | {{< bg "19" "postgresql-19-timescaledb-tsl" "red" >}} {{< bg "18" "postgresql-18-timescaledb-tsl" "green" >}} {{< bg "17" "postgresql-17-timescaledb-tsl" "green" >}} {{< bg "16" "postgresql-16-timescaledb-tsl" "green" >}} {{< bg "15" "postgresql-15-timescaledb-tsl" "green" >}} | `postgresql-$v-timescaledb-tsl` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "timescaledb-tsl_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_18 : AVAIL 28" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_17 : AVAIL 47" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_16 : AVAIL 53" "green" >}} | {{< bg "PIGSTY 2.28.3" "timescaledb-tsl_15 : AVAIL 42" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "timescaledb-tsl_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.28.3" "timescaledb-tsl_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "timescaledb-tsl_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.28.3" "timescaledb-tsl_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "timescaledb-tsl_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.28.3" "timescaledb-tsl_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "timescaledb-tsl_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 2.28.3" "timescaledb-tsl_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "timescaledb-tsl_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.32.1" "timescaledb-tsl_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.28.3" "timescaledb-tsl_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-19-timescaledb-tsl : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-18-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-17-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.32.1" "postgresql-16-timescaledb-tsl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.28.3" "postgresql-15-timescaledb-tsl : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_18` | `2.32.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_18-2.32.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_18-2.32.1-1PGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_18-2.32.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.32.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_18-2.32.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.32.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.31.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_18-2.31.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_18-2.31.0-1PGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.31.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_18-2.31.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.31.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.30.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.30.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.30.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.30.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.30.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.30.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.30.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.30.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.30.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.29.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.29.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.29.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.29.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.29.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.29.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.29.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.29.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.29.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.28.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.28.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.28.3-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.28.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.28.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.28.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.28.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.28.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.28.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.28.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_18-2.28.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.28.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.27.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 770.2 KiB | [timescaledb-tsl_18-2.27.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.27.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.27.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 769.9 KiB | [timescaledb-tsl_18-2.27.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.27.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.27.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 768.2 KiB | [timescaledb-tsl_18-2.27.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.27.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.26.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 753.8 KiB | [timescaledb-tsl_18-2.26.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.26.4-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.26.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 750.8 KiB | [timescaledb-tsl_18-2.26.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.26.3-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.26.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 751.5 KiB | [timescaledb-tsl_18-2.26.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.26.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.26.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 750.5 KiB | [timescaledb-tsl_18-2.26.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.26.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.26.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 749.8 KiB | [timescaledb-tsl_18-2.26.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.26.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.25.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 730.2 KiB | [timescaledb-tsl_18-2.25.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.25.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.25.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 727.5 KiB | [timescaledb-tsl_18-2.25.1-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.25.1-2PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.25.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 726.5 KiB | [timescaledb-tsl_18-2.25.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.25.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.23.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 733.6 KiB | [timescaledb-tsl_18-2.23.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.23.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.23.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 733.3 KiB | [timescaledb-tsl_18-2.23.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/timescaledb-tsl_18-2.23.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_18-2.32.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_18-2.32.1-1PGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_18` | `2.31.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_18-2.31.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_18-2.31.0-1PGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_18` | `2.32.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [timescaledb-tsl_18-2.32.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_18-2.32.1-1PGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 1.5 MiB | [timescaledb-tsl_18-2.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/timescaledb-tsl_18-2.32.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 1.5 MiB | [timescaledb-tsl_18-2.32.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/timescaledb-tsl_18-2.32.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.31.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [timescaledb-tsl_18-2.31.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_18-2.31.0-1PGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_18-2.32.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_18-2.32.1-1PGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_18` | `2.31.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_18-2.31.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_18-2.31.0-1PGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_18` | `2.32.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_18-2.32.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-tsl_18-2.32.1-1PGSTY.el10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_18-2.32.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/timescaledb-tsl_18-2.32.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_18-2.32.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/timescaledb-tsl_18-2.32.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.31.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_18-2.31.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-tsl_18-2.31.0-1PGSTY.el10.x86_64.rpm) |
| `timescaledb-tsl_18` | `2.32.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_18-2.32.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-tsl_18-2.32.1-1PGSTY.el10.aarch64.rpm) |
| `timescaledb-tsl_18` | `2.31.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_18-2.31.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-tsl_18-2.31.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.4 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.3 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.4 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.5 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.3 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-timescaledb-tsl` | `2.32.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/timescaledb-tsl/postgresql-18-timescaledb-tsl_2.32.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_17` | `2.32.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_17-2.32.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_17-2.32.1-1PGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_17-2.32.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.32.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_17-2.32.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.32.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.31.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_17-2.31.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_17-2.31.0-1PGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.31.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_17-2.31.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.31.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.30.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.30.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.30.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.30.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.30.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.30.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.30.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.30.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.30.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.29.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.29.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.29.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.29.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.29.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.29.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.29.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.29.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.29.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.28.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.28.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.28.3-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.28.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.28.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.28.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.28.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.28.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.28.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.28.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_17-2.28.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.28.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.27.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 769.4 KiB | [timescaledb-tsl_17-2.27.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.27.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.27.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 769.0 KiB | [timescaledb-tsl_17-2.27.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.27.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.27.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 767.8 KiB | [timescaledb-tsl_17-2.27.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.27.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.26.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 752.8 KiB | [timescaledb-tsl_17-2.26.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.26.4-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.26.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 749.8 KiB | [timescaledb-tsl_17-2.26.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.26.3-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.26.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 750.9 KiB | [timescaledb-tsl_17-2.26.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.26.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.26.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 749.7 KiB | [timescaledb-tsl_17-2.26.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.26.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.26.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 748.8 KiB | [timescaledb-tsl_17-2.26.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.26.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.25.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 729.1 KiB | [timescaledb-tsl_17-2.25.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.25.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.25.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 726.5 KiB | [timescaledb-tsl_17-2.25.1-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.25.1-2PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.25.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 726.0 KiB | [timescaledb-tsl_17-2.25.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.25.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.23.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 733.2 KiB | [timescaledb-tsl_17-2.23.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.23.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.23.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 732.7 KiB | [timescaledb-tsl_17-2.23.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.23.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.22.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 733.1 KiB | [timescaledb-tsl_17-2.22.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.22.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.22.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 730.7 KiB | [timescaledb-tsl_17-2.22.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.22.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.21.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 738.9 KiB | [timescaledb-tsl_17-2.21.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.21.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.21.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 738.7 KiB | [timescaledb-tsl_17-2.21.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.21.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.21.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 737.8 KiB | [timescaledb-tsl_17-2.21.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.21.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.21.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 737.0 KiB | [timescaledb-tsl_17-2.21.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.21.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.20.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 722.7 KiB | [timescaledb-tsl_17-2.20.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.20.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.20.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 721.1 KiB | [timescaledb-tsl_17-2.20.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.20.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.20.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 720.2 KiB | [timescaledb-tsl_17-2.20.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.20.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.19.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 700.0 KiB | [timescaledb-tsl_17-2.19.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.19.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.19.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 698.9 KiB | [timescaledb-tsl_17-2.19.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.19.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.19.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 698.3 KiB | [timescaledb-tsl_17-2.19.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.19.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.19.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 773.5 KiB | [timescaledb-tsl_17-2.19.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.19.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.18.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 746.3 KiB | [timescaledb-tsl_17-2.18.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.18.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.18.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 745.6 KiB | [timescaledb-tsl_17-2.18.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.18.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.18.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 743.6 KiB | [timescaledb-tsl_17-2.18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.18.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.17.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 686.4 KiB | [timescaledb-tsl_17-2.17.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.17.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.17.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 685.9 KiB | [timescaledb-tsl_17-2.17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.17.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.17.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 683.4 KiB | [timescaledb-tsl_17-2.17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/timescaledb-tsl_17-2.17.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_17-2.32.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_17-2.32.1-1PGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_17` | `2.31.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_17-2.31.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_17-2.31.0-1PGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_17` | `2.32.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [timescaledb-tsl_17-2.32.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_17-2.32.1-1PGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 1.5 MiB | [timescaledb-tsl_17-2.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.32.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 1.5 MiB | [timescaledb-tsl_17-2.32.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/timescaledb-tsl_17-2.32.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.31.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [timescaledb-tsl_17-2.31.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_17-2.31.0-1PGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_17-2.32.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_17-2.32.1-1PGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_17` | `2.31.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_17-2.31.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_17-2.31.0-1PGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_17` | `2.32.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_17-2.32.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-tsl_17-2.32.1-1PGSTY.el10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_17-2.32.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.32.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_17-2.32.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/timescaledb-tsl_17-2.32.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.31.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_17-2.31.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-tsl_17-2.31.0-1PGSTY.el10.x86_64.rpm) |
| `timescaledb-tsl_17` | `2.32.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_17-2.32.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-tsl_17-2.32.1-1PGSTY.el10.aarch64.rpm) |
| `timescaledb-tsl_17` | `2.31.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_17-2.31.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-tsl_17-2.31.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.4 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.4 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.5 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.3 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-timescaledb-tsl` | `2.32.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/timescaledb-tsl/postgresql-17-timescaledb-tsl_2.32.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_16` | `2.32.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_16-2.32.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_16-2.32.1-1PGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_16-2.32.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.32.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_16-2.32.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.32.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.31.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_16-2.31.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_16-2.31.0-1PGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.31.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_16-2.31.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.31.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.30.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.30.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.30.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.30.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.30.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.30.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.30.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.30.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.30.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.29.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.29.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.29.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.29.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.29.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.29.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.29.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.29.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.29.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.28.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.28.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.28.3-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.28.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.28.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.28.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.28.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.28.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.28.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.28.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_16-2.28.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.28.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.27.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 769.7 KiB | [timescaledb-tsl_16-2.27.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.27.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.27.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 769.4 KiB | [timescaledb-tsl_16-2.27.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.27.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.27.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 768.1 KiB | [timescaledb-tsl_16-2.27.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.27.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.26.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 753.0 KiB | [timescaledb-tsl_16-2.26.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.26.4-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.26.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 750.0 KiB | [timescaledb-tsl_16-2.26.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.26.3-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.26.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 750.8 KiB | [timescaledb-tsl_16-2.26.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.26.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.26.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 749.7 KiB | [timescaledb-tsl_16-2.26.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.26.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.26.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 748.8 KiB | [timescaledb-tsl_16-2.26.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.26.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.25.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 729.2 KiB | [timescaledb-tsl_16-2.25.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.25.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.25.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 726.2 KiB | [timescaledb-tsl_16-2.25.1-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.25.1-2PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.25.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 725.6 KiB | [timescaledb-tsl_16-2.25.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.25.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.23.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 732.3 KiB | [timescaledb-tsl_16-2.23.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.23.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.23.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 731.8 KiB | [timescaledb-tsl_16-2.23.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.23.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.22.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 732.6 KiB | [timescaledb-tsl_16-2.22.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.22.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.22.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 730.0 KiB | [timescaledb-tsl_16-2.22.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.22.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.21.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 737.4 KiB | [timescaledb-tsl_16-2.21.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.21.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.21.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 736.7 KiB | [timescaledb-tsl_16-2.21.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.21.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.21.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 735.8 KiB | [timescaledb-tsl_16-2.21.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.21.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.21.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 735.5 KiB | [timescaledb-tsl_16-2.21.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.21.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.20.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 721.2 KiB | [timescaledb-tsl_16-2.20.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.20.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.20.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 720.0 KiB | [timescaledb-tsl_16-2.20.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.20.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.20.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 719.4 KiB | [timescaledb-tsl_16-2.20.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.20.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.19.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 699.6 KiB | [timescaledb-tsl_16-2.19.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.19.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.19.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 698.4 KiB | [timescaledb-tsl_16-2.19.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.19.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.19.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 697.6 KiB | [timescaledb-tsl_16-2.19.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.19.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.19.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 772.7 KiB | [timescaledb-tsl_16-2.19.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.19.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.18.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 745.8 KiB | [timescaledb-tsl_16-2.18.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.18.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.18.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 744.7 KiB | [timescaledb-tsl_16-2.18.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.18.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.18.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 742.4 KiB | [timescaledb-tsl_16-2.18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.18.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.17.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 686.8 KiB | [timescaledb-tsl_16-2.17.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.17.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.17.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 686.3 KiB | [timescaledb-tsl_16-2.17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.17.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.17.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 683.7 KiB | [timescaledb-tsl_16-2.17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.17.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.16.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 641.5 KiB | [timescaledb-tsl_16-2.16.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.16.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.15.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 638.4 KiB | [timescaledb-tsl_16-2.15.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.15.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.15.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 637.7 KiB | [timescaledb-tsl_16-2.15.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.15.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.15.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 635.3 KiB | [timescaledb-tsl_16-2.15.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.15.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.13.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 758.6 KiB | [timescaledb-tsl_16-2.13.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.13.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.13.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 757.4 KiB | [timescaledb-tsl_16-2.13.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/timescaledb-tsl_16-2.13.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_16-2.32.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_16-2.32.1-1PGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_16` | `2.31.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_16-2.31.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_16-2.31.0-1PGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_16` | `2.32.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [timescaledb-tsl_16-2.32.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_16-2.32.1-1PGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 1.5 MiB | [timescaledb-tsl_16-2.32.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.32.1-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 1.5 MiB | [timescaledb-tsl_16-2.32.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/timescaledb-tsl_16-2.32.0-1PGDG.rhel9.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.31.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [timescaledb-tsl_16-2.31.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_16-2.31.0-1PGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_16-2.32.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_16-2.32.1-1PGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_16` | `2.31.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_16-2.31.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_16-2.31.0-1PGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_16` | `2.32.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_16-2.32.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-tsl_16-2.32.1-1PGSTY.el10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_16-2.32.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.32.1-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 1.4 MiB | [timescaledb-tsl_16-2.32.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/timescaledb-tsl_16-2.32.0-1PGDG.rhel10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.31.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_16-2.31.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-tsl_16-2.31.0-1PGSTY.el10.x86_64.rpm) |
| `timescaledb-tsl_16` | `2.32.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_16-2.32.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-tsl_16-2.32.1-1PGSTY.el10.aarch64.rpm) |
| `timescaledb-tsl_16` | `2.31.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_16-2.31.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-tsl_16-2.31.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.4 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.3 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.4 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.3 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-timescaledb-tsl` | `2.32.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/timescaledb-tsl/postgresql-16-timescaledb-tsl_2.32.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-tsl_15` | `2.28.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.4 MiB | [timescaledb-tsl_15-2.28.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-tsl_15-2.28.3-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.28.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_15-2.28.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.28.3-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.28.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_15-2.28.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.28.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.28.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_15-2.28.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.28.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.28.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [timescaledb-tsl_15-2.28.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.28.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.27.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 764.4 KiB | [timescaledb-tsl_15-2.27.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.27.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.27.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 763.9 KiB | [timescaledb-tsl_15-2.27.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.27.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.27.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 762.5 KiB | [timescaledb-tsl_15-2.27.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.27.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.26.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 751.5 KiB | [timescaledb-tsl_15-2.26.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.26.4-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.26.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 748.5 KiB | [timescaledb-tsl_15-2.26.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.26.3-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.26.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 749.2 KiB | [timescaledb-tsl_15-2.26.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.26.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.26.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 748.5 KiB | [timescaledb-tsl_15-2.26.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.26.1-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.26.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 747.7 KiB | [timescaledb-tsl_15-2.26.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.26.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.25.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 727.5 KiB | [timescaledb-tsl_15-2.25.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.25.2-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.25.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 725.0 KiB | [timescaledb-tsl_15-2.25.1-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.25.1-2PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.25.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 724.5 KiB | [timescaledb-tsl_15-2.25.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.25.0-1PGDG.rhel8.10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.23.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 731.2 KiB | [timescaledb-tsl_15-2.23.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.23.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.23.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 731.0 KiB | [timescaledb-tsl_15-2.23.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.23.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.22.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 730.8 KiB | [timescaledb-tsl_15-2.22.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.22.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.22.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 728.2 KiB | [timescaledb-tsl_15-2.22.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.22.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.21.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 736.2 KiB | [timescaledb-tsl_15-2.21.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.21.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.21.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 735.8 KiB | [timescaledb-tsl_15-2.21.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.21.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.21.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 734.6 KiB | [timescaledb-tsl_15-2.21.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.21.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.21.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 734.1 KiB | [timescaledb-tsl_15-2.21.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.21.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.20.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 719.9 KiB | [timescaledb-tsl_15-2.20.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.20.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.20.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 718.6 KiB | [timescaledb-tsl_15-2.20.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.20.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.20.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 718.0 KiB | [timescaledb-tsl_15-2.20.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.20.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.19.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 699.1 KiB | [timescaledb-tsl_15-2.19.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.19.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.19.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 697.5 KiB | [timescaledb-tsl_15-2.19.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.19.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.19.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 696.9 KiB | [timescaledb-tsl_15-2.19.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.19.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.19.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 772.4 KiB | [timescaledb-tsl_15-2.19.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.19.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.18.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 744.9 KiB | [timescaledb-tsl_15-2.18.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.18.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.18.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 743.9 KiB | [timescaledb-tsl_15-2.18.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.18.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.18.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 742.0 KiB | [timescaledb-tsl_15-2.18.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.18.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.17.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 686.1 KiB | [timescaledb-tsl_15-2.17.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.17.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.17.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 685.5 KiB | [timescaledb-tsl_15-2.17.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.17.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.17.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 682.7 KiB | [timescaledb-tsl_15-2.17.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.17.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.16.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 640.9 KiB | [timescaledb-tsl_15-2.16.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.16.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.15.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 637.4 KiB | [timescaledb-tsl_15-2.15.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.15.3-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.15.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 637.0 KiB | [timescaledb-tsl_15-2.15.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.15.2-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.15.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 633.4 KiB | [timescaledb-tsl_15-2.15.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.15.0-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.13.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 768.2 KiB | [timescaledb-tsl_15-2.13.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/timescaledb-tsl_15-2.13.1-1PGDG.rhel8.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.28.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [timescaledb-tsl_15-2.28.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-tsl_15-2.28.3-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-tsl_15` | `2.28.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 778.6 KiB | [timescaledb-tsl_15-2.28.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-tsl_15-2.28.3-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.28.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 738.7 KiB | [timescaledb-tsl_15-2.28.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-tsl_15-2.28.3-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-tsl_15` | `2.28.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 804.6 KiB | [timescaledb-tsl_15-2.28.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-tsl_15-2.28.3-1PIGSTY.el10.x86_64.rpm) |
| `timescaledb-tsl_15` | `2.28.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 758.3 KiB | [timescaledb-tsl_15-2.28.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-tsl_15-2.28.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.4 MiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.3 MiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 810.2 KiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 778.7 KiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 800.3 KiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 777.9 KiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 804.9 KiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-timescaledb-tsl` | `2.28.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 796.4 KiB | [postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/timescaledb-tsl/postgresql-15-timescaledb-tsl_2.28.3-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/timescaledb" title="Repository" icon="github" subtitle="github.com/timescale/timescaledb" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="timescaledb-2.32.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg timescaledb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install timescaledb;		# install via package name, for the active PG version

pig install timescaledb -v 18;   # install for PG 18
pig install timescaledb -v 17;   # install for PG 17
pig install timescaledb -v 16;   # install for PG 16

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'timescaledb';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION timescaledb;
```

## Usage

Sources:

- [Official source](https://github.com/timescale/timescaledb/releases/tag/2.32.1)
- [Official source](https://github.com/timescale/timescaledb/blob/2.32.1/version.config)
- [Official source](https://github.com/timescale/timescaledb/blob/2.32.1/CHANGELOG.md)

- [TimescaleDB v2.30.2 README](https://github.com/timescale/timescaledb/blob/2.30.2/README.md)
- [TimescaleDB 2.29.0 release](https://github.com/timescale/timescaledb/releases/tag/2.29.0)
- [TimescaleDB 2.30.2 release](https://github.com/timescale/timescaledb/releases/tag/2.30.2)
- [Version 2.30.2 changelog](https://github.com/timescale/timescaledb/blob/2.30.2/CHANGELOG.md)
- [TimescaleDB v2.30.2 control file](https://github.com/timescale/timescaledb/blob/2.30.2/timescaledb.control.in)
- [CREATE TABLE API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_table/)
- [create_hypertable() API](https://www.tigerdata.com/docs/reference/timescaledb/hypertables/create_hypertable/)
- [Continuous aggregate API](https://www.tigerdata.com/docs/reference/timescaledb/continuous-aggregates/create_materialized_view/)
- [add_columnstore_policy() API](https://www.tigerdata.com/docs/reference/timescaledb/hypercore/add_columnstore_policy/)
- [TimescaleDB GUCs](https://www.tigerdata.com/docs/reference/timescaledb/configuration/gucs/)
- [2.32.0 release and compatibility limits](https://github.com/timescale/timescaledb/releases/tag/2.32.0)
- [2.32.0 version declaration](https://github.com/timescale/timescaledb/blob/2.32.0/version.config)
- [2.31 to 2.32 SQL upgrade](https://github.com/timescale/timescaledb/blob/2.32.0/sql/updates/2.31.0--2.32.0.sql)

`timescaledb` is a PostgreSQL extension for time-series and event analytics. The current docs emphasize `CREATE TABLE ... WITH (tsdb.hypertable)`, continuous aggregates, automation jobs, and moving chunks into the columnstore.

### Enable the Extension

Append `timescaledb` to the existing preload list and restart PostgreSQL before creating the extension. Version `2.30.2` supports PostgreSQL 16, 17, and 18.

```ini
shared_preload_libraries = 'timescaledb'
```

### Hypertables

```sql
CREATE EXTENSION timescaledb;

CREATE TABLE ts_test (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
) WITH (
  tsdb.hypertable,
  tsdb.orderby = 'ts DESC'
);
```

To convert an existing PostgreSQL table, use the generalized hypertable API:

```sql
CREATE TABLE ts_existing (
  ts timestamptz NOT NULL,
  id bigint,
  v integer
);
SELECT create_hypertable('ts_existing', by_range('ts'));
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` has been documented since TimescaleDB 2.20.0 and is the best-practice path for new hypertables.
- For TimescaleDB 2.23.0 and later, the first `TIMESTAMP` or `TIMESTAMPTZ` column is selected automatically as the partition column unless more than one candidate makes the choice ambiguous.
- `create_hypertable()` still works for converting existing tables.

### Continuous aggregates and jobs

```sql
CREATE MATERIALIZED VIEW ts_hourly
WITH (timescaledb.continuous) AS
SELECT time_bucket('1 hour', ts) AS bucket,
       count(*) AS cnt,
       avg(v)   AS avg_v
FROM ts_test
GROUP BY bucket;

SELECT add_continuous_aggregate_policy(
  'ts_hourly',
  start_offset => INTERVAL '3 hours',
  end_offset => INTERVAL '1 hour',
  schedule_interval => INTERVAL '1 hour'
);

```

- Continuous aggregates require `time_bucket(...)` on the hypertable's time dimension.
- The continuous aggregate `WITH` clause supports `timescaledb.materialized_only`; the current API default is `TRUE`, so real-time aggregation is not enabled unless configured otherwise.
- TimescaleDB 2.28.0 lets manual `refresh_continuous_aggregate()` calls run incrementally in batches. Use `buckets_per_batch`, `max_batches_per_execution`, and `refresh_newest_first` to break large manual refreshes into smaller work units.
- TimescaleDB 2.28.0 also allows adding a new generated aggregate column to an existing continuous aggregate with `ALTER MATERIALIZED VIEW ... ADD COLUMN ... GENERATED ALWAYS AS (...) STORED`; existing rows are `NULL` until refreshed.

### Columnstore

```sql
CREATE TABLE crypto_ticks (
  "time" timestamptz,
  symbol text,
  price double precision,
  day_volume numeric
) WITH (
  tsdb.hypertable,
  tsdb.segmentby = 'symbol',
  tsdb.orderby = 'time DESC'
);

CALL add_columnstore_policy('crypto_ticks', after => INTERVAL '60 days');
```

- `CREATE TABLE ... WITH (tsdb.hypertable)` enables columnstore by default unless `tsdb.columnstore = false`.
- `add_columnstore_policy()` replaces the older `add_compression_policy()` API and requires either `after` or `created_before`, not both.
- Bloom filters are enabled by default for new columnstore chunks. Existing chunks need recompression before they have bloom indexes.

### Relevant GUCs

```sql
SET timescaledb.enable_direct_compress_insert = on;
SET timescaledb.enable_cagg_rewrites = on;
SET timescaledb.enable_columnar_scan_filter_pushdown = on;
```

`timescaledb.enable_direct_compress_insert` and `timescaledb.enable_direct_compress_copy` enable tech-preview direct compression during ingestion. TimescaleDB 2.27.0 adds `timescaledb.enable_cagg_rewrites` and `timescaledb.cagg_rewrites_debug_info`, and documents `timescaledb.enable_columnar_scan_filter_pushdown` as enabled by default.

### Version 2.30.2 and Upgrades

- TimescaleDB 2.29 supports PostgreSQL 16, 17, and 18. PostgreSQL 15 support ended with the 2.28 line, so upgrade PostgreSQL before moving a PG15 database to 2.29.
- Version 2.29.0 adds `compact_chunk()` and a compaction policy for merging small columnstore batches, plus optimized DML chunk exclusion and small-`LIMIT` columnstore scans. Review the release notes before enabling compaction policies on existing workloads.
- The 2.29 line adds `alter_job(..., config_merge => ...)`, direct-compression and unordered-recompression controls, and concurrent refresh policies for hierarchical continuous aggregates.
- The 2.30 line adds `DeferredChunkAppend` for LIMIT queries and concurrent compaction DML. Version 2.30.2 fixes chunk-merge and compressed-query crashes, orphaned compressed chunks after schema removal, and non-deterministic text-collation handling. Granular-refresh option names now use the `timescaledb.cagg_granular_refresh_*` prefix; review any existing granular-refresh configuration.
- The control file marks `timescaledb` trusted and non-relocatable. The server library still has to be preloaded and PostgreSQL restarted according to the packaged deployment configuration.

Install the matching package, then use a fresh `psql -X` connection for the migration so startup commands do not load an older library first. Apply the update in each database and verify its SQL version:

```sql
ALTER EXTENSION timescaledb UPDATE TO '2.30.2';
SELECT extversion FROM pg_extension WHERE extname = 'timescaledb';
```

### Version 2.32.0 upgrade boundaries

The 2.32.0 release is compatible only up to PostgreSQL 18.6, 17.11 and 16.15. Upstream instructs users to reach TimescaleDB 2.32.1 before upgrading beyond those PostgreSQL minor versions; do not use 2.32.0 to approve an unrestricted PostgreSQL minor upgrade. Install matching extension files and restart PostgreSQL for the loader changes, then update each database's extension through its supported path.

This release checks database ownership before stopping or restarting background workers and runs the scheduler and telemetry jobs as the database owner. Review that role's privileges. It also fixes concurrent `UPDATE`/`DELETE` skipping rows under `READ COMMITTED`, missing final continuous-aggregate buckets, and compression crashes. Adaptive Integer Compression is enabled by default. The upgrade removes the internal `bloom1_contains` helper; application code should not depend on internal APIs.

```sql
ALTER EXTENSION timescaledb UPDATE TO '2.32.0';
SELECT extversion FROM pg_extension WHERE extname = 'timescaledb';
```

### Version 2.32.1

Upgrade TimescaleDB to 2.32.1 before moving PostgreSQL to 16.16, 17.12 or 18.7 and later minor versions. Earlier TimescaleDB patches can fail every scheduler job on those server versions. Version 2.32.1 also remains compatible with the preceding minor releases; it supersedes the 2.32.0 restriction above.

The release repairs job-crash backoff and commits inside `run_job`, and performs chunk/index/continuous-aggregate operations using the appropriate caller or owner instead of unnecessarily switching to the catalog owner. Ownership is checked before validating a unique index on a compressed chunk. Review application-role permissions and scheduled-job results after the upgrade.

Install matching files, follow the preload-library restart procedure, and use a fresh `psql -X` session in each database:

```sql
ALTER EXTENSION timescaledb UPDATE TO '2.32.1';
SELECT extversion FROM pg_extension WHERE extname = 'timescaledb';
```
