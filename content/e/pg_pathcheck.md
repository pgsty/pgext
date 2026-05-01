---
title: "pg_pathcheck"
linkTitle: "pg_pathcheck"
description: "Validate planner Path trees for freed or corrupt memory"
weight: 5250
categories: ["ADMIN"]
width: full
---

[**pg_pathcheck**](https://github.com/danolivo/pg_pathcheck) : Validate planner Path trees for freed or corrupt memory


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5250** | {{< badge content="pg_pathcheck" link="https://github.com/danolivo/pg_pathcheck" >}} | {{< ext "pg_pathcheck" >}} | `0.9.1` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_catcheck" >}} {{< ext "pg_checksums" >}} {{< ext "amcheck" >}} {{< ext "pg_surgery" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "pg_repack" >}} |

> [!Note] preload-only module; no CREATE EXTENSION objects; pg17-18 branch


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_pathcheck` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "pg_pathcheck_18" "green" >}} {{< bg "17" "pg_pathcheck_17" "green" >}} {{< bg "16" "pg_pathcheck_16" "red" >}} {{< bg "15" "pg_pathcheck_15" "red" >}} {{< bg "14" "pg_pathcheck_14" "red" >}} | `pg_pathcheck_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "postgresql-18-pg-pathcheck" "green" >}} {{< bg "17" "postgresql-17-pg-pathcheck" "green" >}} {{< bg "16" "postgresql-16-pg-pathcheck" "red" >}} {{< bg "15" "postgresql-15-pg-pathcheck" "red" >}} {{< bg "14" "postgresql-14-pg-pathcheck" "red" >}} | `postgresql-$v-pg-pathcheck` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_pathcheck_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_pathcheck_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_pathcheck_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_pathcheck_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_pathcheck_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pg_pathcheck_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_pathcheck_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_pathcheck_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pg-pathcheck : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pg-pathcheck : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-pathcheck : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-pathcheck : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pathcheck_18` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.2 KiB | [pg_pathcheck_18-0.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pathcheck_18-0.9.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_pathcheck_18` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.7 KiB | [pg_pathcheck_18-0.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pathcheck_18-0.9.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_pathcheck_18` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.6 KiB | [pg_pathcheck_18-0.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pathcheck_18-0.9.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_pathcheck_18` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.4 KiB | [pg_pathcheck_18-0.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pathcheck_18-0.9.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_pathcheck_18` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.8 KiB | [pg_pathcheck_18-0.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pathcheck_18-0.9.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_pathcheck_18` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.4 KiB | [pg_pathcheck_18-0.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pathcheck_18-0.9.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.0 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.9 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.8 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.0 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 67.2 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 67.0 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.1 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.9 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 62.2 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-pathcheck` | `0.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 62.3 KiB | [postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pathcheck/postgresql-18-pg-pathcheck_0.9.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_pathcheck_17` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.2 KiB | [pg_pathcheck_17-0.9.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_pathcheck_17-0.9.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_pathcheck_17` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.7 KiB | [pg_pathcheck_17-0.9.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_pathcheck_17-0.9.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_pathcheck_17` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.5 KiB | [pg_pathcheck_17-0.9.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_pathcheck_17-0.9.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_pathcheck_17` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.3 KiB | [pg_pathcheck_17-0.9.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_pathcheck_17-0.9.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_pathcheck_17` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.7 KiB | [pg_pathcheck_17-0.9.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_pathcheck_17-0.9.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_pathcheck_17` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.6 KiB | [pg_pathcheck_17-0.9.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_pathcheck_17-0.9.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 59.5 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 59.5 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 59.3 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 59.6 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 70.4 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 71.0 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 62.4 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 62.5 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 61.8 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-pathcheck` | `0.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 61.9 KiB | [postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-pathcheck/postgresql-17-pg-pathcheck_0.9.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/danolivo/pg_pathcheck" title="Repository" icon="github" subtitle="github.com/danolivo/pg_pathcheck" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_pathcheck-0.9.1-pg17-18.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_pathcheck;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_pathcheck;		# install via package name, for the active PG version

pig install pg_pathcheck -v 18;   # install for PG 18
pig install pg_pathcheck -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_pathcheck';
```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources: [README](https://github.com/danolivo/pg_pathcheck/blob/main/README.md), [0.9.1 PG17/18 release](https://github.com/danolivo/pg_pathcheck/releases/tag/v0.9.1_pg17-18), [PGXN metadata](https://api.pgxn.org/dist/pg_pathcheck.json), [source](https://api.pgxn.org/src/pg_pathcheck/pg_pathcheck-0.9.1/pg_pathcheck.c)

`pg_pathcheck` is a PostgreSQL planner diagnostics module that validates reachable planner `Path` trees and reports pointers that look freed, corrupt, or re-used for the wrong relation. It is a preload-only shared library: it registers planner hooks and custom GUCs, but it does not create SQL functions, operators, views, or tables.

### Loading The Module

Build `pg_pathcheck` against the PostgreSQL source line you want to test. The upstream README documents separate long-running branches for PostgreSQL 17/18 and PostgreSQL master/19devel; version `0.9.1` is published for the PG17/18 branch and the PGXN metadata also describes the 0.9.1 source package.

Add the module to `shared_preload_libraries` and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_pathcheck'
```

There is no `CREATE EXTENSION pg_pathcheck` step. After preload, run ordinary SQL, `EXPLAIN`, regression tests, or PostgreSQL test suites; `pg_pathcheck` checks planner paths as queries are planned.

For one temporary cluster:

```bash
initdb -D pgdata
echo "shared_preload_libraries = 'pg_pathcheck'" >> pgdata/postgresql.conf
pg_ctl -D pgdata -l pgdata/logfile start

psql -c 'EXPLAIN SELECT ...'
```

For PostgreSQL test clusters spawned by `make check-world`, use `TEMP_CONFIG`:

```bash
cat > /tmp/pg_pathcheck.conf <<'EOF'
shared_preload_libraries = 'pg_pathcheck'
EOF

TEMP_CONFIG=/tmp/pg_pathcheck.conf make check-world
```

### Configuration

`pg_pathcheck.elevel` controls the severity used when a bad `Path` is detected:

```sql
SET pg_pathcheck.elevel = 'log';
SET pg_pathcheck.elevel = 'warning';  -- default
SET pg_pathcheck.elevel = 'error';
SET pg_pathcheck.elevel = 'panic';
```

Use `warning` or `log` for broad coverage while tests continue, `error` to stop on the first offending query, and `panic` only when a backend crash and core dump are useful for post-mortem debugging.

`pg_pathcheck.stage_checks` enables additional per-stage checks:

```sql
SET pg_pathcheck.stage_checks = off;  -- default
SET pg_pathcheck.stage_checks = on;
```

When enabled, the module checks pathlists at base-rel, join-rel, and upper-rel hook boundaries so a finding can be tied to a narrower planner stage. Leave it off for routine runs because the extra walks add planner overhead, especially for join-heavy queries.

### What Gets Checked

The module walks planner roots, relation pathlists, partial pathlists, cheapest path slots, parameterized paths, subquery subroots, subplan subroots, and nested `Path` fields such as join children, append children, sort subpaths, bitmap paths, and similar compound-path members.

It reports two main classes of problem:

- Invalid `NodeTag`: the pointer no longer looks like a PostgreSQL `Path` node.
- Parent mismatch: a valid-looking `Path` on a base or join relation points at a different `RelOptInfo`, which can indicate same-size memory reuse after a stale path pointer survived.

A typical finding includes the bad tag or mismatch, the containing slot such as `pathlist`, `partial_pathlist`, `cheapest_total_path`, or a nested path field, the relation names that could be resolved, pathlist detail, and the query string available through PostgreSQL debug state.

### Caveats

`pg_pathcheck` is a debug aid for PostgreSQL planner development and extension testing, not a user-facing SQL extension. A PostgreSQL cassert/debug build gives better signal because freed memory is easier to recognize. The upstream README notes coverage differences between the PG17/18 branch and the master branch: PG17/18 checks before later planning stages such as `create_plan` and `setrefs.c`, while master can use the newer planner shutdown hook.
