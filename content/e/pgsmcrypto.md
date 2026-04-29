---
title: "pgsmcrypto"
linkTitle: "pgsmcrypto"
description: "PostgreSQL SM Algorithm Extension"
weight: 7080
categories: ["SEC"]
width: full
---

[**pgsmcrypto**](https://github.com/zhuobie/pgsmcrypto) : PostgreSQL SM Algorithm Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7080** | {{< badge content="pgsmcrypto" link="https://github.com/zhuobie/pgsmcrypto" >}} | {{< ext "pgsmcrypto" >}} | `0.1.1` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} {{< ext "pg_tde" >}} {{< ext "sslutils" >}} {{< ext "faker" >}} {{< ext "uuid-ossp" >}} {{< ext "lo" >}} |

> [!Note] manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgsmcrypto` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "pgsmcrypto_18" "green" >}} {{< bg "17" "pgsmcrypto_17" "green" >}} {{< bg "16" "pgsmcrypto_16" "green" >}} {{< bg "15" "pgsmcrypto_15" "green" >}} {{< bg "14" "pgsmcrypto_14" "green" >}} | `pgsmcrypto_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-pgsmcrypto" "green" >}} {{< bg "17" "postgresql-17-pgsmcrypto" "green" >}} {{< bg "16" "postgresql-16-pgsmcrypto" "green" >}} {{< bg "15" "postgresql-15-pgsmcrypto" "green" >}} {{< bg "14" "postgresql-14-pgsmcrypto" "green" >}} | `postgresql-$v-pgsmcrypto` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgsmcrypto_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-16-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-15-pgsmcrypto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-14-pgsmcrypto : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgsmcrypto : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgsmcrypto : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgsmcrypto : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_18` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 845.9 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_18-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_18` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 667.4 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_18-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_18` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 877.0 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_18-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_18` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 727.2 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_18-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_18` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 884.8 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_18-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_18` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 736.7 KiB | [pgsmcrypto_18-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_18-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgsmcrypto` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 705.9 KiB | [postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgsmcrypto` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 531.7 KiB | [postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgsmcrypto` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 707.7 KiB | [postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgsmcrypto` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 532.1 KiB | [postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgsmcrypto` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 800.4 KiB | [postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgsmcrypto` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 644.1 KiB | [postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgsmcrypto` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 799.4 KiB | [postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgsmcrypto` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 631.2 KiB | [postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-18-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_17` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 845.5 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_17-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_17` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 667.3 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_17-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_17` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 877.2 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_17-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_17` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 727.2 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_17-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_17` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 884.5 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_17-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_17` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 736.6 KiB | [pgsmcrypto_17-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_17-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgsmcrypto` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 706.8 KiB | [postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgsmcrypto` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 531.8 KiB | [postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgsmcrypto` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 707.7 KiB | [postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgsmcrypto` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 532.0 KiB | [postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgsmcrypto` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 801.2 KiB | [postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgsmcrypto` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 642.9 KiB | [postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgsmcrypto` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 797.3 KiB | [postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgsmcrypto` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 631.1 KiB | [postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-17-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_16` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 845.5 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_16-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_16` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 667.3 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_16-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_16` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 877.3 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_16-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_16` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 727.1 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_16-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_16` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 883.1 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_16-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_16` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 735.7 KiB | [pgsmcrypto_16-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_16-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgsmcrypto` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 709.2 KiB | [postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgsmcrypto` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 532.1 KiB | [postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgsmcrypto` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 709.1 KiB | [postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgsmcrypto` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 532.3 KiB | [postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgsmcrypto` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 802.0 KiB | [postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgsmcrypto` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 643.9 KiB | [postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgsmcrypto` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 793.9 KiB | [postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgsmcrypto` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 631.1 KiB | [postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-16-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_15` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 845.8 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_15-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_15` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 667.3 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_15-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_15` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 875.5 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_15-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_15` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 726.7 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_15-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_15` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 880.1 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_15-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_15` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 735.8 KiB | [pgsmcrypto_15-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_15-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgsmcrypto` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 706.2 KiB | [postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgsmcrypto` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 532.2 KiB | [postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgsmcrypto` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 707.5 KiB | [postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgsmcrypto` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 532.1 KiB | [postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgsmcrypto` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 803.1 KiB | [postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgsmcrypto` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 643.2 KiB | [postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgsmcrypto` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 795.9 KiB | [postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgsmcrypto` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 631.2 KiB | [postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-15-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgsmcrypto_14` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 845.7 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsmcrypto_14-0.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgsmcrypto_14` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 667.2 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsmcrypto_14-0.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgsmcrypto_14` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 875.8 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsmcrypto_14-0.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgsmcrypto_14` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 727.2 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsmcrypto_14-0.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgsmcrypto_14` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 884.8 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgsmcrypto_14-0.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgsmcrypto_14` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 736.7 KiB | [pgsmcrypto_14-0.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgsmcrypto_14-0.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgsmcrypto` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 705.8 KiB | [postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgsmcrypto` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 531.5 KiB | [postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgsmcrypto` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 705.7 KiB | [postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgsmcrypto` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 532.1 KiB | [postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgsmcrypto` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 800.1 KiB | [postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgsmcrypto` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 644.4 KiB | [postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgsmcrypto` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 796.9 KiB | [postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgsmcrypto` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 631.2 KiB | [postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsmcrypto/postgresql-14-pgsmcrypto_0.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/zhuobie/pgsmcrypto" title="Repository" icon="github" subtitle="github.com/zhuobie/pgsmcrypto" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsmcrypto-0.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgsmcrypto;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgsmcrypto;		# install via package name, for the active PG version

pig install pgsmcrypto -v 18;   # install for PG 18
pig install pgsmcrypto -v 17;   # install for PG 17
pig install pgsmcrypto -v 16;   # install for PG 16
pig install pgsmcrypto -v 15;   # install for PG 15
pig install pgsmcrypto -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgsmcrypto;
```




## Usage

> [pgsmcrypto: SM national cryptographic algorithm extension for PostgreSQL](https://github.com/zhuobie/pgsmcrypto)

`pgsmcrypto` provides Chinese national cryptographic (SM series) algorithms for PostgreSQL, including SM3 hashing, SM2 asymmetric encryption/signing, and SM4 symmetric encryption.

```sql
CREATE EXTENSION pgsmcrypto;
```

### SM3 Message Digest

```sql
SELECT sm3_hash_string('abc');              -- Returns 64-char hex string (32 bytes)
SELECT sm3_hash('abc'::bytea);              -- Hash bytea input
SELECT sm3_hash(E'\\x616263');              -- Hash raw hex input
```

### SM2 Asymmetric Encryption

#### Key Generation

```sql
SELECT sm2_gen_keypair();                   -- Returns {private_key, public_key} array
SELECT sm2_privkey_valid('f774...');        -- Validate private key (1=valid)
SELECT sm2_pubkey_valid('8093...');         -- Validate public key (1=valid)
SELECT sm2_pk_from_sk('f774...');           -- Derive public key from private key
```

#### Key Export/Import (PEM)

```sql
SELECT sm2_keypair_to_pem_bytes('f774...');       -- Private key to PEM
SELECT sm2_pubkey_to_pem_bytes('8093...');        -- Public key to PEM
SELECT sm2_keypair_from_pem_bytes(pem_bytes);     -- Import from PEM
SELECT sm2_pubkey_from_pem_bytes(pem_bytes);      -- Import public key from PEM
```

#### Sign and Verify

```sql
-- Raw sign/verify (signs message directly)
WITH s AS (
    SELECT sm2_sign_raw('abc'::bytea, 'f774...') AS sig
)
SELECT sm2_verify_raw('abc'::bytea, sig, '8093...') FROM s;

-- Standard sign/verify (SM2 specification with id + SM3 digest)
WITH s AS (
    SELECT sm2_sign('myid'::bytea, 'abc'::bytea, 'f774...') AS sig
)
SELECT sm2_verify('myid'::bytea, 'abc'::bytea, sig, '8093...') FROM s;
```

#### Encrypt and Decrypt

```sql
-- Standard encrypt/decrypt
WITH c AS (
    SELECT sm2_encrypt('abc'::bytea, '8093...') AS enc
)
SELECT sm2_decrypt(enc, 'f774...') FROM c;

-- Also available: sm2_encrypt_c1c2c3, sm2_encrypt_asna1, sm2_encrypt_hex, sm2_encrypt_base64
-- with corresponding decrypt variants
```

### SM4 Symmetric Encryption

```sql
-- ECB mode (key must be 16 bytes)
SELECT sm4_encrypt_ecb('abc'::bytea, '1234567812345678'::bytea);
SELECT sm4_decrypt_ecb(encrypted, '1234567812345678'::bytea);

-- CBC mode (key and IV must be 16 bytes)
SELECT sm4_encrypt_cbc('abc'::bytea, '1234567812345678'::bytea, '0000000000000000'::bytea);
SELECT sm4_decrypt_cbc(encrypted, '1234567812345678'::bytea, '0000000000000000'::bytea);
```
