---
title: "anon"
linkTitle: "anon"
description: "PostgreSQL Anonymizer (anon) extension"
weight: 7050
categories: ["SEC"]
width: full
---

[**pg_anon**](https://gitlab.com/dalibo/postgresql_anonymizer/) : PostgreSQL Anonymizer (anon) extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7050** | {{< badge content="anon" link="https://gitlab.com/dalibo/postgresql_anonymizer/" >}} | {{< ext "anon" "pg_anon" >}} | `2.5.1` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "faker" >}} {{< ext "pgsodium" >}} {{< ext "pgcrypto" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_tde" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_anon` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.1` | {{< bg "18" "pg_anon_18" "green" >}} {{< bg "17" "pg_anon_17" "green" >}} {{< bg "16" "pg_anon_16" "green" >}} {{< bg "15" "pg_anon_15" "green" >}} {{< bg "14" "pg_anon_14" "green" >}} {{< bg "13" "pg_anon_13" "green" >}} | `pg_anon_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.1` | {{< bg "18" "postgresql-18-pg-anon" "green" >}} {{< bg "17" "postgresql-17-pg-anon" "green" >}} {{< bg "16" "postgresql-16-pg-anon" "green" >}} {{< bg "15" "postgresql-15-pg-anon" "green" >}} {{< bg "14" "postgresql-14-pg-anon" "green" >}} {{< bg "13" "postgresql-13-pg-anon" "green" >}} | `postgresql-$v-pg-anon` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "pg_anon_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-14-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.1" "postgresql-13-pg-anon : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_18` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [pg_anon_18-2.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_18-2.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_18` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.5 MiB | [pg_anon_18-2.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_18-2.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_18` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_anon_18-2.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_18-2.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_18` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.7 MiB | [pg_anon_18-2.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_18-2.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_18` | `2.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_anon_18-2.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_18-2.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_18` | `2.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [pg_anon_18-2.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_18-2.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-anon` | `2.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-18-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-anon` | `2.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-18-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-18-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-anon` | `2.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-18-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-anon` | `2.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-18-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-18-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-anon` | `2.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-18-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-18-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-anon` | `2.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-18-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-18-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-anon` | `2.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-18-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-anon` | `2.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-18-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-18-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_17` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [pg_anon_17-2.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_17-2.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_17` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.5 MiB | [pg_anon_17-2.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_17-2.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_17` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_anon_17-2.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_17-2.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_17` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.7 MiB | [pg_anon_17-2.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_17-2.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_17` | `2.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_anon_17-2.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_17-2.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_17` | `2.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [pg_anon_17-2.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_17-2.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-anon` | `2.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-17-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-anon` | `2.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-17-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-anon` | `2.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-17-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-17-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-anon` | `2.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-17-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-17-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-anon` | `2.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-17-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-anon` | `2.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-17-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-anon` | `2.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-anon` | `2.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-17-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_16` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [pg_anon_16-2.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_16-2.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_16` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.5 MiB | [pg_anon_16-2.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_16-2.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_16` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_anon_16-2.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_16-2.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_16` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.7 MiB | [pg_anon_16-2.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_16-2.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_16` | `2.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_anon_16-2.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_16-2.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_16` | `2.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [pg_anon_16-2.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_16-2.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-anon` | `2.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-16-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-anon` | `2.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-16-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-anon` | `2.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-16-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-16-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-anon` | `2.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-16-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-16-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-anon` | `2.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-16-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-anon` | `2.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-16-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-anon` | `2.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-anon` | `2.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.5 MiB | [postgresql-16-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_15` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [pg_anon_15-2.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_15-2.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_15` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.5 MiB | [pg_anon_15-2.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_15-2.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_15` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_anon_15-2.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_15-2.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_15` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.7 MiB | [pg_anon_15-2.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_15-2.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_15` | `2.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_anon_15-2.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_15-2.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_15` | `2.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [pg_anon_15-2.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_15-2.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-anon` | `2.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-15-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-anon` | `2.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-15-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-anon` | `2.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-15-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-15-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-anon` | `2.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-15-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-15-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-anon` | `2.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-15-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-anon` | `2.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-15-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-anon` | `2.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-anon` | `2.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.6 MiB | [postgresql-15-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_14` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [pg_anon_14-2.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_14-2.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_14` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.5 MiB | [pg_anon_14-2.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_14-2.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_14` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_anon_14-2.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_14-2.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_14` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.7 MiB | [pg_anon_14-2.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_14-2.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_14` | `2.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_anon_14-2.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_14-2.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_14` | `2.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [pg_anon_14-2.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_14-2.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-anon` | `2.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-14-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-anon` | `2.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-14-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-anon` | `2.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-14-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-14-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-anon` | `2.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-14-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-14-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-anon` | `2.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-14-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-anon` | `2.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-14-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-anon` | `2.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-anon` | `2.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.5 MiB | [postgresql-14-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_13` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.9 MiB | [pg_anon_13-2.5.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_13-2.5.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_anon_13` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.5 MiB | [pg_anon_13-2.5.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_13-2.5.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_anon_13` | `2.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_anon_13-2.5.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_13-2.5.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_anon_13` | `2.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.7 MiB | [pg_anon_13-2.5.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_13-2.5.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_anon_13` | `2.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_anon_13-2.5.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_13-2.5.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_anon_13` | `2.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [pg_anon_13-2.5.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_13-2.5.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-anon` | `2.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-13-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-13-pg-anon_2.5.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-anon` | `2.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.2 MiB | [postgresql-13-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-13-pg-anon_2.5.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-anon` | `2.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-13-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-13-pg-anon_2.5.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-anon` | `2.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.2 MiB | [postgresql-13-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-13-pg-anon_2.5.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-anon` | `2.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.9 MiB | [postgresql-13-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-13-pg-anon_2.5.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-anon` | `2.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.6 MiB | [postgresql-13-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-13-pg-anon_2.5.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-anon` | `2.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-13-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-13-pg-anon_2.5.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-anon` | `2.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.5 MiB | [postgresql-13-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-13-pg-anon_2.5.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitlab.com/dalibo/postgresql_anonymizer/" title="Repository" icon="link" subtitle="gitlab.com/dalibo/postgresql_anonymizer/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql_anonymizer-2.5.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_anon;		# build rpm / deb with pig
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
pig install anon -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'anon';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION anon;
```
