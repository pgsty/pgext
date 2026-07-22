---
title: "pg_session_jwt"
linkTitle: "pg_session_jwt"
description: "Manage authentication sessions using JWTs"
weight: 7060
categories: ["SEC"]
width: full
---

[**pg_session_jwt**](https://github.com/neondatabase/pg_session_jwt) : Manage authentication sessions using JWTs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7060** | {{< badge content="pg_session_jwt" link="https://github.com/neondatabase/pg_session_jwt" >}} | {{< ext "pg_session_jwt" >}} | `0.5.0` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `auth` |
|   **See Also**    | {{< ext "pgjwt" >}} {{< ext "pgaudit" >}} {{< ext "pgsodium" >}} {{< ext "supabase_vault" >}} {{< ext "anon" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_session_jwt` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "pg_session_jwt_18" "green" >}} {{< bg "17" "pg_session_jwt_17" "green" >}} {{< bg "16" "pg_session_jwt_16" "green" >}} {{< bg "15" "pg_session_jwt_15" "green" >}} {{< bg "14" "pg_session_jwt_14" "green" >}} | `pg_session_jwt_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "postgresql-18-pg-session-jwt" "green" >}} {{< bg "17" "postgresql-17-pg-session-jwt" "green" >}} {{< bg "16" "postgresql-16-pg-session-jwt" "green" >}} {{< bg "15" "postgresql-15-pg-session-jwt" "green" >}} {{< bg "14" "postgresql-14-pg-session-jwt" "green" >}} | `postgresql-$v-pg-session-jwt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_session_jwt_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pg-session-jwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pg-session-jwt : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_18` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1007.5 KiB | [pg_session_jwt_18-0.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_18-0.5.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_18` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 929.1 KiB | [pg_session_jwt_18-0.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_18-0.5.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_18` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1023.2 KiB | [pg_session_jwt_18-0.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_18-0.5.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_18` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 985.6 KiB | [pg_session_jwt_18-0.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_18-0.5.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_18` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1023.4 KiB | [pg_session_jwt_18-0.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_18-0.5.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_18` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 964.8 KiB | [pg_session_jwt_18-0.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_18-0.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 803.9 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 688.1 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 803.6 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 688.9 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 900.1 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 817.0 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 892.2 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 807.3 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 885.7 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-session-jwt` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 806.2 KiB | [postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-18-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_17` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1003.4 KiB | [pg_session_jwt_17-0.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_17-0.5.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_17` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 926.0 KiB | [pg_session_jwt_17-0.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_17-0.5.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_17` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1022.6 KiB | [pg_session_jwt_17-0.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_17-0.5.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_17` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 982.8 KiB | [pg_session_jwt_17-0.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_17-0.5.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_17` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1016.9 KiB | [pg_session_jwt_17-0.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_17-0.5.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_17` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 964.7 KiB | [pg_session_jwt_17-0.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_17-0.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 800.3 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 687.4 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 800.1 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 687.0 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 895.0 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 813.1 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 887.2 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 804.9 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 882.8 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-session-jwt` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 803.8 KiB | [postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-17-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_16` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1005.5 KiB | [pg_session_jwt_16-0.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_16-0.5.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_16` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 925.3 KiB | [pg_session_jwt_16-0.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_16-0.5.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_16` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1020.1 KiB | [pg_session_jwt_16-0.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_16-0.5.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_16` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 981.8 KiB | [pg_session_jwt_16-0.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_16-0.5.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_16` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1020.2 KiB | [pg_session_jwt_16-0.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_16-0.5.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_16` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 964.6 KiB | [pg_session_jwt_16-0.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_16-0.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 800.9 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 686.2 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 801.9 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 687.0 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 896.7 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 814.9 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 888.2 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 803.9 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 883.0 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-session-jwt` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 802.8 KiB | [postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-16-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_15` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 996.3 KiB | [pg_session_jwt_15-0.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_15-0.5.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_15` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 917.0 KiB | [pg_session_jwt_15-0.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_15-0.5.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_15` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1006.8 KiB | [pg_session_jwt_15-0.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_15-0.5.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_15` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 971.3 KiB | [pg_session_jwt_15-0.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_15-0.5.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_15` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1006.8 KiB | [pg_session_jwt_15-0.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_15-0.5.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_15` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 959.9 KiB | [pg_session_jwt_15-0.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_15-0.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 796.0 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 681.4 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 794.8 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 682.6 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 887.7 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 808.3 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 880.4 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 798.8 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 877.2 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-session-jwt` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 796.7 KiB | [postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-15-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_session_jwt_14` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 992.6 KiB | [pg_session_jwt_14-0.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_session_jwt_14-0.5.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_session_jwt_14` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 914.4 KiB | [pg_session_jwt_14-0.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_session_jwt_14-0.5.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_session_jwt_14` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1001.9 KiB | [pg_session_jwt_14-0.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_session_jwt_14-0.5.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_session_jwt_14` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 969.8 KiB | [pg_session_jwt_14-0.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_session_jwt_14-0.5.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_session_jwt_14` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1001.7 KiB | [pg_session_jwt_14-0.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_session_jwt_14-0.5.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_session_jwt_14` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 958.6 KiB | [pg_session_jwt_14-0.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_session_jwt_14-0.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 791.0 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 679.4 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 792.7 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 679.5 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 882.1 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 806.9 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 873.7 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 797.3 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 869.8 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-session-jwt` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 794.1 KiB | [postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-session-jwt/postgresql-14-pg-session-jwt_0.5.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/neondatabase/pg_session_jwt" title="Repository" icon="github" subtitle="github.com/neondatabase/pg_session_jwt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_session_jwt-0.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_session_jwt;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_session_jwt;		# install via package name, for the active PG version

pig install pg_session_jwt -v 18;   # install for PG 18
pig install pg_session_jwt -v 17;   # install for PG 17
pig install pg_session_jwt -v 16;   # install for PG 16
pig install pg_session_jwt -v 15;   # install for PG 15
pig install pg_session_jwt -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_session_jwt;
```




## Usage

Sources: [README](https://github.com/neondatabase/pg_session_jwt/blob/v0.5.0/README.md), [v0.5.0 tag](https://github.com/neondatabase/pg_session_jwt/tree/v0.5.0), [control file](https://github.com/neondatabase/pg_session_jwt/blob/v0.5.0/pg_session_jwt.control)

`pg_session_jwt` handles authenticated sessions through JWTs. When configured with a JWK, it verifies JWT authenticity. Without a JWK, it falls back to PostgREST-compatible `request.jwt.claims`.

```sql
CREATE EXTENSION pg_session_jwt;
```

### Mode 1: JWK Validation

Set the JWK at connection time via libpq options:

```bash
export PGOPTIONS="-c pg_session_jwt.jwk=$MY_JWK"
```

Then within the session:

```sql
SELECT auth.init();                        -- Initialize with JWK
SELECT auth.jwt_session_init('eyJ...');    -- Set and validate the JWT
SELECT auth.user_id();                     -- Get the 'sub' claim
SELECT auth.session();                     -- Get full JWT payload as JSONB
```

### Mode 2: PostgREST-Compatible (No JWK)

Works out of the box with PostgREST. No initialization needed:

```sql
SELECT auth.user_id();   -- Returns 'sub' from request.jwt.claims
SELECT auth.session();   -- Returns full claims as JSONB
```

### Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `auth.init()` | `void` | Initialize session using JWK |
| `auth.jwt_session_init(jwt text)` | `void` | Set and validate a JWT |
| `auth.session()` | `jsonb` | Get JWT payload or fallback claims |
| `auth.jwt()` | `jsonb` | Alias for `auth.session()` |
| `auth.user_id()` | `text` | Get the `sub` claim |
| `auth.uid()` | `uuid` | Get `sub` as UUID (or NULL) |
| `auth.organization()` | `jsonb` | Neon Auth organization claim helper |
| `auth.organization_id()` | `uuid` | Neon Auth organization id helper |

### Configuration

| Parameter | Description |
|-----------|-------------|
| `pg_session_jwt.jwk` | JWK for JWT validation (set at startup or connection) |
| `pg_session_jwt.audit_log` | Enable audit logging (`on`/`off`) |

### RLS Example

```sql
CREATE POLICY user_isolation ON my_table
    USING (user_id = auth.user_id());
```

For Neon Auth organization-scoped policies, use the `o` claim helpers:

```sql
CREATE POLICY team_select ON team
  FOR SELECT
  USING (organization_id = auth.organization_id());
```

### Version Notes

The v0.5.0 README adds Neon Auth organization helpers and explicitly separates portable helpers such as `auth.jwt()`, `auth.user_id()`, and `auth.uid()` from Neon-specific `auth.organization()` and `auth.organization_id()`. Other auth providers should use `auth.jwt()` and extract provider-specific claims directly.
