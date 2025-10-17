---
title: "pg_session_jwt"
linkTitle: "pg_session_jwt"
description: "Manage authentication sessions using JWTs"
weight: 7040
categories: ["Sec"]
width: full
---

Manage authentication sessions using JWTs

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7040** | {{< badge content="pg_session_jwt" link="https://github.com/neondatabase/pg_session_jwt" >}} | {{< ext "pg_session_jwt" "pg_session_jwt" >}} | `0.3.1` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "supautils" >}} {{< ext "pgsodium" >}} {{< ext "supabase_vault" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgaudit" >}} |

> [!Note] pgrx=0.12.9 (0.12.6)


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_session_jwt" >}} | `0.3.1` | {{< badge content="18" color="red" alt="pg_session_jwt_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_session_jwt_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_session_jwt" >}} | `0.3.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-session-jwt" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-session-jwt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_session_jwt_18" >}}     | {{< pkg "pg_session_jwt_17" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_17-0.3.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_session_jwt_16" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_16-0.3.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_session_jwt_15" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_15-0.3.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_session_jwt_14" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_14-0.3.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_session_jwt_18" >}}     | {{< pkg "pg_session_jwt_17" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_17-0.3.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_session_jwt_16" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_16-0.3.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_session_jwt_15" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_15-0.3.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_session_jwt_14" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_14-0.3.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_session_jwt_18" >}}     | {{< pkg "pg_session_jwt_17" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_17-0.3.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_session_jwt_16" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_16-0.3.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_session_jwt_15" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_15-0.3.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_session_jwt_14" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_14-0.3.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_session_jwt_18" >}}     | {{< pkg "pg_session_jwt_17" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_17-0.3.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_session_jwt_16" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_16-0.3.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_session_jwt_15" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_15-0.3.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_session_jwt_14" "0.3.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_14-0.3.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-session-jwt" >}}     | {{< pkg "postgresql-17-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-session-jwt" >}}     | {{< pkg "postgresql-17-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-session-jwt" >}}     | {{< pkg "postgresql-17-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-session-jwt" >}}     | {{< pkg "postgresql-17-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-session-jwt" >}}     | {{< pkg "postgresql-17-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-session-jwt" >}}     | {{< pkg "postgresql-17-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-session-jwt" "0.3.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_session_jwt_17` | 0.3.1 | `el8.aarch64` | pigsty | 311.8 KiB | [pg_session_jwt_17-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_17-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_17` | 0.3.1 | `el8.x86_64` | pigsty | 348.2 KiB | [pg_session_jwt_17-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_17-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_17` | 0.3.1 | `el9.aarch64` | pigsty | 332.2 KiB | [pg_session_jwt_17-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_17-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_17` | 0.3.1 | `el9.x86_64` | pigsty | 353.2 KiB | [pg_session_jwt_17-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_17-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `d12.x86_64` | pigsty | 289.0 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `d12.aarch64` | pigsty | 249.2 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `u22.aarch64` | pigsty | 288.6 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `u22.x86_64` | pigsty | 316.1 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `u24.x86_64` | pigsty | 313.6 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `u24.aarch64` | pigsty | 285.8 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_session_jwt_16` | 0.3.1 | `el8.x86_64` | pigsty | 348.1 KiB | [pg_session_jwt_16-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_16-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_16` | 0.3.1 | `el8.aarch64` | pigsty | 311.8 KiB | [pg_session_jwt_16-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_16-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_16` | 0.3.1 | `el9.aarch64` | pigsty | 332.2 KiB | [pg_session_jwt_16-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_16-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_16` | 0.3.1 | `el9.x86_64` | pigsty | 353.0 KiB | [pg_session_jwt_16-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_16-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `d12.aarch64` | pigsty | 249.2 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `d12.x86_64` | pigsty | 288.9 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `u22.x86_64` | pigsty | 316.1 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `u22.aarch64` | pigsty | 288.6 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `u24.aarch64` | pigsty | 285.5 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `u24.x86_64` | pigsty | 313.5 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_session_jwt_15` | 0.3.1 | `el8.x86_64` | pigsty | 348.0 KiB | [pg_session_jwt_15-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_15-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_15` | 0.3.1 | `el8.aarch64` | pigsty | 311.9 KiB | [pg_session_jwt_15-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_15-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_15` | 0.3.1 | `el9.x86_64` | pigsty | 352.4 KiB | [pg_session_jwt_15-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_15-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_15` | 0.3.1 | `el9.aarch64` | pigsty | 332.6 KiB | [pg_session_jwt_15-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_15-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `d12.x86_64` | pigsty | 289.0 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `d12.aarch64` | pigsty | 249.1 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `u22.aarch64` | pigsty | 288.6 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `u22.x86_64` | pigsty | 316.2 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `u24.aarch64` | pigsty | 285.6 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `u24.x86_64` | pigsty | 313.6 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_session_jwt_14` | 0.3.1 | `el8.aarch64` | pigsty | 311.8 KiB | [pg_session_jwt_14-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_14-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_14` | 0.3.1 | `el8.x86_64` | pigsty | 348.0 KiB | [pg_session_jwt_14-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_14-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_14` | 0.3.1 | `el9.x86_64` | pigsty | 352.5 KiB | [pg_session_jwt_14-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_14-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_14` | 0.3.1 | `el9.aarch64` | pigsty | 332.3 KiB | [pg_session_jwt_14-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_14-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `d12.aarch64` | pigsty | 249.0 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `d12.x86_64` | pigsty | 288.8 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `u22.x86_64` | pigsty | 316.1 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `u22.aarch64` | pigsty | 288.6 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `u24.aarch64` | pigsty | 285.9 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `u24.x86_64` | pigsty | 313.6 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/neondatabase/pg_session_jwt" title="Repository" icon="github" subtitle="github.com/neondatabase/pg_session_jwt" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_session_jwt-0.3.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_session_jwt; # get pg_session_jwt source code
pig build dep pg_session_jwt; # install build dependencies
pig build pkg pg_session_jwt; # build extension rpm or deb
pig build ext pg_session_jwt; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_session_jwt; # install by extension name, for the current active PG version
pig ext install pg_session_jwt; # install via package alias, for the active PG version
pig ext install pg_session_jwt -v 17;   # install for PG 17
pig ext install pg_session_jwt -v 16;   # install for PG 16
pig ext install pg_session_jwt -v 15;   # install for PG 15
pig ext install pg_session_jwt -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_session_jwt CASCADE SCHEMA auth;
```

