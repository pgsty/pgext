---
title: "ulak"
linkTitle: "ulak"
description: "Transactional Outbox extension for PostgreSQL with reliable asynchronous delivery"
weight: 2680
categories: ["FEAT"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**ulak**](https://github.com/zeybek/ulak) : Transactional Outbox extension for PostgreSQL with reliable asynchronous delivery


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2680** | {{< badge content="ulak" link="https://github.com/zeybek/ulak" >}} | {{< ext "ulak" >}} | `0.3.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `ulak` |
|   **See Also**    | {{< ext "pgmq" >}} {{< ext "pgmb" >}} {{< ext "pg_durable" >}} {{< ext "redis" >}} {{< ext "pgq" >}} {{< ext "pgmqtt" >}} {{< ext "pg_task" >}} {{< ext "tcn" >}} {{< ext "pg_background" >}} {{< ext "fsm_core" >}} |

> [!Note] preload required; rm README from RPM/DEB


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `ulak` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "19" "ulak_19" "green" >}} {{< bg "18" "ulak_18" "green" >}} {{< bg "17" "ulak_17" "green" >}} {{< bg "16" "ulak_16" "green" >}} {{< bg "15" "ulak_15" "green" >}} | `ulak_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "19" "postgresql-19-ulak" "green" >}} {{< bg "18" "postgresql-18-ulak" "green" >}} {{< bg "17" "postgresql-17-ulak" "green" >}} {{< bg "16" "postgresql-16-ulak" "green" >}} {{< bg "15" "postgresql-15-ulak" "green" >}} | `postgresql-$v-ulak` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "ulak_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "ulak_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_15 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "ulak_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.3.0" "ulak_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_15 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "ulak_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_15 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "ulak_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_15 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "ulak_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_15 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "ulak_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "ulak_15 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-19-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-ulak : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-ulak : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_19` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 357.6 KiB | [ulak_19-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_19-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_19` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 353.4 KiB | [ulak_19-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_19-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_19` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 360.5 KiB | [ulak_19-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_19-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_19` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 355.5 KiB | [ulak_19-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_19-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-19-ulak` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 306.3 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 299.2 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 337.1 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.3 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 321.3 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 316.8 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 344.1 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 340.0 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 342.6 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-ulak` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 337.8 KiB | [postgresql-19-ulak_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-19-ulak_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 330.9 KiB | [ulak_18-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_18-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `ulak_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 308.7 KiB | [ulak_18-0.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_18-0.2.0-1PGSTY.el8.x86_64.rpm) |
| `ulak_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 322.5 KiB | [ulak_18-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_18-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `ulak_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 303.1 KiB | [ulak_18-0.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_18-0.2.0-1PGSTY.el8.aarch64.rpm) |
| `ulak_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 358.5 KiB | [ulak_18-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_18-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 336.6 KiB | [ulak_18-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_18-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 354.0 KiB | [ulak_18-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_18-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 332.5 KiB | [ulak_18-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_18-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 361.5 KiB | [ulak_18-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_18-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 337.3 KiB | [ulak_18-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_18-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 356.1 KiB | [ulak_18-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_18-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `ulak_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 333.3 KiB | [ulak_18-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_18-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-ulak` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 307.4 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 300.0 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 337.9 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.8 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 322.5 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 317.2 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 345.0 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 340.6 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 343.9 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-ulak` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 338.5 KiB | [postgresql-18-ulak_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-18-ulak_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 330.8 KiB | [ulak_17-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_17-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `ulak_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 308.5 KiB | [ulak_17-0.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_17-0.2.0-1PGSTY.el8.x86_64.rpm) |
| `ulak_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 322.4 KiB | [ulak_17-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_17-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `ulak_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 303.1 KiB | [ulak_17-0.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_17-0.2.0-1PGSTY.el8.aarch64.rpm) |
| `ulak_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 358.5 KiB | [ulak_17-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_17-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 336.4 KiB | [ulak_17-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_17-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 353.7 KiB | [ulak_17-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_17-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 332.4 KiB | [ulak_17-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_17-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 361.4 KiB | [ulak_17-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_17-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 337.2 KiB | [ulak_17-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_17-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 356.2 KiB | [ulak_17-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_17-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `ulak_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 333.5 KiB | [ulak_17-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_17-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-ulak` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 307.3 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 299.7 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 337.6 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.2 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 348.3 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 344.6 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 345.2 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 340.5 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 343.9 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-ulak` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 338.2 KiB | [postgresql-17-ulak_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-17-ulak_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 330.8 KiB | [ulak_16-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_16-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `ulak_16` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 308.6 KiB | [ulak_16-0.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_16-0.2.0-1PGSTY.el8.x86_64.rpm) |
| `ulak_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 322.4 KiB | [ulak_16-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_16-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `ulak_16` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 302.9 KiB | [ulak_16-0.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_16-0.2.0-1PGSTY.el8.aarch64.rpm) |
| `ulak_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 358.5 KiB | [ulak_16-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_16-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_16` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 336.5 KiB | [ulak_16-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_16-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 353.8 KiB | [ulak_16-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_16-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 332.5 KiB | [ulak_16-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_16-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 361.4 KiB | [ulak_16-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_16-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_16` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 337.2 KiB | [ulak_16-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_16-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 356.1 KiB | [ulak_16-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_16-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `ulak_16` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 333.3 KiB | [ulak_16-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_16-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-ulak` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 307.2 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 299.6 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 337.8 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 329.1 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 348.0 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 344.2 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 345.2 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 340.5 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 343.8 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-ulak` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 338.2 KiB | [postgresql-16-ulak_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-16-ulak_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ulak_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 332.9 KiB | [ulak_15-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_15-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `ulak_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 310.7 KiB | [ulak_15-0.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ulak_15-0.2.0-1PGSTY.el8.x86_64.rpm) |
| `ulak_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 324.3 KiB | [ulak_15-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_15-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `ulak_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 304.6 KiB | [ulak_15-0.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ulak_15-0.2.0-1PGSTY.el8.aarch64.rpm) |
| `ulak_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 366.7 KiB | [ulak_15-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_15-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 343.4 KiB | [ulak_15-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ulak_15-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `ulak_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 361.4 KiB | [ulak_15-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_15-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 340.0 KiB | [ulak_15-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ulak_15-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `ulak_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 368.4 KiB | [ulak_15-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_15-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_15` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 344.5 KiB | [ulak_15-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ulak_15-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `ulak_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 363.3 KiB | [ulak_15-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_15-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `ulak_15` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 340.7 KiB | [ulak_15-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ulak_15-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-ulak` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 310.6 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 302.9 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 342.2 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 332.9 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 354.9 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 352.0 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 352.0 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 349.0 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 350.6 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-ulak` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 346.5 KiB | [postgresql-15-ulak_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/u/ulak/postgresql-15-ulak_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/zeybek/ulak" title="Repository" icon="github" subtitle="github.com/zeybek/ulak" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="ulak-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg ulak;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install ulak;		# install via package name, for the active PG version

pig install ulak -v 19;   # install for PG 19
pig install ulak -v 18;   # install for PG 18
pig install ulak -v 17;   # install for PG 17
pig install ulak -v 16;   # install for PG 16
pig install ulak -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'ulak';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION ulak;
```

## Usage

Sources:

- [README](https://api.pgxn.org/src/ulak/ulak-0.3.0/README.md)
- [Control](https://api.pgxn.org/src/ulak/ulak-0.3.0/ulak.control)
- [Changelog](https://api.pgxn.org/src/ulak/ulak-0.3.0/CHANGELOG.md)
- [Install SQL](https://api.pgxn.org/src/ulak/ulak-0.3.0/sql/ulak--0.3.0.sql)
- [Upgrade SQL](https://api.pgxn.org/src/ulak/ulak-0.3.0/sql/ulak--0.2.0--0.3.0.sql)

`ulak` implements a transactional outbox: enqueue messages in the same transaction as application data, then deliver them through background workers. Delivery to remote systems is at least once; consumers must tolerate duplicates.

### Enablement

PostgreSQL 14–18 is supported. Set `shared_preload_libraries` and `ulak.database`, restart PostgreSQL, then create the extension in that database. HTTP is built in; Kafka, MQTT, Redis Streams, AMQP and NATS require their optional compiled adapters and client libraries.

```ini
shared_preload_libraries = 'ulak'
ulak.database = 'appdb'
```

```sql
CREATE EXTENSION ulak;
```

### Enqueue and publish

Create endpoints as an administrator; grant application access through `ulak_application`. `ulak_admin` manages endpoints and operations; `ulak_monitor` reads monitoring data. Use `ulak.endpoint_status` for endpoint status without exposing endpoint credentials. A rolled-back transaction leaves no message.

```sql
SELECT ulak.create_endpoint('orders-http', 'http',
  '{"url":"https://example.com/webhook","method":"POST"}'::jsonb);
BEGIN;
SELECT ulak.send('orders-http', '{"event":"order.created","id":123}'::jsonb);
COMMIT;

SELECT ulak.send_batch('orders-http', ARRAY['{"id":1}'::jsonb, '{"id":2}'::jsonb]);
SELECT ulak.create_event_type('order.created', 'Order created');
SELECT ulak.subscribe('order.created', 'orders-http');
SELECT ulak.publish('order.created', '{"id":123}'::jsonb);
```

### Delivery limits and retries

`ulak.send_with_options` adds priority, scheduling, idempotency, expiry and ordering controls. In 0.3.0 all enqueue APIs enforce `ulak.max_payload_size`; oversized payloads raise `54000`. Batch APIs count each row against `ulak.max_queue_size` and raise `53400` on overflow. Stale processing rows consume retry budget and eventually enter the dead-letter queue. Workers use `synchronous_commit = local` and `lock_timeout = 10s`.

### Protocol and upgrade changes

0.3.0 retries HTTP 408/425, Redis OOM/READONLY and unroutable mandatory AMQP messages. Redis TLS now checks hostnames. `ulak.mqtt_default_qos` defaults to 1; multiple workers suffix configured MQTT client IDs with their worker number. NATS message IDs include cluster and database identity; AMQP message IDs use the queue row ID. Review downstream deduplication assumptions. Install the matching library and SQL, restart to load the worker library, and upgrade the extension. External listeners need `ulak.wake_notify = on` to receive `ulak_new_msg`; native workers use shared-memory wakeups.

```sql
ALTER EXTENSION ulak UPDATE TO '0.3.0';
```

### Monitoring and recovery

Tune `ulak.workers`, `ulak.poll_interval`, `ulak.batch_size`, retry and circuit-breaker settings for the workload. Inspect failures before redriving; redrive and replay can deliver an already-observed message again. HTTP SSRF restrictions apply unless explicitly relaxed.

```sql
SELECT * FROM ulak.health_check();
SELECT * FROM ulak.get_worker_status();
SELECT * FROM ulak.get_endpoint_health();
SELECT * FROM ulak.dlq_summary();
SELECT ulak.redrive_endpoint('orders-http');
```
