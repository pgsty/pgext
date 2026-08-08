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

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_describe;		# install via package name, for the active PG version

pig install pg_describe -v 18;   # install for PG 18
pig install pg_describe -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_describe;
```
