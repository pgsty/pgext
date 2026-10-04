---
title: "pg_vault_tde"
linkTitle: "pg_vault_tde"
description: "Transparent Data Encryption for PostgreSQL through custom table and index access methods"
weight: 7510
categories: ["SEC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_vault_tde**](https://github.com/labmiriade/pg_vault_tde) : Transparent Data Encryption for PostgreSQL through custom table and index access methods


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7510** | {{< badge content="pg_vault_tde" link="https://github.com/labmiriade/pg_vault_tde" >}} | {{< ext "pg_vault_tde" >}} | `1.7.2` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_tde" >}} {{< ext "supabase_vault" >}} {{< ext "pgsodium" >}} {{< ext "column_encrypt" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} |

> [!Note] Requires PostgreSQL 17+, OpenSSL 3, libcurl, and shared_preload_libraries=pg_vault_tde; RPM excludes EL8; includes pg_dump_tde, pg_restore_tde, and pg_basebackup_tde. Package 1.7.2; SQL control version 1.7. Export existing 1.7.0 encrypted TOAST data with the old binary before upgrading. Rewrite existing encrypted tables after upgrade; custom WAL users must coordinate primary/standby upgrades for resource-manager ID 128 to 161.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_vault_tde` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.2` | {{< bg "18" "pg_vault_tde_18" "green" >}} {{< bg "17" "pg_vault_tde_17" "green" >}} {{< bg "16" "pg_vault_tde_16" "red" >}} {{< bg "15" "pg_vault_tde_15" "red" >}} {{< bg "14" "pg_vault_tde_14" "red" >}} | `pg_vault_tde_$v` | `openssl-libs`, `libcurl` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.2` | {{< bg "18" "postgresql-18-pg-vault-tde" "green" >}} {{< bg "17" "postgresql-17-pg-vault-tde" "green" >}} {{< bg "16" "postgresql-16-pg-vault-tde" "red" >}} {{< bg "15" "postgresql-15-pg-vault-tde" "red" >}} {{< bg "14" "postgresql-14-pg-vault-tde" "red" >}} | `postgresql-$v-pg-vault-tde` | `libssl3 | libssl3t64`, `libcurl4 | libcurl4t64` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_vault_tde_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_vault_tde_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7.2" "pg_vault_tde_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "pg_vault_tde_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7.2" "pg_vault_tde_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "pg_vault_tde_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7.2" "pg_vault_tde_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "pg_vault_tde_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7.2" "pg_vault_tde_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "pg_vault_tde_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_vault_tde_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vault_tde_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-18-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.2" "postgresql-17-pg-vault-tde : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-vault-tde : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-vault-tde : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vault_tde_18` | `1.7.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 424.8 KiB | [pg_vault_tde_18-1.7.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vault_tde_18-1.7.2-1PGSTY.el9.x86_64.rpm) |
| `pg_vault_tde_18` | `1.7.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 416.3 KiB | [pg_vault_tde_18-1.7.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vault_tde_18-1.7.2-1PGSTY.el9.aarch64.rpm) |
| `pg_vault_tde_18` | `1.7.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 427.1 KiB | [pg_vault_tde_18-1.7.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vault_tde_18-1.7.2-1PGSTY.el10.x86_64.rpm) |
| `pg_vault_tde_18` | `1.7.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 418.2 KiB | [pg_vault_tde_18-1.7.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vault_tde_18-1.7.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 425.2 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 411.5 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 426.0 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 412.0 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 446.2 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 435.2 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 434.9 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 426.9 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 432.0 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-vault-tde` | `1.7.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 423.5 KiB | [postgresql-18-pg-vault-tde_1.7.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vault-tde/postgresql-18-pg-vault-tde_1.7.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vault_tde_17` | `1.7.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 426.3 KiB | [pg_vault_tde_17-1.7.2-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vault_tde_17-1.7.2-1PGSTY.el9.x86_64.rpm) |
| `pg_vault_tde_17` | `1.7.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 417.9 KiB | [pg_vault_tde_17-1.7.2-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vault_tde_17-1.7.2-1PGSTY.el9.aarch64.rpm) |
| `pg_vault_tde_17` | `1.7.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 428.5 KiB | [pg_vault_tde_17-1.7.2-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vault_tde_17-1.7.2-1PGSTY.el10.x86_64.rpm) |
| `pg_vault_tde_17` | `1.7.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 420.0 KiB | [pg_vault_tde_17-1.7.2-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vault_tde_17-1.7.2-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 426.8 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 413.0 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 427.8 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 413.4 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 491.9 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 483.2 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 436.1 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 428.6 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 433.6 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-vault-tde` | `1.7.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 424.8 KiB | [postgresql-17-pg-vault-tde_1.7.2-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vault-tde/postgresql-17-pg-vault-tde_1.7.2-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/labmiriade/pg_vault_tde" title="Repository" icon="github" subtitle="github.com/labmiriade/pg_vault_tde" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_vault_tde-1.7.2.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_vault_tde;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

## Usage

Sources:

- [pg_vault_tde.control](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/pg_vault_tde.control)
- [README.md](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/README.md)
- [doc/pg_vault_tde.md](https://github.com/labmiriade/pg_vault_tde/blob/2520a0d80fc3c4d905f57e47e2d871ed7e6b0e7d/doc/pg_vault_tde.md)

`pg_vault_tde` distribution 1.7.2 encrypts table values with AES-256-GCM through `encrypted_heap`, using Vault/OpenBao, a local PKCS#12 wallet or PKCS#11. SQL/control version remains 1.7. PostgreSQL 17–18, OpenSSL 3 and libcurl are required. Configure the key provider and authentication, preload the library, and restart before creating it as a superuser.

### Core Workflow

```ini
shared_preload_libraries = 'pg_vault_tde'
```

```sql
CREATE EXTENSION pg_vault_tde;
SELECT * FROM pg_vault_tde_health_check();
CREATE TABLE customer_secrets (id bigint, secret text) USING encrypted_heap;
CREATE INDEX customer_secrets_id_idx ON customer_secrets USING tde_btree (id);
```

### Operational Boundaries

This example assumes a configured provider. `pg_vault_tde_health_check`, `pg_vault_tde_verify_integrity`, `pg_vault_tde_get_rotation_status` and encrypted-size helpers expose operational state. `tde_btree` supports encrypted equality lookup, not ordering/range or index-only scans. Numeric and nondeterministic-collation encrypted indexes are unsupported; audit old unique/exclusion constraints because their checks may be ineffective. Plain indexes can expose keys, and pg_dump/COPY produce decrypted output. Protect external wallets/KMS and use the matching encrypted-backup tools where supported. Tuple headers, statistics, logs and query results are outside this protection.

Before replacing an older library or restarting, follow the upstream recovery checklist: earlier rotated tables may depend on keys held only in memory; concurrent rotation may require copying data while still readable, and rotated out-of-line values need the specified rewrite. Version 1.7.0 TOAST upgrades also have a separate export requirement. Check rotated/partial indexes and wallet ownership. After upgrading to 1.7.2, every pre-existing encrypted table needs VACUUM FULL to migrate its tuple layout; reserve a maintenance window and extra disk space. When `pg_vault_tde.toast_custom_rmgr` is enabled, the WAL ID changes from 128 to 161: clean shutdown, coordinated primary/standby upgrade, slot drainage and a new base backup are required. There is no rolling upgrade across those WAL formats.

Run key-management operations one at a time. Table rotation keeps reads available but blocks writes for one transaction; logical slots must decode the rotation before a restart or another rotation. Version 1.7.2 tightens caller privileges, but upstream documents remaining SECURITY DEFINER wrapper limitations; follow the listed grants/revokes. Do not toggle `pg_vault_tde.enabled` on a populated encrypted table. These catalog updates do not perform an operational upgrade or change the Pigsty package baseline.
