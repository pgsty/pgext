---
title: "pg_hashids"
linkTitle: "pg_hashids"
description: "Short unique id generator for PostgreSQL, using hashids"
weight: 4540
categories: ["FUNC"]
width: full
---

[**pg_hashids**](https://github.com/iCyberon/pg_hashids) : Short unique id generator for PostgreSQL, using hashids


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4540** | {{< badge content="pg_hashids" link="https://github.com/iCyberon/pg_hashids" >}} | {{< ext "pg_hashids" >}} | `1.3` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pg_base58" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "permuteseq" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_hashids` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3` | {{< bg "18" "pg_hashids_18" "green" >}} {{< bg "17" "pg_hashids_17" "green" >}} {{< bg "16" "pg_hashids_16" "green" >}} {{< bg "15" "pg_hashids_15" "green" >}} {{< bg "14" "pg_hashids_14" "green" >}} | `pg_hashids_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3` | {{< bg "18" "postgresql-18-pg-hashids" "green" >}} {{< bg "17" "postgresql-17-pg-hashids" "green" >}} {{< bg "16" "postgresql-16-pg-hashids" "green" >}} {{< bg "15" "postgresql-15-pg-hashids" "green" >}} {{< bg "14" "postgresql-14-pg-hashids" "green" >}} | `postgresql-$v-pg-hashids` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "pg_hashids_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-pg-hashids : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-pg-hashids : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-pg-hashids : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-pg-hashids : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-pg-hashids : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-pg-hashids : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-pg-hashids : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.3" "postgresql-18-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-17-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-16-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-15-pg-hashids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3" "postgresql-14-pg-hashids : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-hashids : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-hashids : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-hashids : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-hashids : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-hashids : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-hashids : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-hashids : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-hashids : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-hashids : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-hashids : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashids_18` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.6 KiB | [pg_hashids_18-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_18-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_18` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.9 KiB | [pg_hashids_18-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_18-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_18` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.2 KiB | [pg_hashids_18-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_18-1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashids_18` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.6 KiB | [pg_hashids_18-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_18-1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashids_18` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.2 KiB | [pg_hashids_18-1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashids_18-1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_hashids_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.7 KiB | [pg_hashids_18-1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashids_18-1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-hashids` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.2 KiB | [postgresql-18-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-18-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-hashids` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.4 KiB | [postgresql-18-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-18-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-hashids` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.3 KiB | [postgresql-18-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-18-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-hashids` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.5 KiB | [postgresql-18-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-18-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-hashids` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.5 KiB | [postgresql-18-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-18-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-hashids` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.6 KiB | [postgresql-18-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-18-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-hashids` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-18-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-18-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-hashids` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.7 KiB | [postgresql-18-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-18-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashids_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.6 KiB | [pg_hashids_17-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_17-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.9 KiB | [pg_hashids_17-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_17-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_17` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.2 KiB | [pg_hashids_17-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_17-1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashids_17` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.6 KiB | [pg_hashids_17-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_17-1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashids_17` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.2 KiB | [pg_hashids_17-1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashids_17-1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_hashids_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.7 KiB | [pg_hashids_17-1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashids_17-1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-hashids` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.2 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-hashids` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.5 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-hashids` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.2 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-hashids` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.5 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-hashids` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 28.4 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-hashids` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.6 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-hashids` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-hashids` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.7 KiB | [postgresql-17-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-17-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashids_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.6 KiB | [pg_hashids_16-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_16-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.9 KiB | [pg_hashids_16-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_16-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_16` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.2 KiB | [pg_hashids_16-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_16-1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashids_16` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.6 KiB | [pg_hashids_16-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_16-1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashids_16` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.2 KiB | [pg_hashids_16-1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashids_16-1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_hashids_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.7 KiB | [pg_hashids_16-1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashids_16-1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-hashids` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.2 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-hashids` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.4 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-hashids` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.3 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-hashids` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.5 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-hashids` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 28.4 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-hashids` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.6 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-hashids` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-hashids` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.7 KiB | [postgresql-16-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-16-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashids_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.9 KiB | [pg_hashids_15-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_15-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.2 KiB | [pg_hashids_15-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_15-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_15` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.6 KiB | [pg_hashids_15-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_15-1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashids_15` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.1 KiB | [pg_hashids_15-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_15-1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashids_15` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.7 KiB | [pg_hashids_15-1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashids_15-1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_hashids_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.4 KiB | [pg_hashids_15-1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashids_15-1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-hashids` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.5 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-hashids` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.7 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-hashids` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.6 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-hashids` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.8 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-hashids` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.7 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-hashids` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.9 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-hashids` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.8 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-hashids` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.0 KiB | [postgresql-15-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-15-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_hashids_14` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.9 KiB | [pg_hashids_14-1.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_hashids_14-1.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_hashids_14` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.2 KiB | [pg_hashids_14-1.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_hashids_14-1.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_hashids_14` | `1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 18.6 KiB | [pg_hashids_14-1.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_hashids_14-1.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_hashids_14` | `1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.0 KiB | [pg_hashids_14-1.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_hashids_14-1.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_hashids_14` | `1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.7 KiB | [pg_hashids_14-1.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_hashids_14-1.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_hashids_14` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 18.5 KiB | [pg_hashids_14-1.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_hashids_14-1.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-hashids` | `1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.4 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-hashids` | `1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 26.7 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-hashids` | `1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.5 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-hashids` | `1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 26.8 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-hashids` | `1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.6 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-hashids` | `1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.9 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-hashids` | `1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.8 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-hashids` | `1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 28.0 KiB | [postgresql-14-pg-hashids_1.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-hashids/postgresql-14-pg-hashids_1.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/iCyberon/pg_hashids" title="Repository" icon="github" subtitle="github.com/iCyberon/pg_hashids" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_hashids-1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_hashids;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_hashids;		# install via package name, for the active PG version

pig install pg_hashids -v 18;   # install for PG 18
pig install pg_hashids -v 17;   # install for PG 17
pig install pg_hashids -v 16;   # install for PG 16
pig install pg_hashids -v 15;   # install for PG 15
pig install pg_hashids -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_hashids;
```



## Usage

> [pg_hashids: generate short unique ids from integers](https://github.com/iCyberon/pg_hashids)

Converts numbers like 347 into strings like "yr8". Useful for obfuscating database primary keys.

```sql
CREATE EXTENSION pg_hashids;
```

### Functions

| Function | Description |
|---|---|
| `id_encode(number [, salt [, min_length [, alphabet]]])` | Encode an integer to a hashid string |
| `id_decode(hash, salt, min_length [, alphabet])` | Decode a hashid string back to integer array |
| `id_decode_once(hash [, salt [, min_length [, alphabet]]])` | Decode a hashid string back to a single integer |

### Examples

```sql
-- Basic encoding
SELECT id_encode(1001);                                    -- 'jNl'
SELECT id_encode(1234567, 'This is my salt');              -- 'Pdzxp'
SELECT id_encode(1234567, 'This is my salt', 10);          -- 'PlRPdzxpR7'

-- Custom alphabet
SELECT id_encode(1234567, 'This is my salt', 10, 'abcdefghijABCDxFGHIJ1234567890');
-- '3GJ956J9B9'

-- Decoding (use the same salt!)
SELECT id_decode('PlRPdzxpR7', 'This is my salt', 10);     -- 1234567
SELECT id_decode_once('jNl');                               -- 1001
SELECT id_decode_once('Pdzxp', 'This is my salt');          -- 1234567
```
