---
title: "pg_column_tetris"
linkTitle: "pg_column_tetris"
description: "Enforce optimal column alignment to minimize row padding"
weight: 5280
categories: ["ADMIN"]
width: full
---

[**pg_column_tetris**](https://github.com/rogerwelin/pg_column_tetris) : Enforce optimal column alignment to minimize row padding


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5280** | {{< badge content="pg_column_tetris" link="https://github.com/rogerwelin/pg_column_tetris" >}} | {{< ext "pg_column_tetris" >}} | `0.1.0` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `column_tetris` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pg_repack" >}} {{< ext "pgstattuple" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_dirtyread" >}} |

> [!Note] Upstream has no release or tag; source archive is normalized from commit e70f9867c63e932cdaf87b2d34b6504adad9ce12.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_column_tetris` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_column_tetris_18" "green" >}} {{< bg "17" "pg_column_tetris_17" "green" >}} {{< bg "16" "pg_column_tetris_16" "green" >}} {{< bg "15" "pg_column_tetris_15" "green" >}} {{< bg "14" "pg_column_tetris_14" "green" >}} | `pg_column_tetris_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-column-tetris" "green" >}} {{< bg "17" "postgresql-17-pg-column-tetris" "green" >}} {{< bg "16" "postgresql-16-pg-column-tetris" "green" >}} {{< bg "15" "postgresql-15-pg-column-tetris" "green" >}} {{< bg "14" "postgresql-14-pg-column-tetris" "green" >}} | `postgresql-$v-pg-column-tetris` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_column_tetris_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-column-tetris : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-column-tetris : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_column_tetris_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [pg_column_tetris_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_column_tetris_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_column_tetris_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [pg_column_tetris_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_column_tetris_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [pg_column_tetris_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_column_tetris_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.9 KiB | [pg_column_tetris_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_column_tetris_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_column_tetris_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_column_tetris_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.0 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.0 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.0 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.0 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.2 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.2 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.2 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.2 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.2 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pg-column-tetris` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.2 KiB | [postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-18-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_column_tetris_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [pg_column_tetris_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_column_tetris_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_column_tetris_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [pg_column_tetris_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_column_tetris_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [pg_column_tetris_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_column_tetris_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.9 KiB | [pg_column_tetris_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_column_tetris_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_column_tetris_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_column_tetris_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.0 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.0 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.0 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.0 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.2 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.2 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.2 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.2 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.2 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pg-column-tetris` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.2 KiB | [postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-17-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_column_tetris_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [pg_column_tetris_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_column_tetris_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_column_tetris_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [pg_column_tetris_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_column_tetris_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [pg_column_tetris_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_column_tetris_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.9 KiB | [pg_column_tetris_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_column_tetris_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_column_tetris_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_column_tetris_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.0 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.0 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.0 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.0 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.2 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.2 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.2 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.2 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.2 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pg-column-tetris` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.2 KiB | [postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-16-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_column_tetris_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [pg_column_tetris_15-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_column_tetris_15-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_15-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_column_tetris_15-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [pg_column_tetris_15-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_column_tetris_15-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [pg_column_tetris_15-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_column_tetris_15-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.9 KiB | [pg_column_tetris_15-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_column_tetris_15-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_column_tetris_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_15-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_column_tetris_15-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.0 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.0 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.0 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.0 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.2 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.2 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.2 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.2 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.2 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pg-column-tetris` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.2 KiB | [postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-15-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_column_tetris_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.0 KiB | [pg_column_tetris_14-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_column_tetris_14-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_14-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_column_tetris_14-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `pg_column_tetris_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.7 KiB | [pg_column_tetris_14-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_column_tetris_14-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [pg_column_tetris_14-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_column_tetris_14-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `pg_column_tetris_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.9 KiB | [pg_column_tetris_14-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_column_tetris_14-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `pg_column_tetris_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [pg_column_tetris_14-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_column_tetris_14-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.0 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.0 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.0 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.0 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.2 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.2 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.2 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.2 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.2 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pg-column-tetris` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.2 KiB | [postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-column-tetris/postgresql-14-pg-column-tetris_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rogerwelin/pg_column_tetris" title="Repository" icon="github" subtitle="github.com/rogerwelin/pg_column_tetris" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_column_tetris-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_column_tetris;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_column_tetris;		# install via package name, for the active PG version

pig install pg_column_tetris -v 18;   # install for PG 18
pig install pg_column_tetris -v 17;   # install for PG 17
pig install pg_column_tetris -v 16;   # install for PG 16
pig install pg_column_tetris -v 15;   # install for PG 15
pig install pg_column_tetris -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_column_tetris CASCADE; -- requires plpgsql
```

## Usage

Sources:

- [Project README](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/README.md)
- [Extension control file](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/pg_column_tetris.control)
- [Version 0.1.0 SQL implementation](https://github.com/rogerwelin/pg_column_tetris/blob/e70f9867c63e932cdaf87b2d34b6504adad9ce12/pg_column_tetris--0.1.0.sql)

`pg_column_tetris` 0.1.0 is a pure-SQL extension for PostgreSQL 14 through 18. An event trigger estimates alignment padding after `CREATE TABLE` and can warn about or reject inefficient column order. It also provides inspection and rewrite-suggestion functions.

### Inspect and choose enforcement

The default mode is `warn`; `strict` rejects a newly created table that the estimator considers suboptimal, and `off` disables the event-trigger check.

```sql
CREATE EXTENSION pg_column_tetris;

SELECT column_tetris.mode();
SELECT * FROM column_tetris.check('public.measurement'::regclass);
SELECT column_tetris.padding_wasted('public.measurement'::regclass);

SELECT column_tetris.set_mode('warn');
```

Use `column_tetris.exclude()` for tables that must not be checked. Temporary and system tables are skipped, and the event trigger checks table creation rather than every later alteration.

### Treat estimates and rewrites as advisory

The estimator models tuple headers and type alignment, but it cannot fully predict real storage for null bitmaps, variable-length or toasted values, compression, and workload-specific row populations. A reported byte saving is therefore a design signal, not measured disk reclamation.

`column_tetris.suggest_rewrite()` returns a migration script; it does not preserve every foreign key, index, trigger, or default. The generated sequence renames the original table, creates and copies a replacement, and drops the old table, which can require an exclusive lock and downtime. Never execute that output without reviewing dependent objects, privileges, identity and sequence behavior, replication, rollback, and a realistic staging rehearsal. Column order can also be part of application contracts such as positional inserts and row decoding.
