---
title: "pg_turbovec"
linkTitle: "pg_turbovec"
description: "TurboQuant-compressed vector type and ANN index access method for PostgreSQL."
weight: 1980
categories: ["RAG"]
languages: ["Rust"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_turbovec**](https://codeberg.org/gregburd/pg_turbovec) : TurboQuant-compressed vector type and ANN index access method for PostgreSQL.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1980** | {{< badge content="pg_turbovec" link="https://codeberg.org/gregburd/pg_turbovec" >}} | {{< ext "pg_turbovec" >}} | `2.14.1` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `turbovec` |
|   **See Also**    | {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} |

> [!Note] 2.14.1: recycle backends; REINDEX old IVF inserted into after VACUUM, or any index emptied by VACUUM with a writer open; PG19 experimental.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.14.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_turbovec` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.14.1` | {{< bg "18" "pg_turbovec_18" "green" >}} {{< bg "17" "pg_turbovec_17" "green" >}} {{< bg "16" "pg_turbovec_16" "green" >}} {{< bg "15" "pg_turbovec_15" "green" >}} {{< bg "14" "pg_turbovec_14" "green" >}} | `pg_turbovec_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.14.1` | {{< bg "18" "postgresql-18-pg-turbovec" "green" >}} {{< bg "17" "postgresql-17-pg-turbovec" "green" >}} {{< bg "16" "postgresql-16-pg-turbovec" "green" >}} {{< bg "15" "postgresql-15-pg-turbovec" "green" >}} {{< bg "14" "postgresql-14-pg-turbovec" "green" >}} | `postgresql-$v-pg-turbovec` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.14.1" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_18` | `2.14.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_18-2.14.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_18-2.14.1-1PGSTY.el8.x86_64.rpm) |
| `pg_turbovec_18` | `2.14.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.7 MiB | [pg_turbovec_18-2.14.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_18-2.14.1-1PGSTY.el8.aarch64.rpm) |
| `pg_turbovec_18` | `2.14.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_18-2.14.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_18-2.14.1-1PGSTY.el9.x86_64.rpm) |
| `pg_turbovec_18` | `2.14.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_18-2.14.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_18-2.14.1-1PGSTY.el9.aarch64.rpm) |
| `pg_turbovec_18` | `2.14.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_18-2.14.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_18-2.14.1-1PGSTY.el10.x86_64.rpm) |
| `pg_turbovec_18` | `2.14.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_18-2.14.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_18-2.14.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.5 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.0 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.0 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-turbovec` | `2.14.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-18-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_17` | `2.14.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_17-2.14.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_17-2.14.1-1PGSTY.el8.x86_64.rpm) |
| `pg_turbovec_17` | `2.14.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.7 MiB | [pg_turbovec_17-2.14.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_17-2.14.1-1PGSTY.el8.aarch64.rpm) |
| `pg_turbovec_17` | `2.14.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_17-2.14.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_17-2.14.1-1PGSTY.el9.x86_64.rpm) |
| `pg_turbovec_17` | `2.14.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_17-2.14.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_17-2.14.1-1PGSTY.el9.aarch64.rpm) |
| `pg_turbovec_17` | `2.14.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_17-2.14.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_17-2.14.1-1PGSTY.el10.x86_64.rpm) |
| `pg_turbovec_17` | `2.14.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_17-2.14.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_17-2.14.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.5 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.0 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.0 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-turbovec` | `2.14.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-17-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_16` | `2.14.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_16-2.14.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_16-2.14.1-1PGSTY.el8.x86_64.rpm) |
| `pg_turbovec_16` | `2.14.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.7 MiB | [pg_turbovec_16-2.14.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_16-2.14.1-1PGSTY.el8.aarch64.rpm) |
| `pg_turbovec_16` | `2.14.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_16-2.14.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_16-2.14.1-1PGSTY.el9.x86_64.rpm) |
| `pg_turbovec_16` | `2.14.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_16-2.14.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_16-2.14.1-1PGSTY.el9.aarch64.rpm) |
| `pg_turbovec_16` | `2.14.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_16-2.14.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_16-2.14.1-1PGSTY.el10.x86_64.rpm) |
| `pg_turbovec_16` | `2.14.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_16-2.14.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_16-2.14.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.5 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.0 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.9 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-turbovec` | `2.14.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-16-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_15` | `2.14.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_15-2.14.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_15-2.14.1-1PGSTY.el8.x86_64.rpm) |
| `pg_turbovec_15` | `2.14.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.7 MiB | [pg_turbovec_15-2.14.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_15-2.14.1-1PGSTY.el8.aarch64.rpm) |
| `pg_turbovec_15` | `2.14.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_15-2.14.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_15-2.14.1-1PGSTY.el9.x86_64.rpm) |
| `pg_turbovec_15` | `2.14.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_15-2.14.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_15-2.14.1-1PGSTY.el9.aarch64.rpm) |
| `pg_turbovec_15` | `2.14.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_15-2.14.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_15-2.14.1-1PGSTY.el10.x86_64.rpm) |
| `pg_turbovec_15` | `2.14.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_15-2.14.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_15-2.14.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.5 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-turbovec` | `2.14.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-15-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_14` | `2.14.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_14-2.14.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_14-2.14.1-1PGSTY.el8.x86_64.rpm) |
| `pg_turbovec_14` | `2.14.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.7 MiB | [pg_turbovec_14-2.14.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_14-2.14.1-1PGSTY.el8.aarch64.rpm) |
| `pg_turbovec_14` | `2.14.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_14-2.14.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_14-2.14.1-1PGSTY.el9.x86_64.rpm) |
| `pg_turbovec_14` | `2.14.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_14-2.14.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_14-2.14.1-1PGSTY.el9.aarch64.rpm) |
| `pg_turbovec_14` | `2.14.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_14-2.14.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_14-2.14.1-1PGSTY.el10.x86_64.rpm) |
| `pg_turbovec_14` | `2.14.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.8 MiB | [pg_turbovec_14-2.14.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_14-2.14.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.5 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.5 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.6 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.6 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-turbovec` | `2.14.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.6 MiB | [postgresql-14-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_2.14.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_turbovec" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_turbovec" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_turbovec-2.14.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_turbovec;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_turbovec;		# install via package name, for the active PG version

pig install pg_turbovec -v 18;   # install for PG 18
pig install pg_turbovec -v 17;   # install for PG 17
pig install pg_turbovec -v 16;   # install for PG 16
pig install pg_turbovec -v 15;   # install for PG 15
pig install pg_turbovec -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_turbovec;
```

