---
title: "pg_command_fw"
linkTitle: "pg_command_fw"
description: "DDL and utility command firewall for PostgreSQL"
weight: 7400
categories: ["SEC"]
width: full
---

[**pg_command_fw**](https://github.com/rustwizard/pg_command_fw) : DDL and utility command firewall for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7400** | {{< badge content="pg_command_fw" link="https://github.com/rustwizard/pg_command_fw" >}} | {{< ext "pg_command_fw" >}} | `0.1.0` | {{< category "SEC" >}} | {{< license "BSD-3-Clause" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgaudit" >}} {{< ext "pgextwlist" >}} {{< ext "login_hook" >}} {{< ext "set_user" >}} |

> [!Note] Requires shared_preload_libraries = pg_command_fw to activate hooks for all sessions.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_command_fw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_command_fw_18" "green" >}} {{< bg "17" "pg_command_fw_17" "green" >}} {{< bg "16" "pg_command_fw_16" "green" >}} {{< bg "15" "pg_command_fw_15" "green" >}} {{< bg "14" "pg_command_fw_14" "red" >}} | `pg_command_fw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-command-fw" "green" >}} {{< bg "17" "postgresql-17-pg-command-fw" "green" >}} {{< bg "16" "postgresql-16-pg-command-fw" "green" >}} {{< bg "15" "postgresql-15-pg-command-fw" "green" >}} {{< bg "14" "postgresql-14-pg-command-fw" "red" >}} | `postgresql-$v-pg-command-fw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_command_fw_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_command_fw_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_command_fw_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_command_fw_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_command_fw_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_command_fw_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-command-fw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-command-fw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-command-fw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-command-fw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-command-fw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-command-fw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-command-fw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-command-fw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-command-fw : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_command_fw_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 313.1 KiB | [pg_command_fw_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_command_fw_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 204.4 KiB | [pg_command_fw_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_command_fw_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 329.1 KiB | [pg_command_fw_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_command_fw_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 216.5 KiB | [pg_command_fw_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_command_fw_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 329.4 KiB | [pg_command_fw_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_command_fw_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 217.4 KiB | [pg_command_fw_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_command_fw_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 255.4 KiB | [postgresql-18-pg-command-fw_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 155.2 KiB | [postgresql-18-pg-command-fw_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 255.2 KiB | [postgresql-18-pg-command-fw_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 155.3 KiB | [postgresql-18-pg-command-fw_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 288.7 KiB | [postgresql-18-pg-command-fw_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 179.7 KiB | [postgresql-18-pg-command-fw_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 286.2 KiB | [postgresql-18-pg-command-fw_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 178.3 KiB | [postgresql-18-pg-command-fw_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_command_fw_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 313.4 KiB | [pg_command_fw_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_command_fw_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 204.5 KiB | [pg_command_fw_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_command_fw_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 329.3 KiB | [pg_command_fw_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_command_fw_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 216.4 KiB | [pg_command_fw_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_command_fw_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 329.3 KiB | [pg_command_fw_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_command_fw_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 217.4 KiB | [pg_command_fw_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_command_fw_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 255.2 KiB | [postgresql-17-pg-command-fw_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 155.2 KiB | [postgresql-17-pg-command-fw_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 254.5 KiB | [postgresql-17-pg-command-fw_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 155.1 KiB | [postgresql-17-pg-command-fw_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 288.7 KiB | [postgresql-17-pg-command-fw_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 179.7 KiB | [postgresql-17-pg-command-fw_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 285.8 KiB | [postgresql-17-pg-command-fw_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 178.1 KiB | [postgresql-17-pg-command-fw_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_command_fw_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 313.0 KiB | [pg_command_fw_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_command_fw_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 204.4 KiB | [pg_command_fw_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_command_fw_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 329.0 KiB | [pg_command_fw_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_command_fw_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 216.4 KiB | [pg_command_fw_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_command_fw_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 328.9 KiB | [pg_command_fw_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_command_fw_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 217.5 KiB | [pg_command_fw_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_command_fw_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 255.2 KiB | [postgresql-16-pg-command-fw_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 155.2 KiB | [postgresql-16-pg-command-fw_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 255.1 KiB | [postgresql-16-pg-command-fw_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 155.2 KiB | [postgresql-16-pg-command-fw_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 288.4 KiB | [postgresql-16-pg-command-fw_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 179.7 KiB | [postgresql-16-pg-command-fw_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 286.1 KiB | [postgresql-16-pg-command-fw_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 178.3 KiB | [postgresql-16-pg-command-fw_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_command_fw_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 311.5 KiB | [pg_command_fw_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_command_fw_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 202.6 KiB | [pg_command_fw_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_command_fw_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 327.3 KiB | [pg_command_fw_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_command_fw_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 215.1 KiB | [pg_command_fw_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_command_fw_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 327.5 KiB | [pg_command_fw_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_command_fw_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 216.0 KiB | [pg_command_fw_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_command_fw_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 254.2 KiB | [postgresql-15-pg-command-fw_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 153.8 KiB | [postgresql-15-pg-command-fw_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 254.2 KiB | [postgresql-15-pg-command-fw_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 153.8 KiB | [postgresql-15-pg-command-fw_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 287.0 KiB | [postgresql-15-pg-command-fw_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 178.0 KiB | [postgresql-15-pg-command-fw_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 284.1 KiB | [postgresql-15-pg-command-fw_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 176.3 KiB | [postgresql-15-pg-command-fw_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rustwizard/pg_command_fw" title="Repository" icon="github" subtitle="github.com/rustwizard/pg_command_fw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_command_fw-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_command_fw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_command_fw;		# install via package name, for the active PG version

pig install pg_command_fw -v 18;   # install for PG 18
pig install pg_command_fw -v 17;   # install for PG 17
pig install pg_command_fw -v 16;   # install for PG 16
pig install pg_command_fw -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_command_fw';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_command_fw;
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_command_fw;
> ALTER SYSTEM SET pg_command_fw.block_truncate = on;
> ALTER SYSTEM SET pg_command_fw.production_schemas = 'public,payments';
> SELECT pg_reload_conf();
> ```
>
> Source: [README](https://github.com/rustwizard/pg_command_fw)

`pg_command_fw` is a PostgreSQL command firewall. It intercepts DDL and utility commands through the `ProcessUtility` hook and blocks selected built-in file-reading functions through the post-parse analyze hook. Each command category is controlled by its own GUC.

## Setup

The extension must be preloaded:

```ini
shared_preload_libraries = 'pg_command_fw'
```

Then enable it in the database:

```sql
CREATE EXTENSION pg_command_fw;
```

## Command Categories

The upstream README documents these firewall categories:

- `TRUNCATE`
- `DROP TABLE`
- `ALTER SYSTEM`
- `LOAD`
- `COPY ... PROGRAM`
- plain `COPY`
- `pg_read_file()`, `pg_read_binary_file()`, and `pg_stat_file()`

Some categories block only non-superusers, while others block everyone including superusers. Superusers are only exempt from non-superuser categories unless they are explicitly listed in `pg_command_fw.blocked_roles`.

## Important GUCs

- `pg_command_fw.enabled` to enable or disable all checks
- `pg_command_fw.block_truncate`
- `pg_command_fw.block_drop_table`
- `pg_command_fw.production_schemas`
- `pg_command_fw.block_alter_system`
- `pg_command_fw.block_load`
- `pg_command_fw.block_copy_program`
- `pg_command_fw.block_copy`
- `pg_command_fw.block_read_file`
- `pg_command_fw.blocked_roles`
- `pg_command_fw.hint`
- `pg_command_fw.audit_log_enabled`

## Audit Log

The extension records intercepted commands in `command_fw.audit_log`. The README documents columns such as:

- timestamp
- session and current user names
- original query text
- command type
- target schema or object
- client address
- whether the command was blocked
- internal block reason

## Examples

Block `TRUNCATE` and `DROP TABLE` in production schemas:

```sql
ALTER SYSTEM SET pg_command_fw.block_truncate = on;
ALTER SYSTEM SET pg_command_fw.block_drop_table = on;
ALTER SYSTEM SET pg_command_fw.production_schemas = 'public,payments';
ALTER SYSTEM SET pg_command_fw.hint = 'Contact your DBA to request access';
SELECT pg_reload_conf();
```

Block a specific role from any governed command:

```sql
ALTER SYSTEM SET pg_command_fw.blocked_roles = 'app_deploy';
SELECT pg_reload_conf();
```
