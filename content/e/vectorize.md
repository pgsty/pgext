---
title: "vectorize"
linkTitle: "vectorize"
description: "The simplest way to do vector search on Postgres"
weight: 1830
categories: ["RAG"]
width: full
---

The simplest way to do vector search on Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1830** | {{< badge content="vectorize" link="https://github.com/ChuckHend/pg_vectorize" >}} | {{< ext "vectorize" "pg_vectorize" >}} | `0.25.0` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_cron" >}} {{< ext "pgmq" >}} {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "pg_later" >}} {{< ext "pg_similarity" >}} |

> [!Note] pgrx=0.16.1


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/vectorize" >}} | `0.25.0` | {{< bg "18" "pg_vectorize_18" "green" >}} {{< bg "17" "pg_vectorize_17" "green" >}} {{< bg "16" "pg_vectorize_16" "green" >}} {{< bg "15" "pg_vectorize_15" "green" >}} {{< bg "14" "pg_vectorize_14" "green" >}} {{< bg "13" "pg_vectorize_13" "red" >}} | `pg_vectorize_$v` | `pgmq_$v`, `pg_cron_$v`, `pgvector_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/vectorize" >}} | `0.25.0` | {{< bg "18" "postgresql-18-pg-vectorize" "green" >}} {{< bg "17" "postgresql-17-pg-vectorize" "green" >}} {{< bg "16" "postgresql-16-pg-vectorize" "green" >}} {{< bg "15" "postgresql-15-pg-vectorize" "green" >}} {{< bg "14" "postgresql-14-pg-vectorize" "green" >}} {{< bg "13" "postgresql-13-pg-vectorize" "red" >}} | `postgresql-$v-pg-vectorize` | `postgresql-$v-pgmq`, `postgresql-$v-pg-cron`, `postgresql-$v-pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_vectorize_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_vectorize_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_vectorize_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_vectorize_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_vectorize_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_vectorize_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_vectorize_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "pg_vectorize_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_vectorize_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_vectorize_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_vectorize_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_vectorize_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-vectorize : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-vectorize : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-vectorize : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-vectorize : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-vectorize : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-vectorize : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-vectorize : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.22.2" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.22.2" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-vectorize : MISS 0" "red" >}}      |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_17` | 0.22.2 | `el8.x86_64` | pigsty | 7.2 MiB | [pg_vectorize_17-0.22.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_17-0.22.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_17` | 0.22.2 | `el8.aarch64` | pigsty | 7.0 MiB | [pg_vectorize_17-0.22.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_17-0.22.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_17` | 0.22.2 | `el9.x86_64` | pigsty | 7.0 MiB | [pg_vectorize_17-0.22.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_17-0.22.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_17` | 0.22.2 | `el9.aarch64` | pigsty | 6.9 MiB | [pg_vectorize_17-0.22.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_17-0.22.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-vectorize` | 0.22.2 | `d12.x86_64` | pigsty | 6.0 MiB | [postgresql-17-pg-vectorize_0.22.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.22.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-vectorize` | 0.22.2 | `d12.aarch64` | pigsty | 5.7 MiB | [postgresql-17-pg-vectorize_0.22.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.22.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-vectorize` | 0.22.2 | `u22.x86_64` | pigsty | 6.7 MiB | [postgresql-17-pg-vectorize_0.22.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.22.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-vectorize` | 0.22.2 | `u22.aarch64` | pigsty | 6.6 MiB | [postgresql-17-pg-vectorize_0.22.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.22.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-vectorize` | 0.22.2 | `u24.x86_64` | pigsty | 6.7 MiB | [postgresql-17-pg-vectorize_0.22.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.22.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-vectorize` | 0.22.2 | `u24.aarch64` | pigsty | 6.6 MiB | [postgresql-17-pg-vectorize_0.22.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.22.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_16` | 0.22.2 | `el8.x86_64` | pigsty | 7.2 MiB | [pg_vectorize_16-0.22.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_16-0.22.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_16` | 0.22.2 | `el8.aarch64` | pigsty | 7.0 MiB | [pg_vectorize_16-0.22.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_16-0.22.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_16` | 0.22.2 | `el9.x86_64` | pigsty | 7.0 MiB | [pg_vectorize_16-0.22.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_16-0.22.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_16` | 0.22.2 | `el9.aarch64` | pigsty | 6.9 MiB | [pg_vectorize_16-0.22.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_16-0.22.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-vectorize` | 0.22.2 | `d12.x86_64` | pigsty | 6.0 MiB | [postgresql-16-pg-vectorize_0.22.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.22.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-vectorize` | 0.22.2 | `d12.aarch64` | pigsty | 5.7 MiB | [postgresql-16-pg-vectorize_0.22.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.22.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-vectorize` | 0.22.2 | `u22.x86_64` | pigsty | 6.7 MiB | [postgresql-16-pg-vectorize_0.22.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.22.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-vectorize` | 0.22.2 | `u22.aarch64` | pigsty | 6.6 MiB | [postgresql-16-pg-vectorize_0.22.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.22.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-vectorize` | 0.22.2 | `u24.x86_64` | pigsty | 6.7 MiB | [postgresql-16-pg-vectorize_0.22.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.22.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-vectorize` | 0.22.2 | `u24.aarch64` | pigsty | 6.6 MiB | [postgresql-16-pg-vectorize_0.22.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.22.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_15` | 0.22.2 | `el8.x86_64` | pigsty | 7.2 MiB | [pg_vectorize_15-0.22.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_15-0.22.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_15` | 0.22.2 | `el8.aarch64` | pigsty | 7.0 MiB | [pg_vectorize_15-0.22.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_15-0.22.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_15` | 0.22.2 | `el9.x86_64` | pigsty | 7.0 MiB | [pg_vectorize_15-0.22.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_15-0.22.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_15` | 0.22.2 | `el9.aarch64` | pigsty | 6.9 MiB | [pg_vectorize_15-0.22.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_15-0.22.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-vectorize` | 0.22.2 | `d12.x86_64` | pigsty | 6.0 MiB | [postgresql-15-pg-vectorize_0.22.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.22.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-vectorize` | 0.22.2 | `d12.aarch64` | pigsty | 5.7 MiB | [postgresql-15-pg-vectorize_0.22.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.22.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-vectorize` | 0.22.2 | `u22.x86_64` | pigsty | 6.7 MiB | [postgresql-15-pg-vectorize_0.22.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.22.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-vectorize` | 0.22.2 | `u22.aarch64` | pigsty | 6.6 MiB | [postgresql-15-pg-vectorize_0.22.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.22.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-vectorize` | 0.22.2 | `u24.x86_64` | pigsty | 6.7 MiB | [postgresql-15-pg-vectorize_0.22.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.22.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-vectorize` | 0.22.2 | `u24.aarch64` | pigsty | 6.6 MiB | [postgresql-15-pg-vectorize_0.22.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.22.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_14` | 0.22.2 | `el8.x86_64` | pigsty | 7.2 MiB | [pg_vectorize_14-0.22.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_14-0.22.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_14` | 0.22.2 | `el8.aarch64` | pigsty | 7.0 MiB | [pg_vectorize_14-0.22.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_14-0.22.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_14` | 0.22.2 | `el9.x86_64` | pigsty | 7.0 MiB | [pg_vectorize_14-0.22.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_14-0.22.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_14` | 0.22.2 | `el9.aarch64` | pigsty | 6.9 MiB | [pg_vectorize_14-0.22.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_14-0.22.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-vectorize` | 0.22.2 | `d12.x86_64` | pigsty | 6.0 MiB | [postgresql-14-pg-vectorize_0.22.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.22.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-vectorize` | 0.22.2 | `d12.aarch64` | pigsty | 5.7 MiB | [postgresql-14-pg-vectorize_0.22.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.22.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-vectorize` | 0.22.2 | `u22.x86_64` | pigsty | 6.7 MiB | [postgresql-14-pg-vectorize_0.22.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.22.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-vectorize` | 0.22.2 | `u22.aarch64` | pigsty | 6.6 MiB | [postgresql-14-pg-vectorize_0.22.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.22.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-vectorize` | 0.22.2 | `u24.x86_64` | pigsty | 6.7 MiB | [postgresql-14-pg-vectorize_0.22.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.22.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-vectorize` | 0.22.2 | `u24.aarch64` | pigsty | 6.6 MiB | [postgresql-14-pg-vectorize_0.22.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.22.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChuckHend/pg_vectorize" title="Repository" icon="github" subtitle="github.com/ChuckHend/pg_vectorize" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_vectorize-0.25.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get vectorize; # get vectorize source code
pig build dep vectorize; # install build dependencies
pig build pkg vectorize; # build extension rpm or deb
pig build ext vectorize; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install vectorize; # install by extension name, for the current active PG version
pig ext install pg_vectorize; # install via package alias, for the active PG version
pig ext install vectorize -v 18;   # install for PG 18
pig ext install vectorize -v 17;   # install for PG 17
pig ext install vectorize -v 16;   # install for PG 16
pig ext install vectorize -v 15;   # install for PG 15
pig ext install vectorize -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION vectorize CASCADE SCHEMA vectorize;
```

