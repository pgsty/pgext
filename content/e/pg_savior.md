---
title: "pg_savior"
linkTitle: "pg_savior"
description: "Postgres extension to save OOPS mistakes"
weight: 5810
categories: ["ADMIN"]
width: full
---

[**pg_savior**](https://github.com/viggy28/pg_savior) : Postgres extension to save OOPS mistakes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5810** | {{< badge content="pg_savior" link="https://github.com/viggy28/pg_savior" >}} | {{< ext "pg_savior" >}} | `0.0.1` | {{< category "ADMIN" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_upless" >}} {{< ext "safeupdate" >}} {{< ext "pg_drop_events" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "table_log" >}} {{< ext "pg_snakeoil" >}} {{< ext "pg_auditor" >}} {{< ext "temporal_tables" >}} |

> [!Note] -tuplestore_donestoring , breaks on pg18 @ el


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_savior` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_savior_18*" "green" >}} {{< bg "17" "pg_savior_17*" "green" >}} {{< bg "16" "pg_savior_16*" "green" >}} {{< bg "15" "pg_savior_15*" "green" >}} {{< bg "14" "pg_savior_14*" "green" >}} {{< bg "13" "pg_savior_13*" "green" >}} | `pg_savior_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-savior" "green" >}} {{< bg "17" "postgresql-17-pg-savior" "green" >}} {{< bg "16" "postgresql-16-pg-savior" "green" >}} {{< bg "15" "postgresql-15-pg-savior" "green" >}} {{< bg "14" "postgresql-14-pg-savior" "green" >}} {{< bg "13" "postgresql-13-pg-savior" "green" >}} | `postgresql-$v-pg-savior` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_savior_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-savior : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-savior : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-savior : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-savior : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pg_savior_18-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_18-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_savior_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [pg_savior_18-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_18-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_savior_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [pg_savior_18-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_18-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_savior_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [pg_savior_18-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_18-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_savior_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.4 KiB | [pg_savior_18-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_18-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_savior_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pg_savior_18-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_18-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-savior` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.7 KiB | [postgresql-18-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-18-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-savior` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.6 KiB | [postgresql-18-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-18-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-savior` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.7 KiB | [postgresql-18-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-18-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-savior` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.7 KiB | [postgresql-18-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-18-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-savior` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.7 KiB | [postgresql-18-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-18-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-savior` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.5 KiB | [postgresql-18-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-18-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-savior` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.6 KiB | [postgresql-18-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-18-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-savior` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.2 KiB | [postgresql-18-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-18-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pg_savior_17-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_17-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_savior_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [pg_savior_17-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_17-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_savior_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pg_savior_17-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_17-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_savior_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [pg_savior_17-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_17-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_savior_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.4 KiB | [pg_savior_17-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_17-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_savior_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 KiB | [pg_savior_17-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_17-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-savior` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.6 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-savior` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.8 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-savior` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.6 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-savior` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.9 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-savior` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 20.3 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-savior` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.0 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-savior` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-savior` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-17-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pg_savior_16-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_16-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_savior_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [pg_savior_16-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_16-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_savior_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pg_savior_16-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_16-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_savior_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [pg_savior_16-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_16-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_savior_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.4 KiB | [pg_savior_16-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_16-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_savior_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 KiB | [pg_savior_16-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_16-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-savior` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.6 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-savior` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.8 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-savior` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.6 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-savior` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.9 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-savior` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.8 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-savior` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.5 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-savior` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-savior` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-16-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pg_savior_15-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_15-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_savior_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [pg_savior_15-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_15-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_savior_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pg_savior_15-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_15-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_savior_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [pg_savior_15-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_15-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_savior_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.5 KiB | [pg_savior_15-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_15-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_savior_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 KiB | [pg_savior_15-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_15-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-savior` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.6 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-savior` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.9 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-savior` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.6 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-savior` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.9 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-savior` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 19.8 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-savior` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 19.5 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-savior` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-savior` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-15-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pg_savior_14-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_14-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_savior_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [pg_savior_14-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_14-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_savior_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pg_savior_14-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_14-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_savior_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [pg_savior_14-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_14-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_savior_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.5 KiB | [pg_savior_14-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_14-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_savior_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 KiB | [pg_savior_14-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_14-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-savior` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.6 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-savior` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.9 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-savior` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.6 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-savior` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.9 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-savior` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.6 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-savior` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.3 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-savior` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.4 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-savior` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-14-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_13` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.4 KiB | [pg_savior_13-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_13-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_savior_13` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [pg_savior_13-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_13-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_savior_13` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pg_savior_13-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_13-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_savior_13` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [pg_savior_13-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_13-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_savior_13` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.5 KiB | [pg_savior_13-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_13-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_savior_13` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pg_savior_13-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_13-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-savior` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.5 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-savior` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.6 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-savior` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.5 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-savior` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.7 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-savior` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.5 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-savior` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.2 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-savior` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.4 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-savior` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.2 KiB | [postgresql-13-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-13-pg-savior_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/viggy28/pg_savior" title="Repository" icon="github" subtitle="github.com/viggy28/pg_savior" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_savior-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_savior;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_savior;		# install via package name, for the active PG version

pig install pg_savior -v 18;   # install for PG 18
pig install pg_savior -v 17;   # install for PG 17
pig install pg_savior -v 16;   # install for PG 16
pig install pg_savior -v 15;   # install for PG 15
pig install pg_savior -v 14;   # install for PG 14
pig install pg_savior -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_savior;
```
