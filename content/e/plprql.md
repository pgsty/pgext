---
title: "plprql"
linkTitle: "plprql"
description: "Use PRQL in PostgreSQL - Pipelined Relational Query Language"
weight: 3040
categories: ["LANG"]
width: full
---

Use PRQL in PostgreSQL - Pipelined Relational Query Language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3040** | {{< badge content="plprql" link="https://github.com/kaspermarstal/plprql" >}} | {{< ext "plprql" >}} | `18.0.0` | {{< category "LANG" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_tle" >}} {{< ext "plpgsql" >}} {{< ext "plv8" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pllua" >}} {{< ext "hstore_pllua" >}} {{< ext "plluau" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/plprql" >}} | `18.0.0` | {{< bg "18" "plprql_18" "green" >}} {{< bg "17" "plprql_17" "green" >}} {{< bg "16" "plprql_16" "green" >}} {{< bg "15" "plprql_15" "green" >}} {{< bg "14" "plprql_14" "green" >}} {{< bg "13" "plprql_13" "green" >}} | `plprql_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/plprql" >}} | `18.0.0` | {{< bg "18" "postgresql-18-plprql" "green" >}} {{< bg "17" "postgresql-17-plprql" "green" >}} {{< bg "16" "postgresql-16-plprql" "green" >}} {{< bg "15" "postgresql-15-plprql" "green" >}} {{< bg "14" "postgresql-14-plprql" "green" >}} {{< bg "13" "postgresql-13-plprql" "green" >}} | `postgresql-$v-plprql` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "plprql_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "plprql_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "plprql_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "plprql_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_17 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "plprql_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "plprql_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "plprql_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "plprql_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "plprql_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-plprql : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-plprql : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-plprql : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-plprql : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-plprql : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-plprql : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-plprql : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-plprql : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-plprql : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-16-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-plprql : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-plprql : AVAIL 1" "green" >}} |


