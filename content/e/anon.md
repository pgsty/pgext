---
title: "anon"
linkTitle: "anon"
description: "PostgreSQL Anonymizer (anon) extension"
weight: 7070
categories: ["SEC"]
languages: ["Rust"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_anon**](https://gitlab.com/dalibo/postgresql_anonymizer/) : PostgreSQL Anonymizer (anon) extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7070** | {{< badge content="anon" link="https://gitlab.com/dalibo/postgresql_anonymizer/" >}} | {{< ext "anon" "pg_anon" >}} | `3.3.2` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `anon` |
|   **See Also**    | {{< ext "column_encrypt" >}} {{< ext "pg_tde" >}} {{< ext "supabase_vault" >}} {{< ext "pg_enigma" >}} {{< ext "pgsodium" >}} {{< ext "pgcryptokey" >}} |

> [!Note] No in-place update; drop/recreate and migrate masking rules. The 3.3 series adds security hardening; 3.3.2 restores the ARM build.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.3.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_anon` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.3.2` | {{< bg "18" "pg_anon_18" "green" >}} {{< bg "17" "pg_anon_17" "green" >}} {{< bg "16" "pg_anon_16" "green" >}} {{< bg "15" "pg_anon_15" "green" >}} {{< bg "14" "pg_anon_14" "green" >}} | `pg_anon_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.3` | {{< bg "18" "postgresql-18-pg-anon" "green" >}} {{< bg "17" "postgresql-17-pg-anon" "green" >}} {{< bg "16" "postgresql-16-pg-anon" "green" >}} {{< bg "15" "postgresql-15-pg-anon" "green" >}} {{< bg "14" "postgresql-14-pg-anon" "green" >}} | `postgresql-$v-pg-anon` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.3.2" "pg_anon_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-18-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-17-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-16-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-15-pg-anon : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.3" "postgresql-14-pg-anon : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_18` | `3.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.4 MiB | [pg_anon_18-3.3.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_18-3.3.2-1PGSTY.el8.x86_64.rpm) |
| `pg_anon_18` | `3.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_18-3.3.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_18-3.3.2-1PGSTY.el8.aarch64.rpm) |
| `pg_anon_18` | `3.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.2 MiB | [pg_anon_18-3.3.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_18-3.3.2-1PGSTY.el9.x86_64.rpm) |
| `pg_anon_18` | `3.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_18-3.3.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_18-3.3.2-1PGSTY.el9.aarch64.rpm) |
| `pg_anon_18` | `3.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_anon_18-3.3.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_18-3.3.2-1PGSTY.el10.x86_64.rpm) |
| `pg_anon_18` | `3.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_18-3.3.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_18-3.3.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-anon` | `3.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.1 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.0 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.0 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-anon` | `3.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-18-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-18-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_17` | `3.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.4 MiB | [pg_anon_17-3.3.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_17-3.3.2-1PGSTY.el8.x86_64.rpm) |
| `pg_anon_17` | `3.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_17-3.3.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_17-3.3.2-1PGSTY.el8.aarch64.rpm) |
| `pg_anon_17` | `3.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.2 MiB | [pg_anon_17-3.3.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_17-3.3.2-1PGSTY.el9.x86_64.rpm) |
| `pg_anon_17` | `3.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_17-3.3.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_17-3.3.2-1PGSTY.el9.aarch64.rpm) |
| `pg_anon_17` | `3.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_anon_17-3.3.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_17-3.3.2-1PGSTY.el10.x86_64.rpm) |
| `pg_anon_17` | `3.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_17-3.3.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_17-3.3.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-anon` | `3.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.0 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.0 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-anon` | `3.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-17-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-17-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_16` | `3.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.4 MiB | [pg_anon_16-3.3.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_16-3.3.2-1PGSTY.el8.x86_64.rpm) |
| `pg_anon_16` | `3.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_16-3.3.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_16-3.3.2-1PGSTY.el8.aarch64.rpm) |
| `pg_anon_16` | `3.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.2 MiB | [pg_anon_16-3.3.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_16-3.3.2-1PGSTY.el9.x86_64.rpm) |
| `pg_anon_16` | `3.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_16-3.3.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_16-3.3.2-1PGSTY.el9.aarch64.rpm) |
| `pg_anon_16` | `3.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_anon_16-3.3.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_16-3.3.2-1PGSTY.el10.x86_64.rpm) |
| `pg_anon_16` | `3.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_16-3.3.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_16-3.3.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-anon` | `3.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.0 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.0 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-anon` | `3.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-16-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-16-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_15` | `3.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.4 MiB | [pg_anon_15-3.3.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_15-3.3.2-1PGSTY.el8.x86_64.rpm) |
| `pg_anon_15` | `3.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_15-3.3.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_15-3.3.2-1PGSTY.el8.aarch64.rpm) |
| `pg_anon_15` | `3.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.2 MiB | [pg_anon_15-3.3.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_15-3.3.2-1PGSTY.el9.x86_64.rpm) |
| `pg_anon_15` | `3.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_15-3.3.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_15-3.3.2-1PGSTY.el9.aarch64.rpm) |
| `pg_anon_15` | `3.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_anon_15-3.3.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_15-3.3.2-1PGSTY.el10.x86_64.rpm) |
| `pg_anon_15` | `3.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_15-3.3.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_15-3.3.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-anon` | `3.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.0 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.0 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-anon` | `3.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-15-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-15-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_anon_14` | `3.3.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.4 MiB | [pg_anon_14-3.3.2-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_anon_14-3.3.2-1PGSTY.el8.x86_64.rpm) |
| `pg_anon_14` | `3.3.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_anon_14-3.3.2-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_anon_14-3.3.2-1PGSTY.el8.aarch64.rpm) |
| `pg_anon_14` | `3.3.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.2 MiB | [pg_anon_14-3.3.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_anon_14-3.3.2-1PGSTY.el9.x86_64.rpm) |
| `pg_anon_14` | `3.3.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.0 MiB | [pg_anon_14-3.3.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_anon_14-3.3.2-1PGSTY.el9.aarch64.rpm) |
| `pg_anon_14` | `3.3.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.2 MiB | [pg_anon_14-3.3.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_anon_14-3.3.2-1PGSTY.el10.x86_64.rpm) |
| `pg_anon_14` | `3.3.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.0 MiB | [pg_anon_14-3.3.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_anon_14-3.3.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-anon` | `3.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.4 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.4 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.0 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.0 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.0 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-anon` | `3.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.8 MiB | [postgresql-14-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-anon/postgresql-14-pg-anon_3.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitlab.com/dalibo/postgresql_anonymizer/" title="Repository" icon="link" subtitle="gitlab.com/dalibo/postgresql_anonymizer/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgresql_anonymizer-3.3.2.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_anon;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_anon;		# install via package name, for the active PG version
pig install anon;		# install by extension name, for the current active PG version

pig install anon -v 18;   # install for PG 18
pig install anon -v 17;   # install for PG 17
pig install anon -v 16;   # install for PG 16
pig install anon -v 15;   # install for PG 15
pig install anon -v 14;   # install for PG 14

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'anon';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION anon;
```

