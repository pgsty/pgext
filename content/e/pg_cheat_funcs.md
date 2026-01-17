---
title: "pg_cheat_funcs"
linkTitle: "pg_cheat_funcs"
description: "Provides cheat (but useful) functions"
weight: 5220
categories: ["ADMIN"]
width: full
---

[**pg_cheat_funcs**](https://github.com/MasaoFujii/pg_cheat_funcs) : Provides cheat (but useful) functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5220** | {{< badge content="pg_cheat_funcs" link="https://github.com/MasaoFujii/pg_cheat_funcs" >}} | {{< ext "pg_cheat_funcs" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_crash" >}} {{< ext "pg_snakeoil" >}} {{< ext "pg_dirtyread" >}} {{< ext "pg_savior" >}} {{< ext "pg_surgery" >}} {{< ext "adminpack" >}} {{< ext "pageinspect" >}} {{< ext "pg_repack" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_cheat_funcs` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_cheat_funcs_18" "green" >}} {{< bg "17" "pg_cheat_funcs_17" "green" >}} {{< bg "16" "pg_cheat_funcs_16" "green" >}} {{< bg "15" "pg_cheat_funcs_15" "green" >}} {{< bg "14" "pg_cheat_funcs_14" "green" >}} {{< bg "13" "pg_cheat_funcs_13" "green" >}} | `pg_cheat_funcs_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-cheat-funcs" "green" >}} {{< bg "17" "postgresql-17-pg-cheat-funcs" "green" >}} {{< bg "16" "postgresql-16-pg-cheat-funcs" "green" >}} {{< bg "15" "postgresql-15-pg-cheat-funcs" "green" >}} {{< bg "14" "postgresql-14-pg-cheat-funcs" "green" >}} {{< bg "13" "postgresql-13-pg-cheat-funcs" "green" >}} | `postgresql-$v-pg-cheat-funcs` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_cheat_funcs_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-cheat-funcs : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-cheat-funcs : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-cheat-funcs : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-cheat-funcs : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-cheat-funcs : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-cheat-funcs : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-cheat-funcs : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-cheat-funcs : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-cheat-funcs : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cheat_funcs_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.7 KiB | [pg_cheat_funcs_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 48.2 KiB | [pg_cheat_funcs_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.2 KiB | [pg_cheat_funcs_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.0 KiB | [pg_cheat_funcs_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cheat_funcs_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 53.3 KiB | [pg_cheat_funcs_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cheat_funcs_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cheat_funcs_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.3 KiB | [pg_cheat_funcs_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cheat_funcs_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-cheat-funcs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 93.1 KiB | [postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-cheat-funcs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 92.3 KiB | [postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-cheat-funcs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 93.1 KiB | [postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-cheat-funcs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 92.2 KiB | [postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-cheat-funcs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 118.9 KiB | [postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-cheat-funcs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 118.1 KiB | [postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-cheat-funcs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.4 KiB | [postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-cheat-funcs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 110.8 KiB | [postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-18-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cheat_funcs_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.8 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 48.2 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.2 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.2 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cheat_funcs_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 53.2 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cheat_funcs_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cheat_funcs_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.3 KiB | [pg_cheat_funcs_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cheat_funcs_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-cheat-funcs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 93.9 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-cheat-funcs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 93.1 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-cheat-funcs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 93.8 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-cheat-funcs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 93.1 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-cheat-funcs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 124.5 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-cheat-funcs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.9 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-cheat-funcs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.4 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-cheat-funcs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 110.9 KiB | [postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-17-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cheat_funcs_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.8 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 48.2 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.1 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.3 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cheat_funcs_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 53.2 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cheat_funcs_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cheat_funcs_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.4 KiB | [pg_cheat_funcs_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cheat_funcs_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-cheat-funcs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 93.7 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-cheat-funcs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 93.1 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-cheat-funcs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 93.8 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-cheat-funcs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 93.1 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-cheat-funcs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 123.8 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-cheat-funcs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.1 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-cheat-funcs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.4 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-cheat-funcs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 110.9 KiB | [postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-16-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cheat_funcs_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.0 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 48.5 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.3 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.6 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cheat_funcs_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 53.4 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cheat_funcs_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cheat_funcs_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.6 KiB | [pg_cheat_funcs_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cheat_funcs_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-cheat-funcs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 94.4 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-cheat-funcs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 93.7 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-cheat-funcs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 94.2 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-cheat-funcs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 93.5 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-cheat-funcs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 124.3 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-cheat-funcs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 124.1 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-cheat-funcs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 112.0 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-cheat-funcs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 111.2 KiB | [postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-15-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cheat_funcs_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 48.9 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 48.4 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.2 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.5 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cheat_funcs_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 53.4 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cheat_funcs_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cheat_funcs_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.6 KiB | [pg_cheat_funcs_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cheat_funcs_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-cheat-funcs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 93.4 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-cheat-funcs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 92.6 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-cheat-funcs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 93.2 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-cheat-funcs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 92.5 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-cheat-funcs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 124.6 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-cheat-funcs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.7 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-cheat-funcs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 111.8 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-cheat-funcs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 110.9 KiB | [postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-14-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cheat_funcs_13` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.4 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cheat_funcs_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cheat_funcs_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cheat_funcs_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cheat_funcs_13` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.0 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cheat_funcs_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cheat_funcs_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.0 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cheat_funcs_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cheat_funcs_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 54.0 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cheat_funcs_13-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cheat_funcs_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.2 KiB | [pg_cheat_funcs_13-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cheat_funcs_13-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-cheat-funcs` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 96.0 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-cheat-funcs` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 95.3 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-cheat-funcs` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 96.1 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-cheat-funcs` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 95.4 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-cheat-funcs` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 126.1 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-cheat-funcs` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 125.9 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-cheat-funcs` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.4 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-cheat-funcs` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 113.1 KiB | [postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cheat-funcs/postgresql-13-pg-cheat-funcs_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MasaoFujii/pg_cheat_funcs" title="Repository" icon="github" subtitle="github.com/MasaoFujii/pg_cheat_funcs" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_cheat_funcs-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_cheat_funcs;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_cheat_funcs;		# install via package name, for the active PG version

pig install pg_cheat_funcs -v 18;   # install for PG 18
pig install pg_cheat_funcs -v 17;   # install for PG 17
pig install pg_cheat_funcs -v 16;   # install for PG 16
pig install pg_cheat_funcs -v 15;   # install for PG 15
pig install pg_cheat_funcs -v 14;   # install for PG 14
pig install pg_cheat_funcs -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_cheat_funcs;
```
