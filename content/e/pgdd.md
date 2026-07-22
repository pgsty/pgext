---
title: "pgdd"
linkTitle: "pgdd"
description: "Introspect pg data dictionary via standard SQL"
weight: 5130
categories: ["ADMIN"]
width: full
---

[**pgdd**](https://github.com/rustprooflabs/pgdd) : Introspect pg data dictionary via standard SQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5130** | {{< badge content="pgdd" link="https://github.com/rustprooflabs/pgdd" >}} | {{< ext "pgdd" >}} | `0.6.1` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `dd` |
|   **See Also**    | {{< ext "pg_catcheck" >}} {{< ext "pg_orphaned" >}} {{< ext "pg_checksums" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgdd` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "pgdd_18" "green" >}} {{< bg "17" "pgdd_17" "green" >}} {{< bg "16" "pgdd_16" "green" >}} {{< bg "15" "pgdd_15" "green" >}} {{< bg "14" "pgdd_14" "green" >}} | `pgdd_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "postgresql-18-pgdd" "green" >}} {{< bg "17" "postgresql-17-pgdd" "green" >}} {{< bg "16" "postgresql-16-pgdd" "green" >}} {{< bg "15" "postgresql-15-pgdd" "green" >}} {{< bg "14" "postgresql-14-pgdd" "green" >}} | `postgresql-$v-pgdd` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_18` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 846.3 KiB | [pgdd_18-0.6.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_18-0.6.1-3PIGSTY.el8.x86_64.rpm) |
| `pgdd_18` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 756.7 KiB | [pgdd_18-0.6.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_18-0.6.1-3PIGSTY.el8.aarch64.rpm) |
| `pgdd_18` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 850.9 KiB | [pgdd_18-0.6.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_18-0.6.1-3PIGSTY.el9.x86_64.rpm) |
| `pgdd_18` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 803.1 KiB | [pgdd_18-0.6.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_18-0.6.1-3PIGSTY.el9.aarch64.rpm) |
| `pgdd_18` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 851.2 KiB | [pgdd_18-0.6.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_18-0.6.1-3PIGSTY.el10.x86_64.rpm) |
| `pgdd_18` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 781.8 KiB | [pgdd_18-0.6.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_18-0.6.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 673.0 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 561.8 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 672.1 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 562.1 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 746.2 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 664.2 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 739.9 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 654.5 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 736.3 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 654.3 KiB | [postgresql-18-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_17` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 843.2 KiB | [pgdd_17-0.6.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_17-0.6.1-3PIGSTY.el8.x86_64.rpm) |
| `pgdd_17` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 754.1 KiB | [pgdd_17-0.6.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_17-0.6.1-3PIGSTY.el8.aarch64.rpm) |
| `pgdd_17` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 847.2 KiB | [pgdd_17-0.6.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_17-0.6.1-3PIGSTY.el9.x86_64.rpm) |
| `pgdd_17` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 800.3 KiB | [pgdd_17-0.6.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_17-0.6.1-3PIGSTY.el9.aarch64.rpm) |
| `pgdd_17` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 847.3 KiB | [pgdd_17-0.6.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_17-0.6.1-3PIGSTY.el10.x86_64.rpm) |
| `pgdd_17` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 781.1 KiB | [pgdd_17-0.6.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_17-0.6.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 670.0 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 560.0 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 670.2 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 560.5 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 743.8 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 663.2 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 738.0 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 653.2 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 731.2 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 652.0 KiB | [postgresql-17-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_16` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 841.2 KiB | [pgdd_16-0.6.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_16-0.6.1-3PIGSTY.el8.x86_64.rpm) |
| `pgdd_16` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 752.4 KiB | [pgdd_16-0.6.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_16-0.6.1-3PIGSTY.el8.aarch64.rpm) |
| `pgdd_16` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 846.6 KiB | [pgdd_16-0.6.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_16-0.6.1-3PIGSTY.el9.x86_64.rpm) |
| `pgdd_16` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 798.4 KiB | [pgdd_16-0.6.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_16-0.6.1-3PIGSTY.el9.aarch64.rpm) |
| `pgdd_16` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 846.8 KiB | [pgdd_16-0.6.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_16-0.6.1-3PIGSTY.el10.x86_64.rpm) |
| `pgdd_16` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 778.8 KiB | [pgdd_16-0.6.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_16-0.6.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 669.6 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 559.5 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 669.0 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 559.9 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 744.9 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 662.4 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 736.8 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 652.6 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 732.9 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 650.7 KiB | [postgresql-16-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_15` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 831.8 KiB | [pgdd_15-0.6.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_15-0.6.1-3PIGSTY.el8.x86_64.rpm) |
| `pgdd_15` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 743.6 KiB | [pgdd_15-0.6.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_15-0.6.1-3PIGSTY.el8.aarch64.rpm) |
| `pgdd_15` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 837.6 KiB | [pgdd_15-0.6.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_15-0.6.1-3PIGSTY.el9.x86_64.rpm) |
| `pgdd_15` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 787.7 KiB | [pgdd_15-0.6.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_15-0.6.1-3PIGSTY.el9.aarch64.rpm) |
| `pgdd_15` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 836.1 KiB | [pgdd_15-0.6.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_15-0.6.1-3PIGSTY.el10.x86_64.rpm) |
| `pgdd_15` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 775.8 KiB | [pgdd_15-0.6.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_15-0.6.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 660.8 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 555.3 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 663.5 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 555.1 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 737.8 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 657.5 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 730.4 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 646.4 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 724.9 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 645.0 KiB | [postgresql-15-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_14` | `0.6.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 829.3 KiB | [pgdd_14-0.6.1-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_14-0.6.1-3PIGSTY.el8.x86_64.rpm) |
| `pgdd_14` | `0.6.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 740.9 KiB | [pgdd_14-0.6.1-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_14-0.6.1-3PIGSTY.el8.aarch64.rpm) |
| `pgdd_14` | `0.6.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 834.0 KiB | [pgdd_14-0.6.1-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_14-0.6.1-3PIGSTY.el9.x86_64.rpm) |
| `pgdd_14` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 784.8 KiB | [pgdd_14-0.6.1-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_14-0.6.1-3PIGSTY.el9.aarch64.rpm) |
| `pgdd_14` | `0.6.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 833.7 KiB | [pgdd_14-0.6.1-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_14-0.6.1-3PIGSTY.el10.x86_64.rpm) |
| `pgdd_14` | `0.6.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 775.9 KiB | [pgdd_14-0.6.1-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_14-0.6.1-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 660.9 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 553.5 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 660.8 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 554.0 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 735.3 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 656.7 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 727.8 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 645.3 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 724.0 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 642.9 KiB | [postgresql-14-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rustprooflabs/pgdd" title="Repository" icon="github" subtitle="github.com/rustprooflabs/pgdd" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgdd-0.6.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgdd;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgdd;		# install via package name, for the active PG version

pig install pgdd -v 18;   # install for PG 18
pig install pgdd -v 17;   # install for PG 17
pig install pgdd -v 16;   # install for PG 16
pig install pgdd -v 15;   # install for PG 15
pig install pgdd -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgdd;
```




## Usage

> [pgdd: Introspect pg data dictionary via standard SQL](https://github.com/rustprooflabs/pgdd)

PgDD provides data dictionary views in the `dd` schema for introspecting database objects via standard SQL.

### Database Overview

```sql
SELECT * FROM dd.database;
```

Returns: `db_name`, `db_size`, `schema_count`, `table_count`, `size_in_tables`, `view_count`, `size_in_views`, `extension_count`.

### Schemas

```sql
SELECT s_name, table_count, view_count, function_count, size_plus_indexes, description
  FROM dd.schemas;
```

### Tables

```sql
SELECT t_name, size_pretty, rows, bytes_per_row
  FROM dd.tables
  WHERE s_name = 'public';
```

### Views

```sql
SELECT s_name, v_name, description FROM dd.views;
```

### Columns

```sql
SELECT source_type, s_name, t_name, c_name, data_type
  FROM dd.columns
  WHERE data_type LIKE 'int%';
```

### Functions

```sql
SELECT s_name, f_name, argument_data_types, result_data_types FROM dd.functions;
```

### Partitioned Tables

```sql
SELECT * FROM dd.partition_parents WHERE s_name = 'public';
SELECT * FROM dd.partition_children WHERE s_name = 'public';
```

The `partition_parents` view shows aggregate partition stats (count, total size, total rows). The `partition_children` view shows per-partition details with percentage calculations against the parent.

System objects are excluded by default. To include them, query the underlying functions directly: `SELECT * FROM dd.tables() WHERE system_object;`
