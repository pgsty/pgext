---
title: "column_encrypt"
linkTitle: "column_encrypt"
description: "Transparent column-level encryption with encrypted_text and encrypted_bytea types"
weight: 7030
categories: ["SEC"]
width: full
---

[**column_encrypt**](https://github.com/vibhorkum/column_encrypt) : Transparent column-level encryption with encrypted_text and encrypted_bytea types


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7030** | {{< badge content="column_encrypt" link="https://github.com/vibhorkum/column_encrypt" >}} | {{< ext "column_encrypt" >}} | `4.0` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `encrypt` |
|   **Requires**    | {{< ext "pgcrypto" >}} |
|   **See Also**    | {{< ext "pg_enigma" >}} {{< ext "pgsodium" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} {{< ext "pg_tde" >}} {{< ext "pgsmcrypto" >}} {{< ext "sslutils" >}} |

> [!Note] fixed encrypt schema; create schema encrypt before CREATE EXTENSION; preload column_encrypt;


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `column_encrypt` | `pgcrypto` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0` | {{< bg "18" "column_encrypt_18" "green" >}} {{< bg "17" "column_encrypt_17" "green" >}} {{< bg "16" "column_encrypt_16" "green" >}} {{< bg "15" "column_encrypt_15" "green" >}} {{< bg "14" "column_encrypt_14" "green" >}} | `column_encrypt_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.0` | {{< bg "18" "postgresql-18-column-encrypt" "green" >}} {{< bg "17" "postgresql-17-column-encrypt" "green" >}} {{< bg "16" "postgresql-16-column-encrypt" "green" >}} {{< bg "15" "postgresql-15-column-encrypt" "green" >}} {{< bg "14" "postgresql-14-column-encrypt" "green" >}} | `postgresql-$v-column-encrypt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "column_encrypt_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 4.0" "postgresql-18-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-17-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-16-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-15-column-encrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.0" "postgresql-14-column-encrypt : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `column_encrypt_18` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 55.2 KiB | [column_encrypt_18-4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/column_encrypt_18-4.0-1PIGSTY.el8.x86_64.rpm) |
| `column_encrypt_18` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.9 KiB | [column_encrypt_18-4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/column_encrypt_18-4.0-1PIGSTY.el8.aarch64.rpm) |
| `column_encrypt_18` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 51.4 KiB | [column_encrypt_18-4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/column_encrypt_18-4.0-1PIGSTY.el9.x86_64.rpm) |
| `column_encrypt_18` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 51.0 KiB | [column_encrypt_18-4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/column_encrypt_18-4.0-1PIGSTY.el9.aarch64.rpm) |
| `column_encrypt_18` | `4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 51.6 KiB | [column_encrypt_18-4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/column_encrypt_18-4.0-1PIGSTY.el10.x86_64.rpm) |
| `column_encrypt_18` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 51.3 KiB | [column_encrypt_18-4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/column_encrypt_18-4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-column-encrypt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 61.9 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.2 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 61.9 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.2 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 63.4 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 62.9 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.5 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 61.5 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 62.7 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-column-encrypt` | `4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 62.1 KiB | [postgresql-18-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-18-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `column_encrypt_17` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 55.2 KiB | [column_encrypt_17-4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/column_encrypt_17-4.0-1PIGSTY.el8.x86_64.rpm) |
| `column_encrypt_17` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.9 KiB | [column_encrypt_17-4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/column_encrypt_17-4.0-1PIGSTY.el8.aarch64.rpm) |
| `column_encrypt_17` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 51.5 KiB | [column_encrypt_17-4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/column_encrypt_17-4.0-1PIGSTY.el9.x86_64.rpm) |
| `column_encrypt_17` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 51.0 KiB | [column_encrypt_17-4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/column_encrypt_17-4.0-1PIGSTY.el9.aarch64.rpm) |
| `column_encrypt_17` | `4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 51.6 KiB | [column_encrypt_17-4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/column_encrypt_17-4.0-1PIGSTY.el10.x86_64.rpm) |
| `column_encrypt_17` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 51.3 KiB | [column_encrypt_17-4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/column_encrypt_17-4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-column-encrypt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 61.9 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.1 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 61.9 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.2 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.8 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 64.2 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.6 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 61.7 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 62.7 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-column-encrypt` | `4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 62.1 KiB | [postgresql-17-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-17-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `column_encrypt_16` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 55.2 KiB | [column_encrypt_16-4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/column_encrypt_16-4.0-1PIGSTY.el8.x86_64.rpm) |
| `column_encrypt_16` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.9 KiB | [column_encrypt_16-4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/column_encrypt_16-4.0-1PIGSTY.el8.aarch64.rpm) |
| `column_encrypt_16` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 51.5 KiB | [column_encrypt_16-4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/column_encrypt_16-4.0-1PIGSTY.el9.x86_64.rpm) |
| `column_encrypt_16` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 51.0 KiB | [column_encrypt_16-4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/column_encrypt_16-4.0-1PIGSTY.el9.aarch64.rpm) |
| `column_encrypt_16` | `4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 51.6 KiB | [column_encrypt_16-4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/column_encrypt_16-4.0-1PIGSTY.el10.x86_64.rpm) |
| `column_encrypt_16` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 51.4 KiB | [column_encrypt_16-4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/column_encrypt_16-4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-column-encrypt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 61.9 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.1 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 62.0 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.2 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.8 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 64.2 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.5 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 61.5 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 62.7 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-column-encrypt` | `4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 62.1 KiB | [postgresql-16-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-16-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `column_encrypt_15` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 55.5 KiB | [column_encrypt_15-4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/column_encrypt_15-4.0-1PIGSTY.el8.x86_64.rpm) |
| `column_encrypt_15` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [column_encrypt_15-4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/column_encrypt_15-4.0-1PIGSTY.el8.aarch64.rpm) |
| `column_encrypt_15` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 51.9 KiB | [column_encrypt_15-4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/column_encrypt_15-4.0-1PIGSTY.el9.x86_64.rpm) |
| `column_encrypt_15` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 51.5 KiB | [column_encrypt_15-4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/column_encrypt_15-4.0-1PIGSTY.el9.aarch64.rpm) |
| `column_encrypt_15` | `4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 52.1 KiB | [column_encrypt_15-4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/column_encrypt_15-4.0-1PIGSTY.el10.x86_64.rpm) |
| `column_encrypt_15` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 52.0 KiB | [column_encrypt_15-4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/column_encrypt_15-4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-column-encrypt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 62.0 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.5 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 62.1 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.6 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 65.1 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 64.5 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.6 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.2 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 62.8 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-column-encrypt` | `4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 62.4 KiB | [postgresql-15-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-15-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `column_encrypt_14` | `4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 55.5 KiB | [column_encrypt_14-4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/column_encrypt_14-4.0-1PIGSTY.el8.x86_64.rpm) |
| `column_encrypt_14` | `4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 55.1 KiB | [column_encrypt_14-4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/column_encrypt_14-4.0-1PIGSTY.el8.aarch64.rpm) |
| `column_encrypt_14` | `4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 51.9 KiB | [column_encrypt_14-4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/column_encrypt_14-4.0-1PIGSTY.el9.x86_64.rpm) |
| `column_encrypt_14` | `4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 51.5 KiB | [column_encrypt_14-4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/column_encrypt_14-4.0-1PIGSTY.el9.aarch64.rpm) |
| `column_encrypt_14` | `4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 52.1 KiB | [column_encrypt_14-4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/column_encrypt_14-4.0-1PIGSTY.el10.x86_64.rpm) |
| `column_encrypt_14` | `4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 51.9 KiB | [column_encrypt_14-4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/column_encrypt_14-4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-column-encrypt` | `4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 61.9 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 61.5 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 62.0 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 61.6 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 65.1 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 64.5 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.4 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.2 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 62.8 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-column-encrypt` | `4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 62.4 KiB | [postgresql-14-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/column-encrypt/postgresql-14-column-encrypt_4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/vibhorkum/column_encrypt" title="Repository" icon="github" subtitle="github.com/vibhorkum/column_encrypt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="column_encrypt-4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg column_encrypt;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install column_encrypt;		# install via package name, for the active PG version

pig install column_encrypt -v 18;   # install for PG 18
pig install column_encrypt -v 17;   # install for PG 17
pig install column_encrypt -v 16;   # install for PG 16
pig install column_encrypt -v 15;   # install for PG 15
pig install column_encrypt -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'column_encrypt';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION column_encrypt CASCADE; -- requires pgcrypto
```
