---
title: "block_copy_command"
linkTitle: "block_copy_command"
description: "Block COPY commands via a configurable ProcessUtility hook"
weight: 7430
categories: ["SEC"]
width: full
---

[**block_copy_command**](https://github.com/rustwizard/block_copy_command) : Block COPY commands via a configurable ProcessUtility hook


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7430** | {{< badge content="block_copy_command" link="https://github.com/rustwizard/block_copy_command" >}} | {{< ext "block_copy_command" >}} | `0.1.5` | {{< category "SEC" >}} | {{< license "BSD-3-Clause" >}} | {{< language "Rust" >}} |


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
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-block-copy-command : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-block-copy-command : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_18` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 857.1 KiB | [block_copy_command_18-0.1.5-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_18-0.1.5-3PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 771.0 KiB | [block_copy_command_18-0.1.5-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_18-0.1.5-3PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 865.8 KiB | [block_copy_command_18-0.1.5-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_18-0.1.5-3PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 817.4 KiB | [block_copy_command_18-0.1.5-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_18-0.1.5-3PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 865.4 KiB | [block_copy_command_18-0.1.5-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_18-0.1.5-3PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_18` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 796.6 KiB | [block_copy_command_18-0.1.5-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_18-0.1.5-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 679.7 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 569.3 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 679.8 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 569.0 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 759.0 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 671.5 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 749.4 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 662.9 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 747.0 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-block-copy-command` | `0.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 661.9 KiB | [postgresql-18-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-18-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_17` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 854.4 KiB | [block_copy_command_17-0.1.5-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_17-0.1.5-3PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 767.8 KiB | [block_copy_command_17-0.1.5-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_17-0.1.5-3PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 862.4 KiB | [block_copy_command_17-0.1.5-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_17-0.1.5-3PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 814.6 KiB | [block_copy_command_17-0.1.5-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_17-0.1.5-3PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 860.1 KiB | [block_copy_command_17-0.1.5-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_17-0.1.5-3PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_17` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 795.2 KiB | [block_copy_command_17-0.1.5-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_17-0.1.5-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 678.0 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 568.1 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 678.4 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 567.6 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 757.0 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 670.0 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 747.8 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 663.2 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 744.0 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-block-copy-command` | `0.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 660.3 KiB | [postgresql-17-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-17-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_16` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 854.0 KiB | [block_copy_command_16-0.1.5-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_16-0.1.5-3PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 766.4 KiB | [block_copy_command_16-0.1.5-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_16-0.1.5-3PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 861.7 KiB | [block_copy_command_16-0.1.5-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_16-0.1.5-3PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 813.6 KiB | [block_copy_command_16-0.1.5-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_16-0.1.5-3PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 862.0 KiB | [block_copy_command_16-0.1.5-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_16-0.1.5-3PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_16` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 795.2 KiB | [block_copy_command_16-0.1.5-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_16-0.1.5-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 678.0 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 566.9 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 678.4 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 567.7 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 755.7 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 669.2 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 747.8 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 661.2 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 743.4 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-block-copy-command` | `0.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 659.7 KiB | [postgresql-16-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-16-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_15` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 843.9 KiB | [block_copy_command_15-0.1.5-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_15-0.1.5-3PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 757.3 KiB | [block_copy_command_15-0.1.5-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_15-0.1.5-3PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 851.5 KiB | [block_copy_command_15-0.1.5-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_15-0.1.5-3PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 803.1 KiB | [block_copy_command_15-0.1.5-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_15-0.1.5-3PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 851.2 KiB | [block_copy_command_15-0.1.5-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_15-0.1.5-3PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_15` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 789.8 KiB | [block_copy_command_15-0.1.5-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_15-0.1.5-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 672.1 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 562.5 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 672.1 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 562.8 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 748.3 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 664.2 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 740.0 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 656.2 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 736.4 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-block-copy-command` | `0.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 653.5 KiB | [postgresql-15-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-15-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `block_copy_command_14` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 840.4 KiB | [block_copy_command_14-0.1.5-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/block_copy_command_14-0.1.5-3PIGSTY.el8.x86_64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 754.8 KiB | [block_copy_command_14-0.1.5-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/block_copy_command_14-0.1.5-3PIGSTY.el8.aarch64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 847.5 KiB | [block_copy_command_14-0.1.5-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/block_copy_command_14-0.1.5-3PIGSTY.el9.x86_64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 800.0 KiB | [block_copy_command_14-0.1.5-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/block_copy_command_14-0.1.5-3PIGSTY.el9.aarch64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 848.7 KiB | [block_copy_command_14-0.1.5-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/block_copy_command_14-0.1.5-3PIGSTY.el10.x86_64.rpm) |
| `block_copy_command_14` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 788.6 KiB | [block_copy_command_14-0.1.5-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/block_copy_command_14-0.1.5-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-block-copy-command` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 668.6 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 560.8 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 668.9 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 561.4 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 744.0 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 663.9 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 735.7 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 653.4 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 732.7 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-block-copy-command` | `0.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 651.5 KiB | [postgresql-14-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/b/block-copy-command/postgresql-14-block-copy-command_0.1.5-3PIGSTY~resolute_arm64.deb) |

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

- Source: [README](https://github.com/rustwizard/block_copy_command/blob/master/README.md)

`block_copy_command` installs a `ProcessUtility` hook that intercepts `COPY` statements. The hook is cluster-wide once the library is loaded, while `CREATE EXTENSION` only registers metadata in a database.

### Enable It

```conf
shared_preload_libraries = 'block_copy_command'
```

```sql
CREATE EXTENSION block_copy_command;
```

### Blocking Rules

By default, non-superusers cannot run `COPY TO` or `COPY FROM`:

```sql
COPY my_table TO STDOUT;
COPY my_table FROM STDIN;
COPY (SELECT * FROM my_table) TO '/tmp/out.csv';
```

Priority is documented as:

- `block_copy_command.blocked_roles`: always blocked, even superusers.
- `block_copy_command.block_program = on`: blocks `COPY ... PROGRAM` for everyone.
- `block_copy_command.enabled = off`: allows `COPY` for roles not in `blocked_roles`.
- Superusers otherwise bypass direction blocking.
- `block_copy_command.block_to` and `block_copy_command.block_from` control export/import blocking for non-superusers.

### Main Settings

- `block_copy_command.enabled`: master switch for non-superuser blocking.
- `block_copy_command.block_to`: block `COPY TO`.
- `block_copy_command.block_from`: block `COPY FROM`.
- `block_copy_command.block_program`: block `COPY TO/FROM PROGRAM` for all users.
- `block_copy_command.hint`: append a custom `HINT` to blocked-command errors.
- `block_copy_command.blocked_roles`: comma-separated always-blocked roles.
- `block_copy_command.audit_log_enabled`: write intercepted events to the audit table.

### Audit And Caveats

Allowed and blocked attempts are intercepted, and the extension defines `block_copy_command.audit_log` plus server-log entries for blocked events. The README notes one important caveat: blocked audit rows are inserted before the error is raised, so they are rolled back with the transaction. In practice, PostgreSQL server logs are the authoritative record for blocked `COPY` attempts.
