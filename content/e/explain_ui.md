---
title: "explain_ui"
linkTitle: "explain_ui"
description: "easily jump into a visual plan UI for any SQL query"
weight: 6370
categories: ["STAT"]
width: full
---

easily jump into a visual plan UI for any SQL query


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6370** | {{< badge content="explain_ui" link="https://github.com/davidgomes/pg-explain-ui" >}} | {{< ext "explain_ui" "pg_explain_ui" >}} | `0.0.2` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_store_plans" >}} {{< ext "pg_profile" >}} {{< ext "powa" >}} |

> [!Note] pgrx=0.16.1, manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/explain_ui" >}} | `0.0.2` | {{< bg "18" "pg_explain_ui_18" "green" >}} {{< bg "17" "pg_explain_ui_17" "green" >}} {{< bg "16" "pg_explain_ui_16" "green" >}} {{< bg "15" "pg_explain_ui_15" "green" >}} {{< bg "14" "pg_explain_ui_14" "green" >}} {{< bg "13" "pg_explain_ui_13" "green" >}} | `pg_explain_ui_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/explain_ui" >}} | `0.0.2` | {{< bg "18" "postgresql-18-pg-explain-ui" "green" >}} {{< bg "17" "postgresql-17-pg-explain-ui" "green" >}} {{< bg "16" "postgresql-16-pg-explain-ui" "green" >}} {{< bg "15" "postgresql-15-pg-explain-ui" "green" >}} {{< bg "14" "postgresql-14-pg-explain-ui" "green" >}} {{< bg "13" "postgresql-13-pg-explain-ui" "green" >}} | `postgresql-$v-pg-explain-ui` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_explain_ui_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_explain_ui_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_explain_ui_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_explain_ui_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_explain_ui_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_explain_ui_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_explain_ui_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_explain_ui_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-explain-ui : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.0" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-explain-ui : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.0" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-explain-ui : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-explain-ui : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-explain-ui : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-explain-ui : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.0" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-explain-ui : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.0" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-explain-ui : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.0" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-explain-ui : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.0" "postgresql-17-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-16-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-15-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-14-pg-explain-ui : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.0" "postgresql-13-pg-explain-ui : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_17` | 0.0.1 | `el8.x86_64` | pigsty | 990.1 KiB | [pg_explain_ui_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_17` | 0.0.1 | `el8.aarch64` | pigsty | 928.2 KiB | [pg_explain_ui_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_17` | 0.0.1 | `el9.x86_64` | pigsty | 987.3 KiB | [pg_explain_ui_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_17` | 0.0.1 | `el9.aarch64` | pigsty | 979.5 KiB | [pg_explain_ui_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-explain-ui` | 0.0.0 | `d12.x86_64` | pigsty | 792.6 KiB | [postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-explain-ui` | 0.0.0 | `d12.aarch64` | pigsty | 708.9 KiB | [postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-explain-ui` | 0.0.0 | `u22.x86_64` | pigsty | 874.4 KiB | [postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-explain-ui` | 0.0.0 | `u22.aarch64` | pigsty | 838.8 KiB | [postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-explain-ui` | 0.0.0 | `u24.x86_64` | pigsty | 866.9 KiB | [postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-explain-ui` | 0.0.0 | `u24.aarch64` | pigsty | 829.3 KiB | [postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-17-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_16` | 0.0.1 | `el8.x86_64` | pigsty | 990.4 KiB | [pg_explain_ui_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_16` | 0.0.1 | `el8.aarch64` | pigsty | 928.2 KiB | [pg_explain_ui_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_16` | 0.0.1 | `el9.x86_64` | pigsty | 989.7 KiB | [pg_explain_ui_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_16` | 0.0.1 | `el9.aarch64` | pigsty | 980.1 KiB | [pg_explain_ui_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-explain-ui` | 0.0.0 | `d12.x86_64` | pigsty | 792.6 KiB | [postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-explain-ui` | 0.0.0 | `d12.aarch64` | pigsty | 708.8 KiB | [postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-explain-ui` | 0.0.0 | `u22.x86_64` | pigsty | 875.3 KiB | [postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-explain-ui` | 0.0.0 | `u22.aarch64` | pigsty | 838.6 KiB | [postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-explain-ui` | 0.0.0 | `u24.x86_64` | pigsty | 866.6 KiB | [postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-explain-ui` | 0.0.0 | `u24.aarch64` | pigsty | 827.5 KiB | [postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-16-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_15` | 0.0.1 | `el8.x86_64` | pigsty | 990.1 KiB | [pg_explain_ui_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_15` | 0.0.1 | `el8.aarch64` | pigsty | 928.1 KiB | [pg_explain_ui_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_15` | 0.0.1 | `el9.x86_64` | pigsty | 989.2 KiB | [pg_explain_ui_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_15` | 0.0.1 | `el9.aarch64` | pigsty | 979.4 KiB | [pg_explain_ui_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-explain-ui` | 0.0.0 | `d12.x86_64` | pigsty | 792.6 KiB | [postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-explain-ui` | 0.0.0 | `d12.aarch64` | pigsty | 710.9 KiB | [postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-explain-ui` | 0.0.0 | `u22.x86_64` | pigsty | 875.3 KiB | [postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-explain-ui` | 0.0.0 | `u22.aarch64` | pigsty | 839.0 KiB | [postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-explain-ui` | 0.0.0 | `u24.x86_64` | pigsty | 866.5 KiB | [postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-explain-ui` | 0.0.0 | `u24.aarch64` | pigsty | 828.3 KiB | [postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-15-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_14` | 0.0.1 | `el8.x86_64` | pigsty | 990.4 KiB | [pg_explain_ui_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_14` | 0.0.1 | `el8.aarch64` | pigsty | 928.3 KiB | [pg_explain_ui_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_14` | 0.0.1 | `el9.x86_64` | pigsty | 989.9 KiB | [pg_explain_ui_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_14` | 0.0.1 | `el9.aarch64` | pigsty | 979.6 KiB | [pg_explain_ui_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-explain-ui` | 0.0.0 | `d12.x86_64` | pigsty | 792.6 KiB | [postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-explain-ui` | 0.0.0 | `d12.aarch64` | pigsty | 708.7 KiB | [postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-explain-ui` | 0.0.0 | `u22.x86_64` | pigsty | 875.4 KiB | [postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-explain-ui` | 0.0.0 | `u22.aarch64` | pigsty | 840.6 KiB | [postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-explain-ui` | 0.0.0 | `u24.x86_64` | pigsty | 866.6 KiB | [postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-explain-ui` | 0.0.0 | `u24.aarch64` | pigsty | 828.1 KiB | [postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-14-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_explain_ui_13` | 0.0.1 | `el8.x86_64` | pigsty | 990.4 KiB | [pg_explain_ui_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_explain_ui_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_explain_ui_13` | 0.0.1 | `el8.aarch64` | pigsty | 928.2 KiB | [pg_explain_ui_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_explain_ui_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_explain_ui_13` | 0.0.1 | `el9.x86_64` | pigsty | 991.8 KiB | [pg_explain_ui_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_explain_ui_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_explain_ui_13` | 0.0.1 | `el9.aarch64` | pigsty | 979.6 KiB | [pg_explain_ui_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_explain_ui_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-explain-ui` | 0.0.0 | `d12.x86_64` | pigsty | 1.4 KiB | [postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-explain-ui` | 0.0.0 | `d12.aarch64` | pigsty | 1.4 KiB | [postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-explain-ui` | 0.0.0 | `u22.x86_64` | pigsty | 1.2 KiB | [postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-explain-ui` | 0.0.0 | `u22.aarch64` | pigsty | 1.2 KiB | [postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-explain-ui` | 0.0.0 | `u24.x86_64` | pigsty | 1.2 KiB | [postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-explain-ui` | 0.0.0 | `u24.aarch64` | pigsty | 1.2 KiB | [postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-explain-ui/postgresql-13-pg-explain-ui_0.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/davidgomes/pg-explain-ui" title="Repository" icon="github" subtitle="github.com/davidgomes/pg-explain-ui" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_explain_ui-0.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get explain_ui; # get explain_ui source code
pig build dep explain_ui; # install build dependencies
pig build pkg explain_ui; # build extension rpm or deb
pig build ext explain_ui; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install explain_ui; # install by extension name, for the current active PG version
pig ext install pg_explain_ui; # install via package alias, for the active PG version
pig ext install explain_ui -v 18;   # install for PG 18
pig ext install explain_ui -v 17;   # install for PG 17
pig ext install explain_ui -v 16;   # install for PG 16
pig ext install explain_ui -v 15;   # install for PG 15
pig ext install explain_ui -v 14;   # install for PG 14
pig ext install explain_ui -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION explain_ui;
```

