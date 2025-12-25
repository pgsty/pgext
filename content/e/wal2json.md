---
title: "wal2json"
linkTitle: "wal2json"
description: "Changing data capture in JSON format"
weight: 9630
categories: ["ETL"]
width: full
---

[**wal2json**](https://github.com/eulerto/wal2json) : Changing data capture in JSON format


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9630** | {{< badge content="wal2json" link="https://github.com/eulerto/wal2json" >}} | {{< ext "wal2json" >}} | `2.6` | {{< category "ETL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "wal2mongo" >}} {{< ext "decoderbufs" >}} {{< ext "decoder_raw" >}} {{< ext "kafka_fdw" >}} {{< ext "pglogical_origin" >}} {{< ext "pglogical_ticker" >}} {{< ext "pg_failover_slots" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `wal2json` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.6` | {{< bg "18" "wal2json_18" "green" >}} {{< bg "17" "wal2json_17" "green" >}} {{< bg "16" "wal2json_16" "green" >}} {{< bg "15" "wal2json_15" "green" >}} {{< bg "14" "wal2json_14" "green" >}} {{< bg "13" "wal2json_13" "green" >}} | `wal2json_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.6` | {{< bg "18" "postgresql-18-wal2json" "green" >}} {{< bg "17" "postgresql-17-wal2json" "green" >}} {{< bg "16" "postgresql-16-wal2json" "green" >}} {{< bg "15" "postgresql-15-wal2json" "green" >}} {{< bg "14" "postgresql-14-wal2json" "green" >}} {{< bg "13" "postgresql-13-wal2json" "green" >}} | `postgresql-$v-wal2json` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.6" "wal2json_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.6" "wal2json_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.6" "wal2json_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.5" "wal2json_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.5" "wal2json_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.5" "wal2json_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.5" "wal2json_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.6" "wal2json_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.6" "wal2json_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.6" "wal2json_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "wal2json_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.6" "postgresql-18-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-17-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-16-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-15-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-14-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-13-wal2json : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.6" "postgresql-18-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-17-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-16-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-15-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-14-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-13-wal2json : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.6" "postgresql-18-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-17-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-16-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-15-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-14-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-13-wal2json : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.6" "postgresql-18-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-17-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-16-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-15-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-14-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-13-wal2json : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.6" "postgresql-18-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-17-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-16-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-15-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-14-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-13-wal2json : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.6" "postgresql-18-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-17-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-16-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-15-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-14-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-13-wal2json : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.6" "postgresql-18-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-17-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-16-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-15-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-14-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-13-wal2json : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.6" "postgresql-18-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-17-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-16-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-15-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-14-wal2json : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.6" "postgresql-13-wal2json : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2json_18` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.3 KiB | [wal2json_18-2.6-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/wal2json_18-2.6-3PGDG.rhel8.x86_64.rpm) |
| `wal2json_18` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.4 KiB | [wal2json_18-2.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wal2json_18-2.6-1PIGSTY.el8.x86_64.rpm) |
| `wal2json_18` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.4 KiB | [wal2json_18-2.6-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/wal2json_18-2.6-3PGDG.rhel8.aarch64.rpm) |
| `wal2json_18` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.5 KiB | [wal2json_18-2.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wal2json_18-2.6-1PIGSTY.el8.aarch64.rpm) |
| `wal2json_18` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.1 KiB | [wal2json_18-2.6-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/wal2json_18-2.6-3PGDG.rhel9.x86_64.rpm) |
| `wal2json_18` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [wal2json_18-2.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wal2json_18-2.6-1PIGSTY.el9.x86_64.rpm) |
| `wal2json_18` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.4 KiB | [wal2json_18-2.6-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/wal2json_18-2.6-3PGDG.rhel9.aarch64.rpm) |
| `wal2json_18` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.1 KiB | [wal2json_18-2.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wal2json_18-2.6-1PIGSTY.el9.aarch64.rpm) |
| `wal2json_18` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.9 KiB | [wal2json_18-2.6-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/wal2json_18-2.6-3PGDG.rhel10.x86_64.rpm) |
| `wal2json_18` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [wal2json_18-2.6-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/wal2json_18-2.6-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-wal2json` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 56.2 KiB | [postgresql-18-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-18-wal2json` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 53.9 KiB | [postgresql-18-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-18-wal2json` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 55.9 KiB | [postgresql-18-wal2json_2.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg13+1_amd64.deb) |
| `postgresql-18-wal2json` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 54.1 KiB | [postgresql-18-wal2json_2.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg13+1_arm64.deb) |
| `postgresql-18-wal2json` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 57.6 KiB | [postgresql-18-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-wal2json` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 54.9 KiB | [postgresql-18-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-wal2json` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 56.1 KiB | [postgresql-18-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-wal2json` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 53.9 KiB | [postgresql-18-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2json_17` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.3 KiB | [wal2json_17-2.6-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/wal2json_17-2.6-2PGDG.rhel8.x86_64.rpm) |
| `wal2json_17` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [wal2json_17-2.6-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/wal2json_17-2.6-2PGDG.rhel8.aarch64.rpm) |
| `wal2json_17` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [wal2json_17-2.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/wal2json_17-2.6-2PGDG.rhel9.x86_64.rpm) |
| `wal2json_17` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.9 KiB | [wal2json_17-2.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/wal2json_17-2.6-2PGDG.rhel9.aarch64.rpm) |
| `wal2json_17` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [wal2json_17-2.6-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/wal2json_17-2.6-3PGDG.rhel10.x86_64.rpm) |
| `wal2json_17` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [wal2json_17-2.6-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/wal2json_17-2.6-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-wal2json` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 56.0 KiB | [postgresql-17-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-17-wal2json` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 53.8 KiB | [postgresql-17-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-17-wal2json` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 55.9 KiB | [postgresql-17-wal2json_2.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg13+1_amd64.deb) |
| `postgresql-17-wal2json` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 54.0 KiB | [postgresql-17-wal2json_2.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg13+1_arm64.deb) |
| `postgresql-17-wal2json` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 63.9 KiB | [postgresql-17-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-wal2json` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 61.4 KiB | [postgresql-17-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-wal2json` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 55.9 KiB | [postgresql-17-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-wal2json` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 53.8 KiB | [postgresql-17-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2json_16` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.1 KiB | [wal2json_16-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/wal2json_16-2.6-1PGDG.rhel8.x86_64.rpm) |
| `wal2json_16` | `2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.7 KiB | [wal2json_16-2.5-3.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/wal2json_16-2.5-3.rhel8.1.x86_64.rpm) |
| `wal2json_16` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.5 KiB | [wal2json_16-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/wal2json_16-2.6-1PGDG.rhel8.aarch64.rpm) |
| `wal2json_16` | `2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.0 KiB | [wal2json_16-2.5-3.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/wal2json_16-2.5-3.rhel8.1.aarch64.rpm) |
| `wal2json_16` | `2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.8 KiB | [wal2json_16-2.5-3.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/wal2json_16-2.5-3.rhel9.1.x86_64.rpm) |
| `wal2json_16` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [wal2json_16-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/wal2json_16-2.6-1PGDG.rhel9.aarch64.rpm) |
| `wal2json_16` | `2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [wal2json_16-2.5-3.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/wal2json_16-2.5-3.rhel9.1.aarch64.rpm) |
| `wal2json_16` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [wal2json_16-2.6-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/wal2json_16-2.6-3PGDG.rhel10.x86_64.rpm) |
| `wal2json_16` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.3 KiB | [wal2json_16-2.6-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/wal2json_16-2.6-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-wal2json` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 56.0 KiB | [postgresql-16-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-16-wal2json` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 53.8 KiB | [postgresql-16-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-16-wal2json` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 55.9 KiB | [postgresql-16-wal2json_2.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg13+1_amd64.deb) |
| `postgresql-16-wal2json` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 54.0 KiB | [postgresql-16-wal2json_2.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg13+1_arm64.deb) |
| `postgresql-16-wal2json` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 63.6 KiB | [postgresql-16-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-wal2json` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 61.1 KiB | [postgresql-16-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-wal2json` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 56.0 KiB | [postgresql-16-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-wal2json` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 53.8 KiB | [postgresql-16-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2json_15` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.2 KiB | [wal2json_15-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/wal2json_15-2.6-1PGDG.rhel8.x86_64.rpm) |
| `wal2json_15` | `2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.3 KiB | [wal2json_15-2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/wal2json_15-2.5-1.rhel8.x86_64.rpm) |
| `wal2json_15` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [wal2json_15-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2json_15-2.6-1PGDG.rhel8.aarch64.rpm) |
| `wal2json_15` | `2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.6 KiB | [wal2json_15-2.5-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2json_15-2.5-2.rhel8.aarch64.rpm) |
| `wal2json_15` | `2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.6 KiB | [wal2json_15-2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2json_15-2.5-1.rhel8.aarch64.rpm) |
| `wal2json_15` | `2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.1 KiB | [wal2json_15-2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/wal2json_15-2.5-1.rhel9.x86_64.rpm) |
| `wal2json_15` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [wal2json_15-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2json_15-2.6-1PGDG.rhel9.aarch64.rpm) |
| `wal2json_15` | `2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.6 KiB | [wal2json_15-2.5-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2json_15-2.5-2.rhel9.aarch64.rpm) |
| `wal2json_15` | `2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.5 KiB | [wal2json_15-2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2json_15-2.5-1.rhel9.aarch64.rpm) |
| `wal2json_15` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [wal2json_15-2.6-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/wal2json_15-2.6-3PGDG.rhel10.x86_64.rpm) |
| `wal2json_15` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.4 KiB | [wal2json_15-2.6-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/wal2json_15-2.6-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-wal2json` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 56.6 KiB | [postgresql-15-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-15-wal2json` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 54.1 KiB | [postgresql-15-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-15-wal2json` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 56.4 KiB | [postgresql-15-wal2json_2.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg13+1_amd64.deb) |
| `postgresql-15-wal2json` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 54.2 KiB | [postgresql-15-wal2json_2.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg13+1_arm64.deb) |
| `postgresql-15-wal2json` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.2 KiB | [postgresql-15-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-wal2json` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 61.5 KiB | [postgresql-15-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-wal2json` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 56.6 KiB | [postgresql-15-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-wal2json` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 54.1 KiB | [postgresql-15-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2json_14` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.2 KiB | [wal2json_14-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2json_14-2.6-1PGDG.rhel8.x86_64.rpm) |
| `wal2json_14` | `2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.4 KiB | [wal2json_14-2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2json_14-2.5-1.rhel8.x86_64.rpm) |
| `wal2json_14` | `2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.4 KiB | [wal2json_14-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2json_14-2.4-1.rhel8.x86_64.rpm) |
| `wal2json_14` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.2 KiB | [wal2json_14-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2json_14-2.6-1PGDG.rhel8.aarch64.rpm) |
| `wal2json_14` | `2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.6 KiB | [wal2json_14-2.5-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2json_14-2.5-2.rhel8.aarch64.rpm) |
| `wal2json_14` | `2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.5 KiB | [wal2json_14-2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2json_14-2.5-1.rhel8.aarch64.rpm) |
| `wal2json_14` | `2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.1 KiB | [wal2json_14-2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/wal2json_14-2.5-1.rhel9.x86_64.rpm) |
| `wal2json_14` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [wal2json_14-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2json_14-2.6-1PGDG.rhel9.aarch64.rpm) |
| `wal2json_14` | `2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.6 KiB | [wal2json_14-2.5-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2json_14-2.5-2.rhel9.aarch64.rpm) |
| `wal2json_14` | `2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.6 KiB | [wal2json_14-2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2json_14-2.5-1.rhel9.aarch64.rpm) |
| `wal2json_14` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.9 KiB | [wal2json_14-2.6-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/wal2json_14-2.6-3PGDG.rhel10.x86_64.rpm) |
| `wal2json_14` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.5 KiB | [wal2json_14-2.6-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/wal2json_14-2.6-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-wal2json` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 56.2 KiB | [postgresql-14-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-14-wal2json` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 53.8 KiB | [postgresql-14-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-14-wal2json` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 56.1 KiB | [postgresql-14-wal2json_2.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg13+1_amd64.deb) |
| `postgresql-14-wal2json` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 53.9 KiB | [postgresql-14-wal2json_2.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg13+1_arm64.deb) |
| `postgresql-14-wal2json` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 64.3 KiB | [postgresql-14-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-wal2json` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 61.5 KiB | [postgresql-14-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-wal2json` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 56.2 KiB | [postgresql-14-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-wal2json` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 53.7 KiB | [postgresql-14-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2json_13` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.1 KiB | [wal2json_13-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/wal2json_13-2.6-1PGDG.rhel8.x86_64.rpm) |
| `wal2json_13` | `2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.2 KiB | [wal2json_13-2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/wal2json_13-2.5-1.rhel8.x86_64.rpm) |
| `wal2json_13` | `2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.5 KiB | [wal2json_13-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/wal2json_13-2.4-1.rhel8.x86_64.rpm) |
| `wal2json_13` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [wal2json_13-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/wal2json_13-2.6-1PGDG.rhel8.aarch64.rpm) |
| `wal2json_13` | `2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.6 KiB | [wal2json_13-2.5-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/wal2json_13-2.5-2.rhel8.aarch64.rpm) |
| `wal2json_13` | `2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.5 KiB | [wal2json_13-2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/wal2json_13-2.5-1.rhel8.aarch64.rpm) |
| `wal2json_13` | `2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.1 KiB | [wal2json_13-2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/wal2json_13-2.5-1.rhel9.x86_64.rpm) |
| `wal2json_13` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.7 KiB | [wal2json_13-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/wal2json_13-2.6-1PGDG.rhel9.aarch64.rpm) |
| `wal2json_13` | `2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.6 KiB | [wal2json_13-2.5-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/wal2json_13-2.5-2.rhel9.aarch64.rpm) |
| `wal2json_13` | `2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.5 KiB | [wal2json_13-2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/wal2json_13-2.5-1.rhel9.aarch64.rpm) |
| `wal2json_13` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.0 KiB | [wal2json_13-2.6-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/wal2json_13-2.6-3PGDG.rhel10.x86_64.rpm) |
| `wal2json_13` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.4 KiB | [wal2json_13-2.6-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/wal2json_13-2.6-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-wal2json` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 55.9 KiB | [postgresql-13-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-13-wal2json` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 53.4 KiB | [postgresql-13-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-13-wal2json` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 55.8 KiB | [postgresql-13-wal2json_2.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg13+1_amd64.deb) |
| `postgresql-13-wal2json` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 53.6 KiB | [postgresql-13-wal2json_2.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg13+1_arm64.deb) |
| `postgresql-13-wal2json` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 63.1 KiB | [postgresql-13-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-wal2json` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 60.3 KiB | [postgresql-13-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-wal2json` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 56.1 KiB | [postgresql-13-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-wal2json` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 53.3 KiB | [postgresql-13-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/eulerto/wal2json" title="Repository" icon="github" subtitle="github.com/eulerto/wal2json" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="wal2json-2.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg wal2json;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install wal2json;		# install via package name, for the active PG version

pig install wal2json -v 18;   # install for PG 18
pig install wal2json -v 17;   # install for PG 17
pig install wal2json -v 16;   # install for PG 16
pig install wal2json -v 15;   # install for PG 15
pig install wal2json -v 14;   # install for PG 14
pig install wal2json -v 13;   # install for PG 13

```


This extension does not need `CREATE EXTENSION` DDL command


