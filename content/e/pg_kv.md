---
title: "pg_kv"
linkTitle: "pg_kv"
description: "Shared-memory key/value cache, token-bucket rate limits and sorted sets."
weight: 2895
categories: ["FEAT"]
languages: ["Rust"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_kv**](https://github.com/Murat7Ay/pg_kv) : Shared-memory key/value cache, token-bucket rate limits and sorted sets.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2895** | {{< badge content="pg_kv" link="https://github.com/Murat7Ay/pg_kv" >}} | {{< ext "pg_kv" >}} | `0.1.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `kv` |

> [!Note] Early release, not production ready. Preload/restart and superuser installation required. Cluster-wide keyspace: nontransactional, not durable or replicated. Grant kv schema usage deliberately; prefixing keys is not tenant access control. Capacity eviction can reset rate-limit allowance.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_kv` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_kv_18" "green" >}} {{< bg "17" "pg_kv_17" "green" >}} {{< bg "16" "pg_kv_16" "green" >}} {{< bg "15" "pg_kv_15" "green" >}} {{< bg "14" "pg_kv_14" "red" >}} | `pg_kv_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-kv" "green" >}} {{< bg "17" "postgresql-17-pg-kv" "green" >}} {{< bg "16" "postgresql-16-pg-kv" "green" >}} {{< bg "15" "postgresql-15-pg-kv" "green" >}} {{< bg "14" "postgresql-14-pg-kv" "red" >}} | `postgresql-$v-pg-kv` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_kv_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_kv_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_kv_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_kv_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_kv_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_kv_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_kv_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pg-kv : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pg-kv : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-kv : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kv_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 242.2 KiB | [pg_kv_18-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kv_18-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_kv_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 213.2 KiB | [pg_kv_18-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kv_18-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_kv_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 246.2 KiB | [pg_kv_18-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kv_18-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_kv_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 229.6 KiB | [pg_kv_18-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kv_18-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_kv_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 246.2 KiB | [pg_kv_18-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kv_18-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_kv_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 229.9 KiB | [pg_kv_18-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kv_18-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-kv` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 209.5 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 181.5 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 209.5 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 181.5 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 224.8 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 208.4 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 223.5 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 206.6 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 223.0 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-kv` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 206.1 KiB | [postgresql-18-pg-kv_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kv/postgresql-18-pg-kv_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kv_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 242.1 KiB | [pg_kv_17-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kv_17-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_kv_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 213.3 KiB | [pg_kv_17-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kv_17-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_kv_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 246.1 KiB | [pg_kv_17-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kv_17-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_kv_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 229.4 KiB | [pg_kv_17-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kv_17-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_kv_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 246.2 KiB | [pg_kv_17-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kv_17-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_kv_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 229.8 KiB | [pg_kv_17-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kv_17-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-kv` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 209.5 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 181.7 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 209.6 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 181.4 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 224.7 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 208.2 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 223.4 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 206.2 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 223.0 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-kv` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 206.1 KiB | [postgresql-17-pg-kv_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kv/postgresql-17-pg-kv_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kv_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 242.2 KiB | [pg_kv_16-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kv_16-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_kv_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 213.3 KiB | [pg_kv_16-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kv_16-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_kv_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 246.3 KiB | [pg_kv_16-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kv_16-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_kv_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 229.5 KiB | [pg_kv_16-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kv_16-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_kv_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 246.2 KiB | [pg_kv_16-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kv_16-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_kv_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 229.8 KiB | [pg_kv_16-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kv_16-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-kv` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 209.7 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 181.3 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 209.7 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 181.3 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 224.9 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 208.2 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 223.5 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 206.6 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 223.1 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-kv` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 205.7 KiB | [postgresql-16-pg-kv_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kv/postgresql-16-pg-kv_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_kv_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 242.1 KiB | [pg_kv_15-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_kv_15-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_kv_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 213.3 KiB | [pg_kv_15-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_kv_15-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_kv_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 246.2 KiB | [pg_kv_15-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_kv_15-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_kv_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 229.4 KiB | [pg_kv_15-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_kv_15-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_kv_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 246.1 KiB | [pg_kv_15-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_kv_15-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_kv_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 229.8 KiB | [pg_kv_15-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_kv_15-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-kv` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 209.3 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 181.3 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 209.4 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 181.3 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 224.7 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 208.1 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 223.3 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 206.4 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 223.0 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-kv` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 205.7 KiB | [postgresql-15-pg-kv_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-kv/postgresql-15-pg-kv_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Murat7Ay/pg_kv" title="Repository" icon="github" subtitle="github.com/Murat7Ay/pg_kv" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_kv-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_kv;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_kv;		# install via package name, for the active PG version

pig install pg_kv -v 18;   # install for PG 18
pig install pg_kv -v 17;   # install for PG 17
pig install pg_kv -v 16;   # install for PG 16
pig install pg_kv -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_kv';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_kv;
```

## Usage

Sources:

- [Official source](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/README.md)
- [Official source](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/pg_kv.control)
- [Official source](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/Cargo.toml)
- [Official source](https://github.com/Murat7Ay/pg_kv/blob/bf45bbfd6f29782701a4e6c9bf9c672a0c02bc8f/src/lib.rs)

`pg_kv` 0.1.0 provides a shared-memory cache, counters, token-bucket rate limiting and sorted sets. This early release targets PostgreSQL 15–18 and is not production ready.

### Core Workflow

Add the library to the existing preload list and restart PostgreSQL before installing the extension as a superuser:

```ini
shared_preload_libraries = 'pg_kv'
```

```sql
CREATE EXTENSION pg_kv;
SELECT kv.set('app:session:42', '{"user":1}', interval '30 minutes');
SELECT kv.get('app:session:42');
SELECT kv.ttl('app:session:42');
SELECT kv.incr('app:views:7');
SELECT kv.rate_limit('app:api:42', 100, interval '1 minute');
SELECT kv.zadd('app:leaderboard', 'player:9', 1520);
SELECT * FROM kv.zrange('app:leaderboard', 0, 9, reverse => true);
SELECT * FROM kv.stats();
```

Installation fails if the library was not preloaded. Grant schema `kv` usage only to intended callers. The keyspace is shared across every database in the cluster; prefixes organize keys but do not enforce tenant isolation.

### Operations and Capacity

`kv.ttl` returns milliseconds, -1 for a permanent key and -2 for a missing key. `kv.expire` changes expiry, `kv.del` removes a key, and `kv.rate_limit_info` reports allowance and retry timing. Sorted sets also support `kv.zscore`, `kv.zrank`, `kv.zrevrank`, `kv.zincrby`, `kv.zcard` and `kv.zrem`.

Restart settings include `pg_kv.max_entries`, `pg_kv.max_key_bytes`, `pg_kv.max_value_bytes`, `pg_kv.shards`, `pg_kv.sweep_interval_ms`, `pg_kv.max_zset_members` and `pg_kv.max_member_bytes`. Storage is bounded; oversize values fail, and pressure can evict cached keys. `kv.flush` is an administrative operation unavailable to the public role by default.

### Durability Boundary

Cache changes survive transaction rollback but disappear on a server restart or crash. They are not WAL logged or replicated; a standby has its own cache. Use this for disposable state, not authoritative records or a globally coordinated limiter across independent servers.
