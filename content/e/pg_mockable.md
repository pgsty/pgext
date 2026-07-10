---
title: "pg_mockable"
linkTitle: "pg_mockable"
description: "Create mockable wrappers for PostgreSQL functions in tests"
weight: 3120
categories: ["LANG"]
width: full
---

[**pg_mockable**](https://github.com/bigsmoke/pg_mockable) : Create mockable wrappers for PostgreSQL functions in tests


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3120** | {{< badge content="pg_mockable" link="https://github.com/bigsmoke/pg_mockable" >}} | {{< ext "pg_mockable" >}} | `1.1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `mockable` |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_mockable` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "pg_mockable_18" "green" >}} {{< bg "17" "pg_mockable_17" "green" >}} {{< bg "16" "pg_mockable_16" "green" >}} {{< bg "15" "pg_mockable_15" "green" >}} {{< bg "14" "pg_mockable_14" "green" >}} | `pg_mockable_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-pg-mockable" "green" >}} {{< bg "17" "postgresql-17-pg-mockable" "green" >}} {{< bg "16" "postgresql-16-pg-mockable" "green" >}} {{< bg "15" "postgresql-15-pg-mockable" "green" >}} {{< bg "14" "postgresql-14-pg-mockable" "green" >}} | `postgresql-$v-pg-mockable` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_mockable_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-mockable : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-mockable : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mockable_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [pg_mockable_18-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mockable_18-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.3 KiB | [pg_mockable_18-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mockable_18-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.7 KiB | [pg_mockable_18-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mockable_18-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [pg_mockable_18-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mockable_18-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [pg_mockable_18-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mockable_18-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_mockable_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.8 KiB | [pg_mockable_18-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mockable_18-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-mockable` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 22.3 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.3 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 22.3 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.3 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.0 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.0 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.8 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.8 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.8 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pg-mockable` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.8 KiB | [postgresql-18-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-18-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mockable_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [pg_mockable_17-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mockable_17-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.3 KiB | [pg_mockable_17-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mockable_17-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.7 KiB | [pg_mockable_17-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mockable_17-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [pg_mockable_17-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mockable_17-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [pg_mockable_17-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mockable_17-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_mockable_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.8 KiB | [pg_mockable_17-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mockable_17-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-mockable` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 22.3 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.3 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 22.3 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.3 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.0 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.0 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.8 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.8 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.8 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pg-mockable` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.8 KiB | [postgresql-17-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-17-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mockable_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [pg_mockable_16-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mockable_16-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.3 KiB | [pg_mockable_16-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mockable_16-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.7 KiB | [pg_mockable_16-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mockable_16-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [pg_mockable_16-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mockable_16-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [pg_mockable_16-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mockable_16-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_mockable_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.8 KiB | [pg_mockable_16-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mockable_16-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-mockable` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 22.3 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.3 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 22.3 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.3 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.0 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.0 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.8 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.8 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.8 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pg-mockable` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.8 KiB | [postgresql-16-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-16-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mockable_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [pg_mockable_15-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mockable_15-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.3 KiB | [pg_mockable_15-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mockable_15-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.7 KiB | [pg_mockable_15-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mockable_15-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [pg_mockable_15-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mockable_15-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [pg_mockable_15-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mockable_15-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_mockable_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.8 KiB | [pg_mockable_15-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mockable_15-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-mockable` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 22.3 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.3 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 22.3 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.3 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.0 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.0 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.8 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.8 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.8 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pg-mockable` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.8 KiB | [postgresql-15-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-15-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mockable_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 28.4 KiB | [pg_mockable_14-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mockable_14-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.3 KiB | [pg_mockable_14-1.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mockable_14-1.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_mockable_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.7 KiB | [pg_mockable_14-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mockable_14-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [pg_mockable_14-1.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mockable_14-1.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_mockable_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.8 KiB | [pg_mockable_14-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mockable_14-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_mockable_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.8 KiB | [pg_mockable_14-1.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mockable_14-1.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-mockable` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 22.3 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.3 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 22.3 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.3 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.0 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.0 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.8 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.8 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.8 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pg-mockable` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.8 KiB | [postgresql-14-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mockable/postgresql-14-pg-mockable_1.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_mockable" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_mockable" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_mockable-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_mockable;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_mockable;		# install via package name, for the active PG version

pig install pg_mockable -v 18;   # install for PG 18
pig install pg_mockable -v 17;   # install for PG 17
pig install pg_mockable -v 16;   # install for PG 16
pig install pg_mockable -v 15;   # install for PG 15
pig install pg_mockable -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_mockable;
```




