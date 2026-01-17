---
title: "pg_idkit"
linkTitle: "pg_idkit"
description: "multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID)"
weight: 4500
categories: ["FUNC"]
width: full
---

[**pg_idkit**](https://github.com/VADOSWARE/pg_idkit) : multi-tool for generating new/niche universally unique identifiers (ex. UUIDv6, ULID, KSUID)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4500** | {{< badge content="pg_idkit" link="https://github.com/VADOSWARE/pg_idkit" >}} | {{< ext "pg_idkit" >}} | `0.4.0` | {{< category "FUNC" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "uuid-ossp" >}} {{< ext "permuteseq" >}} {{< ext "pg_cardano" >}} {{< ext "pg_base58" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_idkit` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "pg_idkit_18" "green" >}} {{< bg "17" "pg_idkit_17" "green" >}} {{< bg "16" "pg_idkit_16" "green" >}} {{< bg "15" "pg_idkit_15" "green" >}} {{< bg "14" "pg_idkit_14" "green" >}} {{< bg "13" "pg_idkit_13" "green" >}} | `pg_idkit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "postgresql-18-pg-idkit" "green" >}} {{< bg "17" "postgresql-17-pg-idkit" "green" >}} {{< bg "16" "postgresql-16-pg-idkit" "green" >}} {{< bg "15" "postgresql-15-pg-idkit" "green" >}} {{< bg "14" "postgresql-14-pg-idkit" "green" >}} {{< bg "13" "postgresql-13-pg-idkit" "green" >}} | `postgresql-$v-pg-idkit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-13-pg-idkit : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_18` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 459.4 KiB | [pg_idkit_18-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_18-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 355.0 KiB | [pg_idkit_18-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_18-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 475.0 KiB | [pg_idkit_18-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_18-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 377.4 KiB | [pg_idkit_18-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_18-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 474.5 KiB | [pg_idkit_18-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_18-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 376.8 KiB | [pg_idkit_18-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_18-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 385.9 KiB | [postgresql-18-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 285.4 KiB | [postgresql-18-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 386.1 KiB | [postgresql-18-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 285.3 KiB | [postgresql-18-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 428.7 KiB | [postgresql-18-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 328.4 KiB | [postgresql-18-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 424.2 KiB | [postgresql-18-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 324.3 KiB | [postgresql-18-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_17` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 459.7 KiB | [pg_idkit_17-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_17-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 355.0 KiB | [pg_idkit_17-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_17-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 474.7 KiB | [pg_idkit_17-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_17-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 377.0 KiB | [pg_idkit_17-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_17-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 475.2 KiB | [pg_idkit_17-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_17-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 376.6 KiB | [pg_idkit_17-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_17-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 385.6 KiB | [postgresql-17-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 285.2 KiB | [postgresql-17-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 386.1 KiB | [postgresql-17-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 285.2 KiB | [postgresql-17-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 428.1 KiB | [postgresql-17-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 328.2 KiB | [postgresql-17-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 423.7 KiB | [postgresql-17-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 324.0 KiB | [postgresql-17-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_16` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 459.2 KiB | [pg_idkit_16-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_16-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 355.5 KiB | [pg_idkit_16-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_16-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 474.8 KiB | [pg_idkit_16-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_16-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 377.2 KiB | [pg_idkit_16-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_16-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 474.7 KiB | [pg_idkit_16-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_16-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 376.6 KiB | [pg_idkit_16-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_16-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 385.6 KiB | [postgresql-16-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 285.0 KiB | [postgresql-16-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 386.1 KiB | [postgresql-16-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 285.3 KiB | [postgresql-16-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 428.6 KiB | [postgresql-16-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 328.1 KiB | [postgresql-16-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 424.3 KiB | [postgresql-16-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 323.6 KiB | [postgresql-16-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_15` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 459.2 KiB | [pg_idkit_15-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_15-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 355.2 KiB | [pg_idkit_15-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_15-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 474.4 KiB | [pg_idkit_15-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_15-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 377.1 KiB | [pg_idkit_15-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_15-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 475.2 KiB | [pg_idkit_15-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_15-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 376.7 KiB | [pg_idkit_15-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_15-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 385.4 KiB | [postgresql-15-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 285.1 KiB | [postgresql-15-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 386.0 KiB | [postgresql-15-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 285.3 KiB | [postgresql-15-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 428.6 KiB | [postgresql-15-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 328.5 KiB | [postgresql-15-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 424.3 KiB | [postgresql-15-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 324.1 KiB | [postgresql-15-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_14` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 458.8 KiB | [pg_idkit_14-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_14-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 355.5 KiB | [pg_idkit_14-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_14-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 474.5 KiB | [pg_idkit_14-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_14-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 377.1 KiB | [pg_idkit_14-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_14-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 474.8 KiB | [pg_idkit_14-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_14-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 376.8 KiB | [pg_idkit_14-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_14-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 385.6 KiB | [postgresql-14-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 285.2 KiB | [postgresql-14-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 385.7 KiB | [postgresql-14-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 285.3 KiB | [postgresql-14-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 428.9 KiB | [postgresql-14-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 328.4 KiB | [postgresql-14-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 424.0 KiB | [postgresql-14-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 324.1 KiB | [postgresql-14-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_13` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 459.2 KiB | [pg_idkit_13-0.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_13-0.4.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_13` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 355.1 KiB | [pg_idkit_13-0.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_13-0.4.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_13` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 475.0 KiB | [pg_idkit_13-0.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_13-0.4.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_13` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 377.0 KiB | [pg_idkit_13-0.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_13-0.4.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_13` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 475.1 KiB | [pg_idkit_13-0.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_13-0.4.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_13` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 376.7 KiB | [pg_idkit_13-0.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_13-0.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 385.6 KiB | [postgresql-13-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 285.2 KiB | [postgresql-13-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 385.9 KiB | [postgresql-13-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 285.3 KiB | [postgresql-13-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 428.9 KiB | [postgresql-13-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 328.3 KiB | [postgresql-13-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 424.5 KiB | [postgresql-13-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 324.1 KiB | [postgresql-13-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-13-pg-idkit_0.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/VADOSWARE/pg_idkit" title="Repository" icon="github" subtitle="github.com/VADOSWARE/pg_idkit" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_idkit-0.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_idkit;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_idkit;		# install via package name, for the active PG version

pig install pg_idkit -v 18;   # install for PG 18
pig install pg_idkit -v 17;   # install for PG 17
pig install pg_idkit -v 16;   # install for PG 16
pig install pg_idkit -v 15;   # install for PG 15
pig install pg_idkit -v 14;   # install for PG 14
pig install pg_idkit -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_idkit;
```
