---
title: "pg_summarize"
linkTitle: "pg_summarize"
description: "Text Summarization using LLMs. Built using pgrx"
weight: 1860
categories: ["RAG"]
width: full
---

[**pg_summarize**](https://github.com/HexaCluster/pg_summarize) : Text Summarization using LLMs. Built using pgrx


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1860** | {{< badge content="pg_summarize" link="https://github.com/HexaCluster/pg_summarize" >}} | {{< ext "pg_summarize" >}} | `0.0.1` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vectorize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "pg_net" >}} |

> [!Note] PG18 fix by https://github.com/Vonng/pg_summarize


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_summarize` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_summarize_18" "green" >}} {{< bg "17" "pg_summarize_17" "green" >}} {{< bg "16" "pg_summarize_16" "green" >}} {{< bg "15" "pg_summarize_15" "green" >}} {{< bg "14" "pg_summarize_14" "green" >}} | `pg_summarize_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-summarize" "green" >}} {{< bg "17" "postgresql-17-pg-summarize" "green" >}} {{< bg "16" "postgresql-16-pg-summarize" "green" >}} {{< bg "15" "postgresql-15-pg-summarize" "green" >}} {{< bg "14" "postgresql-14-pg-summarize" "green" >}} | `postgresql-$v-pg-summarize` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_summarize_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-summarize : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-summarize : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-summarize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-summarize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-summarize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-summarize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-summarize : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-summarize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-summarize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-summarize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-summarize : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-summarize : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_18-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_18-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.5 KiB | [pg_summarize_18-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_18-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_18-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_18-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.3 KiB | [pg_summarize_18-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_18-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_18-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_18-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1001.1 KiB | [pg_summarize_18-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_18-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 879.2 KiB | [postgresql-18-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 711.8 KiB | [postgresql-18-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 879.5 KiB | [postgresql-18-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 712.0 KiB | [postgresql-18-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 980.2 KiB | [postgresql-18-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 842.8 KiB | [postgresql-18-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 970.8 KiB | [postgresql-18-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 834.7 KiB | [postgresql-18-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-18-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.4 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.2 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_17-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_17-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1002.2 KiB | [pg_summarize_17-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_17-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 879.1 KiB | [postgresql-17-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 711.8 KiB | [postgresql-17-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 879.1 KiB | [postgresql-17-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 712.3 KiB | [postgresql-17-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 980.3 KiB | [postgresql-17-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 843.1 KiB | [postgresql-17-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 971.0 KiB | [postgresql-17-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 835.3 KiB | [postgresql-17-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-17-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.5 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.2 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_16-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_16-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1001.3 KiB | [pg_summarize_16-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_16-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 878.8 KiB | [postgresql-16-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 711.7 KiB | [postgresql-16-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 879.8 KiB | [postgresql-16-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 711.8 KiB | [postgresql-16-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 980.1 KiB | [postgresql-16-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 843.3 KiB | [postgresql-16-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 971.0 KiB | [postgresql-16-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 834.8 KiB | [postgresql-16-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-16-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.6 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1002.8 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_15-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_15-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1001.1 KiB | [pg_summarize_15-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_15-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 878.9 KiB | [postgresql-15-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 713.3 KiB | [postgresql-15-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 878.9 KiB | [postgresql-15-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 712.1 KiB | [postgresql-15-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 979.9 KiB | [postgresql-15-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 842.7 KiB | [postgresql-15-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 970.5 KiB | [postgresql-15-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 834.8 KiB | [postgresql-15-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-15-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_summarize_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_summarize_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_summarize_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 948.5 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_summarize_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_summarize_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_summarize_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1001.7 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_summarize_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_summarize_14-0.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_summarize_14-0.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_summarize_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1001.2 KiB | [pg_summarize_14-0.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_summarize_14-0.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-summarize` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 879.0 KiB | [postgresql-14-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 712.0 KiB | [postgresql-14-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 879.6 KiB | [postgresql-14-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 712.3 KiB | [postgresql-14-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 980.1 KiB | [postgresql-14-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 843.2 KiB | [postgresql-14-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 970.9 KiB | [postgresql-14-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-summarize` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 834.9 KiB | [postgresql-14-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-summarize/postgresql-14-pg-summarize_0.0.1-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_summarize" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_summarize" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_summarize-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_summarize;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_summarize;		# install via package name, for the active PG version

pig install pg_summarize -v 18;   # install for PG 18
pig install pg_summarize -v 17;   # install for PG 17
pig install pg_summarize -v 16;   # install for PG 16
pig install pg_summarize -v 15;   # install for PG 15
pig install pg_summarize -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_summarize;
```



## Usage

> [pg_summarize](https://github.com/HexaCluster/pg_summarize): Text Summarization using LLMs, built using pgrx.
> Source: [README.md](https://raw.githubusercontent.com/HexaCluster/pg_summarize/main/README.md)

`pg_summarize` is a PostgreSQL extension written in Rust (using `pgrx`) that integrates with the OpenAI API. It includes a basic "Hello, pg_summarize!" function and a `summarize` function that summarizes text using OpenAI's models.


--------

### Getting Started

```sql
CREATE EXTENSION pg_summarize;

-- Test the hello function
SELECT hello_pg_summarize();
--  hello_pg_summarize
-- ----------------------
--  Hello, pg_summarize
```


--------

### Configuration

The extension retrieves configuration from PostgreSQL settings. Set the following before using the `summarize` function:

```sql
-- Set the OpenAI API key (required)
ALTER SYSTEM SET pg_summarizer.api_key = 'your_openai_api_key';

-- Optionally set the model (default: gpt-3.5-turbo)
ALTER SYSTEM SET pg_summarizer.model = 'gpt-3.5-turbo';

-- Or set the prompt at session level
SET pg_summarizer.prompt = 'Your custom prompt here';

-- Reload the configuration if set at SYSTEM level
SELECT pg_reload_conf();
```


--------

### Summarize Function

The `summarize` function takes text input, sends it to the OpenAI API, and returns a summary:

```sql
-- Summarize a text input
SELECT summarize('<This is the text to be summarized.>');

-- Create a summary table from existing data
CREATE TABLE blogs_summary AS
  SELECT blog_url, summarize(blogs_text)
  FROM hexacluster_blogs;

-- Use a different model
SET pg_summarizer.model = 'gpt-4o';
CREATE TABLE blogs_summary_4o AS
  SELECT blog_url, summarize(blogs_text)
  FROM hexacluster_blogs;
```


--------

### How It Works

- **Configuration Retrieval**: The `summarize` function retrieves settings (API key, model, prompt) from PostgreSQL using `current_setting()`. Defaults are used if settings are not found.
- **Default Prompt**: A built-in prompt instructs the AI to summarize text from `<text>` tags, focusing on capturing the most important information concisely.
- **API Call**: The function sends a POST request to the OpenAI chat completions endpoint with the configured model and prompt, returning the summary content.
