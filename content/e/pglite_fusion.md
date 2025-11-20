---
title: "pglite_fusion"
linkTitle: "pglite_fusion"
description: "Embed an SQLite database in your PostgreSQL table"
weight: 3540
categories: ["TYPE"]
width: full
---

[**pglite_fusion**](https://github.com/frectonz/pglite-fusion) : Embed an SQLite database in your PostgreSQL table


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3540** | {{< badge content="pglite_fusion" link="https://github.com/frectonz/pglite-fusion" >}} | {{< ext "pglite_fusion" >}} | `0.0.6` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "duckdb_fdw" >}} {{< ext "sqlite_fdw" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pglite_fusion` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.6` | {{< bg "18" "pglite_fusion_18" "green" >}} {{< bg "17" "pglite_fusion_17" "green" >}} {{< bg "16" "pglite_fusion_16" "green" >}} {{< bg "15" "pglite_fusion_15" "green" >}} {{< bg "14" "pglite_fusion_14" "green" >}} {{< bg "13" "pglite_fusion_13" "green" >}} | `pglite_fusion_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.6` | {{< bg "18" "postgresql-18-pglite-fusion" "green" >}} {{< bg "17" "postgresql-17-pglite-fusion" "green" >}} {{< bg "16" "postgresql-16-pglite-fusion" "green" >}} {{< bg "15" "postgresql-15-pglite-fusion" "green" >}} {{< bg "14" "postgresql-14-pglite-fusion" "green" >}} {{< bg "13" "postgresql-13-pglite-fusion" "green" >}} | `postgresql-$v-pglite-fusion` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "pglite_fusion_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-18-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-17-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-16-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-15-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-14-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-13-pglite-fusion : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-18-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-17-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-16-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-15-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-14-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-13-pglite-fusion : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-18-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-17-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-16-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-15-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-14-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-13-pglite-fusion : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-18-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-17-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-16-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-15-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-14-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-13-pglite-fusion : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-18-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-17-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-16-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-15-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-14-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-13-pglite-fusion : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-18-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-17-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-16-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-15-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-14-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-13-pglite-fusion : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-18-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-17-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-16-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-15-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-14-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-13-pglite-fusion : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-18-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-17-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-16-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-15-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-14-pglite-fusion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.6" "postgresql-13-pglite-fusion : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglite_fusion_18` | `0.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.2 MiB | [pglite_fusion_18-0.0.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglite_fusion_18-0.0.6-2PIGSTY.el8.x86_64.rpm) |
| `pglite_fusion_18` | `0.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_18-0.0.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglite_fusion_18-0.0.6-2PIGSTY.el8.aarch64.rpm) |
| `pglite_fusion_18` | `0.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pglite_fusion_18-0.0.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglite_fusion_18-0.0.6-2PIGSTY.el9.x86_64.rpm) |
| `pglite_fusion_18` | `0.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pglite_fusion_18-0.0.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglite_fusion_18-0.0.6-2PIGSTY.el9.aarch64.rpm) |
| `pglite_fusion_18` | `0.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pglite_fusion_18-0.0.6-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglite_fusion_18-0.0.6-2PIGSTY.el10.x86_64.rpm) |
| `pglite_fusion_18` | `0.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_18-0.0.6-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglite_fusion_18-0.0.6-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pglite-fusion` | `0.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.0 KiB | [postgresql-18-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-18-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pglite-fusion` | `0.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.0 KiB | [postgresql-18-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-18-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pglite-fusion` | `0.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.0 KiB | [postgresql-18-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-18-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pglite-fusion` | `0.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.0 KiB | [postgresql-18-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-18-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pglite-fusion` | `0.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 KiB | [postgresql-18-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-18-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pglite-fusion` | `0.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 KiB | [postgresql-18-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-18-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pglite-fusion` | `0.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 KiB | [postgresql-18-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-18-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pglite-fusion` | `0.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 KiB | [postgresql-18-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-18-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglite_fusion_17` | `0.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.2 MiB | [pglite_fusion_17-0.0.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglite_fusion_17-0.0.6-2PIGSTY.el8.x86_64.rpm) |
| `pglite_fusion_17` | `0.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_17-0.0.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglite_fusion_17-0.0.6-2PIGSTY.el8.aarch64.rpm) |
| `pglite_fusion_17` | `0.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pglite_fusion_17-0.0.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglite_fusion_17-0.0.6-2PIGSTY.el9.x86_64.rpm) |
| `pglite_fusion_17` | `0.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pglite_fusion_17-0.0.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglite_fusion_17-0.0.6-2PIGSTY.el9.aarch64.rpm) |
| `pglite_fusion_17` | `0.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pglite_fusion_17-0.0.6-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglite_fusion_17-0.0.6-2PIGSTY.el10.x86_64.rpm) |
| `pglite_fusion_17` | `0.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_17-0.0.6-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglite_fusion_17-0.0.6-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pglite-fusion` | `0.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-17-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-17-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pglite-fusion` | `0.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 952.6 KiB | [postgresql-17-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-17-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pglite-fusion` | `0.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-17-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-17-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pglite-fusion` | `0.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 976.1 KiB | [postgresql-17-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-17-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pglite-fusion` | `0.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-17-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-17-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pglite-fusion` | `0.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-17-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-17-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pglite-fusion` | `0.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-17-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-17-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pglite-fusion` | `0.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-17-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-17-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglite_fusion_16` | `0.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.2 MiB | [pglite_fusion_16-0.0.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglite_fusion_16-0.0.6-2PIGSTY.el8.x86_64.rpm) |
| `pglite_fusion_16` | `0.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_16-0.0.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglite_fusion_16-0.0.6-2PIGSTY.el8.aarch64.rpm) |
| `pglite_fusion_16` | `0.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pglite_fusion_16-0.0.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglite_fusion_16-0.0.6-2PIGSTY.el9.x86_64.rpm) |
| `pglite_fusion_16` | `0.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pglite_fusion_16-0.0.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglite_fusion_16-0.0.6-2PIGSTY.el9.aarch64.rpm) |
| `pglite_fusion_16` | `0.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pglite_fusion_16-0.0.6-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglite_fusion_16-0.0.6-2PIGSTY.el10.x86_64.rpm) |
| `pglite_fusion_16` | `0.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_16-0.0.6-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglite_fusion_16-0.0.6-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pglite-fusion` | `0.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-16-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-16-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pglite-fusion` | `0.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 952.7 KiB | [postgresql-16-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-16-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pglite-fusion` | `0.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-16-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-16-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pglite-fusion` | `0.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 975.8 KiB | [postgresql-16-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-16-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pglite-fusion` | `0.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-16-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-16-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pglite-fusion` | `0.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-16-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-16-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pglite-fusion` | `0.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-16-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-16-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pglite-fusion` | `0.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-16-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-16-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglite_fusion_15` | `0.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.2 MiB | [pglite_fusion_15-0.0.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglite_fusion_15-0.0.6-2PIGSTY.el8.x86_64.rpm) |
| `pglite_fusion_15` | `0.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_15-0.0.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglite_fusion_15-0.0.6-2PIGSTY.el8.aarch64.rpm) |
| `pglite_fusion_15` | `0.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pglite_fusion_15-0.0.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglite_fusion_15-0.0.6-2PIGSTY.el9.x86_64.rpm) |
| `pglite_fusion_15` | `0.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pglite_fusion_15-0.0.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglite_fusion_15-0.0.6-2PIGSTY.el9.aarch64.rpm) |
| `pglite_fusion_15` | `0.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pglite_fusion_15-0.0.6-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglite_fusion_15-0.0.6-2PIGSTY.el10.x86_64.rpm) |
| `pglite_fusion_15` | `0.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_15-0.0.6-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglite_fusion_15-0.0.6-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pglite-fusion` | `0.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-15-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-15-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pglite-fusion` | `0.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 952.5 KiB | [postgresql-15-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-15-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pglite-fusion` | `0.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-15-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-15-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pglite-fusion` | `0.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 975.8 KiB | [postgresql-15-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-15-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pglite-fusion` | `0.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-15-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-15-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pglite-fusion` | `0.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-15-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pglite-fusion` | `0.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-15-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-15-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pglite-fusion` | `0.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-15-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-15-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglite_fusion_14` | `0.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.2 MiB | [pglite_fusion_14-0.0.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglite_fusion_14-0.0.6-2PIGSTY.el8.x86_64.rpm) |
| `pglite_fusion_14` | `0.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_14-0.0.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglite_fusion_14-0.0.6-2PIGSTY.el8.aarch64.rpm) |
| `pglite_fusion_14` | `0.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pglite_fusion_14-0.0.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglite_fusion_14-0.0.6-2PIGSTY.el9.x86_64.rpm) |
| `pglite_fusion_14` | `0.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pglite_fusion_14-0.0.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglite_fusion_14-0.0.6-2PIGSTY.el9.aarch64.rpm) |
| `pglite_fusion_14` | `0.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pglite_fusion_14-0.0.6-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglite_fusion_14-0.0.6-2PIGSTY.el10.x86_64.rpm) |
| `pglite_fusion_14` | `0.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_14-0.0.6-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglite_fusion_14-0.0.6-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pglite-fusion` | `0.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-14-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-14-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pglite-fusion` | `0.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 952.7 KiB | [postgresql-14-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-14-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pglite-fusion` | `0.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-14-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-14-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pglite-fusion` | `0.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 975.8 KiB | [postgresql-14-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-14-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pglite-fusion` | `0.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-14-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-14-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pglite-fusion` | `0.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-14-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pglite-fusion` | `0.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-14-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-14-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pglite-fusion` | `0.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-14-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-14-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglite_fusion_13` | `0.0.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.2 MiB | [pglite_fusion_13-0.0.6-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglite_fusion_13-0.0.6-2PIGSTY.el8.x86_64.rpm) |
| `pglite_fusion_13` | `0.0.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_13-0.0.6-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglite_fusion_13-0.0.6-2PIGSTY.el8.aarch64.rpm) |
| `pglite_fusion_13` | `0.0.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [pglite_fusion_13-0.0.6-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglite_fusion_13-0.0.6-2PIGSTY.el9.x86_64.rpm) |
| `pglite_fusion_13` | `0.0.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pglite_fusion_13-0.0.6-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglite_fusion_13-0.0.6-2PIGSTY.el9.aarch64.rpm) |
| `pglite_fusion_13` | `0.0.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pglite_fusion_13-0.0.6-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglite_fusion_13-0.0.6-2PIGSTY.el10.x86_64.rpm) |
| `pglite_fusion_13` | `0.0.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pglite_fusion_13-0.0.6-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglite_fusion_13-0.0.6-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pglite-fusion` | `0.0.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.1 MiB | [postgresql-13-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-13-pglite-fusion_0.0.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pglite-fusion` | `0.0.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 952.9 KiB | [postgresql-13-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglite-fusion/postgresql-13-pglite-fusion_0.0.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pglite-fusion` | `0.0.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-13-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-13-pglite-fusion_0.0.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pglite-fusion` | `0.0.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 975.7 KiB | [postgresql-13-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglite-fusion/postgresql-13-pglite-fusion_0.0.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pglite-fusion` | `0.0.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.2 MiB | [postgresql-13-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-13-pglite-fusion_0.0.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pglite-fusion` | `0.0.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-13-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglite-fusion/postgresql-13-pglite-fusion_0.0.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pglite-fusion` | `0.0.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.2 MiB | [postgresql-13-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-13-pglite-fusion_0.0.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pglite-fusion` | `0.0.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.1 MiB | [postgresql-13-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglite-fusion/postgresql-13-pglite-fusion_0.0.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/frectonz/pglite-fusion" title="Repository" icon="github" subtitle="github.com/frectonz/pglite-fusion" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglite-fusion-0.0.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pglite_fusion;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pglite_fusion;		# install via package name, for the active PG version

pig install pglite_fusion -v 18;   # install for PG 18
pig install pglite_fusion -v 17;   # install for PG 17
pig install pglite_fusion -v 16;   # install for PG 16
pig install pglite_fusion -v 15;   # install for PG 15
pig install pglite_fusion -v 14;   # install for PG 14
pig install pglite_fusion -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pglite_fusion';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pglite_fusion;
```


## Usage

> https://github.com/frectonz/pglite-fusion/blob/main/README.md


Here's some demo usage.

```sql
-- Load PG extension
CREATE EXTENSION pglite_fusion;

-- Create a table with an SQLite column
CREATE TABLE people (
                        name     TEXT NOT NULL,
                        database SQLITE DEFAULT init_sqlite('CREATE TABLE todos (task TEXT)')
);

-- Insert a row into the people table
INSERT INTO people VALUES ('frectonz');

-- Create a todo for "frectonz"
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''solve multitenancy'')'
               )
WHERE name = 'frectonz';

-- Create a todo for "frectonz"
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''buy milk'')'
               )
WHERE name = 'frectonz';

-- Fetch frectonz's info
SELECT
    name,
    (
        SELECT json_agg(get_sqlite_text(sqlite_row, 0))
        FROM query_sqlite(
                database,
                'SELECT * FROM todos'
             )
    ) AS todos
FROM
    people
WHERE
    name = 'frectonz';
```


--------

## API Doc


### `empty_sqlite`

Creates an empty SQLite database and returns it as a binary object. This can be used to initialize an empty SQLite database in a PostgreSQL column.

#### Example Usage:

```sql
SELECT empty_sqlite();
```

-----

### `query_sqlite`

Executes a SQL query on a SQLite database stored as a binary object and returns the result as a table of JSON-encoded rows. This function is useful for querying SQLite databases stored in PostgreSQL columns.

#### Parameters:

- `sqlite`: The SQLite database to query, stored as a binary object.
- `query`: The SQL query string to execute on the SQLite database.

#### Example Usage:

```sql
SELECT * FROM query_sqlite(
        database,
        'SELECT * FROM todos'
              );
```

-----

### `execute_sqlite`

Executes a SQL statement (such as `INSERT`, `UPDATE`, or `DELETE`) on a SQLite database stored as a binary object. The updated SQLite database is returned as a binary object, allowing further operations on it.

#### Parameters:

- `sqlite`: The SQLite database to execute the SQL query on, stored as a binary object.
- `query`: The SQL statement to execute on the SQLite database.

##### Example Usage:

```sql
UPDATE people
SET database = execute_sqlite(
        database,
        'INSERT INTO todos VALUES (''solve multitenancy'')'
               )
WHERE name = 'frectonz';
```

-----

### `init_sqlite`

Creates an SQLite database with an initialization query already applied on it. This can be used to initialize a SQLite database with the expected tables already created.

#### Parameters:

- `query`: The SQL statement to execute on the SQLite database.

##### Example Usage:

```sql

CREATE TABLE people (
                        name     TEXT NOT NULL,
                        database SQLITE DEFAULT init_sqlite('CREATE TABLE todos (task TEXT)')
);
```

-----

### `get_sqlite_text`
Extracts a text value from a specific column in a row returned by `query_sqlite`. Use this function to retrieve text values from query results.

#### Parameters:

- `sqlite_row`: A row from the results of `query_sqlite`.
- `index`: The index of the column to extract from the row.

#### Example Usage:

```sql
SELECT get_sqlite_text(sqlite_row, 0)
FROM query_sqlite(database, 'SELECT * FROM todos');
```

----

### `get_sqlite_integer`

Extracts an integer value from a specific column in a row returned by `query_sqlite`. Use this function to retrieve integer values from query results.

#### Parameters:

- `sqlite_row`: A row from the results of `query_sqlite`.
- `index`: The index of the column to extract from the row.

#### Example Usage:

```sql
SELECT get_sqlite_integer(sqlite_row, 1)
FROM query_sqlite(database, 'SELECT * FROM todos');
```

----

### `get_sqlite_real`

Extracts a real (floating-point) value from a specific column in a row returned by `query_sqlite`. Use this function to retrieve real number values from query results.

#### Parameters:

- `sqlite_row`: A row from the results of `query_sqlite`.
- `index`: The index of the column to extract from the row.

#### Example Usage:

```sql
SELECT get_sqlite_real(sqlite_row, 2)
FROM query_sqlite(database, 'SELECT * FROM todos');
```