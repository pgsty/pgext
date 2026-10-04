---
title: "libx509pq"
linkTitle: "libx509pq"
description: "X.509 certificate parsing and inspection through OpenSSL"
weight: 7420
categories: ["SEC"]
languages: ["C"]
licenses: ["GPL-3.0-or-later"]
repos: ["PIGSTY"]
page_width: full
---

[**libx509pq**](https://github.com/crtsh/libx509pq) : X.509 certificate parsing and inspection through OpenSSL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7420** | {{< badge content="libx509pq" link="https://github.com/crtsh/libx509pq" >}} | {{< ext "libx509pq" >}} | `1.3` | {{< category "SEC" >}} | {{< license "GPL-3.0-or-later" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "sslinfo" >}} {{< ext "sslutils" >}} {{< ext "pg_snakeoil" >}} {{< ext "pgcrypto" >}} |

> [!Note] Requires OpenSSL 1.1.1+; GPL-3.0-or-later is explicit in source headers. EL9 default crypto policy rejects a legacy SHA-1 test fixture; SHA-256 verification passed without changing policy.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `libx509pq` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3` | {{< bg "18" "libx509pq_18" "green" >}} {{< bg "17" "libx509pq_17" "green" >}} {{< bg "16" "libx509pq_16" "green" >}} {{< bg "15" "libx509pq_15" "green" >}} {{< bg "14" "libx509pq_14" "green" >}} | `libx509pq_$v` | `openssl-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3` | {{< bg "18" "postgresql-18-libx509pq" "green" >}} {{< bg "17" "postgresql-17-libx509pq" "green" >}} {{< bg "16" "postgresql-16-libx509pq" "green" >}} {{< bg "15" "postgresql-15-libx509pq" "green" >}} {{< bg "14" "postgresql-14-libx509pq" "green" >}} | `postgresql-$v-libx509pq` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3" "libx509pq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3" "libx509pq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3" "libx509pq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3" "libx509pq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3" "libx509pq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3" "libx509pq_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "libx509pq_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-libx509pq : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-libx509pq : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `libx509pq_18` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.2 KiB | [libx509pq_18-1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/libx509pq_18-1.3-1PGSTY.el8.x86_64.rpm) |
| `libx509pq_18` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 91.2 KiB | [libx509pq_18-1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/libx509pq_18-1.3-1PGSTY.el8.aarch64.rpm) |
| `libx509pq_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 92.8 KiB | [libx509pq_18-1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/libx509pq_18-1.3-1PGSTY.el9.x86_64.rpm) |
| `libx509pq_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.4 KiB | [libx509pq_18-1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/libx509pq_18-1.3-1PGSTY.el9.aarch64.rpm) |
| `libx509pq_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 93.1 KiB | [libx509pq_18-1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/libx509pq_18-1.3-1PGSTY.el10.x86_64.rpm) |
| `libx509pq_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 92.4 KiB | [libx509pq_18-1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/libx509pq_18-1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-libx509pq` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 76.6 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 75.9 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 76.5 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 75.7 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.0 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 82.5 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 80.5 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 80.5 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 80.2 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-libx509pq` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 79.6 KiB | [postgresql-18-libx509pq_1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-18-libx509pq_1.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `libx509pq_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.2 KiB | [libx509pq_17-1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/libx509pq_17-1.3-1PGSTY.el8.x86_64.rpm) |
| `libx509pq_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 91.2 KiB | [libx509pq_17-1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/libx509pq_17-1.3-1PGSTY.el8.aarch64.rpm) |
| `libx509pq_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 93.0 KiB | [libx509pq_17-1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/libx509pq_17-1.3-1PGSTY.el9.x86_64.rpm) |
| `libx509pq_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.5 KiB | [libx509pq_17-1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/libx509pq_17-1.3-1PGSTY.el9.aarch64.rpm) |
| `libx509pq_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 93.2 KiB | [libx509pq_17-1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/libx509pq_17-1.3-1PGSTY.el10.x86_64.rpm) |
| `libx509pq_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 92.6 KiB | [libx509pq_17-1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/libx509pq_17-1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-libx509pq` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 76.6 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 75.9 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 76.5 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 75.7 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 97.2 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 96.6 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 80.6 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 80.6 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 80.2 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-libx509pq` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 79.7 KiB | [postgresql-17-libx509pq_1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-17-libx509pq_1.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `libx509pq_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.2 KiB | [libx509pq_16-1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/libx509pq_16-1.3-1PGSTY.el8.x86_64.rpm) |
| `libx509pq_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 91.2 KiB | [libx509pq_16-1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/libx509pq_16-1.3-1PGSTY.el8.aarch64.rpm) |
| `libx509pq_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 93.0 KiB | [libx509pq_16-1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/libx509pq_16-1.3-1PGSTY.el9.x86_64.rpm) |
| `libx509pq_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.5 KiB | [libx509pq_16-1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/libx509pq_16-1.3-1PGSTY.el9.aarch64.rpm) |
| `libx509pq_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 93.2 KiB | [libx509pq_16-1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/libx509pq_16-1.3-1PGSTY.el10.x86_64.rpm) |
| `libx509pq_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 92.6 KiB | [libx509pq_16-1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/libx509pq_16-1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-libx509pq` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 76.6 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 75.9 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 76.5 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 75.7 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 96.4 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 95.8 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 80.6 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 80.6 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 80.3 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-libx509pq` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 79.6 KiB | [postgresql-16-libx509pq_1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-16-libx509pq_1.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `libx509pq_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.4 KiB | [libx509pq_15-1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/libx509pq_15-1.3-1PGSTY.el8.x86_64.rpm) |
| `libx509pq_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 91.3 KiB | [libx509pq_15-1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/libx509pq_15-1.3-1PGSTY.el8.aarch64.rpm) |
| `libx509pq_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 93.6 KiB | [libx509pq_15-1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/libx509pq_15-1.3-1PGSTY.el9.x86_64.rpm) |
| `libx509pq_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.9 KiB | [libx509pq_15-1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/libx509pq_15-1.3-1PGSTY.el9.aarch64.rpm) |
| `libx509pq_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 93.7 KiB | [libx509pq_15-1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/libx509pq_15-1.3-1PGSTY.el10.x86_64.rpm) |
| `libx509pq_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 93.0 KiB | [libx509pq_15-1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/libx509pq_15-1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-libx509pq` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 77.0 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 76.2 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 76.9 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 75.9 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 96.9 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 96.3 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 81.1 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 81.1 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 80.7 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-libx509pq` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 80.1 KiB | [postgresql-15-libx509pq_1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-15-libx509pq_1.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `libx509pq_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 92.2 KiB | [libx509pq_14-1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/libx509pq_14-1.3-1PGSTY.el8.x86_64.rpm) |
| `libx509pq_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 91.2 KiB | [libx509pq_14-1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/libx509pq_14-1.3-1PGSTY.el8.aarch64.rpm) |
| `libx509pq_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 93.5 KiB | [libx509pq_14-1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/libx509pq_14-1.3-1PGSTY.el9.x86_64.rpm) |
| `libx509pq_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.8 KiB | [libx509pq_14-1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/libx509pq_14-1.3-1PGSTY.el9.aarch64.rpm) |
| `libx509pq_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 93.6 KiB | [libx509pq_14-1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/libx509pq_14-1.3-1PGSTY.el10.x86_64.rpm) |
| `libx509pq_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 92.9 KiB | [libx509pq_14-1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/libx509pq_14-1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-libx509pq` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.9 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.1 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.8 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 80.8 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 99.4 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.9 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 85.8 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.0 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 85.5 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-libx509pq` | `1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 85.0 KiB | [postgresql-14-libx509pq_1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/libx/libx509pq/postgresql-14-libx509pq_1.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/crtsh/libx509pq" title="Repository" icon="github" subtitle="github.com/crtsh/libx509pq" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="libx509pq-1.3-bed05c6e87d7.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg libx509pq;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install libx509pq;		# install via package name, for the active PG version

pig install libx509pq -v 18;   # install for PG 18
pig install libx509pq -v 17;   # install for PG 17
pig install libx509pq -v 16;   # install for PG 16
pig install libx509pq -v 15;   # install for PG 15
pig install libx509pq -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION libx509pq;
```

## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/README.md)
- [Version 1.3 SQL objects](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/libx509pq--1.3.sql)
- [OpenSSL-backed implementation](https://github.com/crtsh/libx509pq/blob/bed05c6e87d79a98d79e42a1fac16386eead0c13/libx509pq.c)

`libx509pq` exposes OpenSSL X.509 certificate parsing to SQL. It accepts DER certificates as `bytea` and extracts subject, issuer, validity, serial number, key and signature information, extensions, alternate names, fingerprints, and related fields.

```sql
CREATE EXTENSION libx509pq;
SELECT x509_commonName(der),
       x509_issuerName(der),
       x509_notBefore(der),
       x509_notAfter(der),
       x509_keyAlgorithm(der),
       x509_keySize(der)
FROM certificate_store;
```

For multiple basic fields, `x509_basic_info(der)` avoids parsing the same certificate repeatedly. Other APIs enumerate extensions and names, expose public keys, identify selected weak-key patterns, and verify one certificate signature against a supplied key.

Parsing or signature verification is not full PKIX path validation: it does not by itself establish trust, hostname validity, revocation status, policy, or current acceptability. Malformed DER reaches native OpenSSL code; cap input size, fuzz and regression-test untrusted corpora, and keep PostgreSQL and OpenSSL patched. Pin the OpenSSL ABI used at build and runtime, review sentinel-versus-null error behavior, restrict expensive functions, and benchmark bulk parsing.
