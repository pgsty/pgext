---
title: "pgml"
linkTitle: "pgml"
description: "Run AL/ML workloads with SQL interface"
weight: 1890
categories: ["RAG"]
width: full
---

Run AL/ML workloads with SQL interface


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1890** | {{< badge content="pgml" link="https://github.com/postgresml/postgresml" >}} | {{< ext "pgml" >}} | `2.10.0` | {{< category "RAG" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg4ml" >}} {{< ext "vectorize" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} {{< ext "pg_strom" >}} |

> [!Note] pgrx=0.12.9


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgml" >}} | `2.10.0` | {{< bg "18" "pgml_18" "red" >}} {{< bg "17" "pgml_17" "green" >}} {{< bg "16" "pgml_16" "green" >}} {{< bg "15" "pgml_15" "green" >}} {{< bg "14" "pgml_14" "green" >}} {{< bg "13" "pgml_13" "green" >}} | `pgml_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgml" >}} | `2.10.0` | {{< bg "18" "postgresql-18-pgml" "red" >}} {{< bg "17" "postgresql-17-pgml" "green" >}} {{< bg "16" "postgresql-16-pgml" "green" >}} {{< bg "15" "postgresql-15-pgml" "green" >}} {{< bg "14" "postgresql-14-pgml" "green" >}} {{< bg "13" "postgresql-13-pgml" "red" >}} | `postgresql-$v-pgml` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pgml_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "pgml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pgml_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "pgml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pgml_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "pgml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.9.3" "pgml_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.9.3" "pgml_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.9.3" "pgml_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pgml_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "pgml_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "pgml_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pgml_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pgml_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgml_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgml : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-pgml : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgml : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-pgml : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgml : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgml : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgml : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgml : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-pgml : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgml : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-pgml : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgml : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-pgml : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgml : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-pgml : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-pgml : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgml_17` | 2.10.0 | `el8.x86_64` | pigsty | 5.7 MiB | [pgml_17-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgml_17-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pgml_17` | 2.10.0 | `el8.aarch64` | pigsty | 4.8 MiB | [pgml_17-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgml_17-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pgml_17` | 2.10.0 | `el9.x86_64` | pigsty | 5.3 MiB | [pgml_17-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgml_17-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pgml_17` | 2.10.0 | `el9.aarch64` | pigsty | 5.1 MiB | [pgml_17-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgml_17-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pgml` | 2.10.0 | `d12.x86_64` | pigsty | 4.7 MiB | [postgresql-17-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-17-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgml` | 2.10.0 | `d12.aarch64` | pigsty | 4.0 MiB | [postgresql-17-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-17-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgml` | 2.10.0 | `u22.x86_64` | pigsty | 5.1 MiB | [postgresql-17-pgml_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-17-pgml_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgml` | 2.10.0 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-17-pgml_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-17-pgml_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgml` | 2.10.0 | `u24.x86_64` | pigsty | 5.4 MiB | [postgresql-17-pgml_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-17-pgml_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgml` | 2.10.0 | `u24.aarch64` | pigsty | 4.9 MiB | [postgresql-17-pgml_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-17-pgml_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgml_16` | 2.10.0 | `el8.x86_64` | pigsty | 5.7 MiB | [pgml_16-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgml_16-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pgml_16` | 2.10.0 | `el8.aarch64` | pigsty | 4.8 MiB | [pgml_16-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgml_16-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pgml_16` | 2.9.3 | `el9.x86_64` | pigsty | 4.6 MiB | [pgml_16-2.9.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgml_16-2.9.3-1PIGSTY.el9.x86_64.rpm) |
| `pgml_16` | 2.10.0 | `el9.x86_64` | pigsty | 5.3 MiB | [pgml_16-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgml_16-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pgml_16` | 2.10.0 | `el9.aarch64` | pigsty | 5.1 MiB | [pgml_16-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgml_16-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgml` | 2.10.0 | `d12.x86_64` | pigsty | 4.7 MiB | [postgresql-16-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-16-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgml` | 2.10.0 | `d12.aarch64` | pigsty | 4.0 MiB | [postgresql-16-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-16-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgml` | 2.10.0 | `u22.x86_64` | pigsty | 5.1 MiB | [postgresql-16-pgml_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-16-pgml_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgml` | 2.10.0 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-16-pgml_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-16-pgml_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgml` | 2.10.0 | `u24.x86_64` | pigsty | 5.4 MiB | [postgresql-16-pgml_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-16-pgml_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgml` | 2.10.0 | `u24.aarch64` | pigsty | 4.9 MiB | [postgresql-16-pgml_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-16-pgml_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgml_15` | 2.10.0 | `el8.x86_64` | pigsty | 5.7 MiB | [pgml_15-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgml_15-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pgml_15` | 2.10.0 | `el8.aarch64` | pigsty | 4.8 MiB | [pgml_15-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgml_15-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pgml_15` | 2.9.3 | `el9.x86_64` | pigsty | 4.6 MiB | [pgml_15-2.9.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgml_15-2.9.3-1PIGSTY.el9.x86_64.rpm) |
| `pgml_15` | 2.10.0 | `el9.x86_64` | pigsty | 5.3 MiB | [pgml_15-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgml_15-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pgml_15` | 2.10.0 | `el9.aarch64` | pigsty | 5.1 MiB | [pgml_15-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgml_15-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgml` | 2.10.0 | `d12.x86_64` | pigsty | 4.7 MiB | [postgresql-15-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-15-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgml` | 2.10.0 | `d12.aarch64` | pigsty | 4.0 MiB | [postgresql-15-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-15-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgml` | 2.10.0 | `u22.x86_64` | pigsty | 5.1 MiB | [postgresql-15-pgml_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-15-pgml_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgml` | 2.10.0 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-15-pgml_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-15-pgml_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgml` | 2.10.0 | `u24.x86_64` | pigsty | 5.4 MiB | [postgresql-15-pgml_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-15-pgml_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgml` | 2.10.0 | `u24.aarch64` | pigsty | 4.9 MiB | [postgresql-15-pgml_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-15-pgml_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgml_14` | 2.10.0 | `el8.x86_64` | pigsty | 5.7 MiB | [pgml_14-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgml_14-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pgml_14` | 2.10.0 | `el8.aarch64` | pigsty | 4.8 MiB | [pgml_14-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgml_14-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pgml_14` | 2.9.3 | `el9.x86_64` | pigsty | 4.6 MiB | [pgml_14-2.9.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgml_14-2.9.3-1PIGSTY.el9.x86_64.rpm) |
| `pgml_14` | 2.10.0 | `el9.x86_64` | pigsty | 5.3 MiB | [pgml_14-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgml_14-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pgml_14` | 2.10.0 | `el9.aarch64` | pigsty | 5.1 MiB | [pgml_14-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgml_14-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgml` | 2.10.0 | `d12.x86_64` | pigsty | 4.7 MiB | [postgresql-14-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-14-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgml` | 2.10.0 | `d12.aarch64` | pigsty | 4.0 MiB | [postgresql-14-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-14-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgml` | 2.10.0 | `u22.x86_64` | pigsty | 5.1 MiB | [postgresql-14-pgml_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-14-pgml_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgml` | 2.10.0 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-14-pgml_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-14-pgml_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgml` | 2.10.0 | `u24.x86_64` | pigsty | 5.4 MiB | [postgresql-14-pgml_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-14-pgml_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgml` | 2.10.0 | `u24.aarch64` | pigsty | 4.9 MiB | [postgresql-14-pgml_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-14-pgml_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgml_13` | 2.10.0 | `el8.x86_64` | pigsty | 5.7 MiB | [pgml_13-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgml_13-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `pgml_13` | 2.10.0 | `el8.aarch64` | pigsty | 4.8 MiB | [pgml_13-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgml_13-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `pgml_13` | 2.10.0 | `el9.x86_64` | pigsty | 5.3 MiB | [pgml_13-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgml_13-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `pgml_13` | 2.10.0 | `el9.aarch64` | pigsty | 5.1 MiB | [pgml_13-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgml_13-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pgml` | 2.10.0 | `d12.x86_64` | pigsty | 4.7 MiB | [postgresql-13-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-13-pgml_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgml` | 2.10.0 | `d12.aarch64` | pigsty | 4.0 MiB | [postgresql-13-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgml/postgresql-13-pgml_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgml` | 2.10.0 | `u22.x86_64` | pigsty | 5.1 MiB | [postgresql-13-pgml_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-13-pgml_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgml` | 2.10.0 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-13-pgml_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgml/postgresql-13-pgml_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgml` | 2.10.0 | `u24.x86_64` | pigsty | 5.4 MiB | [postgresql-13-pgml_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-13-pgml_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgml` | 2.10.0 | `u24.aarch64` | pigsty | 4.9 MiB | [postgresql-13-pgml_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgml/postgresql-13-pgml_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgresml/postgresml" title="Repository" icon="github" subtitle="github.com/postgresml/postgresml" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgml-2.10.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgml; # get pgml source code
pig build dep pgml; # install build dependencies
pig build pkg pgml; # build extension rpm or deb
pig build ext pgml; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgml; # install by extension name, for the current active PG version
pig ext install pgml; # install via package alias, for the active PG version
pig ext install pgml -v 17;   # install for PG 17
pig ext install pgml -v 16;   # install for PG 16
pig ext install pgml -v 15;   # install for PG 15
pig ext install pgml -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgml CASCADE SCHEMA pgml;
```



--------

## Usage

<Callout title="This extension is lack of maintenance" type="warn">

The PGML team seems not maintaining this extension for a while.

</Callout>

After installing the `pgml` extension and python dependencies on all cluster nodes, you can enable `pgml` on the PostgreSQL cluster.

[Configure](/pgsql/admin/#config-cluster) cluster with `patronictl` command and add `pgml` to `shared_preload_libraries`, and specify your `venv` dir in `pgml.venv`:

```yaml
shared_preload_libraries: pgml, timescaledb, pg_stat_statements, auto_explain
pgml.venv: '/data/pgml'
```

After that, restart database cluster, and create extension with SQL command:

```sql
CREATE EXTENSION vector;        -- nice to have pgvector installed too!
CREATE EXTENSION pgml;          -- create PostgresML in current database
SELECT pgml.version();          -- print PostgresML version string
```

If it works, you should see something like:

```bash
# create extension pgml;
INFO:  Python version: 3.11.2 (main, Oct  5 2023, 16:06:03) [GCC 8.5.0 20210514 (Red Hat 8.5.0-18)]
INFO:  Scikit-learn 1.3.0, XGBoost 2.0.0, LightGBM 4.1.0, NumPy 1.26.1
CREATE EXTENSION

# SELECT pgml.version(); -- print PostgresML version string
 version
---------
 2.7.8
```

You are all set! Check PostgresML for more details: https://postgresml.org/docs/guides/use-cases/


