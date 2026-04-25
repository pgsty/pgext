---
title: "pg_protobuf"
linkTitle: "pg_protobuf"
description: "Protobuf support for PostgreSQL"
weight: 4120
categories: ["UTIL"]
width: full
---

[**pg_protobuf**](https://github.com/afiskon/pg_protobuf) : Protobuf support for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4120** | {{< badge content="pg_protobuf" link="https://github.com/afiskon/pg_protobuf" >}} | {{< ext "pg_protobuf" >}} | `1.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgjq" >}} {{< ext "pgqr" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_protobuf` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_protobuf_18" "green" >}} {{< bg "17" "pg_protobuf_17" "green" >}} {{< bg "16" "pg_protobuf_16" "green" >}} {{< bg "15" "pg_protobuf_15" "green" >}} {{< bg "14" "pg_protobuf_14" "green" >}} | `pg_protobuf_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-protobuf" "green" >}} {{< bg "17" "postgresql-17-pg-protobuf" "green" >}} {{< bg "16" "postgresql-16-pg-protobuf" "green" >}} {{< bg "15" "postgresql-15-pg-protobuf" "green" >}} {{< bg "14" "postgresql-14-pg-protobuf" "green" >}} | `postgresql-$v-pg-protobuf` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_protobuf_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-protobuf : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-protobuf : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-protobuf : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-protobuf : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-protobuf : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-protobuf : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-protobuf : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-protobuf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-protobuf : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-protobuf : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-protobuf : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-protobuf : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-protobuf : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-protobuf : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-protobuf : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-protobuf : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-protobuf : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-protobuf : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-protobuf : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_protobuf_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.8 KiB | [pg_protobuf_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_protobuf_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_protobuf_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.0 KiB | [pg_protobuf_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_protobuf_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_protobuf_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.1 KiB | [pg_protobuf_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_protobuf_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_protobuf_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.1 KiB | [pg_protobuf_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_protobuf_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_protobuf_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.2 KiB | [pg_protobuf_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_protobuf_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_protobuf_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.2 KiB | [pg_protobuf_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_protobuf_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-protobuf` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.6 KiB | [postgresql-18-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-18-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-protobuf` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 41.5 KiB | [postgresql-18-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-18-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-protobuf` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.6 KiB | [postgresql-18-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-18-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-protobuf` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 41.6 KiB | [postgresql-18-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-18-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-protobuf` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.6 KiB | [postgresql-18-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-18-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-protobuf` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 44.5 KiB | [postgresql-18-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-18-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-protobuf` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 42.5 KiB | [postgresql-18-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-18-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-protobuf` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 42.7 KiB | [postgresql-18-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-18-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_protobuf_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.8 KiB | [pg_protobuf_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_protobuf_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_protobuf_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.0 KiB | [pg_protobuf_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_protobuf_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_protobuf_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.1 KiB | [pg_protobuf_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_protobuf_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_protobuf_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.1 KiB | [pg_protobuf_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_protobuf_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_protobuf_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.2 KiB | [pg_protobuf_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_protobuf_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_protobuf_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.2 KiB | [pg_protobuf_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_protobuf_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-protobuf` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.0 KiB | [postgresql-17-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-17-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-protobuf` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.7 KiB | [postgresql-17-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-17-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-protobuf` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 39.9 KiB | [postgresql-17-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-17-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-protobuf` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.8 KiB | [postgresql-17-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-17-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-protobuf` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.9 KiB | [postgresql-17-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-17-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-protobuf` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 44.8 KiB | [postgresql-17-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-17-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-protobuf` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 41.9 KiB | [postgresql-17-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-17-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-protobuf` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 42.1 KiB | [postgresql-17-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-17-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_protobuf_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.7 KiB | [pg_protobuf_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_protobuf_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_protobuf_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.0 KiB | [pg_protobuf_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_protobuf_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_protobuf_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.1 KiB | [pg_protobuf_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_protobuf_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_protobuf_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.1 KiB | [pg_protobuf_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_protobuf_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_protobuf_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.2 KiB | [pg_protobuf_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_protobuf_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_protobuf_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.2 KiB | [pg_protobuf_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_protobuf_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-protobuf` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 39.8 KiB | [postgresql-16-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-16-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-protobuf` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.5 KiB | [postgresql-16-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-16-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-protobuf` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 39.7 KiB | [postgresql-16-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-16-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-protobuf` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.7 KiB | [postgresql-16-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-16-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-protobuf` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.6 KiB | [postgresql-16-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-16-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-protobuf` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 44.5 KiB | [postgresql-16-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-16-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-protobuf` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 41.6 KiB | [postgresql-16-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-16-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-protobuf` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 41.9 KiB | [postgresql-16-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-16-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_protobuf_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [pg_protobuf_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_protobuf_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_protobuf_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.2 KiB | [pg_protobuf_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_protobuf_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_protobuf_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.5 KiB | [pg_protobuf_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_protobuf_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_protobuf_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.6 KiB | [pg_protobuf_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_protobuf_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_protobuf_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.5 KiB | [pg_protobuf_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_protobuf_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_protobuf_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_protobuf_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_protobuf_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-protobuf` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 38.7 KiB | [postgresql-15-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-15-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-protobuf` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 39.4 KiB | [postgresql-15-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-15-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-protobuf` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.6 KiB | [postgresql-15-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-15-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-protobuf` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 39.5 KiB | [postgresql-15-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-15-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-protobuf` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 43.6 KiB | [postgresql-15-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-15-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-protobuf` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 43.6 KiB | [postgresql-15-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-15-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-protobuf` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 40.7 KiB | [postgresql-15-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-15-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-protobuf` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 41.0 KiB | [postgresql-15-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-15-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_protobuf_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [pg_protobuf_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_protobuf_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_protobuf_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.2 KiB | [pg_protobuf_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_protobuf_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_protobuf_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.5 KiB | [pg_protobuf_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_protobuf_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_protobuf_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.6 KiB | [pg_protobuf_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_protobuf_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_protobuf_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.5 KiB | [pg_protobuf_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_protobuf_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_protobuf_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.8 KiB | [pg_protobuf_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_protobuf_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-protobuf` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 38.5 KiB | [postgresql-14-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-14-pg-protobuf_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-protobuf` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 39.3 KiB | [postgresql-14-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-protobuf/postgresql-14-pg-protobuf_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-protobuf` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 38.5 KiB | [postgresql-14-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-14-pg-protobuf_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-protobuf` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 39.4 KiB | [postgresql-14-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-protobuf/postgresql-14-pg-protobuf_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-protobuf` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 43.5 KiB | [postgresql-14-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-14-pg-protobuf_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-protobuf` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 43.4 KiB | [postgresql-14-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-protobuf/postgresql-14-pg-protobuf_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-protobuf` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 40.5 KiB | [postgresql-14-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-14-pg-protobuf_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-protobuf` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 40.8 KiB | [postgresql-14-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-protobuf/postgresql-14-pg-protobuf_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/afiskon/pg_protobuf" title="Repository" icon="github" subtitle="github.com/afiskon/pg_protobuf" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_protobuf-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_protobuf;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_protobuf;		# install via package name, for the active PG version

pig install pg_protobuf -v 18;   # install for PG 18
pig install pg_protobuf -v 17;   # install for PG 17
pig install pg_protobuf -v 16;   # install for PG 16
pig install pg_protobuf -v 15;   # install for PG 15
pig install pg_protobuf -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_protobuf;
```




## Usage

> [pg_protobuf: Protocol Buffers support for PostgreSQL](https://github.com/afiskon/pg_protobuf)

Provides functions to decode Protocol Buffer binary data directly in SQL without schema definitions.

### Functions

- `protobuf_decode(bytea) RETURNS cstring` -- Decode protobuf binary into a human-readable format
- `protobuf_get_int(bytea, int) RETURNS bigint` -- Extract an integer field by field number
- `protobuf_get_string(bytea, int) RETURNS text` -- Extract a string field by field number
- `protobuf_get_bytes(bytea, int) RETURNS bytea` -- Extract raw bytes by field number
- `protobuf_get_bool(bytea, int) RETURNS boolean` -- Extract a boolean field by field number
- `protobuf_get_float(bytea, int) RETURNS real` -- Extract a float field by field number
- `protobuf_get_double(bytea, int) RETURNS double precision` -- Extract a double field by field number
- `protobuf_get_sfixed32(bytea, int) RETURNS int` -- Extract a signed fixed 32-bit field
- `protobuf_get_sfixed64(bytea, int) RETURNS bigint` -- Extract a signed fixed 64-bit field
- `protobuf_get_int_multi(bytea, int) RETURNS bigint[]` -- Extract repeated integer fields
- `protobuf_get_string_multi(bytea, int) RETURNS text[]` -- Extract repeated string fields
- `protobuf_get_bytes_multi(bytea, int) RETURNS bytea[]` -- Extract repeated bytes fields
- `protobuf_get_bool_multi(bytea, int) RETURNS boolean[]` -- Extract repeated boolean fields
- `protobuf_get_float_multi(bytea, int) RETURNS real[]` -- Extract repeated float fields
- `protobuf_get_double_multi(bytea, int) RETURNS double precision[]` -- Extract repeated double fields
- `protobuf_get_sfixed32_multi(bytea, int) RETURNS int[]` -- Extract repeated sfixed32 fields
- `protobuf_get_sfixed64_multi(bytea, int) RETURNS bigint[]` -- Extract repeated sfixed64 fields

### Example

```sql
CREATE EXTENSION pg_protobuf;

-- Create a table with protobuf data
CREATE TABLE heroes (x bytea);

-- Define accessor functions for specific fields
CREATE FUNCTION hero_name(x bytea) RETURNS text AS $$
BEGIN
  RETURN protobuf_get_string(x, 512);
END $$ LANGUAGE plpgsql IMMUTABLE;

CREATE FUNCTION hero_hp(x bytea) RETURNS bigint AS $$
BEGIN
  RETURN protobuf_get_int(x, 2);
END $$ LANGUAGE plpgsql IMMUTABLE;

-- Create an index on a protobuf field
CREATE INDEX hero_name_idx ON heroes USING btree(hero_name(x));

-- Query by protobuf field
SELECT hero_name(x) FROM heroes ORDER BY hero_name(x) LIMIT 10;
```

### Limitations

- No modification of Protobuf data
- Enums readable via `protobuf_get_int`
- Unsigned types not directly supported (no unsigned integers in PostgreSQL)
- `[packed=true]` not supported by `*_multi` procedures (use `protobuf_get_bytes*` instead)
