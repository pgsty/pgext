---
title: "financial"
linkTitle: "financial"
description: "Financial aggregate functions"
weight: 4840
categories: ["Func"]
width: full
---

Financial aggregate functions

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4840** | {{< badge content="financial" link="https://github.com/intgr/pg_financial" >}} | {{< ext "financial" "pg_financial" >}} | `1.0.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "count_distinct" >}} {{< ext "omnisketch" >}} {{< ext "ddsketch" >}} {{< ext "tdigest" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/financial" >}} | `1.0.1` | {{< badge content="18" color="red" alt="pg_financial_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_financial_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/financial" >}} | `1.0.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-financial" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-financial` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_financial_18" >}}     | {{< pkg "pg_financial_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_17-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_financial_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_16-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_financial_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_15-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_financial_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_14-1.0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_financial_18" >}}     | {{< pkg "pg_financial_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_17-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_financial_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_16-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_financial_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_15-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_financial_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_14-1.0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_financial_18" >}}     | {{< pkg "pg_financial_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_17-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_financial_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_16-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_financial_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_15-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_financial_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_14-1.0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_financial_18" >}}     | {{< pkg "pg_financial_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_17-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_financial_16" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_16-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_financial_15" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_15-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_financial_14" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_14-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-financial" >}}     | {{< pkg "postgresql-17-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-financial" >}}     | {{< pkg "postgresql-17-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-financial" >}}     | {{< pkg "postgresql-17-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-financial" >}}     | {{< pkg "postgresql-17-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-financial" >}}     | {{< pkg "postgresql-17-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-financial" >}}     | {{< pkg "postgresql-17-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-financial" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_financial_17` | 1.0.1 | `el8.x86_64` | pigsty | 14.5 KiB | [pg_financial_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_17` | 1.0.1 | `el8.aarch64` | pigsty | 14.6 KiB | [pg_financial_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_17` | 1.0.1 | `el9.aarch64` | pigsty | 14.6 KiB | [pg_financial_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_financial_17` | 1.0.1 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_financial_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-financial` | 1.0.1 | `d12.x86_64` | pigsty | 13.1 KiB | [postgresql-17-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-financial` | 1.0.1 | `d12.aarch64` | pigsty | 13.2 KiB | [postgresql-17-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-financial` | 1.0.1 | `u22.x86_64` | pigsty | 13.8 KiB | [postgresql-17-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-financial` | 1.0.1 | `u22.aarch64` | pigsty | 13.9 KiB | [postgresql-17-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-financial` | 1.0.1 | `u24.x86_64` | pigsty | 13.5 KiB | [postgresql-17-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-financial` | 1.0.1 | `u24.aarch64` | pigsty | 13.4 KiB | [postgresql-17-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_financial_16` | 1.0.1 | `el8.x86_64` | pigsty | 14.5 KiB | [pg_financial_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_16` | 1.0.1 | `el8.aarch64` | pigsty | 14.6 KiB | [pg_financial_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_16` | 1.0.1 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_financial_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_16` | 1.0.1 | `el9.aarch64` | pigsty | 14.6 KiB | [pg_financial_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-financial` | 1.0.1 | `d12.x86_64` | pigsty | 13.1 KiB | [postgresql-16-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-financial` | 1.0.1 | `d12.aarch64` | pigsty | 13.2 KiB | [postgresql-16-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-financial` | 1.0.1 | `u22.aarch64` | pigsty | 13.9 KiB | [postgresql-16-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-financial` | 1.0.1 | `u22.x86_64` | pigsty | 13.8 KiB | [postgresql-16-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-financial` | 1.0.1 | `u24.x86_64` | pigsty | 13.5 KiB | [postgresql-16-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-financial` | 1.0.1 | `u24.aarch64` | pigsty | 13.3 KiB | [postgresql-16-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_financial_15` | 1.0.1 | `el8.x86_64` | pigsty | 14.5 KiB | [pg_financial_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_15` | 1.0.1 | `el8.aarch64` | pigsty | 14.5 KiB | [pg_financial_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_15` | 1.0.1 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_financial_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_15` | 1.0.1 | `el9.aarch64` | pigsty | 14.6 KiB | [pg_financial_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-financial` | 1.0.1 | `d12.aarch64` | pigsty | 13.2 KiB | [postgresql-15-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-financial` | 1.0.1 | `d12.x86_64` | pigsty | 13.1 KiB | [postgresql-15-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-financial` | 1.0.1 | `u22.aarch64` | pigsty | 13.9 KiB | [postgresql-15-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-financial` | 1.0.1 | `u22.x86_64` | pigsty | 13.8 KiB | [postgresql-15-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-financial` | 1.0.1 | `u24.x86_64` | pigsty | 13.5 KiB | [postgresql-15-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-financial` | 1.0.1 | `u24.aarch64` | pigsty | 13.4 KiB | [postgresql-15-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_financial_14` | 1.0.1 | `el8.x86_64` | pigsty | 14.5 KiB | [pg_financial_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_14` | 1.0.1 | `el8.aarch64` | pigsty | 14.5 KiB | [pg_financial_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_14` | 1.0.1 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_financial_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_14` | 1.0.1 | `el9.aarch64` | pigsty | 14.6 KiB | [pg_financial_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-financial` | 1.0.1 | `d12.x86_64` | pigsty | 13.1 KiB | [postgresql-14-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-financial` | 1.0.1 | `d12.aarch64` | pigsty | 13.2 KiB | [postgresql-14-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-financial` | 1.0.1 | `u22.x86_64` | pigsty | 13.7 KiB | [postgresql-14-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-financial` | 1.0.1 | `u22.aarch64` | pigsty | 13.9 KiB | [postgresql-14-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-financial` | 1.0.1 | `u24.x86_64` | pigsty | 13.5 KiB | [postgresql-14-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-financial` | 1.0.1 | `u24.aarch64` | pigsty | 13.3 KiB | [postgresql-14-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_financial_13` | 1.0.1 | `el8.aarch64` | pigsty | 14.5 KiB | [pg_financial_13-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_13-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_13` | 1.0.1 | `el8.x86_64` | pigsty | 14.4 KiB | [pg_financial_13-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_13-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_13` | 1.0.1 | `el9.aarch64` | pigsty | 14.5 KiB | [pg_financial_13-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_13-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_financial_13` | 1.0.1 | `el9.x86_64` | pigsty | 14.6 KiB | [pg_financial_13-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_13-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-financial` | 1.0.1 | `d12.aarch64` | pigsty | 13.3 KiB | [postgresql-13-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-financial` | 1.0.1 | `d12.x86_64` | pigsty | 13.0 KiB | [postgresql-13-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-financial` | 1.0.1 | `u22.aarch64` | pigsty | 13.7 KiB | [postgresql-13-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-financial` | 1.0.1 | `u22.x86_64` | pigsty | 13.5 KiB | [postgresql-13-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-financial` | 1.0.1 | `u24.aarch64` | pigsty | 13.4 KiB | [postgresql-13-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-financial` | 1.0.1 | `u24.x86_64` | pigsty | 13.4 KiB | [postgresql-13-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/intgr/pg_financial" title="Repository" icon="github" subtitle="github.com/intgr/pg_financial" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_financial-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get financial; # get financial source code
pig build dep financial; # install build dependencies
pig build pkg financial; # build extension rpm or deb
pig build ext financial; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install financial; # install by extension name, for the current active PG version
pig ext install pg_financial; # install via package alias, for the active PG version
pig ext install financial -v 17;   # install for PG 17
pig ext install financial -v 16;   # install for PG 16
pig ext install financial -v 15;   # install for PG 15
pig ext install financial -v 14;   # install for PG 14
pig ext install financial -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION financial;
```