## Usage

> Sources: [pg_mockable upstream README](https://github.com/bigsmoke/pg_mockable/blob/v1.1.0/README.md), [v1.1.0 tag](https://github.com/bigsmoke/pg_mockable/tree/v1.1.0), [PGXN pg_mockable](https://pgxn.org/dist/pg_mockable/), local source tarball `pg_mockable-1.1.0.tar.gz`.

`pg_mockable` creates mockable wrapper functions for PostgreSQL routines. It is mainly useful in database tests where application code should call a stable wrapper, while tests temporarily replace the wrapper's return value.

```sql
CREATE EXTENSION pg_mockable CASCADE;
```

The extension installs into the fixed `mockable` schema and is not relocatable.

### Mock Built-In Time Functions

`mockable.now()` is pre-created because mocking `now()` also covers the related current-time wrappers exposed by this extension.

```sql
SELECT mockable.now();

SELECT mockable.mock(
  'pg_catalog.now()',
  '2026-06-17 10:00:00+08'::timestamptz
);

SELECT mockable.now();
SELECT mockable.current_timestamp();
SELECT mockable.current_date();

CALL mockable.unmock('pg_catalog.now()');
```

`mockable.mock(regprocedure, anyelement)` stores the mock value and returns it. `mockable.unmock(regprocedure)` clears the mock and restores the wrapper to call the original routine.

### Wrap Application Functions

Use `mockable.wrap_function()` to create a thin wrapper in the `mockable` schema:

```sql
CREATE SCHEMA app;

CREATE FUNCTION app.answer()
RETURNS int
LANGUAGE sql
RETURN 42;

SELECT mockable.wrap_function('app.answer()');

SELECT mockable.answer();                 -- 42
SELECT mockable.mock('app.answer()', 7);   -- 7
SELECT mockable.answer();                 -- 7

CALL mockable.unmock('app.answer()');
SELECT mockable.answer();                 -- 42
```

The first argument is a `regprocedure`, so include argument types when the function is overloaded:

```sql
SELECT mockable.wrap_function('pg_catalog.current_setting(text, boolean)');
```

If automatic wrapper generation is not sufficient, pass the exact `CREATE OR REPLACE FUNCTION` statement as the second argument:

```sql
SELECT mockable.wrap_function(
  'app.answer()',
  $$
  CREATE OR REPLACE FUNCTION mockable.answer()
  RETURNS int
  LANGUAGE sql
  RETURN app.answer();
  $$
);
```

Version 1.1.0 also adds optional debug logging for wrapped/mockable routines through `raise_debug_messages$` on `mockable.wrap_function(...)` and the `mock_memory.raise_debug_messages` column.

### Mock Lifetime

The default mock lifetime is transaction-scoped. For values that must survive dump/restore or later transactions, create the wrapper with a persistent lifetime:

```sql
SELECT mockable.wrap_function(
  'app.answer()',
  mock_duration$ => 'PERSISTENT'
);
```

Persistent mocks should be explicitly cleared when the test fixture no longer needs them:

```sql
CALL mockable.unmock('app.answer()');
```

### Search Path Caveat

Application code must actually call the wrapper, for example `mockable.now()` or `mockable.answer()`, for the mock to apply. Some PL/pgSQL code can be redirected by adjusting `search_path`, but expressions such as table defaults are compiled to function OIDs; adding `mockable` to `search_path` later does not rewrite those references. Prefer explicit `mockable.*` calls in code that is meant to be testable.

### Caveats

- Version 1.1.0 supports PostgreSQL 14-18. It is a SQL extension and does not need `shared_preload_libraries`.
- `pg_mockable` owns the `mockable` schema; installing it in another schema is not supported by the control file.
- Wrapper privileges are derived from the wrapped routine. The tests verify that wrapping a private function does not grant execute privilege to roles that could not call the original function.
