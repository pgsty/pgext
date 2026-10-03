---
title: "snowflake"
linkTitle: "snowflake"
description: "Snowflake-style 64-bit ID generator and sequence utilities for PostgreSQL"
weight: 4590
categories: ["FUNC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**snowflake**](https://github.com/pgEdge/snowflake) : Snowflake-style 64-bit ID generator and sequence utilities for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4590** | {{< badge content="snowflake" link="https://github.com/pgEdge/snowflake" >}} | {{< ext "snowflake" >}} | `2.6.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `snowflake` |
|   **See Also**    | {{< ext "pg_uuidv7" >}} {{< ext "sequential_uuids" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuid_v8" >}} {{< ext "uuid-ossp" >}} {{< ext "typeid" >}} {{< ext "permuteseq" >}} |

> [!Note] Bundled with pgEdge 15.19/16.15/17.11/18.6; set snowflake.node.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.6.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `snowflake` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.6` | {{< bg "18" "pgedge-18" "green" >}} {{< bg "17" "pgedge-17" "green" >}} {{< bg "16" "pgedge-16" "green" >}} {{< bg "15" "pgedge-15" "green" >}} {{< bg "14" "pgedge-14" "red" >}} | `pgedge-$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.6` | {{< bg "18" "pgedge-18" "green" >}} {{< bg "17" "pgedge-17" "green" >}} {{< bg "16" "pgedge-16" "green" >}} {{< bg "15" "pgedge-15" "green" >}} {{< bg "14" "pgedge-14" "red" >}} | `pgedge-$v` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-18` | `18.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.6 MiB | [pgedge-18-18.6-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-18-18.6-2PGSTY.el8.x86_64.rpm) |
| `pgedge-18` | `18.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.2 MiB | [pgedge-18-18.6-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-18-18.6-2PGSTY.el8.aarch64.rpm) |
| `pgedge-18` | `18.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.0 MiB | [pgedge-18-18.6-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-18-18.6-2PGSTY.el9.x86_64.rpm) |
| `pgedge-18` | `18.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.7 MiB | [pgedge-18-18.6-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-18-18.6-2PGSTY.el9.aarch64.rpm) |
| `pgedge-18` | `18.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.1 MiB | [pgedge-18-18.6-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-18-18.6-2PGSTY.el10.x86_64.rpm) |
| `pgedge-18` | `18.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.9 MiB | [pgedge-18-18.6-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-18-18.6-2PGSTY.el10.aarch64.rpm) |
| `pgedge-18` | `18.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.3 MiB | [pgedge-18_18.6-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~bookworm_amd64.deb) |
| `pgedge-18` | `18.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [pgedge-18_18.6-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~bookworm_arm64.deb) |
| `pgedge-18` | `18.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.3 MiB | [pgedge-18_18.6-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~trixie_amd64.deb) |
| `pgedge-18` | `18.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.8 MiB | [pgedge-18_18.6-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~trixie_arm64.deb) |
| `pgedge-18` | `18.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.6 MiB | [pgedge-18_18.6-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~jammy_amd64.deb) |
| `pgedge-18` | `18.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.4 MiB | [pgedge-18_18.6-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~jammy_arm64.deb) |
| `pgedge-18` | `18.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.4 MiB | [pgedge-18_18.6-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~noble_amd64.deb) |
| `pgedge-18` | `18.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 MiB | [pgedge-18_18.6-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~noble_arm64.deb) |
| `pgedge-18` | `18.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.5 MiB | [pgedge-18_18.6-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~resolute_amd64.deb) |
| `pgedge-18` | `18.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.3 MiB | [pgedge-18_18.6-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-17` | `17.11` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.2 MiB | [pgedge-17-17.11-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-17-17.11-2PGSTY.el8.x86_64.rpm) |
| `pgedge-17` | `17.11` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.8 MiB | [pgedge-17-17.11-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-17-17.11-2PGSTY.el8.aarch64.rpm) |
| `pgedge-17` | `17.11` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.7 MiB | [pgedge-17-17.11-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-17-17.11-2PGSTY.el9.x86_64.rpm) |
| `pgedge-17` | `17.11` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 MiB | [pgedge-17-17.11-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-17-17.11-2PGSTY.el9.aarch64.rpm) |
| `pgedge-17` | `17.11` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.8 MiB | [pgedge-17-17.11-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-17-17.11-2PGSTY.el10.x86_64.rpm) |
| `pgedge-17` | `17.11` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 MiB | [pgedge-17-17.11-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-17-17.11-2PGSTY.el10.aarch64.rpm) |
| `pgedge-17` | `17.11` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.0 MiB | [pgedge-17_17.11-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~bookworm_amd64.deb) |
| `pgedge-17` | `17.11` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.5 MiB | [pgedge-17_17.11-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~bookworm_arm64.deb) |
| `pgedge-17` | `17.11` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.0 MiB | [pgedge-17_17.11-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~trixie_amd64.deb) |
| `pgedge-17` | `17.11` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.5 MiB | [pgedge-17_17.11-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~trixie_arm64.deb) |
| `pgedge-17` | `17.11` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.3 MiB | [pgedge-17_17.11-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~jammy_amd64.deb) |
| `pgedge-17` | `17.11` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.1 MiB | [pgedge-17_17.11-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~jammy_arm64.deb) |
| `pgedge-17` | `17.11` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.2 MiB | [pgedge-17_17.11-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~noble_amd64.deb) |
| `pgedge-17` | `17.11` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.0 MiB | [pgedge-17_17.11-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~noble_arm64.deb) |
| `pgedge-17` | `17.11` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.2 MiB | [pgedge-17_17.11-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~resolute_amd64.deb) |
| `pgedge-17` | `17.11` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.9 MiB | [pgedge-17_17.11-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-16` | `16.15` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.5 MiB | [pgedge-16-16.15-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-16-16.15-2PGSTY.el8.x86_64.rpm) |
| `pgedge-16` | `16.15` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.1 MiB | [pgedge-16-16.15-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-16-16.15-2PGSTY.el8.aarch64.rpm) |
| `pgedge-16` | `16.15` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.2 MiB | [pgedge-16-16.15-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-16-16.15-2PGSTY.el9.x86_64.rpm) |
| `pgedge-16` | `16.15` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.9 MiB | [pgedge-16-16.15-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-16-16.15-2PGSTY.el9.aarch64.rpm) |
| `pgedge-16` | `16.15` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.3 MiB | [pgedge-16-16.15-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-16-16.15-2PGSTY.el10.x86_64.rpm) |
| `pgedge-16` | `16.15` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.1 MiB | [pgedge-16-16.15-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-16-16.15-2PGSTY.el10.aarch64.rpm) |
| `pgedge-16` | `16.15` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.5 MiB | [pgedge-16_16.15-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~bookworm_amd64.deb) |
| `pgedge-16` | `16.15` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.0 MiB | [pgedge-16_16.15-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~bookworm_arm64.deb) |
| `pgedge-16` | `16.15` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.5 MiB | [pgedge-16_16.15-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~trixie_amd64.deb) |
| `pgedge-16` | `16.15` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.1 MiB | [pgedge-16_16.15-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~trixie_arm64.deb) |
| `pgedge-16` | `16.15` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.8 MiB | [pgedge-16_16.15-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~jammy_amd64.deb) |
| `pgedge-16` | `16.15` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.6 MiB | [pgedge-16_16.15-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~jammy_arm64.deb) |
| `pgedge-16` | `16.15` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.7 MiB | [pgedge-16_16.15-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~noble_amd64.deb) |
| `pgedge-16` | `16.15` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.5 MiB | [pgedge-16_16.15-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~noble_arm64.deb) |
| `pgedge-16` | `16.15` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.7 MiB | [pgedge-16_16.15-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~resolute_amd64.deb) |
| `pgedge-16` | `16.15` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.4 MiB | [pgedge-16_16.15-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-15` | `15.19` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 MiB | [pgedge-15-15.19-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-15-15.19-2PGSTY.el8.x86_64.rpm) |
| `pgedge-15` | `15.19` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.9 MiB | [pgedge-15-15.19-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-15-15.19-2PGSTY.el8.aarch64.rpm) |
| `pgedge-15` | `15.19` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.2 MiB | [pgedge-15-15.19-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-15-15.19-2PGSTY.el9.x86_64.rpm) |
| `pgedge-15` | `15.19` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.0 MiB | [pgedge-15-15.19-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-15-15.19-2PGSTY.el9.aarch64.rpm) |
| `pgedge-15` | `15.19` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.3 MiB | [pgedge-15-15.19-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-15-15.19-2PGSTY.el10.x86_64.rpm) |
| `pgedge-15` | `15.19` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.1 MiB | [pgedge-15-15.19-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-15-15.19-2PGSTY.el10.aarch64.rpm) |
| `pgedge-15` | `15.19` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.5 MiB | [pgedge-15_15.19-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~bookworm_amd64.deb) |
| `pgedge-15` | `15.19` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.2 MiB | [pgedge-15_15.19-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~bookworm_arm64.deb) |
| `pgedge-15` | `15.19` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.6 MiB | [pgedge-15_15.19-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~trixie_amd64.deb) |
| `pgedge-15` | `15.19` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.2 MiB | [pgedge-15_15.19-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~trixie_arm64.deb) |
| `pgedge-15` | `15.19` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.9 MiB | [pgedge-15_15.19-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~jammy_amd64.deb) |
| `pgedge-15` | `15.19` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.7 MiB | [pgedge-15_15.19-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~jammy_arm64.deb) |
| `pgedge-15` | `15.19` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.8 MiB | [pgedge-15_15.19-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~noble_amd64.deb) |
| `pgedge-15` | `15.19` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.6 MiB | [pgedge-15_15.19-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~noble_arm64.deb) |
| `pgedge-15` | `15.19` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.8 MiB | [pgedge-15_15.19-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~resolute_amd64.deb) |
| `pgedge-15` | `15.19` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.6 MiB | [pgedge-15_15.19-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgEdge/snowflake" title="Repository" icon="github" subtitle="github.com/pgEdge/snowflake" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="snowflake-2.6.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg snowflake;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install snowflake;		# install via package name, for the active PG version

pig install snowflake -v 18;   # install for PG 18
pig install snowflake -v 17;   # install for PG 17
pig install snowflake -v 16;   # install for PG 16
pig install snowflake -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION snowflake;
```




## Usage

Sources:

- [snowflake v2.6.0 README](https://github.com/pgEdge/snowflake/blob/v2.6.0/README.md)
- [Creating a Snowflake sequence](https://github.com/pgEdge/snowflake/blob/v2.6.0/docs/creating.md)
- [Converting PostgreSQL sequences](https://github.com/pgEdge/snowflake/blob/v2.6.0/docs/converting.md)
- [Function reference](https://github.com/pgEdge/snowflake/blob/v2.6.0/docs/snowflake_functions.md)
- [v2.6.0 changelog](https://github.com/pgEdge/snowflake/blob/v2.6.0/docs/changelog.md)
- [v2.6.0 C implementation](https://github.com/pgEdge/snowflake/blob/v2.6.0/snowflake.c)
- [Tagged SQL API definitions](https://github.com/pgEdge/snowflake/blob/v2.6.0/snowflake--2.3.sql)
- [v2.6.0 control file](https://github.com/pgEdge/snowflake/blob/v2.6.0/snowflake.control)

`snowflake` generates distributed `bigint` identifiers from a timestamp, a per-node identifier, and an in-millisecond counter. Existing PostgreSQL sequences can be converted so table defaults continue using `nextval(...)` while producing Snowflake IDs.

```sql
CREATE EXTENSION snowflake;
```

### Configuration

Assign every writable node a distinct identifier in `postgresql.conf`, then reload PostgreSQL:

```ini
snowflake.node = 1
```

Reusing a node identifier on concurrently writable servers can generate duplicate IDs.

### Convert a Sequence

Create a normal PostgreSQL sequence, then convert its definition. Existing values in referencing columns are not rewritten.

```sql
CREATE TABLE orders (
  id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  payload jsonb NOT NULL
);

SELECT snowflake.convert_sequence_to_snowflake(
  pg_get_serial_sequence('orders', 'id')::regclass
);

INSERT INTO orders (payload) VALUES ('{"status":"new"}');
SELECT id, snowflake.format(id) FROM orders;
```

### Functions

| Function | Description |
|---|---|
| `snowflake.nextval([sequence regclass])` | Generate the next Snowflake ID (uses internal sequence if none specified) |
| `snowflake.currval([sequence regclass])` | Return the current value of the sequence |
| `snowflake.get_epoch(snowflake int8)` | Return Unix epoch seconds since 1970-01-01, with millisecond precision |
| `snowflake.get_count(snowflake int8)` | Extract the count part (resets per millisecond) |
| `snowflake.get_node(snowflake int8)` | Extract the node identifier |
| `snowflake.format(snowflake int8)` | Return JSONB fields `id` (node identifier), `ts`, and `count` |

### Examples

```sql
-- Generate a snowflake ID
SELECT snowflake.nextval();

-- Use an already converted named sequence
SELECT snowflake.nextval('orders_id_seq'::regclass);

-- Decode a fixed example: node 1, count 0
SET TIME ZONE 'UTC';
SELECT snowflake.get_epoch(136169504773246976);
-- 1704996539.845

SELECT to_timestamp(snowflake.get_epoch(136169504773246976));
-- 2024-01-11 18:08:59.845+00

SELECT snowflake.get_node(136169504773246976);
-- 1

SELECT snowflake.get_count(136169504773246976);
-- 0

SELECT snowflake.format(136169504773246976);
-- {"id": 1, "ts": "2024-01-11 18:08:59.845+00", "count": 0}

-- Use as default column
CREATE TABLE direct_ids (
  id int8 DEFAULT snowflake.nextval() PRIMARY KEY,
  data text
);
```

### Review and Upgrade

Review converted table defaults with `psql \d+` or the PostgreSQL catalogs. They should call `snowflake.nextval(...)` rather than the original `nextval(...)`:

```sql
SELECT table_schema, table_name, column_name, column_default
FROM information_schema.columns
WHERE column_default LIKE 'snowflake.nextval(%';
```

Version `2.5.0` fixes dump/restore of converted sequences whose `MAXVALUE` was left at the normal `bigint` maximum. It also repairs conversion SQL for affected sequences and adds PostgreSQL 18 support.

### Caveats

- Conversion changes the sequence definition, not IDs already stored in table rows.
- A Snowflake generator can emit at most 4096 counter values per millisecond. Do not configure a sequence increment above 4096.
- Keep the node identifier stable and unique for the lifetime of concurrent writers; record it as part of cluster provisioning and failover procedures.
- Install the same extension version on every node before logical replication or rolling changes involving converted sequence definitions.

Version 2.6.0 adds PostgreSQL 19 support; the current Pigsty pgEdge bundle covers PostgreSQL 15–18. Assign `snowflake.node` within 1–1023; its default is intentionally invalid. The control file fixes schema `snowflake` and does not require shared preload. The 2.5.0 dump/restore repair remains relevant when upgrading converted sequences from earlier releases.

Use an administrator to install the C extension and convert identity definitions. ID generation requires USAGE or UPDATE on the selected sequence; the current value is defined only after that sequence has generated an ID in the same session.
