---
title: "pg_bikram_sambat"
linkTitle: "pg_bikram_sambat"
description: "Bikram Sambat date type and AD/BS conversion functions"
weight: 3860
categories: ["TYPE"]
width: full
---

[**pg_bikram_sambat**](https://github.com/LeohangRai/pg_bikram_sambat) : Bikram Sambat date type and AD/BS conversion functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3860** | {{< badge content="pg_bikram_sambat" link="https://github.com/LeohangRai/pg_bikram_sambat" >}} | {{< ext "pg_bikram_sambat" >}} | `0.1.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_duration" >}} {{< ext "pg_rrule" >}} {{< ext "pgcalendar" >}} {{< ext "timestamp9" >}} {{< ext "pg_extra_time" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "country" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_bikram_sambat` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_bikram_sambat_18" "green" >}} {{< bg "17" "pg_bikram_sambat_17" "green" >}} {{< bg "16" "pg_bikram_sambat_16" "green" >}} {{< bg "15" "pg_bikram_sambat_15" "green" >}} {{< bg "14" "pg_bikram_sambat_14" "green" >}} | `pg_bikram_sambat_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-bikram-sambat" "green" >}} {{< bg "17" "postgresql-17-pg-bikram-sambat" "green" >}} {{< bg "16" "postgresql-16-pg-bikram-sambat" "green" >}} {{< bg "15" "postgresql-15-pg-bikram-sambat" "green" >}} {{< bg "14" "postgresql-14-pg-bikram-sambat" "green" >}} | `postgresql-$v-pg-bikram-sambat` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_bikram_sambat_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-bikram-sambat : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pg-bikram-sambat : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bikram_sambat_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.1 KiB | [pg_bikram_sambat_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bikram_sambat_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bikram_sambat_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.6 KiB | [pg_bikram_sambat_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bikram_sambat_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bikram_sambat_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.0 KiB | [pg_bikram_sambat_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bikram_sambat_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bikram_sambat_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.9 KiB | [pg_bikram_sambat_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bikram_sambat_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bikram_sambat_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.2 KiB | [pg_bikram_sambat_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bikram_sambat_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bikram_sambat_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.1 KiB | [pg_bikram_sambat_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bikram_sambat_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.4 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.1 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.3 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.3 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 73.4 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 73.2 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.5 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.5 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.6 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-bikram-sambat` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.6 KiB | [postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-18-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bikram_sambat_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.0 KiB | [pg_bikram_sambat_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bikram_sambat_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bikram_sambat_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.6 KiB | [pg_bikram_sambat_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bikram_sambat_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bikram_sambat_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.0 KiB | [pg_bikram_sambat_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bikram_sambat_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bikram_sambat_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.9 KiB | [pg_bikram_sambat_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bikram_sambat_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bikram_sambat_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.2 KiB | [pg_bikram_sambat_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bikram_sambat_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bikram_sambat_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.1 KiB | [pg_bikram_sambat_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bikram_sambat_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.4 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.1 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.5 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.2 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 74.5 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 74.4 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.0 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.8 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 71.0 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-bikram-sambat` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.1 KiB | [postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-17-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bikram_sambat_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.0 KiB | [pg_bikram_sambat_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bikram_sambat_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bikram_sambat_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.6 KiB | [pg_bikram_sambat_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bikram_sambat_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bikram_sambat_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.0 KiB | [pg_bikram_sambat_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bikram_sambat_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bikram_sambat_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.8 KiB | [pg_bikram_sambat_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bikram_sambat_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bikram_sambat_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.2 KiB | [pg_bikram_sambat_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bikram_sambat_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bikram_sambat_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.1 KiB | [pg_bikram_sambat_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bikram_sambat_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.1 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 66.9 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.2 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.0 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 75.6 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 75.5 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 71.7 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.5 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 73.0 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-bikram-sambat` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 72.9 KiB | [postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-16-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bikram_sambat_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.7 KiB | [pg_bikram_sambat_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bikram_sambat_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bikram_sambat_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.1 KiB | [pg_bikram_sambat_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bikram_sambat_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bikram_sambat_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.3 KiB | [pg_bikram_sambat_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bikram_sambat_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bikram_sambat_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.3 KiB | [pg_bikram_sambat_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bikram_sambat_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bikram_sambat_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.5 KiB | [pg_bikram_sambat_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bikram_sambat_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bikram_sambat_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.3 KiB | [pg_bikram_sambat_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bikram_sambat_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 67.5 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.1 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 67.7 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.2 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 68.7 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 68.7 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 72.0 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.9 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 70.7 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-bikram-sambat` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 70.8 KiB | [postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-15-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bikram_sambat_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.7 KiB | [pg_bikram_sambat_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bikram_sambat_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_bikram_sambat_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.1 KiB | [pg_bikram_sambat_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bikram_sambat_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_bikram_sambat_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.3 KiB | [pg_bikram_sambat_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bikram_sambat_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_bikram_sambat_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.2 KiB | [pg_bikram_sambat_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bikram_sambat_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_bikram_sambat_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.5 KiB | [pg_bikram_sambat_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bikram_sambat_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_bikram_sambat_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.3 KiB | [pg_bikram_sambat_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bikram_sambat_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 65.8 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 65.5 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 66.0 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 65.5 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 70.7 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 70.7 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 70.0 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 70.0 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 69.3 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-bikram-sambat` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 69.3 KiB | [postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bikram-sambat/postgresql-14-pg-bikram-sambat_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/LeohangRai/pg_bikram_sambat" title="Repository" icon="github" subtitle="github.com/LeohangRai/pg_bikram_sambat" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bikram_sambat-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_bikram_sambat;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_bikram_sambat;		# install via package name, for the active PG version

pig install pg_bikram_sambat -v 18;   # install for PG 18
pig install pg_bikram_sambat -v 17;   # install for PG 17
pig install pg_bikram_sambat -v 16;   # install for PG 16
pig install pg_bikram_sambat -v 15;   # install for PG 15
pig install pg_bikram_sambat -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_bikram_sambat;
```

## Usage

Sources: [PGXN metadata](https://api.pgxn.org/dist/pg_bikram_sambat.json), [PGXN source tree](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/), [type SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/types.sql), [function SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/functions.sql), [operator SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/operators.sql), [cast SQL](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/sql/casts.sql), [regression examples](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/tests/pg_regress/sql/003_functions.sql), [TODO](https://api.pgxn.org/src/pg_bikram_sambat/pg_bikram_sambat-0.1.0/todos.txt)

`pg_bikram_sambat` adds a `bs_date` type for Bikram Sambat dates plus conversion, formatting, comparison, and btree indexing support. Install it as a normal PostgreSQL extension:

```sql
CREATE EXTENSION pg_bikram_sambat;
```

### Date Type

`bs_date` stores a Bikram Sambat date and displays it as `YYYY-MM-DD`. Text input accepts year/month/day values separated with `/`, `-`, or `.`; the input parser also accepts day-first strings when the year appears in the last field.

```sql
SELECT '2057/10/19'::bs_date;
SELECT CAST('2057-10-19' AS bs_date);
SELECT '19.10.2057'::bs_date;
```

Use it in tables like any other PostgreSQL type:

```sql
CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  ad_date date,
  bs bs_date NOT NULL
);

