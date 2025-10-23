---
title: "pg_tiktoken"
linkTitle: "pg_tiktoken"
description: "tiktoken tokenizer for use with OpenAI models in postgres"
weight: 1870
categories: ["RAG"]
width: full
---

tiktoken tokenizer for use with OpenAI models in postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1870** | {{< badge content="pg_tiktoken" link="https://github.com/kelvich/pg_tiktoken" >}} | {{< ext "pg_tiktoken" >}} | `0.0.1` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vectorize" >}} {{< ext "pg_summarize" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "pg_graphql" >}} |

> [!Note] pgrx=0.12.6


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_tiktoken" >}} | `0.0.1` | {{< bg "18" "pg_tiktoken_18" "red" >}} {{< bg "17" "pg_tiktoken_17" "green" >}} {{< bg "16" "pg_tiktoken_16" "green" >}} {{< bg "15" "pg_tiktoken_15" "green" >}} {{< bg "14" "pg_tiktoken_14" "green" >}} | `pg_tiktoken_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_tiktoken" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-tiktoken" "red" >}} {{< bg "17" "postgresql-17-pg-tiktoken" "green" >}} {{< bg "16" "postgresql-16-pg-tiktoken" "green" >}} {{< bg "15" "postgresql-15-pg-tiktoken" "green" >}} {{< bg "14" "postgresql-14-pg-tiktoken" "green" >}} | `postgresql-$v-pg-tiktoken` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_tiktoken_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_tiktoken_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_tiktoken_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_tiktoken_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-tiktoken : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-tiktoken : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-tiktoken : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-tiktoken : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-tiktoken : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-tiktoken : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_17` | 0.0.1 | `el8.x86_64` | pigsty | 1.6 MiB | [pg_tiktoken_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_17` | 0.0.1 | `el8.aarch64` | pigsty | 1.6 MiB | [pg_tiktoken_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_17` | 0.0.1 | `el9.x86_64` | pigsty | 1.6 MiB | [pg_tiktoken_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_17` | 0.0.1 | `el9.aarch64` | pigsty | 1.5 MiB | [pg_tiktoken_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-tiktoken` | 0.0.1 | `d12.x86_64` | pigsty | 4.9 KiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tiktoken` | 0.0.1 | `d12.aarch64` | pigsty | 4.9 KiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tiktoken` | 0.0.1 | `u22.x86_64` | pigsty | 4.9 KiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tiktoken` | 0.0.1 | `u22.aarch64` | pigsty | 4.9 KiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tiktoken` | 0.0.1 | `u24.x86_64` | pigsty | 4.9 KiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tiktoken` | 0.0.1 | `u24.aarch64` | pigsty | 4.9 KiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_16` | 0.0.1 | `el8.x86_64` | pigsty | 1.6 MiB | [pg_tiktoken_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_16` | 0.0.1 | `el8.aarch64` | pigsty | 1.6 MiB | [pg_tiktoken_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_16` | 0.0.1 | `el9.x86_64` | pigsty | 1.6 MiB | [pg_tiktoken_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_16` | 0.0.1 | `el9.aarch64` | pigsty | 1.5 MiB | [pg_tiktoken_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-tiktoken` | 0.0.1 | `d12.x86_64` | pigsty | 1.3 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tiktoken` | 0.0.1 | `d12.aarch64` | pigsty | 1.2 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tiktoken` | 0.0.1 | `u22.x86_64` | pigsty | 1.5 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tiktoken` | 0.0.1 | `u22.aarch64` | pigsty | 1.5 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tiktoken` | 0.0.1 | `u24.x86_64` | pigsty | 1.5 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tiktoken` | 0.0.1 | `u24.aarch64` | pigsty | 1.4 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_15` | 0.0.1 | `el8.x86_64` | pigsty | 1.6 MiB | [pg_tiktoken_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_15` | 0.0.1 | `el8.aarch64` | pigsty | 1.6 MiB | [pg_tiktoken_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_15` | 0.0.1 | `el9.x86_64` | pigsty | 1.6 MiB | [pg_tiktoken_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_15` | 0.0.1 | `el9.aarch64` | pigsty | 1.5 MiB | [pg_tiktoken_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-tiktoken` | 0.0.1 | `d12.x86_64` | pigsty | 1.3 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tiktoken` | 0.0.1 | `d12.aarch64` | pigsty | 1.2 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tiktoken` | 0.0.1 | `u22.x86_64` | pigsty | 1.5 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tiktoken` | 0.0.1 | `u22.aarch64` | pigsty | 1.5 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tiktoken` | 0.0.1 | `u24.x86_64` | pigsty | 1.5 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tiktoken` | 0.0.1 | `u24.aarch64` | pigsty | 1.4 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_14` | 0.0.1 | `el8.x86_64` | pigsty | 1.6 MiB | [pg_tiktoken_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_14` | 0.0.1 | `el8.aarch64` | pigsty | 1.6 MiB | [pg_tiktoken_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_14` | 0.0.1 | `el9.x86_64` | pigsty | 1.6 MiB | [pg_tiktoken_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_14` | 0.0.1 | `el9.aarch64` | pigsty | 1.5 MiB | [pg_tiktoken_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-tiktoken` | 0.0.1 | `d12.x86_64` | pigsty | 1.3 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tiktoken` | 0.0.1 | `d12.aarch64` | pigsty | 1.2 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tiktoken` | 0.0.1 | `u22.x86_64` | pigsty | 1.5 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tiktoken` | 0.0.1 | `u22.aarch64` | pigsty | 1.5 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tiktoken` | 0.0.1 | `u24.x86_64` | pigsty | 1.5 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tiktoken` | 0.0.1 | `u24.aarch64` | pigsty | 1.4 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kelvich/pg_tiktoken" title="Repository" icon="github" subtitle="github.com/kelvich/pg_tiktoken" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tiktoken-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_tiktoken; # get pg_tiktoken source code
pig build dep pg_tiktoken; # install build dependencies
pig build pkg pg_tiktoken; # build extension rpm or deb
pig build ext pg_tiktoken; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_tiktoken; # install by extension name, for the current active PG version
pig ext install pg_tiktoken; # install via package alias, for the active PG version
pig ext install pg_tiktoken -v 17;   # install for PG 17
pig ext install pg_tiktoken -v 16;   # install for PG 16
pig ext install pg_tiktoken -v 15;   # install for PG 15
pig ext install pg_tiktoken -v 14;   # install for PG 14
pig ext install pg_tiktoken -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_tiktoken;
```

