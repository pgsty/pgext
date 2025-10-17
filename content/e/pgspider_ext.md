---
title: "pgspider_ext"
linkTitle: "pgspider_ext"
description: "foreign-data wrapper for remote PGSpider servers"
weight: 8540
categories: ["Fdw"]
width: full
---

foreign-data wrapper for remote PGSpider servers

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8540** | {{< badge content="pgspider_ext" link="https://github.com/pgspider/pgspider_ext" >}} | {{< ext "pgspider_ext" "pgspider_ext" >}} | `1.3.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plproxy" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "postgres_fdw" >}} {{< ext "citus" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "mongo_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgspider_ext" >}} | `1.3.0` | {{< badge content="18" color="red" alt="pgspider_ext_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="pgspider_ext_14*" >}} | `pgspider_ext_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgspider_ext" >}} | `1.3.0` | {{< badge content="18" color="red" alt="postgresql-18-pgspider-ext" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="postgresql-14-pgspider-ext" >}} | `postgresql-$v-pgspider-ext` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgspider_ext_18" >}}     | {{< pkg "pgspider_ext_17" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgspider_ext_17-1.3.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgspider_ext_16" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgspider_ext_16-1.3.0-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgspider_ext_15" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgspider_ext_15-1.3.0-1PGDG.rhel8.x86_64.rpm" >}} |    {{< pkg "pgspider_ext_14" >}}     |
|    `el8.aarch64`    |    {{< pkg "pgspider_ext_18" >}}     | {{< pkg "pgspider_ext_17" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgspider_ext_17-1.3.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgspider_ext_16" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgspider_ext_16-1.3.0-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgspider_ext_15" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgspider_ext_15-1.3.0-1PGDG.rhel8.aarch64.rpm" >}} |    {{< pkg "pgspider_ext_14" >}}     |
|    `el9.x86_64`    |    {{< pkg "pgspider_ext_18" >}}     | {{< pkg "pgspider_ext_17" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgspider_ext_17-1.3.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgspider_ext_16" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgspider_ext_16-1.3.0-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgspider_ext_15" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgspider_ext_15-1.3.0-1PGDG.rhel9.x86_64.rpm" >}} |    {{< pkg "pgspider_ext_14" >}}     |
|    `el9.aarch64`    |    {{< pkg "pgspider_ext_18" >}}     | {{< pkg "pgspider_ext_17" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgspider_ext_17-1.3.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgspider_ext_16" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgspider_ext_16-1.3.0-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgspider_ext_15" "1.3.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgspider_ext_15-1.3.0-1PGDG.rhel9.aarch64.rpm" >}} |    {{< pkg "pgspider_ext_14" >}}     |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgspider-ext" >}}     | {{< pkg "postgresql-17-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb" >}} |    {{< pkg "postgresql-14-pgspider-ext" >}}     |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgspider-ext" >}}     | {{< pkg "postgresql-17-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb" >}} |    {{< pkg "postgresql-14-pgspider-ext" >}}     |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgspider-ext" >}}     | {{< pkg "postgresql-17-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb" >}} |    {{< pkg "postgresql-14-pgspider-ext" >}}     |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgspider-ext" >}}     | {{< pkg "postgresql-17-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb" >}} |    {{< pkg "postgresql-14-pgspider-ext" >}}     |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgspider-ext" >}}     | {{< pkg "postgresql-17-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb" >}} |    {{< pkg "postgresql-14-pgspider-ext" >}}     |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgspider-ext" >}}     | {{< pkg "postgresql-17-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgspider-ext" "1.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb" >}} |    {{< pkg "postgresql-14-pgspider-ext" >}}     |