INSERT INTO events (ad_date, bs)
VALUES
  ('2001-02-01', '2057/10/19'),
  ('1972-02-17', '2028/11/05');
```

### Conversion Functions

`ad_to_bs(date)` converts a Gregorian `date` to `bs_date`:

```sql
SELECT ad_to_bs('2001-02-01'::date);  -- 2057-10-19
SELECT ad_to_bs('1972-02-17'::date);  -- 2028-11-05
```

`current_bs_date()` returns the current transaction timestamp converted to `bs_date`, so repeated calls inside the same transaction are stable:

```sql
SELECT current_bs_date();
SELECT pg_typeof(current_bs_date());  -- bs_date
```

Version `0.1.0` does not expose a SQL `bs_to_ad()` function or direct `bs_date` to `date` cast; the upstream TODO file lists those as future work.

### Formatting

The extension overloads PostgreSQL `to_char` for `bs_date`:

```sql
SELECT to_char('2057/10/19'::bs_date, 'YYYY-MM-DD');
SELECT to_char('2057/10/19'::bs_date, 'DD/MM/YYYY');
SELECT to_char('2057/10/19'::bs_date, 'Month DD, YYYY');
SELECT to_char('2057/10/19'::bs_date, 'Day, DD Month YYYY');
```

Supported date-format tokens are `YYYY`, `YY`, `Month`, `Mon`, `MM`, `Day`, `Dy`, and `DD`. Month and weekday names follow the casing of the format token, so `MONTH`, `Month`, and `month` produce upper-case, title-case, and lower-case English names.

Pass `dev` as the third argument for Devanagari digits, month names, and weekday names:

```sql
SELECT to_char('2057/10/19'::bs_date, 'YYYY-MM-DD', 'dev');
SELECT to_char('2057/10/19'::bs_date, 'Day, DD Month YYYY', 'dev');
```

### Operators And Indexes

`bs_date` supports the comparison operators `=`, `<>`, `>`, `>=`, `<`, and `<=`. The default btree operator class `bs_date_ops` enables ordinary btree indexes, range predicates, and ordering:

```sql
CREATE INDEX events_bs_idx ON events (bs);

SELECT * FROM events WHERE bs >= '2057/01/01' ORDER BY bs;
SELECT * FROM events WHERE bs BETWEEN '2056/01/01' AND '2058/12/12';
```

### Caveats

The packaged conversion dataset covers BS years `2000` through `2100`, with `1943-04-14` AD as the reference date for `2000-01-01` BS. Dates before the reference date or beyond the mapped BS range raise PostgreSQL errors. The extension defines an implicit cast from `text` to `bs_date`, but it does not define casts from arbitrary numeric types.
