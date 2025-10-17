---
title: "pg_jsonschema"
linkTitle: "pg_jsonschema"
description: "PostgreSQL extension providing JSON Schema validation"
weight: 2800
categories: ["Feat"]
width: full
---

PostgreSQL extension providing JSON Schema validation

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2800** | {{< badge content="pg_jsonschema" link="https://github.com/supabase/pg_jsonschema" >}} | {{< ext "pg_jsonschema" "pg_jsonschema" >}} | `0.3.3` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_graphql" >}} {{< ext "jsquery" >}} {{< ext "plv8" >}} {{< ext "jsonb_plperl" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} |

> [!Note] pgrx=0.12.9


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_jsonschema" >}} | `0.3.3` | {{< badge content="18" color="red" alt="pg_jsonschema_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_jsonschema_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_jsonschema" >}} | `0.3.3` | {{< badge content="18" color="red" alt="postgresql-18-pg-jsonschema" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-jsonschema` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_jsonschema_18" >}}     | {{< pkg "pg_jsonschema_17" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_17-0.3.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_jsonschema_16" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_16-0.3.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_jsonschema_15" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_15-0.3.3-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_jsonschema_14" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_14-0.3.3-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_jsonschema_18" >}}     | {{< pkg "pg_jsonschema_17" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_17-0.3.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_jsonschema_16" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_16-0.3.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_jsonschema_15" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_15-0.3.3-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_jsonschema_14" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_14-0.3.3-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_jsonschema_18" >}}     | {{< pkg "pg_jsonschema_17" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_17-0.3.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_jsonschema_16" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_16-0.3.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_jsonschema_15" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_15-0.3.3-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_jsonschema_14" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_14-0.3.3-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_jsonschema_18" >}}     | {{< pkg "pg_jsonschema_17" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_17-0.3.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_jsonschema_16" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_16-0.3.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_jsonschema_15" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_15-0.3.3-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_jsonschema_14" "0.3.3" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_14-0.3.3-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-jsonschema" >}}     | {{< pkg "postgresql-17-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-jsonschema" >}}     | {{< pkg "postgresql-17-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-jsonschema" >}}     | {{< pkg "postgresql-17-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-jsonschema" >}}     | {{< pkg "postgresql-17-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-jsonschema" >}}     | {{< pkg "postgresql-17-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-jsonschema" >}}     | {{< pkg "postgresql-17-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-jsonschema" "0.3.3" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_jsonschema_17` | 0.3.3 | `el8.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_17-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_17-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_jsonschema_17` | 0.3.3 | `el8.aarch64` | pigsty | 970.0 KiB | [pg_jsonschema_17-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_17-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_jsonschema_17` | 0.3.3 | `el9.aarch64` | pigsty | 1.0 MiB | [pg_jsonschema_17-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_17-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_jsonschema_17` | 0.3.3 | `el9.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_17-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_17-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-jsonschema` | 0.3.3 | `d12.x86_64` | pigsty | 923.3 KiB | [postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-jsonschema` | 0.3.3 | `d12.aarch64` | pigsty | 783.8 KiB | [postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-jsonschema` | 0.3.3 | `u22.x86_64` | pigsty | 1.0 MiB | [postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-jsonschema` | 0.3.3 | `u22.aarch64` | pigsty | 937.2 KiB | [postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-jsonschema` | 0.3.3 | `u24.x86_64` | pigsty | 1023.6 KiB | [postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-jsonschema` | 0.3.3 | `u24.aarch64` | pigsty | 926.9 KiB | [postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-17-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_jsonschema_16` | 0.3.3 | `el8.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_16-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_16-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_jsonschema_16` | 0.3.3 | `el8.aarch64` | pigsty | 970.0 KiB | [pg_jsonschema_16-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_16-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_jsonschema_16` | 0.3.3 | `el9.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_16-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_16-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_jsonschema_16` | 0.3.3 | `el9.aarch64` | pigsty | 1.0 MiB | [pg_jsonschema_16-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_16-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-jsonschema` | 0.3.3 | `d12.x86_64` | pigsty | 923.4 KiB | [postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-jsonschema` | 0.3.3 | `d12.aarch64` | pigsty | 783.5 KiB | [postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-jsonschema` | 0.3.3 | `u22.aarch64` | pigsty | 932.7 KiB | [postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-jsonschema` | 0.3.3 | `u22.x86_64` | pigsty | 1.0 MiB | [postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-jsonschema` | 0.3.3 | `u24.x86_64` | pigsty | 1023.3 KiB | [postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-jsonschema` | 0.3.3 | `u24.aarch64` | pigsty | 925.0 KiB | [postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-16-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_jsonschema_15` | 0.3.3 | `el8.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_15-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_15-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_jsonschema_15` | 0.3.3 | `el8.aarch64` | pigsty | 969.8 KiB | [pg_jsonschema_15-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_15-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_jsonschema_15` | 0.3.3 | `el9.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_15-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_15-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_jsonschema_15` | 0.3.3 | `el9.aarch64` | pigsty | 1.0 MiB | [pg_jsonschema_15-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_15-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-jsonschema` | 0.3.3 | `d12.aarch64` | pigsty | 783.6 KiB | [postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-jsonschema` | 0.3.3 | `d12.x86_64` | pigsty | 923.1 KiB | [postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-jsonschema` | 0.3.3 | `u22.aarch64` | pigsty | 936.4 KiB | [postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-jsonschema` | 0.3.3 | `u22.x86_64` | pigsty | 1.0 MiB | [postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-jsonschema` | 0.3.3 | `u24.x86_64` | pigsty | 1023.6 KiB | [postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-jsonschema` | 0.3.3 | `u24.aarch64` | pigsty | 928.3 KiB | [postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-15-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_jsonschema_14` | 0.3.3 | `el8.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_14-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_14-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_jsonschema_14` | 0.3.3 | `el8.aarch64` | pigsty | 969.8 KiB | [pg_jsonschema_14-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_14-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_jsonschema_14` | 0.3.3 | `el9.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_14-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_14-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_jsonschema_14` | 0.3.3 | `el9.aarch64` | pigsty | 1.0 MiB | [pg_jsonschema_14-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_14-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-jsonschema` | 0.3.3 | `d12.x86_64` | pigsty | 923.3 KiB | [postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-jsonschema` | 0.3.3 | `d12.aarch64` | pigsty | 783.8 KiB | [postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-jsonschema` | 0.3.3 | `u22.x86_64` | pigsty | 1.0 MiB | [postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-jsonschema` | 0.3.3 | `u22.aarch64` | pigsty | 937.3 KiB | [postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-jsonschema` | 0.3.3 | `u24.x86_64` | pigsty | 1023.6 KiB | [postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-jsonschema` | 0.3.3 | `u24.aarch64` | pigsty | 924.2 KiB | [postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-14-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_jsonschema_13` | 0.3.3 | `el8.aarch64` | pigsty | 970.0 KiB | [pg_jsonschema_13-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_jsonschema_13-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_jsonschema_13` | 0.3.3 | `el8.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_13-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_jsonschema_13-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_jsonschema_13` | 0.3.3 | `el9.aarch64` | pigsty | 1.0 MiB | [pg_jsonschema_13-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_jsonschema_13-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_jsonschema_13` | 0.3.3 | `el9.x86_64` | pigsty | 1.1 MiB | [pg_jsonschema_13-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_jsonschema_13-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-jsonschema` | 0.3.3 | `d12.aarch64` | pigsty | 783.6 KiB | [postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-jsonschema` | 0.3.3 | `d12.x86_64` | pigsty | 923.7 KiB | [postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-jsonschema/postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-jsonschema` | 0.3.3 | `u22.aarch64` | pigsty | 937.3 KiB | [postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-jsonschema` | 0.3.3 | `u22.x86_64` | pigsty | 1.0 MiB | [postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-jsonschema/postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-jsonschema` | 0.3.3 | `u24.aarch64` | pigsty | 924.3 KiB | [postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-jsonschema` | 0.3.3 | `u24.x86_64` | pigsty | 1023.0 KiB | [postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-jsonschema/postgresql-13-pg-jsonschema_0.3.3-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/pg_jsonschema" title="Repository" icon="github" subtitle="github.com/supabase/pg_jsonschema" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_jsonschema-0.3.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_jsonschema; # get pg_jsonschema source code
pig build dep pg_jsonschema; # install build dependencies
pig build pkg pg_jsonschema; # build extension rpm or deb
pig build ext pg_jsonschema; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_jsonschema; # install by extension name, for the current active PG version
pig ext install pg_jsonschema; # install via package alias, for the active PG version
pig ext install pg_jsonschema -v 17;   # install for PG 17
pig ext install pg_jsonschema -v 16;   # install for PG 16
pig ext install pg_jsonschema -v 15;   # install for PG 15
pig ext install pg_jsonschema -v 14;   # install for PG 14
pig ext install pg_jsonschema -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_jsonschema;
```

