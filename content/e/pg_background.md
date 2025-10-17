---
title: "pg_background"
linkTitle: "pg_background"
description: "Run SQL queries in the background"
weight: 1100
categories: ["Time"]
width: full
---

Run SQL queries in the background

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1100** | {{< badge content="pg_background" link="https://github.com/vibhorkum/pg_background" >}} | {{< ext "pg_background" "pg_background" >}} | `1.3` | {{< category "TIME" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_later" >}} {{< ext "pgq" >}} {{< ext "timescaledb" >}} {{< ext "timescaledb_toolkit" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_background" >}} | `1.3` | {{< badge content="18" color="red" alt="pg_background_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_background_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_background" >}} | `1.3` | {{< badge content="18" color="red" alt="postgresql-18-pg-background" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-background` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_background_18" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_background_17" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_background_16" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_background_15" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_background_14" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_background_18" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_background_17" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_background_16" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_background_15" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_background_14" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_background_18" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.5-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_background_17" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.3-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_background_16" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.3-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_background_15" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.3-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_background_14" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.3-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_background_18" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.5-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_background_17" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.3-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_background_16" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.3-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_background_15" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.3-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_background_14" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-background" >}}     | {{< pkg "postgresql-17-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-background" >}}     | {{< pkg "postgresql-17-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-background" >}}     | {{< pkg "postgresql-17-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-background" >}}     | {{< pkg "postgresql-17-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-background" >}}     | {{< pkg "postgresql-17-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-background" >}}     | {{< pkg "postgresql-17-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-background" "1.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_background_18` | 1.5 | `el8.x86_64` | pgdg | 22.5 KiB | [pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_background_18-1.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_18` | 1.5 | `el8.aarch64` | pgdg | 22.0 KiB | [pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_background_18-1.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_18` | 1.5 | `el9.aarch64` | pgdg | 22.0 KiB | [pg_background_18-1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_background_18-1.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_18` | 1.5 | `el9.x86_64` | pgdg | 22.7 KiB | [pg_background_18-1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_background_18-1.5-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_background_17` | 1.3 | `el8.aarch64` | pgdg | 21.3 KiB | [pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | 1.3 | `el8.x86_64` | pgdg | 21.9 KiB | [pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | 1.2 | `el8.x86_64` | pgdg | 20.1 KiB | [pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_background_17-1.2-2PGDG.rhel8.x86_64.rpm) |
| `pg_background_17` | 1.2 | `el8.aarch64` | pgdg | 19.6 KiB | [pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_background_17-1.2-2PGDG.rhel8.aarch64.rpm) |
| `pg_background_17` | 1.3 | `el9.x86_64` | pgdg | 22.3 KiB | [pg_background_17-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_17` | 1.3 | `el9.aarch64` | pgdg | 21.6 KiB | [pg_background_17-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_17` | 1.2 | `el9.x86_64` | pgdg | 20.5 KiB | [pg_background_17-1.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_background_17-1.2-2PGDG.rhel9.x86_64.rpm) |
| `pg_background_17` | 1.2 | `el9.aarch64` | pgdg | 19.9 KiB | [pg_background_17-1.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_background_17-1.2-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pg-background` | 1.3 | `d12.x86_64` | pigsty | 45.6 KiB | [postgresql-17-pg-background_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-background` | 1.3 | `d12.aarch64` | pigsty | 44.5 KiB | [postgresql-17-pg-background_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-background` | 1.3 | `u22.x86_64` | pigsty | 48.7 KiB | [postgresql-17-pg-background_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-background` | 1.3 | `u22.aarch64` | pigsty | 47.9 KiB | [postgresql-17-pg-background_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-background` | 1.3 | `u24.x86_64` | pigsty | 43.4 KiB | [postgresql-17-pg-background_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-background` | 1.3 | `u24.aarch64` | pigsty | 42.5 KiB | [postgresql-17-pg-background_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-17-pg-background_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_background_16` | 1.3 | `el8.aarch64` | pgdg | 21.3 KiB | [pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | 1.3 | `el8.x86_64` | pgdg | 21.9 KiB | [pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | 1.2 | `el8.x86_64` | pgdg | 19.7 KiB | [pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_background_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_16` | 1.2 | `el8.aarch64` | pgdg | 19.2 KiB | [pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_background_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_16` | 1.3 | `el9.aarch64` | pgdg | 21.6 KiB | [pg_background_16-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_16` | 1.3 | `el9.x86_64` | pgdg | 22.3 KiB | [pg_background_16-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_16` | 1.2 | `el9.x86_64` | pgdg | 19.9 KiB | [pg_background_16-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_background_16-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_16` | 1.2 | `el9.aarch64` | pgdg | 19.2 KiB | [pg_background_16-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_background_16-1.2-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pg-background` | 1.3 | `d12.aarch64` | pigsty | 44.5 KiB | [postgresql-16-pg-background_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-background` | 1.3 | `d12.x86_64` | pigsty | 45.4 KiB | [postgresql-16-pg-background_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-background` | 1.3 | `u22.x86_64` | pigsty | 48.6 KiB | [postgresql-16-pg-background_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-background` | 1.3 | `u22.aarch64` | pigsty | 47.8 KiB | [postgresql-16-pg-background_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-background` | 1.3 | `u24.aarch64` | pigsty | 42.6 KiB | [postgresql-16-pg-background_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-background` | 1.3 | `u24.x86_64` | pigsty | 43.3 KiB | [postgresql-16-pg-background_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-16-pg-background_1.3-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_background_15` | 1.3 | `el8.aarch64` | pgdg | 21.2 KiB | [pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | 1.3 | `el8.x86_64` | pgdg | 21.9 KiB | [pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | 1.2 | `el8.aarch64` | pgdg | 19.2 KiB | [pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_15` | 1.2 | `el8.x86_64` | pgdg | 19.6 KiB | [pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_15` | 1.0 | `el8.x86_64` | pgdg | 39.3 KiB | [pg_background_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_background_15-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_15` | 1.0 | `el8.aarch64` | pgdg | 38.7 KiB | [pg_background_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_background_15-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_15` | 1.3 | `el9.x86_64` | pgdg | 22.2 KiB | [pg_background_15-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_15` | 1.3 | `el9.aarch64` | pgdg | 21.6 KiB | [pg_background_15-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_15` | 1.2 | `el9.aarch64` | pgdg | 19.1 KiB | [pg_background_15-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_15` | 1.2 | `el9.x86_64` | pgdg | 19.9 KiB | [pg_background_15-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_15` | 1.0 | `el9.x86_64` | pgdg | 40.6 KiB | [pg_background_15-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_background_15-1.0-1.rhel9.x86_64.rpm) |
| `pg_background_15` | 1.0 | `el9.aarch64` | pgdg | 39.7 KiB | [pg_background_15-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_background_15-1.0-1.rhel9.aarch64.rpm) |
| `postgresql-15-pg-background` | 1.3 | `d12.aarch64` | pigsty | 44.4 KiB | [postgresql-15-pg-background_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-background` | 1.3 | `d12.x86_64` | pigsty | 45.4 KiB | [postgresql-15-pg-background_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-background` | 1.3 | `u22.aarch64` | pigsty | 47.7 KiB | [postgresql-15-pg-background_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-background` | 1.3 | `u22.x86_64` | pigsty | 48.6 KiB | [postgresql-15-pg-background_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-background` | 1.3 | `u24.x86_64` | pigsty | 43.3 KiB | [postgresql-15-pg-background_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-background` | 1.3 | `u24.aarch64` | pigsty | 42.5 KiB | [postgresql-15-pg-background_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-15-pg-background_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_background_14` | 1.3 | `el8.x86_64` | pgdg | 21.9 KiB | [pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_14` | 1.3 | `el8.aarch64` | pgdg | 21.2 KiB | [pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_14` | 1.0 | `el8.aarch64` | pgdg | 38.6 KiB | [pg_background_14-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_background_14-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_14` | 1.0 | `el8.x86_64` | pgdg | 39.6 KiB | [pg_background_14-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_background_14-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_14` | 1.3 | `el9.aarch64` | pgdg | 21.6 KiB | [pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_14` | 1.3 | `el9.x86_64` | pgdg | 22.2 KiB | [pg_background_14-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_14` | 1.2 | `el9.x86_64` | pgdg | 19.9 KiB | [pg_background_14-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_background_14-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_14` | 1.0 | `el9.aarch64` | pgdg | 39.5 KiB | [pg_background_14-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_background_14-1.0-1.rhel9.aarch64.rpm) |
| `postgresql-14-pg-background` | 1.3 | `d12.x86_64` | pigsty | 45.4 KiB | [postgresql-14-pg-background_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-background` | 1.3 | `d12.aarch64` | pigsty | 44.5 KiB | [postgresql-14-pg-background_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-background` | 1.3 | `u22.aarch64` | pigsty | 47.7 KiB | [postgresql-14-pg-background_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-background` | 1.3 | `u22.x86_64` | pigsty | 48.6 KiB | [postgresql-14-pg-background_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-background` | 1.3 | `u24.aarch64` | pigsty | 42.5 KiB | [postgresql-14-pg-background_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-background` | 1.3 | `u24.x86_64` | pigsty | 43.3 KiB | [postgresql-14-pg-background_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-14-pg-background_1.3-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_background_13` | 1.3 | `el8.x86_64` | pgdg | 21.8 KiB | [pg_background_13-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_background_13-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_background_13` | 1.3 | `el8.aarch64` | pgdg | 21.2 KiB | [pg_background_13-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_background_13-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_background_13` | 1.0 | `el8.aarch64` | pgdg | 38.4 KiB | [pg_background_13-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_background_13-1.0-1.rhel8.aarch64.rpm) |
| `pg_background_13` | 1.0 | `el8.x86_64` | pgdg | 39.1 KiB | [pg_background_13-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_background_13-1.0-1.rhel8.x86_64.rpm) |
| `pg_background_13` | 1.3 | `el9.aarch64` | pgdg | 21.6 KiB | [pg_background_13-1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_background_13-1.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_background_13` | 1.3 | `el9.x86_64` | pgdg | 22.2 KiB | [pg_background_13-1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_background_13-1.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_13` | 1.2 | `el9.x86_64` | pgdg | 19.9 KiB | [pg_background_13-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_background_13-1.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_background_13` | 1.0 | `el9.aarch64` | pgdg | 39.4 KiB | [pg_background_13-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_background_13-1.0-1.rhel9.aarch64.rpm) |
| `postgresql-13-pg-background` | 1.3 | `d12.x86_64` | pigsty | 45.1 KiB | [postgresql-13-pg-background_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-13-pg-background_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-background` | 1.3 | `d12.aarch64` | pigsty | 44.2 KiB | [postgresql-13-pg-background_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-background/postgresql-13-pg-background_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-background` | 1.3 | `u22.x86_64` | pigsty | 48.4 KiB | [postgresql-13-pg-background_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-13-pg-background_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-background` | 1.3 | `u22.aarch64` | pigsty | 47.2 KiB | [postgresql-13-pg-background_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-background/postgresql-13-pg-background_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-background` | 1.3 | `u24.x86_64` | pigsty | 43.1 KiB | [postgresql-13-pg-background_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-13-pg-background_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-background` | 1.3 | `u24.aarch64` | pigsty | 42.4 KiB | [postgresql-13-pg-background_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-background/postgresql-13-pg-background_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/vibhorkum/pg_background" title="Repository" icon="github" subtitle="github.com/vibhorkum/pg_background" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_background-1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_background; # get pg_background source code
pig build dep pg_background; # install build dependencies
pig build pkg pg_background; # build extension rpm or deb
pig build ext pg_background; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_background; # install by extension name, for the current active PG version
pig ext install pg_background; # install via package alias, for the active PG version
pig ext install pg_background -v 17;   # install for PG 17
pig ext install pg_background -v 16;   # install for PG 16
pig ext install pg_background -v 15;   # install for PG 15
pig ext install pg_background -v 14;   # install for PG 14
pig ext install pg_background -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_background;
```

