---
title: "pg_ducklake"
linkTitle: "pg_ducklake"
description: "DuckLake lakehouse extension for PostgreSQL, backed by DuckDB and Parquet"
weight: 2490
categories: ["OLAP"]
width: full
---

[**pg_ducklake**](https://github.com/relytcloud/pg_ducklake) : DuckLake lakehouse extension for PostgreSQL, backed by DuckDB and Parquet


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2490** | {{< badge content="pg_ducklake" link="https://github.com/relytcloud/pg_ducklake" >}} | {{< ext "pg_ducklake" >}} | `1.0.0` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `ducklake` |
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_analytics" >}} {{< ext "pg_parquet" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_ducklake` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_ducklake_18" "green" >}} {{< bg "17" "pg_ducklake_17" "green" >}} {{< bg "16" "pg_ducklake_16" "green" >}} {{< bg "15" "pg_ducklake_15" "green" >}} {{< bg "14" "pg_ducklake_14" "green" >}} | `pg_ducklake_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-ducklake" "green" >}} {{< bg "17" "postgresql-17-pg-ducklake" "green" >}} {{< bg "16" "postgresql-16-pg-ducklake" "green" >}} {{< bg "15" "postgresql-15-pg-ducklake" "green" >}} {{< bg "14" "postgresql-14-pg-ducklake" "green" >}} | `postgresql-$v-pg-ducklake` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_ducklake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ducklake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ducklake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ducklake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ducklake_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_ducklake_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ducklake_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ducklake_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ducklake_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ducklake_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_ducklake_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pg-ducklake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pg-ducklake : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ducklake_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 MiB | [pg_ducklake_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ducklake_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ducklake_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.6 MiB | [pg_ducklake_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ducklake_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ducklake_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 MiB | [pg_ducklake_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ducklake_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ducklake_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.9 MiB | [pg_ducklake_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ducklake_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.6 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.5 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.2 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.5 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.5 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-ducklake` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 MiB | [postgresql-18-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-18-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ducklake_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 MiB | [pg_ducklake_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ducklake_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ducklake_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.6 MiB | [pg_ducklake_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ducklake_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ducklake_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 MiB | [pg_ducklake_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ducklake_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ducklake_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.9 MiB | [pg_ducklake_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ducklake_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.6 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.5 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.2 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.5 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.5 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-ducklake` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 MiB | [postgresql-17-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-17-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ducklake_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 MiB | [pg_ducklake_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ducklake_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ducklake_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.6 MiB | [pg_ducklake_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ducklake_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ducklake_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 MiB | [pg_ducklake_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ducklake_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ducklake_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.9 MiB | [pg_ducklake_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ducklake_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.6 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.5 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.2 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.5 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.5 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-ducklake` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 MiB | [postgresql-16-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-16-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ducklake_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.1 MiB | [pg_ducklake_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ducklake_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ducklake_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.6 MiB | [pg_ducklake_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ducklake_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ducklake_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 MiB | [pg_ducklake_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ducklake_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ducklake_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.9 MiB | [pg_ducklake_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ducklake_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.6 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.5 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.3 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.6 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.5 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-ducklake` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 MiB | [postgresql-15-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-15-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ducklake_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.1 MiB | [pg_ducklake_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ducklake_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ducklake_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.6 MiB | [pg_ducklake_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ducklake_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ducklake_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.6 MiB | [pg_ducklake_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ducklake_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ducklake_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.9 MiB | [pg_ducklake_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ducklake_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.5 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.6 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.5 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.4 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.3 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.3 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.6 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.5 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.1 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-ducklake` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 MiB | [postgresql-14-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-ducklake/postgresql-14-pg-ducklake_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/relytcloud/pg_ducklake" title="Repository" icon="github" subtitle="github.com/relytcloud/pg_ducklake" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_ducklake-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_ducklake;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_ducklake;		# install via package name, for the active PG version

pig install pg_ducklake -v 18;   # install for PG 18
pig install pg_ducklake -v 17;   # install for PG 17
pig install pg_ducklake -v 16;   # install for PG 16
pig install pg_ducklake -v 15;   # install for PG 15
pig install pg_ducklake -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_ducklake';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_ducklake;
```

## Usage

Sources: [README](https://github.com/relytcloud/pg_ducklake/blob/v1.0.0/README.md), [v1.0.0 release](https://github.com/relytcloud/pg_ducklake/releases/tag/v1.0.0), [project docs](https://github.com/relytcloud/pg_ducklake/tree/v1.0.0/pg_ducklake/docs)

`pg_ducklake` adds DuckLake tables to PostgreSQL. DuckLake metadata lives in PostgreSQL while table data is stored in Parquet and queried through DuckDB, giving PostgreSQL SQL clients access to lakehouse features such as snapshots, time travel, partitioning, sort keys, and external object storage.

### Create A DuckLake Table

```sql
CREATE EXTENSION pg_ducklake;

CREATE TABLE events (
  id int,
  kind text,
  ts timestamptz
) USING ducklake;

INSERT INTO events VALUES
  (1, 'login', now()),
  (2, 'click', now());

SELECT * FROM events ORDER BY id;
```

Set a table path explicitly when data should live outside the default path:

```sql
CREATE TABLE lake_events (
  id int,
  payload jsonb
) WITH (
  ducklake.table_path = 's3://my-bucket/prefix/'
) USING ducklake;
```

### Time Travel

Each commit creates a snapshot. Capture a snapshot id before a change, then query the older state:

```sql
SELECT max(snapshot_id) AS before_delete
FROM ducklake.ducklake_snapshot \gset

DELETE FROM events WHERE id = 1;

SELECT * FROM ducklake.time_travel('events'::regclass, :before_delete);
```

### Convert And Load Data

Create DuckLake tables from existing PostgreSQL heap tables or external data readers:

```sql
CREATE TABLE row_store AS
SELECT i AS id, 'hello pg_ducklake' AS msg
FROM generate_series(1, 10000) AS i;

CREATE TABLE col_store USING ducklake AS
SELECT * FROM row_store;

CREATE TABLE titanic USING ducklake AS
SELECT * FROM ducklake.read_csv('https://raw.githubusercontent.com/datasciencedojo/datasets/master/titanic.csv');
```

### Inlining, Partitioning, And Maintenance

Small writes are inlined in metadata by default to avoid creating many tiny Parquet files. Tune the row limit or flush explicitly:

```sql
CALL ducklake.set_option('data_inlining_row_limit', 100);
SELECT * FROM ducklake.flush_inlined_data('events'::regclass);
```

Partition and sort tables for pruning and analytics:

```sql
CALL ducklake.set_partition('events'::regclass, 'bucket(4, id)', 'month(ts)');
CREATE INDEX ON events USING ducklake_sorted (id, ts);
```

Run maintenance on demand when automatic background maintenance is not enough:

```sql
SELECT * FROM ducklake.merge_adjacent_files('events'::regclass);
CALL ducklake.set_option('expire_older_than', '7 days');
SELECT * FROM ducklake.expire_snapshots();
SELECT * FROM ducklake.cleanup_old_files();
```

### External DuckDB Access

DuckDB clients can attach the same DuckLake metadata:

```sql
INSTALL ducklake;
LOAD ducklake;
ATTACH 'ducklake:postgres:dbname=postgres host=localhost' AS my_ducklake
  (METADATA_SCHEMA 'ducklake');

SELECT * FROM my_ducklake.public.events;
```

### Caveats

- Version 1.0.0 supports PostgreSQL 14-18.
- The README lists Ubuntu 22.04-24.04 and macOS as source-build targets.
- Cloud credentials are stored through a `ducklake_secret` foreign server and per-user mappings; protect those catalog objects like other database secrets.
- For incremental heap-to-DuckLake conversion, upstream points to the separate `pg_duckpipe` project.
