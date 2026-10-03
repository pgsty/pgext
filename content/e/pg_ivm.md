---
title: "pg_ivm"
linkTitle: "pg_ivm"
description: "incremental view maintenance on PostgreSQL"
weight: 2840
categories: ["FEAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_ivm**](https://github.com/sraoss/pg_ivm) : incremental view maintenance on PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2840** | {{< badge content="pg_ivm" link="https://github.com/sraoss/pg_ivm" >}} | {{< ext "pg_ivm" >}} | `1.16` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|   **See Also**    | {{< ext "pg_incremental" >}} {{< ext "pg_trickle" >}} {{< ext "timescaledb" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_partman" >}} {{< ext "pg_ttl_index" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_lake" >}} |

> [!Note] Add pg_ivm to shared_preload_libraries or session_preload_libraries for correct maintenance.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.16` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_ivm` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.16` | {{< bg "18" "pg_ivm_18" "green" >}} {{< bg "17" "pg_ivm_17" "green" >}} {{< bg "16" "pg_ivm_16" "green" >}} {{< bg "15" "pg_ivm_15" "green" >}} {{< bg "14" "pg_ivm_14" "green" >}} | `pg_ivm_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.16` | {{< bg "18" "postgresql-18-pg-ivm" "green" >}} {{< bg "17" "postgresql-17-pg-ivm" "green" >}} {{< bg "16" "postgresql-16-pg-ivm" "green" >}} {{< bg "15" "postgresql-15-pg-ivm" "green" >}} {{< bg "14" "postgresql-14-pg-ivm" "green" >}} | `postgresql-$v-pg-ivm` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_16 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_15 : AVAIL 14" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_14 : AVAIL 18" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_17 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_16 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_15 : AVAIL 14" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_14 : AVAIL 14" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_18 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_17 : AVAIL 10" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_16 : AVAIL 11" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_15 : AVAIL 16" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_14 : AVAIL 19" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_18 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_17 : AVAIL 10" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_16 : AVAIL 11" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_15 : AVAIL 16" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_14 : AVAIL 16" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_18 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_16 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_15 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_14 : AVAIL 9" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_18 : AVAIL 8" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_16 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_15 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 1.16" "pg_ivm_14 : AVAIL 9" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.16" "postgresql-18-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-17-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-16-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-15-pg-ivm : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.16" "postgresql-14-pg-ivm : AVAIL 4" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_18` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 137.1 KiB | [pg_ivm_18-1.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_18-1.16-1PGSTY.el8.x86_64.rpm) |
| `pg_ivm_18` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.6 KiB | [pg_ivm_18-1.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.16-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_18` | `1.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.4 KiB | [pg_ivm_18-1.15-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.15-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.3 KiB | [pg_ivm_18-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_ivm_18-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_ivm_18-1.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_ivm_18-1.12-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_18` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 134.3 KiB | [pg_ivm_18-1.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_18-1.16-1PGSTY.el8.aarch64.rpm) |
| `pg_ivm_18` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 51.0 KiB | [pg_ivm_18-1.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.16-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_18` | `1.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [pg_ivm_18-1.15-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.15-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.1 KiB | [pg_ivm_18-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_18` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.5 KiB | [pg_ivm_18-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [pg_ivm_18-1.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_ivm_18-1.12-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_18` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.8 KiB | [pg_ivm_18-1.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_18-1.16-1PGSTY.el9.x86_64.rpm) |
| `pg_ivm_18` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.0 KiB | [pg_ivm_18-1.16-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.16-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_18` | `1.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [pg_ivm_18-1.15-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.15-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [pg_ivm_18-1.14-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.14-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_18-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_18-1.14-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.14-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_18-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.3 KiB | [pg_ivm_18-1.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_ivm_18-1.12-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_18` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.3 KiB | [pg_ivm_18-1.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_18-1.16-1PGSTY.el9.aarch64.rpm) |
| `pg_ivm_18` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.9 KiB | [pg_ivm_18-1.16-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.16-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_18` | `1.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.8 KiB | [pg_ivm_18-1.15-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.15-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [pg_ivm_18-1.14-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.14-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [pg_ivm_18-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [pg_ivm_18-1.14-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.14-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_ivm_18` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.1 KiB | [pg_ivm_18-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_ivm_18-1.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_ivm_18-1.12-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_18` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.1 KiB | [pg_ivm_18-1.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_18-1.16-1PGSTY.el10.x86_64.rpm) |
| `pg_ivm_18` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.3 KiB | [pg_ivm_18-1.16-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.16-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_18` | `1.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.3 KiB | [pg_ivm_18-1.15-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.15-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.8 KiB | [pg_ivm_18-1.14-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.14-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.8 KiB | [pg_ivm_18-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.2 KiB | [pg_ivm_18-1.14-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.14-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_ivm_18` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.6 KiB | [pg_ivm_18-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_18` | `1.12` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.1 KiB | [pg_ivm_18-1.12-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_ivm_18-1.12-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_18` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 139.4 KiB | [pg_ivm_18-1.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_18-1.16-1PGSTY.el10.aarch64.rpm) |
| `pg_ivm_18` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.1 KiB | [pg_ivm_18-1.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.16-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_18` | `1.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.1 KiB | [pg_ivm_18-1.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.15-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.5 KiB | [pg_ivm_18-1.14-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.14-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.5 KiB | [pg_ivm_18-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_18` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.5 KiB | [pg_ivm_18-1.14-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.14-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_ivm_18` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_18-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_18` | `1.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.8 KiB | [pg_ivm_18-1.12-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_ivm_18-1.12-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-ivm` | `1.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 125.3 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.4 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.8 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 118.7 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 121.7 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.7 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 121.0 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 115.4 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 125.5 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.5 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.8 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 118.8 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 121.6 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.7 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.9 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 114.9 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 137.2 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 127.4 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 127.7 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 121.6 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 134.8 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 123.6 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 124.0 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 117.9 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 130.7 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.4 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.6 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 118.7 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 129.5 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.5 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.9 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 114.9 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 129.5 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 122.9 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 123.4 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 117.1 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pg-ivm` | `1.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 127.7 KiB | [postgresql-18-pg-ivm_1.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 119.2 KiB | [postgresql-18-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 119.5 KiB | [postgresql-18-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 113.6 KiB | [postgresql-18-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-18-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_17` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 137.0 KiB | [pg_ivm_17-1.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_17-1.16-1PGSTY.el8.x86_64.rpm) |
| `pg_ivm_17` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.5 KiB | [pg_ivm_17-1.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.16-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_17` | `1.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.3 KiB | [pg_ivm_17-1.15-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.15-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.2 KiB | [pg_ivm_17-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [pg_ivm_17-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.6 KiB | [pg_ivm_17-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [pg_ivm_17-1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_ivm_17-1.9-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_17` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 133.8 KiB | [pg_ivm_17-1.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_17-1.16-1PGSTY.el8.aarch64.rpm) |
| `pg_ivm_17` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 51.0 KiB | [pg_ivm_17-1.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.16-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_17` | `1.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [pg_ivm_17-1.15-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.15-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.0 KiB | [pg_ivm_17-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_17` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.3 KiB | [pg_ivm_17-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.5 KiB | [pg_ivm_17-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [pg_ivm_17-1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_ivm_17-1.9-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_17` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.6 KiB | [pg_ivm_17-1.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_17-1.16-1PGSTY.el9.x86_64.rpm) |
| `pg_ivm_17` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.0 KiB | [pg_ivm_17-1.16-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.16-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_17` | `1.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [pg_ivm_17-1.15-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.15-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_17-1.14-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.14-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_17-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.8 KiB | [pg_ivm_17-1.14-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.14-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_17-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.2 KiB | [pg_ivm_17-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.9 KiB | [pg_ivm_17-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [pg_ivm_17-1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_ivm_17-1.9-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_17` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 137.5 KiB | [pg_ivm_17-1.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_17-1.16-1PGSTY.el9.aarch64.rpm) |
| `pg_ivm_17` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.8 KiB | [pg_ivm_17-1.16-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.16-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_17` | `1.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.8 KiB | [pg_ivm_17-1.15-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.15-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [pg_ivm_17-1.14-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.14-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [pg_ivm_17-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [pg_ivm_17-1.14-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.14-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_ivm_17` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.0 KiB | [pg_ivm_17-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.9 KiB | [pg_ivm_17-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.6 KiB | [pg_ivm_17-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.9 KiB | [pg_ivm_17-1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_ivm_17-1.9-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_17` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.9 KiB | [pg_ivm_17-1.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_17-1.16-1PGSTY.el10.x86_64.rpm) |
| `pg_ivm_17` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.3 KiB | [pg_ivm_17-1.16-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.16-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_17` | `1.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.2 KiB | [pg_ivm_17-1.15-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.15-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.9 KiB | [pg_ivm_17-1.14-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.14-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.9 KiB | [pg_ivm_17-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.3 KiB | [pg_ivm_17-1.14-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.14-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_ivm_17` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.7 KiB | [pg_ivm_17-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.2 KiB | [pg_ivm_17-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.9 KiB | [pg_ivm_17-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_ivm_17-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_17` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.8 KiB | [pg_ivm_17-1.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_17-1.16-1PGSTY.el10.aarch64.rpm) |
| `pg_ivm_17` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.2 KiB | [pg_ivm_17-1.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.16-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_17` | `1.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.2 KiB | [pg_ivm_17-1.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.15-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_17-1.14-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.14-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_17-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_17` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_17-1.14-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.14-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_ivm_17` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.6 KiB | [pg_ivm_17-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_17` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.7 KiB | [pg_ivm_17-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_17` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_ivm_17-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_ivm_17-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-ivm` | `1.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.8 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 123.8 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.2 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 118.3 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 121.3 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.3 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.9 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 115.2 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 124.9 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 123.9 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.3 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 118.2 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 121.1 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.1 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.5 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 114.8 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 157.7 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 147.6 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 148.0 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 141.1 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 155.5 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 143.5 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 144.1 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 137.8 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 130.4 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.2 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.5 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 118.3 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 129.2 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.3 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.7 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 114.7 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 129.0 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 122.8 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 123.1 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 116.8 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pg-ivm` | `1.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 126.9 KiB | [postgresql-17-pg-ivm_1.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 118.6 KiB | [postgresql-17-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 119.1 KiB | [postgresql-17-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 113.6 KiB | [postgresql-17-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-17-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_16` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 137.2 KiB | [pg_ivm_16-1.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_16-1.16-1PGSTY.el8.x86_64.rpm) |
| `pg_ivm_16` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.6 KiB | [pg_ivm_16-1.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.16-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_16` | `1.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.5 KiB | [pg_ivm_16-1.15-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.15-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.4 KiB | [pg_ivm_16-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_ivm_16-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.0 KiB | [pg_ivm_16-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.7 KiB | [pg_ivm_16-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.9 KiB | [pg_ivm_16-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [pg_ivm_16-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_ivm_16-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_16` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 134.2 KiB | [pg_ivm_16-1.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_16-1.16-1PGSTY.el8.aarch64.rpm) |
| `pg_ivm_16` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 51.2 KiB | [pg_ivm_16-1.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.16-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_16` | `1.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [pg_ivm_16-1.15-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.15-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.1 KiB | [pg_ivm_16-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_16` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.4 KiB | [pg_ivm_16-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_ivm_16-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.6 KiB | [pg_ivm_16-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.9 KiB | [pg_ivm_16-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.7 KiB | [pg_ivm_16-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_ivm_16-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_16` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.7 KiB | [pg_ivm_16-1.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_16-1.16-1PGSTY.el9.x86_64.rpm) |
| `pg_ivm_16` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.1 KiB | [pg_ivm_16-1.16-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.16-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_16` | `1.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.2 KiB | [pg_ivm_16-1.15-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.15-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_16-1.14-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.14-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.7 KiB | [pg_ivm_16-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.8 KiB | [pg_ivm_16-1.14-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.14-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [pg_ivm_16-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.2 KiB | [pg_ivm_16-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.0 KiB | [pg_ivm_16-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.7 KiB | [pg_ivm_16-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.6 KiB | [pg_ivm_16-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_ivm_16-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_16` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 137.8 KiB | [pg_ivm_16-1.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_16-1.16-1PGSTY.el9.aarch64.rpm) |
| `pg_ivm_16` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.9 KiB | [pg_ivm_16-1.16-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.16-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_16` | `1.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.9 KiB | [pg_ivm_16-1.15-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.15-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [pg_ivm_16-1.14-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.14-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [pg_ivm_16-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [pg_ivm_16-1.14-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.14-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_ivm_16` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.2 KiB | [pg_ivm_16-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_ivm_16-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.7 KiB | [pg_ivm_16-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.3 KiB | [pg_ivm_16-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.4 KiB | [pg_ivm_16-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_ivm_16-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_16` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.0 KiB | [pg_ivm_16-1.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_16-1.16-1PGSTY.el10.x86_64.rpm) |
| `pg_ivm_16` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.4 KiB | [pg_ivm_16-1.16-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.16-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_16` | `1.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.3 KiB | [pg_ivm_16-1.15-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.15-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.9 KiB | [pg_ivm_16-1.14-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.14-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.9 KiB | [pg_ivm_16-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.3 KiB | [pg_ivm_16-1.14-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.14-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_ivm_16` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.7 KiB | [pg_ivm_16-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [pg_ivm_16-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 43.9 KiB | [pg_ivm_16-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_ivm_16-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_16` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.9 KiB | [pg_ivm_16-1.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_16-1.16-1PGSTY.el10.aarch64.rpm) |
| `pg_ivm_16` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.2 KiB | [pg_ivm_16-1.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.16-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_16` | `1.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.2 KiB | [pg_ivm_16-1.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.15-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_16-1.14-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.14-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_16-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_16` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_16-1.14-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.14-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_ivm_16` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [pg_ivm_16-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_16` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.7 KiB | [pg_ivm_16-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_16` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 42.4 KiB | [pg_ivm_16-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_ivm_16-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-ivm` | `1.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 125.0 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.0 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.3 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 118.1 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 121.7 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.5 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 121.2 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 115.2 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 125.1 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 123.9 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.3 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 118.1 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 121.2 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.1 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.6 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 114.8 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 156.8 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 146.5 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 147.3 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 140.2 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 154.3 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 142.4 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 143.0 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 136.5 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 130.4 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.0 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.5 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 118.1 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 129.4 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.4 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.9 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 114.7 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 129.1 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 122.6 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 123.2 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 117.0 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pg-ivm` | `1.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 127.0 KiB | [postgresql-16-pg-ivm_1.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 118.8 KiB | [postgresql-16-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 119.4 KiB | [postgresql-16-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 113.5 KiB | [postgresql-16-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-16-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_15` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 136.9 KiB | [pg_ivm_15-1.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_15-1.16-1PGSTY.el8.x86_64.rpm) |
| `pg_ivm_15` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.9 KiB | [pg_ivm_15-1.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.16-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_15` | `1.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.6 KiB | [pg_ivm_15-1.15-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.15-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.5 KiB | [pg_ivm_15-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_15` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.8 KiB | [pg_ivm_15-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_ivm_15-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.0 KiB | [pg_ivm_15-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.3 KiB | [pg_ivm_15-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [pg_ivm_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [pg_ivm_15-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.1 KiB | [pg_ivm_15-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.2 KiB | [pg_ivm_15-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.5-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 38.3 KiB | [pg_ivm_15-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.4-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.8 KiB | [pg_ivm_15-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_ivm_15-1.3-1.rhel8.x86_64.rpm) |
| `pg_ivm_15` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 134.2 KiB | [pg_ivm_15-1.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_15-1.16-1PGSTY.el8.aarch64.rpm) |
| `pg_ivm_15` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 51.3 KiB | [pg_ivm_15-1.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.16-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_15` | `1.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.4 KiB | [pg_ivm_15-1.15-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.15-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.3 KiB | [pg_ivm_15-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_15` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.6 KiB | [pg_ivm_15-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.3 KiB | [pg_ivm_15-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.9 KiB | [pg_ivm_15-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.1 KiB | [pg_ivm_15-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [pg_ivm_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [pg_ivm_15-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.5 KiB | [pg_ivm_15-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 37.6 KiB | [pg_ivm_15-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.5-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.4 KiB | [pg_ivm_15-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.4-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.0 KiB | [pg_ivm_15-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_ivm_15-1.3-1.rhel8.aarch64.rpm) |
| `pg_ivm_15` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 140.2 KiB | [pg_ivm_15-1.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_15-1.16-1PGSTY.el9.x86_64.rpm) |
| `pg_ivm_15` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.5 KiB | [pg_ivm_15-1.16-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.16-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_15` | `1.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.6 KiB | [pg_ivm_15-1.15-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.15-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.2 KiB | [pg_ivm_15-1.14-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.14-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.2 KiB | [pg_ivm_15-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.3 KiB | [pg_ivm_15-1.14-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.14-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_ivm_15` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [pg_ivm_15-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.0 KiB | [pg_ivm_15-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.7 KiB | [pg_ivm_15-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.6 KiB | [pg_ivm_15-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.6 KiB | [pg_ivm_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.4 KiB | [pg_ivm_15-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [pg_ivm_15-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.1 KiB | [pg_ivm_15-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.5-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [pg_ivm_15-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.4-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [pg_ivm_15-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_ivm_15-1.3-1.rhel9.x86_64.rpm) |
| `pg_ivm_15` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.3 KiB | [pg_ivm_15-1.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_15-1.16-1PGSTY.el9.aarch64.rpm) |
| `pg_ivm_15` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.6 KiB | [pg_ivm_15-1.16-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.16-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_15` | `1.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 51.5 KiB | [pg_ivm_15-1.15-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.15-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.2 KiB | [pg_ivm_15-1.14-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.14-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.4 KiB | [pg_ivm_15-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.5 KiB | [pg_ivm_15-1.14-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.14-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_ivm_15` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.8 KiB | [pg_ivm_15-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.8 KiB | [pg_ivm_15-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.5 KiB | [pg_ivm_15-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.9 KiB | [pg_ivm_15-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.2 KiB | [pg_ivm_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_ivm_15-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [pg_ivm_15-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [pg_ivm_15-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.5-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.7 KiB | [pg_ivm_15-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.4-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.3 KiB | [pg_ivm_15-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_ivm_15-1.3-1.rhel9.aarch64.rpm) |
| `pg_ivm_15` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.3 KiB | [pg_ivm_15-1.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_15-1.16-1PGSTY.el10.x86_64.rpm) |
| `pg_ivm_15` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 54.8 KiB | [pg_ivm_15-1.16-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.16-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_15` | `1.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.9 KiB | [pg_ivm_15-1.15-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.15-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.5 KiB | [pg_ivm_15-1.14-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.14-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.5 KiB | [pg_ivm_15-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [pg_ivm_15-1.14-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.14-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_ivm_15` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.5 KiB | [pg_ivm_15-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.0 KiB | [pg_ivm_15-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.7 KiB | [pg_ivm_15-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_ivm_15-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_15` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 139.1 KiB | [pg_ivm_15-1.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_15-1.16-1PGSTY.el10.aarch64.rpm) |
| `pg_ivm_15` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 53.3 KiB | [pg_ivm_15-1.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.16-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_15` | `1.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.4 KiB | [pg_ivm_15-1.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.15-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.1 KiB | [pg_ivm_15-1.14-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.14-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.3 KiB | [pg_ivm_15-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_15` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.3 KiB | [pg_ivm_15-1.14-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.14-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_ivm_15` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.4 KiB | [pg_ivm_15-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_15` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.5 KiB | [pg_ivm_15-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_15` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 43.1 KiB | [pg_ivm_15-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_ivm_15-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-ivm` | `1.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.9 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.0 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.7 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 118.7 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 121.2 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.2 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.6 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 115.2 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 125.0 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.0 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.6 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 118.4 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 120.9 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 119.9 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.6 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 114.9 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 156.2 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 146.5 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 146.8 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 140.3 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 153.9 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 142.5 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 143.1 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 136.3 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 130.4 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.0 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.4 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 118.3 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 129.0 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.5 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 120.9 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 115.2 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 128.8 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 122.6 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 123.3 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 117.1 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pg-ivm` | `1.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 127.1 KiB | [postgresql-15-pg-ivm_1.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 118.9 KiB | [postgresql-15-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 119.5 KiB | [postgresql-15-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 113.6 KiB | [postgresql-15-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-15-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ivm_14` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 228.2 KiB | [pg_ivm_14-1.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ivm_14-1.16-1PGSTY.el8.x86_64.rpm) |
| `pg_ivm_14` | `1.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 82.1 KiB | [pg_ivm_14-1.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.16-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_14` | `1.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 80.9 KiB | [pg_ivm_14-1.15-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.15-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.8 KiB | [pg_ivm_14-1.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.14-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_ivm_14` | `1.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.0 KiB | [pg_ivm_14-1.13-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.13-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.8 KiB | [pg_ivm_14-1.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.11-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.5 KiB | [pg_ivm_14-1.10-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.10-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.6 KiB | [pg_ivm_14-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.6 KiB | [pg_ivm_14-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 71.4 KiB | [pg_ivm_14-1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.6-1PGDG.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.0 KiB | [pg_ivm_14-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.1 KiB | [pg_ivm_14-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.5-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.2 KiB | [pg_ivm_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.4-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 67.6 KiB | [pg_ivm_14-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.3-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 66.2 KiB | [pg_ivm_14-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.2-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.4 KiB | [pg_ivm_14-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.1-1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.8 KiB | [pg_ivm_14-1.0-.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.0-.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.9 KiB | [pg_ivm_14-1.0-alpha1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_ivm_14-1.0-alpha1.rhel8.x86_64.rpm) |
| `pg_ivm_14` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 222.8 KiB | [pg_ivm_14-1.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ivm_14-1.16-1PGSTY.el8.aarch64.rpm) |
| `pg_ivm_14` | `1.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 77.1 KiB | [pg_ivm_14-1.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.16-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_14` | `1.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 76.1 KiB | [pg_ivm_14-1.15-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.15-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.1 KiB | [pg_ivm_14-1.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.14-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_ivm_14` | `1.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.4 KiB | [pg_ivm_14-1.13-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.13-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.1 KiB | [pg_ivm_14-1.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.11-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.8 KiB | [pg_ivm_14-1.10-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.10-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.0 KiB | [pg_ivm_14-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 67.0 KiB | [pg_ivm_14-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 66.8 KiB | [pg_ivm_14-1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.6-1PGDG.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.7 KiB | [pg_ivm_14-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.9 KiB | [pg_ivm_14-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.5-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.8 KiB | [pg_ivm_14-1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.4-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.4 KiB | [pg_ivm_14-1.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_ivm_14-1.3-1.rhel8.aarch64.rpm) |
| `pg_ivm_14` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 233.5 KiB | [pg_ivm_14-1.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ivm_14-1.16-1PGSTY.el9.x86_64.rpm) |
| `pg_ivm_14` | `1.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 82.8 KiB | [pg_ivm_14-1.16-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.16-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_14` | `1.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 81.9 KiB | [pg_ivm_14-1.15-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.15-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 79.8 KiB | [pg_ivm_14-1.14-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.14-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 79.8 KiB | [pg_ivm_14-1.14-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.14-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.x86_64](/os/el9.x86_64) | pgdg | 80.0 KiB | [pg_ivm_14-1.14-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.14-1PGDG.rhel9.6.x86_64.rpm) |
| `pg_ivm_14` | `1.13` | [el9.x86_64](/os/el9.x86_64) | pgdg | 79.3 KiB | [pg_ivm_14-1.13-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.13-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.5 KiB | [pg_ivm_14-1.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.11-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.2 KiB | [pg_ivm_14-1.10-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.10-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.0 KiB | [pg_ivm_14-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.4 KiB | [pg_ivm_14-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.6 KiB | [pg_ivm_14-1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.6-1PGDG.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.3 KiB | [pg_ivm_14-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.4 KiB | [pg_ivm_14-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.5-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.5 KiB | [pg_ivm_14-1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.4-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 71.0 KiB | [pg_ivm_14-1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.3-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.4 KiB | [pg_ivm_14-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.2-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [pg_ivm_14-1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.1-1.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.2 KiB | [pg_ivm_14-1.0-.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_ivm_14-1.0-.rhel9.x86_64.rpm) |
| `pg_ivm_14` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 229.6 KiB | [pg_ivm_14-1.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ivm_14-1.16-1PGSTY.el9.aarch64.rpm) |
| `pg_ivm_14` | `1.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 80.3 KiB | [pg_ivm_14-1.16-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.16-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_14` | `1.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 79.2 KiB | [pg_ivm_14-1.15-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.15-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 77.3 KiB | [pg_ivm_14-1.14-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.14-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 77.4 KiB | [pg_ivm_14-1.14-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.14-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el9.aarch64](/os/el9.aarch64) | pgdg | 77.5 KiB | [pg_ivm_14-1.14-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.14-1PGDG.rhel9.6.aarch64.rpm) |
| `pg_ivm_14` | `1.13` | [el9.aarch64](/os/el9.aarch64) | pgdg | 77.0 KiB | [pg_ivm_14-1.13-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.13-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.11` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.9 KiB | [pg_ivm_14-1.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.11-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.10` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.6 KiB | [pg_ivm_14-1.10-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.10-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.2 KiB | [pg_ivm_14-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.6 KiB | [pg_ivm_14-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.4 KiB | [pg_ivm_14-1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.6-1PGDG.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.2 KiB | [pg_ivm_14-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.3 KiB | [pg_ivm_14-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.5-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.3 KiB | [pg_ivm_14-1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.4-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 68.0 KiB | [pg_ivm_14-1.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_ivm_14-1.3-1.rhel9.aarch64.rpm) |
| `pg_ivm_14` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 234.3 KiB | [pg_ivm_14-1.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ivm_14-1.16-1PGSTY.el10.x86_64.rpm) |
| `pg_ivm_14` | `1.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 84.2 KiB | [pg_ivm_14-1.16-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.16-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_14` | `1.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 83.4 KiB | [pg_ivm_14-1.15-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.15-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 81.1 KiB | [pg_ivm_14-1.14-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.14-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 81.1 KiB | [pg_ivm_14-1.14-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.14-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.x86_64](/os/el10.x86_64) | pgdg | 81.4 KiB | [pg_ivm_14-1.14-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.14-1PGDG.rhel10.0.x86_64.rpm) |
| `pg_ivm_14` | `1.13` | [el10.x86_64](/os/el10.x86_64) | pgdg | 80.8 KiB | [pg_ivm_14-1.13-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.13-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.11` | [el10.x86_64](/os/el10.x86_64) | pgdg | 74.9 KiB | [pg_ivm_14-1.11-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.11-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.10` | [el10.x86_64](/os/el10.x86_64) | pgdg | 74.6 KiB | [pg_ivm_14-1.10-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_ivm_14-1.10-1PGDG.rhel10.x86_64.rpm) |
| `pg_ivm_14` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 231.3 KiB | [pg_ivm_14-1.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ivm_14-1.16-1PGSTY.el10.aarch64.rpm) |
| `pg_ivm_14` | `1.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 81.8 KiB | [pg_ivm_14-1.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.16-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_14` | `1.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 81.0 KiB | [pg_ivm_14-1.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.15-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 78.8 KiB | [pg_ivm_14-1.14-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.14-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 78.9 KiB | [pg_ivm_14-1.14-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.14-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_ivm_14` | `1.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 78.9 KiB | [pg_ivm_14-1.14-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.14-1PGDG.rhel10.0.aarch64.rpm) |
| `pg_ivm_14` | `1.13` | [el10.aarch64](/os/el10.aarch64) | pgdg | 79.1 KiB | [pg_ivm_14-1.13-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.13-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_14` | `1.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 72.4 KiB | [pg_ivm_14-1.11-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.11-1PGDG.rhel10.aarch64.rpm) |
| `pg_ivm_14` | `1.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 72.1 KiB | [pg_ivm_14-1.10-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_ivm_14-1.10-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-ivm` | `1.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 215.2 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 214.3 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 214.5 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d12.x86_64](/os/d12.x86_64) | pgdg | 209.0 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 207.5 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 206.8 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 207.6 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d12.aarch64](/os/d12.aarch64) | pgdg | 201.9 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 215.0 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 213.9 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 214.2 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d13.x86_64](/os/d13.x86_64) | pgdg | 208.6 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 207.8 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 207.1 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 207.4 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [d13.aarch64](/os/d13.aarch64) | pgdg | 201.9 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 259.9 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 244.7 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 245.3 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u22.x86_64](/os/u22.x86_64) | pgdg | 238.7 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 255.8 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 236.6 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 237.4 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u22.aarch64](/os/u22.aarch64) | pgdg | 230.9 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 224.7 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 214.2 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 214.7 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u24.x86_64](/os/u24.x86_64) | pgdg | 208.9 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 221.6 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 207.1 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 207.6 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u24.aarch64](/os/u24.aarch64) | pgdg | 202.0 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 221.6 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 211.4 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 211.5 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u26.x86_64](/os/u26.x86_64) | pgdg | 206.0 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pg-ivm` | `1.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 217.9 KiB | [postgresql-14-pg-ivm_1.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 203.5 KiB | [postgresql-14-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 204.0 KiB | [postgresql-14-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.15-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pg-ivm` | `1.13` | [u26.aarch64](/os/u26.aarch64) | pgdg | 198.6 KiB | [postgresql-14-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-ivm/postgresql-14-pg-ivm_1.13-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/sraoss/pg_ivm" title="Repository" icon="github" subtitle="github.com/sraoss/pg_ivm" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_ivm-1.16.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_ivm;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_ivm;		# install via package name, for the active PG version

pig install pg_ivm -v 18;   # install for PG 18
pig install pg_ivm -v 17;   # install for PG 17
pig install pg_ivm -v 16;   # install for PG 16
pig install pg_ivm -v 15;   # install for PG 15
pig install pg_ivm -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_ivm';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_ivm;
```

## Usage

Sources:

- [Official v1.16 README](https://github.com/sraoss/pg_ivm/blob/v1.16/README.md)
- [v1.16 release notes](https://github.com/sraoss/pg_ivm/releases/tag/v1.16)
- [1.15 to 1.16 upgrade SQL](https://github.com/sraoss/pg_ivm/blob/v1.16/pg_ivm--1.15--1.16.sql)
- [pg_ivm_dump_metadata utility](https://github.com/sraoss/pg_ivm/blob/v1.16/scripts/pg_ivm_dump_metadata)

`pg_ivm` provides immediate incremental view maintenance for PostgreSQL. An Incrementally Maintainable Materialized View (IMMV) is stored as a table with triggers and metadata in the `pgivm` schema; base-table changes update the IMMV inside the same transaction instead of recomputing the complete query.

### Enable and Create an IMMV

Load the library for every session that can modify an IMMV's base tables. A cluster-wide setup requires a restart:

```conf
shared_preload_libraries = 'pg_ivm'
```

`session_preload_libraries = 'pg_ivm'` is also supported when managed consistently for all relevant sessions.

```sql
CREATE EXTENSION pg_ivm;

CREATE TABLE accounts (
    account_id bigint PRIMARY KEY,
    branch_id integer NOT NULL,
    balance numeric NOT NULL
);
INSERT INTO accounts VALUES (42, 1, 1000);

SELECT pgivm.create_immv(
    'account_totals',
    'SELECT branch_id, count(*) AS accounts, sum(balance) AS balance
     FROM accounts
     GROUP BY branch_id'
);

UPDATE accounts
SET balance = balance + 100
WHERE account_id = 42;

SELECT * FROM account_totals;
```

### Manage and Inspect IMMVs

- `pgivm.create_immv(name, query)`: creates and populates an IMMV, returning its row count.
- `pgivm.refresh_immv(name, with_data)`: fully rebuilds the IMMV; `false` disables maintenance until a later populated refresh.
- `pgivm.get_immv_def(regclass)`: returns the stored view definition.
- `pgivm.restore_immv(name, query, populate)`: version 1.15 function that reconstructs metadata, triggers, and indexes for an existing IMMV table.
- `pgivm.get_create_immv_commands()` and `pgivm.get_restore_immv_commands()`: emit SQL for rebuilding IMMVs or restoring their metadata.

Version 1.15 includes a helper for dump or `pg_upgrade` workflows:

```shell
pg_ivm_dump_metadata -d application > pg_ivm_metadata.sql
```

The script emits `pgivm.restore_immv()` calls. Restore the table data first, then execute the saved metadata SQL so incremental maintenance resumes without recreating the tables.

### Restrictions and Operational Caveats

- Supported definitions include selected joins, `DISTINCT`, simple subqueries/CTEs, and built-in `count`, `sum`, `avg`, `min`, and `max` aggregates. Unsupported constructs include `HAVING`, window functions, `ORDER BY`, `LIMIT/OFFSET`, set operations, `DISTINCT ON`, and user-defined aggregates.
- Efficient maintenance depends on a suitable unique index. `create_immv()` creates one automatically only when the definition supplies usable grouping, distinct, or base-table primary-key columns.
- Creation and refresh take `AccessExclusiveLock`. Upstream warns about consistency risks for creation under `REPEATABLE READ` or `SERIALIZABLE`; use `READ COMMITTED` or refresh afterward.
- `restore_immv()` fails when the relation is already registered or its table definition does not match the supplied query.
- Version 1.15 also fixes incorrect maintenance after repeated trigger-driven modifications and a v1.14 outer-join maintenance crash.

### Upgrade to 1.16

Version 1.16 adds PostgreSQL 19 support and fixes a crash when maintaining and dropping an IMMV in one transaction, maintenance involving columns without default equality operators, large OID handling, and maintenance lock-release timing. Install matching files and ensure modifying sessions load the new library, then run:

```sql
ALTER EXTENSION pg_ivm UPDATE TO '1.16';
```
