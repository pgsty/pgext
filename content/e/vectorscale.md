---
title: "vectorscale"
linkTitle: "vectorscale"
description: "Advanced indexing for vector data with DiskANN"
weight: 1820
categories: ["RAG"]
width: full
---

[**pgvectorscale**](https://github.com/timescale/pgvectorscale)


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

> [!Note] not an official release tag


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/vectorscale" >}} | `0.8.0` | {{< bg "18" "pgvectorscale_18" "green" >}} {{< bg "17" "pgvectorscale_17" "green" >}} {{< bg "16" "pgvectorscale_16" "green" >}} {{< bg "15" "pgvectorscale_15" "green" >}} {{< bg "14" "pgvectorscale_14" "green" >}} {{< bg "13" "pgvectorscale_13" "green" >}} | `pgvectorscale_$v` | `pgvector_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/vectorscale" >}} | `0.8.0` | {{< bg "18" "postgresql-18-pgvectorscale" "green" >}} {{< bg "17" "postgresql-17-pgvectorscale" "green" >}} {{< bg "16" "postgresql-16-pgvectorscale" "green" >}} {{< bg "15" "postgresql-15-pgvectorscale" "green" >}} {{< bg "14" "postgresql-14-pgvectorscale" "green" >}} {{< bg "13" "postgresql-13-pgvectorscale" "green" >}} | `postgresql-$v-pgvectorscale` | `postgresql-$v-pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pgvectorscale_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-13-pgvectorscale : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_18` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 503.9 KiB | [pgvectorscale_18-0.8.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_18-0.8.0-2PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_18` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 371.4 KiB | [pgvectorscale_18-0.8.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_18-0.8.0-2PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_18` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 520.6 KiB | [pgvectorscale_18-0.8.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_18-0.8.0-2PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_18` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 395.9 KiB | [pgvectorscale_18-0.8.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_18-0.8.0-2PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_18` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 521.1 KiB | [pgvectorscale_18-0.8.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_18-0.8.0-2PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_18` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 395.7 KiB | [pgvectorscale_18-0.8.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_18-0.8.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgvectorscale` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.0 KiB | [postgresql-18-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.0 KiB | [postgresql-18-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.0 KiB | [postgresql-18-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.0 KiB | [postgresql-18-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 KiB | [postgresql-18-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 KiB | [postgresql-18-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 KiB | [postgresql-18-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 KiB | [postgresql-18-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_17` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 504.1 KiB | [pgvectorscale_17-0.8.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_17-0.8.0-2PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_17` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 371.4 KiB | [pgvectorscale_17-0.8.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_17-0.8.0-2PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_17` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 520.6 KiB | [pgvectorscale_17-0.8.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_17-0.8.0-2PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_17` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 395.4 KiB | [pgvectorscale_17-0.8.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_17-0.8.0-2PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_17` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 520.4 KiB | [pgvectorscale_17-0.8.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_17-0.8.0-2PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_17` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 395.7 KiB | [pgvectorscale_17-0.8.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_17-0.8.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgvectorscale` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 413.9 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 286.2 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 414.2 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 286.2 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 463.2 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 336.7 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 459.0 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 332.4 KiB | [postgresql-17-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_16` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 504.1 KiB | [pgvectorscale_16-0.8.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_16-0.8.0-2PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_16` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 371.2 KiB | [pgvectorscale_16-0.8.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_16-0.8.0-2PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_16` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 520.9 KiB | [pgvectorscale_16-0.8.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_16-0.8.0-2PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_16` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 395.6 KiB | [pgvectorscale_16-0.8.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_16-0.8.0-2PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_16` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 521.0 KiB | [pgvectorscale_16-0.8.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_16-0.8.0-2PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_16` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 395.4 KiB | [pgvectorscale_16-0.8.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_16-0.8.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgvectorscale` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 414.3 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 286.0 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 413.4 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 286.3 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 462.9 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 336.9 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 458.8 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 332.2 KiB | [postgresql-16-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_15` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 503.9 KiB | [pgvectorscale_15-0.8.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_15-0.8.0-2PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_15` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 371.2 KiB | [pgvectorscale_15-0.8.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_15-0.8.0-2PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_15` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 520.4 KiB | [pgvectorscale_15-0.8.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_15-0.8.0-2PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_15` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 395.6 KiB | [pgvectorscale_15-0.8.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_15-0.8.0-2PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_15` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 520.4 KiB | [pgvectorscale_15-0.8.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_15-0.8.0-2PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_15` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 395.6 KiB | [pgvectorscale_15-0.8.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_15-0.8.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgvectorscale` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 414.1 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 285.9 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 413.0 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 286.3 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 463.0 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 336.8 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 458.5 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 332.3 KiB | [postgresql-15-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_14` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 503.8 KiB | [pgvectorscale_14-0.8.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_14-0.8.0-2PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_14` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 370.8 KiB | [pgvectorscale_14-0.8.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_14-0.8.0-2PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_14` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 519.5 KiB | [pgvectorscale_14-0.8.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_14-0.8.0-2PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_14` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 395.1 KiB | [pgvectorscale_14-0.8.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_14-0.8.0-2PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_14` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 520.0 KiB | [pgvectorscale_14-0.8.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_14-0.8.0-2PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_14` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 394.8 KiB | [pgvectorscale_14-0.8.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_14-0.8.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgvectorscale` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 413.2 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 285.9 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 412.6 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 285.9 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 462.0 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 336.3 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 457.3 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 331.7 KiB | [postgresql-14-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_13` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 504.4 KiB | [pgvectorscale_13-0.8.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_13-0.8.0-2PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_13` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 370.9 KiB | [pgvectorscale_13-0.8.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_13-0.8.0-2PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_13` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 520.3 KiB | [pgvectorscale_13-0.8.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_13-0.8.0-2PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_13` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 395.4 KiB | [pgvectorscale_13-0.8.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_13-0.8.0-2PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_13` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 520.3 KiB | [pgvectorscale_13-0.8.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_13-0.8.0-2PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_13` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 394.9 KiB | [pgvectorscale_13-0.8.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_13-0.8.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgvectorscale` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 413.6 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgvectorscale` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 286.1 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgvectorscale` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 413.4 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgvectorscale` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 285.8 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgvectorscale` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 462.7 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgvectorscale` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 336.4 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgvectorscale` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 458.7 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgvectorscale` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 332.4 KiB | [postgresql-13-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-13-pgvectorscale_0.8.0-1PIGSTY~noble_arm64.deb) |

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
pig ext install vectorscale -v 18;   # install for PG 18
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

