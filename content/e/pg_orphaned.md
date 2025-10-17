---
title: "pg_orphaned"
linkTitle: "pg_orphaned"
description: "Deal with orphaned files"
weight: 5200
categories: ["Admin"]
width: full
---

Deal with orphaned files

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5200** | {{< badge content="pg_orphaned" link="https://github.com/bdrouvot/pg_orphaned" >}} | {{< ext "pg_orphaned" "pg_orphaned" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_dirtyread" >}} {{< ext "pg_surgery" >}} {{< ext "amcheck" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pg_checksums" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_repack" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_orphaned" >}} | `1.0` | {{< badge content="18" color="red" alt="pg_orphaned_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_orphaned_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_orphaned" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-orphaned" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-orphaned` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_orphaned_18" >}}     | {{< pkg "pg_orphaned_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_orphaned_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_orphaned_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_orphaned_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_orphaned_18" >}}     | {{< pkg "pg_orphaned_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_orphaned_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_orphaned_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_orphaned_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_orphaned_18" >}}     | {{< pkg "pg_orphaned_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_orphaned_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_orphaned_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_orphaned_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_orphaned_18" >}}     | {{< pkg "pg_orphaned_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_orphaned_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_orphaned_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_orphaned_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-orphaned" >}}     | {{< pkg "postgresql-17-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-orphaned" >}}     | {{< pkg "postgresql-17-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-orphaned" >}}     | {{< pkg "postgresql-17-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-orphaned" >}}     | {{< pkg "postgresql-17-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-orphaned" >}}     | {{< pkg "postgresql-17-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-orphaned" >}}     | {{< pkg "postgresql-17-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-orphaned" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_orphaned_17` | 1.0 | `el8.x86_64` | pigsty | 20.7 KiB | [pg_orphaned_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_17` | 1.0 | `el8.aarch64` | pigsty | 20.5 KiB | [pg_orphaned_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_17` | 1.0 | `el9.aarch64` | pigsty | 21.0 KiB | [pg_orphaned_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_17` | 1.0 | `el9.x86_64` | pigsty | 20.9 KiB | [pg_orphaned_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 33.4 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 33.4 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.8 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.9 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.5 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.5 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_orphaned_16` | 1.0 | `el8.x86_64` | pigsty | 20.7 KiB | [pg_orphaned_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_16` | 1.0 | `el8.aarch64` | pigsty | 20.5 KiB | [pg_orphaned_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_16` | 1.0 | `el9.x86_64` | pigsty | 20.9 KiB | [pg_orphaned_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_16` | 1.0 | `el9.aarch64` | pigsty | 21.0 KiB | [pg_orphaned_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 33.0 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 33.1 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.5 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.4 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.5 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.6 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_orphaned_15` | 1.0 | `el8.x86_64` | pigsty | 20.8 KiB | [pg_orphaned_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_15` | 1.0 | `el8.aarch64` | pigsty | 20.6 KiB | [pg_orphaned_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_15` | 1.0 | `el9.x86_64` | pigsty | 21.1 KiB | [pg_orphaned_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_15` | 1.0 | `el9.aarch64` | pigsty | 21.1 KiB | [pg_orphaned_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 33.1 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 33.1 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.6 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.5 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.6 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.6 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_orphaned_14` | 1.0 | `el8.x86_64` | pigsty | 20.8 KiB | [pg_orphaned_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_14` | 1.0 | `el8.aarch64` | pigsty | 20.6 KiB | [pg_orphaned_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_14` | 1.0 | `el9.x86_64` | pigsty | 21.0 KiB | [pg_orphaned_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_14` | 1.0 | `el9.aarch64` | pigsty | 21.1 KiB | [pg_orphaned_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 33.0 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 32.9 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.4 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.5 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.6 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.6 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_orphaned_13` | 1.0 | `el8.aarch64` | pigsty | 20.5 KiB | [pg_orphaned_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_13` | 1.0 | `el8.x86_64` | pigsty | 20.5 KiB | [pg_orphaned_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_13` | 1.0 | `el9.aarch64` | pigsty | 21.0 KiB | [pg_orphaned_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_13` | 1.0 | `el9.x86_64` | pigsty | 21.0 KiB | [pg_orphaned_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-orphaned` | 1.0 | `d12.aarch64` | pigsty | 32.8 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `d12.x86_64` | pigsty | 32.9 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `u22.aarch64` | pigsty | 34.2 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `u22.x86_64` | pigsty | 34.4 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `u24.aarch64` | pigsty | 30.5 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-orphaned` | 1.0 | `u24.x86_64` | pigsty | 30.3 KiB | [postgresql-13-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-13-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bdrouvot/pg_orphaned" title="Repository" icon="github" subtitle="github.com/bdrouvot/pg_orphaned" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_orphaned-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_orphaned; # get pg_orphaned source code
pig build dep pg_orphaned; # install build dependencies
pig build pkg pg_orphaned; # build extension rpm or deb
pig build ext pg_orphaned; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_orphaned; # install by extension name, for the current active PG version
pig ext install pg_orphaned; # install via package alias, for the active PG version
pig ext install pg_orphaned -v 18;   # install for PG 18
pig ext install pg_orphaned -v 17;   # install for PG 17
pig ext install pg_orphaned -v 16;   # install for PG 16
pig ext install pg_orphaned -v 15;   # install for PG 15
pig ext install pg_orphaned -v 14;   # install for PG 14
pig ext install pg_orphaned -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_orphaned;
```

