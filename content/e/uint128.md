---
title: "uint128"
linkTitle: "uint128"
description: "Native uint128 type"
weight: 3740
categories: ["TYPE"]
width: full
---

[**pg_uint128**](https://github.com/pg-uint/pg-uint128) : Native uint128 type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3740** | {{< badge content="uint128" link="https://github.com/pg-uint/pg-uint128" >}} | {{< ext "uint128" "pg_uint128" >}} | `1.2.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] breaks on el8 since 1.1 ,fix el8 build problem by adding __has_builtin marco


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_uint128` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "pg_uint128_18" "green" >}} {{< bg "17" "pg_uint128_17" "green" >}} {{< bg "16" "pg_uint128_16" "green" >}} {{< bg "15" "pg_uint128_15" "green" >}} {{< bg "14" "pg_uint128_14" "green" >}} {{< bg "13" "pg_uint128_13" "green" >}} | `pg_uint128_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "postgresql-18-pg-uint128" "green" >}} {{< bg "17" "postgresql-17-pg-uint128" "green" >}} {{< bg "16" "postgresql-16-pg-uint128" "green" >}} {{< bg "15" "postgresql-15-pg-uint128" "green" >}} {{< bg "14" "postgresql-14-pg-uint128" "green" >}} {{< bg "13" "postgresql-13-pg-uint128" "green" >}} | `postgresql-$v-pg-uint128` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-13-pg-uint128 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_18` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.2 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_18-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.5 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_18-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 170.9 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_18-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.4 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_18-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 169.3 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_18-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 163.4 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_18-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 328.2 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 318.4 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 327.0 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 319.3 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 365.8 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 357.9 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 355.4 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.3 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.3 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_17-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.6 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_17-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 169.9 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_17-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.5 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_17-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 169.3 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_17-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 163.3 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_17-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 327.7 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 317.5 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 327.1 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 320.1 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 384.2 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 373.3 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 355.4 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.3 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.1 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_16-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.3 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_16-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 170.7 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_16-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.4 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_16-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 169.3 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_16-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 164.1 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_16-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 327.5 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 318.1 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 326.9 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 319.7 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 381.3 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 371.7 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 354.6 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.0 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.7 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_15-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.8 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_15-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 170.2 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_15-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.3 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_15-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 172.0 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_15-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 164.7 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_15-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 331.8 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 322.1 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 331.3 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 324.3 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 383.8 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 374.0 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 359.3 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 356.3 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.7 KiB | [pg_uint128_14-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_14-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.7 KiB | [pg_uint128_14-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_14-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 171.5 KiB | [pg_uint128_14-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_14-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.7 KiB | [pg_uint128_14-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_14-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_14` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 171.8 KiB | [pg_uint128_14-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_14-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_14` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 165.1 KiB | [pg_uint128_14-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_14-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 331.7 KiB | [postgresql-14-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 322.1 KiB | [postgresql-14-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 331.0 KiB | [postgresql-14-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 324.1 KiB | [postgresql-14-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 383.7 KiB | [postgresql-14-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 374.0 KiB | [postgresql-14-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 359.4 KiB | [postgresql-14-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 356.3 KiB | [postgresql-14-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-14-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_13` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 185.4 KiB | [pg_uint128_13-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_13-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_13` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.6 KiB | [pg_uint128_13-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_13-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_13` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 170.9 KiB | [pg_uint128_13-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_13-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_13` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 160.7 KiB | [pg_uint128_13-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_13-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_13` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 171.8 KiB | [pg_uint128_13-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_13-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_13` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 165.2 KiB | [pg_uint128_13-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_13-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 329.3 KiB | [postgresql-13-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 322.0 KiB | [postgresql-13-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 330.0 KiB | [postgresql-13-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 323.0 KiB | [postgresql-13-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 382.9 KiB | [postgresql-13-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 372.7 KiB | [postgresql-13-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 359.7 KiB | [postgresql-13-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 355.5 KiB | [postgresql-13-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-13-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pg-uint/pg-uint128" title="Repository" icon="github" subtitle="github.com/pg-uint/pg-uint128" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-uint128-1.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_uint128;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_uint128;		# install via package name, for the active PG version
pig install uint128;		# install by extension name, for the current active PG version

pig install uint128 -v 18;   # install for PG 18
pig install uint128 -v 17;   # install for PG 17
pig install uint128 -v 16;   # install for PG 16
pig install uint128 -v 15;   # install for PG 15
pig install uint128 -v 14;   # install for PG 14
pig install uint128 -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION uint128;
```
