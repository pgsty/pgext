---
title: "pg_tiktoken"
linkTitle: "pg_tiktoken"
description: "tiktoken tokenizer for use with OpenAI models in postgres"
weight: 1870
categories: ["RAG"]
width: full
---

[**pg_tiktoken**](https://github.com/kelvich/pg_tiktoken) : tiktoken tokenizer for use with OpenAI models in postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1870** | {{< badge content="pg_tiktoken" link="https://github.com/kelvich/pg_tiktoken" >}} | {{< ext "pg_tiktoken" >}} | `0.0.1` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vectorize" >}} {{< ext "pg_summarize" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "pg_graphql" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_tiktoken` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_tiktoken_18" "green" >}} {{< bg "17" "pg_tiktoken_17" "green" >}} {{< bg "16" "pg_tiktoken_16" "green" >}} {{< bg "15" "pg_tiktoken_15" "green" >}} {{< bg "14" "pg_tiktoken_14" "green" >}} | `pg_tiktoken_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-tiktoken" "green" >}} {{< bg "17" "postgresql-17-pg-tiktoken" "green" >}} {{< bg "16" "postgresql-16-pg-tiktoken" "green" >}} {{< bg "15" "postgresql-15-pg-tiktoken" "green" >}} {{< bg "14" "postgresql-14-pg-tiktoken" "green" >}} | `postgresql-$v-pg-tiktoken` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_tiktoken_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-tiktoken : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-tiktoken : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-tiktoken : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-tiktoken : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-tiktoken : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-tiktoken : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-tiktoken : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-tiktoken : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-tiktoken : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-tiktoken : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-tiktoken : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-tiktoken : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_18-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_18-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [pg_tiktoken_18-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_18-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_tiktoken_18-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_18-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_18-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_18-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_18-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_18-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_18-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_18-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tiktoken` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tiktoken` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tiktoken` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tiktoken` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tiktoken` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tiktoken` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tiktoken` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tiktoken` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-18-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_17-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_17-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [pg_tiktoken_17-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_17-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_tiktoken_17-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_17-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_17-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_17-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_17-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_17-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_17-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_17-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tiktoken` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tiktoken` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tiktoken` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tiktoken` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tiktoken` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tiktoken` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tiktoken` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tiktoken` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-17-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_16-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_16-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [pg_tiktoken_16-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_16-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_tiktoken_16-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_16-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_16-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_16-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_16-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_16-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_16-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_16-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-tiktoken` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tiktoken` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tiktoken` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-tiktoken` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-tiktoken` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tiktoken` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tiktoken` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tiktoken` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-16-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_15-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_15-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [pg_tiktoken_15-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_15-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_tiktoken_15-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_15-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_15-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_15-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_15-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_15-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_15-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_15-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-tiktoken` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tiktoken` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tiktoken` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-tiktoken` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-tiktoken` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tiktoken` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tiktoken` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tiktoken` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-15-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_14-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_14-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.6 MiB | [pg_tiktoken_14-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_14-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_tiktoken_14-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_14-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_14-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_14-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.7 MiB | [pg_tiktoken_14-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_14-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.5 MiB | [pg_tiktoken_14-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_14-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-tiktoken` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tiktoken` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tiktoken` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-tiktoken` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-tiktoken` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tiktoken` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.4 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tiktoken` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.6 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tiktoken` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.4 MiB | [postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken/postgresql-14-pg-tiktoken_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kelvich/pg_tiktoken" title="Repository" icon="github" subtitle="github.com/kelvich/pg_tiktoken" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tiktoken-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_tiktoken;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_tiktoken;		# install via package name, for the active PG version

pig install pg_tiktoken -v 18;   # install for PG 18
pig install pg_tiktoken -v 17;   # install for PG 17
pig install pg_tiktoken -v 16;   # install for PG 16
pig install pg_tiktoken -v 15;   # install for PG 15
pig install pg_tiktoken -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_tiktoken;
```



## Usage

> [pg_tiktoken](https://github.com/kelvich/pg_tiktoken): tiktoken tokenizer for use with OpenAI models in PostgreSQL.
> Source: [README.md](https://raw.githubusercontent.com/kelvich/pg_tiktoken/main/README.md)

`pg_tiktoken` is a PostgreSQL extension that provides input tokenization using OpenAI's [tiktoken](https://github.com/openai/tiktoken) library. It allows you to count and encode tokens directly in SQL, which is useful for managing input length limits when working with OpenAI models.


--------

### Functions

#### tiktoken_count

Count the number of tokens for a given encoding or model:

```sql
SELECT tiktoken_count('p50k_edit', 'A long time ago in a galaxy far, far away');
 tiktoken_count
----------------
             11
(1 row)
```

#### tiktoken_encode

Get the token IDs for a given encoding or model:

```sql
SELECT tiktoken_encode('cl100k_base', 'A long time ago in a galaxy far, far away');
                  tiktoken_encode
----------------------------------------------------
 {32,1317,892,4227,304,264,34261,3117,11,3117,3201}
(1 row)
```

Both `tiktoken_count` and `tiktoken_encode` accept either an encoding name or an OpenAI model name as the first argument.


--------

### Supported Models

| Encoding name | OpenAI models |
|---|---|
| `cl100k_base` | ChatGPT models, `text-embedding-ada-002` |
| `p50k_base` | Code models, `text-davinci-002`, `text-davinci-003` |
| `p50k_edit` | Edit models like `text-davinci-edit-001`, `code-davinci-edit-001` |
| `r50k_base` (or `gpt2`) | GPT-3 models like `davinci` |
