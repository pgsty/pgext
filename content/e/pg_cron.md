---
title: "pg_cron"
linkTitle: "pg_cron"
description: "Job scheduler for PostgreSQL"
weight: 1070
categories: ["TIME"]
width: full
---

[**pg_cron**](https://github.com/citusdata/pg_cron) : Job scheduler for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1070** | {{< badge content="pg_cron" link="https://github.com/citusdata/pg_cron" >}} | {{< ext "pg_cron" >}} | `1.6.7` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} {{< ext "pg_incremental" >}} {{< ext "timeseries" >}} {{< ext "vectorize" >}} |
|   **See Also**    | {{< ext "timescaledb_toolkit" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} |

> [!Note] require cron.database_name


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_cron` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.7` | {{< bg "18" "pg_cron_18" "green" >}} {{< bg "17" "pg_cron_17" "green" >}} {{< bg "16" "pg_cron_16" "green" >}} {{< bg "15" "pg_cron_15" "green" >}} {{< bg "14" "pg_cron_14" "green" >}} {{< bg "13" "pg_cron_13" "green" >}} | `pg_cron_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6.7` | {{< bg "18" "postgresql-18-cron" "green" >}} {{< bg "17" "postgresql-17-cron" "green" >}} {{< bg "16" "postgresql-16-cron" "green" >}} {{< bg "15" "postgresql-15-cron" "green" >}} {{< bg "14" "postgresql-14-cron" "green" >}} {{< bg "13" "postgresql-13-cron" "green" >}} | `postgresql-$v-cron` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.6.7" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_14 : AVAIL 12" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_13 : AVAIL 13" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.6.7" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_13 : AVAIL 10" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.6.7" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_13 : AVAIL 10" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.6.7" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_13 : AVAIL 10" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.6.7" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.6.7" "pg_cron_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.6.7" "pg_cron_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.6.7" "postgresql-18-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-17-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-16-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-15-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-14-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-13-cron : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.6.7" "postgresql-18-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-17-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-16-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-15-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-14-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-13-cron : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.6.7" "postgresql-18-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-17-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-16-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-15-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-14-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-13-cron : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.6.7" "postgresql-18-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-17-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-16-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-15-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-14-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-13-cron : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.6.7" "postgresql-18-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-17-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-16-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-15-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-14-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-13-cron : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.6.7" "postgresql-18-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-17-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-16-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-15-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-14-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-13-cron : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.6.7" "postgresql-18-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-17-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-16-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-15-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-14-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-13-cron : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.6.7" "postgresql-18-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-17-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-16-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-15-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-14-cron : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.7" "postgresql-13-cron : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_18` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.8 KiB | [pg_cron_18-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_cron_18-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_18` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.2 KiB | [pg_cron_18-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_cron_18-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_18` | `1.6.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.1 KiB | [pg_cron_18-1.6.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_cron_18-1.6.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_18` | `1.6.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.6 KiB | [pg_cron_18-1.6.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_cron_18-1.6.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_18` | `1.6.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.3 KiB | [pg_cron_18-1.6.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_cron_18-1.6.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_18` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [pg_cron_18-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_cron_18-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-cron` | `1.6.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 89.7 KiB | [postgresql-18-cron_1.6.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.7-2.pgdg12+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 87.4 KiB | [postgresql-18-cron_1.6.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.7-2.pgdg12+1_arm64.deb) |
| `postgresql-18-cron` | `1.6.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 89.6 KiB | [postgresql-18-cron_1.6.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.7-2.pgdg13+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 87.3 KiB | [postgresql-18-cron_1.6.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.7-2.pgdg13+1_arm64.deb) |
| `postgresql-18-cron` | `1.6.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 88.7 KiB | [postgresql-18-cron_1.6.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 86.1 KiB | [postgresql-18-cron_1.6.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-cron` | `1.6.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 87.7 KiB | [postgresql-18-cron_1.6.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-cron` | `1.6.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 84.6 KiB | [postgresql-18-cron_1.6.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-18-cron_1.6.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_17` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.8 KiB | [pg_cron_17-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_cron_17-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_17` | `1.6.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.2 KiB | [pg_cron_17-1.6.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_cron_17-1.6.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_17` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [pg_cron_17-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_cron_17-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_17` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.2 KiB | [pg_cron_17-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_cron_17-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_17` | `1.6.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.6 KiB | [pg_cron_17-1.6.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_cron_17-1.6.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_17` | `1.6.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.9 KiB | [pg_cron_17-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_cron_17-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_17` | `1.6.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.2 KiB | [pg_cron_17-1.6.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_cron_17-1.6.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_17` | `1.6.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [pg_cron_17-1.6.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_cron_17-1.6.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_17` | `1.6.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.3 KiB | [pg_cron_17-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_cron_17-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_17` | `1.6.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.6 KiB | [pg_cron_17-1.6.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_cron_17-1.6.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_17` | `1.6.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.6 KiB | [pg_cron_17-1.6.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_cron_17-1.6.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_17` | `1.6.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.0 KiB | [pg_cron_17-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_cron_17-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_17` | `1.6.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.3 KiB | [pg_cron_17-1.6.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_cron_17-1.6.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_17` | `1.6.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.0 KiB | [pg_cron_17-1.6.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_cron_17-1.6.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_17` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [pg_cron_17-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_cron_17-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_cron_17` | `1.6.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [pg_cron_17-1.6.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_cron_17-1.6.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-cron` | `1.6.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 89.6 KiB | [postgresql-17-cron_1.6.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.7-2.pgdg12+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 87.3 KiB | [postgresql-17-cron_1.6.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.7-2.pgdg12+1_arm64.deb) |
| `postgresql-17-cron` | `1.6.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 89.6 KiB | [postgresql-17-cron_1.6.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.7-2.pgdg13+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 87.4 KiB | [postgresql-17-cron_1.6.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.7-2.pgdg13+1_arm64.deb) |
| `postgresql-17-cron` | `1.6.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 100.5 KiB | [postgresql-17-cron_1.6.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 97.9 KiB | [postgresql-17-cron_1.6.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-cron` | `1.6.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 87.6 KiB | [postgresql-17-cron_1.6.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-cron` | `1.6.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 84.5 KiB | [postgresql-17-cron_1.6.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-17-cron_1.6.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_16` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.8 KiB | [pg_cron_16-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.2 KiB | [pg_cron_16-1.6.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [pg_cron_16-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.2 KiB | [pg_cron_16-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.9 KiB | [pg_cron_16-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.7 KiB | [pg_cron_16-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.3 KiB | [pg_cron_16-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_cron_16-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_16` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.2 KiB | [pg_cron_16-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.6 KiB | [pg_cron_16-1.6.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.9 KiB | [pg_cron_16-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.7 KiB | [pg_cron_16-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.4 KiB | [pg_cron_16-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pg_cron_16-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.8 KiB | [pg_cron_16-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_cron_16-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_16` | `1.6.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.1 KiB | [pg_cron_16-1.6.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_cron_16-1.6.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_16` | `1.6.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [pg_cron_16-1.6.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_cron_16-1.6.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_16` | `1.6.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.3 KiB | [pg_cron_16-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_cron_16-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_16` | `1.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.2 KiB | [pg_cron_16-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_cron_16-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_16` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.8 KiB | [pg_cron_16-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_cron_16-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_16` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.6 KiB | [pg_cron_16-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_cron_16-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.4 KiB | [pg_cron_16-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_cron_16-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_16` | `1.6.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.7 KiB | [pg_cron_16-1.6.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_cron_16-1.6.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_16` | `1.6.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.6 KiB | [pg_cron_16-1.6.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_cron_16-1.6.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_16` | `1.6.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.0 KiB | [pg_cron_16-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_cron_16-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_16` | `1.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.8 KiB | [pg_cron_16-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_cron_16-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_16` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.2 KiB | [pg_cron_16-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_cron_16-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_16` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.0 KiB | [pg_cron_16-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_cron_16-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 41.8 KiB | [pg_cron_16-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_cron_16-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_16` | `1.6.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.3 KiB | [pg_cron_16-1.6.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_cron_16-1.6.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_16` | `1.6.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.0 KiB | [pg_cron_16-1.6.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_cron_16-1.6.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_16` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.6 KiB | [pg_cron_16-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_cron_16-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_cron_16` | `1.6.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.4 KiB | [pg_cron_16-1.6.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_cron_16-1.6.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-cron` | `1.6.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 89.6 KiB | [postgresql-16-cron_1.6.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.7-2.pgdg12+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 87.3 KiB | [postgresql-16-cron_1.6.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.7-2.pgdg12+1_arm64.deb) |
| `postgresql-16-cron` | `1.6.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 89.6 KiB | [postgresql-16-cron_1.6.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.7-2.pgdg13+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 87.4 KiB | [postgresql-16-cron_1.6.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.7-2.pgdg13+1_arm64.deb) |
| `postgresql-16-cron` | `1.6.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 100.0 KiB | [postgresql-16-cron_1.6.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 97.5 KiB | [postgresql-16-cron_1.6.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-cron` | `1.6.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 87.4 KiB | [postgresql-16-cron_1.6.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-cron` | `1.6.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 84.5 KiB | [postgresql-16-cron_1.6.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-16-cron_1.6.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
| `pg_cron_15` | `1.6.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_cron_15-1.6.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.6.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.6.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_cron_15-1.6.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.6.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.6.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.2 KiB | [pg_cron_15-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.2 KiB | [pg_cron_15-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.0 KiB | [pg_cron_15-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.8 KiB | [pg_cron_15-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [pg_cron_15-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.0 KiB | [pg_cron_15-1.5.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.5.2-1.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.7 KiB | [pg_cron_15-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.1 KiB | [pg_cron_15-1.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_cron_15-1.4.2-1.rhel9.x86_64.rpm) |
| `pg_cron_15` | `1.6.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.6 KiB | [pg_cron_15-1.6.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.6.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.6.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.9 KiB | [pg_cron_15-1.6.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.6.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.6.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.4 KiB | [pg_cron_15-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.2 KiB | [pg_cron_15-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.6 KiB | [pg_cron_15-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.4 KiB | [pg_cron_15-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.2 KiB | [pg_cron_15-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.4 KiB | [pg_cron_15-1.5.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.5.2-1.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.3 KiB | [pg_cron_15-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 109.9 KiB | [pg_cron_15-1.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_cron_15-1.4.2-1.rhel9.aarch64.rpm) |
| `pg_cron_15` | `1.6.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.2 KiB | [pg_cron_15-1.6.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_cron_15-1.6.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_15` | `1.6.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.1 KiB | [pg_cron_15-1.6.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_cron_15-1.6.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_15` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.1 KiB | [pg_cron_15-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_cron_15-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_cron_15` | `1.6.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.2 KiB | [pg_cron_15-1.6.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_cron_15-1.6.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-cron` | `1.6.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.0 KiB | [postgresql-15-cron_1.6.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.7-2.pgdg12+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.8 KiB | [postgresql-15-cron_1.6.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.7-2.pgdg12+1_arm64.deb) |
| `postgresql-15-cron` | `1.6.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.1 KiB | [postgresql-15-cron_1.6.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.7-2.pgdg13+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 88.7 KiB | [postgresql-15-cron_1.6.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.7-2.pgdg13+1_arm64.deb) |
| `postgresql-15-cron` | `1.6.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.4 KiB | [postgresql-15-cron_1.6.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 100.0 KiB | [postgresql-15-cron_1.6.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-cron` | `1.6.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 88.9 KiB | [postgresql-15-cron_1.6.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-cron` | `1.6.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.4 KiB | [postgresql-15-cron_1.6.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-15-cron_1.6.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_14` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.0 KiB | [pg_cron_14-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.6.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.6 KiB | [pg_cron_14-1.6.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.6.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.9 KiB | [pg_cron_14-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.7 KiB | [pg_cron_14-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.4 KiB | [pg_cron_14-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.2 KiB | [pg_cron_14-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.9 KiB | [pg_cron_14-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.1 KiB | [pg_cron_14-1.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.5.2-1.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.9 KiB | [pg_cron_14-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.1 KiB | [pg_cron_14-1.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.4.2-1.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.5 KiB | [pg_cron_14-1.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.4.1-1.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 97.5 KiB | [pg_cron_14-1.3.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_cron_14-1.3.1-2.rhel8.x86_64.rpm) |
| `pg_cron_14` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pg_cron_14-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.6.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.0 KiB | [pg_cron_14-1.6.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.6.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.6.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.2 KiB | [pg_cron_14-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.0 KiB | [pg_cron_14-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.7 KiB | [pg_cron_14-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.5 KiB | [pg_cron_14-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.2 KiB | [pg_cron_14-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.3 KiB | [pg_cron_14-1.5.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.5.2-1.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [pg_cron_14-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.7 KiB | [pg_cron_14-1.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_cron_14-1.4.2-1.rhel8.aarch64.rpm) |
| `pg_cron_14` | `1.6.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.7 KiB | [pg_cron_14-1.6.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.6.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.6.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_cron_14-1.6.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.6.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.6.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.2 KiB | [pg_cron_14-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.2 KiB | [pg_cron_14-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [pg_cron_14-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.8 KiB | [pg_cron_14-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.5 KiB | [pg_cron_14-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.0 KiB | [pg_cron_14-1.5.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.5.2-1.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.7 KiB | [pg_cron_14-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.9 KiB | [pg_cron_14-1.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_cron_14-1.4.2-1.rhel9.x86_64.rpm) |
| `pg_cron_14` | `1.6.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.6 KiB | [pg_cron_14-1.6.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.6.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.6.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.9 KiB | [pg_cron_14-1.6.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.6.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.6.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_cron_14-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.2 KiB | [pg_cron_14-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.5 KiB | [pg_cron_14-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.3 KiB | [pg_cron_14-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.1 KiB | [pg_cron_14-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.5 KiB | [pg_cron_14-1.5.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.5.2-1.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.2 KiB | [pg_cron_14-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 109.7 KiB | [pg_cron_14-1.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_cron_14-1.4.2-1.rhel9.aarch64.rpm) |
| `pg_cron_14` | `1.6.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.2 KiB | [pg_cron_14-1.6.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_cron_14-1.6.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_14` | `1.6.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.1 KiB | [pg_cron_14-1.6.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_cron_14-1.6.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_14` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.1 KiB | [pg_cron_14-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_cron_14-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_cron_14` | `1.6.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.1 KiB | [pg_cron_14-1.6.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_cron_14-1.6.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-cron` | `1.6.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.8 KiB | [postgresql-14-cron_1.6.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-14-cron_1.6.7-2.pgdg12+1_amd64.deb) |
| `postgresql-14-cron` | `1.6.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.9 KiB | [postgresql-14-cron_1.6.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-14-cron_1.6.7-2.pgdg12+1_arm64.deb) |
| `postgresql-14-cron` | `1.6.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.9 KiB | [postgresql-14-cron_1.6.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-14-cron_1.6.7-2.pgdg13+1_amd64.deb) |
| `postgresql-14-cron` | `1.6.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 88.8 KiB | [postgresql-14-cron_1.6.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-14-cron_1.6.7-2.pgdg13+1_arm64.deb) |
| `postgresql-14-cron` | `1.6.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 102.0 KiB | [postgresql-14-cron_1.6.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-14-cron_1.6.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-cron` | `1.6.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.8 KiB | [postgresql-14-cron_1.6.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-14-cron_1.6.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-cron` | `1.6.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 88.9 KiB | [postgresql-14-cron_1.6.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-14-cron_1.6.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-cron` | `1.6.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.4 KiB | [postgresql-14-cron_1.6.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-14-cron_1.6.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cron_13` | `1.6.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.7 KiB | [pg_cron_13-1.6.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.6.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.6.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.2 KiB | [pg_cron_13-1.6.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.6.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.5 KiB | [pg_cron_13-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.3 KiB | [pg_cron_13-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.1 KiB | [pg_cron_13-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.9 KiB | [pg_cron_13-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 43.5 KiB | [pg_cron_13-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.7 KiB | [pg_cron_13-1.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.5.2-1.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.5 KiB | [pg_cron_13-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.2 KiB | [pg_cron_13-1.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.4.2-1.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.6 KiB | [pg_cron_13-1.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.4.1-1.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 96.5 KiB | [pg_cron_13-1.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.3.1-1.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 96.4 KiB | [pg_cron_13-1.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_cron_13-1.3.0-1.rhel8.x86_64.rpm) |
| `pg_cron_13` | `1.6.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.4 KiB | [pg_cron_13-1.6.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.6.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.6.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.9 KiB | [pg_cron_13-1.6.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.6.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.6.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.2 KiB | [pg_cron_13-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.0 KiB | [pg_cron_13-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.7 KiB | [pg_cron_13-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.5 KiB | [pg_cron_13-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 42.1 KiB | [pg_cron_13-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.3 KiB | [pg_cron_13-1.5.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.5.2-1.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [pg_cron_13-1.5.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.5.1-1.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.0 KiB | [pg_cron_13-1.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_cron_13-1.4.2-1.rhel8.aarch64.rpm) |
| `pg_cron_13` | `1.6.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.7 KiB | [pg_cron_13-1.6.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.6.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.6.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.6 KiB | [pg_cron_13-1.6.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.6.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.6.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.5 KiB | [pg_cron_13-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.1 KiB | [pg_cron_13-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.8 KiB | [pg_cron_13-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.6 KiB | [pg_cron_13-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.4 KiB | [pg_cron_13-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.9 KiB | [pg_cron_13-1.5.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.5.2-1.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 43.8 KiB | [pg_cron_13-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.5.1-1.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.2 KiB | [pg_cron_13-1.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_cron_13-1.4.2-1.rhel9.x86_64.rpm) |
| `pg_cron_13` | `1.6.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.6 KiB | [pg_cron_13-1.6.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.6.7-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.6.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.9 KiB | [pg_cron_13-1.6.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.6.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.6.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_cron_13-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.2 KiB | [pg_cron_13-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.5 KiB | [pg_cron_13-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.4 KiB | [pg_cron_13-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.2 KiB | [pg_cron_13-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.4 KiB | [pg_cron_13-1.5.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.5.2-1.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.1 KiB | [pg_cron_13-1.5.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.5.1-1.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 109.1 KiB | [pg_cron_13-1.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_cron_13-1.4.2-1.rhel9.aarch64.rpm) |
| `pg_cron_13` | `1.6.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.2 KiB | [pg_cron_13-1.6.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_cron_13-1.6.7-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_13` | `1.6.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.3 KiB | [pg_cron_13-1.6.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_cron_13-1.6.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_cron_13` | `1.6.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.2 KiB | [pg_cron_13-1.6.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_cron_13-1.6.7-1PGDG.rhel10.aarch64.rpm) |
| `pg_cron_13` | `1.6.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 45.2 KiB | [pg_cron_13-1.6.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_cron_13-1.6.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-cron` | `1.6.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.6 KiB | [postgresql-13-cron_1.6.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-13-cron_1.6.7-2.pgdg12+1_amd64.deb) |
| `postgresql-13-cron` | `1.6.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.7 KiB | [postgresql-13-cron_1.6.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-13-cron_1.6.7-2.pgdg12+1_arm64.deb) |
| `postgresql-13-cron` | `1.6.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.9 KiB | [postgresql-13-cron_1.6.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-13-cron_1.6.7-2.pgdg13+1_amd64.deb) |
| `postgresql-13-cron` | `1.6.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 88.8 KiB | [postgresql-13-cron_1.6.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-13-cron_1.6.7-2.pgdg13+1_arm64.deb) |
| `postgresql-13-cron` | `1.6.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 101.7 KiB | [postgresql-13-cron_1.6.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-13-cron_1.6.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-cron` | `1.6.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 99.4 KiB | [postgresql-13-cron_1.6.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-13-cron_1.6.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-cron` | `1.6.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 88.8 KiB | [postgresql-13-cron_1.6.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-13-cron_1.6.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-cron` | `1.6.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 86.5 KiB | [postgresql-13-cron_1.6.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-cron/postgresql-13-cron_1.6.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/citusdata/pg_cron" title="Repository" icon="github" subtitle="github.com/citusdata/pg_cron" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_cron;		# install via package name, for the active PG version

pig install pg_cron -v 18;   # install for PG 18
pig install pg_cron -v 17;   # install for PG 17
pig install pg_cron -v 16;   # install for PG 16
pig install pg_cron -v 15;   # install for PG 15
pig install pg_cron -v 14;   # install for PG 14
pig install pg_cron -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_cron';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_cron;
```


## Usage

beware that `cron.database` has to be set before adding to `shared_preload_libraries`

```
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

```
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