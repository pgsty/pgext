---
title: "pg_base58"
linkTitle: "pg_base58"
description: "Base58 Encoder/Decoder Extension for PostgreSQL"
weight: 4830
categories: ["FUNC"]
width: full
---

[**pg_base58**](https://github.com/Fell-x27/pg_base58) : Base58 Encoder/Decoder Extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4830** | {{< badge content="pg_base58" link="https://github.com/Fell-x27/pg_base58" >}} | {{< ext "pg_base58" >}} | `0.0.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "url_encode" >}} {{< ext "pg_cardano" >}} {{< ext "base36" >}} {{< ext "base62" >}} {{< ext "pg_polyline" >}} {{< ext "uri" >}} {{< ext "pg_curl" >}} {{< ext "pg_rewrite" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_base58` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_base58_18" "green" >}} {{< bg "17" "pg_base58_17" "green" >}} {{< bg "16" "pg_base58_16" "green" >}} {{< bg "15" "pg_base58_15" "green" >}} {{< bg "14" "pg_base58_14" "green" >}} | `pg_base58_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-base58" "green" >}} {{< bg "17" "postgresql-17-pg-base58" "green" >}} {{< bg "16" "postgresql-16-pg-base58" "green" >}} {{< bg "15" "postgresql-15-pg-base58" "green" >}} {{< bg "14" "postgresql-14-pg-base58" "green" >}} | `postgresql-$v-pg-base58` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_base58_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-base58 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-base58 : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 850.3 KiB | [pg_base58_18-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_18-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_base58_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 763.5 KiB | [pg_base58_18-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_18-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_base58_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 860.5 KiB | [pg_base58_18-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_18-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_base58_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 809.2 KiB | [pg_base58_18-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_18-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_base58_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 859.9 KiB | [pg_base58_18-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base58_18-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_base58_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 786.5 KiB | [pg_base58_18-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base58_18-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-base58` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 681.1 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 567.9 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 680.2 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 568.2 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 758.2 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 672.3 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 748.4 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 661.1 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 744.7 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-base58` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 660.3 KiB | [postgresql-18-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-18-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 847.7 KiB | [pg_base58_17-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_17-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_base58_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 760.5 KiB | [pg_base58_17-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_17-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_base58_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 857.5 KiB | [pg_base58_17-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_17-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_base58_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 806.6 KiB | [pg_base58_17-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_17-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_base58_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 857.7 KiB | [pg_base58_17-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base58_17-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_base58_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 786.2 KiB | [pg_base58_17-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base58_17-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-base58` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 679.3 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 566.5 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 678.8 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 566.9 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 755.7 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 669.2 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 747.7 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 659.6 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 743.5 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-base58` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 658.0 KiB | [postgresql-17-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-17-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 846.0 KiB | [pg_base58_16-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_16-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_base58_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 758.4 KiB | [pg_base58_16-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_16-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_base58_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 855.3 KiB | [pg_base58_16-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_16-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_base58_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 804.9 KiB | [pg_base58_16-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_16-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_base58_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 854.8 KiB | [pg_base58_16-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base58_16-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_base58_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 785.3 KiB | [pg_base58_16-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base58_16-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-base58` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 678.2 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 566.2 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 678.6 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 566.4 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 753.5 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 668.5 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 745.0 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 658.6 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 742.8 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-base58` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 656.9 KiB | [postgresql-16-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-16-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 836.7 KiB | [pg_base58_15-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_15-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_base58_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 749.6 KiB | [pg_base58_15-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_15-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_base58_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 846.8 KiB | [pg_base58_15-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_15-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_base58_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 792.7 KiB | [pg_base58_15-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_15-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_base58_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 847.2 KiB | [pg_base58_15-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base58_15-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_base58_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 783.3 KiB | [pg_base58_15-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base58_15-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-base58` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 671.9 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 561.4 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 672.1 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 561.1 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 749.3 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 661.9 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 740.0 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 652.2 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 735.7 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-base58` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 650.9 KiB | [postgresql-15-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-15-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_base58_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 833.8 KiB | [pg_base58_14-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_base58_14-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_base58_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 747.5 KiB | [pg_base58_14-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_base58_14-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_base58_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 843.8 KiB | [pg_base58_14-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_base58_14-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_base58_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 792.1 KiB | [pg_base58_14-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_base58_14-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_base58_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 844.1 KiB | [pg_base58_14-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_base58_14-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_base58_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 782.2 KiB | [pg_base58_14-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_base58_14-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-base58` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 669.3 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 559.6 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 670.8 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 560.2 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 747.3 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 660.5 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 737.6 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 650.6 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 734.0 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-base58` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 649.2 KiB | [postgresql-14-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-base58/postgresql-14-pg-base58_0.0.1-5PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Fell-x27/pg_base58" title="Repository" icon="github" subtitle="github.com/Fell-x27/pg_base58" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_base58-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_base58;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_base58;		# install via package name, for the active PG version

pig install pg_base58 -v 18;   # install for PG 18
pig install pg_base58 -v 17;   # install for PG 17
pig install pg_base58 -v 16;   # install for PG 16
pig install pg_base58 -v 15;   # install for PG 15
pig install pg_base58 -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_base58;
```




## Usage

> [pg_base58: base58 encoding/decoding for PostgreSQL](https://github.com/Fell-x27/pg_base58)

Provides functions to encode and decode data using Base58 encoding.

```sql
CREATE EXTENSION pg_base58;
```

### Functions

| Function | Description |
|---|---|
| `base58_encode(bytea)` | Encode bytea data to Base58 text |
| `base58_decode(text)` | Decode Base58 text back to bytea |

### Examples

```sql
-- Encode a string to Base58
SELECT base58_encode('hello'::bytea);
-- 'Cn8eVZg'

-- Decode a Base58 string back
SELECT base58_decode('Cn8eVZg');
-- '\x68656c6c6f'  (i.e., 'hello')

-- Round-trip
SELECT convert_from(base58_decode(base58_encode('hello'::bytea)), 'UTF8');
-- 'hello'
```
