---
title: "pg_fkpart"
linkTitle: "pg_fkpart"
description: "Table partitioning by foreign key utility"
weight: 2500
categories: ["Olap"]
width: full
---

Table partitioning by foreign key utility

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2500** | {{< badge content="pg_fkpart" link="https://github.com/lemoineat/pg_fkpart" >}} | {{< ext "pg_fkpart" "pg_fkpart" >}} | `1.7.0` | {{< category "OLAP" >}} | {{< license "GPL-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "citus" >}} {{< ext "pg_partman" >}} {{< ext "timescaledb" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "btree_gist" >}} {{< ext "emaj" >}} {{< ext "table_version" >}} |

> [!Note] no pg13,12 on el8 pgdg repo


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_fkpart" >}} | `1.7.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_fkpart_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_fkpart" >}} | `1.7.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-fkpart" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-fkpart` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_fkpart_18" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_fkpart_18-1.7.0-6PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_fkpart_17" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_fkpart_17-1.7.0-6PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_fkpart_16" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_fkpart_16-1.7.0-4.rhel8.noarch.rpm" >}} | {{< pkg "pg_fkpart_15" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_fkpart_15-1.7.0-3.rhel8.noarch.rpm" >}} | {{< pkg "pg_fkpart_14" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_fkpart_14-1.7.0-3.rhel8.noarch.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_fkpart_18" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_fkpart_18-1.7.0-6PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_fkpart_17" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_fkpart_17-1.7.0-6PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_fkpart_16" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_fkpart_16-1.7.0-4.rhel8.noarch.rpm" >}} | {{< pkg "pg_fkpart_15" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_fkpart_15-1.7.0-3.rhel8.noarch.rpm" >}} | {{< pkg "pg_fkpart_14" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_fkpart_14-1.7.0-3.rhel8.noarch.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_fkpart_18" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_fkpart_18-1.7.0-6PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_fkpart_17" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_fkpart_17-1.7.0-6PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_fkpart_16" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_fkpart_16-1.7.0-4.rhel9.noarch.rpm" >}} | {{< pkg "pg_fkpart_15" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_fkpart_15-1.7.0-3.rhel9.noarch.rpm" >}} | {{< pkg "pg_fkpart_14" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_fkpart_14-1.7.0-3.rhel9.noarch.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_fkpart_18" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_fkpart_18-1.7.0-6PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_fkpart_17" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_fkpart_17-1.7.0-6PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_fkpart_16" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_fkpart_16-1.7.0-4.rhel9.noarch.rpm" >}} | {{< pkg "pg_fkpart_15" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_fkpart_15-1.7.0-3.rhel9.noarch.rpm" >}} | {{< pkg "pg_fkpart_14" "1.7.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_fkpart_14-1.7.0-3.rhel9.noarch.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-fkpart" >}}     | {{< pkg "postgresql-17-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-fkpart" >}}     | {{< pkg "postgresql-17-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-fkpart" >}}     | {{< pkg "postgresql-17-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-fkpart" >}}     | {{< pkg "postgresql-17-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-fkpart" >}}     | {{< pkg "postgresql-17-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-fkpart" >}}     | {{< pkg "postgresql-17-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-fkpart" "1.7.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fkpart_18` | 1.7.0 | `el8.x86_64` | pgdg | 24.4 KiB | [pg_fkpart_18-1.7.0-6PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_fkpart_18-1.7.0-6PGDG.rhel8.noarch.rpm) |
| `pg_fkpart_18` | 1.7.0 | `el8.aarch64` | pgdg | 24.3 KiB | [pg_fkpart_18-1.7.0-6PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_fkpart_18-1.7.0-6PGDG.rhel8.noarch.rpm) |
| `pg_fkpart_18` | 1.7.0 | `el9.x86_64` | pgdg | 22.9 KiB | [pg_fkpart_18-1.7.0-6PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_fkpart_18-1.7.0-6PGDG.rhel9.noarch.rpm) |
| `pg_fkpart_18` | 1.7.0 | `el9.aarch64` | pgdg | 22.8 KiB | [pg_fkpart_18-1.7.0-6PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_fkpart_18-1.7.0-6PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fkpart_17` | 1.7.0 | `el8.x86_64` | pigsty | 23.1 KiB | [pg_fkpart_17-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fkpart_17-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fkpart_17` | 1.7.0 | `el8.aarch64` | pgdg | 24.3 KiB | [pg_fkpart_17-1.7.0-6PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_fkpart_17-1.7.0-6PGDG.rhel8.noarch.rpm) |
| `pg_fkpart_17` | 1.7.0 | `el8.aarch64` | pigsty | 23.1 KiB | [pg_fkpart_17-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fkpart_17-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fkpart_17` | 1.7.0 | `el8.x86_64` | pgdg | 24.4 KiB | [pg_fkpart_17-1.7.0-6PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_fkpart_17-1.7.0-6PGDG.rhel8.noarch.rpm) |
| `pg_fkpart_17` | 1.7.0 | `el9.x86_64` | pigsty | 22.6 KiB | [pg_fkpart_17-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fkpart_17-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fkpart_17` | 1.7.0 | `el9.aarch64` | pigsty | 22.6 KiB | [pg_fkpart_17-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fkpart_17-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fkpart_17` | 1.7.0 | `el9.x86_64` | pgdg | 22.9 KiB | [pg_fkpart_17-1.7.0-6PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_fkpart_17-1.7.0-6PGDG.rhel9.noarch.rpm) |
| `pg_fkpart_17` | 1.7.0 | `el9.aarch64` | pgdg | 22.8 KiB | [pg_fkpart_17-1.7.0-6PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_fkpart_17-1.7.0-6PGDG.rhel9.noarch.rpm) |
| `postgresql-17-pg-fkpart` | 1.7.0 | `d12.x86_64` | pigsty | 15.4 KiB | [postgresql-17-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-fkpart` | 1.7.0 | `d12.aarch64` | pigsty | 15.4 KiB | [postgresql-17-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-fkpart` | 1.7.0 | `u22.x86_64` | pigsty | 15.7 KiB | [postgresql-17-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-fkpart` | 1.7.0 | `u22.aarch64` | pigsty | 15.7 KiB | [postgresql-17-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-fkpart` | 1.7.0 | `u24.x86_64` | pigsty | 15.6 KiB | [postgresql-17-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-fkpart` | 1.7.0 | `u24.aarch64` | pigsty | 15.6 KiB | [postgresql-17-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-17-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fkpart_16` | 1.7.0 | `el8.aarch64` | pgdg | 24.1 KiB | [pg_fkpart_16-1.7.0-4.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_fkpart_16-1.7.0-4.rhel8.noarch.rpm) |
| `pg_fkpart_16` | 1.7.0 | `el8.aarch64` | pigsty | 23.1 KiB | [pg_fkpart_16-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fkpart_16-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fkpart_16` | 1.7.0 | `el8.x86_64` | pgdg | 24.2 KiB | [pg_fkpart_16-1.7.0-4.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_fkpart_16-1.7.0-4.rhel8.noarch.rpm) |
| `pg_fkpart_16` | 1.7.0 | `el8.x86_64` | pigsty | 23.1 KiB | [pg_fkpart_16-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fkpart_16-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fkpart_16` | 1.7.0 | `el9.x86_64` | pigsty | 22.6 KiB | [pg_fkpart_16-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fkpart_16-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fkpart_16` | 1.7.0 | `el9.x86_64` | pgdg | 22.8 KiB | [pg_fkpart_16-1.7.0-4.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_fkpart_16-1.7.0-4.rhel9.noarch.rpm) |
| `pg_fkpart_16` | 1.7.0 | `el9.aarch64` | pgdg | 22.6 KiB | [pg_fkpart_16-1.7.0-4.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_fkpart_16-1.7.0-4.rhel9.noarch.rpm) |
| `pg_fkpart_16` | 1.7.0 | `el9.aarch64` | pigsty | 22.6 KiB | [pg_fkpart_16-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fkpart_16-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-fkpart` | 1.7.0 | `d12.x86_64` | pigsty | 15.4 KiB | [postgresql-16-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-fkpart` | 1.7.0 | `d12.aarch64` | pigsty | 15.4 KiB | [postgresql-16-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-fkpart` | 1.7.0 | `u22.aarch64` | pigsty | 15.7 KiB | [postgresql-16-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-fkpart` | 1.7.0 | `u22.x86_64` | pigsty | 15.7 KiB | [postgresql-16-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-fkpart` | 1.7.0 | `u24.aarch64` | pigsty | 15.6 KiB | [postgresql-16-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-fkpart` | 1.7.0 | `u24.x86_64` | pigsty | 15.6 KiB | [postgresql-16-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-16-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fkpart_15` | 1.7.0 | `el8.x86_64` | pgdg | 24.1 KiB | [pg_fkpart_15-1.7.0-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_fkpart_15-1.7.0-3.rhel8.noarch.rpm) |
| `pg_fkpart_15` | 1.7.0 | `el8.x86_64` | pigsty | 23.1 KiB | [pg_fkpart_15-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fkpart_15-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fkpart_15` | 1.7.0 | `el8.aarch64` | pigsty | 23.1 KiB | [pg_fkpart_15-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fkpart_15-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fkpart_15` | 1.7.0 | `el8.aarch64` | pgdg | 24.0 KiB | [pg_fkpart_15-1.7.0-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_fkpart_15-1.7.0-3.rhel8.noarch.rpm) |
| `pg_fkpart_15` | 1.7.0 | `el9.aarch64` | pgdg | 22.9 KiB | [pg_fkpart_15-1.7.0-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_fkpart_15-1.7.0-3.rhel9.noarch.rpm) |
| `pg_fkpart_15` | 1.7.0 | `el9.aarch64` | pigsty | 22.6 KiB | [pg_fkpart_15-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fkpart_15-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fkpart_15` | 1.7.0 | `el9.x86_64` | pgdg | 23.1 KiB | [pg_fkpart_15-1.7.0-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_fkpart_15-1.7.0-3.rhel9.noarch.rpm) |
| `pg_fkpart_15` | 1.7.0 | `el9.x86_64` | pigsty | 22.6 KiB | [pg_fkpart_15-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fkpart_15-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-pg-fkpart` | 1.7.0 | `d12.aarch64` | pigsty | 15.4 KiB | [postgresql-15-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-fkpart` | 1.7.0 | `d12.x86_64` | pigsty | 15.4 KiB | [postgresql-15-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-fkpart` | 1.7.0 | `u22.x86_64` | pigsty | 15.7 KiB | [postgresql-15-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-fkpart` | 1.7.0 | `u22.aarch64` | pigsty | 15.7 KiB | [postgresql-15-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-fkpart` | 1.7.0 | `u24.x86_64` | pigsty | 15.6 KiB | [postgresql-15-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-fkpart` | 1.7.0 | `u24.aarch64` | pigsty | 15.6 KiB | [postgresql-15-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-15-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fkpart_14` | 1.7.0 | `el8.x86_64` | pgdg | 24.1 KiB | [pg_fkpart_14-1.7.0-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_fkpart_14-1.7.0-3.rhel8.noarch.rpm) |
| `pg_fkpart_14` | 1.7.0 | `el8.aarch64` | pigsty | 23.1 KiB | [pg_fkpart_14-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fkpart_14-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fkpart_14` | 1.7.0 | `el8.aarch64` | pgdg | 24.0 KiB | [pg_fkpart_14-1.7.0-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_fkpart_14-1.7.0-3.rhel8.noarch.rpm) |
| `pg_fkpart_14` | 1.7.0 | `el8.x86_64` | pigsty | 23.1 KiB | [pg_fkpart_14-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fkpart_14-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fkpart_14` | 1.7.0 | `el9.aarch64` | pgdg | 22.9 KiB | [pg_fkpart_14-1.7.0-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_fkpart_14-1.7.0-3.rhel9.noarch.rpm) |
| `pg_fkpart_14` | 1.7.0 | `el9.x86_64` | pigsty | 22.6 KiB | [pg_fkpart_14-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fkpart_14-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fkpart_14` | 1.7.0 | `el9.x86_64` | pgdg | 23.1 KiB | [pg_fkpart_14-1.7.0-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_fkpart_14-1.7.0-3.rhel9.noarch.rpm) |
| `pg_fkpart_14` | 1.7.0 | `el9.aarch64` | pigsty | 22.6 KiB | [pg_fkpart_14-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fkpart_14-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-fkpart` | 1.7.0 | `d12.aarch64` | pigsty | 15.4 KiB | [postgresql-14-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-fkpart` | 1.7.0 | `d12.x86_64` | pigsty | 15.4 KiB | [postgresql-14-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-fkpart` | 1.7.0 | `u22.aarch64` | pigsty | 15.7 KiB | [postgresql-14-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-fkpart` | 1.7.0 | `u22.x86_64` | pigsty | 15.7 KiB | [postgresql-14-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-fkpart` | 1.7.0 | `u24.x86_64` | pigsty | 15.6 KiB | [postgresql-14-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-fkpart` | 1.7.0 | `u24.aarch64` | pigsty | 15.6 KiB | [postgresql-14-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-14-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_fkpart_13` | 1.7.0 | `el8.aarch64` | pgdg | 24.0 KiB | [pg_fkpart_13-1.7.0-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_fkpart_13-1.7.0-3.rhel8.noarch.rpm) |
| `pg_fkpart_13` | 1.7.0 | `el8.aarch64` | pigsty | 23.1 KiB | [pg_fkpart_13-1.7.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_fkpart_13-1.7.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_fkpart_13` | 1.7.0 | `el8.x86_64` | pigsty | 23.1 KiB | [pg_fkpart_13-1.7.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_fkpart_13-1.7.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_fkpart_13` | 1.7.0 | `el9.x86_64` | pigsty | 22.6 KiB | [pg_fkpart_13-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_fkpart_13-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_fkpart_13` | 1.7.0 | `el9.aarch64` | pigsty | 22.6 KiB | [pg_fkpart_13-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_fkpart_13-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_fkpart_13` | 1.7.0 | `el9.x86_64` | pgdg | 23.1 KiB | [pg_fkpart_13-1.7.0-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_fkpart_13-1.7.0-3.rhel9.noarch.rpm) |
| `pg_fkpart_13` | 1.7.0 | `el9.aarch64` | pgdg | 22.9 KiB | [pg_fkpart_13-1.7.0-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_fkpart_13-1.7.0-3.rhel9.noarch.rpm) |
| `postgresql-13-pg-fkpart` | 1.7.0 | `d12.x86_64` | pigsty | 15.4 KiB | [postgresql-13-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-13-pg-fkpart_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-fkpart` | 1.7.0 | `d12.aarch64` | pigsty | 15.4 KiB | [postgresql-13-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-fkpart/postgresql-13-pg-fkpart_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-fkpart` | 1.7.0 | `u22.aarch64` | pigsty | 15.7 KiB | [postgresql-13-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-13-pg-fkpart_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-fkpart` | 1.7.0 | `u22.x86_64` | pigsty | 15.7 KiB | [postgresql-13-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-fkpart/postgresql-13-pg-fkpart_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-fkpart` | 1.7.0 | `u24.x86_64` | pigsty | 15.6 KiB | [postgresql-13-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-13-pg-fkpart_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-fkpart` | 1.7.0 | `u24.aarch64` | pigsty | 15.6 KiB | [postgresql-13-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-fkpart/postgresql-13-pg-fkpart_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lemoineat/pg_fkpart" title="Repository" icon="github" subtitle="github.com/lemoineat/pg_fkpart" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_fkpart-1.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_fkpart; # get pg_fkpart source code
pig build dep pg_fkpart; # install build dependencies
pig build pkg pg_fkpart; # build extension rpm or deb
pig build ext pg_fkpart; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_fkpart; # install by extension name, for the current active PG version
pig ext install pg_fkpart; # install via package alias, for the active PG version
pig ext install pg_fkpart -v 18;   # install for PG 18
pig ext install pg_fkpart -v 17;   # install for PG 17
pig ext install pg_fkpart -v 16;   # install for PG 16
pig ext install pg_fkpart -v 15;   # install for PG 15
pig ext install pg_fkpart -v 14;   # install for PG 14
pig ext install pg_fkpart -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_fkpart CASCADE SCHEMA pgfkpart;
```

