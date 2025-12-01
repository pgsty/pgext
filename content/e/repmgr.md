---
title: "repmgr"
linkTitle: "repmgr"
description: "Replication manager for PostgreSQL"
weight: 9710
categories: ["ETL"]
width: full
---

[**repmgr**](https://github.com/EnterpriseDB/repmgr) : Replication manager for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9710** | {{< badge content="repmgr" link="https://github.com/EnterpriseDB/repmgr" >}} | {{< ext "repmgr" >}} | `5.5.0` | {{< category "ETL" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "bgw_replstatus" >}} {{< ext "postgres_fdw" >}} {{< ext "pglogical_origin" >}} {{< ext "pglogical_ticker" >}} {{< ext "dblink" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `repmgr` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.5.0` | {{< bg "18" "repmgr_18*" "red" >}} {{< bg "17" "repmgr_17*" "green" >}} {{< bg "16" "repmgr_16*" "green" >}} {{< bg "15" "repmgr_15*" "green" >}} {{< bg "14" "repmgr_14*" "green" >}} {{< bg "13" "repmgr_13*" "green" >}} | `repmgr_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.5.0` | {{< bg "18" "postgresql-18-repmgr" "green" >}} {{< bg "17" "postgresql-17-repmgr" "green" >}} {{< bg "16" "postgresql-16-repmgr" "green" >}} {{< bg "15" "postgresql-15-repmgr" "green" >}} {{< bg "14" "postgresql-14-repmgr" "green" >}} {{< bg "13" "postgresql-13-repmgr" "green" >}} | `postgresql-$v-repmgr` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "repmgr_18 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.0" "repmgr_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_14 : AVAIL 7" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_13 : AVAIL 8" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "repmgr_18 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.0" "repmgr_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_13 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "repmgr_18 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.0" "repmgr_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "repmgr_18 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.0" "repmgr_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "repmgr_18 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.0" "repmgr_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "repmgr_18 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.0" "repmgr_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "repmgr_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 5.5.0" "postgresql-18-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-17-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-16-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-15-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-14-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-13-repmgr : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 5.5.0" "postgresql-18-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-17-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-16-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-15-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-14-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-13-repmgr : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 5.5.0" "postgresql-18-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-17-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-16-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-15-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-14-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-13-repmgr : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 5.5.0" "postgresql-18-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-17-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-16-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-15-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-14-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-13-repmgr : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 5.5.0" "postgresql-18-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-17-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-16-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-15-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-14-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-13-repmgr : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 5.5.0" "postgresql-18-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-17-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-16-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-15-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-14-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-13-repmgr : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 5.5.0" "postgresql-18-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-17-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-16-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-15-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-14-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-13-repmgr : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 5.5.0" "postgresql-18-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-17-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-16-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-15-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-14-repmgr : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.0" "postgresql-13-repmgr : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-repmgr` | `5.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 246.7 KiB | [postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb) |
| `postgresql-18-repmgr` | `5.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.1 KiB | [postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb) |
| `postgresql-18-repmgr` | `5.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 246.2 KiB | [postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb) |
| `postgresql-18-repmgr` | `5.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 224.6 KiB | [postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb) |
| `postgresql-18-repmgr` | `5.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 235.2 KiB | [postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-repmgr` | `5.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 211.2 KiB | [postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-repmgr` | `5.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 231.3 KiB | [postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-repmgr` | `5.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 208.1 KiB | [postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-18-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `repmgr_17` | `5.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 295.1 KiB | [repmgr_17-5.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/repmgr_17-5.5.0-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_17` | `5.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 285.2 KiB | [repmgr_17-5.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/repmgr_17-5.5.0-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_17` | `5.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 267.9 KiB | [repmgr_17-5.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/repmgr_17-5.5.0-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_17` | `5.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 261.0 KiB | [repmgr_17-5.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/repmgr_17-5.5.0-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_17` | `5.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 270.4 KiB | [repmgr_17-5.5.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/repmgr_17-5.5.0-3PGDG.rhel10.x86_64.rpm) |
| `repmgr_17` | `5.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 261.6 KiB | [repmgr_17-5.5.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/repmgr_17-5.5.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-repmgr` | `5.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 246.6 KiB | [postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb) |
| `postgresql-17-repmgr` | `5.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.5 KiB | [postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb) |
| `postgresql-17-repmgr` | `5.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 246.5 KiB | [postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb) |
| `postgresql-17-repmgr` | `5.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 224.6 KiB | [postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb) |
| `postgresql-17-repmgr` | `5.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 236.2 KiB | [postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-repmgr` | `5.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 212.0 KiB | [postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-repmgr` | `5.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 231.4 KiB | [postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-repmgr` | `5.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 207.5 KiB | [postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-17-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `repmgr_16` | `5.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 293.0 KiB | [repmgr_16-5.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/repmgr_16-5.5.0-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_16` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 291.6 KiB | [repmgr_16-5.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/repmgr_16-5.4.1-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_16` | `5.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 282.6 KiB | [repmgr_16-5.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/repmgr_16-5.5.0-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_16` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 280.9 KiB | [repmgr_16-5.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/repmgr_16-5.4.1-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_16` | `5.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 268.1 KiB | [repmgr_16-5.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/repmgr_16-5.5.0-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_16` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 266.8 KiB | [repmgr_16-5.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/repmgr_16-5.4.1-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_16` | `5.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 261.2 KiB | [repmgr_16-5.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/repmgr_16-5.5.0-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_16` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 261.1 KiB | [repmgr_16-5.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/repmgr_16-5.4.1-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_16` | `5.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 270.4 KiB | [repmgr_16-5.5.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/repmgr_16-5.5.0-3PGDG.rhel10.x86_64.rpm) |
| `repmgr_16` | `5.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 261.6 KiB | [repmgr_16-5.5.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/repmgr_16-5.5.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-repmgr` | `5.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 244.8 KiB | [postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb) |
| `postgresql-16-repmgr` | `5.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 223.3 KiB | [postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb) |
| `postgresql-16-repmgr` | `5.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 244.5 KiB | [postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb) |
| `postgresql-16-repmgr` | `5.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 222.6 KiB | [postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb) |
| `postgresql-16-repmgr` | `5.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 233.7 KiB | [postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-repmgr` | `5.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 209.9 KiB | [postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-repmgr` | `5.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 229.5 KiB | [postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-repmgr` | `5.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 205.4 KiB | [postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-16-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `repmgr_15` | `5.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 291.7 KiB | [repmgr_15-5.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/repmgr_15-5.5.0-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_15` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 290.4 KiB | [repmgr_15-5.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/repmgr_15-5.4.1-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_15` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 289.9 KiB | [repmgr_15-5.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/repmgr_15-5.4.0-1.rhel8.x86_64.rpm) |
| `repmgr_15` | `5.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 286.4 KiB | [repmgr_15-5.3.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/repmgr_15-5.3.3-1.rhel8.x86_64.rpm) |
| `repmgr_15` | `5.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 282.3 KiB | [repmgr_15-5.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/repmgr_15-5.5.0-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_15` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 280.8 KiB | [repmgr_15-5.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/repmgr_15-5.4.1-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_15` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 280.4 KiB | [repmgr_15-5.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/repmgr_15-5.4.0-1.rhel8.aarch64.rpm) |
| `repmgr_15` | `5.3.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 276.5 KiB | [repmgr_15-5.3.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/repmgr_15-5.3.3-1.rhel8.aarch64.rpm) |
| `repmgr_15` | `5.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 270.1 KiB | [repmgr_15-5.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/repmgr_15-5.5.0-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_15` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 268.7 KiB | [repmgr_15-5.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/repmgr_15-5.4.1-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_15` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 268.6 KiB | [repmgr_15-5.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/repmgr_15-5.4.0-1.rhel9.x86_64.rpm) |
| `repmgr_15` | `5.3.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 266.8 KiB | [repmgr_15-5.3.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/repmgr_15-5.3.3-1.rhel9.x86_64.rpm) |
| `repmgr_15` | `5.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 263.0 KiB | [repmgr_15-5.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/repmgr_15-5.5.0-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_15` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 263.0 KiB | [repmgr_15-5.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/repmgr_15-5.4.1-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_15` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 262.8 KiB | [repmgr_15-5.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/repmgr_15-5.4.0-1.rhel9.aarch64.rpm) |
| `repmgr_15` | `5.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 270.8 KiB | [repmgr_15-5.5.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/repmgr_15-5.5.0-3PGDG.rhel10.x86_64.rpm) |
| `repmgr_15` | `5.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 262.3 KiB | [repmgr_15-5.5.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/repmgr_15-5.5.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-repmgr` | `5.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 243.8 KiB | [postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb) |
| `postgresql-15-repmgr` | `5.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 222.1 KiB | [postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb) |
| `postgresql-15-repmgr` | `5.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 242.4 KiB | [postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb) |
| `postgresql-15-repmgr` | `5.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 221.9 KiB | [postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb) |
| `postgresql-15-repmgr` | `5.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 234.8 KiB | [postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-repmgr` | `5.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 211.2 KiB | [postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-repmgr` | `5.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 230.3 KiB | [postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-repmgr` | `5.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 206.1 KiB | [postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-15-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `repmgr_14` | `5.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 289.7 KiB | [repmgr_14-5.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/repmgr_14-5.5.0-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_14` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 288.5 KiB | [repmgr_14-5.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/repmgr_14-5.4.1-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_14` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 288.0 KiB | [repmgr_14-5.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/repmgr_14-5.4.0-1.rhel8.x86_64.rpm) |
| `repmgr_14` | `5.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 284.4 KiB | [repmgr_14-5.3.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/repmgr_14-5.3.3-1.rhel8.x86_64.rpm) |
| `repmgr_14` | `5.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 296.4 KiB | [repmgr_14-5.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/repmgr_14-5.3.2-1.rhel8.x86_64.rpm) |
| `repmgr_14` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 296.6 KiB | [repmgr_14-5.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/repmgr_14-5.3.1-1.rhel8.x86_64.rpm) |
| `repmgr_14` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 296.0 KiB | [repmgr_14-5.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/repmgr_14-5.3.0-1.rhel8.x86_64.rpm) |
| `repmgr_14` | `5.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 280.9 KiB | [repmgr_14-5.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/repmgr_14-5.5.0-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_14` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.2 KiB | [repmgr_14-5.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/repmgr_14-5.4.1-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_14` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.9 KiB | [repmgr_14-5.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/repmgr_14-5.4.0-1.rhel8.aarch64.rpm) |
| `repmgr_14` | `5.3.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 275.3 KiB | [repmgr_14-5.3.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/repmgr_14-5.3.3-1.rhel8.aarch64.rpm) |
| `repmgr_14` | `5.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 269.3 KiB | [repmgr_14-5.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/repmgr_14-5.5.0-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_14` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 268.2 KiB | [repmgr_14-5.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/repmgr_14-5.4.1-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_14` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 267.9 KiB | [repmgr_14-5.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/repmgr_14-5.4.0-1.rhel9.x86_64.rpm) |
| `repmgr_14` | `5.3.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 266.0 KiB | [repmgr_14-5.3.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/repmgr_14-5.3.3-1.rhel9.x86_64.rpm) |
| `repmgr_14` | `5.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 279.2 KiB | [repmgr_14-5.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/repmgr_14-5.3.2-1.rhel9.x86_64.rpm) |
| `repmgr_14` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 278.7 KiB | [repmgr_14-5.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/repmgr_14-5.3.1-1.rhel9.x86_64.rpm) |
| `repmgr_14` | `5.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 262.4 KiB | [repmgr_14-5.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/repmgr_14-5.5.0-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_14` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 262.3 KiB | [repmgr_14-5.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/repmgr_14-5.4.1-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_14` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 262.1 KiB | [repmgr_14-5.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/repmgr_14-5.4.0-1.rhel9.aarch64.rpm) |
| `repmgr_14` | `5.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 269.9 KiB | [repmgr_14-5.5.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/repmgr_14-5.5.0-3PGDG.rhel10.x86_64.rpm) |
| `repmgr_14` | `5.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 261.2 KiB | [repmgr_14-5.5.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/repmgr_14-5.5.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-repmgr` | `5.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 243.0 KiB | [postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb) |
| `postgresql-14-repmgr` | `5.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 222.2 KiB | [postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb) |
| `postgresql-14-repmgr` | `5.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 241.6 KiB | [postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb) |
| `postgresql-14-repmgr` | `5.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 221.1 KiB | [postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb) |
| `postgresql-14-repmgr` | `5.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 233.3 KiB | [postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-repmgr` | `5.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 209.9 KiB | [postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-repmgr` | `5.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 229.6 KiB | [postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-repmgr` | `5.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 205.5 KiB | [postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-14-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `repmgr_13` | `5.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 289.6 KiB | [repmgr_13-5.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/repmgr_13-5.5.0-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_13` | `5.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 288.4 KiB | [repmgr_13-5.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/repmgr_13-5.4.1-1PGDG.rhel8.x86_64.rpm) |
| `repmgr_13` | `5.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 287.8 KiB | [repmgr_13-5.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/repmgr_13-5.4.0-1.rhel8.x86_64.rpm) |
| `repmgr_13` | `5.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 284.3 KiB | [repmgr_13-5.3.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/repmgr_13-5.3.3-1.rhel8.x86_64.rpm) |
| `repmgr_13` | `5.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 296.1 KiB | [repmgr_13-5.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/repmgr_13-5.3.2-1.rhel8.x86_64.rpm) |
| `repmgr_13` | `5.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 296.3 KiB | [repmgr_13-5.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/repmgr_13-5.3.1-1.rhel8.x86_64.rpm) |
| `repmgr_13` | `5.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 295.8 KiB | [repmgr_13-5.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/repmgr_13-5.3.0-1.rhel8.x86_64.rpm) |
| `repmgr_13` | `5.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 292.6 KiB | [repmgr_13-5.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/repmgr_13-5.2.1-1.rhel8.x86_64.rpm) |
| `repmgr_13` | `5.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 280.8 KiB | [repmgr_13-5.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/repmgr_13-5.5.0-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_13` | `5.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 279.1 KiB | [repmgr_13-5.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/repmgr_13-5.4.1-1PGDG.rhel8.aarch64.rpm) |
| `repmgr_13` | `5.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 278.9 KiB | [repmgr_13-5.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/repmgr_13-5.4.0-1.rhel8.aarch64.rpm) |
| `repmgr_13` | `5.3.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 275.1 KiB | [repmgr_13-5.3.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/repmgr_13-5.3.3-1.rhel8.aarch64.rpm) |
| `repmgr_13` | `5.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 269.4 KiB | [repmgr_13-5.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/repmgr_13-5.5.0-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_13` | `5.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 268.2 KiB | [repmgr_13-5.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/repmgr_13-5.4.1-1PGDG.rhel9.x86_64.rpm) |
| `repmgr_13` | `5.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 268.1 KiB | [repmgr_13-5.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/repmgr_13-5.4.0-1.rhel9.x86_64.rpm) |
| `repmgr_13` | `5.3.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 266.2 KiB | [repmgr_13-5.3.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/repmgr_13-5.3.3-1.rhel9.x86_64.rpm) |
| `repmgr_13` | `5.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 279.3 KiB | [repmgr_13-5.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/repmgr_13-5.3.2-1.rhel9.x86_64.rpm) |
| `repmgr_13` | `5.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 278.9 KiB | [repmgr_13-5.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/repmgr_13-5.3.1-1.rhel9.x86_64.rpm) |
| `repmgr_13` | `5.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 262.7 KiB | [repmgr_13-5.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/repmgr_13-5.5.0-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_13` | `5.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 262.6 KiB | [repmgr_13-5.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/repmgr_13-5.4.1-1PGDG.rhel9.aarch64.rpm) |
| `repmgr_13` | `5.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 262.2 KiB | [repmgr_13-5.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/repmgr_13-5.4.0-1.rhel9.aarch64.rpm) |
| `repmgr_13` | `5.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 270.2 KiB | [repmgr_13-5.5.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/repmgr_13-5.5.0-3PGDG.rhel10.x86_64.rpm) |
| `repmgr_13` | `5.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 261.5 KiB | [repmgr_13-5.5.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/repmgr_13-5.5.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-repmgr` | `5.5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 243.4 KiB | [postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg12+1_amd64.deb) |
| `postgresql-13-repmgr` | `5.5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 221.3 KiB | [postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg12+1_arm64.deb) |
| `postgresql-13-repmgr` | `5.5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 241.8 KiB | [postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg13+1_amd64.deb) |
| `postgresql-13-repmgr` | `5.5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 220.6 KiB | [postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg13+1_arm64.deb) |
| `postgresql-13-repmgr` | `5.5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 232.9 KiB | [postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-repmgr` | `5.5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 209.7 KiB | [postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-repmgr` | `5.5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 229.7 KiB | [postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-repmgr` | `5.5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 206.3 KiB | [postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/repmgr/postgresql-13-repmgr_5.5.0+debpgdg-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/repmgr" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/repmgr" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install repmgr;		# install via package name, for the active PG version

pig install repmgr -v 18;   # install for PG 18
pig install repmgr -v 17;   # install for PG 17
pig install repmgr -v 16;   # install for PG 16
pig install repmgr -v 15;   # install for PG 15
pig install repmgr -v 14;   # install for PG 14
pig install repmgr -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION repmgr;
```
