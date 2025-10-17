---
title: "pg_readonly"
linkTitle: "pg_readonly"
description: "cluster database read only"
weight: 5120
categories: ["Admin"]
width: full
---

cluster database read only

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5120** | {{< badge content="pg_readonly" link="https://github.com/pierreforstmann/pg_readonly" >}} | {{< ext "pg_readonly" "pg_readonly" >}} | `1.0.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_permissions" >}} {{< ext "pg_upless" >}} {{< ext "safeupdate" >}} {{< ext "set_user" >}} {{< ext "pgaudit" >}} {{< ext "noset" >}} {{< ext "sepgsql" >}} {{< ext "login_hook" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_readonly" >}} | `1.0.3` | {{< badge content="18" color="red" alt="pg_readonly_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_readonly_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_readonly" >}} | `1.0.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-readonly" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-readonly` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_readonly_18" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_readonly_18-1.0.3-5PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_readonly_17" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_readonly_17-1.0.3-4PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_readonly_16" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_readonly_16-1.0.3-2.rhel8.1.x86_64.rpm" >}} | {{< pkg "pg_readonly_15" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readonly_15-1.0.3-1.rhel8.x86_64.rpm" >}} | {{< pkg "pg_readonly_14" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readonly_14-1.0.3-1.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_readonly_18" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_readonly_18-1.0.3-5PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_readonly_17" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_readonly_17-1.0.3-4PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_readonly_16" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_readonly_16-1.0.3-2.rhel8.1.aarch64.rpm" >}} | {{< pkg "pg_readonly_15" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_readonly_15-1.0.3-1.rhel8.aarch64.rpm" >}} | {{< pkg "pg_readonly_14" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_readonly_14-1.0.3-1.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_readonly_18" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_readonly_18-1.0.3-5PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_readonly_17" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_readonly_17-1.0.3-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_readonly_16" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_readonly_16-1.0.3-2.rhel9.1.x86_64.rpm" >}} | {{< pkg "pg_readonly_15" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readonly_15-1.0.3-1.rhel9.x86_64.rpm" >}} | {{< pkg "pg_readonly_14" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readonly_14-1.0.3-1.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_readonly_18" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_readonly_18-1.0.3-5PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_readonly_17" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_readonly_17-1.0.3-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_readonly_16" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_readonly_16-1.0.3-2.rhel9.1.aarch64.rpm" >}} | {{< pkg "pg_readonly_15" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_readonly_15-1.0.3-1.rhel9.aarch64.rpm" >}} | {{< pkg "pg_readonly_14" "1.0.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_readonly_14-1.0.3-1.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-readonly" >}}     | {{< pkg "postgresql-17-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-readonly" >}}     | {{< pkg "postgresql-17-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-readonly" >}}     | {{< pkg "postgresql-17-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-readonly" >}}     | {{< pkg "postgresql-17-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-readonly" >}}     | {{< pkg "postgresql-17-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-readonly" >}}     | {{< pkg "postgresql-17-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-readonly" "1.0.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readonly_18` | 1.0.3 | `el8.aarch64` | pgdg | 16.5 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_readonly_18-1.0.3-5PGDG.rhel8.aarch64.rpm) |
| `pg_readonly_18` | 1.0.3 | `el8.x86_64` | pgdg | 16.6 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_readonly_18-1.0.3-5PGDG.rhel8.x86_64.rpm) |
| `pg_readonly_18` | 1.0.3 | `el9.aarch64` | pgdg | 16.0 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_readonly_18-1.0.3-5PGDG.rhel9.aarch64.rpm) |
| `pg_readonly_18` | 1.0.3 | `el9.x86_64` | pgdg | 16.5 KiB | [pg_readonly_18-1.0.3-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_readonly_18-1.0.3-5PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readonly_17` | 1.0.3 | `el8.x86_64` | pgdg | 16.6 KiB | [pg_readonly_17-1.0.3-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_readonly_17-1.0.3-4PGDG.rhel8.x86_64.rpm) |
| `pg_readonly_17` | 1.0.3 | `el8.aarch64` | pgdg | 16.4 KiB | [pg_readonly_17-1.0.3-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_readonly_17-1.0.3-4PGDG.rhel8.aarch64.rpm) |
| `pg_readonly_17` | 1.0.3 | `el9.x86_64` | pgdg | 16.7 KiB | [pg_readonly_17-1.0.3-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_readonly_17-1.0.3-4PGDG.rhel9.x86_64.rpm) |
| `pg_readonly_17` | 1.0.3 | `el9.aarch64` | pgdg | 16.4 KiB | [pg_readonly_17-1.0.3-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_readonly_17-1.0.3-4PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pg-readonly` | 1.0.0 | `d12.aarch64` | pigsty | 22.1 KiB | [postgresql-17-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-readonly` | 1.0.0 | `d12.x86_64` | pigsty | 22.4 KiB | [postgresql-17-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-readonly` | 1.0.0 | `u22.x86_64` | pigsty | 23.1 KiB | [postgresql-17-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-readonly` | 1.0.0 | `u22.aarch64` | pigsty | 22.9 KiB | [postgresql-17-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-readonly` | 1.0.0 | `u24.aarch64` | pigsty | 19.9 KiB | [postgresql-17-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-readonly` | 1.0.0 | `u24.x86_64` | pigsty | 20.2 KiB | [postgresql-17-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-17-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readonly_16` | 1.0.3 | `el8.x86_64` | pgdg | 16.4 KiB | [pg_readonly_16-1.0.3-2.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_readonly_16-1.0.3-2.rhel8.1.x86_64.rpm) |
| `pg_readonly_16` | 1.0.3 | `el8.aarch64` | pgdg | 16.2 KiB | [pg_readonly_16-1.0.3-2.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_readonly_16-1.0.3-2.rhel8.1.aarch64.rpm) |
| `pg_readonly_16` | 1.0.3 | `el9.x86_64` | pgdg | 16.5 KiB | [pg_readonly_16-1.0.3-2.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_readonly_16-1.0.3-2.rhel9.1.x86_64.rpm) |
| `pg_readonly_16` | 1.0.3 | `el9.aarch64` | pgdg | 16.1 KiB | [pg_readonly_16-1.0.3-2.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_readonly_16-1.0.3-2.rhel9.1.aarch64.rpm) |
| `postgresql-16-pg-readonly` | 1.0.0 | `d12.aarch64` | pigsty | 22.1 KiB | [postgresql-16-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-readonly` | 1.0.0 | `d12.x86_64` | pigsty | 22.3 KiB | [postgresql-16-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-readonly` | 1.0.0 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-16-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-readonly` | 1.0.0 | `u22.x86_64` | pigsty | 23.0 KiB | [postgresql-16-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-readonly` | 1.0.0 | `u24.x86_64` | pigsty | 20.2 KiB | [postgresql-16-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-readonly` | 1.0.0 | `u24.aarch64` | pigsty | 19.9 KiB | [postgresql-16-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-16-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readonly_15` | 1.0.3 | `el8.x86_64` | pgdg | 30.7 KiB | [pg_readonly_15-1.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readonly_15-1.0.3-1.rhel8.x86_64.rpm) |
| `pg_readonly_15` | 1.0.3 | `el8.aarch64` | pgdg | 30.4 KiB | [pg_readonly_15-1.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_readonly_15-1.0.3-1.rhel8.aarch64.rpm) |
| `pg_readonly_15` | 1.0.1 | `el8.x86_64` | pgdg | 29.2 KiB | [pg_readonly_15-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_readonly_15-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_readonly_15` | 1.0.3 | `el9.x86_64` | pgdg | 31.3 KiB | [pg_readonly_15-1.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readonly_15-1.0.3-1.rhel9.x86_64.rpm) |
| `pg_readonly_15` | 1.0.3 | `el9.aarch64` | pgdg | 31.0 KiB | [pg_readonly_15-1.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_readonly_15-1.0.3-1.rhel9.aarch64.rpm) |
| `pg_readonly_15` | 1.0.1 | `el9.x86_64` | pgdg | 29.9 KiB | [pg_readonly_15-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_readonly_15-1.0.1-1.rhel9.x86_64.rpm) |
| `postgresql-15-pg-readonly` | 1.0.0 | `d12.x86_64` | pigsty | 22.3 KiB | [postgresql-15-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-readonly` | 1.0.0 | `d12.aarch64` | pigsty | 22.0 KiB | [postgresql-15-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-readonly` | 1.0.0 | `u22.x86_64` | pigsty | 23.0 KiB | [postgresql-15-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-readonly` | 1.0.0 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-15-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-readonly` | 1.0.0 | `u24.x86_64` | pigsty | 20.2 KiB | [postgresql-15-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-readonly` | 1.0.0 | `u24.aarch64` | pigsty | 19.9 KiB | [postgresql-15-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-15-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readonly_14` | 1.0.3 | `el8.aarch64` | pgdg | 30.1 KiB | [pg_readonly_14-1.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_readonly_14-1.0.3-1.rhel8.aarch64.rpm) |
| `pg_readonly_14` | 1.0.3 | `el8.x86_64` | pgdg | 30.5 KiB | [pg_readonly_14-1.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readonly_14-1.0.3-1.rhel8.x86_64.rpm) |
| `pg_readonly_14` | 1.0.1 | `el8.x86_64` | pgdg | 29.7 KiB | [pg_readonly_14-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_readonly_14-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_readonly_14` | 1.0.3 | `el9.x86_64` | pgdg | 31.1 KiB | [pg_readonly_14-1.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readonly_14-1.0.3-1.rhel9.x86_64.rpm) |
| `pg_readonly_14` | 1.0.3 | `el9.aarch64` | pgdg | 30.7 KiB | [pg_readonly_14-1.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_readonly_14-1.0.3-1.rhel9.aarch64.rpm) |
| `pg_readonly_14` | 1.0.1 | `el9.x86_64` | pgdg | 29.9 KiB | [pg_readonly_14-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_readonly_14-1.0.1-1.rhel9.x86_64.rpm) |
| `postgresql-14-pg-readonly` | 1.0.0 | `d12.x86_64` | pigsty | 22.1 KiB | [postgresql-14-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-readonly` | 1.0.0 | `d12.aarch64` | pigsty | 21.8 KiB | [postgresql-14-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-readonly` | 1.0.0 | `u22.x86_64` | pigsty | 22.7 KiB | [postgresql-14-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-readonly` | 1.0.0 | `u22.aarch64` | pigsty | 22.6 KiB | [postgresql-14-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-readonly` | 1.0.0 | `u24.aarch64` | pigsty | 19.7 KiB | [postgresql-14-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-readonly` | 1.0.0 | `u24.x86_64` | pigsty | 20.0 KiB | [postgresql-14-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-14-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_readonly_13` | 1.0.3 | `el8.aarch64` | pgdg | 29.9 KiB | [pg_readonly_13-1.0.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_readonly_13-1.0.3-1.rhel8.aarch64.rpm) |
| `pg_readonly_13` | 1.0.3 | `el8.x86_64` | pgdg | 30.2 KiB | [pg_readonly_13-1.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_readonly_13-1.0.3-1.rhel8.x86_64.rpm) |
| `pg_readonly_13` | 1.0.1 | `el8.x86_64` | pgdg | 29.3 KiB | [pg_readonly_13-1.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_readonly_13-1.0.1-1.rhel8.x86_64.rpm) |
| `pg_readonly_13` | 1.0.3 | `el9.aarch64` | pgdg | 30.4 KiB | [pg_readonly_13-1.0.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_readonly_13-1.0.3-1.rhel9.aarch64.rpm) |
| `pg_readonly_13` | 1.0.3 | `el9.x86_64` | pgdg | 30.7 KiB | [pg_readonly_13-1.0.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_readonly_13-1.0.3-1.rhel9.x86_64.rpm) |
| `pg_readonly_13` | 1.0.1 | `el9.x86_64` | pgdg | 29.6 KiB | [pg_readonly_13-1.0.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_readonly_13-1.0.1-1.rhel9.x86_64.rpm) |
| `postgresql-13-pg-readonly` | 1.0.0 | `d12.aarch64` | pigsty | 21.6 KiB | [postgresql-13-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-readonly` | 1.0.0 | `d12.x86_64` | pigsty | 21.7 KiB | [postgresql-13-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-readonly` | 1.0.0 | `u22.aarch64` | pigsty | 22.3 KiB | [postgresql-13-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-readonly` | 1.0.0 | `u22.x86_64` | pigsty | 22.4 KiB | [postgresql-13-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-readonly` | 1.0.0 | `u24.x86_64` | pigsty | 19.6 KiB | [postgresql-13-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-readonly` | 1.0.0 | `u24.aarch64` | pigsty | 19.5 KiB | [postgresql-13-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-readonly/postgresql-13-pg-readonly_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pierreforstmann/pg_readonly" title="Repository" icon="github" subtitle="github.com/pierreforstmann/pg_readonly" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_readonly-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_readonly; # get pg_readonly source code
pig build dep pg_readonly; # install build dependencies
pig build pkg pg_readonly; # build extension rpm or deb
pig build ext pg_readonly; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_readonly; # install by extension name, for the current active PG version
pig ext install pg_readonly; # install via package alias, for the active PG version
pig ext install pg_readonly -v 17;   # install for PG 17
pig ext install pg_readonly -v 16;   # install for PG 16
pig ext install pg_readonly -v 15;   # install for PG 15
pig ext install pg_readonly -v 14;   # install for PG 14
pig ext install pg_readonly -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_readonly;
```

