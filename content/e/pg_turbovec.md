---
title: "pg_turbovec"
linkTitle: "pg_turbovec"
description: "TurboQuant-compressed vector type and ANN index access method for PostgreSQL."
weight: 1980
categories: ["RAG"]
width: full
---

[**pg_turbovec**](https://codeberg.org/gregburd/pg_turbovec) : TurboQuant-compressed vector type and ANN index access method for PostgreSQL.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1980** | {{< badge content="pg_turbovec" link="https://codeberg.org/gregburd/pg_turbovec" >}} | {{< ext "pg_turbovec" >}} | `1.29.0` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `turbovec` |
|   **See Also**    | {{< ext "pg_turboquant" >}} {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "vectorscale" >}} |

> [!Note] Upstream v1.29.0 supports PostgreSQL 13-18 with PG19 beta experimental; PIGSTY RPM, DEB, and source remain at 1.28.3 for PostgreSQL 14-18 with pgrx 0.19.1 and OpenBLAS. Upstream 1.28.4 fixes a persisted row-count corruption bug present in 1.28.3.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.29.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_turbovec` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.28.3` | {{< bg "18" "pg_turbovec_18" "green" >}} {{< bg "17" "pg_turbovec_17" "green" >}} {{< bg "16" "pg_turbovec_16" "green" >}} {{< bg "15" "pg_turbovec_15" "green" >}} {{< bg "14" "pg_turbovec_14" "green" >}} | `pg_turbovec_$v` | `openblas` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.28.3` | {{< bg "18" "postgresql-18-pg-turbovec" "green" >}} {{< bg "17" "postgresql-17-pg-turbovec" "green" >}} {{< bg "16" "postgresql-16-pg-turbovec" "green" >}} {{< bg "15" "postgresql-15-pg-turbovec" "green" >}} {{< bg "14" "postgresql-14-pg-turbovec" "green" >}} | `postgresql-$v-pg-turbovec` | `libopenblas0` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "pg_turbovec_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-18-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-17-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-16-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-15-pg-turbovec : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.28.3" "postgresql-14-pg-turbovec : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_18` | `1.28.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_18-1.28.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_18-1.28.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_turbovec_18` | `1.28.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [pg_turbovec_18-1.28.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_18-1.28.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_turbovec_18` | `1.28.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_18-1.28.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_18-1.28.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_turbovec_18` | `1.28.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_18-1.28.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_18-1.28.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_turbovec_18` | `1.28.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_18-1.28.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_18-1.28.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_turbovec_18` | `1.28.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_18-1.28.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_18-1.28.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.9 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-turbovec` | `1.28.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.9 MiB | [postgresql-18-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-18-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_17` | `1.28.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_17-1.28.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_17-1.28.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_turbovec_17` | `1.28.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [pg_turbovec_17-1.28.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_17-1.28.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_turbovec_17` | `1.28.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_17-1.28.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_17-1.28.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_turbovec_17` | `1.28.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_17-1.28.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_17-1.28.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_turbovec_17` | `1.28.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_17-1.28.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_17-1.28.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_turbovec_17` | `1.28.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_17-1.28.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_17-1.28.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.9 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-turbovec` | `1.28.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.9 MiB | [postgresql-17-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-17-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_16` | `1.28.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_16-1.28.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_16-1.28.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_turbovec_16` | `1.28.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [pg_turbovec_16-1.28.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_16-1.28.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_turbovec_16` | `1.28.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_16-1.28.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_16-1.28.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_turbovec_16` | `1.28.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_16-1.28.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_16-1.28.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_turbovec_16` | `1.28.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_16-1.28.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_16-1.28.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_turbovec_16` | `1.28.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_16-1.28.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_16-1.28.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.0 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.9 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-turbovec` | `1.28.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.9 MiB | [postgresql-16-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-16-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_15` | `1.28.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_15-1.28.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_15-1.28.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_turbovec_15` | `1.28.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [pg_turbovec_15-1.28.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_15-1.28.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_turbovec_15` | `1.28.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_15-1.28.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_15-1.28.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_turbovec_15` | `1.28.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_15-1.28.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_15-1.28.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_turbovec_15` | `1.28.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_15-1.28.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_15-1.28.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_turbovec_15` | `1.28.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_15-1.28.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_15-1.28.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-turbovec` | `1.28.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.9 MiB | [postgresql-15-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-15-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_turbovec_14` | `1.28.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_14-1.28.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_turbovec_14-1.28.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_turbovec_14` | `1.28.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.0 MiB | [pg_turbovec_14-1.28.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_turbovec_14-1.28.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_turbovec_14` | `1.28.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_14-1.28.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_turbovec_14-1.28.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_turbovec_14` | `1.28.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_14-1.28.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_turbovec_14-1.28.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_turbovec_14` | `1.28.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [pg_turbovec_14-1.28.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_turbovec_14-1.28.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_turbovec_14` | `1.28.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.1 MiB | [pg_turbovec_14-1.28.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_turbovec_14-1.28.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-turbovec` | `1.28.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.9 MiB | [postgresql-14-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-turbovec/postgresql-14-pg-turbovec_1.28.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_turbovec" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_turbovec" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_turbovec-1.28.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_turbovec;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_turbovec;		# install via package name, for the active PG version

pig install pg_turbovec -v 18;   # install for PG 18
pig install pg_turbovec -v 17;   # install for PG 17
pig install pg_turbovec -v 16;   # install for PG 16
pig install pg_turbovec -v 15;   # install for PG 15
pig install pg_turbovec -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_turbovec;
```
