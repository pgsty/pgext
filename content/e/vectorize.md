---
title: "vectorize"
linkTitle: "vectorize"
description: "The simplest way to do vector search on Postgres"
weight: 1830
categories: ["RAG"]
width: full
---

[**pg_vectorize**](https://github.com/ChuckHend/pg_vectorize) : The simplest way to do vector search on Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1830** | {{< badge content="vectorize" link="https://github.com/ChuckHend/pg_vectorize" >}} | {{< ext "vectorize" "pg_vectorize" >}} | `0.26.1` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `vectorize` |
|   **Requires**    | {{< ext "pg_cron" >}} {{< ext "pgmq" >}} {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "pg_later" >}} {{< ext "pg_similarity" >}} |

> [!Note] manually upgraded PGRX from 0.16.1 to 0.17.0 by Vonng; shared_preload_libraries should include vectorize and pg_cron.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.26.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_vectorize` | `pg_cron`, `pgmq`, `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.26.1` | {{< bg "18" "pg_vectorize_18" "green" >}} {{< bg "17" "pg_vectorize_17" "green" >}} {{< bg "16" "pg_vectorize_16" "green" >}} {{< bg "15" "pg_vectorize_15" "green" >}} {{< bg "14" "pg_vectorize_14" "green" >}} | `pg_vectorize_$v` | `pgmq_$v`, `pg_cron_$v`, `pgvector_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.26.1` | {{< bg "18" "postgresql-18-pg-vectorize" "green" >}} {{< bg "17" "postgresql-17-pg-vectorize" "green" >}} {{< bg "16" "postgresql-16-pg-vectorize" "green" >}} {{< bg "15" "postgresql-15-pg-vectorize" "green" >}} {{< bg "14" "postgresql-14-pg-vectorize" "green" >}} | `postgresql-$v-pg-vectorize` | `postgresql-$v-pgmq`, `postgresql-$v-pg-cron`, `postgresql-$v-pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "pg_vectorize_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-18-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-18-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-18-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-18-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-18-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-18-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-18-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-18-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-17-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-16-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-15-pg-vectorize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.26.1" "postgresql-14-pg-vectorize : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_18` | `0.26.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_vectorize_18-0.26.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_18-0.26.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_18` | `0.26.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pg_vectorize_18-0.26.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_18-0.26.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_18` | `0.26.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_18-0.26.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_18-0.26.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_18` | `0.26.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.7 MiB | [pg_vectorize_18-0.26.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_18-0.26.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_vectorize_18` | `0.26.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_18-0.26.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vectorize_18-0.26.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_vectorize_18` | `0.26.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.3 MiB | [pg_vectorize_18-0.26.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vectorize_18-0.26.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-vectorize` | `0.26.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-18-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-18-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-vectorize` | `0.26.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 MiB | [postgresql-18-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-18-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-vectorize` | `0.26.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.9 MiB | [postgresql-18-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-18-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-vectorize` | `0.26.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.5 MiB | [postgresql-18-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-18-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-vectorize` | `0.26.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-18-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-18-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-vectorize` | `0.26.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 MiB | [postgresql-18-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-18-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-vectorize` | `0.26.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-18-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-18-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-vectorize` | `0.26.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 MiB | [postgresql-18-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-18-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_17` | `0.26.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_vectorize_17-0.26.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_17-0.26.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_17` | `0.26.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pg_vectorize_17-0.26.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_17-0.26.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_17` | `0.26.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_17-0.26.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_17-0.26.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_17` | `0.26.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.7 MiB | [pg_vectorize_17-0.26.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_17-0.26.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_vectorize_17` | `0.26.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_17-0.26.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vectorize_17-0.26.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_vectorize_17` | `0.26.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.3 MiB | [pg_vectorize_17-0.26.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vectorize_17-0.26.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-vectorize` | `0.26.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.9 MiB | [postgresql-17-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-vectorize` | `0.26.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 MiB | [postgresql-17-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-vectorize` | `0.26.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.9 MiB | [postgresql-17-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-vectorize` | `0.26.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 MiB | [postgresql-17-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-vectorize` | `0.26.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-17-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-vectorize` | `0.26.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 MiB | [postgresql-17-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-vectorize` | `0.26.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-17-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-vectorize` | `0.26.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 MiB | [postgresql-17-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-17-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_16` | `0.26.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_vectorize_16-0.26.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_16-0.26.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_16` | `0.26.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pg_vectorize_16-0.26.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_16-0.26.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_16` | `0.26.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_16-0.26.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_16-0.26.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_16` | `0.26.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.7 MiB | [pg_vectorize_16-0.26.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_16-0.26.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_vectorize_16` | `0.26.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_16-0.26.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vectorize_16-0.26.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_vectorize_16` | `0.26.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.3 MiB | [pg_vectorize_16-0.26.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vectorize_16-0.26.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-vectorize` | `0.26.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.9 MiB | [postgresql-16-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-vectorize` | `0.26.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 MiB | [postgresql-16-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-vectorize` | `0.26.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-16-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-vectorize` | `0.26.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 MiB | [postgresql-16-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-vectorize` | `0.26.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-16-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-vectorize` | `0.26.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 MiB | [postgresql-16-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-vectorize` | `0.26.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-16-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-vectorize` | `0.26.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 MiB | [postgresql-16-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-16-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_15` | `0.26.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_vectorize_15-0.26.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_15-0.26.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_15` | `0.26.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pg_vectorize_15-0.26.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_15-0.26.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_15` | `0.26.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_15-0.26.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_15-0.26.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_15` | `0.26.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.7 MiB | [pg_vectorize_15-0.26.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_15-0.26.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_vectorize_15` | `0.26.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_15-0.26.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vectorize_15-0.26.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_vectorize_15` | `0.26.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.4 MiB | [pg_vectorize_15-0.26.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vectorize_15-0.26.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-vectorize` | `0.26.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.9 MiB | [postgresql-15-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-vectorize` | `0.26.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 MiB | [postgresql-15-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-vectorize` | `0.26.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.9 MiB | [postgresql-15-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-vectorize` | `0.26.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.5 MiB | [postgresql-15-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-vectorize` | `0.26.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-15-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-vectorize` | `0.26.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 MiB | [postgresql-15-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-vectorize` | `0.26.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-15-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-vectorize` | `0.26.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 MiB | [postgresql-15-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-15-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorize_14` | `0.26.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.2 MiB | [pg_vectorize_14-0.26.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorize_14-0.26.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_vectorize_14` | `0.26.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pg_vectorize_14-0.26.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorize_14-0.26.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_vectorize_14` | `0.26.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_14-0.26.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorize_14-0.26.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_vectorize_14` | `0.26.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.7 MiB | [pg_vectorize_14-0.26.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorize_14-0.26.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_vectorize_14` | `0.26.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pg_vectorize_14-0.26.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vectorize_14-0.26.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_vectorize_14` | `0.26.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.3 MiB | [pg_vectorize_14-0.26.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vectorize_14-0.26.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-vectorize` | `0.26.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.9 MiB | [postgresql-14-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.26.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-vectorize` | `0.26.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 MiB | [postgresql-14-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.26.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-vectorize` | `0.26.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.9 MiB | [postgresql-14-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.26.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-vectorize` | `0.26.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.5 MiB | [postgresql-14-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.26.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-vectorize` | `0.26.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.7 MiB | [postgresql-14-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.26.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-vectorize` | `0.26.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 MiB | [postgresql-14-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.26.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-vectorize` | `0.26.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.7 MiB | [postgresql-14-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.26.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-vectorize` | `0.26.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 MiB | [postgresql-14-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorize/postgresql-14-pg-vectorize_0.26.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChuckHend/pg_vectorize" title="Repository" icon="github" subtitle="github.com/ChuckHend/pg_vectorize" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_vectorize-0.26.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_vectorize;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_vectorize;		# install via package name, for the active PG version
pig install vectorize;		# install by extension name, for the current active PG version

