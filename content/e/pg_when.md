---
title: "pg_when"
linkTitle: "pg_when"
description: "Natural language time parsing for PostgreSQL"
weight: 1120
categories: ["TIME"]
width: full
---

[**pg_when**](https://github.com/frectonz/pg-when) : Natural language time parsing for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1120** | {{< badge content="pg_when" link="https://github.com/frectonz/pg-when" >}} | {{< ext "pg_when" >}} | `0.1.9` | {{< category "TIME" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] manually upgraded PGRX from 0.15.0 to 0.17.0 by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_when` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "pg_when_18" "green" >}} {{< bg "17" "pg_when_17" "green" >}} {{< bg "16" "pg_when_16" "green" >}} {{< bg "15" "pg_when_15" "green" >}} {{< bg "14" "pg_when_14" "green" >}} | `pg_when_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "postgresql-18-pg-when" "green" >}} {{< bg "17" "postgresql-17-pg-when" "green" >}} {{< bg "16" "postgresql-16-pg-when" "green" >}} {{< bg "15" "postgresql-15-pg-when" "green" >}} {{< bg "14" "postgresql-14-pg-when" "green" >}} | `postgresql-$v-pg-when` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-when : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-when : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-when : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-when : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-when : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-when : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-when : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-when : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-when : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-when : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_18` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 439.5 KiB | [pg_when_18-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_18-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `pg_when_18` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 328.5 KiB | [pg_when_18-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_18-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `pg_when_18` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 456.1 KiB | [pg_when_18-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_18-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `pg_when_18` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 347.9 KiB | [pg_when_18-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_18-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `pg_when_18` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 456.2 KiB | [pg_when_18-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_18-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `pg_when_18` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 347.6 KiB | [pg_when_18-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_18-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 364.2 KiB | [postgresql-18-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 259.7 KiB | [postgresql-18-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 364.2 KiB | [postgresql-18-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 259.6 KiB | [postgresql-18-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.7 KiB | [postgresql-18-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 298.4 KiB | [postgresql-18-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 403.4 KiB | [postgresql-18-pg-when_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 296.1 KiB | [postgresql-18-pg-when_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-18-pg-when_0.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_17` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 439.7 KiB | [pg_when_17-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_17-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `pg_when_17` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 328.5 KiB | [pg_when_17-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_17-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `pg_when_17` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 456.1 KiB | [pg_when_17-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_17-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `pg_when_17` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 348.0 KiB | [pg_when_17-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_17-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `pg_when_17` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 456.3 KiB | [pg_when_17-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_17-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `pg_when_17` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 347.5 KiB | [pg_when_17-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_17-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 364.5 KiB | [postgresql-17-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 259.7 KiB | [postgresql-17-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 364.7 KiB | [postgresql-17-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 259.6 KiB | [postgresql-17-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.4 KiB | [postgresql-17-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 298.5 KiB | [postgresql-17-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 403.5 KiB | [postgresql-17-pg-when_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 296.1 KiB | [postgresql-17-pg-when_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-17-pg-when_0.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_16` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 439.8 KiB | [pg_when_16-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_16-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `pg_when_16` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 328.5 KiB | [pg_when_16-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_16-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `pg_when_16` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 456.2 KiB | [pg_when_16-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_16-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `pg_when_16` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 348.0 KiB | [pg_when_16-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_16-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `pg_when_16` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 456.4 KiB | [pg_when_16-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_16-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `pg_when_16` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 347.6 KiB | [pg_when_16-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_16-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 363.9 KiB | [postgresql-16-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 259.6 KiB | [postgresql-16-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 364.0 KiB | [postgresql-16-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 259.6 KiB | [postgresql-16-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.5 KiB | [postgresql-16-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 298.4 KiB | [postgresql-16-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 403.2 KiB | [postgresql-16-pg-when_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 296.1 KiB | [postgresql-16-pg-when_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-16-pg-when_0.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_15` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 439.2 KiB | [pg_when_15-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_15-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `pg_when_15` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 328.6 KiB | [pg_when_15-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_15-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `pg_when_15` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 456.1 KiB | [pg_when_15-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_15-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `pg_when_15` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 348.3 KiB | [pg_when_15-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_15-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `pg_when_15` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 456.2 KiB | [pg_when_15-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_15-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `pg_when_15` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 347.6 KiB | [pg_when_15-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_15-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 364.1 KiB | [postgresql-15-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 259.7 KiB | [postgresql-15-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 363.9 KiB | [postgresql-15-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 259.5 KiB | [postgresql-15-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.5 KiB | [postgresql-15-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 298.6 KiB | [postgresql-15-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 403.3 KiB | [postgresql-15-pg-when_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 296.1 KiB | [postgresql-15-pg-when_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-15-pg-when_0.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_14` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 439.3 KiB | [pg_when_14-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_14-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `pg_when_14` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 328.5 KiB | [pg_when_14-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_14-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `pg_when_14` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 455.7 KiB | [pg_when_14-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_14-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `pg_when_14` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 348.0 KiB | [pg_when_14-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_14-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `pg_when_14` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 456.0 KiB | [pg_when_14-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_14-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `pg_when_14` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 347.6 KiB | [pg_when_14-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_14-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-when` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 364.2 KiB | [postgresql-14-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 259.6 KiB | [postgresql-14-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 364.2 KiB | [postgresql-14-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 259.5 KiB | [postgresql-14-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 406.2 KiB | [postgresql-14-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 298.4 KiB | [postgresql-14-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 403.2 KiB | [postgresql-14-pg-when_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 296.1 KiB | [postgresql-14-pg-when_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-14-pg-when_0.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/frectonz/pg-when" title="Repository" icon="github" subtitle="github.com/frectonz/pg-when" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_when-0.1.9.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_when;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_when;		# install via package name, for the active PG version

pig install pg_when -v 18;   # install for PG 18
pig install pg_when -v 17;   # install for PG 17
pig install pg_when -v 16;   # install for PG 16
pig install pg_when -v 15;   # install for PG 15
pig install pg_when -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_when;
```

## Usage

Sources: [official README](https://github.com/frectonz/pg-when/blob/main/README.md), [official repo](https://github.com/frectonz/pg-when)

`pg-when` parses a constrained natural-language time expression and returns either a PostgreSQL timestamp with time zone or Unix epoch values at different resolutions.

```sql
CREATE EXTENSION pg_when;

SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT seconds_at('next friday at 8:00 pm in America/New_York');
SELECT millis_at('next friday at 8:00 pm in America/New_York');
SELECT micros_at('next friday at 8:00 pm in America/New_York');
SELECT nanos_at('next friday at 8:00 pm in America/New_York');
```

### Supported Query Shape

The parser accepts up to three parts:

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

If no timezone is provided, upstream says the default is UTC.

### Common Inputs

- relative dates: `today`, `tomorrow`, `last month`, `this friday`, `5 days ago`, `in 2 years`
- exact dates: `YYYY-MM-DD`, `DD/MM/YYYY`, `January 10, 2004`, `10 Jan 2004`
- relative times: `noon`, `midnight`, `morning`, `evening`, `next hour`
- exact times: `8:30 pm`, `15:45`
- time zones: `America/New_York`, `Europe/London`, `UTC-08:00`, `UTC+05:30`

### Examples

```sql
SELECT when_is('5 days ago at this hour in Asia/Tokyo');
SELECT when_is('in 2 months at midnight in UTC-8');
SELECT when_is('December 31, 2026 at evening');
```

### Caveats

- The extension is aimed at the documented grammar above, not arbitrary English.
- Upstream distributes ready-made Docker images for PostgreSQL 13 through 18, but the stub should focus on SQL usage rather than container setup.
