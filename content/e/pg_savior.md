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
| **5810** | {{< badge content="pg_savior" link="https://github.com/viggy28/pg_savior" >}} | {{< ext "pg_savior" >}} | `0.1.0` | {{< category "ADMIN" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_savior` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_savior_18" "green" >}} {{< bg "17" "pg_savior_17" "green" >}} {{< bg "16" "pg_savior_16" "green" >}} {{< bg "15" "pg_savior_15" "green" >}} {{< bg "14" "pg_savior_14" "green" >}} | `pg_savior_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-savior" "green" >}} {{< bg "17" "postgresql-17-pg-savior" "green" >}} {{< bg "16" "postgresql-16-pg-savior" "green" >}} {{< bg "15" "postgresql-15-pg-savior" "green" >}} {{< bg "14" "postgresql-14-pg-savior" "green" >}} | `postgresql-$v-pg-savior` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_savior_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-savior : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-savior : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.3 KiB | [pg_savior_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_savior_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_savior_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_savior_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.0 KiB | [pg_savior_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_savior_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.1 KiB | [pg_savior_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_savior_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.4 KiB | [pg_savior_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-savior` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.7 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.9 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.7 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 18.6 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.4 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.4 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.6 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-savior` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.0 KiB | [postgresql-18-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-18-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.3 KiB | [pg_savior_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_savior_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_savior_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_savior_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.0 KiB | [pg_savior_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_savior_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.1 KiB | [pg_savior_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_savior_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.4 KiB | [pg_savior_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-savior` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.7 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.7 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.9 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.2 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.1 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.3 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.3 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.6 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-savior` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.0 KiB | [postgresql-17-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-17-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.3 KiB | [pg_savior_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_savior_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_savior_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_savior_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.0 KiB | [pg_savior_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_savior_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.1 KiB | [pg_savior_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_savior_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.4 KiB | [pg_savior_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-savior` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.7 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.7 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.9 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.7 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.6 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.4 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.3 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.6 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-savior` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.0 KiB | [postgresql-16-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-16-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.3 KiB | [pg_savior_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_savior_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_savior_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_savior_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.0 KiB | [pg_savior_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_savior_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.1 KiB | [pg_savior_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_savior_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.4 KiB | [pg_savior_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-savior` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.7 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.9 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.7 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.6 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.3 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.3 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.6 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-savior` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 19.0 KiB | [postgresql-15-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-15-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_savior_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 17.3 KiB | [pg_savior_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_savior_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_savior_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.4 KiB | [pg_savior_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_savior_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_savior_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 17.1 KiB | [pg_savior_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_savior_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_savior_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 17.0 KiB | [pg_savior_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_savior_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_savior_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 17.1 KiB | [pg_savior_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_savior_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_savior_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.4 KiB | [pg_savior_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_savior_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-savior` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.7 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.7 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.7 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.9 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.7 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.6 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 18.3 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 18.3 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 18.6 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-savior` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 18.9 KiB | [postgresql-14-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-savior/postgresql-14-pg-savior_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/viggy28/pg_savior" title="Repository" icon="github" subtitle="github.com/viggy28/pg_savior" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_savior-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_savior;		# build rpm/deb
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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_savior;
```

## Usage

Sources: [README](https://github.com/viggy28/pg_savior/blob/v0.1.0/README.md), [release 0.1.0](https://github.com/viggy28/pg_savior/releases/tag/v0.1.0), [PGXN 0.1.0](https://pgxn.org/dist/pg_savior/0.1.0/), [SQL file](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior--0.1.0.sql), [C source](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior.c), [pg_savior.control](https://github.com/viggy28/pg_savior/blob/v0.1.0/pg_savior.control)

`pg_savior` is a PostgreSQL safety extension for blocking specific high-risk DML and DDL statements before they run. Version `0.1.0` is a deliberate PGXN release and a major rewrite from `0.0.1`; the README still labels it pre-1.0 and not production-ready.

### Activation

`CREATE EXTENSION` alone does not activate the checks. The SQL file only documents that protection lives in the loaded shared library, so each backend must load `pg_savior` by one of the upstream-supported paths:

Cluster-wide activation uses `shared_preload_libraries` and requires a PostgreSQL restart:

```conf
shared_preload_libraries = 'pg_savior'
```

Per-session activation for new connections can use `session_preload_libraries` after a config reload:

```conf
session_preload_libraries = 'pg_savior'
```

For development or test sessions, load the library manually:

```sql
LOAD 'pg_savior';
CREATE EXTENSION pg_savior;
```

Once the library is loaded, `_PG_init` installs `post_parse_analyze_hook`, `ExecutorStart_hook`, and `ProcessUtility_hook` for that backend.

### DML Guards

`pg_savior` blocks `DELETE` and `UPDATE` statements that have no `WHERE` clause. The parser hook checks the analyzed `Query` tree and raises `ERROR`, so the transaction aborts and the application sees the failure.

```sql
CREATE TABLE emp (id int);
INSERT INTO emp VALUES (1), (2), (3);

DELETE FROM emp;
-- ERROR: pg_savior: DELETE without WHERE clause is blocked

UPDATE emp SET id = id + 1;
-- ERROR: pg_savior: UPDATE without WHERE clause is blocked

DELETE FROM emp WHERE id = 1;
-- allowed
```

The optional row-count guard applies to `DELETE` and `UPDATE` statements whose planner estimate exceeds `pg_savior.max_rows_affected`. It runs from `ExecutorStart_hook`, after planning and before tuples are touched.

```sql
SET pg_savior.max_rows_affected = 100;

DELETE FROM emp WHERE id > 0;
-- blocked if the planner estimate is greater than 100 rows
```

### DDL Guards

The `ProcessUtility_hook` guards only the DDL operations listed by upstream:

- `CREATE INDEX` without `CONCURRENTLY` is always blocked.
- `DROP DATABASE` is always blocked.
- `ALTER TABLE ADD COLUMN ... DEFAULT` is blocked when the target table is larger than `pg_savior.large_table_threshold_rows`.
- `ALTER TABLE ALTER COLUMN TYPE` is blocked for large tables.
- `TRUNCATE` is blocked when any target table is large.
- `DROP TABLE` is blocked when any target table is large.

Large-table checks use `pg_class.reltuples > pg_savior.large_table_threshold_rows`.

```sql
CREATE INDEX emp_idx ON emp (id);
-- ERROR: pg_savior: CREATE INDEX without CONCURRENTLY is blocked

CREATE INDEX CONCURRENTLY emp_idx ON emp (id);
-- allowed by this guard

ALTER TABLE big_emp ADD COLUMN status text DEFAULT 'active';
-- blocked when big_emp is over pg_savior.large_table_threshold_rows

TRUNCATE big_emp;
-- blocked when big_emp is over pg_savior.large_table_threshold_rows
```

### Configuration

All documented GUCs are session-scoped `USERSET` variables:

| GUC | Default | Effect |
|---|---:|---|
| `pg_savior.enabled` | `on` | Master switch; when `off`, checks do not run. |
| `pg_savior.bypass` | `off` | Allows the current session through the guards. |
| `pg_savior.max_rows_affected` | `0` | Blocks estimated `DELETE`/`UPDATE` row counts above this value; `0` disables the check. |
| `pg_savior.large_table_threshold_rows` | `1000000` | Defines "large" for the guarded large-table DDL operations. |

Use `SET LOCAL` for a deliberate one-transaction bypass:

```sql
BEGIN;
SET LOCAL pg_savior.bypass = on;
DELETE FROM staging_table;
COMMIT;
```

### Caveats

- The library must be loaded in the backend before protection exists; `CREATE EXTENSION pg_savior` only registers extension metadata.
- The row-count and large-table guards depend on planner/catalog estimates. Run `ANALYZE` when recent changes make estimates stale.
- `UPDATE` coverage is limited to unguarded `UPDATE` and the optional planner-estimate threshold; the README does not claim semantic review of every `WHERE` predicate.
- DDL coverage is limited to the listed `ProcessUtility_hook` cases. Do not assume other schema changes are blocked.
- The `ADD COLUMN ... DEFAULT` guard is conservative and blocks any default on a large table, including non-volatile defaults that newer PostgreSQL versions may handle without a full table rewrite.
