---
title: "envvar"
linkTitle: "envvar"
description: "Fetch the value of an environment variable"
weight: 4270
categories: ["Util"]
width: full
---

Fetch the value of an environment variable

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4270** | {{< badge content="envvar" link="https://github.com/theory/pg-envvar" >}} | {{< ext "envvar" "envvar" >}} | `1.0.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "session_variable" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/envvar" >}} | `1.0.1` | {{< badge content="18" color="red" alt="pg_envvar_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_envvar_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/envvar" >}} | `1.0.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-envvar" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-envvar` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_envvar_18" >}}     | {{< pkg "pg_envvar_17" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_17-1.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_envvar_16" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_16-1.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_envvar_15" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_15-1.0.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_envvar_14" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_14-1.0.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_envvar_18" >}}     | {{< pkg "pg_envvar_17" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_17-1.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_envvar_16" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_16-1.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_envvar_15" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_15-1.0.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_envvar_14" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_14-1.0.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_envvar_18" >}}     | {{< pkg "pg_envvar_17" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_17-1.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_envvar_16" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_16-1.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_envvar_15" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_15-1.0.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_envvar_14" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_14-1.0.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_envvar_18" >}}     | {{< pkg "pg_envvar_17" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_17-1.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_envvar_16" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_16-1.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_envvar_15" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_15-1.0.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_envvar_14" "1.0.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_14-1.0.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-envvar" >}}     | {{< pkg "postgresql-17-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-envvar" >}}     | {{< pkg "postgresql-17-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-envvar" >}}     | {{< pkg "postgresql-17-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-envvar" >}}     | {{< pkg "postgresql-17-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-envvar" >}}     | {{< pkg "postgresql-17-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-envvar" >}}     | {{< pkg "postgresql-17-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-envvar" "1.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_envvar_17` | 1.0.1 | `el8.aarch64` | pigsty | 12.8 KiB | [pg_envvar_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_17` | 1.0.1 | `el8.x86_64` | pigsty | 12.8 KiB | [pg_envvar_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_17` | 1.0.0 | `el8.aarch64` | pigsty | 12.6 KiB | [pg_envvar_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_17` | 1.0.0 | `el8.x86_64` | pigsty | 12.6 KiB | [pg_envvar_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_17` | 1.0.1 | `el9.x86_64` | pigsty | 12.8 KiB | [pg_envvar_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_envvar_17` | 1.0.1 | `el9.aarch64` | pigsty | 12.7 KiB | [pg_envvar_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_17` | 1.0.0 | `el9.aarch64` | pigsty | 12.4 KiB | [pg_envvar_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_17` | 1.0.0 | `el9.x86_64` | pigsty | 12.6 KiB | [pg_envvar_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-envvar` | 1.0.1 | `d12.aarch64` | pigsty | 9.5 KiB | [postgresql-17-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-envvar` | 1.0.1 | `d12.x86_64` | pigsty | 9.3 KiB | [postgresql-17-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-envvar` | 1.0.1 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-17-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-envvar` | 1.0.1 | `u22.x86_64` | pigsty | 9.0 KiB | [postgresql-17-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-envvar` | 1.0.1 | `u24.x86_64` | pigsty | 9.3 KiB | [postgresql-17-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-envvar` | 1.0.1 | `u24.aarch64` | pigsty | 9.3 KiB | [postgresql-17-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-17-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_envvar_16` | 1.0.1 | `el8.x86_64` | pigsty | 12.8 KiB | [pg_envvar_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_16` | 1.0.1 | `el8.aarch64` | pigsty | 12.8 KiB | [pg_envvar_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_16` | 1.0.0 | `el8.x86_64` | pigsty | 12.6 KiB | [pg_envvar_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_16` | 1.0.0 | `el8.aarch64` | pigsty | 12.6 KiB | [pg_envvar_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_16` | 1.0.1 | `el9.x86_64` | pigsty | 12.8 KiB | [pg_envvar_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_envvar_16` | 1.0.1 | `el9.aarch64` | pigsty | 12.6 KiB | [pg_envvar_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_16` | 1.0.0 | `el9.aarch64` | pigsty | 12.4 KiB | [pg_envvar_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_16` | 1.0.0 | `el9.x86_64` | pigsty | 12.6 KiB | [pg_envvar_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-pg-envvar` | 1.0.1 | `d12.x86_64` | pigsty | 9.3 KiB | [postgresql-16-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-envvar` | 1.0.1 | `d12.aarch64` | pigsty | 9.6 KiB | [postgresql-16-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-envvar` | 1.0.1 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-16-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-envvar` | 1.0.1 | `u22.x86_64` | pigsty | 9.0 KiB | [postgresql-16-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-envvar` | 1.0.1 | `u24.aarch64` | pigsty | 9.3 KiB | [postgresql-16-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-envvar` | 1.0.1 | `u24.x86_64` | pigsty | 9.3 KiB | [postgresql-16-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-16-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_envvar_15` | 1.0.1 | `el8.aarch64` | pigsty | 12.8 KiB | [pg_envvar_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_15` | 1.0.1 | `el8.x86_64` | pigsty | 12.8 KiB | [pg_envvar_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_15` | 1.0.0 | `el8.aarch64` | pigsty | 12.6 KiB | [pg_envvar_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_15` | 1.0.0 | `el8.x86_64` | pigsty | 12.6 KiB | [pg_envvar_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_15` | 1.0.1 | `el9.aarch64` | pigsty | 12.7 KiB | [pg_envvar_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_15` | 1.0.1 | `el9.x86_64` | pigsty | 12.8 KiB | [pg_envvar_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_envvar_15` | 1.0.0 | `el9.aarch64` | pigsty | 12.4 KiB | [pg_envvar_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_15` | 1.0.0 | `el9.x86_64` | pigsty | 12.6 KiB | [pg_envvar_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-15-pg-envvar` | 1.0.1 | `d12.x86_64` | pigsty | 9.3 KiB | [postgresql-15-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-envvar` | 1.0.1 | `d12.aarch64` | pigsty | 9.6 KiB | [postgresql-15-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-envvar` | 1.0.1 | `u22.x86_64` | pigsty | 9.0 KiB | [postgresql-15-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-envvar` | 1.0.1 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-15-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-envvar` | 1.0.1 | `u24.x86_64` | pigsty | 9.3 KiB | [postgresql-15-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-envvar` | 1.0.1 | `u24.aarch64` | pigsty | 9.3 KiB | [postgresql-15-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-15-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_envvar_14` | 1.0.1 | `el8.aarch64` | pigsty | 12.8 KiB | [pg_envvar_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_14` | 1.0.1 | `el8.x86_64` | pigsty | 12.8 KiB | [pg_envvar_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_14` | 1.0.0 | `el8.x86_64` | pigsty | 12.6 KiB | [pg_envvar_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_14` | 1.0.0 | `el8.aarch64` | pigsty | 12.6 KiB | [pg_envvar_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_14` | 1.0.1 | `el9.aarch64` | pigsty | 12.6 KiB | [pg_envvar_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_14` | 1.0.1 | `el9.x86_64` | pigsty | 12.8 KiB | [pg_envvar_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_envvar_14` | 1.0.0 | `el9.aarch64` | pigsty | 12.4 KiB | [pg_envvar_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_14` | 1.0.0 | `el9.x86_64` | pigsty | 12.6 KiB | [pg_envvar_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-pg-envvar` | 1.0.1 | `d12.aarch64` | pigsty | 9.5 KiB | [postgresql-14-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-envvar` | 1.0.1 | `d12.x86_64` | pigsty | 9.3 KiB | [postgresql-14-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-envvar` | 1.0.1 | `u22.x86_64` | pigsty | 9.0 KiB | [postgresql-14-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-envvar` | 1.0.1 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-14-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-envvar` | 1.0.1 | `u24.x86_64` | pigsty | 9.2 KiB | [postgresql-14-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-envvar` | 1.0.1 | `u24.aarch64` | pigsty | 9.3 KiB | [postgresql-14-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-14-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_envvar_13` | 1.0.1 | `el8.aarch64` | pigsty | 12.8 KiB | [pg_envvar_13-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_13-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_13` | 1.0.1 | `el8.x86_64` | pigsty | 12.8 KiB | [pg_envvar_13-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_13-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_13` | 1.0.0 | `el8.aarch64` | pigsty | 12.6 KiB | [pg_envvar_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_envvar_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_envvar_13` | 1.0.0 | `el8.x86_64` | pigsty | 12.6 KiB | [pg_envvar_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_envvar_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_envvar_13` | 1.0.1 | `el9.aarch64` | pigsty | 12.6 KiB | [pg_envvar_13-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_13-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_envvar_13` | 1.0.1 | `el9.x86_64` | pigsty | 12.8 KiB | [pg_envvar_13-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_13-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_envvar_13` | 1.0.0 | `el9.x86_64` | pigsty | 12.6 KiB | [pg_envvar_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_envvar_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_envvar_13` | 1.0.0 | `el9.aarch64` | pigsty | 12.4 KiB | [pg_envvar_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_envvar_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-envvar` | 1.0.1 | `d12.aarch64` | pigsty | 9.5 KiB | [postgresql-13-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-13-pg-envvar_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-envvar` | 1.0.1 | `d12.x86_64` | pigsty | 9.3 KiB | [postgresql-13-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-envvar/postgresql-13-pg-envvar_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-envvar` | 1.0.1 | `u22.x86_64` | pigsty | 9.0 KiB | [postgresql-13-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-13-pg-envvar_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-envvar` | 1.0.1 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-13-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-envvar/postgresql-13-pg-envvar_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-envvar` | 1.0.1 | `u24.x86_64` | pigsty | 9.2 KiB | [postgresql-13-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-13-pg-envvar_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-envvar` | 1.0.1 | `u24.aarch64` | pigsty | 9.3 KiB | [postgresql-13-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-envvar/postgresql-13-pg-envvar_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theory/pg-envvar" title="Repository" icon="github" subtitle="github.com/theory/pg-envvar" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg-envvar-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get envvar; # get envvar source code
pig build dep envvar; # install build dependencies
pig build pkg envvar; # build extension rpm or deb
pig build ext envvar; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install envvar; # install by extension name, for the current active PG version
pig ext install envvar; # install via package alias, for the active PG version
pig ext install envvar -v 18;   # install for PG 18
pig ext install envvar -v 17;   # install for PG 17
pig ext install envvar -v 16;   # install for PG 16
pig ext install envvar -v 15;   # install for PG 15
pig ext install envvar -v 14;   # install for PG 14
pig ext install envvar -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION envvar;
```

