---
title: "wal2json"
linkTitle: "wal2json"
description: "Changing data capture in JSON format"
weight: 9630
categories: ["Etl"]
width: full
---

Changing data capture in JSON format

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9630** | {{< badge content="wal2json" link="https://github.com/eulerto/wal2json" >}} | {{< ext "wal2json" "wal2json" >}} | `2.6` | {{< category "ETL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s----" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "wal2mongo" >}} {{< ext "decoderbufs" >}} {{< ext "decoder_raw" >}} {{< ext "kafka_fdw" >}} {{< ext "pglogical_origin" >}} {{< ext "pglogical_ticker" >}} {{< ext "pg_failover_slots" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/wal2json" >}} | `2.6` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `wal2json_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/wal2json" >}} | `2.6` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-wal2json` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "wal2json_18" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/wal2json_18-2.6-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "wal2json_17" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/wal2json_17-2.6-2PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "wal2json_16" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/wal2json_16-2.6-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "wal2json_15" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/wal2json_15-2.6-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "wal2json_14" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2json_14-2.6-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "wal2json_18" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/wal2json_18-2.6-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "wal2json_17" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/wal2json_17-2.6-2PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "wal2json_16" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/wal2json_16-2.6-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "wal2json_15" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2json_15-2.6-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "wal2json_14" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2json_14-2.6-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "wal2json_18" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/wal2json_18-2.6-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "wal2json_17" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/wal2json_17-2.6-2PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "wal2json_16" "2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/wal2json_16-2.5-3.rhel9.1.x86_64.rpm" >}} | {{< pkg "wal2json_15" "2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/wal2json_15-2.5-1.rhel9.x86_64.rpm" >}} | {{< pkg "wal2json_14" "2.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/wal2json_14-2.5-1.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "wal2json_18" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/wal2json_18-2.6-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "wal2json_17" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/wal2json_17-2.6-2PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "wal2json_16" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/wal2json_16-2.6-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "wal2json_15" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2json_15-2.6-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "wal2json_14" "2.6" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2json_14-2.6-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-wal2json" "2.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2json_18` | 2.6 | `el8.aarch64` | pgdg | 31.4 KiB | [wal2json_18-2.6-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/wal2json_18-2.6-3PGDG.rhel8.aarch64.rpm) |
| `wal2json_18` | 2.6 | `el8.x86_64` | pgdg | 33.3 KiB | [wal2json_18-2.6-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/wal2json_18-2.6-3PGDG.rhel8.x86_64.rpm) |
| `wal2json_18` | 2.6 | `el8.aarch64` | pigsty | 29.5 KiB | [wal2json_18-2.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wal2json_18-2.6-1PIGSTY.el8.aarch64.rpm) |
| `wal2json_18` | 2.6 | `el8.x86_64` | pigsty | 31.4 KiB | [wal2json_18-2.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wal2json_18-2.6-1PIGSTY.el8.x86_64.rpm) |
| `wal2json_18` | 2.6 | `el9.x86_64` | pgdg | 32.1 KiB | [wal2json_18-2.6-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/wal2json_18-2.6-3PGDG.rhel9.x86_64.rpm) |
| `wal2json_18` | 2.6 | `el9.x86_64` | pigsty | 31.8 KiB | [wal2json_18-2.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wal2json_18-2.6-1PIGSTY.el9.x86_64.rpm) |
| `wal2json_18` | 2.6 | `el9.aarch64` | pgdg | 30.4 KiB | [wal2json_18-2.6-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/wal2json_18-2.6-3PGDG.rhel9.aarch64.rpm) |
| `wal2json_18` | 2.6 | `el9.aarch64` | pigsty | 30.1 KiB | [wal2json_18-2.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wal2json_18-2.6-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-18-wal2json` | 2.6 | `d12.x86_64` | pgdg | 56.2 KiB | [postgresql-18-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-18-wal2json` | 2.6 | `d12.aarch64` | pgdg | 53.9 KiB | [postgresql-18-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-18-wal2json` | 2.6 | `u22.x86_64` | pgdg | 57.6 KiB | [postgresql-18-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-wal2json` | 2.6 | `u22.aarch64` | pgdg | 54.9 KiB | [postgresql-18-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-wal2json` | 2.6 | `u24.x86_64` | pgdg | 56.1 KiB | [postgresql-18-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-wal2json` | 2.6 | `u24.aarch64` | pgdg | 53.9 KiB | [postgresql-18-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-18-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2json_17` | 2.6 | `el8.aarch64` | pgdg | 31.5 KiB | [wal2json_17-2.6-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/wal2json_17-2.6-2PGDG.rhel8.aarch64.rpm) |
| `wal2json_17` | 2.6 | `el8.x86_64` | pgdg | 33.3 KiB | [wal2json_17-2.6-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/wal2json_17-2.6-2PGDG.rhel8.x86_64.rpm) |
| `wal2json_17` | 2.6 | `el9.aarch64` | pgdg | 30.9 KiB | [wal2json_17-2.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/wal2json_17-2.6-2PGDG.rhel9.aarch64.rpm) |
| `wal2json_17` | 2.6 | `el9.x86_64` | pgdg | 32.5 KiB | [wal2json_17-2.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/wal2json_17-2.6-2PGDG.rhel9.x86_64.rpm) |
| `postgresql-17-wal2json` | 2.6 | `d12.aarch64` | pgdg | 53.8 KiB | [postgresql-17-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-17-wal2json` | 2.6 | `d12.x86_64` | pgdg | 56.0 KiB | [postgresql-17-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-17-wal2json` | 2.6 | `u22.x86_64` | pgdg | 63.9 KiB | [postgresql-17-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-wal2json` | 2.6 | `u22.aarch64` | pgdg | 61.4 KiB | [postgresql-17-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-wal2json` | 2.6 | `u24.aarch64` | pgdg | 53.8 KiB | [postgresql-17-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-wal2json` | 2.6 | `u24.x86_64` | pgdg | 55.9 KiB | [postgresql-17-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-17-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2json_16` | 2.6 | `el8.aarch64` | pgdg | 31.5 KiB | [wal2json_16-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/wal2json_16-2.6-1PGDG.rhel8.aarch64.rpm) |
| `wal2json_16` | 2.6 | `el8.x86_64` | pgdg | 33.1 KiB | [wal2json_16-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/wal2json_16-2.6-1PGDG.rhel8.x86_64.rpm) |
| `wal2json_16` | 2.5 | `el8.x86_64` | pgdg | 32.7 KiB | [wal2json_16-2.5-3.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/wal2json_16-2.5-3.rhel8.1.x86_64.rpm) |
| `wal2json_16` | 2.5 | `el8.aarch64` | pgdg | 31.0 KiB | [wal2json_16-2.5-3.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/wal2json_16-2.5-3.rhel8.1.aarch64.rpm) |
| `wal2json_16` | 2.6 | `el9.aarch64` | pgdg | 30.8 KiB | [wal2json_16-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/wal2json_16-2.6-1PGDG.rhel9.aarch64.rpm) |
| `wal2json_16` | 2.5 | `el9.x86_64` | pgdg | 31.8 KiB | [wal2json_16-2.5-3.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/wal2json_16-2.5-3.rhel9.1.x86_64.rpm) |
| `wal2json_16` | 2.5 | `el9.aarch64` | pgdg | 30.3 KiB | [wal2json_16-2.5-3.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/wal2json_16-2.5-3.rhel9.1.aarch64.rpm) |
| `postgresql-16-wal2json` | 2.6 | `d12.aarch64` | pgdg | 53.8 KiB | [postgresql-16-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-16-wal2json` | 2.6 | `d12.x86_64` | pgdg | 56.0 KiB | [postgresql-16-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-16-wal2json` | 2.6 | `u22.aarch64` | pgdg | 61.1 KiB | [postgresql-16-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-wal2json` | 2.6 | `u22.x86_64` | pgdg | 63.6 KiB | [postgresql-16-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-wal2json` | 2.6 | `u24.aarch64` | pgdg | 53.8 KiB | [postgresql-16-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-wal2json` | 2.6 | `u24.x86_64` | pgdg | 56.0 KiB | [postgresql-16-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-16-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2json_15` | 2.6 | `el8.x86_64` | pgdg | 33.2 KiB | [wal2json_15-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/wal2json_15-2.6-1PGDG.rhel8.x86_64.rpm) |
| `wal2json_15` | 2.6 | `el8.aarch64` | pgdg | 31.3 KiB | [wal2json_15-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2json_15-2.6-1PGDG.rhel8.aarch64.rpm) |
| `wal2json_15` | 2.5 | `el8.aarch64` | pgdg | 30.6 KiB | [wal2json_15-2.5-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2json_15-2.5-2.rhel8.aarch64.rpm) |
| `wal2json_15` | 2.5 | `el8.aarch64` | pgdg | 30.6 KiB | [wal2json_15-2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2json_15-2.5-1.rhel8.aarch64.rpm) |
| `wal2json_15` | 2.5 | `el8.x86_64` | pgdg | 32.3 KiB | [wal2json_15-2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/wal2json_15-2.5-1.rhel8.x86_64.rpm) |
| `wal2json_15` | 2.6 | `el9.aarch64` | pgdg | 30.7 KiB | [wal2json_15-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2json_15-2.6-1PGDG.rhel9.aarch64.rpm) |
| `wal2json_15` | 2.5 | `el9.x86_64` | pgdg | 32.1 KiB | [wal2json_15-2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/wal2json_15-2.5-1.rhel9.x86_64.rpm) |
| `wal2json_15` | 2.5 | `el9.aarch64` | pgdg | 30.5 KiB | [wal2json_15-2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2json_15-2.5-1.rhel9.aarch64.rpm) |
| `wal2json_15` | 2.5 | `el9.aarch64` | pgdg | 30.6 KiB | [wal2json_15-2.5-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2json_15-2.5-2.rhel9.aarch64.rpm) |
| `postgresql-15-wal2json` | 2.6 | `d12.aarch64` | pgdg | 54.1 KiB | [postgresql-15-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-15-wal2json` | 2.6 | `d12.x86_64` | pgdg | 56.6 KiB | [postgresql-15-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-15-wal2json` | 2.6 | `u22.x86_64` | pgdg | 64.2 KiB | [postgresql-15-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-wal2json` | 2.6 | `u22.aarch64` | pgdg | 61.5 KiB | [postgresql-15-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-wal2json` | 2.6 | `u24.aarch64` | pgdg | 54.1 KiB | [postgresql-15-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-wal2json` | 2.6 | `u24.x86_64` | pgdg | 56.6 KiB | [postgresql-15-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-15-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2json_14` | 2.6 | `el8.x86_64` | pgdg | 33.2 KiB | [wal2json_14-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2json_14-2.6-1PGDG.rhel8.x86_64.rpm) |
| `wal2json_14` | 2.6 | `el8.aarch64` | pgdg | 31.2 KiB | [wal2json_14-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2json_14-2.6-1PGDG.rhel8.aarch64.rpm) |
| `wal2json_14` | 2.5 | `el8.aarch64` | pgdg | 30.6 KiB | [wal2json_14-2.5-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2json_14-2.5-2.rhel8.aarch64.rpm) |
| `wal2json_14` | 2.5 | `el8.x86_64` | pgdg | 32.4 KiB | [wal2json_14-2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2json_14-2.5-1.rhel8.x86_64.rpm) |
| `wal2json_14` | 2.5 | `el8.aarch64` | pgdg | 30.5 KiB | [wal2json_14-2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2json_14-2.5-1.rhel8.aarch64.rpm) |
| `wal2json_14` | 2.4 | `el8.x86_64` | pgdg | 76.4 KiB | [wal2json_14-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2json_14-2.4-1.rhel8.x86_64.rpm) |
| `wal2json_14` | 2.6 | `el9.aarch64` | pgdg | 30.7 KiB | [wal2json_14-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2json_14-2.6-1PGDG.rhel9.aarch64.rpm) |
| `wal2json_14` | 2.5 | `el9.aarch64` | pgdg | 30.6 KiB | [wal2json_14-2.5-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2json_14-2.5-2.rhel9.aarch64.rpm) |
| `wal2json_14` | 2.5 | `el9.aarch64` | pgdg | 30.6 KiB | [wal2json_14-2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2json_14-2.5-1.rhel9.aarch64.rpm) |
| `wal2json_14` | 2.5 | `el9.x86_64` | pgdg | 32.1 KiB | [wal2json_14-2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/wal2json_14-2.5-1.rhel9.x86_64.rpm) |
| `postgresql-14-wal2json` | 2.6 | `d12.x86_64` | pgdg | 56.2 KiB | [postgresql-14-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-14-wal2json` | 2.6 | `d12.aarch64` | pgdg | 53.8 KiB | [postgresql-14-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-14-wal2json` | 2.6 | `u22.x86_64` | pgdg | 64.3 KiB | [postgresql-14-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-wal2json` | 2.6 | `u22.aarch64` | pgdg | 61.5 KiB | [postgresql-14-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-wal2json` | 2.6 | `u24.x86_64` | pgdg | 56.2 KiB | [postgresql-14-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-wal2json` | 2.6 | `u24.aarch64` | pgdg | 53.7 KiB | [postgresql-14-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-14-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2json_13` | 2.6 | `el8.x86_64` | pgdg | 33.1 KiB | [wal2json_13-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/wal2json_13-2.6-1PGDG.rhel8.x86_64.rpm) |
| `wal2json_13` | 2.6 | `el8.aarch64` | pgdg | 31.3 KiB | [wal2json_13-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/wal2json_13-2.6-1PGDG.rhel8.aarch64.rpm) |
| `wal2json_13` | 2.5 | `el8.x86_64` | pgdg | 32.2 KiB | [wal2json_13-2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/wal2json_13-2.5-1.rhel8.x86_64.rpm) |
| `wal2json_13` | 2.5 | `el8.aarch64` | pgdg | 30.6 KiB | [wal2json_13-2.5-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/wal2json_13-2.5-2.rhel8.aarch64.rpm) |
| `wal2json_13` | 2.5 | `el8.aarch64` | pgdg | 30.5 KiB | [wal2json_13-2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/wal2json_13-2.5-1.rhel8.aarch64.rpm) |
| `wal2json_13` | 2.4 | `el8.x86_64` | pgdg | 74.5 KiB | [wal2json_13-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/wal2json_13-2.4-1.rhel8.x86_64.rpm) |
| `wal2json_13` | 2.6 | `el9.aarch64` | pgdg | 30.7 KiB | [wal2json_13-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/wal2json_13-2.6-1PGDG.rhel9.aarch64.rpm) |
| `wal2json_13` | 2.5 | `el9.aarch64` | pgdg | 30.5 KiB | [wal2json_13-2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/wal2json_13-2.5-1.rhel9.aarch64.rpm) |
| `wal2json_13` | 2.5 | `el9.x86_64` | pgdg | 32.1 KiB | [wal2json_13-2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/wal2json_13-2.5-1.rhel9.x86_64.rpm) |
| `wal2json_13` | 2.5 | `el9.aarch64` | pgdg | 30.6 KiB | [wal2json_13-2.5-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/wal2json_13-2.5-2.rhel9.aarch64.rpm) |
| `postgresql-13-wal2json` | 2.6 | `d12.aarch64` | pgdg | 53.4 KiB | [postgresql-13-wal2json_2.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg12+1_arm64.deb) |
| `postgresql-13-wal2json` | 2.6 | `d12.x86_64` | pgdg | 55.9 KiB | [postgresql-13-wal2json_2.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg12+1_amd64.deb) |
| `postgresql-13-wal2json` | 2.6 | `u22.aarch64` | pgdg | 60.3 KiB | [postgresql-13-wal2json_2.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-wal2json` | 2.6 | `u22.x86_64` | pgdg | 63.1 KiB | [postgresql-13-wal2json_2.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-wal2json` | 2.6 | `u24.x86_64` | pgdg | 56.1 KiB | [postgresql-13-wal2json_2.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-wal2json` | 2.6 | `u24.aarch64` | pgdg | 53.3 KiB | [postgresql-13-wal2json_2.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/w/wal2json/postgresql-13-wal2json_2.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/eulerto/wal2json" title="Repository" icon="github" subtitle="github.com/eulerto/wal2json" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="wal2json-2.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build get wal2json; # get wal2json source code
pig build dep wal2json; # install build dependencies
pig build pkg wal2json; # build extension rpm or deb
pig build ext wal2json; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install wal2json; # install by extension name, for the current active PG version
pig ext install wal2json; # install via package alias, for the active PG version
pig ext install wal2json -v 18;   # install for PG 18
pig ext install wal2json -v 17;   # install for PG 17
pig ext install wal2json -v 16;   # install for PG 16
pig ext install wal2json -v 15;   # install for PG 15
pig ext install wal2json -v 14;   # install for PG 14
pig ext install wal2json -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION wal2json;
```

