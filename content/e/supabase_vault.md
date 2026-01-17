---
title: "supabase_vault"
linkTitle: "supabase_vault"
description: "Supabase Vault Extension"
weight: 7030
categories: ["SEC"]
width: full
---

[**pg_vault**](https://github.com/supabase/vault) : Supabase Vault Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7030** | {{< badge content="supabase_vault" link="https://github.com/supabase/vault" >}} | {{< ext "supabase_vault" "pg_vault" >}} | `0.3.1` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `vault` |
|   **Requires**    | {{< ext "pgsodium" >}} |
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "supautils" >}} {{< ext "pg_session_jwt" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgaudit" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_vault` | `pgsodium` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.1` | {{< bg "18" "vault_18" "green" >}} {{< bg "17" "vault_17" "green" >}} {{< bg "16" "vault_16" "green" >}} {{< bg "15" "vault_15" "green" >}} {{< bg "14" "vault_14" "green" >}} {{< bg "13" "vault_13" "green" >}} | `vault_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.1` | {{< bg "18" "postgresql-18-vault" "green" >}} {{< bg "17" "postgresql-17-vault" "green" >}} {{< bg "16" "postgresql-16-vault" "green" >}} {{< bg "15" "postgresql-15-vault" "green" >}} {{< bg "14" "postgresql-14-vault" "green" >}} {{< bg "13" "postgresql-13-vault" "green" >}} | `postgresql-$v-vault` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-13-vault : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-13-vault : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-13-vault : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-13-vault : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-13-vault : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-13-vault : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-13-vault : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-13-vault : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vault_18` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.2 KiB | [vault_18-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vault_18-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `vault_18` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.1 KiB | [vault_18-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vault_18-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `vault_18` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.5 KiB | [vault_18-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vault_18-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `vault_18` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [vault_18-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vault_18-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `vault_18` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [vault_18-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vault_18-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `vault_18` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [vault_18-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vault_18-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-vault` | `0.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.8 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-vault` | `0.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-vault` | `0.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.8 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-vault` | `0.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.7 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-vault` | `0.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.4 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-vault` | `0.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.2 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-vault` | `0.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.8 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-vault` | `0.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.6 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vault_17` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.2 KiB | [vault_17-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vault_17-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `vault_17` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.1 KiB | [vault_17-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vault_17-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `vault_17` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.5 KiB | [vault_17-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vault_17-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `vault_17` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [vault_17-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vault_17-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `vault_17` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [vault_17-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vault_17-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `vault_17` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [vault_17-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vault_17-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-vault` | `0.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.7 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-vault` | `0.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-vault` | `0.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.8 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-vault` | `0.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.7 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-vault` | `0.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.5 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-vault` | `0.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.3 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-vault` | `0.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.7 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-vault` | `0.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.6 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vault_16` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.2 KiB | [vault_16-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vault_16-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `vault_16` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.1 KiB | [vault_16-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vault_16-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `vault_16` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.5 KiB | [vault_16-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vault_16-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `vault_16` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [vault_16-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vault_16-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `vault_16` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [vault_16-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vault_16-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `vault_16` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [vault_16-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vault_16-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-vault` | `0.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.7 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-vault` | `0.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-vault` | `0.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.8 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-vault` | `0.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.7 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-vault` | `0.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.5 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-vault` | `0.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.2 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-vault` | `0.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.7 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-vault` | `0.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.6 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vault_15` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.4 KiB | [vault_15-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vault_15-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `vault_15` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.1 KiB | [vault_15-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vault_15-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `vault_15` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.9 KiB | [vault_15-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vault_15-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `vault_15` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.7 KiB | [vault_15-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vault_15-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `vault_15` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [vault_15-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vault_15-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `vault_15` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.8 KiB | [vault_15-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vault_15-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-vault` | `0.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.8 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-vault` | `0.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.8 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-vault` | `0.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.9 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-vault` | `0.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.8 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-vault` | `0.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.8 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-vault` | `0.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.6 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-vault` | `0.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.0 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-vault` | `0.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.0 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vault_14` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.4 KiB | [vault_14-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vault_14-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `vault_14` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.1 KiB | [vault_14-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vault_14-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `vault_14` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.9 KiB | [vault_14-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vault_14-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `vault_14` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.7 KiB | [vault_14-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vault_14-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `vault_14` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.8 KiB | [vault_14-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vault_14-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `vault_14` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.8 KiB | [vault_14-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vault_14-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-vault` | `0.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.7 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-vault` | `0.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-vault` | `0.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.8 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-vault` | `0.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.7 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-vault` | `0.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.7 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-vault` | `0.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.6 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-vault` | `0.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.9 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-vault` | `0.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.0 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vault_13` | `0.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.2 KiB | [vault_13-0.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vault_13-0.3.1-1PIGSTY.el8.x86_64.rpm) |
| `vault_13` | `0.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.1 KiB | [vault_13-0.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vault_13-0.3.1-1PIGSTY.el8.aarch64.rpm) |
| `vault_13` | `0.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.9 KiB | [vault_13-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vault_13-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `vault_13` | `0.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.6 KiB | [vault_13-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vault_13-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `vault_13` | `0.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.7 KiB | [vault_13-0.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vault_13-0.3.1-1PIGSTY.el10.x86_64.rpm) |
| `vault_13` | `0.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.7 KiB | [vault_13-0.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vault_13-0.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-vault` | `0.3.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.5 KiB | [postgresql-13-vault_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-13-vault_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-vault` | `0.3.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.6 KiB | [postgresql-13-vault_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vault/postgresql-13-vault_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-vault` | `0.3.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.6 KiB | [postgresql-13-vault_0.3.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-13-vault_0.3.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-vault` | `0.3.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.6 KiB | [postgresql-13-vault_0.3.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vault/postgresql-13-vault_0.3.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-vault` | `0.3.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.4 KiB | [postgresql-13-vault_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-13-vault_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-vault` | `0.3.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.3 KiB | [postgresql-13-vault_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vault/postgresql-13-vault_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-vault` | `0.3.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.7 KiB | [postgresql-13-vault_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-13-vault_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-vault` | `0.3.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.8 KiB | [postgresql-13-vault_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vault/postgresql-13-vault_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/vault" title="Repository" icon="github" subtitle="github.com/supabase/vault" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="vault-0.3.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_vault;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_vault;		# install via package name, for the active PG version
pig install supabase_vault;		# install by extension name, for the current active PG version

pig install supabase_vault -v 18;   # install for PG 18
pig install supabase_vault -v 17;   # install for PG 17
pig install supabase_vault -v 16;   # install for PG 16
pig install supabase_vault -v 15;   # install for PG 15
pig install supabase_vault -v 14;   # install for PG 14
pig install supabase_vault -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION supabase_vault CASCADE; -- requires pgsodium
```
