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
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-vectorize : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-vectorize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-vectorize : MISS 0" "red" >}}      |


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

- Sources: [repo README](https://github.com/ChuckHend/pg_vectorize/blob/main/README.md), [extension README](https://github.com/ChuckHend/pg_vectorize/blob/main/extension/README.md), [v0.26.1 release](https://github.com/ChuckHend/pg_vectorize/releases/tag/v0.26.1)

`vectorize` is the PostgreSQL extension from `pg_vectorize`. Upstream documents two modes: a standalone HTTP service and the in-database SQL extension. For the packaged extension here, the SQL workflow is the relevant one.

### Enable The Extension

```sql
ALTER SYSTEM SET shared_preload_libraries = 'vectorize,pg_cron';
ALTER SYSTEM SET cron.database_name = 'postgres';

CREATE EXTENSION vectorize CASCADE;
```

The extension README lists `pg_cron`, `pgmq`, and `pgvector` as dependencies, plus `vectorize.embedding_service_url` for the embedding service.

### Create A Search Job

The high-level SQL API starts with `vectorize.table()`:

```sql
SELECT vectorize.table(
  job_name    => 'product_search_hf',
  relation    => 'products',
  primary_key => 'product_id',
  columns     => ARRAY['product_name', 'description'],
  transformer => 'sentence-transformers/all-MiniLM-L6-v2',
  schedule    => 'realtime'
);
```

The extension README says this creates and maintains an embeddings column for the source table.

### Search, RAG, And Direct Model Calls

Search with:

```sql
SELECT * FROM vectorize.search(
  job_name       => 'product_search_hf',
  query          => 'accessories for mobile devices',
  return_columns => ARRAY['product_id', 'product_name'],
  num_results    => 3
);
```

Upstream also documents:

- `vectorize.rag()` for retrieval-augmented answers.
- `vectorize.generate()` for text generation.
- `vectorize.encode()` for direct embedding generation.
- `vectorize.import_embeddings()` for loading precomputed vectors.

### Update Behavior And v0.26.1 Note

The extension README says `schedule => '* * * * *'` checks for updates every minute, while `schedule => 'realtime'` creates triggers for immediate refresh on inserts and updates.

The v0.26.1 release note only says "update dependencies", so there is no upstream user-facing SQL/API delta to document beyond the existing README surface.
