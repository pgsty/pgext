---
title: "hashtypes"
linkTitle: "hashtypes"
description: "sha1, md5 and other data types for PostgreSQL"
weight: 3750
categories: ["Type"]
width: full
---

sha1, md5 and other data types for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3750** | {{< badge content="hashtypes" link="https://github.com/adjust/hashtypes/" >}} | {{< ext "hashtypes" "hashtypes" >}} | `0.1.5` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/hashtypes" >}} | `0.1.5` | {{< badge content="18" color="red" alt="hashtypes_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `hashtypes_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/hashtypes" >}} | `0.1.5` | {{< badge content="18" color="red" alt="postgresql-18-hashtypes" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-hashtypes` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "hashtypes_18" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_18-0.1.5-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "hashtypes_17" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_17-0.1.5-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "hashtypes_16" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_16-0.1.5-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "hashtypes_15" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_15-0.1.5-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "hashtypes_14" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_14-0.1.5-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "hashtypes_18" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_18-0.1.5-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "hashtypes_17" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_17-0.1.5-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "hashtypes_16" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_16-0.1.5-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "hashtypes_15" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_15-0.1.5-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "hashtypes_14" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_14-0.1.5-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "hashtypes_18" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_18-0.1.5-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "hashtypes_17" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_17-0.1.5-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "hashtypes_16" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_16-0.1.5-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "hashtypes_15" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_15-0.1.5-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "hashtypes_14" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_14-0.1.5-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "hashtypes_18" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_18-0.1.5-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "hashtypes_17" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_17-0.1.5-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "hashtypes_16" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_16-0.1.5-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "hashtypes_15" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_15-0.1.5-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "hashtypes_14" "0.1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_14-0.1.5-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-hashtypes" >}}     | {{< pkg "postgresql-17-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-hashtypes" >}}     | {{< pkg "postgresql-17-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-hashtypes" >}}     | {{< pkg "postgresql-17-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-hashtypes" >}}     | {{< pkg "postgresql-17-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-hashtypes" >}}     | {{< pkg "postgresql-17-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-hashtypes" >}}     | {{< pkg "postgresql-17-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-hashtypes" "0.1.5" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hashtypes_18` | 0.1.5 | `el8.aarch64` | pigsty | 23.2 KiB | [hashtypes_18-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_18-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_18` | 0.1.5 | `el8.x86_64` | pigsty | 23.4 KiB | [hashtypes_18-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_18-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_18` | 0.1.5 | `el9.x86_64` | pigsty | 23.6 KiB | [hashtypes_18-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_18-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `hashtypes_18` | 0.1.5 | `el9.aarch64` | pigsty | 23.2 KiB | [hashtypes_18-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_18-0.1.5-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hashtypes_17` | 0.1.5 | `el8.aarch64` | pigsty | 23.2 KiB | [hashtypes_17-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_17-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_17` | 0.1.5 | `el8.x86_64` | pigsty | 23.4 KiB | [hashtypes_17-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_17-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_17` | 0.1.5 | `el9.aarch64` | pigsty | 23.2 KiB | [hashtypes_17-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_17-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `hashtypes_17` | 0.1.5 | `el9.x86_64` | pigsty | 23.7 KiB | [hashtypes_17-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_17-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-hashtypes` | 0.1.5 | `d12.aarch64` | pigsty | 34.0 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-hashtypes` | 0.1.5 | `d12.x86_64` | pigsty | 34.3 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-hashtypes` | 0.1.5 | `u22.aarch64` | pigsty | 35.8 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-hashtypes` | 0.1.5 | `u22.x86_64` | pigsty | 36.4 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-hashtypes` | 0.1.5 | `u24.x86_64` | pigsty | 35.1 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-hashtypes` | 0.1.5 | `u24.aarch64` | pigsty | 34.8 KiB | [postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-17-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hashtypes_16` | 0.1.5 | `el8.aarch64` | pigsty | 23.1 KiB | [hashtypes_16-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_16-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_16` | 0.1.5 | `el8.x86_64` | pigsty | 23.4 KiB | [hashtypes_16-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_16-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_16` | 0.1.5 | `el9.x86_64` | pigsty | 23.7 KiB | [hashtypes_16-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_16-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `hashtypes_16` | 0.1.5 | `el9.aarch64` | pigsty | 23.2 KiB | [hashtypes_16-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_16-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-hashtypes` | 0.1.5 | `d12.x86_64` | pigsty | 34.3 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-hashtypes` | 0.1.5 | `d12.aarch64` | pigsty | 34.0 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-hashtypes` | 0.1.5 | `u22.aarch64` | pigsty | 35.8 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-hashtypes` | 0.1.5 | `u22.x86_64` | pigsty | 36.4 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-hashtypes` | 0.1.5 | `u24.aarch64` | pigsty | 34.8 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-hashtypes` | 0.1.5 | `u24.x86_64` | pigsty | 35.1 KiB | [postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-16-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hashtypes_15` | 0.1.5 | `el8.x86_64` | pigsty | 23.5 KiB | [hashtypes_15-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_15-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_15` | 0.1.5 | `el8.aarch64` | pigsty | 23.2 KiB | [hashtypes_15-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_15-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_15` | 0.1.5 | `el9.x86_64` | pigsty | 23.7 KiB | [hashtypes_15-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_15-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `hashtypes_15` | 0.1.5 | `el9.aarch64` | pigsty | 23.5 KiB | [hashtypes_15-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_15-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-hashtypes` | 0.1.5 | `d12.x86_64` | pigsty | 34.5 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-hashtypes` | 0.1.5 | `d12.aarch64` | pigsty | 34.3 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-hashtypes` | 0.1.5 | `u22.x86_64` | pigsty | 36.7 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-hashtypes` | 0.1.5 | `u22.aarch64` | pigsty | 36.1 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-hashtypes` | 0.1.5 | `u24.aarch64` | pigsty | 34.9 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-hashtypes` | 0.1.5 | `u24.x86_64` | pigsty | 35.1 KiB | [postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-15-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hashtypes_14` | 0.1.5 | `el8.x86_64` | pigsty | 23.5 KiB | [hashtypes_14-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_14-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_14` | 0.1.5 | `el8.aarch64` | pigsty | 23.2 KiB | [hashtypes_14-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_14-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_14` | 0.1.5 | `el9.aarch64` | pigsty | 23.4 KiB | [hashtypes_14-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_14-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `hashtypes_14` | 0.1.5 | `el9.x86_64` | pigsty | 23.7 KiB | [hashtypes_14-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_14-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-hashtypes` | 0.1.5 | `d12.x86_64` | pigsty | 34.4 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-hashtypes` | 0.1.5 | `d12.aarch64` | pigsty | 34.2 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-hashtypes` | 0.1.5 | `u22.x86_64` | pigsty | 36.7 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-hashtypes` | 0.1.5 | `u22.aarch64` | pigsty | 36.0 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-hashtypes` | 0.1.5 | `u24.aarch64` | pigsty | 34.9 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-hashtypes` | 0.1.5 | `u24.x86_64` | pigsty | 35.1 KiB | [postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-14-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `hashtypes_13` | 0.1.5 | `el8.x86_64` | pigsty | 23.4 KiB | [hashtypes_13-0.1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hashtypes_13-0.1.5-1PIGSTY.el8.x86_64.rpm) |
| `hashtypes_13` | 0.1.5 | `el8.aarch64` | pigsty | 23.2 KiB | [hashtypes_13-0.1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hashtypes_13-0.1.5-1PIGSTY.el8.aarch64.rpm) |
| `hashtypes_13` | 0.1.5 | `el9.aarch64` | pigsty | 23.3 KiB | [hashtypes_13-0.1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hashtypes_13-0.1.5-1PIGSTY.el9.aarch64.rpm) |
| `hashtypes_13` | 0.1.5 | `el9.x86_64` | pigsty | 23.7 KiB | [hashtypes_13-0.1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hashtypes_13-0.1.5-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-hashtypes` | 0.1.5 | `d12.x86_64` | pigsty | 34.4 KiB | [postgresql-13-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-13-hashtypes_0.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-hashtypes` | 0.1.5 | `d12.aarch64` | pigsty | 34.3 KiB | [postgresql-13-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hashtypes/postgresql-13-hashtypes_0.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-hashtypes` | 0.1.5 | `u22.aarch64` | pigsty | 36.1 KiB | [postgresql-13-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-13-hashtypes_0.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-hashtypes` | 0.1.5 | `u22.x86_64` | pigsty | 36.3 KiB | [postgresql-13-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hashtypes/postgresql-13-hashtypes_0.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-hashtypes` | 0.1.5 | `u24.x86_64` | pigsty | 35.1 KiB | [postgresql-13-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-13-hashtypes_0.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-hashtypes` | 0.1.5 | `u24.aarch64` | pigsty | 35.0 KiB | [postgresql-13-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hashtypes/postgresql-13-hashtypes_0.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/hashtypes/" title="Repository" icon="github" subtitle="github.com/adjust/hashtypes/" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="hashtypes-0.1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get hashtypes; # get hashtypes source code
pig build dep hashtypes; # install build dependencies
pig build pkg hashtypes; # build extension rpm or deb
pig build ext hashtypes; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install hashtypes; # install by extension name, for the current active PG version
pig ext install hashtypes; # install via package alias, for the active PG version
pig ext install hashtypes -v 18;   # install for PG 18
pig ext install hashtypes -v 17;   # install for PG 17
pig ext install hashtypes -v 16;   # install for PG 16
pig ext install hashtypes -v 15;   # install for PG 15
pig ext install hashtypes -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION hashtypes;
```

