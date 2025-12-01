---
title: "pg_upless"
linkTitle: "pg_upless"
description: "Detect Useless UPDATE"
weight: 5180
categories: ["ADMIN"]
width: full
---

[**pg_upless**](https://github.com/rodo/pg_upless) : Detect Useless UPDATE


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5180** | {{< badge content="pg_upless" link="https://github.com/rodo/pg_upless" >}} | {{< ext "pg_upless" >}} | `0.0.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_savior" >}} {{< ext "safeupdate" >}} {{< ext "pg_permissions" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_drop_events" >}} {{< ext "table_log" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_upless` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "pg_upless_18" "green" >}} {{< bg "17" "pg_upless_17" "green" >}} {{< bg "16" "pg_upless_16" "green" >}} {{< bg "15" "pg_upless_15" "green" >}} {{< bg "14" "pg_upless_14" "green" >}} {{< bg "13" "pg_upless_13" "green" >}} | `pg_upless_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.3` | {{< bg "18" "postgresql-18-pg-upless" "green" >}} {{< bg "17" "postgresql-17-pg-upless" "green" >}} {{< bg "16" "postgresql-16-pg-upless" "green" >}} {{< bg "15" "postgresql-15-pg-upless" "green" >}} {{< bg "14" "postgresql-14-pg-upless" "green" >}} {{< bg "13" "postgresql-13-pg-upless" "green" >}} | `postgresql-$v-pg-upless` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "pg_upless_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-upless : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-upless : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-upless : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-upless : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-upless : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-upless : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-upless : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-18-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-17-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-16-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-15-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-14-pg-upless : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.3" "postgresql-13-pg-upless : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_upless_18` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [pg_upless_18-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_18-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_18` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.6 KiB | [pg_upless_18-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_18-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_18` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.6 KiB | [pg_upless_18-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_18-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_18` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [pg_upless_18-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_18-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_upless_18` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [pg_upless_18-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_upless_18-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_upless_18` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [pg_upless_18-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_upless_18-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-upless` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.1 KiB | [postgresql-18-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-18-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-upless` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 KiB | [postgresql-18-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-18-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-upless` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.1 KiB | [postgresql-18-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-18-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-upless` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 KiB | [postgresql-18-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-18-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-upless` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.9 KiB | [postgresql-18-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-18-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-upless` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.9 KiB | [postgresql-18-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-18-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-upless` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.8 KiB | [postgresql-18-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-18-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-upless` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.9 KiB | [postgresql-18-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-18-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_upless_17` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [pg_upless_17-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_17-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_17` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.6 KiB | [pg_upless_17-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_17-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_17` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.6 KiB | [pg_upless_17-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_17-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_17` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [pg_upless_17-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_17-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_upless_17` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [pg_upless_17-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_upless_17-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_upless_17` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [pg_upless_17-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_upless_17-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-upless` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.1 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-upless` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-upless` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.1 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-upless` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-upless` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.9 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-upless` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.9 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-upless` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-upless` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.9 KiB | [postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-17-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_upless_16` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [pg_upless_16-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_16-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_16` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.6 KiB | [pg_upless_16-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_16-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_16` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.6 KiB | [pg_upless_16-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_16-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_16` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [pg_upless_16-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_16-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_upless_16` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [pg_upless_16-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_upless_16-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_upless_16` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [pg_upless_16-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_upless_16-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-upless` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.1 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-upless` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-upless` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.1 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-upless` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-upless` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.9 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-upless` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.9 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-upless` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-upless` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.9 KiB | [postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-16-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_upless_15` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [pg_upless_15-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_15-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_15` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.6 KiB | [pg_upless_15-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_15-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_15` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.6 KiB | [pg_upless_15-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_15-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_15` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [pg_upless_15-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_15-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_upless_15` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [pg_upless_15-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_upless_15-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_upless_15` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [pg_upless_15-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_upless_15-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-upless` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.1 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-upless` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-upless` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.1 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-upless` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-upless` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.9 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-upless` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.9 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-upless` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-upless` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.9 KiB | [postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-15-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_upless_14` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [pg_upless_14-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_14-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_14` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.6 KiB | [pg_upless_14-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_14-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_14` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.6 KiB | [pg_upless_14-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_14-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_14` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [pg_upless_14-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_14-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_upless_14` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [pg_upless_14-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_upless_14-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_upless_14` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [pg_upless_14-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_upless_14-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-upless` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.1 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-upless` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-upless` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.1 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-upless` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-upless` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.9 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-upless` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.9 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-upless` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-upless` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.9 KiB | [postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-14-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_upless_13` | `0.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [pg_upless_13-0.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_upless_13-0.0.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_upless_13` | `0.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.6 KiB | [pg_upless_13-0.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_upless_13-0.0.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_upless_13` | `0.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.6 KiB | [pg_upless_13-0.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_upless_13-0.0.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_upless_13` | `0.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [pg_upless_13-0.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_upless_13-0.0.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_upless_13` | `0.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [pg_upless_13-0.0.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_upless_13-0.0.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_upless_13` | `0.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [pg_upless_13-0.0.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_upless_13-0.0.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-upless` | `0.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.1 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-upless` | `0.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-upless` | `0.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.1 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-upless` | `0.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-upless` | `0.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.9 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-upless` | `0.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.9 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-upless` | `0.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.9 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-upless` | `0.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.9 KiB | [postgresql-13-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-upless/postgresql-13-pg-upless_0.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rodo/pg_upless" title="Repository" icon="github" subtitle="github.com/rodo/pg_upless" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_upless-0.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_upless;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_upless;		# install via package name, for the active PG version

pig install pg_upless -v 18;   # install for PG 18
pig install pg_upless -v 17;   # install for PG 17
pig install pg_upless -v 16;   # install for PG 16
pig install pg_upless -v 15;   # install for PG 15
pig install pg_upless -v 14;   # install for PG 14
pig install pg_upless -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_upless CASCADE; -- requires plpgsql
```
