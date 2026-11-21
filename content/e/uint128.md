---
title: "uint128"
linkTitle: "uint128"
description: "Native uint128 type"
weight: 3760
categories: ["TYPE"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_uint128**](https://github.com/pg-uint/pg-uint128) : Native uint128 type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3760** | {{< badge content="uint128" link="https://github.com/pg-uint/pg-uint128" >}} | {{< ext "uint128" "pg_uint128" >}} | `1.3.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "uint" >}} {{< ext "pgmp" >}} {{< ext "numeral" >}} {{< ext "unit" >}} {{< ext "pg_rational" >}} {{< ext "seg" >}} {{< ext "cube" >}} |

> [!Note] breaks on el8 since 1.1 ,fix el8 build problem by adding __has_builtin marco


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_uint128` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.0` | {{< bg "19" "pg_uint128_19" "green" >}} {{< bg "18" "pg_uint128_18" "green" >}} {{< bg "17" "pg_uint128_17" "green" >}} {{< bg "16" "pg_uint128_16" "green" >}} {{< bg "15" "pg_uint128_15" "green" >}} | `pg_uint128_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.0` | {{< bg "19" "postgresql-19-pg-uint128" "green" >}} {{< bg "18" "postgresql-18-pg-uint128" "green" >}} {{< bg "17" "postgresql-17-pg-uint128" "green" >}} {{< bg "16" "postgresql-16-pg-uint128" "green" >}} {{< bg "15" "postgresql-15-pg-uint128" "green" >}} | `postgresql-$v-pg-uint128` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_uint128_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_uint128_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "pg_uint128_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "MISS" "pg_uint128_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "MISS" "pg_uint128_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "MISS" "pg_uint128_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_uint128_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-pg-uint128 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-uint128 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-uint128 : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_18` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.2 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_18-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.5 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_18-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 170.9 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_18-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.4 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_18-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 169.3 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_18-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_18` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 163.4 KiB | [pg_uint128_18-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_18-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 328.2 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 318.4 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 327.0 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 319.3 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 365.8 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 357.9 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 355.4 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.3 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 357.5 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-uint128` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 350.9 KiB | [postgresql-18-pg-uint128_1.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uint128/postgresql-18-pg-uint128_1.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.3 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_17-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.6 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_17-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 169.9 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_17-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.5 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_17-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 169.3 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_17-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 163.3 KiB | [pg_uint128_17-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_17-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 327.7 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 317.5 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 327.1 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 320.1 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 384.2 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 373.3 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 355.4 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.3 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 357.4 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-uint128` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 350.6 KiB | [postgresql-17-pg-uint128_1.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uint128/postgresql-17-pg-uint128_1.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.1 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_16-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.3 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_16-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 170.7 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_16-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.4 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_16-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 169.3 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_16-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 164.1 KiB | [pg_uint128_16-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_16-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 327.5 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 318.1 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 326.9 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 319.7 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 381.3 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 371.7 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 354.6 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.0 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 356.9 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-uint128` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 350.0 KiB | [postgresql-16-pg-uint128_1.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uint128/postgresql-16-pg-uint128_1.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uint128_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.7 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uint128_15-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 176.8 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uint128_15-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 170.2 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uint128_15-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 161.3 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uint128_15-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 172.0 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uint128_15-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_uint128_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 164.7 KiB | [pg_uint128_15-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uint128_15-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-uint128` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 331.8 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 322.1 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 331.3 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 324.3 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 383.8 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 374.0 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 359.3 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 356.3 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 358.7 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-uint128` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 351.7 KiB | [postgresql-15-pg-uint128_1.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-uint128/postgresql-15-pg-uint128_1.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pg-uint/pg-uint128" title="Repository" icon="github" subtitle="github.com/pg-uint/pg-uint128" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-uint128-1.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_uint128;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_uint128;		# install via package name, for the active PG version
pig install uint128;		# install by extension name, for the current active PG version

pig install uint128 -v 19;   # install for PG 19
pig install uint128 -v 18;   # install for PG 18
pig install uint128 -v 17;   # install for PG 17
pig install uint128 -v 16;   # install for PG 16
pig install uint128 -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION uint128;
```

## Usage

Sources:

- [README.md](https://api.pgxn.org/src/uint128/uint128-1.3.0/README.md)
- [uint128.control](https://api.pgxn.org/src/uint128/uint128-1.3.0/uint128.control)
- [uint128--1.2.0--1.3.0.sql](https://api.pgxn.org/src/uint128/uint128-1.3.0/uint128--1.2.0--1.3.0.sql)
- [1.3.0](https://github.com/pg-uint/pg-uint128/releases/tag/1.3.0)


`uint128` adds unsigned integers and signed 8-bit and 128-bit integer types. Version 1.3.0 supports PostgreSQL 12–18 and fixes operator metadata that could make indexed comparisons return incorrect results.

### Store and Query Integers

Installation requires a superuser; the control file does not mark the extension trusted. It is relocatable and requires no preload.

```sql
CREATE EXTENSION uint128;
CREATE TABLE ids (id uint16 PRIMARY KEY);
INSERT INTO ids VALUES ('340282366920938463463374607431768211455'::uint16);
SELECT * FROM ids WHERE id > '18446744073709551615'::uint16;
SELECT * FROM generate_series(1::uint4, 3::uint4);
```

The unsigned types `uint1`, `uint2`, `uint4`, `uint8` and `uint16` occupy 1, 2, 4, 8 and 16 bytes. Their range is zero through 2 raised to the bit width minus one. `int1` and `int16` are signed 1-byte and 16-byte types.

### Operators and Indexes

The types provide arithmetic, comparisons, bitwise operations, B-tree and hash indexes, binary I/O, casts and `generate_series`. Corresponding range types, including `uint1range` and `uint16range`, support GiST indexes; multiranges depend on server support. `SUM` and `AVG` use `numeric` internally; `MIN` and `MAX` are also available.

Mixed signed/unsigned arithmetic is sensitive to operand order and result type. Cast explicitly when a negative result is possible: an unsigned result cannot represent it. Conversions and arithmetic may raise range errors; do not assume wraparound.

### Upgrade to 1.3.0

Install the matching files, then run as a superuser:

```sql
ALTER EXTENSION uint128 UPDATE TO '1.3.0';
```

This corrects comparison commutators and removes invalid mixed-type addition/subtraction commutators. It also fixes installation with the stricter operator checks in PostgreSQL 17.12. The supplied SQL updates operator metadata in place while preserving data, indexes, views and materialized views. Check the installed version and representative indexed comparisons after upgrading.
