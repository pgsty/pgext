---
title: "pg_living_assertions"
linkTitle: "pg_living_assertions"
description: "Executable SQL checks with verdict dates and assertion history"
weight: 5300
categories: ["ADMIN"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_living_assertions**](https://github.com/Manuelreyesbravo/pg_living_assertions) : Executable SQL checks with verdict dates and assertion history


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5300** | {{< badge content="pg_living_assertions" link="https://github.com/Manuelreyesbravo/pg_living_assertions" >}} | {{< ext "pg_living_assertions" >}} | `0.5.1` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `living_assertions` |
|    **Need By**    | {{< ext "pg_grammar_guard" >}} |

> [!Note] On-demand SQL checks run in a read-only subtransaction that is always rolled back; not per-write SQL ASSERTION constraints.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_living_assertions` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "18" "pg_living_assertions_18" "green" >}} {{< bg "17" "pg_living_assertions_17" "green" >}} {{< bg "16" "pg_living_assertions_16" "green" >}} {{< bg "15" "pg_living_assertions_15" "green" >}} {{< bg "14" "pg_living_assertions_14" "green" >}} | `pg_living_assertions_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "18" "postgresql-18-pg-living-assertions" "green" >}} {{< bg "17" "postgresql-17-pg-living-assertions" "green" >}} {{< bg "16" "postgresql-16-pg-living-assertions" "green" >}} {{< bg "15" "postgresql-15-pg-living-assertions" "green" >}} {{< bg "14" "postgresql-14-pg-living-assertions" "green" >}} | `postgresql-$v-pg-living-assertions` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_living_assertions_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-living-assertions : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-14-pg-living-assertions : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_18` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_18-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_18-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_18-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_18-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_18-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_18` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_18-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_18-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-18-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-18-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-18-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_17` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_17-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_17-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_17-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_17-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_17-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_17-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_17-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-17-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-17-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-17-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_16-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_16-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_16-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_16-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_16-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_16-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_16-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-16-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-16-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-16-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_15-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_15-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_15-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_15-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_15-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_15-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_15-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-15-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-15-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-15-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_living_assertions_14` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.1 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_living_assertions_14-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.1 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_living_assertions_14-0.5.1-1PGSTY.el8.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.8 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_living_assertions_14-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_living_assertions_14-0.5.1-1PGSTY.el9.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.9 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_living_assertions_14-0.5.1-1PGSTY.el10.noarch.rpm) |
| `pg_living_assertions_14` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.9 KiB | [pg_living_assertions_14-0.5.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_living_assertions_14-0.5.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.2 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 23.2 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.0 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.9 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.9 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.9 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
| `postgresql-14-pg-living-assertions` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 23.9 KiB | [postgresql-14-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-living-assertions/postgresql-14-pg-living-assertions_0.5.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Manuelreyesbravo/pg_living_assertions" title="Repository" icon="github" subtitle="github.com/Manuelreyesbravo/pg_living_assertions" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_living_assertions-0.5.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_living_assertions;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_living_assertions;		# install via package name, for the active PG version

pig install pg_living_assertions -v 18;   # install for PG 18
pig install pg_living_assertions -v 17;   # install for PG 17
pig install pg_living_assertions -v 16;   # install for PG 16
pig install pg_living_assertions -v 15;   # install for PG 15
pig install pg_living_assertions -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_living_assertions;
```

## Usage

Sources:

- [README.md](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/README.md)
- [pg_living_assertions.control](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions.control)
- [pg_living_assertions--0.4.1--0.5.0.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions--0.4.1--0.5.0.sql)
- [pg_living_assertions--0.5.0--0.5.1.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions--0.5.0--0.5.1.sql)
- [test/sql/read_only.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/test/sql/read_only.sql)

`pg_living_assertions` 0.5.1 stores SQL checks, verdicts, verification times and replacement history. These are checks run on demand, not SQL ASSERTION constraints evaluated on every write. It is a pure SQL extension with no preload requirement.

### Register and Verify

```sql
CREATE EXTENSION pg_living_assertions;
SELECT living_assertions.declare(
  'simple_check', 'one equals one',
  $$SELECT 1 = 1 AS holds, 'arithmetic check'::text AS detail$$);
SELECT living_assertions.run('simple_check');
SELECT name, state, age FROM living_assertions.status;
```

### Results and History

A check must return exactly one row with boolean `holds` and optional text `detail`. `living_assertions.run_all()` evaluates registered checks. `living_assertions.state()` distinguishes holds, broken, unknown, erroring, unchecked, retired and unregistered; `living_assertions.stale()` separates never-checked assertions from old results. `living_assertions.declare_unchanged()` records an expression for later text comparison, so the author must canonicalize its output. Definitions are superseded with a reason; results and registry data are included in database dumps.

### Execution and Privileges

Since 0.5.0 the evaluator runs read-only inside a subtransaction that is always rolled back, preserving its verdict. This repairs the older STABLE-only evaluator, which did not stop side effects through volatile functions. It is not a sandbox for untrusted SQL: temporary-sequence changes, session advisory locks and external effects can survive. Only trusted administrators should register checks; they execute with the privileges of the later caller. Registry tables belong to the extension owner and write functions are revoked from `PUBLIC` by default.

### Upgrade

After installing the matching files, use `ALTER EXTENSION pg_living_assertions UPDATE TO '0.5.1'`. The 0.4.1→0.5.0→0.5.1 chain replaces evaluator functions without changing registry tables. The final patch qualifies row types in `run()` so type-cache invalidation does not resolve them under an assertion's unrelated search path.

Fresh installation also uses an earlier base SQL script followed by the packaged upgrade chain; keep the complete set of matching scripts installed.
