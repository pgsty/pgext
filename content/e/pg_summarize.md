---
title: "pg_summarize"
linkTitle: "pg_summarize"
description: "Text Summarization using LLMs. Built using pgrx"
weight: 1860
categories: ["RAG"]
width: full
---

Text Summarization using LLMs. Built using pgrx


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1860** | {{< badge content="pg_summarize" link="https://github.com/HexaCluster/pg_summarize" >}} | {{< ext "pg_summarize" >}} | `0.0.1` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vectorize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "pg_net" >}} |

> [!Note] pgrx=0.12.4


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_summarize" >}} | `0.0.1` | {{< bg "18" "pg_summarize_18" "red" >}} {{< bg "17" "pg_summarize_17" "green" >}} {{< bg "16" "pg_summarize_16" "green" >}} {{< bg "15" "pg_summarize_15" "green" >}} {{< bg "14" "pg_summarize_14" "green" >}} | `pg_summarize_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_summarize" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-summarize" "red" >}} {{< bg "17" "postgresql-17-pg-summarize" "green" >}} {{< bg "16" "postgresql-16-pg-summarize" "green" >}} {{< bg "15" "postgresql-15-pg-summarize" "green" >}} {{< bg "14" "postgresql-14-pg-summarize" "green" >}} | `postgresql-$v-pg-summarize` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_summarize_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_summarize_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_summarize_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_summarize_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-summarize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-summarize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-summarize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-summarize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-summarize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-summarize : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_17` | 0.0.1 | `el8.x86_64` | pigsty | 964.2 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_17` | 0.0.1 | `el8.aarch64` | pigsty | 899.7 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_17` | 0.0.1 | `el9.x86_64` | pigsty | 963.8 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_17` | 0.0.1 | `el9.aarch64` | pigsty | 951.5 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-summarize` | 0.0.1 | `d12.x86_64` | pigsty | 773.4 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-summarize` | 0.0.1 | `d12.aarch64` | pigsty | 690.3 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-summarize` | 0.0.1 | `u22.x86_64` | pigsty | 852.1 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-summarize` | 0.0.1 | `u22.aarch64` | pigsty | 815.1 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-summarize` | 0.0.1 | `u24.x86_64` | pigsty | 845.5 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-summarize` | 0.0.1 | `u24.aarch64` | pigsty | 806.8 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_16` | 0.0.1 | `el8.x86_64` | pigsty | 964.1 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_16` | 0.0.1 | `el8.aarch64` | pigsty | 899.6 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_16` | 0.0.1 | `el9.x86_64` | pigsty | 963.6 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_16` | 0.0.1 | `el9.aarch64` | pigsty | 951.3 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-summarize` | 0.0.1 | `d12.x86_64` | pigsty | 773.5 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-summarize` | 0.0.1 | `d12.aarch64` | pigsty | 690.2 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-summarize` | 0.0.1 | `u22.x86_64` | pigsty | 853.0 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-summarize` | 0.0.1 | `u22.aarch64` | pigsty | 814.8 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-summarize` | 0.0.1 | `u24.x86_64` | pigsty | 845.5 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-summarize` | 0.0.1 | `u24.aarch64` | pigsty | 806.6 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_15` | 0.0.1 | `el8.x86_64` | pigsty | 964.1 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_15` | 0.0.1 | `el8.aarch64` | pigsty | 899.5 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_15` | 0.0.1 | `el9.x86_64` | pigsty | 963.8 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_15` | 0.0.1 | `el9.aarch64` | pigsty | 951.2 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-summarize` | 0.0.1 | `d12.x86_64` | pigsty | 774.0 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-summarize` | 0.0.1 | `d12.aarch64` | pigsty | 690.3 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-summarize` | 0.0.1 | `u22.x86_64` | pigsty | 851.8 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-summarize` | 0.0.1 | `u22.aarch64` | pigsty | 815.1 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-summarize` | 0.0.1 | `u24.x86_64` | pigsty | 846.0 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-summarize` | 0.0.1 | `u24.aarch64` | pigsty | 806.7 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_14` | 0.0.1 | `el8.x86_64` | pigsty | 964.1 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_14` | 0.0.1 | `el8.aarch64` | pigsty | 899.6 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_14` | 0.0.1 | `el9.x86_64` | pigsty | 963.8 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_14` | 0.0.1 | `el9.aarch64` | pigsty | 951.2 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-summarize` | 0.0.1 | `d12.x86_64` | pigsty | 773.4 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-summarize` | 0.0.1 | `d12.aarch64` | pigsty | 690.9 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-summarize` | 0.0.1 | `u22.x86_64` | pigsty | 853.1 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-summarize` | 0.0.1 | `u22.aarch64` | pigsty | 815.1 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-summarize` | 0.0.1 | `u24.x86_64` | pigsty | 845.5 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-summarize` | 0.0.1 | `u24.aarch64` | pigsty | 806.9 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_summarize" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_summarize" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_summarize-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_summarize; # get pg_summarize source code
pig build dep pg_summarize; # install build dependencies
pig build pkg pg_summarize; # build extension rpm or deb
pig build ext pg_summarize; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_summarize; # install by extension name, for the current active PG version
pig ext install pg_summarize; # install via package alias, for the active PG version
pig ext install pg_summarize -v 17;   # install for PG 17
pig ext install pg_summarize -v 16;   # install for PG 16
pig ext install pg_summarize -v 15;   # install for PG 15
pig ext install pg_summarize -v 14;   # install for PG 14
pig ext install pg_summarize -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_summarize;
```

