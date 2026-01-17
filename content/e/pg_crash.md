---
title: "pg_crash"
linkTitle: "pg_crash"
description: "Send random signals to random processes"
weight: 5210
categories: ["ADMIN"]
width: full
---

[**pg_crash**](https://github.com/cybertec-postgresql/pg_crash) : Send random signals to random processes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5210** | {{< badge content="pg_crash" link="https://github.com/cybertec-postgresql/pg_crash" >}} | {{< ext "pg_crash" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_snakeoil" >}} {{< ext "pg_cheat_funcs" >}} {{< ext "pg_savior" >}} {{< ext "pg_dirtyread" >}} {{< ext "pg_surgery" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_crash` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_crash_18" "green" >}} {{< bg "17" "pg_crash_17" "green" >}} {{< bg "16" "pg_crash_16" "green" >}} {{< bg "15" "pg_crash_15" "green" >}} {{< bg "14" "pg_crash_14" "green" >}} {{< bg "13" "pg_crash_13" "green" >}} | `pg_crash_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-crash" "green" >}} {{< bg "17" "postgresql-17-pg-crash" "green" >}} {{< bg "16" "postgresql-16-pg-crash" "green" >}} {{< bg "15" "postgresql-15-pg-crash" "green" >}} {{< bg "14" "postgresql-14-pg-crash" "green" >}} {{< bg "13" "postgresql-13-pg-crash" "green" >}} | `postgresql-$v-pg-crash` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_crash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_crash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_crash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_crash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_crash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_crash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_crash_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-crash : AVAIL 2" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-crash : AVAIL 2" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-17-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-16-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-15-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-14-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-13-pg-crash : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-17-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-16-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-15-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-14-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.3" "postgresql-13-pg-crash : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-crash : AVAIL 2" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-crash : AVAIL 2" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.3" "postgresql-18-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-crash : AVAIL 2" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.3" "postgresql-18-pg-crash : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-crash : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-crash : AVAIL 2" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_crash_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 KiB | [pg_crash_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.2 KiB | [pg_crash_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [pg_crash_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 12.9 KiB | [pg_crash_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [pg_crash_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_crash_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_crash_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.2 KiB | [pg_crash_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_crash_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-crash` | `0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-18-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-crash` | `0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-18-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-crash` | `0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-18-pg-crash_0.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-crash` | `0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 13.0 KiB | [postgresql-18-pg-crash_0.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-crash` | `0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.7 KiB | [postgresql-18-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-crash` | `0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.8 KiB | [postgresql-18-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-crash` | `0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-18-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-crash` | `0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.0 KiB | [postgresql-18-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-18-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_crash_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 KiB | [pg_crash_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.2 KiB | [pg_crash_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [pg_crash_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.0 KiB | [pg_crash_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [pg_crash_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_crash_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_crash_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.2 KiB | [pg_crash_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_crash_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-crash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-crash` | `0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-17-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-crash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-crash` | `0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-17-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-crash` | `0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-17-pg-crash_0.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-crash` | `0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.9 KiB | [postgresql-17-pg-crash_0.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-crash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-crash` | `0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.0 KiB | [postgresql-17-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-crash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.4 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-crash` | `0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.1 KiB | [postgresql-17-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-crash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.0 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-crash` | `0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-17-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-crash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.0 KiB | [postgresql-17-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-17-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-crash` | `0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.0 KiB | [postgresql-17-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-17-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_crash_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 KiB | [pg_crash_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.2 KiB | [pg_crash_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [pg_crash_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.0 KiB | [pg_crash_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [pg_crash_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_crash_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_crash_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.1 KiB | [pg_crash_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_crash_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-crash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-crash` | `0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-16-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-crash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-crash` | `0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.9 KiB | [postgresql-16-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-crash` | `0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-16-pg-crash_0.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-crash` | `0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 13.0 KiB | [postgresql-16-pg-crash_0.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-crash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-crash` | `0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.0 KiB | [postgresql-16-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-crash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.4 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-crash` | `0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.1 KiB | [postgresql-16-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-crash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.0 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-crash` | `0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.9 KiB | [postgresql-16-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-crash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.0 KiB | [postgresql-16-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-16-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-crash` | `0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.0 KiB | [postgresql-16-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-16-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_crash_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 KiB | [pg_crash_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.2 KiB | [pg_crash_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [pg_crash_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.0 KiB | [pg_crash_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [pg_crash_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_crash_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_crash_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.2 KiB | [pg_crash_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_crash_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-crash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-crash` | `0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-15-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-crash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-crash` | `0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-15-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-crash` | `0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-15-pg-crash_0.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-crash` | `0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 13.0 KiB | [postgresql-15-pg-crash_0.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-crash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-crash` | `0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.0 KiB | [postgresql-15-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-crash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.4 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-crash` | `0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.1 KiB | [postgresql-15-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-crash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.0 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-crash` | `0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.9 KiB | [postgresql-15-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-crash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.0 KiB | [postgresql-15-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-15-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-crash` | `0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.0 KiB | [postgresql-15-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-15-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_crash_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 KiB | [pg_crash_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.2 KiB | [pg_crash_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [pg_crash_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.0 KiB | [pg_crash_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [pg_crash_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_crash_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_crash_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.2 KiB | [pg_crash_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_crash_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-crash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-crash` | `0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-14-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-crash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-crash` | `0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-14-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-crash` | `0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-14-pg-crash_0.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-crash` | `0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 13.0 KiB | [postgresql-14-pg-crash_0.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-crash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-crash` | `0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.0 KiB | [postgresql-14-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-crash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.4 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-crash` | `0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.1 KiB | [postgresql-14-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-crash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.0 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-crash` | `0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-14-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-crash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.0 KiB | [postgresql-14-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-14-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-crash` | `0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.0 KiB | [postgresql-14-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-14-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_crash_13` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.9 KiB | [pg_crash_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_crash_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_crash_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.2 KiB | [pg_crash_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_crash_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_crash_13` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [pg_crash_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_crash_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_crash_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.0 KiB | [pg_crash_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_crash_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_crash_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.0 KiB | [pg_crash_13-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_crash_13-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_crash_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.2 KiB | [pg_crash_13-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_crash_13-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-crash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-crash` | `0.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.5 KiB | [postgresql-13-pg-crash_0.3-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-crash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-crash` | `0.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.7 KiB | [postgresql-13-pg-crash_0.3-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-crash` | `0.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.5 KiB | [postgresql-13-pg-crash_0.3-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg13+1_amd64.deb) |
| `postgresql-13-pg-crash` | `0.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-13-pg-crash_0.3-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg13+1_arm64.deb) |
| `postgresql-13-pg-crash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-crash` | `0.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.8 KiB | [postgresql-13-pg-crash_0.3-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-crash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.4 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-crash` | `0.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.9 KiB | [postgresql-13-pg-crash_0.3-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-crash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.0 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-crash` | `0.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.6 KiB | [postgresql-13-pg-crash_0.3-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-crash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.0 KiB | [postgresql-13-pg-crash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-crash/postgresql-13-pg-crash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-crash` | `0.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-13-pg-crash_0.3-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-crash/postgresql-13-pg-crash_0.3-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_crash" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_crash" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_crash-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_crash;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_crash;		# install via package name, for the active PG version

pig install pg_crash -v 18;   # install for PG 18
pig install pg_crash -v 17;   # install for PG 17
pig install pg_crash -v 16;   # install for PG 16
pig install pg_crash -v 15;   # install for PG 15
pig install pg_crash -v 14;   # install for PG 14
pig install pg_crash -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_crash';
```


This extension does not need `CREATE EXTENSION` DDL command


