---
title: "pg_circuit"
linkTitle: "pg_circuit"
description: "Runtime observation, warnings and blocking for dangerous SQL statements"
weight: 5815
categories: ["ADMIN"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_circuit**](https://github.com/PG-Circuit/pg-circuit) : Runtime observation, warnings and blocking for dangerous SQL statements


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5815** | {{< badge content="pg_circuit" link="https://github.com/PG-Circuit/pg-circuit" >}} | {{< ext "pg_circuit" >}} | `0.1.0` | {{< category "ADMIN" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Requires shared_preload_libraries and restart; Community runtime pressure is informational.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_circuit` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_circuit_18" "green" >}} {{< bg "17" "pg_circuit_17" "green" >}} {{< bg "16" "pg_circuit_16" "green" >}} {{< bg "15" "pg_circuit_15" "red" >}} {{< bg "14" "pg_circuit_14" "red" >}} | `pg_circuit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-circuit" "green" >}} {{< bg "17" "postgresql-17-pg-circuit" "green" >}} {{< bg "16" "postgresql-16-pg-circuit" "green" >}} {{< bg "15" "postgresql-15-pg-circuit" "red" >}} {{< bg "14" "postgresql-14-pg-circuit" "red" >}} | `postgresql-$v-pg-circuit` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_circuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_circuit_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_circuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_circuit_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_circuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_circuit_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_circuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_circuit_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_circuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_circuit_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_circuit_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_circuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_circuit_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-circuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-circuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-circuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-circuit : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_circuit_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 102.7 KiB | [pg_circuit_18-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_circuit_18-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_circuit_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 101.1 KiB | [pg_circuit_18-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_circuit_18-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_circuit_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 100.9 KiB | [pg_circuit_18-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_circuit_18-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_circuit_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 100.0 KiB | [pg_circuit_18-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_circuit_18-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_circuit_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 101.3 KiB | [pg_circuit_18-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_circuit_18-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_circuit_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.6 KiB | [pg_circuit_18-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_circuit_18-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-circuit` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 92.3 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 91.3 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 93.5 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 92.2 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 97.3 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 96.7 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 93.8 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 93.6 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 92.6 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-circuit` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 91.8 KiB | [postgresql-18-pg-circuit_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-circuit/postgresql-18-pg-circuit_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_circuit_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 102.7 KiB | [pg_circuit_17-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_circuit_17-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_circuit_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 101.1 KiB | [pg_circuit_17-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_circuit_17-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_circuit_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 101.0 KiB | [pg_circuit_17-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_circuit_17-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_circuit_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 100.1 KiB | [pg_circuit_17-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_circuit_17-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_circuit_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 101.4 KiB | [pg_circuit_17-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_circuit_17-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_circuit_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.6 KiB | [pg_circuit_17-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_circuit_17-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-circuit` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 92.2 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 91.0 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 93.4 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 92.2 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 109.3 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 109.0 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 93.8 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 93.7 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 92.7 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-circuit` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 92.0 KiB | [postgresql-17-pg-circuit_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-circuit/postgresql-17-pg-circuit_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_circuit_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 102.8 KiB | [pg_circuit_16-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_circuit_16-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_circuit_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 101.2 KiB | [pg_circuit_16-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_circuit_16-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_circuit_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 101.0 KiB | [pg_circuit_16-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_circuit_16-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_circuit_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 100.2 KiB | [pg_circuit_16-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_circuit_16-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_circuit_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 101.4 KiB | [pg_circuit_16-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_circuit_16-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_circuit_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.7 KiB | [pg_circuit_16-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_circuit_16-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-circuit` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 92.3 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 91.1 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 93.5 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 92.3 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 109.1 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 108.6 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 93.8 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 93.7 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 92.9 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-circuit` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 92.1 KiB | [postgresql-16-pg-circuit_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-circuit/postgresql-16-pg-circuit_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/PG-Circuit/pg-circuit" title="Repository" icon="github" subtitle="github.com/PG-Circuit/pg-circuit" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_circuit-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_circuit;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_circuit;		# install via package name, for the active PG version

pig install pg_circuit -v 18;   # install for PG 18
pig install pg_circuit -v 17;   # install for PG 17
pig install pg_circuit -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_circuit';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_circuit;
```

## Usage

Sources:

- [README v0.1.0](https://github.com/PG-Circuit/pg-circuit/blob/v0.1.0/README.md)

`pg_circuit` inspects potentially dangerous DML and DDL on PostgreSQL 16–18. The Community edition can observe, warn or block statements. Add it to `shared_preload_libraries`, restart PostgreSQL, then create the extension as a superuser.

### Basic protection

```conf
shared_preload_libraries = 'pg_circuit'
```

```sql
CREATE EXTENSION pg_circuit;
CREATE TABLE circuit_demo (id integer);
INSERT INTO circuit_demo VALUES (1);
SET pg_circuit.mode = 'enforce';
DELETE FROM circuit_demo; -- blocked
DELETE FROM circuit_demo WHERE id = 1;
SELECT * FROM pg_circuit_status();
```

### Configuration and diagnostics

`pg_circuit.mode` defaults to warn; observe collects risk without warnings, and enforce blocks scores at or above the configured threshold. `pg_circuit_runtime_state()` reports pressure signals and `pg_circuit_events()` exposes recent events.

Community always reports effective runtime mode NORMAL. Pressure readings do not automatically escalate enforcement. This is a policy aid; application transactions, authorization and backups still determine data protection. Review rules and thresholds on representative queries before enabling enforcement.
