---
title: "pglogical"
linkTitle: "pglogical"
description: "PostgreSQL Logical Replication"
weight: 9500
categories: ["ETL"]
width: full
---

[**pglogical**](https://github.com/2ndQuadrant/pglogical) : PostgreSQL Logical Replication


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9500** | {{< badge content="pglogical" link="https://github.com/2ndQuadrant/pglogical" >}} | {{< ext "pglogical" >}} | `2.4.6` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pglogical` |
|    **Need By**    | {{< ext "pgl_ddl_deploy" >}} {{< ext "pglogical_ticker" >}} |
|   **See Also**    | {{< ext "decoderbufs" >}} {{< ext "wal2json" >}} {{< ext "dblink" >}} {{< ext "postgres_fdw" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "repmgr" >}} {{< ext "kafka_fdw" >}} |
|    **Siblings**   | {{< ext "pglogical_origin" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pglogical` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "pglogical_18" "green" >}} {{< bg "17" "pglogical_17" "green" >}} {{< bg "16" "pglogical_16" "green" >}} {{< bg "15" "pglogical_15" "green" >}} {{< bg "14" "pglogical_14" "green" >}} | `pglogical_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "postgresql-18-pglogical" "green" >}} {{< bg "17" "postgresql-17-pglogical" "green" >}} {{< bg "16" "postgresql-16-pglogical" "green" >}} {{< bg "15" "postgresql-15-pglogical" "green" >}} {{< bg "14" "postgresql-14-pglogical" "green" >}} | `postgresql-$v-pglogical` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_18` | `2.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 154.6 KiB | [pglogical_18-2.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pglogical_18-2.4.6-1PGDG.rhel8.x86_64.rpm) |
| `pglogical_18` | `2.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.0 KiB | [pglogical_18-2.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pglogical_18-2.4.6-1PGDG.rhel8.aarch64.rpm) |
| `pglogical_18` | `2.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.3 KiB | [pglogical_18-2.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pglogical_18-2.4.6-1PGDG.rhel9.x86_64.rpm) |
| `pglogical_18` | `2.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.1 KiB | [pglogical_18-2.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pglogical_18-2.4.6-1PGDG.rhel9.aarch64.rpm) |
| `pglogical_18` | `2.4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 148.4 KiB | [pglogical_18-2.4.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pglogical_18-2.4.6-1PGDG.rhel10.x86_64.rpm) |
| `pglogical_18` | `2.4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.1 KiB | [pglogical_18-2.4.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pglogical_18-2.4.6-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 345.7 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 336.5 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 346.5 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 337.7 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 357.3 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 345.8 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 344.7 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 335.5 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 343.4 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 332.5 KiB | [postgresql-18-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_17` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.8 KiB | [pglogical_17-2.4.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pglogical_17-2.4.5-1PGDG.rhel8.x86_64.rpm) |
| `pglogical_17` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.2 KiB | [pglogical_17-2.4.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pglogical_17-2.4.5-1PGDG.rhel8.aarch64.rpm) |
| `pglogical_17` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.5 KiB | [pglogical_17-2.4.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pglogical_17-2.4.5-1PGDG.rhel9.x86_64.rpm) |
| `pglogical_17` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.5 KiB | [pglogical_17-2.4.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pglogical_17-2.4.5-1PGDG.rhel9.aarch64.rpm) |
| `pglogical_17` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 148.3 KiB | [pglogical_17-2.4.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pglogical_17-2.4.5-3PGDG.rhel10.x86_64.rpm) |
| `pglogical_17` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 144.7 KiB | [pglogical_17-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pglogical_17-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 345.7 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 336.0 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 346.1 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 337.1 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 434.2 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 422.2 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 344.9 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 334.8 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 343.7 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 332.0 KiB | [postgresql-17-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_16` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.1 KiB | [pglogical_16-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pglogical_16-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pglogical_16` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.6 KiB | [pglogical_16-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pglogical_16-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pglogical_16` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.1 KiB | [pglogical_16-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pglogical_16-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pglogical_16` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.0 KiB | [pglogical_16-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pglogical_16-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pglogical_16` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 148.5 KiB | [pglogical_16-2.4.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pglogical_16-2.4.5-3PGDG.rhel10.x86_64.rpm) |
| `pglogical_16` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 144.9 KiB | [pglogical_16-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pglogical_16-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 344.5 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 334.6 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 345.8 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 336.0 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 431.0 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 419.8 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 343.4 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 334.2 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 342.1 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 330.7 KiB | [postgresql-16-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_15` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.2 KiB | [pglogical_15-2.4.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pglogical_15-2.4.3-1.rhel8.x86_64.rpm) |
| `pglogical_15` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.5 KiB | [pglogical_15-2.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pglogical_15-2.4.2-1.rhel8.x86_64.rpm) |
| `pglogical_15` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.4 KiB | [pglogical_15-2.4.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pglogical_15-2.4.3-1.rhel8.aarch64.rpm) |
| `pglogical_15` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.6 KiB | [pglogical_15-2.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pglogical_15-2.4.2-1.rhel8.aarch64.rpm) |
| `pglogical_15` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.6 KiB | [pglogical_15-2.4.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pglogical_15-2.4.3-1.rhel9.x86_64.rpm) |
| `pglogical_15` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.1 KiB | [pglogical_15-2.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pglogical_15-2.4.2-1.rhel9.x86_64.rpm) |
| `pglogical_15` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.0 KiB | [pglogical_15-2.4.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pglogical_15-2.4.3-1.rhel9.aarch64.rpm) |
| `pglogical_15` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.1 KiB | [pglogical_15-2.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pglogical_15-2.4.2-1.rhel9.aarch64.rpm) |
| `pglogical_15` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 151.9 KiB | [pglogical_15-2.4.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pglogical_15-2.4.5-3PGDG.rhel10.x86_64.rpm) |
| `pglogical_15` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.2 KiB | [pglogical_15-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pglogical_15-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 346.9 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 335.7 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 348.1 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 336.9 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 436.4 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 424.8 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 347.0 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 336.9 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 345.3 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 333.1 KiB | [postgresql-15-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_14` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.7 KiB | [pglogical_14-2.4.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pglogical_14-2.4.3-1.rhel8.x86_64.rpm) |
| `pglogical_14` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.9 KiB | [pglogical_14-2.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pglogical_14-2.4.2-1.rhel8.x86_64.rpm) |
| `pglogical_14` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.4 KiB | [pglogical_14-2.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pglogical_14-2.4.1-1.rhel8.x86_64.rpm) |
| `pglogical_14` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 149.9 KiB | [pglogical_14-2.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pglogical_14-2.4.0-1.rhel8.x86_64.rpm) |
| `pglogical_14` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.3 KiB | [pglogical_14-2.4.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pglogical_14-2.4.3-1.rhel8.aarch64.rpm) |
| `pglogical_14` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.6 KiB | [pglogical_14-2.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pglogical_14-2.4.2-1.rhel8.aarch64.rpm) |
| `pglogical_14` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.1 KiB | [pglogical_14-2.4.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pglogical_14-2.4.3-1.rhel9.x86_64.rpm) |
| `pglogical_14` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.1 KiB | [pglogical_14-2.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pglogical_14-2.4.2-1.rhel9.x86_64.rpm) |
| `pglogical_14` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.7 KiB | [pglogical_14-2.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pglogical_14-2.4.1-1.rhel9.x86_64.rpm) |
| `pglogical_14` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 145.6 KiB | [pglogical_14-2.4.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pglogical_14-2.4.3-1.rhel9.aarch64.rpm) |
| `pglogical_14` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 145.6 KiB | [pglogical_14-2.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pglogical_14-2.4.2-1.rhel9.aarch64.rpm) |
| `pglogical_14` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 151.3 KiB | [pglogical_14-2.4.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pglogical_14-2.4.5-3PGDG.rhel10.x86_64.rpm) |
| `pglogical_14` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.7 KiB | [pglogical_14-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pglogical_14-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 347.2 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 335.5 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 347.2 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 336.8 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 434.5 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 423.2 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 347.2 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 336.0 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 344.4 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 333.3 KiB | [postgresql-14-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/2ndQuadrant/pglogical" title="Repository" icon="github" subtitle="github.com/2ndQuadrant/pglogical" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglogical-2.4.6.tar.gz" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pglogical;		# install via package name, for the active PG version

pig install pglogical -v 18;   # install for PG 18
pig install pglogical -v 17;   # install for PG 17
pig install pglogical -v 16;   # install for PG 16
pig install pglogical -v 15;   # install for PG 15
pig install pglogical -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pglogical';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pglogical;
```



## Usage

> [pglogical: PostgreSQL Logical Replication](https://github.com/2ndQuadrant/pglogical)

A logical replication system for PostgreSQL using a publish/subscribe model. Requires no triggers or external programs.

### Enabling

Add to `postgresql.conf`:

```ini
wal_level = 'logical'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
shared_preload_libraries = 'pglogical'
```

```sql
CREATE EXTENSION pglogical;
```

### Provider (Publisher) Setup

```sql
-- Create a node on the provider
SELECT pglogical.create_node(
    node_name := 'provider1',
    dsn := 'host=providerhost port=5432 dbname=mydb'
);

-- Add all tables in public schema to default replication set
SELECT pglogical.replication_set_add_all_tables('default', ARRAY['public']);

-- Add all sequences in public schema
SELECT pglogical.replication_set_add_all_sequences('default', ARRAY['public']);
```

### Subscriber Setup

```sql
-- Create a node on the subscriber
SELECT pglogical.create_node(
    node_name := 'subscriber1',
    dsn := 'host=subscriberhost port=5432 dbname=mydb'
);

-- Create a subscription to the provider
SELECT pglogical.create_subscription(
    subscription_name := 'subscription1',
    provider_dsn := 'host=providerhost port=5432 dbname=mydb'
);
```

### Replication Set Management

```sql
-- Create a custom replication set
SELECT pglogical.create_replication_set('my_set');

-- Add a specific table
SELECT pglogical.replication_set_add_table('my_set', 'my_table', true);

-- Remove a table
SELECT pglogical.replication_set_remove_table('my_set', 'my_table');
```

### Row and Column Filtering

```sql
-- Row filtering: only replicate rows matching a condition
SELECT pglogical.replication_set_add_table(
    set_name := 'default',
    relation := 'my_table',
    row_filter := 'id > 1000'
);

-- Column filtering: only replicate specific columns
SELECT pglogical.replication_set_add_table(
    set_name := 'default',
    relation := 'my_table',
    columns := '{id, name, updated_at}'
);
```

### Subscription Management

```sql
-- Check subscription status
SELECT * FROM pglogical.show_subscription_status();

-- Drop subscription
SELECT pglogical.drop_subscription('subscription1');
```

### Key Features

- Selective replication (per-table, row filtering, column filtering)
- Replication between different PostgreSQL major versions
- Delayed replication
- No need for superuser on subscriber
