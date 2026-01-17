---
title: "pg_enigma"
linkTitle: "pg_enigma"
description: "Encrypted postgres data type"
weight: 7070
categories: ["SEC"]
width: full
---

[**pg_enigma**](https://github.com/SoftwareLibreMx/pg_enigma) : Encrypted postgres data type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7070** | {{< badge content="pg_enigma" link="https://github.com/SoftwareLibreMx/pg_enigma" >}} | {{< ext "pg_enigma" >}} | `0.5.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} {{< ext "pg_tde" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_enigma` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "pg_enigma_18" "green" >}} {{< bg "17" "pg_enigma_17" "green" >}} {{< bg "16" "pg_enigma_16" "green" >}} {{< bg "15" "pg_enigma_15" "green" >}} {{< bg "14" "pg_enigma_14" "green" >}} {{< bg "13" "pg_enigma_13" "red" >}} | `pg_enigma_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "postgresql-18-enigma" "green" >}} {{< bg "17" "postgresql-17-enigma" "green" >}} {{< bg "16" "postgresql-16-enigma" "green" >}} {{< bg "15" "postgresql-15-enigma" "green" >}} {{< bg "14" "postgresql-14-enigma" "green" >}} {{< bg "13" "postgresql-13-enigma" "red" >}} | `postgresql-$v-enigma` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_enigma_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_enigma_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_enigma_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_enigma_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_enigma_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pg_enigma_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_enigma_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-enigma : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-enigma : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-enigma : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-enigma : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-enigma : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-enigma : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-enigma : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-enigma : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-enigma : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-enigma : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-enigma : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-enigma : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-enigma : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-enigma : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-enigma : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-enigma : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-enigma : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_enigma_18` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.6 MiB | [pg_enigma_18-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_enigma_18-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_enigma_18` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [pg_enigma_18-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_enigma_18-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_enigma_18` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_enigma_18-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_enigma_18-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_enigma_18` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_enigma_18-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_enigma_18-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_enigma_18` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pg_enigma_18-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_enigma_18-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_enigma_18` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_enigma_18-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_enigma_18-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-enigma` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.1 KiB | [postgresql-18-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-18-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-enigma` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 KiB | [postgresql-18-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-18-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-enigma` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.1 KiB | [postgresql-18-enigma_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-18-enigma_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-enigma` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 KiB | [postgresql-18-enigma_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-18-enigma_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-enigma` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 KiB | [postgresql-18-enigma_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-18-enigma_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-enigma` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 KiB | [postgresql-18-enigma_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-18-enigma_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-enigma` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 KiB | [postgresql-18-enigma_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-18-enigma_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-enigma` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 KiB | [postgresql-18-enigma_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-18-enigma_0.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_enigma_17` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.6 MiB | [pg_enigma_17-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_enigma_17-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_enigma_17` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [pg_enigma_17-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_enigma_17-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_enigma_17` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_enigma_17-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_enigma_17-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_enigma_17` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_enigma_17-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_enigma_17-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_enigma_17` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pg_enigma_17-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_enigma_17-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_enigma_17` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_enigma_17-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_enigma_17-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-enigma` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.1 KiB | [postgresql-17-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-17-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-enigma` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 KiB | [postgresql-17-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-17-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-enigma` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.1 KiB | [postgresql-17-enigma_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-17-enigma_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-enigma` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 KiB | [postgresql-17-enigma_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-17-enigma_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-enigma` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 KiB | [postgresql-17-enigma_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-17-enigma_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-enigma` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 KiB | [postgresql-17-enigma_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-17-enigma_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-enigma` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 KiB | [postgresql-17-enigma_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-17-enigma_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-enigma` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 KiB | [postgresql-17-enigma_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-17-enigma_0.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_enigma_16` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.6 MiB | [pg_enigma_16-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_enigma_16-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_enigma_16` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [pg_enigma_16-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_enigma_16-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_enigma_16` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_enigma_16-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_enigma_16-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_enigma_16` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_enigma_16-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_enigma_16-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_enigma_16` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pg_enigma_16-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_enigma_16-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_enigma_16` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_enigma_16-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_enigma_16-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-enigma` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.1 KiB | [postgresql-16-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-16-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-enigma` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 KiB | [postgresql-16-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-16-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-enigma` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.1 KiB | [postgresql-16-enigma_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-16-enigma_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-enigma` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 KiB | [postgresql-16-enigma_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-16-enigma_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-enigma` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 KiB | [postgresql-16-enigma_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-16-enigma_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-enigma` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 KiB | [postgresql-16-enigma_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-16-enigma_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-enigma` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 KiB | [postgresql-16-enigma_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-16-enigma_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-enigma` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 KiB | [postgresql-16-enigma_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-16-enigma_0.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_enigma_15` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.6 MiB | [pg_enigma_15-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_enigma_15-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_enigma_15` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [pg_enigma_15-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_enigma_15-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_enigma_15` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_enigma_15-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_enigma_15-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_enigma_15` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_enigma_15-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_enigma_15-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_enigma_15` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pg_enigma_15-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_enigma_15-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_enigma_15` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_enigma_15-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_enigma_15-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-enigma` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.1 KiB | [postgresql-15-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-15-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-enigma` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 KiB | [postgresql-15-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-15-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-enigma` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.1 KiB | [postgresql-15-enigma_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-15-enigma_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-enigma` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 KiB | [postgresql-15-enigma_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-15-enigma_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-enigma` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 KiB | [postgresql-15-enigma_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-15-enigma_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-enigma` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 KiB | [postgresql-15-enigma_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-15-enigma_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-enigma` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 KiB | [postgresql-15-enigma_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-15-enigma_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-enigma` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 KiB | [postgresql-15-enigma_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-15-enigma_0.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_enigma_14` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.6 MiB | [pg_enigma_14-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_enigma_14-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_enigma_14` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [pg_enigma_14-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_enigma_14-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_enigma_14` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [pg_enigma_14-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_enigma_14-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_enigma_14` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.4 MiB | [pg_enigma_14-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_enigma_14-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_enigma_14` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pg_enigma_14-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_enigma_14-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_enigma_14` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pg_enigma_14-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_enigma_14-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-enigma` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.1 KiB | [postgresql-14-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-14-enigma_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-enigma` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 KiB | [postgresql-14-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-enigma/postgresql-14-enigma_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-enigma` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.1 KiB | [postgresql-14-enigma_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-14-enigma_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-enigma` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 KiB | [postgresql-14-enigma_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-enigma/postgresql-14-enigma_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-enigma` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 KiB | [postgresql-14-enigma_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-14-enigma_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-enigma` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 KiB | [postgresql-14-enigma_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-enigma/postgresql-14-enigma_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-enigma` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 KiB | [postgresql-14-enigma_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-14-enigma_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-enigma` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 KiB | [postgresql-14-enigma_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-enigma/postgresql-14-enigma_0.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/SoftwareLibreMx/pg_enigma" title="Repository" icon="github" subtitle="github.com/SoftwareLibreMx/pg_enigma" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_enigma-0.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_enigma;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_enigma;		# install via package name, for the active PG version

pig install pg_enigma -v 18;   # install for PG 18
pig install pg_enigma -v 17;   # install for PG 17
pig install pg_enigma -v 16;   # install for PG 16
pig install pg_enigma -v 15;   # install for PG 15
pig install pg_enigma -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_enigma;
```
