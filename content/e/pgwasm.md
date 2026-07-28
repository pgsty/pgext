---
title: "pgwasm"
linkTitle: "pgwasm"
description: "Run sandboxed WebAssembly components as strongly typed PostgreSQL SQL functions."
weight: 3150
categories: ["LANG"]
width: full
---

[**pgwasm**](https://github.com/jnicholls/pgwasm) : Run sandboxed WebAssembly components as strongly typed PostgreSQL SQL functions.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3150** | {{< badge content="pgwasm" link="https://github.com/jnicholls/pgwasm" >}} | {{< ext "pgwasm" >}} | `0.1.0` | {{< category "LANG" >}} | {{< license "BSD-3-Clause" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgwasm` |
|   **See Also**    | {{< ext "wasm" >}} {{< ext "pg_extism" >}} {{< ext "plrust" >}} |

> [!Note] No upstream tag or release; package pins commit 535b5336, ports pgrx 0.18 to 0.19.1, and supports PostgreSQL 14-18. Preloading is optional and enables shared metrics.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgwasm` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pgwasm_18" "green" >}} {{< bg "17" "pgwasm_17" "green" >}} {{< bg "16" "pgwasm_16" "green" >}} {{< bg "15" "pgwasm_15" "green" >}} {{< bg "14" "pgwasm_14" "green" >}} | `pgwasm_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pgwasm" "green" >}} {{< bg "17" "postgresql-17-pgwasm" "green" >}} {{< bg "16" "postgresql-16-pgwasm" "green" >}} {{< bg "15" "postgresql-15-pgwasm" "green" >}} {{< bg "14" "postgresql-14-pgwasm" "green" >}} | `postgresql-$v-pgwasm` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgwasm_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgwasm : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgwasm : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgwasm_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [pgwasm_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgwasm_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgwasm_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgwasm_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgwasm_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgwasm_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgwasm_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgwasm_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgwasm_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.3 MiB | [pgwasm_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgwasm_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgwasm_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pgwasm_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgwasm_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgwasm_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgwasm_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgwasm_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgwasm` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.8 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.8 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.5 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.7 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.5 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.5 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgwasm` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-18-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-18-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgwasm_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [pgwasm_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgwasm_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgwasm_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgwasm_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgwasm_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgwasm_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgwasm_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgwasm_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgwasm_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.3 MiB | [pgwasm_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgwasm_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgwasm_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pgwasm_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgwasm_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgwasm_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgwasm_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgwasm_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgwasm` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.8 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.8 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.5 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.7 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.5 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.5 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgwasm` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-17-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-17-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgwasm_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [pgwasm_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgwasm_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgwasm_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgwasm_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgwasm_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgwasm_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgwasm_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgwasm_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgwasm_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.3 MiB | [pgwasm_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgwasm_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgwasm_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pgwasm_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgwasm_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgwasm_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgwasm_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgwasm_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgwasm` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.8 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.8 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.5 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.7 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.5 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.5 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgwasm` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-16-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-16-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgwasm_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [pgwasm_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgwasm_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgwasm_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgwasm_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgwasm_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgwasm_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgwasm_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgwasm_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgwasm_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.3 MiB | [pgwasm_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgwasm_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgwasm_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pgwasm_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgwasm_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgwasm_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgwasm_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgwasm_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgwasm` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.8 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.8 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.5 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.7 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.5 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.5 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgwasm` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-15-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-15-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgwasm_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.3 MiB | [pgwasm_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgwasm_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgwasm_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgwasm_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgwasm_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgwasm_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgwasm_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgwasm_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgwasm_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.3 MiB | [pgwasm_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgwasm_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgwasm_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.0 MiB | [pgwasm_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgwasm_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgwasm_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgwasm_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgwasm_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgwasm` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.8 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.8 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.5 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.7 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.5 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.5 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgwasm` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-14-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgwasm/postgresql-14-pgwasm_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jnicholls/pgwasm" title="Repository" icon="github" subtitle="github.com/jnicholls/pgwasm" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgwasm-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgwasm;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgwasm;		# install via package name, for the active PG version

pig install pgwasm -v 18;   # install for PG 18
pig install pgwasm -v 17;   # install for PG 17
pig install pgwasm -v 16;   # install for PG 16
pig install pgwasm -v 15;   # install for PG 15
pig install pgwasm -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgwasm;
```

## Usage

Sources:

- [pgwasm README at the documented revision](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/README.md)
- [pgwasm architecture and SQL lifecycle](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/architecture.md)
- [pgwasm GUC reference](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/guc.md)
- [pgwasm WIT type mapping](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/docs/wit-mapping.md)
- [pgwasm control file](https://github.com/jnicholls/pgwasm/blob/535b53363f8208af139e757e508e66c46309ee29/pgwasm/pgwasm.control)

`pgwasm` loads WebAssembly components into PostgreSQL and registers WIT exports as typed PostgreSQL functions. Compiled artifacts are stored under the cluster data directory and reused by backend-local instance pools. This document follows pinned revision `535b53363f8208af139e757e508e66c46309ee29`; the source declares version 0.1.0 but does not provide a tagged 0.1.0 release.

### Core Workflow

Create the extension as a superuser. The following file-based workflow must be enabled and confined by an administrator before a loader role can use it:

```sql
CREATE EXTENSION pgwasm;

ALTER SYSTEM SET pgwasm.allow_load_from_file = on;
ALTER SYSTEM SET pgwasm.module_path = '/srv/pgwasm';
ALTER SYSTEM SET pgwasm.allowed_path_prefixes = '/srv/pgwasm';
SELECT pg_reload_conf();

GRANT pgwasm_loader TO app_runtime;

SELECT pgwasm.pgwasm_load(
    'arith',
    '{"path":"arith.component.wasm"}'::json,
    '{}'::json
);

SELECT * FROM pgwasm.pgwasm_functions();
SELECT * FROM pgwasm.pgwasm_modules();

SELECT pgwasm.pgwasm_unload('arith');
```

`pgwasm_load(module_name text, bytes_or_path json, options json)` accepts exactly one `bytes` or `path` source. File loading is off by default. A module name becomes the durable catalog key and the prefix for sanitized generated SQL function names.

### Lifecycle and Type Mapping

- `pgwasm_load` validates, resolves policy, creates required PostgreSQL types and functions, compiles an AOT artifact, and records the module.
- `pgwasm_reload` replaces module bytes while preserving stable identities when signatures remain compatible.
- `pgwasm_reconfigure` narrows or changes policy and resource limits.
- `pgwasm_unload` removes generated functions, types, catalog rows, and artifacts; dependencies block removal unless cascade is explicitly selected.
- WIT records map to composite types, enums to PostgreSQL enums, lists to arrays or `bytea`, and supported variants, flags, options, results, and resources to documented PostgreSQL representations.
- `pgwasm_modules()`, `pgwasm_functions()`, `pgwasm_wit_types()`, `pgwasm_policy_effective()`, and `pgwasm_stats()` provide inspection.

Review the generated function signatures before granting execute privileges or calling a newly loaded component. Reloads with breaking WIT changes require an explicit policy decision and dependency review.

### Sandbox and Privileges

The extension creates `pgwasm_loader` for lifecycle mutation and `pgwasm_reader` for observability. Loading, reloading, reconfiguring, and unloading require superuser or loader-role membership.

WASI filesystem, environment, sockets, HTTP, and SPI host-query access are all disabled by default. An administrator sets the cluster ceiling through `pgwasm.*` GUCs; per-module options can narrow that ceiling but cannot broaden it. Keep `pgwasm.allowed_hosts`, path prefixes, and filesystem preopens explicit and minimal.

### Resource and Operational Boundaries

- The default module-size limit is 32 MiB, invocation memory is 1,024 WebAssembly pages, and the wall-clock deadline is 5 seconds. Fuel metering is available but off by default.
- Artifacts under `$PGDATA/pgwasm/<module_id>/` are derived from module bytes and the Wasmtime build. Recompile them after incompatible Wasmtime or PostgreSQL upgrades rather than copying them as authoritative data.
- Shared counters depend on postmaster-time shared-memory allocation. Preload `pgwasm` when shared metrics are required; otherwise observability can fall back to non-shared counters and reports that state.
- The source exposes build features for PostgreSQL 13 through 18 and defaults to PostgreSQL 17, but the pinned revision has no published support matrix. Validate the exact PostgreSQL-major build and all required WIT mappings before deployment.
- Treat guest code as privileged database-adjacent code even with sandboxing: limit who can load modules, bound every capability and resource, and test traps, cancellation, reload, restart, and rollback behavior.