{{< tabs items="PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_16` | 1.0.0 | `el8.x86_64` | pigsty | 2.1 MiB | [plprql_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_16` | 1.0.0 | `el8.aarch64` | pigsty | 1.9 MiB | [plprql_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_16` | 1.0.0 | `el9.x86_64` | pigsty | 2.0 MiB | [plprql_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_16` | 1.0.0 | `el9.aarch64` | pigsty | 1.9 MiB | [plprql_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-plprql` | 1.0.0 | `d12.x86_64` | pigsty | 1.7 MiB | [postgresql-16-plprql_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-16-plprql_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plprql` | 1.0.0 | `d12.aarch64` | pigsty | 1.5 MiB | [postgresql-16-plprql_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-16-plprql_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plprql` | 1.0.0 | `u22.x86_64` | pigsty | 1.9 MiB | [postgresql-16-plprql_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-16-plprql_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plprql` | 1.0.0 | `u22.aarch64` | pigsty | 1.8 MiB | [postgresql-16-plprql_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-16-plprql_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plprql` | 1.0.0 | `u24.x86_64` | pigsty | 1.9 MiB | [postgresql-16-plprql_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-16-plprql_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plprql` | 1.0.0 | `u24.aarch64` | pigsty | 1.8 MiB | [postgresql-16-plprql_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-16-plprql_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_15` | 1.0.0 | `el8.x86_64` | pigsty | 2.1 MiB | [plprql_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_15` | 1.0.0 | `el8.aarch64` | pigsty | 1.9 MiB | [plprql_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_15` | 1.0.0 | `el9.x86_64` | pigsty | 2.0 MiB | [plprql_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_15` | 1.0.0 | `el9.aarch64` | pigsty | 1.9 MiB | [plprql_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-plprql` | 1.0.0 | `d12.x86_64` | pigsty | 1.7 MiB | [postgresql-15-plprql_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-15-plprql_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plprql` | 1.0.0 | `d12.aarch64` | pigsty | 1.5 MiB | [postgresql-15-plprql_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-15-plprql_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plprql` | 1.0.0 | `u22.x86_64` | pigsty | 1.9 MiB | [postgresql-15-plprql_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-15-plprql_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plprql` | 1.0.0 | `u22.aarch64` | pigsty | 1.8 MiB | [postgresql-15-plprql_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-15-plprql_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plprql` | 1.0.0 | `u24.x86_64` | pigsty | 1.9 MiB | [postgresql-15-plprql_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-15-plprql_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plprql` | 1.0.0 | `u24.aarch64` | pigsty | 1.8 MiB | [postgresql-15-plprql_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-15-plprql_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_14` | 1.0.0 | `el8.x86_64` | pigsty | 2.1 MiB | [plprql_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_14` | 1.0.0 | `el8.aarch64` | pigsty | 1.9 MiB | [plprql_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_14` | 1.0.0 | `el9.x86_64` | pigsty | 2.0 MiB | [plprql_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_14` | 1.0.0 | `el9.aarch64` | pigsty | 1.9 MiB | [plprql_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-plprql` | 1.0.0 | `d12.x86_64` | pigsty | 1.7 MiB | [postgresql-14-plprql_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-14-plprql_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plprql` | 1.0.0 | `d12.aarch64` | pigsty | 1.5 MiB | [postgresql-14-plprql_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-14-plprql_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plprql` | 1.0.0 | `u22.x86_64` | pigsty | 1.9 MiB | [postgresql-14-plprql_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-14-plprql_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plprql` | 1.0.0 | `u22.aarch64` | pigsty | 1.8 MiB | [postgresql-14-plprql_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-14-plprql_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plprql` | 1.0.0 | `u24.x86_64` | pigsty | 1.9 MiB | [postgresql-14-plprql_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-14-plprql_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plprql` | 1.0.0 | `u24.aarch64` | pigsty | 1.8 MiB | [postgresql-14-plprql_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-14-plprql_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprql_13` | 1.0.0 | `el8.x86_64` | pigsty | 2.1 MiB | [plprql_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plprql_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `plprql_13` | 1.0.0 | `el8.aarch64` | pigsty | 1.9 MiB | [plprql_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plprql_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `plprql_13` | 1.0.0 | `el9.x86_64` | pigsty | 2.0 MiB | [plprql_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plprql_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `plprql_13` | 1.0.0 | `el9.aarch64` | pigsty | 1.9 MiB | [plprql_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plprql_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-plprql` | 1.0.0 | `d12.x86_64` | pigsty | 1.7 MiB | [postgresql-13-plprql_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-13-plprql_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-plprql` | 1.0.0 | `d12.aarch64` | pigsty | 1.5 MiB | [postgresql-13-plprql_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plprql/postgresql-13-plprql_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-plprql` | 1.0.0 | `u22.x86_64` | pigsty | 1.9 MiB | [postgresql-13-plprql_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-13-plprql_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-plprql` | 1.0.0 | `u22.aarch64` | pigsty | 1.8 MiB | [postgresql-13-plprql_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plprql/postgresql-13-plprql_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-plprql` | 1.0.0 | `u24.x86_64` | pigsty | 1.9 MiB | [postgresql-13-plprql_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-13-plprql_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-plprql` | 1.0.0 | `u24.aarch64` | pigsty | 1.8 MiB | [postgresql-13-plprql_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plprql/postgresql-13-plprql_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kaspermarstal/plprql" title="Repository" icon="github" subtitle="github.com/kaspermarstal/plprql" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plprql-18.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get plprql; # get plprql source code
pig build dep plprql; # install build dependencies
pig build pkg plprql; # build extension rpm or deb
pig build ext plprql; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install plprql; # install by extension name, for the current active PG version
pig ext install plprql; # install via package alias, for the active PG version
pig ext install plprql -v 18;   # install for PG 18
pig ext install plprql -v 17;   # install for PG 17
pig ext install plprql -v 16;   # install for PG 16
pig ext install plprql -v 15;   # install for PG 15
pig ext install plprql -v 14;   # install for PG 14
pig ext install plprql -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION plprql;
```

