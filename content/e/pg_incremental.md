---
title: "pg_incremental"
linkTitle: "pg_incremental"
description: "Incremental Processing by Crunchy Data"
weight: 2850
categories: ["FEAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_incremental**](https://github.com/CrunchyData/pg_incremental) : Incremental Processing by Crunchy Data


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2850** | {{< badge content="pg_incremental" link="https://github.com/CrunchyData/pg_incremental" >}} | {{< ext "pg_incremental" >}} | `1.6` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "pg_ivm" >}} {{< ext "pg_trickle" >}} {{< ext "timescaledb" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_partman" >}} {{< ext "pg_ttl_index" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_lake" >}} |

> [!Note] Package 1.6.1; SQL/control 1.6. Fixes _drop_trigger search_path injection. Sequence pipelines require SELECT or USAGE at creation and execution; file-list functions run as the pipeline executor. pg_cron is optional except for scheduled pipelines.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_incremental` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.6.1` | {{< bg "18" "pg_incremental_18" "green" >}} {{< bg "17" "pg_incremental_17" "green" >}} {{< bg "16" "pg_incremental_16" "green" >}} {{< bg "15" "pg_incremental_15" "red" >}} {{< bg "14" "pg_incremental_14" "red" >}} | `pg_incremental_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.6.1` | {{< bg "18" "postgresql-18-pg-incremental" "green" >}} {{< bg "17" "postgresql-17-pg-incremental" "green" >}} {{< bg "16" "postgresql-16-pg-incremental" "green" >}} {{< bg "15" "postgresql-15-pg-incremental" "red" >}} {{< bg "14" "postgresql-14-pg-incremental" "red" >}} | `postgresql-$v-pg-incremental` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_16 : AVAIL 4" "blue" >}} | {{< bg "N/A" "pg_incremental_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_incremental_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_16 : AVAIL 4" "blue" >}} | {{< bg "N/A" "pg_incremental_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_incremental_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_16 : AVAIL 5" "blue" >}} | {{< bg "N/A" "pg_incremental_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_incremental_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_16 : AVAIL 5" "blue" >}} | {{< bg "N/A" "pg_incremental_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_incremental_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_16 : AVAIL 5" "blue" >}} | {{< bg "N/A" "pg_incremental_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_incremental_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.1" "pg_incremental_16 : AVAIL 5" "blue" >}} | {{< bg "N/A" "pg_incremental_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_incremental_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-18-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-17-pg-incremental : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6.1" "postgresql-16-pg-incremental : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-incremental : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-incremental : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_incremental_18` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.5 KiB | [pg_incremental_18-1.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_incremental_18-1.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.5 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_incremental_18-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.4 KiB | [pg_incremental_18-1.5.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_incremental_18-1.5.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_incremental_18-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_incremental_18` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.3 KiB | [pg_incremental_18-1.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_incremental_18-1.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_incremental_18-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.6 KiB | [pg_incremental_18-1.5.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_incremental_18-1.5.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.0 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_incremental_18-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_incremental_18` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.8 KiB | [pg_incremental_18-1.6.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_incremental_18-1.6.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_incremental_18-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [pg_incremental_18-1.5.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_incremental_18-1.5.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [pg_incremental_18-1.0.0-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_incremental_18-1.0.0-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_incremental_18-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_incremental_18` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.1 KiB | [pg_incremental_18-1.6.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_incremental_18-1.6.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.6 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_incremental_18-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.0 KiB | [pg_incremental_18-1.5.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_incremental_18-1.5.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.6 KiB | [pg_incremental_18-1.0.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_incremental_18-1.0.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.3 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_incremental_18-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_incremental_18` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.3 KiB | [pg_incremental_18-1.6.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_incremental_18-1.6.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.4 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_incremental_18-1.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [pg_incremental_18-1.5.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_incremental_18-1.5.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [pg_incremental_18-1.0.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_incremental_18-1.0.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_incremental_18-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_incremental_18` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.4 KiB | [pg_incremental_18-1.6.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_incremental_18-1.6.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.9 KiB | [pg_incremental_18-1.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_incremental_18-1.5.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_incremental_18` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [pg_incremental_18-1.5.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_incremental_18-1.5.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.0 KiB | [pg_incremental_18-1.0.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_incremental_18-1.0.0-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.1 KiB | [pg_incremental_18-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_incremental_18-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-incremental` | `1.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 70.3 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 69.0 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 70.2 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 69.0 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 71.7 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 70.9 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 70.0 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 69.9 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 69.7 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-incremental` | `1.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 69.4 KiB | [postgresql-18-pg-incremental_1.6.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-18-pg-incremental_1.6.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_incremental_17` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.5 KiB | [pg_incremental_17-1.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_incremental_17-1.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.5 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_incremental_17-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.4 KiB | [pg_incremental_17-1.5.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_incremental_17-1.5.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_incremental_17-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_incremental_17` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.3 KiB | [pg_incremental_17-1.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_incremental_17-1.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.0 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_incremental_17-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.6 KiB | [pg_incremental_17-1.5.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_incremental_17-1.5.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.0 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_incremental_17-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_incremental_17` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.8 KiB | [pg_incremental_17-1.6.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_incremental_17-1.6.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_incremental_17-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [pg_incremental_17-1.5.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_incremental_17-1.5.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [pg_incremental_17-1.0.0-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_incremental_17-1.0.0-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_incremental_17-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_incremental_17` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.1 KiB | [pg_incremental_17-1.6.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_incremental_17-1.6.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.6 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_incremental_17-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.0 KiB | [pg_incremental_17-1.5.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_incremental_17-1.5.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.7 KiB | [pg_incremental_17-1.0.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_incremental_17-1.0.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.5 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_incremental_17-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_incremental_17` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.3 KiB | [pg_incremental_17-1.6.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_incremental_17-1.6.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.4 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_incremental_17-1.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [pg_incremental_17-1.5.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_incremental_17-1.5.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.7 KiB | [pg_incremental_17-1.0.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_incremental_17-1.0.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_incremental_17-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_incremental_17` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.4 KiB | [pg_incremental_17-1.6.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_incremental_17-1.6.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.8 KiB | [pg_incremental_17-1.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_incremental_17-1.5.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_incremental_17` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [pg_incremental_17-1.5.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_incremental_17-1.5.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.0 KiB | [pg_incremental_17-1.0.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_incremental_17-1.0.0-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.1 KiB | [pg_incremental_17-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_incremental_17-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-incremental` | `1.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 70.3 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 69.0 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 70.3 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 69.1 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 78.1 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 77.1 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 70.0 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 69.9 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 69.7 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-incremental` | `1.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 69.4 KiB | [postgresql-17-pg-incremental_1.6.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-17-pg-incremental_1.6.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_incremental_16` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.5 KiB | [pg_incremental_16-1.6.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_incremental_16-1.6.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.5 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_incremental_16-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [pg_incremental_16-1.5.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_incremental_16-1.5.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.7 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_incremental_16-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_incremental_16` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.3 KiB | [pg_incremental_16-1.6.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_incremental_16-1.6.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 33.1 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_incremental_16-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.6 KiB | [pg_incremental_16-1.5.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_incremental_16-1.5.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.0 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_incremental_16-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_incremental_16` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 37.8 KiB | [pg_incremental_16-1.6.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_incremental_16-1.6.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_incremental_16-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.4 KiB | [pg_incremental_16-1.5.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_incremental_16-1.5.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [pg_incremental_16-1.0.0-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_incremental_16-1.0.0-3PGDG.rhel9.8.x86_64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_incremental_16-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_incremental_16` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 37.1 KiB | [pg_incremental_16-1.6.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_incremental_16-1.6.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.6 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_incremental_16-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.0 KiB | [pg_incremental_16-1.5.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_incremental_16-1.5.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.6 KiB | [pg_incremental_16-1.0.0-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_incremental_16-1.0.0-3PGDG.rhel9.8.aarch64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.5 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_incremental_16-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_incremental_16` | `1.6.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 38.3 KiB | [pg_incremental_16-1.6.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_incremental_16-1.6.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.4 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_incremental_16-1.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.1 KiB | [pg_incremental_16-1.5.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_incremental_16-1.5.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [pg_incremental_16-1.0.0-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_incremental_16-1.0.0-3PGDG.rhel10.2.x86_64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_incremental_16-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_incremental_16` | `1.6.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 37.4 KiB | [pg_incremental_16-1.6.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_incremental_16-1.6.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.8 KiB | [pg_incremental_16-1.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_incremental_16-1.5.0-1PIGSTY.el10.aarch64.rpm) |
| `pg_incremental_16` | `1.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.3 KiB | [pg_incremental_16-1.5.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_incremental_16-1.5.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.0 KiB | [pg_incremental_16-1.0.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_incremental_16-1.0.0-3PGDG.rhel10.2.aarch64.rpm) |
| `pg_incremental_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.1 KiB | [pg_incremental_16-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_incremental_16-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-incremental` | `1.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 70.4 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 69.1 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 70.3 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 69.1 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 78.1 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 77.2 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 70.1 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 70.2 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 69.7 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-incremental` | `1.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 69.4 KiB | [postgresql-16-pg-incremental_1.6.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-incremental/postgresql-16-pg-incremental_1.6.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pg_incremental" title="Repository" icon="github" subtitle="github.com/CrunchyData/pg_incremental" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_incremental-1.6.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_incremental;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_incremental;		# install via package name, for the active PG version

pig install pg_incremental -v 18;   # install for PG 18
pig install pg_incremental -v 17;   # install for PG 17
pig install pg_incremental -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_incremental;
```

## Usage

Sources:

- [Official README](https://github.com/CrunchyData/pg_incremental/blob/v1.6.1/README.md)
- [v1.6.1 release](https://github.com/CrunchyData/pg_incremental/releases/tag/v1.6.1)
- [Control file](https://github.com/CrunchyData/pg_incremental/blob/v1.6.1/pg_incremental.control)

`pg_incremental` defines exactly-once incremental pipelines for append-only tables and file feeds. It depends on `pg_cron` and supports sequence, time-interval and file-list pipelines on PostgreSQL 17–18 in the reviewed upstream test matrix.

### Install and schedule

```sql
CREATE EXTENSION pg_incremental CASCADE;
GRANT USAGE ON SCHEMA cron TO app_user;
```

Pipelines execute immediately by default and then run from `pg_cron`; set `execute_immediately := false` when the initial run should be deferred. The progress update and the pipeline command commit together, so a failed command does not advance the processed range.

### Pipeline types

```sql
SELECT incremental.create_sequence_pipeline('event-aggregation', 'events', $$
  INSERT INTO events_agg
  SELECT date_trunc('day', event_time), count(*)
  FROM events
  WHERE event_id BETWEEN $1 AND $2
  GROUP BY 1
  ON CONFLICT (day) DO UPDATE
  SET event_count = events_agg.event_count + excluded.event_count
$$);

SELECT incremental.create_time_interval_pipeline('hourly-export', '1 hour', $$
  SELECT export_events($1, $2)
$$);

SELECT incremental.create_file_list_pipeline('event-import', 's3://bucket/events/*.csv', $$
  SELECT import_events($1)
$$);
```

Sequence pipelines pass a safe `$1`/`$2` range; time pipelines pass interval boundaries; file-list pipelines pass a newly discovered file. `max_batch_size`, `max_batches_per_run` and `batched := false` tune workload shape. `incremental.execute_pipeline`, `reset_pipeline`, `drop_pipeline` and the pipeline status views provide operational control.

### Security fixes in 1.6.1

Release 1.6.1 fixes a `_drop_trigger` `search_path` injection issue, requires `SELECT` or `USAGE` on a sequence when a sequence pipeline is created and run, and executes file-list functions as the caller. Review ownership and grants before upgrading; existing package facts are separate from this upstream release.
