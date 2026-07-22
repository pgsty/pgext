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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_idkit` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "pg_idkit_18" "green" >}} {{< bg "17" "pg_idkit_17" "green" >}} {{< bg "16" "pg_idkit_16" "green" >}} {{< bg "15" "pg_idkit_15" "green" >}} {{< bg "14" "pg_idkit_14" "green" >}} | `pg_idkit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.0` | {{< bg "18" "postgresql-18-pg-idkit" "green" >}} {{< bg "17" "postgresql-17-pg-idkit" "green" >}} {{< bg "16" "postgresql-16-pg-idkit" "green" >}} {{< bg "15" "postgresql-15-pg-idkit" "green" >}} {{< bg "14" "postgresql-14-pg-idkit" "green" >}} | `postgresql-$v-pg-idkit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "pg_idkit_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-18-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-17-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-16-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-15-pg-idkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.0" "postgresql-14-pg-idkit : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_18` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_idkit_18-0.4.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_18-0.4.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pg_idkit_18-0.4.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_18-0.4.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_idkit_18-0.4.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_18-0.4.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_idkit_18-0.4.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_18-0.4.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_idkit_18-0.4.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_18-0.4.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_18` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pg_idkit_18-0.4.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_18-0.4.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 910.5 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 788.3 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 909.9 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 788.8 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1004.9 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 919.3 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 993.6 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 910.2 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 990.4 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-idkit` | `0.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 908.0 KiB | [postgresql-18-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-18-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_17` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_idkit_17-0.4.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_17-0.4.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pg_idkit_17-0.4.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_17-0.4.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_idkit_17-0.4.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_17-0.4.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_idkit_17-0.4.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_17-0.4.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_idkit_17-0.4.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_17-0.4.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_17` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pg_idkit_17-0.4.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_17-0.4.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 908.0 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 787.1 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 908.2 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 787.9 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1001.8 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 917.8 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 991.6 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 909.3 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 988.5 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-idkit` | `0.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 905.7 KiB | [postgresql-17-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-17-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_16` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_idkit_16-0.4.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_16-0.4.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.0 MiB | [pg_idkit_16-0.4.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_16-0.4.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_idkit_16-0.4.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_16-0.4.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_idkit_16-0.4.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_16-0.4.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_idkit_16-0.4.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_16-0.4.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_16` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pg_idkit_16-0.4.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_16-0.4.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 907.6 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 786.8 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 907.5 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 788.2 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1002.4 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 917.0 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 990.7 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 907.8 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 988.4 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-idkit` | `0.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 906.2 KiB | [postgresql-16-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-16-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_15` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_idkit_15-0.4.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_15-0.4.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1021.1 KiB | [pg_idkit_15-0.4.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_15-0.4.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_idkit_15-0.4.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_15-0.4.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_idkit_15-0.4.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_15-0.4.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_idkit_15-0.4.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_15-0.4.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_15` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pg_idkit_15-0.4.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_15-0.4.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 902.1 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 781.6 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 901.7 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 782.4 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 993.7 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 910.7 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 983.6 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 901.6 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 979.0 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-idkit` | `0.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 899.6 KiB | [postgresql-15-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-15-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_idkit_14` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pg_idkit_14-0.4.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_idkit_14-0.4.0-3PIGSTY.el8.x86_64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1018.9 KiB | [pg_idkit_14-0.4.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_idkit_14-0.4.0-3PIGSTY.el8.aarch64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_idkit_14-0.4.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_idkit_14-0.4.0-3PIGSTY.el9.x86_64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.1 MiB | [pg_idkit_14-0.4.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_idkit_14-0.4.0-3PIGSTY.el9.aarch64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_idkit_14-0.4.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_idkit_14-0.4.0-3PIGSTY.el10.x86_64.rpm) |
| `pg_idkit_14` | `0.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.0 MiB | [pg_idkit_14-0.4.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_idkit_14-0.4.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-idkit` | `0.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 899.5 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 779.8 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 899.9 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 781.0 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 991.8 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 908.5 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 981.7 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 899.9 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 978.8 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-idkit` | `0.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 896.7 KiB | [postgresql-14-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-idkit/postgresql-14-pg-idkit_0.4.0-3PIGSTY~resolute_arm64.deb) |

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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_idkit;
```




