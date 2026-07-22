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
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_command_fw_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_command_fw_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_command_fw_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_command_fw_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_command_fw_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_command_fw_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_command_fw_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-command-fw : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-command-fw : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_command_fw_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 864.7 KiB | [pg_command_fw_18-0.1.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_command_fw_18-0.1.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 775.1 KiB | [pg_command_fw_18-0.1.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_command_fw_18-0.1.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 872.7 KiB | [pg_command_fw_18-0.1.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_command_fw_18-0.1.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 822.2 KiB | [pg_command_fw_18-0.1.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_command_fw_18-0.1.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 873.1 KiB | [pg_command_fw_18-0.1.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_command_fw_18-0.1.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_command_fw_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 800.7 KiB | [pg_command_fw_18-0.1.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_command_fw_18-0.1.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 687.9 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 573.3 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 687.2 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 574.1 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 766.0 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 678.5 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 757.6 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 668.3 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 752.8 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-command-fw` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 666.8 KiB | [postgresql-18-pg-command-fw_0.1.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-command-fw/postgresql-18-pg-command-fw_0.1.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_command_fw_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 861.5 KiB | [pg_command_fw_17-0.1.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_command_fw_17-0.1.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 772.7 KiB | [pg_command_fw_17-0.1.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_command_fw_17-0.1.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 869.4 KiB | [pg_command_fw_17-0.1.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_command_fw_17-0.1.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 819.1 KiB | [pg_command_fw_17-0.1.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_command_fw_17-0.1.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 869.5 KiB | [pg_command_fw_17-0.1.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_command_fw_17-0.1.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_command_fw_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 800.4 KiB | [pg_command_fw_17-0.1.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_command_fw_17-0.1.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 685.4 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 571.9 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 685.3 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 572.8 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 763.3 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 675.5 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 754.4 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 666.6 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 749.7 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-command-fw` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 665.2 KiB | [postgresql-17-pg-command-fw_0.1.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-command-fw/postgresql-17-pg-command-fw_0.1.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_command_fw_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 861.0 KiB | [pg_command_fw_16-0.1.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_command_fw_16-0.1.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 770.9 KiB | [pg_command_fw_16-0.1.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_command_fw_16-0.1.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 869.6 KiB | [pg_command_fw_16-0.1.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_command_fw_16-0.1.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 818.3 KiB | [pg_command_fw_16-0.1.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_command_fw_16-0.1.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 869.6 KiB | [pg_command_fw_16-0.1.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_command_fw_16-0.1.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_command_fw_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 800.5 KiB | [pg_command_fw_16-0.1.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_command_fw_16-0.1.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 685.8 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 571.6 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 685.4 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 571.7 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 763.5 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 674.2 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 755.1 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 665.8 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 750.4 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-command-fw` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 664.5 KiB | [postgresql-16-pg-command-fw_0.1.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-command-fw/postgresql-16-pg-command-fw_0.1.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_command_fw_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 849.0 KiB | [pg_command_fw_15-0.1.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_command_fw_15-0.1.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 760.1 KiB | [pg_command_fw_15-0.1.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_command_fw_15-0.1.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 857.0 KiB | [pg_command_fw_15-0.1.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_command_fw_15-0.1.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 806.2 KiB | [pg_command_fw_15-0.1.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_command_fw_15-0.1.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 856.5 KiB | [pg_command_fw_15-0.1.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_command_fw_15-0.1.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_command_fw_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 794.3 KiB | [pg_command_fw_15-0.1.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_command_fw_15-0.1.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 677.3 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 565.4 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 677.6 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 565.9 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 754.1 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 667.3 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 745.2 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 658.7 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 741.7 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-command-fw` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 656.5 KiB | [postgresql-15-pg-command-fw_0.1.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-command-fw/postgresql-15-pg-command-fw_0.1.0-3PIGSTY~resolute_arm64.deb) |

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

- Source: [README](https://github.com/rustwizard/pg_command_fw/blob/master/README.md)

`pg_command_fw` is a PostgreSQL command firewall. It intercepts DDL and utility commands through the `ProcessUtility` hook and blocks selected built-in file-reading functions through the post-parse analyze hook. Each command category is controlled by its own GUC.

### Enable It

The extension must be preloaded:

```ini
shared_preload_libraries = 'pg_command_fw'
```

Then enable it in the database:

```sql
CREATE EXTENSION pg_command_fw;
```

Pigsty package metadata is version `0.1.0` for PostgreSQL 15-18 and notes that preloading is required to activate hooks for all sessions. The upstream README also documents PostgreSQL 15-18 support.

### Command Categories

The upstream README documents these firewall categories:

- `TRUNCATE`: `pg_command_fw.block_truncate`, default `on`, blocks non-superusers.
- `DROP TABLE`: `pg_command_fw.block_drop_table`, default `off`, blocks non-superusers when enabled.
- `ALTER SYSTEM`: `pg_command_fw.block_alter_system`, default `on`, blocks everyone.
- `LOAD`: `pg_command_fw.block_load`, default `on`, blocks everyone.
- `COPY ... PROGRAM`: `pg_command_fw.block_copy_program`, default `on`, blocks everyone.
- plain `COPY`: `pg_command_fw.block_copy`, default `off`, blocks non-superusers when enabled.
- `pg_read_file()`, `pg_read_binary_file()`, and `pg_stat_file()`: `pg_command_fw.block_read_file`, default `on`, blocks everyone.

Some categories block only non-superusers, while others block everyone including superusers. Superusers are only exempt from non-superuser categories unless they are explicitly listed in `pg_command_fw.blocked_roles`.

### Important GUCs

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

When `production_schemas` is set, `DROP TABLE` checks are limited to schema-qualified table names in those schemas; the README says unqualified names are not resolved through `search_path`.

### Audit Log

The extension records intercepted commands in `command_fw.audit_log`. The README documents columns such as:

- timestamp
- session and current user names
- original query text
- command type
- target schema or object
- client address
- whether the command was blocked
- internal block reason

Blocked audit inserts are best-effort because the row is rolled back with the blocked transaction; use the PostgreSQL server log as the authoritative record for blocked events.

### Examples

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

Temporarily disable the firewall in a maintenance session:

```sql
SET pg_command_fw.enabled = off;
TRUNCATE big_table;
SET pg_command_fw.enabled = on;
```
