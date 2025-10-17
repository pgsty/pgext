---
title: "pg_bulkload"
linkTitle: "pg_bulkload"
description: "pg_bulkload is a high speed data loading utility for PostgreSQL"
weight: 9830
categories: ["Etl"]
width: full
---

pg_bulkload is a high speed data loading utility for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9830** | {{< badge content="pg_bulkload" link="https://github.com/ossc-db/pg_bulkload" >}} | {{< ext "pg_bulkload" "pg_bulkload" >}} | `3.1.22` | {{< category "ETL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--bs-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "file_fdw" >}} {{< ext "aws_s3" >}} {{< ext "db_migrator" >}} {{< ext "pg_fact_loader" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "pglogical" >}} |

> [!Note] pg18 breaks


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_bulkload" >}} | `3.1.22` | {{< badge content="18" color="red" alt="pg_bulkload_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_bulkload_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_bulkload" >}} | `3.1.22` | {{< badge content="18" color="red" alt="postgresql-18-pg-bulkload" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-bulkload` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_bulkload_18" >}}     | {{< pkg "pg_bulkload_17" "3.1.22" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_17-3.1.22-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_bulkload_16" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_bulkload_15" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_bulkload_14" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.21-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_bulkload_18" >}}     | {{< pkg "pg_bulkload_17" "3.1.22" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_17-3.1.22-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_bulkload_16" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_bulkload_15" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_bulkload_14" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.21-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_bulkload_18" >}}     | {{< pkg "pg_bulkload_17" "3.1.22" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_17-3.1.22-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_bulkload_16" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_bulkload_15" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_bulkload_14" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bulkload_14-3.1.21-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_bulkload_18" >}}     | {{< pkg "pg_bulkload_17" "3.1.22" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_17-3.1.22-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_bulkload_16" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_bulkload_15" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_bulkload_14" "3.1.21" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.21-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-bulkload" >}}     | {{< pkg "postgresql-17-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-bulkload" >}}     | {{< pkg "postgresql-17-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-bulkload" >}}     | {{< pkg "postgresql-17-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-bulkload" >}}     | {{< pkg "postgresql-17-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-bulkload" >}}     | {{< pkg "postgresql-17-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-bulkload" >}}     | {{< pkg "postgresql-17-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-bulkload" "3.1.22" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bulkload_17` | 3.1.22 | `el8.aarch64` | pigsty | 60.0 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_17-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_17` | 3.1.22 | `el8.x86_64` | pigsty | 64.9 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_17-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_17` | 3.1.22 | `el9.aarch64` | pigsty | 58.4 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_17-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_17` | 3.1.22 | `el9.x86_64` | pigsty | 61.0 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_17-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-bulkload` | 3.1.22 | `d12.x86_64` | pigsty | 289.2 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bulkload` | 3.1.22 | `d12.aarch64` | pigsty | 281.5 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bulkload` | 3.1.22 | `u22.aarch64` | pigsty | 291.7 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bulkload` | 3.1.22 | `u22.x86_64` | pigsty | 296.8 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bulkload` | 3.1.22 | `u24.x86_64` | pigsty | 244.8 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bulkload` | 3.1.22 | `u24.aarch64` | pigsty | 240.6 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bulkload_16` | 3.1.22 | `el8.aarch64` | pigsty | 60.5 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_16-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_16` | 3.1.22 | `el8.x86_64` | pigsty | 65.2 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_16-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_16` | 3.1.21 | `el8.aarch64` | pgdg | 71.4 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_16` | 3.1.21 | `el8.x86_64` | pgdg | 76.4 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_16` | 3.1.22 | `el9.aarch64` | pigsty | 58.8 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_16-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_16` | 3.1.22 | `el9.x86_64` | pigsty | 61.3 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_16-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_16` | 3.1.21 | `el9.x86_64` | pgdg | 69.7 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_16` | 3.1.21 | `el9.aarch64` | pgdg | 66.9 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pg-bulkload` | 3.1.22 | `d12.aarch64` | pigsty | 279.5 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bulkload` | 3.1.22 | `d12.x86_64` | pigsty | 287.9 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bulkload` | 3.1.22 | `u22.x86_64` | pigsty | 295.2 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bulkload` | 3.1.22 | `u22.aarch64` | pigsty | 289.4 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bulkload` | 3.1.22 | `u24.aarch64` | pigsty | 241.9 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-bulkload` | 3.1.22 | `u24.x86_64` | pigsty | 245.7 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bulkload_15` | 3.1.22 | `el8.x86_64` | pigsty | 65.1 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_15-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_15` | 3.1.22 | `el8.aarch64` | pigsty | 60.3 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_15-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_15` | 3.1.21 | `el8.x86_64` | pgdg | 76.2 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_15` | 3.1.21 | `el8.aarch64` | pgdg | 71.1 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_15` | 3.1.20 | `el8.aarch64` | pgdg | 71.6 KiB | [pg_bulkload_15-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_15` | 3.1.20 | `el8.x86_64` | pgdg | 76.7 KiB | [pg_bulkload_15-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_15` | 3.1.22 | `el9.x86_64` | pigsty | 61.2 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_15-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_15` | 3.1.22 | `el9.aarch64` | pigsty | 58.6 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_15-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_15` | 3.1.21 | `el9.x86_64` | pgdg | 69.6 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_15` | 3.1.21 | `el9.aarch64` | pgdg | 66.8 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_15` | 3.1.20 | `el9.aarch64` | pgdg | 67.6 KiB | [pg_bulkload_15-3.1.20-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bulkload_15-3.1.20-1.rhel9.aarch64.rpm) |
| `postgresql-15-pg-bulkload` | 3.1.22 | `d12.x86_64` | pigsty | 288.5 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bulkload` | 3.1.22 | `d12.aarch64` | pigsty | 280.2 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bulkload` | 3.1.22 | `u22.aarch64` | pigsty | 290.5 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bulkload` | 3.1.22 | `u22.x86_64` | pigsty | 296.1 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bulkload` | 3.1.22 | `u24.aarch64` | pigsty | 242.4 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-bulkload` | 3.1.22 | `u24.x86_64` | pigsty | 246.3 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bulkload_14` | 3.1.22 | `el8.x86_64` | pigsty | 65.0 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_14-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_14` | 3.1.22 | `el8.aarch64` | pigsty | 60.2 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_14-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_14` | 3.1.21 | `el8.x86_64` | pgdg | 76.1 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | 3.1.21 | `el8.aarch64` | pgdg | 71.1 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | 3.1.20 | `el8.aarch64` | pgdg | 71.6 KiB | [pg_bulkload_14-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | 3.1.20 | `el8.x86_64` | pgdg | 76.6 KiB | [pg_bulkload_14-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | 3.1.19 | `el8.aarch64` | pgdg | 71.4 KiB | [pg_bulkload_14-3.1.19-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.19-2.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | 3.1.19 | `el8.x86_64` | pgdg | 265.7 KiB | [pg_bulkload_14-3.1.19-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.19-1.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | 3.1.22 | `el9.x86_64` | pigsty | 61.0 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_14-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_14` | 3.1.22 | `el9.aarch64` | pigsty | 58.5 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_14-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_14` | 3.1.21 | `el9.aarch64` | pgdg | 66.8 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_14` | 3.1.21 | `el9.x86_64` | pgdg | 69.7 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bulkload_14-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_14` | 3.1.20 | `el9.aarch64` | pgdg | 67.5 KiB | [pg_bulkload_14-3.1.20-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.20-1.rhel9.aarch64.rpm) |
| `pg_bulkload_14` | 3.1.19 | `el9.aarch64` | pgdg | 67.6 KiB | [pg_bulkload_14-3.1.19-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.19-2.rhel9.aarch64.rpm) |
| `postgresql-14-pg-bulkload` | 3.1.22 | `d12.aarch64` | pigsty | 280.1 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-bulkload` | 3.1.22 | `d12.x86_64` | pigsty | 288.0 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-bulkload` | 3.1.22 | `u22.aarch64` | pigsty | 289.7 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-bulkload` | 3.1.22 | `u22.x86_64` | pigsty | 294.9 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-bulkload` | 3.1.22 | `u24.x86_64` | pigsty | 244.9 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-bulkload` | 3.1.22 | `u24.aarch64` | pigsty | 241.1 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_bulkload_13` | 3.1.22 | `el8.aarch64` | pigsty | 59.8 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_13-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_13` | 3.1.22 | `el8.x86_64` | pigsty | 64.1 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_13-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_13` | 3.1.21 | `el8.aarch64` | pgdg | 70.7 KiB | [pg_bulkload_13-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bulkload_13-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_13` | 3.1.21 | `el8.x86_64` | pgdg | 74.5 KiB | [pg_bulkload_13-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | 3.1.20 | `el8.x86_64` | pgdg | 74.9 KiB | [pg_bulkload_13-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | 3.1.20 | `el8.aarch64` | pgdg | 71.2 KiB | [pg_bulkload_13-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bulkload_13-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_13` | 3.1.19 | `el8.aarch64` | pgdg | 71.0 KiB | [pg_bulkload_13-3.1.19-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bulkload_13-3.1.19-2.rhel8.aarch64.rpm) |
| `pg_bulkload_13` | 3.1.18 | `el8.x86_64` | pgdg | 260.3 KiB | [pg_bulkload_13-3.1.18-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.18-1.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | 3.1.18 | `el8.x86_64` | pgdg | 260.4 KiB | [pg_bulkload_13-3.1.18-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.18-2.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | 3.1.17 | `el8.x86_64` | pgdg | 522.1 KiB | [pg_bulkload_13-3.1.17-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.17-1.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | 3.1.22 | `el9.aarch64` | pigsty | 58.5 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_13-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_13` | 3.1.22 | `el9.x86_64` | pigsty | 60.8 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_13-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_13` | 3.1.21 | `el9.aarch64` | pgdg | 66.5 KiB | [pg_bulkload_13-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bulkload_13-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_13` | 3.1.21 | `el9.x86_64` | pgdg | 69.3 KiB | [pg_bulkload_13-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_bulkload_13-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_13` | 3.1.20 | `el9.aarch64` | pgdg | 67.3 KiB | [pg_bulkload_13-3.1.20-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bulkload_13-3.1.20-1.rhel9.aarch64.rpm) |
| `pg_bulkload_13` | 3.1.19 | `el9.aarch64` | pgdg | 67.4 KiB | [pg_bulkload_13-3.1.19-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bulkload_13-3.1.19-2.rhel9.aarch64.rpm) |
| `postgresql-13-pg-bulkload` | 3.1.22 | `d12.x86_64` | pigsty | 282.8 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-bulkload` | 3.1.22 | `d12.aarch64` | pigsty | 274.5 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-bulkload` | 3.1.22 | `u22.aarch64` | pigsty | 284.1 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-bulkload` | 3.1.22 | `u22.x86_64` | pigsty | 289.3 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-bulkload` | 3.1.22 | `u24.aarch64` | pigsty | 238.5 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-bulkload` | 3.1.22 | `u24.x86_64` | pigsty | 242.8 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ossc-db/pg_bulkload" title="Repository" icon="github" subtitle="github.com/ossc-db/pg_bulkload" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_bulkload-VERSION3_1_22.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_bulkload; # get pg_bulkload source code
pig build dep pg_bulkload; # install build dependencies
pig build pkg pg_bulkload; # build extension rpm or deb
pig build ext pg_bulkload; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_bulkload; # install by extension name, for the current active PG version
pig ext install pg_bulkload; # install via package alias, for the active PG version
pig ext install pg_bulkload -v 17;   # install for PG 17
pig ext install pg_bulkload -v 16;   # install for PG 16
pig ext install pg_bulkload -v 15;   # install for PG 15
pig ext install pg_bulkload -v 14;   # install for PG 14
pig ext install pg_bulkload -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_bulkload;
```

