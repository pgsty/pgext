---
title: "anon"
linkTitle: "anon"
description: "PostgreSQL Anonymizer (anon) extension"
weight: 7050
categories: ["Sec"]
width: full
---

PostgreSQL Anonymizer (anon) extension

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7050** | {{< badge content="anon" link="https://gitlab.com/dalibo/postgresql_anonymizer/" >}} | {{< ext "anon" "pg_anon" >}} | `2.3.0` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "faker" >}} {{< ext "pgsodium" >}} {{< ext "pg_tde" >}} {{< ext "pgcrypto" >}} {{< ext "pg_permissions" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_drop_events" >}} |

> [!Note] pgrx=0.14.3


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/anon" >}} | `2.3.0` | {{< badge content="18" color="red" alt="pg_anon_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_anon_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/anon" >}} | `2.3.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-anon" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-anon` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_anon_18" >}}     | {{< pkg "pg_anon_17" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_17-2.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_anon_16" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_16-2.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_anon_15" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_15-2.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_anon_14" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_14-2.3.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_anon_18" >}}     | {{< pkg "pg_anon_17" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_17-2.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_anon_16" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_16-2.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_anon_15" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_15-2.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_anon_14" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_14-2.3.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_anon_18" >}}     | {{< pkg "pg_anon_17" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_17-2.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_anon_16" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_16-2.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_anon_15" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_15-2.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_anon_14" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_14-2.3.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_anon_18" >}}     | {{< pkg "pg_anon_17" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_17-2.3.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_anon_16" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_16-2.3.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_anon_15" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_15-2.3.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_anon_14" "2.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_14-2.3.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-anon" >}}     | {{< pkg "postgresql-17-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-anon" >}}     | {{< pkg "postgresql-17-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-anon" >}}     | {{< pkg "postgresql-17-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-anon" >}}     | {{< pkg "postgresql-17-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-anon" >}}     | {{< pkg "postgresql-17-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-anon" >}}     | {{< pkg "postgresql-17-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-anon" "2.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_anon_17` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_17-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_17-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_17` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_17-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_17-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_17` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_17-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_17-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_17` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_17-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_17-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.3 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.6 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_anon_16` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_16-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_16-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_16` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_16-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_16-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_16` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_16-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_16-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_16` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_16-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_16-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.3 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.6 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_anon_15` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_15-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_15-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_15` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_15-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_15-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_15` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_15-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_15-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_15` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_15-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_15-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.3 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.6 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_anon_14` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_14-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_14-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_14` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_14-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_14-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_14` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_14-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_14-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_14` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_14-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_14-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.3 MiB | [postgresql-14-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-14-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.6 MiB | [postgresql-14-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-14-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-14-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-14-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_anon_13` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_13-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_13-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_13` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_13-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_13-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_13` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_13-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_13-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_13` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_13-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_13-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.4 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.7 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitlab.com/dalibo/postgresql_anonymizer/" title="Repository" icon="link" subtitle="gitlab.com/dalibo/postgresql_anonymizer/" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="postgresql_anonymizer-2.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get anon; # get anon source code
pig build dep anon; # install build dependencies
pig build pkg anon; # build extension rpm or deb
pig build ext anon; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install anon; # install by extension name, for the current active PG version
pig ext install pg_anon; # install via package alias, for the active PG version
pig ext install anon -v 17;   # install for PG 17
pig ext install anon -v 16;   # install for PG 16
pig ext install anon -v 15;   # install for PG 15
pig ext install anon -v 14;   # install for PG 14
pig ext install anon -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION anon CASCADE SCHEMA anon;
```

