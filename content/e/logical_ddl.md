---
title: "logical_ddl"
linkTitle: "logical_ddl"
description: "Replicate supported DDL changes over PostgreSQL logical replication"
weight: 9530
categories: ["ETL"]
width: full
---

[**logical_ddl**](https://github.com/samedyildirim/logical_ddl) : Replicate supported DDL changes over PostgreSQL logical replication


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9530** | {{< badge content="logical_ddl" link="https://github.com/samedyildirim/logical_ddl" >}} | {{< ext "logical_ddl" >}} | `0.1.0` | {{< category "ETL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `logical_ddl` |

> [!Note] Pigsty carries the upstream RAISE WARNING typo fix for 0.1.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `logical_ddl` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "logical_ddl_18" "green" >}} {{< bg "17" "logical_ddl_17" "green" >}} {{< bg "16" "logical_ddl_16" "green" >}} {{< bg "15" "logical_ddl_15" "green" >}} {{< bg "14" "logical_ddl_14" "green" >}} | `logical_ddl_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-logical-ddl" "green" >}} {{< bg "17" "postgresql-17-logical-ddl" "green" >}} {{< bg "16" "postgresql-16-logical-ddl" "green" >}} {{< bg "15" "postgresql-15-logical-ddl" "green" >}} {{< bg "14" "postgresql-14-logical-ddl" "green" >}} | `postgresql-$v-logical-ddl` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "logical_ddl_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-logical-ddl : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-logical-ddl : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-logical-ddl : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-logical-ddl : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-logical-ddl : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-logical-ddl : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-logical-ddl : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-logical-ddl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-logical-ddl : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-logical-ddl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-logical-ddl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-logical-ddl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-logical-ddl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-logical-ddl : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-logical-ddl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-logical-ddl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-logical-ddl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-logical-ddl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-logical-ddl : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logical_ddl_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [logical_ddl_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/logical_ddl_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `logical_ddl_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.3 KiB | [logical_ddl_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/logical_ddl_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `logical_ddl_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [logical_ddl_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/logical_ddl_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `logical_ddl_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [logical_ddl_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/logical_ddl_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `logical_ddl_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.2 KiB | [logical_ddl_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/logical_ddl_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `logical_ddl_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 KiB | [logical_ddl_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/logical_ddl_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-logical-ddl` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 KiB | [postgresql-18-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-18-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-logical-ddl` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.3 KiB | [postgresql-18-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-18-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-logical-ddl` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.3 KiB | [postgresql-18-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-18-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-logical-ddl` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.3 KiB | [postgresql-18-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-18-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-logical-ddl` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.0 KiB | [postgresql-18-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-18-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-logical-ddl` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.1 KiB | [postgresql-18-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-18-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-logical-ddl` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.0 KiB | [postgresql-18-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-18-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-logical-ddl` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.7 KiB | [postgresql-18-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-18-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logical_ddl_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [logical_ddl_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/logical_ddl_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `logical_ddl_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.3 KiB | [logical_ddl_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/logical_ddl_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `logical_ddl_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [logical_ddl_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/logical_ddl_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `logical_ddl_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [logical_ddl_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/logical_ddl_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `logical_ddl_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.2 KiB | [logical_ddl_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/logical_ddl_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `logical_ddl_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 KiB | [logical_ddl_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/logical_ddl_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-logical-ddl` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 KiB | [postgresql-17-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-17-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-logical-ddl` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.3 KiB | [postgresql-17-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-17-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-logical-ddl` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.3 KiB | [postgresql-17-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-17-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-logical-ddl` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.3 KiB | [postgresql-17-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-17-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-logical-ddl` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.6 KiB | [postgresql-17-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-17-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-logical-ddl` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.7 KiB | [postgresql-17-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-17-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-logical-ddl` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.0 KiB | [postgresql-17-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-17-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-logical-ddl` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.7 KiB | [postgresql-17-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-17-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logical_ddl_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [logical_ddl_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/logical_ddl_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `logical_ddl_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.3 KiB | [logical_ddl_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/logical_ddl_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `logical_ddl_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [logical_ddl_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/logical_ddl_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `logical_ddl_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [logical_ddl_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/logical_ddl_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `logical_ddl_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.2 KiB | [logical_ddl_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/logical_ddl_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `logical_ddl_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 KiB | [logical_ddl_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/logical_ddl_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-logical-ddl` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 KiB | [postgresql-16-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-16-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-logical-ddl` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.2 KiB | [postgresql-16-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-16-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-logical-ddl` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.3 KiB | [postgresql-16-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-16-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-logical-ddl` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.3 KiB | [postgresql-16-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-16-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-logical-ddl` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.5 KiB | [postgresql-16-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-16-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-logical-ddl` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.6 KiB | [postgresql-16-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-16-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-logical-ddl` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.0 KiB | [postgresql-16-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-16-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-logical-ddl` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.7 KiB | [postgresql-16-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-16-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logical_ddl_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [logical_ddl_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/logical_ddl_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `logical_ddl_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.3 KiB | [logical_ddl_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/logical_ddl_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `logical_ddl_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [logical_ddl_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/logical_ddl_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `logical_ddl_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [logical_ddl_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/logical_ddl_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `logical_ddl_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.2 KiB | [logical_ddl_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/logical_ddl_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `logical_ddl_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 KiB | [logical_ddl_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/logical_ddl_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-logical-ddl` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 KiB | [postgresql-15-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-15-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-logical-ddl` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.3 KiB | [postgresql-15-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-15-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-logical-ddl` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.3 KiB | [postgresql-15-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-15-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-logical-ddl` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.3 KiB | [postgresql-15-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-15-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-logical-ddl` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.5 KiB | [postgresql-15-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-15-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-logical-ddl` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.6 KiB | [postgresql-15-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-15-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-logical-ddl` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.0 KiB | [postgresql-15-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-15-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-logical-ddl` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.8 KiB | [postgresql-15-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-15-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logical_ddl_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.1 KiB | [logical_ddl_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/logical_ddl_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `logical_ddl_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.3 KiB | [logical_ddl_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/logical_ddl_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `logical_ddl_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.0 KiB | [logical_ddl_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/logical_ddl_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `logical_ddl_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.1 KiB | [logical_ddl_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/logical_ddl_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `logical_ddl_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.2 KiB | [logical_ddl_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/logical_ddl_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `logical_ddl_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.3 KiB | [logical_ddl_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/logical_ddl_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-logical-ddl` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 KiB | [postgresql-14-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-14-logical-ddl_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-logical-ddl` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.3 KiB | [postgresql-14-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logical-ddl/postgresql-14-logical-ddl_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-logical-ddl` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.3 KiB | [postgresql-14-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-14-logical-ddl_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-logical-ddl` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.3 KiB | [postgresql-14-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logical-ddl/postgresql-14-logical-ddl_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-logical-ddl` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.5 KiB | [postgresql-14-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-14-logical-ddl_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-logical-ddl` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.5 KiB | [postgresql-14-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logical-ddl/postgresql-14-logical-ddl_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-logical-ddl` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.9 KiB | [postgresql-14-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-14-logical-ddl_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-logical-ddl` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.7 KiB | [postgresql-14-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logical-ddl/postgresql-14-logical-ddl_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/samedyildirim/logical_ddl" title="Repository" icon="github" subtitle="github.com/samedyildirim/logical_ddl" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="logical_ddl-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg logical_ddl;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install logical_ddl;		# install via package name, for the active PG version

pig install logical_ddl -v 18;   # install for PG 18
pig install logical_ddl -v 17;   # install for PG 17
pig install logical_ddl -v 16;   # install for PG 16
pig install logical_ddl -v 15;   # install for PG 15
pig install logical_ddl -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION logical_ddl;
```

## Usage

Source: [README](https://github.com/samedyildirim/logical_ddl/blob/master/README.md)

`logical_ddl` captures a limited set of `ALTER TABLE` changes, writes them into a replicated shadow table, and replays equivalent DDL on logical-replication subscribers.

### Supported DDL

- `ALTER TABLE ... RENAME TO ...`
- `ALTER TABLE ... RENAME COLUMN ... TO ...`
- `ALTER TABLE ... ADD COLUMN ...`
- `ALTER TABLE ... ALTER COLUMN ... TYPE ...`
- `ALTER TABLE ... DROP COLUMN ...`

Built-in types, arrays, composite types, domains, and enums are supported as column types, but the extension does not replicate the definitions of those custom types themselves.

### Publisher and subscriber setup

```sql
CREATE EXTENSION logical_ddl;

-- Publisher
INSERT INTO logical_ddl.settings VALUES (true, 'source1');
INSERT INTO logical_ddl.publish_tablelist (relid) VALUES ('table1'::regclass);
ALTER PUBLICATION log_pub_1 ADD TABLE logical_ddl.shadow_table;

-- Subscriber
INSERT INTO logical_ddl.settings VALUES (false, 'source1');
INSERT INTO logical_ddl.subscribe_tablelist (source, relid)
VALUES ('source1', 'table1'::regclass);
ALTER SUBSCRIPTION log_sub_1 REFRESH PUBLICATION;
```

### Main tables

- `logical_ddl.settings`: declares publisher/subscriber role and source name.
- `logical_ddl.publish_tablelist`: tables and command kinds to capture.
- `logical_ddl.subscribe_tablelist`: target tables and command kinds to replay.
- `logical_ddl.shadow_table`: replicated command log.
- `logical_ddl.applied_commands`: replay history and failure tracking.

### Caveats

- The extension works under superuser privileges.
- `USING` expressions for type changes, default expressions, constraints, and indexes are not implemented.
- Pigsty notes an upstream `RAISE WARNING` typo fix for `0.1.0`; that does not change the user-facing usage documented here.
