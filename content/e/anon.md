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
| **7070** | {{< badge content="anon" link="https://gitlab.com/dalibo/postgresql_anonymizer/" >}} | {{< ext "anon" "pg_anon" >}} | `3.1.3` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `anon` |
|   **See Also**    | {{< ext "faker" >}} {{< ext "pgsodium" >}} {{< ext "pgcrypto" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_tde" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_anon` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.3` | {{< bg "18" "pg_anon_18" "green" >}} {{< bg "17" "pg_anon_17" "green" >}} {{< bg "16" "pg_anon_16" "green" >}} {{< bg "15" "pg_anon_15" "green" >}} {{< bg "14" "pg_anon_14" "green" >}} | `pg_anon_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.3` | {{< bg "18" "postgresql-18-pg-anon" "green" >}} {{< bg "17" "postgresql-17-pg-anon" "green" >}} {{< bg "16" "postgresql-16-pg-anon" "green" >}} {{< bg "15" "postgresql-15-pg-anon" "green" >}} {{< bg "14" "postgresql-14-pg-anon" "green" >}} | `postgresql-$v-pg-anon` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_18` | `3.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.3 MiB | [pg_anon_18-3.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_18-3.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_anon_18` | `3.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [pg_anon_18-3.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_18-3.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_anon_18` | `3.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_anon_18-3.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_18-3.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_anon_18` | `3.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_18-3.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_18-3.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_anon_18` | `3.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.1 MiB | [pg_anon_18-3.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_18-3.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_anon_18` | `3.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_18-3.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_18-3.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-anon` | `3.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.9 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.9 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-anon` | `3.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-18-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-18-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_17` | `3.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.3 MiB | [pg_anon_17-3.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_17-3.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_anon_17` | `3.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [pg_anon_17-3.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_17-3.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_anon_17` | `3.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_anon_17-3.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_17-3.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_anon_17` | `3.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_17-3.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_17-3.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_anon_17` | `3.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.1 MiB | [pg_anon_17-3.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_17-3.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_anon_17` | `3.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [pg_anon_17-3.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_17-3.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-anon` | `3.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.9 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.9 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-anon` | `3.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-17-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-17-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_16` | `3.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.3 MiB | [pg_anon_16-3.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_16-3.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_anon_16` | `3.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [pg_anon_16-3.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_16-3.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_anon_16` | `3.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_anon_16-3.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_16-3.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_anon_16` | `3.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_16-3.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_16-3.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_anon_16` | `3.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.1 MiB | [pg_anon_16-3.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_16-3.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_anon_16` | `3.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_16-3.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_16-3.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-anon` | `3.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.9 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.9 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-anon` | `3.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-16-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-16-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_15` | `3.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.3 MiB | [pg_anon_15-3.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_15-3.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_anon_15` | `3.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [pg_anon_15-3.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_15-3.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_anon_15` | `3.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_anon_15-3.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_15-3.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_anon_15` | `3.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_15-3.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_15-3.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_anon_15` | `3.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.1 MiB | [pg_anon_15-3.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_15-3.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_anon_15` | `3.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.9 MiB | [pg_anon_15-3.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_15-3.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-anon` | `3.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.3 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.9 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.9 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-anon` | `3.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-15-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-15-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_14` | `3.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.3 MiB | [pg_anon_14-3.1.3-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_14-3.1.3-2PIGSTY.el8.x86_64.rpm) |
| `pg_anon_14` | `3.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.9 MiB | [pg_anon_14-3.1.3-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_14-3.1.3-2PIGSTY.el8.aarch64.rpm) |
| `pg_anon_14` | `3.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_anon_14-3.1.3-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_14-3.1.3-2PIGSTY.el9.x86_64.rpm) |
| `pg_anon_14` | `3.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_14-3.1.3-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_14-3.1.3-2PIGSTY.el9.aarch64.rpm) |
| `pg_anon_14` | `3.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.1 MiB | [pg_anon_14-3.1.3-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_14-3.1.3-2PIGSTY.el10.x86_64.rpm) |
| `pg_anon_14` | `3.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_14-3.1.3-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_14-3.1.3-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-anon` | `3.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.7 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.7 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.3 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.9 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.9 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-anon` | `3.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.7 MiB | [postgresql-14-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-14-pg-anon_3.1.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitlab.com/dalibo/postgresql_anonymizer/" title="Repository" icon="link" subtitle="gitlab.com/dalibo/postgresql_anonymizer/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql_anonymizer-3.1.3.tar.gz" >}}
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

Sources:

- [PostgreSQL Anonymizer 3.1.3 README](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/README.md)
- [Masking functions](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/docs/masking_functions.md)
- [3.1.3 changelog](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.1.3/CHANGELOG.md)
- [Official documentation](https://postgresql-anonymizer.readthedocs.io/en/stable/)

`anon` is PostgreSQL Anonymizer. It applies declarative masking rules for protected query access, produces anonymized data sets, and provides pseudonymization and randomized-response helpers. Use it when realistic data must remain useful without exposing the original sensitive values; treat masking policy, role grants, and access to the unmasked database as part of the security boundary.

### Core Workflow

Load `anon` for sessions in the target database, install the extension, and enable transparent dynamic masking. New connections pick up database-level settings.

```sql
ALTER DATABASE app SET session_preload_libraries = 'anon';

\connect app
CREATE EXTENSION anon;
ALTER DATABASE app SET anon.transparent_dynamic_masking = true;
```

Mark a login as masked and attach masking rules to sensitive columns:

```sql
CREATE ROLE reporting LOGIN;
SECURITY LABEL FOR anon ON ROLE reporting IS 'MASKED';
GRANT pg_read_all_data TO reporting;

SECURITY LABEL FOR anon ON COLUMN customer.last_name
  IS 'MASKED WITH FUNCTION anon.dummy_last_name()';
SECURITY LABEL FOR anon ON COLUMN customer.phone
  IS 'MASKED WITH FUNCTION anon.partial(phone, 2, $$******$$, 2)';
```

Queries made as `reporting` see the transformed values. Privileged users still see the originals, so do not grant masked roles a path around the policy.

### Masking Strategies

- **Dynamic masking** transforms results for roles labeled `MASKED` without rewriting the table.
- **Static masking** permanently rewrites selected data and is appropriate for disposable development or test copies.
- **Anonymous dumps and replicas** produce sanitized exports or downstream copies.
- **Masking views and data wrappers** expose a deliberately reduced or transformed projection.
- **Pseudonymization** uses deterministic transforms when joins or repeated values must remain consistent.

### Important Objects

- `anon.dummy_*`, `anon.random_*`, and `anon.partial(...)` generate or partially conceal values.
- `anon.hash(text)` and `anon.digest(text, text, text)` provide deterministic transformations. In 3.1.2 they were marked `RESTRICTED` to limit brute-force exposure.
- `anon.ldp_grrm(value, epsilon, max_v)` and `anon.ldp_grrm_pttt(value, truth_probability, max_v)` implement generalized randomized response for local differential privacy.
- `anon.ldp_truth_probability(...)` and `anon.ldp_lie_probability(...)` help inspect randomized-response probabilities.
- Security labels on roles and columns define who is masked and how each value is transformed.

### Operational Notes

`anon` is superuser-installed and non-relocatable. Test every policy with the same grants and connection path used by the intended consumer. Randomization is not automatically deterministic; use a confirmed pseudonymization function when stable equality is required. Static anonymization is destructive, so run it on a copy and verify constraints and application behavior afterward.

Version 3.1.3 reruns missing ARM builds and changes release metadata, with no new SQL workflow. The material delta since 3.1.1 is the 3.1.2 security hardening for `anon.hash` and `anon.digest`; deployments using those functions should upgrade rather than relying on the old labels.
