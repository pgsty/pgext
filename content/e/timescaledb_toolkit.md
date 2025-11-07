---
title: "timescaledb_toolkit"
linkTitle: "timescaledb_toolkit"
description: "Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities"
weight: 1010
categories: ["TIME"]
width: full
---

Library of analytical hyperfunctions, time-series pipelining, and other SQL utilities


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1010** | {{< badge content="timescaledb_toolkit" link="https://github.com/timescale/timescaledb-toolkit" >}} | {{< ext "timescaledb_toolkit" >}} | `1.22.0` | {{< category "TIME" >}} | {{< license "Timescale" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "timescaledb" >}} {{< ext "timeseries" >}} {{< ext "periods" >}} {{< ext "temporal_tables" >}} {{< ext "emaj" >}} {{< ext "pg_cron" >}} {{< ext "pg_partman" >}} {{< ext "table_version" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/timescaledb_toolkit" >}} | `1.22.0` | {{< bg "18" "timescaledb-toolkit_18" "green" >}} {{< bg "17" "timescaledb-toolkit_17" "green" >}} {{< bg "16" "timescaledb-toolkit_16" "green" >}} {{< bg "15" "timescaledb-toolkit_15" "green" >}} {{< bg "14" "timescaledb-toolkit_14" "red" >}} {{< bg "13" "timescaledb-toolkit_13" "red" >}} | `timescaledb-toolkit_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/timescaledb_toolkit" >}} | `1.22.0` | {{< bg "18" "postgresql-18-timescaledb-toolkit" "green" >}} {{< bg "17" "postgresql-17-timescaledb-toolkit" "green" >}} {{< bg "16" "postgresql-16-timescaledb-toolkit" "green" >}} {{< bg "15" "postgresql-15-timescaledb-toolkit" "green" >}} {{< bg "14" "postgresql-14-timescaledb-toolkit" "red" >}} {{< bg "13" "postgresql-13-timescaledb-toolkit" "red" >}} | `postgresql-$v-timescaledb-toolkit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.19.0" "timescaledb-toolkit_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "timescaledb-toolkit_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.19.0" "timescaledb-toolkit_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "timescaledb-toolkit_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19.0" "timescaledb-toolkit_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "timescaledb-toolkit_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.19.0" "timescaledb-toolkit_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "timescaledb-toolkit_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "timescaledb-toolkit_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "timescaledb-toolkit_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "timescaledb-toolkit_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "timescaledb-toolkit_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "timescaledb-toolkit_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-18-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-17-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-16-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-15-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19.0" "postgresql-14-timescaledb-toolkit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-toolkit : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-18-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-17-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-16-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-15-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19.0" "postgresql-14-timescaledb-toolkit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-toolkit : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-timescaledb-toolkit : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-timescaledb-toolkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-timescaledb-toolkit : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-18-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-17-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-16-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-15-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19.0" "postgresql-14-timescaledb-toolkit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-toolkit : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-18-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-17-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-16-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-15-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19.0" "postgresql-14-timescaledb-toolkit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-toolkit : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-18-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-17-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-16-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-15-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19.0" "postgresql-14-timescaledb-toolkit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-toolkit : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-18-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-17-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-16-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.22.0" "postgresql-15-timescaledb-toolkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19.0" "postgresql-14-timescaledb-toolkit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-timescaledb-toolkit : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-toolkit_18` | 1.22.0 | `el8.x86_64` | pigsty | 3.3 MiB | [timescaledb-toolkit_18-1.22.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_18-1.22.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_18` | 1.22.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_18-1.22.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_18-1.22.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_18` | 1.22.0 | `el9.x86_64` | pigsty | 3.3 MiB | [timescaledb-toolkit_18-1.22.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-toolkit_18-1.22.0-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-toolkit_18` | 1.22.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_18-1.22.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_18-1.22.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_18` | 1.22.0 | `el10.x86_64` | pigsty | 3.4 MiB | [timescaledb-toolkit_18-1.22.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-toolkit_18-1.22.0-1PIGSTY.el10.x86_64.rpm) |
| `timescaledb-toolkit_18` | 1.22.0 | `el10.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_18-1.22.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-toolkit_18-1.22.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-timescaledb-toolkit` | 1.22.0 | `d12.x86_64` | pigsty | 2.8 MiB | [postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-timescaledb-toolkit` | 1.22.0 | `d12.aarch64` | pigsty | 11.2 KiB | [postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-timescaledb-toolkit` | 1.22.0 | `u22.x86_64` | pigsty | 3.1 MiB | [postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-timescaledb-toolkit` | 1.22.0 | `u22.aarch64` | pigsty | 2.7 MiB | [postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-timescaledb-toolkit` | 1.22.0 | `u24.x86_64` | pigsty | 3.1 MiB | [postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-timescaledb-toolkit` | 1.22.0 | `u24.aarch64` | pigsty | 2.6 MiB | [postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-18-timescaledb-toolkit_1.22.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-toolkit_17` | 1.22.0 | `el8.x86_64` | pigsty | 3.3 MiB | [timescaledb-toolkit_17-1.22.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_17-1.22.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_17` | 1.21.0 | `el8.x86_64` | pigsty | 3.2 MiB | [timescaledb-toolkit_17-1.21.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_17-1.21.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_17` | 1.19.0 | `el8.x86_64` | pigsty | 3.2 MiB | [timescaledb-toolkit_17-1.19.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_17-1.19.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_17` | 1.22.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_17-1.22.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_17-1.22.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_17` | 1.21.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_17-1.21.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_17-1.21.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_17` | 1.19.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_17-1.19.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_17-1.19.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_17` | 1.22.0 | `el9.x86_64` | pigsty | 3.3 MiB | [timescaledb-toolkit_17-1.22.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-toolkit_17-1.22.0-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-toolkit_17` | 1.22.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_17-1.22.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_17-1.22.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_17` | 1.21.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_17-1.21.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_17-1.21.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_17` | 1.19.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_17-1.19.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_17-1.19.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_17` | 1.22.0 | `el10.x86_64` | pigsty | 3.4 MiB | [timescaledb-toolkit_17-1.22.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-toolkit_17-1.22.0-1PIGSTY.el10.x86_64.rpm) |
| `timescaledb-toolkit_17` | 1.22.0 | `el10.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_17-1.22.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-toolkit_17-1.22.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-timescaledb-toolkit` | 1.22.0 | `d12.x86_64` | pigsty | 2.8 MiB | [postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-timescaledb-toolkit` | 1.22.0 | `d12.aarch64` | pigsty | 2.3 MiB | [postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-timescaledb-toolkit` | 1.22.0 | `u22.x86_64` | pigsty | 3.1 MiB | [postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-timescaledb-toolkit` | 1.22.0 | `u22.aarch64` | pigsty | 2.7 MiB | [postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-timescaledb-toolkit` | 1.22.0 | `u24.x86_64` | pigsty | 3.1 MiB | [postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-timescaledb-toolkit` | 1.22.0 | `u24.aarch64` | pigsty | 2.6 MiB | [postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-17-timescaledb-toolkit_1.22.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-toolkit_16` | 1.22.0 | `el8.x86_64` | pigsty | 3.3 MiB | [timescaledb-toolkit_16-1.22.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_16-1.22.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_16` | 1.21.0 | `el8.x86_64` | pigsty | 3.2 MiB | [timescaledb-toolkit_16-1.21.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_16-1.21.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_16` | 1.19.0 | `el8.x86_64` | pigsty | 3.2 MiB | [timescaledb-toolkit_16-1.19.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_16-1.19.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_16` | 1.22.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_16-1.22.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_16-1.22.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_16` | 1.21.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_16-1.21.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_16-1.21.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_16` | 1.19.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_16-1.19.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_16-1.19.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_16` | 1.22.0 | `el9.x86_64` | pigsty | 3.3 MiB | [timescaledb-toolkit_16-1.22.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-toolkit_16-1.22.0-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-toolkit_16` | 1.22.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_16-1.22.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_16-1.22.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_16` | 1.21.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_16-1.21.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_16-1.21.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_16` | 1.19.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_16-1.19.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_16-1.19.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_16` | 1.22.0 | `el10.x86_64` | pigsty | 3.4 MiB | [timescaledb-toolkit_16-1.22.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-toolkit_16-1.22.0-1PIGSTY.el10.x86_64.rpm) |
| `timescaledb-toolkit_16` | 1.22.0 | `el10.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_16-1.22.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-toolkit_16-1.22.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-timescaledb-toolkit` | 1.22.0 | `d12.x86_64` | pigsty | 2.8 MiB | [postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-timescaledb-toolkit` | 1.22.0 | `d12.aarch64` | pigsty | 2.3 MiB | [postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-timescaledb-toolkit` | 1.22.0 | `u22.x86_64` | pigsty | 3.1 MiB | [postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-timescaledb-toolkit` | 1.22.0 | `u22.aarch64` | pigsty | 2.7 MiB | [postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-timescaledb-toolkit` | 1.22.0 | `u24.x86_64` | pigsty | 3.1 MiB | [postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-timescaledb-toolkit` | 1.22.0 | `u24.aarch64` | pigsty | 2.6 MiB | [postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-16-timescaledb-toolkit_1.22.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-toolkit_15` | 1.22.0 | `el8.x86_64` | pigsty | 3.3 MiB | [timescaledb-toolkit_15-1.22.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_15-1.22.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_15` | 1.21.0 | `el8.x86_64` | pigsty | 3.2 MiB | [timescaledb-toolkit_15-1.21.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_15-1.21.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_15` | 1.19.0 | `el8.x86_64` | pigsty | 3.2 MiB | [timescaledb-toolkit_15-1.19.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_15-1.19.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_15` | 1.22.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_15-1.22.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_15-1.22.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_15` | 1.21.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_15-1.21.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_15-1.21.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_15` | 1.19.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_15-1.19.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_15-1.19.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_15` | 1.22.0 | `el9.x86_64` | pigsty | 3.3 MiB | [timescaledb-toolkit_15-1.22.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-toolkit_15-1.22.0-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-toolkit_15` | 1.22.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_15-1.22.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_15-1.22.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_15` | 1.21.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_15-1.21.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_15-1.21.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_15` | 1.19.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_15-1.19.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_15-1.19.0-1PIGSTY.el9.aarch64.rpm) |
| `timescaledb-toolkit_15` | 1.22.0 | `el10.x86_64` | pigsty | 3.4 MiB | [timescaledb-toolkit_15-1.22.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/timescaledb-toolkit_15-1.22.0-1PIGSTY.el10.x86_64.rpm) |
| `timescaledb-toolkit_15` | 1.22.0 | `el10.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_15-1.22.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/timescaledb-toolkit_15-1.22.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-timescaledb-toolkit` | 1.22.0 | `d12.x86_64` | pigsty | 2.8 MiB | [postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-timescaledb-toolkit` | 1.22.0 | `d12.aarch64` | pigsty | 2.3 MiB | [postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-timescaledb-toolkit` | 1.22.0 | `u22.x86_64` | pigsty | 3.1 MiB | [postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-timescaledb-toolkit` | 1.22.0 | `u22.aarch64` | pigsty | 2.7 MiB | [postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-timescaledb-toolkit` | 1.22.0 | `u24.x86_64` | pigsty | 3.1 MiB | [postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-timescaledb-toolkit` | 1.22.0 | `u24.aarch64` | pigsty | 2.6 MiB | [postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-15-timescaledb-toolkit_1.22.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timescaledb-toolkit_14` | 1.19.0 | `el8.x86_64` | pigsty | 3.2 MiB | [timescaledb-toolkit_14-1.19.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/timescaledb-toolkit_14-1.19.0-1PIGSTY.el8.x86_64.rpm) |
| `timescaledb-toolkit_14` | 1.19.0 | `el8.aarch64` | pigsty | 2.8 MiB | [timescaledb-toolkit_14-1.19.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/timescaledb-toolkit_14-1.19.0-1PIGSTY.el8.aarch64.rpm) |
| `timescaledb-toolkit_14` | 1.19.0 | `el9.x86_64` | pigsty | 3.2 MiB | [timescaledb-toolkit_14-1.19.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/timescaledb-toolkit_14-1.19.0-1PIGSTY.el9.x86_64.rpm) |
| `timescaledb-toolkit_14` | 1.19.0 | `el9.aarch64` | pigsty | 3.0 MiB | [timescaledb-toolkit_14-1.19.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/timescaledb-toolkit_14-1.19.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-timescaledb-toolkit` | 1.19.0 | `d12.x86_64` | pigsty | 2.7 MiB | [postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-timescaledb-toolkit` | 1.19.0 | `d12.aarch64` | pigsty | 2.3 MiB | [postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timescaledb-toolkit/postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-timescaledb-toolkit` | 1.19.0 | `u22.x86_64` | pigsty | 3.0 MiB | [postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-timescaledb-toolkit` | 1.19.0 | `u22.aarch64` | pigsty | 2.7 MiB | [postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timescaledb-toolkit/postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-timescaledb-toolkit` | 1.19.0 | `u24.x86_64` | pigsty | 2.9 MiB | [postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-timescaledb-toolkit` | 1.19.0 | `u24.aarch64` | pigsty | 2.7 MiB | [postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timescaledb-toolkit/postgresql-14-timescaledb-toolkit_1.19.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/timescaledb-toolkit" title="Repository" icon="github" subtitle="github.com/timescale/timescaledb-toolkit" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="timescaledb-toolkit-1.22.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get timescaledb_toolkit; # get timescaledb_toolkit source code
pig build dep timescaledb_toolkit; # install build dependencies
pig build pkg timescaledb_toolkit; # build extension rpm or deb
pig build ext timescaledb_toolkit; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install timescaledb_toolkit; # install by extension name, for the current active PG version
pig ext install timescaledb_toolkit; # install via package alias, for the active PG version
pig ext install timescaledb_toolkit -v 18;   # install for PG 18
pig ext install timescaledb_toolkit -v 17;   # install for PG 17
pig ext install timescaledb_toolkit -v 16;   # install for PG 16
pig ext install timescaledb_toolkit -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION timescaledb_toolkit;
```



--------

## Usage

TimescaleDB Toolkit provides specialized functions for time-series analytics using a **two-step aggregation pattern**. Most functions create intermediate representations that accessor functions then query, enabling efficient reuse and multiple analyses.

### Approximate Analytics

#### HyperLogLog - Distinct Count Estimation

Probabilistic distinct counting with configurable precision for high-cardinality datasets.

```sql
-- Estimate unique users per day
SELECT 
    date_trunc('day', timestamp) as day,
    distinct_count(hyperloglog(64, user_id)) as unique_users
FROM events 
GROUP BY day;

-- Combine counts across partitions
SELECT distinct_count(rollup(hll))
FROM (SELECT hyperloglog(32, session_id) as hll FROM events_2023
      UNION ALL 
      SELECT hyperloglog(32, session_id) FROM events_2024) t;
```

#### T-Digest - Quantile Approximation

High-accuracy percentile estimation optimized for tail quantiles (P95, P99).

```sql
-- Track response time percentiles
SELECT 
    service_name,
    approx_percentile(0.50, tdigest(100, response_time)) as p50,
    approx_percentile(0.95, tdigest(100, response_time)) as p95,
    approx_percentile(0.99, tdigest(100, response_time)) as p99
FROM api_metrics 
GROUP BY service_name;

-- Hourly percentiles with continuous aggregation
CREATE MATERIALIZED VIEW hourly_percentiles AS
SELECT 
    time_bucket('1 hour', timestamp) as hour,
    tdigest(200, response_time) as digest
FROM requests GROUP BY hour;
```

#### UddSketch - Bounded Error Quantiles

Quantile estimation with guaranteed maximum relative error bounds.

```sql
-- CPU utilization percentiles with 1% max error
SELECT 
    host_id,
    approx_percentile(0.95, uddsketch(100, 0.01, cpu_percent)) as p95_cpu,
    error(uddsketch(100, 0.01, cpu_percent)) as actual_error
FROM system_metrics 
GROUP BY host_id;
```

### Counter Analytics

#### Counter Aggregates - Monotonic Metrics

Handle counters that increase monotonically with automatic reset detection.

```sql
-- Request rate calculation
SELECT 
    time_bucket('5 min', timestamp) as bucket,
    rate(counter_agg(timestamp, request_count)) as requests_per_sec,
    delta(counter_agg(timestamp, request_count)) as total_requests
FROM metrics 
GROUP BY bucket;

-- Extrapolated rate for partial buckets
SELECT 
    extrapolated_rate(
        counter_agg(timestamp, bytes_sent, 
                   bounds => time_bucket_range('1 hour', timestamp))
    ) as bytes_per_second
FROM network_stats;
```

#### Gauge Aggregates - Varying Metrics  

Analytics for metrics that vary up and down (temperature, memory usage).

```sql
-- Temperature change analysis
SELECT 
    sensor_id,
    delta(gauge_agg(timestamp, temperature)) as temp_delta,
    rate(gauge_agg(timestamp, temperature)) as temp_rate_per_sec
FROM weather_data 
GROUP BY sensor_id;
```

### Time-Weighted Analytics

#### Time-Weighted Averages

Handle irregularly sampled data with interpolation methods (LOCF, Linear).

```sql
-- Weighted average for irregular sensor readings
SELECT 
    device_id,
    average(time_weight('LOCF', timestamp, sensor_value)) as weighted_avg,
    average(time_weight('Linear', timestamp, sensor_value)) as linear_avg
FROM iot_readings 
GROUP BY device_id;

-- Combining multiple time ranges
SELECT average(rollup(tw))
FROM (SELECT time_weight('LOCF', ts, val) as tw FROM readings_2023
      UNION ALL
      SELECT time_weight('LOCF', ts, val) FROM readings_2024) t;
```

### Data Visualization

#### LTTB Downsampling

Downsample time series while preserving visual similarity for charts.

```sql
-- Reduce 100K points to 1K for visualization  
SELECT time, value
FROM unnest((
    SELECT lttb(timestamp, price, 1000)
    FROM stock_prices 
    WHERE symbol = 'AAPL'
));
```

#### ASAP Smoothing

Generate human-readable graphs by reducing noise while preserving trends.

```sql
-- Smooth daily data to weekly resolution
SELECT time, value 
FROM unnest((
    SELECT asap_smooth(date, daily_sales, 52)
    FROM sales_data
    WHERE date >= '2023-01-01'
));
```

### Statistical Analysis

#### Stats Aggregates

Comprehensive statistical analysis with 1D and 2D regression capabilities.

```sql
-- Multi-variable analysis
SELECT 
    -- Basic statistics
    average(stats_agg(response_time)) as avg_response,
    stddev(stats_agg(response_time)) as response_stddev,
    
    -- Regression analysis
    slope(stats_agg(response_time, request_size)) as size_impact,
    corr(stats_agg(response_time, request_size)) as correlation,
    determination_coeff(stats_agg(response_time, request_size)) as r_squared
FROM performance_data;
```

### Timevector Data Type

Efficient intermediate representation for time series operations.

```sql
-- Create and manipulate timevector
CREATE VIEW cpu_series AS 
SELECT host_id, timevector(timestamp, cpu_percent) as ts
FROM system_metrics GROUP BY host_id;

-- Chain operations on timevector
SELECT host_id, unnest(lttb(ts, 100)) 
FROM cpu_series;
```

## Integration Patterns

### Continuous Aggregation Support

Most toolkit functions work seamlessly with TimescaleDB continuous aggregates:

```sql
CREATE MATERIALIZED VIEW hourly_analytics AS
SELECT 
    time_bucket('1 hour', timestamp) as hour,
    service_name,
    tdigest(100, response_time) as response_digest,
    counter_agg(timestamp, request_count) as request_counter,
    hyperloglog(64, user_id) as unique_users
FROM api_events
GROUP BY hour, service_name;

-- Query pre-computed aggregates
SELECT 
    hour,
    approx_percentile(0.95, response_digest) as p95_response,
    rate(request_counter) as req_per_sec,
    distinct_count(unique_users) as unique_users
FROM hourly_analytics
WHERE hour >= NOW() - INTERVAL '24 hours';
```

### Two-Step Analysis Pattern

Store intermediate aggregates for multiple analyses:

```sql
-- Step 1: Create aggregates
CREATE TABLE daily_summaries AS
SELECT 
    date_trunc('day', timestamp) as day,
    tdigest(200, response_time) as response_digest,
    stats_agg(response_time, request_size) as stats
FROM requests GROUP BY day;

-- Step 2: Multiple analyses from same data
SELECT 
    day,
    approx_percentile(0.50, response_digest) as median,
    approx_percentile(0.99, response_digest) as p99,
    average(stats) as avg_response,
    slope(stats) as size_correlation
FROM daily_summaries;
```

All functions in the **experimental** schema (`toolkit_experimental`) may change between versions. Use stable functions for production workloads requiring API stability.
