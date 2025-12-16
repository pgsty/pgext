---
title: "pg_ai_query"
linkTitle: "pg_ai_query"
description: "AI-powered SQL query generation for PostgreSQL"
weight: 2760
categories: ["FEAT"]
width: full
---

[**pg_ai_query**](https://github.com/benodiwal/pg_ai_query) : AI-powered SQL query generation for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2760** | {{< badge content="pg_ai_query" link="https://github.com/benodiwal/pg_ai_query" >}} | {{< ext "pg_ai_query" >}} | `0.1.1` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgml" >}} {{< ext "pg4ml" >}} {{< ext "vectorize" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_ai_query` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "pg_ai_query_18*" "green" >}} {{< bg "17" "pg_ai_query_17*" "green" >}} {{< bg "16" "pg_ai_query_16*" "green" >}} {{< bg "15" "pg_ai_query_15*" "green" >}} {{< bg "14" "pg_ai_query_14*" "green" >}} {{< bg "13" "pg_ai_query_13*" "red" >}} | `pg_ai_query_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-ai-query" "green" >}} {{< bg "17" "postgresql-17-ai-query" "green" >}} {{< bg "16" "postgresql-16-ai-query" "green" >}} {{< bg "15" "postgresql-15-ai-query" "green" >}} {{< bg "14" "postgresql-14-ai-query" "green" >}} {{< bg "13" "postgresql-13-ai-query" "red" >}} | `postgresql-$v-ai-query` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_ai_query_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_ai_query_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ai_query_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_ai_query_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_ai_query_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_ai_query_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pg_ai_query_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_ai_query_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ai-query : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ai-query : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-ai-query : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-ai-query : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-ai-query : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-ai-query : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ai-query : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-ai-query : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ai-query : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-ai-query : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-ai-query : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-ai-query : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-ai-query : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-ai-query : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ai_query_18` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 967.9 KiB | [pg_ai_query_18-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ai_query_18-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_ai_query_18` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 880.6 KiB | [pg_ai_query_18-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ai_query_18-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_ai_query_18` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 899.5 KiB | [pg_ai_query_18-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ai_query_18-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_ai_query_18` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 846.4 KiB | [pg_ai_query_18-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ai_query_18-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-ai-query` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 842.3 KiB | [postgresql-18-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-18-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-ai-query` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 760.5 KiB | [postgresql-18-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-18-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-ai-query` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.0 KiB | [postgresql-18-ai-query_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-18-ai-query_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-ai-query` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 771.8 KiB | [postgresql-18-ai-query_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-18-ai-query_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ai_query_17` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 967.8 KiB | [pg_ai_query_17-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ai_query_17-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_ai_query_17` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 881.9 KiB | [pg_ai_query_17-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ai_query_17-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_ai_query_17` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 899.4 KiB | [pg_ai_query_17-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ai_query_17-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_ai_query_17` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 846.4 KiB | [pg_ai_query_17-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ai_query_17-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-ai-query` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 845.8 KiB | [postgresql-17-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-17-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-ai-query` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 760.1 KiB | [postgresql-17-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-17-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-ai-query` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.0 KiB | [postgresql-17-ai-query_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-17-ai-query_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-ai-query` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 771.8 KiB | [postgresql-17-ai-query_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-17-ai-query_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ai_query_16` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 967.8 KiB | [pg_ai_query_16-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ai_query_16-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_ai_query_16` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 881.0 KiB | [pg_ai_query_16-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ai_query_16-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_ai_query_16` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 899.5 KiB | [pg_ai_query_16-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ai_query_16-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_ai_query_16` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 846.4 KiB | [pg_ai_query_16-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ai_query_16-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-ai-query` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 841.1 KiB | [postgresql-16-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-16-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-ai-query` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 759.0 KiB | [postgresql-16-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-16-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-ai-query` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.6 KiB | [postgresql-16-ai-query_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-16-ai-query_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-ai-query` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 771.8 KiB | [postgresql-16-ai-query_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-16-ai-query_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ai_query_15` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 967.4 KiB | [pg_ai_query_15-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ai_query_15-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_ai_query_15` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 881.0 KiB | [pg_ai_query_15-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ai_query_15-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_ai_query_15` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 899.3 KiB | [pg_ai_query_15-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ai_query_15-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_ai_query_15` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 846.4 KiB | [pg_ai_query_15-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ai_query_15-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-ai-query` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 841.4 KiB | [postgresql-15-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-15-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-ai-query` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 758.6 KiB | [postgresql-15-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-15-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-ai-query` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.2 KiB | [postgresql-15-ai-query_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-15-ai-query_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-ai-query` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 771.6 KiB | [postgresql-15-ai-query_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-15-ai-query_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ai_query_14` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 967.8 KiB | [pg_ai_query_14-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ai_query_14-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_ai_query_14` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 881.0 KiB | [pg_ai_query_14-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ai_query_14-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_ai_query_14` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 899.3 KiB | [pg_ai_query_14-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ai_query_14-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_ai_query_14` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 846.4 KiB | [pg_ai_query_14-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ai_query_14-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-ai-query` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 841.6 KiB | [postgresql-14-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-14-ai-query_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-ai-query` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 760.7 KiB | [postgresql-14-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ai-query/postgresql-14-ai-query_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-ai-query` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.2 KiB | [postgresql-14-ai-query_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-14-ai-query_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-ai-query` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 771.7 KiB | [postgresql-14-ai-query_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ai-query/postgresql-14-ai-query_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/benodiwal/pg_ai_query" title="Repository" icon="github" subtitle="github.com/benodiwal/pg_ai_query" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_ai_query-0.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_ai_query;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_ai_query;		# install via package name, for the active PG version

pig install pg_ai_query -v 18;   # install for PG 18
pig install pg_ai_query -v 17;   # install for PG 17
pig install pg_ai_query -v 16;   # install for PG 16
pig install pg_ai_query -v 15;   # install for PG 15
pig install pg_ai_query -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_ai_query;
```
