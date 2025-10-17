---
title: "pgcozy"
linkTitle: "pgcozy"
description: "Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL."
weight: 5190
categories: ["Admin"]
width: full
---

Pre-warming shared buffers according to previous pg_buffercache snapshots for PostgreSQL.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5190** | {{< badge content="pgcozy" link="https://github.com/vventirozos/pgcozy" >}} | {{< ext "pgcozy" "pgcozy" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgfincore" >}} {{< ext "pg_cooldown" >}} {{< ext "pg_prewarm" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_visibility" >}} {{< ext "system_stats" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgcozy" >}} | `1.0` | {{< badge content="18" color="red" alt="pgcozy_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgcozy_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgcozy" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-pgcozy" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgcozy` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pgcozy_18" >}}     | {{< pkg "pgcozy_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgcozy_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgcozy_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pgcozy_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pgcozy_18" >}}     | {{< pkg "pgcozy_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgcozy_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgcozy_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pgcozy_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pgcozy_18" >}}     | {{< pkg "pgcozy_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgcozy_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgcozy_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pgcozy_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pgcozy_18" >}}     | {{< pkg "pgcozy_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgcozy_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgcozy_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pgcozy_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgcozy" >}}     | {{< pkg "postgresql-17-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgcozy" >}}     | {{< pkg "postgresql-17-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgcozy" >}}     | {{< pkg "postgresql-17-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgcozy" >}}     | {{< pkg "postgresql-17-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgcozy" >}}     | {{< pkg "postgresql-17-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgcozy" >}}     | {{< pkg "postgresql-17-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pgcozy" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgcozy_17` | 1.0 | `el8.x86_64` | pigsty | 10.7 KiB | [pgcozy_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcozy_17` | 1.0 | `el8.aarch64` | pigsty | 10.7 KiB | [pgcozy_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcozy_17` | 1.0 | `el9.aarch64` | pigsty | 10.7 KiB | [pgcozy_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcozy_17` | 1.0 | `el9.x86_64` | pigsty | 10.7 KiB | [pgcozy_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pgcozy` | 1.0 | `d12.x86_64` | pigsty | 8.2 KiB | [postgresql-17-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgcozy` | 1.0 | `d12.aarch64` | pigsty | 8.2 KiB | [postgresql-17-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgcozy` | 1.0 | `u22.x86_64` | pigsty | 8.2 KiB | [postgresql-17-pgcozy_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgcozy` | 1.0 | `u22.aarch64` | pigsty | 8.2 KiB | [postgresql-17-pgcozy_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgcozy` | 1.0 | `u24.x86_64` | pigsty | 8.2 KiB | [postgresql-17-pgcozy_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgcozy` | 1.0 | `u24.aarch64` | pigsty | 8.2 KiB | [postgresql-17-pgcozy_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-17-pgcozy_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgcozy_16` | 1.0 | `el8.x86_64` | pigsty | 10.7 KiB | [pgcozy_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcozy_16` | 1.0 | `el8.aarch64` | pigsty | 10.7 KiB | [pgcozy_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcozy_16` | 1.0 | `el9.x86_64` | pigsty | 10.7 KiB | [pgcozy_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcozy_16` | 1.0 | `el9.aarch64` | pigsty | 10.7 KiB | [pgcozy_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgcozy` | 1.0 | `d12.x86_64` | pigsty | 8.2 KiB | [postgresql-16-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgcozy` | 1.0 | `d12.aarch64` | pigsty | 8.2 KiB | [postgresql-16-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgcozy` | 1.0 | `u22.aarch64` | pigsty | 8.2 KiB | [postgresql-16-pgcozy_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgcozy` | 1.0 | `u22.x86_64` | pigsty | 8.2 KiB | [postgresql-16-pgcozy_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgcozy` | 1.0 | `u24.x86_64` | pigsty | 8.2 KiB | [postgresql-16-pgcozy_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgcozy` | 1.0 | `u24.aarch64` | pigsty | 8.2 KiB | [postgresql-16-pgcozy_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-16-pgcozy_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgcozy_15` | 1.0 | `el8.x86_64` | pigsty | 10.7 KiB | [pgcozy_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcozy_15` | 1.0 | `el8.aarch64` | pigsty | 10.7 KiB | [pgcozy_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcozy_15` | 1.0 | `el9.x86_64` | pigsty | 10.7 KiB | [pgcozy_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcozy_15` | 1.0 | `el9.aarch64` | pigsty | 10.7 KiB | [pgcozy_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgcozy` | 1.0 | `d12.aarch64` | pigsty | 8.2 KiB | [postgresql-15-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgcozy` | 1.0 | `d12.x86_64` | pigsty | 8.2 KiB | [postgresql-15-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgcozy` | 1.0 | `u22.aarch64` | pigsty | 8.2 KiB | [postgresql-15-pgcozy_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgcozy` | 1.0 | `u22.x86_64` | pigsty | 8.2 KiB | [postgresql-15-pgcozy_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgcozy` | 1.0 | `u24.x86_64` | pigsty | 8.2 KiB | [postgresql-15-pgcozy_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgcozy` | 1.0 | `u24.aarch64` | pigsty | 8.2 KiB | [postgresql-15-pgcozy_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-15-pgcozy_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgcozy_14` | 1.0 | `el8.x86_64` | pigsty | 10.7 KiB | [pgcozy_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcozy_14` | 1.0 | `el8.aarch64` | pigsty | 10.7 KiB | [pgcozy_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcozy_14` | 1.0 | `el9.x86_64` | pigsty | 10.7 KiB | [pgcozy_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcozy_14` | 1.0 | `el9.aarch64` | pigsty | 10.7 KiB | [pgcozy_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgcozy` | 1.0 | `d12.x86_64` | pigsty | 8.2 KiB | [postgresql-14-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgcozy` | 1.0 | `d12.aarch64` | pigsty | 8.2 KiB | [postgresql-14-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgcozy` | 1.0 | `u22.x86_64` | pigsty | 8.2 KiB | [postgresql-14-pgcozy_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgcozy` | 1.0 | `u22.aarch64` | pigsty | 8.2 KiB | [postgresql-14-pgcozy_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgcozy` | 1.0 | `u24.x86_64` | pigsty | 8.2 KiB | [postgresql-14-pgcozy_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgcozy` | 1.0 | `u24.aarch64` | pigsty | 8.2 KiB | [postgresql-14-pgcozy_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-14-pgcozy_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgcozy_13` | 1.0 | `el8.aarch64` | pigsty | 10.7 KiB | [pgcozy_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcozy_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcozy_13` | 1.0 | `el8.x86_64` | pigsty | 10.7 KiB | [pgcozy_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcozy_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcozy_13` | 1.0 | `el9.aarch64` | pigsty | 10.7 KiB | [pgcozy_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcozy_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcozy_13` | 1.0 | `el9.x86_64` | pigsty | 10.7 KiB | [pgcozy_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcozy_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pgcozy` | 1.0 | `d12.aarch64` | pigsty | 8.2 KiB | [postgresql-13-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-13-pgcozy_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgcozy` | 1.0 | `d12.x86_64` | pigsty | 8.2 KiB | [postgresql-13-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcozy/postgresql-13-pgcozy_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgcozy` | 1.0 | `u22.aarch64` | pigsty | 8.2 KiB | [postgresql-13-pgcozy_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-13-pgcozy_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgcozy` | 1.0 | `u22.x86_64` | pigsty | 8.2 KiB | [postgresql-13-pgcozy_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcozy/postgresql-13-pgcozy_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgcozy` | 1.0 | `u24.aarch64` | pigsty | 8.2 KiB | [postgresql-13-pgcozy_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-13-pgcozy_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pgcozy` | 1.0 | `u24.x86_64` | pigsty | 8.2 KiB | [postgresql-13-pgcozy_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcozy/postgresql-13-pgcozy_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/vventirozos/pgcozy" title="Repository" icon="github" subtitle="github.com/vventirozos/pgcozy" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgcozy-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgcozy; # get pgcozy source code
pig build dep pgcozy; # install build dependencies
pig build pkg pgcozy; # build extension rpm or deb
pig build ext pgcozy; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgcozy; # install by extension name, for the current active PG version
pig ext install pgcozy; # install via package alias, for the active PG version
pig ext install pgcozy -v 18;   # install for PG 18
pig ext install pgcozy -v 17;   # install for PG 17
pig ext install pgcozy -v 16;   # install for PG 16
pig ext install pgcozy -v 15;   # install for PG 15
pig ext install pgcozy -v 14;   # install for PG 14
pig ext install pgcozy -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgcozy;
```

