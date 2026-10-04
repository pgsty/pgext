---
title: "edtf_postgres"
linkTitle: "edtf_postgres"
description: "EDTF validation, normalization, date bounds, and temporal relations"
weight: 1130
categories: ["TIME"]
languages: ["Rust"]
licenses: ["MIT OR Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**edtf_postgres**](https://github.com/monumental-archive/edtf/tree/edtf-postgres-v1.2.3/crates/edtf-postgres) : EDTF validation, normalization, date bounds, and temporal relations


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1130** | {{< badge content="edtf_postgres" link="https://github.com/monumental-archive/edtf/tree/edtf-postgres-v1.2.3/crates/edtf-postgres" >}} | {{< ext "edtf_postgres" >}} | `1.2.3` | {{< category "TIME" >}} | {{< license "MIT OR Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_when" >}} {{< ext "pgcalendar" >}} {{< ext "pg_rrule" >}} {{< ext "pg_duration" >}} |

> [!Note] Trusted extension; upstream Cargo requests pgrx 0.19.1 and the locked build uses 0.19.2.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `edtf_postgres` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "edtf_postgres_18" "green" >}} {{< bg "17" "edtf_postgres_17" "green" >}} {{< bg "16" "edtf_postgres_16" "green" >}} {{< bg "15" "edtf_postgres_15" "green" >}} {{< bg "14" "edtf_postgres_14" "green" >}} | `edtf_postgres_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.3` | {{< bg "18" "postgresql-18-edtf-postgres" "green" >}} {{< bg "17" "postgresql-17-edtf-postgres" "green" >}} {{< bg "16" "postgresql-16-edtf-postgres" "green" >}} {{< bg "15" "postgresql-15-edtf-postgres" "green" >}} {{< bg "14" "postgresql-14-edtf-postgres" "green" >}} | `postgresql-$v-edtf-postgres` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "edtf_postgres_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-18-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-17-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-16-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-15-edtf-postgres : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.3" "postgresql-14-edtf-postgres : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `edtf_postgres_18` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_18-1.2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/edtf_postgres_18-1.2.3-1PGSTY.el8.x86_64.rpm) |
| `edtf_postgres_18` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [edtf_postgres_18-1.2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/edtf_postgres_18-1.2.3-1PGSTY.el8.aarch64.rpm) |
| `edtf_postgres_18` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.6 MiB | [edtf_postgres_18-1.2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/edtf_postgres_18-1.2.3-1PGSTY.el9.x86_64.rpm) |
| `edtf_postgres_18` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_18-1.2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/edtf_postgres_18-1.2.3-1PGSTY.el9.aarch64.rpm) |
| `edtf_postgres_18` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [edtf_postgres_18-1.2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/edtf_postgres_18-1.2.3-1PGSTY.el10.x86_64.rpm) |
| `edtf_postgres_18` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_18-1.2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/edtf_postgres_18-1.2.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.1 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.1 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.3 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-edtf-postgres` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-18-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-18-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `edtf_postgres_17` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_17-1.2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/edtf_postgres_17-1.2.3-1PGSTY.el8.x86_64.rpm) |
| `edtf_postgres_17` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [edtf_postgres_17-1.2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/edtf_postgres_17-1.2.3-1PGSTY.el8.aarch64.rpm) |
| `edtf_postgres_17` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_17-1.2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/edtf_postgres_17-1.2.3-1PGSTY.el9.x86_64.rpm) |
| `edtf_postgres_17` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_17-1.2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/edtf_postgres_17-1.2.3-1PGSTY.el9.aarch64.rpm) |
| `edtf_postgres_17` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_17-1.2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/edtf_postgres_17-1.2.3-1PGSTY.el10.x86_64.rpm) |
| `edtf_postgres_17` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_17-1.2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/edtf_postgres_17-1.2.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.1 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.1 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.3 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-edtf-postgres` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-17-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-17-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `edtf_postgres_16` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_16-1.2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/edtf_postgres_16-1.2.3-1PGSTY.el8.x86_64.rpm) |
| `edtf_postgres_16` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [edtf_postgres_16-1.2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/edtf_postgres_16-1.2.3-1PGSTY.el8.aarch64.rpm) |
| `edtf_postgres_16` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_16-1.2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/edtf_postgres_16-1.2.3-1PGSTY.el9.x86_64.rpm) |
| `edtf_postgres_16` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_16-1.2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/edtf_postgres_16-1.2.3-1PGSTY.el9.aarch64.rpm) |
| `edtf_postgres_16` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_16-1.2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/edtf_postgres_16-1.2.3-1PGSTY.el10.x86_64.rpm) |
| `edtf_postgres_16` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_16-1.2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/edtf_postgres_16-1.2.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.1 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.1 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.3 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-edtf-postgres` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-16-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-16-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `edtf_postgres_15` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_15-1.2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/edtf_postgres_15-1.2.3-1PGSTY.el8.x86_64.rpm) |
| `edtf_postgres_15` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [edtf_postgres_15-1.2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/edtf_postgres_15-1.2.3-1PGSTY.el8.aarch64.rpm) |
| `edtf_postgres_15` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_15-1.2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/edtf_postgres_15-1.2.3-1PGSTY.el9.x86_64.rpm) |
| `edtf_postgres_15` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_15-1.2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/edtf_postgres_15-1.2.3-1PGSTY.el9.aarch64.rpm) |
| `edtf_postgres_15` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_15-1.2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/edtf_postgres_15-1.2.3-1PGSTY.el10.x86_64.rpm) |
| `edtf_postgres_15` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_15-1.2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/edtf_postgres_15-1.2.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.1 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.1 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.3 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-edtf-postgres` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-15-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-15-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `edtf_postgres_14` | `1.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_14-1.2.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/edtf_postgres_14-1.2.3-1PGSTY.el8.x86_64.rpm) |
| `edtf_postgres_14` | `1.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.5 MiB | [edtf_postgres_14-1.2.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/edtf_postgres_14-1.2.3-1PGSTY.el8.aarch64.rpm) |
| `edtf_postgres_14` | `1.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_14-1.2.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/edtf_postgres_14-1.2.3-1PGSTY.el9.x86_64.rpm) |
| `edtf_postgres_14` | `1.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_14-1.2.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/edtf_postgres_14-1.2.3-1PGSTY.el9.aarch64.rpm) |
| `edtf_postgres_14` | `1.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.5 MiB | [edtf_postgres_14-1.2.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/edtf_postgres_14-1.2.3-1PGSTY.el10.x86_64.rpm) |
| `edtf_postgres_14` | `1.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.6 MiB | [edtf_postgres_14-1.2.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/edtf_postgres_14-1.2.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.2 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.1 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.1 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.3 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.3 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~noble_amd64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~noble_arm64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 1.3 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-edtf-postgres` | `1.2.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.3 MiB | [postgresql-14-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/e/edtf-postgres/postgresql-14-edtf-postgres_1.2.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/monumental-archive/edtf/tree/edtf-postgres-v1.2.3/crates/edtf-postgres" title="Repository" icon="github" subtitle="github.com/monumental-archive/edtf/tree/edtf-postgres-v1.2.3/crates/edtf-postgres" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="edtf_postgres-1.2.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg edtf_postgres;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install edtf_postgres;		# install via package name, for the active PG version

pig install edtf_postgres -v 18;   # install for PG 18
pig install edtf_postgres -v 17;   # install for PG 17
pig install edtf_postgres -v 16;   # install for PG 16
pig install edtf_postgres -v 15;   # install for PG 15
pig install edtf_postgres -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION edtf_postgres;
```

## Usage

Sources:

- [Official extension README](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/README.md)
- [Extension control file](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/edtf_postgres.control)
- [pgrx manifest](https://github.com/monumental-archive/edtf/blob/edtf-postgres-v1.2.3/crates/edtf-postgres/Cargo.toml)

`edtf_postgres` exposes Extended Date/Time Format validation, normalization, bounds, and temporal relations for ISO 8601-2:2019 Annex A strings.

### Enablement

Release 1.2.3 publishes artifacts for PostgreSQL 14–18 on amd64 and arm64. Install the matching artifact, then run:

```sql
CREATE EXTENSION edtf_postgres;
```

The extension is relocatable and trusted, so a role with database `CREATE` privilege can install it. It does not require preloading.

### Validation and Canonical Form

Use `edtf_valid` before accepting external text, `edtf_level` to inspect the supported EDTF level, and `edtf_canonical` to normalize equivalent forms.

```sql
SELECT edtf_valid('1985-04-12');
SELECT edtf_level('1985-04-12/..');
SELECT edtf_canonical('1985-04-12');
```

`edtf_min` and `edtf_max` return index-friendly date bounds. `edtf_relation` compares two expressions and returns a three-valued temporal relation.

```sql
SELECT edtf_min('1985-04'), edtf_max('1985-04');
SELECT edtf_relation('1985', '1986');
```

### Data Modeling Boundary

The extension operates on text and derived dates; it does not replace application validation or introduce a stored EDTF base type. Preserve the original expression when uncertainty and qualifiers are significant, and index derived bounds only for the query semantics your application intends.

