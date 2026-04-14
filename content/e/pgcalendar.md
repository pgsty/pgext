---
title: "pgcalendar"
linkTitle: "pgcalendar"
description: "Recurring calendar, schedule, and exception management for PostgreSQL"
weight: 3890
categories: ["TYPE"]
width: full
---

[**pgcalendar**](https://github.com/h4kbas/pgcalendar) : Recurring calendar, schedule, and exception management for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3890** | {{< badge content="pgcalendar" link="https://github.com/h4kbas/pgcalendar" >}} | {{< ext "pgcalendar" >}} | `1.1.0` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgcalendar` |
|   **See Also**    | {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "timeseries" >}} {{< ext "pg_cron" >}} |

> [!Note] Deb/RPM recipes patch the stale upstream 1.1.0 control metadata (default_version/module_pathname).


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgcalendar` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "pgcalendar_18" "green" >}} {{< bg "17" "pgcalendar_17" "green" >}} {{< bg "16" "pgcalendar_16" "green" >}} {{< bg "15" "pgcalendar_15" "green" >}} {{< bg "14" "pgcalendar_14" "green" >}} | `pgcalendar_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-pgcalendar" "green" >}} {{< bg "17" "postgresql-17-pgcalendar" "green" >}} {{< bg "16" "postgresql-16-pgcalendar" "green" >}} {{< bg "15" "postgresql-15-pgcalendar" "green" >}} {{< bg "14" "postgresql-14-pgcalendar" "green" >}} | `postgresql-$v-pgcalendar` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcalendar_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pgcalendar : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pgcalendar : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pgcalendar : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pgcalendar : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pgcalendar : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pgcalendar : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pgcalendar : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pgcalendar : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pgcalendar : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcalendar_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pgcalendar_18-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcalendar_18-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcalendar_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.5 KiB | [pgcalendar_18-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcalendar_18-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcalendar_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pgcalendar_18-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcalendar_18-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcalendar_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pgcalendar_18-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcalendar_18-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcalendar_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pgcalendar_18-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcalendar_18-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcalendar_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pgcalendar_18-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcalendar_18-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgcalendar` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.3 KiB | [postgresql-18-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-18-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgcalendar` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 KiB | [postgresql-18-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-18-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgcalendar` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.3 KiB | [postgresql-18-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-18-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgcalendar` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.3 KiB | [postgresql-18-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-18-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgcalendar` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-18-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgcalendar` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-18-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-18-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgcalendar` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-18-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-18-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgcalendar` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-18-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-18-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcalendar_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pgcalendar_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcalendar_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcalendar_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.5 KiB | [pgcalendar_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcalendar_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcalendar_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pgcalendar_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcalendar_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcalendar_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pgcalendar_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcalendar_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcalendar_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pgcalendar_17-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcalendar_17-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcalendar_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pgcalendar_17-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcalendar_17-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgcalendar` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.3 KiB | [postgresql-17-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-17-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgcalendar` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 KiB | [postgresql-17-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-17-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgcalendar` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.3 KiB | [postgresql-17-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-17-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgcalendar` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.3 KiB | [postgresql-17-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-17-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgcalendar` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-17-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgcalendar` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-17-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-17-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgcalendar` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-17-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-17-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgcalendar` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-17-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-17-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcalendar_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pgcalendar_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcalendar_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcalendar_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.5 KiB | [pgcalendar_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcalendar_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcalendar_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pgcalendar_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcalendar_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcalendar_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pgcalendar_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcalendar_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcalendar_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pgcalendar_16-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcalendar_16-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcalendar_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pgcalendar_16-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcalendar_16-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgcalendar` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.3 KiB | [postgresql-16-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-16-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgcalendar` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 KiB | [postgresql-16-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-16-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgcalendar` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.3 KiB | [postgresql-16-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-16-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgcalendar` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.3 KiB | [postgresql-16-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-16-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgcalendar` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-16-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgcalendar` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-16-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-16-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgcalendar` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-16-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-16-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgcalendar` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-16-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-16-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcalendar_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pgcalendar_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcalendar_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcalendar_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.5 KiB | [pgcalendar_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcalendar_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcalendar_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pgcalendar_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcalendar_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcalendar_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pgcalendar_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcalendar_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcalendar_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pgcalendar_15-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcalendar_15-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcalendar_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pgcalendar_15-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcalendar_15-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgcalendar` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.3 KiB | [postgresql-15-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-15-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgcalendar` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 KiB | [postgresql-15-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-15-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgcalendar` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.3 KiB | [postgresql-15-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-15-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgcalendar` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.3 KiB | [postgresql-15-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-15-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgcalendar` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-15-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgcalendar` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-15-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-15-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgcalendar` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-15-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-15-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgcalendar` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-15-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-15-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcalendar_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [pgcalendar_14-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcalendar_14-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcalendar_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.5 KiB | [pgcalendar_14-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcalendar_14-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcalendar_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.5 KiB | [pgcalendar_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcalendar_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcalendar_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [pgcalendar_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcalendar_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcalendar_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [pgcalendar_14-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcalendar_14-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcalendar_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [pgcalendar_14-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcalendar_14-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgcalendar` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 7.3 KiB | [postgresql-14-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-14-pgcalendar_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgcalendar` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 KiB | [postgresql-14-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcalendar/postgresql-14-pgcalendar_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgcalendar` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 7.3 KiB | [postgresql-14-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-14-pgcalendar_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgcalendar` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.3 KiB | [postgresql-14-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcalendar/postgresql-14-pgcalendar_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgcalendar` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 7.4 KiB | [postgresql-14-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-14-pgcalendar_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgcalendar` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 7.4 KiB | [postgresql-14-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcalendar/postgresql-14-pgcalendar_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgcalendar` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 7.4 KiB | [postgresql-14-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-14-pgcalendar_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgcalendar` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 7.4 KiB | [postgresql-14-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcalendar/postgresql-14-pgcalendar_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/h4kbas/pgcalendar" title="Repository" icon="github" subtitle="github.com/h4kbas/pgcalendar" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcalendar-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgcalendar;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgcalendar;		# install via package name, for the active PG version

