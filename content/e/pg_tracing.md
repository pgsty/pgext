---
title: "pg_tracing"
linkTitle: "pg_tracing"
description: "Distributed Tracing for PostgreSQL"
weight: 6010
categories: ["STAT"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_tracing**](https://github.com/DataDog/pg_tracing) : Distributed Tracing for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6010** | {{< badge content="pg_tracing" link="https://github.com/DataDog/pg_tracing" >}} | {{< ext "pg_tracing" >}} | `0.1.3` | {{< category "STAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_stat_log" >}} {{< ext "pg_stat_backtrace" >}} {{< ext "pgfr_record" >}} {{< ext "logerrors" >}} {{< ext "pgelog" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.3` | {{< bg "19" "" "red" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_tracing` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.3` | {{< bg "19" "pg_tracing_19" "red" >}} {{< bg "18" "pg_tracing_18" "green" >}} {{< bg "17" "pg_tracing_17" "green" >}} {{< bg "16" "pg_tracing_16" "green" >}} {{< bg "15" "pg_tracing_15" "green" >}} | `pg_tracing_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.3` | {{< bg "19" "postgresql-19-pg-tracing" "red" >}} {{< bg "18" "postgresql-18-pg-tracing" "green" >}} {{< bg "17" "postgresql-17-pg-tracing" "green" >}} {{< bg "16" "postgresql-16-pg-tracing" "green" >}} {{< bg "15" "postgresql-15-pg-tracing" "green" >}} | `postgresql-$v-pg-tracing` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_tracing_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_tracing_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "pg_tracing_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "pg_tracing_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "pg_tracing_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "pg_tracing_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "pg_tracing_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-tracing : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-18-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-17-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-16-pg-tracing : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.3" "postgresql-15-pg-tracing : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_18` | `0.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 117.1 KiB | [pg_tracing_18-0.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_18-0.1.3-1PGSTY.el8.x86_64.rpm) |
| `pg_tracing_18` | `0.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 114.5 KiB | [pg_tracing_18-0.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_18-0.1.3-1PGSTY.el8.aarch64.rpm) |
| `pg_tracing_18` | `0.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 118.4 KiB | [pg_tracing_18-0.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_18-0.1.3-1PGSTY.el9.x86_64.rpm) |
| `pg_tracing_18` | `0.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 117.0 KiB | [pg_tracing_18-0.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_18-0.1.3-1PGSTY.el9.aarch64.rpm) |
| `pg_tracing_18` | `0.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 119.2 KiB | [pg_tracing_18-0.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_18-0.1.3-1PGSTY.el10.x86_64.rpm) |
| `pg_tracing_18` | `0.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 117.4 KiB | [pg_tracing_18-0.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_18-0.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tracing` | `0.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 105.5 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 102.9 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 105.5 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 103.5 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 114.7 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 113.8 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 110.2 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 109.2 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 109.6 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-tracing` | `0.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 108.1 KiB | [postgresql-18-pg-tracing_0.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tracing/postgresql-18-pg-tracing_0.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_17` | `0.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 117.2 KiB | [pg_tracing_17-0.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_17-0.1.3-1PGSTY.el8.x86_64.rpm) |
| `pg_tracing_17` | `0.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 114.5 KiB | [pg_tracing_17-0.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_17-0.1.3-1PGSTY.el8.aarch64.rpm) |
| `pg_tracing_17` | `0.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 118.3 KiB | [pg_tracing_17-0.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_17-0.1.3-1PGSTY.el9.x86_64.rpm) |
| `pg_tracing_17` | `0.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 117.0 KiB | [pg_tracing_17-0.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_17-0.1.3-1PGSTY.el9.aarch64.rpm) |
| `pg_tracing_17` | `0.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 119.3 KiB | [pg_tracing_17-0.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_17-0.1.3-1PGSTY.el10.x86_64.rpm) |
| `pg_tracing_17` | `0.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 117.4 KiB | [pg_tracing_17-0.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_17-0.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tracing` | `0.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 105.4 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 103.1 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 105.6 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 103.8 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 129.4 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 128.2 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 110.3 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 109.3 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 109.7 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-tracing` | `0.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 108.1 KiB | [postgresql-17-pg-tracing_0.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tracing/postgresql-17-pg-tracing_0.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_16` | `0.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 117.2 KiB | [pg_tracing_16-0.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_16-0.1.3-1PGSTY.el8.x86_64.rpm) |
| `pg_tracing_16` | `0.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 114.4 KiB | [pg_tracing_16-0.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_16-0.1.3-1PGSTY.el8.aarch64.rpm) |
| `pg_tracing_16` | `0.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 118.3 KiB | [pg_tracing_16-0.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_16-0.1.3-1PGSTY.el9.x86_64.rpm) |
| `pg_tracing_16` | `0.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 116.9 KiB | [pg_tracing_16-0.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_16-0.1.3-1PGSTY.el9.aarch64.rpm) |
| `pg_tracing_16` | `0.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 119.3 KiB | [pg_tracing_16-0.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_16-0.1.3-1PGSTY.el10.x86_64.rpm) |
| `pg_tracing_16` | `0.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 117.3 KiB | [pg_tracing_16-0.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_16-0.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-tracing` | `0.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 105.5 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 103.0 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 105.7 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 103.6 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 128.8 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 127.7 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 110.3 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 109.2 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 109.7 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-tracing` | `0.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 108.0 KiB | [postgresql-16-pg-tracing_0.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tracing/postgresql-16-pg-tracing_0.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tracing_15` | `0.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 119.7 KiB | [pg_tracing_15-0.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tracing_15-0.1.3-1PGSTY.el8.x86_64.rpm) |
| `pg_tracing_15` | `0.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 116.7 KiB | [pg_tracing_15-0.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tracing_15-0.1.3-1PGSTY.el8.aarch64.rpm) |
| `pg_tracing_15` | `0.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 121.3 KiB | [pg_tracing_15-0.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tracing_15-0.1.3-1PGSTY.el9.x86_64.rpm) |
| `pg_tracing_15` | `0.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 120.1 KiB | [pg_tracing_15-0.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tracing_15-0.1.3-1PGSTY.el9.aarch64.rpm) |
| `pg_tracing_15` | `0.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 122.3 KiB | [pg_tracing_15-0.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tracing_15-0.1.3-1PGSTY.el10.x86_64.rpm) |
| `pg_tracing_15` | `0.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 120.8 KiB | [pg_tracing_15-0.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tracing_15-0.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-tracing` | `0.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 108.1 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.4 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 108.2 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 106.0 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 132.1 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 130.9 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.3 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 112.5 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 112.2 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-tracing` | `0.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 111.1 KiB | [postgresql-15-pg-tracing_0.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tracing/postgresql-15-pg-tracing_0.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/DataDog/pg_tracing" title="Repository" icon="github" subtitle="github.com/DataDog/pg_tracing" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tracing-0.1.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_tracing;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_tracing;		# install via package name, for the active PG version

pig install pg_tracing -v 18;   # install for PG 18
pig install pg_tracing -v 17;   # install for PG 17
pig install pg_tracing -v 16;   # install for PG 16
pig install pg_tracing -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_tracing';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_tracing;
```




## Usage

> [pg_tracing: distributed tracing for PostgreSQL](https://github.com/DataDog/pg_tracing)

pg_tracing generates server-side spans for distributed tracing. It creates spans for planner, executor, plan nodes, nested queries, triggers, parallel workers, and transaction commits.

### Trace Propagation via SQLCommenter

Pass trace context as a SQL comment. Queries with a sampled flag will generate spans:

```sql
/*traceparent='00-00000000000000000000000000000123-0000000000000123-01'*/ SELECT 1;
```

### Trace Propagation via GUC

```sql
BEGIN;
SET LOCAL pg_tracing.trace_context='traceparent=''00-00000000000000000000000000000005-0000000000000005-01''';
UPDATE pgbench_accounts SET abalance=1 WHERE aid=1;
COMMIT;
```

### Standalone Sampling

Randomly sample queries without external trace context:

```sql
SET pg_tracing.sample_rate = 1.0;  -- trace all queries
SELECT 1;
```

### Viewing Spans

```sql
-- Consume spans (removes them from buffer)
SELECT trace_id, parent_id, span_id, span_start, span_end, span_type, span_operation
FROM pg_tracing_consume_spans ORDER BY span_start;

-- Peek at spans (non-destructive)
SELECT * FROM pg_tracing_peek_spans;

-- Export as OTLP JSON
SELECT pg_tracing_json_spans();
```

### Statistics

```sql
SELECT * FROM pg_tracing_info;
SELECT pg_tracing_reset();
```

### Sending Spans to OpenTelemetry Collector

Configure in `postgresql.conf`:

```ini
pg_tracing.otel_endpoint = http://127.0.0.1:4318/v1/traces
pg_tracing.otel_naptime = 2000
```

### Key GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pg_tracing.max_span` | 10000 | Maximum spans in shared memory |
| `pg_tracing.track` | `all` | Which statements to track |
| `pg_tracing.sample_rate` | 0 | Fraction of queries to sample randomly |
| `pg_tracing.otel_endpoint` | (empty) | OTLP HTTP endpoint for span export |
