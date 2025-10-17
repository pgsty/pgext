---
title: "pgsentinel"
linkTitle: "pgsentinel"
description: "active session history"
weight: 6280
categories: ["Stat"]
width: full
---

active session history

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6280** | {{< badge content="pgsentinel" link="https://github.com/pgsentinel/pgsentinel" >}} | {{< ext "pgsentinel" "pgsentinel" >}} | `1.2.0` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "system_stats" >}} {{< ext "pgnodemx" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_wait_sampling" >}} {{< ext "bgw_replstatus" >}} {{< ext "pg_profile" >}} {{< ext "pg_proctab" >}} {{< ext "powa" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgsentinel" >}} | `1.2.0` | {{< badge content="18" color="red" alt="pgsentinel_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgsentinel_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgsentinel" >}} | `1.2.0` | {{< badge content="18" color="red" alt="postgresql-18-pgsentinel" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgsentinel` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgsentinel_18" >}}     | {{< pkg "pgsentinel_17" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsentinel_16" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsentinel_15" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsentinel_14" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgsentinel_18" >}}     | {{< pkg "pgsentinel_17" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsentinel_16" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsentinel_15" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsentinel_14" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgsentinel_18" >}}     | {{< pkg "pgsentinel_17" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsentinel_16" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsentinel_15" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsentinel_14" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgsentinel_18" >}}     | {{< pkg "pgsentinel_17" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsentinel_16" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsentinel_15" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsentinel_14" "1.2.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-pgsentinel" "1.3.0" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `postgresql-18-pgsentinel` | 1.3.0 | `d12.x86_64` | pgdg | 44.2 KiB | [postgresql-18-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgsentinel` | 1.3.0 | `d12.aarch64` | pgdg | 42.7 KiB | [postgresql-18-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgsentinel` | 1.3.0 | `u22.x86_64` | pgdg | 44.7 KiB | [postgresql-18-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgsentinel` | 1.3.0 | `u22.aarch64` | pgdg | 43.6 KiB | [postgresql-18-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgsentinel` | 1.3.0 | `u24.x86_64` | pgdg | 44.2 KiB | [postgresql-18-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgsentinel` | 1.3.0 | `u24.aarch64` | pgdg | 42.8 KiB | [postgresql-18-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-18-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsentinel_17` | 1.2.0 | `el8.x86_64` | pgdg | 23.6 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_17` | 1.2.0 | `el8.aarch64` | pgdg | 22.8 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_17` | 1.1.0 | `el8.aarch64` | pigsty | 21.1 KiB | [pgsentinel_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_17` | 1.1.0 | `el8.x86_64` | pigsty | 21.9 KiB | [pgsentinel_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_17` | 1.2.0 | `el9.x86_64` | pgdg | 24.1 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsentinel_17-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_17` | 1.2.0 | `el9.aarch64` | pgdg | 23.2 KiB | [pgsentinel_17-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsentinel_17-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_17` | 1.1.0 | `el9.aarch64` | pigsty | 21.8 KiB | [pgsentinel_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgsentinel_17` | 1.1.0 | `el9.x86_64` | pigsty | 22.5 KiB | [pgsentinel_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pgsentinel` | 1.3.0 | `d12.x86_64` | pgdg | 44.2 KiB | [postgresql-17-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgsentinel` | 1.3.0 | `d12.aarch64` | pgdg | 42.7 KiB | [postgresql-17-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgsentinel` | 1.2.0 | `d12.aarch64` | pigsty | 38.8 KiB | [postgresql-17-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgsentinel` | 1.2.0 | `d12.x86_64` | pigsty | 40.4 KiB | [postgresql-17-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgsentinel` | 1.3.0 | `u22.x86_64` | pgdg | 52.7 KiB | [postgresql-17-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgsentinel` | 1.3.0 | `u22.aarch64` | pgdg | 51.6 KiB | [postgresql-17-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgsentinel` | 1.2.0 | `u22.x86_64` | pigsty | 51.6 KiB | [postgresql-17-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgsentinel` | 1.2.0 | `u22.aarch64` | pigsty | 50.8 KiB | [postgresql-17-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgsentinel` | 1.3.0 | `u24.x86_64` | pgdg | 44.1 KiB | [postgresql-17-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgsentinel` | 1.3.0 | `u24.aarch64` | pgdg | 42.8 KiB | [postgresql-17-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgsentinel` | 1.2.0 | `u24.aarch64` | pigsty | 41.6 KiB | [postgresql-17-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgsentinel` | 1.2.0 | `u24.x86_64` | pigsty | 42.7 KiB | [postgresql-17-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-17-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsentinel_16` | 1.2.0 | `el8.x86_64` | pgdg | 23.6 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_16` | 1.2.0 | `el8.aarch64` | pgdg | 22.8 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_16` | 1.1.0 | `el8.x86_64` | pigsty | 21.9 KiB | [pgsentinel_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_16` | 1.1.0 | `el8.aarch64` | pigsty | 21.1 KiB | [pgsentinel_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_16` | 1.2.0 | `el9.aarch64` | pgdg | 23.2 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsentinel_16-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_16` | 1.2.0 | `el9.x86_64` | pgdg | 24.1 KiB | [pgsentinel_16-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsentinel_16-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_16` | 1.1.0 | `el9.x86_64` | pigsty | 22.5 KiB | [pgsentinel_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_16` | 1.1.0 | `el9.aarch64` | pigsty | 21.8 KiB | [pgsentinel_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgsentinel` | 1.3.0 | `d12.x86_64` | pgdg | 44.2 KiB | [postgresql-16-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgsentinel` | 1.3.0 | `d12.aarch64` | pgdg | 42.6 KiB | [postgresql-16-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgsentinel` | 1.2.0 | `d12.x86_64` | pigsty | 40.3 KiB | [postgresql-16-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgsentinel` | 1.2.0 | `d12.aarch64` | pigsty | 38.8 KiB | [postgresql-16-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgsentinel` | 1.3.0 | `u22.aarch64` | pgdg | 51.5 KiB | [postgresql-16-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgsentinel` | 1.3.0 | `u22.x86_64` | pgdg | 52.6 KiB | [postgresql-16-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgsentinel` | 1.2.0 | `u22.x86_64` | pigsty | 51.4 KiB | [postgresql-16-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgsentinel` | 1.2.0 | `u22.aarch64` | pigsty | 50.6 KiB | [postgresql-16-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgsentinel` | 1.3.0 | `u24.aarch64` | pgdg | 42.8 KiB | [postgresql-16-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgsentinel` | 1.3.0 | `u24.x86_64` | pgdg | 44.1 KiB | [postgresql-16-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgsentinel` | 1.2.0 | `u24.x86_64` | pigsty | 42.6 KiB | [postgresql-16-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgsentinel` | 1.2.0 | `u24.aarch64` | pigsty | 41.6 KiB | [postgresql-16-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-16-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsentinel_15` | 1.2.0 | `el8.aarch64` | pgdg | 22.8 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_15` | 1.2.0 | `el8.x86_64` | pgdg | 23.8 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_15` | 1.1.0 | `el8.aarch64` | pigsty | 21.1 KiB | [pgsentinel_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_15` | 1.1.0 | `el8.x86_64` | pigsty | 22.0 KiB | [pgsentinel_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_15` | 1.2.0 | `el9.aarch64` | pgdg | 23.4 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsentinel_15-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_15` | 1.2.0 | `el9.x86_64` | pgdg | 24.2 KiB | [pgsentinel_15-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsentinel_15-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_15` | 1.1.0 | `el9.x86_64` | pigsty | 22.7 KiB | [pgsentinel_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_15` | 1.1.0 | `el9.aarch64` | pigsty | 21.9 KiB | [pgsentinel_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgsentinel` | 1.3.0 | `d12.aarch64` | pgdg | 42.6 KiB | [postgresql-15-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgsentinel` | 1.3.0 | `d12.x86_64` | pgdg | 44.2 KiB | [postgresql-15-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgsentinel` | 1.2.0 | `d12.aarch64` | pigsty | 38.8 KiB | [postgresql-15-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgsentinel` | 1.2.0 | `d12.x86_64` | pigsty | 40.3 KiB | [postgresql-15-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgsentinel` | 1.3.0 | `u22.x86_64` | pgdg | 52.5 KiB | [postgresql-15-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgsentinel` | 1.3.0 | `u22.aarch64` | pgdg | 51.3 KiB | [postgresql-15-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgsentinel` | 1.2.0 | `u22.x86_64` | pigsty | 51.4 KiB | [postgresql-15-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgsentinel` | 1.2.0 | `u22.aarch64` | pigsty | 50.5 KiB | [postgresql-15-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgsentinel` | 1.3.0 | `u24.aarch64` | pgdg | 42.8 KiB | [postgresql-15-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgsentinel` | 1.3.0 | `u24.x86_64` | pgdg | 44.2 KiB | [postgresql-15-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgsentinel` | 1.2.0 | `u24.aarch64` | pigsty | 41.6 KiB | [postgresql-15-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgsentinel` | 1.2.0 | `u24.x86_64` | pigsty | 42.6 KiB | [postgresql-15-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-15-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsentinel_14` | 1.2.0 | `el8.aarch64` | pgdg | 22.7 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_14` | 1.2.0 | `el8.x86_64` | pgdg | 23.7 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_14` | 1.1.0 | `el8.aarch64` | pigsty | 21.0 KiB | [pgsentinel_14-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_14-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_14` | 1.1.0 | `el8.x86_64` | pigsty | 21.9 KiB | [pgsentinel_14-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_14-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_14` | 1.2.0 | `el9.aarch64` | pgdg | 23.3 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsentinel_14-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_14` | 1.2.0 | `el9.x86_64` | pgdg | 24.1 KiB | [pgsentinel_14-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsentinel_14-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_14` | 1.1.0 | `el9.x86_64` | pigsty | 22.6 KiB | [pgsentinel_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_14` | 1.1.0 | `el9.aarch64` | pigsty | 21.8 KiB | [pgsentinel_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgsentinel` | 1.3.0 | `d12.x86_64` | pgdg | 43.9 KiB | [postgresql-14-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgsentinel` | 1.3.0 | `d12.aarch64` | pgdg | 42.3 KiB | [postgresql-14-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgsentinel` | 1.2.0 | `d12.x86_64` | pigsty | 40.1 KiB | [postgresql-14-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgsentinel` | 1.2.0 | `d12.aarch64` | pigsty | 38.5 KiB | [postgresql-14-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgsentinel` | 1.3.0 | `u22.x86_64` | pgdg | 50.3 KiB | [postgresql-14-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgsentinel` | 1.3.0 | `u22.aarch64` | pgdg | 49.2 KiB | [postgresql-14-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgsentinel` | 1.2.0 | `u22.aarch64` | pigsty | 48.1 KiB | [postgresql-14-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgsentinel` | 1.2.0 | `u22.x86_64` | pigsty | 48.9 KiB | [postgresql-14-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgsentinel` | 1.3.0 | `u24.aarch64` | pgdg | 42.5 KiB | [postgresql-14-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgsentinel` | 1.3.0 | `u24.x86_64` | pgdg | 43.9 KiB | [postgresql-14-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.3.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgsentinel` | 1.2.0 | `u24.x86_64` | pigsty | 42.4 KiB | [postgresql-14-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgsentinel` | 1.2.0 | `u24.aarch64` | pigsty | 41.2 KiB | [postgresql-14-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-14-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsentinel_13` | 1.2.0 | `el8.aarch64` | pgdg | 22.7 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsentinel_13-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pgsentinel_13` | 1.2.0 | `el8.x86_64` | pgdg | 23.6 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsentinel_13-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pgsentinel_13` | 1.1.0 | `el8.aarch64` | pigsty | 21.0 KiB | [pgsentinel_13-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsentinel_13-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgsentinel_13` | 1.1.0 | `el8.x86_64` | pigsty | 21.8 KiB | [pgsentinel_13-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsentinel_13-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgsentinel_13` | 1.2.0 | `el9.aarch64` | pgdg | 23.3 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsentinel_13-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pgsentinel_13` | 1.2.0 | `el9.x86_64` | pgdg | 24.1 KiB | [pgsentinel_13-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsentinel_13-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pgsentinel_13` | 1.1.0 | `el9.x86_64` | pigsty | 22.6 KiB | [pgsentinel_13-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsentinel_13-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgsentinel_13` | 1.1.0 | `el9.aarch64` | pigsty | 21.8 KiB | [pgsentinel_13-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsentinel_13-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pgsentinel` | 1.2.0 | `d12.x86_64` | pigsty | 39.9 KiB | [postgresql-13-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgsentinel` | 1.2.0 | `d12.aarch64` | pigsty | 38.4 KiB | [postgresql-13-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgsentinel` | 1.2.0 | `u22.x86_64` | pigsty | 48.3 KiB | [postgresql-13-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgsentinel` | 1.2.0 | `u22.aarch64` | pigsty | 47.6 KiB | [postgresql-13-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgsentinel` | 1.2.0 | `u24.x86_64` | pigsty | 42.1 KiB | [postgresql-13-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgsentinel` | 1.2.0 | `u24.aarch64` | pigsty | 41.1 KiB | [postgresql-13-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsentinel/postgresql-13-pgsentinel_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsentinel/pgsentinel" title="Repository" icon="github" subtitle="github.com/pgsentinel/pgsentinel" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsentinel-1.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgsentinel; # get pgsentinel source code
pig build dep pgsentinel; # install build dependencies
pig build pkg pgsentinel; # build extension rpm or deb
pig build ext pgsentinel; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgsentinel; # install by extension name, for the current active PG version
pig ext install pgsentinel; # install via package alias, for the active PG version
pig ext install pgsentinel -v 18;   # install for PG 18
pig ext install pgsentinel -v 17;   # install for PG 17
pig ext install pgsentinel -v 16;   # install for PG 16
pig ext install pgsentinel -v 15;   # install for PG 15
pig ext install pgsentinel -v 14;   # install for PG 14
pig ext install pgsentinel -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgsentinel;
```

