---
title: "pgs3"
linkTitle: "pgs3"
description: "S3-compatible object storage endpoint implemented inside PostgreSQL"
weight: 9430
categories: ["SIM"]
languages: ["Rust"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pgs3**](https://github.com/pgsty/pgs3) : S3-compatible object storage endpoint implemented inside PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9430** | {{< badge content="pgs3" link="https://github.com/pgsty/pgs3" >}} | {{< ext "pgs3" >}} | `0.1.1` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgs3` |
|   **See Also**    | {{< ext "aws_s3" >}} {{< ext "pg_lake" >}} {{< ext "pg_parquet" >}} {{< ext "pg_ducklake" >}} {{< ext "omni_aws" >}} |

> [!Note] Early alpha; PG17-18; endpoint startup requires preload or pgs3.start(); TLS terminates externally; small-object and 100,000-object Fork performance gates remain unmet.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pgs3` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "pgs3_18" "green" >}} {{< bg "17" "pgs3_17" "green" >}} {{< bg "16" "pgs3_16" "red" >}} {{< bg "15" "pgs3_15" "red" >}} {{< bg "14" "pgs3_14" "red" >}} | `pgs3_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.1` | {{< bg "18" "postgresql-18-pgs3" "green" >}} {{< bg "17" "postgresql-17-pgs3" "green" >}} {{< bg "16" "postgresql-16-pgs3" "red" >}} {{< bg "15" "postgresql-15-pgs3" "red" >}} {{< bg "14" "postgresql-14-pgs3" "red" >}} | `postgresql-$v-pgs3` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgs3_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgs3_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgs3_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgs3_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgs3_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "pgs3_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgs3_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pgs3_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-18-pgs3 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.1" "postgresql-17-pgs3 : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pgs3 : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pgs3 : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgs3_18` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 919.2 KiB | [pgs3_18-0.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgs3_18-0.1.1-1PGSTY.el8.x86_64.rpm) |
| `pgs3_18` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 740.1 KiB | [pgs3_18-0.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgs3_18-0.1.1-1PGSTY.el8.aarch64.rpm) |
| `pgs3_18` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 892.7 KiB | [pgs3_18-0.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgs3_18-0.1.1-1PGSTY.el9.x86_64.rpm) |
| `pgs3_18` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 795.8 KiB | [pgs3_18-0.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgs3_18-0.1.1-1PGSTY.el9.aarch64.rpm) |
| `pgs3_18` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 892.7 KiB | [pgs3_18-0.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgs3_18-0.1.1-1PGSTY.el10.x86_64.rpm) |
| `pgs3_18` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 796.7 KiB | [pgs3_18-0.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgs3_18-0.1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgs3` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 796.1 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 655.4 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 796.4 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 657.6 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 863.0 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 767.4 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 858.2 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 762.8 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 856.0 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pgs3` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 760.0 KiB | [postgresql-18-pgs3_0.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgs3/postgresql-18-pgs3_0.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgs3_17` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 919.1 KiB | [pgs3_17-0.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgs3_17-0.1.1-1PGSTY.el8.x86_64.rpm) |
| `pgs3_17` | `0.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 740.0 KiB | [pgs3_17-0.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgs3_17-0.1.1-1PGSTY.el8.aarch64.rpm) |
| `pgs3_17` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 892.9 KiB | [pgs3_17-0.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgs3_17-0.1.1-1PGSTY.el9.x86_64.rpm) |
| `pgs3_17` | `0.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 796.2 KiB | [pgs3_17-0.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgs3_17-0.1.1-1PGSTY.el9.aarch64.rpm) |
| `pgs3_17` | `0.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 892.9 KiB | [pgs3_17-0.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgs3_17-0.1.1-1PGSTY.el10.x86_64.rpm) |
| `pgs3_17` | `0.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 794.9 KiB | [pgs3_17-0.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgs3_17-0.1.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgs3` | `0.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 796.4 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 654.9 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 796.4 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 654.8 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 862.9 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 767.5 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 858.2 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 762.1 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 856.0 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pgs3` | `0.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 759.9 KiB | [postgresql-17-pgs3_0.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgs3/postgresql-17-pgs3_0.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsty/pgs3" title="Repository" icon="github" subtitle="github.com/pgsty/pgs3" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgs3-0.1.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgs3;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgs3;		# install via package name, for the active PG version

pig install pgs3 -v 18;   # install for PG 18
pig install pgs3 -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgs3';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgs3;
```

## Usage

Sources:

- [Official release v0.1.1](https://github.com/pgsty/pgs3/releases/tag/v0.1.1)
- [Official README v0.1.1](https://github.com/pgsty/pgs3/blob/v0.1.1/README.md)
- [Extension control file](https://github.com/pgsty/pgs3/blob/v0.1.1/pgs3.control)
- [Configuration reference](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/guc.md)
- [Operations guide](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/operations.md)
- [Known limitations](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/known-limitations.md)

`pgs3` 0.1.1 turns one PostgreSQL database into a path-style S3-compatible endpoint. PostgreSQL background workers authenticate SigV4 requests and execute versioned object operations against ordinary SQL tables, so object metadata, payloads, authorization, WAL, physical backup, and recovery remain inside PostgreSQL. It targets PostgreSQL 17 and 18 and remains early-alpha software.

### Core Workflow

Install the extension in the database that will own the object store, create a restricted tenant role and credential, then start the worker pool:

```sql
CREATE EXTENSION pgs3;

CREATE ROLE tenant_app
  NOLOGIN NOINHERIT NOSUPERUSER NOCREATEDB NOCREATEROLE
  NOREPLICATION NOBYPASSRLS;

SELECT pgs3.create_credential(
  'tenant-access-key', 'replace-with-a-secret', 'tenant_app'::name, true
);

SELECT pgs3.start();
TABLE pgs3.worker_state;
TABLE pgs3.stats;
```

For automatic startup, preload the library and configure the target database before restarting PostgreSQL:

```conf
shared_preload_libraries = 'pgs3'
pgs3.enabled = on
pgs3.target_database = 'artifacts'
pgs3.listen_addr = '127.0.0.1'
pgs3.port = 9000
pgs3.workers = 4
```

Adding or removing `shared_preload_libraries`, or changing `pgs3.target_database`, requires a PostgreSQL restart. Other documented settings use SIGHUP semantics, but operational readiness must be checked through `pgs3.worker_state` and logs rather than only an open TCP port. A manually started pool can be stopped with `pgs3.stop()`.

### Client and Storage Behavior

Clients must use path-style addressing and an explicit endpoint:

```bash
export PGS3_ENDPOINT='https://s3.example.com'
export AWS_ACCESS_KEY_ID='<access-key>'
export AWS_SECRET_ACCESS_KEY='<secret-key>'
export AWS_DEFAULT_REGION='us-east-1'

aws --endpoint-url "$PGS3_ENDPOINT" s3api list-buckets
```

Always pass `--endpoint-url`; otherwise an AWS client can silently send a request to AWS. The supported client paths include AWS CLI, boto3, rclone, s3fs, and DuckDB `httpfs`. Bucket and object operations include range and conditional reads/writes, ListObjectsV2, permanent version history, delete markers, CopyObject, and multipart upload.

The canonical payload lives in `pgs3.blob`. Object versions, CopyObject, SQL Restore, and metadata-only Fork operations can share the same blob rather than duplicating bytes. One endpoint serves one configured database.

### Operations and Security

Credential access keys map to PostgreSQL tenant roles. Keep those roles `NOLOGIN`, `NOINHERIT`, and `NOBYPASSRLS`; do not use `pgs3.server_role` as an application identity. Row-level security is the tenant-isolation boundary, while credential-management and worker-control functions remain operator-only.

SigV4 requires reversibly stored secrets, so database backups and replicas contain credential material and must be encrypted and access-controlled. `pgs3` serves cleartext HTTP; production deployments need an external TLS proxy that preserves the signed path, host, and headers. Use physical backups: logical dump/restore is not supported for the extension-owned object state.

### Compatibility and Limitations

The packaged and upstream-supported paths are PostgreSQL 17 and 18 with pgrx 0.19.2. Version 0.1.1 includes the tested `0.1.0 -> 0.1.1` extension upgrade edge.

`pgs3` is not a general-purpose production S3 replacement. Virtual-host addressing, built-in TLS, IAM or bucket-policy languages, ACLs, lifecycle rules, and cross-database routing are not implemented. Small-object GET/PUT targets and the 100,000-object Fork target are not met, and full-object GET currently materializes the response in memory. Set a deployment-specific object-size limit and review the upstream limitations before exposing an endpoint.
