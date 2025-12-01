---
title: "explain_ui"
linkTitle: "explain_ui"
description: "easily jump into a visual plan UI for any SQL query"
weight: 6370
categories: ["STAT"]
width: full
---

[**pg_explain_ui**](https://github.com/davidgomes/pg-explain-ui) : easily jump into a visual plan UI for any SQL query


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6370** | {{< badge content="explain_ui" link="https://github.com/davidgomes/pg-explain-ui" >}} | {{< ext "explain_ui" "pg_explain_ui" >}} | `0.0.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_profile" >}} {{< ext "powa" >}} |

> [!Note] manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_explain_ui` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "pg_explain_ui_18" "green" >}} {{< bg "17" "pg_explain_ui_17" "green" >}} {{< bg "16" "pg_explain_ui_16" "green" >}} {{< bg "15" "pg_explain_ui_15" "green" >}} {{< bg "14" "pg_explain_ui_14" "green" >}} {{< bg "13" "pg_explain_ui_13" "green" >}} | `pg_explain_ui_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "postgresql-18-pg-explain-ui" "green" >}} {{< bg "17" "postgresql-17-pg-explain-ui" "green" >}} {{< bg "16" "postgresql-16-pg-explain-ui" "green" >}} {{< bg "15" "postgresql-15-pg-explain-ui" "green" >}} {{< bg "14" "postgresql-14-pg-explain-ui" "green" >}} {{< bg "13" "postgresql-13-pg-explain-ui" "green" >}} | `postgresql-$v-pg-explain-ui` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_18` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_18-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_18-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_18` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 935.6 KiB | [pg_explain_ui_18-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_18-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_18` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_18-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_18-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_18` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 988.3 KiB | [pg_explain_ui_18-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_18-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_explain_ui_18` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_18-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_explain_ui_18-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_explain_ui_18` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 987.6 KiB | [pg_explain_ui_18-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_explain_ui_18-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-explain-ui` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 880.7 KiB | [postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-explain-ui` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 717.0 KiB | [postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-explain-ui` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 879.9 KiB | [postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-explain-ui` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 716.6 KiB | [postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-explain-ui` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 981.7 KiB | [postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-explain-ui` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 846.9 KiB | [postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-explain-ui` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 968.9 KiB | [postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-explain-ui` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 836.3 KiB | [postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-18-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_17` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_17-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_17-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_17` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 935.6 KiB | [pg_explain_ui_17-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_17-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_17` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_17-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_17-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_17` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 989.0 KiB | [pg_explain_ui_17-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_17-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_explain_ui_17` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_17-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_explain_ui_17-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_explain_ui_17` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 988.7 KiB | [pg_explain_ui_17-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_explain_ui_17-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-explain-ui` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 879.8 KiB | [postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-explain-ui` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 716.2 KiB | [postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-explain-ui` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 880.0 KiB | [postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-explain-ui` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 716.9 KiB | [postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-explain-ui` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 981.3 KiB | [postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-explain-ui` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 846.8 KiB | [postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-explain-ui` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 970.2 KiB | [postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-explain-ui` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 836.2 KiB | [postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_16` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_16-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_16-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_16` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 935.6 KiB | [pg_explain_ui_16-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_16-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_16` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_16-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_16-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_16` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 988.1 KiB | [pg_explain_ui_16-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_16-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_explain_ui_16` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_16-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_explain_ui_16-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_explain_ui_16` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 988.7 KiB | [pg_explain_ui_16-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_explain_ui_16-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-explain-ui` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 880.4 KiB | [postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-explain-ui` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 716.7 KiB | [postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-explain-ui` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 880.2 KiB | [postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-explain-ui` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 716.3 KiB | [postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-explain-ui` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 979.8 KiB | [postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-explain-ui` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 846.7 KiB | [postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-explain-ui` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 970.4 KiB | [postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-explain-ui` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 837.1 KiB | [postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_15` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_15-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_15-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_15` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 935.5 KiB | [pg_explain_ui_15-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_15-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_15` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_15-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_15-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_15` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 987.8 KiB | [pg_explain_ui_15-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_15-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_explain_ui_15` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_15-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_explain_ui_15-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_explain_ui_15` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 987.8 KiB | [pg_explain_ui_15-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_explain_ui_15-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-explain-ui` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 880.2 KiB | [postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-explain-ui` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 716.6 KiB | [postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-explain-ui` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 880.1 KiB | [postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-explain-ui` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 717.1 KiB | [postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-explain-ui` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 979.6 KiB | [postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-explain-ui` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 847.4 KiB | [postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-explain-ui` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 969.8 KiB | [postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-explain-ui` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 837.1 KiB | [postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_14` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_14-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_14-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_14` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 935.5 KiB | [pg_explain_ui_14-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_14-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_14` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_14-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_14-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_14` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 987.3 KiB | [pg_explain_ui_14-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_14-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_explain_ui_14` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_14-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_explain_ui_14-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_explain_ui_14` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 987.6 KiB | [pg_explain_ui_14-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_explain_ui_14-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-explain-ui` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 879.2 KiB | [postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-explain-ui` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 717.2 KiB | [postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-explain-ui` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 879.9 KiB | [postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-explain-ui` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 716.6 KiB | [postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-explain-ui` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 978.2 KiB | [postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-explain-ui` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 847.2 KiB | [postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-explain-ui` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 969.7 KiB | [postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-explain-ui` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 837.8 KiB | [postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_13` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_13-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_13-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_13` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 935.4 KiB | [pg_explain_ui_13-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_13-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_13` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_13-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_13-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_13` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 988.3 KiB | [pg_explain_ui_13-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_13-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_explain_ui_13` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_explain_ui_13-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_explain_ui_13-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_explain_ui_13` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 987.7 KiB | [pg_explain_ui_13-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_explain_ui_13-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-explain-ui` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 880.3 KiB | [postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-explain-ui` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 717.2 KiB | [postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-explain-ui` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 880.9 KiB | [postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-explain-ui` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 716.9 KiB | [postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-explain-ui` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 979.1 KiB | [postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-explain-ui` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 847.1 KiB | [postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-explain-ui` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 970.2 KiB | [postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-explain-ui` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 838.0 KiB | [postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/davidgomes/pg-explain-ui" title="Repository" icon="github" subtitle="github.com/davidgomes/pg-explain-ui" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_explain_ui-0.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_explain_ui;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_explain_ui;		# install via package name, for the active PG version
pig install explain_ui;		# install by extension name, for the current active PG version

pig install explain_ui -v 18;   # install for PG 18
pig install explain_ui -v 17;   # install for PG 17
pig install explain_ui -v 16;   # install for PG 16
pig install explain_ui -v 15;   # install for PG 15
pig install explain_ui -v 14;   # install for PG 14
pig install explain_ui -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION explain_ui;
```
