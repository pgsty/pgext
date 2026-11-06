---
title: "pg_lion"
linkTitle: "pg_lion"
description: "Roaring bitmap inverted indexes for filtering, counts and aggregates"
weight: 2715
categories: ["FEAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_lion**](https://github.com/cgkanchi/pg_lion) : Roaring bitmap inverted indexes for filtering, counts and aggregates


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2715** | {{< badge content="pg_lion" link="https://github.com/cgkanchi/pg_lion" >}} | {{< ext "pg_lion" >}} | `0.3.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_lion_citext" >}} {{< ext "pg_lion_vector" >}} |
|    **Siblings**   | {{< ext "pg_lion_citext" >}} {{< ext "pg_lion_vector" >}} |

> [!Note] Custom WAL requires preload.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_lion` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "pg_lion_18" "green" >}} {{< bg "17" "pg_lion_17" "green" >}} {{< bg "16" "pg_lion_16" "green" >}} {{< bg "15" "pg_lion_15" "red" >}} {{< bg "14" "pg_lion_14" "red" >}} | `pg_lion_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-pg-lion" "green" >}} {{< bg "17" "postgresql-17-pg-lion" "green" >}} {{< bg "16" "postgresql-16-pg-lion" "green" >}} {{< bg "15" "postgresql-15-pg-lion" "red" >}} {{< bg "14" "postgresql-14-pg-lion" "red" >}} | `postgresql-$v-pg-lion` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_lion_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_lion_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lion_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-lion : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-lion : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-lion : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-lion : N/A 0" "gray" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lion_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [pg_lion_18-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lion_18-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_lion_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.6 MiB | [pg_lion_18-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lion_18-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_lion_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [pg_lion_18-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lion_18-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_lion_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [pg_lion_18-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lion_18-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_lion_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.8 MiB | [pg_lion_18-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lion_18-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_lion_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [pg_lion_18-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lion_18-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-lion` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.8 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.7 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-lion` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.6 MiB | [postgresql-18-pg-lion_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lion/postgresql-18-pg-lion_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lion_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.8 MiB | [pg_lion_17-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lion_17-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_lion_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.6 MiB | [pg_lion_17-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lion_17-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_lion_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [pg_lion_17-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lion_17-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_lion_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [pg_lion_17-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lion_17-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_lion_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.8 MiB | [pg_lion_17-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lion_17-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_lion_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.7 MiB | [pg_lion_17-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lion_17-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-lion` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.7 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-lion` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.6 MiB | [postgresql-17-pg-lion_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lion/postgresql-17-pg-lion_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lion_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.7 MiB | [pg_lion_16-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lion_16-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_lion_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.6 MiB | [pg_lion_16-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lion_16-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_lion_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.7 MiB | [pg_lion_16-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lion_16-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_lion_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.6 MiB | [pg_lion_16-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lion_16-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_lion_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.7 MiB | [pg_lion_16-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lion_16-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_lion_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.6 MiB | [pg_lion_16-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lion_16-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-lion` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.6 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.6 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.7 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-lion` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.6 MiB | [postgresql-16-pg-lion_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-lion/postgresql-16-pg-lion_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cgkanchi/pg_lion" title="Repository" icon="github" subtitle="github.com/cgkanchi/pg_lion" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_lion-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_lion;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_lion;		# install via package name, for the active PG version

pig install pg_lion -v 18;   # install for PG 18
pig install pg_lion -v 17;   # install for PG 17
pig install pg_lion -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_lion;
```

## Usage

Sources:

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [Reference](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [Control file](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion.control)

`pg_lion` 0.3.0 adds roaring bitmap inverted indexes for filtering, counts and grouped aggregates on PostgreSQL 16 and later. It works on heap tables and uses ordinary PostgreSQL operators; it does not require preloading for its default generic WAL mode.

### Core Workflow

Create `lion` indexes on the columns used by filters and groups. Vacuumed, all-visible pages let counts avoid heap visits; recently changed pages may need visibility checks in the heap.

```sql
CREATE EXTENSION pg_lion;
CREATE TABLE lion_events (country text, amount integer);
INSERT INTO lion_events VALUES ('JP', 10), ('JP', 20), ('NZ', 30);
CREATE INDEX lion_country ON lion_events USING lion (country);
CREATE INDEX lion_amount ON lion_events USING lion (amount) WITH (slices = on);
VACUUM ANALYZE lion_events;
SELECT count(*), sum(amount) FROM lion_events WHERE country = 'JP';
SELECT country, count(*) FROM lion_events GROUP BY country;
```

Scalar equality, ranges, `IN`, null tests and combined predicates are supported. `lion` also indexes arrays, `jsonb` and `tsvector`; `lion_bm25_score` ranks full-text matches. An expression index on a JSON field can support grouping that a default `jsonb` index cannot.

### Options and Diagnostics

| Object | Purpose |
| --- | --- |
| `summaries = auto` | Pre-merge bitmaps to accelerate ranges with many distinct values. |
| `slices = on` | Store bit slices for filtered aggregates on integer, date and time values; adds write and storage cost. |
| `lion_index_count(index, key)` | Count a key through a named index. |
| `lion_index_stats(index)` | Inspect index statistics; initially granted to `pg_stat_scan_tables`. |
| `lion_index_verify(index, true)` | Check structure and heap coverage; restricted diagnostic function. |
| `lion_index_wal_mode(index)` | Report `generic` or `rmgr`; restricted diagnostic function. |

`pg_lion_citext` adds the `citext` operator class. `pg_lion_vector` adds the `lionvec` access method and requires pgvector 0.7 or later. Both companions ship with the parent distribution and are enabled separately.

### WAL and Operational Boundaries

Creating the extensions requires a superuser. Optional `shared_preload_libraries = 'pg_lion'` plus a server restart enables the custom WAL resource manager (default ID 170). Indexes built in that mode require the preload for writes and `VACUUM` until rebuilt; check `lion_index_wal_mode` before changing the setting.

Keep the library preloaded wherever old custom WAL may still be replayed, including standbys and archive recovery. `REINDEX` does not rewrite historical WAL. Before removing the preload, rebuild every `rmgr` index with `wal_mode = generic` while the library is still preloaded, and account for WAL replay consumers.

There are no unique indexes or `INCLUDE` columns, and only heap tables are supported. Keep B-tree indexes for uniqueness and ordinary ordered scans. Loading the library installs planner hooks and can affect estimates for other queries. Measure write-heavy workloads before adding indexes or slices.

The on-disk format may change before 1.0; follow release-specific `REINDEX` requirements. Builds predating 0.3.0 have no extension update path: upstream requires recreating the extensions and their indexes. Version 0.3.0 supplies installation scripts only.
