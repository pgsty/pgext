---
title: "pg_tle"
linkTitle: "pg_tle"
description: "Trusted Language Extensions for PostgreSQL"
weight: 3000
categories: ["Lang"]
width: full
---

Trusted Language Extensions for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3000** | {{< badge content="pg_tle" link="https://github.com/aws/pg_tle" >}} | {{< ext "pg_tle" "pg_tle" >}} | `1.5.1` | {{< category "LANG" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "pljava" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "plprql" >}} {{< ext "plsh" >}} |

> [!Note] pg18rc1 support with 3c99c51086ae2d1ec7aeb0ecf186a1a29f465d2c


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_tle" >}} | `1.5.1` | {{< badge content="18" color="red" alt="pg_tle_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_tle_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_tle" >}} | `1.5.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-tle" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-tle` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_tle_18" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_tle_18-1.5.2-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_tle_17" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_tle_17-1.5.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_tle_16" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_tle_16-1.5.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_tle_15" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_tle_15-1.5.1-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_tle_14" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_tle_14-1.5.1-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_tle_18" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_tle_18-1.5.2-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_tle_17" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_tle_17-1.5.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_tle_16" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_tle_16-1.5.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_tle_15" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_tle_15-1.5.1-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_tle_14" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_tle_14-1.5.1-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_tle_18" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_tle_18-1.5.2-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_tle_17" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_tle_17-1.5.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_tle_16" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_tle_16-1.5.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_tle_15" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_tle_15-1.5.1-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_tle_14" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_tle_14-1.5.1-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_tle_18" "1.5.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_tle_18-1.5.2-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_tle_17" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_tle_17-1.5.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_tle_16" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_tle_16-1.5.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_tle_15" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_tle_15-1.5.1-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_tle_14" "1.5.1" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_tle_14-1.5.1-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-tle" >}}     | {{< pkg "postgresql-17-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-tle" >}}     | {{< pkg "postgresql-17-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-tle" >}}     | {{< pkg "postgresql-17-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-tle" >}}     | {{< pkg "postgresql-17-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-tle" >}}     | {{< pkg "postgresql-17-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-tle" >}}     | {{< pkg "postgresql-17-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-tle" "1.5.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_tle_18` | 1.5.2 | `el8.aarch64` | pgdg | 65.3 KiB | [pg_tle_18-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_tle_18-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_18` | 1.5.2 | `el8.x86_64` | pgdg | 68.7 KiB | [pg_tle_18-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_tle_18-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_18` | 1.5.1 | `el8.x86_64` | pigsty | 64.1 KiB | [pg_tle_18-1.5.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_18-1.5.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tle_18` | 1.5.1 | `el8.aarch64` | pigsty | 60.7 KiB | [pg_tle_18-1.5.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_18-1.5.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tle_18` | 1.5.2 | `el9.aarch64` | pgdg | 62.1 KiB | [pg_tle_18-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_tle_18-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_18` | 1.5.2 | `el9.x86_64` | pgdg | 65.1 KiB | [pg_tle_18-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_tle_18-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_18` | 1.5.1 | `el9.aarch64` | pigsty | 57.5 KiB | [pg_tle_18-1.5.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_18-1.5.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tle_18` | 1.5.1 | `el9.x86_64` | pigsty | 60.8 KiB | [pg_tle_18-1.5.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_18-1.5.1-2PIGSTY.el9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_tle_17` | 1.5.1 | `el8.x86_64` | pigsty | 64.1 KiB | [pg_tle_17-1.5.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_17-1.5.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tle_17` | 1.5.1 | `el8.aarch64` | pigsty | 60.7 KiB | [pg_tle_17-1.5.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_17-1.5.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tle_17` | 1.5.1 | `el8.x86_64` | pgdg | 68.2 KiB | [pg_tle_17-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_tle_17-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_17` | 1.5.1 | `el8.aarch64` | pgdg | 64.8 KiB | [pg_tle_17-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_tle_17-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_17` | 1.5.0 | `el8.aarch64` | pgdg | 64.7 KiB | [pg_tle_17-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_tle_17-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_17` | 1.5.0 | `el8.x86_64` | pigsty | 64.0 KiB | [pg_tle_17-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_17-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_17` | 1.5.0 | `el8.x86_64` | pgdg | 68.1 KiB | [pg_tle_17-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_tle_17-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_17` | 1.5.0 | `el8.aarch64` | pigsty | 60.6 KiB | [pg_tle_17-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_17-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_17` | 1.2.0 | `el8.aarch64` | pgdg | 59.4 KiB | [pg_tle_17-1.2.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_tle_17-1.2.0-2PGDG.rhel8.aarch64.rpm) |
| `pg_tle_17` | 1.2.0 | `el8.x86_64` | pgdg | 63.1 KiB | [pg_tle_17-1.2.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_tle_17-1.2.0-2PGDG.rhel8.x86_64.rpm) |
| `pg_tle_17` | 1.5.1 | `el9.x86_64` | pigsty | 60.9 KiB | [pg_tle_17-1.5.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_17-1.5.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tle_17` | 1.5.1 | `el9.x86_64` | pgdg | 64.7 KiB | [pg_tle_17-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_tle_17-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_17` | 1.5.1 | `el9.aarch64` | pigsty | 57.5 KiB | [pg_tle_17-1.5.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_17-1.5.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tle_17` | 1.5.1 | `el9.aarch64` | pgdg | 61.8 KiB | [pg_tle_17-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_tle_17-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_17` | 1.5.0 | `el9.x86_64` | pigsty | 60.8 KiB | [pg_tle_17-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_17-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_17` | 1.5.0 | `el9.x86_64` | pgdg | 64.5 KiB | [pg_tle_17-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_tle_17-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_17` | 1.5.0 | `el9.aarch64` | pgdg | 61.7 KiB | [pg_tle_17-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_tle_17-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_17` | 1.5.0 | `el9.aarch64` | pigsty | 57.7 KiB | [pg_tle_17-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_17-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_17` | 1.2.0 | `el9.x86_64` | pgdg | 59.4 KiB | [pg_tle_17-1.2.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_tle_17-1.2.0-2PGDG.rhel9.x86_64.rpm) |
| `pg_tle_17` | 1.2.0 | `el9.aarch64` | pgdg | 56.2 KiB | [pg_tle_17-1.2.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_tle_17-1.2.0-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pg-tle` | 1.5.1 | `d12.x86_64` | pigsty | 159.6 KiB | [postgresql-17-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tle` | 1.5.1 | `d12.aarch64` | pigsty | 154.8 KiB | [postgresql-17-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tle` | 1.5.1 | `u22.x86_64` | pigsty | 183.7 KiB | [postgresql-17-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tle` | 1.5.1 | `u22.aarch64` | pigsty | 179.8 KiB | [postgresql-17-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tle` | 1.5.1 | `u24.aarch64` | pigsty | 159.8 KiB | [postgresql-17-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-tle` | 1.5.1 | `u24.x86_64` | pigsty | 162.3 KiB | [postgresql-17-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-17-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_tle_16` | 1.5.1 | `el8.x86_64` | pgdg | 68.2 KiB | [pg_tle_16-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_tle_16-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_16` | 1.5.1 | `el8.aarch64` | pigsty | 60.7 KiB | [pg_tle_16-1.5.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_16-1.5.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tle_16` | 1.5.1 | `el8.x86_64` | pigsty | 64.1 KiB | [pg_tle_16-1.5.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_16-1.5.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tle_16` | 1.5.1 | `el8.aarch64` | pgdg | 64.8 KiB | [pg_tle_16-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_tle_16-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_16` | 1.5.0 | `el8.aarch64` | pgdg | 64.7 KiB | [pg_tle_16-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_tle_16-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_16` | 1.5.0 | `el8.x86_64` | pgdg | 68.1 KiB | [pg_tle_16-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_tle_16-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_16` | 1.5.0 | `el8.x86_64` | pigsty | 64.0 KiB | [pg_tle_16-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_16-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_16` | 1.5.0 | `el8.aarch64` | pigsty | 60.6 KiB | [pg_tle_16-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_16-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_16` | 1.2.0 | `el8.x86_64` | pgdg | 63.0 KiB | [pg_tle_16-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_tle_16-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_16` | 1.2.0 | `el8.aarch64` | pgdg | 59.3 KiB | [pg_tle_16-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_tle_16-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_16` | 1.5.1 | `el9.aarch64` | pgdg | 61.9 KiB | [pg_tle_16-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_tle_16-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_16` | 1.5.1 | `el9.x86_64` | pigsty | 60.9 KiB | [pg_tle_16-1.5.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_16-1.5.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tle_16` | 1.5.1 | `el9.x86_64` | pgdg | 64.7 KiB | [pg_tle_16-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_tle_16-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_16` | 1.5.1 | `el9.aarch64` | pigsty | 57.6 KiB | [pg_tle_16-1.5.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_16-1.5.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tle_16` | 1.5.0 | `el9.aarch64` | pgdg | 61.7 KiB | [pg_tle_16-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_tle_16-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_16` | 1.5.0 | `el9.aarch64` | pigsty | 57.7 KiB | [pg_tle_16-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_16-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_16` | 1.5.0 | `el9.x86_64` | pgdg | 64.6 KiB | [pg_tle_16-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_tle_16-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_16` | 1.5.0 | `el9.x86_64` | pigsty | 60.8 KiB | [pg_tle_16-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_16-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_16` | 1.2.0 | `el9.aarch64` | pgdg | 56.2 KiB | [pg_tle_16-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_tle_16-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_16` | 1.2.0 | `el9.x86_64` | pgdg | 59.2 KiB | [pg_tle_16-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_tle_16-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-pg-tle` | 1.5.1 | `d12.aarch64` | pigsty | 155.1 KiB | [postgresql-16-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tle` | 1.5.1 | `d12.x86_64` | pigsty | 159.6 KiB | [postgresql-16-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tle` | 1.5.1 | `u22.aarch64` | pigsty | 179.7 KiB | [postgresql-16-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tle` | 1.5.1 | `u22.x86_64` | pigsty | 183.5 KiB | [postgresql-16-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tle` | 1.5.1 | `u24.aarch64` | pigsty | 159.9 KiB | [postgresql-16-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-tle` | 1.5.1 | `u24.x86_64` | pigsty | 162.4 KiB | [postgresql-16-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-16-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_tle_15` | 1.5.1 | `el8.aarch64` | pgdg | 65.5 KiB | [pg_tle_15-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_tle_15-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_15` | 1.5.1 | `el8.x86_64` | pigsty | 65.0 KiB | [pg_tle_15-1.5.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_15-1.5.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tle_15` | 1.5.1 | `el8.aarch64` | pigsty | 61.5 KiB | [pg_tle_15-1.5.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_15-1.5.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tle_15` | 1.5.1 | `el8.x86_64` | pgdg | 69.1 KiB | [pg_tle_15-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_tle_15-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_15` | 1.5.0 | `el8.aarch64` | pgdg | 65.4 KiB | [pg_tle_15-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_tle_15-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_15` | 1.5.0 | `el8.x86_64` | pgdg | 69.0 KiB | [pg_tle_15-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_tle_15-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_15` | 1.5.0 | `el8.x86_64` | pigsty | 64.9 KiB | [pg_tle_15-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_15-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_15` | 1.5.0 | `el8.aarch64` | pigsty | 61.3 KiB | [pg_tle_15-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_15-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_15` | 1.2.0 | `el8.x86_64` | pgdg | 63.9 KiB | [pg_tle_15-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_tle_15-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_15` | 1.2.0 | `el8.aarch64` | pgdg | 60.2 KiB | [pg_tle_15-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_tle_15-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_15` | 1.5.1 | `el9.x86_64` | pgdg | 70.5 KiB | [pg_tle_15-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_tle_15-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_15` | 1.5.1 | `el9.x86_64` | pigsty | 66.5 KiB | [pg_tle_15-1.5.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_15-1.5.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tle_15` | 1.5.1 | `el9.aarch64` | pigsty | 63.8 KiB | [pg_tle_15-1.5.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_15-1.5.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tle_15` | 1.5.1 | `el9.aarch64` | pgdg | 68.1 KiB | [pg_tle_15-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_tle_15-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_15` | 1.5.0 | `el9.aarch64` | pgdg | 67.9 KiB | [pg_tle_15-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_tle_15-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_15` | 1.5.0 | `el9.aarch64` | pigsty | 63.8 KiB | [pg_tle_15-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_15-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_15` | 1.5.0 | `el9.x86_64` | pigsty | 66.4 KiB | [pg_tle_15-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_15-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_15` | 1.5.0 | `el9.x86_64` | pgdg | 70.8 KiB | [pg_tle_15-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_tle_15-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_15` | 1.2.0 | `el9.x86_64` | pgdg | 65.5 KiB | [pg_tle_15-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_tle_15-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_15` | 1.2.0 | `el9.aarch64` | pgdg | 62.3 KiB | [pg_tle_15-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_tle_15-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pg-tle` | 1.5.1 | `d12.aarch64` | pigsty | 156.2 KiB | [postgresql-15-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tle` | 1.5.1 | `d12.x86_64` | pigsty | 161.2 KiB | [postgresql-15-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tle` | 1.5.1 | `u22.aarch64` | pigsty | 186.0 KiB | [postgresql-15-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tle` | 1.5.1 | `u22.x86_64` | pigsty | 189.9 KiB | [postgresql-15-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tle` | 1.5.1 | `u24.x86_64` | pigsty | 168.5 KiB | [postgresql-15-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tle` | 1.5.1 | `u24.aarch64` | pigsty | 166.2 KiB | [postgresql-15-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-15-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_tle_14` | 1.5.1 | `el8.aarch64` | pgdg | 65.7 KiB | [pg_tle_14-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_tle_14-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_14` | 1.5.1 | `el8.aarch64` | pigsty | 61.6 KiB | [pg_tle_14-1.5.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_14-1.5.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tle_14` | 1.5.1 | `el8.x86_64` | pigsty | 65.2 KiB | [pg_tle_14-1.5.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_14-1.5.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tle_14` | 1.5.1 | `el8.x86_64` | pgdg | 69.3 KiB | [pg_tle_14-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_tle_14-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_14` | 1.5.0 | `el8.aarch64` | pigsty | 61.5 KiB | [pg_tle_14-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_14-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_14` | 1.5.0 | `el8.x86_64` | pgdg | 69.1 KiB | [pg_tle_14-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_tle_14-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_14` | 1.5.0 | `el8.aarch64` | pgdg | 65.5 KiB | [pg_tle_14-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_tle_14-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_14` | 1.5.0 | `el8.x86_64` | pigsty | 65.0 KiB | [pg_tle_14-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_14-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_14` | 1.2.0 | `el8.aarch64` | pgdg | 60.3 KiB | [pg_tle_14-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_tle_14-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_14` | 1.2.0 | `el8.x86_64` | pgdg | 64.1 KiB | [pg_tle_14-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_tle_14-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_14` | 1.5.1 | `el9.x86_64` | pgdg | 70.7 KiB | [pg_tle_14-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_tle_14-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_14` | 1.5.1 | `el9.aarch64` | pigsty | 63.9 KiB | [pg_tle_14-1.5.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_14-1.5.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tle_14` | 1.5.1 | `el9.x86_64` | pigsty | 66.6 KiB | [pg_tle_14-1.5.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_14-1.5.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tle_14` | 1.5.1 | `el9.aarch64` | pgdg | 68.3 KiB | [pg_tle_14-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_tle_14-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_14` | 1.5.0 | `el9.aarch64` | pigsty | 64.1 KiB | [pg_tle_14-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_14-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_14` | 1.5.0 | `el9.aarch64` | pgdg | 68.1 KiB | [pg_tle_14-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_tle_14-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_14` | 1.5.0 | `el9.x86_64` | pigsty | 66.7 KiB | [pg_tle_14-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_14-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_14` | 1.5.0 | `el9.x86_64` | pgdg | 70.6 KiB | [pg_tle_14-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_tle_14-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_14` | 1.2.0 | `el9.x86_64` | pgdg | 65.6 KiB | [pg_tle_14-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_tle_14-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_14` | 1.2.0 | `el9.aarch64` | pgdg | 62.5 KiB | [pg_tle_14-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_tle_14-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-pg-tle` | 1.5.1 | `d12.aarch64` | pigsty | 156.9 KiB | [postgresql-14-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tle` | 1.5.1 | `d12.x86_64` | pigsty | 162.1 KiB | [postgresql-14-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tle` | 1.5.1 | `u22.x86_64` | pigsty | 189.9 KiB | [postgresql-14-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tle` | 1.5.1 | `u22.aarch64` | pigsty | 186.0 KiB | [postgresql-14-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tle` | 1.5.1 | `u24.x86_64` | pigsty | 169.3 KiB | [postgresql-14-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tle` | 1.5.1 | `u24.aarch64` | pigsty | 167.0 KiB | [postgresql-14-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-14-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_tle_13` | 1.5.1 | `el8.x86_64` | pigsty | 64.3 KiB | [pg_tle_13-1.5.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_13-1.5.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tle_13` | 1.5.1 | `el8.x86_64` | pgdg | 68.4 KiB | [pg_tle_13-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_tle_13-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_13` | 1.5.1 | `el8.aarch64` | pgdg | 65.7 KiB | [pg_tle_13-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_tle_13-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_13` | 1.5.1 | `el8.aarch64` | pigsty | 61.6 KiB | [pg_tle_13-1.5.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_13-1.5.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tle_13` | 1.5.0 | `el8.aarch64` | pigsty | 61.5 KiB | [pg_tle_13-1.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tle_13-1.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_tle_13` | 1.5.0 | `el8.x86_64` | pgdg | 68.2 KiB | [pg_tle_13-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_tle_13-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_13` | 1.5.0 | `el8.x86_64` | pigsty | 64.1 KiB | [pg_tle_13-1.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tle_13-1.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_tle_13` | 1.5.0 | `el8.aarch64` | pgdg | 65.6 KiB | [pg_tle_13-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_tle_13-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_13` | 1.2.0 | `el8.x86_64` | pgdg | 63.2 KiB | [pg_tle_13-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_tle_13-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_tle_13` | 1.2.0 | `el8.aarch64` | pgdg | 60.4 KiB | [pg_tle_13-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_tle_13-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_tle_13` | 1.5.1 | `el9.x86_64` | pigsty | 67.0 KiB | [pg_tle_13-1.5.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_13-1.5.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tle_13` | 1.5.1 | `el9.aarch64` | pgdg | 68.5 KiB | [pg_tle_13-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_tle_13-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_13` | 1.5.1 | `el9.aarch64` | pigsty | 64.2 KiB | [pg_tle_13-1.5.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_13-1.5.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tle_13` | 1.5.1 | `el9.x86_64` | pgdg | 71.3 KiB | [pg_tle_13-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_tle_13-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_13` | 1.5.0 | `el9.x86_64` | pgdg | 71.2 KiB | [pg_tle_13-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_tle_13-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_13` | 1.5.0 | `el9.x86_64` | pigsty | 66.6 KiB | [pg_tle_13-1.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tle_13-1.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_tle_13` | 1.5.0 | `el9.aarch64` | pigsty | 64.2 KiB | [pg_tle_13-1.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tle_13-1.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_tle_13` | 1.5.0 | `el9.aarch64` | pgdg | 68.4 KiB | [pg_tle_13-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_tle_13-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_tle_13` | 1.2.0 | `el9.x86_64` | pgdg | 65.8 KiB | [pg_tle_13-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_tle_13-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_tle_13` | 1.2.0 | `el9.aarch64` | pgdg | 62.9 KiB | [pg_tle_13-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_tle_13-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-13-pg-tle` | 1.5.1 | `d12.aarch64` | pigsty | 157.1 KiB | [postgresql-13-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-tle` | 1.5.1 | `d12.x86_64` | pigsty | 162.2 KiB | [postgresql-13-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-tle` | 1.5.1 | `u22.aarch64` | pigsty | 186.0 KiB | [postgresql-13-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-tle` | 1.5.1 | `u22.x86_64` | pigsty | 189.5 KiB | [postgresql-13-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-tle` | 1.5.1 | `u24.x86_64` | pigsty | 169.7 KiB | [postgresql-13-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-tle` | 1.5.1 | `u24.aarch64` | pigsty | 167.1 KiB | [postgresql-13-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tle/postgresql-13-pg-tle_1.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/pg_tle" title="Repository" icon="github" subtitle="github.com/aws/pg_tle" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_tle-1.5.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_tle; # get pg_tle source code
pig build dep pg_tle; # install build dependencies
pig build pkg pg_tle; # build extension rpm or deb
pig build ext pg_tle; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_tle; # install by extension name, for the current active PG version
pig ext install pg_tle; # install via package alias, for the active PG version
pig ext install pg_tle -v 18;   # install for PG 18
pig ext install pg_tle -v 17;   # install for PG 17
pig ext install pg_tle -v 16;   # install for PG 16
pig ext install pg_tle -v 15;   # install for PG 15
pig ext install pg_tle -v 14;   # install for PG 14
pig ext install pg_tle -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_tle CASCADE SCHEMA pgtle;
```

