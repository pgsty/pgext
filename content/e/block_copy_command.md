---
title: "block_copy_command"
linkTitle: "block_copy_command"
description: "Block COPY commands via a configurable ProcessUtility hook"
weight: 7405
categories: ["SEC"]
width: full
---

[**block_copy_command**](https://github.com/rustwizard/block_copy_command) : Block COPY commands via a configurable ProcessUtility hook


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7405** | {{< badge content="block_copy_command" link="https://github.com/rustwizard/block_copy_command" >}} | {{< ext "block_copy_command" >}} | `0.1.5` | {{< category "SEC" >}} | {{< license "BSD 3-Clause" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Requires shared_preload_libraries = block_copy_command.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `block_copy_command` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "block_copy_command_18" "green" >}} {{< bg "17" "block_copy_command_17" "green" >}} {{< bg "16" "block_copy_command_16" "green" >}} {{< bg "15" "block_copy_command_15" "green" >}} {{< bg "14" "block_copy_command_14" "green" >}} | `block_copy_command_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "postgresql-18-block-copy-command" "green" >}} {{< bg "17" "postgresql-17-block-copy-command" "green" >}} {{< bg "16" "postgresql-16-block-copy-command" "green" >}} {{< bg "15" "postgresql-15-block-copy-command" "green" >}} {{< bg "14" "postgresql-14-block-copy-command" "green" >}} | `postgresql-$v-block-copy-command` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "block_copy_command_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_18` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 306.1 KiB | [block_copy_command_18-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_18-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 199.0 KiB | [block_copy_command_18-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_18-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 321.7 KiB | [block_copy_command_18-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_18-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 212.1 KiB | [block_copy_command_18-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_18-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 321.9 KiB | [block_copy_command_18-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_18-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 212.1 KiB | [block_copy_command_18-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_18-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 248.3 KiB | [postgresql-18-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 149.9 KiB | [postgresql-18-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 248.3 KiB | [postgresql-18-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 149.9 KiB | [postgresql-18-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 281.0 KiB | [postgresql-18-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 173.9 KiB | [postgresql-18-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 278.5 KiB | [postgresql-18-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 172.4 KiB | [postgresql-18-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_17` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 306.0 KiB | [block_copy_command_17-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_17-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 199.1 KiB | [block_copy_command_17-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_17-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 321.6 KiB | [block_copy_command_17-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_17-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 212.1 KiB | [block_copy_command_17-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_17-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 321.7 KiB | [block_copy_command_17-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_17-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 212.1 KiB | [block_copy_command_17-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_17-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 247.4 KiB | [postgresql-17-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 149.9 KiB | [postgresql-17-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 247.6 KiB | [postgresql-17-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.0 KiB | [postgresql-17-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 280.9 KiB | [postgresql-17-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 173.9 KiB | [postgresql-17-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 278.4 KiB | [postgresql-17-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 172.8 KiB | [postgresql-17-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_16` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 305.9 KiB | [block_copy_command_16-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_16-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 199.1 KiB | [block_copy_command_16-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_16-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 321.3 KiB | [block_copy_command_16-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_16-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 212.6 KiB | [block_copy_command_16-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_16-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 321.5 KiB | [block_copy_command_16-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_16-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 212.1 KiB | [block_copy_command_16-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_16-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 248.1 KiB | [postgresql-16-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 149.9 KiB | [postgresql-16-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 248.2 KiB | [postgresql-16-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 149.8 KiB | [postgresql-16-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 281.0 KiB | [postgresql-16-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 174.0 KiB | [postgresql-16-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 278.4 KiB | [postgresql-16-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 172.8 KiB | [postgresql-16-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_15` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 305.9 KiB | [block_copy_command_15-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_15-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 199.1 KiB | [block_copy_command_15-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_15-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 321.8 KiB | [block_copy_command_15-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_15-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 212.0 KiB | [block_copy_command_15-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_15-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 322.0 KiB | [block_copy_command_15-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_15-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 212.1 KiB | [block_copy_command_15-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_15-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 247.4 KiB | [postgresql-15-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 149.9 KiB | [postgresql-15-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 247.4 KiB | [postgresql-15-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 149.9 KiB | [postgresql-15-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 280.7 KiB | [postgresql-15-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 173.9 KiB | [postgresql-15-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 278.0 KiB | [postgresql-15-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 173.0 KiB | [postgresql-15-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_14` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 305.6 KiB | [block_copy_command_14-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_14-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 199.0 KiB | [block_copy_command_14-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_14-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 321.9 KiB | [block_copy_command_14-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_14-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 212.5 KiB | [block_copy_command_14-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_14-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 322.1 KiB | [block_copy_command_14-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_14-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 212.2 KiB | [block_copy_command_14-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_14-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 248.3 KiB | [postgresql-14-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 149.9 KiB | [postgresql-14-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 247.9 KiB | [postgresql-14-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 149.9 KiB | [postgresql-14-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 280.7 KiB | [postgresql-14-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 174.0 KiB | [postgresql-14-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 278.2 KiB | [postgresql-14-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 172.8 KiB | [postgresql-14-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rustwizard/block_copy_command" title="Repository" icon="github" subtitle="github.com/rustwizard/block_copy_command" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="block_copy_command-0.1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg block_copy_command;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install block_copy_command;		# install via package name, for the active PG version

pig install block_copy_command -v 18;   # install for PG 18
pig install block_copy_command -v 17;   # install for PG 17
pig install block_copy_command -v 16;   # install for PG 16
pig install block_copy_command -v 15;   # install for PG 15
pig install block_copy_command -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'block_copy_command';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION block_copy_command;
```

## Usage
- GitHub Repo: [`rustwizard/block_copy_command`](https://github.com/rustwizard/block_copy_command)
- README: [rustwizard/block_copy_command/blob/master/README.md](https://github.com/rustwizard/block_copy_command/blob/master/README.md)

`block_copy_command` blocks `COPY` commands cluster-wide by installing a `ProcessUtility` hook. It is loaded with `shared_preload_libraries`, and `CREATE EXTENSION` only registers the extension metadata in each database.

This extension is intended for deployments that want to stop `COPY TO` and `COPY FROM` by default for non-superusers, while still allowing finer-grained policy through GUCs and an audit table.

### Setup

```conf
shared_preload_libraries = 'block_copy_command'
```

```sql
CREATE EXTENSION block_copy_command;
```

The README says the hook becomes active for the whole cluster as soon as the library is loaded.

### Blocking Rules

By default, non-superusers are blocked from running `COPY`.

```sql
COPY my_table TO STDOUT;
COPY my_table FROM STDIN;
COPY (SELECT * FROM my_table) TO '/tmp/out.csv';
```

Superusers bypass the block unless they are listed in `block_copy_command.blocked_roles` or `block_copy_command.block_program` is enabled. `COPY ... PROGRAM` is blocked for everyone by default.

### Settings

- `block_copy_command.enabled` toggles blocking for non-superusers.
- `block_copy_command.block_to` controls whether `COPY TO` is blocked.
- `block_copy_command.block_from` controls whether `COPY FROM` is blocked.
- `block_copy_command.block_program` blocks `COPY TO/FROM PROGRAM` for all users.
- `block_copy_command.hint` appends a custom `HINT:` to blocked commands.
- `block_copy_command.blocked_roles` permanently blocks named roles, including superusers.
- `block_copy_command.audit_log_enabled` controls whether intercepted `COPY` events are written to `block_copy_command.audit_log`.

### Audit Log

The extension records intercepted `COPY` activity in `block_copy_command.audit_log` and also writes blocked events to the PostgreSQL server log at `LOG` level.

Typical monitoring queries from the README include listing recent events, filtering blocked events, and grouping by user.

### Scope

The upstream README covers requirements, enablement, blocking behavior, the main GUCs, the audit table, and test coverage. No separate project homepage or docs site was needed for this stub.
