---
title: "gzip"
linkTitle: "gzip"
description: "gzip and gunzip functions."
weight: 4010
categories: ["UTIL"]
width: full
---

[**pg_gzip**](https://github.com/pramsey/pgsql-gzip) : gzip and gunzip functions.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4010** | {{< badge content="gzip" link="https://github.com/pramsey/pgsql-gzip" >}} | {{< ext "gzip" "pg_gzip" >}} | `1.0.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} {{< ext "pgjwt" >}} {{< ext "pg_smtp_client" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_gzip` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.0` | {{< bg "18" "pg_gzip_18*" "green" >}} {{< bg "17" "pg_gzip_17*" "green" >}} {{< bg "16" "pg_gzip_16*" "green" >}} {{< bg "15" "pg_gzip_15*" "green" >}} {{< bg "14" "pg_gzip_14*" "green" >}} {{< bg "13" "pg_gzip_13*" "green" >}} | `pg_gzip_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-gzip" "green" >}} {{< bg "17" "postgresql-17-gzip" "green" >}} {{< bg "16" "postgresql-16-gzip" "green" >}} {{< bg "15" "postgresql-15-gzip" "green" >}} {{< bg "14" "postgresql-14-gzip" "green" >}} {{< bg "13" "postgresql-13-gzip" "green" >}} | `postgresql-$v-gzip` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_gzip_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-gzip : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-gzip : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-gzip : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-gzip : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-gzip : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-gzip : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-gzip : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-gzip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-gzip : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_gzip_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [pg_gzip_18-1.0.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_gzip_18-1.0.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_gzip_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_gzip_18-1.0.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_gzip_18-1.0.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_gzip_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_gzip_18-1.0.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_gzip_18-1.0.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_gzip_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_gzip_18-1.0.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_gzip_18-1.0.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_gzip_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_gzip_18-1.0.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_gzip_18-1.0.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_gzip_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 KiB | [pg_gzip_18-1.0.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_gzip_18-1.0.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-gzip` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.3 KiB | [postgresql-18-gzip_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-18-gzip_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-gzip` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.2 KiB | [postgresql-18-gzip_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-18-gzip_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-gzip` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.3 KiB | [postgresql-18-gzip_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-18-gzip_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-gzip` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 KiB | [postgresql-18-gzip_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-18-gzip_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-gzip` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.7 KiB | [postgresql-18-gzip_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-18-gzip_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-gzip` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.7 KiB | [postgresql-18-gzip_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-18-gzip_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-gzip` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.8 KiB | [postgresql-18-gzip_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-18-gzip_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-gzip` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.6 KiB | [postgresql-18-gzip_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-18-gzip_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_gzip_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [pg_gzip_17-1.0.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_gzip_17-1.0.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_gzip_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_gzip_17-1.0.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_gzip_17-1.0.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_gzip_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_gzip_17-1.0.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_gzip_17-1.0.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_gzip_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_gzip_17-1.0.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_gzip_17-1.0.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_gzip_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_gzip_17-1.0.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_gzip_17-1.0.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_gzip_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 KiB | [pg_gzip_17-1.0.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_gzip_17-1.0.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-gzip` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-gzip` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-gzip` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.3 KiB | [postgresql-17-gzip_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-gzip` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 KiB | [postgresql-17-gzip_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-gzip` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-gzip` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.2 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-gzip` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.9 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-gzip` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-17-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-17-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_gzip_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [pg_gzip_16-1.0.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_gzip_16-1.0.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_gzip_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_gzip_16-1.0.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_gzip_16-1.0.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_gzip_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.7 KiB | [pg_gzip_16-1.0.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_gzip_16-1.0.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_gzip_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_gzip_16-1.0.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_gzip_16-1.0.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_gzip_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_gzip_16-1.0.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_gzip_16-1.0.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_gzip_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 KiB | [pg_gzip_16-1.0.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_gzip_16-1.0.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-gzip` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-gzip` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-gzip` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.3 KiB | [postgresql-16-gzip_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-gzip` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 KiB | [postgresql-16-gzip_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-gzip` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-gzip` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.2 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-gzip` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.9 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-gzip` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-16-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-16-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_gzip_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [pg_gzip_15-1.0.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_gzip_15-1.0.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_gzip_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_gzip_15-1.0.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_gzip_15-1.0.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_gzip_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_gzip_15-1.0.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_gzip_15-1.0.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_gzip_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_gzip_15-1.0.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_gzip_15-1.0.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_gzip_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_gzip_15-1.0.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_gzip_15-1.0.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_gzip_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 KiB | [pg_gzip_15-1.0.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_gzip_15-1.0.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-gzip` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-gzip` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-gzip` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.3 KiB | [postgresql-15-gzip_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-gzip` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 KiB | [postgresql-15-gzip_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-gzip` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-gzip` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.2 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-gzip` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.9 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-gzip` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-15-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-15-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_gzip_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.7 KiB | [pg_gzip_14-1.0.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_gzip_14-1.0.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_gzip_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_gzip_14-1.0.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_gzip_14-1.0.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_gzip_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_gzip_14-1.0.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_gzip_14-1.0.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_gzip_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_gzip_14-1.0.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_gzip_14-1.0.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_gzip_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_gzip_14-1.0.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_gzip_14-1.0.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_gzip_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 KiB | [pg_gzip_14-1.0.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_gzip_14-1.0.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-gzip` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-gzip` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-gzip` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.3 KiB | [postgresql-14-gzip_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-gzip` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 KiB | [postgresql-14-gzip_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-gzip` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-gzip` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-gzip` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.8 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-gzip` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-14-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-14-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_gzip_13` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.6 KiB | [pg_gzip_13-1.0.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_gzip_13-1.0.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_gzip_13` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [pg_gzip_13-1.0.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_gzip_13-1.0.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_gzip_13` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [pg_gzip_13-1.0.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_gzip_13-1.0.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_gzip_13` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_gzip_13-1.0.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_gzip_13-1.0.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_gzip_13` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pg_gzip_13-1.0.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_gzip_13-1.0.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_gzip_13` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 KiB | [pg_gzip_13-1.0.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_gzip_13-1.0.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-gzip` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.6 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-gzip` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.7 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-gzip` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.1 KiB | [postgresql-13-gzip_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-gzip` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 KiB | [postgresql-13-gzip_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-gzip` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.0 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-gzip` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-gzip` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.7 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-gzip` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.6 KiB | [postgresql-13-gzip_1.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsql-gzip/postgresql-13-gzip_1.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-gzip" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-gzip" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsql-gzip-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_gzip;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_gzip;		# install via package name, for the active PG version
pig install gzip;		# install by extension name, for the current active PG version

pig install gzip -v 18;   # install for PG 18
pig install gzip -v 17;   # install for PG 17
pig install gzip -v 16;   # install for PG 16
pig install gzip -v 15;   # install for PG 15
pig install gzip -v 14;   # install for PG 14
pig install gzip -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION gzip;
```


## Usage

Sometimes you just need to compress your `bytea` object before you return it to the client.

Sometimes you receive a compressed `bytea` from the client, and you have to uncompress it before you can work with it.

This extension is for that.

This extension is **not** for storage compression. PostgreSQL already does [tuple compression](https://www.postgresql.org/docs/current/storage-toast.html) on the fly if your tuple gets large enough, manually pre-compressing your data using this function won't make things smaller.


* `gzip(uncompressed BYTEA, [compression_level INTEGER])` returns `BYTEA`
* `gzip(uncompressed TEXT, [compression_level INTEGER])` returns `BYTEA`
* `gunzip(compressed BYTEA)` returns `BYTEA`


### Examples

    > SELECT gzip('this is my this is my this is my this is my text');

                                       gzip
    --------------------------------------------------------------------------
     \x1f8b08000000000000132bc9c82c5600a2dc4a851282ccd48a12002e7a22ff30000000

Wait, what, the compressed output is longer?!? No, it only **looks** that way, because in hex every byte is represented with two hex digits. The original string looks like this in hex:

    > SELECT 'this is my this is my this is my this is my text'::bytea;

                                                   bytea
    ----------------------------------------------------------------------------------------------------
     \x74686973206973206d792074686973206973206d792074686973206973206d792074686973206973206d792074657874

For really long, repetitive things, compression naturally works like a charm:

    > SELECT gzip(repeat('this is my ', 100));

                                                   bytea
    ----------------------------------------------------------------------------------------------------
     \x1f8b08000000000000132bc9c82c5600a2dc4a859251e628739439ca24970900d1341c5c4c040000

To convert a `bytea` back into an equivalent `text` you must use the `encode()` function with the `escape` encoding.

    > SELECT encode('test text'::bytea, 'escape');
       encode
    -----------
     test text

    > SELECT encode(gunzip(gzip('this text has been compressed and then decompressed')), 'escape')

                          encode
    -----------------------------------------------------
     this text has been compressed and then decompressed

