---
title: "pg_vault_tde"
linkTitle: "pg_vault_tde"
description: "Transparent Data Encryption for PostgreSQL through custom table and index access methods"
weight: 7510
categories: ["SEC"]
width: full
---

[**pg_vault_tde**](https://github.com/labmiriade/pg_vault_tde) : Transparent Data Encryption for PostgreSQL through custom table and index access methods


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7510** | {{< badge content="pg_vault_tde" link="https://github.com/labmiriade/pg_vault_tde" >}} | {{< ext "pg_vault_tde" >}} | `1.7.0` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_tde" >}} {{< ext "supabase_vault" >}} {{< ext "pgsodium" >}} {{< ext "column_encrypt" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} |

> [!Note] Requires PostgreSQL 17+, OpenSSL 3, libcurl, and shared_preload_libraries=pg_vault_tde; RPM excludes EL8; includes pg_dump_tde, pg_restore_tde, and pg_basebackup_tde.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_vault_tde` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "pg_vault_tde_18" "green" >}} {{< bg "17" "pg_vault_tde_17" "green" >}} {{< bg "16" "pg_vault_tde_16" "red" >}} {{< bg "15" "pg_vault_tde_15" "red" >}} {{< bg "14" "pg_vault_tde_14" "red" >}} | `pg_vault_tde_$v` | `openssl-libs`, `libcurl` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "postgresql-18-pg-vault-tde" "green" >}} {{< bg "17" "postgresql-17-pg-vault-tde" "green" >}} {{< bg "16" "postgresql-16-pg-vault-tde" "red" >}} {{< bg "15" "postgresql-15-pg-vault-tde" "red" >}} {{< bg "14" "postgresql-14-pg-vault-tde" "red" >}} | `postgresql-$v-pg-vault-tde` | `libssl3 | libssl3t64`, `libcurl4 | libcurl4t64` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_vault_tde_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_vault_tde_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "pg_vault_tde_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_vault_tde_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "pg_vault_tde_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_vault_tde_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "pg_vault_tde_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_vault_tde_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "pg_vault_tde_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "pg_vault_tde_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vault_tde_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 161.6 KiB | [pg_vault_tde_18-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vault_tde_18-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_vault_tde_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 158.0 KiB | [pg_vault_tde_18-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vault_tde_18-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_vault_tde_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 164.5 KiB | [pg_vault_tde_18-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vault_tde_18-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_vault_tde_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 159.2 KiB | [pg_vault_tde_18-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vault_tde_18-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 319.6 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 309.2 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 320.8 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 309.5 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 342.9 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 334.0 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 333.8 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 327.4 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 332.0 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 323.5 KiB | [postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vault_tde_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 162.1 KiB | [pg_vault_tde_17-1.7.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vault_tde_17-1.7.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_vault_tde_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 158.3 KiB | [pg_vault_tde_17-1.7.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vault_tde_17-1.7.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_vault_tde_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 165.1 KiB | [pg_vault_tde_17-1.7.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vault_tde_17-1.7.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_vault_tde_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 159.9 KiB | [pg_vault_tde_17-1.7.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vault_tde_17-1.7.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 321.4 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 311.2 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 323.1 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 312.3 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 384.8 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 376.2 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 335.8 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 329.6 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 334.0 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 325.7 KiB | [postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/labmiriade/pg_vault_tde" title="Repository" icon="github" subtitle="github.com/labmiriade/pg_vault_tde" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_vault_tde-1.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_vault_tde;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_vault_tde;		# install via package name, for the active PG version

pig install pg_vault_tde -v 18;   # install for PG 18
pig install pg_vault_tde -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_vault_tde';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_vault_tde;
```
