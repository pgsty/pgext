---
title: "documentdb_core"
linkTitle: "documentdb_core"
description: "Core API surface for DocumentDB for PostgreSQL"
weight: 9010
categories: ["Sim"]
width: full
---

Core API surface for DocumentDB for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9010** | {{< badge content="documentdb_core" link="https://github.com/microsoft/documentdb" >}} | {{< ext "documentdb_core" "documentdb" >}} | `0.106` | {{< category "SIM" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "documentdb" >}} |
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "rum" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_cron" >}} {{< ext "postgis" >}} {{< ext "vector" >}} |
|    **Siblings**   | {{< ext "documentdb" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] FerretDB Fork, work with FerretDB 2.5


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/documentdb" >}} | `0.106` | {{< badge content="18" color="red" alt="documentdb_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="documentdb_14*" >}} | `documentdb_$v*` | `postgresql$v-contrib`, `pg_cron_$v`, `pgvector_$v`, `rum_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/documentdb" >}} | `0.106` | {{< badge content="18" color="red" alt="postgresql-18-documentdb" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="postgresql-14-documentdb" >}} | `postgresql-$v-documentdb` | `postgresql-$v-cron`, `postgresql-$v-pgvector`, `postgresql-$v-rum` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "documentdb_18" >}}     | {{< pkg "documentdb_17" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-0.105-0PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "documentdb_16" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-0.105-0PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "documentdb_15" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-0.105-0PIGSTY.el8.x86_64.rpm" >}} |    {{< pkg "documentdb_14" >}}     |
|    `el8.aarch64`    |    {{< pkg "documentdb_18" >}}     | {{< pkg "documentdb_17" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-0.105-0PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "documentdb_16" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-0.105-0PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "documentdb_15" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-0.105-0PIGSTY.el8.aarch64.rpm" >}} |    {{< pkg "documentdb_14" >}}     |
|    `el9.x86_64`    |    {{< pkg "documentdb_18" >}}     | {{< pkg "documentdb_17" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-0.105-0PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "documentdb_16" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-0.105-0PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "documentdb_15" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-0.105-0PIGSTY.el9.x86_64.rpm" >}} |    {{< pkg "documentdb_14" >}}     |
|    `el9.aarch64`    |    {{< pkg "documentdb_18" >}}     | {{< pkg "documentdb_17" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-0.105-0PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "documentdb_16" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-0.105-0PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "documentdb_15" "0.105" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-0.105-0PIGSTY.el9.aarch64.rpm" >}} |    {{< pkg "documentdb_14" >}}     |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-documentdb" >}}     | {{< pkg "postgresql-17-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~bookworm_amd64.deb" >}} |    {{< pkg "postgresql-14-documentdb" >}}     |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-documentdb" >}}     | {{< pkg "postgresql-17-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~bookworm_arm64.deb" >}} |    {{< pkg "postgresql-14-documentdb" >}}     |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-documentdb" >}}     | {{< pkg "postgresql-17-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~jammy_amd64.deb" >}} |    {{< pkg "postgresql-14-documentdb" >}}     |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-documentdb" >}}     | {{< pkg "postgresql-17-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~jammy_arm64.deb" >}} |    {{< pkg "postgresql-14-documentdb" >}}     |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-documentdb" >}}     | {{< pkg "postgresql-17-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~noble_amd64.deb" >}} |    {{< pkg "postgresql-14-documentdb" >}}     |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-documentdb" >}}     | {{< pkg "postgresql-17-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-documentdb" "0.106" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~noble_arm64.deb" >}} |    {{< pkg "postgresql-14-documentdb" >}}     |


{{< tabs items="PG17,PG16,PG15" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `documentdb_17` | 0.106 | `el8.aarch64` | pigsty | 2.4 MiB | [documentdb_17-0.106-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-0.106-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_17` | 0.106 | `el8.x86_64` | pigsty | 2.6 MiB | [documentdb_17-0.106-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-0.106-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_17` | 0.105 | `el8.x86_64` | pigsty | 2.5 MiB | [documentdb_17-0.105-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_17-0.105-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_17` | 0.105 | `el8.aarch64` | pigsty | 2.4 MiB | [documentdb_17-0.105-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_17-0.105-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_17` | 0.106 | `el9.aarch64` | pigsty | 2.3 MiB | [documentdb_17-0.106-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-0.106-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_17` | 0.106 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_17-0.106-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-0.106-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_17` | 0.105 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_17-0.105-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_17-0.105-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_17` | 0.105 | `el9.aarch64` | pigsty | 2.4 MiB | [documentdb_17-0.105-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_17-0.105-0PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-documentdb` | 0.106 | `d12.x86_64` | pigsty | 4.3 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-documentdb` | 0.106 | `d12.aarch64` | pigsty | 4.1 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-documentdb` | 0.106 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~jammy_arm64.deb) |
| `postgresql-17-documentdb` | 0.106 | `u22.x86_64` | pigsty | 4.8 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~jammy_amd64.deb) |
| `postgresql-17-documentdb` | 0.106 | `u24.x86_64` | pigsty | 4.4 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~noble_amd64.deb) |
| `postgresql-17-documentdb` | 0.106 | `u24.aarch64` | pigsty | 4.4 MiB | [postgresql-17-documentdb_0.106-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-17-documentdb_0.106-0PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `documentdb_16` | 0.106 | `el8.x86_64` | pigsty | 2.6 MiB | [documentdb_16-0.106-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-0.106-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_16` | 0.106 | `el8.aarch64` | pigsty | 2.5 MiB | [documentdb_16-0.106-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-0.106-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_16` | 0.105 | `el8.x86_64` | pigsty | 2.5 MiB | [documentdb_16-0.105-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_16-0.105-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_16` | 0.105 | `el8.aarch64` | pigsty | 2.4 MiB | [documentdb_16-0.105-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_16-0.105-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_16` | 0.106 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_16-0.106-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-0.106-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_16` | 0.106 | `el9.aarch64` | pigsty | 2.3 MiB | [documentdb_16-0.106-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-0.106-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_16` | 0.105 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_16-0.105-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_16-0.105-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_16` | 0.105 | `el9.aarch64` | pigsty | 2.4 MiB | [documentdb_16-0.105-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_16-0.105-0PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-documentdb` | 0.106 | `d12.x86_64` | pigsty | 4.3 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-documentdb` | 0.106 | `d12.aarch64` | pigsty | 4.1 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-documentdb` | 0.106 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~jammy_arm64.deb) |
| `postgresql-16-documentdb` | 0.106 | `u22.x86_64` | pigsty | 4.8 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~jammy_amd64.deb) |
| `postgresql-16-documentdb` | 0.106 | `u24.aarch64` | pigsty | 4.4 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~noble_arm64.deb) |
| `postgresql-16-documentdb` | 0.106 | `u24.x86_64` | pigsty | 4.4 MiB | [postgresql-16-documentdb_0.106-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-16-documentdb_0.106-0PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `documentdb_15` | 0.106 | `el8.x86_64` | pigsty | 2.6 MiB | [documentdb_15-0.106-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-0.106-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_15` | 0.106 | `el8.aarch64` | pigsty | 2.5 MiB | [documentdb_15-0.106-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-0.106-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_15` | 0.105 | `el8.x86_64` | pigsty | 2.5 MiB | [documentdb_15-0.105-0PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/documentdb_15-0.105-0PIGSTY.el8.x86_64.rpm) |
| `documentdb_15` | 0.105 | `el8.aarch64` | pigsty | 2.4 MiB | [documentdb_15-0.105-0PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/documentdb_15-0.105-0PIGSTY.el8.aarch64.rpm) |
| `documentdb_15` | 0.106 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_15-0.106-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-0.106-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_15` | 0.106 | `el9.aarch64` | pigsty | 2.3 MiB | [documentdb_15-0.106-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-0.106-0PIGSTY.el9.aarch64.rpm) |
| `documentdb_15` | 0.105 | `el9.x86_64` | pigsty | 2.5 MiB | [documentdb_15-0.105-0PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/documentdb_15-0.105-0PIGSTY.el9.x86_64.rpm) |
| `documentdb_15` | 0.105 | `el9.aarch64` | pigsty | 2.4 MiB | [documentdb_15-0.105-0PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/documentdb_15-0.105-0PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-documentdb` | 0.106 | `d12.x86_64` | pigsty | 4.3 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-documentdb` | 0.106 | `d12.aarch64` | pigsty | 4.1 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-documentdb` | 0.106 | `u22.x86_64` | pigsty | 4.8 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~jammy_amd64.deb) |
| `postgresql-15-documentdb` | 0.106 | `u22.aarch64` | pigsty | 4.8 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~jammy_arm64.deb) |
| `postgresql-15-documentdb` | 0.106 | `u24.aarch64` | pigsty | 4.4 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~noble_arm64.deb) |
| `postgresql-15-documentdb` | 0.106 | `u24.x86_64` | pigsty | 4.4 MiB | [postgresql-15-documentdb_0.106-0PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/documentdb/postgresql-15-documentdb_0.106-0PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/microsoft/documentdb" title="Repository" icon="github" subtitle="github.com/microsoft/documentdb" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="documentdb-0.105.0-ferretdb-2.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get documentdb_core; # get documentdb_core source code
pig build dep documentdb_core; # install build dependencies
pig build pkg documentdb_core; # build extension rpm or deb
pig build ext documentdb_core; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install documentdb_core; # install by extension name, for the current active PG version
pig ext install documentdb; # install via package alias, for the active PG version
pig ext install documentdb_core -v 17;   # install for PG 17
pig ext install documentdb_core -v 16;   # install for PG 16
pig ext install documentdb_core -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION documentdb_core;
```

