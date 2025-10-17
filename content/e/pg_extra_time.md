---
title: "pg_extra_time"
linkTitle: "pg_extra_time"
description: "Some date time functions and operators that,"
weight: 4220
categories: ["Util"]
width: full
---

Some date time functions and operators that,

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4220** | {{< badge content="pg_extra_time" link="https://github.com/bigsmoke/pg_extra_time" >}} | {{< ext "pg_extra_time" "pg_extra_time" >}} | `2.0.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dtr" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsql_tweaks" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "pg_cron" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_extra_time" >}} | `2.0.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_extra_time_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_extra_time" >}} | `2.0.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-extra-time" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-extra-time` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_extra_time_18" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_extra_time_17" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_extra_time_16" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_extra_time_15" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_extra_time_14" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_extra_time_18" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_extra_time_17" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_extra_time_16" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_extra_time_15" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_extra_time_14" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_extra_time_18" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_extra_time_18-2.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_extra_time_17" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_extra_time_17-2.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_extra_time_16" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_extra_time_16-2.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_extra_time_15" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_extra_time_15-2.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_extra_time_14" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_extra_time_14-2.0.0-1PGDG.rhel9.noarch.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_extra_time_18" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_extra_time_18-2.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_extra_time_17" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_extra_time_17-2.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_extra_time_16" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_extra_time_16-2.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_extra_time_15" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_extra_time_15-2.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_extra_time_14" "2.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_extra_time_14-2.0.0-1PGDG.rhel9.noarch.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-extra-time" >}}     | {{< pkg "postgresql-17-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-extra-time" >}}     | {{< pkg "postgresql-17-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-extra-time" >}}     | {{< pkg "postgresql-17-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-extra-time" >}}     | {{< pkg "postgresql-17-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-extra-time" >}}     | {{< pkg "postgresql-17-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-extra-time" >}}     | {{< pkg "postgresql-17-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-extra-time" "2.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_extra_time_18` | 2.0.0 | `el8.aarch64` | pgdg | 33.7 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_18` | 2.0.0 | `el8.x86_64` | pgdg | 33.8 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_extra_time_18-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_18` | 2.0.0 | `el9.aarch64` | pgdg | 32.3 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_extra_time_18-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_18` | 2.0.0 | `el9.x86_64` | pgdg | 32.3 KiB | [pg_extra_time_18-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_extra_time_18-2.0.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_extra_time_17` | 2.0.0 | `el8.aarch64` | pgdg | 33.7 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_17` | 2.0.0 | `el8.x86_64` | pgdg | 33.8 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_extra_time_17-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_17` | 1.1.3 | `el8.x86_64` | pgdg | 18.8 KiB | [pg_extra_time_17-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_extra_time_17-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_17` | 1.1.3 | `el8.aarch64` | pgdg | 18.8 KiB | [pg_extra_time_17-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_extra_time_17-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_17` | 2.0.0 | `el9.x86_64` | pgdg | 32.3 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_extra_time_17-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_17` | 2.0.0 | `el9.aarch64` | pgdg | 32.3 KiB | [pg_extra_time_17-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_extra_time_17-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_17` | 1.1.3 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_extra_time_17-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_extra_time_17-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_17` | 1.1.3 | `el9.x86_64` | pgdg | 18.6 KiB | [pg_extra_time_17-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_extra_time_17-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `postgresql-17-pg-extra-time` | 2.0.0 | `d12.aarch64` | pigsty | 39.9 KiB | [postgresql-17-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-extra-time` | 2.0.0 | `d12.x86_64` | pigsty | 40.0 KiB | [postgresql-17-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-extra-time` | 2.0.0 | `u22.aarch64` | pigsty | 36.4 KiB | [postgresql-17-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-extra-time` | 2.0.0 | `u22.x86_64` | pigsty | 36.4 KiB | [postgresql-17-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-extra-time` | 2.0.0 | `u24.aarch64` | pigsty | 36.2 KiB | [postgresql-17-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-extra-time` | 2.0.0 | `u24.x86_64` | pigsty | 36.2 KiB | [postgresql-17-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-17-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_extra_time_16` | 2.0.0 | `el8.x86_64` | pgdg | 33.8 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | 2.0.0 | `el8.aarch64` | pgdg | 33.7 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_extra_time_16-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | 1.1.3 | `el8.aarch64` | pgdg | 18.8 KiB | [pg_extra_time_16-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_extra_time_16-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | 1.1.3 | `el8.x86_64` | pgdg | 18.8 KiB | [pg_extra_time_16-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_extra_time_16-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | 1.1.2 | `el8.x86_64` | pgdg | 18.3 KiB | [pg_extra_time_16-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_extra_time_16-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | 1.1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_extra_time_16-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_extra_time_16-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_16` | 2.0.0 | `el9.aarch64` | pgdg | 32.3 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_extra_time_16-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_16` | 2.0.0 | `el9.x86_64` | pgdg | 32.3 KiB | [pg_extra_time_16-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_extra_time_16-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_16` | 1.1.3 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_extra_time_16-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_extra_time_16-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_16` | 1.1.3 | `el9.x86_64` | pgdg | 18.6 KiB | [pg_extra_time_16-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_extra_time_16-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_16` | 1.1.2 | `el9.aarch64` | pgdg | 18.0 KiB | [pg_extra_time_16-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_extra_time_16-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_16` | 1.1.2 | `el9.x86_64` | pgdg | 18.1 KiB | [pg_extra_time_16-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_extra_time_16-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `postgresql-16-pg-extra-time` | 2.0.0 | `d12.x86_64` | pigsty | 39.9 KiB | [postgresql-16-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-extra-time` | 2.0.0 | `d12.aarch64` | pigsty | 40.0 KiB | [postgresql-16-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-extra-time` | 2.0.0 | `u22.x86_64` | pigsty | 36.4 KiB | [postgresql-16-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-extra-time` | 2.0.0 | `u22.aarch64` | pigsty | 36.4 KiB | [postgresql-16-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-extra-time` | 2.0.0 | `u24.x86_64` | pigsty | 36.2 KiB | [postgresql-16-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-extra-time` | 2.0.0 | `u24.aarch64` | pigsty | 36.2 KiB | [postgresql-16-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-16-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_extra_time_15` | 2.0.0 | `el8.x86_64` | pgdg | 33.8 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | 2.0.0 | `el8.aarch64` | pgdg | 33.7 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_extra_time_15-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | 1.1.3 | `el8.aarch64` | pgdg | 18.8 KiB | [pg_extra_time_15-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_extra_time_15-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | 1.1.3 | `el8.x86_64` | pgdg | 18.8 KiB | [pg_extra_time_15-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_extra_time_15-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | 1.1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_extra_time_15-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_extra_time_15-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | 1.1.2 | `el8.x86_64` | pgdg | 18.3 KiB | [pg_extra_time_15-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_extra_time_15-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_15` | 2.0.0 | `el9.x86_64` | pgdg | 32.3 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_extra_time_15-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_15` | 2.0.0 | `el9.aarch64` | pgdg | 32.3 KiB | [pg_extra_time_15-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_extra_time_15-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_15` | 1.1.3 | `el9.x86_64` | pgdg | 18.6 KiB | [pg_extra_time_15-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_extra_time_15-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_15` | 1.1.3 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_extra_time_15-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_extra_time_15-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_15` | 1.1.2 | `el9.x86_64` | pgdg | 18.1 KiB | [pg_extra_time_15-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_extra_time_15-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_15` | 1.1.2 | `el9.aarch64` | pgdg | 18.0 KiB | [pg_extra_time_15-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_extra_time_15-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `postgresql-15-pg-extra-time` | 2.0.0 | `d12.aarch64` | pigsty | 39.9 KiB | [postgresql-15-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-extra-time` | 2.0.0 | `d12.x86_64` | pigsty | 39.9 KiB | [postgresql-15-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-extra-time` | 2.0.0 | `u22.aarch64` | pigsty | 36.4 KiB | [postgresql-15-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-extra-time` | 2.0.0 | `u22.x86_64` | pigsty | 36.4 KiB | [postgresql-15-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-extra-time` | 2.0.0 | `u24.aarch64` | pigsty | 36.2 KiB | [postgresql-15-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-extra-time` | 2.0.0 | `u24.x86_64` | pigsty | 36.2 KiB | [postgresql-15-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-15-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_extra_time_14` | 2.0.0 | `el8.aarch64` | pgdg | 33.7 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | 2.0.0 | `el8.x86_64` | pgdg | 33.8 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_extra_time_14-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | 1.1.3 | `el8.aarch64` | pgdg | 18.8 KiB | [pg_extra_time_14-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_extra_time_14-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | 1.1.3 | `el8.x86_64` | pgdg | 18.8 KiB | [pg_extra_time_14-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_extra_time_14-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | 1.1.2 | `el8.x86_64` | pgdg | 18.3 KiB | [pg_extra_time_14-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_extra_time_14-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | 1.1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_extra_time_14-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_extra_time_14-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_14` | 2.0.0 | `el9.x86_64` | pgdg | 32.3 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_extra_time_14-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_14` | 2.0.0 | `el9.aarch64` | pgdg | 32.3 KiB | [pg_extra_time_14-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_extra_time_14-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_14` | 1.1.3 | `el9.x86_64` | pgdg | 18.6 KiB | [pg_extra_time_14-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_extra_time_14-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_14` | 1.1.3 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_extra_time_14-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_extra_time_14-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_14` | 1.1.2 | `el9.aarch64` | pgdg | 18.0 KiB | [pg_extra_time_14-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_extra_time_14-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_14` | 1.1.2 | `el9.x86_64` | pgdg | 18.1 KiB | [pg_extra_time_14-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_extra_time_14-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `postgresql-14-pg-extra-time` | 2.0.0 | `d12.x86_64` | pigsty | 39.9 KiB | [postgresql-14-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-extra-time` | 2.0.0 | `d12.aarch64` | pigsty | 39.9 KiB | [postgresql-14-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-extra-time` | 2.0.0 | `u22.x86_64` | pigsty | 36.4 KiB | [postgresql-14-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-extra-time` | 2.0.0 | `u22.aarch64` | pigsty | 36.4 KiB | [postgresql-14-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-extra-time` | 2.0.0 | `u24.aarch64` | pigsty | 36.2 KiB | [postgresql-14-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-extra-time` | 2.0.0 | `u24.x86_64` | pigsty | 36.2 KiB | [postgresql-14-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-14-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_extra_time_13` | 2.0.0 | `el8.x86_64` | pgdg | 33.8 KiB | [pg_extra_time_13-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_extra_time_13-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_13` | 2.0.0 | `el8.aarch64` | pgdg | 33.7 KiB | [pg_extra_time_13-2.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_extra_time_13-2.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_13` | 1.1.3 | `el8.x86_64` | pgdg | 18.8 KiB | [pg_extra_time_13-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_extra_time_13-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_13` | 1.1.3 | `el8.aarch64` | pgdg | 18.8 KiB | [pg_extra_time_13-1.1.3-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_extra_time_13-1.1.3-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_13` | 1.1.2 | `el8.aarch64` | pgdg | 18.3 KiB | [pg_extra_time_13-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_extra_time_13-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_13` | 1.1.2 | `el8.x86_64` | pgdg | 18.3 KiB | [pg_extra_time_13-1.1.2-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_extra_time_13-1.1.2-1PGDG.rhel8.noarch.rpm) |
| `pg_extra_time_13` | 2.0.0 | `el9.aarch64` | pgdg | 32.3 KiB | [pg_extra_time_13-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_extra_time_13-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_13` | 2.0.0 | `el9.x86_64` | pgdg | 32.3 KiB | [pg_extra_time_13-2.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_extra_time_13-2.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_13` | 1.1.3 | `el9.x86_64` | pgdg | 18.6 KiB | [pg_extra_time_13-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_extra_time_13-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_13` | 1.1.3 | `el9.aarch64` | pgdg | 18.6 KiB | [pg_extra_time_13-1.1.3-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_extra_time_13-1.1.3-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_13` | 1.1.2 | `el9.aarch64` | pgdg | 18.0 KiB | [pg_extra_time_13-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_extra_time_13-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `pg_extra_time_13` | 1.1.2 | `el9.x86_64` | pgdg | 18.1 KiB | [pg_extra_time_13-1.1.2-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_extra_time_13-1.1.2-1PGDG.rhel9.noarch.rpm) |
| `postgresql-13-pg-extra-time` | 2.0.0 | `d12.aarch64` | pigsty | 39.9 KiB | [postgresql-13-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-13-pg-extra-time_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-extra-time` | 2.0.0 | `d12.x86_64` | pigsty | 39.9 KiB | [postgresql-13-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-extra-time/postgresql-13-pg-extra-time_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-extra-time` | 2.0.0 | `u22.aarch64` | pigsty | 36.4 KiB | [postgresql-13-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-13-pg-extra-time_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-extra-time` | 2.0.0 | `u22.x86_64` | pigsty | 36.4 KiB | [postgresql-13-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-extra-time/postgresql-13-pg-extra-time_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-extra-time` | 2.0.0 | `u24.aarch64` | pigsty | 36.2 KiB | [postgresql-13-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-13-pg-extra-time_2.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-extra-time` | 2.0.0 | `u24.x86_64` | pigsty | 36.2 KiB | [postgresql-13-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-extra-time/postgresql-13-pg-extra-time_2.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_extra_time" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_extra_time" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_extra_time-2.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_extra_time; # get pg_extra_time source code
pig build dep pg_extra_time; # install build dependencies
pig build pkg pg_extra_time; # build extension rpm or deb
pig build ext pg_extra_time; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_extra_time; # install by extension name, for the current active PG version
pig ext install pg_extra_time; # install via package alias, for the active PG version
pig ext install pg_extra_time -v 18;   # install for PG 18
pig ext install pg_extra_time -v 17;   # install for PG 17
pig ext install pg_extra_time -v 16;   # install for PG 16
pig ext install pg_extra_time -v 15;   # install for PG 15
pig ext install pg_extra_time -v 14;   # install for PG 14
pig ext install pg_extra_time -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_extra_time;
```

