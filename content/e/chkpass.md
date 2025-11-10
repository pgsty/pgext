---
title: "chkpass"
linkTitle: "chkpass"
description: "data type for auto-encrypted passwords"
weight: 3920
categories: ["TYPE"]
width: full
---

[**chkpass**](https://github.com/lacanoid/chkpass) : data type for auto-encrypted passwords


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3920** | {{< badge content="chkpass" link="https://github.com/lacanoid/chkpass" >}} | {{< ext "chkpass" >}} | `1.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `chkpass` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "chkpass_18*" "green" >}} {{< bg "17" "chkpass_17*" "green" >}} {{< bg "16" "chkpass_16*" "green" >}} {{< bg "15" "chkpass_15*" "green" >}} {{< bg "14" "chkpass_14*" "green" >}} {{< bg "13" "chkpass_13*" "green" >}} | `chkpass_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-chkpass" "green" >}} {{< bg "17" "postgresql-17-chkpass" "green" >}} {{< bg "16" "postgresql-16-chkpass" "green" >}} {{< bg "15" "postgresql-15-chkpass" "green" >}} {{< bg "14" "postgresql-14-chkpass" "green" >}} {{< bg "13" "postgresql-13-chkpass" "green" >}} | `postgresql-$v-chkpass` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "chkpass_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "chkpass_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "chkpass_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "chkpass_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "chkpass_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "chkpass_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "chkpass_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-chkpass : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-chkpass : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-chkpass : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-chkpass : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-chkpass : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-chkpass : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-chkpass : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-chkpass : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-chkpass : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `chkpass_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.4 KiB | [chkpass_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [chkpass_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [chkpass_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [chkpass_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `chkpass_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [chkpass_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/chkpass_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `chkpass_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [chkpass_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/chkpass_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-chkpass` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.8 KiB | [postgresql-18-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-18-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-chkpass` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-18-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-18-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-chkpass` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-18-chkpass_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-18-chkpass_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-chkpass` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.7 KiB | [postgresql-18-chkpass_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-18-chkpass_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-chkpass` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.6 KiB | [postgresql-18-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-18-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-chkpass` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.8 KiB | [postgresql-18-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-18-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-chkpass` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.1 KiB | [postgresql-18-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-18-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-chkpass` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 KiB | [postgresql-18-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-18-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `chkpass_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.4 KiB | [chkpass_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [chkpass_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [chkpass_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [chkpass_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `chkpass_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [chkpass_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/chkpass_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `chkpass_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [chkpass_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/chkpass_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-chkpass` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.8 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-chkpass` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-chkpass` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-chkpass` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.7 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-chkpass` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.8 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-chkpass` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.0 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-chkpass` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.1 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-chkpass` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 KiB | [postgresql-17-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-17-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `chkpass_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.4 KiB | [chkpass_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [chkpass_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [chkpass_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [chkpass_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `chkpass_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [chkpass_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/chkpass_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `chkpass_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [chkpass_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/chkpass_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-chkpass` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.8 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-chkpass` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-chkpass` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-chkpass` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.7 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-chkpass` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.8 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-chkpass` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.0 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-chkpass` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.1 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-chkpass` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 KiB | [postgresql-16-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-16-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `chkpass_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.4 KiB | [chkpass_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [chkpass_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [chkpass_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [chkpass_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `chkpass_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [chkpass_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/chkpass_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `chkpass_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [chkpass_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/chkpass_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-chkpass` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.8 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-chkpass` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-chkpass` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-chkpass` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.7 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-chkpass` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-chkpass` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.0 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-chkpass` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.1 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-chkpass` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 KiB | [postgresql-15-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-15-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `chkpass_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.4 KiB | [chkpass_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [chkpass_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [chkpass_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [chkpass_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `chkpass_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [chkpass_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/chkpass_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `chkpass_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [chkpass_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/chkpass_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-chkpass` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.7 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-chkpass` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.4 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-chkpass` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-chkpass` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.7 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-chkpass` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.8 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-chkpass` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.0 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-chkpass` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.1 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-chkpass` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.2 KiB | [postgresql-14-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-14-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `chkpass_13` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.3 KiB | [chkpass_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/chkpass_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `chkpass_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [chkpass_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/chkpass_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `chkpass_13` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [chkpass_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/chkpass_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `chkpass_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.3 KiB | [chkpass_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/chkpass_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `chkpass_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [chkpass_13-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/chkpass_13-1.0-1PIGSTY.el10.x86_64.rpm) |
| `chkpass_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [chkpass_13-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/chkpass_13-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-chkpass` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.3 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-chkpass` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.5 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-chkpass` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.6 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-chkpass` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.0 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-chkpass` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.9 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-chkpass` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.0 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-chkpass` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.9 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-chkpass` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.7 KiB | [postgresql-13-chkpass_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/chkpass/postgresql-13-chkpass_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/lacanoid/chkpass" title="Repository" icon="github" subtitle="github.com/lacanoid/chkpass" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="chkpass-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg chkpass;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install chkpass;		# install via package name, for the active PG version

pig install chkpass -v 18;   # install for PG 18
pig install chkpass -v 17;   # install for PG 17
pig install chkpass -v 16;   # install for PG 16
pig install chkpass -v 15;   # install for PG 15
pig install chkpass -v 14;   # install for PG 14
pig install chkpass -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION chkpass;
```
