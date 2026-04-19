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
| **4130** | {{< badge content="pgproto" link="https://github.com/Apaezmx/pgproto" >}} | {{< ext "pgproto" >}} | `0.3.3` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgproto` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.3` | {{< bg "18" "pgproto_18" "green" >}} {{< bg "17" "pgproto_17" "green" >}} {{< bg "16" "pgproto_16" "green" >}} {{< bg "15" "pgproto_15" "green" >}} {{< bg "14" "pgproto_14" "green" >}} | `pgproto_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.3` | {{< bg "18" "postgresql-18-pgproto" "green" >}} {{< bg "17" "postgresql-17-pgproto" "green" >}} {{< bg "16" "postgresql-16-pgproto" "green" >}} {{< bg "15" "postgresql-15-pgproto" "green" >}} {{< bg "14" "postgresql-14-pgproto" "green" >}} | `postgresql-$v-pgproto` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "pgproto_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-14-pgproto : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-18-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-17-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-16-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-15-pgproto : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.3" "postgresql-14-pgproto : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_18` | `0.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 228.6 KiB | [pgproto_18-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_18-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_18` | `0.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 213.0 KiB | [pgproto_18-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_18-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_18` | `0.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 127.3 KiB | [pgproto_18-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_18-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_18` | `0.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 124.9 KiB | [pgproto_18-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_18-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_18` | `0.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 129.2 KiB | [pgproto_18-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_18-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_18` | `0.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 125.0 KiB | [pgproto_18-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_18-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgproto` | `0.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 697.5 KiB | [postgresql-18-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-18-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgproto` | `0.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 693.2 KiB | [postgresql-18-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-18-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgproto` | `0.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 703.2 KiB | [postgresql-18-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-18-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgproto` | `0.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 696.1 KiB | [postgresql-18-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-18-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgproto` | `0.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 755.0 KiB | [postgresql-18-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-18-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgproto` | `0.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 754.3 KiB | [postgresql-18-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-18-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgproto` | `0.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 733.0 KiB | [postgresql-18-pgproto_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-18-pgproto_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgproto` | `0.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 742.9 KiB | [postgresql-18-pgproto_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-18-pgproto_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_17` | `0.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 228.6 KiB | [pgproto_17-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_17-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_17` | `0.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 213.0 KiB | [pgproto_17-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_17-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_17` | `0.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 127.3 KiB | [pgproto_17-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_17-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_17` | `0.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 124.9 KiB | [pgproto_17-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_17-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_17` | `0.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 129.2 KiB | [pgproto_17-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_17-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_17` | `0.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 125.0 KiB | [pgproto_17-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_17-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgproto` | `0.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 697.4 KiB | [postgresql-17-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-17-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgproto` | `0.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 693.0 KiB | [postgresql-17-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-17-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgproto` | `0.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 702.9 KiB | [postgresql-17-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-17-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgproto` | `0.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 696.3 KiB | [postgresql-17-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-17-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgproto` | `0.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 814.3 KiB | [postgresql-17-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-17-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgproto` | `0.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 812.8 KiB | [postgresql-17-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-17-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgproto` | `0.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 733.0 KiB | [postgresql-17-pgproto_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-17-pgproto_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgproto` | `0.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 742.8 KiB | [postgresql-17-pgproto_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-17-pgproto_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_16` | `0.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 228.7 KiB | [pgproto_16-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_16-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_16` | `0.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 213.0 KiB | [pgproto_16-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_16-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_16` | `0.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 127.3 KiB | [pgproto_16-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_16-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_16` | `0.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 124.9 KiB | [pgproto_16-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_16-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_16` | `0.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 129.2 KiB | [pgproto_16-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_16-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_16` | `0.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 125.0 KiB | [pgproto_16-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_16-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgproto` | `0.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 697.3 KiB | [postgresql-16-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-16-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgproto` | `0.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 693.1 KiB | [postgresql-16-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-16-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgproto` | `0.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 703.0 KiB | [postgresql-16-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-16-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgproto` | `0.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 696.2 KiB | [postgresql-16-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-16-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgproto` | `0.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 814.3 KiB | [postgresql-16-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-16-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgproto` | `0.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 812.8 KiB | [postgresql-16-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-16-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgproto` | `0.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 733.0 KiB | [postgresql-16-pgproto_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-16-pgproto_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgproto` | `0.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 742.9 KiB | [postgresql-16-pgproto_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-16-pgproto_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_15` | `0.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 239.7 KiB | [pgproto_15-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_15-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_15` | `0.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 223.5 KiB | [pgproto_15-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_15-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_15` | `0.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 219.1 KiB | [pgproto_15-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_15-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_15` | `0.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 212.0 KiB | [pgproto_15-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_15-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_15` | `0.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 223.4 KiB | [pgproto_15-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_15-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_15` | `0.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 213.7 KiB | [pgproto_15-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_15-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgproto` | `0.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 707.0 KiB | [postgresql-15-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-15-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgproto` | `0.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 702.2 KiB | [postgresql-15-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-15-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgproto` | `0.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 712.7 KiB | [postgresql-15-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-15-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgproto` | `0.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 705.8 KiB | [postgresql-15-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-15-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgproto` | `0.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 822.2 KiB | [postgresql-15-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-15-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgproto` | `0.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 819.9 KiB | [postgresql-15-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-15-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgproto` | `0.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 741.3 KiB | [postgresql-15-pgproto_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-15-pgproto_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgproto` | `0.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 749.8 KiB | [postgresql-15-pgproto_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-15-pgproto_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgproto_14` | `0.3.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 239.7 KiB | [pgproto_14-0.3.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgproto_14-0.3.3-1PIGSTY.el8.x86_64.rpm) |
| `pgproto_14` | `0.3.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 223.5 KiB | [pgproto_14-0.3.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgproto_14-0.3.3-1PIGSTY.el8.aarch64.rpm) |
| `pgproto_14` | `0.3.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 219.2 KiB | [pgproto_14-0.3.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgproto_14-0.3.3-1PIGSTY.el9.x86_64.rpm) |
| `pgproto_14` | `0.3.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 212.0 KiB | [pgproto_14-0.3.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgproto_14-0.3.3-1PIGSTY.el9.aarch64.rpm) |
| `pgproto_14` | `0.3.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 223.0 KiB | [pgproto_14-0.3.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgproto_14-0.3.3-1PIGSTY.el10.x86_64.rpm) |
| `pgproto_14` | `0.3.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 215.5 KiB | [pgproto_14-0.3.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgproto_14-0.3.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgproto` | `0.3.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 706.9 KiB | [postgresql-14-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-14-pgproto_0.3.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgproto` | `0.3.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 701.9 KiB | [postgresql-14-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgproto/postgresql-14-pgproto_0.3.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgproto` | `0.3.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 712.5 KiB | [postgresql-14-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-14-pgproto_0.3.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgproto` | `0.3.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 705.7 KiB | [postgresql-14-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgproto/postgresql-14-pgproto_0.3.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgproto` | `0.3.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 822.2 KiB | [postgresql-14-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-14-pgproto_0.3.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgproto` | `0.3.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 819.9 KiB | [postgresql-14-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgproto/postgresql-14-pgproto_0.3.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgproto` | `0.3.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 741.3 KiB | [postgresql-14-pgproto_0.3.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-14-pgproto_0.3.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgproto` | `0.3.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 749.8 KiB | [postgresql-14-pgproto_0.3.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgproto/postgresql-14-pgproto_0.3.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Apaezmx/pgproto" title="Repository" icon="github" subtitle="github.com/Apaezmx/pgproto" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgproto-0.3.3.tar.gz" >}}
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

> Syntax:
>
> ```sql
> CREATE EXTENSION pgproto;
> INSERT INTO pb_schemas (name, data) VALUES ('MySchema', '\x...');
> CREATE TABLE items (id serial PRIMARY KEY, data protobuf);
> SELECT data #> '{Outer, inner, id}'::text[] FROM items;
> ```
>
> Source: [README](https://github.com/Apaezmx/pgproto)

`pgproto` adds native Protocol Buffers support to PostgreSQL. It provides a `protobuf` type, runtime schema registration, nested field extraction, update helpers, and indexing support for schema-aware access to protobuf payloads.

## Setup

Enable the extension:

```sql
CREATE EXTENSION pgproto;
```

Register protobuf schemas by loading `FileDescriptorSet` blobs:

```sql
INSERT INTO pb_schemas (name, data) VALUES ('MySchema', '\x...');
```

Create a table using the custom `protobuf` type:

```sql
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    data protobuf
);
```

## Querying

The README highlights nested field extraction with PostgreSQL-style operators:

```sql
SELECT data #> '{Outer, inner, id}'::text[] FROM items;
SELECT data #> '{Outer, tags, mykey}'::text[] FROM items;
```

It also mentions custom operators such as `->` and `#>` for schema-aware navigation.

## Modification Functions

`pgproto` includes pure functions that return a new protobuf value:

- `pb_set(...)`
- `pb_insert(...)`
- `pb_delete(...)`

Because they return modified values rather than mutating in place, they are normally used in `UPDATE` statements:

```sql
UPDATE items SET data = pb_set(data, ARRAY['Outer', 'a'], '42');
UPDATE items SET data = pb_insert(data, ARRAY['Outer', 'scores', '0'], '100');
UPDATE items SET data = pb_delete(data, ARRAY['Outer', 'a']);
```

The `||` operator merges two protobuf messages of the same type.

## Indexing

The README documents B-tree expression indexes on extracted fields:

```sql
CREATE INDEX idx_pb_id ON items ((data #> '{Outer, inner, id}'::text[]));
```

The project also advertises GIN support for retrieval workflows.

## Notes

The upstream README positions `pgproto` as more storage-efficient than JSONB for protobuf-native payloads and highlights protobuf schema evolution, enums, `oneof`, and map/repeated field access as supported use cases.
