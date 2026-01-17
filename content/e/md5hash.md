---
title: "md5hash"
linkTitle: "md5hash"
description: "type for storing 128-bit binary data inline"
weight: 3550
categories: ["TYPE"]
width: full
---

[**md5hash**](https://github.com/tvondra/md5hash) : type for storing 128-bit binary data inline


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3550** | {{< badge content="md5hash" link="https://github.com/tvondra/md5hash" >}} | {{< ext "md5hash" >}} | `1.0.1` | {{< category "TYPE" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hashlib" >}} {{< ext "xxhash" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `md5hash` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "md5hash_18" "green" >}} {{< bg "17" "md5hash_17" "green" >}} {{< bg "16" "md5hash_16" "green" >}} {{< bg "15" "md5hash_15" "green" >}} {{< bg "14" "md5hash_14" "green" >}} {{< bg "13" "md5hash_13" "green" >}} | `md5hash_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "postgresql-18-md5hash" "green" >}} {{< bg "17" "postgresql-17-md5hash" "green" >}} {{< bg "16" "postgresql-16-md5hash" "green" >}} {{< bg "15" "postgresql-15-md5hash" "green" >}} {{< bg "14" "postgresql-14-md5hash" "green" >}} {{< bg "13" "postgresql-13-md5hash" "green" >}} | `postgresql-$v-md5hash` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "md5hash_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-md5hash : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-md5hash : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-md5hash : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-md5hash : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-md5hash : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-md5hash : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-md5hash : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-md5hash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-md5hash : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `md5hash_18` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.9 KiB | [md5hash_18-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/md5hash_18-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `md5hash_18` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [md5hash_18-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/md5hash_18-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `md5hash_18` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [md5hash_18-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/md5hash_18-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `md5hash_18` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [md5hash_18-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/md5hash_18-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `md5hash_18` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.7 KiB | [md5hash_18-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/md5hash_18-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `md5hash_18` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.0 KiB | [md5hash_18-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/md5hash_18-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-md5hash` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 KiB | [postgresql-18-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-18-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-md5hash` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.6 KiB | [postgresql-18-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-18-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-md5hash` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.5 KiB | [postgresql-18-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-18-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-md5hash` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.7 KiB | [postgresql-18-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-18-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-md5hash` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.2 KiB | [postgresql-18-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-18-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-md5hash` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.4 KiB | [postgresql-18-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-18-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-md5hash` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-18-md5hash_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-18-md5hash_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-md5hash` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.2 KiB | [postgresql-18-md5hash_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-18-md5hash_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `md5hash_17` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.8 KiB | [md5hash_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/md5hash_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `md5hash_17` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [md5hash_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/md5hash_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `md5hash_17` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [md5hash_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/md5hash_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `md5hash_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [md5hash_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/md5hash_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `md5hash_17` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.7 KiB | [md5hash_17-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/md5hash_17-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `md5hash_17` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.0 KiB | [md5hash_17-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/md5hash_17-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-md5hash` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 KiB | [postgresql-17-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-17-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-md5hash` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.6 KiB | [postgresql-17-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-17-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-md5hash` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.5 KiB | [postgresql-17-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-17-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-md5hash` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 KiB | [postgresql-17-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-17-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-md5hash` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.4 KiB | [postgresql-17-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-17-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-md5hash` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 KiB | [postgresql-17-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-17-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-md5hash` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-17-md5hash_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-17-md5hash_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-md5hash` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.2 KiB | [postgresql-17-md5hash_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-17-md5hash_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `md5hash_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.8 KiB | [md5hash_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/md5hash_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `md5hash_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [md5hash_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/md5hash_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `md5hash_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [md5hash_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/md5hash_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `md5hash_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [md5hash_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/md5hash_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `md5hash_16` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.7 KiB | [md5hash_16-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/md5hash_16-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `md5hash_16` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.0 KiB | [md5hash_16-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/md5hash_16-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-md5hash` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 KiB | [postgresql-16-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-16-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-md5hash` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.6 KiB | [postgresql-16-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-16-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-md5hash` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.5 KiB | [postgresql-16-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-16-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-md5hash` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 KiB | [postgresql-16-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-16-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-md5hash` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.4 KiB | [postgresql-16-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-16-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-md5hash` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 KiB | [postgresql-16-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-16-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-md5hash` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-16-md5hash_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-16-md5hash_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-md5hash` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.2 KiB | [postgresql-16-md5hash_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-16-md5hash_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `md5hash_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.9 KiB | [md5hash_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/md5hash_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `md5hash_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [md5hash_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/md5hash_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `md5hash_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [md5hash_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/md5hash_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `md5hash_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [md5hash_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/md5hash_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `md5hash_15` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.7 KiB | [md5hash_15-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/md5hash_15-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `md5hash_15` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.0 KiB | [md5hash_15-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/md5hash_15-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-md5hash` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 KiB | [postgresql-15-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-15-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-md5hash` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.6 KiB | [postgresql-15-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-15-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-md5hash` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.5 KiB | [postgresql-15-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-15-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-md5hash` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 KiB | [postgresql-15-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-15-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-md5hash` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.4 KiB | [postgresql-15-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-15-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-md5hash` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 KiB | [postgresql-15-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-15-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-md5hash` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.3 KiB | [postgresql-15-md5hash_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-15-md5hash_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-md5hash` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.2 KiB | [postgresql-15-md5hash_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-15-md5hash_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `md5hash_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.8 KiB | [md5hash_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/md5hash_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `md5hash_14` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [md5hash_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/md5hash_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `md5hash_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [md5hash_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/md5hash_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `md5hash_14` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [md5hash_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/md5hash_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `md5hash_14` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.7 KiB | [md5hash_14-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/md5hash_14-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `md5hash_14` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [md5hash_14-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/md5hash_14-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-md5hash` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 KiB | [postgresql-14-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-14-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-md5hash` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.6 KiB | [postgresql-14-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-14-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-md5hash` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.5 KiB | [postgresql-14-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-14-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-md5hash` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 KiB | [postgresql-14-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-14-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-md5hash` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.4 KiB | [postgresql-14-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-14-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-md5hash` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 KiB | [postgresql-14-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-14-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-md5hash` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.2 KiB | [postgresql-14-md5hash_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-14-md5hash_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-md5hash` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.2 KiB | [postgresql-14-md5hash_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-14-md5hash_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `md5hash_13` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.8 KiB | [md5hash_13-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/md5hash_13-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `md5hash_13` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.1 KiB | [md5hash_13-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/md5hash_13-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `md5hash_13` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [md5hash_13-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/md5hash_13-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `md5hash_13` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.8 KiB | [md5hash_13-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/md5hash_13-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `md5hash_13` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.7 KiB | [md5hash_13-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/md5hash_13-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `md5hash_13` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [md5hash_13-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/md5hash_13-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-md5hash` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 KiB | [postgresql-13-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-13-md5hash_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-md5hash` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.7 KiB | [postgresql-13-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/md5hash/postgresql-13-md5hash_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-md5hash` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 13.6 KiB | [postgresql-13-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-13-md5hash_1.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-md5hash` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.7 KiB | [postgresql-13-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/md5hash/postgresql-13-md5hash_1.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-md5hash` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.2 KiB | [postgresql-13-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-13-md5hash_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-md5hash` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.3 KiB | [postgresql-13-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/md5hash/postgresql-13-md5hash_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-md5hash` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 14.4 KiB | [postgresql-13-md5hash_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-13-md5hash_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-md5hash` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.3 KiB | [postgresql-13-md5hash_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/md5hash/postgresql-13-md5hash_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/md5hash" title="Repository" icon="github" subtitle="github.com/tvondra/md5hash" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="md5hash-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg md5hash;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install md5hash;		# install via package name, for the active PG version

pig install md5hash -v 18;   # install for PG 18
pig install md5hash -v 17;   # install for PG 17
pig install md5hash -v 16;   # install for PG 16
pig install md5hash -v 15;   # install for PG 15
pig install md5hash -v 14;   # install for PG 14
pig install md5hash -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION md5hash;
```
