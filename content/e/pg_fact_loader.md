---
title: "pg_fact_loader"
linkTitle: "pg_fact_loader"
description: "build fact tables with Postgres"
weight: 9820
categories: ["ETL"]
width: full
---

build fact tables with Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9820** | {{< badge content="pg_fact_loader" link="https://github.com/enova/pg_fact_loader" >}} | {{< ext "pg_fact_loader" >}} | `2.0.1` | {{< category "ETL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "pg_jobmon" >}} {{< ext "mimeo" >}} {{< ext "timescaledb" >}} {{< ext "citus" >}} {{< ext "tablefunc" >}} {{< ext "pg_bulkload" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_fact_loader" >}} | `2.0.1` | {{< bg "18" "pg_fact_loader_18*" "green" >}} {{< bg "17" "pg_fact_loader_17*" "green" >}} {{< bg "16" "pg_fact_loader_16*" "green" >}} {{< bg "15" "pg_fact_loader_15*" "green" >}} {{< bg "14" "pg_fact_loader_14*" "green" >}} | `pg_fact_loader_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_fact_loader" >}} | `2.0.1` | {{< bg "18" "postgresql-18-pg-fact-loader" "red" >}} {{< bg "17" "postgresql-17-pg-fact-loader" "green" >}} {{< bg "16" "postgresql-16-pg-fact-loader" "green" >}} {{< bg "15" "postgresql-15-pg-fact-loader" "green" >}} {{< bg "14" "postgresql-14-pg-fact-loader" "green" >}} | `postgresql-$v-pg-fact-loader` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 3" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 2" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 2" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_18` | 2.0.1 | `el8.x86_64` | pgdg | 36.2 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_fact_loader_18-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_18` | 2.0.1 | `el8.aarch64` | pgdg | 36.1 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_fact_loader_18-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_18` | 2.0.1 | `el9.x86_64` | pgdg | 34.6 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_fact_loader_18-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_18` | 2.0.1 | `el9.aarch64` | pgdg | 34.5 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_fact_loader_18-2.0.1-3PGDG.rhel9.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_17` | 2.0.1 | `el8.x86_64` | pgdg | 36.2 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_fact_loader_17-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_17` | 2.0.1 | `el8.x86_64` | pgdg | 36.0 KiB | [pg_fact_loader_17-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_fact_loader_17-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_17` | 2.0.1 | `el8.aarch64` | pgdg | 36.1 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_fact_loader_17-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_17` | 2.0.1 | `el8.aarch64` | pgdg | 36.0 KiB | [pg_fact_loader_17-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_fact_loader_17-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_17` | 2.0.1 | `el9.x86_64` | pgdg | 34.6 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_fact_loader_17-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_17` | 2.0.1 | `el9.x86_64` | pgdg | 34.5 KiB | [pg_fact_loader_17-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_fact_loader_17-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_17` | 2.0.1 | `el9.aarch64` | pgdg | 34.4 KiB | [pg_fact_loader_17-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_fact_loader_17-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_17` | 2.0.1 | `el9.aarch64` | pgdg | 34.5 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_fact_loader_17-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `postgresql-17-pg-fact-loader` | 2.0.1 | `d12.x86_64` | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb) |
| `postgresql-17-pg-fact-loader` | 2.0.1 | `d12.aarch64` | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb) |
| `postgresql-17-pg-fact-loader` | 2.0.1 | `u22.x86_64` | pgdg | 40.6 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-fact-loader` | 2.0.1 | `u22.aarch64` | pgdg | 40.6 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-fact-loader` | 2.0.1 | `u24.x86_64` | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-fact-loader` | 2.0.1 | `u24.aarch64` | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_16` | 2.0.1 | `el8.x86_64` | pgdg | 36.2 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_fact_loader_16-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_16` | 2.0.1 | `el8.x86_64` | pgdg | 36.0 KiB | [pg_fact_loader_16-2.0.1-1PGDG.f42.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_fact_loader_16-2.0.1-1PGDG.f42.noarch.rpm) |
| `pg_fact_loader_16` | 2.0.1 | `el8.x86_64` | pgdg | 36.0 KiB | [pg_fact_loader_16-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_fact_loader_16-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_16` | 2.0.1 | `el8.aarch64` | pgdg | 36.1 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_fact_loader_16-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_16` | 2.0.1 | `el8.aarch64` | pgdg | 36.0 KiB | [pg_fact_loader_16-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_fact_loader_16-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_16` | 2.0.1 | `el9.x86_64` | pgdg | 34.5 KiB | [pg_fact_loader_16-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_fact_loader_16-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_16` | 2.0.1 | `el9.x86_64` | pgdg | 34.6 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_fact_loader_16-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_16` | 2.0.1 | `el9.aarch64` | pgdg | 34.3 KiB | [pg_fact_loader_16-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_fact_loader_16-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_16` | 2.0.1 | `el9.aarch64` | pgdg | 34.5 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_fact_loader_16-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `postgresql-16-pg-fact-loader` | 2.0.1 | `d12.x86_64` | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb) |
| `postgresql-16-pg-fact-loader` | 2.0.1 | `d12.aarch64` | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb) |
| `postgresql-16-pg-fact-loader` | 2.0.1 | `u22.x86_64` | pgdg | 40.6 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-fact-loader` | 2.0.1 | `u22.aarch64` | pgdg | 40.6 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-fact-loader` | 2.0.1 | `u24.x86_64` | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-fact-loader` | 2.0.1 | `u24.aarch64` | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_15` | 2.0.1 | `el8.x86_64` | pgdg | 36.0 KiB | [pg_fact_loader_15-2.0.1-1PGDG.f42.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_fact_loader_15-2.0.1-1PGDG.f42.noarch.rpm) |
| `pg_fact_loader_15` | 2.0.1 | `el8.x86_64` | pgdg | 36.0 KiB | [pg_fact_loader_15-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_fact_loader_15-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_15` | 2.0.1 | `el8.x86_64` | pgdg | 36.2 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_fact_loader_15-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_15` | 2.0.1 | `el8.aarch64` | pgdg | 36.1 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_fact_loader_15-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_15` | 2.0.1 | `el8.aarch64` | pgdg | 36.0 KiB | [pg_fact_loader_15-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_fact_loader_15-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_15` | 2.0.1 | `el9.x86_64` | pgdg | 34.5 KiB | [pg_fact_loader_15-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_fact_loader_15-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_15` | 2.0.1 | `el9.x86_64` | pgdg | 34.6 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_fact_loader_15-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_15` | 2.0.1 | `el9.aarch64` | pgdg | 34.3 KiB | [pg_fact_loader_15-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_fact_loader_15-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_15` | 2.0.1 | `el9.aarch64` | pgdg | 34.6 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_fact_loader_15-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `postgresql-15-pg-fact-loader` | 2.0.1 | `d12.x86_64` | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb) |
| `postgresql-15-pg-fact-loader` | 2.0.1 | `d12.aarch64` | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb) |
| `postgresql-15-pg-fact-loader` | 2.0.1 | `u22.x86_64` | pgdg | 40.6 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-fact-loader` | 2.0.1 | `u22.aarch64` | pgdg | 40.6 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-fact-loader` | 2.0.1 | `u24.x86_64` | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-fact-loader` | 2.0.1 | `u24.aarch64` | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_14` | 2.0.1 | `el8.x86_64` | pgdg | 36.2 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_fact_loader_14-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_14` | 2.0.1 | `el8.x86_64` | pgdg | 36.0 KiB | [pg_fact_loader_14-2.0.1-1PGDG.f42.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_fact_loader_14-2.0.1-1PGDG.f42.noarch.rpm) |
| `pg_fact_loader_14` | 2.0.1 | `el8.x86_64` | pgdg | 36.0 KiB | [pg_fact_loader_14-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_fact_loader_14-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_14` | 2.0.1 | `el8.aarch64` | pgdg | 36.0 KiB | [pg_fact_loader_14-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_fact_loader_14-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_14` | 2.0.1 | `el8.aarch64` | pgdg | 36.1 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_fact_loader_14-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_14` | 2.0.1 | `el9.x86_64` | pgdg | 34.5 KiB | [pg_fact_loader_14-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_fact_loader_14-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_14` | 2.0.1 | `el9.x86_64` | pgdg | 34.6 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_fact_loader_14-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_14` | 2.0.1 | `el9.aarch64` | pgdg | 34.5 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_fact_loader_14-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_14` | 2.0.1 | `el9.aarch64` | pgdg | 34.3 KiB | [pg_fact_loader_14-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_fact_loader_14-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `postgresql-14-pg-fact-loader` | 2.0.1 | `d12.x86_64` | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb) |
| `postgresql-14-pg-fact-loader` | 2.0.1 | `d12.aarch64` | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb) |
| `postgresql-14-pg-fact-loader` | 2.0.1 | `u22.x86_64` | pgdg | 40.6 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-fact-loader` | 2.0.1 | `u22.aarch64` | pgdg | 40.6 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-fact-loader` | 2.0.1 | `u24.x86_64` | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-fact-loader` | 2.0.1 | `u24.aarch64` | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/enova/pg_fact_loader" title="Repository" icon="github" subtitle="github.com/enova/pg_fact_loader" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_fact_loader; # install by extension name, for the current active PG version
pig ext install pg_fact_loader; # install via package alias, for the active PG version
pig ext install pg_fact_loader -v 18;   # install for PG 18
pig ext install pg_fact_loader -v 17;   # install for PG 17
pig ext install pg_fact_loader -v 16;   # install for PG 16
pig ext install pg_fact_loader -v 15;   # install for PG 15
pig ext install pg_fact_loader -v 14;   # install for PG 14
pig ext install pg_fact_loader -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_fact_loader CASCADE SCHEMA fact_loader;
```

