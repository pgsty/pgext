---
title: "pg_when"
linkTitle: "pg_when"
description: "Natural language time parsing for PostgreSQL"
weight: 1120
categories: ["TIME"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_when**](https://github.com/frectonz/pg-when) : Natural language time parsing for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1120** | {{< badge content="pg_when" link="https://github.com/frectonz/pg-when" >}} | {{< ext "pg_when" >}} | `0.1.10` | {{< category "TIME" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pgcalendar" >}} {{< ext "pg_rrule" >}} {{< ext "cron_utils" >}} {{< ext "pgagent" >}} {{< ext "pg_task" >}} {{< ext "pg_dbms_job" >}} {{< ext "pg_duration" >}} {{< ext "pg_bikram_sambat" >}} {{< ext "pg_dispatch" >}} |

> [!Note] Upstream 0.1.10 supports PostgreSQL 13-18 and pins pgrx 0.18.1; PIGSTY packages PostgreSQL 14-18 with a locked pgrx 0.19.1 compatibility update.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.10` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_when` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.10` | {{< bg "18" "pg_when_18" "green" >}} {{< bg "17" "pg_when_17" "green" >}} {{< bg "16" "pg_when_16" "green" >}} {{< bg "15" "pg_when_15" "green" >}} {{< bg "14" "pg_when_14" "green" >}} | `pg_when_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.10` | {{< bg "18" "postgresql-18-pg-when" "green" >}} {{< bg "17" "postgresql-17-pg-when" "green" >}} {{< bg "16" "postgresql-16-pg-when" "green" >}} {{< bg "15" "postgresql-15-pg-when" "green" >}} {{< bg "14" "postgresql-14-pg-when" "green" >}} | `postgresql-$v-pg-when` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "pg_when_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-18-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-17-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-16-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-15-pg-when : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.10" "postgresql-14-pg-when : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_18` | `0.1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pg_when_18-0.1.10-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_18-0.1.10-1PGSTY.el8.x86_64.rpm) |
| `pg_when_18` | `0.1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 973.7 KiB | [pg_when_18-0.1.10-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_18-0.1.10-1PGSTY.el8.aarch64.rpm) |
| `pg_when_18` | `0.1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_when_18-0.1.10-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_18-0.1.10-1PGSTY.el9.x86_64.rpm) |
| `pg_when_18` | `0.1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_when_18-0.1.10-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_18-0.1.10-1PGSTY.el9.aarch64.rpm) |
| `pg_when_18` | `0.1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_when_18-0.1.10-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_18-0.1.10-1PGSTY.el10.x86_64.rpm) |
| `pg_when_18` | `0.1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1016.2 KiB | [pg_when_18-0.1.10-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_18-0.1.10-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-when` | `0.1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 882.4 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 755.0 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 882.8 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 756.4 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 977.1 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 887.3 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 967.6 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 877.5 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [u26.x86_64](/os/u26.x86_64) | pigsty | 964.5 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-when` | `0.1.10` | [u26.aarch64](/os/u26.aarch64) | pigsty | 875.7 KiB | [postgresql-18-pg-when_0.1.10-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-18-pg-when_0.1.10-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_17` | `0.1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pg_when_17-0.1.10-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_17-0.1.10-1PGSTY.el8.x86_64.rpm) |
| `pg_when_17` | `0.1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 970.8 KiB | [pg_when_17-0.1.10-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_17-0.1.10-1PGSTY.el8.aarch64.rpm) |
| `pg_when_17` | `0.1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_when_17-0.1.10-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_17-0.1.10-1PGSTY.el9.x86_64.rpm) |
| `pg_when_17` | `0.1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_when_17-0.1.10-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_17-0.1.10-1PGSTY.el9.aarch64.rpm) |
| `pg_when_17` | `0.1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_when_17-0.1.10-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_17-0.1.10-1PGSTY.el10.x86_64.rpm) |
| `pg_when_17` | `0.1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1015.4 KiB | [pg_when_17-0.1.10-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_17-0.1.10-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-when` | `0.1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 881.4 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 753.8 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 881.4 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 754.7 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 976.1 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 884.1 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 967.1 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 874.8 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [u26.x86_64](/os/u26.x86_64) | pigsty | 962.8 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-when` | `0.1.10` | [u26.aarch64](/os/u26.aarch64) | pigsty | 873.7 KiB | [postgresql-17-pg-when_0.1.10-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-17-pg-when_0.1.10-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_16` | `0.1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pg_when_16-0.1.10-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_16-0.1.10-1PGSTY.el8.x86_64.rpm) |
| `pg_when_16` | `0.1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 969.4 KiB | [pg_when_16-0.1.10-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_16-0.1.10-1PGSTY.el8.aarch64.rpm) |
| `pg_when_16` | `0.1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pg_when_16-0.1.10-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_16-0.1.10-1PGSTY.el9.x86_64.rpm) |
| `pg_when_16` | `0.1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pg_when_16-0.1.10-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_16-0.1.10-1PGSTY.el9.aarch64.rpm) |
| `pg_when_16` | `0.1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pg_when_16-0.1.10-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_16-0.1.10-1PGSTY.el10.x86_64.rpm) |
| `pg_when_16` | `0.1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1015.0 KiB | [pg_when_16-0.1.10-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_16-0.1.10-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-when` | `0.1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 880.5 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 753.1 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 880.7 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 754.5 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 974.4 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 884.4 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 965.9 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 875.1 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [u26.x86_64](/os/u26.x86_64) | pigsty | 961.5 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-when` | `0.1.10` | [u26.aarch64](/os/u26.aarch64) | pigsty | 872.2 KiB | [postgresql-16-pg-when_0.1.10-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-16-pg-when_0.1.10-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_15` | `0.1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pg_when_15-0.1.10-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_15-0.1.10-1PGSTY.el8.x86_64.rpm) |
| `pg_when_15` | `0.1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 960.2 KiB | [pg_when_15-0.1.10-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_15-0.1.10-1PGSTY.el8.aarch64.rpm) |
| `pg_when_15` | `0.1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.0 MiB | [pg_when_15-0.1.10-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_15-0.1.10-1PGSTY.el9.x86_64.rpm) |
| `pg_when_15` | `0.1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1022.5 KiB | [pg_when_15-0.1.10-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_15-0.1.10-1PGSTY.el9.aarch64.rpm) |
| `pg_when_15` | `0.1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_when_15-0.1.10-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_15-0.1.10-1PGSTY.el10.x86_64.rpm) |
| `pg_when_15` | `0.1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1012.2 KiB | [pg_when_15-0.1.10-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_15-0.1.10-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-when` | `0.1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 875.0 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 748.7 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 874.7 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 749.3 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 970.2 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 877.2 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 959.9 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 868.4 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [u26.x86_64](/os/u26.x86_64) | pigsty | 953.1 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-when` | `0.1.10` | [u26.aarch64](/os/u26.aarch64) | pigsty | 866.1 KiB | [postgresql-15-pg-when_0.1.10-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-15-pg-when_0.1.10-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_when_14` | `0.1.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pg_when_14-0.1.10-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_when_14-0.1.10-1PGSTY.el8.x86_64.rpm) |
| `pg_when_14` | `0.1.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 957.8 KiB | [pg_when_14-0.1.10-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_when_14-0.1.10-1PGSTY.el8.aarch64.rpm) |
| `pg_when_14` | `0.1.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.0 MiB | [pg_when_14-0.1.10-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_when_14-0.1.10-1PGSTY.el9.x86_64.rpm) |
| `pg_when_14` | `0.1.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1020.8 KiB | [pg_when_14-0.1.10-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_when_14-0.1.10-1PGSTY.el9.aarch64.rpm) |
| `pg_when_14` | `0.1.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.0 MiB | [pg_when_14-0.1.10-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_when_14-0.1.10-1PGSTY.el10.x86_64.rpm) |
| `pg_when_14` | `0.1.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1010.2 KiB | [pg_when_14-0.1.10-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_when_14-0.1.10-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-when` | `0.1.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 872.2 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 746.9 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 872.3 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 747.9 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 965.2 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 875.4 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 955.5 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 866.2 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [u26.x86_64](/os/u26.x86_64) | pigsty | 952.0 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-when` | `0.1.10` | [u26.aarch64](/os/u26.aarch64) | pigsty | 863.7 KiB | [postgresql-14-pg-when_0.1.10-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-when/postgresql-14-pg-when_0.1.10-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/frectonz/pg-when" title="Repository" icon="github" subtitle="github.com/frectonz/pg-when" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_when-0.1.10.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_when;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_when;		# install via package name, for the active PG version

pig install pg_when -v 18;   # install for PG 18
pig install pg_when -v 17;   # install for PG 17
pig install pg_when -v 16;   # install for PG 16
pig install pg_when -v 15;   # install for PG 15
pig install pg_when -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_when;
```

## Usage

Sources:

- [pg_when 0.1.10 on PGXN](https://pgxn.org/dist/pg_when/0.1.10/)
- [pg_when 0.1.10 README](https://github.com/frectonz/pg-when/blob/0.1.10/README.md)
- [pg_when 0.1.10 Cargo manifest](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/Cargo.toml)
- [pg_when 0.1.10 control file](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/pg_when.control)
- [pg_when 0.1.10 exported functions](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/src/when_is.rs)
- [pg_when 0.1.10 relative-date implementation](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/src/when_relative_date.rs)

`pg_when` 0.1.10 parses a constrained natural-language date and time expression and returns either a PostgreSQL `timestamptz` value or a Unix epoch value at a selected precision.

```sql
CREATE EXTENSION pg_when;

SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT seconds_at('5 days ago at this hour in Asia/Tokyo');
SELECT millis_at('in 2 months at midnight in UTC-8');
SELECT micros_at('December 31, 2026 at evening');
SELECT nanos_at('last monday at 22:30');
```

### Query Shape

A query can contain a date, a time, and a time zone, joined by `at` and `in`:

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<date> in <timezone>');
SELECT when_is('<time>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

If the time zone is omitted, the parser uses UTC. Supported inputs include relative dates such as `tomorrow`, `last month`, and `5 days ago`; exact dates in common numeric and month-name forms; relative times such as `noon`, `midnight`, and `next hour`; clock times; IANA time-zone names; and UTC offsets.

### Function Index

- `when_is(text)` returns `timestamptz`.
- `seconds_at(text)` returns Unix epoch seconds.
- `millis_at(text)` returns Unix epoch milliseconds.
- `micros_at(text)` returns Unix epoch microseconds.
- `nanos_at(text)` returns Unix epoch nanoseconds.

### Compatibility and Boundaries

- The parser implements the documented grammar; it is not a general-purpose natural-language interpreter.
- Upstream 0.1.10 declares PostgreSQL 13–18 features and pins pgrx 0.18.1. Pigsty packages cover PostgreSQL 14–18 and apply a locked pgrx 0.19.1 compatibility update.
- `pg_when` is not relocatable and its control file requires a superuser for `CREATE EXTENSION`.
- Invalid text raises an error. All five functions are `STRICT`, so a null input returns null; `nanos_at(text)` also errors when the epoch nanoseconds cannot fit in `bigint`.
- The 0.1.10 SQL functions are declared `IMMUTABLE`, but relative expressions such as `now`, `tomorrow`, and `5 days ago` read the wall clock. Do not use relative-input calls in expression indexes or generated columns, and do not rely on them being reevaluated in cached plans; only fully specified date, time, and time-zone inputs are time-independent.
