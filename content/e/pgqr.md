---
title: "pgqr"
linkTitle: "pgqr"
description: "QR Code generator from PostgreSQL"
weight: 4250
categories: ["UTIL"]
width: full
---

[**pgqr**](https://github.com/AbdulYadi/pgqr) : QR Code generator from PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4250** | {{< badge content="pgqr" link="https://github.com/AbdulYadi/pgqr" >}} | {{< ext "pgqr" >}} | `1.0` | {{< category "UTIL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_protobuf" >}} {{< ext "base36" >}} {{< ext "base62" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgqr` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pgqr_18" "green" >}} {{< bg "17" "pgqr_17" "green" >}} {{< bg "16" "pgqr_16" "green" >}} {{< bg "15" "pgqr_15" "green" >}} {{< bg "14" "pgqr_14" "green" >}} {{< bg "13" "pgqr_13" "green" >}} | `pgqr_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pgqr" "green" >}} {{< bg "17" "postgresql-17-pgqr" "green" >}} {{< bg "16" "postgresql-16-pgqr" "green" >}} {{< bg "15" "postgresql-15-pgqr" "green" >}} {{< bg "14" "postgresql-14-pgqr" "green" >}} {{< bg "13" "postgresql-13-pgqr" "green" >}} | `postgresql-$v-pgqr` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pgqr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pgqr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pgqr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pgqr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pgqr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pgqr_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pgqr : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pgqr : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pgqr : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pgqr : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pgqr : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pgqr : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pgqr : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pgqr : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.5 KiB | [pgqr_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.8 KiB | [pgqr_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.0 KiB | [pgqr_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.6 KiB | [pgqr_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgqr_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.8 KiB | [pgqr_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgqr_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgqr_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.9 KiB | [pgqr_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgqr_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgqr` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 47.9 KiB | [postgresql-18-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-18-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgqr` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.6 KiB | [postgresql-18-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-18-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgqr` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.2 KiB | [postgresql-18-pgqr_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-18-pgqr_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgqr` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.7 KiB | [postgresql-18-pgqr_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-18-pgqr_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgqr` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 50.2 KiB | [postgresql-18-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-18-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgqr` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.8 KiB | [postgresql-18-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-18-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgqr` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.0 KiB | [postgresql-18-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-18-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgqr` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.0 KiB | [postgresql-18-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-18-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.5 KiB | [pgqr_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.8 KiB | [pgqr_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.0 KiB | [pgqr_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.6 KiB | [pgqr_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgqr_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.8 KiB | [pgqr_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgqr_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgqr_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.9 KiB | [pgqr_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgqr_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgqr` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 47.8 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgqr` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.5 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgqr` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.1 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgqr` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.6 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgqr` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 50.9 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgqr` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 49.6 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgqr` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 49.9 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgqr` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.0 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.5 KiB | [pgqr_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.8 KiB | [pgqr_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.8 KiB | [pgqr_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.6 KiB | [pgqr_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgqr_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.8 KiB | [pgqr_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgqr_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgqr_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.9 KiB | [pgqr_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgqr_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgqr` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 47.8 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgqr` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.5 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgqr` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.2 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgqr` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.7 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgqr` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 50.9 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgqr` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 49.6 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgqr` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.0 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgqr` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.0 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.8 KiB | [pgqr_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.2 KiB | [pgqr_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.8 KiB | [pgqr_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.5 KiB | [pgqr_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgqr_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.7 KiB | [pgqr_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgqr_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgqr_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.0 KiB | [pgqr_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgqr_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgqr` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 47.7 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgqr` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.5 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgqr` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.3 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgqr` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.9 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgqr` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 51.3 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgqr` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.1 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgqr` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.6 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgqr` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.5 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.8 KiB | [pgqr_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.2 KiB | [pgqr_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.8 KiB | [pgqr_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.5 KiB | [pgqr_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgqr_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.7 KiB | [pgqr_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgqr_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgqr_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.9 KiB | [pgqr_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgqr_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgqr` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 47.7 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgqr` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.5 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgqr` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.3 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgqr` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.8 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgqr` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 51.3 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgqr` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.1 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgqr` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.5 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgqr` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.5 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_13` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.7 KiB | [pgqr_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.2 KiB | [pgqr_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_13` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.8 KiB | [pgqr_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.5 KiB | [pgqr_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgqr_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.7 KiB | [pgqr_13-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgqr_13-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgqr_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.9 KiB | [pgqr_13-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgqr_13-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgqr` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 47.8 KiB | [postgresql-13-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-13-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgqr` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.6 KiB | [postgresql-13-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-13-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgqr` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 48.5 KiB | [postgresql-13-pgqr_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-13-pgqr_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgqr` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 47.0 KiB | [postgresql-13-pgqr_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgqr/postgresql-13-pgqr_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgqr` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 51.2 KiB | [postgresql-13-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-13-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgqr` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 50.0 KiB | [postgresql-13-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-13-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgqr` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 50.7 KiB | [postgresql-13-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-13-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgqr` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 49.7 KiB | [postgresql-13-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-13-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/AbdulYadi/pgqr" title="Repository" icon="github" subtitle="github.com/AbdulYadi/pgqr" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgqr-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgqr;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgqr;		# install via package name, for the active PG version

pig install pgqr -v 18;   # install for PG 18
pig install pgqr -v 17;   # install for PG 17
pig install pgqr -v 16;   # install for PG 16
pig install pgqr -v 15;   # install for PG 15
pig install pgqr -v 14;   # install for PG 14
pig install pgqr -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgqr;
```
