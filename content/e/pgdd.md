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
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_catcheck" >}} {{< ext "pg_orphaned" >}} {{< ext "pg_checksums" >}} |

> [!Note] manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgdd` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "pgdd_18" "green" >}} {{< bg "17" "pgdd_17" "green" >}} {{< bg "16" "pgdd_16" "green" >}} {{< bg "15" "pgdd_15" "green" >}} {{< bg "14" "pgdd_14" "green" >}} {{< bg "13" "pgdd_13" "green" >}} | `pgdd_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.1` | {{< bg "18" "postgresql-18-pgdd" "green" >}} {{< bg "17" "postgresql-17-pgdd" "green" >}} {{< bg "16" "postgresql-16-pgdd" "green" >}} {{< bg "15" "postgresql-15-pgdd" "green" >}} {{< bg "14" "postgresql-14-pgdd" "green" >}} {{< bg "13" "postgresql-13-pgdd" "green" >}} | `postgresql-$v-pgdd` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.0" "pgdd_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-13-pgdd : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-18-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-17-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-16-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-15-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-14-pgdd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.1" "postgresql-13-pgdd : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_18` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 286.1 KiB | [pgdd_18-0.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_18-0.6.0-2PIGSTY.el8.x86_64.rpm) |
| `pgdd_18` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 183.1 KiB | [pgdd_18-0.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_18-0.6.0-2PIGSTY.el8.aarch64.rpm) |
| `pgdd_18` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 300.2 KiB | [pgdd_18-0.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_18-0.6.0-2PIGSTY.el9.x86_64.rpm) |
| `pgdd_18` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 194.3 KiB | [pgdd_18-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_18-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pgdd_18` | `0.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 300.2 KiB | [pgdd_18-0.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_18-0.6.0-2PIGSTY.el10.x86_64.rpm) |
| `pgdd_18` | `0.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 194.9 KiB | [pgdd_18-0.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_18-0.6.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 236.1 KiB | [postgresql-18-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 140.6 KiB | [postgresql-18-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 235.9 KiB | [postgresql-18-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.6 KiB | [postgresql-18-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 265.9 KiB | [postgresql-18-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 162.7 KiB | [postgresql-18-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 263.1 KiB | [postgresql-18-pgdd_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 161.5 KiB | [postgresql-18-pgdd_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-18-pgdd_0.6.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_17` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 286.6 KiB | [pgdd_17-0.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_17-0.6.0-2PIGSTY.el8.x86_64.rpm) |
| `pgdd_17` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 183.7 KiB | [pgdd_17-0.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_17-0.6.0-2PIGSTY.el8.aarch64.rpm) |
| `pgdd_17` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 300.5 KiB | [pgdd_17-0.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_17-0.6.0-2PIGSTY.el9.x86_64.rpm) |
| `pgdd_17` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 194.4 KiB | [pgdd_17-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_17-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pgdd_17` | `0.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 300.8 KiB | [pgdd_17-0.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_17-0.6.0-2PIGSTY.el10.x86_64.rpm) |
| `pgdd_17` | `0.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 194.9 KiB | [pgdd_17-0.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_17-0.6.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 235.7 KiB | [postgresql-17-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 140.6 KiB | [postgresql-17-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 235.5 KiB | [postgresql-17-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.6 KiB | [postgresql-17-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 265.9 KiB | [postgresql-17-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 163.0 KiB | [postgresql-17-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 263.0 KiB | [postgresql-17-pgdd_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 161.6 KiB | [postgresql-17-pgdd_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-17-pgdd_0.6.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_16` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 286.6 KiB | [pgdd_16-0.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_16-0.6.0-2PIGSTY.el8.x86_64.rpm) |
| `pgdd_16` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 183.7 KiB | [pgdd_16-0.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_16-0.6.0-2PIGSTY.el8.aarch64.rpm) |
| `pgdd_16` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 300.4 KiB | [pgdd_16-0.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_16-0.6.0-2PIGSTY.el9.x86_64.rpm) |
| `pgdd_16` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 194.6 KiB | [pgdd_16-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_16-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pgdd_16` | `0.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 300.7 KiB | [pgdd_16-0.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_16-0.6.0-2PIGSTY.el10.x86_64.rpm) |
| `pgdd_16` | `0.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 194.8 KiB | [pgdd_16-0.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_16-0.6.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 235.7 KiB | [postgresql-16-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 140.6 KiB | [postgresql-16-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 235.6 KiB | [postgresql-16-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.6 KiB | [postgresql-16-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 265.9 KiB | [postgresql-16-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 162.7 KiB | [postgresql-16-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 263.0 KiB | [postgresql-16-pgdd_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 161.7 KiB | [postgresql-16-pgdd_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-16-pgdd_0.6.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_15` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 286.2 KiB | [pgdd_15-0.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_15-0.6.0-2PIGSTY.el8.x86_64.rpm) |
| `pgdd_15` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 183.6 KiB | [pgdd_15-0.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_15-0.6.0-2PIGSTY.el8.aarch64.rpm) |
| `pgdd_15` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 300.4 KiB | [pgdd_15-0.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_15-0.6.0-2PIGSTY.el9.x86_64.rpm) |
| `pgdd_15` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 194.3 KiB | [pgdd_15-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_15-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pgdd_15` | `0.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 300.4 KiB | [pgdd_15-0.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_15-0.6.0-2PIGSTY.el10.x86_64.rpm) |
| `pgdd_15` | `0.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 194.7 KiB | [pgdd_15-0.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_15-0.6.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 235.6 KiB | [postgresql-15-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 140.6 KiB | [postgresql-15-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 235.4 KiB | [postgresql-15-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.7 KiB | [postgresql-15-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 266.0 KiB | [postgresql-15-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 162.9 KiB | [postgresql-15-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 263.0 KiB | [postgresql-15-pgdd_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 161.8 KiB | [postgresql-15-pgdd_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-15-pgdd_0.6.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_14` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 286.2 KiB | [pgdd_14-0.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_14-0.6.0-2PIGSTY.el8.x86_64.rpm) |
| `pgdd_14` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 183.7 KiB | [pgdd_14-0.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_14-0.6.0-2PIGSTY.el8.aarch64.rpm) |
| `pgdd_14` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 300.4 KiB | [pgdd_14-0.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_14-0.6.0-2PIGSTY.el9.x86_64.rpm) |
| `pgdd_14` | `0.6.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 194.4 KiB | [pgdd_14-0.6.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_14-0.6.1-1PIGSTY.el9.aarch64.rpm) |
| `pgdd_14` | `0.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 300.0 KiB | [pgdd_14-0.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_14-0.6.0-2PIGSTY.el10.x86_64.rpm) |
| `pgdd_14` | `0.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 194.8 KiB | [pgdd_14-0.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_14-0.6.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 235.3 KiB | [postgresql-14-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 140.6 KiB | [postgresql-14-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 234.8 KiB | [postgresql-14-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.7 KiB | [postgresql-14-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 265.6 KiB | [postgresql-14-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 162.9 KiB | [postgresql-14-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 262.9 KiB | [postgresql-14-pgdd_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 161.8 KiB | [postgresql-14-pgdd_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-14-pgdd_0.6.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgdd_13` | `0.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 286.8 KiB | [pgdd_13-0.6.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgdd_13-0.6.0-2PIGSTY.el8.x86_64.rpm) |
| `pgdd_13` | `0.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 183.7 KiB | [pgdd_13-0.6.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgdd_13-0.6.0-2PIGSTY.el8.aarch64.rpm) |
| `pgdd_13` | `0.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 300.9 KiB | [pgdd_13-0.6.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgdd_13-0.6.0-2PIGSTY.el9.x86_64.rpm) |
| `pgdd_13` | `0.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 195.8 KiB | [pgdd_13-0.6.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgdd_13-0.6.0-2PIGSTY.el9.aarch64.rpm) |
| `pgdd_13` | `0.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 300.8 KiB | [pgdd_13-0.6.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgdd_13-0.6.0-2PIGSTY.el10.x86_64.rpm) |
| `pgdd_13` | `0.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 194.8 KiB | [pgdd_13-0.6.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgdd_13-0.6.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgdd` | `0.6.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.1 KiB | [postgresql-13-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-13-pgdd_0.6.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgdd` | `0.6.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.1 KiB | [postgresql-13-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgdd/postgresql-13-pgdd_0.6.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgdd` | `0.6.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.1 KiB | [postgresql-13-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-13-pgdd_0.6.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgdd` | `0.6.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.1 KiB | [postgresql-13-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgdd/postgresql-13-pgdd_0.6.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgdd` | `0.6.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 KiB | [postgresql-13-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-13-pgdd_0.6.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgdd` | `0.6.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.0 KiB | [postgresql-13-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgdd/postgresql-13-pgdd_0.6.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgdd` | `0.6.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.0 KiB | [postgresql-13-pgdd_0.6.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-13-pgdd_0.6.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgdd` | `0.6.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.0 KiB | [postgresql-13-pgdd_0.6.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgdd/postgresql-13-pgdd_0.6.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rustprooflabs/pgdd" title="Repository" icon="github" subtitle="github.com/rustprooflabs/pgdd" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgdd-0.6.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgdd;		# build rpm / deb with pig
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
pig install pgdd -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgdd;
```
