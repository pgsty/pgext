---
title: "vectorscale"
linkTitle: "vectorscale"
description: "Advanced indexing for vector data with DiskANN"
weight: 1820
categories: ["RAG"]
width: full
---

Advanced indexing for vector data with DiskANN


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1820** | {{< badge content="vectorscale" link="https://github.com/timescale/pgvectorscale" >}} | {{< ext "vectorscale" "pgvectorscale" >}} | `0.8.0` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vchord" >}} {{< ext "vectorize" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_similarity" >}} |

> [!Note] pgrx=0.12.9


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/vectorscale" >}} | `0.8.0` | {{< bg "18" "pgvectorscale_18" "red" >}} {{< bg "17" "pgvectorscale_17" "green" >}} {{< bg "16" "pgvectorscale_16" "green" >}} {{< bg "15" "pgvectorscale_15" "green" >}} {{< bg "14" "pgvectorscale_14" "green" >}} {{< bg "13" "pgvectorscale_13" "green" >}} | `pgvectorscale_$v` | `pgvector_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/vectorscale" >}} | `0.8.0` | {{< bg "18" "postgresql-18-pgvectorscale" "red" >}} {{< bg "17" "postgresql-17-pgvectorscale" "green" >}} {{< bg "16" "postgresql-16-pgvectorscale" "green" >}} {{< bg "15" "postgresql-15-pgvectorscale" "green" >}} {{< bg "14" "postgresql-14-pgvectorscale" "green" >}} {{< bg "13" "postgresql-13-pgvectorscale" "green" >}} | `postgresql-$v-pgvectorscale` | `postgresql-$v-pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pgvectorscale_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pgvectorscale_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pgvectorscale_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pgvectorscale_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pgvectorscale_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pgvectorscale_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgvectorscale_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgvectorscale : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgvectorscale : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgvectorscale : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgvectorscale : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgvectorscale : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgvectorscale : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgvectorscale : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgvectorscale : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgvectorscale : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_17` | 0.8.0 | `el8.x86_64` | pigsty | 403.5 KiB | [pgvectorscale_17-0.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_17-0.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_17` | 0.8.0 | `el8.aarch64` | pigsty | 364.4 KiB | [pgvectorscale_17-0.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_17-0.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_17` | 0.8.0 | `el9.x86_64` | pigsty | 411.4 KiB | [pgvectorscale_17-0.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_17-0.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_17` | 0.8.0 | `el9.aarch64` | pigsty | 388.7 KiB | [pgvectorscale_17-0.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_17-0.8.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pgvectorscale` | 0.8.0 | `d12.x86_64` | pigsty | 327.2 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgvectorscale` | 0.8.0 | `d12.aarch64` | pigsty | 280.2 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgvectorscale` | 0.8.0 | `u22.x86_64` | pigsty | 360.1 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgvectorscale` | 0.8.0 | `u22.aarch64` | pigsty | 328.9 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgvectorscale` | 0.8.0 | `u24.x86_64` | pigsty | 356.9 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgvectorscale` | 0.8.0 | `u24.aarch64` | pigsty | 325.7 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_16` | 0.8.0 | `el8.x86_64` | pigsty | 403.5 KiB | [pgvectorscale_16-0.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_16-0.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_16` | 0.8.0 | `el8.aarch64` | pigsty | 364.3 KiB | [pgvectorscale_16-0.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_16-0.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_16` | 0.8.0 | `el9.x86_64` | pigsty | 410.4 KiB | [pgvectorscale_16-0.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_16-0.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_16` | 0.8.0 | `el9.aarch64` | pigsty | 388.4 KiB | [pgvectorscale_16-0.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_16-0.8.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgvectorscale` | 0.8.0 | `d12.x86_64` | pigsty | 327.2 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgvectorscale` | 0.8.0 | `d12.aarch64` | pigsty | 280.0 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgvectorscale` | 0.8.0 | `u22.x86_64` | pigsty | 359.7 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgvectorscale` | 0.8.0 | `u22.aarch64` | pigsty | 328.9 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgvectorscale` | 0.8.0 | `u24.x86_64` | pigsty | 356.9 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgvectorscale` | 0.8.0 | `u24.aarch64` | pigsty | 325.5 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_15` | 0.8.0 | `el8.x86_64` | pigsty | 403.3 KiB | [pgvectorscale_15-0.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_15-0.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_15` | 0.8.0 | `el8.aarch64` | pigsty | 364.2 KiB | [pgvectorscale_15-0.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_15-0.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_15` | 0.8.0 | `el9.x86_64` | pigsty | 411.4 KiB | [pgvectorscale_15-0.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_15-0.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_15` | 0.8.0 | `el9.aarch64` | pigsty | 388.5 KiB | [pgvectorscale_15-0.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_15-0.8.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgvectorscale` | 0.8.0 | `d12.x86_64` | pigsty | 327.1 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgvectorscale` | 0.8.0 | `d12.aarch64` | pigsty | 280.1 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgvectorscale` | 0.8.0 | `u22.x86_64` | pigsty | 360.5 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgvectorscale` | 0.8.0 | `u22.aarch64` | pigsty | 329.3 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgvectorscale` | 0.8.0 | `u24.x86_64` | pigsty | 356.8 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgvectorscale` | 0.8.0 | `u24.aarch64` | pigsty | 325.5 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_14` | 0.8.0 | `el8.x86_64` | pigsty | 403.1 KiB | [pgvectorscale_14-0.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_14-0.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_14` | 0.8.0 | `el8.aarch64` | pigsty | 364.2 KiB | [pgvectorscale_14-0.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_14-0.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_14` | 0.8.0 | `el9.x86_64` | pigsty | 410.2 KiB | [pgvectorscale_14-0.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_14-0.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_14` | 0.8.0 | `el9.aarch64` | pigsty | 388.2 KiB | [pgvectorscale_14-0.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_14-0.8.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgvectorscale` | 0.8.0 | `d12.x86_64` | pigsty | 326.9 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgvectorscale` | 0.8.0 | `d12.aarch64` | pigsty | 279.9 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgvectorscale` | 0.8.0 | `u22.x86_64` | pigsty | 359.9 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgvectorscale` | 0.8.0 | `u22.aarch64` | pigsty | 328.8 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgvectorscale` | 0.8.0 | `u24.x86_64` | pigsty | 356.4 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgvectorscale` | 0.8.0 | `u24.aarch64` | pigsty | 325.0 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_13` | 0.8.0 | `el8.x86_64` | pigsty | 403.1 KiB | [pgvectorscale_13-0.8.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_13-0.8.0-1PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_13` | 0.8.0 | `el8.aarch64` | pigsty | 364.2 KiB | [pgvectorscale_13-0.8.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_13-0.8.0-1PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_13` | 0.8.0 | `el9.x86_64` | pigsty | 410.1 KiB | [pgvectorscale_13-0.8.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_13-0.8.0-1PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_13` | 0.8.0 | `el9.aarch64` | pigsty | 388.3 KiB | [pgvectorscale_13-0.8.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_13-0.8.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pgvectorscale` | 0.8.0 | `d12.x86_64` | pigsty | 326.7 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgvectorscale` | 0.8.0 | `d12.aarch64` | pigsty | 279.7 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgvectorscale` | 0.8.0 | `u22.x86_64` | pigsty | 360.0 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgvectorscale` | 0.8.0 | `u22.aarch64` | pigsty | 329.0 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgvectorscale` | 0.8.0 | `u24.x86_64` | pigsty | 356.2 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgvectorscale` | 0.8.0 | `u24.aarch64` | pigsty | 324.9 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pgvectorscale" title="Repository" icon="github" subtitle="github.com/timescale/pgvectorscale" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgvectorscale-0.8.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get vectorscale; # get vectorscale source code
pig build dep vectorscale; # install build dependencies
pig build pkg vectorscale; # build extension rpm or deb
pig build ext vectorscale; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install vectorscale; # install by extension name, for the current active PG version
pig ext install pgvectorscale; # install via package alias, for the active PG version
pig ext install vectorscale -v 17;   # install for PG 17
pig ext install vectorscale -v 16;   # install for PG 16
pig ext install vectorscale -v 15;   # install for PG 15
pig ext install vectorscale -v 14;   # install for PG 14
pig ext install vectorscale -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION vectorscale;
```



--------

## Usage


```sql
CREATE EXTENSION vectorscale CASCADE;

CREATE TABLE IF NOT EXISTS document_embedding  (
    id BIGINT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    metadata JSONB,
    contents TEXT,
    embedding VECTOR(1536)
);
  
CREATE INDEX document_embedding_idx ON document_embedding
USING diskann (embedding);

SELECT *
FROM document_embedding
ORDER BY embedding <=> $1
LIMIT 10
```

This fdw is read-only for now.

