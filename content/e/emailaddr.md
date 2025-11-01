---
title: "emailaddr"
linkTitle: "emailaddr"
description: "Email address type for PostgreSQL"
weight: 3850
categories: ["TYPE"]
width: full
---

Email address type for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3850** | {{< badge content="emailaddr" link="https://github.com/petere/pgemailaddr" >}} | {{< ext "emailaddr" "pg_emailaddr" >}} | `0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] +varatt.h


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/emailaddr" >}} | `0` | {{< bg "18" "pg_emailaddr_18*" "red" >}} {{< bg "17" "pg_emailaddr_17*" "green" >}} {{< bg "16" "pg_emailaddr_16*" "green" >}} {{< bg "15" "pg_emailaddr_15*" "green" >}} {{< bg "14" "pg_emailaddr_14*" "green" >}} {{< bg "13" "pg_emailaddr_13*" "green" >}} | `pg_emailaddr_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/emailaddr" >}} | `0` | {{< bg "18" "postgresql-18-pg-emailaddr" "red" >}} {{< bg "17" "postgresql-17-pg-emailaddr" "green" >}} {{< bg "16" "postgresql-16-pg-emailaddr" "green" >}} {{< bg "15" "postgresql-15-pg-emailaddr" "green" >}} {{< bg "14" "postgresql-14-pg-emailaddr" "green" >}} {{< bg "13" "postgresql-13-pg-emailaddr" "green" >}} | `postgresql-$v-pg-emailaddr` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_emailaddr_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_emailaddr_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_emailaddr_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_emailaddr_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 0" "pg_emailaddr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 0" "pg_emailaddr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "pg_emailaddr_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-emailaddr : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-13-pg-emailaddr : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-emailaddr : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-13-pg-emailaddr : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-emailaddr : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-emailaddr : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-emailaddr : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-emailaddr : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-13-pg-emailaddr : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-emailaddr : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-13-pg-emailaddr : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-emailaddr : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-13-pg-emailaddr : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-emailaddr : MISS 0" "red" >}}      | {{< bg "PIGSTY 0" "postgresql-17-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-16-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-15-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-14-pg-emailaddr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0" "postgresql-13-pg-emailaddr : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_18` | 0 | `el10.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_18-0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_18-0-1PIGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_18` | 0 | `el10.aarch64` | pigsty | 13.5 KiB | [pg_emailaddr_18-0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_18-0-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_17` | 0 | `el8.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_17-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_17-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_17` | 0 | `el8.aarch64` | pigsty | 13.7 KiB | [pg_emailaddr_17-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_17-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_17` | 0 | `el9.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_17-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_17-0-1PIGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_17` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_17-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_17-0-1PIGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_17` | 0 | `el10.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_17-0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_17-0-1PIGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_17` | 0 | `el10.aarch64` | pigsty | 13.6 KiB | [pg_emailaddr_17-0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_17-0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.5 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-17-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-17-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_16` | 0 | `el8.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_16-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_16-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_16` | 0 | `el8.aarch64` | pigsty | 13.7 KiB | [pg_emailaddr_16-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_16-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_16` | 0 | `el9.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_16-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_16-0-1PIGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_16` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_16-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_16-0-1PIGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_16` | 0 | `el10.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_16-0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_16-0-1PIGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_16` | 0 | `el10.aarch64` | pigsty | 13.6 KiB | [pg_emailaddr_16-0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_16-0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.5 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-16-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-16-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_15` | 0 | `el8.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_15-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_15-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_15` | 0 | `el8.aarch64` | pigsty | 13.7 KiB | [pg_emailaddr_15-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_15-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_15` | 0 | `el9.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_15-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_15-0-1PIGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_15` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_15-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_15-0-1PIGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_15` | 0 | `el10.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_15-0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_15-0-1PIGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_15` | 0 | `el10.aarch64` | pigsty | 13.6 KiB | [pg_emailaddr_15-0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_15-0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.5 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.0 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.8 KiB | [postgresql-15-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-15-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_14` | 0 | `el8.x86_64` | pigsty | 13.8 KiB | [pg_emailaddr_14-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_14-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_14` | 0 | `el8.aarch64` | pigsty | 13.7 KiB | [pg_emailaddr_14-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_14-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_14` | 0 | `el9.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_14-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_14-0-1PIGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_14` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_14-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_14-0-1PIGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_14` | 0 | `el10.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_14-0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_14-0-1PIGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_14` | 0 | `el10.aarch64` | pigsty | 13.6 KiB | [pg_emailaddr_14-0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_14-0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.5 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.7 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.7 KiB | [postgresql-14-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-14-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_emailaddr_13` | 0 | `el8.x86_64` | pigsty | 13.7 KiB | [pg_emailaddr_13-0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_emailaddr_13-0-1PIGSTY.el8.x86_64.rpm) |
| `pg_emailaddr_13` | 0 | `el8.aarch64` | pigsty | 13.7 KiB | [pg_emailaddr_13-0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_emailaddr_13-0-1PIGSTY.el8.aarch64.rpm) |
| `pg_emailaddr_13` | 0 | `el9.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_13-0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_emailaddr_13-0-1PIGSTY.el9.x86_64.rpm) |
| `pg_emailaddr_13` | 0 | `el9.aarch64` | pigsty | 13.4 KiB | [pg_emailaddr_13-0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_emailaddr_13-0-1PIGSTY.el9.aarch64.rpm) |
| `pg_emailaddr_13` | 0 | `el10.x86_64` | pigsty | 13.6 KiB | [pg_emailaddr_13-0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_emailaddr_13-0-1PIGSTY.el10.x86_64.rpm) |
| `pg_emailaddr_13` | 0 | `el10.aarch64` | pigsty | 13.6 KiB | [pg_emailaddr_13-0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_emailaddr_13-0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-emailaddr` | 0 | `d12.x86_64` | pigsty | 12.6 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `d12.aarch64` | pigsty | 12.5 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `u22.x86_64` | pigsty | 13.0 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `u22.aarch64` | pigsty | 12.6 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-emailaddr` | 0 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-13-pg-emailaddr_0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-emailaddr/postgresql-13-pg-emailaddr_0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/pgemailaddr" title="Repository" icon="github" subtitle="github.com/petere/pgemailaddr" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgemailaddr-0.tar.gz" >}}
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
pig ext install pg_emailaddr; # install via package alias, for the active PG version
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