pig install vectorize -v 18;   # install for PG 18
pig install vectorize -v 17;   # install for PG 17
pig install vectorize -v 16;   # install for PG 16
pig install vectorize -v 15;   # install for PG 15
pig install vectorize -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_cron, vectorize';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION vectorize CASCADE; -- requires pg_cron, pgmq, vector
```



## Usage

> [pg_vectorize](https://github.com/ChuckHend/pg_vectorize): The simplest way to do vector search on Postgres.
> Source: [README.md](https://raw.githubusercontent.com/ChuckHend/pg_vectorize/main/README.md)

A Postgres extension that automates the transformation and orchestration of text to embeddings and provides hooks into the most popular LLMs. This allows you to get up and running and automate maintenance for vector search, full text search, and hybrid search, which enables you to quickly build RAG and search engines on Postgres.

This project relies heavily on [pgvector](https://github.com/pgvector/pgvector) for vector similarity search, [pgmq](https://github.com/pgmq/pgmq) for orchestration in background workers, and [SentenceTransformers](https://huggingface.co/sentence-transformers).

**API Documentation**: https://chuckhend.github.io/pg_vectorize/

--------

### Overview

pg_vectorize provides two ways to add semantic, full text, and hybrid search to any Postgres, making it easy to build retrieval-augmented generation (RAG) on Postgres.

Modes at a glance:

- **HTTP server** (recommended for managed DBs): run a standalone service that connects to Postgres and exposes a REST API (`POST /api/v1/table`, `GET /api/v1/search`).
- **Postgres extension** (SQL): install the extension into Postgres and use SQL functions like `vectorize.table()` and `vectorize.search()` (requires filesystem access to Postgres).

--------

### Quick Start -- HTTP Server

Run Postgres and the HTTP servers locally using docker compose:

```bash
# runs Postgres, the embeddings server, and the management API
docker compose up -d
```

Load the example dataset into Postgres (optional):

```bash
psql postgres://postgres:postgres@localhost:5432/postgres -f server/sql/example.sql
```

Create an embedding job via the HTTP API. This generates embeddings for the existing data and continuously watches for updates or new data:

```bash
curl -X POST http://localhost:8080/api/v1/table -d '{
		"job_name": "my_job",
		"src_table": "my_products",
		"src_schema": "public",
		"src_columns": ["product_name", "description"],
		"primary_key": "product_id",
		"update_time_col": "updated_at",
		"model": "sentence-transformers/all-MiniLM-L6-v2"
	}' -H "Content-Type: application/json"
