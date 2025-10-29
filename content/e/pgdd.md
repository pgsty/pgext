---
title: "pgdd"
linkTitle: "pgdd"
description: "Introspect pg data dictionary via standard SQL"
weight: 5130
categories: ["ADMIN"]
width: full
---

Introspect pg data dictionary via standard SQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5130** | {{< badge content="pgdd" link="https://github.com/rustprooflabs/pgdd" >}} | {{< ext "pgdd" >}} | `0.6.0` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "ddlx" >}} {{< ext "pg_catcheck" >}} {{< ext "pageinspect" >}} {{< ext "pg_permissions" >}} {{< ext "adminpack" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} |

> [!Note] pgrx=0.16.1, manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgdd" >}} | `0.6.0` | {{< bg "18" "pgdd_18" "green" >}} {{< bg "17" "pgdd_17" "green" >}} {{< bg "16" "pgdd_16" "green" >}} {{< bg "15" "pgdd_15" "green" >}} {{< bg "14" "pgdd_14" "green" >}} {{< bg "13" "pgdd_13" "green" >}} | `pgdd_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgdd" >}} | `0.6.0` | {{< bg "18" "postgresql-18-pgdd" "green" >}} {{< bg "17" "postgresql-17-pgdd" "green" >}} {{< bg "16" "postgresql-16-pgdd" "green" >}} {{< bg "15" "postgresql-15-pgdd" "green" >}} {{< bg "14" "postgresql-14-pgdd" "green" >}} {{< bg "13" "postgresql-13-pgdd" "green" >}} | `postgresql-$v-pgdd` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pgdd_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pgdd_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pgdd_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pgdd_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pgdd_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pgdd_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pgdd_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgdd : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgdd : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgdd : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgdd : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgdd : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgdd : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgdd : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgdd : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgdd : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.0" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "postgresql-13-pgdd : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_17` | 0.6.0 | `el8.x86_64` | pigsty | 195.8 KiB | [pgdd_17-0.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_17-0.6.0-1PIGSTY.el8.x86_64.rpm) |
| `pgdd_17` | 0.6.0 | `el8.aarch64` | pigsty | 185.7 KiB | [pgdd_17-0.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_17-0.6.0-1PIGSTY.el8.aarch64.rpm) |
| `pgdd_17` | 0.6.0 | `el9.x86_64` | pigsty | 200.4 KiB | [pgdd_17-0.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_17-0.6.0-1PIGSTY.el9.x86_64.rpm) |
| `pgdd_17` | 0.6.0 | `el9.aarch64` | pigsty | 198.2 KiB | [pgdd_17-0.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_17-0.6.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pgdd` | 0.6.0 | `d12.x86_64` | pigsty | 158.9 KiB | [postgresql-17-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-17-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgdd` | 0.6.0 | `d12.aarch64` | pigsty | 144.4 KiB | [postgresql-17-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-17-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgdd` | 0.6.0 | `u22.x86_64` | pigsty | 173.5 KiB | [postgresql-17-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-17-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgdd` | 0.6.0 | `u22.aarch64` | pigsty | 167.2 KiB | [postgresql-17-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-17-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgdd` | 0.6.0 | `u24.x86_64` | pigsty | 172.1 KiB | [postgresql-17-pgdd_0.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-17-pgdd_0.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgdd` | 0.6.0 | `u24.aarch64` | pigsty | 165.8 KiB | [postgresql-17-pgdd_0.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-17-pgdd_0.6.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_16` | 0.6.0 | `el8.x86_64` | pigsty | 195.8 KiB | [pgdd_16-0.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_16-0.6.0-1PIGSTY.el8.x86_64.rpm) |
| `pgdd_16` | 0.6.0 | `el8.aarch64` | pigsty | 185.7 KiB | [pgdd_16-0.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_16-0.6.0-1PIGSTY.el8.aarch64.rpm) |
| `pgdd_16` | 0.6.0 | `el9.x86_64` | pigsty | 200.3 KiB | [pgdd_16-0.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_16-0.6.0-1PIGSTY.el9.x86_64.rpm) |
| `pgdd_16` | 0.6.0 | `el9.aarch64` | pigsty | 198.2 KiB | [pgdd_16-0.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_16-0.6.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgdd` | 0.6.0 | `d12.x86_64` | pigsty | 158.9 KiB | [postgresql-16-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-16-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgdd` | 0.6.0 | `d12.aarch64` | pigsty | 144.3 KiB | [postgresql-16-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-16-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgdd` | 0.6.0 | `u22.x86_64` | pigsty | 173.5 KiB | [postgresql-16-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-16-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgdd` | 0.6.0 | `u22.aarch64` | pigsty | 167.0 KiB | [postgresql-16-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-16-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgdd` | 0.6.0 | `u24.x86_64` | pigsty | 172.1 KiB | [postgresql-16-pgdd_0.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-16-pgdd_0.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgdd` | 0.6.0 | `u24.aarch64` | pigsty | 165.8 KiB | [postgresql-16-pgdd_0.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-16-pgdd_0.6.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_15` | 0.6.0 | `el8.x86_64` | pigsty | 195.7 KiB | [pgdd_15-0.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_15-0.6.0-1PIGSTY.el8.x86_64.rpm) |
| `pgdd_15` | 0.6.0 | `el8.aarch64` | pigsty | 185.7 KiB | [pgdd_15-0.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_15-0.6.0-1PIGSTY.el8.aarch64.rpm) |
| `pgdd_15` | 0.6.0 | `el9.x86_64` | pigsty | 200.3 KiB | [pgdd_15-0.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_15-0.6.0-1PIGSTY.el9.x86_64.rpm) |
| `pgdd_15` | 0.6.0 | `el9.aarch64` | pigsty | 198.1 KiB | [pgdd_15-0.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_15-0.6.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgdd` | 0.6.0 | `d12.x86_64` | pigsty | 158.9 KiB | [postgresql-15-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-15-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgdd` | 0.6.0 | `d12.aarch64` | pigsty | 144.3 KiB | [postgresql-15-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-15-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgdd` | 0.6.0 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-15-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-15-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgdd` | 0.6.0 | `u22.aarch64` | pigsty | 167.0 KiB | [postgresql-15-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-15-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgdd` | 0.6.0 | `u24.x86_64` | pigsty | 172.1 KiB | [postgresql-15-pgdd_0.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-15-pgdd_0.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgdd` | 0.6.0 | `u24.aarch64` | pigsty | 166.0 KiB | [postgresql-15-pgdd_0.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-15-pgdd_0.6.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_14` | 0.6.0 | `el8.x86_64` | pigsty | 195.8 KiB | [pgdd_14-0.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_14-0.6.0-1PIGSTY.el8.x86_64.rpm) |
| `pgdd_14` | 0.6.0 | `el8.aarch64` | pigsty | 185.7 KiB | [pgdd_14-0.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_14-0.6.0-1PIGSTY.el8.aarch64.rpm) |
| `pgdd_14` | 0.6.0 | `el9.x86_64` | pigsty | 199.9 KiB | [pgdd_14-0.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_14-0.6.0-1PIGSTY.el9.x86_64.rpm) |
| `pgdd_14` | 0.6.0 | `el9.aarch64` | pigsty | 198.1 KiB | [pgdd_14-0.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_14-0.6.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgdd` | 0.6.0 | `d12.x86_64` | pigsty | 159.0 KiB | [postgresql-14-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-14-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgdd` | 0.6.0 | `d12.aarch64` | pigsty | 144.3 KiB | [postgresql-14-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-14-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgdd` | 0.6.0 | `u22.x86_64` | pigsty | 173.6 KiB | [postgresql-14-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-14-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgdd` | 0.6.0 | `u22.aarch64` | pigsty | 166.9 KiB | [postgresql-14-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-14-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgdd` | 0.6.0 | `u24.x86_64` | pigsty | 172.6 KiB | [postgresql-14-pgdd_0.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-14-pgdd_0.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgdd` | 0.6.0 | `u24.aarch64` | pigsty | 166.1 KiB | [postgresql-14-pgdd_0.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-14-pgdd_0.6.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_13` | 0.6.0 | `el8.x86_64` | pigsty | 195.8 KiB | [pgdd_13-0.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_13-0.6.0-1PIGSTY.el8.x86_64.rpm) |
| `pgdd_13` | 0.6.0 | `el8.aarch64` | pigsty | 185.7 KiB | [pgdd_13-0.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_13-0.6.0-1PIGSTY.el8.aarch64.rpm) |
| `pgdd_13` | 0.6.0 | `el9.x86_64` | pigsty | 200.0 KiB | [pgdd_13-0.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_13-0.6.0-1PIGSTY.el9.x86_64.rpm) |
| `pgdd_13` | 0.6.0 | `el9.aarch64` | pigsty | 198.2 KiB | [pgdd_13-0.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_13-0.6.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pgdd` | 0.6.0 | `d12.x86_64` | pigsty | 158.9 KiB | [postgresql-13-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-13-pgdd_0.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgdd` | 0.6.0 | `d12.aarch64` | pigsty | 144.3 KiB | [postgresql-13-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-13-pgdd_0.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgdd` | 0.6.0 | `u22.x86_64` | pigsty | 173.8 KiB | [postgresql-13-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-13-pgdd_0.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgdd` | 0.6.0 | `u22.aarch64` | pigsty | 167.0 KiB | [postgresql-13-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-13-pgdd_0.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgdd` | 0.6.0 | `u24.x86_64` | pigsty | 172.7 KiB | [postgresql-13-pgdd_0.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-13-pgdd_0.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgdd` | 0.6.0 | `u24.aarch64` | pigsty | 166.1 KiB | [postgresql-13-pgdd_0.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-13-pgdd_0.6.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rustprooflabs/pgdd" title="Repository" icon="github" subtitle="github.com/rustprooflabs/pgdd" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgdd-0.6.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgdd; # get pgdd source code
pig build dep pgdd; # install build dependencies
pig build pkg pgdd; # build extension rpm or deb
pig build ext pgdd; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgdd; # install by extension name, for the current active PG version
pig ext install pgdd; # install via package alias, for the active PG version
pig ext install pgdd -v 18;   # install for PG 18
pig ext install pgdd -v 17;   # install for PG 17
pig ext install pgdd -v 16;   # install for PG 16
pig ext install pgdd -v 15;   # install for PG 15
pig ext install pgdd -v 14;   # install for PG 14
pig ext install pgdd -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgdd CASCADE SCHEMA dd;
```

