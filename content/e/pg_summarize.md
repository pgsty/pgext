---
title: "pg_summarize"
linkTitle: "pg_summarize"
description: "Text Summarization using LLMs. Built using pgrx"
weight: 1860
categories: ["RAG"]
width: full
---

[**pg_summarize**](https://github.com/HexaCluster/pg_summarize)


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

> [!Note] PG18 fix by https://github.com/Vonng/pg_summarize


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_summarize" >}} | `0.0.1` | {{< bg "18" "pg_summarize_18" "green" >}} {{< bg "17" "pg_summarize_17" "green" >}} {{< bg "16" "pg_summarize_16" "green" >}} {{< bg "15" "pg_summarize_15" "green" >}} {{< bg "14" "pg_summarize_14" "green" >}} {{< bg "13" "pg_summarize_13" "green" >}} | `pg_summarize_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_summarize" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-summarize" "green" >}} {{< bg "17" "postgresql-17-pg-summarize" "green" >}} {{< bg "16" "postgresql-16-pg-summarize" "green" >}} {{< bg "15" "postgresql-15-pg-summarize" "green" >}} {{< bg "14" "postgresql-14-pg-summarize" "green" >}} {{< bg "13" "postgresql-13-pg-summarize" "green" >}} | `postgresql-$v-pg-summarize` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-summarize : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.5 KiB | [pg_summarize_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.3 KiB | [pg_summarize_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_18-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_18-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1001.1 KiB | [pg_summarize_18-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_18-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 KiB | [postgresql-18-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 KiB | [postgresql-18-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.5 KiB | [postgresql-18-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.5 KiB | [postgresql-18-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.4 KiB | [postgresql-18-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 KiB | [postgresql-18-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 KiB | [postgresql-18-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 KiB | [postgresql-18-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.4 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.2 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_17-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_17-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1002.2 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_17-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 893.6 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 730.8 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 893.7 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 728.6 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 992.2 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 858.0 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 984.2 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 849.3 KiB | [postgresql-17-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.5 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.2 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_16-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_16-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1001.3 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_16-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 894.0 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 728.4 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 893.8 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 728.5 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 992.3 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 858.1 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 984.2 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 849.2 KiB | [postgresql-16-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.6 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1002.8 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_15-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_15-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1001.1 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_15-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 893.6 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 729.9 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 893.4 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 728.4 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 991.7 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 858.8 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 983.8 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 849.3 KiB | [postgresql-15-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.5 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.7 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_14-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_14-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1001.2 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_14-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 893.8 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 728.3 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 893.0 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 728.7 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 990.7 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 858.8 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 983.6 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 849.1 KiB | [postgresql-14-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_13` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_13` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.6 KiB | [pg_summarize_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_13` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_13` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.2 KiB | [pg_summarize_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_13` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_13-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_13-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_13` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1000.9 KiB | [pg_summarize_13-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_13-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 893.6 KiB | [postgresql-13-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-13-pg-summarize_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 728.3 KiB | [postgresql-13-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-13-pg-summarize_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 894.1 KiB | [postgresql-13-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-13-pg-summarize_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 729.2 KiB | [postgresql-13-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-13-pg-summarize_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 990.7 KiB | [postgresql-13-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-13-pg-summarize_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 858.9 KiB | [postgresql-13-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-13-pg-summarize_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 984.5 KiB | [postgresql-13-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-13-pg-summarize_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 849.4 KiB | [postgresql-13-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-13-pg-summarize_0.0.1-1PIGSTY~noble_arm64.deb) |

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
pig ext install pg_summarize -v 18;   # install for PG 18
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

