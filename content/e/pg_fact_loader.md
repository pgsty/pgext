---
title: "pg_fact_loader"
linkTitle: "pg_fact_loader"
description: "build fact tables with Postgres"
weight: 9820
categories: ["ETL"]
width: full
---

[**pg_fact_loader**](https://github.com/enova/pg_fact_loader) : build fact tables with Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9820** | {{< badge content="pg_fact_loader" link="https://github.com/enova/pg_fact_loader" >}} | {{< ext "pg_fact_loader" >}} | `2.0.1` | {{< category "ETL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `fact_loader` |
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "pg_jobmon" >}} {{< ext "mimeo" >}} {{< ext "timescaledb" >}} {{< ext "citus" >}} {{< ext "tablefunc" >}} {{< ext "pg_bulkload" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_fact_loader` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.1` | {{< bg "18" "pg_fact_loader_18" "green" >}} {{< bg "17" "pg_fact_loader_17" "green" >}} {{< bg "16" "pg_fact_loader_16" "green" >}} {{< bg "15" "pg_fact_loader_15" "green" >}} {{< bg "14" "pg_fact_loader_14" "green" >}} | `pg_fact_loader_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.1` | {{< bg "18" "postgresql-18-pg-fact-loader" "red" >}} {{< bg "17" "postgresql-17-pg-fact-loader" "green" >}} {{< bg "16" "postgresql-16-pg-fact-loader" "green" >}} {{< bg "15" "postgresql-15-pg-fact-loader" "green" >}} {{< bg "14" "postgresql-14-pg-fact-loader" "green" >}} | `postgresql-$v-pg-fact-loader` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "pg_fact_loader_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-fact-loader : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.1" "postgresql-17-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-pg-fact-loader : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-pg-fact-loader : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_18` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.2 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_fact_loader_18-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_18` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_fact_loader_18-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_18` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.6 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_fact_loader_18-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_18` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.5 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_fact_loader_18-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_18` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.1 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_fact_loader_18-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `pg_fact_loader_18` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 35.0 KiB | [pg_fact_loader_18-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_fact_loader_18-2.0.1-3PGDG.rhel10.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_17` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.2 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_fact_loader_17-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_fact_loader_17-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_fact_loader_17-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_fact_loader_17-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.0 KiB | [pg_fact_loader_17-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_fact_loader_17-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.6 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_fact_loader_17-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [pg_fact_loader_17-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_fact_loader_17-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.5 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_fact_loader_17-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.4 KiB | [pg_fact_loader_17-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_fact_loader_17-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.1 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_fact_loader_17-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `pg_fact_loader_17` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 35.0 KiB | [pg_fact_loader_17-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_fact_loader_17-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `postgresql-17-pg-fact-loader` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb) |
| `postgresql-17-pg-fact-loader` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb) |
| `postgresql-17-pg-fact-loader` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg130+1_amd64.deb) |
| `postgresql-17-pg-fact-loader` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg130+1_arm64.deb) |
| `postgresql-17-pg-fact-loader` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.6 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-fact-loader` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 40.6 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-fact-loader` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-fact-loader` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 40.5 KiB | [postgresql-17-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-17-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_16` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.2 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_fact_loader_16-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_fact_loader_16-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_fact_loader_16-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_fact_loader_16-2.0.1-1PGDG.f42.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_fact_loader_16-2.0.1-1PGDG.f42.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_fact_loader_16-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.0 KiB | [pg_fact_loader_16-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_fact_loader_16-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.6 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_fact_loader_16-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [pg_fact_loader_16-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_fact_loader_16-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.5 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_fact_loader_16-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.3 KiB | [pg_fact_loader_16-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_fact_loader_16-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.1 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_fact_loader_16-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `pg_fact_loader_16` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 35.0 KiB | [pg_fact_loader_16-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_fact_loader_16-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `postgresql-16-pg-fact-loader` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb) |
| `postgresql-16-pg-fact-loader` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb) |
| `postgresql-16-pg-fact-loader` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg130+1_amd64.deb) |
| `postgresql-16-pg-fact-loader` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg130+1_arm64.deb) |
| `postgresql-16-pg-fact-loader` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.6 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-fact-loader` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 40.6 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-fact-loader` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-fact-loader` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 40.5 KiB | [postgresql-16-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-16-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.2 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_fact_loader_15-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_fact_loader_15-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_fact_loader_15-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_fact_loader_15-2.0.1-1PGDG.f42.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_fact_loader_15-2.0.1-1PGDG.f42.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_fact_loader_15-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.0 KiB | [pg_fact_loader_15-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_fact_loader_15-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.6 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_fact_loader_15-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [pg_fact_loader_15-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_fact_loader_15-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.6 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_fact_loader_15-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.3 KiB | [pg_fact_loader_15-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_fact_loader_15-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.1 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_fact_loader_15-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `pg_fact_loader_15` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 35.0 KiB | [pg_fact_loader_15-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_fact_loader_15-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `postgresql-15-pg-fact-loader` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb) |
| `postgresql-15-pg-fact-loader` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb) |
| `postgresql-15-pg-fact-loader` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg130+1_amd64.deb) |
| `postgresql-15-pg-fact-loader` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg130+1_arm64.deb) |
| `postgresql-15-pg-fact-loader` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.6 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-fact-loader` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 40.6 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-fact-loader` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-fact-loader` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 40.5 KiB | [postgresql-15-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-15-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_fact_loader_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.2 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_fact_loader_14-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_fact_loader_14-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_fact_loader_14-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 36.0 KiB | [pg_fact_loader_14-2.0.1-1PGDG.f42.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_fact_loader_14-2.0.1-1PGDG.f42.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.1 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_fact_loader_14-2.0.1-3PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 36.0 KiB | [pg_fact_loader_14-2.0.1-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_fact_loader_14-2.0.1-2PGDG.rhel8.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.6 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_fact_loader_14-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [pg_fact_loader_14-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_fact_loader_14-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.5 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_fact_loader_14-2.0.1-3PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.3 KiB | [pg_fact_loader_14-2.0.1-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_fact_loader_14-2.0.1-2PGDG.rhel9.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 35.1 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_fact_loader_14-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `pg_fact_loader_14` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 35.0 KiB | [pg_fact_loader_14-2.0.1-3PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_fact_loader_14-2.0.1-3PGDG.rhel10.noarch.rpm) |
| `postgresql-14-pg-fact-loader` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg120+1_amd64.deb) |
| `postgresql-14-pg-fact-loader` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg120+1_arm64.deb) |
| `postgresql-14-pg-fact-loader` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg130+1_amd64.deb) |
| `postgresql-14-pg-fact-loader` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg130+1_arm64.deb) |
| `postgresql-14-pg-fact-loader` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.6 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-fact-loader` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 40.6 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-fact-loader` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-fact-loader` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 40.5 KiB | [postgresql-14-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-fact-loader/postgresql-14-pg-fact-loader_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/enova/pg_fact_loader" title="Repository" icon="github" subtitle="github.com/enova/pg_fact_loader" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_fact_loader;		# install via package name, for the active PG version

pig install pg_fact_loader -v 18;   # install for PG 18
pig install pg_fact_loader -v 17;   # install for PG 17
pig install pg_fact_loader -v 16;   # install for PG 16
pig install pg_fact_loader -v 15;   # install for PG 15
pig install pg_fact_loader -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_fact_loader;
```



## Usage

> [pg_fact_loader: build fact tables with Postgres](https://github.com/enova/pg_fact_loader)

Build and maintain fact tables using queue-based change data capture. Processes audit/change log tables to incrementally update fact tables.

### Enabling

```sql
CREATE EXTENSION pg_fact_loader;
```

Optionally with pglogical for replica-based setup:

```sql
CREATE EXTENSION pglogical;
CREATE EXTENSION pglogical_ticker;
CREATE EXTENSION pg_fact_loader;
```

### Workflow

1. **Replicate source tables** to a reporting database (via pglogical or other means)
2. **Create audit/change log tables** on the OLTP system for source tables
3. **Create a fact table** structure for aggregated data
4. **Create a merge function** that takes a key ID and returns one row of the fact table
5. **Configure pg_fact_loader** to wire queue tables to fact tables
6. **Backfill** the fact table initially
7. **Schedule** the worker to process changes continuously

### Configuration Tables

```sql
-- Register a fact table
INSERT INTO fact_loader.fact_tables (fact_table_relid, fact_table_agg_proid, ...)
VALUES ('public.customers_fact'::regclass, 'customers_fact_merge'::regproc, ...);

-- Register queue (audit) tables
INSERT INTO fact_loader.queue_tables (queue_table_relid, queue_of_base_table_relid, ...)
VALUES ('audit.customers_audit'::regclass, 'public.customers'::regclass, ...);

-- Connect queue tables to fact tables with merge functions
INSERT INTO fact_loader.queue_table_deps
    (fact_table_id, queue_table_id, insert_merge_proid, update_merge_proid, delete_merge_proid)
VALUES (1, 1, 'customers_fact_merge'::regproc, 'customers_fact_merge'::regproc, 'customers_fact_merge'::regproc);

-- Define how to retrieve the key from queue entries
INSERT INTO fact_loader.key_retrieval_sequences
    (queue_table_dep_id, return_columns, is_fact_key)
VALUES (1, '{customer_id}', true);
```

### Running the Worker

```sql
-- Process pending changes
SELECT fact_loader.worker();

-- Schedule this to run periodically (e.g., every few seconds via pg_cron)
```

### Initial Backfill

```sql
-- Run the merge function for every existing row
SELECT customers_fact_merge(customer_id) FROM customers;
```

### Adding Batch ID Fields

```sql
SELECT fact_loader.add_batch_id_fields();
```

### Key Features

- Queue-based incremental fact table updates
- Supports insert, update, and delete events
- Handles multi-level key retrieval (joins through multiple tables)
- Fact table dependency chains (child facts updated after parent)
- Checks replication lag before processing (when used with pglogical)
