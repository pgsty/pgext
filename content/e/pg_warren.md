---
title: "pg_warren"
linkTitle: "pg_warren"
description: "Versioned graph traversal and computation over PostgreSQL tables"
weight: 2645
categories: ["FEAT"]
languages: ["Rust"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_warren**](https://github.com/ExaDev/pg_warren) : Versioned graph traversal and computation over PostgreSQL tables


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2645** | {{< badge content="pg_warren" link="https://github.com/ExaDev/pg_warren" >}} | {{< ext "pg_warren" >}} | `0.8.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `warren` |

> [!Note] Shared cache preload on PG15–16


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_warren` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.0` | {{< bg "18" "pg_warren_18" "green" >}} {{< bg "17" "pg_warren_17" "green" >}} {{< bg "16" "pg_warren_16" "green" >}} {{< bg "15" "pg_warren_15" "green" >}} {{< bg "14" "pg_warren_14" "red" >}} | `pg_warren_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.0` | {{< bg "18" "postgresql-18-pg-warren" "green" >}} {{< bg "17" "postgresql-17-pg-warren" "green" >}} {{< bg "16" "postgresql-16-pg-warren" "green" >}} {{< bg "15" "postgresql-15-pg-warren" "green" >}} {{< bg "14" "postgresql-14-pg-warren" "red" >}} | `postgresql-$v-pg-warren` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_warren_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_warren_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_warren_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_warren_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_warren_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "pg_warren_15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_warren_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-18-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-17-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-16-pg-warren : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.0" "postgresql-15-pg-warren : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-14-pg-warren : N/A 0" "gray" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_warren_18` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_warren_18-0.8.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_warren_18-0.8.0-1PGSTY.el8.x86_64.rpm) |
| `pg_warren_18` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [pg_warren_18-0.8.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_warren_18-0.8.0-1PGSTY.el8.aarch64.rpm) |
| `pg_warren_18` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_warren_18-0.8.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_warren_18-0.8.0-1PGSTY.el9.x86_64.rpm) |
| `pg_warren_18` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [pg_warren_18-0.8.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_warren_18-0.8.0-1PGSTY.el9.aarch64.rpm) |
| `pg_warren_18` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_warren_18-0.8.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_warren_18-0.8.0-1PGSTY.el10.x86_64.rpm) |
| `pg_warren_18` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.6 MiB | [pg_warren_18-0.8.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_warren_18-0.8.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-warren` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.1 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.1 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.5 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.4 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-warren` | `0.8.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.4 MiB | [postgresql-18-pg-warren_0.8.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-warren/postgresql-18-pg-warren_0.8.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_warren_17` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_warren_17-0.8.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_warren_17-0.8.0-1PGSTY.el8.x86_64.rpm) |
| `pg_warren_17` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [pg_warren_17-0.8.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_warren_17-0.8.0-1PGSTY.el8.aarch64.rpm) |
| `pg_warren_17` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_warren_17-0.8.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_warren_17-0.8.0-1PGSTY.el9.x86_64.rpm) |
| `pg_warren_17` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [pg_warren_17-0.8.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_warren_17-0.8.0-1PGSTY.el9.aarch64.rpm) |
| `pg_warren_17` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_warren_17-0.8.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_warren_17-0.8.0-1PGSTY.el10.x86_64.rpm) |
| `pg_warren_17` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.6 MiB | [pg_warren_17-0.8.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_warren_17-0.8.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-warren` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.1 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.1 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.5 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.4 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-warren` | `0.8.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.4 MiB | [postgresql-17-pg-warren_0.8.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-warren/postgresql-17-pg-warren_0.8.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_warren_16` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_warren_16-0.8.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_warren_16-0.8.0-1PGSTY.el8.x86_64.rpm) |
| `pg_warren_16` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [pg_warren_16-0.8.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_warren_16-0.8.0-1PGSTY.el8.aarch64.rpm) |
| `pg_warren_16` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_warren_16-0.8.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_warren_16-0.8.0-1PGSTY.el9.x86_64.rpm) |
| `pg_warren_16` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [pg_warren_16-0.8.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_warren_16-0.8.0-1PGSTY.el9.aarch64.rpm) |
| `pg_warren_16` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_warren_16-0.8.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_warren_16-0.8.0-1PGSTY.el10.x86_64.rpm) |
| `pg_warren_16` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.6 MiB | [pg_warren_16-0.8.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_warren_16-0.8.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-warren` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.1 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.1 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.5 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.4 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-warren` | `0.8.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.4 MiB | [postgresql-16-pg-warren_0.8.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-warren/postgresql-16-pg-warren_0.8.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_warren_15` | `0.8.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.0 MiB | [pg_warren_15-0.8.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_warren_15-0.8.0-1PGSTY.el8.x86_64.rpm) |
| `pg_warren_15` | `0.8.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.4 MiB | [pg_warren_15-0.8.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_warren_15-0.8.0-1PGSTY.el8.aarch64.rpm) |
| `pg_warren_15` | `0.8.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.9 MiB | [pg_warren_15-0.8.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_warren_15-0.8.0-1PGSTY.el9.x86_64.rpm) |
| `pg_warren_15` | `0.8.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [pg_warren_15-0.8.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_warren_15-0.8.0-1PGSTY.el9.aarch64.rpm) |
| `pg_warren_15` | `0.8.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.9 MiB | [pg_warren_15-0.8.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_warren_15-0.8.0-1PGSTY.el10.x86_64.rpm) |
| `pg_warren_15` | `0.8.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.6 MiB | [pg_warren_15-0.8.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_warren_15-0.8.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-warren` | `0.8.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.1 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.1 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.5 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.4 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.8 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-warren` | `0.8.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.4 MiB | [postgresql-15-pg-warren_0.8.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-warren/postgresql-15-pg-warren_0.8.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ExaDev/pg_warren" title="Repository" icon="github" subtitle="github.com/ExaDev/pg_warren" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_warren-0.8.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_warren;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_warren;		# install via package name, for the active PG version

pig install pg_warren -v 18;   # install for PG 18
pig install pg_warren -v 17;   # install for PG 17
pig install pg_warren -v 16;   # install for PG 16
pig install pg_warren -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_warren;
```

## Usage

Sources:

- [README](https://github.com/ExaDev/pg_warren/blob/v0.8.0/README.md)
- [Control file](https://github.com/ExaDev/pg_warren/blob/v0.8.0/crates/pg_warren/pg_warren.control)
- [SQL API](https://github.com/ExaDev/pg_warren/blob/v0.8.0/sql/pg_warren--0.8.0.sql)
- [Limits and shared cache](https://github.com/ExaDev/pg_warren/blob/v0.8.0/docs/limits.md)

`pg_warren` 0.8.0 builds versioned adjacency snapshots from ordinary PostgreSQL tables for graph traversal and computation. It supports PostgreSQL 15–18 and keeps graph versions, vectors and their changes within PostgreSQL transactions. Traversal reads a built snapshot; source-table changes require an explicit refresh.

### Core Workflow

This is a trusted extension: a role with `CREATE` privilege on the database can install it. A basic graph needs node and edge queries:

```sql
CREATE EXTENSION pg_warren;
CREATE TABLE people (id bigint PRIMARY KEY, name text);
CREATE TABLE follows (src bigint, dst bigint);
INSERT INTO people VALUES (1, 'Ada'), (2, 'Bob'), (3, 'Cy');
INSERT INTO follows VALUES (1, 2), (2, 3);
SELECT warren.create_graph('social',
  'SELECT id FROM public.people', 'SELECT src, dst FROM public.follows');
SELECT warren.refresh('social');
SELECT * FROM warren.bfs('social', ARRAY[1::bigint], 3);
SELECT warren.pagerank('social', 'rank');
SELECT * FROM warren.vector_topk('social', 'rank', 3, NULL::bigint);
```

The graph stores its source queries. Run `warren.refresh` after changing the source tables; use `warren.info` to inspect versions and `warren.diff` to compare them. Building and promoting versions rolls back with the surrounding transaction.

### Traversal and Computation

| Interface | Purpose |
| --- | --- |
| `warren.bfs`, `warren.shortest_path` | Breadth-first traversal and shortest paths |
| `warren.pagerank`, `warren.components` | Rank nodes and find connected components |
| `warren.leiden` | Detect communities |
| `warren.spmv`, `warren.gather` | Compose graph calculations with SQL |
| `warren.vector_topk` | Read the largest values from a stored vector |
| `warren.shared_cache`, `warren.backends` | Inspect shared entries and backend cache usage |

Identifier types must agree with the graph. Functions without a node identifier argument may require a type witness such as `NULL::bigint`. In 0.8.0, PageRank and Leiden use native kernels; they write their results to named vectors.

### Memory and Preload

PostgreSQL 17–18 initialize the shared cache on first use. PostgreSQL 15–16 need the following setting and a server restart to share graph data between backends:

```ini
shared_preload_libraries = 'pg_warren'
```

Without preload on PostgreSQL 15–16, graph operations still work using private backend caches. Add the library to any existing preload list.

`warren.shared_cache_memory` limits the cluster-wide shared cache and takes effect on configuration reload. `warren.max_cache_memory`, `warren.max_compute_memory` and `warren.max_build_memory` limit distinct private-cache, computation and build allocations; they are not one combined memory budget. `warren.threads` controls computation parallelism. On Linux, shared entries also need sufficient `/dev/shm`; entries that do not fit fall back to private caches.

### Access, Backups and Upgrades

Source queries run as the graph owner with the stored `search_path`. Use schema-qualified source objects and controlled schemas. Public functions enforce graph access; installation does not grant direct access to private catalog tables.

Dump as a superuser to include graph catalogs, versions and vectors, and restore the required owner/reader roles first. Rebuilding from source tables does not preserve historical graph versions. After installing newer library files, run `ALTER EXTENSION pg_warren UPDATE` in each database using the extension. Dropping it removes stored graphs. The project remains pre-1.0 and does not maintain graph snapshots incrementally.
