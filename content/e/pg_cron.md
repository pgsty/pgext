---
title: "pg_cron"
linkTitle: "pg_cron"
description: "Job scheduler for PostgreSQL"
weight: 1070
categories: ["TIME"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_cron**](https://github.com/citusdata/pg_cron) : Job scheduler for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1070** | {{< badge content="pg_cron" link="https://github.com/citusdata/pg_cron" >}} | {{< ext "pg_cron" >}} | `1.6.8` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|    **Need By**    | {{< ext "documentdb" >}} {{< ext "pg_dispatch" >}} {{< ext "pgfr_record" >}} {{< ext "pglock" >}} {{< ext "pgmb" >}} {{< ext "timeseries" >}} {{< ext "vectorize" >}} |
|   **See Also**    | {{< ext "pgagent" >}} {{< ext "pg_task" >}} {{< ext "pg_dbms_job" >}} {{< ext "pg_dispatch" >}} {{< ext "pg_later" >}} {{< ext "pg_jobmon" >}} {{< ext "pg_durable" >}} {{< ext "pg_background" >}} {{< ext "pg_when" >}} {{< ext "cron_utils" >}} |

> [!Note] require cron.database_name


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.8` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_cron` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.8` | {{< bg "19" "pg_cron_19" "green" >}} {{< bg "18" "pg_cron_18" "green" >}} {{< bg "17" "pg_cron_17" "green" >}} {{< bg "16" "pg_cron_16" "green" >}} {{< bg "15" "pg_cron_15" "green" >}} | `pg_cron_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.8` | {{< bg "19" "postgresql-19-cron" "green" >}} {{< bg "18" "postgresql-18-cron" "green" >}} {{< bg "17" "postgresql-17-cron" "green" >}} {{< bg "16" "postgresql-16-cron" "green" >}} {{< bg "15" "postgresql-15-cron" "green" >}} | `postgresql-$v-cron` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_cron_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.6.8" "pg_cron_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_15 : AVAIL 12" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_cron_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.6.8" "pg_cron_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_15 : AVAIL 12" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.6.8" "pg_cron_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_15 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.6.8" "pg_cron_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.6.8" "pg_cron_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.6.8" "pg_cron_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.6.8" "pg_cron_15 : AVAIL 6" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.6.8" "postgresql-19-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-18-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-17-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-16-cron : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.8" "postgresql-15-cron : AVAIL 3" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_19` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.7 KiB | [pg_cron_19-1.6.8-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pg_cron_19-1.6.8-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_cron_19` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.6 KiB | [pg_cron_19-1.6.8-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pg_cron_19-1.6.8-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_cron_19` | `1.6.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.8 KiB | [pg_cron_19-1.6.8-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pg_cron_19-1.6.8-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_cron_19` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.2 KiB | [pg_cron_19-1.6.8-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pg_cron_19-1.6.8-4PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.6 KiB | [postgresql-19-cron_1.6.8-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg12+3_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.4 KiB | [postgresql-19-cron_1.6.8-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg12+2_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.1 KiB | [postgresql-19-cron_1.6.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg12+1_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.3 KiB | [postgresql-19-cron_1.6.8-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg12+3_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.0 KiB | [postgresql-19-cron_1.6.8-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg12+2_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.1 KiB | [postgresql-19-cron_1.6.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg12+1_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.7 KiB | [postgresql-19-cron_1.6.8-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg13+3_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.6 KiB | [postgresql-19-cron_1.6.8-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg13+2_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.3 KiB | [postgresql-19-cron_1.6.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg13+1_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.6 KiB | [postgresql-19-cron_1.6.8-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg13+3_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.1 KiB | [postgresql-19-cron_1.6.8-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg13+2_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.1 KiB | [postgresql-19-cron_1.6.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg13+1_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.1 KiB | [postgresql-19-cron_1.6.8-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg22.04+3_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.9 KiB | [postgresql-19-cron_1.6.8-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg22.04+2_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.9 KiB | [postgresql-19-cron_1.6.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 86.6 KiB | [postgresql-19-cron_1.6.8-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg22.04+3_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.0 KiB | [postgresql-19-cron_1.6.8-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg22.04+2_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.1 KiB | [postgresql-19-cron_1.6.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 88.6 KiB | [postgresql-19-cron_1.6.8-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg24.04+3_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 88.5 KiB | [postgresql-19-cron_1.6.8-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg24.04+2_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 88.4 KiB | [postgresql-19-cron_1.6.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.0 KiB | [postgresql-19-cron_1.6.8-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg24.04+3_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 85.7 KiB | [postgresql-19-cron_1.6.8-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg24.04+2_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 85.8 KiB | [postgresql-19-cron_1.6.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.6 KiB | [postgresql-19-cron_1.6.8-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg26.04+3_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.1 KiB | [postgresql-19-cron_1.6.8-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg26.04+2_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.2 KiB | [postgresql-19-cron_1.6.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 84.7 KiB | [postgresql-19-cron_1.6.8-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg26.04+3_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 84.5 KiB | [postgresql-19-cron_1.6.8-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg26.04+2_arm64.deb) |
| `postgresql-19-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 84.4 KiB | [postgresql-19-cron_1.6.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-19-cron_1.6.8-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_18` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_cron_18-1.6.8-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_cron_18-1.6.8-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_cron_18` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.0 KiB | [pg_cron_18-1.6.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_cron_18-1.6.8-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_cron_18` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.8 KiB | [pg_cron_18-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_cron_18-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_18` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.9 KiB | [pg_cron_18-1.6.8-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_cron_18-1.6.8-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_cron_18` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.4 KiB | [pg_cron_18-1.6.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_cron_18-1.6.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_cron_18` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.2 KiB | [pg_cron_18-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_cron_18-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_18` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.8 KiB | [pg_cron_18-1.6.8-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_cron_18-1.6.8-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_cron_18` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.8 KiB | [pg_cron_18-1.6.8-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_cron_18-1.6.8-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_cron_18` | `1.6.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 49.0 KiB | [pg_cron_18-1.6.8-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_cron_18-1.6.8-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_cron_18` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.4 KiB | [pg_cron_18-1.6.8-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_cron_18-1.6.8-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_cron_18` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.4 KiB | [pg_cron_18-1.6.8-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_cron_18-1.6.8-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_18` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.5 KiB | [pg_cron_18-1.6.8-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_cron_18-1.6.8-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_18` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.2 KiB | [pg_cron_18-1.6.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_cron_18-1.6.7-5PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_18` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [pg_cron_18-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_cron_18-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.2 KiB | [postgresql-18-cron_1.6.8-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg12+3_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.0 KiB | [postgresql-18-cron_1.6.8-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg12+2_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.0 KiB | [postgresql-18-cron_1.6.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg12+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 90.2 KiB | [postgresql-18-cron_1.6.8-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg12+3_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.9 KiB | [postgresql-18-cron_1.6.8-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg12+2_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.9 KiB | [postgresql-18-cron_1.6.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg12+1_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.3 KiB | [postgresql-18-cron_1.6.8-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg13+3_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.9 KiB | [postgresql-18-cron_1.6.8-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg13+2_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.0 KiB | [postgresql-18-cron_1.6.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg13+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 90.2 KiB | [postgresql-18-cron_1.6.8-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg13+3_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.9 KiB | [postgresql-18-cron_1.6.8-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg13+2_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.9 KiB | [postgresql-18-cron_1.6.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg13+1_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 89.9 KiB | [postgresql-18-cron_1.6.8-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg22.04+3_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 90.7 KiB | [postgresql-18-cron_1.6.8-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 90.6 KiB | [postgresql-18-cron_1.6.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.4 KiB | [postgresql-18-cron_1.6.8-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg22.04+3_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 87.9 KiB | [postgresql-18-cron_1.6.8-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 88.0 KiB | [postgresql-18-cron_1.6.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.4 KiB | [postgresql-18-cron_1.6.8-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg24.04+3_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.5 KiB | [postgresql-18-cron_1.6.8-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.4 KiB | [postgresql-18-cron_1.6.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.8 KiB | [postgresql-18-cron_1.6.8-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg24.04+3_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.5 KiB | [postgresql-18-cron_1.6.8-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.5 KiB | [postgresql-18-cron_1.6.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.2 KiB | [postgresql-18-cron_1.6.8-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg26.04+3_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.9 KiB | [postgresql-18-cron_1.6.8-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.8 KiB | [postgresql-18-cron_1.6.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.3 KiB | [postgresql-18-cron_1.6.8-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg26.04+3_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.1 KiB | [postgresql-18-cron_1.6.8-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg26.04+2_arm64.deb) |
| `postgresql-18-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.1 KiB | [postgresql-18-cron_1.6.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.8-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_17` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_cron_17-1.6.8-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_cron_17-1.6.8-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_cron_17` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.0 KiB | [pg_cron_17-1.6.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_cron_17-1.6.8-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_cron_17` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.8 KiB | [pg_cron_17-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_cron_17-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_17` | `1.6.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.2 KiB | [pg_cron_17-1.6.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_cron_17-1.6.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_17` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [pg_cron_17-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_cron_17-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_17` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.9 KiB | [pg_cron_17-1.6.8-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_cron_17-1.6.8-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_cron_17` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.4 KiB | [pg_cron_17-1.6.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_cron_17-1.6.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_cron_17` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.2 KiB | [pg_cron_17-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_cron_17-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_17` | `1.6.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.6 KiB | [pg_cron_17-1.6.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_cron_17-1.6.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_17` | `1.6.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.9 KiB | [pg_cron_17-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_cron_17-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_17` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.8 KiB | [pg_cron_17-1.6.8-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_cron_17-1.6.8-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_cron_17` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.8 KiB | [pg_cron_17-1.6.8-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_cron_17-1.6.8-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_cron_17` | `1.6.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 49.0 KiB | [pg_cron_17-1.6.8-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_cron_17-1.6.8-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_cron_17` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.4 KiB | [pg_cron_17-1.6.8-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_cron_17-1.6.8-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_cron_17` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.4 KiB | [pg_cron_17-1.6.8-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_cron_17-1.6.8-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_17` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.5 KiB | [pg_cron_17-1.6.8-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_cron_17-1.6.8-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_17` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.3 KiB | [pg_cron_17-1.6.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_cron_17-1.6.7-5PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_17` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [pg_cron_17-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_cron_17-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_cron_17` | `1.6.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [pg_cron_17-1.6.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_cron_17-1.6.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.2 KiB | [postgresql-17-cron_1.6.8-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg12+3_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.0 KiB | [postgresql-17-cron_1.6.8-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg12+2_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.0 KiB | [postgresql-17-cron_1.6.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg12+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 90.1 KiB | [postgresql-17-cron_1.6.8-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg12+3_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.8 KiB | [postgresql-17-cron_1.6.8-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg12+2_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.8 KiB | [postgresql-17-cron_1.6.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg12+1_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.3 KiB | [postgresql-17-cron_1.6.8-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg13+3_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.3 KiB | [postgresql-17-cron_1.6.8-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg13+2_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.3 KiB | [postgresql-17-cron_1.6.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg13+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 90.2 KiB | [postgresql-17-cron_1.6.8-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg13+3_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.8 KiB | [postgresql-17-cron_1.6.8-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg13+2_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.8 KiB | [postgresql-17-cron_1.6.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg13+1_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 101.9 KiB | [postgresql-17-cron_1.6.8-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg22.04+3_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.7 KiB | [postgresql-17-cron_1.6.8-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg22.04+2_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.7 KiB | [postgresql-17-cron_1.6.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.3 KiB | [postgresql-17-cron_1.6.8-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg22.04+3_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.8 KiB | [postgresql-17-cron_1.6.8-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg22.04+2_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.9 KiB | [postgresql-17-cron_1.6.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.3 KiB | [postgresql-17-cron_1.6.8-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg24.04+3_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.0 KiB | [postgresql-17-cron_1.6.8-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg24.04+2_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.0 KiB | [postgresql-17-cron_1.6.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.8 KiB | [postgresql-17-cron_1.6.8-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg24.04+3_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.5 KiB | [postgresql-17-cron_1.6.8-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg24.04+2_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.5 KiB | [postgresql-17-cron_1.6.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.2 KiB | [postgresql-17-cron_1.6.8-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg26.04+3_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.1 KiB | [postgresql-17-cron_1.6.8-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.8 KiB | [postgresql-17-cron_1.6.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.3 KiB | [postgresql-17-cron_1.6.8-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg26.04+3_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.2 KiB | [postgresql-17-cron_1.6.8-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg26.04+2_arm64.deb) |
| `postgresql-17-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.0 KiB | [postgresql-17-cron_1.6.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.8-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_16` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [pg_cron_16-1.6.8-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.8-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_cron_16` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.0 KiB | [pg_cron_16-1.6.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.8-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_cron_16` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.8 KiB | [pg_cron_16-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.2 KiB | [pg_cron_16-1.6.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [pg_cron_16-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.2 KiB | [pg_cron_16-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.9 KiB | [pg_cron_16-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_cron_16-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_cron_16-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.9 KiB | [pg_cron_16-1.6.8-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.8-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_cron_16` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.3 KiB | [pg_cron_16-1.6.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_cron_16` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.2 KiB | [pg_cron_16-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.6 KiB | [pg_cron_16-1.6.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.9 KiB | [pg_cron_16-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.7 KiB | [pg_cron_16-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.4 KiB | [pg_cron_16-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pg_cron_16-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.8 KiB | [pg_cron_16-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.9 KiB | [pg_cron_16-1.6.8-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_cron_16-1.6.8-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_cron_16` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 45.8 KiB | [pg_cron_16-1.6.8-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_cron_16-1.6.8-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_cron_16` | `1.6.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 49.0 KiB | [pg_cron_16-1.6.8-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_cron_16-1.6.8-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_cron_16` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.4 KiB | [pg_cron_16-1.6.8-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_cron_16-1.6.8-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_cron_16` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.4 KiB | [pg_cron_16-1.6.8-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_cron_16-1.6.8-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_16` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 46.6 KiB | [pg_cron_16-1.6.8-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_cron_16-1.6.8-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_16` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.3 KiB | [pg_cron_16-1.6.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_cron_16-1.6.7-5PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_16` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [pg_cron_16-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_cron_16-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_cron_16` | `1.6.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.4 KiB | [pg_cron_16-1.6.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_cron_16-1.6.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.1 KiB | [postgresql-16-cron_1.6.8-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg12+3_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.9 KiB | [postgresql-16-cron_1.6.8-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg12+2_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.9 KiB | [postgresql-16-cron_1.6.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg12+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 90.1 KiB | [postgresql-16-cron_1.6.8-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg12+3_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.8 KiB | [postgresql-16-cron_1.6.8-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg12+2_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.8 KiB | [postgresql-16-cron_1.6.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg12+1_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.4 KiB | [postgresql-16-cron_1.6.8-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg13+3_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.3 KiB | [postgresql-16-cron_1.6.8-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg13+2_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.3 KiB | [postgresql-16-cron_1.6.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg13+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 90.1 KiB | [postgresql-16-cron_1.6.8-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg13+3_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.9 KiB | [postgresql-16-cron_1.6.8-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg13+2_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.9 KiB | [postgresql-16-cron_1.6.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg13+1_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 101.9 KiB | [postgresql-16-cron_1.6.8-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg22.04+3_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.4 KiB | [postgresql-16-cron_1.6.8-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg22.04+2_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.4 KiB | [postgresql-16-cron_1.6.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.2 KiB | [postgresql-16-cron_1.6.8-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg22.04+3_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.5 KiB | [postgresql-16-cron_1.6.8-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg22.04+2_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.5 KiB | [postgresql-16-cron_1.6.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.4 KiB | [postgresql-16-cron_1.6.8-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg24.04+3_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.0 KiB | [postgresql-16-cron_1.6.8-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg24.04+2_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.1 KiB | [postgresql-16-cron_1.6.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.7 KiB | [postgresql-16-cron_1.6.8-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg24.04+3_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.5 KiB | [postgresql-16-cron_1.6.8-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg24.04+2_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.5 KiB | [postgresql-16-cron_1.6.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.3 KiB | [postgresql-16-cron_1.6.8-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg26.04+3_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.9 KiB | [postgresql-16-cron_1.6.8-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.9 KiB | [postgresql-16-cron_1.6.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.3 KiB | [postgresql-16-cron_1.6.8-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg26.04+3_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.0 KiB | [postgresql-16-cron_1.6.8-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg26.04+2_arm64.deb) |
| `postgresql-16-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.1 KiB | [postgresql-16-cron_1.6.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.8-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_15` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.7 KiB | [pg_cron_15-1.6.8-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.8-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_cron_15` | `1.6.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.2 KiB | [pg_cron_15-1.6.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.8-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_cron_15` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.9 KiB | [pg_cron_15-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.6.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.5 KiB | [pg_cron_15-1.6.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.8 KiB | [pg_cron_15-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_cron_15-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [pg_cron_15-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.3 KiB | [pg_cron_15-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.9 KiB | [pg_cron_15-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.2 KiB | [pg_cron_15-1.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.5.2-1.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.0 KiB | [pg_cron_15-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.3 KiB | [pg_cron_15-1.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_cron_15-1.4.2-1.rhel8.x86_64.rpm) |
| `pg_cron_15` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.2 KiB | [pg_cron_15-1.6.8-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.8-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_cron_15` | `1.6.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.6 KiB | [pg_cron_15-1.6.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.8-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_cron_15` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.5 KiB | [pg_cron_15-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.6.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.0 KiB | [pg_cron_15-1.6.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.6.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.2 KiB | [pg_cron_15-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.1 KiB | [pg_cron_15-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.8 KiB | [pg_cron_15-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.6 KiB | [pg_cron_15-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pg_cron_15-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.4 KiB | [pg_cron_15-1.5.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.5.2-1.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [pg_cron_15-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.9 KiB | [pg_cron_15-1.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_cron_15-1.4.2-1.rhel8.aarch64.rpm) |
| `pg_cron_15` | `1.6.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.6 KiB | [pg_cron_15-1.6.8-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.6.8-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_cron_15` | `1.6.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 47.8 KiB | [pg_cron_15-1.6.8-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.6.8-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_cron_15` | `1.6.8` | [el10.x86_64](/os/el10.x86_64) | pgdg | 48.9 KiB | [pg_cron_15-1.6.8-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_cron_15-1.6.8-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_cron_15` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.9 KiB | [pg_cron_15-1.6.8-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_cron_15-1.6.8-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_cron_15` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.9 KiB | [pg_cron_15-1.6.8-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_cron_15-1.6.8-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_15` | `1.6.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 47.1 KiB | [pg_cron_15-1.6.8-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_cron_15-1.6.8-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_15` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.8 KiB | [pg_cron_15-1.6.7-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_cron_15-1.6.7-5PGDG.rhel10.2.aarch64.rpm) |
| `pg_cron_15` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.1 KiB | [pg_cron_15-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_cron_15-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_cron_15` | `1.6.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.2 KiB | [pg_cron_15-1.6.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_cron_15-1.6.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 92.0 KiB | [postgresql-15-cron_1.6.8-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg12+3_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 93.6 KiB | [postgresql-15-cron_1.6.8-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg12+2_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 93.5 KiB | [postgresql-15-cron_1.6.8-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg12+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 90.0 KiB | [postgresql-15-cron_1.6.8-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg12+3_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 91.4 KiB | [postgresql-15-cron_1.6.8-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg12+2_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 91.4 KiB | [postgresql-15-cron_1.6.8-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg12+1_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 92.2 KiB | [postgresql-15-cron_1.6.8-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg13+3_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 93.8 KiB | [postgresql-15-cron_1.6.8-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg13+2_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 93.8 KiB | [postgresql-15-cron_1.6.8-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg13+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 90.0 KiB | [postgresql-15-cron_1.6.8-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg13+3_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 91.8 KiB | [postgresql-15-cron_1.6.8-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg13+2_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 91.7 KiB | [postgresql-15-cron_1.6.8-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg13+1_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.4 KiB | [postgresql-15-cron_1.6.8-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg22.04+3_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 104.4 KiB | [postgresql-15-cron_1.6.8-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 104.3 KiB | [postgresql-15-cron_1.6.8-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 100.1 KiB | [postgresql-15-cron_1.6.8-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg22.04+3_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 102.0 KiB | [postgresql-15-cron_1.6.8-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 101.9 KiB | [postgresql-15-cron_1.6.8-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.1 KiB | [postgresql-15-cron_1.6.8-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg24.04+3_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.8 KiB | [postgresql-15-cron_1.6.8-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.7 KiB | [postgresql-15-cron_1.6.8-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.0 KiB | [postgresql-15-cron_1.6.8-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg24.04+3_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.5 KiB | [postgresql-15-cron_1.6.8-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg24.04+2_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.5 KiB | [postgresql-15-cron_1.6.8-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 87.6 KiB | [postgresql-15-cron_1.6.8-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg26.04+3_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.9 KiB | [postgresql-15-cron_1.6.8-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.8 KiB | [postgresql-15-cron_1.6.8-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 85.2 KiB | [postgresql-15-cron_1.6.8-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg26.04+3_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 86.3 KiB | [postgresql-15-cron_1.6.8-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg26.04+2_arm64.deb) |
| `postgresql-15-cron` | `1.6.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 86.3 KiB | [postgresql-15-cron_1.6.8-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.8-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/pg_cron" title="Repository" icon="github" subtitle="github.com/citusdata/pg_cron" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_cron;		# install via package name, for the active PG version

pig install pg_cron -v 19;   # install for PG 19
pig install pg_cron -v 18;   # install for PG 18
pig install pg_cron -v 17;   # install for PG 17
pig install pg_cron -v 16;   # install for PG 16
pig install pg_cron -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_cron';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_cron;
```




## Usage

beware that `cron.database` has to be set before adding to `shared_preload_libraries`

```sql
-- Delete old data on Saturday at 3:30am (GMT)
SELECT cron.schedule('30 3 * * 6', $$DELETE FROM events WHERE event_time < now() - interval '1 week'$$);
 schedule
----------
       42

-- Vacuum every day at 10:00am (GMT)
SELECT cron.schedule('nightly-vacuum', '0 10 * * *', 'VACUUM');
 schedule
----------
       43

-- Change to vacuum at 3:00am (GMT)
SELECT cron.schedule('nightly-vacuum', '0 3 * * *', 'VACUUM');
 schedule
----------
       43

-- Stop scheduling jobs
SELECT cron.unschedule('nightly-vacuum' );
 unschedule 
------------
 t

SELECT cron.unschedule(42);
 unschedule
------------
          t

-- Vacuum every Sunday at 4:00am (GMT) in a database other than the one pg_cron is installed in
SELECT cron.schedule_in_database('weekly-vacuum', '0 4 * * 0', 'VACUUM', 'some_other_database');
 schedule
----------
       44

-- Call a stored procedure every 5 seconds
SELECT cron.schedule('process-updates', '5 seconds', 'CALL process_updates()');

-- Process payroll at 12:00 of the last day of each month
SELECT cron.schedule('process-payroll', '0 12 $ * *', 'CALL process_payroll()');
```

Crontab format:

```text
 ┌───────────── min (0 - 59)
 │ ┌────────────── hour (0 - 23)
 │ │ ┌─────────────── day of month (1 - 31) or last day of the month ($)
 │ │ │ ┌──────────────── month (1 - 12)
 │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
 │ │ │ │ │                  Saturday, or use names; 7 is also Sunday)
 │ │ │ │ │
 │ │ │ │ │
 * * * * *
```