## Usage

Sources:

- [PostgreSQL Anonymizer 3.3.2 README](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.3.2/README.md)
- [Masking functions](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.3.2/docs/masking_functions.md)
- [3.3.2 changelog](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.3.2/CHANGELOG.md)
- [Official documentation](https://postgresql-anonymizer.readthedocs.io/en/stable/)
- [3.3.2 upgrade guide](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.3.2/docs/UPGRADE.md)
- [3.3 series release announcement](https://gitlab.com/dalibo/postgresql_anonymizer/-/blob/3.3.2/NEWS.md)

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

### Version 3.3.2

Version 3.3.2 is the follow-up release to the 3.3 security series and restores the missing ARM build. The 3.3 series forbids subqueries in masking rules, disables inlining for transparent dynamic masking, rejects SQL-standard-body functions for masked roles, and changes `anon.salt`. There is no in-place update path: preserve and migrate masking rules, then drop and recreate the extension in each database as directed by upstream.

Version 3.2.0 introduces localized deterministic `anon.seeded_*` functions; the older `anon.pseudo_*` family is deprecated. It fixes three upstream-reported security vulnerabilities and defaults to refusing masking operations executed by superusers. Use a dedicated non-superuser masking role. Setting `anon.nosuperuser` to false removes that barrier and is not the recommended upgrade path. The masking-rule JSON format changed; re-export rules after upgrading. This documentation describes release 3.2.3.

Version 3.2.3 is a security update for CVE-2026-97469: dynamic masking now checks for `RESTRICTED` functions inside subqueries. Upgrade the installed library, then follow the upstream drop-and-recreate procedure in each database; `ALTER EXTENSION ... UPDATE` is unsupported. Preserve and migrate masking rules before recreating the extension; review masked-role queries that previously reached restricted functions through subqueries.
