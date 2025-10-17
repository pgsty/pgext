---
title: "pg_store_plans"
linkTitle: "pg_store_plans"
description: "track plan statistics of all SQL statements executed"
weight: 6250
categories: ["Stat"]
width: full
---

track plan statistics of all SQL statements executed

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6250** | {{< badge content="pg_store_plans" link="https://github.com/ossc-db/pg_store_plans" >}} | {{< ext "pg_store_plans" "pg_store_plans" >}} | `1.8` | {{< category "STAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_show_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "pre_prepare" >}} {{< ext "pg_stat_monitor" >}} {{< ext "explain_ui" >}} {{< ext "plprofiler" >}} |

> [!Note] pg18 breaks


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_store_plans" >}} | `1.8` | {{< badge content="18" color="red" alt="pg_store_plans_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_store_plans_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_store_plans" >}} | `1.8` | {{< badge content="18" color="red" alt="postgresql-18-pg-store-plan" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-store-plan` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_store_plans_18" >}}     | {{< pkg "pg_store_plans_17" "1.8" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_17-1.8-2PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_store_plans_16" "1.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_store_plans_16-1.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_store_plans_15" "1.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_store_plans_15-1.8-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_store_plans_14" "1.8" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_14-1.8-2PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_store_plans_18" >}}     | {{< pkg "pg_store_plans_17" "1.8" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_17-1.8-2PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_store_plans_16" "1.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_store_plans_16-1.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_store_plans_15" "1.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_store_plans_15-1.8-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_store_plans_14" "1.8" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_14-1.8-2PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_store_plans_18" >}}     | {{< pkg "pg_store_plans_17" "1.8" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_17-1.8-2PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_store_plans_16" "1.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_store_plans_16-1.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_store_plans_15" "1.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_store_plans_15-1.8-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_store_plans_14" "1.8" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_14-1.8-2PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_store_plans_18" >}}     | {{< pkg "pg_store_plans_17" "1.8" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_17-1.8-2PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_store_plans_16" "1.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_store_plans_16-1.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_store_plans_15" "1.8" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_store_plans_15-1.8-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_store_plans_14" "1.8" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_14-1.8-2PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-store-plan" >}}     | {{< pkg "postgresql-17-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-store-plan" >}}     | {{< pkg "postgresql-17-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-store-plan" >}}     | {{< pkg "postgresql-17-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-store-plan" >}}     | {{< pkg "postgresql-17-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-store-plan" >}}     | {{< pkg "postgresql-17-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-store-plan" >}}     | {{< pkg "postgresql-17-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-store-plan" "1.8" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_store_plans_17` | 1.8 | `el8.aarch64` | pigsty | 39.4 KiB | [pg_store_plans_17-1.8-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_17-1.8-2PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_17` | 1.8 | `el8.x86_64` | pigsty | 40.7 KiB | [pg_store_plans_17-1.8-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_17-1.8-2PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_17` | 1.8 | `el9.aarch64` | pigsty | 41.5 KiB | [pg_store_plans_17-1.8-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_17-1.8-2PIGSTY.el9.aarch64.rpm) |
| `pg_store_plans_17` | 1.8 | `el9.x86_64` | pigsty | 42.4 KiB | [pg_store_plans_17-1.8-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_17-1.8-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-store-plan` | 1.8 | `d12.aarch64` | pigsty | 116.3 KiB | [postgresql-17-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-store-plan` | 1.8 | `d12.x86_64` | pigsty | 119.2 KiB | [postgresql-17-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-store-plan` | 1.8 | `u22.aarch64` | pigsty | 127.3 KiB | [postgresql-17-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-store-plan` | 1.8 | `u22.x86_64` | pigsty | 128.9 KiB | [postgresql-17-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-store-plan` | 1.8 | `u24.x86_64` | pigsty | 114.0 KiB | [postgresql-17-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-store-plan` | 1.8 | `u24.aarch64` | pigsty | 112.5 KiB | [postgresql-17-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-17-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_store_plans_16` | 1.8 | `el8.x86_64` | pgdg | 45.4 KiB | [pg_store_plans_16-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_store_plans_16-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_store_plans_16` | 1.8 | `el8.aarch64` | pigsty | 39.4 KiB | [pg_store_plans_16-1.8-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_16-1.8-2PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_16` | 1.8 | `el8.aarch64` | pgdg | 44.1 KiB | [pg_store_plans_16-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_store_plans_16-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_store_plans_16` | 1.8 | `el8.x86_64` | pigsty | 40.7 KiB | [pg_store_plans_16-1.8-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_16-1.8-2PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_16` | 1.8 | `el9.aarch64` | pigsty | 41.4 KiB | [pg_store_plans_16-1.8-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_16-1.8-2PIGSTY.el9.aarch64.rpm) |
| `pg_store_plans_16` | 1.8 | `el9.aarch64` | pgdg | 45.1 KiB | [pg_store_plans_16-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_store_plans_16-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_store_plans_16` | 1.8 | `el9.x86_64` | pgdg | 46.1 KiB | [pg_store_plans_16-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_store_plans_16-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_store_plans_16` | 1.8 | `el9.x86_64` | pigsty | 42.3 KiB | [pg_store_plans_16-1.8-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_16-1.8-2PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-pg-store-plan` | 1.8 | `d12.x86_64` | pigsty | 119.0 KiB | [postgresql-16-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-store-plan` | 1.8 | `d12.aarch64` | pigsty | 116.2 KiB | [postgresql-16-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-store-plan` | 1.8 | `u22.aarch64` | pigsty | 127.0 KiB | [postgresql-16-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-store-plan` | 1.8 | `u22.x86_64` | pigsty | 128.7 KiB | [postgresql-16-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-store-plan` | 1.8 | `u24.x86_64` | pigsty | 113.9 KiB | [postgresql-16-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-store-plan` | 1.8 | `u24.aarch64` | pigsty | 112.6 KiB | [postgresql-16-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-16-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_store_plans_15` | 1.8 | `el8.aarch64` | pigsty | 40.3 KiB | [pg_store_plans_15-1.8-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_15-1.8-2PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_15` | 1.8 | `el8.x86_64` | pgdg | 46.3 KiB | [pg_store_plans_15-1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_store_plans_15-1.8-1PGDG.rhel8.x86_64.rpm) |
| `pg_store_plans_15` | 1.8 | `el8.x86_64` | pigsty | 41.7 KiB | [pg_store_plans_15-1.8-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_15-1.8-2PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_15` | 1.8 | `el8.aarch64` | pgdg | 45.0 KiB | [pg_store_plans_15-1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_store_plans_15-1.8-1PGDG.rhel8.aarch64.rpm) |
| `pg_store_plans_15` | 1.7 | `el8.x86_64` | pgdg | 46.3 KiB | [pg_store_plans_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_store_plans_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pg_store_plans_15` | 1.7 | `el8.aarch64` | pgdg | 44.8 KiB | [pg_store_plans_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_store_plans_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pg_store_plans_15` | 1.8 | `el9.x86_64` | pigsty | 43.9 KiB | [pg_store_plans_15-1.8-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_15-1.8-2PIGSTY.el9.x86_64.rpm) |
| `pg_store_plans_15` | 1.8 | `el9.aarch64` | pigsty | 43.1 KiB | [pg_store_plans_15-1.8-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_15-1.8-2PIGSTY.el9.aarch64.rpm) |
| `pg_store_plans_15` | 1.8 | `el9.aarch64` | pgdg | 46.1 KiB | [pg_store_plans_15-1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_store_plans_15-1.8-1PGDG.rhel9.aarch64.rpm) |
| `pg_store_plans_15` | 1.8 | `el9.x86_64` | pgdg | 47.2 KiB | [pg_store_plans_15-1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_store_plans_15-1.8-1PGDG.rhel9.x86_64.rpm) |
| `pg_store_plans_15` | 1.7 | `el9.x86_64` | pgdg | 47.3 KiB | [pg_store_plans_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_store_plans_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pg_store_plans_15` | 1.7 | `el9.aarch64` | pgdg | 45.9 KiB | [pg_store_plans_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_store_plans_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pg-store-plan` | 1.8 | `d12.x86_64` | pigsty | 120.6 KiB | [postgresql-15-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-store-plan` | 1.8 | `d12.aarch64` | pigsty | 117.5 KiB | [postgresql-15-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-store-plan` | 1.8 | `u22.aarch64` | pigsty | 128.8 KiB | [postgresql-15-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-store-plan` | 1.8 | `u22.x86_64` | pigsty | 130.6 KiB | [postgresql-15-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-store-plan` | 1.8 | `u24.x86_64` | pigsty | 115.8 KiB | [postgresql-15-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-store-plan` | 1.8 | `u24.aarch64` | pigsty | 114.4 KiB | [postgresql-15-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-15-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_store_plans_14` | 1.8 | `el8.x86_64` | pigsty | 41.6 KiB | [pg_store_plans_14-1.8-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_14-1.8-2PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_14` | 1.8 | `el8.aarch64` | pigsty | 40.2 KiB | [pg_store_plans_14-1.8-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_14-1.8-2PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_14` | 1.8 | `el9.x86_64` | pigsty | 43.6 KiB | [pg_store_plans_14-1.8-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_14-1.8-2PIGSTY.el9.x86_64.rpm) |
| `pg_store_plans_14` | 1.8 | `el9.aarch64` | pigsty | 43.0 KiB | [pg_store_plans_14-1.8-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_14-1.8-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-store-plan` | 1.8 | `d12.x86_64` | pigsty | 120.2 KiB | [postgresql-14-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-store-plan` | 1.8 | `d12.aarch64` | pigsty | 117.1 KiB | [postgresql-14-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-store-plan` | 1.8 | `u22.x86_64` | pigsty | 130.0 KiB | [postgresql-14-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-store-plan` | 1.8 | `u22.aarch64` | pigsty | 128.2 KiB | [postgresql-14-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-store-plan` | 1.8 | `u24.aarch64` | pigsty | 114.4 KiB | [postgresql-14-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-store-plan` | 1.8 | `u24.x86_64` | pigsty | 115.5 KiB | [postgresql-14-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-14-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_store_plans_13` | 1.8 | `el8.x86_64` | pigsty | 41.6 KiB | [pg_store_plans_13-1.8-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_store_plans_13-1.8-2PIGSTY.el8.x86_64.rpm) |
| `pg_store_plans_13` | 1.8 | `el8.aarch64` | pigsty | 40.2 KiB | [pg_store_plans_13-1.8-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_store_plans_13-1.8-2PIGSTY.el8.aarch64.rpm) |
| `pg_store_plans_13` | 1.8 | `el9.x86_64` | pigsty | 43.5 KiB | [pg_store_plans_13-1.8-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_store_plans_13-1.8-2PIGSTY.el9.x86_64.rpm) |
| `pg_store_plans_13` | 1.8 | `el9.aarch64` | pigsty | 42.8 KiB | [pg_store_plans_13-1.8-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_store_plans_13-1.8-2PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-store-plan` | 1.8 | `d12.aarch64` | pigsty | 116.5 KiB | [postgresql-13-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-13-pg-store-plan_1.8-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-store-plan` | 1.8 | `d12.x86_64` | pigsty | 119.6 KiB | [postgresql-13-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-store-plan/postgresql-13-pg-store-plan_1.8-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-store-plan` | 1.8 | `u22.x86_64` | pigsty | 129.4 KiB | [postgresql-13-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-13-pg-store-plan_1.8-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-store-plan` | 1.8 | `u22.aarch64` | pigsty | 127.8 KiB | [postgresql-13-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-store-plan/postgresql-13-pg-store-plan_1.8-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-store-plan` | 1.8 | `u24.x86_64` | pigsty | 115.3 KiB | [postgresql-13-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-13-pg-store-plan_1.8-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-store-plan` | 1.8 | `u24.aarch64` | pigsty | 113.4 KiB | [postgresql-13-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-store-plan/postgresql-13-pg-store-plan_1.8-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ossc-db/pg_store_plans" title="Repository" icon="github" subtitle="github.com/ossc-db/pg_store_plans" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_store_plans-1.8.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_store_plans; # get pg_store_plans source code
pig build dep pg_store_plans; # install build dependencies
pig build pkg pg_store_plans; # build extension rpm or deb
pig build ext pg_store_plans; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_store_plans; # install by extension name, for the current active PG version
pig ext install pg_store_plans; # install via package alias, for the active PG version
pig ext install pg_store_plans -v 17;   # install for PG 17
pig ext install pg_store_plans -v 16;   # install for PG 16
pig ext install pg_store_plans -v 15;   # install for PG 15
pig ext install pg_store_plans -v 14;   # install for PG 14
pig ext install pg_store_plans -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_store_plans;
```