pig install pgcalendar -v 18;   # install for PG 18
pig install pgcalendar -v 17;   # install for PG 17
pig install pgcalendar -v 16;   # install for PG 16
pig install pgcalendar -v 15;   # install for PG 15
pig install pgcalendar -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgcalendar;
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pgcalendar;
> INSERT INTO pgcalendar.events (name, description, category)
> VALUES ('Daily Standup', 'Team daily standup meeting', 'meeting');
> SELECT * FROM pgcalendar.get_event_projections(1, '2024-01-01'::date, '2024-01-07'::date);
> ```
>
> Source: [README](https://github.com/h4kbas/pgcalendar)

`pgcalendar` is a recurring calendar extension for PostgreSQL. It models events, schedules, exceptions, and projections, and generates calendar occurrences across arbitrary date ranges.

## Data Model

The README describes four main concepts:

- `events` as logical objects such as meetings or tasks
- `schedules` as non-overlapping recurrence definitions
- `exceptions` as per-occurrence cancellations or modifications
- `projections` as the actual generated calendar occurrences

## Quick Start

Create an event:

```sql
INSERT INTO pgcalendar.events (name, description, category)
VALUES ('Daily Standup', 'Team daily standup meeting', 'meeting');
```

Create a schedule:

```sql
INSERT INTO pgcalendar.schedules (
    event_id, start_date, end_date, recurrence_type, recurrence_interval
) VALUES (
    1, '2024-01-01 09:00:00', '2024-01-07 23:59:59', 'daily', 1
);
```

Get projections:

```sql
SELECT * FROM pgcalendar.get_event_projections(
    1, '2024-01-01'::date, '2024-01-07'::date
);
```

## Recurrence Types

The README shows schedule examples for:

- daily recurrence
- weekly recurrence with `recurrence_day_of_week`
- monthly recurrence with `recurrence_day_of_month`
- yearly recurrence with `recurrence_month` and `recurrence_day_of_month`

## Exceptions

Exceptions can cancel or modify a single occurrence:

```sql
INSERT INTO pgcalendar.exceptions (
    schedule_id, exception_date, exception_type, notes
) VALUES (
    1, '2024-01-15', 'cancelled', 'Holiday - meeting cancelled'
);
```

Modified occurrences can also change date and time.

## Functions and Views

The README documents:

- `get_event_projections(event_id, start_date, end_date)`
- `get_events_detailed(start_date, end_date)`
- `transition_event_schedule(...)`
- `check_schedule_overlap(event_id, start_date, end_date)`
- `pgcalendar.event_calendar`

`transition_event_schedule(...)` safely switches an event to a new schedule definition, while `check_schedule_overlap(...)` validates that new schedules do not overlap existing ones.
