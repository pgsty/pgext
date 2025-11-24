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
|    **Need By**    | {{< ext "pgl_ddl_deploy" >}} {{< ext "pglogical_ticker" >}} |
|   **See Also**    | {{< ext "decoderbufs" >}} {{< ext "wal2json" >}} {{< ext "dblink" >}} {{< ext "postgres_fdw" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "repmgr" >}} {{< ext "kafka_fdw" >}} |
|    **Siblings**   | {{< ext "pglogical_origin" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pglogical` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "pglogical_18*" "green" >}} {{< bg "17" "pglogical_17*" "green" >}} {{< bg "16" "pglogical_16*" "green" >}} {{< bg "15" "pglogical_15*" "green" >}} {{< bg "14" "pglogical_14*" "green" >}} {{< bg "13" "pglogical_13*" "green" >}} | `pglogical_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.6` | {{< bg "18" "postgresql-18-pglogical" "green" >}} {{< bg "17" "postgresql-17-pglogical" "green" >}} {{< bg "16" "postgresql-16-pglogical" "green" >}} {{< bg "15" "postgresql-15-pglogical" "green" >}} {{< bg "14" "postgresql-14-pglogical" "green" >}} {{< bg "13" "postgresql-13-pglogical" "green" >}} | `postgresql-$v-pglogical` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_13 : AVAIL 5" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-13-pglogical : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_18` | `2.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 154.6 KiB | [pglogical_18-2.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pglogical_18-2.4.6-1PGDG.rhel8.x86_64.rpm) |
| `pglogical_18` | `2.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.0 KiB | [pglogical_18-2.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pglogical_18-2.4.6-1PGDG.rhel8.aarch64.rpm) |
| `pglogical_18` | `2.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.3 KiB | [pglogical_18-2.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pglogical_18-2.4.6-1PGDG.rhel9.x86_64.rpm) |
| `pglogical_18` | `2.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.1 KiB | [pglogical_18-2.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pglogical_18-2.4.6-1PGDG.rhel9.aarch64.rpm) |
| `pglogical_18` | `2.4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 148.4 KiB | [pglogical_18-2.4.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pglogical_18-2.4.6-1PGDG.rhel10.x86_64.rpm) |
| `pglogical_18` | `2.4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.1 KiB | [pglogical_18-2.4.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10.0-aarch64/pglogical_18-2.4.6-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 344.8 KiB | [postgresql-18-pglogical_2.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 335.8 KiB | [postgresql-18-pglogical_2.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 345.4 KiB | [postgresql-18-pglogical_2.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 336.7 KiB | [postgresql-18-pglogical_2.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 355.8 KiB | [postgresql-18-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 344.9 KiB | [postgresql-18-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 343.4 KiB | [postgresql-18-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 334.2 KiB | [postgresql-18-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-18-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_17` | `2.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.8 KiB | [pglogical_17-2.4.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pglogical_17-2.4.5-1PGDG.rhel8.x86_64.rpm) |
| `pglogical_17` | `2.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.2 KiB | [pglogical_17-2.4.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pglogical_17-2.4.5-1PGDG.rhel8.aarch64.rpm) |
| `pglogical_17` | `2.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.5 KiB | [pglogical_17-2.4.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pglogical_17-2.4.5-1PGDG.rhel9.x86_64.rpm) |
| `pglogical_17` | `2.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.5 KiB | [pglogical_17-2.4.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pglogical_17-2.4.5-1PGDG.rhel9.aarch64.rpm) |
| `pglogical_17` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 148.3 KiB | [pglogical_17-2.4.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pglogical_17-2.4.5-3PGDG.rhel10.x86_64.rpm) |
| `pglogical_17` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 144.7 KiB | [pglogical_17-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/pglogical_17-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 344.7 KiB | [postgresql-17-pglogical_2.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 334.8 KiB | [postgresql-17-pglogical_2.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 345.2 KiB | [postgresql-17-pglogical_2.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 336.6 KiB | [postgresql-17-pglogical_2.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 432.4 KiB | [postgresql-17-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 420.8 KiB | [postgresql-17-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 343.6 KiB | [postgresql-17-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 333.7 KiB | [postgresql-17-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-17-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_16` | `2.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.1 KiB | [pglogical_16-2.4.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pglogical_16-2.4.4-1PGDG.rhel8.x86_64.rpm) |
| `pglogical_16` | `2.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.6 KiB | [pglogical_16-2.4.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pglogical_16-2.4.4-1PGDG.rhel8.aarch64.rpm) |
| `pglogical_16` | `2.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.1 KiB | [pglogical_16-2.4.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pglogical_16-2.4.4-1PGDG.rhel9.x86_64.rpm) |
| `pglogical_16` | `2.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.0 KiB | [pglogical_16-2.4.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pglogical_16-2.4.4-1PGDG.rhel9.aarch64.rpm) |
| `pglogical_16` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 148.5 KiB | [pglogical_16-2.4.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pglogical_16-2.4.5-3PGDG.rhel10.x86_64.rpm) |
| `pglogical_16` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 144.9 KiB | [pglogical_16-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/pglogical_16-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 343.5 KiB | [postgresql-16-pglogical_2.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 333.7 KiB | [postgresql-16-pglogical_2.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 344.4 KiB | [postgresql-16-pglogical_2.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 335.4 KiB | [postgresql-16-pglogical_2.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 429.9 KiB | [postgresql-16-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 417.9 KiB | [postgresql-16-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 342.4 KiB | [postgresql-16-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 332.8 KiB | [postgresql-16-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-16-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

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
| `pglogical_15` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.2 KiB | [pglogical_15-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/pglogical_15-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 346.2 KiB | [postgresql-15-pglogical_2.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 335.2 KiB | [postgresql-15-pglogical_2.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 347.0 KiB | [postgresql-15-pglogical_2.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 336.2 KiB | [postgresql-15-pglogical_2.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 435.7 KiB | [postgresql-15-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 423.9 KiB | [postgresql-15-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 346.3 KiB | [postgresql-15-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 335.3 KiB | [postgresql-15-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-15-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

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
| `pglogical_14` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.7 KiB | [pglogical_14-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/pglogical_14-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 346.6 KiB | [postgresql-14-pglogical_2.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 334.7 KiB | [postgresql-14-pglogical_2.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 346.2 KiB | [postgresql-14-pglogical_2.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 335.7 KiB | [postgresql-14-pglogical_2.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 433.6 KiB | [postgresql-14-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 422.4 KiB | [postgresql-14-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 345.9 KiB | [postgresql-14-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 335.1 KiB | [postgresql-14-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-14-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglogical_13` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 149.9 KiB | [pglogical_13-2.4.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pglogical_13-2.4.3-1.rhel8.x86_64.rpm) |
| `pglogical_13` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 149.1 KiB | [pglogical_13-2.4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pglogical_13-2.4.2-1.rhel8.x86_64.rpm) |
| `pglogical_13` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.6 KiB | [pglogical_13-2.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pglogical_13-2.4.1-1.rhel8.x86_64.rpm) |
| `pglogical_13` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.9 KiB | [pglogical_13-2.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pglogical_13-2.4.0-1.rhel8.x86_64.rpm) |
| `pglogical_13` | `2.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 460.2 KiB | [pglogical_13-2.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pglogical_13-2.3.4-1.rhel8.x86_64.rpm) |
| `pglogical_13` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.2 KiB | [pglogical_13-2.4.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pglogical_13-2.4.3-1.rhel8.aarch64.rpm) |
| `pglogical_13` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.4 KiB | [pglogical_13-2.4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pglogical_13-2.4.2-1.rhel8.aarch64.rpm) |
| `pglogical_13` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.3 KiB | [pglogical_13-2.4.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pglogical_13-2.4.3-1.rhel9.x86_64.rpm) |
| `pglogical_13` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.0 KiB | [pglogical_13-2.4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pglogical_13-2.4.2-1.rhel9.x86_64.rpm) |
| `pglogical_13` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.0 KiB | [pglogical_13-2.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pglogical_13-2.4.1-1.rhel9.x86_64.rpm) |
| `pglogical_13` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 145.6 KiB | [pglogical_13-2.4.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pglogical_13-2.4.3-1.rhel9.aarch64.rpm) |
| `pglogical_13` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 145.5 KiB | [pglogical_13-2.4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pglogical_13-2.4.2-1.rhel9.aarch64.rpm) |
| `pglogical_13` | `2.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 151.2 KiB | [pglogical_13-2.4.5-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pglogical_13-2.4.5-3PGDG.rhel10.x86_64.rpm) |
| `pglogical_13` | `2.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.3 KiB | [pglogical_13-2.4.5-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/pglogical_13-2.4.5-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pglogical` | `2.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 345.3 KiB | [postgresql-13-pglogical_2.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-13-pglogical_2.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-13-pglogical` | `2.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 334.0 KiB | [postgresql-13-pglogical_2.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-13-pglogical_2.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-13-pglogical` | `2.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 345.4 KiB | [postgresql-13-pglogical_2.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-13-pglogical_2.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-13-pglogical` | `2.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 335.6 KiB | [postgresql-13-pglogical_2.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-13-pglogical_2.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-13-pglogical` | `2.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 430.1 KiB | [postgresql-13-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-13-pglogical_2.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pglogical` | `2.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 419.2 KiB | [postgresql-13-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-13-pglogical_2.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pglogical` | `2.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 345.1 KiB | [postgresql-13-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-13-pglogical_2.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pglogical` | `2.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 334.9 KiB | [postgresql-13-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pglogical/postgresql-13-pglogical_2.4.6-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/2ndQuadrant/pglogical" title="Repository" icon="github" subtitle="github.com/2ndQuadrant/pglogical" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglogical_ticker-1.4.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pglogical;		# build spec not ready
```


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
pig install pglogical -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pglogical';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pglogical;
```
