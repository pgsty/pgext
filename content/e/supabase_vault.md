---
title: "supabase_vault"
linkTitle: "supabase_vault"
description: "Supabase Vault Extension"
weight: 7050
categories: ["SEC"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_vault**](https://github.com/supabase/vault) : Supabase Vault Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7050** | {{< badge content="supabase_vault" link="https://github.com/supabase/vault" >}} | {{< ext "supabase_vault" "pg_vault" >}} | `0.3.1` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `vault` |
|   **See Also**    | {{< ext "column_encrypt" >}} {{< ext "pg_tde" >}} {{< ext "pg_enigma" >}} {{< ext "pgsodium" >}} {{< ext "pgcryptokey" >}} {{< ext "anon" >}} {{< ext "pgcrypto" >}} {{< ext "pgsmcrypto" >}} |

> [!Note] Upstream 0.3.2/pkg 0.3.1; restrict decrypted view access; keep plaintext out of logs; pgsodium no longer required.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_vault` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.1` | {{< bg "18" "vault_18" "green" >}} {{< bg "17" "vault_17" "green" >}} {{< bg "16" "vault_16" "green" >}} {{< bg "15" "vault_15" "green" >}} {{< bg "14" "vault_14" "green" >}} | `vault_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.1` | {{< bg "18" "postgresql-18-vault" "green" >}} {{< bg "17" "postgresql-17-vault" "green" >}} {{< bg "16" "postgresql-16-vault" "green" >}} {{< bg "15" "postgresql-15-vault" "green" >}} {{< bg "14" "postgresql-14-vault" "green" >}} | `postgresql-$v-vault` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "vault_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "vault_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-18-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-17-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-vault : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-vault : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

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
| `postgresql-18-vault` | `0.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.4 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-vault` | `0.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.1 KiB | [postgresql-18-vault_0.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-18-vault_0.3.1-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

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
| `postgresql-17-vault` | `0.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.4 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-vault` | `0.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.1 KiB | [postgresql-17-vault_0.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-17-vault_0.3.1-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

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
| `postgresql-16-vault` | `0.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.4 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-vault` | `0.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.1 KiB | [postgresql-16-vault_0.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-16-vault_0.3.1-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

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
| `postgresql-15-vault` | `0.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.8 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-vault` | `0.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.5 KiB | [postgresql-15-vault_0.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-15-vault_0.3.1-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

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
| `postgresql-14-vault` | `0.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.8 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-vault` | `0.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.5 KiB | [postgresql-14-vault_0.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/v/vault/postgresql-14-vault_0.3.1-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/vault" title="Repository" icon="github" subtitle="github.com/supabase/vault" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="vault-0.3.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_vault;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_vault;		# install via package name, for the active PG version
pig install supabase_vault;		# install by extension name, for the current active PG version

pig install supabase_vault -v 18;   # install for PG 18
pig install supabase_vault -v 17;   # install for PG 17
pig install supabase_vault -v 16;   # install for PG 16
pig install supabase_vault -v 15;   # install for PG 15
pig install supabase_vault -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION supabase_vault;
```

## Usage

Sources:

- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/README.md)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/supabase_vault.control.in)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/Makefile)
- [Official source](https://github.com/supabase/vault/blob/2f05b9be4dcac774a458d5a9e6acae794eb6b8c4/sql/supabase_vault--0.3.1--0.3.2.sql)

`supabase_vault` 0.3.2 stores encrypted secrets in PostgreSQL and exposes decrypted values through a privileged view. Its fixed schema is `vault`; the current release no longer requires `pgsodium`. Upstream tests PostgreSQL 14–19. Install as a superuser; no preload or restart is required.

### Core Workflow

The following values are demonstration strings, not real credentials:

```sql
CREATE EXTENSION supabase_vault;
SELECT vault.create_secret('example-secret', 'api_token', 'Demo only');
SELECT name, decrypted_secret FROM vault.decrypted_secrets WHERE name = 'api_token';
SELECT vault.update_secret(id, 'replacement-example', 'api_token', 'Rotated demo')
FROM vault.secrets WHERE name = 'api_token';
```

`vault.create_secret` returns the secret UUID and accepts optional name and description. `vault.update_secret` updates a selected UUID. `vault.secrets` holds encrypted values and metadata; `vault.decrypted_secrets` exposes plaintext to roles allowed to read it. Restrict both helper and view privileges deliberately.

### Security and Upgrade

Version 0.3.2 moves `pg_temp` to the end of the search path used by secret creation/update helpers. After installing its files, run `ALTER EXTENSION supabase_vault UPDATE TO '0.3.2'` in each database.

Encryption at rest does not prevent authorized queries from returning plaintext. Secret literals and returned values can appear in database logs, client histories or application telemetry; use the documented secret workflow and configure those systems appropriately. Protect the encryption-key environment and verify backup/restore behavior. Do not treat access to the decrypted view as safe merely because the stored table is encrypted.
