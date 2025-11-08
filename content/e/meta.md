---
title: "meta"
linkTitle: "meta"
description: "Normalized, friendlier system catalog for PostgreSQL"
weight: 6300
categories: ["STAT"]
width: full
---

Normalized, friendlier system catalog for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6300** | {{< badge content="meta" link="https://github.com/aquameta/meta" >}} | {{< ext "meta" "pg_meta" >}} | `0.4.0` | {{< category "STAT" >}} | {{< license "BSD 2-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_profile" >}} {{< ext "pg_tracing" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_stat_kcache" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_track_settings" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/meta" >}} | `0.4.0` | {{< bg "18" "pg_meta_18" "green" >}} {{< bg "17" "pg_meta_17" "green" >}} {{< bg "16" "pg_meta_16" "green" >}} {{< bg "15" "pg_meta_15" "green" >}} {{< bg "14" "pg_meta_14" "green" >}} {{< bg "13" "pg_meta_13" "green" >}} | `pg_meta_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/meta" >}} | `0.4.0` | {{< bg "18" "postgresql-18-pg-meta" "green" >}} {{< bg "17" "postgresql-17-pg-meta" "green" >}} {{< bg "16" "postgresql-16-pg-meta" "green" >}} {{< bg "15" "postgresql-15-pg-meta" "green" >}} {{< bg "14" "postgresql-14-pg-meta" "green" >}} {{< bg "13" "postgresql-13-pg-meta" "green" >}} | `postgresql-$v-pg-meta` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_meta_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-meta : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-meta : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-meta : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-meta : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-meta : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-meta : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-meta : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-meta : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-meta : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_meta_18` | 0.4.0 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_meta_18-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_meta_18-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_meta_18` | 0.4.0 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_meta_18-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_meta_18-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_meta_18` | 0.4.0 | `el9.x86_64` | pigsty | 15.8 KiB | [pg_meta_18-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_meta_18-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_meta_18` | 0.4.0 | `el9.aarch64` | pigsty | 15.7 KiB | [pg_meta_18-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_meta_18-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_meta_18` | 0.4.0 | `el10.x86_64` | pigsty | 15.9 KiB | [pg_meta_18-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_meta_18-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_meta_18` | 0.4.0 | `el10.aarch64` | pigsty | 15.8 KiB | [pg_meta_18-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_meta_18-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-meta` | 0.4.0 | `d12.x86_64` | pigsty | 11.4 KiB | [postgresql-18-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-18-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-meta` | 0.4.0 | `d12.aarch64` | pigsty | 11.4 KiB | [postgresql-18-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-18-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-meta` | 0.4.0 | `d13.x86_64` | pigsty | 11.4 KiB | [postgresql-18-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-18-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-meta` | 0.4.0 | `d13.aarch64` | pigsty | 11.4 KiB | [postgresql-18-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-18-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-meta` | 0.4.0 | `u22.x86_64` | pigsty | 11.4 KiB | [postgresql-18-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-18-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-meta` | 0.4.0 | `u22.aarch64` | pigsty | 11.4 KiB | [postgresql-18-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-18-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-meta` | 0.4.0 | `u24.x86_64` | pigsty | 11.4 KiB | [postgresql-18-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-18-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-meta` | 0.4.0 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-18-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-18-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_meta_17` | 0.4.0 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_meta_17-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_meta_17-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_meta_17` | 0.4.0 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_meta_17-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_meta_17-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_meta_17` | 0.4.0 | `el9.x86_64` | pigsty | 15.8 KiB | [pg_meta_17-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_meta_17-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_meta_17` | 0.4.0 | `el9.aarch64` | pigsty | 15.7 KiB | [pg_meta_17-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_meta_17-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_meta_17` | 0.4.0 | `el10.x86_64` | pigsty | 15.9 KiB | [pg_meta_17-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_meta_17-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_meta_17` | 0.4.0 | `el10.aarch64` | pigsty | 15.8 KiB | [pg_meta_17-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_meta_17-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-meta` | 0.4.0 | `d12.x86_64` | pigsty | 11.4 KiB | [postgresql-17-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-17-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-meta` | 0.4.0 | `d12.aarch64` | pigsty | 11.4 KiB | [postgresql-17-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-17-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-meta` | 0.4.0 | `d13.x86_64` | pigsty | 11.4 KiB | [postgresql-17-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-17-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-meta` | 0.4.0 | `d13.aarch64` | pigsty | 11.4 KiB | [postgresql-17-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-17-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-meta` | 0.4.0 | `u22.x86_64` | pigsty | 11.4 KiB | [postgresql-17-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-17-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-meta` | 0.4.0 | `u22.aarch64` | pigsty | 11.4 KiB | [postgresql-17-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-17-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-meta` | 0.4.0 | `u24.x86_64` | pigsty | 11.4 KiB | [postgresql-17-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-17-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-meta` | 0.4.0 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-17-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-17-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_meta_16` | 0.4.0 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_meta_16-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_meta_16-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_meta_16` | 0.4.0 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_meta_16-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_meta_16-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_meta_16` | 0.4.0 | `el9.x86_64` | pigsty | 15.8 KiB | [pg_meta_16-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_meta_16-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_meta_16` | 0.4.0 | `el9.aarch64` | pigsty | 15.7 KiB | [pg_meta_16-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_meta_16-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_meta_16` | 0.4.0 | `el10.x86_64` | pigsty | 15.9 KiB | [pg_meta_16-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_meta_16-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_meta_16` | 0.4.0 | `el10.aarch64` | pigsty | 15.8 KiB | [pg_meta_16-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_meta_16-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-meta` | 0.4.0 | `d12.x86_64` | pigsty | 11.4 KiB | [postgresql-16-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-16-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-meta` | 0.4.0 | `d12.aarch64` | pigsty | 11.4 KiB | [postgresql-16-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-16-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-meta` | 0.4.0 | `d13.x86_64` | pigsty | 11.4 KiB | [postgresql-16-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-16-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-meta` | 0.4.0 | `d13.aarch64` | pigsty | 11.4 KiB | [postgresql-16-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-16-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-meta` | 0.4.0 | `u22.x86_64` | pigsty | 11.4 KiB | [postgresql-16-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-16-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-meta` | 0.4.0 | `u22.aarch64` | pigsty | 11.4 KiB | [postgresql-16-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-16-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-meta` | 0.4.0 | `u24.x86_64` | pigsty | 11.4 KiB | [postgresql-16-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-16-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-meta` | 0.4.0 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-16-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-16-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_meta_15` | 0.4.0 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_meta_15-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_meta_15-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_meta_15` | 0.4.0 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_meta_15-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_meta_15-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_meta_15` | 0.4.0 | `el9.x86_64` | pigsty | 15.8 KiB | [pg_meta_15-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_meta_15-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_meta_15` | 0.4.0 | `el9.aarch64` | pigsty | 15.7 KiB | [pg_meta_15-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_meta_15-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_meta_15` | 0.4.0 | `el10.x86_64` | pigsty | 15.9 KiB | [pg_meta_15-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_meta_15-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_meta_15` | 0.4.0 | `el10.aarch64` | pigsty | 15.8 KiB | [pg_meta_15-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_meta_15-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-meta` | 0.4.0 | `d12.x86_64` | pigsty | 11.4 KiB | [postgresql-15-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-15-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-meta` | 0.4.0 | `d12.aarch64` | pigsty | 11.4 KiB | [postgresql-15-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-15-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-meta` | 0.4.0 | `d13.x86_64` | pigsty | 11.4 KiB | [postgresql-15-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-15-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-meta` | 0.4.0 | `d13.aarch64` | pigsty | 11.4 KiB | [postgresql-15-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-15-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-meta` | 0.4.0 | `u22.x86_64` | pigsty | 11.4 KiB | [postgresql-15-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-15-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-meta` | 0.4.0 | `u22.aarch64` | pigsty | 11.4 KiB | [postgresql-15-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-15-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-meta` | 0.4.0 | `u24.x86_64` | pigsty | 11.4 KiB | [postgresql-15-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-15-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-meta` | 0.4.0 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-15-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-15-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_meta_14` | 0.4.0 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_meta_14-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_meta_14-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_meta_14` | 0.4.0 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_meta_14-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_meta_14-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_meta_14` | 0.4.0 | `el9.x86_64` | pigsty | 15.8 KiB | [pg_meta_14-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_meta_14-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_meta_14` | 0.4.0 | `el9.aarch64` | pigsty | 15.7 KiB | [pg_meta_14-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_meta_14-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_meta_14` | 0.4.0 | `el10.x86_64` | pigsty | 15.9 KiB | [pg_meta_14-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_meta_14-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_meta_14` | 0.4.0 | `el10.aarch64` | pigsty | 15.8 KiB | [pg_meta_14-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_meta_14-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-meta` | 0.4.0 | `d12.x86_64` | pigsty | 11.4 KiB | [postgresql-14-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-14-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-meta` | 0.4.0 | `d12.aarch64` | pigsty | 11.4 KiB | [postgresql-14-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-14-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-meta` | 0.4.0 | `d13.x86_64` | pigsty | 11.4 KiB | [postgresql-14-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-14-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-meta` | 0.4.0 | `d13.aarch64` | pigsty | 11.4 KiB | [postgresql-14-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-14-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-meta` | 0.4.0 | `u22.x86_64` | pigsty | 11.4 KiB | [postgresql-14-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-14-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-meta` | 0.4.0 | `u22.aarch64` | pigsty | 11.4 KiB | [postgresql-14-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-14-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-meta` | 0.4.0 | `u24.x86_64` | pigsty | 11.4 KiB | [postgresql-14-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-14-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-meta` | 0.4.0 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-14-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-14-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_meta_13` | 0.4.0 | `el8.x86_64` | pigsty | 16.4 KiB | [pg_meta_13-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_meta_13-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_meta_13` | 0.4.0 | `el8.aarch64` | pigsty | 16.4 KiB | [pg_meta_13-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_meta_13-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_meta_13` | 0.4.0 | `el9.x86_64` | pigsty | 15.8 KiB | [pg_meta_13-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_meta_13-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_meta_13` | 0.4.0 | `el9.aarch64` | pigsty | 15.7 KiB | [pg_meta_13-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_meta_13-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_meta_13` | 0.4.0 | `el10.x86_64` | pigsty | 15.9 KiB | [pg_meta_13-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_meta_13-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_meta_13` | 0.4.0 | `el10.aarch64` | pigsty | 15.8 KiB | [pg_meta_13-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_meta_13-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-meta` | 0.4.0 | `d12.x86_64` | pigsty | 11.4 KiB | [postgresql-13-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-13-pg-meta_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-meta` | 0.4.0 | `d12.aarch64` | pigsty | 11.4 KiB | [postgresql-13-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-meta/postgresql-13-pg-meta_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-meta` | 0.4.0 | `d13.x86_64` | pigsty | 11.4 KiB | [postgresql-13-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-13-pg-meta_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-meta` | 0.4.0 | `d13.aarch64` | pigsty | 11.4 KiB | [postgresql-13-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-meta/postgresql-13-pg-meta_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-meta` | 0.4.0 | `u22.x86_64` | pigsty | 11.4 KiB | [postgresql-13-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-13-pg-meta_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-meta` | 0.4.0 | `u22.aarch64` | pigsty | 11.4 KiB | [postgresql-13-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-meta/postgresql-13-pg-meta_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-meta` | 0.4.0 | `u24.x86_64` | pigsty | 11.4 KiB | [postgresql-13-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-13-pg-meta_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-meta` | 0.4.0 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-13-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-meta/postgresql-13-pg-meta_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aquameta/meta" title="Repository" icon="github" subtitle="github.com/aquameta/meta" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="meta-0.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get meta; # get meta source code
pig build dep meta; # install build dependencies
pig build pkg meta; # build extension rpm or deb
pig build ext meta; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install meta; # install by extension name, for the current active PG version
pig ext install pg_meta; # install via package alias, for the active PG version
pig ext install meta -v 18;   # install for PG 18
pig ext install meta -v 17;   # install for PG 17
pig ext install meta -v 16;   # install for PG 16
pig ext install meta -v 15;   # install for PG 15
pig ext install meta -v 14;   # install for PG 14
pig ext install meta -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION meta;
```

