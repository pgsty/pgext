---
title: "mongo_fdw"
linkTitle: "mongo_fdw"
description: "foreign data wrapper for MongoDB access"
weight: 8700
categories: ["FDW"]
width: full
---

[**mongo_fdw**](https://github.com/EnterpriseDB/mongo_fdw) : foreign data wrapper for MongoDB access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8700** | {{< badge content="mongo_fdw" link="https://github.com/EnterpriseDB/mongo_fdw" >}} | {{< ext "mongo_fdw" >}} | `5.5.3` | {{< category "FDW" >}} | {{< license "LGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "redis_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "hdfs_fdw" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} {{< ext "multicorn" >}} {{< ext "jdbc_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.5.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `mongo_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.5.3` | {{< bg "18" "mongo_fdw_18" "green" >}} {{< bg "17" "mongo_fdw_17" "green" >}} {{< bg "16" "mongo_fdw_16" "green" >}} {{< bg "15" "mongo_fdw_15" "green" >}} {{< bg "14" "mongo_fdw_14" "green" >}} | `mongo_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.5.3` | {{< bg "18" "postgresql-18-mongo-fdw" "green" >}} {{< bg "17" "postgresql-17-mongo-fdw" "green" >}} {{< bg "16" "postgresql-16-mongo-fdw" "green" >}} {{< bg "15" "postgresql-15-mongo-fdw" "green" >}} {{< bg "14" "postgresql-14-mongo-fdw" "green" >}} | `postgresql-$v-mongo-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_14 : AVAIL 6" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "mongo_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-18-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-17-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-16-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-15-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-14-mongo-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-18-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-17-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-16-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-15-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-14-mongo-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 5.5.3" "postgresql-18-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-17-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-16-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-15-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-14-mongo-fdw : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 5.5.3" "postgresql-18-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-17-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-16-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-15-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-14-mongo-fdw : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-18-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-17-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-16-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-15-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-14-mongo-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-18-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-17-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-16-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-15-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-14-mongo-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-18-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-17-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-16-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-15-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-14-mongo-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-18-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-17-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-16-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-15-mongo-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.5.3" "postgresql-14-mongo-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 5.5.3" "postgresql-18-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-17-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-16-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-15-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-14-mongo-fdw : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 5.5.3" "postgresql-18-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-17-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-16-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-15-mongo-fdw : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.3" "postgresql-14-mongo-fdw : AVAIL 2" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mongo_fdw_18` | `5.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.8 KiB | [mongo_fdw_18-5.5.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/mongo_fdw_18-5.5.3-2PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 52.4 KiB | [mongo_fdw_18-5.5.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/mongo_fdw_18-5.5.3-2PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.4 KiB | [mongo_fdw_18-5.5.3-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/mongo_fdw_18-5.5.3-3PGDG.rhel9.8.x86_64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.4 KiB | [mongo_fdw_18-5.5.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/mongo_fdw_18-5.5.3-2PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.9 KiB | [mongo_fdw_18-5.5.3-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/mongo_fdw_18-5.5.3-3PGDG.rhel9.8.aarch64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.8 KiB | [mongo_fdw_18-5.5.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/mongo_fdw_18-5.5.3-2PGDG.rhel9.aarch64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.5 KiB | [mongo_fdw_18-5.5.3-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/mongo_fdw_18-5.5.3-3PGDG.rhel10.2.x86_64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.8 KiB | [mongo_fdw_18-5.5.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/mongo_fdw_18-5.5.3-2PGDG.rhel10.x86_64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.0 KiB | [mongo_fdw_18-5.5.3-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/mongo_fdw_18-5.5.3-3PGDG.rhel10.2.aarch64.rpm) |
| `mongo_fdw_18` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.2 KiB | [mongo_fdw_18-5.5.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/mongo_fdw_18-5.5.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 112.8 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 108.9 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 113.3 KiB | [postgresql-18-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 113.2 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 109.6 KiB | [postgresql-18-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.2 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 123.9 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 121.2 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 117.6 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 115.4 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 112.1 KiB | [postgresql-18-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 117.3 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 108.4 KiB | [postgresql-18-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 115.1 KiB | [postgresql-18-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-18-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mongo_fdw_17` | `5.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.7 KiB | [mongo_fdw_17-5.5.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/mongo_fdw_17-5.5.3-2PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 52.2 KiB | [mongo_fdw_17-5.5.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/mongo_fdw_17-5.5.3-2PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.2 KiB | [mongo_fdw_17-5.5.3-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/mongo_fdw_17-5.5.3-3PGDG.rhel9.8.x86_64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.3 KiB | [mongo_fdw_17-5.5.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/mongo_fdw_17-5.5.3-2PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.7 KiB | [mongo_fdw_17-5.5.3-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/mongo_fdw_17-5.5.3-3PGDG.rhel9.8.aarch64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.6 KiB | [mongo_fdw_17-5.5.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/mongo_fdw_17-5.5.3-2PGDG.rhel9.aarch64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.4 KiB | [mongo_fdw_17-5.5.3-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/mongo_fdw_17-5.5.3-3PGDG.rhel10.2.x86_64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.7 KiB | [mongo_fdw_17-5.5.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/mongo_fdw_17-5.5.3-2PGDG.rhel10.x86_64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.8 KiB | [mongo_fdw_17-5.5.3-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/mongo_fdw_17-5.5.3-3PGDG.rhel10.2.aarch64.rpm) |
| `mongo_fdw_17` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.1 KiB | [mongo_fdw_17-5.5.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/mongo_fdw_17-5.5.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 112.5 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 108.7 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 113.0 KiB | [postgresql-17-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 112.8 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 109.3 KiB | [postgresql-17-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.0 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 144.0 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 141.1 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 117.4 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 115.1 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 111.6 KiB | [postgresql-17-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 117.0 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 108.0 KiB | [postgresql-17-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 114.6 KiB | [postgresql-17-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-17-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mongo_fdw_16` | `5.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 54.7 KiB | [mongo_fdw_16-5.5.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/mongo_fdw_16-5.5.3-2PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_16` | `5.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.2 KiB | [mongo_fdw_16-5.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/mongo_fdw_16-5.5.1-1PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 52.2 KiB | [mongo_fdw_16-5.5.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/mongo_fdw_16-5.5.3-2PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_16` | `5.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.7 KiB | [mongo_fdw_16-5.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/mongo_fdw_16-5.5.1-1PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.2 KiB | [mongo_fdw_16-5.5.3-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/mongo_fdw_16-5.5.3-3PGDG.rhel9.8.x86_64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.3 KiB | [mongo_fdw_16-5.5.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/mongo_fdw_16-5.5.3-2PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_16` | `5.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.2 KiB | [mongo_fdw_16-5.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/mongo_fdw_16-5.5.1-1PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.7 KiB | [mongo_fdw_16-5.5.3-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/mongo_fdw_16-5.5.3-3PGDG.rhel9.8.aarch64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.6 KiB | [mongo_fdw_16-5.5.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/mongo_fdw_16-5.5.3-2PGDG.rhel9.aarch64.rpm) |
| `mongo_fdw_16` | `5.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 63.2 KiB | [mongo_fdw_16-5.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/mongo_fdw_16-5.5.1-1PGDG.rhel9.aarch64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.3 KiB | [mongo_fdw_16-5.5.3-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/mongo_fdw_16-5.5.3-3PGDG.rhel10.2.x86_64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 53.6 KiB | [mongo_fdw_16-5.5.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/mongo_fdw_16-5.5.3-2PGDG.rhel10.x86_64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.3 KiB | [mongo_fdw_16-5.5.3-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/mongo_fdw_16-5.5.3-3PGDG.rhel10.2.aarch64.rpm) |
| `mongo_fdw_16` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 52.1 KiB | [mongo_fdw_16-5.5.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/mongo_fdw_16-5.5.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 112.4 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 108.6 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 112.9 KiB | [postgresql-16-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 112.8 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 109.3 KiB | [postgresql-16-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.1 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 142.5 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 139.6 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 117.2 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 115.1 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 111.6 KiB | [postgresql-16-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 116.9 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 107.8 KiB | [postgresql-16-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 114.8 KiB | [postgresql-16-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-16-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mongo_fdw_15` | `5.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.1 KiB | [mongo_fdw_15-5.5.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mongo_fdw_15-5.5.3-2PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.5 KiB | [mongo_fdw_15-5.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mongo_fdw_15-5.5.1-1PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.5 KiB | [mongo_fdw_15-5.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mongo_fdw_15-5.5.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_15` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.3 KiB | [mongo_fdw_15-5.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mongo_fdw_15-5.4.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.5 KiB | [mongo_fdw_15-5.5.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/mongo_fdw_15-5.5.3-2PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_15` | `5.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.8 KiB | [mongo_fdw_15-5.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/mongo_fdw_15-5.5.1-1PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 55.4 KiB | [mongo_fdw_15-5.5.3-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mongo_fdw_15-5.5.3-3PGDG.rhel9.8.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 55.4 KiB | [mongo_fdw_15-5.5.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mongo_fdw_15-5.5.3-2PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 79.2 KiB | [mongo_fdw_15-5.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mongo_fdw_15-5.5.1-1PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.9 KiB | [mongo_fdw_15-5.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mongo_fdw_15-5.5.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_15` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.6 KiB | [mongo_fdw_15-5.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mongo_fdw_15-5.4.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.4 KiB | [mongo_fdw_15-5.5.3-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/mongo_fdw_15-5.5.3-3PGDG.rhel9.8.aarch64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.2 KiB | [mongo_fdw_15-5.5.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/mongo_fdw_15-5.5.3-2PGDG.rhel9.aarch64.rpm) |
| `mongo_fdw_15` | `5.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 75.7 KiB | [mongo_fdw_15-5.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/mongo_fdw_15-5.5.1-1PGDG.rhel9.aarch64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 56.3 KiB | [mongo_fdw_15-5.5.3-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/mongo_fdw_15-5.5.3-3PGDG.rhel10.2.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 56.6 KiB | [mongo_fdw_15-5.5.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/mongo_fdw_15-5.5.3-2PGDG.rhel10.x86_64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.2 KiB | [mongo_fdw_15-5.5.3-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/mongo_fdw_15-5.5.3-3PGDG.rhel10.2.aarch64.rpm) |
| `mongo_fdw_15` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.4 KiB | [mongo_fdw_15-5.5.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/mongo_fdw_15-5.5.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 113.4 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 109.4 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 113.8 KiB | [postgresql-15-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 113.7 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 110.2 KiB | [postgresql-15-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 110.1 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 144.7 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 141.8 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 118.9 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 116.7 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 113.0 KiB | [postgresql-15-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 118.6 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 109.3 KiB | [postgresql-15-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 116.5 KiB | [postgresql-15-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-15-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mongo_fdw_14` | `5.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.1 KiB | [mongo_fdw_14-5.5.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.5.3-2PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.5 KiB | [mongo_fdw_14-5.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.5.1-1PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.4 KiB | [mongo_fdw_14-5.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.5.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.3 KiB | [mongo_fdw_14-5.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.4.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.3 KiB | [mongo_fdw_14-5.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.3.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | `5.2.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.7 KiB | [mongo_fdw_14-5.2.10-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.2.10-2.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.5 KiB | [mongo_fdw_14-5.5.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/mongo_fdw_14-5.5.3-2PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_14` | `5.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.8 KiB | [mongo_fdw_14-5.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/mongo_fdw_14-5.5.1-1PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 55.3 KiB | [mongo_fdw_14-5.5.3-3PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.5.3-3PGDG.rhel9.8.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 55.4 KiB | [mongo_fdw_14-5.5.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.5.3-2PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 79.0 KiB | [mongo_fdw_14-5.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.5.1-1PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.9 KiB | [mongo_fdw_14-5.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.5.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.9 KiB | [mongo_fdw_14-5.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.4.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | `5.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.9 KiB | [mongo_fdw_14-5.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.3.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.4 KiB | [mongo_fdw_14-5.5.3-3PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/mongo_fdw_14-5.5.3-3PGDG.rhel9.8.aarch64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 53.3 KiB | [mongo_fdw_14-5.5.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/mongo_fdw_14-5.5.3-2PGDG.rhel9.aarch64.rpm) |
| `mongo_fdw_14` | `5.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 75.6 KiB | [mongo_fdw_14-5.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/mongo_fdw_14-5.5.1-1PGDG.rhel9.aarch64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 56.3 KiB | [mongo_fdw_14-5.5.3-3PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/mongo_fdw_14-5.5.3-3PGDG.rhel10.2.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 56.6 KiB | [mongo_fdw_14-5.5.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/mongo_fdw_14-5.5.3-2PGDG.rhel10.x86_64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.2 KiB | [mongo_fdw_14-5.5.3-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/mongo_fdw_14-5.5.3-3PGDG.rhel10.2.aarch64.rpm) |
| `mongo_fdw_14` | `5.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 54.4 KiB | [mongo_fdw_14-5.5.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/mongo_fdw_14-5.5.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 113.4 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 109.4 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 113.8 KiB | [postgresql-14-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 114.0 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 110.1 KiB | [postgresql-14-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.9 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 144.5 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 141.6 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 118.9 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 116.8 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 112.9 KiB | [postgresql-14-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 118.6 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 109.3 KiB | [postgresql-14-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-mongo-fdw` | `5.5.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 116.5 KiB | [postgresql-14-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/mongo-fdw/postgresql-14-mongo-fdw_5.5.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/mongo_fdw" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/mongo_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="mongo_fdw-REL-5_5_3.tar.gz" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install mongo_fdw;		# install via package name, for the active PG version

pig install mongo_fdw -v 18;   # install for PG 18
pig install mongo_fdw -v 17;   # install for PG 17
pig install mongo_fdw -v 16;   # install for PG 16
pig install mongo_fdw -v 15;   # install for PG 15
pig install mongo_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION mongo_fdw;
```




## Usage

Sources: [README](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_5_3/README.md), [REL-5_5_3 release](https://github.com/EnterpriseDB/mongo_fdw/releases/tag/REL-5_5_3)

### Create Server

```sql
CREATE EXTENSION mongo_fdw;

CREATE SERVER mongo_server FOREIGN DATA WRAPPER mongo_fdw
  OPTIONS (address '127.0.0.1', port '27017');
```

**Server Options:** `address` (default `127.0.0.1`), `port` (default `27017`), `authentication_database`, `replica_set`, `read_preference` (`primary`, `secondary`, `primaryPreferred`, `secondaryPreferred`, `nearest`), `ssl` (default `false`), `pem_file`, `pem_pwd`, `ca_file`, `ca_dir`, `crl_file`, `weak_cert_validation`, `use_remote_estimate` (default `false`), `enable_join_pushdown` (default `true`), `enable_aggregate_pushdown` (default `true`), `enable_order_by_pushdown` (default `true`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR pguser SERVER mongo_server
  OPTIONS (username 'mongouser', password 'mongopass');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE warehouse (
  _id name,
  warehouse_id int,
  warehouse_name text,
  warehouse_created timestamptz
)
SERVER mongo_server
OPTIONS (database 'mydb', collection 'warehouse');
```

**Important:** The first column must be `_id` of type `name` (MongoDB's object identifier).

**Table Options:** `database` (default `test`), `collection` (defaults to table name), `enable_join_pushdown`, `enable_aggregate_pushdown`, `enable_order_by_pushdown`.

### CRUD Operations

```sql
SELECT warehouse_id, warehouse_name FROM warehouse WHERE warehouse_id > 10;
INSERT INTO warehouse VALUES ('new_id', 100, 'New Warehouse', now());
UPDATE warehouse SET warehouse_name = 'Updated' WHERE warehouse_id = 100;
DELETE FROM warehouse WHERE warehouse_id = 100;
```

### Pushdown Features

mongo_fdw pushes down WHERE clauses, joins between foreign tables on the same server, aggregate functions, ORDER BY, LIMIT, and OFFSET to MongoDB for efficient query execution. Use the `mongo_fdw.enable_join_pushdown`, `mongo_fdw.enable_aggregate_pushdown`, `mongo_fdw.enable_order_by_pushdown`, and `mongo_fdw.log_remote_query` GUCs when diagnosing remote plans.

### Version Notes

`mongo_fdw` 5.5.3, released upstream as `REL-5_5_3`, adds PostgreSQL 18 support, updates bundled `mongoc-driver` and `json-c` libraries for MongoDB 8, adds the `mongo_fdw.log_remote_query` debug GUC, and fixes nested-field, WHERE, ORDER BY, and unsafe join-pushdown cases. Upstream dropped PostgreSQL 12 support in this line.

### Notes

- BSON only supports UTF-8; ensure PostgreSQL database uses UTF-8 encoding
- Column names with uppercase letters or dots (for nested documents) require double-quoting
- PostgreSQL limits column names to 63 characters by default
