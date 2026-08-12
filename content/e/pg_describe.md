---
title: "pg_describe"
linkTitle: "pg_describe"
description: "Report a query's parameters and result columns without executing it"
weight: 4350
categories: ["UTIL"]
width: full
---

[**pg_describe**](https://github.com/sajonaro/pg_describe) : Report a query's parameters and result columns without executing it


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4350** | {{< badge content="pg_describe" link="https://github.com/sajonaro/pg_describe" >}} | {{< ext "pg_describe" >}} | `1.0.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "describe_resultset" >}} {{< ext "colnames" >}} {{< ext "ddlx" >}} {{< ext "pg_readme" >}} {{< ext "pglinter" >}} |

> [!Note] Uses PostgreSQL parser and analyzer without invoking the executor; upstream and PIGSTY packages require PostgreSQL 17 or newer.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_describe` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_describe_18" "green" >}} {{< bg "17" "pg_describe_17" "green" >}} {{< bg "16" "pg_describe_16" "red" >}} {{< bg "15" "pg_describe_15" "red" >}} {{< bg "14" "pg_describe_14" "red" >}} | `pg_describe_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-describe" "green" >}} {{< bg "17" "postgresql-17-pg-describe" "green" >}} {{< bg "16" "postgresql-16-pg-describe" "red" >}} {{< bg "15" "postgresql-15-pg-describe" "red" >}} {{< bg "14" "postgresql-14-pg-describe" "red" >}} | `postgresql-$v-pg-describe` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_describe_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_describe_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_describe_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_describe_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_describe_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_describe_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_describe_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_describe_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-describe : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-describe : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-describe : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-describe : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_describe_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.2 KiB | [pg_describe_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_describe_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_describe_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.3 KiB | [pg_describe_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_describe_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_describe_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.9 KiB | [pg_describe_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_describe_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_describe_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.8 KiB | [pg_describe_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_describe_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_describe_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.9 KiB | [pg_describe_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_describe_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_describe_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.1 KiB | [pg_describe_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_describe_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-describe` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 35.9 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.8 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 35.9 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 35.8 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.8 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.7 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.2 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 36.9 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 37.1 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-describe` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 37.2 KiB | [postgresql-18-pg-describe_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-describe/postgresql-18-pg-describe_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_describe_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 34.2 KiB | [pg_describe_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_describe_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_describe_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 34.3 KiB | [pg_describe_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_describe_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_describe_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.9 KiB | [pg_describe_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_describe_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_describe_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.8 KiB | [pg_describe_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_describe_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_describe_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.9 KiB | [pg_describe_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_describe_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_describe_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 34.1 KiB | [pg_describe_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_describe_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-describe` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 35.9 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 35.7 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 35.9 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 35.8 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.2 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 40.2 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.1 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 36.8 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 37.1 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-describe` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 37.1 KiB | [postgresql-17-pg-describe_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-describe/postgresql-17-pg-describe_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/sajonaro/pg_describe" title="Repository" icon="github" subtitle="github.com/sajonaro/pg_describe" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_describe-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_describe;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_describe;		# install via package name, for the active PG version

pig install pg_describe -v 18;   # install for PG 18
pig install pg_describe -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_describe;
```

## Usage

Sources:

- [pg_describe 1.0.0 README](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/README.md)
- [pg_describe documentation](https://sajonaro.github.io/pg_describe/)
- [pg_describe 1.0.0 control file](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/pg_describe.control)
- [pg_describe 1.0.0 SQL](https://api.pgxn.org/src/pg_describe/pg_describe-1.0.0/sql/pg_describe--1.0.0.sql)

`pg_describe` reports the parameters and result columns of a SQL statement without executing it. It uses PostgreSQL parsing and analysis to infer parameter types, wire-visible result types, source-column provenance, and outer-join-aware nullability. Use it for code generation, migration checks, and query-contract tooling.

### Describe a Query

```sql
CREATE EXTENSION pg_describe;

SELECT *
FROM pg_describe(
  'SELECT id, email FROM users WHERE id = $1'
);
```

Rows with `kind = 'param'` describe `$1`, `$2`, and later parameters. Rows with `kind = 'column'` describe result-column order, name, type OID/name, source table/column, base `NOT NULL` status, and whether the final expression is known non-null.

### Check Join Nullability

```sql
SELECT *
FROM pg_describe($query$
  SELECT o.id, c.email
  FROM orders AS o
  LEFT JOIN customers AS c ON c.id = o.customer_id
  WHERE o.placed_at >= $1
$query$);
```

Even when `customers.email` is declared `NOT NULL`, `result_not_null` is false because a left join can null-extend the row. This distinction is useful when generating nullable client types.

### Execution and Security Boundary

- The statement is parsed and analyzed but not executed. Describing a `DELETE`, volatile function call, or expensive query does not run the statement.
- Normal name resolution and privilege checks still apply. Callers cannot use `pg_describe` to inspect objects they could not reference themselves.
- Parameter types must be inferable from context; ambiguous `$n` parameters still produce PostgreSQL analysis errors.
- The result describes PostgreSQL's analyzed output, not dynamic SQL assembled later by an application.

### Requirements and Caveats

- Upstream 1.0.0 requires PostgreSQL 17; PostgreSQL 16 is described as possibly working but untested. Pigsty packages target PostgreSQL 17 and 18.
- The extension is relocatable and does not require preloading or a restart.
- The companion `pg-describe-gen` TypeScript tool is a separate npm package. The PostgreSQL extension works without it.
- This is a young API. Pin the extension/tool versions in CI and review generated changes alongside schema migrations.
