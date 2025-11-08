---
title: "pgx_ulid"
linkTitle: "pgx_ulid"
description: "ulid type and methods"
weight: 4510
categories: ["FUNC"]
width: full
---

ulid type and methods


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4510** | {{< badge content="pgx_ulid" link="https://github.com/pksunkara/pgx_ulid" >}} | {{< ext "pgx_ulid" >}} | `0.2.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "uuid-ossp" >}} {{< ext "pg_hashids" >}} {{< ext "permuteseq" >}} |

> [!Note] manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgx_ulid" >}} | `0.2.1` | {{< bg "18" "pgx_ulid_18" "green" >}} {{< bg "17" "pgx_ulid_17" "green" >}} {{< bg "16" "pgx_ulid_16" "green" >}} {{< bg "15" "pgx_ulid_15" "green" >}} {{< bg "14" "pgx_ulid_14" "green" >}} {{< bg "13" "pgx_ulid_13" "red" >}} | `pgx_ulid_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgx_ulid" >}} | `0.2.1` | {{< bg "18" "postgresql-18-pgx-ulid" "green" >}} {{< bg "17" "postgresql-17-pgx-ulid" "green" >}} {{< bg "16" "postgresql-16-pgx-ulid" "green" >}} {{< bg "15" "postgresql-15-pgx-ulid" "green" >}} {{< bg "14" "postgresql-14-pgx-ulid" "green" >}} {{< bg "13" "postgresql-13-pgx-ulid" "red" >}} | `postgresql-$v-pgx-ulid` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_18` | 0.2.1 | `el8.x86_64` | pigsty | 360.2 KiB | [pgx_ulid_18-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_18-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_18` | 0.2.1 | `el8.aarch64` | pigsty | 248.3 KiB | [pgx_ulid_18-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_18-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_18` | 0.2.1 | `el9.x86_64` | pigsty | 376.4 KiB | [pgx_ulid_18-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_18-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_18` | 0.2.1 | `el9.aarch64` | pigsty | 265.6 KiB | [pgx_ulid_18-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_18-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_18` | 0.2.1 | `el10.x86_64` | pigsty | 376.6 KiB | [pgx_ulid_18-0.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_18-0.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_18` | 0.2.1 | `el10.aarch64` | pigsty | 265.1 KiB | [pgx_ulid_18-0.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_18-0.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgx-ulid` | 0.2.1 | `d12.x86_64` | pigsty | 2.4 KiB | [postgresql-18-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgx-ulid` | 0.2.1 | `d12.aarch64` | pigsty | 2.4 KiB | [postgresql-18-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgx-ulid` | 0.2.1 | `d13.x86_64` | pigsty | 2.4 KiB | [postgresql-18-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgx-ulid` | 0.2.1 | `d13.aarch64` | pigsty | 2.4 KiB | [postgresql-18-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgx-ulid` | 0.2.1 | `u22.x86_64` | pigsty | 2.0 KiB | [postgresql-18-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgx-ulid` | 0.2.1 | `u22.aarch64` | pigsty | 2.0 KiB | [postgresql-18-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgx-ulid` | 0.2.1 | `u24.x86_64` | pigsty | 2.0 KiB | [postgresql-18-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgx-ulid` | 0.2.1 | `u24.aarch64` | pigsty | 2.0 KiB | [postgresql-18-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_17` | 0.2.1 | `el8.x86_64` | pigsty | 360.1 KiB | [pgx_ulid_17-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_17-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_17` | 0.2.1 | `el8.aarch64` | pigsty | 248.3 KiB | [pgx_ulid_17-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_17-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_17` | 0.2.1 | `el9.x86_64` | pigsty | 376.4 KiB | [pgx_ulid_17-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_17-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_17` | 0.2.1 | `el9.aarch64` | pigsty | 265.8 KiB | [pgx_ulid_17-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_17-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_17` | 0.2.1 | `el10.x86_64` | pigsty | 376.8 KiB | [pgx_ulid_17-0.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_17-0.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_17` | 0.2.1 | `el10.aarch64` | pigsty | 265.0 KiB | [pgx_ulid_17-0.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_17-0.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgx-ulid` | 0.2.1 | `d12.x86_64` | pigsty | 297.0 KiB | [postgresql-17-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.1 | `d12.aarch64` | pigsty | 192.0 KiB | [postgresql-17-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.1 | `d13.x86_64` | pigsty | 297.3 KiB | [postgresql-17-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.1 | `d13.aarch64` | pigsty | 192.1 KiB | [postgresql-17-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.1 | `u22.x86_64` | pigsty | 334.6 KiB | [postgresql-17-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.1 | `u22.aarch64` | pigsty | 224.3 KiB | [postgresql-17-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.1 | `u24.x86_64` | pigsty | 331.5 KiB | [postgresql-17-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgx-ulid` | 0.2.1 | `u24.aarch64` | pigsty | 222.3 KiB | [postgresql-17-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_16` | 0.2.1 | `el8.x86_64` | pigsty | 360.6 KiB | [pgx_ulid_16-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_16-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_16` | 0.2.1 | `el8.aarch64` | pigsty | 248.3 KiB | [pgx_ulid_16-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_16-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_16` | 0.2.1 | `el9.x86_64` | pigsty | 376.3 KiB | [pgx_ulid_16-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_16-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_16` | 0.2.1 | `el9.aarch64` | pigsty | 265.5 KiB | [pgx_ulid_16-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_16-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_16` | 0.2.1 | `el10.x86_64` | pigsty | 376.7 KiB | [pgx_ulid_16-0.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_16-0.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_16` | 0.2.1 | `el10.aarch64` | pigsty | 265.3 KiB | [pgx_ulid_16-0.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_16-0.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgx-ulid` | 0.2.1 | `d12.x86_64` | pigsty | 296.9 KiB | [postgresql-16-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.1 | `d12.aarch64` | pigsty | 192.1 KiB | [postgresql-16-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.1 | `d13.x86_64` | pigsty | 296.7 KiB | [postgresql-16-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.1 | `d13.aarch64` | pigsty | 192.0 KiB | [postgresql-16-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.1 | `u22.x86_64` | pigsty | 334.3 KiB | [postgresql-16-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.1 | `u22.aarch64` | pigsty | 224.2 KiB | [postgresql-16-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.1 | `u24.x86_64` | pigsty | 331.5 KiB | [postgresql-16-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgx-ulid` | 0.2.1 | `u24.aarch64` | pigsty | 222.3 KiB | [postgresql-16-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_15` | 0.2.1 | `el8.x86_64` | pigsty | 359.9 KiB | [pgx_ulid_15-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_15-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_15` | 0.2.1 | `el8.aarch64` | pigsty | 248.4 KiB | [pgx_ulid_15-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_15-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_15` | 0.2.1 | `el9.x86_64` | pigsty | 376.3 KiB | [pgx_ulid_15-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_15-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_15` | 0.2.1 | `el9.aarch64` | pigsty | 265.7 KiB | [pgx_ulid_15-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_15-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_15` | 0.2.1 | `el10.x86_64` | pigsty | 376.3 KiB | [pgx_ulid_15-0.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_15-0.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_15` | 0.2.1 | `el10.aarch64` | pigsty | 265.0 KiB | [pgx_ulid_15-0.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_15-0.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgx-ulid` | 0.2.1 | `d12.x86_64` | pigsty | 296.9 KiB | [postgresql-15-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.1 | `d12.aarch64` | pigsty | 192.0 KiB | [postgresql-15-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.1 | `d13.x86_64` | pigsty | 296.8 KiB | [postgresql-15-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.1 | `d13.aarch64` | pigsty | 192.0 KiB | [postgresql-15-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.1 | `u22.x86_64` | pigsty | 334.7 KiB | [postgresql-15-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.1 | `u22.aarch64` | pigsty | 224.3 KiB | [postgresql-15-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.1 | `u24.x86_64` | pigsty | 331.7 KiB | [postgresql-15-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgx-ulid` | 0.2.1 | `u24.aarch64` | pigsty | 222.4 KiB | [postgresql-15-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_14` | 0.2.1 | `el8.x86_64` | pigsty | 359.2 KiB | [pgx_ulid_14-0.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_14-0.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_14` | 0.2.1 | `el8.aarch64` | pigsty | 247.8 KiB | [pgx_ulid_14-0.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_14-0.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_14` | 0.2.1 | `el9.x86_64` | pigsty | 375.4 KiB | [pgx_ulid_14-0.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_14-0.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_14` | 0.2.1 | `el9.aarch64` | pigsty | 265.3 KiB | [pgx_ulid_14-0.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_14-0.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_14` | 0.2.1 | `el10.x86_64` | pigsty | 376.1 KiB | [pgx_ulid_14-0.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_14-0.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_14` | 0.2.1 | `el10.aarch64` | pigsty | 265.0 KiB | [pgx_ulid_14-0.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_14-0.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgx-ulid` | 0.2.1 | `d12.x86_64` | pigsty | 296.3 KiB | [postgresql-14-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.1 | `d12.aarch64` | pigsty | 191.8 KiB | [postgresql-14-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.1 | `d13.x86_64` | pigsty | 295.9 KiB | [postgresql-14-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.1 | `d13.aarch64` | pigsty | 191.8 KiB | [postgresql-14-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.1 | `u22.x86_64` | pigsty | 333.9 KiB | [postgresql-14-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.1 | `u22.aarch64` | pigsty | 223.6 KiB | [postgresql-14-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.1 | `u24.x86_64` | pigsty | 330.7 KiB | [postgresql-14-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgx-ulid` | 0.2.1 | `u24.aarch64` | pigsty | 222.0 KiB | [postgresql-14-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pksunkara/pgx_ulid" title="Repository" icon="github" subtitle="github.com/pksunkara/pgx_ulid" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgx_ulid-0.2.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgx_ulid; # get pgx_ulid source code
pig build dep pgx_ulid; # install build dependencies
pig build pkg pgx_ulid; # build extension rpm or deb
pig build ext pgx_ulid; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgx_ulid; # install by extension name, for the current active PG version
pig ext install pgx_ulid; # install via package alias, for the active PG version
pig ext install pgx_ulid -v 18;   # install for PG 18
pig ext install pgx_ulid -v 17;   # install for PG 17
pig ext install pgx_ulid -v 16;   # install for PG 16
pig ext install pgx_ulid -v 15;   # install for PG 15
pig ext install pgx_ulid -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgx_ulid;
```

