---
title: "pg_stl"
linkTitle: "pg_stl"
description: "Time series analysis functions for PostgreSQL"
weight: 1130
categories: ["TIME"]
width: full
---

[**pg_stl**](https://github.com/nadyaloseva/pg_ts_analysis) : Time series analysis functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1130** | {{< badge content="pg_stl" link="https://github.com/nadyaloseva/pg_ts_analysis" >}} | {{< ext "pg_stl" >}} | `1.0.0` | {{< category "TIME" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "timescaledb" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} |

> [!Note] ACF, PACF, STL decomposition, and Holt-Winters forecasting.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_stl` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_stl_18" "green" >}} {{< bg "17" "pg_stl_17" "green" >}} {{< bg "16" "pg_stl_16" "green" >}} {{< bg "15" "pg_stl_15" "red" >}} {{< bg "14" "pg_stl_14" "red" >}} | `pg_stl_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-stl" "green" >}} {{< bg "17" "postgresql-17-pg-stl" "green" >}} {{< bg "16" "postgresql-16-pg-stl" "green" >}} {{< bg "15" "postgresql-15-pg-stl" "red" >}} {{< bg "14" "postgresql-14-pg-stl" "red" >}} | `postgresql-$v-pg-stl` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stl_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stl_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stl_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stl_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stl_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stl_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stl_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stl_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stl_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stl_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_stl_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_stl_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_stl_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-stl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-stl : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-pg-stl : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-stl : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stl_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [pg_stl_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stl_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_stl_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_stl_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stl_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_stl_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pg_stl_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stl_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_stl_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_stl_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stl_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_stl_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.7 KiB | [pg_stl_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stl_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_stl_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_stl_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stl_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-stl` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.8 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.8 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.3 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.9 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.5 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.7 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.7 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-stl` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.4 KiB | [postgresql-18-pg-stl_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stl/postgresql-18-pg-stl_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stl_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [pg_stl_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stl_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_stl_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_stl_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stl_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_stl_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pg_stl_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stl_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_stl_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_stl_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stl_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_stl_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.7 KiB | [pg_stl_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stl_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_stl_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_stl_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stl_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-stl` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.7 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.8 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.3 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.0 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.6 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.7 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.7 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-stl` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.4 KiB | [postgresql-17-pg-stl_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stl/postgresql-17-pg-stl_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_stl_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [pg_stl_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_stl_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_stl_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 18.9 KiB | [pg_stl_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_stl_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_stl_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pg_stl_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_stl_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_stl_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_stl_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_stl_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_stl_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.7 KiB | [pg_stl_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_stl_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_stl_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.3 KiB | [pg_stl_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_stl_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-stl` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.7 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.4 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.8 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.3 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.0 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.6 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.7 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.5 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 25.7 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-stl` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 25.4 KiB | [postgresql-16-pg-stl_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-stl/postgresql-16-pg-stl_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/nadyaloseva/pg_ts_analysis" title="Repository" icon="github" subtitle="github.com/nadyaloseva/pg_ts_analysis" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_stl-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_stl;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_stl;		# install via package name, for the active PG version

pig install pg_stl -v 18;   # install for PG 18
pig install pg_stl -v 17;   # install for PG 17
pig install pg_stl -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_stl;
```

## Usage

Sources: [pg_ts_analysis README](https://github.com/nadyaloseva/pg_ts_analysis), [SQL definitions](https://github.com/nadyaloseva/pg_ts_analysis/blob/main/pg_stl--1.0.sql), [control file](https://github.com/nadyaloseva/pg_ts_analysis/blob/main/pg_stl.control).

`pg_stl` provides time-series analysis functions for PostgreSQL: autocorrelation, partial autocorrelation, STL decomposition, and Holt-Winters forecasting. The upstream README and SQL definitions target PostgreSQL 16+.

### Autocorrelation

`acf_array(data double precision[], lags integer)` returns autocorrelation values for lags `1..lags`:

```sql
CREATE EXTENSION pg_stl;

SELECT acf_array(
  array_agg(revenue ORDER BY date)::double precision[],
  28
)
FROM daily_sales;
```

The README describes using peaks at lags such as `7`, `14`, and `21` as a signal for weekly seasonality. The function returns `NULL` when the series is too short, `lags < 1`, or `lags >= n`.

### Partial Autocorrelation

`pacf_array(data double precision[], lags integer)` returns partial autocorrelation values using the Durbin-Levinson recursion:

```sql
WITH series AS (
  SELECT array_agg(value ORDER BY ts)::double precision[] AS values
  FROM measurements
)
SELECT
  unnest(acf_array(values, 20)) AS acf,
  unnest(pacf_array(values, 20)) AS pacf
FROM series;
```

Use PACF when you want the direct lag relationship after accounting for shorter lags.

### STL Decomposition

`stl_decompose` decomposes a series into trend, seasonal, and residual arrays:

```sql
WITH data AS (
  SELECT array_agg(revenue ORDER BY month)::double precision[] AS values
  FROM monthly_revenue
),
decomposed AS (
  SELECT (stl_decompose(values, 12)).*
  FROM data
)
SELECT
  unnest(trend) AS trend,
  unnest(seasonal) AS seasonal,
  unnest(residual) AS residual
FROM decomposed;
```

Signature from the SQL definition:

```sql
stl_decompose(
  y double precision[],
  period integer,
  seasonal integer DEFAULT 7,
  robust boolean DEFAULT true,
  trend integer DEFAULT 0,
  low_pass integer DEFAULT 0,
  inner_iter integer DEFAULT 2,
  outer_iter integer DEFAULT 0
) RETURNS stl_result
```

Use the convenience functions when only one component is needed:

```sql
SELECT stl_trend(values, 12) FROM series;
SELECT stl_seasonal(values, 12) FROM series;
SELECT stl_residual(values, 12) FROM series;
```

### Ordered Collection Helper

The SQL file also defines `stl_collect_ordered(tbl regclass, val text, ord text)` to collect a column into an ordered `double precision[]`:

```sql
SELECT stl_decompose(
  stl_collect_ordered('monthly_revenue'::regclass, 'revenue', 'month'),
  12
);
```

### Holt-Winters Forecasting

`holt_winters_predict(seasonal_type text, period_length int, start_data_array real[])` forecasts one seasonal cycle ahead. `seasonal_type` is `'mult'` for multiplicative seasonality or `'add'` for additive seasonality:

```sql
SELECT *
FROM holt_winters_predict(
  'mult',
  4,
  (SELECT array_agg(revenue ORDER BY date)::real[] FROM sales)
);
```

The SQL implementation chooses smoothing coefficients automatically: first by 500 random initializations, then by refinement in `0.001` steps to minimize squared error. The helper `holt_winters_mse(...)` is present as the error-calculation routine used by the predictor.

### Caveats

- `stl_decompose` expects a `double precision[]` with no `NULL` values.
- The README states the series length must be at least `2 * period`.
- `seasonal` must be an odd integer greater than or equal to `3`.
- Holt-Winters expects a `real[]` input and supports only `'mult'` and `'add'` seasonal types.
