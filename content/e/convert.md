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
| **4850** | {{< badge content="convert" link="https://github.com/rustprooflabs/convert" >}} | {{< ext "convert" "pg_convert" >}} | `0.1.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "unit" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_convert` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_convert_18" "green" >}} {{< bg "17" "pg_convert_17" "green" >}} {{< bg "16" "pg_convert_16" "green" >}} {{< bg "15" "pg_convert_15" "green" >}} {{< bg "14" "pg_convert_14" "green" >}} | `pg_convert_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-convert" "green" >}} {{< bg "17" "postgresql-17-convert" "green" >}} {{< bg "16" "postgresql-16-convert" "green" >}} {{< bg "15" "postgresql-15-convert" "green" >}} {{< bg "14" "postgresql-14-convert" "green" >}} | `postgresql-$v-convert` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_convert_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-convert : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-convert : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-convert : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-convert : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-convert : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-convert : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-convert : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-convert : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-convert : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 293.2 KiB | [pg_convert_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.8 KiB | [pg_convert_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 308.7 KiB | [pg_convert_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_convert_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.2 KiB | [pg_convert_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_convert_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 308.9 KiB | [pg_convert_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_convert_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.2 KiB | [pg_convert_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-convert` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 241.9 KiB | [postgresql-18-convert_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-18-convert_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-convert` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 146.0 KiB | [postgresql-18-convert_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-18-convert_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-convert` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 242.0 KiB | [postgresql-18-convert_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-18-convert_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-convert` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 146.0 KiB | [postgresql-18-convert_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-18-convert_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-convert` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 273.8 KiB | [postgresql-18-convert_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-18-convert_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-convert` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 170.1 KiB | [postgresql-18-convert_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-18-convert_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-convert` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 271.4 KiB | [postgresql-18-convert_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-18-convert_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-convert` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.8 KiB | [postgresql-18-convert_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-18-convert_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 293.0 KiB | [pg_convert_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.9 KiB | [pg_convert_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 308.7 KiB | [pg_convert_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_convert_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.2 KiB | [pg_convert_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_convert_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 309.0 KiB | [pg_convert_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_convert_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.2 KiB | [pg_convert_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-convert` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 242.4 KiB | [postgresql-17-convert_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-17-convert_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-convert` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 146.0 KiB | [postgresql-17-convert_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-17-convert_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-convert` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 242.2 KiB | [postgresql-17-convert_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-17-convert_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-convert` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 145.9 KiB | [postgresql-17-convert_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-17-convert_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-convert` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 273.7 KiB | [postgresql-17-convert_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-17-convert_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-convert` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 170.0 KiB | [postgresql-17-convert_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-17-convert_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-convert` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 271.2 KiB | [postgresql-17-convert_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-17-convert_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-convert` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.7 KiB | [postgresql-17-convert_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-17-convert_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 293.1 KiB | [pg_convert_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.9 KiB | [pg_convert_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 308.5 KiB | [pg_convert_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_convert_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.2 KiB | [pg_convert_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_convert_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 308.7 KiB | [pg_convert_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_convert_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.2 KiB | [pg_convert_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-convert` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 241.9 KiB | [postgresql-16-convert_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-16-convert_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-convert` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 146.0 KiB | [postgresql-16-convert_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-16-convert_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-convert` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 241.6 KiB | [postgresql-16-convert_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-16-convert_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-convert` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 146.0 KiB | [postgresql-16-convert_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-16-convert_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-convert` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 273.5 KiB | [postgresql-16-convert_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-16-convert_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-convert` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.9 KiB | [postgresql-16-convert_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-16-convert_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-convert` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 271.2 KiB | [postgresql-16-convert_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-16-convert_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-convert` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.7 KiB | [postgresql-16-convert_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-16-convert_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 292.9 KiB | [pg_convert_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.8 KiB | [pg_convert_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 308.3 KiB | [pg_convert_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_convert_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.2 KiB | [pg_convert_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_convert_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 308.5 KiB | [pg_convert_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_convert_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.2 KiB | [pg_convert_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-convert` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 242.1 KiB | [postgresql-15-convert_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-15-convert_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-convert` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 145.9 KiB | [postgresql-15-convert_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-15-convert_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-convert` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 242.2 KiB | [postgresql-15-convert_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-15-convert_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-convert` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 146.0 KiB | [postgresql-15-convert_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-15-convert_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-convert` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 273.7 KiB | [postgresql-15-convert_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-15-convert_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-convert` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.9 KiB | [postgresql-15-convert_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-15-convert_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-convert` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 271.0 KiB | [postgresql-15-convert_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-15-convert_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-convert` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.8 KiB | [postgresql-15-convert_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-15-convert_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_convert_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 293.0 KiB | [pg_convert_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_convert_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_convert_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 189.8 KiB | [pg_convert_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_convert_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_convert_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 308.1 KiB | [pg_convert_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_convert_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_convert_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 203.1 KiB | [pg_convert_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_convert_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_convert_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 308.4 KiB | [pg_convert_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_convert_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_convert_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.2 KiB | [pg_convert_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_convert_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-convert` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 241.8 KiB | [postgresql-14-convert_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-14-convert_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-convert` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 146.0 KiB | [postgresql-14-convert_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/convert/postgresql-14-convert_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-convert` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 242.3 KiB | [postgresql-14-convert_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-14-convert_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-convert` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 146.0 KiB | [postgresql-14-convert_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/convert/postgresql-14-convert_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-convert` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 273.2 KiB | [postgresql-14-convert_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-14-convert_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-convert` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 169.9 KiB | [postgresql-14-convert_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/convert/postgresql-14-convert_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-convert` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 270.9 KiB | [postgresql-14-convert_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-14-convert_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-convert` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.9 KiB | [postgresql-14-convert_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/convert/postgresql-14-convert_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rustprooflabs/convert" title="Repository" icon="github" subtitle="github.com/rustprooflabs/convert" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="convert-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_convert;		# build rpm/deb
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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION convert;
```



## Usage

> [convert: common unit conversion functions for PostgreSQL](https://github.com/rustprooflabs/convert)

Provides functions for common unit conversions: distance, speed, time-to-travel, power, area, and temperature.

```sql
CREATE EXTENSION convert;
```

### Distance Functions

| Function | Description |
|---|---|
| `dist_mi_to_ft(miles)` | Miles to feet |
| `dist_ft_to_mi(feet)` | Feet to miles |
| `dist_ft_to_m(feet)` | Feet to meters |
| `dist_m_to_ft(meters)` | Meters to feet |
| `dist_m_to_km(meters)` | Meters to kilometers |
| `dist_km_to_m(km)` | Kilometers to meters |
| `dist_mi_to_km(miles)` | Miles to kilometers |
| `dist_m_to_mi(meters)` | Meters to miles |
| `dist_km_to_mi(km)` | Kilometers to miles |

### Speed Functions

| Function | Description |
|---|---|
| `speed_mph_to_kmhr(mph)` | MPH to km/h |
| `speed_kmhr_to_mph(kmhr)` | km/h to MPH |
| `speed_kmhr_to_m_s(kmhr)` | km/h to m/s |
| `speed_mph_to_m_s(mph)` | MPH to m/s |
| `speed_m_s_to_kmhr(m_s)` | m/s to km/h |
| `speed_m_s_to_mph(m_s)` | m/s to MPH |

### Area Functions

| Function | Description |
|---|---|
| `area_m2_to_km2(m2)` | sq meters to sq km |
| `area_m2_to_ft2(m2)` | sq meters to sq feet |
| `area_ft2_to_m2(ft2)` | sq feet to sq meters |
| `area_ft2_to_mi2(ft2)` | sq feet to sq miles |
| `area_mi2_to_ft2(mi2)` | sq miles to sq feet |
| `area_mi2_to_acre(mi2)` | sq miles to acres |
| `area_acre_to_mi2(acres)` | Acres to sq miles |
| `area_acre_to_km2(acres)` | Acres to sq km |

### Temperature Functions

| Function | Description |
|---|---|
| `temp_c_to_f(celsius)` | Celsius to Fahrenheit |
| `temp_f_to_c(fahrenheit)` | Fahrenheit to Celsius |

### Power Functions

| Function | Description |
|---|---|
| `power_dbm_to_watts(dbm)` | dBm to watts |
| `power_watts_to_dbm(watts)` | Watts to dBm |

### Examples

```sql
SELECT dist_mi_to_km(26.2);      -- 42.16 (marathon in km)
SELECT temp_f_to_c(98.6);         -- 37.0
SELECT speed_mph_to_kmhr(60.0);   -- 96.56
SELECT area_acre_to_km2(640.0);   -- ~2.59
```