## Usage

> [pg_idkit: multi-tool for generating new/niche universally unique identifiers](https://github.com/VADOSWARE/pg_idkit)

```sql
CREATE EXTENSION pg_idkit;
SELECT idkit_uuidv7_generate();
```

### Available Functions

| Methodology | Function | Description |
|---|---|---|
| UUID v6 | `idkit_uuidv6_generate()` | UUID v6 (RFC 4122) |
| | `idkit_uuidv6_generate_uuid()` | UUID v6 as native UUID type |
| | `idkit_uuidv6_extract_timestamptz(TEXT)` | Extract timestamp from UUID v6 |
| UUID v7 | `idkit_uuidv7_generate()` | UUID v7 (RFC 4122) |
| | `idkit_uuidv7_generate_uuid()` | UUID v7 as native UUID type |
| | `idkit_uuidv7_extract_timestamptz(TEXT)` | Extract timestamp from UUID v7 |
| NanoID | `idkit_nanoid_generate()` | NanoID |
| | `idkit_nanoid_custom_generate_text()` | NanoID with custom length and alphabet |
| KSUID | `idkit_ksuid_generate()` | K-Sortable UID |
| | `idkit_ksuid_extract_timestamptz(TEXT)` | Extract timestamp from KSUID |
| | `idkit_ksuidms_generate()` | KSUID with millisecond precision |
| | `idkit_ksuidms_extract_timestamptz(TEXT)` | Extract timestamp from KSUID-ms |
| ULID | `idkit_ulid_generate()` | Universally Unique Lexicographically Sortable ID |
| | `idkit_ulid_extract_timestamptz(TEXT)` | Extract timestamp from ULID |
| Timeflake | `idkit_timeflake_generate()` | Snowflake + Instagram ID + Firebase PushID |
| | `idkit_timeflake_extract_timestamptz(TEXT)` | Extract timestamp from Timeflake |
| PushID | `idkit_pushid_generate()` | Google Firebase PushID |
| XID | `idkit_xid_generate()` | XID |
| | `idkit_xid_extract_timestamptz(TEXT)` | Extract timestamp from XID |
| CUID | `idkit_cuid_generate()` | CUID (deprecated) |
| | `idkit_cuid_extract_timestamptz(TEXT)` | Extract timestamp from CUID |
| CUID2 | `idkit_cuid2_generate()` | CUID2 |
| | `idkit_cuid2_generate_with_len(length)` | CUID2 with custom length |
| TypeID | `idkit_typeid_generate(TEXT)` | TypeID with prefix and UUIDv7 |
| | `idkit_typeid_generate_text(TEXT)` | TypeID as text |
| | `idkit_typeid_from_uuid_v7(TEXT, TEXT)` | TypeID from a given UUID v7 |
| | `idkit_typeid_extract_timestamptz(TEXT)` | Extract timestamp from TypeID |

### Examples

```sql
-- Generate different ID types
SELECT idkit_uuidv7_generate();           -- 018c106f-9304-79bb-b5be-4483b92b036c
SELECT idkit_nanoid_generate();            -- A8jFA0r3NC6FdalR4LEJ0
SELECT idkit_ksuid_generate();             -- 2HMQIBkTJmEN11JI7tvSTMwfYI3
SELECT idkit_ulid_generate();              -- 01HPYV2X17GM5SQP22M3DVFZY3
SELECT idkit_cuid2_generate();             -- clrjx3bwh0000fj3x4c2y1z0s

-- Extract timestamp
SELECT idkit_uuidv7_extract_timestamptz('018c106f-9304-79bb-b5be-4483b92b036c');
```
