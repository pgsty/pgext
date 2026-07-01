---
title: "wrappers"
linkTitle: "wrappers"
description: "Foreign data wrappers developed by Supabase"
weight: 8500
categories: ["FDW"]
width: full
---

[**wrappers**](https://github.com/supabase/wrappers) : Foreign data wrappers developed by Supabase


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8500** | {{< badge content="wrappers" link="https://github.com/supabase/wrappers" >}} | {{< ext "wrappers" >}} | `0.6.1` | {{< category "FDW" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "pgspider_ext" >}} |

> [!Note] pgrx patched to 0.18.1.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `wrappers` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "wrappers_18" "green" >}} {{< bg "17" "wrappers_17" "green" >}} {{< bg "16" "wrappers_16" "green" >}} {{< bg "15" "wrappers_15" "green" >}} {{< bg "14" "wrappers_14" "green" >}} | `wrappers_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "postgresql-18-wrappers" "green" >}} {{< bg "17" "postgresql-17-wrappers" "green" >}} {{< bg "16" "postgresql-16-wrappers" "green" >}} {{< bg "15" "postgresql-15-wrappers" "green" >}} {{< bg "14" "postgresql-14-wrappers" "green" >}} | `postgresql-$v-wrappers` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-wrappers : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_18` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 460.9 KiB | [wrappers_18-0.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_18-0.6.1-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_18` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 440.1 KiB | [wrappers_18-0.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_18-0.6.1-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_18` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 468.0 KiB | [wrappers_18-0.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_18-0.6.1-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_18` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 465.5 KiB | [wrappers_18-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_18-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_18` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 467.4 KiB | [wrappers_18-0.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_18-0.6.1-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_18` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 465.7 KiB | [wrappers_18-0.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_18-0.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-wrappers` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 372.2 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 331.6 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 372.5 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 332.4 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.5 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 388.4 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 405.9 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 383.8 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 400.2 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 381.8 KiB | [postgresql-18-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-18-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_17` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 459.8 KiB | [wrappers_17-0.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_17-0.6.1-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_17` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 440.0 KiB | [wrappers_17-0.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_17-0.6.1-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_17` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 466.7 KiB | [wrappers_17-0.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_17-0.6.1-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_17` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 464.7 KiB | [wrappers_17-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_17-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_17` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 466.8 KiB | [wrappers_17-0.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_17-0.6.1-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_17` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 464.7 KiB | [wrappers_17-0.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_17-0.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-wrappers` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 371.6 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 331.1 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 371.9 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 331.9 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.5 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 388.3 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 405.2 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 383.5 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 400.1 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 381.2 KiB | [postgresql-17-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-17-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_16` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 459.3 KiB | [wrappers_16-0.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_16-0.6.1-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_16` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 439.6 KiB | [wrappers_16-0.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_16-0.6.1-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_16` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 466.3 KiB | [wrappers_16-0.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_16-0.6.1-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_16` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 464.3 KiB | [wrappers_16-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_16-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_16` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 466.9 KiB | [wrappers_16-0.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_16-0.6.1-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_16` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 464.3 KiB | [wrappers_16-0.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_16-0.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-wrappers` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 371.8 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 332.1 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 372.3 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 331.0 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.4 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 388.1 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 405.6 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 383.6 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 400.2 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 381.3 KiB | [postgresql-16-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-16-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_15` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 454.9 KiB | [wrappers_15-0.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_15-0.6.1-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_15` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 435.3 KiB | [wrappers_15-0.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_15-0.6.1-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_15` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 461.8 KiB | [wrappers_15-0.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_15-0.6.1-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_15` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 460.1 KiB | [wrappers_15-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_15-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_15` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 462.0 KiB | [wrappers_15-0.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_15-0.6.1-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_15` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 460.5 KiB | [wrappers_15-0.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_15-0.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-wrappers` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 368.6 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 328.8 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 368.2 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 328.8 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.3 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 385.2 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 401.9 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 380.4 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 400.1 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 378.5 KiB | [postgresql-15-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-15-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_14` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 454.7 KiB | [wrappers_14-0.6.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_14-0.6.1-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_14` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 435.1 KiB | [wrappers_14-0.6.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_14-0.6.1-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_14` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 461.6 KiB | [wrappers_14-0.6.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_14-0.6.1-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_14` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 459.5 KiB | [wrappers_14-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_14-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_14` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 461.8 KiB | [wrappers_14-0.6.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_14-0.6.1-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_14` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 460.4 KiB | [wrappers_14-0.6.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_14-0.6.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-wrappers` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 368.3 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 329.0 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 368.2 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.9 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.4 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 385.6 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 401.9 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 380.5 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 400.1 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 378.7 KiB | [postgresql-14-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-14-wrappers_0.6.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/wrappers" title="Repository" icon="github" subtitle="github.com/supabase/wrappers" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="wrappers-0.6.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg wrappers;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install wrappers;		# install via package name, for the active PG version

pig install wrappers -v 18;   # install for PG 18
pig install wrappers -v 17;   # install for PG 17
pig install wrappers -v 16;   # install for PG 16
pig install wrappers -v 15;   # install for PG 15
pig install wrappers -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION wrappers;
```




## Usage

Sources: [official README](https://github.com/supabase/wrappers/blob/v0.6.1/README.md), [official docs](https://fdw.dev/), [v0.6.1 release](https://github.com/supabase/wrappers/releases/tag/v0.6.1)

`wrappers` is both a Rust framework for writing PostgreSQL foreign data wrappers and a packaged collection of Supabase-maintained FDWs. A single extension installs many wrapper implementations, then each foreign server chooses the specific wrapper type it needs.

```sql
CREATE EXTENSION wrappers;
```

### Typical Workflow

Create a server for one wrapper, then expose remote data through foreign tables:

```sql
CREATE SERVER stripe_server
  FOREIGN DATA WRAPPER stripe_wrapper
  OPTIONS (
    api_key_id 'stripe_api_key',
    api_url 'https://api.stripe.com/v1/'
  );

CREATE FOREIGN TABLE stripe_customers (
  id text,
  email text,
  name text,
  description text,
  created timestamp,
  attrs jsonb
)
  SERVER stripe_server
  OPTIONS (
    object 'customers',
    rowid_column 'id'
  );
```

### What It Covers

Upstream ships wrappers for databases and services such as BigQuery, ClickHouse, DuckDB, DynamoDB, MySQL/Doris, Redis, S3, S3 Vectors, Stripe, Snowflake, Slack, Notion, OpenAPI, Infura, and many others. Read and write support varies by wrapper, but pushdown for `WHERE`, `ORDER BY`, and `LIMIT` is a core framework feature.

### Version Notes

The `v0.6.1` release keeps the same extension model but expands the catalog and wrapper behavior. Official release notes call out:

- new DynamoDB FDW support
- MySQL/Doris support through `mysql_fdw`
- schema evolution support for `iceberg_fdw`
- vault secret lookup by name in `_id` options
- aggregate pushdown support for COUNT, SUM, AVG, MIN, and MAX, including MySQL FDW support
- parameter-state refresh/rescan fixes and dependency/security updates

### Caveats

- Wrapper-specific options, supported objects, and write support differ widely; check the official catalog page for the exact FDW you use.
- The docs warn that logical restores can fail when materialized views depend on foreign tables, so avoid that pattern or rely on physical backups.
