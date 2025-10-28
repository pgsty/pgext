---
title: "anon"
linkTitle: "anon"
description: "PostgreSQL Anonymizer (anon) extension"
weight: 7050
categories: ["SEC"]
width: full
---

PostgreSQL Anonymizer (anon) extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7050** | {{< badge content="anon" link="https://gitlab.com/dalibo/postgresql_anonymizer/" >}} | {{< ext "anon" "pg_anon" >}} | `2.4.1` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "faker" >}} {{< ext "pgsodium" >}} {{< ext "pg_tde" >}} {{< ext "pgcrypto" >}} {{< ext "pg_permissions" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_drop_events" >}} |

> [!Note] pgrx=0.16.1, manual update from 0.16.0


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/anon" >}} | `2.4.1` | {{< bg "18" "pg_anon_18" "green" >}} {{< bg "17" "pg_anon_17" "green" >}} {{< bg "16" "pg_anon_16" "green" >}} {{< bg "15" "pg_anon_15" "green" >}} {{< bg "14" "pg_anon_14" "green" >}} {{< bg "13" "pg_anon_13" "green" >}} | `pg_anon_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/anon" >}} | `2.4.1` | {{< bg "18" "postgresql-18-pg-anon" "green" >}} {{< bg "17" "postgresql-17-pg-anon" "green" >}} {{< bg "16" "postgresql-16-pg-anon" "green" >}} {{< bg "15" "postgresql-15-pg-anon" "green" >}} {{< bg "14" "postgresql-14-pg-anon" "green" >}} {{< bg "13" "postgresql-13-pg-anon" "green" >}} | `postgresql-$v-pg-anon` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_anon_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_anon_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_anon_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_anon_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "pg_anon_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_anon_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_anon_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_anon_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-anon : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-anon : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-anon : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-anon : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-anon : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-anon : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-anon : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-anon : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-anon : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.3.0" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.3.0" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_17` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_17-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_17-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_17` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_17-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_17-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_17` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_17-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_17-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_17` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_17-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_17-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.3 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.6 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_16` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_16-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_16-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_16` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_16-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_16-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_16` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_16-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_16-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_16` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_16-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_16-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.3 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.6 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_15` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_15-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_15-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_15` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_15-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_15-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_15` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_15-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_15-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_15` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_15-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_15-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.3 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.6 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
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
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_13` | 2.3.0 | `el8.x86_64` | pigsty | 2.8 MiB | [pg_anon_13-2.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_13-2.3.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_13` | 2.3.0 | `el8.aarch64` | pigsty | 2.5 MiB | [pg_anon_13-2.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_13-2.3.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_13` | 2.3.0 | `el9.x86_64` | pigsty | 2.8 MiB | [pg_anon_13-2.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_13-2.3.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_13` | 2.3.0 | `el9.aarch64` | pigsty | 2.7 MiB | [pg_anon_13-2.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_13-2.3.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-anon` | 2.3.0 | `d12.x86_64` | pigsty | 2.4 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `d12.aarch64` | pigsty | 2.1 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `u22.x86_64` | pigsty | 2.7 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `u22.aarch64` | pigsty | 2.5 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `u24.x86_64` | pigsty | 2.6 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-anon` | 2.3.0 | `u24.aarch64` | pigsty | 2.4 MiB | [postgresql-13-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-13-pg-anon_2.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitlab.com/dalibo/postgresql_anonymizer/" title="Repository" icon="link" subtitle="gitlab.com/dalibo/postgresql_anonymizer/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql_anonymizer-2.4.1.tar.gz" >}}
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
pig ext install anon -v 18;   # install for PG 18
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

