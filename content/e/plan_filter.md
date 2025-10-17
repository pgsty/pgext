---
title: "plan_filter"
linkTitle: "plan_filter"
description: "filter statements by their execution plans."
weight: 2850
categories: ["Feat"]
width: full
---

filter statements by their execution plans.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2850** | {{< badge content="plan_filter" link="https://github.com/pgexperts/pg_plan_filter" >}} | {{< ext "plan_filter" "pg_plan_filter" >}} | `0.0.1` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/plan_filter" >}} | `0.0.1` | {{< badge content="18" color="red" alt="pg_plan_filter_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_plan_filter_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/plan_filter" >}} | `0.0.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-plan-filter" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-plan-filter` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_plan_filter_18" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_18-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_plan_filter_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_17-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_plan_filter_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_16-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_plan_filter_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_15-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_plan_filter_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_14-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_plan_filter_18" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_18-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_plan_filter_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_17-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_plan_filter_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_16-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_plan_filter_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_15-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_plan_filter_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_14-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_plan_filter_18" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_18-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_plan_filter_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_17-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_plan_filter_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_16-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_plan_filter_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_15-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_plan_filter_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_14-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_plan_filter_18" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_18-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_plan_filter_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_17-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_plan_filter_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_16-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_plan_filter_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_15-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_plan_filter_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_14-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-plan-filter" >}}     | {{< pkg "postgresql-17-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-plan-filter" >}}     | {{< pkg "postgresql-17-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-plan-filter" >}}     | {{< pkg "postgresql-17-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-plan-filter" >}}     | {{< pkg "postgresql-17-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-plan-filter" >}}     | {{< pkg "postgresql-17-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-plan-filter" >}}     | {{< pkg "postgresql-17-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-plan-filter" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_plan_filter_18` | 0.0.1 | `el8.aarch64` | pigsty | 11.0 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_18` | 0.0.1 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_18` | 0.0.1 | `el9.x86_64` | pigsty | 10.7 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_plan_filter_18` | 0.0.1 | `el9.aarch64` | pigsty | 10.5 KiB | [pg_plan_filter_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_plan_filter_17` | 0.0.1 | `el8.aarch64` | pigsty | 11.0 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_17` | 0.0.1 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_17` | 0.0.1 | `el9.aarch64` | pigsty | 10.7 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_plan_filter_17` | 0.0.1 | `el9.x86_64` | pigsty | 10.8 KiB | [pg_plan_filter_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-plan-filter` | 0.0.1 | `d12.aarch64` | pigsty | 10.5 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-plan-filter` | 0.0.1 | `d12.x86_64` | pigsty | 10.4 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-plan-filter` | 0.0.1 | `u22.aarch64` | pigsty | 10.6 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-plan-filter` | 0.0.1 | `u22.x86_64` | pigsty | 10.6 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-plan-filter` | 0.0.1 | `u24.x86_64` | pigsty | 10.2 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-plan-filter` | 0.0.1 | `u24.aarch64` | pigsty | 10.3 KiB | [postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-17-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_plan_filter_16` | 0.0.1 | `el8.aarch64` | pigsty | 11.0 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_16` | 0.0.1 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_16` | 0.0.1 | `el9.x86_64` | pigsty | 10.8 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_plan_filter_16` | 0.0.1 | `el9.aarch64` | pigsty | 10.6 KiB | [pg_plan_filter_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-plan-filter` | 0.0.1 | `d12.x86_64` | pigsty | 10.4 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-plan-filter` | 0.0.1 | `d12.aarch64` | pigsty | 10.5 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-plan-filter` | 0.0.1 | `u22.aarch64` | pigsty | 10.6 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-plan-filter` | 0.0.1 | `u22.x86_64` | pigsty | 10.6 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-plan-filter` | 0.0.1 | `u24.aarch64` | pigsty | 10.3 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-plan-filter` | 0.0.1 | `u24.x86_64` | pigsty | 10.2 KiB | [postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-16-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_plan_filter_15` | 0.0.1 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_15` | 0.0.1 | `el8.aarch64` | pigsty | 11.0 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_15` | 0.0.1 | `el9.x86_64` | pigsty | 10.8 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_plan_filter_15` | 0.0.1 | `el9.aarch64` | pigsty | 10.6 KiB | [pg_plan_filter_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-plan-filter` | 0.0.1 | `d12.x86_64` | pigsty | 10.3 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-plan-filter` | 0.0.1 | `d12.aarch64` | pigsty | 10.5 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-plan-filter` | 0.0.1 | `u22.x86_64` | pigsty | 10.6 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-plan-filter` | 0.0.1 | `u22.aarch64` | pigsty | 10.6 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-plan-filter` | 0.0.1 | `u24.aarch64` | pigsty | 10.3 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-plan-filter` | 0.0.1 | `u24.x86_64` | pigsty | 10.2 KiB | [postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-15-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_plan_filter_14` | 0.0.1 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_14` | 0.0.1 | `el8.aarch64` | pigsty | 11.0 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_14` | 0.0.1 | `el9.aarch64` | pigsty | 10.6 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_plan_filter_14` | 0.0.1 | `el9.x86_64` | pigsty | 10.8 KiB | [pg_plan_filter_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-pg-plan-filter` | 0.0.1 | `d12.x86_64` | pigsty | 10.3 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-plan-filter` | 0.0.1 | `d12.aarch64` | pigsty | 10.5 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-plan-filter` | 0.0.1 | `u22.x86_64` | pigsty | 10.6 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-plan-filter` | 0.0.1 | `u22.aarch64` | pigsty | 10.6 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-plan-filter` | 0.0.1 | `u24.aarch64` | pigsty | 10.3 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-plan-filter` | 0.0.1 | `u24.x86_64` | pigsty | 10.2 KiB | [postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-14-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_plan_filter_13` | 0.0.1 | `el8.x86_64` | pigsty | 11.0 KiB | [pg_plan_filter_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_plan_filter_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_plan_filter_13` | 0.0.1 | `el8.aarch64` | pigsty | 11.0 KiB | [pg_plan_filter_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_plan_filter_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_plan_filter_13` | 0.0.1 | `el9.aarch64` | pigsty | 10.6 KiB | [pg_plan_filter_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_plan_filter_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_plan_filter_13` | 0.0.1 | `el9.x86_64` | pigsty | 10.8 KiB | [pg_plan_filter_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_plan_filter_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-plan-filter` | 0.0.1 | `d12.x86_64` | pigsty | 10.3 KiB | [postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-plan-filter` | 0.0.1 | `d12.aarch64` | pigsty | 10.5 KiB | [postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-plan-filter/postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-plan-filter` | 0.0.1 | `u22.aarch64` | pigsty | 10.5 KiB | [postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-plan-filter` | 0.0.1 | `u22.x86_64` | pigsty | 10.4 KiB | [postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-plan-filter/postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-plan-filter` | 0.0.1 | `u24.x86_64` | pigsty | 10.2 KiB | [postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-plan-filter` | 0.0.1 | `u24.aarch64` | pigsty | 10.3 KiB | [postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-plan-filter/postgresql-13-pg-plan-filter_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgexperts/pg_plan_filter" title="Repository" icon="github" subtitle="github.com/pgexperts/pg_plan_filter" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_plan_filter.tar.gz" >}}
{{< /cards >}}


```bash
pig build get plan_filter; # get plan_filter source code
pig build dep plan_filter; # install build dependencies
pig build pkg plan_filter; # build extension rpm or deb
pig build ext plan_filter; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install plan_filter; # install by extension name, for the current active PG version
pig ext install pg_plan_filter; # install via package alias, for the active PG version
pig ext install plan_filter -v 18;   # install for PG 18
pig ext install plan_filter -v 17;   # install for PG 17
pig ext install plan_filter -v 16;   # install for PG 16
pig ext install plan_filter -v 15;   # install for PG 15
pig ext install plan_filter -v 14;   # install for PG 14
pig ext install plan_filter -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION plan_filter;
```

