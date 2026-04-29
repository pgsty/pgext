---
title: "pg_text_semver"
linkTitle: "pg_text_semver"
description: "Semantic version domain and comparison operators for PostgreSQL"
weight: 3520
categories: ["TYPE"]
width: full
---

[**pg_text_semver**](https://github.com/bigsmoke/pg_text_semver) : Semantic version domain and comparison operators for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3520** | {{< badge content="pg_text_semver" link="https://github.com/bigsmoke/pg_text_semver" >}} | {{< ext "pg_text_semver" >}} | `1.2.1` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_text_semver` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.1` | {{< bg "18" "pg_text_semver_18" "green" >}} {{< bg "17" "pg_text_semver_17" "green" >}} {{< bg "16" "pg_text_semver_16" "green" >}} {{< bg "15" "pg_text_semver_15" "green" >}} {{< bg "14" "pg_text_semver_14" "green" >}} | `pg_text_semver_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.1` | {{< bg "18" "postgresql-18-pg-text-semver" "green" >}} {{< bg "17" "postgresql-17-pg-text-semver" "green" >}} {{< bg "16" "postgresql-16-pg-text-semver" "green" >}} {{< bg "15" "postgresql-15-pg-text-semver" "green" >}} {{< bg "14" "postgresql-14-pg-text-semver" "green" >}} | `postgresql-$v-pg-text-semver` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_text_semver_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-18-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-17-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-16-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-15-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-14-pg-text-semver : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-18-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-17-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-16-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-15-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-14-pg-text-semver : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-18-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-17-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-16-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-15-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-14-pg-text-semver : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-18-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-17-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-16-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-15-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-14-pg-text-semver : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-18-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-17-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-16-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-15-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-14-pg-text-semver : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-18-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-17-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-16-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-15-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-14-pg-text-semver : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-18-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-17-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-16-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-15-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-14-pg-text-semver : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-18-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-17-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-16-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-15-pg-text-semver : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-14-pg-text-semver : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-text-semver : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-text-semver : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-text-semver : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-text-semver : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-text-semver : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-text-semver : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-text-semver : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-text-semver : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-text-semver : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-text-semver : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_text_semver_18` | `1.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.0 KiB | [pg_text_semver_18-1.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_text_semver_18-1.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_text_semver_18` | `1.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.9 KiB | [pg_text_semver_18-1.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_text_semver_18-1.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_text_semver_18` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.6 KiB | [pg_text_semver_18-1.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_text_semver_18-1.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_text_semver_18` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.6 KiB | [pg_text_semver_18-1.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_text_semver_18-1.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_text_semver_18` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [pg_text_semver_18-1.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_text_semver_18-1.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_text_semver_18` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_text_semver_18-1.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_text_semver_18-1.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-text-semver` | `1.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.3 KiB | [postgresql-18-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-18-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-text-semver` | `1.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-18-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-18-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-text-semver` | `1.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.3 KiB | [postgresql-18-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-18-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-text-semver` | `1.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 KiB | [postgresql-18-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-18-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-text-semver` | `1.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.8 KiB | [postgresql-18-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-18-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-text-semver` | `1.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.8 KiB | [postgresql-18-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-18-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-text-semver` | `1.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 KiB | [postgresql-18-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-18-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-text-semver` | `1.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.7 KiB | [postgresql-18-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-18-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_text_semver_17` | `1.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.0 KiB | [pg_text_semver_17-1.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_text_semver_17-1.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_text_semver_17` | `1.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.9 KiB | [pg_text_semver_17-1.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_text_semver_17-1.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_text_semver_17` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.6 KiB | [pg_text_semver_17-1.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_text_semver_17-1.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_text_semver_17` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.6 KiB | [pg_text_semver_17-1.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_text_semver_17-1.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_text_semver_17` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [pg_text_semver_17-1.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_text_semver_17-1.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_text_semver_17` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_text_semver_17-1.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_text_semver_17-1.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-text-semver` | `1.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.3 KiB | [postgresql-17-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-17-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-text-semver` | `1.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-17-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-17-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-text-semver` | `1.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.3 KiB | [postgresql-17-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-17-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-text-semver` | `1.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 KiB | [postgresql-17-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-17-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-text-semver` | `1.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.8 KiB | [postgresql-17-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-17-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-text-semver` | `1.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.8 KiB | [postgresql-17-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-17-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-text-semver` | `1.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 KiB | [postgresql-17-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-17-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-text-semver` | `1.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.7 KiB | [postgresql-17-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-17-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_text_semver_16` | `1.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.0 KiB | [pg_text_semver_16-1.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_text_semver_16-1.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_text_semver_16` | `1.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.9 KiB | [pg_text_semver_16-1.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_text_semver_16-1.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_text_semver_16` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.6 KiB | [pg_text_semver_16-1.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_text_semver_16-1.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_text_semver_16` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.6 KiB | [pg_text_semver_16-1.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_text_semver_16-1.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_text_semver_16` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [pg_text_semver_16-1.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_text_semver_16-1.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_text_semver_16` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_text_semver_16-1.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_text_semver_16-1.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-text-semver` | `1.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.3 KiB | [postgresql-16-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-16-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-text-semver` | `1.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-16-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-16-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-text-semver` | `1.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.3 KiB | [postgresql-16-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-16-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-text-semver` | `1.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 KiB | [postgresql-16-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-16-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-text-semver` | `1.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.8 KiB | [postgresql-16-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-16-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-text-semver` | `1.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.8 KiB | [postgresql-16-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-16-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-text-semver` | `1.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 KiB | [postgresql-16-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-16-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-text-semver` | `1.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.7 KiB | [postgresql-16-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-16-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_text_semver_15` | `1.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.0 KiB | [pg_text_semver_15-1.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_text_semver_15-1.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_text_semver_15` | `1.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.9 KiB | [pg_text_semver_15-1.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_text_semver_15-1.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_text_semver_15` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.6 KiB | [pg_text_semver_15-1.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_text_semver_15-1.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_text_semver_15` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.6 KiB | [pg_text_semver_15-1.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_text_semver_15-1.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_text_semver_15` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [pg_text_semver_15-1.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_text_semver_15-1.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_text_semver_15` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_text_semver_15-1.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_text_semver_15-1.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-text-semver` | `1.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.3 KiB | [postgresql-15-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-15-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-text-semver` | `1.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-15-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-15-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-text-semver` | `1.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.3 KiB | [postgresql-15-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-15-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-text-semver` | `1.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 KiB | [postgresql-15-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-15-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-text-semver` | `1.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.9 KiB | [postgresql-15-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-15-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-text-semver` | `1.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.9 KiB | [postgresql-15-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-15-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-text-semver` | `1.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.8 KiB | [postgresql-15-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-15-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-text-semver` | `1.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.8 KiB | [postgresql-15-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-15-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_text_semver_14` | `1.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.0 KiB | [pg_text_semver_14-1.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_text_semver_14-1.2.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_text_semver_14` | `1.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.9 KiB | [pg_text_semver_14-1.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_text_semver_14-1.2.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_text_semver_14` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.6 KiB | [pg_text_semver_14-1.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_text_semver_14-1.2.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_text_semver_14` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.6 KiB | [pg_text_semver_14-1.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_text_semver_14-1.2.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_text_semver_14` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [pg_text_semver_14-1.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_text_semver_14-1.2.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_text_semver_14` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_text_semver_14-1.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_text_semver_14-1.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-text-semver` | `1.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.3 KiB | [postgresql-14-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-14-pg-text-semver_1.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-text-semver` | `1.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.3 KiB | [postgresql-14-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-text-semver/postgresql-14-pg-text-semver_1.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-text-semver` | `1.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.3 KiB | [postgresql-14-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-14-pg-text-semver_1.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-text-semver` | `1.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.3 KiB | [postgresql-14-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-text-semver/postgresql-14-pg-text-semver_1.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-text-semver` | `1.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.8 KiB | [postgresql-14-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-14-pg-text-semver_1.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-text-semver` | `1.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.8 KiB | [postgresql-14-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-text-semver/postgresql-14-pg-text-semver_1.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-text-semver` | `1.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.7 KiB | [postgresql-14-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-14-pg-text-semver_1.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-text-semver` | `1.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.7 KiB | [postgresql-14-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-text-semver/postgresql-14-pg-text-semver_1.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_text_semver" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_text_semver" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_text_semver-1.2.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_text_semver;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_text_semver;		# install via package name, for the active PG version

pig install pg_text_semver -v 18;   # install for PG 18
pig install pg_text_semver -v 17;   # install for PG 17
pig install pg_text_semver -v 16;   # install for PG 16
pig install pg_text_semver -v 15;   # install for PG 15
pig install pg_text_semver -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_text_semver;
```

## Usage

Source: [README](https://github.com/bigsmoke/pg_text_semver/blob/master/README.md), [META.json](https://github.com/bigsmoke/pg_text_semver/blob/master/META.json), [Tag v1.2.1](https://github.com/bigsmoke/pg_text_semver/tree/v1.2.1)

`pg_text_semver` implements Semantic Versioning 2.0.0 on top of PostgreSQL `text` using a `semver` domain rather than a custom C type.

### Core types and functions

```sql
CREATE EXTENSION pg_text_semver;

SELECT '1.2.3'::semver < '2.0.0'::semver;
SELECT semver_cmp('1.2.3'::semver, '1.2.4'::semver);
SELECT semver_regexp(true);
SELECT '1.2.3-alpha.1+build5'::semver::semver_parsed;
```

- `semver`: domain over `text` with SemVer validation.
- `semver_parsed`: parsed composite type that supports sorting and indexing.
- `semver_prerelease`: domain for prerelease identifiers.
- `semver_cmp(...)`: comparison function for `semver` and `semver_parsed`.
- `semver_regexp(include_captures boolean)`: exposes the validation regex.

### Extra helpers

The current README also documents PGXN-version-range helpers:

- `meta_pgxn_version_range(text)`
- `meta_pgxn_version_range_cmp(text, text)`
- `nonsemver_cmp(text, text, text)`

### Caveats

- This extension favors a spec-oriented, text-backed implementation over the lower-level C-based alternatives.
- The upstream README remains the authoritative user-facing reference; the current stub already matched that surface closely, so this refresh mainly aligns it with the documented 1.2.1 helper set.
