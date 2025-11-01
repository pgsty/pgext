---
title: "pg_session_jwt"
linkTitle: "pg_session_jwt"
description: "Manage authentication sessions using JWTs"
weight: 7040
categories: ["SEC"]
width: full
---

Manage authentication sessions using JWTs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7040** | {{< badge content="pg_session_jwt" link="https://github.com/neondatabase/pg_session_jwt" >}} | {{< ext "pg_session_jwt" >}} | `0.3.3` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "supautils" >}} {{< ext "pgsodium" >}} {{< ext "supabase_vault" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgaudit" >}} |

> [!Note] pgrx=0.16.1, manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_session_jwt" >}} | `0.3.3` | {{< bg "18" "pg_session_jwt_18" "green" >}} {{< bg "17" "pg_session_jwt_17" "green" >}} {{< bg "16" "pg_session_jwt_16" "green" >}} {{< bg "15" "pg_session_jwt_15" "green" >}} {{< bg "14" "pg_session_jwt_14" "green" >}} {{< bg "13" "pg_session_jwt_13" "red" >}} | `pg_session_jwt_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_session_jwt" >}} | `0.3.3` | {{< bg "18" "postgresql-18-pg-session-jwt" "green" >}} {{< bg "17" "postgresql-17-pg-session-jwt" "green" >}} {{< bg "16" "postgresql-16-pg-session-jwt" "green" >}} {{< bg "15" "postgresql-15-pg-session-jwt" "green" >}} {{< bg "14" "postgresql-14-pg-session-jwt" "green" >}} {{< bg "13" "postgresql-13-pg-session-jwt" "red" >}} | `postgresql-$v-pg-session-jwt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_session_jwt_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_session_jwt_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_session_jwt_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_session_jwt_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_session_jwt_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pg_session_jwt_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_session_jwt_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-session-jwt : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-session-jwt : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-session-jwt : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-session-jwt : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-session-jwt : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-session-jwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-session-jwt : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-session-jwt : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-session-jwt : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-session-jwt : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-session-jwt : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-session-jwt : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-session-jwt : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-session-jwt : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-session-jwt : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_18` | 0.3.3 | `el8.x86_64` | pigsty | 433.0 KiB | [pg_session_jwt_18-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_18-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_18` | 0.3.3 | `el8.aarch64` | pigsty | 305.8 KiB | [pg_session_jwt_18-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_18-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_18` | 0.3.3 | `el9.x86_64` | pigsty | 448.8 KiB | [pg_session_jwt_18-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_18-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_18` | 0.3.3 | `el9.aarch64` | pigsty | 324.6 KiB | [pg_session_jwt_18-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_18-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_18` | 0.3.3 | `el10.x86_64` | pigsty | 447.3 KiB | [pg_session_jwt_18-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_18-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_18` | 0.3.3 | `el10.aarch64` | pigsty | 324.7 KiB | [pg_session_jwt_18-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_18-0.3.3-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_17` | 0.3.3 | `el8.x86_64` | pigsty | 432.9 KiB | [pg_session_jwt_17-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_17-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_17` | 0.3.3 | `el8.aarch64` | pigsty | 305.8 KiB | [pg_session_jwt_17-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_17-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_17` | 0.3.3 | `el9.x86_64` | pigsty | 448.7 KiB | [pg_session_jwt_17-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_17-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_17` | 0.3.3 | `el9.aarch64` | pigsty | 324.7 KiB | [pg_session_jwt_17-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_17-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_17` | 0.3.3 | `el10.x86_64` | pigsty | 447.1 KiB | [pg_session_jwt_17-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_17-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_17` | 0.3.3 | `el10.aarch64` | pigsty | 324.7 KiB | [pg_session_jwt_17-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_17-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `d12.x86_64` | pigsty | 289.0 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `d12.aarch64` | pigsty | 249.2 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `u22.x86_64` | pigsty | 316.1 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `u22.aarch64` | pigsty | 288.6 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `u24.x86_64` | pigsty | 313.6 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-session-jwt` | 0.3.1 | `u24.aarch64` | pigsty | 285.8 KiB | [postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_16` | 0.3.3 | `el8.x86_64` | pigsty | 432.9 KiB | [pg_session_jwt_16-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_16-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_16` | 0.3.3 | `el8.aarch64` | pigsty | 305.9 KiB | [pg_session_jwt_16-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_16-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_16` | 0.3.3 | `el9.x86_64` | pigsty | 448.9 KiB | [pg_session_jwt_16-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_16-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_16` | 0.3.3 | `el9.aarch64` | pigsty | 324.8 KiB | [pg_session_jwt_16-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_16-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_16` | 0.3.3 | `el10.x86_64` | pigsty | 447.2 KiB | [pg_session_jwt_16-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_16-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_16` | 0.3.3 | `el10.aarch64` | pigsty | 324.8 KiB | [pg_session_jwt_16-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_16-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `d12.x86_64` | pigsty | 288.9 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `d12.aarch64` | pigsty | 249.2 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `u22.x86_64` | pigsty | 316.1 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `u22.aarch64` | pigsty | 288.6 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `u24.x86_64` | pigsty | 313.5 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-session-jwt` | 0.3.1 | `u24.aarch64` | pigsty | 285.5 KiB | [postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_15` | 0.3.3 | `el8.x86_64` | pigsty | 432.5 KiB | [pg_session_jwt_15-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_15-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_15` | 0.3.3 | `el8.aarch64` | pigsty | 305.9 KiB | [pg_session_jwt_15-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_15-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_15` | 0.3.3 | `el9.x86_64` | pigsty | 448.6 KiB | [pg_session_jwt_15-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_15-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_15` | 0.3.3 | `el9.aarch64` | pigsty | 324.8 KiB | [pg_session_jwt_15-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_15-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_15` | 0.3.3 | `el10.x86_64` | pigsty | 446.9 KiB | [pg_session_jwt_15-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_15-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_15` | 0.3.3 | `el10.aarch64` | pigsty | 324.9 KiB | [pg_session_jwt_15-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_15-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `d12.x86_64` | pigsty | 289.0 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `d12.aarch64` | pigsty | 249.1 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `u22.x86_64` | pigsty | 316.2 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `u22.aarch64` | pigsty | 288.6 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `u24.x86_64` | pigsty | 313.6 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-session-jwt` | 0.3.1 | `u24.aarch64` | pigsty | 285.6 KiB | [postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_14` | 0.3.3 | `el8.x86_64` | pigsty | 432.5 KiB | [pg_session_jwt_14-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_14-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_14` | 0.3.3 | `el8.aarch64` | pigsty | 305.8 KiB | [pg_session_jwt_14-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_14-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_14` | 0.3.3 | `el9.x86_64` | pigsty | 448.6 KiB | [pg_session_jwt_14-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_14-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_14` | 0.3.3 | `el9.aarch64` | pigsty | 324.8 KiB | [pg_session_jwt_14-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_14-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_14` | 0.3.3 | `el10.x86_64` | pigsty | 446.4 KiB | [pg_session_jwt_14-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_14-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_14` | 0.3.3 | `el10.aarch64` | pigsty | 324.8 KiB | [pg_session_jwt_14-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_14-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `d12.x86_64` | pigsty | 288.8 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `d12.aarch64` | pigsty | 249.0 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `u22.x86_64` | pigsty | 316.1 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `u22.aarch64` | pigsty | 288.6 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `u24.x86_64` | pigsty | 313.6 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-session-jwt` | 0.3.1 | `u24.aarch64` | pigsty | 285.9 KiB | [postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/neondatabase/pg_session_jwt" title="Repository" icon="github" subtitle="github.com/neondatabase/pg_session_jwt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_session_jwt-0.3.3.tar.gz" >}}
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
pig ext install pg_session_jwt -v 18;   # install for PG 18
pig ext install pg_session_jwt -v 17;   # install for PG 17
pig ext install pg_session_jwt -v 16;   # install for PG 16
pig ext install pg_session_jwt -v 15;   # install for PG 15
pig ext install pg_session_jwt -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_session_jwt CASCADE SCHEMA auth;
```

