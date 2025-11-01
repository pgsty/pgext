---
title: "pgsmcrypto"
linkTitle: "pgsmcrypto"
description: "PostgreSQL SM Algorithm Extension"
weight: 7070
categories: ["SEC"]
width: full
---

PostgreSQL SM Algorithm Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7070** | {{< badge content="pgsmcrypto" link="https://github.com/zhuobie/pgsmcrypto" >}} | {{< ext "pgsmcrypto" >}} | `0.1.1` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} {{< ext "pg_tde" >}} {{< ext "sslutils" >}} {{< ext "faker" >}} {{< ext "uuid-ossp" >}} {{< ext "lo" >}} |

> [!Note] pgrx=0.16.1, manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgsmcrypto" >}} | `0.1.1` | {{< bg "18" "pgsmcrypto_18" "green" >}} {{< bg "17" "pgsmcrypto_17" "green" >}} {{< bg "16" "pgsmcrypto_16" "green" >}} {{< bg "15" "pgsmcrypto_15" "green" >}} {{< bg "14" "pgsmcrypto_14" "green" >}} {{< bg "13" "pgsmcrypto_13" "green" >}} | `pgsmcrypto_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgsmcrypto" >}} | `0.1.1` | {{< bg "18" "postgresql-18-pgsmcrypto" "green" >}} {{< bg "17" "postgresql-17-pgsmcrypto" "green" >}} {{< bg "16" "postgresql-16-pgsmcrypto" "green" >}} {{< bg "15" "postgresql-15-pgsmcrypto" "green" >}} {{< bg "14" "postgresql-14-pgsmcrypto" "green" >}} {{< bg "13" "postgresql-13-pgsmcrypto" "green" >}} | `postgresql-$v-pgsmcrypto` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-pgsmcrypto : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-pgsmcrypto : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgsmcrypto : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgsmcrypto : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-pgsmcrypto : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-pgsmcrypto : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-pgsmcrypto : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-13-pgsmcrypto : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_18` | 0.1.1 | `el8.x86_64` | pigsty | 845.9 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_18-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_18` | 0.1.1 | `el8.aarch64` | pigsty | 667.4 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_18-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_18` | 0.1.1 | `el9.x86_64` | pigsty | 877.0 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_18-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_18` | 0.1.1 | `el9.aarch64` | pigsty | 727.2 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_18-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_18` | 0.1.1 | `el10.x86_64` | pigsty | 884.8 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_18-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_18` | 0.1.1 | `el10.aarch64` | pigsty | 736.7 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_18-0.1.1-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_17` | 0.1.1 | `el8.x86_64` | pigsty | 845.5 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_17-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_17` | 0.1.1 | `el8.aarch64` | pigsty | 667.3 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_17-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_17` | 0.1.1 | `el9.x86_64` | pigsty | 877.2 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_17-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_17` | 0.1.1 | `el9.aarch64` | pigsty | 727.2 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_17-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_17` | 0.1.1 | `el10.x86_64` | pigsty | 884.5 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_17-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_17` | 0.1.1 | `el10.aarch64` | pigsty | 736.6 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_17-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 639.8 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.5 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 719.0 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 656.2 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 715.9 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 646.8 KiB | [postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_16` | 0.1.1 | `el8.x86_64` | pigsty | 845.5 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_16-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_16` | 0.1.1 | `el8.aarch64` | pigsty | 667.3 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_16-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_16` | 0.1.1 | `el9.x86_64` | pigsty | 877.3 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_16-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_16` | 0.1.1 | `el9.aarch64` | pigsty | 727.1 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_16-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_16` | 0.1.1 | `el10.x86_64` | pigsty | 883.1 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_16-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_16` | 0.1.1 | `el10.aarch64` | pigsty | 735.7 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_16-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 640.3 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.4 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 720.7 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 656.0 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 715.9 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 643.7 KiB | [postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_15` | 0.1.1 | `el8.x86_64` | pigsty | 845.8 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_15-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_15` | 0.1.1 | `el8.aarch64` | pigsty | 667.3 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_15-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_15` | 0.1.1 | `el9.x86_64` | pigsty | 875.5 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_15-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_15` | 0.1.1 | `el9.aarch64` | pigsty | 726.7 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_15-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_15` | 0.1.1 | `el10.x86_64` | pigsty | 880.1 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_15-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_15` | 0.1.1 | `el10.aarch64` | pigsty | 735.8 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_15-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 640.3 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.5 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 721.4 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 656.0 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 716.0 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 644.7 KiB | [postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_14` | 0.1.1 | `el8.x86_64` | pigsty | 845.7 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_14-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_14` | 0.1.1 | `el8.aarch64` | pigsty | 667.2 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_14-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_14` | 0.1.1 | `el9.x86_64` | pigsty | 875.8 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_14-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_14` | 0.1.1 | `el9.aarch64` | pigsty | 727.2 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_14-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_14` | 0.1.1 | `el10.x86_64` | pigsty | 884.8 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_14-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_14` | 0.1.1 | `el10.aarch64` | pigsty | 736.7 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_14-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 640.1 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.2 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 721.4 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 655.9 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 716.2 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 646.7 KiB | [postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_13` | 0.1.1 | `el8.x86_64` | pigsty | 846.0 KiB | [pgsmcrypto_13-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_13-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_13` | 0.1.1 | `el8.aarch64` | pigsty | 667.1 KiB | [pgsmcrypto_13-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_13-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_13` | 0.1.1 | `el9.x86_64` | pigsty | 879.9 KiB | [pgsmcrypto_13-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_13-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_13` | 0.1.1 | `el9.aarch64` | pigsty | 727.2 KiB | [pgsmcrypto_13-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_13-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_13` | 0.1.1 | `el10.x86_64` | pigsty | 885.4 KiB | [pgsmcrypto_13-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_13-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_13` | 0.1.1 | `el10.aarch64` | pigsty | 736.5 KiB | [pgsmcrypto_13-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_13-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `d12.x86_64` | pigsty | 639.9 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `d12.aarch64` | pigsty | 540.2 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `u22.x86_64` | pigsty | 721.3 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `u22.aarch64` | pigsty | 655.9 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `u24.x86_64` | pigsty | 716.0 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgsmcrypto` | 0.1.0 | `u24.aarch64` | pigsty | 646.6 KiB | [postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-13-pgsmcrypto_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/zhuobie/pgsmcrypto" title="Repository" icon="github" subtitle="github.com/zhuobie/pgsmcrypto" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsmcrypto-0.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgsmcrypto; # get pgsmcrypto source code
pig build dep pgsmcrypto; # install build dependencies
pig build pkg pgsmcrypto; # build extension rpm or deb
pig build ext pgsmcrypto; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgsmcrypto; # install by extension name, for the current active PG version
pig ext install pgsmcrypto; # install via package alias, for the active PG version
pig ext install pgsmcrypto -v 18;   # install for PG 18
pig ext install pgsmcrypto -v 17;   # install for PG 17
pig ext install pgsmcrypto -v 16;   # install for PG 16
pig ext install pgsmcrypto -v 15;   # install for PG 15
pig ext install pgsmcrypto -v 14;   # install for PG 14
pig ext install pgsmcrypto -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgsmcrypto;
```