{{< tabs items="PG17,PG16,PG15" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgspider_ext_17` | 1.3.0 | `el8.aarch64` | pgdg | 27.7 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgspider_ext_17-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgspider_ext_17` | 1.3.0 | `el8.x86_64` | pgdg | 28.6 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgspider_ext_17-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgspider_ext_17` | 1.3.0 | `el9.x86_64` | pgdg | 29.1 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgspider_ext_17-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgspider_ext_17` | 1.3.0 | `el9.aarch64` | pgdg | 28.4 KiB | [pgspider_ext_17-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgspider_ext_17-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pgspider-ext` | 1.3.0 | `d12.x86_64` | pigsty | 48.8 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgspider-ext` | 1.3.0 | `d12.aarch64` | pigsty | 47.1 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgspider-ext` | 1.3.0 | `u22.x86_64` | pigsty | 61.7 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgspider-ext` | 1.3.0 | `u22.aarch64` | pigsty | 60.7 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgspider-ext` | 1.3.0 | `u24.x86_64` | pigsty | 50.9 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgspider-ext` | 1.3.0 | `u24.aarch64` | pigsty | 49.6 KiB | [postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-17-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgspider_ext_16` | 1.3.0 | `el8.x86_64` | pgdg | 28.6 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgspider_ext_16-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgspider_ext_16` | 1.3.0 | `el8.aarch64` | pgdg | 27.8 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgspider_ext_16-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgspider_ext_16` | 1.3.0 | `el9.aarch64` | pgdg | 28.5 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgspider_ext_16-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pgspider_ext_16` | 1.3.0 | `el9.x86_64` | pgdg | 29.2 KiB | [pgspider_ext_16-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgspider_ext_16-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-pgspider-ext` | 1.3.0 | `d12.aarch64` | pigsty | 47.0 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgspider-ext` | 1.3.0 | `d12.x86_64` | pigsty | 48.7 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgspider-ext` | 1.3.0 | `u22.x86_64` | pigsty | 61.3 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgspider-ext` | 1.3.0 | `u22.aarch64` | pigsty | 60.3 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgspider-ext` | 1.3.0 | `u24.aarch64` | pigsty | 49.6 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgspider-ext` | 1.3.0 | `u24.x86_64` | pigsty | 50.8 KiB | [postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-16-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgspider_ext_15` | 1.3.0 | `el8.x86_64` | pgdg | 29.0 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgspider_ext_15-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pgspider_ext_15` | 1.3.0 | `el8.aarch64` | pgdg | 28.0 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgspider_ext_15-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pgspider_ext_15` | 1.3.0 | `el9.x86_64` | pgdg | 29.6 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgspider_ext_15-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pgspider_ext_15` | 1.3.0 | `el9.aarch64` | pgdg | 28.8 KiB | [pgspider_ext_15-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgspider_ext_15-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pgspider-ext` | 1.3.0 | `d12.aarch64` | pigsty | 47.0 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgspider-ext` | 1.3.0 | `d12.x86_64` | pigsty | 48.9 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgspider-ext` | 1.3.0 | `u22.x86_64` | pigsty | 61.3 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgspider-ext` | 1.3.0 | `u22.aarch64` | pigsty | 60.3 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgspider-ext` | 1.3.0 | `u24.aarch64` | pigsty | 49.7 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgspider-ext` | 1.3.0 | `u24.x86_64` | pigsty | 51.0 KiB | [postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgspider-ext/postgresql-15-pgspider-ext_1.3.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgspider/pgspider_ext" title="Repository" icon="github" subtitle="github.com/pgspider/pgspider_ext" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgspider_ext-1.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgspider_ext; # get pgspider_ext source code
pig build dep pgspider_ext; # install build dependencies
pig build pkg pgspider_ext; # build extension rpm or deb
pig build ext pgspider_ext; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgspider_ext; # install by extension name, for the current active PG version
pig ext install pgspider_ext; # install via package alias, for the active PG version
pig ext install pgspider_ext -v 17;   # install for PG 17
pig ext install pgspider_ext -v 16;   # install for PG 16
pig ext install pgspider_ext -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgspider_ext;
```

