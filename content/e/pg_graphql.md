---
title: "pg_graphql"
linkTitle: "pg_graphql"
description: "Add in-database GraphQL support"
weight: 2790
categories: ["Feat"]
width: full
---

Add in-database GraphQL support

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2790** | {{< badge content="pg_graphql" link="https://github.com/supabase/pg_graphql" >}} | {{< ext "pg_graphql" "pg_graphql" >}} | `1.5.11` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_net" >}} {{< ext "http" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "wrappers" >}} |

> [!Note] pgrx=0.12.9


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_graphql" >}} | `1.5.11` | {{< badge content="18" color="red" alt="pg_graphql_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_graphql_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_graphql" >}} | `1.5.11` | {{< badge content="18" color="red" alt="postgresql-18-pg-graphql" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-graphql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_graphql_18" >}}     | {{< pkg "pg_graphql_17" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_17-1.5.11-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_graphql_16" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_16-1.5.11-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_graphql_15" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_15-1.5.11-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_graphql_14" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_14-1.5.11-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_graphql_18" >}}     | {{< pkg "pg_graphql_17" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_17-1.5.11-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_graphql_16" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_16-1.5.11-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_graphql_15" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_15-1.5.11-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_graphql_14" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_14-1.5.11-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_graphql_18" >}}     | {{< pkg "pg_graphql_17" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_17-1.5.11-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_graphql_16" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_16-1.5.11-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_graphql_15" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_15-1.5.11-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_graphql_14" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_14-1.5.11-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_graphql_18" >}}     | {{< pkg "pg_graphql_17" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_17-1.5.11-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_graphql_16" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_16-1.5.11-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_graphql_15" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_15-1.5.11-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_graphql_14" "1.5.11" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_14-1.5.11-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-graphql" >}}     | {{< pkg "postgresql-17-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-graphql" >}}     | {{< pkg "postgresql-17-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-graphql" >}}     | {{< pkg "postgresql-17-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-graphql" >}}     | {{< pkg "postgresql-17-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-graphql" >}}     | {{< pkg "postgresql-17-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-graphql" >}}     | {{< pkg "postgresql-17-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-graphql" "1.5.11" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_graphql_17` | 1.5.11 | `el8.aarch64` | pigsty | 694.2 KiB | [pg_graphql_17-1.5.11-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_17-1.5.11-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_17` | 1.5.11 | `el8.x86_64` | pigsty | 782.2 KiB | [pg_graphql_17-1.5.11-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_17-1.5.11-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_17` | 1.5.11 | `el9.aarch64` | pigsty | 741.0 KiB | [pg_graphql_17-1.5.11-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_17-1.5.11-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_17` | 1.5.11 | `el9.x86_64` | pigsty | 780.6 KiB | [pg_graphql_17-1.5.11-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_17-1.5.11-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-graphql` | 1.5.11 | `d12.x86_64` | pigsty | 664.1 KiB | [postgresql-17-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-graphql` | 1.5.11 | `d12.aarch64` | pigsty | 572.7 KiB | [postgresql-17-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-graphql` | 1.5.11 | `u22.aarch64` | pigsty | 671.1 KiB | [postgresql-17-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-graphql` | 1.5.11 | `u22.x86_64` | pigsty | 726.2 KiB | [postgresql-17-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-graphql` | 1.5.11 | `u24.x86_64` | pigsty | 719.6 KiB | [postgresql-17-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-graphql` | 1.5.11 | `u24.aarch64` | pigsty | 664.3 KiB | [postgresql-17-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-17-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_graphql_16` | 1.5.11 | `el8.x86_64` | pigsty | 782.1 KiB | [pg_graphql_16-1.5.11-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_16-1.5.11-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_16` | 1.5.11 | `el8.aarch64` | pigsty | 694.1 KiB | [pg_graphql_16-1.5.11-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_16-1.5.11-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_16` | 1.5.11 | `el9.aarch64` | pigsty | 740.8 KiB | [pg_graphql_16-1.5.11-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_16-1.5.11-1PIGSTY.el9.aarch64.rpm) |
| `pg_graphql_16` | 1.5.11 | `el9.x86_64` | pigsty | 781.3 KiB | [pg_graphql_16-1.5.11-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_16-1.5.11-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-pg-graphql` | 1.5.11 | `d12.aarch64` | pigsty | 573.2 KiB | [postgresql-16-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-graphql` | 1.5.11 | `d12.x86_64` | pigsty | 663.9 KiB | [postgresql-16-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-graphql` | 1.5.11 | `u22.x86_64` | pigsty | 726.0 KiB | [postgresql-16-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-graphql` | 1.5.11 | `u22.aarch64` | pigsty | 671.3 KiB | [postgresql-16-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-graphql` | 1.5.11 | `u24.aarch64` | pigsty | 664.1 KiB | [postgresql-16-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-graphql` | 1.5.11 | `u24.x86_64` | pigsty | 719.3 KiB | [postgresql-16-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-16-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_graphql_15` | 1.5.11 | `el8.x86_64` | pigsty | 782.2 KiB | [pg_graphql_15-1.5.11-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_15-1.5.11-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_15` | 1.5.11 | `el8.aarch64` | pigsty | 694.4 KiB | [pg_graphql_15-1.5.11-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_15-1.5.11-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_15` | 1.5.11 | `el9.x86_64` | pigsty | 780.3 KiB | [pg_graphql_15-1.5.11-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_15-1.5.11-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_15` | 1.5.11 | `el9.aarch64` | pigsty | 741.0 KiB | [pg_graphql_15-1.5.11-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_15-1.5.11-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-graphql` | 1.5.11 | `d12.x86_64` | pigsty | 664.0 KiB | [postgresql-15-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-graphql` | 1.5.11 | `d12.aarch64` | pigsty | 573.0 KiB | [postgresql-15-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-graphql` | 1.5.11 | `u22.aarch64` | pigsty | 671.0 KiB | [postgresql-15-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-graphql` | 1.5.11 | `u22.x86_64` | pigsty | 726.2 KiB | [postgresql-15-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-graphql` | 1.5.11 | `u24.aarch64` | pigsty | 664.4 KiB | [postgresql-15-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-graphql` | 1.5.11 | `u24.x86_64` | pigsty | 717.6 KiB | [postgresql-15-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-15-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_graphql_14` | 1.5.11 | `el8.aarch64` | pigsty | 694.2 KiB | [pg_graphql_14-1.5.11-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_graphql_14-1.5.11-1PIGSTY.el8.aarch64.rpm) |
| `pg_graphql_14` | 1.5.11 | `el8.x86_64` | pigsty | 782.3 KiB | [pg_graphql_14-1.5.11-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_graphql_14-1.5.11-1PIGSTY.el8.x86_64.rpm) |
| `pg_graphql_14` | 1.5.11 | `el9.x86_64` | pigsty | 780.7 KiB | [pg_graphql_14-1.5.11-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_graphql_14-1.5.11-1PIGSTY.el9.x86_64.rpm) |
| `pg_graphql_14` | 1.5.11 | `el9.aarch64` | pigsty | 740.7 KiB | [pg_graphql_14-1.5.11-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_graphql_14-1.5.11-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-graphql` | 1.5.11 | `d12.aarch64` | pigsty | 573.0 KiB | [postgresql-14-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-graphql` | 1.5.11 | `d12.x86_64` | pigsty | 663.8 KiB | [postgresql-14-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-graphql` | 1.5.11 | `u22.x86_64` | pigsty | 725.9 KiB | [postgresql-14-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-graphql` | 1.5.11 | `u22.aarch64` | pigsty | 670.5 KiB | [postgresql-14-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-graphql` | 1.5.11 | `u24.aarch64` | pigsty | 664.1 KiB | [postgresql-14-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-graphql` | 1.5.11 | `u24.x86_64` | pigsty | 718.4 KiB | [postgresql-14-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-graphql/postgresql-14-pg-graphql_1.5.11-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/pg_graphql" title="Repository" icon="github" subtitle="github.com/supabase/pg_graphql" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_graphql-1.5.11.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_graphql; # get pg_graphql source code
pig build dep pg_graphql; # install build dependencies
pig build pkg pg_graphql; # build extension rpm or deb
pig build ext pg_graphql; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_graphql; # install by extension name, for the current active PG version
pig ext install pg_graphql; # install via package alias, for the active PG version
pig ext install pg_graphql -v 17;   # install for PG 17
pig ext install pg_graphql -v 16;   # install for PG 16
pig ext install pg_graphql -v 15;   # install for PG 15
pig ext install pg_graphql -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_graphql CASCADE SCHEMA graphql;
```

