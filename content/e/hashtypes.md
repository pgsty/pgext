---
title: "hashtypes"
linkTitle: "hashtypes"
description: "sha1, md5 and other data types for PostgreSQL"
weight: 3750
categories: ["TYPE"]
width: full
---

[**hashtypes**](https://github.com/adjust/hashtypes/) : sha1, md5 and other data types for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3750** | {{< badge content="hashtypes" link="https://github.com/adjust/hashtypes/" >}} | {{< ext "hashtypes" >}} | `0.1.5` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `hashtypes` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "hashtypes_18" "green" >}} {{< bg "17" "hashtypes_17" "green" >}} {{< bg "16" "hashtypes_16" "green" >}} {{< bg "15" "hashtypes_15" "green" >}} {{< bg "14" "hashtypes_14" "green" >}} | `hashtypes_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.5` | {{< bg "18" "postgresql-18-hashtypes" "green" >}} {{< bg "17" "postgresql-17-hashtypes" "green" >}} {{< bg "16" "postgresql-16-hashtypes" "green" >}} {{< bg "15" "postgresql-15-hashtypes" "green" >}} {{< bg "14" "postgresql-14-hashtypes" "green" >}} | `postgresql-$v-hashtypes` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "hashtypes_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-hashtypes : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-hashtypes : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-hashtypes : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-hashtypes : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-hashtypes : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-hashtypes : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-hashtypes : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-18-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-17-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-16-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-15-hashtypes : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.5" "postgresql-14-hashtypes : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-hashtypes : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hashtypes : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-hashtypes : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-hashtypes : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-hashtypes : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-hashtypes : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-hashtypes : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-hashtypes : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-hashtypes : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-hashtypes : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hashtypes_18` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.0 KiB | [hashtypes_18-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_18-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_18` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.9 KiB | [hashtypes_18-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_18-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_18` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.2 KiB | [hashtypes_18-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_18-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `hashtypes_18` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.7 KiB | [hashtypes_18-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_18-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `hashtypes_18` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.0 KiB | [hashtypes_18-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hashtypes_18-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `hashtypes_18` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.0 KiB | [hashtypes_18-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hashtypes_18-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-hashtypes` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 33.5 KiB | [postgresql-18-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-18-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-hashtypes` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 33.5 KiB | [postgresql-18-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-18-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-hashtypes` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 33.5 KiB | [postgresql-18-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-18-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-hashtypes` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.7 KiB | [postgresql-18-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-18-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-hashtypes` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.8 KiB | [postgresql-18-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-18-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-hashtypes` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.2 KiB | [postgresql-18-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-18-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-hashtypes` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.0 KiB | [postgresql-18-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-18-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-hashtypes` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.1 KiB | [postgresql-18-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-18-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hashtypes_17` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.0 KiB | [hashtypes_17-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_17-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_17` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.9 KiB | [hashtypes_17-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_17-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_17` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.2 KiB | [hashtypes_17-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_17-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `hashtypes_17` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.7 KiB | [hashtypes_17-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_17-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `hashtypes_17` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.0 KiB | [hashtypes_17-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hashtypes_17-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `hashtypes_17` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.0 KiB | [hashtypes_17-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hashtypes_17-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-hashtypes` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 33.5 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-hashtypes` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 33.5 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-hashtypes` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 33.6 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-hashtypes` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.6 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-hashtypes` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.4 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-hashtypes` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.9 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-hashtypes` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.0 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-hashtypes` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.2 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hashtypes_16` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.0 KiB | [hashtypes_16-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_16-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_16` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.9 KiB | [hashtypes_16-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_16-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_16` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.2 KiB | [hashtypes_16-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_16-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `hashtypes_16` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.7 KiB | [hashtypes_16-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_16-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `hashtypes_16` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.0 KiB | [hashtypes_16-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hashtypes_16-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `hashtypes_16` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.0 KiB | [hashtypes_16-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hashtypes_16-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-hashtypes` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 33.5 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-hashtypes` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 33.4 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-hashtypes` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 33.5 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-hashtypes` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.6 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-hashtypes` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.4 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-hashtypes` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.9 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-hashtypes` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.0 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-hashtypes` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hashtypes_15` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.0 KiB | [hashtypes_15-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_15-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_15` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.9 KiB | [hashtypes_15-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_15-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_15` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.3 KiB | [hashtypes_15-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_15-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `hashtypes_15` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.0 KiB | [hashtypes_15-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_15-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `hashtypes_15` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.2 KiB | [hashtypes_15-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hashtypes_15-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `hashtypes_15` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.2 KiB | [hashtypes_15-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hashtypes_15-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-hashtypes` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 33.5 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-hashtypes` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 33.5 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-hashtypes` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 33.6 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-hashtypes` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.5 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-hashtypes` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.7 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-hashtypes` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.1 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-hashtypes` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.1 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-hashtypes` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hashtypes_14` | `0.1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.0 KiB | [hashtypes_14-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_14-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_14` | `0.1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 25.9 KiB | [hashtypes_14-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_14-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_14` | `0.1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.2 KiB | [hashtypes_14-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_14-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `hashtypes_14` | `0.1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.0 KiB | [hashtypes_14-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_14-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `hashtypes_14` | `0.1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.2 KiB | [hashtypes_14-0.1.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hashtypes_14-0.1.5-1PIGSTY.el10.x86_64.rpm) |
| `hashtypes_14` | `0.1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.2 KiB | [hashtypes_14-0.1.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hashtypes_14-0.1.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-hashtypes` | `0.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 33.5 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-hashtypes` | `0.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 33.4 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-hashtypes` | `0.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 33.6 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-hashtypes` | `0.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.4 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-hashtypes` | `0.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.7 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-hashtypes` | `0.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.0 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-hashtypes` | `0.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 34.9 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-hashtypes` | `0.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/hashtypes/" title="Repository" icon="github" subtitle="github.com/adjust/hashtypes/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hashtypes-0.1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg hashtypes;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install hashtypes;		# install via package name, for the active PG version

pig install hashtypes -v 18;   # install for PG 18
pig install hashtypes -v 17;   # install for PG 17
pig install hashtypes -v 16;   # install for PG 16
pig install hashtypes -v 15;   # install for PG 15
pig install hashtypes -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hashtypes;
```



## Usage

> [hashtypes: hash and checksum data types (SHA, CRC32)](https://github.com/adjust/hashtypes/)

The `hashtypes` extension provides native data types for storing hash values and checksums in their binary representation, saving storage compared to text.

```sql
CREATE EXTENSION hashtypes;
```

### Data Types

| Type | Storage | Description |
|------|---------|-------------|
| `sha1` | 20 bytes | SHA-1 hash (160-bit) |
| `sha224` | 28 bytes | SHA-224 hash (224-bit) |
| `sha256` | 32 bytes | SHA-256 hash (256-bit) |
| `sha384` | 48 bytes | SHA-384 hash (384-bit) |
| `sha512` | 64 bytes | SHA-512 hash (512-bit) |
| `crc32` | 4 bytes | CRC-32 checksum |

### Usage

```sql
CREATE TABLE objects (
    hash sha256 PRIMARY KEY,
    data bytea
);

INSERT INTO objects VALUES (
    'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855',
    '\x'
);

SELECT * FROM objects WHERE hash = 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855';
```

### Operators

All types support comparison operators: `=`, `<>`, `<`, `>`, `<=`, `>=`.

### Index Support

Btree and hash index operator classes are provided for all types.

### Casts

All types support bidirectional casts to/from `text` and `bytea`.
