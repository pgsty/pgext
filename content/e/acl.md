---
title: "acl"
linkTitle: "acl"
description: "ACL Data type"
weight: 3860
categories: ["TYPE"]
width: full
---

[**pg_acl**](https://github.com/arkhipov/acl) : ACL Data type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3860** | {{< badge content="acl" link="https://github.com/arkhipov/acl" >}} | {{< ext "acl" "pg_acl" >}} | `1.0.4` | {{< category "TYPE" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] +cast pg_uuid_t


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_acl` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.4` | {{< bg "18" "acl_18" "green" >}} {{< bg "17" "acl_17" "green" >}} {{< bg "16" "acl_16" "green" >}} {{< bg "15" "acl_15" "green" >}} {{< bg "14" "acl_14" "green" >}} {{< bg "13" "acl_13" "green" >}} | `acl_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.4` | {{< bg "18" "postgresql-18-acl" "green" >}} {{< bg "17" "postgresql-17-acl" "green" >}} {{< bg "16" "postgresql-16-acl" "green" >}} {{< bg "15" "postgresql-15-acl" "green" >}} {{< bg "14" "postgresql-14-acl" "green" >}} {{< bg "13" "postgresql-13-acl" "green" >}} | `postgresql-$v-acl` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "acl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "acl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "acl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "acl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "acl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "acl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-18-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-13-acl : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-18-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-13-acl : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-18-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-13-acl : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-18-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-13-acl : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-18-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-13-acl : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-18-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-13-acl : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-18-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-13-acl : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-18-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-13-acl : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_18` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [acl_18-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_18-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_18` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.7 KiB | [acl_18-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_18-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_18` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.5 KiB | [acl_18-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_18-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_18` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.1 KiB | [acl_18-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_18-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `acl_18` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.9 KiB | [acl_18-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acl_18-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `acl_18` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.2 KiB | [acl_18-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acl_18-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-acl` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.0 KiB | [postgresql-18-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-18-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-acl` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.6 KiB | [postgresql-18-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-18-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-acl` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.2 KiB | [postgresql-18-acl_1.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-18-acl_1.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-acl` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.9 KiB | [postgresql-18-acl_1.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-18-acl_1.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-acl` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 47.5 KiB | [postgresql-18-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-18-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-acl` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 47.5 KiB | [postgresql-18-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-18-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-acl` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.2 KiB | [postgresql-18-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-18-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-acl` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.6 KiB | [postgresql-18-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-18-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_17` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [acl_17-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_17-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_17` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.7 KiB | [acl_17-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_17-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_17` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.4 KiB | [acl_17-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_17-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_17` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.1 KiB | [acl_17-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_17-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `acl_17` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.9 KiB | [acl_17-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acl_17-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `acl_17` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.2 KiB | [acl_17-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acl_17-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-acl` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.0 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-acl` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.6 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-acl` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.2 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-acl` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.7 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-acl` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 50.1 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-acl` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.0 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-acl` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.2 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-acl` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.6 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_16` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [acl_16-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_16-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_16` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.7 KiB | [acl_16-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_16-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_16` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.5 KiB | [acl_16-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_16-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_16` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.1 KiB | [acl_16-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_16-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `acl_16` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.9 KiB | [acl_16-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acl_16-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `acl_16` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.2 KiB | [acl_16-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acl_16-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-acl` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.0 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-acl` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.6 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-acl` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.2 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-acl` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.8 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-acl` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 50.1 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-acl` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.0 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-acl` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 47.2 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-acl` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.6 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_15` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [acl_15-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_15-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_15` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.8 KiB | [acl_15-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_15-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_15` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.6 KiB | [acl_15-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_15-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_15` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.1 KiB | [acl_15-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_15-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `acl_15` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [acl_15-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acl_15-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `acl_15` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.4 KiB | [acl_15-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acl_15-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-acl` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-acl` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.7 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-acl` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.3 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-acl` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.9 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-acl` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 50.3 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-acl` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 49.9 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-acl` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.9 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-acl` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.5 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_14` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [acl_14-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_14-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_14` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.8 KiB | [acl_14-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_14-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_14` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.6 KiB | [acl_14-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_14-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_14` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.1 KiB | [acl_14-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_14-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `acl_14` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.6 KiB | [acl_14-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acl_14-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `acl_14` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [acl_14-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acl_14-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-acl` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-acl` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.8 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-acl` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.3 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-acl` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.9 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-acl` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 50.2 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-acl` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 49.8 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-acl` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.9 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-acl` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.4 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_13` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.2 KiB | [acl_13-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_13-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_13` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.9 KiB | [acl_13-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_13-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_13` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.5 KiB | [acl_13-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_13-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_13` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.0 KiB | [acl_13-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_13-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `acl_13` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.5 KiB | [acl_13-1.0.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/acl_13-1.0.4-1PIGSTY.el10.x86_64.rpm) |
| `acl_13` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [acl_13-1.0.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/acl_13-1.0.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-acl` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.0 KiB | [postgresql-13-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-13-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-acl` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.5 KiB | [postgresql-13-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-13-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-acl` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.2 KiB | [postgresql-13-acl_1.0.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-13-acl_1.0.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-acl` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.6 KiB | [postgresql-13-acl_1.0.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/acl/postgresql-13-acl_1.0.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-acl` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 50.3 KiB | [postgresql-13-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-13-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-acl` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 49.5 KiB | [postgresql-13-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-13-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-acl` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.8 KiB | [postgresql-13-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-13-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-acl` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.2 KiB | [postgresql-13-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-13-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/arkhipov/acl" title="Repository" icon="github" subtitle="github.com/arkhipov/acl" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="acl-1.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_acl;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_acl;		# install via package name, for the active PG version
pig install acl;		# install by extension name, for the current active PG version

pig install acl -v 18;   # install for PG 18
pig install acl -v 17;   # install for PG 17
pig install acl -v 16;   # install for PG 16
pig install acl -v 15;   # install for PG 15
pig install acl -v 14;   # install for PG 14
pig install acl -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION acl;
```