```

```json
{"id":"16b80184-2e8e-4ee6-b7e2-1a068ff4b314"}
```

Search using the HTTP API:

```bash
curl -G \
  "http://localhost:8080/api/v1/search" \
  --data-urlencode "job_name=my_job" \
  --data-urlencode "query=camping backpack" \
  --data-urlencode "limit=1" \
  | jq .
```

```json
[
  {
    "description": "Storage solution for carrying personal items on ones back",
    "fts_rank": 1,
    "price": 45.0,
    "product_category": "accessories",
    "product_id": 6,
    "product_name": "Backpack",
    "rrf_score": 0.03278688524590164,
    "semantic_rank": 1,
    "similarity_score": 0.6296013593673706,
    "updated_at": "2025-10-05T00:14:39.220893+00:00"
  }
]
```

--------

### Which Mode Should I Pick?

- Use the **HTTP server** when your Postgres is managed (RDS, Cloud SQL, etc.) or you cannot install extensions. It requires only that `pgvector` is available in the database. You run the HTTP services separately.
- Use the **Postgres extension** when you self-host Postgres and can install extensions. This provides an in-database experience and direct SQL APIs for vectorization and RAG.

### Quick Start -- Postgres Extension (SQL)

```sql
CREATE EXTENSION vectorize CASCADE;
```

Use `vectorize.table()` to create an embedding job and `vectorize.search()` to perform semantic search directly from SQL.
