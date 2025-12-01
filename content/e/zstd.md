---
title: "zstd"
linkTitle: "zstd"
description: "Zstandard compression algorithm implementation in PostgreSQL"
weight: 4030
categories: ["UTIL"]
width: full
---

[**pg_zstd**](https://github.com/grahamedgecombe/pgzstd) : Zstandard compression algorithm implementation in PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4030** | {{< badge content="zstd" link="https://github.com/grahamedgecombe/pgzstd" >}} | {{< ext "zstd" "pg_zstd" >}} | `1.1.2` | {{< category "UTIL" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} |

> [!Note] +varatt.h


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_zstd` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "pg_zstd_18*" "green" >}} {{< bg "17" "pg_zstd_17*" "green" >}} {{< bg "16" "pg_zstd_16*" "green" >}} {{< bg "15" "pg_zstd_15*" "green" >}} {{< bg "14" "pg_zstd_14*" "green" >}} {{< bg "13" "pg_zstd_13*" "green" >}} | `pg_zstd_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.2` | {{< bg "18" "postgresql-18-zstd" "green" >}} {{< bg "17" "postgresql-17-zstd" "green" >}} {{< bg "16" "postgresql-16-zstd" "green" >}} {{< bg "15" "postgresql-15-zstd" "green" >}} {{< bg "14" "postgresql-14-zstd" "green" >}} {{< bg "13" "postgresql-13-zstd" "green" >}} | `postgresql-$v-zstd` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "pg_zstd_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-13-zstd : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-13-zstd : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-13-zstd : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-13-zstd : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-13-zstd : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-13-zstd : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-13-zstd : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-18-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-17-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-16-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-15-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-14-zstd : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.2" "postgresql-13-zstd : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_zstd_18` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.2 KiB | [pg_zstd_18-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_18-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_18` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.2 KiB | [pg_zstd_18-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_18-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_18` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.8 KiB | [pg_zstd_18-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_18-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_18` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.7 KiB | [pg_zstd_18-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_18-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_zstd_18` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.8 KiB | [pg_zstd_18-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_zstd_18-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_zstd_18` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.8 KiB | [pg_zstd_18-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_zstd_18-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-zstd` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.9 KiB | [postgresql-18-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-18-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-zstd` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.8 KiB | [postgresql-18-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-18-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-zstd` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.9 KiB | [postgresql-18-zstd_1.1.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-18-zstd_1.1.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-zstd` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.8 KiB | [postgresql-18-zstd_1.1.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-18-zstd_1.1.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-zstd` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.2 KiB | [postgresql-18-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-18-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-zstd` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.0 KiB | [postgresql-18-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-18-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-zstd` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-18-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-18-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-zstd` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.1 KiB | [postgresql-18-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-18-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_zstd_17` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.2 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_17-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_17` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.2 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_17-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_17` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.0 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_17-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_17` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.8 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_17-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_zstd_17` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.9 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_zstd_17-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_zstd_17` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.0 KiB | [pg_zstd_17-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_zstd_17-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-zstd` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.9 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-zstd` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.8 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-zstd` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.9 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-zstd` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.8 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-zstd` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.6 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-zstd` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.4 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-zstd` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-zstd` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-17-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-17-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_zstd_16` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.2 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_16-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_16` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.2 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_16-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_16` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.0 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_16-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_16` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.8 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_16-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_zstd_16` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.9 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_zstd_16-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_zstd_16` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.9 KiB | [pg_zstd_16-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_zstd_16-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-zstd` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.9 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-zstd` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.8 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-zstd` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.9 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-zstd` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.9 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-zstd` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.6 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-zstd` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.4 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-zstd` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.4 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-zstd` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-16-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-16-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_zstd_15` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.2 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_15-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_15` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.2 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_15-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_15` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.0 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_15-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_15` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.8 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_15-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_zstd_15` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.9 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_zstd_15-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_zstd_15` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.9 KiB | [pg_zstd_15-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_zstd_15-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-zstd` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.8 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-zstd` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.8 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-zstd` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.9 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-zstd` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.8 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-zstd` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.6 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-zstd` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.4 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-zstd` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.4 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-zstd` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-15-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-15-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_zstd_14` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.2 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_14-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_14` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.2 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_14-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_14` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.0 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_14-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_14` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.8 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_14-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_zstd_14` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.9 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_zstd_14-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_zstd_14` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.9 KiB | [pg_zstd_14-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_zstd_14-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-zstd` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.8 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-zstd` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.8 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-zstd` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.9 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-zstd` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.8 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-zstd` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.6 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-zstd` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.4 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-zstd` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-zstd` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-14-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-14-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_zstd_13` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.1 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_zstd_13-1.1.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_zstd_13` | `1.1.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.2 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_zstd_13-1.1.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_zstd_13` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.9 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_zstd_13-1.1.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_zstd_13` | `1.1.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.8 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_zstd_13-1.1.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_zstd_13` | `1.1.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.9 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_zstd_13-1.1.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_zstd_13` | `1.1.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.9 KiB | [pg_zstd_13-1.1.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_zstd_13-1.1.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-zstd` | `1.1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.8 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-zstd` | `1.1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.7 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-zstd` | `1.1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.8 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-13-zstd` | `1.1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.8 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-13-zstd` | `1.1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.4 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-zstd` | `1.1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-zstd` | `1.1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-zstd` | `1.1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-13-zstd_1.1.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-zstd/postgresql-13-zstd_1.1.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/grahamedgecombe/pgzstd" title="Repository" icon="github" subtitle="github.com/grahamedgecombe/pgzstd" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgzstd-1.1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_zstd;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_zstd;		# install via package name, for the active PG version
pig install zstd;		# install by extension name, for the current active PG version

pig install zstd -v 18;   # install for PG 18
pig install zstd -v 17;   # install for PG 17
pig install zstd -v 16;   # install for PG 16
pig install zstd -v 15;   # install for PG 15
pig install zstd -v 14;   # install for PG 14
pig install zstd -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION zstd;
```


## Usage

| Function                                                                             | Return Type |
|--------------------------------------------------------------------------------------|-------------|
| <code>zstd_compress(*data* bytea [, *dictionary* bytea [, *level* integer ]])</code> | `bytea`     |
| <code>zstd_decompress(*data* bytea [, *dictionary* bytea ])</code>                   | `bytea`     |
| <code>zstd_length(*data* bytea)</code>                                               | `integer`   |

`zstd_compress` compresses the provided `data` and returns a Zstandard frame. A
preset `dictionary` may also be provided. The default compression `level` may
also be overriden, valid values range from `1` (best speed) to `22` (best
compression). The default level is `3`.

If you want to override the compression level without using a dictionary, set
`dictionary` to `NULL`.

`zstd_decompress` decompresses the provided Zstandard frame in `data` and
returns the uncompressed data. A preset `dictionary`, matching the dictionary
used to compress the data, may also be provided.

`zstd_length` returns the decompressed length of the provided Zstandard frame.
If `ZSTD_getFrameContentSize()` is available it returns `NULL` if the length is
unknown. If unavailable, it isn't possible to distinguish the error, unknown
decompressed length and zero decompressed length cases.


### Example

Basic compress/decompress example:

```sql
CREATE EXTENSION zstd;

SELECT zstd_compress('hello world');
-- zstd_compress
-- --------------------------------------------
-- \x28b52ffd200b59000068656c6c6f20776f726c64

SELECT convert_from(zstd_decompress('\x28b52ffd200b59000068656c6c6f20776f726c64'), 'utf-8');
-- convert_from
-- --------------
--  hello world
```

Compress with level (`1` for best speed, `22` for best compression, default to `3`)

```sql
CREATE EXTENSION zstd;

SELECT zstd_compress('hello world',  NULL, 10);
-- zstd_compress
-- --------------------------------------------
-- \x28b52ffd200b59000068656c6c6f20776f726c64

SELECT convert_from(zstd_decompress('\x28b52ffd200b59000068656c6c6f20776f726c64'), 'utf-8');
-- convert_from
-- --------------
--  hello world
```
