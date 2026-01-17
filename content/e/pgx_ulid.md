---
title: "pgx_ulid"
linkTitle: "pgx_ulid"
description: "ulid type and methods"
weight: 4510
categories: ["FUNC"]
width: full
---

[**pgx_ulid**](https://github.com/pksunkara/pgx_ulid) : ulid type and methods


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4510** | {{< badge content="pgx_ulid" link="https://github.com/pksunkara/pgx_ulid" >}} | {{< ext "pgx_ulid" >}} | `0.2.2` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "uuid-ossp" >}} {{< ext "pg_hashids" >}} {{< ext "permuteseq" >}} |

> [!Note] manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pgx_ulid` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "pgx_ulid_18" "green" >}} {{< bg "17" "pgx_ulid_17" "green" >}} {{< bg "16" "pgx_ulid_16" "green" >}} {{< bg "15" "pgx_ulid_15" "green" >}} {{< bg "14" "pgx_ulid_14" "green" >}} {{< bg "13" "pgx_ulid_13" "red" >}} | `pgx_ulid_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.2` | {{< bg "18" "postgresql-18-pgx-ulid" "green" >}} {{< bg "17" "postgresql-17-pgx-ulid" "green" >}} {{< bg "16" "postgresql-16-pgx-ulid" "green" >}} {{< bg "15" "postgresql-15-pgx-ulid" "green" >}} {{< bg "14" "postgresql-14-pgx-ulid" "green" >}} {{< bg "13" "postgresql-13-pgx-ulid" "red" >}} | `postgresql-$v-pgx-ulid` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "pgx_ulid_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgx_ulid_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-18-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-17-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-16-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-15-pgx-ulid : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-14-pgx-ulid : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pgx-ulid : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_18` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 382.6 KiB | [pgx_ulid_18-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_18-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_18` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 268.0 KiB | [pgx_ulid_18-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_18-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_18` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 399.2 KiB | [pgx_ulid_18-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_18-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_18` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 286.7 KiB | [pgx_ulid_18-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_18-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_18` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 399.1 KiB | [pgx_ulid_18-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_18-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_18` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 287.0 KiB | [pgx_ulid_18-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_18-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgx-ulid` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 315.7 KiB | [postgresql-18-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgx-ulid` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 207.1 KiB | [postgresql-18-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgx-ulid` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 315.6 KiB | [postgresql-18-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgx-ulid` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 207.2 KiB | [postgresql-18-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgx-ulid` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 354.6 KiB | [postgresql-18-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgx-ulid` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 241.3 KiB | [postgresql-18-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgx-ulid` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 351.3 KiB | [postgresql-18-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgx-ulid` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 239.5 KiB | [postgresql-18-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-18-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_17` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 383.1 KiB | [pgx_ulid_17-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_17-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_17` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 268.1 KiB | [pgx_ulid_17-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_17-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_17` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 398.6 KiB | [pgx_ulid_17-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_17-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_17` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 287.0 KiB | [pgx_ulid_17-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_17-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_17` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 398.8 KiB | [pgx_ulid_17-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_17-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_17` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 286.9 KiB | [pgx_ulid_17-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_17-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgx-ulid` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 315.3 KiB | [postgresql-17-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgx-ulid` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 206.9 KiB | [postgresql-17-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgx-ulid` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 315.4 KiB | [postgresql-17-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgx-ulid` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 207.1 KiB | [postgresql-17-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgx-ulid` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 354.6 KiB | [postgresql-17-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgx-ulid` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 240.7 KiB | [postgresql-17-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgx-ulid` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 351.1 KiB | [postgresql-17-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgx-ulid` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 239.0 KiB | [postgresql-17-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-17-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_16` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 383.3 KiB | [pgx_ulid_16-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_16-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_16` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 268.2 KiB | [pgx_ulid_16-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_16-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_16` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 399.7 KiB | [pgx_ulid_16-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_16-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_16` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 287.0 KiB | [pgx_ulid_16-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_16-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_16` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 399.6 KiB | [pgx_ulid_16-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_16-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_16` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 287.1 KiB | [pgx_ulid_16-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_16-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgx-ulid` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 315.7 KiB | [postgresql-16-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgx-ulid` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 207.0 KiB | [postgresql-16-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgx-ulid` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 315.7 KiB | [postgresql-16-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgx-ulid` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 207.2 KiB | [postgresql-16-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgx-ulid` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 354.7 KiB | [postgresql-16-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgx-ulid` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 241.1 KiB | [postgresql-16-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgx-ulid` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 351.2 KiB | [postgresql-16-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgx-ulid` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 239.5 KiB | [postgresql-16-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-16-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_15` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 382.7 KiB | [pgx_ulid_15-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_15-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_15` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 268.0 KiB | [pgx_ulid_15-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_15-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_15` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 398.2 KiB | [pgx_ulid_15-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_15-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_15` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 286.5 KiB | [pgx_ulid_15-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_15-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_15` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 398.2 KiB | [pgx_ulid_15-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_15-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_15` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 286.9 KiB | [pgx_ulid_15-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_15-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgx-ulid` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 315.5 KiB | [postgresql-15-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgx-ulid` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 206.9 KiB | [postgresql-15-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgx-ulid` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 315.5 KiB | [postgresql-15-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgx-ulid` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 207.1 KiB | [postgresql-15-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgx-ulid` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 354.3 KiB | [postgresql-15-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgx-ulid` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 240.9 KiB | [postgresql-15-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgx-ulid` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 350.9 KiB | [postgresql-15-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgx-ulid` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 238.7 KiB | [postgresql-15-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-15-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgx_ulid_14` | `0.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 381.9 KiB | [pgx_ulid_14-0.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgx_ulid_14-0.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pgx_ulid_14` | `0.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 267.6 KiB | [pgx_ulid_14-0.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgx_ulid_14-0.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pgx_ulid_14` | `0.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 397.8 KiB | [pgx_ulid_14-0.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgx_ulid_14-0.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pgx_ulid_14` | `0.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 286.7 KiB | [pgx_ulid_14-0.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgx_ulid_14-0.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pgx_ulid_14` | `0.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 397.8 KiB | [pgx_ulid_14-0.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgx_ulid_14-0.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pgx_ulid_14` | `0.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 286.7 KiB | [pgx_ulid_14-0.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgx_ulid_14-0.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgx-ulid` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 314.8 KiB | [postgresql-14-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgx-ulid` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 206.7 KiB | [postgresql-14-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgx-ulid` | `0.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 315.0 KiB | [postgresql-14-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgx-ulid` | `0.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 206.7 KiB | [postgresql-14-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgx-ulid` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 353.5 KiB | [postgresql-14-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgx-ulid` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 240.3 KiB | [postgresql-14-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgx-ulid` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 350.5 KiB | [postgresql-14-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgx-ulid` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 238.4 KiB | [postgresql-14-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgx-ulid/postgresql-14-pgx-ulid_0.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pksunkara/pgx_ulid" title="Repository" icon="github" subtitle="github.com/pksunkara/pgx_ulid" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgx_ulid-0.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgx_ulid;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgx_ulid;		# install via package name, for the active PG version

pig install pgx_ulid -v 18;   # install for PG 18
pig install pgx_ulid -v 17;   # install for PG 17
pig install pgx_ulid -v 16;   # install for PG 16
pig install pgx_ulid -v 15;   # install for PG 15
pig install pgx_ulid -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pgx_ulid';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgx_ulid;
```
