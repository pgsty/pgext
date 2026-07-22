---
title: "pg_roast"
linkTitle: "pg_roast"
description: "Opinionated PostgreSQL database auditor"
weight: 7120
categories: ["SEC"]
width: full
---

[**pg_roast**](https://github.com/samirketema/pg_roast) : Opinionated PostgreSQL database auditor


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7120** | {{< badge content="pg_roast" link="https://github.com/samirketema/pg_roast" >}} | {{< ext "pg_roast" >}} | `1.0` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `roast` |
|   **See Also**    | {{< ext "pglinter" >}} {{< ext "pg_profile" >}} {{< ext "pg_stat_statements" >}} |

> [!Note] Upstream has no release tag; package pins main commit ccbf012. Manual audits work normally; the periodic background worker requires shared_preload_libraries=pg_roast.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_roast` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_roast_18" "green" >}} {{< bg "17" "pg_roast_17" "green" >}} {{< bg "16" "pg_roast_16" "green" >}} {{< bg "15" "pg_roast_15" "green" >}} {{< bg "14" "pg_roast_14" "green" >}} | `pg_roast_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-roast" "green" >}} {{< bg "17" "postgresql-17-pg-roast" "green" >}} {{< bg "16" "postgresql-16-pg-roast" "green" >}} {{< bg "15" "postgresql-15-pg-roast" "green" >}} {{< bg "14" "postgresql-14-pg-roast" "green" >}} | `postgresql-$v-pg-roast` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_roast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_roast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_roast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_roast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_roast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_roast_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_roast_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-roast : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-roast : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roast_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.6 KiB | [pg_roast_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roast_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_roast_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.0 KiB | [pg_roast_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roast_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_roast_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.2 KiB | [pg_roast_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roast_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_roast_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.3 KiB | [pg_roast_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roast_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_roast_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.3 KiB | [pg_roast_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roast_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_roast_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.6 KiB | [pg_roast_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roast_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-roast` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.1 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.0 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 31.1 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 31.1 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.5 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.2 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.0 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.8 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.8 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-roast` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.8 KiB | [postgresql-18-pg-roast_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-18-pg-roast_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roast_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.6 KiB | [pg_roast_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roast_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_roast_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.0 KiB | [pg_roast_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roast_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_roast_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.2 KiB | [pg_roast_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roast_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_roast_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.3 KiB | [pg_roast_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roast_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_roast_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.3 KiB | [pg_roast_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roast_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_roast_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.6 KiB | [pg_roast_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roast_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-roast` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.0 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.0 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 31.1 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 31.1 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 33.0 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.6 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.9 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.8 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.8 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-roast` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.8 KiB | [postgresql-17-pg-roast_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-17-pg-roast_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roast_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.6 KiB | [pg_roast_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roast_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_roast_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.0 KiB | [pg_roast_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roast_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_roast_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.2 KiB | [pg_roast_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roast_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_roast_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.3 KiB | [pg_roast_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roast_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_roast_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.3 KiB | [pg_roast_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roast_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_roast_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.6 KiB | [pg_roast_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roast_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-roast` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.0 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.0 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 31.0 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 31.0 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 33.0 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.6 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.9 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.8 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.7 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-roast` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.8 KiB | [postgresql-16-pg-roast_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-16-pg-roast_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roast_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.7 KiB | [pg_roast_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roast_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_roast_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.0 KiB | [pg_roast_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roast_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_roast_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.3 KiB | [pg_roast_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roast_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_roast_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.3 KiB | [pg_roast_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roast_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_roast_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.4 KiB | [pg_roast_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roast_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_roast_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.6 KiB | [pg_roast_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roast_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-roast` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.1 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.1 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 31.1 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 31.1 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 33.1 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.6 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.1 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.8 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.9 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-roast` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.8 KiB | [postgresql-15-pg-roast_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-15-pg-roast_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roast_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.7 KiB | [pg_roast_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roast_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_roast_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.0 KiB | [pg_roast_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roast_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_roast_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.3 KiB | [pg_roast_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roast_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_roast_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.3 KiB | [pg_roast_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roast_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_roast_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.4 KiB | [pg_roast_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roast_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_roast_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.6 KiB | [pg_roast_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roast_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-roast` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 31.1 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 31.0 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 31.1 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 31.0 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 33.0 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.6 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.1 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.8 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.9 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-roast` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.8 KiB | [postgresql-14-pg-roast_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-roast/postgresql-14-pg-roast_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/samirketema/pg_roast" title="Repository" icon="github" subtitle="github.com/samirketema/pg_roast" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_roast-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_roast;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_roast;		# install via package name, for the active PG version

pig install pg_roast -v 18;   # install for PG 18
pig install pg_roast -v 17;   # install for PG 17
pig install pg_roast -v 16;   # install for PG 16
pig install pg_roast -v 15;   # install for PG 15
pig install pg_roast -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_roast';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_roast;
```

## Usage

Sources:

- [Upstream README](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/README.md)
- [Version 1.0 install SQL](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/pg_roast--1.0.sql)
- [Background-worker implementation](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/pg_roast_bgw.c)

pg_roast runs opinionated PostgreSQL health checks and stores findings in its fixed roast schema. It checks configuration, schema design, indexes, vacuum and bloat indicators, security posture, replication, connections, and workload signals. Version 1.0 targets PostgreSQL 14 and later.

### Manual audits

Manual mode does not require preloading:

```sql
CREATE EXTENSION pg_roast;

SELECT * FROM roast.run();
SELECT severity, check_name, object_name, roast
FROM roast.latest
ORDER BY severity, check_name;

SELECT * FROM roast.summary;
```

Each run persists audit history and findings. Use the latest view for the newest run and the summary view for grouped results.

### Scheduled audits

The optional background worker requires preload configuration and a restart:

```ini
shared_preload_libraries = 'pg_roast'
pg_roast.database = 'mydb'
pg_roast.interval = 3600
```

The database setting is fixed at server start. Review the upstream settings before enabling automatic audits across a production workload.

### Caveats

- Findings are heuristic advice, not automatic proof of a defect. Review workload context, maintenance windows, and PostgreSQL documentation before applying any recommendation.
- Audits execute catalog and statistics queries and write history tables. Measure overhead on large catalogs and protect the roast schema from untrusted users.
- Results can expose object names, configuration, security findings, and query-related operational details. Apply least privilege and an explicit retention policy.
- Preloading is unnecessary for manual runs but mandatory for the background worker; changing startup-only settings requires a restart.
