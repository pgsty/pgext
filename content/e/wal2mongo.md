---
title: "wal2mongo"
linkTitle: "wal2mongo"
description: "PostgreSQL logical decoding output plugin for MongoDB"
weight: 9640
categories: ["Etl"]
width: full
---

PostgreSQL logical decoding output plugin for MongoDB

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9640** | {{< badge content="wal2mongo" link="https://github.com/HighgoSoftware/wal2mongo" >}} | {{< ext "wal2mongo" "wal2mongo" >}} | `1.0.7` | {{< category "ETL" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s----" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "decoder_raw" >}} {{< ext "documentdb" >}} {{< ext "pglogical" >}} {{< ext "test_decoding" >}} {{< ext "pgoutput" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/wal2mongo" >}} | `1.0.7` | {{< badge content="18" color="red" alt="wal2mongo_18*" >}} {{< badge content="17" color="red" alt="wal2mongo_17*" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `wal2mongo_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/wal2mongo" >}} | `1.0.7` | {{< badge content="18" color="red" alt="postgresql-18-wal2mongo" >}} {{< badge content="17" color="red" alt="postgresql-17-wal2mongo" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-wal2mongo` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "wal2mongo_18" >}}     |    {{< pkg "wal2mongo_17" >}}     | {{< pkg "wal2mongo_16" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/wal2mongo_16-1.0.7-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "wal2mongo_15" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/wal2mongo_15-1.0.7-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "wal2mongo_14" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2mongo_14-1.0.7-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "wal2mongo_18" >}}     |    {{< pkg "wal2mongo_17" >}}     | {{< pkg "wal2mongo_16" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/wal2mongo_16-1.0.7-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "wal2mongo_15" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2mongo_15-1.0.7-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "wal2mongo_14" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2mongo_14-1.0.7-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "wal2mongo_18" >}}     |    {{< pkg "wal2mongo_17" >}}     | {{< pkg "wal2mongo_16" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/wal2mongo_16-1.0.7-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "wal2mongo_15" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/wal2mongo_15-1.0.7-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "wal2mongo_14" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/wal2mongo_14-1.0.7-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "wal2mongo_18" >}}     |    {{< pkg "wal2mongo_17" >}}     | {{< pkg "wal2mongo_16" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/wal2mongo_16-1.0.7-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "wal2mongo_15" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2mongo_15-1.0.7-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "wal2mongo_14" "1.0.7" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2mongo_14-1.0.7-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-wal2mongo" >}}     |    {{< pkg "postgresql-17-wal2mongo" >}}     | {{< pkg "postgresql-16-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-wal2mongo" >}}     |    {{< pkg "postgresql-17-wal2mongo" >}}     | {{< pkg "postgresql-16-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-wal2mongo" >}}     |    {{< pkg "postgresql-17-wal2mongo" >}}     | {{< pkg "postgresql-16-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-wal2mongo" >}}     |    {{< pkg "postgresql-17-wal2mongo" >}}     | {{< pkg "postgresql-16-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-wal2mongo" >}}     |    {{< pkg "postgresql-17-wal2mongo" >}}     | {{< pkg "postgresql-16-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-wal2mongo" >}}     |    {{< pkg "postgresql-17-wal2mongo" >}}     | {{< pkg "postgresql-16-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-wal2mongo" "1.0.7" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2mongo_16` | 1.0.7 | `el8.aarch64` | pgdg | 20.0 KiB | [wal2mongo_16-1.0.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/wal2mongo_16-1.0.7-1PGDG.rhel8.aarch64.rpm) |
| `wal2mongo_16` | 1.0.7 | `el8.x86_64` | pgdg | 20.2 KiB | [wal2mongo_16-1.0.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/wal2mongo_16-1.0.7-1PGDG.rhel8.x86_64.rpm) |
| `wal2mongo_16` | 1.0.7 | `el9.aarch64` | pgdg | 19.7 KiB | [wal2mongo_16-1.0.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/wal2mongo_16-1.0.7-1PGDG.rhel9.aarch64.rpm) |
| `wal2mongo_16` | 1.0.7 | `el9.x86_64` | pgdg | 20.1 KiB | [wal2mongo_16-1.0.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/wal2mongo_16-1.0.7-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-wal2mongo` | 1.0.7 | `d12.x86_64` | pigsty | 37.7 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-wal2mongo` | 1.0.7 | `d12.aarch64` | pigsty | 37.3 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-wal2mongo` | 1.0.7 | `u22.aarch64` | pigsty | 39.0 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-wal2mongo` | 1.0.7 | `u22.x86_64` | pigsty | 39.5 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-wal2mongo` | 1.0.7 | `u24.x86_64` | pigsty | 36.2 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-wal2mongo` | 1.0.7 | `u24.aarch64` | pigsty | 36.0 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2mongo_15` | 1.0.7 | `el8.x86_64` | pgdg | 20.2 KiB | [wal2mongo_15-1.0.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/wal2mongo_15-1.0.7-1PGDG.rhel8.x86_64.rpm) |
| `wal2mongo_15` | 1.0.7 | `el8.aarch64` | pgdg | 19.9 KiB | [wal2mongo_15-1.0.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2mongo_15-1.0.7-1PGDG.rhel8.aarch64.rpm) |
| `wal2mongo_15` | 1.0.7 | `el9.aarch64` | pgdg | 19.6 KiB | [wal2mongo_15-1.0.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2mongo_15-1.0.7-1PGDG.rhel9.aarch64.rpm) |
| `wal2mongo_15` | 1.0.7 | `el9.x86_64` | pgdg | 20.0 KiB | [wal2mongo_15-1.0.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/wal2mongo_15-1.0.7-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-15-wal2mongo` | 1.0.7 | `d12.aarch64` | pigsty | 37.0 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-wal2mongo` | 1.0.7 | `d12.x86_64` | pigsty | 37.4 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-wal2mongo` | 1.0.7 | `u22.x86_64` | pigsty | 39.3 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-wal2mongo` | 1.0.7 | `u22.aarch64` | pigsty | 38.7 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-wal2mongo` | 1.0.7 | `u24.aarch64` | pigsty | 35.8 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-wal2mongo` | 1.0.7 | `u24.x86_64` | pigsty | 36.0 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2mongo_14` | 1.0.7 | `el8.x86_64` | pgdg | 20.2 KiB | [wal2mongo_14-1.0.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2mongo_14-1.0.7-1PGDG.rhel8.x86_64.rpm) |
| `wal2mongo_14` | 1.0.7 | `el8.aarch64` | pgdg | 19.9 KiB | [wal2mongo_14-1.0.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2mongo_14-1.0.7-1PGDG.rhel8.aarch64.rpm) |
| `wal2mongo_14` | 1.0.7 | `el9.x86_64` | pgdg | 20.0 KiB | [wal2mongo_14-1.0.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/wal2mongo_14-1.0.7-1PGDG.rhel9.x86_64.rpm) |
| `wal2mongo_14` | 1.0.7 | `el9.aarch64` | pgdg | 19.6 KiB | [wal2mongo_14-1.0.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2mongo_14-1.0.7-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-wal2mongo` | 1.0.7 | `d12.x86_64` | pigsty | 37.4 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-wal2mongo` | 1.0.7 | `d12.aarch64` | pigsty | 37.0 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-wal2mongo` | 1.0.7 | `u22.aarch64` | pigsty | 38.6 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-wal2mongo` | 1.0.7 | `u22.x86_64` | pigsty | 39.2 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-wal2mongo` | 1.0.7 | `u24.aarch64` | pigsty | 35.7 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-wal2mongo` | 1.0.7 | `u24.x86_64` | pigsty | 36.0 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `wal2mongo_13` | 1.0.7 | `el8.aarch64` | pgdg | 19.9 KiB | [wal2mongo_13-1.0.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/wal2mongo_13-1.0.7-1PGDG.rhel8.aarch64.rpm) |
| `wal2mongo_13` | 1.0.7 | `el8.x86_64` | pgdg | 20.1 KiB | [wal2mongo_13-1.0.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/wal2mongo_13-1.0.7-1PGDG.rhel8.x86_64.rpm) |
| `wal2mongo_13` | 1.0.7 | `el9.x86_64` | pgdg | 20.0 KiB | [wal2mongo_13-1.0.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/wal2mongo_13-1.0.7-1PGDG.rhel9.x86_64.rpm) |
| `wal2mongo_13` | 1.0.7 | `el9.aarch64` | pgdg | 19.6 KiB | [wal2mongo_13-1.0.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/wal2mongo_13-1.0.7-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-13-wal2mongo` | 1.0.7 | `d12.aarch64` | pigsty | 36.4 KiB | [postgresql-13-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-13-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-wal2mongo` | 1.0.7 | `d12.x86_64` | pigsty | 37.2 KiB | [postgresql-13-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-13-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-wal2mongo` | 1.0.7 | `u22.x86_64` | pigsty | 38.8 KiB | [postgresql-13-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-13-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-wal2mongo` | 1.0.7 | `u22.aarch64` | pigsty | 38.3 KiB | [postgresql-13-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-13-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-wal2mongo` | 1.0.7 | `u24.aarch64` | pigsty | 35.3 KiB | [postgresql-13-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-13-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-wal2mongo` | 1.0.7 | `u24.x86_64` | pigsty | 35.9 KiB | [postgresql-13-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-13-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HighgoSoftware/wal2mongo" title="Repository" icon="github" subtitle="github.com/HighgoSoftware/wal2mongo" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="wal2mongo-1.0.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build get wal2mongo; # get wal2mongo source code
pig build dep wal2mongo; # install build dependencies
pig build pkg wal2mongo; # build extension rpm or deb
pig build ext wal2mongo; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install wal2mongo; # install by extension name, for the current active PG version
pig ext install wal2mongo; # install via package alias, for the active PG version
pig ext install wal2mongo -v 16;   # install for PG 16
pig ext install wal2mongo -v 15;   # install for PG 15
pig ext install wal2mongo -v 14;   # install for PG 14
pig ext install wal2mongo -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION wal2mongo;
```

