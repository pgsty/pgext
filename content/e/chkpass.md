---
title: "chkpass"
linkTitle: "chkpass"
description: "data type for auto-encrypted passwords"
weight: 3920
categories: ["Type"]
width: full
---

data type for auto-encrypted passwords

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3920** | {{< badge content="chkpass" link="https://github.com/lacanoid/chkpass" >}} | {{< ext "chkpass" "chkpass" >}} | `1.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/chkpass" >}} | `1.0` | {{< badge content="18" color="red" alt="chkpass_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `chkpass_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/chkpass" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-chkpass" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-chkpass` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "chkpass_18" >}}     | {{< pkg "chkpass_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "chkpass_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "chkpass_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "chkpass_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "chkpass_18" >}}     | {{< pkg "chkpass_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "chkpass_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "chkpass_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "chkpass_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "chkpass_18" >}}     | {{< pkg "chkpass_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "chkpass_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "chkpass_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "chkpass_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "chkpass_18" >}}     | {{< pkg "chkpass_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "chkpass_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "chkpass_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "chkpass_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-chkpass" >}}     | {{< pkg "postgresql-17-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-chkpass" >}}     | {{< pkg "postgresql-17-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-chkpass" >}}     | {{< pkg "postgresql-17-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-chkpass" >}}     | {{< pkg "postgresql-17-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-chkpass" >}}     | {{< pkg "postgresql-17-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-chkpass" >}}     | {{< pkg "postgresql-17-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-chkpass" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `chkpass_17` | 1.0 | `el8.x86_64` | pigsty | 13.4 KiB | [chkpass_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_17` | 1.0 | `el8.aarch64` | pigsty | 13.4 KiB | [chkpass_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_17` | 1.0 | `el9.aarch64` | pigsty | 13.4 KiB | [chkpass_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `chkpass_17` | 1.0 | `el9.x86_64` | pigsty | 13.5 KiB | [chkpass_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-chkpass` | 1.0 | `d12.x86_64` | pigsty | 10.7 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-chkpass` | 1.0 | `d12.aarch64` | pigsty | 10.9 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-chkpass` | 1.0 | `u22.x86_64` | pigsty | 10.8 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-chkpass` | 1.0 | `u22.aarch64` | pigsty | 11.0 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-chkpass` | 1.0 | `u24.x86_64` | pigsty | 10.9 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-chkpass` | 1.0 | `u24.aarch64` | pigsty | 11.2 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `chkpass_16` | 1.0 | `el8.x86_64` | pigsty | 13.4 KiB | [chkpass_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_16` | 1.0 | `el8.aarch64` | pigsty | 13.4 KiB | [chkpass_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_16` | 1.0 | `el9.x86_64` | pigsty | 13.5 KiB | [chkpass_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_16` | 1.0 | `el9.aarch64` | pigsty | 13.3 KiB | [chkpass_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-chkpass` | 1.0 | `d12.x86_64` | pigsty | 10.7 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-chkpass` | 1.0 | `d12.aarch64` | pigsty | 10.9 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-chkpass` | 1.0 | `u22.aarch64` | pigsty | 11.0 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-chkpass` | 1.0 | `u22.x86_64` | pigsty | 10.8 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-chkpass` | 1.0 | `u24.x86_64` | pigsty | 10.9 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-chkpass` | 1.0 | `u24.aarch64` | pigsty | 11.2 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `chkpass_15` | 1.0 | `el8.x86_64` | pigsty | 13.4 KiB | [chkpass_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_15` | 1.0 | `el8.aarch64` | pigsty | 13.4 KiB | [chkpass_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_15` | 1.0 | `el9.x86_64` | pigsty | 13.5 KiB | [chkpass_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_15` | 1.0 | `el9.aarch64` | pigsty | 13.4 KiB | [chkpass_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-chkpass` | 1.0 | `d12.aarch64` | pigsty | 11.0 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-chkpass` | 1.0 | `d12.x86_64` | pigsty | 10.8 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-chkpass` | 1.0 | `u22.aarch64` | pigsty | 11.0 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-chkpass` | 1.0 | `u22.x86_64` | pigsty | 10.9 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-chkpass` | 1.0 | `u24.x86_64` | pigsty | 11.0 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-chkpass` | 1.0 | `u24.aarch64` | pigsty | 11.2 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `chkpass_14` | 1.0 | `el8.x86_64` | pigsty | 13.4 KiB | [chkpass_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_14` | 1.0 | `el8.aarch64` | pigsty | 13.4 KiB | [chkpass_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_14` | 1.0 | `el9.x86_64` | pigsty | 13.5 KiB | [chkpass_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_14` | 1.0 | `el9.aarch64` | pigsty | 13.4 KiB | [chkpass_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-chkpass` | 1.0 | `d12.x86_64` | pigsty | 10.8 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-chkpass` | 1.0 | `d12.aarch64` | pigsty | 11.0 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-chkpass` | 1.0 | `u22.x86_64` | pigsty | 10.8 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-chkpass` | 1.0 | `u22.aarch64` | pigsty | 11.0 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-chkpass` | 1.0 | `u24.x86_64` | pigsty | 10.9 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-chkpass` | 1.0 | `u24.aarch64` | pigsty | 11.2 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `chkpass_13` | 1.0 | `el8.aarch64` | pigsty | 13.4 KiB | [chkpass_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_13` | 1.0 | `el8.x86_64` | pigsty | 13.3 KiB | [chkpass_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_13` | 1.0 | `el9.aarch64` | pigsty | 13.4 KiB | [chkpass_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `chkpass_13` | 1.0 | `el9.x86_64` | pigsty | 13.5 KiB | [chkpass_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-chkpass` | 1.0 | `d12.aarch64` | pigsty | 10.6 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-chkpass` | 1.0 | `d12.x86_64` | pigsty | 10.6 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-chkpass` | 1.0 | `u22.aarch64` | pigsty | 11.0 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-chkpass` | 1.0 | `u22.x86_64` | pigsty | 11.0 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-chkpass` | 1.0 | `u24.aarch64` | pigsty | 10.9 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-chkpass` | 1.0 | `u24.x86_64` | pigsty | 11.0 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lacanoid/chkpass" title="Repository" icon="github" subtitle="github.com/lacanoid/chkpass" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="chkpass-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get chkpass; # get chkpass source code
pig build dep chkpass; # install build dependencies
pig build pkg chkpass; # build extension rpm or deb
pig build ext chkpass; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install chkpass; # install by extension name, for the current active PG version
pig ext install chkpass; # install via package alias, for the active PG version
pig ext install chkpass -v 18;   # install for PG 18
pig ext install chkpass -v 17;   # install for PG 17
pig ext install chkpass -v 16;   # install for PG 16
pig ext install chkpass -v 15;   # install for PG 15
pig ext install chkpass -v 14;   # install for PG 14
pig ext install chkpass -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION chkpass;
```

