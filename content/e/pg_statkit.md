---
title: "pg_statkit"
linkTitle: "pg_statkit"
description: "Descriptive statistics and clinical effect-size functions over arrays"
weight: 4690
categories: ["FUNC"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_statkit**](https://github.com/kirdmi/pg_statkit) : Descriptive statistics and clinical effect-size functions over arrays


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4690** | {{< badge content="pg_statkit" link="https://github.com/kirdmi/pg_statkit" >}} | {{< ext "pg_statkit" >}} | `1.1.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "weighted_statistics" >}} {{< ext "quantile" >}} {{< ext "pg_math" >}} {{< ext "xicor" >}} {{< ext "plr" >}} |

> [!Note] Release Cargo/CI support PostgreSQL 13-17 with pgrx 0.18.1; no PG18 feature. This is statistical computation, not server monitoring.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_statkit` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "pg_statkit_18" "red" >}} {{< bg "17" "pg_statkit_17" "green" >}} {{< bg "16" "pg_statkit_16" "green" >}} {{< bg "15" "pg_statkit_15" "green" >}} {{< bg "14" "pg_statkit_14" "green" >}} | `pg_statkit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-pg-statkit" "red" >}} {{< bg "17" "postgresql-17-pg-statkit" "green" >}} {{< bg "16" "postgresql-16-pg-statkit" "green" >}} {{< bg "15" "postgresql-15-pg-statkit" "green" >}} {{< bg "14" "postgresql-14-pg-statkit" "green" >}} | `postgresql-$v-pg-statkit` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_statkit_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_statkit_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "pg_statkit_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "pg_statkit_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "pg_statkit_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "pg_statkit_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_statkit_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-18-pg-statkit : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-statkit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-statkit : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statkit_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 196.4 KiB | [pg_statkit_17-1.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statkit_17-1.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statkit_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 175.7 KiB | [pg_statkit_17-1.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statkit_17-1.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statkit_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 200.1 KiB | [pg_statkit_17-1.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statkit_17-1.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statkit_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 188.9 KiB | [pg_statkit_17-1.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statkit_17-1.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statkit_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 200.1 KiB | [pg_statkit_17-1.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statkit_17-1.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statkit_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 189.2 KiB | [pg_statkit_17-1.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statkit_17-1.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-statkit` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 164.2 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 143.4 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 163.9 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 143.5 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 177.9 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.6 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 174.8 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 166.0 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 174.6 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-statkit` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 165.6 KiB | [postgresql-17-pg-statkit_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statkit/postgresql-17-pg-statkit_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statkit_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 196.4 KiB | [pg_statkit_16-1.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statkit_16-1.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statkit_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 175.7 KiB | [pg_statkit_16-1.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statkit_16-1.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statkit_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 200.0 KiB | [pg_statkit_16-1.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statkit_16-1.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statkit_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 189.0 KiB | [pg_statkit_16-1.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statkit_16-1.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statkit_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 200.1 KiB | [pg_statkit_16-1.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statkit_16-1.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statkit_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 189.2 KiB | [pg_statkit_16-1.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statkit_16-1.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-statkit` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 164.2 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 144.0 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 164.2 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 144.0 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 177.9 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.7 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 174.8 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 166.1 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 174.5 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-statkit` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 165.6 KiB | [postgresql-16-pg-statkit_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statkit/postgresql-16-pg-statkit_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statkit_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 196.4 KiB | [pg_statkit_15-1.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statkit_15-1.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statkit_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 175.7 KiB | [pg_statkit_15-1.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statkit_15-1.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statkit_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 200.0 KiB | [pg_statkit_15-1.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statkit_15-1.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statkit_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 188.9 KiB | [pg_statkit_15-1.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statkit_15-1.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statkit_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 200.1 KiB | [pg_statkit_15-1.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statkit_15-1.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statkit_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 189.2 KiB | [pg_statkit_15-1.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statkit_15-1.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-statkit` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 164.2 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 143.7 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 164.2 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 144.0 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 177.9 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.8 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 174.8 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 166.1 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 174.6 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-statkit` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 165.6 KiB | [postgresql-15-pg-statkit_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statkit/postgresql-15-pg-statkit_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_statkit_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 196.4 KiB | [pg_statkit_14-1.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_statkit_14-1.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_statkit_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 175.7 KiB | [pg_statkit_14-1.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_statkit_14-1.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_statkit_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 200.0 KiB | [pg_statkit_14-1.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_statkit_14-1.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_statkit_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 189.0 KiB | [pg_statkit_14-1.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_statkit_14-1.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_statkit_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 200.1 KiB | [pg_statkit_14-1.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_statkit_14-1.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_statkit_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 189.2 KiB | [pg_statkit_14-1.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_statkit_14-1.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-statkit` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 166.1 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 146.2 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 166.1 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 146.2 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 180.1 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 168.8 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 176.8 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 168.2 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 176.7 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-statkit` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 167.7 KiB | [postgresql-14-pg-statkit_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-statkit/postgresql-14-pg-statkit_1.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/kirdmi/pg_statkit" title="Repository" icon="github" subtitle="github.com/kirdmi/pg_statkit" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_statkit-1.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_statkit;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_statkit;		# install via package name, for the active PG version

pig install pg_statkit -v 17;   # install for PG 17
pig install pg_statkit -v 16;   # install for PG 16
pig install pg_statkit -v 15;   # install for PG 15
pig install pg_statkit -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_statkit;
```

## Usage

Sources:

- [Official README](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/README.md)
- [Extension control file](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/pg_statkit.control)
- [pgrx manifest](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/Cargo.toml)

`pg_statkit` provides descriptive statistics and clinical effect-size functions over PostgreSQL arrays without exporting source data to a separate analysis process.

### Enablement

Release 1.1.0 builds and tests PostgreSQL 13–17 with pgrx 0.18.1. Install the matching native artifact, then create the non-relocatable, untrusted extension as a superuser:

```sql
CREATE EXTENSION pg_statkit;
```

It does not require preloading or a restart.

### Descriptive Statistics

Functions accept `double precision[]`; explicitly cast array literals because decimal literals otherwise form `numeric[]`.

```sql
SELECT statkit_mean('{1,2,3,4,5}'::float8[]);
SELECT statkit_median('{1,2,3,4,5}'::float8[]);
SELECT statkit_iqr('{1,2,3,4,5}'::float8[]);
SELECT statkit_percentile('{1,2,3,4,5}'::float8[], 40);

SELECT patient_id,
       statkit_median(array_agg(value::float8)) AS median_value
FROM measurements
GROUP BY patient_id;
```

The surface also includes variance, population/sample standard deviation, MAD, coefficient of variation, and standard error.

### Effect Sizes and Interpretation

`statkit_risk_difference`, `statkit_risk_ratio`, and `statkit_odds_ratio` accept a 2×2 table as four `bigint` counts. `statkit_cohens_d`, `statkit_hedges_g`, and `statkit_glass_delta` accept two samples.

```sql
SELECT statkit_odds_ratio(a => 8, b => 2, c => 2, d => 8);

SELECT statkit_cohens_d(
  array_agg(value::float8) FILTER (WHERE arm = 'treatment'),
  array_agg(value::float8) FILTER (WHERE arm = 'control')
)
FROM trial;
```

The odds-ratio function applies a 0.5 correction when any cell is zero. Undefined inputs return `NULL`. Version 1.1.0 does not provide confidence intervals, hypothesis tests, or assumption checks; use a full statistical package when those are required for research conclusions.

