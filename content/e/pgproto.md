---
title: "pgproto"
linkTitle: "pgproto"
description: "Native Protobuf parsing, mutation, indexing, and JSON conversion support"
weight: 4130
categories: ["UTIL"]
width: full
---

[**pgproto**](https://github.com/Apaezmx/pgproto) : Native Protobuf parsing, mutation, indexing, and JSON conversion support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4130** | {{< badge content="pgproto" link="https://github.com/Apaezmx/pgproto" >}} | {{< ext "pgproto" >}} | `0.5.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_protobuf" >}} {{< ext "pg_jsonschema" >}} {{< ext "pg_csv" >}} |

> [!Note] release 0.3.3; SQL v1.0


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgproto` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "pgproto_18" "green" >}} {{< bg "17" "pgproto_17" "green" >}} {{< bg "16" "pgproto_16" "green" >}} {{< bg "15" "pgproto_15" "green" >}} {{< bg "14" "pgproto_14" "green" >}} | `pgproto_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.0` | {{< bg "18" "postgresql-18-pgproto" "green" >}} {{< bg "17" "postgresql-17-pgproto" "green" >}} {{< bg "16" "postgresql-16-pgproto" "green" >}} {{< bg "15" "postgresql-15-pgproto" "green" >}} {{< bg "14" "postgresql-14-pgproto" "green" >}} | `postgresql-$v-pgproto` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.0" "postgresql-14-pgproto : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_18` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.4 KiB | [pgproto_18-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_18-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_18` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.1 KiB | [pgproto_18-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_18-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_18` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.1 KiB | [pgproto_18-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_18-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_18` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 KiB | [pgproto_18-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_18-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_18` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.7 KiB | [pgproto_18-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_18-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_18` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.6 KiB | [pgproto_18-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_18-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgproto` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.4 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.8 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.3 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.9 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 54.7 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 53.7 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.8 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.6 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.2 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgproto` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.4 KiB | [postgresql-18-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-18-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_17` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.4 KiB | [pgproto_17-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_17-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_17` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.1 KiB | [pgproto_17-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_17-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_17` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.1 KiB | [pgproto_17-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_17-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_17` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 KiB | [pgproto_17-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_17-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_17` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.7 KiB | [pgproto_17-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_17-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_17` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.5 KiB | [pgproto_17-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_17-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgproto` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.4 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.8 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.3 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.9 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.5 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.3 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.7 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.6 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.2 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgproto` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.4 KiB | [postgresql-17-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-17-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_16` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.4 KiB | [pgproto_16-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_16-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_16` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.1 KiB | [pgproto_16-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_16-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_16` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.1 KiB | [pgproto_16-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_16-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_16` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 KiB | [pgproto_16-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_16-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_16` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.7 KiB | [pgproto_16-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_16-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_16` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.5 KiB | [pgproto_16-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_16-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgproto` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.4 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.8 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.3 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 50.9 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.5 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.3 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.7 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.6 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.2 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgproto` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.4 KiB | [postgresql-16-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-16-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_15` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.5 KiB | [pgproto_15-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_15-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_15` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.2 KiB | [pgproto_15-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_15-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_15` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.3 KiB | [pgproto_15-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_15-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_15` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.1 KiB | [pgproto_15-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_15-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_15` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.0 KiB | [pgproto_15-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_15-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_15` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.6 KiB | [pgproto_15-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_15-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgproto` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.5 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.8 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.4 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.0 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.6 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.4 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.7 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.6 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.2 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgproto` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.4 KiB | [postgresql-15-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-15-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_14` | `0.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.4 KiB | [pgproto_14-0.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_14-0.5.0-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_14` | `0.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.2 KiB | [pgproto_14-0.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_14-0.5.0-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_14` | `0.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.3 KiB | [pgproto_14-0.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_14-0.5.0-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_14` | `0.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 KiB | [pgproto_14-0.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_14-0.5.0-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_14` | `0.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.0 KiB | [pgproto_14-0.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_14-0.5.0-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_14` | `0.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.6 KiB | [pgproto_14-0.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_14-0.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgproto` | `0.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.4 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 50.8 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.4 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.0 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 57.5 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.3 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.7 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.6 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 54.2 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgproto` | `0.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 53.4 KiB | [postgresql-14-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgproto/postgresql-14-pgproto_0.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Apaezmx/pgproto" title="Repository" icon="github" subtitle="github.com/Apaezmx/pgproto" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgproto-0.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgproto;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgproto;		# install via package name, for the active PG version

pig install pgproto -v 18;   # install for PG 18
pig install pgproto -v 17;   # install for PG 17
pig install pgproto -v 16;   # install for PG 16
pig install pgproto -v 15;   # install for PG 15
pig install pgproto -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgproto;
```

## Usage

Sources: [README](https://github.com/Apaezmx/pgproto/blob/v0.5.0/README.md), [release 0.5.0](https://github.com/Apaezmx/pgproto/releases/tag/v0.5.0), [PGXN 0.5.0](https://pgxn.org/dist/pgproto/0.5.0/), [SQL definitions](https://github.com/Apaezmx/pgproto/blob/v0.5.0/sql/pgproto--1.0.sql), [Makefile](https://github.com/Apaezmx/pgproto/blob/v0.5.0/Makefile), [pgproto.control](https://github.com/Apaezmx/pgproto/blob/v0.5.0/pgproto.control)

`pgproto` stores Protocol Buffers `proto3` payloads in PostgreSQL as a native `protobuf` type, with schema-aware extraction, update helpers, containment/index support, and text/integer path operators. The upstream package version is `0.5.0`; the extension SQL/control default version remains `1.0`.

The current upstream source is a C/PGXS extension: the official `Makefile` sets `MODULE_big = pgproto`, builds C objects from `src/*.o`, and includes `$(PGXS)`. The README describes the implementation as pure C with no external Protobuf library dependency.

```sql
CREATE EXTENSION pgproto;
```

### Schema Registry and Storage

`pgproto` needs runtime protobuf descriptors before name/path-based extraction can interpret a binary payload. Register a serialized `FileDescriptorSet` in `pb_schemas`, or call the SQL registration helper when that fits your workflow:

```sql
INSERT INTO pb_schemas (name, data)
VALUES ('MySchema', '\x...');

SELECT pb_register_schema('MySchema', '\x...');
```

Store serialized protobuf bytes in a `protobuf` column:

```sql
CREATE TABLE items (
  id serial PRIMARY KEY,
  data protobuf
);

INSERT INTO items (data) VALUES ('\x0a02082a');
```

The 0.5.0 SQL also installs a convenience cast from `protobuf` to `bytea`, so byte-oriented functions such as `length(data::bytea)` can be used when needed.

### Querying

Use the path operators for nested, repeated, and map fields:

```sql
-- Integer accessor: returns int4
SELECT data #> '{Outer, inner, id}'::text[] FROM items;

-- Text accessor: returns text
SELECT data #>> '{Outer, tags, mykey}'::text[] FROM items;

-- Array index lookup
SELECT data #> '{Outer, scores, 0}'::text[] FROM items;
```

Other user-facing extraction helpers and operators defined by the extension include:

- `pb_get_int32(protobuf, int4)` for tag-based `int4` extraction.
- `pb_get_int32_by_name(protobuf, text, text)` and `pb_get_int32_by_name_dot(protobuf, text)` for name-based integer extraction.
- `->` as shorthand for dot-path integer lookup through `pb_get_int32_by_name_dot`.
- `pb_get_int32_by_path(protobuf, text[])` behind `#>`.
- `pb_get_text_by_path(protobuf, text[])` behind `#>>`.
- `pb_to_json(protobuf, text)` for text JSON conversion when a message name is supplied.

### Updates and Merge

`pb_set`, `pb_insert`, and `pb_delete` are pure functions: they return a new `protobuf` value, so persist changes with `UPDATE ... SET`. Upstream 0.5.0 documents automatic compaction for these mutations to remove stale tags.

```sql
UPDATE items
SET data = pb_set(data, ARRAY['Outer', 'a'], '42');

UPDATE items
SET data = pb_insert(data, ARRAY['Outer', 'scores', '0'], '100');

UPDATE items
SET data = pb_insert(data, ARRAY['Outer', 'tags', 'key1'], 'value1');

UPDATE items
SET data = pb_delete(data, ARRAY['Outer', 'a']);
```

Merge two protobuf values with the `||` operator, which calls `pb_merge`:

```sql
UPDATE items
SET data = data || other.data
FROM other
WHERE items.id = other.id;
```

### Indexing and Containment

Use ordinary expression indexes on extracted fields:

```sql
CREATE INDEX idx_items_pb_id
ON items ((data #> '{Outer, inner, id}'::text[]));

SELECT *
FROM items
WHERE (data #> '{Outer, inner, id}'::text[]) = 42;
```

The SQL definitions also expose protobuf containment with `@>` and a default `protobuf_gin_ops` operator class for GIN indexes:

```sql
CREATE INDEX idx_items_data_gin
ON items USING gin (data protobuf_gin_ops);

SELECT * FROM items WHERE data @> '\x0a02082a'::protobuf;
```

### Schema Evolution

The README frames schema evolution as a normal use case: added fields read as `NULL` from older messages, deprecated or unknown fields are skipped during traversal, enums are read as standard varints, and unset `oneof` fields return `NULL`.

### Caveats

- Runtime schemas are required for schema-aware path navigation; without registered descriptors, the extension cannot resolve message field names.
- `#>` returns `int4` and `#>>` returns `text`; choose the operator/function that matches the expected field type.
- Mutator helpers do not update rows in place; the returned value must be assigned back to the column.
- The README benchmark numbers are upstream project benchmarks, not independent performance guarantees.
