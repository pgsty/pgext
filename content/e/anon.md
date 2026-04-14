---
title: "anon"
linkTitle: "anon"
description: "PostgreSQL Anonymizer (anon) extension"
weight: 7070
categories: ["SEC"]
width: full
---

[**pg_anon**](https://gitlab.com/dalibo/postgresql_anonymizer/) : PostgreSQL Anonymizer (anon) extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7070** | {{< badge content="anon" link="https://gitlab.com/dalibo/postgresql_anonymizer/" >}} | {{< ext "anon" "pg_anon" >}} | `3.0.13` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `anon` |
|   **See Also**    | {{< ext "faker" >}} {{< ext "pgsodium" >}} {{< ext "pgcrypto" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_tde" >}} |

> [!Note] manually upgraded PGRX from 0.16.1 to 0.17.0 by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.13` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_anon` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.13` | {{< bg "18" "pg_anon_18" "green" >}} {{< bg "17" "pg_anon_17" "green" >}} {{< bg "16" "pg_anon_16" "green" >}} {{< bg "15" "pg_anon_15" "green" >}} {{< bg "14" "pg_anon_14" "green" >}} | `pg_anon_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.13` | {{< bg "18" "postgresql-18-pg-anon" "green" >}} {{< bg "17" "postgresql-17-pg-anon" "green" >}} {{< bg "16" "postgresql-16-pg-anon" "green" >}} {{< bg "15" "postgresql-15-pg-anon" "green" >}} {{< bg "14" "postgresql-14-pg-anon" "green" >}} | `postgresql-$v-pg-anon` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.13" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_18` | `3.0.13` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_anon_18-3.0.13-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_18-3.0.13-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_18` | `3.0.13` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_18-3.0.13-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_18-3.0.13-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_18` | `3.0.13` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.3 MiB | [pg_anon_18-3.0.13-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_18-3.0.13-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_18` | `3.0.13` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_18-3.0.13-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_18-3.0.13-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_18` | `3.0.13` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.3 MiB | [pg_anon_18-3.0.13-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_18-3.0.13-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_18` | `3.0.13` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_anon_18-3.0.13-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_18-3.0.13-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-anon` | `3.0.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-18-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-anon` | `3.0.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-18-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-18-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-anon` | `3.0.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-18-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-anon` | `3.0.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-18-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-18-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-anon` | `3.0.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-18-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-18-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-anon` | `3.0.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-18-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-anon` | `3.0.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-18-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-18-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-anon` | `3.0.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-18-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_17` | `3.0.13` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_anon_17-3.0.13-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_17-3.0.13-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_17` | `3.0.13` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_17-3.0.13-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_17-3.0.13-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_17` | `3.0.13` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.3 MiB | [pg_anon_17-3.0.13-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_17-3.0.13-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_17` | `3.0.13` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_anon_17-3.0.13-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_17-3.0.13-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_17` | `3.0.13` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.3 MiB | [pg_anon_17-3.0.13-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_17-3.0.13-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_17` | `3.0.13` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_anon_17-3.0.13-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_17-3.0.13-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-anon` | `3.0.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-anon` | `3.0.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-17-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-anon` | `3.0.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-17-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-anon` | `3.0.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-17-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-17-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-anon` | `3.0.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-17-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-anon` | `3.0.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-anon` | `3.0.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-17-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-anon` | `3.0.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_16` | `3.0.13` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_anon_16-3.0.13-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_16-3.0.13-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_16` | `3.0.13` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_16-3.0.13-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_16-3.0.13-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_16` | `3.0.13` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.3 MiB | [pg_anon_16-3.0.13-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_16-3.0.13-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_16` | `3.0.13` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_16-3.0.13-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_16-3.0.13-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_16` | `3.0.13` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.3 MiB | [pg_anon_16-3.0.13-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_16-3.0.13-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_16` | `3.0.13` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_16-3.0.13-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_16-3.0.13-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-anon` | `3.0.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-anon` | `3.0.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-16-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-anon` | `3.0.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-16-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-anon` | `3.0.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-16-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-16-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-anon` | `3.0.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-16-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-anon` | `3.0.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-anon` | `3.0.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-16-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-anon` | `3.0.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_15` | `3.0.13` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_anon_15-3.0.13-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_15-3.0.13-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_15` | `3.0.13` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_15-3.0.13-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_15-3.0.13-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_15` | `3.0.13` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.3 MiB | [pg_anon_15-3.0.13-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_15-3.0.13-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_15` | `3.0.13` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_anon_15-3.0.13-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_15-3.0.13-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_15` | `3.0.13` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.3 MiB | [pg_anon_15-3.0.13-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_15-3.0.13-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_15` | `3.0.13` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_anon_15-3.0.13-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_15-3.0.13-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-anon` | `3.0.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-anon` | `3.0.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-15-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-anon` | `3.0.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-15-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-anon` | `3.0.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-15-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-15-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-anon` | `3.0.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-15-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-anon` | `3.0.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-anon` | `3.0.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-15-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-anon` | `3.0.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_14` | `3.0.13` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_anon_14-3.0.13-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_14-3.0.13-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_14` | `3.0.13` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.1 MiB | [pg_anon_14-3.0.13-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_14-3.0.13-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_14` | `3.0.13` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.3 MiB | [pg_anon_14-3.0.13-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_14-3.0.13-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_14` | `3.0.13` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_anon_14-3.0.13-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_14-3.0.13-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_14` | `3.0.13` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.3 MiB | [pg_anon_14-3.0.13-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_14-3.0.13-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_14` | `3.0.13` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_anon_14-3.0.13-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_14-3.0.13-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-anon` | `3.0.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_3.0.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-anon` | `3.0.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-14-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_3.0.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-anon` | `3.0.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-14-pg-anon_3.0.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-anon` | `3.0.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-14-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-14-pg-anon_3.0.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-anon` | `3.0.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-14-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_3.0.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-anon` | `3.0.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_3.0.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-anon` | `3.0.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.1 MiB | [postgresql-14-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_3.0.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-anon` | `3.0.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_3.0.13-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitlab.com/dalibo/postgresql_anonymizer/" title="Repository" icon="link" subtitle="gitlab.com/dalibo/postgresql_anonymizer/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql_anonymizer-3.0.13.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_anon;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_anon;		# install via package name, for the active PG version
pig install anon;		# install by extension name, for the current active PG version

pig install anon -v 18;   # install for PG 18
pig install anon -v 17;   # install for PG 17
pig install anon -v 16;   # install for PG 16
pig install anon -v 15;   # install for PG 15
pig install anon -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'anon';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION anon;
```




## Usage

> [anon: Anonymization & Data Masking for PostgreSQL](https://gitlab.com/dalibo/postgresql_anonymizer/)

`postgresql_anonymizer` (extension name: `anon`) masks or replaces personally identifiable information (PII) using a declarative approach. Masking rules are defined directly in the database schema using security labels.

```sql
CREATE EXTENSION IF NOT EXISTS anon CASCADE;
SELECT anon.init();
```

### Declaring Masking Rules

```sql
SECURITY LABEL FOR anon ON COLUMN player.name
  IS 'MASKED WITH FUNCTION anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN player.id
  IS 'MASKED WITH VALUE NULL';
```

### Static Masking

Permanently replace PII in the database:

```sql
SECURITY LABEL FOR anon ON COLUMN customer.full_name
  IS 'MASKED WITH FUNCTION anon.fake_first_name() || '' '' || anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN customer.birth
  IS 'MASKED WITH FUNCTION anon.random_date_between(''1920-01-01''::DATE, now())';

SELECT anon.anonymize_database();
-- Also available: anon.anonymize_table(), anon.anonymize_column()
```

### Dynamic Masking

Hide PII from specific roles while others see original data:

```sql
SELECT anon.start_dynamic_masking();

CREATE ROLE skynet LOGIN;
SECURITY LABEL FOR anon ON ROLE skynet IS 'MASKED';

SECURITY LABEL FOR anon ON COLUMN people.lastname
  IS 'MASKED WITH FUNCTION anon.fake_last_name()';

SECURITY LABEL FOR anon ON COLUMN people.phone
  IS 'MASKED WITH FUNCTION anon.partial(phone, 2, $$******$$, 2)';
```

When `skynet` queries the table, masked data is returned automatically.

### Anonymous Dumps

```bash
pg_dump_anon.sh -h localhost -p 5432 -U bob bob_db > dump.sql
```

### Common Masking Functions

| Function | Description |
|----------|-------------|
| `anon.fake_first_name()` | Random first name |
| `anon.fake_last_name()` | Random last name |
| `anon.fake_company()` | Random company name |
| `anon.random_date_between(d1, d2)` | Random date in range |
| `anon.random_zip()` | Random zip code |
| `anon.partial(value, prefix, padding, suffix)` | Partial scrambling |
| `anon.random_string(n)` | Random string of length n |
| `anon.random_int_between(i1, i2)` | Random integer in range |
