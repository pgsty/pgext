---
title: "convert"
linkTitle: "convert"
description: "conversion functions for spatial, routing and other specialized uses"
weight: 4850
categories: ["FUNC"]
width: full
---

[**pg_convert**](https://github.com/rustprooflabs/convert) : conversion functions for spatial, routing and other specialized uses


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4850** | {{< badge content="convert" link="https://github.com/rustprooflabs/convert" >}} | {{< ext "convert" "pg_convert" >}} | `0.0.4` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "permuteseq" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |

> [!Note] manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_convert` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "pg_convert_18" "green" >}} {{< bg "17" "pg_convert_17" "green" >}} {{< bg "16" "pg_convert_16" "green" >}} {{< bg "15" "pg_convert_15" "green" >}} {{< bg "14" "pg_convert_14" "green" >}} {{< bg "13" "pg_convert_13" "green" >}} | `pg_convert_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.4` | {{< bg "18" "postgresql-18-convert" "green" >}} {{< bg "17" "postgresql-17-convert" "green" >}} {{< bg "16" "postgresql-16-convert" "green" >}} {{< bg "15" "postgresql-15-convert" "green" >}} {{< bg "14" "postgresql-14-convert" "green" >}} {{< bg "13" "postgresql-13-convert" "green" >}} | `postgresql-$v-convert` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "pg_convert_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-13-convert : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-13-convert : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-13-convert : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-13-convert : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-13-convert : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-13-convert : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-13-convert : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-14-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.4" "postgresql-13-convert : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_18` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 292.1 KiB | [pg_convert_18-0.0.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_18-0.0.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_convert_18` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.5 KiB | [pg_convert_18-0.0.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_18-0.0.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_convert_18` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 307.2 KiB | [pg_convert_18-0.0.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_18-0.0.4-2PIGSTY.el9.x86_64.rpm) |
| `pg_convert_18` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.1 KiB | [pg_convert_18-0.0.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_18-0.0.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_convert_18` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 309.5 KiB | [pg_convert_18-0.0.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_18-0.0.4-2PIGSTY.el10.x86_64.rpm) |
| `pg_convert_18` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 202.9 KiB | [pg_convert_18-0.0.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_18-0.0.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-convert` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 240.9 KiB | [postgresql-18-convert_0.0.4-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-18-convert_0.0.4-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-convert` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 145.9 KiB | [postgresql-18-convert_0.0.4-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-18-convert_0.0.4-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-convert` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.6 KiB | [postgresql-18-convert_0.0.4-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-18-convert_0.0.4-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-convert` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 146.0 KiB | [postgresql-18-convert_0.0.4-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-18-convert_0.0.4-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-convert` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 272.6 KiB | [postgresql-18-convert_0.0.4-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-18-convert_0.0.4-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-convert` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.4 KiB | [postgresql-18-convert_0.0.4-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-18-convert_0.0.4-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-convert` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 270.1 KiB | [postgresql-18-convert_0.0.4-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-18-convert_0.0.4-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-convert` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.4 KiB | [postgresql-18-convert_0.0.4-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-18-convert_0.0.4-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_17` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 292.1 KiB | [pg_convert_17-0.0.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_17-0.0.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_convert_17` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.4 KiB | [pg_convert_17-0.0.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_17-0.0.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_convert_17` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 307.2 KiB | [pg_convert_17-0.0.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_17-0.0.4-2PIGSTY.el9.x86_64.rpm) |
| `pg_convert_17` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.0 KiB | [pg_convert_17-0.0.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_17-0.0.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_convert_17` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 309.5 KiB | [pg_convert_17-0.0.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_17-0.0.4-2PIGSTY.el10.x86_64.rpm) |
| `pg_convert_17` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.0 KiB | [pg_convert_17-0.0.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_17-0.0.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-convert` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 240.2 KiB | [postgresql-17-convert_0.0.4-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-17-convert_0.0.4-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-convert` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 145.9 KiB | [postgresql-17-convert_0.0.4-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-17-convert_0.0.4-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-convert` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.3 KiB | [postgresql-17-convert_0.0.4-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-17-convert_0.0.4-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-convert` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 145.9 KiB | [postgresql-17-convert_0.0.4-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-17-convert_0.0.4-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-convert` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 272.5 KiB | [postgresql-17-convert_0.0.4-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-17-convert_0.0.4-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-convert` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.3 KiB | [postgresql-17-convert_0.0.4-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-17-convert_0.0.4-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-convert` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 270.1 KiB | [postgresql-17-convert_0.0.4-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-17-convert_0.0.4-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-convert` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.3 KiB | [postgresql-17-convert_0.0.4-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-17-convert_0.0.4-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_16` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 292.1 KiB | [pg_convert_16-0.0.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_16-0.0.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_convert_16` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.4 KiB | [pg_convert_16-0.0.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_16-0.0.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_convert_16` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 307.2 KiB | [pg_convert_16-0.0.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_16-0.0.4-2PIGSTY.el9.x86_64.rpm) |
| `pg_convert_16` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.0 KiB | [pg_convert_16-0.0.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_16-0.0.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_convert_16` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 309.5 KiB | [pg_convert_16-0.0.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_16-0.0.4-2PIGSTY.el10.x86_64.rpm) |
| `pg_convert_16` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.0 KiB | [pg_convert_16-0.0.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_16-0.0.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-convert` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 241.0 KiB | [postgresql-16-convert_0.0.4-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-16-convert_0.0.4-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-convert` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 145.8 KiB | [postgresql-16-convert_0.0.4-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-16-convert_0.0.4-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-convert` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.4 KiB | [postgresql-16-convert_0.0.4-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-16-convert_0.0.4-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-convert` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 145.9 KiB | [postgresql-16-convert_0.0.4-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-16-convert_0.0.4-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-convert` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 272.5 KiB | [postgresql-16-convert_0.0.4-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-16-convert_0.0.4-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-convert` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.4 KiB | [postgresql-16-convert_0.0.4-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-16-convert_0.0.4-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-convert` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 270.0 KiB | [postgresql-16-convert_0.0.4-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-16-convert_0.0.4-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-convert` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.4 KiB | [postgresql-16-convert_0.0.4-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-16-convert_0.0.4-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_15` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 292.0 KiB | [pg_convert_15-0.0.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_15-0.0.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_convert_15` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.5 KiB | [pg_convert_15-0.0.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_15-0.0.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_convert_15` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 307.3 KiB | [pg_convert_15-0.0.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_15-0.0.4-2PIGSTY.el9.x86_64.rpm) |
| `pg_convert_15` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 202.8 KiB | [pg_convert_15-0.0.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_15-0.0.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_convert_15` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 309.3 KiB | [pg_convert_15-0.0.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_15-0.0.4-2PIGSTY.el10.x86_64.rpm) |
| `pg_convert_15` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.0 KiB | [pg_convert_15-0.0.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_15-0.0.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-convert` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 240.2 KiB | [postgresql-15-convert_0.0.4-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-15-convert_0.0.4-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-convert` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 145.9 KiB | [postgresql-15-convert_0.0.4-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-15-convert_0.0.4-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-convert` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.6 KiB | [postgresql-15-convert_0.0.4-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-15-convert_0.0.4-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-convert` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 145.9 KiB | [postgresql-15-convert_0.0.4-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-15-convert_0.0.4-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-convert` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 272.0 KiB | [postgresql-15-convert_0.0.4-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-15-convert_0.0.4-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-convert` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.4 KiB | [postgresql-15-convert_0.0.4-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-15-convert_0.0.4-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-convert` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 270.0 KiB | [postgresql-15-convert_0.0.4-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-15-convert_0.0.4-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-convert` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.1 KiB | [postgresql-15-convert_0.0.4-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-15-convert_0.0.4-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_14` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 291.9 KiB | [pg_convert_14-0.0.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_14-0.0.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_convert_14` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.5 KiB | [pg_convert_14-0.0.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_14-0.0.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_convert_14` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 307.0 KiB | [pg_convert_14-0.0.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_14-0.0.4-2PIGSTY.el9.x86_64.rpm) |
| `pg_convert_14` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.1 KiB | [pg_convert_14-0.0.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_14-0.0.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_convert_14` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 309.0 KiB | [pg_convert_14-0.0.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_14-0.0.4-2PIGSTY.el10.x86_64.rpm) |
| `pg_convert_14` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.0 KiB | [pg_convert_14-0.0.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_14-0.0.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-convert` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 240.3 KiB | [postgresql-14-convert_0.0.4-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-14-convert_0.0.4-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-convert` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 145.9 KiB | [postgresql-14-convert_0.0.4-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-14-convert_0.0.4-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-convert` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.1 KiB | [postgresql-14-convert_0.0.4-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-14-convert_0.0.4-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-convert` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 145.9 KiB | [postgresql-14-convert_0.0.4-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-14-convert_0.0.4-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-convert` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 272.3 KiB | [postgresql-14-convert_0.0.4-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-14-convert_0.0.4-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-convert` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.7 KiB | [postgresql-14-convert_0.0.4-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-14-convert_0.0.4-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-convert` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 269.9 KiB | [postgresql-14-convert_0.0.4-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-14-convert_0.0.4-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-convert` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.3 KiB | [postgresql-14-convert_0.0.4-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-14-convert_0.0.4-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_13` | `0.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 292.2 KiB | [pg_convert_13-0.0.4-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_13-0.0.4-2PIGSTY.el8.x86_64.rpm) |
| `pg_convert_13` | `0.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.4 KiB | [pg_convert_13-0.0.4-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_13-0.0.4-2PIGSTY.el8.aarch64.rpm) |
| `pg_convert_13` | `0.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 307.7 KiB | [pg_convert_13-0.0.4-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_13-0.0.4-2PIGSTY.el9.x86_64.rpm) |
| `pg_convert_13` | `0.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.1 KiB | [pg_convert_13-0.0.4-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_13-0.0.4-2PIGSTY.el9.aarch64.rpm) |
| `pg_convert_13` | `0.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 309.8 KiB | [pg_convert_13-0.0.4-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_13-0.0.4-2PIGSTY.el10.x86_64.rpm) |
| `pg_convert_13` | `0.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 202.9 KiB | [pg_convert_13-0.0.4-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_13-0.0.4-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-convert` | `0.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 240.9 KiB | [postgresql-13-convert_0.0.4-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-13-convert_0.0.4-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-convert` | `0.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 145.9 KiB | [postgresql-13-convert_0.0.4-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-13-convert_0.0.4-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-convert` | `0.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.6 KiB | [postgresql-13-convert_0.0.4-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-13-convert_0.0.4-3PIGSTY~trixie_amd64.deb) |
| `postgresql-13-convert` | `0.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 145.9 KiB | [postgresql-13-convert_0.0.4-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-13-convert_0.0.4-3PIGSTY~trixie_arm64.deb) |
| `postgresql-13-convert` | `0.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 272.3 KiB | [postgresql-13-convert_0.0.4-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-13-convert_0.0.4-3PIGSTY~jammy_amd64.deb) |
| `postgresql-13-convert` | `0.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.5 KiB | [postgresql-13-convert_0.0.4-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-13-convert_0.0.4-3PIGSTY~jammy_arm64.deb) |
| `postgresql-13-convert` | `0.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 270.4 KiB | [postgresql-13-convert_0.0.4-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-13-convert_0.0.4-3PIGSTY~noble_amd64.deb) |
| `postgresql-13-convert` | `0.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.2 KiB | [postgresql-13-convert_0.0.4-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-13-convert_0.0.4-3PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rustprooflabs/convert" title="Repository" icon="github" subtitle="github.com/rustprooflabs/convert" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="convert-0.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_convert;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_convert;		# install via package name, for the active PG version
pig install convert;		# install by extension name, for the current active PG version

pig install convert -v 18;   # install for PG 18
pig install convert -v 17;   # install for PG 17
pig install convert -v 16;   # install for PG 16
pig install convert -v 15;   # install for PG 15
pig install convert -v 14;   # install for PG 14
pig install convert -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION convert;
```