## Usage

Sources:

- [Threads per query](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/docs/PRODUCTION.md)
- [v2.14.1 README](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/README.md)
- [v2.14.1 changelog](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/CHANGELOG.md)
- [Upgrade matrix](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/docs/UPGRADING.md)
- [Control file](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/pg_turbovec.control)
- [Filtering guide](https://api.pgxn.org/src/pg_turbovec/pg_turbovec-2.14.1/docs/FILTERING.md)

`pg_turbovec` 2.14.1 provides the `turbovec.vector` type and compact vector indexes with candidate reranking against the original heap vectors. Flat search scans quantized codes; IVF additionally searches selected cells. Both are approximate unless the candidate set contains every true neighbour. **Upgrading any 1.x index to 2.x requires rebuilding it.**

### Create and Query Vectors

```sql
CREATE EXTENSION pg_turbovec;
SET search_path = public, turbovec;

CREATE TABLE items (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  embedding turbovec.vector CHECK (turbovec.vector_dims(embedding) = 8)
);
INSERT INTO items (embedding) VALUES
  ('[1,2,3,4,5,6,7,8]'),
  ('[2,1,3,5,4,7,6,8]'),
  ('[8,7,6,5,4,3,2,1]');
CREATE INDEX items_embedding_idx ON items
USING turbovec (embedding turbovec.vec_cosine_ops)
WITH (bit_width = 4);

SELECT id, embedding <=> '[1,2,3,4,5,6,7,8]'::turbovec.vector AS distance
FROM items
ORDER BY embedding <=> '[1,2,3,4,5,6,7,8]'::turbovec.vector
LIMIT 3;
```

Indexed vectors must have a consistent dimension that is a multiple of 8; the type accepts up to 16,000 coordinates. The example uses eight-dimensional vectors so it is valid for indexing. Use `turbovec.vec_cosine_ops` with `<=>` and `turbovec.vec_ip_ops` with `<#>`; `<->` and `<+>` also provide exact distance operators.

### Choose an Index and Tune Queries

- `lists = 0` is the default flat quantized scan. Candidate reranking does not guarantee exact recall for every dataset.
- `WITH (lists = N)` enables IVF. Train on enough representative rows, choose the cell count for the corpus, and measure recall against an exact baseline; increasing the cell count alone does not ensure faster queries.
- `bit_width = 4` is the default. Two- and three-bit quantization are also available. `bit_width = 1` uses centered sign binary quantization, supports IVF, and reranks from the heap; it is a different scheme from TurboQuant.
- `WITH (graph = true)` is deprecated. Prefer flat or IVF for new indexes and consult the upstream migration guide for an existing graph index.

Indexes and queries can run without preload, but upstream recommends adding the library to `shared_preload_libraries` for its tuning GUCs. Merge it with existing entries and restart PostgreSQL; do not replace other required libraries:

```conf
shared_preload_libraries = 'pg_turbovec'
```

```sql
SELECT count(*) FROM pg_settings WHERE name LIKE 'turbovec.%';
SET turbovec.probes = 16;
SET turbovec.search_k = 64;
```

The settings query must return a nonzero count. A custom dotted parameter accepted by SET is not proof that the extension's GUC registered. Important controls are `turbovec.probes`, `turbovec.search_k`, `turbovec.oversample`, `turbovec.iterative_scan` and `turbovec.cache_size_mb`. Use partial indexes for stable filters or the documented allowlist/iterative-scan paths; compare filtered results with an exact baseline. Native partitions have separate indexes that must each be maintained.

### Upgrade from 1.x to 2.12.1

Schedule a maintenance window and retain the heap vectors plus a verified backup. Install the matching new binaries, update the SQL extension, and restart PostgreSQL so every backend uses the new library:

```sql
ALTER EXTENSION pg_turbovec UPDATE TO '2.12.1';
```

The index format changed from 7 to 8 in 2.0.0. Old indexes cannot be read in place: **rebuild every TurboVec index from the heap before resuming queries**, including every partition's index. In psql, generate and run the rebuild statements:

```sql
SELECT format('REINDEX INDEX %I.%I;', n.nspname, c.relname)
FROM pg_class c
JOIN pg_am a ON a.oid = c.relam
JOIN pg_namespace n ON n.oid = c.relnamespace
WHERE a.amname = 'turbovec';
\gexec
```

This is ordinary blocking REINDEX. If choosing concurrent rebuilding, issue each command as a separate top-level statement, outside a transaction block or DO function, and plan for old-format scans to fail until the rebuilt index is available. Do not treat the SQL extension update alone as a completed 1.x migration.

The 2.10.2-to-2.10.3 patch preserves format 8 and does not itself require reindexing. Earlier transitions still need the actions in the upstream upgrade matrix. An already corrupted index needs repair even when a patch preserves its format; `turbovec.turbovec_check(regclass)` reports detected corruption and a reason.

### Compatibility and Maintenance

The control file fixes objects in schema `turbovec`, sets `superuser = false`, and is not relocatable. Upstream covers PostgreSQL 13–18 and labels PostgreSQL 19 experimental in this release; current Pigsty packages cover 14–18. Binary replacement, tuning preload and major-format migration have separate restart/rebuild requirements. Version 2.10.3 parallelizes cold-backend code repacking; it does not change existing index bytes. Budget index-build memory, temporary space, WAL and vacuum work from the actual corpus, rather than treating upstream benchmark numbers as guarantees.

### 2.11.0 Upgrade

Version 2.11.0 adopts staged 2/4-bit search on eligible hardware and large indexes; the approximate candidate set can change. It also fixes concurrent-write index corruption and per-scan memory leaks present in 2.10.3 and earlier. When upgrading a healthy format-8 index from 2.x, install matching files, update the extension and restart: the release changes neither SQL surface nor wire format and does not require routine REINDEX. Existing corruption still requires repair; this does not change the mandatory rebuild for 1.x indexes.

### Distance Rounding in 2.12.0

Vectorized exact-distance kernels change rounding for vector and halfvec inputs with eight or more dimensions. Very close neighbors can change order or cross a distance threshold; sparsevec operators are unchanged. Format 8 and existing index bytes are unchanged: update the SQL extension and reconnect or restart to load the new library, without a routine index rebuild.

Expression indexes, partial-index predicates, stored generated columns and materialized views may retain older rounded results. Rebuild indexes, refresh views and recompute stored derived values if exact agreement with new calculations matters. Existing corruption still needs repair, and all 1.x indexes still require rebuilding. IVF indexes created empty or truncated do not train cells; load data and follow the upstream rebuild procedure.

### 2.12.1 Recall Fix

Version 2.12.1 updates turbovec to 1.1.2, correcting staged-search recall loss on corpora with a strong common vector direction. Candidate selection can change on eligible hardware; exact reranking scores are unchanged. It also fixes Nix builds and updates upstream pgrx to 0.19.3, including PostgreSQL 19beta4 CI. SQL objects, GUCs and format 8 are unchanged. Update the extension and reconnect to load the library; healthy 2.x indexes do not need REINDEX. The 1.x rebuild and existing-corruption requirements above still apply.

### Bounded Scan Threads in 2.13.0

`turbovec.scan_parallelism` now bounds flat, IVF, ColBERT ANN and 1-bit scans, `turbovec.knn()`, and first-open repacking. The deprecated graph path is unchanged. In 2.13.0, the default 0 meant up to four cores; reducing the setting takes effect in the session. Review this tradeoff when tuning concurrent throughput versus latency for a single query.

The inert `turbovec.search_concurrency` setting was removed. A successful SET of that dotted name can be an unregistered PostgreSQL custom parameter and does not mean it is active. Version 2.13.0 also fixes contention in the out-of-core IVF codebook cache. Install matching files, run `ALTER EXTENSION pg_turbovec UPDATE TO '2.13.0'`, and reconnect or restart. Healthy format-8 indexes do not require REINDEX; the earlier 1.x rebuild requirement still applies.

### Concurrency Defaults and Reranking in 2.14.0

`turbovec.scan_parallelism = 0` now defaults to one thread to reduce CPU oversubscription under concurrency. This does not remove the memory cost of private decoded caches in each backend. Explicitly increasing the thread count remains a low-concurrency latency option; measure the throughput tradeoff.

New `turbovec.rerank_multiplier` defaults to 16.0, with a range of 1.0–1024.0. The exact-rerank candidate window now scales with the requested count and oversample instead of using a dimension-based floor. Candidate selection can change, so validate recall. Install matching files, run `ALTER EXTENSION pg_turbovec UPDATE TO '2.14.0'`, and reconnect to load the library and new defaults. Format 8 is unchanged: healthy 2.x indexes need no routine REINDEX, while the earlier 1.x migration and corruption-repair requirements remain.

### 2.14.1 Write and VACUUM Correctness

Version 2.14.1 fixes stale backend caches after overlapping commits, loss of IVF tombstones during insert flushes, resurrection of rows when VACUUM empties an index while a writer is open, and tombstone handling in degraded IVF scans. Exact scores and format 8 are unchanged. Install matching files, run `ALTER EXTENSION pg_turbovec UPDATE TO '2.14.1'`, and recycle long-lived database connections so their libraries and cached copies are replaced.

Two histories require a one-time REINDEX: an IVF index that received inserts after VACUUM on an earlier release, and any index emptied by VACUUM while a writer transaction remained open. On a quiet table, `VACUUM (INDEX_CLEANUP ON)` followed by `turbovec.turbovec_check(regclass)` can help compare the estimated live index entries with the covered heap rows. Include the non-null vector condition and any partial-index predicate; excess index entries require REINDEX. Other healthy 2.x indexes need no routine rebuild. This does not remove the mandatory 1.x rebuild or repair of existing corruption.
