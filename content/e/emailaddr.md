---
title: "emailaddr"
linkTitle: "emailaddr"
description: "Email address type for PostgreSQL"
weight: 3850
categories: ["Type"]
width: full
---

Email address type for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3850** | {{< badge content="emailaddr" link="https://github.com/petere/pgemailaddr" >}} | {{< ext "emailaddr" "pgemailaddr" >}} | `0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] +varatt.h


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/emailaddr" >}} | `0` | {{< badge content="18" color="red" alt="pg_emailaddr_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_emailaddr_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/emailaddr" >}} | `0` | {{< badge content="18" color="red" alt="postgresql-18-pg-emailaddr" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-emailaddr` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_emailaddr_18" >}}     | {{< pkg "pg_emailaddr_17" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_17-0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_emailaddr_16" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_16-0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_emailaddr_15" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_15-0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_emailaddr_14" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_14-0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_emailaddr_18" >}}     | {{< pkg "pg_emailaddr_17" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_17-0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_emailaddr_16" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_16-0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_emailaddr_15" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_15-0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_emailaddr_14" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_14-0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_emailaddr_18" >}}     | {{< pkg "pg_emailaddr_17" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_17-0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_emailaddr_16" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_16-0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_emailaddr_15" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_15-0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_emailaddr_14" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_14-0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_emailaddr_18" >}}     | {{< pkg "pg_emailaddr_17" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_17-0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_emailaddr_16" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_16-0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_emailaddr_15" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_15-0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_emailaddr_14" "0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_14-0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-emailaddr" >}}     | {{< pkg "postgresql-17-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-emailaddr" >}}     | {{< pkg "postgresql-17-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-emailaddr" >}}     | {{< pkg "postgresql-17-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-emailaddr" >}}     | {{< pkg "postgresql-17-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-emailaddr" >}}     | {{< pkg "postgresql-17-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-emailaddr" >}}     | {{< pkg "postgresql-17-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-emailaddr" "0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_emailaddr_17` | 0 | `el8.x86_64` | pigsty | 13.7 KiB | [pg_emailaddr_17-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_17-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_17` | 0 | `el8.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_17-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_17-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_17` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_17-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_17-0-1PIGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_17` | 0 | `el9.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_17-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_17-0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.5 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_emailaddr_16` | 0 | `el8.x86_64` | pigsty | 13.7 KiB | [pg_emailaddr_16-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_16-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_16` | 0 | `el8.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_16-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_16-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_16` | 0 | `el9.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_16-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_16-0-1PIGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_16` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_16-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_16-0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.5 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_emailaddr_15` | 0 | `el8.x86_64` | pigsty | 13.7 KiB | [pg_emailaddr_15-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_15-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_15` | 0 | `el8.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_15-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_15-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_15` | 0 | `el9.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_15-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_15-0-1PIGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_15` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_15-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_15-0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.5 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.0 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_emailaddr_14` | 0 | `el8.x86_64` | pigsty | 13.7 KiB | [pg_emailaddr_14-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_14-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_14` | 0 | `el8.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_14-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_14-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_14` | 0 | `el9.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_14-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_14-0-1PIGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_14` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_14-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_14-0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.5 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.7 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.7 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_emailaddr_13` | 0 | `el8.aarch64` | pigsty | 13.5 KiB | [pg_emailaddr_13-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_13-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_13` | 0 | `el8.x86_64` | pigsty | 13.7 KiB | [pg_emailaddr_13-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_13-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_13` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_13-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_13-0-1PIGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_13` | 0 | `el9.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_13-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_13-0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.5 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.6 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.6 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.0 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/pgemailaddr" title="Repository" icon="github" subtitle="github.com/petere/pgemailaddr" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgemailaddr-0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get emailaddr; # get emailaddr source code
pig build dep emailaddr; # install build dependencies
pig build pkg emailaddr; # build extension rpm or deb
pig build ext emailaddr; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install emailaddr; # install by extension name, for the current active PG version
pig ext install pgemailaddr; # install via package alias, for the active PG version
pig ext install emailaddr -v 17;   # install for PG 17
pig ext install emailaddr -v 16;   # install for PG 16
pig ext install emailaddr -v 15;   # install for PG 15
pig ext install emailaddr -v 14;   # install for PG 14
pig ext install emailaddr -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION emailaddr;
```

